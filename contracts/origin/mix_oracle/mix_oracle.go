// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mix_oracle

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MixOracleABI is the input ABI used to generate the binding from.
const MixOracleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxDrift\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDrift\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimGovernance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGovernor\",\"type\":\"address\"}],\"name\":\"transferGovernance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ethUsdOracles\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxDrift\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minDrift\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousGovernor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"PendingGovernorshipTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousGovernor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"GovernorshipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxDrift\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minDrift\",\"type\":\"uint256\"}],\"name\":\"setMinMaxDrift\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"registerEthUsdOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"ethOracles\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"usdOracles\",\"type\":\"address[]\"}],\"name\":\"registerTokenOracles\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"priceMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"priceMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MixOracleBin is the compiled bytecode used for deploying new contracts.
var MixOracleBin = "0x608060405234801561001057600080fd5b50604051611a04380380611a048339818101604052604081101561003357600080fd5b81019080805190602001909291908051906020019092919050505061005d336100da60201b60201c565b61006b61010960201b60201c565b73ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fc7c0c772add429241571afb3805861fb3cfa2af374534088b76cdb4325a87e9a60405160405180910390a38160358190555080603681905550505061013a565b60007f7bea13895fa79d2831e0a9e28edede30099005a50d652d8957cf8a607ee6ca4a60001b90508181555050565b6000807f7bea13895fa79d2831e0a9e28edede30099005a50d652d8957cf8a607ee6ca4a60001b9050805491505090565b6118bb806101496000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80637bf0c215116100715780637bf0c2151461020e57806398698d3c1461029b578063c7af3352146103be578063d38bfff4146103e0578063df3be62c14610424578063f92ebeac14610468576100b4565b80630c340a24146100b957806319af6bf01461010357806330d5fa4c1461019057806339295f9f146101ae57806344347397146101e65780635d36b19014610204575b600080fd5b6100c16104d6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61017a6004803603602081101561011957600080fd5b810190808035906020019064010000000081111561013657600080fd5b82018360208201111561014857600080fd5b8035906020019184600183028401116401000000008311171561016a57600080fd5b90919293919293905050506104e5565b6040518082815260200191505060405180910390f35b610198610ac4565b6040518082815260200191505060405180910390f35b6101e4600480360360408110156101c457600080fd5b810190808035906020019092919080359060200190929190505050610aca565b005b6101ee610b56565b6040518082815260200191505060405180910390f35b61020c610b5c565b005b6102856004803603602081101561022457600080fd5b810190808035906020019064010000000081111561024157600080fd5b82018360208201111561025357600080fd5b8035906020019184600183028401116401000000008311171561027557600080fd5b9091929391929390505050610bf2565b6040518082815260200191505060405180910390f35b6103bc600480360360608110156102b157600080fd5b81019080803590602001906401000000008111156102ce57600080fd5b8201836020820111156102e057600080fd5b8035906020019184600183028401116401000000008311171561030257600080fd5b90919293919293908035906020019064010000000081111561032357600080fd5b82018360208201111561033557600080fd5b8035906020019184602083028401116401000000008311171561035757600080fd5b90919293919293908035906020019064010000000081111561037857600080fd5b82018360208201111561038a57600080fd5b803590602001918460208302840111640100000000831117156103ac57600080fd5b9091929391929390505050611117565b005b6103c6611208565b604051808215151515815260200191505060405180910390f35b610422600480360360208110156103f657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611245565b005b6104666004803603602081101561043a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061132c565b005b6104946004803603602081101561047e57600080fd5b8101908080359060200190929190505050611509565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60006104e0611545565b905090565b60008060346000858560405160200180838380828437808301925050509250505060405160208183030381529060405280519060200120815260200190815260200160002090506000807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff935060008360010180549050111561081e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff915060008090505b83600101805490508110156106a8578360010181815481106105a957fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632d9ff29688886040518363ffffffff1660e01b815260040180806020018281038252848482818152602001925080828437600081840152601f19601f8201169050808301925050509350505050602060405180830381600087803b15801561065257600080fd5b505af1158015610666573d6000803e3d6000fd5b505050506040513d602081101561067c57600080fd5b810190808051906020019092919050505091508183111561069b578192505b808060010191505061058b565b508193507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff915060008090505b6033805490508110156107b557603381815481106106ef57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639478ab8c6040518163ffffffff1660e01b815260040160206040518083038186803b15801561075f57600080fd5b505afa158015610773573d6000803e3d6000fd5b505050506040513d602081101561078957600080fd5b81019080805190602001909291905050509150818311156107a8578192505b80806001019150506106d5565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff841415801561080657507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214155b1561081d57620f42408285028161081957fe5b0493505b5b6000836000018054905011156109545760008090505b836000018054905081101561095257606484600001828154811061085457fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fe2c619889896040518363ffffffff1660e01b815260040180806020018281038252848482818152602001925080828437600081840152601f19601f820116905080830192505050935050505060206040518083038186803b1580156108fb57600080fd5b505afa15801561090f573d6000803e3d6000fd5b505050506040513d602081101561092557600080fd5b810190808051906020019092919050505002915081851115610945578194505b8080600101915050610834565b505b60355484106109cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f50726963652065786365656473206d61782076616c75652e000000000000000081525060200191505060405180910390fd5b6036548411610a42576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f5072696365206c6f776572207468616e206d696e2076616c75652e000000000081525060200191505060405180910390fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff841415610abb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001806117f9602f913960400191505060405180910390fd5b50505092915050565b60355481565b610ad2611208565b610b44576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f43616c6c6572206973206e6f742074686520476f7665726e6f7200000000000081525060200191505060405180910390fd5b81603581905550806036819055505050565b60365481565b610b64611576565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610be7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806118286030913960400191505060405180910390fd5b610bf0336115a7565b565b600080603460008585604051602001808383808284378083019250505092505050604051602081830303815290604052805190602001208152602001908152602001600020905060008060009350600083600101805490501115610e90576000915060008090505b8360010180549050811015610d7757836001018181548110610c7857fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16632d9ff29688886040518363ffffffff1660e01b815260040180806020018281038252848482818152602001925080828437600081840152601f19601f8201169050808301925050509350505050602060405180830381600087803b158015610d2157600080fd5b505af1158015610d35573d6000803e3d6000fd5b505050506040513d6020811015610d4b57600080fd5b8101908080519060200190929190505050915081831015610d6a578192505b8080600101915050610c5a565b508193506000915060008090505b603380549050811015610e655760338181548110610d9f57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639478ab8c6040518163ffffffff1660e01b815260040160206040518083038186803b158015610e0f57600080fd5b505afa158015610e23573d6000803e3d6000fd5b505050506040513d6020811015610e3957600080fd5b8101908080519060200190929190505050915081831015610e58578192505b8080600101915050610d85565b5060008414158015610e78575060008214155b15610e8f57620f424082850281610e8b57fe5b0493505b5b600083600001805490501115610fc65760008090505b8360000180549050811015610fc4576064846000018281548110610ec657fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fe2c619889896040518363ffffffff1660e01b815260040180806020018281038252848482818152602001925080828437600081840152601f19601f820116905080830192505050935050505060206040518083038186803b158015610f6d57600080fd5b505afa158015610f81573d6000803e3d6000fd5b505050506040513d6020811015610f9757600080fd5b810190808051906020019092919050505002915081851015610fb7578194505b8080600101915050610ea6565b505b603554841061103d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f50726963652061626f7665206d61782076616c75652e0000000000000000000081525060200191505060405180910390fd5b60365484116110b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f50726963652062656c6f77206d696e2076616c75652e0000000000000000000081525060200191505060405180910390fd5b600084141561110e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f815260200180611858602f913960400191505060405180910390fd5b50505092915050565b61111f611208565b611191576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f43616c6c6572206973206e6f742074686520476f7665726e6f7200000000000081525060200191505060405180910390fd5b6000603460008888604051602001808383808284378083019250505092505050604051602081830303815290604052805190602001208152602001908152602001600020905084848260010191906111ea929190611715565b5082828260000191906111fe929190611715565b5050505050505050565b6000611212611545565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b61124d611208565b6112bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f43616c6c6572206973206e6f742074686520476f7665726e6f7200000000000081525060200191505060405180910390fd5b6112c8816116b7565b8073ffffffffffffffffffffffffffffffffffffffff166112e7611545565b73ffffffffffffffffffffffffffffffffffffffff167fa39cc5eb22d0f34d8beaefee8a3f17cc229c1a1d1ef87a5ad47313487b1c4f0d60405160405180910390a350565b611334611208565b6113a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f43616c6c6572206973206e6f742074686520476f7665726e6f7200000000000081525060200191505060405180910390fd5b60008090505b60338054905081101561149f578173ffffffffffffffffffffffffffffffffffffffff16603382815481106113dd57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611492576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f4f7261636c6520616c726561647920726567697374657265642e00000000000081525060200191505060405180910390fd5b80806001019150506113ac565b5060338190806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6033818154811061151657fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000807f7bea13895fa79d2831e0a9e28edede30099005a50d652d8957cf8a607ee6ca4a60001b9050805491505090565b6000807f44c4d30b2eaad5130ad70c3ba6972730566f3e6359ab83e800d905c61b1c51db60001b9050805491505090565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561164a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f4e657720476f7665726e6f72206973206164647265737328302900000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16611669611545565b73ffffffffffffffffffffffffffffffffffffffff167fc7c0c772add429241571afb3805861fb3cfa2af374534088b76cdb4325a87e9a60405160405180910390a36116b4816116e6565b50565b60007f44c4d30b2eaad5130ad70c3ba6972730566f3e6359ab83e800d905c61b1c51db60001b90508181555050565b60007f7bea13895fa79d2831e0a9e28edede30099005a50d652d8957cf8a607ee6ca4a60001b90508181555050565b8280548282559060005260206000209081019282156117a4579160200282015b828111156117a357823573ffffffffffffffffffffffffffffffffffffffff168260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190611735565b5b5090506117b191906117b5565b5090565b6117f591905b808211156117f157600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055506001016117bb565b5090565b9056fe4e6f6e65206f66206f7572206f7261636c65732072657475726e656420612076616c6964206d696e207072696365214f6e6c79207468652070656e64696e6720476f7665726e6f722063616e20636f6d706c6574652074686520636c61696d4e6f6e65206f66206f7572206f7261636c65732072657475726e656420612076616c6964206d617820707269636521a265627a7a7231582003dfe794ff19e9c3b3cc60bc92202ee9523cd7ecc4e28fb2bf97cd59ce1ef17464736f6c634300050b0032"

// DeployMixOracle deploys a new Ethereum contract, binding an instance of MixOracle to it.
func DeployMixOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _maxDrift *big.Int, _minDrift *big.Int) (common.Address, *types.Transaction, *MixOracle, error) {
	parsed, err := abi.JSON(strings.NewReader(MixOracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MixOracleBin), backend, _maxDrift, _minDrift)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MixOracle{MixOracleCaller: MixOracleCaller{contract: contract}, MixOracleTransactor: MixOracleTransactor{contract: contract}, MixOracleFilterer: MixOracleFilterer{contract: contract}}, nil
}

// MixOracle is an auto generated Go binding around an Ethereum contract.
type MixOracle struct {
	MixOracleCaller     // Read-only binding to the contract
	MixOracleTransactor // Write-only binding to the contract
	MixOracleFilterer   // Log filterer for contract events
}

// MixOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MixOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MixOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MixOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MixOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MixOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MixOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MixOracleSession struct {
	Contract     *MixOracle        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MixOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MixOracleCallerSession struct {
	Contract *MixOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MixOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MixOracleTransactorSession struct {
	Contract     *MixOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MixOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MixOracleRaw struct {
	Contract *MixOracle // Generic contract binding to access the raw methods on
}

// MixOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MixOracleCallerRaw struct {
	Contract *MixOracleCaller // Generic read-only contract binding to access the raw methods on
}

// MixOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MixOracleTransactorRaw struct {
	Contract *MixOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMixOracle creates a new instance of MixOracle, bound to a specific deployed contract.
func NewMixOracle(address common.Address, backend bind.ContractBackend) (*MixOracle, error) {
	contract, err := bindMixOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MixOracle{MixOracleCaller: MixOracleCaller{contract: contract}, MixOracleTransactor: MixOracleTransactor{contract: contract}, MixOracleFilterer: MixOracleFilterer{contract: contract}}, nil
}

// NewMixOracleCaller creates a new read-only instance of MixOracle, bound to a specific deployed contract.
func NewMixOracleCaller(address common.Address, caller bind.ContractCaller) (*MixOracleCaller, error) {
	contract, err := bindMixOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MixOracleCaller{contract: contract}, nil
}

// NewMixOracleTransactor creates a new write-only instance of MixOracle, bound to a specific deployed contract.
func NewMixOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*MixOracleTransactor, error) {
	contract, err := bindMixOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MixOracleTransactor{contract: contract}, nil
}

// NewMixOracleFilterer creates a new log filterer instance of MixOracle, bound to a specific deployed contract.
func NewMixOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*MixOracleFilterer, error) {
	contract, err := bindMixOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MixOracleFilterer{contract: contract}, nil
}

// bindMixOracle binds a generic wrapper to an already deployed contract.
func bindMixOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MixOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MixOracle *MixOracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MixOracle.Contract.MixOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MixOracle *MixOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MixOracle.Contract.MixOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MixOracle *MixOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MixOracle.Contract.MixOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MixOracle *MixOracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MixOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MixOracle *MixOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MixOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MixOracle *MixOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MixOracle.Contract.contract.Transact(opts, method, params...)
}

// EthUsdOracles is a free data retrieval call binding the contract method 0xf92ebeac.
//
// Solidity: function ethUsdOracles(uint256 ) view returns(address)
func (_MixOracle *MixOracleCaller) EthUsdOracles(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MixOracle.contract.Call(opts, out, "ethUsdOracles", arg0)
	return *ret0, err
}

// EthUsdOracles is a free data retrieval call binding the contract method 0xf92ebeac.
//
// Solidity: function ethUsdOracles(uint256 ) view returns(address)
func (_MixOracle *MixOracleSession) EthUsdOracles(arg0 *big.Int) (common.Address, error) {
	return _MixOracle.Contract.EthUsdOracles(&_MixOracle.CallOpts, arg0)
}

// EthUsdOracles is a free data retrieval call binding the contract method 0xf92ebeac.
//
// Solidity: function ethUsdOracles(uint256 ) view returns(address)
func (_MixOracle *MixOracleCallerSession) EthUsdOracles(arg0 *big.Int) (common.Address, error) {
	return _MixOracle.Contract.EthUsdOracles(&_MixOracle.CallOpts, arg0)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_MixOracle *MixOracleCaller) Governor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MixOracle.contract.Call(opts, out, "governor")
	return *ret0, err
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_MixOracle *MixOracleSession) Governor() (common.Address, error) {
	return _MixOracle.Contract.Governor(&_MixOracle.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_MixOracle *MixOracleCallerSession) Governor() (common.Address, error) {
	return _MixOracle.Contract.Governor(&_MixOracle.CallOpts)
}

// IsGovernor is a free data retrieval call binding the contract method 0xc7af3352.
//
// Solidity: function isGovernor() view returns(bool)
func (_MixOracle *MixOracleCaller) IsGovernor(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MixOracle.contract.Call(opts, out, "isGovernor")
	return *ret0, err
}

// IsGovernor is a free data retrieval call binding the contract method 0xc7af3352.
//
// Solidity: function isGovernor() view returns(bool)
func (_MixOracle *MixOracleSession) IsGovernor() (bool, error) {
	return _MixOracle.Contract.IsGovernor(&_MixOracle.CallOpts)
}

// IsGovernor is a free data retrieval call binding the contract method 0xc7af3352.
//
// Solidity: function isGovernor() view returns(bool)
func (_MixOracle *MixOracleCallerSession) IsGovernor() (bool, error) {
	return _MixOracle.Contract.IsGovernor(&_MixOracle.CallOpts)
}

// MaxDrift is a free data retrieval call binding the contract method 0x30d5fa4c.
//
// Solidity: function maxDrift() view returns(uint256)
func (_MixOracle *MixOracleCaller) MaxDrift(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MixOracle.contract.Call(opts, out, "maxDrift")
	return *ret0, err
}

// MaxDrift is a free data retrieval call binding the contract method 0x30d5fa4c.
//
// Solidity: function maxDrift() view returns(uint256)
func (_MixOracle *MixOracleSession) MaxDrift() (*big.Int, error) {
	return _MixOracle.Contract.MaxDrift(&_MixOracle.CallOpts)
}

// MaxDrift is a free data retrieval call binding the contract method 0x30d5fa4c.
//
// Solidity: function maxDrift() view returns(uint256)
func (_MixOracle *MixOracleCallerSession) MaxDrift() (*big.Int, error) {
	return _MixOracle.Contract.MaxDrift(&_MixOracle.CallOpts)
}

// MinDrift is a free data retrieval call binding the contract method 0x44347397.
//
// Solidity: function minDrift() view returns(uint256)
func (_MixOracle *MixOracleCaller) MinDrift(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MixOracle.contract.Call(opts, out, "minDrift")
	return *ret0, err
}

// MinDrift is a free data retrieval call binding the contract method 0x44347397.
//
// Solidity: function minDrift() view returns(uint256)
func (_MixOracle *MixOracleSession) MinDrift() (*big.Int, error) {
	return _MixOracle.Contract.MinDrift(&_MixOracle.CallOpts)
}

// MinDrift is a free data retrieval call binding the contract method 0x44347397.
//
// Solidity: function minDrift() view returns(uint256)
func (_MixOracle *MixOracleCallerSession) MinDrift() (*big.Int, error) {
	return _MixOracle.Contract.MinDrift(&_MixOracle.CallOpts)
}

// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
//
// Solidity: function claimGovernance() returns()
func (_MixOracle *MixOracleTransactor) ClaimGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MixOracle.contract.Transact(opts, "claimGovernance")
}

// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
//
// Solidity: function claimGovernance() returns()
func (_MixOracle *MixOracleSession) ClaimGovernance() (*types.Transaction, error) {
	return _MixOracle.Contract.ClaimGovernance(&_MixOracle.TransactOpts)
}

// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
//
// Solidity: function claimGovernance() returns()
func (_MixOracle *MixOracleTransactorSession) ClaimGovernance() (*types.Transaction, error) {
	return _MixOracle.Contract.ClaimGovernance(&_MixOracle.TransactOpts)
}

// PriceMax is a paid mutator transaction binding the contract method 0x7bf0c215.
//
// Solidity: function priceMax(string symbol) returns(uint256 price)
func (_MixOracle *MixOracleTransactor) PriceMax(opts *bind.TransactOpts, symbol string) (*types.Transaction, error) {
	return _MixOracle.contract.Transact(opts, "priceMax", symbol)
}

// PriceMax is a paid mutator transaction binding the contract method 0x7bf0c215.
//
// Solidity: function priceMax(string symbol) returns(uint256 price)
func (_MixOracle *MixOracleSession) PriceMax(symbol string) (*types.Transaction, error) {
	return _MixOracle.Contract.PriceMax(&_MixOracle.TransactOpts, symbol)
}

// PriceMax is a paid mutator transaction binding the contract method 0x7bf0c215.
//
// Solidity: function priceMax(string symbol) returns(uint256 price)
func (_MixOracle *MixOracleTransactorSession) PriceMax(symbol string) (*types.Transaction, error) {
	return _MixOracle.Contract.PriceMax(&_MixOracle.TransactOpts, symbol)
}

// PriceMin is a paid mutator transaction binding the contract method 0x19af6bf0.
//
// Solidity: function priceMin(string symbol) returns(uint256 price)
func (_MixOracle *MixOracleTransactor) PriceMin(opts *bind.TransactOpts, symbol string) (*types.Transaction, error) {
	return _MixOracle.contract.Transact(opts, "priceMin", symbol)
}

// PriceMin is a paid mutator transaction binding the contract method 0x19af6bf0.
//
// Solidity: function priceMin(string symbol) returns(uint256 price)
func (_MixOracle *MixOracleSession) PriceMin(symbol string) (*types.Transaction, error) {
	return _MixOracle.Contract.PriceMin(&_MixOracle.TransactOpts, symbol)
}

// PriceMin is a paid mutator transaction binding the contract method 0x19af6bf0.
//
// Solidity: function priceMin(string symbol) returns(uint256 price)
func (_MixOracle *MixOracleTransactorSession) PriceMin(symbol string) (*types.Transaction, error) {
	return _MixOracle.Contract.PriceMin(&_MixOracle.TransactOpts, symbol)
}

// RegisterEthUsdOracle is a paid mutator transaction binding the contract method 0xdf3be62c.
//
// Solidity: function registerEthUsdOracle(address oracle) returns()
func (_MixOracle *MixOracleTransactor) RegisterEthUsdOracle(opts *bind.TransactOpts, oracle common.Address) (*types.Transaction, error) {
	return _MixOracle.contract.Transact(opts, "registerEthUsdOracle", oracle)
}

// RegisterEthUsdOracle is a paid mutator transaction binding the contract method 0xdf3be62c.
//
// Solidity: function registerEthUsdOracle(address oracle) returns()
func (_MixOracle *MixOracleSession) RegisterEthUsdOracle(oracle common.Address) (*types.Transaction, error) {
	return _MixOracle.Contract.RegisterEthUsdOracle(&_MixOracle.TransactOpts, oracle)
}

// RegisterEthUsdOracle is a paid mutator transaction binding the contract method 0xdf3be62c.
//
// Solidity: function registerEthUsdOracle(address oracle) returns()
func (_MixOracle *MixOracleTransactorSession) RegisterEthUsdOracle(oracle common.Address) (*types.Transaction, error) {
	return _MixOracle.Contract.RegisterEthUsdOracle(&_MixOracle.TransactOpts, oracle)
}

// RegisterTokenOracles is a paid mutator transaction binding the contract method 0x98698d3c.
//
// Solidity: function registerTokenOracles(string symbol, address[] ethOracles, address[] usdOracles) returns()
func (_MixOracle *MixOracleTransactor) RegisterTokenOracles(opts *bind.TransactOpts, symbol string, ethOracles []common.Address, usdOracles []common.Address) (*types.Transaction, error) {
	return _MixOracle.contract.Transact(opts, "registerTokenOracles", symbol, ethOracles, usdOracles)
}

// RegisterTokenOracles is a paid mutator transaction binding the contract method 0x98698d3c.
//
// Solidity: function registerTokenOracles(string symbol, address[] ethOracles, address[] usdOracles) returns()
func (_MixOracle *MixOracleSession) RegisterTokenOracles(symbol string, ethOracles []common.Address, usdOracles []common.Address) (*types.Transaction, error) {
	return _MixOracle.Contract.RegisterTokenOracles(&_MixOracle.TransactOpts, symbol, ethOracles, usdOracles)
}

// RegisterTokenOracles is a paid mutator transaction binding the contract method 0x98698d3c.
//
// Solidity: function registerTokenOracles(string symbol, address[] ethOracles, address[] usdOracles) returns()
func (_MixOracle *MixOracleTransactorSession) RegisterTokenOracles(symbol string, ethOracles []common.Address, usdOracles []common.Address) (*types.Transaction, error) {
	return _MixOracle.Contract.RegisterTokenOracles(&_MixOracle.TransactOpts, symbol, ethOracles, usdOracles)
}

// SetMinMaxDrift is a paid mutator transaction binding the contract method 0x39295f9f.
//
// Solidity: function setMinMaxDrift(uint256 _maxDrift, uint256 _minDrift) returns()
func (_MixOracle *MixOracleTransactor) SetMinMaxDrift(opts *bind.TransactOpts, _maxDrift *big.Int, _minDrift *big.Int) (*types.Transaction, error) {
	return _MixOracle.contract.Transact(opts, "setMinMaxDrift", _maxDrift, _minDrift)
}

// SetMinMaxDrift is a paid mutator transaction binding the contract method 0x39295f9f.
//
// Solidity: function setMinMaxDrift(uint256 _maxDrift, uint256 _minDrift) returns()
func (_MixOracle *MixOracleSession) SetMinMaxDrift(_maxDrift *big.Int, _minDrift *big.Int) (*types.Transaction, error) {
	return _MixOracle.Contract.SetMinMaxDrift(&_MixOracle.TransactOpts, _maxDrift, _minDrift)
}

// SetMinMaxDrift is a paid mutator transaction binding the contract method 0x39295f9f.
//
// Solidity: function setMinMaxDrift(uint256 _maxDrift, uint256 _minDrift) returns()
func (_MixOracle *MixOracleTransactorSession) SetMinMaxDrift(_maxDrift *big.Int, _minDrift *big.Int) (*types.Transaction, error) {
	return _MixOracle.Contract.SetMinMaxDrift(&_MixOracle.TransactOpts, _maxDrift, _minDrift)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _newGovernor) returns()
func (_MixOracle *MixOracleTransactor) TransferGovernance(opts *bind.TransactOpts, _newGovernor common.Address) (*types.Transaction, error) {
	return _MixOracle.contract.Transact(opts, "transferGovernance", _newGovernor)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _newGovernor) returns()
func (_MixOracle *MixOracleSession) TransferGovernance(_newGovernor common.Address) (*types.Transaction, error) {
	return _MixOracle.Contract.TransferGovernance(&_MixOracle.TransactOpts, _newGovernor)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _newGovernor) returns()
func (_MixOracle *MixOracleTransactorSession) TransferGovernance(_newGovernor common.Address) (*types.Transaction, error) {
	return _MixOracle.Contract.TransferGovernance(&_MixOracle.TransactOpts, _newGovernor)
}

// MixOracleGovernorshipTransferredIterator is returned from FilterGovernorshipTransferred and is used to iterate over the raw logs and unpacked data for GovernorshipTransferred events raised by the MixOracle contract.
type MixOracleGovernorshipTransferredIterator struct {
	Event *MixOracleGovernorshipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MixOracleGovernorshipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MixOracleGovernorshipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MixOracleGovernorshipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MixOracleGovernorshipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MixOracleGovernorshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MixOracleGovernorshipTransferred represents a GovernorshipTransferred event raised by the MixOracle contract.
type MixOracleGovernorshipTransferred struct {
	PreviousGovernor common.Address
	NewGovernor      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterGovernorshipTransferred is a free log retrieval operation binding the contract event 0xc7c0c772add429241571afb3805861fb3cfa2af374534088b76cdb4325a87e9a.
//
// Solidity: event GovernorshipTransferred(address indexed previousGovernor, address indexed newGovernor)
func (_MixOracle *MixOracleFilterer) FilterGovernorshipTransferred(opts *bind.FilterOpts, previousGovernor []common.Address, newGovernor []common.Address) (*MixOracleGovernorshipTransferredIterator, error) {

	var previousGovernorRule []interface{}
	for _, previousGovernorItem := range previousGovernor {
		previousGovernorRule = append(previousGovernorRule, previousGovernorItem)
	}
	var newGovernorRule []interface{}
	for _, newGovernorItem := range newGovernor {
		newGovernorRule = append(newGovernorRule, newGovernorItem)
	}

	logs, sub, err := _MixOracle.contract.FilterLogs(opts, "GovernorshipTransferred", previousGovernorRule, newGovernorRule)
	if err != nil {
		return nil, err
	}
	return &MixOracleGovernorshipTransferredIterator{contract: _MixOracle.contract, event: "GovernorshipTransferred", logs: logs, sub: sub}, nil
}

// WatchGovernorshipTransferred is a free log subscription operation binding the contract event 0xc7c0c772add429241571afb3805861fb3cfa2af374534088b76cdb4325a87e9a.
//
// Solidity: event GovernorshipTransferred(address indexed previousGovernor, address indexed newGovernor)
func (_MixOracle *MixOracleFilterer) WatchGovernorshipTransferred(opts *bind.WatchOpts, sink chan<- *MixOracleGovernorshipTransferred, previousGovernor []common.Address, newGovernor []common.Address) (event.Subscription, error) {

	var previousGovernorRule []interface{}
	for _, previousGovernorItem := range previousGovernor {
		previousGovernorRule = append(previousGovernorRule, previousGovernorItem)
	}
	var newGovernorRule []interface{}
	for _, newGovernorItem := range newGovernor {
		newGovernorRule = append(newGovernorRule, newGovernorItem)
	}

	logs, sub, err := _MixOracle.contract.WatchLogs(opts, "GovernorshipTransferred", previousGovernorRule, newGovernorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MixOracleGovernorshipTransferred)
				if err := _MixOracle.contract.UnpackLog(event, "GovernorshipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernorshipTransferred is a log parse operation binding the contract event 0xc7c0c772add429241571afb3805861fb3cfa2af374534088b76cdb4325a87e9a.
//
// Solidity: event GovernorshipTransferred(address indexed previousGovernor, address indexed newGovernor)
func (_MixOracle *MixOracleFilterer) ParseGovernorshipTransferred(log types.Log) (*MixOracleGovernorshipTransferred, error) {
	event := new(MixOracleGovernorshipTransferred)
	if err := _MixOracle.contract.UnpackLog(event, "GovernorshipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MixOraclePendingGovernorshipTransferIterator is returned from FilterPendingGovernorshipTransfer and is used to iterate over the raw logs and unpacked data for PendingGovernorshipTransfer events raised by the MixOracle contract.
type MixOraclePendingGovernorshipTransferIterator struct {
	Event *MixOraclePendingGovernorshipTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MixOraclePendingGovernorshipTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MixOraclePendingGovernorshipTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MixOraclePendingGovernorshipTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MixOraclePendingGovernorshipTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MixOraclePendingGovernorshipTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MixOraclePendingGovernorshipTransfer represents a PendingGovernorshipTransfer event raised by the MixOracle contract.
type MixOraclePendingGovernorshipTransfer struct {
	PreviousGovernor common.Address
	NewGovernor      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPendingGovernorshipTransfer is a free log retrieval operation binding the contract event 0xa39cc5eb22d0f34d8beaefee8a3f17cc229c1a1d1ef87a5ad47313487b1c4f0d.
//
// Solidity: event PendingGovernorshipTransfer(address indexed previousGovernor, address indexed newGovernor)
func (_MixOracle *MixOracleFilterer) FilterPendingGovernorshipTransfer(opts *bind.FilterOpts, previousGovernor []common.Address, newGovernor []common.Address) (*MixOraclePendingGovernorshipTransferIterator, error) {

	var previousGovernorRule []interface{}
	for _, previousGovernorItem := range previousGovernor {
		previousGovernorRule = append(previousGovernorRule, previousGovernorItem)
	}
	var newGovernorRule []interface{}
	for _, newGovernorItem := range newGovernor {
		newGovernorRule = append(newGovernorRule, newGovernorItem)
	}

	logs, sub, err := _MixOracle.contract.FilterLogs(opts, "PendingGovernorshipTransfer", previousGovernorRule, newGovernorRule)
	if err != nil {
		return nil, err
	}
	return &MixOraclePendingGovernorshipTransferIterator{contract: _MixOracle.contract, event: "PendingGovernorshipTransfer", logs: logs, sub: sub}, nil
}

// WatchPendingGovernorshipTransfer is a free log subscription operation binding the contract event 0xa39cc5eb22d0f34d8beaefee8a3f17cc229c1a1d1ef87a5ad47313487b1c4f0d.
//
// Solidity: event PendingGovernorshipTransfer(address indexed previousGovernor, address indexed newGovernor)
func (_MixOracle *MixOracleFilterer) WatchPendingGovernorshipTransfer(opts *bind.WatchOpts, sink chan<- *MixOraclePendingGovernorshipTransfer, previousGovernor []common.Address, newGovernor []common.Address) (event.Subscription, error) {

	var previousGovernorRule []interface{}
	for _, previousGovernorItem := range previousGovernor {
		previousGovernorRule = append(previousGovernorRule, previousGovernorItem)
	}
	var newGovernorRule []interface{}
	for _, newGovernorItem := range newGovernor {
		newGovernorRule = append(newGovernorRule, newGovernorItem)
	}

	logs, sub, err := _MixOracle.contract.WatchLogs(opts, "PendingGovernorshipTransfer", previousGovernorRule, newGovernorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MixOraclePendingGovernorshipTransfer)
				if err := _MixOracle.contract.UnpackLog(event, "PendingGovernorshipTransfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePendingGovernorshipTransfer is a log parse operation binding the contract event 0xa39cc5eb22d0f34d8beaefee8a3f17cc229c1a1d1ef87a5ad47313487b1c4f0d.
//
// Solidity: event PendingGovernorshipTransfer(address indexed previousGovernor, address indexed newGovernor)
func (_MixOracle *MixOracleFilterer) ParsePendingGovernorshipTransfer(log types.Log) (*MixOraclePendingGovernorshipTransfer, error) {
	event := new(MixOraclePendingGovernorshipTransfer)
	if err := _MixOracle.contract.UnpackLog(event, "PendingGovernorshipTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
