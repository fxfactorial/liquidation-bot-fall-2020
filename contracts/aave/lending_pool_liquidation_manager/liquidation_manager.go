// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lending_pool_liquidation_manager

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

// LendingPoolLiquidationManagerABI is the input ABI used to generate the binding from.
const LendingPoolLiquidationManagerABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"addressesProvider\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_collateral\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_feeLiquidated\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_liquidatedCollateralForFee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"OriginationFeeLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_collateral\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_purchaseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_liquidatedCollateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_accruedBorrowInterest\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_liquidator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_receiveAToken\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"LiquidationCall\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_collateral\",\"type\":\"address\"},{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_purchaseAmount\",\"type\":\"uint256\"},{\"name\":\"_receiveAToken\",\"type\":\"bool\"}],\"name\":\"liquidationCall\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// LendingPoolLiquidationManagerBin is the compiled bytecode used for deploying new contracts.
var LendingPoolLiquidationManagerBin = "0x60806040526000600181815590556114bf8061001c6000396000f3fe6080604052600436106100285760003560e01c8062a718a91461002d578063c72c4d10146100f0575b600080fd5b610071600480360360a081101561004357600080fd5b506001600160a01b038135811691602081013582169160408201351690606081013590608001351515610121565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156100b457818101518382015260200161009c565b50505050905090810190601f1680156100e15780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b3480156100fc57600080fd5b50610105610de5565b604080516001600160a01b039092168252519081900360200190f35b6000606061012d61130c565b603754604080517f2c6d0e9b0000000000000000000000000000000000000000000000000000000081526001600160a01b03898116600483015291519190921691632c6d0e9b91602480830192610100929190829003018186803b15801561019457600080fd5b505afa1580156101a8573d6000803e3d6000fd5b505050506040513d6101008110156101bf57600080fd5b5060e0015115156101e082018190526101f7576004604051806060016040528060288152602001611442602891399250925050610ddb565b603654604080517f18a4dbca0000000000000000000000000000000000000000000000000000000081526001600160a01b038b811660048301528981166024830152915191909216916318a4dbca916044808301926020929190829003018186803b15801561026557600080fd5b505afa158015610279573d6000803e3d6000fd5b505050506040513d602081101561028f57600080fd5b505180825215156102d957505060408051808201909152601f81527f496e76616c696420636f6c6c61746572616c20746f206c697175696461746500602082015260019150610ddb565b603654604080517f18f9bbae0000000000000000000000000000000000000000000000000000000081526001600160a01b038b81166004830152915191909216916318f9bbae916024808301926020929190829003018186803b15801561033f57600080fd5b505afa158015610353573d6000803e3d6000fd5b505050506040513d602081101561036957600080fd5b5051801561040d5750603654604080517f9e3c4f3b0000000000000000000000000000000000000000000000000000000081526001600160a01b038b81166004830152898116602483015291519190921691639e3c4f3b916044808301926020929190829003018186803b1580156103e057600080fd5b505afa1580156103f4573d6000803e3d6000fd5b505050506040513d602081101561040a57600080fd5b50515b15156101c082018190526104405760026040518060600160405280602a8152602001611418602a91399250925050610ddb565b603654604080517f9fb8afcd0000000000000000000000000000000000000000000000000000000081526001600160a01b038a81166004830152898116602483015291519190921691639fb8afcd916044808301926060929190829003018186803b1580156104ae57600080fd5b505afa1580156104c2573d6000803e3d6000fd5b505050506040513d60608110156104d857600080fd5b50602080820151604092830151928401929092528201819052151561051c5760036040518060600160405280602a815260200161146a602a91399250925050610ddb565b610545606461053960328460200151610df490919063ffffffff16565b9063ffffffff610e7416565b606082018190528511610558578461055e565b80606001515b608082018190528151600091829161057a918c918c9190610eb6565b603654604080517ffeab31ac0000000000000000000000000000000000000000000000000000000081526001600160a01b038e811660048301528d811660248301529151949650929450169163feab31ac91604480820192602092909190829003018186803b1580156105ec57600080fd5b505afa158015610600573d6000803e3d6000fd5b505050506040513d602081101561061657600080fd5b50516101208401819052156106595761064c8a8a85610120015161064786886000015161119590919063ffffffff16565b610eb6565b6101408501526101608401525b826080015181101561066d57608083018190525b85151561073e57603654604080517fe24030190000000000000000000000000000000000000000000000000000000081526001600160a01b038d811660048301529151600093929092169163e240301991602480820192602092909190829003018186803b1580156106de57600080fd5b505afa1580156106f2573d6000803e3d6000fd5b505050506040513d602081101561070857600080fd5b505190508281101561073c5760056040518060600160405280603381526020016113e5603391399550955050505050610ddb565b505b603660009054906101000a90046001600160a01b03166001600160a01b03166368beb4d68a8c8b8760800151878961014001518a61016001518b604001518f6040518a63ffffffff1660e01b8152600401808a6001600160a01b03166001600160a01b03168152602001896001600160a01b03166001600160a01b03168152602001886001600160a01b03166001600160a01b03168152602001878152602001868152602001858152602001848152602001838152602001821515151581526020019950505050505050505050600060405180830381600087803b15801561082557600080fd5b505af1158015610839573d6000803e3d6000fd5b5050603654604080517f34b3beee0000000000000000000000000000000000000000000000000000000081526001600160a01b038f811660048301529151600095509190921692506334b3beee91602480820192602092909190829003018186803b1580156108a757600080fd5b505afa1580156108bb573d6000803e3d6000fd5b505050506040513d60208110156108d157600080fd5b50519050861561096957604080517ff866c3190000000000000000000000000000000000000000000000000000000081526001600160a01b038b811660048301523360248301526044820186905291519183169163f866c3199160648082019260009290919082900301818387803b15801561094c57600080fd5b505af1158015610960573d6000803e3d6000fd5b50505050610a89565b806001600160a01b0316633edb7cb88a856040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b1580156109c957600080fd5b505af11580156109dd573d6000803e3d6000fd5b50505050603660009054906101000a90046001600160a01b03166001600160a01b031663fa93b2a58c33866040518463ffffffff1660e01b815260040180846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b031681526020018281526020019350505050600060405180830381600087803b158015610a7057600080fd5b505af1158015610a84573d6000803e3d6000fd5b505050505b6036546080850151604080517f28fcf4d30000000000000000000000000000000000000000000000000000000081526001600160a01b038e811660048301523360248301526044820193909352905191909216916328fcf4d391349160648082019260009290919082900301818588803b158015610b0657600080fd5b505af1158015610b1a573d6000803e3d6000fd5b505050505060008461014001511115610d2857806001600160a01b0316633edb7cb88a8661016001516040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050600060405180830381600087803b158015610b9257600080fd5b505af1158015610ba6573d6000803e3d6000fd5b50505050603660009054906101000a90046001600160a01b03166001600160a01b0316638f385c228c866101600151603560009054906101000a90046001600160a01b03166001600160a01b031663ee8912966040518163ffffffff1660e01b815260040160206040518083038186803b158015610c2357600080fd5b505afa158015610c37573d6000803e3d6000fd5b505050506040513d6020811015610c4d57600080fd5b50516040805163ffffffff861660e01b81526001600160a01b0394851660048201526024810193909352921660448201529051606480830192600092919082900301818387803b158015610ca057600080fd5b505af1158015610cb4573d6000803e3d6000fd5b50505050886001600160a01b03168a6001600160a01b03168c6001600160a01b03167f36ca8b16d61dc13b1062adff83e3778ab92d14f9e35bfe9fd1283e02b13fb0a18761014001518861016001514260405180848152602001838152602001828152602001935050505060405180910390a45b60808085015160408087015181519283526020830187905282820152336060830152891515928201929092524260a082015290516001600160a01b03808c16928d821692918f16917f56864757fd5b1fc9f38f5f3a981cd8ae512ce41b902cf73fc506ee369c6bc2379181900360c00190a460006040518060400160405280600981526020017f4e6f206572726f7273000000000000000000000000000000000000000000000081525095509550505050505b9550959350505050565b6035546001600160a01b031681565b6000821515610e0557506000610e6e565b828202828482811515610e1457fe5b0414610e6b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806113c46021913960400191505060405180910390fd5b90505b92915050565b6000610e6b83836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f0000000000008152506111d7565b603554604080517ffca513a80000000000000000000000000000000000000000000000000000000081529051600092839283926001600160a01b039092169163fca513a891600480820192602092909190829003018186803b158015610f1b57600080fd5b505afa158015610f2f573d6000803e3d6000fd5b505050506040513d6020811015610f4557600080fd5b50519050610f51611394565b816001600160a01b031663b3596f07896040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015610fa757600080fd5b505afa158015610fbb573d6000803e3d6000fd5b505050506040513d6020811015610fd157600080fd5b505160408083019190915280517fb3596f070000000000000000000000000000000000000000000000000000000081526001600160a01b03898116600483015291519184169163b3596f0791602480820192602092909190829003018186803b15801561103d57600080fd5b505afa158015611051573d6000803e3d6000fd5b505050506040513d602081101561106757600080fd5b50516060820152603654604080517fc76a6c9c0000000000000000000000000000000000000000000000000000000081526001600160a01b038b811660048301529151919092169163c76a6c9c916024808301926020929190829003018186803b1580156110d457600080fd5b505afa1580156110e8573d6000803e3d6000fd5b505050506040513d60208110156110fe57600080fd5b5051602082018190526040820151606083015161113b926064926105399261112f919084908d63ffffffff610df416565b9063ffffffff610df416565b6080820181905285101561117f578493506111788160200151610539606461112f85606001516105398a8860400151610df490919063ffffffff16565b925061118a565b806080015193508592505b505094509492505050565b6000610e6b83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611297565b60008183151561127f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561124457818101518382015260200161122c565b50505050905090810190601f1680156112715780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000838581151561128d57fe5b0495945050505050565b60008184841115611304576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360008381101561124457818101518382015260200161122c565b505050900390565b6040518061020001604052806000815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000600281111561137957fe5b81526000602082018190526040820181905260609091015290565b6040518060a001604052806000815260200160008152602001600081526020016000815260200160008152509056fe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f7754686572652069736e277420656e6f756768206c697175696469747920617661696c61626c6520746f206c697175696461746554686520636f6c6c61746572616c2063686f73656e2063616e6e6f74206265206c6971756964617465644865616c746820666163746f72206973206e6f742062656c6f7720746865207468726573686f6c645573657220646964206e6f7420626f72726f7720746865207370656369666965642063757272656e6379a165627a7a72305820e94cfeb88b1c49abce68668c02045853bc5bda8ffbdaf289fec57eaa796351fa0029"

// DeployLendingPoolLiquidationManager deploys a new Ethereum contract, binding an instance of LendingPoolLiquidationManager to it.
func DeployLendingPoolLiquidationManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LendingPoolLiquidationManager, error) {
	parsed, err := abi.JSON(strings.NewReader(LendingPoolLiquidationManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LendingPoolLiquidationManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LendingPoolLiquidationManager{LendingPoolLiquidationManagerCaller: LendingPoolLiquidationManagerCaller{contract: contract}, LendingPoolLiquidationManagerTransactor: LendingPoolLiquidationManagerTransactor{contract: contract}, LendingPoolLiquidationManagerFilterer: LendingPoolLiquidationManagerFilterer{contract: contract}}, nil
}

// LendingPoolLiquidationManager is an auto generated Go binding around an Ethereum contract.
type LendingPoolLiquidationManager struct {
	LendingPoolLiquidationManagerCaller     // Read-only binding to the contract
	LendingPoolLiquidationManagerTransactor // Write-only binding to the contract
	LendingPoolLiquidationManagerFilterer   // Log filterer for contract events
}

// LendingPoolLiquidationManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LendingPoolLiquidationManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolLiquidationManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LendingPoolLiquidationManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolLiquidationManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LendingPoolLiquidationManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolLiquidationManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LendingPoolLiquidationManagerSession struct {
	Contract     *LendingPoolLiquidationManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// LendingPoolLiquidationManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LendingPoolLiquidationManagerCallerSession struct {
	Contract *LendingPoolLiquidationManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// LendingPoolLiquidationManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LendingPoolLiquidationManagerTransactorSession struct {
	Contract     *LendingPoolLiquidationManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// LendingPoolLiquidationManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LendingPoolLiquidationManagerRaw struct {
	Contract *LendingPoolLiquidationManager // Generic contract binding to access the raw methods on
}

// LendingPoolLiquidationManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LendingPoolLiquidationManagerCallerRaw struct {
	Contract *LendingPoolLiquidationManagerCaller // Generic read-only contract binding to access the raw methods on
}

// LendingPoolLiquidationManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LendingPoolLiquidationManagerTransactorRaw struct {
	Contract *LendingPoolLiquidationManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLendingPoolLiquidationManager creates a new instance of LendingPoolLiquidationManager, bound to a specific deployed contract.
func NewLendingPoolLiquidationManager(address common.Address, backend bind.ContractBackend) (*LendingPoolLiquidationManager, error) {
	contract, err := bindLendingPoolLiquidationManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LendingPoolLiquidationManager{LendingPoolLiquidationManagerCaller: LendingPoolLiquidationManagerCaller{contract: contract}, LendingPoolLiquidationManagerTransactor: LendingPoolLiquidationManagerTransactor{contract: contract}, LendingPoolLiquidationManagerFilterer: LendingPoolLiquidationManagerFilterer{contract: contract}}, nil
}

// NewLendingPoolLiquidationManagerCaller creates a new read-only instance of LendingPoolLiquidationManager, bound to a specific deployed contract.
func NewLendingPoolLiquidationManagerCaller(address common.Address, caller bind.ContractCaller) (*LendingPoolLiquidationManagerCaller, error) {
	contract, err := bindLendingPoolLiquidationManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LendingPoolLiquidationManagerCaller{contract: contract}, nil
}

// NewLendingPoolLiquidationManagerTransactor creates a new write-only instance of LendingPoolLiquidationManager, bound to a specific deployed contract.
func NewLendingPoolLiquidationManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*LendingPoolLiquidationManagerTransactor, error) {
	contract, err := bindLendingPoolLiquidationManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LendingPoolLiquidationManagerTransactor{contract: contract}, nil
}

// NewLendingPoolLiquidationManagerFilterer creates a new log filterer instance of LendingPoolLiquidationManager, bound to a specific deployed contract.
func NewLendingPoolLiquidationManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*LendingPoolLiquidationManagerFilterer, error) {
	contract, err := bindLendingPoolLiquidationManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LendingPoolLiquidationManagerFilterer{contract: contract}, nil
}

// bindLendingPoolLiquidationManager binds a generic wrapper to an already deployed contract.
func bindLendingPoolLiquidationManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LendingPoolLiquidationManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingPoolLiquidationManager *LendingPoolLiquidationManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LendingPoolLiquidationManager.Contract.LendingPoolLiquidationManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingPoolLiquidationManager *LendingPoolLiquidationManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPoolLiquidationManager.Contract.LendingPoolLiquidationManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingPoolLiquidationManager *LendingPoolLiquidationManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingPoolLiquidationManager.Contract.LendingPoolLiquidationManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingPoolLiquidationManager *LendingPoolLiquidationManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LendingPoolLiquidationManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingPoolLiquidationManager *LendingPoolLiquidationManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPoolLiquidationManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingPoolLiquidationManager *LendingPoolLiquidationManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingPoolLiquidationManager.Contract.contract.Transact(opts, method, params...)
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_LendingPoolLiquidationManager *LendingPoolLiquidationManagerCaller) AddressesProvider(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LendingPoolLiquidationManager.contract.Call(opts, out, "addressesProvider")
	return *ret0, err
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_LendingPoolLiquidationManager *LendingPoolLiquidationManagerSession) AddressesProvider() (common.Address, error) {
	return _LendingPoolLiquidationManager.Contract.AddressesProvider(&_LendingPoolLiquidationManager.CallOpts)
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_LendingPoolLiquidationManager *LendingPoolLiquidationManagerCallerSession) AddressesProvider() (common.Address, error) {
	return _LendingPoolLiquidationManager.Contract.AddressesProvider(&_LendingPoolLiquidationManager.CallOpts)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address _collateral, address _reserve, address _user, uint256 _purchaseAmount, bool _receiveAToken) payable returns(uint256, string)
func (_LendingPoolLiquidationManager *LendingPoolLiquidationManagerTransactor) LiquidationCall(opts *bind.TransactOpts, _collateral common.Address, _reserve common.Address, _user common.Address, _purchaseAmount *big.Int, _receiveAToken bool) (*types.Transaction, error) {
	return _LendingPoolLiquidationManager.contract.Transact(opts, "liquidationCall", _collateral, _reserve, _user, _purchaseAmount, _receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address _collateral, address _reserve, address _user, uint256 _purchaseAmount, bool _receiveAToken) payable returns(uint256, string)
func (_LendingPoolLiquidationManager *LendingPoolLiquidationManagerSession) LiquidationCall(_collateral common.Address, _reserve common.Address, _user common.Address, _purchaseAmount *big.Int, _receiveAToken bool) (*types.Transaction, error) {
	return _LendingPoolLiquidationManager.Contract.LiquidationCall(&_LendingPoolLiquidationManager.TransactOpts, _collateral, _reserve, _user, _purchaseAmount, _receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address _collateral, address _reserve, address _user, uint256 _purchaseAmount, bool _receiveAToken) payable returns(uint256, string)
func (_LendingPoolLiquidationManager *LendingPoolLiquidationManagerTransactorSession) LiquidationCall(_collateral common.Address, _reserve common.Address, _user common.Address, _purchaseAmount *big.Int, _receiveAToken bool) (*types.Transaction, error) {
	return _LendingPoolLiquidationManager.Contract.LiquidationCall(&_LendingPoolLiquidationManager.TransactOpts, _collateral, _reserve, _user, _purchaseAmount, _receiveAToken)
}

// LendingPoolLiquidationManagerLiquidationCallIterator is returned from FilterLiquidationCall and is used to iterate over the raw logs and unpacked data for LiquidationCall events raised by the LendingPoolLiquidationManager contract.
type LendingPoolLiquidationManagerLiquidationCallIterator struct {
	Event *LendingPoolLiquidationManagerLiquidationCall // Event containing the contract specifics and raw log

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
func (it *LendingPoolLiquidationManagerLiquidationCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolLiquidationManagerLiquidationCall)
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
		it.Event = new(LendingPoolLiquidationManagerLiquidationCall)
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
func (it *LendingPoolLiquidationManagerLiquidationCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolLiquidationManagerLiquidationCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolLiquidationManagerLiquidationCall represents a LiquidationCall event raised by the LendingPoolLiquidationManager contract.
type LendingPoolLiquidationManagerLiquidationCall struct {
	Collateral                 common.Address
	Reserve                    common.Address
	User                       common.Address
	PurchaseAmount             *big.Int
	LiquidatedCollateralAmount *big.Int
	AccruedBorrowInterest      *big.Int
	Liquidator                 common.Address
	ReceiveAToken              bool
	Timestamp                  *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterLiquidationCall is a free log retrieval operation binding the contract event 0x56864757fd5b1fc9f38f5f3a981cd8ae512ce41b902cf73fc506ee369c6bc237.
//
// Solidity: event LiquidationCall(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _purchaseAmount, uint256 _liquidatedCollateralAmount, uint256 _accruedBorrowInterest, address _liquidator, bool _receiveAToken, uint256 _timestamp)
func (_LendingPoolLiquidationManager *LendingPoolLiquidationManagerFilterer) FilterLiquidationCall(opts *bind.FilterOpts, _collateral []common.Address, _reserve []common.Address, _user []common.Address) (*LendingPoolLiquidationManagerLiquidationCallIterator, error) {

	var _collateralRule []interface{}
	for _, _collateralItem := range _collateral {
		_collateralRule = append(_collateralRule, _collateralItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _LendingPoolLiquidationManager.contract.FilterLogs(opts, "LiquidationCall", _collateralRule, _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolLiquidationManagerLiquidationCallIterator{contract: _LendingPoolLiquidationManager.contract, event: "LiquidationCall", logs: logs, sub: sub}, nil
}

// WatchLiquidationCall is a free log subscription operation binding the contract event 0x56864757fd5b1fc9f38f5f3a981cd8ae512ce41b902cf73fc506ee369c6bc237.
//
// Solidity: event LiquidationCall(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _purchaseAmount, uint256 _liquidatedCollateralAmount, uint256 _accruedBorrowInterest, address _liquidator, bool _receiveAToken, uint256 _timestamp)
func (_LendingPoolLiquidationManager *LendingPoolLiquidationManagerFilterer) WatchLiquidationCall(opts *bind.WatchOpts, sink chan<- *LendingPoolLiquidationManagerLiquidationCall, _collateral []common.Address, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _collateralRule []interface{}
	for _, _collateralItem := range _collateral {
		_collateralRule = append(_collateralRule, _collateralItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _LendingPoolLiquidationManager.contract.WatchLogs(opts, "LiquidationCall", _collateralRule, _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolLiquidationManagerLiquidationCall)
				if err := _LendingPoolLiquidationManager.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
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

// ParseLiquidationCall is a log parse operation binding the contract event 0x56864757fd5b1fc9f38f5f3a981cd8ae512ce41b902cf73fc506ee369c6bc237.
//
// Solidity: event LiquidationCall(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _purchaseAmount, uint256 _liquidatedCollateralAmount, uint256 _accruedBorrowInterest, address _liquidator, bool _receiveAToken, uint256 _timestamp)
func (_LendingPoolLiquidationManager *LendingPoolLiquidationManagerFilterer) ParseLiquidationCall(log types.Log) (*LendingPoolLiquidationManagerLiquidationCall, error) {
	event := new(LendingPoolLiquidationManagerLiquidationCall)
	if err := _LendingPoolLiquidationManager.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingPoolLiquidationManagerOriginationFeeLiquidatedIterator is returned from FilterOriginationFeeLiquidated and is used to iterate over the raw logs and unpacked data for OriginationFeeLiquidated events raised by the LendingPoolLiquidationManager contract.
type LendingPoolLiquidationManagerOriginationFeeLiquidatedIterator struct {
	Event *LendingPoolLiquidationManagerOriginationFeeLiquidated // Event containing the contract specifics and raw log

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
func (it *LendingPoolLiquidationManagerOriginationFeeLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolLiquidationManagerOriginationFeeLiquidated)
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
		it.Event = new(LendingPoolLiquidationManagerOriginationFeeLiquidated)
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
func (it *LendingPoolLiquidationManagerOriginationFeeLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolLiquidationManagerOriginationFeeLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolLiquidationManagerOriginationFeeLiquidated represents a OriginationFeeLiquidated event raised by the LendingPoolLiquidationManager contract.
type LendingPoolLiquidationManagerOriginationFeeLiquidated struct {
	Collateral                 common.Address
	Reserve                    common.Address
	User                       common.Address
	FeeLiquidated              *big.Int
	LiquidatedCollateralForFee *big.Int
	Timestamp                  *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterOriginationFeeLiquidated is a free log retrieval operation binding the contract event 0x36ca8b16d61dc13b1062adff83e3778ab92d14f9e35bfe9fd1283e02b13fb0a1.
//
// Solidity: event OriginationFeeLiquidated(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _feeLiquidated, uint256 _liquidatedCollateralForFee, uint256 _timestamp)
func (_LendingPoolLiquidationManager *LendingPoolLiquidationManagerFilterer) FilterOriginationFeeLiquidated(opts *bind.FilterOpts, _collateral []common.Address, _reserve []common.Address, _user []common.Address) (*LendingPoolLiquidationManagerOriginationFeeLiquidatedIterator, error) {

	var _collateralRule []interface{}
	for _, _collateralItem := range _collateral {
		_collateralRule = append(_collateralRule, _collateralItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _LendingPoolLiquidationManager.contract.FilterLogs(opts, "OriginationFeeLiquidated", _collateralRule, _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolLiquidationManagerOriginationFeeLiquidatedIterator{contract: _LendingPoolLiquidationManager.contract, event: "OriginationFeeLiquidated", logs: logs, sub: sub}, nil
}

// WatchOriginationFeeLiquidated is a free log subscription operation binding the contract event 0x36ca8b16d61dc13b1062adff83e3778ab92d14f9e35bfe9fd1283e02b13fb0a1.
//
// Solidity: event OriginationFeeLiquidated(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _feeLiquidated, uint256 _liquidatedCollateralForFee, uint256 _timestamp)
func (_LendingPoolLiquidationManager *LendingPoolLiquidationManagerFilterer) WatchOriginationFeeLiquidated(opts *bind.WatchOpts, sink chan<- *LendingPoolLiquidationManagerOriginationFeeLiquidated, _collateral []common.Address, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _collateralRule []interface{}
	for _, _collateralItem := range _collateral {
		_collateralRule = append(_collateralRule, _collateralItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _LendingPoolLiquidationManager.contract.WatchLogs(opts, "OriginationFeeLiquidated", _collateralRule, _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolLiquidationManagerOriginationFeeLiquidated)
				if err := _LendingPoolLiquidationManager.contract.UnpackLog(event, "OriginationFeeLiquidated", log); err != nil {
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

// ParseOriginationFeeLiquidated is a log parse operation binding the contract event 0x36ca8b16d61dc13b1062adff83e3778ab92d14f9e35bfe9fd1283e02b13fb0a1.
//
// Solidity: event OriginationFeeLiquidated(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _feeLiquidated, uint256 _liquidatedCollateralForFee, uint256 _timestamp)
func (_LendingPoolLiquidationManager *LendingPoolLiquidationManagerFilterer) ParseOriginationFeeLiquidated(log types.Log) (*LendingPoolLiquidationManagerOriginationFeeLiquidated, error) {
	event := new(LendingPoolLiquidationManagerOriginationFeeLiquidated)
	if err := _LendingPoolLiquidationManager.contract.UnpackLog(event, "OriginationFeeLiquidated", log); err != nil {
		return nil, err
	}
	return event, nil
}
