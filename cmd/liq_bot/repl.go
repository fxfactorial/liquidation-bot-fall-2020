package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"watch_for_arb/contracts/compound/comptroller_g4"
	"watch_for_arb/contracts/custom_contracts/uniswap_flashloan_liquidate_comp_based"
	"watch_for_arb/contracts/erc20"
	"watch_for_arb/mainnet_addrs"

	"github.com/c-bata/go-prompt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/syndtr/goleveldb/leveldb/util"
)

const (
	inspect_aave_record           = "inspect_aave_record"
	inspect_comp_record           = "inspect_comp_record"
	inspect_claimed               = "inspect_claimed"
	add_cream_user                = "add_cream_user"
	show_all_aave_records         = "show_all_aave_records"
	show_all_comp_records         = "show_all_comp_records"
	show_all_cream_records        = "show_all_cream_records"
	delete_key                    = "delete_key"
	show_all_claimed              = "show_all_claimed"
	show_all_aave_liquidations    = "show_all_aave_liquidations"
	show_all_comp_liquidations    = "show_all_comp_liquidations"
	kill_contract                 = "kill_contract"
	transfer_asset_weth_cream_liq = "transfer_asset_weth_cream_liq"
	delete_all_claimed            = "delete_all_claimed"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: inspect_aave_record, Description: "check aave record"},
		{Text: inspect_comp_record, Description: "check comp record"},
		{Text: delete_key, Description: "delete a record from the db"},
		{Text: add_cream_user, Description: "add cream addr "},
		{Text: show_all_aave_records, Description: "show all the aave keys "},
		{Text: show_all_comp_records, Description: "show all the comp keys"},
		{Text: show_all_cream_records, Description: "show all the cream users keys"},
		{Text: show_all_claimed, Description: "show all the claimed liquidations"},
		{Text: inspect_claimed, Description: "inspect the claimed txn"},
		{Text: delete_all_claimed, Description: "delete all the claimed liquidations"},
		{Text: kill_contract, Description: "end the contract"},
		{Text: transfer_asset_weth_cream_liq, Description: "move out an asset"},
		{Text: show_all_comp_liquidations, Description: "dump all the comp liquidations in db"},
		{Text: show_all_aave_liquidations, Description: "dump all the aave liquidations in db"},
		{Text: "exit", Description: "exit the repl"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

var bot_handle_repl *bot

func executor(in string) {
	in = strings.TrimSpace(in)
	split := strings.Split(in, " ")
	var keys []common.Address
	var links []string

	switch split[0] {
	case delete_key, inspect_aave_record,
		add_cream_user,
		inspect_comp_record, inspect_claimed:
		if len(split) < 2 {
			log.Println("need the argument of key")
			return
		}
	}

	switch split[0] {

	case add_cream_user:
		cream_troll, err := comptroller_g4.NewComptrollerG4(
			mainnet_addrs.CREAM_TROLLER_ADDR, bot_handle_repl.clients.cream,
		)
		if err != nil {
			fmt.Println("some error ", err)
			return
		}
		user := common.HexToAddress(split[1])
		error, liquidity, shortfall, err := cream_troll.GetAccountLiquidity(nil, user)

		if error.Cmp(common.Big0) != 0 {
			fmt.Println("some error", err)
			return
		}

		var key []byte
		key = append(key, user_cream_prefix...)
		key = append(key, user[:]...)

		value, err := json.Marshal(comp_user_account{
			UserAddr:  user,
			Liquidity: liquidity,
			ShortFall: shortfall,
		})

		fmt.Println("user ", liquidity, shortfall)

		if err != nil {
			fmt.Println("some error ", err)
			return
		}

		if err := bot_handle_repl.db.Put(key, value, nil); err != nil {
			fmt.Println("some error writing ", err)
		}

	case delete_all_claimed:
		iter := bot_handle_repl.db.NewIterator(util.BytesPrefix(claimed_liqs_prefix), nil)
		for iter.Next() {
			bot_handle_repl.db.Delete(iter.Key(), nil)
		}
		iter.Release()
		fmt.Println("deleted all claimed")

	case inspect_claimed:
		payload, err := bot_handle_repl.db.Get([]byte(split[1]), nil)
		if err != nil {
			fmt.Println("does not exist in db")
			return
		}
		hash := common.BytesToHash(payload)
		fmt.Println("txn hash is ", hash.Hex())

	case show_all_claimed:
		iter := bot_handle_repl.db.NewIterator(util.BytesPrefix(claimed_liqs_prefix), nil)
		var keys []string
		for iter.Next() {
			keys = append(keys, string(iter.Key()))
		}
		iter.Release()
		s, _ := json.MarshalIndent(keys, "", "\t")
		fmt.Println(string(s))

	case transfer_asset_weth_cream_liq:
		handle, err := uniswap_flashloan_liquidate_comp_based.NewUniswapFlashloanLiquidateCompBased(
			mainnet_addrs.FLASHLOAN_LIQ_COMP_BASED,
			bot_handle_repl.clients.otherwise,
		)

		if err != nil {
			fmt.Println("something wrong", err)
		}

		opts.GasPrice = big.NewInt(20e9)

		addr := mainnet_addrs.TOKEN_ADDRS["swag"]

		token, err := erc20.NewErc20(addr, bot_handle_repl.clients.otherwise)
		if err != nil {
			fmt.Println("thinkg", err)
			return
		}
		weth, _ := erc20.NewErc20(mainnet_addrs.WETH_ADDR, bot_handle_repl.clients.otherwise)

		bal, err := token.BalanceOf(nil, mainnet_addrs.FLASHLOAN_LIQ_COMP_BASED)
		bal_weth, _ := weth.BalanceOf(nil, mainnet_addrs.FLASHLOAN_LIQ_COMP_BASED)

		if bal.Cmp(common.Big0) == 0 || bal_weth.Cmp(common.Big0) == 0 {
			fmt.Println("doing a 0 amount transfer is dumb")
			return
		}

		if err != nil {
			fmt.Println("thinkg", err)
			return
		}

		txn, _ := handle.Payout(opts, []uniswap_flashloan_liquidate_comp_based.WorkerPayouts{
			{addr, bal},
			{mainnet_addrs.WETH_ADDR, bal_weth},
		})

		fmt.Println("see txn", etherscan_of_hash(txn.Hash()))

	case kill_contract:
		cream_handle, _ := uniswap_flashloan_liquidate_comp_based.NewUniswapFlashloanLiquidateCompBased(
			mainnet_addrs.DEAD_CONTRACT, bot_handle_repl.clients.otherwise,
		)
		opts.GasPrice = big.NewInt(21e9)
		txn, err := cream_handle.KillYourself(opts)
		fmt.Println(txn.Hash().Hex(), err)

	case show_all_comp_records:
		iter := bot_handle_repl.db.NewIterator(util.BytesPrefix(user_comp_prefix), nil)
		prefix_length := len(user_comp_prefix)
		for iter.Next() {
			just_key := iter.Key()[prefix_length:]
			user_addr := common.BytesToAddress(just_key)
			keys = append(keys, user_addr)
		}
		iter.Release()
		s, _ := json.MarshalIndent(keys, "", "\t")
		fmt.Println(string(s))

	case show_all_cream_records:
		iter := bot_handle_repl.db.NewIterator(util.BytesPrefix(user_cream_prefix), nil)
		prefix_length := len(user_cream_prefix)
		for iter.Next() {
			just_key := iter.Key()[prefix_length:]
			user_addr := common.BytesToAddress(just_key)
			keys = append(keys, user_addr)
		}
		iter.Release()
		s, _ := json.MarshalIndent(keys, "", "\t")
		fmt.Println(string(s))

	case delete_key:
		if err := bot_handle_repl.db.Delete(
			[]byte(split[1]), nil,
		); err != nil {
			fmt.Println("some problem on delete the key ", err)
			return
		}

	case inspect_aave_record:

		payload, err := bot_handle_repl.db.Get([]byte(split[1]), nil)
		if err != nil {
			fmt.Println("this key did not exist in db ", split[1])
			return
		}

		var u aave_user_account
		if err := json.Unmarshal(payload, &u); err != nil {
			fmt.Println("some problem with decoding this ", err)
			return
		}
		s, _ := json.MarshalIndent(u, "", "\t")
		fmt.Println(string(s))
	case show_all_aave_records:
		iter := bot_handle_repl.db.NewIterator(util.BytesPrefix(user_aave_prefix), nil)
		prefix_length := len(user_aave_prefix)

		for iter.Next() {
			just_key := iter.Key()[prefix_length:]
			user_addr := common.BytesToAddress(just_key)
			keys = append(keys, user_addr)
		}

		iter.Release()
		s, _ := json.MarshalIndent(keys, "", "\t")
		fmt.Println(string(s))
	case inspect_comp_record:
		var key []byte
		key = append(key, user_comp_prefix...)
		key = append(key, []byte(split[1])...)

		payload, err := bot_handle_repl.db.Get(key, nil)
		if err != nil {
			fmt.Println("this key did not exist in db ", string(key))
			return
		}

		var u comp_user_account
		if err := json.Unmarshal(payload, &u); err != nil {
			fmt.Println("some problem with decoding this ", err)
			return
		}
		s, _ := json.MarshalIndent(u, "", "\t")
		fmt.Println(string(s))

	case show_all_comp_liquidations:
		iter := bot_handle_repl.db.NewIterator(util.BytesPrefix(comp_liquidations), nil)

		for iter.Next() {
			links = append(links, etherscan_of_hash(common.BytesToHash(iter.Value())))
		}

		iter.Release()
		s, _ := json.MarshalIndent(links, "", "\t")
		fmt.Println(string(s))

	case show_all_aave_liquidations:
		iter := bot_handle_repl.db.NewIterator(util.BytesPrefix(aave_liquidations), nil)

		for iter.Next() {
			links = append(links, etherscan_of_hash(common.BytesToHash(iter.Value())))
		}

		iter.Release()
		s, _ := json.MarshalIndent(links, "", "\t")
		fmt.Println(string(s))

	case "exit":
		if err := bot_handle_repl.close_resources(); err != nil {
			log.Print("something wrong with closing resources", err)
		}

		os.Exit(0)
	default:
		fmt.Println("unknown input")
	}

}
