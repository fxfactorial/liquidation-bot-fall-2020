// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simple_price_oracle

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

// SimplePriceOracleABI is the input ABI used to generate the binding from.
const SimplePriceOracleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousPriceMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestedPriceMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPriceMantissa\",\"type\":\"uint256\"}],\"name\":\"PricePosted\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPriceOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"getUnderlyingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingPriceMantissa\",\"type\":\"uint256\"}],\"name\":\"setUnderlyingPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setDirectPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"assetPrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SimplePriceOracleBin is the compiled bytecode used for deploying new contracts.
var SimplePriceOracleBin = "0x608060405234801561001057600080fd5b506105d7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c80635e9a523c116100505780635e9a523c146100c657806366331bba146100fe578063fc57d4df1461011a57610067565b806309a8acb01461006c578063127ffda01461009a575b600080fd5b6100986004803603604081101561008257600080fd5b506001600160a01b038135169060200135610140565b005b610098600480360360408110156100b057600080fd5b506001600160a01b0381351690602001356101b8565b6100ec600480360360208110156100dc57600080fd5b50356001600160a01b031661029a565b60408051918252519081900360200190f35b6101066102b9565b604080519115158252519081900360200190f35b6100ec6004803603602081101561013057600080fd5b50356001600160a01b03166102be565b6001600160a01b038216600081815260208181526040918290205482519384529083015281810183905260608201839052517fdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae79181900360800190a16001600160a01b03909116600090815260208190526040902055565b6000826001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b1580156101f357600080fd5b505afa158015610207573d6000803e3d6000fd5b505050506040513d602081101561021d57600080fd5b50516001600160a01b038116600081815260208181526040918290205482519384529083015281810185905260608201859052519192507fdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae7919081900360800190a16001600160a01b031660009081526020819052604090205550565b6001600160a01b0381166000908152602081905260409020545b919050565b600181565b600061041e826001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b1580156102fc57600080fd5b505afa158015610310573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561033957600080fd5b810190808051604051939291908464010000000082111561035957600080fd5b90830190602082018581111561036e57600080fd5b825164010000000081118282018810171561038857600080fd5b82525081516020918201929091019080838360005b838110156103b557818101518382015260200161039d565b50505050905090810190601f1680156103e25780820380516001836020036101000a031916815260200191505b5060408181019052600481527f6345544800000000000000000000000000000000000000000000000000000000602082015292506104bb915050565b156104325750670de0b6b3a76400006102b4565b600080836001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b15801561046e57600080fd5b505afa158015610482573d6000803e3d6000fd5b505050506040513d602081101561049857600080fd5b50516001600160a01b0316815260208101919091526040016000205490506102b4565b6000816040516020018082805190602001908083835b602083106104f05780518252601f1990920191602091820191016104d1565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120836040516020018082805190602001908083835b6020831061055e5780518252601f19909201916020918201910161053f565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001201490509291505056fea265627a7a72315820bbe13382ebcfaf9b268c160bd0567beca0c519fcae82d5ad45ee5e5170361b4564736f6c63430005100032"

// DeploySimplePriceOracle deploys a new Ethereum contract, binding an instance of SimplePriceOracle to it.
func DeploySimplePriceOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimplePriceOracle, error) {
	parsed, err := abi.JSON(strings.NewReader(SimplePriceOracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimplePriceOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimplePriceOracle{SimplePriceOracleCaller: SimplePriceOracleCaller{contract: contract}, SimplePriceOracleTransactor: SimplePriceOracleTransactor{contract: contract}, SimplePriceOracleFilterer: SimplePriceOracleFilterer{contract: contract}}, nil
}

// SimplePriceOracle is an auto generated Go binding around an Ethereum contract.
type SimplePriceOracle struct {
	SimplePriceOracleCaller     // Read-only binding to the contract
	SimplePriceOracleTransactor // Write-only binding to the contract
	SimplePriceOracleFilterer   // Log filterer for contract events
}

// SimplePriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimplePriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplePriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimplePriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplePriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimplePriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplePriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimplePriceOracleSession struct {
	Contract     *SimplePriceOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SimplePriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimplePriceOracleCallerSession struct {
	Contract *SimplePriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// SimplePriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimplePriceOracleTransactorSession struct {
	Contract     *SimplePriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SimplePriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimplePriceOracleRaw struct {
	Contract *SimplePriceOracle // Generic contract binding to access the raw methods on
}

// SimplePriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimplePriceOracleCallerRaw struct {
	Contract *SimplePriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// SimplePriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimplePriceOracleTransactorRaw struct {
	Contract *SimplePriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimplePriceOracle creates a new instance of SimplePriceOracle, bound to a specific deployed contract.
func NewSimplePriceOracle(address common.Address, backend bind.ContractBackend) (*SimplePriceOracle, error) {
	contract, err := bindSimplePriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimplePriceOracle{SimplePriceOracleCaller: SimplePriceOracleCaller{contract: contract}, SimplePriceOracleTransactor: SimplePriceOracleTransactor{contract: contract}, SimplePriceOracleFilterer: SimplePriceOracleFilterer{contract: contract}}, nil
}

// NewSimplePriceOracleCaller creates a new read-only instance of SimplePriceOracle, bound to a specific deployed contract.
func NewSimplePriceOracleCaller(address common.Address, caller bind.ContractCaller) (*SimplePriceOracleCaller, error) {
	contract, err := bindSimplePriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimplePriceOracleCaller{contract: contract}, nil
}

// NewSimplePriceOracleTransactor creates a new write-only instance of SimplePriceOracle, bound to a specific deployed contract.
func NewSimplePriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*SimplePriceOracleTransactor, error) {
	contract, err := bindSimplePriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimplePriceOracleTransactor{contract: contract}, nil
}

// NewSimplePriceOracleFilterer creates a new log filterer instance of SimplePriceOracle, bound to a specific deployed contract.
func NewSimplePriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*SimplePriceOracleFilterer, error) {
	contract, err := bindSimplePriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimplePriceOracleFilterer{contract: contract}, nil
}

// bindSimplePriceOracle binds a generic wrapper to an already deployed contract.
func bindSimplePriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimplePriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimplePriceOracle *SimplePriceOracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimplePriceOracle.Contract.SimplePriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimplePriceOracle *SimplePriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimplePriceOracle.Contract.SimplePriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimplePriceOracle *SimplePriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimplePriceOracle.Contract.SimplePriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimplePriceOracle *SimplePriceOracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SimplePriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimplePriceOracle *SimplePriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimplePriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimplePriceOracle *SimplePriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimplePriceOracle.Contract.contract.Transact(opts, method, params...)
}

// AssetPrices is a free data retrieval call binding the contract method 0x5e9a523c.
//
// Solidity: function assetPrices(address asset) view returns(uint256)
func (_SimplePriceOracle *SimplePriceOracleCaller) AssetPrices(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimplePriceOracle.contract.Call(opts, out, "assetPrices", asset)
	return *ret0, err
}

// AssetPrices is a free data retrieval call binding the contract method 0x5e9a523c.
//
// Solidity: function assetPrices(address asset) view returns(uint256)
func (_SimplePriceOracle *SimplePriceOracleSession) AssetPrices(asset common.Address) (*big.Int, error) {
	return _SimplePriceOracle.Contract.AssetPrices(&_SimplePriceOracle.CallOpts, asset)
}

// AssetPrices is a free data retrieval call binding the contract method 0x5e9a523c.
//
// Solidity: function assetPrices(address asset) view returns(uint256)
func (_SimplePriceOracle *SimplePriceOracleCallerSession) AssetPrices(asset common.Address) (*big.Int, error) {
	return _SimplePriceOracle.Contract.AssetPrices(&_SimplePriceOracle.CallOpts, asset)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address cToken) view returns(uint256)
func (_SimplePriceOracle *SimplePriceOracleCaller) GetUnderlyingPrice(opts *bind.CallOpts, cToken common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SimplePriceOracle.contract.Call(opts, out, "getUnderlyingPrice", cToken)
	return *ret0, err
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address cToken) view returns(uint256)
func (_SimplePriceOracle *SimplePriceOracleSession) GetUnderlyingPrice(cToken common.Address) (*big.Int, error) {
	return _SimplePriceOracle.Contract.GetUnderlyingPrice(&_SimplePriceOracle.CallOpts, cToken)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address cToken) view returns(uint256)
func (_SimplePriceOracle *SimplePriceOracleCallerSession) GetUnderlyingPrice(cToken common.Address) (*big.Int, error) {
	return _SimplePriceOracle.Contract.GetUnderlyingPrice(&_SimplePriceOracle.CallOpts, cToken)
}

// IsPriceOracle is a free data retrieval call binding the contract method 0x66331bba.
//
// Solidity: function isPriceOracle() view returns(bool)
func (_SimplePriceOracle *SimplePriceOracleCaller) IsPriceOracle(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SimplePriceOracle.contract.Call(opts, out, "isPriceOracle")
	return *ret0, err
}

// IsPriceOracle is a free data retrieval call binding the contract method 0x66331bba.
//
// Solidity: function isPriceOracle() view returns(bool)
func (_SimplePriceOracle *SimplePriceOracleSession) IsPriceOracle() (bool, error) {
	return _SimplePriceOracle.Contract.IsPriceOracle(&_SimplePriceOracle.CallOpts)
}

// IsPriceOracle is a free data retrieval call binding the contract method 0x66331bba.
//
// Solidity: function isPriceOracle() view returns(bool)
func (_SimplePriceOracle *SimplePriceOracleCallerSession) IsPriceOracle() (bool, error) {
	return _SimplePriceOracle.Contract.IsPriceOracle(&_SimplePriceOracle.CallOpts)
}

// SetDirectPrice is a paid mutator transaction binding the contract method 0x09a8acb0.
//
// Solidity: function setDirectPrice(address asset, uint256 price) returns()
func (_SimplePriceOracle *SimplePriceOracleTransactor) SetDirectPrice(opts *bind.TransactOpts, asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _SimplePriceOracle.contract.Transact(opts, "setDirectPrice", asset, price)
}

// SetDirectPrice is a paid mutator transaction binding the contract method 0x09a8acb0.
//
// Solidity: function setDirectPrice(address asset, uint256 price) returns()
func (_SimplePriceOracle *SimplePriceOracleSession) SetDirectPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _SimplePriceOracle.Contract.SetDirectPrice(&_SimplePriceOracle.TransactOpts, asset, price)
}

// SetDirectPrice is a paid mutator transaction binding the contract method 0x09a8acb0.
//
// Solidity: function setDirectPrice(address asset, uint256 price) returns()
func (_SimplePriceOracle *SimplePriceOracleTransactorSession) SetDirectPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _SimplePriceOracle.Contract.SetDirectPrice(&_SimplePriceOracle.TransactOpts, asset, price)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x127ffda0.
//
// Solidity: function setUnderlyingPrice(address cToken, uint256 underlyingPriceMantissa) returns()
func (_SimplePriceOracle *SimplePriceOracleTransactor) SetUnderlyingPrice(opts *bind.TransactOpts, cToken common.Address, underlyingPriceMantissa *big.Int) (*types.Transaction, error) {
	return _SimplePriceOracle.contract.Transact(opts, "setUnderlyingPrice", cToken, underlyingPriceMantissa)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x127ffda0.
//
// Solidity: function setUnderlyingPrice(address cToken, uint256 underlyingPriceMantissa) returns()
func (_SimplePriceOracle *SimplePriceOracleSession) SetUnderlyingPrice(cToken common.Address, underlyingPriceMantissa *big.Int) (*types.Transaction, error) {
	return _SimplePriceOracle.Contract.SetUnderlyingPrice(&_SimplePriceOracle.TransactOpts, cToken, underlyingPriceMantissa)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x127ffda0.
//
// Solidity: function setUnderlyingPrice(address cToken, uint256 underlyingPriceMantissa) returns()
func (_SimplePriceOracle *SimplePriceOracleTransactorSession) SetUnderlyingPrice(cToken common.Address, underlyingPriceMantissa *big.Int) (*types.Transaction, error) {
	return _SimplePriceOracle.Contract.SetUnderlyingPrice(&_SimplePriceOracle.TransactOpts, cToken, underlyingPriceMantissa)
}

// SimplePriceOraclePricePostedIterator is returned from FilterPricePosted and is used to iterate over the raw logs and unpacked data for PricePosted events raised by the SimplePriceOracle contract.
type SimplePriceOraclePricePostedIterator struct {
	Event *SimplePriceOraclePricePosted // Event containing the contract specifics and raw log

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
func (it *SimplePriceOraclePricePostedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimplePriceOraclePricePosted)
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
		it.Event = new(SimplePriceOraclePricePosted)
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
func (it *SimplePriceOraclePricePostedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimplePriceOraclePricePostedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimplePriceOraclePricePosted represents a PricePosted event raised by the SimplePriceOracle contract.
type SimplePriceOraclePricePosted struct {
	Asset                  common.Address
	PreviousPriceMantissa  *big.Int
	RequestedPriceMantissa *big.Int
	NewPriceMantissa       *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterPricePosted is a free log retrieval operation binding the contract event 0xdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae7.
//
// Solidity: event PricePosted(address asset, uint256 previousPriceMantissa, uint256 requestedPriceMantissa, uint256 newPriceMantissa)
func (_SimplePriceOracle *SimplePriceOracleFilterer) FilterPricePosted(opts *bind.FilterOpts) (*SimplePriceOraclePricePostedIterator, error) {

	logs, sub, err := _SimplePriceOracle.contract.FilterLogs(opts, "PricePosted")
	if err != nil {
		return nil, err
	}
	return &SimplePriceOraclePricePostedIterator{contract: _SimplePriceOracle.contract, event: "PricePosted", logs: logs, sub: sub}, nil
}

// WatchPricePosted is a free log subscription operation binding the contract event 0xdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae7.
//
// Solidity: event PricePosted(address asset, uint256 previousPriceMantissa, uint256 requestedPriceMantissa, uint256 newPriceMantissa)
func (_SimplePriceOracle *SimplePriceOracleFilterer) WatchPricePosted(opts *bind.WatchOpts, sink chan<- *SimplePriceOraclePricePosted) (event.Subscription, error) {

	logs, sub, err := _SimplePriceOracle.contract.WatchLogs(opts, "PricePosted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimplePriceOraclePricePosted)
				if err := _SimplePriceOracle.contract.UnpackLog(event, "PricePosted", log); err != nil {
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

// ParsePricePosted is a log parse operation binding the contract event 0xdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae7.
//
// Solidity: event PricePosted(address asset, uint256 previousPriceMantissa, uint256 requestedPriceMantissa, uint256 newPriceMantissa)
func (_SimplePriceOracle *SimplePriceOracleFilterer) ParsePricePosted(log types.Log) (*SimplePriceOraclePricePosted, error) {
	event := new(SimplePriceOraclePricePosted)
	if err := _SimplePriceOracle.contract.UnpackLog(event, "PricePosted", log); err != nil {
		return nil, err
	}
	return event, nil
}
