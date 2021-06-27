// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lending_pool_parameters_provider

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

// LendingPoolParametersProviderABI is the input ABI used to generate the binding from.
const LendingPoolParametersProviderABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_addressesProvider\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMaxStableRateBorrowSizePercent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRebalanceDownRateDelta\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFlashLoanFeesInBips\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// LendingPoolParametersProviderBin is the compiled bytecode used for deploying new contracts.
var LendingPoolParametersProviderBin = "0x60806040526000805534801561001457600080fd5b50610204806100246000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806346f4f8d114610051578063586feb401461006b578063c4d66de81461008c578063d6b725ac146100c1575b600080fd5b6100596100c9565b60408051918252519081900360200190f35b6100736100d8565b6040805192835260208301919091528051918290030190f35b6100bf600480360360208110156100a257600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166100e1565b005b61005961019a565b6aa56fa5b99019a5c800000090565b6023610bb89091565b60006100eb61019f565b60015490915060ff168061010257506101026101a4565b8061010e575060005481115b1515610165576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001806101ab602e913960400191505060405180910390fd5b60015460ff16158015610184576001805460ff19168117905560008290555b8015610195576001805460ff191690555b505050565b601990565b600190565b303b159056fe436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564a165627a7a723058201b875980cbe2e73f28839dcab5335718c4fcdf181957d453f7d5b27e155a51290029"

// DeployLendingPoolParametersProvider deploys a new Ethereum contract, binding an instance of LendingPoolParametersProvider to it.
func DeployLendingPoolParametersProvider(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LendingPoolParametersProvider, error) {
	parsed, err := abi.JSON(strings.NewReader(LendingPoolParametersProviderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LendingPoolParametersProviderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LendingPoolParametersProvider{LendingPoolParametersProviderCaller: LendingPoolParametersProviderCaller{contract: contract}, LendingPoolParametersProviderTransactor: LendingPoolParametersProviderTransactor{contract: contract}, LendingPoolParametersProviderFilterer: LendingPoolParametersProviderFilterer{contract: contract}}, nil
}

// LendingPoolParametersProvider is an auto generated Go binding around an Ethereum contract.
type LendingPoolParametersProvider struct {
	LendingPoolParametersProviderCaller     // Read-only binding to the contract
	LendingPoolParametersProviderTransactor // Write-only binding to the contract
	LendingPoolParametersProviderFilterer   // Log filterer for contract events
}

// LendingPoolParametersProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type LendingPoolParametersProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolParametersProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LendingPoolParametersProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolParametersProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LendingPoolParametersProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolParametersProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LendingPoolParametersProviderSession struct {
	Contract     *LendingPoolParametersProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// LendingPoolParametersProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LendingPoolParametersProviderCallerSession struct {
	Contract *LendingPoolParametersProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// LendingPoolParametersProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LendingPoolParametersProviderTransactorSession struct {
	Contract     *LendingPoolParametersProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// LendingPoolParametersProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type LendingPoolParametersProviderRaw struct {
	Contract *LendingPoolParametersProvider // Generic contract binding to access the raw methods on
}

// LendingPoolParametersProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LendingPoolParametersProviderCallerRaw struct {
	Contract *LendingPoolParametersProviderCaller // Generic read-only contract binding to access the raw methods on
}

// LendingPoolParametersProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LendingPoolParametersProviderTransactorRaw struct {
	Contract *LendingPoolParametersProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLendingPoolParametersProvider creates a new instance of LendingPoolParametersProvider, bound to a specific deployed contract.
func NewLendingPoolParametersProvider(address common.Address, backend bind.ContractBackend) (*LendingPoolParametersProvider, error) {
	contract, err := bindLendingPoolParametersProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LendingPoolParametersProvider{LendingPoolParametersProviderCaller: LendingPoolParametersProviderCaller{contract: contract}, LendingPoolParametersProviderTransactor: LendingPoolParametersProviderTransactor{contract: contract}, LendingPoolParametersProviderFilterer: LendingPoolParametersProviderFilterer{contract: contract}}, nil
}

// NewLendingPoolParametersProviderCaller creates a new read-only instance of LendingPoolParametersProvider, bound to a specific deployed contract.
func NewLendingPoolParametersProviderCaller(address common.Address, caller bind.ContractCaller) (*LendingPoolParametersProviderCaller, error) {
	contract, err := bindLendingPoolParametersProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LendingPoolParametersProviderCaller{contract: contract}, nil
}

// NewLendingPoolParametersProviderTransactor creates a new write-only instance of LendingPoolParametersProvider, bound to a specific deployed contract.
func NewLendingPoolParametersProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*LendingPoolParametersProviderTransactor, error) {
	contract, err := bindLendingPoolParametersProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LendingPoolParametersProviderTransactor{contract: contract}, nil
}

// NewLendingPoolParametersProviderFilterer creates a new log filterer instance of LendingPoolParametersProvider, bound to a specific deployed contract.
func NewLendingPoolParametersProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*LendingPoolParametersProviderFilterer, error) {
	contract, err := bindLendingPoolParametersProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LendingPoolParametersProviderFilterer{contract: contract}, nil
}

// bindLendingPoolParametersProvider binds a generic wrapper to an already deployed contract.
func bindLendingPoolParametersProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LendingPoolParametersProviderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingPoolParametersProvider *LendingPoolParametersProviderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LendingPoolParametersProvider.Contract.LendingPoolParametersProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingPoolParametersProvider *LendingPoolParametersProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPoolParametersProvider.Contract.LendingPoolParametersProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingPoolParametersProvider *LendingPoolParametersProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingPoolParametersProvider.Contract.LendingPoolParametersProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingPoolParametersProvider *LendingPoolParametersProviderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LendingPoolParametersProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingPoolParametersProvider *LendingPoolParametersProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPoolParametersProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingPoolParametersProvider *LendingPoolParametersProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingPoolParametersProvider.Contract.contract.Transact(opts, method, params...)
}

// GetFlashLoanFeesInBips is a free data retrieval call binding the contract method 0x586feb40.
//
// Solidity: function getFlashLoanFeesInBips() pure returns(uint256, uint256)
func (_LendingPoolParametersProvider *LendingPoolParametersProviderCaller) GetFlashLoanFeesInBips(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _LendingPoolParametersProvider.contract.Call(opts, out, "getFlashLoanFeesInBips")
	return *ret0, *ret1, err
}

// GetFlashLoanFeesInBips is a free data retrieval call binding the contract method 0x586feb40.
//
// Solidity: function getFlashLoanFeesInBips() pure returns(uint256, uint256)
func (_LendingPoolParametersProvider *LendingPoolParametersProviderSession) GetFlashLoanFeesInBips() (*big.Int, *big.Int, error) {
	return _LendingPoolParametersProvider.Contract.GetFlashLoanFeesInBips(&_LendingPoolParametersProvider.CallOpts)
}

// GetFlashLoanFeesInBips is a free data retrieval call binding the contract method 0x586feb40.
//
// Solidity: function getFlashLoanFeesInBips() pure returns(uint256, uint256)
func (_LendingPoolParametersProvider *LendingPoolParametersProviderCallerSession) GetFlashLoanFeesInBips() (*big.Int, *big.Int, error) {
	return _LendingPoolParametersProvider.Contract.GetFlashLoanFeesInBips(&_LendingPoolParametersProvider.CallOpts)
}

// GetMaxStableRateBorrowSizePercent is a free data retrieval call binding the contract method 0xd6b725ac.
//
// Solidity: function getMaxStableRateBorrowSizePercent() pure returns(uint256)
func (_LendingPoolParametersProvider *LendingPoolParametersProviderCaller) GetMaxStableRateBorrowSizePercent(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolParametersProvider.contract.Call(opts, out, "getMaxStableRateBorrowSizePercent")
	return *ret0, err
}

// GetMaxStableRateBorrowSizePercent is a free data retrieval call binding the contract method 0xd6b725ac.
//
// Solidity: function getMaxStableRateBorrowSizePercent() pure returns(uint256)
func (_LendingPoolParametersProvider *LendingPoolParametersProviderSession) GetMaxStableRateBorrowSizePercent() (*big.Int, error) {
	return _LendingPoolParametersProvider.Contract.GetMaxStableRateBorrowSizePercent(&_LendingPoolParametersProvider.CallOpts)
}

// GetMaxStableRateBorrowSizePercent is a free data retrieval call binding the contract method 0xd6b725ac.
//
// Solidity: function getMaxStableRateBorrowSizePercent() pure returns(uint256)
func (_LendingPoolParametersProvider *LendingPoolParametersProviderCallerSession) GetMaxStableRateBorrowSizePercent() (*big.Int, error) {
	return _LendingPoolParametersProvider.Contract.GetMaxStableRateBorrowSizePercent(&_LendingPoolParametersProvider.CallOpts)
}

// GetRebalanceDownRateDelta is a free data retrieval call binding the contract method 0x46f4f8d1.
//
// Solidity: function getRebalanceDownRateDelta() pure returns(uint256)
func (_LendingPoolParametersProvider *LendingPoolParametersProviderCaller) GetRebalanceDownRateDelta(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolParametersProvider.contract.Call(opts, out, "getRebalanceDownRateDelta")
	return *ret0, err
}

// GetRebalanceDownRateDelta is a free data retrieval call binding the contract method 0x46f4f8d1.
//
// Solidity: function getRebalanceDownRateDelta() pure returns(uint256)
func (_LendingPoolParametersProvider *LendingPoolParametersProviderSession) GetRebalanceDownRateDelta() (*big.Int, error) {
	return _LendingPoolParametersProvider.Contract.GetRebalanceDownRateDelta(&_LendingPoolParametersProvider.CallOpts)
}

// GetRebalanceDownRateDelta is a free data retrieval call binding the contract method 0x46f4f8d1.
//
// Solidity: function getRebalanceDownRateDelta() pure returns(uint256)
func (_LendingPoolParametersProvider *LendingPoolParametersProviderCallerSession) GetRebalanceDownRateDelta() (*big.Int, error) {
	return _LendingPoolParametersProvider.Contract.GetRebalanceDownRateDelta(&_LendingPoolParametersProvider.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _addressesProvider) returns()
func (_LendingPoolParametersProvider *LendingPoolParametersProviderTransactor) Initialize(opts *bind.TransactOpts, _addressesProvider common.Address) (*types.Transaction, error) {
	return _LendingPoolParametersProvider.contract.Transact(opts, "initialize", _addressesProvider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _addressesProvider) returns()
func (_LendingPoolParametersProvider *LendingPoolParametersProviderSession) Initialize(_addressesProvider common.Address) (*types.Transaction, error) {
	return _LendingPoolParametersProvider.Contract.Initialize(&_LendingPoolParametersProvider.TransactOpts, _addressesProvider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _addressesProvider) returns()
func (_LendingPoolParametersProvider *LendingPoolParametersProviderTransactorSession) Initialize(_addressesProvider common.Address) (*types.Transaction, error) {
	return _LendingPoolParametersProvider.Contract.Initialize(&_LendingPoolParametersProvider.TransactOpts, _addressesProvider)
}
