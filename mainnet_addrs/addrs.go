package mainnet_addrs

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	router_mainnet_addr  = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
	mainnet_addr         = ""
	arb_contract_addr    = ""
	factory_mainnet_addr = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
)

var (
	MY_MAINNET_ADDR      = common.HexToAddress(mainnet_addr)
	UNISWAP_ROUTER       = common.HexToAddress(router_mainnet_addr)
	ARB_CONTRACT_ADDR    = common.HexToAddress(arb_contract_addr)
	UNISWAP_FACTORY_ADDR = common.HexToAddress(factory_mainnet_addr)
	WETH_ADDR            = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
)

var (
	UNISWAP_FLASHLOAN_AAVE_LIQ_ADDR = common.HexToAddress("")
	// TODO these two gotta go
	// UNISWAP_FLASHLOAN_COMP_LIQ  = common.HexToAddress("")
	// UNISWAP_FLASHLOAN_CREAM_LIQ = common.HexToAddress("")
	FLASHLOAN_LIQ_COMP_BASED = common.HexToAddress("")
)

var (
	// TODO Find a way to reinitialize this contract - pull out the link token

	ARB_SWAPS = common.HexToAddress("")
	// NOTE just a help for when to use the repl mode
	DEAD_CONTRACT = common.HexToAddress("")
	// TODO delete this one
	// UNISWAP_FLASHLOAN_COMP_LIQ  = common.HexToAddress("")
	CREAM_SWAP_ARB = common.HexToAddress("")
)

type CompLiq struct {
	Token         common.Address
	TokenDecimals *big.Int
	Pair          common.Address
}

var (
	AAVE_ADDR_PROVIDER = common.HexToAddress("0x24a42fD28C976A61Df5D00D0599C34c4f90748c8")
	COMP_TROLLER_ADDR  = common.HexToAddress("0x3d9819210A31b4961b30EF54bE2aeD79B9c9Cd3B")
	CREAM_TROLLER_ADDR = common.HexToAddress("0x3d5BC3c8d13dcB8bF317092d84783c2697AE9258")
	cBAT               = common.HexToAddress("0x6c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e")
	CDAI               = common.HexToAddress("0x5d3a536e4d6dbd6114cc1ead35777bab948e3643")
	cETH               = common.HexToAddress("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5")
	CETH               = cETH
	CREP               = common.HexToAddress("0x158079ee67fce2f58472a96584a73c7ab9ac95c1")
	CSAI               = common.HexToAddress("0xf5dce57282a584d2746faf1593d3121fcac444dc")
	cUNI               = common.HexToAddress("0x35a18000230da775cac24873d00ff85bccded550")
	cUSDC              = common.HexToAddress("0x39aa39c021dfbae8fac545936693ac917d5e7563")
	cUSDT              = common.HexToAddress("0xf650c3d88d12db855b8bf7d11be6c55a4e07dcc9")
	CUSDT              = cUSDT
	cWBTC              = common.HexToAddress("0xc11b1268c1a384e55c48c2391d8d480264a3a7f4")
	cZRX               = common.HexToAddress("0xb3319f5d18bc0d84dd1b4825dcde5d5f7266d407")
	CCOMP              = common.HexToAddress("0x70e36f6BF80a52b3B46b3aF8e106CC0ed743E8e4")
	COMP_cTOKENS       = []common.Address{
		cBAT, CDAI, cUNI, cUSDC, cUSDT, cWBTC, cZRX, CETH, CCOMP, CREP, CSAI,
	}
	AAVE_ETH_HANDLE     = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
	AAVE_RESERVE_TOKENS = []common.Address{
		TOKEN_ADDRS["busd"],
		TOKEN_ADDRS["dai"],
		TOKEN_ADDRS["susd"],
		TOKEN_ADDRS["tusd"],
		TOKEN_ADDRS["usdc"],
		TOKEN_ADDRS["usdt"],
		TOKEN_ADDRS["bat"],
		TOKEN_ADDRS["enj"],
		AAVE_ETH_HANDLE,
		TOKEN_ADDRS["knc"],
		TOKEN_ADDRS["aave"],
		TOKEN_ADDRS["lend"],
		TOKEN_ADDRS["link"],
		TOKEN_ADDRS["mana"],
		TOKEN_ADDRS["mkr"],
		TOKEN_ADDRS["ren"],
		TOKEN_ADDRS["rep"],
		TOKEN_ADDRS["snx"],
		TOKEN_ADDRS["wbtc"],
		TOKEN_ADDRS["yfi"],
		TOKEN_ADDRS["zrx"],
	}

	CR_UNI    = common.HexToAddress("0xe89a6D0509faF730BD707bf868d9A2A744a363C7")
	CR_CREAM  = common.HexToAddress("0x892b14321a4fcba80669ae30bd0cd99a7ecf6ac0")
	CR_USDC   = common.HexToAddress("0x44fbebd2f576670a6c33f6fc0b00aa8c5753b322")
	CR_YFI    = common.HexToAddress("0xcbae0a83f4f9926997c8339545fb8ee32edc6b76")
	CR_META   = common.HexToAddress("0x3623387773010d9214b10c551d6e7fc375d31f58")
	CR_ETH    = common.HexToAddress("0xd06527d5e56a3495252a528c4987003b712860ee")
	CR_USDT   = common.HexToAddress("0x797AAB1ce7c01eB727ab980762bA88e7133d2157")
	CR_YCRV   = common.HexToAddress("0x9baF8a5236d44AC410c0186Fe39178d5AAD0Bb87")
	CR_LINK   = common.HexToAddress("0x697256CAA3cCaFD62BB6d3Aa1C7C5671786A5fD9")
	CR_CRV    = common.HexToAddress("0xc7Fd8Dcee4697ceef5a2fd4608a7BD6A94C77480")
	CR_SUSHI  = common.HexToAddress("0x338286C0BC081891A4Bda39C7667ae150bf5D206")
	CR_RENBTC = common.HexToAddress("0x17107f40d70f4470d20CB3f138a052cAE8EbD4bE")
	CR_BAL    = common.HexToAddress("0xcE4Fe9b4b8Ff61949DCfeB7e03bc9FAca59D2Eb3")
	CR_COMP   = common.HexToAddress("0x19D1666f543D42ef17F66E376944A22aEa1a8E46")
	CR_SRM    = common.HexToAddress("0xef58b2d5a1b8d3cde67b8ab054dc5c831e9bc025")
	CR_BUSD   = common.HexToAddress("0x1ff8cdb51219a8838b52e9cac09b71e591bc998e")
	CR_YETH   = common.HexToAddress("0x01da76dea59703578040012357b81ffe62015c2d")
	CR_LEND   = common.HexToAddress("0x8b86e0598616a8d4f1fdae8b59e55fb5bc33d0d6")
	CR_WNXM   = common.HexToAddress("0xeff039c3c1d668f408d09dd7b63008622a77532c")
	CR_FTX    = common.HexToAddress("0x10FDBD1e48eE2fD9336a482D746138AE19e649Db")
	CR_YUSD   = common.HexToAddress("0x4EE15f44c6F0d8d1136c83EfD2e8E4AC768954c6")
	CR_SWAG   = common.HexToAddress("0x22B243B96495C547598D9042B6f94B01C22B2e9E")

	CREAM_LIQUIDATION_PARAMS = map[common.Address]*CompLiq{
		CR_SWAG: {
			TOKEN_ADDRS["swag"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["swag"]],
		},
		CR_ETH: {
			WETH_ADDR, big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["dai"]],
		},
		CR_USDT: {
			TOKEN_ADDRS["usdt"], big.NewInt(1e6), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["usdt"]],
		},
		CR_USDC: {
			TOKEN_ADDRS["usdc"], big.NewInt(1e6), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["usdc"]],
		},
		CR_COMP: {
			TOKEN_ADDRS["comp"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["comp"]],
		},
		CR_BAL: {
			TOKEN_ADDRS["bal"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["bal"]],
		},
		CR_YFI: {
			TOKEN_ADDRS["yfi"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["yfi"]],
		},
		CR_YCRV: {
			TOKEN_ADDRS["ycrv"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["ycrv"]],
		},
		CR_LINK: {
			TOKEN_ADDRS["link"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["link"]],
		},
		CR_CREAM: {
			TOKEN_ADDRS["cream"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["cream"]],
		},
		CR_LEND: {
			TOKEN_ADDRS["lend"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["lend"]],
		},
		CR_CRV: {
			TOKEN_ADDRS["crv"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["crv"]],
		},
		CR_RENBTC: {
			TOKEN_ADDRS["renbtc"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["renbtc"]],
		},
		CR_BUSD: {
			TOKEN_ADDRS["busd"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["busd"]],
		},
		CR_META: {
			TOKEN_ADDRS["mta"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["mta"]],
		},
		CR_YUSD: {
			TOKEN_ADDRS["yusd"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["yusd"]],
		},
		CR_SUSHI: {
			TOKEN_ADDRS["sushi"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["sushi"]],
		},
		CR_FTX: {
			TOKEN_ADDRS["ftx"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["ftx"]],
		},
		CR_UNI: {
			TOKEN_ADDRS["uni"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["uni"]],
		},
		CR_YETH: {
			TOKEN_ADDRS["yweth"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["yweth"]],
		},
		CR_SRM: {
			TOKEN_ADDRS["srm"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["srm"]],
		},
		CR_WNXM: {
			TOKEN_ADDRS["wnxm"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["wnxm"]],
		},
	}

	COMP_LIQUIDATIONS_PARAMS = map[common.Address]*CompLiq{
		CSAI: {
			TOKEN_ADDRS["dai"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["dai"]],
		},
		CREP: {
			TOKEN_ADDRS["rep"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["rep"]],
		},
		CCOMP: {
			TOKEN_ADDRS["comp"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["comp"]],
		},
		CETH: {
			WETH_ADDR,
			big.NewInt(1e18),
			UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["dai"]],
		},
		cBAT: {
			TOKEN_ADDRS["bat"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["bat"]],
		},
		CDAI: {
			TOKEN_ADDRS["dai"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["dai"]],
		},
		cUNI: {
			TOKEN_ADDRS["uni"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["uni"]],
		},
		cUSDC: {
			TOKEN_ADDRS["usdc"], big.NewInt(1e6), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["usdc"]],
		},
		cUSDT: {
			TOKEN_ADDRS["usdt"], big.NewInt(1e6), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["usdt"]],
		},
		cWBTC: {
			TOKEN_ADDRS["wbtc"], big.NewInt(1e8), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["wbtc"]],
		},
		cZRX: {
			TOKEN_ADDRS["zrx"], big.NewInt(1e18), UNISWAP_USABLE_ADDRS[TOKEN_ADDRS["zrx"]],
		},
	}

	DAI  = TOKEN_ADDRS["dai"]
	USDC = TOKEN_ADDRS["usdc"]

	UNISWAP_USABLE_ADDRS = map[common.Address]common.Address{
		TOKEN_ADDRS["susd"]: common.HexToAddress(
			// actually this is ftx-lua pair so need one more swap
			"0xf80758ab42c3b07da84053fd88804bcb6baa4b5c",
		),

		TOKEN_ADDRS["ftx"]: common.HexToAddress(
			// actually this is ftx-lua pair so need one more swap
			"0xf0ec5e8ea37911dec8e8e9bc940e9dba2de60706",
		),

		TOKEN_ADDRS["swag"]: common.HexToAddress(
			"0x2059024d050cdafe4e4850596ff1490cfc40c7bd",
		),
		// HACK
		TOKEN_ADDRS["weth"]: common.HexToAddress(
			"0xa478c2975ab1ea89e8196811f51a7b7ade33eb11",
		),
		// ANOTHER HACK
		AAVE_ETH_HANDLE: common.HexToAddress(
			"0xa478c2975ab1ea89e8196811f51a7b7ade33eb11",
		),
		TOKEN_ADDRS["cream"]: common.HexToAddress(
			"0xddf9b7a31b32ebaf5c064c80900046c9e5b7c65f",
		),
		TOKEN_ADDRS["mta"]: common.HexToAddress(
			"0x0d0d65e7a7db277d3e0f5e1676325e75f3340455",
		),
		TOKEN_ADDRS["sushi"]: common.HexToAddress(
			"0xce84867c3c02b05dc570d0135103d3fb9cc19433",
		),
		TOKEN_ADDRS["renbtc"]: common.HexToAddress(
			"0x81fbef4704776cc5bba0a5df3a90056d2c6900b3",
		),
		TOKEN_ADDRS["usdt"]: common.HexToAddress(
			"0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852",
		),
		TOKEN_ADDRS["dai"]: common.HexToAddress(
			"0xa478c2975ab1ea89e8196811f51a7b7ade33eb11",
		),
		TOKEN_ADDRS["tusd"]: common.HexToAddress(
			"0xb4d0d9df2738abe81b87b66c80851292492d1404",
		),
		TOKEN_ADDRS["usdc"]: common.HexToAddress(
			"0xb4e16d0168e52d35cacd2c6185b44281ec28c9dc",
		),
		TOKEN_ADDRS["bat"]: common.HexToAddress(
			"0xb6909b960dbbe7392d405429eb2b3649752b4838",
		),
		TOKEN_ADDRS["knc"]: common.HexToAddress(
			"0xf49c43ae0faf37217bdcb00df478cf793edd6687",
		),
		TOKEN_ADDRS["link"]: common.HexToAddress(
			"0xa2107fa5b38d9bbd2c461d6edf11b11a50f6b974",
		),
		TOKEN_ADDRS["mkr"]: common.HexToAddress(
			"0xc2adda861f89bbb333c90c492cb837741916a225",
		),
		TOKEN_ADDRS["wbtc"]: common.HexToAddress(
			"0xbb2b8038a1640196fbe3e38816f3e67cba72d940",
		),
		TOKEN_ADDRS["yfi"]: common.HexToAddress(
			"0x2fdbadf3c4d5a8666bc06645b8358ab803996e28",
		),
		TOKEN_ADDRS["enj"]: common.HexToAddress(
			"0xe56c60b5f9f7b5fc70de0eb79c6ee7d00efa2625",
		),
		TOKEN_ADDRS["zrx"]: common.HexToAddress(
			"0xc6f348dd3b91a56d117ec0071c1e9b83c0996de4",
		),
		TOKEN_ADDRS["snx"]: common.HexToAddress(
			"0x43ae24960e5534731fc831386c07755a2dc33d47",
		),
		TOKEN_ADDRS["mana"]: common.HexToAddress(
			"0x11b1f53204d03e5529f09eb3091939e4fd8c9cf3",
		),
		TOKEN_ADDRS["lend"]: common.HexToAddress(
			"0xab3f9bf1d81ddb224a2014e98b238638824bcf20",
		),
		TOKEN_ADDRS["ren"]: common.HexToAddress(
			"0x8bd1661da98ebdd3bd080f0be4e6d9be8ce9858c",
		),
		TOKEN_ADDRS["uni"]: common.HexToAddress(
			"0xd3d2e2692501a5c9ca623199d38826e513033a17",
		),
		TOKEN_ADDRS["rep"]: common.HexToAddress(
			"0xec2d2240d02a8cf63c3fa0b7d2c5a3169a319496",
		),
		TOKEN_ADDRS["bal"]: common.HexToAddress(
			"0xa70d458a4d9bc0e6571565faee18a48da5c0d593",
		),
		TOKEN_ADDRS["comp"]: common.HexToAddress(
			"0xcffdded873554f362ac02f8fb1f02e5ada10516f",
		),
		TOKEN_ADDRS["srm"]: common.HexToAddress(
			"0xcc3d1ecef1f9fd25599dbea2755019dc09db3c54",
		),
		TOKEN_ADDRS["busd"]: common.HexToAddress(
			"0xc2923b8a9683556a3640ccc2961b2f52b5c4459a",
		),
		TOKEN_ADDRS["wnxm"]: common.HexToAddress(
			"0x23bff8ca20aac06efdf23cee3b8ae296a30dfd27",
		),
		TOKEN_ADDRS["crv"]: common.HexToAddress(
			"0x3da1313ae46132a397d90d95b1424a9a7e3e0fce",
		),
		TOKEN_ADDRS["yusd"]: common.HexToAddress(
			"0x9346c20186d1794101b8517177a1b15c49c9ff9b",
		),
		TOKEN_ADDRS["ycrv"]: common.HexToAddress(
			"0x55df969467ebdf954fe33470ed9c3c0f8fab0816",
		),
		TOKEN_ADDRS["yweth"]: common.HexToAddress(
			"0x29c356881d538e6ebe32e28fa3d29e7233ae6347",
		),
		TOKEN_ADDRS["aave"]: common.HexToAddress(
			"0xdfc14d2af169b0d36c4eff567ada9b2e0cae044f",
		),
	}
)

var (
	TOKEN_ADDRS = map[string]common.Address{
		"aave":    common.HexToAddress("0x7fc66500c84a76ad7e9c93437bfc5ac33e2ddae9"),
		"mana":    common.HexToAddress("0x0f5d2fb29fb7d3cfee444a200298f468908cc942"),
		"snx":     common.HexToAddress("0xc011a73ee8576fb46f5e1c5751ca3b9fe0af2a6f"),
		"zrx":     common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498"),
		"enj":     common.HexToAddress("0xf629cbd94d3791c9250152bd8dfbdf380e2a3b9c"),
		"wbtc":    common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"),
		"knc":     common.HexToAddress("0xdd974d5c2e2928dea5f71b9825b8b646686bd200"),
		"mkr":     common.HexToAddress("0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2"),
		"bat":     common.HexToAddress("0x0D8775F648430679A709E98d2b0Cb6250d2887EF"),
		"usdc":    common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
		"tusd":    common.HexToAddress("0x0000000000085d4780B73119b644AE5ecd22b376"),
		"usdt":    common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
		"yfi":     common.HexToAddress("0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e"),
		"weth":    common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		"susd":    common.HexToAddress("0x57Ab1ec28D129707052df4dF418D58a2D46d5f51"),
		"dai":     common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F"),
		"lend":    common.HexToAddress("0x80fB784B7eD66730e8b1DBd9820aFD29931aab03"),
		"ampl":    common.HexToAddress("0xD46bA6D942050d489DBd938a2C909A5d5039A161"),
		"dzar":    common.HexToAddress("0x9cb2f26a23b8d89973f08c957c4d7cdf75cd341c"),
		"srm":     common.HexToAddress("0x476c5e26a75bd202a9683ffd34359c0cc15be0ff"),
		"ren":     common.HexToAddress("0x408e41876cccdc0f92210600ef50372656052a38"),
		"yam_v2":  common.HexToAddress("0xaba8cac6866b83ae4eec97dd07ed254282f6ad8a"),
		"crv":     common.HexToAddress("0xd533a949740bb3306d119cc777fa900ba034cd52"),
		"elf":     common.HexToAddress("0xbf2179859fc6d5bee9bf9158632dc51678a4100e"),
		"sashimi": common.HexToAddress("0xc28e27870558cf22add83540d2126da2e4b464c2"),
		"link":    common.HexToAddress("0x514910771af9ca656af840dff83e8264ecf986ca"),
		"sake":    common.HexToAddress("0x066798d9ef0833ccc719076dab77199ecbd178b0"),
		"print":   common.HexToAddress("0x54b8c98268da0055971652a95f2bfd3a9349a38c"),
		"uni":     common.HexToAddress("0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"),
		"rep":     common.HexToAddress("0x1985365e9f78359a9b6ad760e32412f4a445e862"),
		"comp":    common.HexToAddress("0xc00e94Cb662C3520282E6f5717214004A7f26888"),
		"cream":   common.HexToAddress("0x2ba592f78db6436527729929aaf6c908497cb200"),
		"mta":     common.HexToAddress("0xa3bed4e1c75d00fa6f4e5e6922db7261b5e9acd2"),
		"sushi":   common.HexToAddress("0x6b3595068778dd592e39a122f4f5a5cf09c90fe2"),
		"renbtc":  common.HexToAddress("0xeb4c2781e4eba804ce9a9803c67d0893436bb27d"),
		"bal":     common.HexToAddress("0xba100000625a3754423978a60c9317c58a424e3d"),
		"busd":    common.HexToAddress("0x4fabb145d64652a948d72533023f6e7a623c7c53"),
		"wnxm":    common.HexToAddress("0x0d438f3b5175bebc262bf23753c1e53d03432bde"),
		"yusd":    common.HexToAddress("0x5dbcf33d8c2e976c6b560249878e6f1491bca25c"),
		"ycrv":    common.HexToAddress("0xdf5e0e81dff6faf3a7e52ba697820c5e32d806a8"),
		"yweth":   common.HexToAddress("0xe1237aa7f535b0cc33fd973d66cbf830354d16c7"),
		"swag":    common.HexToAddress("0x87eDfFDe3E14c7a66c9b9724747a1C5696b742e6"),
		"ftx":     common.HexToAddress("0x50d1c9771902476076ecfc8b2a83ad6b9355a4c9"),
	}

	REVERSE_NAMING = map[common.Address]string{}

	REVERSE_NAMING_COMP = map[common.Address]string{}

	// The "older" CToken contracts do not allow repaying and seizing the same asset. This applies to:
	// BAT ETH REP SAI USDC WBTC ZRX
	CANT_BE_REPAY_SAME = map[common.Address]struct{}{
		TOKEN_ADDRS["bat"]:  {},
		TOKEN_ADDRS["weth"]: {},
		TOKEN_ADDRS["rep"]:  {},
		TOKEN_ADDRS["usdc"]: {},
		TOKEN_ADDRS["wbtc"]: {},
		TOKEN_ADDRS["zrx"]:  {},
	}
)

func init() {
	for key, value := range TOKEN_ADDRS {
		REVERSE_NAMING[value] = key
	}
	for key, value := range TOKEN_ADDRS {
		REVERSE_NAMING_COMP[value] = key
	}

}
