package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"reflect"
	"sort"
	"strings"
	"syscall"
	"time"
	"watch_for_arb/adjustment/decimal"
	"watch_for_arb/common_flags"
	"watch_for_arb/contracts/aave/aave_price_oracle"
	"watch_for_arb/contracts/aave/lending_pool"
	"watch_for_arb/contracts/aave/lending_pool_addresses_provider"
	"watch_for_arb/contracts/aave/lending_pool_data_provider"
	"watch_for_arb/contracts/compound/cerc20"
	"watch_for_arb/contracts/compound/comptroller_g4"
	"watch_for_arb/contracts/custom_contracts/uniswap_flashloan_liquidate_aave"
	comp_based "watch_for_arb/contracts/custom_contracts/uniswap_flashloan_liquidate_comp_based"
	"watch_for_arb/contracts/erc20"
	"watch_for_arb/contracts/uniswap/pair"
	"watch_for_arb/contracts/uniswap/router"
	"watch_for_arb/mainnet_addrs"

	"github.com/c-bata/go-prompt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"
	"github.com/pkg/profile"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"golang.org/x/sync/errgroup"
)

type amt_as_token struct {
	CTokenAmount  *big.Int
	Exchange_rate *big.Int
	Ctoken_addr   common.Address
}

type reserves_data struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}

type borrow_side_params struct {
	Price                 *big.Int
	UnderlyingAssetAmount *big.Int
	BorrowToken           common.Address
	RepayCToken           common.Address
	BorrowPair            common.Address
	ExchangeRate          *big.Int
}

type sell_side_params struct {
	Price                 *big.Int
	UnderlyingAssetAmount *big.Int
	Amount                *big.Int
	SellToPair            common.Address
	ReceiveBonusCToken    common.Address
	ReceiveBonusToken     common.Address
	ExchangeRate          *big.Int
}

type comp_bot struct {
	event_subs comp_event_subscriptions
	incoming   comp_incoming_chans
}

type aave_bot struct {
	event_subs aave_event_subscriptions
	incoming   aave_incoming_chans
}

type clients struct {
	comp      *ethclient.Client
	aave      *ethclient.Client
	cream     *ethclient.Client
	otherwise *ethclient.Client
}

type report_msg struct {
	Error    error
	Role     string
	When     time.Time
	ArgsUsed interface{}
	Extra    interface{}
}

type report_logs struct {
	estimate_fail_log chan *report_msg
	not_worth_it_log  chan *report_msg
}

type bot struct {
	db                  *leveldb.DB
	clients             clients
	comp                comp_bot
	cream               comp_bot
	aave                aave_bot
	shutdown            chan struct{}
	log_update_incoming report_logs
}

type comp_incoming_chans struct {
	Market_entered          chan *comptroller_g4.ComptrollerG4MarketEntered
	Market_exited           chan *comptroller_g4.ComptrollerG4MarketExited
	CompDistributedBorrower chan *comptroller_g4.ComptrollerG4DistributedBorrowerComp
	CompDistributedSupplier chan *comptroller_g4.ComptrollerG4DistributedSupplierComp
}

type aave_incoming_chans struct {
	Borrow                         chan *lending_pool.LendingPoolBorrow
	Deposit                        chan *lending_pool.LendingPoolDeposit
	Flashloan                      chan *lending_pool.LendingPoolFlashLoan
	Liquidate                      chan *lending_pool.LendingPoolLiquidationCall
	Origination_fee_liquidate      chan *lending_pool.LendingPoolOriginationFeeLiquidated
	Rebalance                      chan *lending_pool.LendingPoolRebalanceStableBorrowRate
	Redeem_underlying              chan *lending_pool.LendingPoolRedeemUnderlying
	Repay                          chan *lending_pool.LendingPoolRepay
	Reserve_used_as_colat_disabled chan *lending_pool.LendingPoolReserveUsedAsCollateralDisabled
	Reserve_used_as_colat_enabled  chan *lending_pool.LendingPoolReserveUsedAsCollateralEnabled
	Swap                           chan *lending_pool.LendingPoolSwap
}

type aave_user_account_data struct {
	TotalLiquidityETH           *big.Int
	TotalCollateralETH          *big.Int
	TotalBorrowsETH             *big.Int
	TotalFeesETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}

type aave_user_account struct {
	UserAddr common.Address
	aave_user_account_data
}

type comp_user_account struct {
	UserAddr  common.Address
	Liquidity *big.Int
	ShortFall *big.Int
}

type aave_event_subscriptions struct {
	Borrow                event.Subscription
	Deposit               event.Subscription
	Flashloan             event.Subscription
	Liquidation           event.Subscription
	Origination           event.Subscription
	Rebalance             event.Subscription
	Repay                 event.Subscription
	Reserve_colat_disable event.Subscription
	Reserve_colat_enable  event.Subscription
	Swap                  event.Subscription
}

type comp_event_subscriptions struct {
	Market_entered          event.Subscription
	Market_exited           event.Subscription
	CompDistributedSupplier event.Subscription
	CompDistributedBorrower event.Subscription
}

type aave_liq_params struct {
	TokenAddr                    common.Address
	Borrowed_token_balance_repay *big.Int
	Collateral_bonus_amt         *big.Int
	Current_a_token_balance      *big.Int
	Usage_as_collateral_enabled  bool
}

type compliq_with_handle struct {
	*mainnet_addrs.CompLiq
	*pair.Pair
	isFirst bool
	cToken  *cerc20.Cerc20
}

type record struct {
	Addr   common.Address
	Symbol string
	Amount string
}

type array_container struct {
	spot  int
	addrs [][]byte
	err   error
}

type simple_iter interface {
	Release()
	Error() error
	Key() []byte
	Next() bool
}

func new_array_container(addrs [][]byte) simple_iter {
	return &array_container{0, addrs, nil} //
}

func (c *array_container) Error() error {
	return nil
}

func (c *array_container) First() bool {
	return c.spot == 0
}

func (c *array_container) Last() bool {
	return c.spot == len(c.addrs)
}

func (c *array_container) Release() {
	// c.spot = 0 //
	// c.addrs = [][]byte{}
}

func (c *array_container) Next() bool {
	return c.spot < len(c.addrs)
}

func (c *array_container) Key() []byte {
	item := c.addrs[c.spot]
	c.spot++
	return item
}

const (
	ESTIMATE_GAS_FAIL    = "estimate_gas_fail.json"
	REASONS_NOT_WORTH_IT = "reason_why_arg_not_worth_it.json"
	comp_seize_too_much  = "execution reverted: LIQUIDATE_SEIZE_TOO_MUCH"
)

var (
	aave_liquidations   = []byte("aave-liquidations-")
	comp_liquidations   = []byte("comp-liquidations-")
	cream_liquidations  = []byte("cream-liquidations-")
	yfi_last_bought     = []byte("yfi-interaction-")
	user_aave_prefix    = []byte("aave-user-")
	user_comp_prefix    = []byte("comp-user-")
	user_cream_prefix   = []byte("cream-user-")
	claimed_liqs_prefix = []byte("already-claimed-liq-")
	db_path             = flag.String(
		"db_path", "money-market-db", "where is the db",
	)
	repl_mode = flag.Bool(
		"repl_mode", false, "repl mode to inspect DB",
	)
	graph_pricing = flag.Bool(
		"graph", false, "graph mode check",
	)
	dump_stats = flag.Bool(
		"dump_stats", false, "report how many aave & comp users known",
	)
	aave_check = flag.Uint(
		"aave_every", 1, "check aave accounts every how many seconds",
	)
	comp_check = flag.Uint(
		"comp_every", 1, "check comp accounts every how many seconds",
	)
	no_comp_loop = flag.Bool(
		"no_comp_loop", false, "whether to disable the comp loop logic",
	)
	health_factor_one = big.NewInt(1e18)
	factor            = decimal.NewDec(1e18)
	factor_28         = decimal.MustNewDecFromStr(
		"10000000000000000000000000000",
	).TruncateInt()
	factor_8                      = decimal.MustNewDecFromStr("100000000").TruncateInt()
	private_key_contract_owner, _ = crypto.HexToECDSA(
		"",
	)
	enable_debug      = flag.Bool("debug", false, "global debug")
	aave_debug        = flag.Bool("aave_liq_debug", false, "aave debug liquidations")
	comp_debug        = flag.Bool("comp_liq_debug", false, "comp debug liquidations")
	cream_debug       = flag.Bool("cream_liq_debug", false, "cream debug liquidation")
	show_cream_tokens = flag.Bool("show_cream_tokens", false, "whether to dump cream tokens")
	enable_arb_pairs  = flag.Bool(
		"arb_pairs", true, "whether to arb sushiswap, uniswap on yfi",
	)
	enable_est_fail_log = flag.Bool(
		"estimate_fail_log", false, "keeping failed estimates in "+ESTIMATE_GAS_FAIL,
	)
	enable_not_worth_it_log = flag.Bool(
		"not_worth_it_log", false, "keep not worth it log "+REASONS_NOT_WORTH_IT,
	)
	ram_liq_user   = flag.String("ram_user", "", "give a pass on an txn liquizied user")
	opts           = bind.NewKeyedTransactor(private_key_contract_owner)
	dry_run        = flag.Bool("dry_run", false, "dry run ")
	uni_liq_abi, _ = abi.JSON(
		strings.NewReader(
			string(uniswap_flashloan_liquidate_aave.UniswapFlashloanLiquidateAaveABI),
		),
	)

	uni_flashloan_liquidate_comp_abi, _ = abi.JSON(
		strings.NewReader(
			string(comp_based.UniswapFlashloanLiquidateCompBasedABI),
		),
	)
	max_gas_uniswap_strategy = uint64(1_850_000)
	one_erc_20               = big.NewInt(1e18)
	half_yfi                 = big.NewInt(5e17)
	min_ctoken_borrowed      = big.NewInt(1)
	ignore_estimate_error    = flag.Bool(
		"ignore_estimate_error", false, "whether to ignore estimate gas errors",
	)
	ignore_estimate_error_aave = flag.Bool(
		"ignore_estimate_error_aave", false, "whether to ignore estimate gas errors in aave",
	)
	ignore_estimate_error_cream = flag.Bool(
		"ignore_estimate_error_cream", false,
		"whether to ignore estimate gas errors in cream liquidation",
	)
	more_gas              = big.NewInt(1e9)
	only_do_arb_check     = flag.Bool("only_arb_check", false, "arb check only - no trades")
	counter_key           = []byte("counter-key")
	enforce_max_borrows   = flag.Bool("enforce_max_borrow", true, "enforce max borrow limits")
	only_show_min_borrows = flag.Bool(
		"only_min_borrows", false, "only show min borrows",
	)
	max_gas_price = big.NewInt(130e9)
	enable_pprof  = flag.Bool("pprof", false, "pprof profiling")
)

func asSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func twenty_percent_more_gas(current_gas *big.Int) {

	// mutates the gas value
}

func (b *bot) read_counter_key() string {
	s, err := b.db.Get(counter_key, nil)
	if err != nil {
		return string("0")
	}
	return string(s)
}

func (b *bot) set_counter_key(counter string) error {
	return b.db.Put(counter_key, []byte(counter), nil)
}

const (
	ETHERSCAN_TXNHASH_LINK = "https://etherscan.io/tx/"
	ETHERSCAN_ADDR_LINK    = "https://etherscan.io/address/"
)

func etherscan_of_hash(txn_hash common.Hash) string {
	return ETHERSCAN_TXNHASH_LINK + txn_hash.Hex()
}

func etherscan_of_addr(addr common.Address) string {
	return ETHERSCAN_ADDR_LINK + addr.Hex()
}

func pretty_token(
	tk common.Address,
	value *big.Int,
	cl *ethclient.Client,
) (string, string) {

	if tk == mainnet_addrs.AAVE_ETH_HANDLE {
		tk = mainnet_addrs.WETH_ADDR
	}

	token, err := erc20.NewErc20(tk, cl)
	if err != nil {
		return "", ""
	}

	dec, err := token.Decimals(nil)
	if err != nil {
		return "", ""
	}

	symbol, err := token.Symbol(nil)
	if err != nil {
		return "", ""
	}

	one_token := decimal.NewDecFromBigInt(new(big.Int).Exp(
		ten,
		big.NewInt(int64(dec)),
		nil,
	))

	return symbol, decimal.NewDecFromBigInt(
		value,
	).Quo(one_token).String()
}

func pretty_aave(
	params []aave_liq_params,
	client *ethclient.Client,
) string {

	type record struct {
		Addr          common.Address
		Symbol        string
		BorrowBalance string
		ATokenBalance string
	}
	var all_borrows []record

	for _, b_ := range params {

		borrow_bal := "0"
		a_token_bal := "0"
		symbol := ""

		if b_.Current_a_token_balance != nil &&
			b_.Current_a_token_balance.Cmp(common.Big0) == 1 {
			symbol, a_token_bal = pretty_token(
				b_.TokenAddr, b_.Current_a_token_balance, client,
			)
		}

		if b_.Borrowed_token_balance_repay != nil &&
			b_.Borrowed_token_balance_repay.Cmp(common.Big0) == 1 {
			symbol, borrow_bal = pretty_token(
				b_.TokenAddr, b_.Borrowed_token_balance_repay, client,
			)
		}

		all_borrows = append(all_borrows, record{
			Addr:          b_.TokenAddr,
			Symbol:        symbol,
			BorrowBalance: borrow_bal,
			ATokenBalance: a_token_bal,
		})
	}

	s1, _ := json.MarshalIndent(all_borrows, "", "\t")
	return string(s1)
}

func pretty_lists(
	cheapest_price []borrow_side_params,
	most_expensive []sell_side_params,
	client *ethclient.Client,
) (string, string) {
	var (
		all_borrows  []record
		all_supplies []record
	)

	for _, b_ := range cheapest_price {
		symbol, amount := pretty_token(
			b_.BorrowToken, b_.UnderlyingAssetAmount, client,
		)
		all_borrows = append(all_borrows, record{
			Addr:   b_.BorrowToken,
			Symbol: symbol,
			Amount: amount,
		})
	}

	for _, b_ := range most_expensive {
		symbol, amount := pretty_token(
			b_.ReceiveBonusToken, b_.UnderlyingAssetAmount, client,
		)
		all_supplies = append(all_supplies, record{
			Addr:   b_.ReceiveBonusToken,
			Symbol: symbol,
			Amount: amount,
		})
	}

	s1, _ := json.MarshalIndent(all_borrows, "", "\t")
	s2, _ := json.MarshalIndent(all_supplies, "", "\t")
	return string(s1), string(s2)
}

func (b *bot) kick_off_fail_log(file_used string, data_from chan *report_msg) {
	// TODO come back to error handling

	f, err := os.OpenFile(file_used, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Println("come back to this")
		return
	}

	for {
		select {
		case <-b.shutdown:
			f.Close()
			return
		case msg := <-data_from:
			s, err := json.MarshalIndent(msg, "", "\t")
			if err != nil {
				log.Println("come back to this")
				return
			}

			if _, err := f.Write(s); err != nil {
				log.Println("some error on writing to "+file_used, err)
				f.Close()
				return
			}

			f.WriteString("\n")
		}
	}

}

func new_bot(db_path, client_path string) (*bot, error) {
	client, err := ethclient.Dial(client_path)

	if err != nil {
		return nil, err
	}

	client_aave, err := ethclient.Dial(client_path)

	if err != nil {
		return nil, err
	}

	client_comp, err := ethclient.Dial(client_path)

	if err != nil {
		return nil, err
	}

	client_cream, err := ethclient.Dial(client_path)

	if err != nil {
		return nil, err
	}

	db, err := leveldb.RecoverFile(db_path, nil)

	if err != nil {
		return nil, err
	}

	return &bot{
		db: db,
		clients: clients{
			comp:      client_comp,
			aave:      client_aave,
			cream:     client_cream,
			otherwise: client,
		},
		comp: comp_bot{
			incoming: comp_incoming_chans{
				make(chan *comptroller_g4.ComptrollerG4MarketEntered),
				make(chan *comptroller_g4.ComptrollerG4MarketExited),
				make(chan *comptroller_g4.ComptrollerG4DistributedBorrowerComp),
				make(chan *comptroller_g4.ComptrollerG4DistributedSupplierComp),
			},
		},
		cream: comp_bot{
			incoming: comp_incoming_chans{
				make(chan *comptroller_g4.ComptrollerG4MarketEntered),
				make(chan *comptroller_g4.ComptrollerG4MarketExited),
				make(chan *comptroller_g4.ComptrollerG4DistributedBorrowerComp),
				make(chan *comptroller_g4.ComptrollerG4DistributedSupplierComp),
			},
		},
		aave: aave_bot{
			incoming: aave_incoming_chans{
				make(chan *lending_pool.LendingPoolBorrow),
				make(chan *lending_pool.LendingPoolDeposit),
				make(chan *lending_pool.LendingPoolFlashLoan),
				make(chan *lending_pool.LendingPoolLiquidationCall),
				make(chan *lending_pool.LendingPoolOriginationFeeLiquidated),
				make(chan *lending_pool.LendingPoolRebalanceStableBorrowRate),
				make(chan *lending_pool.LendingPoolRedeemUnderlying),
				make(chan *lending_pool.LendingPoolRepay),
				make(chan *lending_pool.LendingPoolReserveUsedAsCollateralDisabled),
				make(chan *lending_pool.LendingPoolReserveUsedAsCollateralEnabled),
				make(chan *lending_pool.LendingPoolSwap),
			},
		},
		log_update_incoming: report_logs{
			make(chan *report_msg), make(chan *report_msg),
		},
		shutdown: make(chan struct{}),
	}, nil
}

func (b *bot) enable_comp_subscriptions(
	troll *comptroller_g4.ComptrollerG4,
) error {

	m_entered_sub, err := troll.WatchMarketEntered(
		nil, b.comp.incoming.Market_entered,
	)

	if err != nil {
		return err
	}

	m_exited_sub, err := troll.WatchMarketExited(nil, b.comp.incoming.Market_exited)

	if err != nil {
		return err
	}

	m_comp_borrower, err := troll.WatchDistributedBorrowerComp(
		nil, b.comp.incoming.CompDistributedBorrower, nil, nil,
	)

	if err != nil {
		return err
	}

	m_comp_supplier, err := troll.WatchDistributedSupplierComp(
		nil, b.comp.incoming.CompDistributedSupplier, nil, nil,
	)

	if err != nil {
		return err
	}

	b.comp.event_subs.Market_entered = m_entered_sub
	b.comp.event_subs.Market_exited = m_exited_sub
	b.comp.event_subs.CompDistributedBorrower = m_comp_borrower
	b.comp.event_subs.CompDistributedSupplier = m_comp_supplier

	return nil
}

func (b *bot) enable_cream_subscriptions(
	troll *comptroller_g4.ComptrollerG4,
) error {

	m_entered_sub, err := troll.WatchMarketEntered(
		nil, b.cream.incoming.Market_entered,
	)

	if err != nil {
		return err
	}

	m_exited_sub, err := troll.WatchMarketExited(nil, b.cream.incoming.Market_exited)

	if err != nil {
		return err
	}

	m_comp_borrower, err := troll.WatchDistributedBorrowerComp(
		nil, b.cream.incoming.CompDistributedBorrower, nil, nil,
	)

	if err != nil {
		return err
	}

	m_comp_supplier, err := troll.WatchDistributedSupplierComp(
		nil, b.cream.incoming.CompDistributedSupplier, nil, nil,
	)

	if err != nil {
		return err
	}

	b.cream.event_subs.Market_entered = m_entered_sub
	b.cream.event_subs.Market_exited = m_exited_sub
	b.cream.event_subs.CompDistributedBorrower = m_comp_borrower
	b.cream.event_subs.CompDistributedSupplier = m_comp_supplier

	return nil
}

func (b *bot) enable_aave_subscriptions(
	pool_handle *lending_pool.LendingPool,
) error {

	sub_borrow, err := pool_handle.WatchBorrow(
		nil, b.aave.incoming.Borrow, nil, nil, nil,
	)

	if err != nil {
		return err
	}

	sub_deposit, err := pool_handle.WatchDeposit(
		nil, b.aave.incoming.Deposit, nil, nil, nil,
	)

	if err != nil {
		return err
	}

	sub_flashloan, err := pool_handle.WatchFlashLoan(
		nil, b.aave.incoming.Flashloan, nil, nil,
	)

	if err != nil {
		return err
	}

	sub_liquidation, err := pool_handle.WatchLiquidationCall(
		nil, b.aave.incoming.Liquidate, nil, nil, nil,
	)

	if err != nil {
		return err
	}

	sub_origination, err := pool_handle.WatchOriginationFeeLiquidated(
		nil, b.aave.incoming.Origination_fee_liquidate, nil, nil, nil,
	)

	if err != nil {
		return err
	}

	sub_rebalance, err := pool_handle.WatchRebalanceStableBorrowRate(
		nil, b.aave.incoming.Rebalance, nil, nil,
	)

	if err != nil {
		return err
	}

	sub_repay, err := pool_handle.WatchRepay(
		nil, b.aave.incoming.Repay, nil, nil, nil,
	)

	if err != nil {
		return err
	}

	sub_reserve_colat_disable, err := pool_handle.WatchReserveUsedAsCollateralDisabled(
		nil, b.aave.incoming.Reserve_used_as_colat_disabled, nil, nil,
	)

	if err != nil {
		return err
	}

	sub_reserve_colat_enable, err := pool_handle.WatchReserveUsedAsCollateralEnabled(
		nil, b.aave.incoming.Reserve_used_as_colat_enabled, nil, nil,
	)

	if err != nil {
		return err
	}

	sub_swap, err := pool_handle.WatchSwap(
		nil, b.aave.incoming.Swap, nil, nil,
	)

	if err != nil {
		return err
	}

	b.aave.event_subs.Borrow = sub_borrow
	b.aave.event_subs.Deposit = sub_deposit
	b.aave.event_subs.Flashloan = sub_flashloan
	b.aave.event_subs.Liquidation = sub_liquidation
	b.aave.event_subs.Origination = sub_origination
	b.aave.event_subs.Rebalance = sub_rebalance
	b.aave.event_subs.Repay = sub_repay
	b.aave.event_subs.Reserve_colat_disable = sub_reserve_colat_disable
	b.aave.event_subs.Reserve_colat_enable = sub_reserve_colat_enable
	b.aave.event_subs.Swap = sub_swap

	return nil
}

var (
	aave_liquidation_bonus = map[common.Address]decimal.Dec{
		mainnet_addrs.TOKEN_ADDRS["busd"]: decimal.MustNewDecFromStr("0.00"),
		mainnet_addrs.TOKEN_ADDRS["susd"]: decimal.MustNewDecFromStr("0.00"),
		mainnet_addrs.TOKEN_ADDRS["usdt"]: decimal.MustNewDecFromStr("0.00"),
		mainnet_addrs.TOKEN_ADDRS["aave"]: decimal.MustNewDecFromStr("0.10"),
		mainnet_addrs.TOKEN_ADDRS["dai"]:  decimal.MustNewDecFromStr("0.05"),
		mainnet_addrs.TOKEN_ADDRS["tusd"]: decimal.MustNewDecFromStr("0.05"),
		mainnet_addrs.TOKEN_ADDRS["usdc"]: decimal.MustNewDecFromStr("0.05"),
		mainnet_addrs.TOKEN_ADDRS["bat"]:  decimal.MustNewDecFromStr("0.10"),
		mainnet_addrs.TOKEN_ADDRS["enj"]:  decimal.MustNewDecFromStr("0.10"),
		mainnet_addrs.TOKEN_ADDRS["knc"]:  decimal.MustNewDecFromStr("0.10"),
		mainnet_addrs.TOKEN_ADDRS["lend"]: decimal.MustNewDecFromStr("0.10"),
		mainnet_addrs.TOKEN_ADDRS["link"]: decimal.MustNewDecFromStr("0.10"),
		mainnet_addrs.TOKEN_ADDRS["ren"]:  decimal.MustNewDecFromStr("0.10"),
		mainnet_addrs.TOKEN_ADDRS["mkr"]:  decimal.MustNewDecFromStr("0.10"),
		mainnet_addrs.TOKEN_ADDRS["snx"]:  decimal.MustNewDecFromStr("0.10"),
		mainnet_addrs.TOKEN_ADDRS["zrx"]:  decimal.MustNewDecFromStr("0.10"),
		mainnet_addrs.TOKEN_ADDRS["wbtc"]: decimal.MustNewDecFromStr("0.15"),
		mainnet_addrs.TOKEN_ADDRS["yfi"]:  decimal.MustNewDecFromStr("0.15"),
		mainnet_addrs.TOKEN_ADDRS["mana"]: decimal.MustNewDecFromStr("0.10"),
		mainnet_addrs.TOKEN_ADDRS["rep"]:  decimal.MustNewDecFromStr("0.10"),
		mainnet_addrs.AAVE_ETH_HANDLE:     decimal.MustNewDecFromStr("0.05"),
	}
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	for _, key := range mainnet_addrs.AAVE_RESERVE_TOKENS {
		if _, exists := aave_liquidation_bonus[key]; !exists {
			log.Println("AAVE-BAD!", key.Hex(), "not in liquidation bonus map")
		}
		if _, exists := min_borrow_amounts_aave[key]; !exists {
			log.Println("AAVE-BAD!", key.Hex(), "not in min borrow map aave")
		}
		if _, exists := mainnet_addrs.UNISWAP_USABLE_ADDRS[key]; !exists {
			log.Println("AAVE-BAD!", key.Hex(), "not in uniswap pair map, from where to borrow!?")
		}
	}

	for _, value := range mainnet_addrs.COMP_LIQUIDATIONS_PARAMS {
		if _, exists := min_borrow_amounts_comp[value.Token]; !exists {
			log.Println("COMP-BAD!", value.Token.Hex(), "not in min borrow map comp")
		}
		if _, exists := mainnet_addrs.UNISWAP_USABLE_ADDRS[value.Token]; !exists {
			log.Println(
				"COMP-BAD!", value.Token.Hex(), "not in uniswap pair map, from where to borrow!?",
			)
		}
	}

	for _, value := range mainnet_addrs.CREAM_LIQUIDATION_PARAMS {
		if _, exists := min_borrow_amounts_cream[value.Token]; !exists {
			log.Println("CREAM-BAD!", value.Token.Hex(), "not in min borrow map cream ")
		}
		if _, exists := mainnet_addrs.UNISWAP_USABLE_ADDRS[value.Token]; !exists {
			log.Println(
				"CREAM-BAD!", value.Token.Hex(), "not in uniswap pair map, from where to borrow!?",
			)
		}
	}

}

func comp_user_event(
	d *leveldb.DB, t *comptroller_g4.ComptrollerG4,
) func(common.Address) error {
	db := d

	return func(user common.Address) error {
		wo := &opt.WriteOptions{
			Sync: false,
		}

		var key []byte
		key = append(key, user_comp_prefix...)
		key = append(key, user[:]...)

		return db.Put(key, user[:], wo)
	}
}

func cream_user_event(
	d *leveldb.DB, t *comptroller_g4.ComptrollerG4,
) func(common.Address) error {

	db := d

	return func(user common.Address) error {
		wo := &opt.WriteOptions{
			Sync: false,
		}

		var key []byte
		key = append(key, user_cream_prefix...)
		key = append(key, user[:]...)

		return db.Put(key, user[:], wo)
	}
}

func aave_user_event(
	d *leveldb.DB, p *lending_pool.LendingPool,
) func(common.Address) error {
	db := d

	return func(user common.Address) error {

		wo := &opt.WriteOptions{
			Sync: false,
		}

		var key []byte
		key = append(key, user_aave_prefix...)
		key = append(key, user[:]...)

		return db.Put(key, user[:], wo)
	}
}

func (b *bot) close_resources() error {
	b.clients.aave.Close()
	b.clients.comp.Close()
	b.clients.cream.Close()
	b.clients.otherwise.Close()

	return b.db.Close()
}

func run_liquidation_bot(
	b *bot,
) (e error) {

	addr_provider, err := lending_pool_addresses_provider.NewLendingPoolAddressesProvider(
		mainnet_addrs.AAVE_ADDR_PROVIDER, b.clients.aave,
	)

	if err != nil {
		return err
	}

	aave_price_oracle_addr, err := addr_provider.GetPriceOracle(nil)

	if err != nil {
		return err
	}

	pool_addr, err := addr_provider.GetLendingPool(nil)

	if err != nil {
		return err
	}

	pool_handle, err := lending_pool.NewLendingPool(
		pool_addr, b.clients.otherwise,
	)

	if err != nil {
		return err
	}

	if err := b.enable_aave_subscriptions(pool_handle); err != nil {
		return err
	}

	troll, err := comptroller_g4.NewComptrollerG4(
		mainnet_addrs.COMP_TROLLER_ADDR, b.clients.comp,
	)

	if err != nil {
		return err
	}

	cream_troll, err := comptroller_g4.NewComptrollerG4(
		mainnet_addrs.CREAM_TROLLER_ADDR, b.clients.cream,
	)

	if err != nil {
		return err
	}

	if err := b.enable_comp_subscriptions(troll); err != nil {
		return err
	}

	if err := b.enable_cream_subscriptions(cream_troll); err != nil {
		return err
	}

	var g errgroup.Group

	if *enable_est_fail_log {
		go b.kick_off_fail_log(ESTIMATE_GAS_FAIL, b.log_update_incoming.estimate_fail_log)
	}

	if *enable_not_worth_it_log {
		go b.kick_off_fail_log(REASONS_NOT_WORTH_IT, b.log_update_incoming.not_worth_it_log)
	}

	g.Go(func() error {
		interrupt := make(chan os.Signal)
		defer signal.Stop(interrupt)
		defer close(interrupt)
		signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-interrupt
		close(b.shutdown)
		time.Sleep(time.Second * 2)
		return nil
	})

	if *only_do_arb_check {
		router, err := router.NewRouter(
			mainnet_addrs.UNISWAP_ROUTER, b.clients.otherwise,
		)

		if err != nil {
			return err
		}

		t := time.NewTicker(time.Millisecond * 200)
		defer t.Stop()

		for {
			select {
			case <-b.shutdown:
				return nil
			case <-t.C:
				if err := all_paths(router); err != nil {
					return err
				}
			}
		}

	}

	g.Go(func() error {
		cream_handler := cream_user_event(b.db, cream_troll)
		for {
			select {
			case <-b.shutdown:
				return nil
			case oops := <-b.cream.event_subs.Market_entered.Err():
				return oops
			case oops := <-b.cream.event_subs.Market_exited.Err():
				return oops
			case oops := <-b.cream.event_subs.CompDistributedBorrower.Err():
				return oops
			case oops := <-b.cream.event_subs.CompDistributedSupplier.Err():
				return oops

			case payload := <-b.cream.incoming.Market_entered:
				if err := cream_handler(payload.Account); err != nil {
					return err
				}
			case payload := <-b.cream.incoming.Market_exited:
				if err := cream_handler(payload.Account); err != nil {
					return err
				}
			case payload := <-b.cream.incoming.CompDistributedBorrower:
				if err := cream_handler(payload.Borrower); err != nil {
					return err
				}
			case payload := <-b.cream.incoming.CompDistributedSupplier:
				if err := cream_handler(payload.Supplier); err != nil {
					return err
				}

			}
		}
		return nil
	})

	g.Go(func() error {
		debug := *cream_debug
		count := 0
		disabled := *no_comp_loop
		var start time.Time
		prefix_length := len(user_cream_prefix)
		cream_liq, err := comp_based.NewUniswapFlashloanLiquidateCompBased(
			mainnet_addrs.FLASHLOAN_LIQ_COMP_BASED, b.clients.cream,
		)

		if err != nil {
			return err
		}

		router, err := router.NewRouter(
			mainnet_addrs.UNISWAP_ROUTER, b.clients.cream,
		)

		if err != nil {
			return err
		}

		ctoken_swapping_handles := map[common.Address]*compliq_with_handle{}

		for key, value := range mainnet_addrs.CREAM_LIQUIDATION_PARAMS {

			if value != nil {
				pair, err := pair.NewPair(
					value.Pair, b.clients.cream,
				)

				if err != nil {
					return err
				}

				token1, err := pair.Token1(nil)

				if err != nil {
					return err
				}

				cToken, err := cerc20.NewCerc20(key, b.clients.cream)

				if err != nil {
					return err
				}

				ctoken_swapping_handles[key] = &compliq_with_handle{
					CompLiq: value,
					Pair:    pair,
					isFirst: bytes.Compare(key.Bytes(), token1.Bytes()) == -1,
					cToken:  cToken,
				}
			}
		}

		write_to_estimate_gas_fail := *enable_est_fail_log
		ignore_estimate_err := *ignore_estimate_error_cream
		write_to_not_worth_it := *enable_not_worth_it_log
		scale_back := big.NewInt(1000000000000000000)
		dry_run := *dry_run
		close_factor, err := cream_troll.CloseFactorMantissa(nil)
		max_borrow := *enforce_max_borrows

		if err != nil {
			return err
		}

		last_time_done := time.Now()
		is_first_time := true
		using_db := true

		var use_these_cached_addrs [][]byte

		for {
			count = 0
			start = time.Now()

			select {
			case <-b.shutdown:

				return nil
			default:
				if disabled {
					continue
				}
			}

			var iter simple_iter
			if is_first_time {
				iter = b.db.NewIterator(util.BytesPrefix(user_cream_prefix), nil)
			} else if is_first_time == false && time.Since(last_time_done) < 5*time.Minute {
				using_db = false
				iter = new_array_container(use_these_cached_addrs) //
			} else {
				use_these_cached_addrs = [][]byte{}
				using_db = true
				iter = b.db.NewIterator(util.BytesPrefix(user_cream_prefix), nil)
			}

		iterating_loop:

			for {

				select {
				case <-b.shutdown:
					iter.Release()
					return nil
				default:
					if iter.Next() {
						count++
						var user_addr common.Address

						just_key := iter.Key()

						if using_db {
							just_key = just_key[prefix_length:]
							k := make([]byte, len(just_key))
							copy(k[:], just_key[:])
							use_these_cached_addrs = append(use_these_cached_addrs, k)
							user_addr = common.BytesToAddress(just_key)
						} else {
							user_addr = common.BytesToAddress(just_key)
						}

						oops, _, shortfall, err := cream_troll.GetAccountLiquidity(nil, user_addr)

						if err != nil {
							log.Println("some problem on getting cream troller call of account liquidty")
							break iterating_loop
						}

						if oops.Cmp(common.Big0) != 0 {
							log.Printf("why some issue %v\n", oops)
							continue
						}

						if *show_cream_tokens {
							var new_ones []common.Address

							ctokens, _ := cream_troll.GetAssetsIn(nil, user_addr)

							for _, addr := range ctokens {
								if _, exists := mainnet_addrs.CREAM_LIQUIDATION_PARAMS[addr]; !exists {
									new_ones = append(new_ones, addr)
								}
							}
							if len(new_ones) != 0 {
								s, _ := json.MarshalIndent(new_ones, "", "\t")
								fmt.Println("cream tokens for ", user_addr.Hex(), string(s))
							}
						}

						if shortfall.Cmp(common.Big0) != 0 {

							ctokens, err := cream_troll.GetAssetsIn(nil, user_addr)

							if err != nil {
								log.Printf("problem getting cream assets in %v\n", err)
								break iterating_loop
							}

							var (
								borrows []amt_as_token
								collats []amt_as_token
							)

							for i := range ctokens {

								cToken, exists := ctoken_swapping_handles[ctokens[i]]

								if !exists {
									log.Println(
										"why this cream collaterized token doesn't exist - maybe new?",
										ctokens[i].Hex(),
									)
									continue
								}

								if cToken.cToken == nil {
									log.Println("okay why this one nil??", ctokens[i])
									continue
								}

								borrowed_amount, err := cToken.cToken.BorrowBalanceStored(nil, user_addr)

								if err != nil {
									log.Println("is there an error", err)
									return err
								}

								exchange_rate, err := cToken.cToken.ExchangeRateStored(nil)
								if err != nil {
									log.Println("some problem on exchange rate stored", err)
									break iterating_loop
								}

								if borrowed_amount.Cmp(min_ctoken_borrowed) == 1 {
									borrows = append(borrows, amt_as_token{
										// borrowed is in the underlying unit's
										CTokenAmount:  borrowed_amount,
										Exchange_rate: exchange_rate, //
										Ctoken_addr:   ctokens[i],
									})
								}

								ctok_bal, err := cToken.cToken.BalanceOf(nil, user_addr)

								if err != nil {
									return err
								}

								if ctok_bal.Cmp(common.Big0) == 1 {

									collats = append(collats, amt_as_token{
										CTokenAmount:  ctok_bal,
										Exchange_rate: exchange_rate,
										Ctoken_addr:   ctokens[i],
									})
								}

							}

							if len(borrows) == 0 {
								continue
							}

							var (
								cheapest_price []borrow_side_params
								most_expensive []sell_side_params
							)

							// find the cheapest to borrow
							for i := range borrows {
								liq_info := ctoken_swapping_handles[borrows[i].Ctoken_addr]
								reserves, err := liq_info.Pair.GetReserves(nil)

								if err != nil {
									log.Printf("some problem on getting reserve %v\n", err)
									break iterating_loop
								}

								var (
									price *big.Int
									err_  error
								)

								// inverted logic - kinda
								if liq_info.isFirst {
									price, err_ = router.GetAmountOut(
										nil, borrows[i].CTokenAmount, reserves.Reserve1, reserves.Reserve0,
									)
								} else {
									price, err_ = router.GetAmountOut(
										nil, borrows[i].CTokenAmount, reserves.Reserve0, reserves.Reserve1,
									)
								}

								if err_ != nil {
									s, _ := json.MarshalIndent(liq_info, "", "\t")
									log.Println(
										"error on borrows loop", err_,
										borrows[i].Ctoken_addr.Hex(), string(s),
									)
									break iterating_loop
								}

								cheapest_price = append(cheapest_price, borrow_side_params{
									Price:                 price,
									UnderlyingAssetAmount: borrows[i].CTokenAmount,
									BorrowToken:           liq_info.CompLiq.Token,
									RepayCToken:           borrows[i].Ctoken_addr,
									BorrowPair:            liq_info.CompLiq.Pair,
									ExchangeRate:          borrows[i].Exchange_rate,
								})
							}

							for i := range collats {
								liq_info := ctoken_swapping_handles[collats[i].Ctoken_addr]

								reserves, err := liq_info.Pair.GetReserves(nil)

								if err != nil {
									log.Printf("some problem on getting reserve %v\n", err)
									break iterating_loop
								}
								var (
									price *big.Int
									err_  error
								)

								amt := new(big.Int).Mul(
									collats[i].CTokenAmount, collats[i].Exchange_rate,
								)

								amt.Div(amt, one_erc_20)

								if liq_info.isFirst {
									price, err_ = router.GetAmountOut(
										nil,
										amt,
										reserves.Reserve0, reserves.Reserve1,
									)

								} else {
									price, err_ = router.GetAmountOut(
										nil,
										amt,
										reserves.Reserve1, reserves.Reserve0,
									)
								}

								if err_ != nil {
									s1, _ := json.MarshalIndent(liq_info, "", "\t")
									s2, _ := json.MarshalIndent(collats[i], "", "\t")
									log.Println(
										"error on comp collats loop", err_,
										user_addr.Hex(),
										string(s1),
										string(s2),
									)
									break iterating_loop
								}

								most_expensive = append(most_expensive, sell_side_params{
									Price:                 price,
									Amount:                collats[i].CTokenAmount,
									UnderlyingAssetAmount: amt,
									SellToPair:            liq_info.CompLiq.Pair,
									ReceiveBonusCToken:    collats[i].Ctoken_addr,
									ReceiveBonusToken:     liq_info.CompLiq.Token,
									ExchangeRate:          collats[i].Exchange_rate,
								})
							}

							if len(borrows) == 0 {
								continue
							}

							sort.SliceStable(cheapest_price, func(i, j int) bool {
								return cheapest_price[i].UnderlyingAssetAmount.Cmp(
									cheapest_price[j].UnderlyingAssetAmount,
								) == 1
							})

							sort.SliceStable(most_expensive, func(i, j int) bool {
								return most_expensive[i].Price.Cmp(most_expensive[j].Price) == 1
							})

							if debug {
								s1, s2 := pretty_lists(cheapest_price, most_expensive, b.clients.cream)
								log.Println(
									"here is cream list, first is borrows, second is most expensive",
									"user is",
									user_addr.Hex(),
									s1, s2,
								)
							}

							var (
								br *borrow_side_params
								s  sell_side_params
							)

							for _, br_ := range cheapest_price {
								if amt, exists := min_borrow_amounts_cream[br_.BorrowToken]; exists {
									if br_.UnderlyingAssetAmount.Cmp(
										amt,
									) == 1 {
										br = &br_
										break
									} else {
										if write_to_not_worth_it {
											go func() {
												b.log_update_incoming.not_worth_it_log <- &report_msg{
													Error: nil,
													Role:  "cream not worth it ",
													When:  time.Now().UTC(),
													ArgsUsed: struct {
														Token  common.Address
														Amount *big.Int
														Symbol string
													}{
														br_.BorrowToken,
														br_.UnderlyingAssetAmount,
														mainnet_addrs.REVERSE_NAMING[br_.BorrowToken],
													},
												}
											}()
										}
									}

								} else {
									log.Println("TODO NEED TO ADD THIS bORROW TOKEN", br_.BorrowToken.Hex())
								}
							}

							if br == nil {
								continue
							}

							if len(collats) > 0 {
								s = most_expensive[0]
							} else {
								s = sell_side_params{
									ReceiveBonusCToken: br.RepayCToken,
									ReceiveBonusToken:  br.BorrowToken,
									SellToPair:         br.BorrowPair,
								}
							}

							if br.BorrowToken == s.ReceiveBonusToken {
								if _, exists := mainnet_addrs.CANT_BE_REPAY_SAME[br.BorrowToken]; exists {
									if debug {
										did_have, _ := json.MarshalIndent(ctokens, "", "\t")
										log.Printf(
											"special case -> can't liquidate, seize same token %s for user %s %s\n",
											br.BorrowToken.Hex(),
											user_addr.Hex(),
											string(did_have),
										)
									}
									continue
								}
							}

							br.UnderlyingAssetAmount.Mul(br.UnderlyingAssetAmount, close_factor)
							br.UnderlyingAssetAmount.Div(br.UnderlyingAssetAmount, scale_back)

							if max_borrow {
								if max_borrow, exists := max_borrows[br.BorrowToken]; exists {
									if br.UnderlyingAssetAmount.Cmp(max_borrow) >= 0 {
										br.UnderlyingAssetAmount = max_borrow
									}
								} else {
									log.Println("fogot this token to add to max borrow", br.BorrowToken.Hex())
								}
							}

							args := comp_based.FlashloanLiquidateCompBasedFlashLiquidateParam{
								BorrowPair:                br.BorrowPair,
								BorrowToken:               br.BorrowToken,
								RepayCToken:               br.RepayCToken,
								ReceiveBonusCToken:        s.ReceiveBonusCToken,
								ReceiveBonusToken:         s.ReceiveBonusToken,
								SellToPair:                s.SellToPair,
								UnderCollaterizedCompUser: user_addr,
								HowMuchBorrow:             br.UnderlyingAssetAmount,
							}

							markets_args := comp_based.FlashloanLiquidateCompBasedSetupMarkets{
								Troller: mainnet_addrs.CREAM_TROLLER_ADDR,
								Markets: []common.Address{
									args.RepayCToken, args.ReceiveBonusCToken,
								},
							}

							if args.ReceiveBonusCToken == mainnet_addrs.TOKEN_ADDRS["ftx"] {
								// lua
								args.OneMoreSellPairTokenSell = common.HexToAddress("0xb1f66997a5760428d3a87d68b90bfe0ae64121cc")
								// lua-eth
								args.OneMoreSwapAfterSell = common.HexToAddress("0xc5d3c66133a6264b0f2e712b8e10ef96efb93eb2")
							}

							new_counter := decimal.MustNewDecFromStr(
								b.read_counter_key(),
							).Add(decimal.OneDec())

							if w_bytes := mainnet_addrs.WETH_ADDR.Bytes(); bytes.Compare(
								br.BorrowToken.Bytes(), w_bytes,
							) == 0 {
								handle := ctoken_swapping_handles[args.RepayCToken]
								token0, _ := handle.Pair.Token0(nil)
								args.AssumeBorrowingWethThenIsItFirst = bytes.Compare(
									token0.Bytes(), w_bytes,
								) == 0
							}

							claimed := asSha256(args)

							var key []byte
							key = append(key, claimed_liqs_prefix...)
							key = append(key, []byte(claimed)...)

							if _, err := b.db.Get(key, nil); err == nil {
								if debug {
									log.Println("already claimed this cream liq opprurunity", string(key))
								}
								continue
							}

							suggested_gas, err := b.clients.cream.SuggestGasPrice(
								context.Background(),
							)

							if err != nil {
								log.Printf("issue on suggesting gas price %v\n", err)
								continue
							}

							encoded_args, err := uni_flashloan_liquidate_comp_abi.Pack(
								"flashloan_uniswap_then_liquidate", args, markets_args,
							)

							if err != nil {
								log.Println(
									"not clear why cream liquidation encoding arg caused error ", err,
								)
								continue
							}

							gas_used, err := b.clients.cream.EstimateGas(
								context.Background(), ethereum.CallMsg{
									From:     opts.From,
									To:       &mainnet_addrs.FLASHLOAN_LIQ_COMP_BASED,
									Gas:      max_gas_uniswap_strategy,
									GasPrice: suggested_gas,
									Value:    big.NewInt(0),
									Data:     encoded_args,
								})

							if err != nil && err.Error() == comp_seize_too_much {
								//								fmt.Println("lame way but it works")
								continue
							}

							if err != nil {
								if debug {
									s, _ := json.MarshalIndent(args, "", "\t")
									log.Println(
										"estimating gas errored for cream liquidation ",
										err, string(s),
									)
								}

								if write_to_estimate_gas_fail {
									go func() {
										b.log_update_incoming.estimate_fail_log <- &report_msg{
											Error:    err,
											Role:     "running cream liquidation",
											When:     time.Now().UTC(),
											ArgsUsed: args,
										}
									}()
								}

								if ignore_estimate_err == false {
									continue
								}

							}

							total_cost := new(big.Int).Mul(
								new(big.Int).SetUint64(gas_used), suggested_gas,
							)

							suggested_gas.Add(suggested_gas, more_gas)

							opts.GasPrice = suggested_gas
							opts.GasLimit = max_gas_uniswap_strategy

							if dry_run {
								continue
							}

							if suggested_gas.Cmp(max_gas_price) >= 0 {
								s2, _ := json.MarshalIndent(args, "", "\t")
								log.Println("max gas hit", string(s2))
								continue
							}

							txn2, err2 := cream_liq.FlashloanUniswapThenLiquidate(
								opts, args, markets_args,
							)

							s2, _ := json.MarshalIndent(args, "", "\t")
							fmt.Println(string(s2))

							if err2 != nil {
								fmt.Println("some error on kicking off txn!?", err2)
								continue
							}

							hash := txn2.Hash()

							if err := b.db.Put(key, hash[:], nil); err != nil {
								break iterating_loop
							}

							log.Println(
								"SUCCESS! ->total cream liqiudation cost estimated at total cost",
								total_cost, string(s2),
								etherscan_of_hash(hash),
								err2,
							)

							bind.WaitMined(context.Background(), b.clients.cream, txn2)
							if err := b.set_counter_key(new_counter.String()); err != nil {
								log.Println("wy counter key writing broke?", err)
							}
						}
					} else {
						break iterating_loop
					}
				}
			}

			iter.Release()

			if debug {
				end := time.Now()
				log.Println(
					"finished a keys loop over ",
					count, "for cream records and took",
					end.Sub(start).String(),
				)
			}

			is_first_time = false

			if using_db {
				last_time_done = time.Now()
			}

			if err := iter.Error(); err != nil {
				return err
			}

		}
		return nil
	})

	g.Go(func() error {
		comp_handler := comp_user_event(b.db, troll)
		for {
			select {
			case <-b.shutdown:
				return nil
			case oops := <-b.comp.event_subs.Market_entered.Err():
				return oops
			case oops := <-b.comp.event_subs.Market_exited.Err():
				return oops
			case oops := <-b.comp.event_subs.CompDistributedBorrower.Err():
				return oops
			case oops := <-b.comp.event_subs.CompDistributedSupplier.Err():
				return oops

			case payload := <-b.comp.incoming.Market_entered:
				if err := comp_handler(payload.Account); err != nil {
					return err
				}
			case payload := <-b.comp.incoming.Market_exited:
				if err := comp_handler(payload.Account); err != nil {
					return err
				}

			case payload := <-b.comp.incoming.CompDistributedBorrower:
				if err := comp_handler(payload.Borrower); err != nil {
					return err
				}
			case payload := <-b.comp.incoming.CompDistributedSupplier:
				if err := comp_handler(payload.Supplier); err != nil {
					return err
				}

			}
		}
		return nil
	})

	g.Go(func() error {
		v := reflect.ValueOf(b.aave.event_subs)

		count := v.NumField()
		cases := make([]reflect.SelectCase, count+1)
		for i := 0; i < count; i++ {
			cases[i] = reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(v.Field(i).Interface().(event.Subscription).Err()),
			}
		}

		cases[count] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(b.shutdown),
		}

		for {
			idx, value, recved_ok := reflect.Select(cases)
			if idx == count {
				return nil
			}

			if recved_ok {
				err, ok := value.Interface().(error)
				if !ok {
					continue
				}
				return err
			}
		}

		return nil
	})

	g.Go(func() error {
		user_handler := aave_user_event(b.db, pool_handle)

		for {
			select {
			case <-b.shutdown:
				return nil
			case payload := <-b.aave.incoming.Borrow:
				if err := user_handler(payload.User); err != nil {
					return err
				}

			case payload := <-b.aave.incoming.Deposit:
				if err := user_handler(payload.User); err != nil {
					return err
				}

			case payload := <-b.aave.incoming.Flashloan:
				log.Printf(
					"aave flashloan %s\n",
					etherscan_of_hash(payload.Raw.TxHash),
				)

			case payload := <-b.aave.incoming.Liquidate:
				if err := user_handler(payload.User); err != nil {
					return err
				}

				var key []byte
				key = append(key, aave_liquidations...)
				key = append(key, payload.Raw.TxHash[:]...)

				if err := b.db.Put(key, payload.Raw.TxHash[:], nil); err != nil {
					log.Printf("problem on writing the txn hash for aave liquidation %v\n", err)
					return err
				}

				key = []byte{}

				key = append(key, user_aave_prefix...)
				key = append(key, payload.User[:]...)

				if _, err := b.db.Get(key, nil); err == nil {
					if payload.Liquidator != mainnet_addrs.MY_MAINNET_ADDR ||
						payload.Liquidator != mainnet_addrs.UNISWAP_FLASHLOAN_AAVE_LIQ_ADDR {
						log.Println(
							"had this aave user - why didn't we liquidation it first",
							"\n user:\t", etherscan_of_addr(payload.User),
							"\n tx-hash:\t", etherscan_of_hash(payload.Raw.TxHash),
						)
					}
				}

			case payload := <-b.aave.incoming.Origination_fee_liquidate:
				if err := user_handler(payload.User); err != nil {
					return err
				}

			case payload := <-b.aave.incoming.Rebalance:
				if err := user_handler(payload.User); err != nil {
					return err
				}

			case payload := <-b.aave.incoming.Repay:
				if err := user_handler(payload.User); err != nil {
					return err
				}

			case payload := <-b.aave.incoming.Reserve_used_as_colat_disabled:
				if err := user_handler(payload.User); err != nil {
					return err
				}

			case payload := <-b.aave.incoming.Reserve_used_as_colat_enabled:
				if err := user_handler(payload.User); err != nil {
					return err
				}

			case payload := <-b.aave.incoming.Swap:
				if err := user_handler(payload.User); err != nil {
					return err
				}
			}
		}
		return nil
	})

	g.Go(func() error {
		debug := *enable_debug
		l := len(mainnet_addrs.CREAM_LIQUIDATION_PARAMS) + 1
		cases := make([]reflect.SelectCase, l)
		i := 0

		for key := range mainnet_addrs.CREAM_LIQUIDATION_PARAMS {
			token, err := cerc20.NewCerc20(key, b.clients.cream)
			if err != nil {
				return err
			}
			c := make(chan *cerc20.Cerc20LiquidateBorrow)
			_, err = token.WatchLiquidateBorrow(nil, c)
			if err != nil {
				return err
			}
			cases[i] = reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			}
			i++
		}

		cream_handler := cream_user_event(b.db, cream_troll)
		cases[l-1] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(b.shutdown),
		}

		for {
			idx, value, recved_ok := reflect.Select(cases)

			if idx == (l - 1) {
				return nil
			}

			if recved_ok {
				payload, ok := value.Interface().(*cerc20.Cerc20LiquidateBorrow)

				if !ok {
					continue
				}

				if err := cream_handler(payload.Borrower); err != nil {
					return err
				}

				var key []byte
				key = append(key, cream_liquidations...)
				key = append(key, payload.Raw.TxHash[:]...)

				if err := b.db.Put(key, payload.Raw.TxHash[:], nil); err != nil {
					if debug {
						log.Println("some issue on writing to db the ctoken liquidation")
					}
				}

				key = []byte{}

				key = append(key, user_cream_prefix...)
				key = append(key, payload.Borrower[:]...)
				if _, err := b.db.Get(key, nil); err == nil {
					if payload.Liquidator != mainnet_addrs.FLASHLOAN_LIQ_COMP_BASED {
						log.Println(
							"had this cream user - why didn't we liquidation it first",
							"\n user:\t", etherscan_of_addr(payload.Borrower),
							"\n tx-hash:\t", etherscan_of_hash(payload.Raw.TxHash),
						)
					}
				}
			}
		}

		return nil
	})

	g.Go(func() error {
		debug := *enable_debug
		l := len(mainnet_addrs.COMP_cTOKENS) + 1
		cases := make([]reflect.SelectCase, l)

		for i, tok := range mainnet_addrs.COMP_cTOKENS {
			token, err := cerc20.NewCerc20(tok, b.clients.comp)
			if err != nil {
				return err
			}
			c := make(chan *cerc20.Cerc20LiquidateBorrow)
			_, err = token.WatchLiquidateBorrow(nil, c)
			if err != nil {
				return err
			}
			cases[i] = reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			}
		}

		comp_handler := comp_user_event(b.db, troll)
		cases[l-1] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(b.shutdown),
		}

		for {
			idx, value, recved_ok := reflect.Select(cases)

			if idx == (l - 1) {
				return nil
			}

			if recved_ok {
				payload, ok := value.Interface().(*cerc20.Cerc20LiquidateBorrow)

				if !ok {
					continue
				}

				if err := comp_handler(payload.Borrower); err != nil {
					return err
				}

				var key []byte
				key = append(key, comp_liquidations...)
				key = append(key, payload.Raw.TxHash[:]...)

				if err := b.db.Put(key, payload.Raw.TxHash[:], nil); err != nil {
					if debug {
						log.Println("some issue on writing to db the ctoken liquidation")
					}
				}

				key = []byte{}

				key = append(key, user_comp_prefix...)
				key = append(key, payload.Borrower[:]...)

				if _, err := b.db.Get(key, nil); err == nil {
					if payload.Liquidator != mainnet_addrs.FLASHLOAN_LIQ_COMP_BASED {
						log.Println(
							"had this comp user - why didn't we liquidation it first",
							"\n user:\t", etherscan_of_addr(payload.Borrower),
							"\n tx-hash:\t", etherscan_of_hash(payload.Raw.TxHash),
						)
					}
				}
			}
		}

		return nil
	})

	g.Go(func() error {
		debug := *comp_debug
		count := 0
		var start time.Time
		prefix_length := len(user_comp_prefix)
		disabled := *no_comp_loop
		comp_liq, err := comp_based.NewUniswapFlashloanLiquidateCompBased(
			mainnet_addrs.FLASHLOAN_LIQ_COMP_BASED, b.clients.comp,
		)

		if err != nil {
			return err
		}

		ctoken_swapping_handles := map[common.Address]*compliq_with_handle{}

		for key, value := range mainnet_addrs.COMP_LIQUIDATIONS_PARAMS {

			pair, err := pair.NewPair(
				value.Pair, b.clients.comp,
			)

			if err != nil {
				return err
			}

			token1, err := pair.Token1(nil)

			if err != nil {
				return err
			}

			cToken, err := cerc20.NewCerc20(key, b.clients.comp)

			if err != nil {
				return err
			}

			ctoken_swapping_handles[key] = &compliq_with_handle{
				CompLiq: value,
				Pair:    pair,
				isFirst: bytes.Compare(key.Bytes(), token1.Bytes()) == -1,
				cToken:  cToken,
			}

		}

		router, err := router.NewRouter(
			mainnet_addrs.UNISWAP_ROUTER, b.clients.otherwise,
		)

		if err != nil {
			return err
		}

		ignore_estimate_err := *ignore_estimate_error
		scale_back := big.NewInt(1000000000000000000)
		write_to_estimate_gas_fail := *enable_est_fail_log
		write_to_not_worth_it := *enable_not_worth_it_log

		close_factor, err := troll.CloseFactorMantissa(nil)
		if err != nil {
			if debug {
				log.Println("something wrong on close factor why?", err)
			}
			return err
		}

		dry_run := *dry_run
		max_borrow := *enforce_max_borrows

		last_time_done := time.Now()
		is_first_time := true
		using_db := true

		var use_these_cached_addrs [][]byte

		for {
			count = 0
			start = time.Now()

			select {
			case <-b.shutdown:
				return nil
			default:
				if disabled {
					continue
				}
			}

			var iter simple_iter
			if is_first_time {
				iter = b.db.NewIterator(util.BytesPrefix(user_cream_prefix), nil)
			} else if is_first_time == false && time.Since(last_time_done) < 5*time.Minute {
				using_db = false
				iter = new_array_container(use_these_cached_addrs) //
			} else {
				use_these_cached_addrs = [][]byte{}
				using_db = true
				iter = b.db.NewIterator(util.BytesPrefix(user_cream_prefix), nil)
			}

		iterating_loop:
			for {
				select {
				case <-b.shutdown:
					iter.Release()
					return nil
				default:
					if iter.Next() {
						count++

						var user_addr common.Address

						just_key := iter.Key()

						if using_db {
							just_key = just_key[prefix_length:]
							k := make([]byte, len(just_key))
							copy(k[:], just_key[:])
							use_these_cached_addrs = append(use_these_cached_addrs, k)
							user_addr = common.BytesToAddress(just_key)
						} else {
							user_addr = common.BytesToAddress(just_key)
						}

						oops, _, shortfall, err := troll.GetAccountLiquidity(nil, user_addr)

						if err != nil {
							log.Println(err)
							continue
						}

						if oops.Cmp(common.Big0) != 0 {
							log.Printf("why some issue %v\n", oops)
							continue
						}

						if shortfall.Cmp(common.Big0) != 0 {

							ctokens, err := troll.GetAssetsIn(nil, user_addr)

							if err != nil {
								log.Println("something wrong ", err)
								continue
							}

							var (
								borrows []amt_as_token
								collats []amt_as_token
							)

							for i := range ctokens {

								cToken, exists := ctoken_swapping_handles[ctokens[i]]

								if !exists {
									log.Println("why this one doesn't exist",
										ctokens[i].Hex(),
									)
									continue
								}

								if cToken.cToken == nil {
									log.Println("okay why this one nil??", ctokens[i])
									continue
								}

								borrowed_amount, err := cToken.cToken.BorrowBalanceStored(nil, user_addr)

								if err != nil {
									log.Println("is there an error", err)
									return err
								}

								exchange_rate, err := cToken.cToken.ExchangeRateStored(nil)
								if err != nil {
									log.Println("some problem on exchange rate stored", err)
									continue
								}

								if borrowed_amount.Cmp(min_ctoken_borrowed) == 1 {
									borrows = append(borrows, amt_as_token{
										// borrowed is in the underlying unit's
										CTokenAmount:  borrowed_amount,
										Exchange_rate: exchange_rate, //
										Ctoken_addr:   ctokens[i],
									})
								}

								ctok_bal, err := cToken.cToken.BalanceOf(nil, user_addr)

								if err != nil {
									return err
								}

								if ctok_bal.Cmp(common.Big0) == 1 {

									collats = append(collats, amt_as_token{
										CTokenAmount:  ctok_bal,
										Exchange_rate: exchange_rate,
										Ctoken_addr:   ctokens[i],
									})
								}

							}

							if len(borrows) == 0 {
								continue
							}

							var (
								cheapest_price []borrow_side_params
								most_expensive []sell_side_params
							)

							// find the cheapest to borrow
							for i := range borrows {
								liq_info := ctoken_swapping_handles[borrows[i].Ctoken_addr]
								reserves, err := liq_info.Pair.GetReserves(nil)

								if err != nil {
									log.Printf("some problem on getting reserve %v\n", err)
									break iterating_loop
								}

								var (
									price *big.Int
									err_  error
								)

								// inverted logic - kinda
								if liq_info.isFirst {
									price, err_ = router.GetAmountOut(
										nil, borrows[i].CTokenAmount, reserves.Reserve1, reserves.Reserve0,
									)
								} else {
									price, err_ = router.GetAmountOut(
										nil, borrows[i].CTokenAmount, reserves.Reserve0, reserves.Reserve1,
									)
								}

								if err_ != nil {
									s, _ := json.MarshalIndent(liq_info, "", "\t")
									log.Println(
										"error on borrows loop", err_,
										borrows[i].Ctoken_addr.Hex(), string(s),
									)
									break iterating_loop
								}

								cheapest_price = append(cheapest_price, borrow_side_params{
									Price:                 price,
									UnderlyingAssetAmount: borrows[i].CTokenAmount,
									BorrowToken:           liq_info.CompLiq.Token,
									RepayCToken:           borrows[i].Ctoken_addr,
									BorrowPair:            liq_info.CompLiq.Pair,
									ExchangeRate:          borrows[i].Exchange_rate,
								})
							}

							for i := range collats {
								liq_info := ctoken_swapping_handles[collats[i].Ctoken_addr]

								reserves, err := liq_info.Pair.GetReserves(nil)

								if err != nil {
									log.Printf("some problem on getting reserve %v\n", err)
									break iterating_loop
								}
								var (
									price *big.Int
									err_  error
								)

								amt := new(big.Int).Mul(
									collats[i].CTokenAmount, collats[i].Exchange_rate,
								)

								amt.Div(amt, one_erc_20)

								if liq_info.isFirst {
									price, err_ = router.GetAmountOut(
										nil,
										amt,
										reserves.Reserve0, reserves.Reserve1,
									)

								} else {
									price, err_ = router.GetAmountOut(
										nil,
										amt,
										reserves.Reserve1, reserves.Reserve0,
									)
								}

								if err_ != nil {
									s1, _ := json.MarshalIndent(liq_info, "", "\t")
									s2, _ := json.MarshalIndent(collats[i], "", "\t")
									log.Println(
										"error on comp collats loop", err_,
										user_addr.Hex(),
										string(s1),
										string(s2),
									)
									break iterating_loop
								}

								most_expensive = append(most_expensive, sell_side_params{
									Price:                 price,
									Amount:                collats[i].CTokenAmount,
									UnderlyingAssetAmount: amt,
									SellToPair:            liq_info.CompLiq.Pair,
									ReceiveBonusCToken:    collats[i].Ctoken_addr,
									ReceiveBonusToken:     liq_info.CompLiq.Token,
									ExchangeRate:          collats[i].Exchange_rate,
								})
							}

							if len(borrows) == 0 {
								continue
							}

							sort.SliceStable(cheapest_price, func(i, j int) bool {
								return cheapest_price[i].UnderlyingAssetAmount.Cmp(
									cheapest_price[j].UnderlyingAssetAmount,
								) == 1
							})

							sort.SliceStable(most_expensive, func(i, j int) bool {
								return most_expensive[i].Price.Cmp(most_expensive[j].Price) == 1
							})

							if debug {
								s1, s2 := pretty_lists(cheapest_price, most_expensive, b.clients.comp)
								log.Println(
									"here is comp list, first is borrows, second is most expensive",
									"user is",
									user_addr.Hex(),
									s1, s2,
								)
							}

							if user_addr == common.HexToAddress("0xF478A6Fc28462fc70496EA5D36491095C8CA6aa5") {

								// args := comp_based.FlashloanLiquidateCompBasedFlashLiquidateParam{
								// 	BorrowPair:                mainnet_addrs.UNISWAP_USABLE_ADDRS[mainnet_addrs.TOKEN_ADDRS["cream"]],
								// 	BorrowToken:               mainnet_addrs.TOKEN_ADDRS["cream"],
								// 	RepayCToken:               mainnet_addrs.CR_CREAM,
								// 	ReceiveBonusCToken:        mainnet_addrs.CR_META,
								// 	ReceiveBonusToken:         mainnet_addrs.TOKEN_ADDRS["mta"],
								// 	SellToPair:                mainnet_addrs.UNISWAP_USABLE_ADDRS[mainnet_addrs.TOKEN_ADDRS["mta"]],
								// 	UnderCollaterizedCompUser: user_addr,
								// 	HowMuchBorrow:             just_cream,
								// }
								// markets_args := comp_based.FlashloanLiquidateCompBasedSetupMarkets{
								// 	Troller: mainnet_addrs.COMP_TROLLER_ADDR,
								// 	Markets: []common.Address{
								// 		args.RepayCToken, args.ReceiveBonusCToken,
								// 	},
								// }
								// suggested_gas, _ := b.clients.otherwise.SuggestGasPrice(
								// 	context.Background(),
								// )

								// opts.GasPrice = suggested_gas

								// txn2, _ := comp_liq.FlashloanUniswapThenLiquidate(
								// 	opts, args, markets_args,
								// )

								// fmt.Println("by hand emergency!", txn2.Hash().Hex())
								// continue
							}

							var (
								br *borrow_side_params
								s  sell_side_params
							)

							for _, br_ := range cheapest_price {
								if amt, exists := min_borrow_amounts_comp[br_.BorrowToken]; exists {
									if br_.UnderlyingAssetAmount.Cmp(
										amt,
									) == 1 {
										br = &br_
										break
									} else {

										if write_to_not_worth_it {
											go func() {
												b.log_update_incoming.not_worth_it_log <- &report_msg{
													Error: nil,
													Role:  "comp not worth it ",
													When:  time.Now().UTC(),
													ArgsUsed: struct {
														Token  common.Address
														Amount *big.Int
														Symbol string
													}{
														br_.BorrowToken,
														br_.UnderlyingAssetAmount,
														mainnet_addrs.REVERSE_NAMING_COMP[br_.BorrowToken],
													},
												}
											}()
										}
									}
								} else {
									fmt.Println("why this borrow token in comp for min_borrow_amounts missing!??",
										br_.BorrowToken.Hex())
								}
							}

							if br == nil {
								continue
							}

							if len(collats) > 0 {
								s = most_expensive[0]
							} else {
								s = sell_side_params{
									ReceiveBonusCToken: br.RepayCToken,
									ReceiveBonusToken:  br.BorrowToken,
									SellToPair:         br.BorrowPair,
								}
							}

							if br.BorrowToken == s.ReceiveBonusToken {
								if _, exists := mainnet_addrs.CANT_BE_REPAY_SAME[br.BorrowToken]; exists {
									if debug {
										did_have, _ := json.MarshalIndent(ctokens, "", "\t")
										log.Printf(
											"special case - can't liquidate, seize same token %s for user %s %s\n",
											br.BorrowToken.Hex(),
											user_addr.Hex(),
											string(did_have),
										)
									}
									continue
								}
							}

							br.UnderlyingAssetAmount.Mul(br.UnderlyingAssetAmount, close_factor)
							br.UnderlyingAssetAmount.Div(br.UnderlyingAssetAmount, scale_back)

							args := comp_based.FlashloanLiquidateCompBasedFlashLiquidateParam{
								BorrowPair:                br.BorrowPair,
								BorrowToken:               br.BorrowToken,
								RepayCToken:               br.RepayCToken,
								ReceiveBonusCToken:        s.ReceiveBonusCToken,
								ReceiveBonusToken:         s.ReceiveBonusToken,
								SellToPair:                s.SellToPair,
								UnderCollaterizedCompUser: user_addr,
								HowMuchBorrow:             br.UnderlyingAssetAmount,
							}

							// min borrow sanity check again apparently

							if amount, exists := min_borrow_amounts_cream[args.BorrowToken]; exists {
								if amount.Cmp(args.HowMuchBorrow) == -1 {
									if debug {
										log.Println("not enough token to borrow")
									}
									continue
								}
							} else {
								log.Println("Not possible!??", args.BorrowToken.Hex())
							}

							if max_borrow {
								if max, exists := max_borrows[args.BorrowToken]; exists {
									if args.HowMuchBorrow.Cmp(max) >= 0 {
										args.HowMuchBorrow = max
									}
								} else {
									log.Println(
										"missing something to add to maxborrow map ",
										args.BorrowToken.Hex(),
									)
								}
							}

							if args.ReceiveBonusToken == mainnet_addrs.DAI {
								args.BorrowPair = mainnet_addrs.USDC
							}

							markets_args := comp_based.FlashloanLiquidateCompBasedSetupMarkets{
								Troller: mainnet_addrs.COMP_TROLLER_ADDR,
								Markets: []common.Address{
									args.RepayCToken, args.ReceiveBonusCToken,
								},
							}

							if w_bytes := mainnet_addrs.WETH_ADDR.Bytes(); bytes.Compare(
								br.BorrowToken.Bytes(), w_bytes,
							) == 0 {
								handle := ctoken_swapping_handles[args.RepayCToken]
								token0, _ := handle.Pair.Token0(nil)
								args.AssumeBorrowingWethThenIsItFirst = bytes.Compare(
									token0.Bytes(), w_bytes,
								) == 0
							}

							claimed := asSha256(args)

							var key []byte
							key = append(key, claimed_liqs_prefix...)
							key = append(key, []byte(claimed)...)

							if _, err := b.db.Get(key, nil); err == nil {
								if debug {
									log.Println("already claimed this comp liq opprurunity", string(key))
								}
								continue
							}

							suggested_gas, err := b.clients.otherwise.SuggestGasPrice(
								context.Background(),
							)

							if err != nil {
								log.Printf("issue on suggesting gas price %v\n", err)
								continue
							}

							encoded_args, err := uni_flashloan_liquidate_comp_abi.Pack(
								"flashloan_uniswap_then_liquidate", args, markets_args,
							)

							if err != nil {
								log.Println(
									"not clear why comp liquidation encoding arg caused error ", err,
								)
								continue
							}

							gas_used, err := b.clients.otherwise.EstimateGas(
								context.Background(), ethereum.CallMsg{
									From:     opts.From,
									To:       &mainnet_addrs.FLASHLOAN_LIQ_COMP_BASED,
									Gas:      max_gas_uniswap_strategy,
									GasPrice: suggested_gas,
									Value:    big.NewInt(0),
									Data:     encoded_args,
								})

							if err != nil {
								if debug {
									s, _ := json.MarshalIndent(args, "", "\t")
									log.Println(
										"estimating gas errored ", err, string(s),
									)
								}

								if write_to_estimate_gas_fail {
									go func() {
										b.log_update_incoming.estimate_fail_log <- &report_msg{
											Error:    err,
											Role:     "running comp liquidation",
											When:     time.Now().UTC(),
											ArgsUsed: args,
											Extra:    markets_args,
										}
									}()
								}

								if ignore_estimate_err == false {
									continue
								}

							}

							total_cost := new(big.Int).Mul(
								new(big.Int).SetUint64(gas_used), suggested_gas,
							)

							suggested_gas.Add(suggested_gas, more_gas)

							opts.GasPrice = suggested_gas
							opts.GasLimit = uint64(1_300_000)

							if dry_run {
								continue
							}

							txn2, err2 := comp_liq.FlashloanUniswapThenLiquidate(
								opts, args, markets_args,
							)

							hash := txn2.Hash()

							if err := b.db.Put(key, hash[:], nil); err != nil {
								break iterating_loop
							}

							s2, _ := json.MarshalIndent(args, "", "\t")

							log.Println(
								"SUCCESS! ->total comp liqiudation cost estimated at total cost",
								total_cost, string(s2),
								etherscan_of_hash(hash),
								err2,
							)

							bind.WaitMined(context.Background(), b.clients.comp, txn2)
						}

					} else {
						// means we can break the loop
						break iterating_loop
					}
				}
			}

			iter.Release()
			is_first_time = false

			if using_db {
				last_time_done = time.Now()
			}

			if debug {
				end := time.Now()
				log.Println(
					"finished a keys loop over ",
					count, "for comp records and took", end.Sub(start).String())
			}

			if err := iter.Error(); err != nil {
				return err
			}
		}

		return nil
	})

	g.Go(func() error {
		ro := &opt.ReadOptions{
			DontFillCache: false,
		}

		uni_liquidate, _ := uniswap_flashloan_liquidate_aave.NewUniswapFlashloanLiquidateAave(
			mainnet_addrs.UNISWAP_FLASHLOAN_AAVE_LIQ_ADDR, b.clients.otherwise,
		)

		write_to_estimate_gas_fail := *enable_est_fail_log
		debug := *aave_debug
		count := 0
		var start time.Time

		data_provider_addr, _ := pool_handle.DataProvider(nil)
		aave_data_provider, _ := lending_pool_data_provider.NewLendingPoolDataProvider(
			data_provider_addr, b.clients.aave,
		)

		oracle, err := aave_price_oracle.NewAavePriceOracle(
			aave_price_oracle_addr, b.clients.aave,
		)

		if err != nil {
			return err
		}

		prefix_length := len(user_aave_prefix)
		pair_map := mainnet_addrs.UNISWAP_USABLE_ADDRS

		borrow_pair_handle, err := pair.NewPair(
			mainnet_addrs.UNISWAP_USABLE_ADDRS[mainnet_addrs.AAVE_ETH_HANDLE], b.clients.aave,
		)
		ignore_estimate_err := *ignore_estimate_error_aave
		w_bytes := mainnet_addrs.WETH_ADDR.Bytes()

		if err != nil {
			return err
		}

		max_borrow := *enforce_max_borrows

		for {
			count = 0
			start = time.Now()

			select {
			case <-b.shutdown:
				return nil
			default:
			}

			var iter iterator.Iterator

			iter = b.db.NewIterator(util.BytesPrefix(user_aave_prefix), ro)

		iterating_loop:
			for {
				select {
				case <-b.shutdown:
					iter.Release()
					return nil
				default:
					if iter.Next() {
						count++

						just_key := iter.Key()[prefix_length:]
						user_addr := common.BytesToAddress(just_key)

						if err != nil {
							log.Printf("problem getting user account data %v\n", err)
							break iterating_loop
						}

						global_data_user, err := aave_data_provider.CalculateUserGlobalData(
							nil, user_addr,
						)

						if err != nil {
							log.Printf("problem getting user global data %v\n", err)
							break iterating_loop
						}

						if global_data_user.HealthFactorBelowThreshold {
							var per_token_values []aave_liq_params
							for i := range mainnet_addrs.AAVE_RESERVE_TOKENS {

								key := mainnet_addrs.AAVE_RESERVE_TOKENS[i]

								user_reserve_data, err := pool_handle.GetUserReserveData(
									nil, key, user_addr,
								)

								if err != nil {
									log.Printf("problem getting user reserve data %v\n", err)
									return errors.Wrapf(
										err, "cant get user reserve data %s %s",
										key.Hex(), user_addr.Hex(),
									)
								}

								if user_reserve_data.CurrentATokenBalance.Cmp(common.Big0) == 0 &&
									user_reserve_data.CurrentBorrowBalance.Cmp(common.Big0) == 0 {
									continue
								}

								p := aave_liq_params{
									TokenAddr:                    key,
									Borrowed_token_balance_repay: user_reserve_data.CurrentBorrowBalance,
									Usage_as_collateral_enabled:  user_reserve_data.UsageAsCollateralEnabled,
									Collateral_bonus_amt:         common.Big0,
									Current_a_token_balance:      common.Big0,
								}

								if user_reserve_data.UsageAsCollateralEnabled &&
									user_reserve_data.CurrentATokenBalance.Cmp(common.Big0) == 1 {

									asset_price, err := oracle.GetAssetPrice(nil, key)

									if err != nil {
										return errors.Wrapf(err, "aave oracle problem for %s", key.Hex())
									}

									bonus, exists := aave_liquidation_bonus[key]

									if !exists {
										log.Println("BAD -> bonus for ", key.Hex(), "missing in aave bonus map")
										break iterating_loop
									}

									max_bonus := decimal.NewDecFromBigInt(
										user_reserve_data.CurrentATokenBalance,
									).Mul(bonus).
										Mul(decimal.NewDecFromBigInt(asset_price).Quo(factor)).
										TruncateInt()
									p.Collateral_bonus_amt = max_bonus
									p.Current_a_token_balance = user_reserve_data.CurrentATokenBalance
								}

								per_token_values = append(per_token_values, p)
							}

							var (
								maxColatBonus    *big.Int
								repaymentAmount  *big.Int
								borrowToken      common.Address
								colatRedeemToken common.Address
							)

							if debug {
								s := pretty_aave(per_token_values, b.clients.aave)
								fmt.Println("aave positions of user", user_addr.Hex(), string(s))
							}

							sort.SliceStable(per_token_values, func(i, j int) bool {
								return per_token_values[i].Borrowed_token_balance_repay.Cmp(
									per_token_values[j].Borrowed_token_balance_repay,
								) == 1
							})

							repaymentAmount = per_token_values[0].Borrowed_token_balance_repay
							if repaymentAmount.Cmp(common.Big0) == 0 {
								s4, _ := json.MarshalIndent(per_token_values, "", "\t")
								fmt.Println("how its possible", user_addr.Hex(), string(s4))
							}

							borrowToken = per_token_values[0].TokenAddr

							sort.SliceStable(per_token_values, func(i, j int) bool {
								return per_token_values[i].Collateral_bonus_amt.Cmp(
									per_token_values[j].Collateral_bonus_amt,
								) == 1
							})

							maxColatBonus = per_token_values[0].Collateral_bonus_amt
							colatRedeemToken = per_token_values[0].TokenAddr

							if borrow_eth_pair, exists := pair_map[borrowToken]; exists {
								if collat_eth_pair, exists := pair_map[colatRedeemToken]; exists {
									repaymentAmount.Div(repaymentAmount, common.Big2)

									args := uniswap_flashloan_liquidate_aave.FlashloanLiquidateAaveFlashLoanParam{
										CollateralToken:        colatRedeemToken,
										CollateralTokenEthPair: collat_eth_pair,
										BorrowToken:            borrowToken,
										BorrowTokenEthPair:     borrow_eth_pair,
										AaveUserLiquidate:      user_addr,
										HowMuchBorrow:          repaymentAmount,
									}

									if amt, exists := min_borrow_amounts_aave[args.BorrowToken]; exists {
										if args.HowMuchBorrow.Cmp(amt) == -1 {
											if *enable_not_worth_it_log {
												go func() {
													b.log_update_incoming.not_worth_it_log <- &report_msg{
														Error:    nil,
														Role:     "not worth while to do the aave trade",
														When:     time.Now().UTC(),
														ArgsUsed: args,
													}
												}()
											}
											continue
										}
									} else {
										log.Println("NEED TO ADD TO AAVE MAP!", args.BorrowToken.Hex())
										continue
									}

									if args.BorrowToken == mainnet_addrs.AAVE_ETH_HANDLE {
										args.BorrowToken = mainnet_addrs.WETH_ADDR
										token0, _ := borrow_pair_handle.Token0(nil)
										args.AssumeBorrowingWethThenIsItFirst = bytes.Compare(
											token0.Bytes(), w_bytes,
										) == 0
									}

									if max_borrow {

										if max, exists := max_borrows[args.BorrowToken]; exists {
											if args.HowMuchBorrow.Cmp(max) >= 0 {
												args.HowMuchBorrow = max
											}
										} else {
											log.Println(
												"missing something to add to maxborrow map ",
												args.BorrowToken.Hex(),
											)
										}
									}
									suggested_gas, err := b.clients.otherwise.SuggestGasPrice(
										context.Background(),
									)

									if err != nil {
										log.Printf("issue on suggesting gas price %v\n", err)
										continue
									}

									if debug {
										s, _ := json.MarshalIndent(args, "", "\t")
										log.Println("so now the aave liq params would be picked?", string(s))
									}

									encoded_args, err := uni_liq_abi.Pack(
										"flash_swap_then_liquidate", args,
									)

									if err != nil {
										log.Println("not clear why encoding arg caused error ", err)
										continue
									}

									gas_used, err := b.clients.otherwise.EstimateGas(
										context.Background(), ethereum.CallMsg{
											From:     opts.From,
											To:       &mainnet_addrs.UNISWAP_FLASHLOAN_AAVE_LIQ_ADDR,
											Gas:      max_gas_uniswap_strategy,
											GasPrice: suggested_gas,
											Value:    big.NewInt(0),
											Data:     encoded_args,
										})

									if err != nil {

										if write_to_estimate_gas_fail {
											go func() {
												b.log_update_incoming.estimate_fail_log <- &report_msg{
													Error:    err,
													Role:     "running aave liquidation",
													When:     time.Now().UTC(),
													ArgsUsed: args,
												}
											}()
										}

										if debug {
											log.Println("aave err on args", err)
										}
										if ignore_estimate_err == false {
											continue
										}

									}

									total_cost := new(big.Int).Mul(
										new(big.Int).SetUint64(gas_used), suggested_gas,
									)

									if total_cost.Cmp(maxColatBonus) == 1 {
										if debug {
											log.Println(
												"total cost aave is",
												decimal.NewDecFromBigInt(total_cost).Quo(factor),
												"but bonus is",
												decimal.NewDecFromBigInt(maxColatBonus).Quo(factor),
											)
										}
										continue
									}

									claimed := asSha256(args)

									var key []byte
									key = append(key, claimed_liqs_prefix...)
									key = append(key, []byte(claimed)...)

									if _, err := b.db.Get(key, nil); err == nil {
										log.Println("already claimed this aave opprurunity", string(key))
										continue
									}

									suggested_gas.Add(suggested_gas, more_gas)

									opts.GasPrice = suggested_gas
									opts.GasLimit = max_gas_uniswap_strategy

									txn, err := uni_liquidate.FlashSwapThenLiquidate(opts, args)

									if err != nil {
										log.Printf(
											"some issue on kicking off the flashswaps and liquidate %v\n",
											err,
										)
										break iterating_loop
									}

									hash := txn.Hash()

									if err := b.db.Put(key, hash[:], nil); err != nil {
										break iterating_loop
									}

									s, _ := json.MarshalIndent(txn, "", "\t")
									log.Println(
										"kicked off a aave liquidation txn",
										string(s),
										etherscan_of_hash(hash),
									)
									bind.WaitMined(context.Background(), b.clients.aave, txn)
								}
							}
						}

					} else {
						// means we can break the loop
						break iterating_loop
					}
				}
			}

			iter.Release()

			if debug {
				end := time.Now()
				log.Println(
					"finished a keys loop over ",
					count, "for aave records", end.Sub(start).String(),
				)
			}

			if err := iter.Error(); err != nil {
				return err
			}

		}
		return nil
	})

	return g.Wait()
}

func main() {
	flag.Parse()

	if *enable_pprof {
		defer profile.Start().Stop()
	}

	bot, err := new_bot(*db_path, *common_flags.Client_dial)

	if err != nil {
		log.Print("something wrong with making a new bot ", err)
		os.Exit(-1)
	}

	if *only_show_min_borrows {
		if err := bot.show_min_borrows(); err != nil {
			fmt.Println("something wrong doesn't make sense", err)
		}
		return
	}

	if *graph_pricing {
		pairs, err := bot.new_all_relevant_pairs()
		if err != nil {
			fmt.Println("issue", err)
			bot.close_resources()
			os.Exit(-1)
			return
		}
		t := time.NewTicker(time.Second)

		for range t.C {
			// start := time.Now()
			if err := bot.create_edges(pairs); err != nil {
				fmt.Println("issue", err)
				bot.close_resources()
				os.Exit(-1)
				return
			}
			// end := time.Now()
			// fmt.Println("took ", end.Sub(start))
		}
		bot.close_resources()
		os.Exit(0)
	}

	if *repl_mode {
		bot_handle_repl = bot
		p := prompt.New(
			executor,
			completer,
			prompt.OptionPrefix(">>> "),
			prompt.OptionTitle(*db_path+" "+"leveldb repl"),
		)
		p.Run()

		if err := bot.close_resources(); err != nil {
			log.Print("something wrong with closing resources", err)
		}

		os.Exit(0)
	}

	if *dump_stats {
		iter := bot.db.NewIterator(util.BytesPrefix(user_aave_prefix), nil)
		count_aave, count_comp, count_cream := 0, 0, 0
		for iter.Next() {
			count_aave++
		}
		iter.Release()
		iter = bot.db.NewIterator(util.BytesPrefix(user_comp_prefix), nil)
		for iter.Next() {
			count_comp++
		}
		iter.Release()
		iter = bot.db.NewIterator(util.BytesPrefix(user_cream_prefix), nil)
		for iter.Next() {
			count_cream++
		}
		fmt.Println(
			count_aave, "aave user records and ",
			count_comp, "comp user records",
			count_cream, "cream user records",
		)
		if err := bot.close_resources(); err != nil {
			log.Print("something wrong with closing resources", err)

		}
		os.Exit(0)
	}

	if err := run_liquidation_bot(bot); err != nil {
		log.Println("something wrong with running the bot", err)
	}

	if err := bot.close_resources(); err != nil {
		log.Print("something wrong with closing resources", err)
	}

}
