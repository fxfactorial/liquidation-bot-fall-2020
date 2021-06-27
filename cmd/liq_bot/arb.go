package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"watch_for_arb/adjustment/decimal"
	"watch_for_arb/contracts/erc20"
	"watch_for_arb/contracts/uniswap/pair"
	"watch_for_arb/contracts/uniswap/router"
	"watch_for_arb/mainnet_addrs"

	"github.com/ALTree/bigfloat"
	"github.com/ethereum/go-ethereum/common"
)

type pair_ struct {
	Token0 common.Address
	Token1 common.Address
	Price  *big.Float
}

var (
	a         = mainnet_addrs.TOKEN_ADDRS
	inf       = new(big.Float).SetInf(true)
	neg_one   = new(big.Float).SetFloat64(-1)
	plain_one = new(big.Float).SetFloat64(1.1)
	one       = new(big.Float).SetFloat64(1)
	nine_nine = new(big.Float).SetFloat64(0.99)
	zero      = new(big.Float).SetFloat64(0)
	ten       = new(big.Int).SetInt64(10)

	table = [][]pair_{
		{
			{a["weth"], a["weth"], zero},
			{a["weth"], a["usdt"], big.NewFloat(232)},
			{a["weth"], a["usdc"], big.NewFloat(234)},
		},
		{
			{a["usdt"], a["weth"], big.NewFloat(1.0 / 230.0)},
			{a["usdt"], a["usdt"], nine_nine},
			{a["usdt"], a["usdc"], zero},
		},
		{
			{a["usdc"], a["weth"], big.NewFloat(1.0 / 230.0)},
			{a["usdc"], a["usdt"], plain_one},
			{a["usdc"], a["usdc"], zero},
		},
	}
)

func arbitrage(table [][]pair_) bool {
	transformed_graph := make([][]*big.Float, len(table))

	// pair table
	for i := range transformed_graph {
		transformed_graph[i] = make([]*big.Float, len(table[i]))
		for j := range table[i] {
			fmt.Println("price", table[i][j].Price)
			lg := bigfloat.Log(table[i][j].Price)
			lg.Mul(lg, neg_one)
			fmt.Println("what is val lg", lg)
			transformed_graph[i][j] = lg
		}
	}

	s_graph, _ := json.MarshalIndent(transformed_graph, "", "\t")
	fmt.Println(string(s_graph))

	source := 0
	n := len(transformed_graph)
	min_dist := make([]*big.Float, n)

	for i := range min_dist {
		min_dist[i] = inf
	}

	min_dist[source] = new(big.Float).SetInt(common.Big0)

	s3, _ := json.MarshalIndent(min_dist, "", "\t")

	fmt.Println("min dist first time", string(s3))

	for i := 0; i < n-1; i++ {
		for v := 0; v < n; v++ {
			for w := 0; w < n; w++ {

				lhs := min_dist[w]

				rhs := new(big.Float).Add(
					min_dist[v], transformed_graph[v][w],
				)

				fmt.Println("rhs->", rhs)
				if lhs.Cmp(rhs) == 1 {
					fmt.Println("setting lhds")
					min_dist[w] = rhs
				}
			}
		}
	}

	s2, _ := json.MarshalIndent(min_dist, "", "\t")

	fmt.Println("min dist second time", string(s2))

	for v := 0; v < n; v++ {
		for w := 0; w < n; w++ {
			v_add_vw := new(big.Float).Add(
				min_dist[v], transformed_graph[v][w],
			)

			if min_dist[w].Cmp(v_add_vw) == 1 {
				return true
			}

		}
	}

	return false
}

type edge struct {
	U, V string
	W    float64
}

func video() {
	// u, v are the nodes - w is the weight
	edges := []edge{
		{"s", "e", 8},
		{"s", "a", 10},
		{"e", "d", 1},
		{"d", "a", -4},
		{"d", "c", -1},
		{"c", "b", -2},
		{"a", "c", 2},
		{"b", "a", 1},
	}

	nodes := map[string]int{
		"s": 0, "a": 1, "b": 2,
		"c": 3, "d": 4, "e": 5,
	}

	// nodes_two := []string{"s", "a", "b", "c", "d", "e"}

	distances := make([]float64, len(nodes))

	for i := range distances {
		distances[i] = math.Inf(1)
	}

	distances[0] = 0
	// _ = nodes_two
	// s2, _ := json.MarshalIndent(edges, "", " ")
	// fmt.Println(string(s2))

	for i := 0; i < len(nodes)-1; i++ {
		for _, edge := range edges {
			if distances[nodes[edge.U]]+edge.W < distances[nodes[edge.V]] {
				distances[nodes[edge.V]] = distances[nodes[edge.U]] + edge.W
			}
		}
	}

	s, _ := json.MarshalIndent(distances, "", " ")
	fmt.Println(string(s))
}

type currency_edge struct {
	FromToken, ToToken string
	Price              decimal.Dec
	PairAddress        common.Address
}

func (b *bot) new_all_relevant_pairs() ([]*pair.Pair, error) {

	var pairs []*pair.Pair
	var use_pair_addrs []common.Address

	for _, value := range mainnet_addrs.UNISWAP_USABLE_ADDRS {
		use_pair_addrs = append(use_pair_addrs, value)
	}

	for _, pair_addr := range append(use_pair_addrs, []common.Address{
		common.HexToAddress("0x3041cbd36888becc7bbcbc0045e3b1f144466f5f"),
		common.HexToAddress("0xb20bd5d04be54f870d5c0d3ca85d82b34b836405"),
		common.HexToAddress("0xb4e16d0168e52d35cacd2c6185b44281ec28c9dc"),
		common.HexToAddress("0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852"),
	}...) {
		pair, err := pair.NewPair(pair_addr, b.clients.otherwise)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, pair)
	}
	return pairs, nil
}

type price_quote struct {
	TokenIn       string
	TokenOut      string
	PriceInToOut  *big.Float
	PriceNegOfLog *big.Float
}

func (b *bot) create_edges(pairs []*pair.Pair) error {

	var quotes []price_quote
	router, _ := router.NewRouter(mainnet_addrs.UNISWAP_ROUTER, b.clients.otherwise)

	for _, pair := range pairs {

		token0, err := pair.Token0(nil)

		if err != nil {
			return err
		}

		token1, err := pair.Token1(nil)

		if err != nil {
			return err
		}

		contract0, err := erc20.NewErc20(token0, b.clients.otherwise)
		if err != nil {
			return err
		}

		contract1, err := erc20.NewErc20(token1, b.clients.otherwise)
		if err != nil {
			return err
		}

		resrs, err := pair.GetReserves(nil)
		if err != nil {
			return err
		}

		symbol0, err := contract0.Symbol(nil)
		if err != nil {
			// exemption for maker
			if token0 == common.HexToAddress("0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2") {
				symbol0 = "MKR"
			} else {
				return err
			}
		}

		symbol1, err := contract1.Symbol(nil)
		if err != nil {
			return err
		}

		dec0, err := contract0.Decimals(nil)

		if err != nil {
			return err
		}

		dec1, err := contract1.Decimals(nil)

		if err != nil {
			return err
		}

		one_token0 := new(big.Int).Exp(ten, big.NewInt(int64(dec0)), nil)
		one_token1 := new(big.Int).Exp(ten, big.NewInt(int64(dec1)), nil)

		// fmt.Println(symbol0, symbol1, one_token0, one_token1, resrs)

		price_0_to_1, err := router.GetAmountOut(
			nil, one_token0, resrs.Reserve0, resrs.Reserve1,
		)

		if err != nil {
			return err
		}

		price_1_to_0, err := router.GetAmountOut(
			nil, one_token1, resrs.Reserve1, resrs.Reserve0,
		)

		if err != nil {
			return err
		}

		p0 := new(big.Float).SetInt(price_0_to_1)
		p0.Quo(p0, new(big.Float).SetInt(one_token1))

		p1 := new(big.Float).SetInt(price_1_to_0)
		p1.Quo(p1, new(big.Float).SetInt(one_token0))

		p0_neg_log := bigfloat.Log(p0)
		p0_neg_log.Mul(p0_neg_log, neg_one)

		p1_neg_log := bigfloat.Log(p1)
		p1_neg_log.Mul(p1_neg_log, neg_one)

		quotes = append(quotes, price_quote{
			TokenIn: symbol0, TokenOut: symbol1,
			PriceInToOut: p0, PriceNegOfLog: p0_neg_log,
		}, price_quote{
			TokenIn: symbol1, TokenOut: symbol0,
			PriceInToOut: p1, PriceNegOfLog: p1_neg_log,
		})
	}

	// fmt.Println(quotes)

	// s2, err := json.MarshalIndent(quotes, "", " ")
	// if err != nil {
	// 	return err
	// }

	// fmt.Println("quotes", string(s2))

	// note the symbols must be upper case - cause that what erc20 give back
	nodes := map[string]int{
		"WETH": 0, "USDT": 1, "DAI": 2, "USDC": 3,
		"REN": 4, "MKR": 5, "renBTC": 6, "SUSHI": 7,
		"FTX Token": 8, "LINK": 9, "wNXM": 10,
		"CRV": 11, "BAL": 12, "SWAG": 13, "UNI": 14,
		"ENJ": 15, "yyDAI+yUSDC+yUSDT+yTUSD": 16,
		"COMP": 17, "LEND": 18, "SRM": 19, "ZRX": 20,
		"SNX": 21, "BAT": 22, "MANA": 23, "LUA": 24,
		"REP": 25, "BUSD": 26, "KNC": 27, "TUSD": 28,
		"WBTC": 29, "yDAI+yUSDC+yUSDT+yTUSD": 30,
		"CREAM": 31, "yWETH": 32, "MTA": 33,
		"YFI": 34, "AAVE": 35, "sUSD": 36,
	}
	distances := make([]float64, len(nodes))

	for i := range distances {
		distances[i] = math.Inf(1)
	}

	distances[0] = 0

	paths := map[string][]string{
		"WETH": {},
		"DAI":  {},
		"USDC": {},
		"USDT": {},
		"WBTC": {},
		"TUSD": {},
		"SNX":  {},
		"SUSD": {},
	}

	for i := 0; i < len(nodes)-1; i++ {
		for _, edge := range quotes {
			cost, _ := edge.PriceNegOfLog.Float64()
			token_in, exists := nodes[edge.TokenIn]

			if !exists {
				fmt.Println("BAD -> assign it", edge.TokenIn)
				continue
			}

			token_out, exists := nodes[edge.TokenOut]

			if !exists {
				fmt.Println("BAD -> assign it", edge.TokenOut)
				continue
			}

			a := distances[token_in]
			b := distances[token_out]

			if a+cost < b {
				// fmt.Println("before distances were", distances)
				distances[nodes[edge.TokenOut]] = a + cost
				// fmt.Println("here", a, cost, b, edge.TokenIn, edge.TokenOut)
			}

		}
	}

	// s4, _ := json.MarshalIndent(paths, "", "\t")
	// fmt.Println("next", string(s4))
	_ = paths
	for i := 0; i < len(nodes); i++ {
		for _, edge := range quotes {
			cost, _ := edge.PriceNegOfLog.Float64()
			a := distances[nodes[edge.TokenIn]]
			b := distances[nodes[edge.TokenOut]]

			if a+cost < b {
				// s, _ := json.MarshalIndent(quotes, "", "\t")
				fmt.Println(
					"relax still possible arbing!?",
					edge.TokenIn, edge.TokenOut,
					a, cost, b, //string(s),
				)
			}
		}
	}

	return nil

}
