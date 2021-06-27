package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"watch_for_arb/adjustment/decimal"
	"watch_for_arb/contracts/uniswap/router"
	"watch_for_arb/mainnet_addrs"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"gonum.org/v1/gonum/stat/combin"
)

var (
	// idea = []string{"usdt", "usdc", "dai", "wbtc", "tusd"}

	eight_hundred = decimal.MustNewDecFromStr(
		"800000000000000000000",
	).TruncateInt()
	one_thousand = big.NewInt(1000e6)
	nine_eth     = big.NewInt(9e18)
	start_amount = map[common.Address]*big.Int{
		mainnet_addrs.TOKEN_ADDRS["usdt"]:   one_thousand,
		mainnet_addrs.TOKEN_ADDRS["usdc"]:   one_thousand,
		mainnet_addrs.WETH_ADDR:             nine_eth,
		mainnet_addrs.TOKEN_ADDRS["dai"]:    eight_hundred,
		mainnet_addrs.TOKEN_ADDRS["wbtc"]:   big.NewInt(1e8),
		mainnet_addrs.TOKEN_ADDRS["tusd"]:   eight_hundred,
		mainnet_addrs.TOKEN_ADDRS["yfi"]:    big.NewInt(1e18),
		mainnet_addrs.TOKEN_ADDRS["renbtc"]: big.NewInt(1e8),
		mainnet_addrs.TOKEN_ADDRS["yweth"]:  nine_eth,
		mainnet_addrs.TOKEN_ADDRS["mana"]:   eight_hundred,
		mainnet_addrs.TOKEN_ADDRS["snx"]:    eight_hundred,
		mainnet_addrs.TOKEN_ADDRS["mkr"]:    nine_eth,
		mainnet_addrs.TOKEN_ADDRS["lend"]:   eight_hundred,
		mainnet_addrs.TOKEN_ADDRS["comp"]:   nine_eth,
		mainnet_addrs.TOKEN_ADDRS["busd"]:   eight_hundred,
		mainnet_addrs.TOKEN_ADDRS["yusd"]:   eight_hundred,
		// mainnet_addrs.TOKEN_ADDRS[""]:     nine_eth,
		// mainnet_addrs.TOKEN_ADDRS[""]:     nine_eth,
		// mainnet_addrs.TOKEN_ADDRS[""]:     nine_eth,
		// mainnet_addrs.TOKEN_ADDRS[""]:     nine_eth,
		// mainnet_addrs.TOKEN_ADDRS[""]:     nine_eth,
		// mainnet_addrs.TOKEN_ADDRS[""]:     nine_eth,
		// mainnet_addrs.TOKEN_ADDRS[""]:     nine_eth,
		// mainnet_addrs.TOKEN_ADDRS[""]:     nine_eth,
	}
)

var (
	idea = []common.Address{
		mainnet_addrs.TOKEN_ADDRS["usdt"],
		mainnet_addrs.TOKEN_ADDRS["usdc"],
		mainnet_addrs.TOKEN_ADDRS["dai"],
		mainnet_addrs.TOKEN_ADDRS["wbtc"],
		mainnet_addrs.TOKEN_ADDRS["tusd"],
		mainnet_addrs.TOKEN_ADDRS["yfi"],
		mainnet_addrs.TOKEN_ADDRS["renbtc"],
		mainnet_addrs.TOKEN_ADDRS["yweth"],
		mainnet_addrs.TOKEN_ADDRS["mana"],
		mainnet_addrs.TOKEN_ADDRS["snx"],
		mainnet_addrs.TOKEN_ADDRS["mkr"],
		mainnet_addrs.TOKEN_ADDRS["lend"],
		mainnet_addrs.TOKEN_ADDRS["comp"],
		mainnet_addrs.TOKEN_ADDRS["busd"],
		mainnet_addrs.TOKEN_ADDRS["yusd"],
	}
	weth = mainnet_addrs.WETH_ADDR
)

func all_paths(r *router.Router) error {

	// for i := range scenarios {

	n := len(idea)
	k := 2

	// gen := combin.NewPermutationGenerator(n, k)
	gen := combin.NewCombinationGenerator(n, k)

	var keep_these [][]common.Address

	unique_paths := map[string]struct{}{}

	for gen.Next() {
		// temp := gen.Permutation(nil)
		temp := gen.Combination(nil)

		k0, k1 := idea[temp[0]], idea[temp[1]]

		for _, p := range [][]common.Address{
			{weth, k0, k1, weth},
			{weth, k1, k0, weth},
			{k0, weth, k1, k0},
			{k1, weth, k0, k1},
		} {
			hash := asSha256(p)
			if _, exists := revert_lookup[hash]; exists {
				continue
			}

			if _, have := unique_paths[hash]; !have {
				keep_these = append(keep_these, p)
				unique_paths[hash] = struct{}{}
			}

		}
	}
	var junk []string

	for i := range keep_these {
		start_with := start_amount[keep_these[i][0]]
		results, err := r.GetAmountsOut(nil, start_with, keep_these[i])
		if err != nil {
			if err.Error() == vm.ErrExecutionReverted.Error() {
				junk = append(junk, asSha256(keep_these[i]))
				continue
			} else {

				junk = append(junk, asSha256(keep_these[i]))
			}
			continue
		}

		if end_with := results[len(results)-1]; end_with.Cmp(
			start_with,
		) == 1 {
			s, err := json.MarshalIndent(results, "", "\t")
			if err != nil {
				return err
			}
			fmt.Println("ARB!", string(s))

		} else {
			// fmt.Println("no arb -> start out with ", start_with, end_with)
		}
	}

	if len(junk) > 0 {
		s3, _ := json.MarshalIndent(junk, "", "\t")
		fmt.Println("revert tbale", string(s3))

	}
	return nil
}
