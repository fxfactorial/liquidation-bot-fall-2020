package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"watch_for_arb/adjustment/decimal"
	"watch_for_arb/mainnet_addrs"

	"github.com/ethereum/go-ethereum/common"
)

var (
	bat_amount = decimal.MustNewDecFromStr(
		"1000000000000000000000",
	).TruncateInt()
	usdc              = mainnet_addrs.TOKEN_ADDRS["usdc"]
	six_places_dollar = big.NewInt(600e6)
	twenty            = decimal.MustNewDecFromStr(
		"50000000000000000000",
	).TruncateInt()
	eth_min = big.NewInt(2e18)

	min_borrow_amounts_comp = map[common.Address]*big.Int{
		mainnet_addrs.TOKEN_ADDRS["rep"]:  twenty,
		mainnet_addrs.TOKEN_ADDRS["comp"]: big.NewInt(5e18),
		mainnet_addrs.WETH_ADDR:           eth_min,
		mainnet_addrs.TOKEN_ADDRS["wbtc"]: big.NewInt(3e6),
		mainnet_addrs.TOKEN_ADDRS["usdt"]: six_places_dollar,
		usdc:                              six_places_dollar,
		mainnet_addrs.TOKEN_ADDRS["dai"]: decimal.MustNewDecFromStr(
			"911000000000000000000",
		).TruncateInt(),
		mainnet_addrs.TOKEN_ADDRS["uni"]: decimal.MustNewDecFromStr(
			"300000000000000000000",
		).TruncateInt(),
		mainnet_addrs.TOKEN_ADDRS["zrx"]: decimal.MustNewDecFromStr(
			"6000000000000000000000",
		).TruncateInt(),
		mainnet_addrs.TOKEN_ADDRS["bat"]: bat_amount,
	}

	two_hundred = decimal.MustNewDecFromStr(
		"900000000000000000000",
	).TruncateInt()

	ren_min = decimal.MustNewDecFromStr(
		"2000000000000000000000",
	).TruncateInt()

	min_borrow_amounts_aave = map[common.Address]*big.Int{
		mainnet_addrs.TOKEN_ADDRS["link"]: twenty,
		mainnet_addrs.TOKEN_ADDRS["knc"]:  two_hundred,
		mainnet_addrs.TOKEN_ADDRS["lend"]: two_hundred,
		mainnet_addrs.TOKEN_ADDRS["enj"]:  two_hundred,
		mainnet_addrs.TOKEN_ADDRS["tusd"]: two_hundred,
		mainnet_addrs.TOKEN_ADDRS["ren"]:  ren_min,
		mainnet_addrs.TOKEN_ADDRS["busd"]: bat_amount,
		mainnet_addrs.TOKEN_ADDRS["susd"]: two_hundred,
		mainnet_addrs.TOKEN_ADDRS["aave"]: two_hundred,
		mainnet_addrs.TOKEN_ADDRS["snx"]:  two_hundred,
		mainnet_addrs.TOKEN_ADDRS["yfi"]:  big.NewInt(3e17),
		mainnet_addrs.TOKEN_ADDRS["mkr"]:  one_erc_20,
		mainnet_addrs.AAVE_ETH_HANDLE:     eth_min,
		mainnet_addrs.WETH_ADDR:           eth_min,
		mainnet_addrs.TOKEN_ADDRS["wbtc"]: big.NewInt(3e6),
		mainnet_addrs.TOKEN_ADDRS["usdt"]: six_places_dollar,
		mainnet_addrs.TOKEN_ADDRS["mana"]: two_hundred,
		mainnet_addrs.TOKEN_ADDRS["rep"]:  two_hundred,
		usdc:                              six_places_dollar,
		mainnet_addrs.TOKEN_ADDRS["dai"]:  two_hundred,
		mainnet_addrs.TOKEN_ADDRS["zrx"]: decimal.MustNewDecFromStr(
			"5000000000000000000000",
		).TruncateInt(),
		mainnet_addrs.TOKEN_ADDRS["bat"]: bat_amount,
	}

	big_number_eighteens_more = decimal.MustNewDecFromStr(
		"3500000000000000000000",
	).TruncateInt()

	crv = decimal.MustNewDecFromStr(
		"6500000000000000000000",
	).TruncateInt()

	hundred = decimal.MustNewDecFromStr(
		"600000000000000000000",
	).TruncateInt()
	thirty = decimal.MustNewDecFromStr(
		"30000000000000000000",
	).TruncateInt()

	min_borrow_amounts_cream = map[common.Address]*big.Int{
		mainnet_addrs.TOKEN_ADDRS["srm"]:    hundred,
		mainnet_addrs.TOKEN_ADDRS["ftx"]:    bat_amount,
		mainnet_addrs.TOKEN_ADDRS["swag"]:   big_number_eighteens_more,
		mainnet_addrs.TOKEN_ADDRS["crv"]:    crv,
		mainnet_addrs.TOKEN_ADDRS["sushi"]:  big_number_eighteens_more,
		mainnet_addrs.TOKEN_ADDRS["yweth"]:  eth_min,
		mainnet_addrs.WETH_ADDR:             eth_min,
		mainnet_addrs.TOKEN_ADDRS["wbtc"]:   big.NewInt(3e6),
		mainnet_addrs.TOKEN_ADDRS["yfi"]:    big.NewInt(5e17),
		mainnet_addrs.TOKEN_ADDRS["usdt"]:   six_places_dollar,
		mainnet_addrs.TOKEN_ADDRS["yusd"]:   hundred,
		mainnet_addrs.TOKEN_ADDRS["ycrv"]:   hundred,
		mainnet_addrs.TOKEN_ADDRS["link"]:   thirty,
		mainnet_addrs.TOKEN_ADDRS["cream"]:  thirty,
		mainnet_addrs.TOKEN_ADDRS["lend"]:   hundred,
		mainnet_addrs.TOKEN_ADDRS["renbtc"]: big.NewInt(3e7),
		mainnet_addrs.TOKEN_ADDRS["mta"]:    hundred,
		mainnet_addrs.TOKEN_ADDRS["wnxm"]:   thirty,
		mainnet_addrs.TOKEN_ADDRS["bal"]:    hundred,
		mainnet_addrs.TOKEN_ADDRS["busd"]:   hundred,
		usdc:                                six_places_dollar,
		mainnet_addrs.TOKEN_ADDRS["comp"]:   big.NewInt(8e18),
		mainnet_addrs.TOKEN_ADDRS["dai"]:    hundred,
		mainnet_addrs.TOKEN_ADDRS["uni"]: decimal.MustNewDecFromStr(
			"200000000000000000000",
		).TruncateInt(),
		mainnet_addrs.TOKEN_ADDRS["zrx"]: decimal.MustNewDecFromStr(
			"6000000000000000000000",
		).TruncateInt(),
		mainnet_addrs.TOKEN_ADDRS["bat"]: bat_amount,
	}

	fifty_five_hundred = decimal.MustNewDecFromStr(
		"5500284445941365090884",
	).TruncateInt()

	twenty_thousands = decimal.MustNewDecFromStr(
		"20000000000000000000000",
	).TruncateInt()
	twenty_thousands_dollar_coins = decimal.MustNewDecFromStr(
		"20000000000",
	).TruncateInt()
	forty_eth = decimal.MustNewDecFromStr(
		"40000000000000000000",
	).TruncateInt()
	three_btc = big.NewInt(3e8)

	eighty_eth = decimal.MustNewDecFromStr(
		"80000000000000000000",
	).TruncateInt()

	just_cream = decimal.MustNewDecFromStr(
		"40000000000000000000",
	).TruncateInt()

	max_borrows = map[common.Address]*big.Int{
		mainnet_addrs.TOKEN_ADDRS["srm"]: twenty_thousands,
		mainnet_addrs.TOKEN_ADDRS["ftx"]: twenty_thousands,
		mainnet_addrs.TOKEN_ADDRS["swag"]: new(big.Int).Mul(
			twenty_thousands, common.Big3,
		),
		mainnet_addrs.TOKEN_ADDRS["crv"]:    twenty_thousands,
		mainnet_addrs.TOKEN_ADDRS["sushi"]:  twenty_thousands,
		mainnet_addrs.TOKEN_ADDRS["yweth"]:  forty_eth,
		mainnet_addrs.WETH_ADDR:             forty_eth,
		mainnet_addrs.TOKEN_ADDRS["wbtc"]:   three_btc,
		mainnet_addrs.TOKEN_ADDRS["yfi"]:    new(big.Int).Div(forty_eth, common.Big2),
		mainnet_addrs.TOKEN_ADDRS["usdt"]:   twenty_thousands_dollar_coins,
		mainnet_addrs.TOKEN_ADDRS["yusd"]:   twenty_thousands,
		mainnet_addrs.TOKEN_ADDRS["ycrv"]:   twenty_thousands,
		mainnet_addrs.TOKEN_ADDRS["link"]:   new(big.Int).Mul(forty_eth, big.NewInt(16)),
		mainnet_addrs.TOKEN_ADDRS["cream"]:  just_cream,
		mainnet_addrs.TOKEN_ADDRS["lend"]:   twenty_thousands,
		mainnet_addrs.TOKEN_ADDRS["renbtc"]: three_btc,
		mainnet_addrs.TOKEN_ADDRS["mta"]:    fifty_five_hundred,
		mainnet_addrs.TOKEN_ADDRS["wnxm"]:   forty_eth,
		mainnet_addrs.TOKEN_ADDRS["bal"]:    new(big.Int).Mul(forty_eth, big.NewInt(16)),
		mainnet_addrs.TOKEN_ADDRS["busd"]:   twenty_thousands,
		usdc:                                twenty_thousands_dollar_coins,
		mainnet_addrs.TOKEN_ADDRS["comp"]:   new(big.Int).Mul(forty_eth, common.Big2),
		mainnet_addrs.TOKEN_ADDRS["dai"]:    twenty_thousands,
		mainnet_addrs.TOKEN_ADDRS["uni"]:    twenty_thousands,
		mainnet_addrs.TOKEN_ADDRS["zrx"]:    twenty_thousands,
		mainnet_addrs.TOKEN_ADDRS["bat"]:    twenty_thousands,
	}
)

func (b *bot) show_min_borrows() error {

	type market struct {
		Name       string
		MinBorrows []record
	}

	aave := market{Name: "aave", MinBorrows: []record{}}
	comp := market{Name: "comp", MinBorrows: []record{}}
	cream := market{Name: "cream", MinBorrows: []record{}}

	for key, amount := range min_borrow_amounts_aave {
		symbol, amount := pretty_token(key, amount, b.clients.otherwise)
		aave.MinBorrows = append(aave.MinBorrows, record{
			Addr:   key,
			Symbol: symbol,
			Amount: amount,
		})
	}

	for key, amount := range min_borrow_amounts_comp {
		symbol, amount := pretty_token(key, amount, b.clients.otherwise)
		comp.MinBorrows = append(comp.MinBorrows, record{
			Addr:   key,
			Symbol: symbol,
			Amount: amount,
		})
	}

	for key, amount := range min_borrow_amounts_cream {
		symbol, amount := pretty_token(key, amount, b.clients.otherwise)
		cream.MinBorrows = append(cream.MinBorrows, record{
			Addr:   key,
			Symbol: symbol,
			Amount: amount,
		})
	}

	s1, err := json.MarshalIndent([]market{aave, comp, cream}, "", "\t")
	if err != nil {
		return err
	}

	fmt.Println(string(s1))
	return nil
}
