// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave_price_oracle

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

// AavePriceOracleABI is the input ABI used to generate the binding from.
const AavePriceOracleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_asset\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AssetPriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"EthPriceUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getAssetPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_asset\",\"type\":\"address\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setAssetPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEthUsdPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setEthUsdPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AavePriceOracleBin is the compiled bytecode used for deploying new contracts.
var AavePriceOracleBin = "0x608060405234801561001057600080fd5b506101fa806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806351323f7214610051578063a0a8045e1461008c578063b3596f07146100a6578063b951883a146100d9575b600080fd5b61008a6004803603604081101561006757600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356100f6565b005b61009461015e565b60408051918252519081900360200190f35b610094600480360360208110156100bc57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610164565b61008a600480360360208110156100ef57600080fd5b503561018c565b73ffffffffffffffffffffffffffffffffffffffff821660008181526020818152604091829020849055815192835282018390524282820152517fce6e0b57367bae95ca7198e1172f653ea64a645c16ab586b4cefa9237bfc2d929181900360600190a15050565b60015490565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b60018190556040805182815242602082015281517fb4f35977939fa8b5ffe552d517a8ff5223046b1fdd3ee0068ae38d1e2b8d0016929181900390910190a15056fea165627a7a72305820c1ae7a0789e2f83118b21508678b640ec2856161a62c47a53b7edbec44205e930029"

// DeployAavePriceOracle deploys a new Ethereum contract, binding an instance of AavePriceOracle to it.
func DeployAavePriceOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AavePriceOracle, error) {
	parsed, err := abi.JSON(strings.NewReader(AavePriceOracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AavePriceOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AavePriceOracle{AavePriceOracleCaller: AavePriceOracleCaller{contract: contract}, AavePriceOracleTransactor: AavePriceOracleTransactor{contract: contract}, AavePriceOracleFilterer: AavePriceOracleFilterer{contract: contract}}, nil
}

// AavePriceOracle is an auto generated Go binding around an Ethereum contract.
type AavePriceOracle struct {
	AavePriceOracleCaller     // Read-only binding to the contract
	AavePriceOracleTransactor // Write-only binding to the contract
	AavePriceOracleFilterer   // Log filterer for contract events
}

// AavePriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type AavePriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AavePriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AavePriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AavePriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AavePriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AavePriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AavePriceOracleSession struct {
	Contract     *AavePriceOracle  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AavePriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AavePriceOracleCallerSession struct {
	Contract *AavePriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AavePriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AavePriceOracleTransactorSession struct {
	Contract     *AavePriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AavePriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type AavePriceOracleRaw struct {
	Contract *AavePriceOracle // Generic contract binding to access the raw methods on
}

// AavePriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AavePriceOracleCallerRaw struct {
	Contract *AavePriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// AavePriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AavePriceOracleTransactorRaw struct {
	Contract *AavePriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAavePriceOracle creates a new instance of AavePriceOracle, bound to a specific deployed contract.
func NewAavePriceOracle(address common.Address, backend bind.ContractBackend) (*AavePriceOracle, error) {
	contract, err := bindAavePriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AavePriceOracle{AavePriceOracleCaller: AavePriceOracleCaller{contract: contract}, AavePriceOracleTransactor: AavePriceOracleTransactor{contract: contract}, AavePriceOracleFilterer: AavePriceOracleFilterer{contract: contract}}, nil
}

// NewAavePriceOracleCaller creates a new read-only instance of AavePriceOracle, bound to a specific deployed contract.
func NewAavePriceOracleCaller(address common.Address, caller bind.ContractCaller) (*AavePriceOracleCaller, error) {
	contract, err := bindAavePriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AavePriceOracleCaller{contract: contract}, nil
}

// NewAavePriceOracleTransactor creates a new write-only instance of AavePriceOracle, bound to a specific deployed contract.
func NewAavePriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*AavePriceOracleTransactor, error) {
	contract, err := bindAavePriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AavePriceOracleTransactor{contract: contract}, nil
}

// NewAavePriceOracleFilterer creates a new log filterer instance of AavePriceOracle, bound to a specific deployed contract.
func NewAavePriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*AavePriceOracleFilterer, error) {
	contract, err := bindAavePriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AavePriceOracleFilterer{contract: contract}, nil
}

// bindAavePriceOracle binds a generic wrapper to an already deployed contract.
func bindAavePriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AavePriceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AavePriceOracle *AavePriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AavePriceOracle.Contract.AavePriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AavePriceOracle *AavePriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AavePriceOracle.Contract.AavePriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AavePriceOracle *AavePriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AavePriceOracle.Contract.AavePriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AavePriceOracle *AavePriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AavePriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AavePriceOracle *AavePriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AavePriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AavePriceOracle *AavePriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AavePriceOracle.Contract.contract.Transact(opts, method, params...)
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address _asset) view returns(uint256)
func (_AavePriceOracle *AavePriceOracleCaller) GetAssetPrice(opts *bind.CallOpts, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AavePriceOracle.contract.Call(opts, &out, "getAssetPrice", _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address _asset) view returns(uint256)
func (_AavePriceOracle *AavePriceOracleSession) GetAssetPrice(_asset common.Address) (*big.Int, error) {
	return _AavePriceOracle.Contract.GetAssetPrice(&_AavePriceOracle.CallOpts, _asset)
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address _asset) view returns(uint256)
func (_AavePriceOracle *AavePriceOracleCallerSession) GetAssetPrice(_asset common.Address) (*big.Int, error) {
	return _AavePriceOracle.Contract.GetAssetPrice(&_AavePriceOracle.CallOpts, _asset)
}

// GetEthUsdPrice is a free data retrieval call binding the contract method 0xa0a8045e.
//
// Solidity: function getEthUsdPrice() view returns(uint256)
func (_AavePriceOracle *AavePriceOracleCaller) GetEthUsdPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AavePriceOracle.contract.Call(opts, &out, "getEthUsdPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthUsdPrice is a free data retrieval call binding the contract method 0xa0a8045e.
//
// Solidity: function getEthUsdPrice() view returns(uint256)
func (_AavePriceOracle *AavePriceOracleSession) GetEthUsdPrice() (*big.Int, error) {
	return _AavePriceOracle.Contract.GetEthUsdPrice(&_AavePriceOracle.CallOpts)
}

// GetEthUsdPrice is a free data retrieval call binding the contract method 0xa0a8045e.
//
// Solidity: function getEthUsdPrice() view returns(uint256)
func (_AavePriceOracle *AavePriceOracleCallerSession) GetEthUsdPrice() (*big.Int, error) {
	return _AavePriceOracle.Contract.GetEthUsdPrice(&_AavePriceOracle.CallOpts)
}

// SetAssetPrice is a paid mutator transaction binding the contract method 0x51323f72.
//
// Solidity: function setAssetPrice(address _asset, uint256 _price) returns()
func (_AavePriceOracle *AavePriceOracleTransactor) SetAssetPrice(opts *bind.TransactOpts, _asset common.Address, _price *big.Int) (*types.Transaction, error) {
	return _AavePriceOracle.contract.Transact(opts, "setAssetPrice", _asset, _price)
}

// SetAssetPrice is a paid mutator transaction binding the contract method 0x51323f72.
//
// Solidity: function setAssetPrice(address _asset, uint256 _price) returns()
func (_AavePriceOracle *AavePriceOracleSession) SetAssetPrice(_asset common.Address, _price *big.Int) (*types.Transaction, error) {
	return _AavePriceOracle.Contract.SetAssetPrice(&_AavePriceOracle.TransactOpts, _asset, _price)
}

// SetAssetPrice is a paid mutator transaction binding the contract method 0x51323f72.
//
// Solidity: function setAssetPrice(address _asset, uint256 _price) returns()
func (_AavePriceOracle *AavePriceOracleTransactorSession) SetAssetPrice(_asset common.Address, _price *big.Int) (*types.Transaction, error) {
	return _AavePriceOracle.Contract.SetAssetPrice(&_AavePriceOracle.TransactOpts, _asset, _price)
}

// SetEthUsdPrice is a paid mutator transaction binding the contract method 0xb951883a.
//
// Solidity: function setEthUsdPrice(uint256 _price) returns()
func (_AavePriceOracle *AavePriceOracleTransactor) SetEthUsdPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _AavePriceOracle.contract.Transact(opts, "setEthUsdPrice", _price)
}

// SetEthUsdPrice is a paid mutator transaction binding the contract method 0xb951883a.
//
// Solidity: function setEthUsdPrice(uint256 _price) returns()
func (_AavePriceOracle *AavePriceOracleSession) SetEthUsdPrice(_price *big.Int) (*types.Transaction, error) {
	return _AavePriceOracle.Contract.SetEthUsdPrice(&_AavePriceOracle.TransactOpts, _price)
}

// SetEthUsdPrice is a paid mutator transaction binding the contract method 0xb951883a.
//
// Solidity: function setEthUsdPrice(uint256 _price) returns()
func (_AavePriceOracle *AavePriceOracleTransactorSession) SetEthUsdPrice(_price *big.Int) (*types.Transaction, error) {
	return _AavePriceOracle.Contract.SetEthUsdPrice(&_AavePriceOracle.TransactOpts, _price)
}

// AavePriceOracleAssetPriceUpdatedIterator is returned from FilterAssetPriceUpdated and is used to iterate over the raw logs and unpacked data for AssetPriceUpdated events raised by the AavePriceOracle contract.
type AavePriceOracleAssetPriceUpdatedIterator struct {
	Event *AavePriceOracleAssetPriceUpdated // Event containing the contract specifics and raw log

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
func (it *AavePriceOracleAssetPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePriceOracleAssetPriceUpdated)
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
		it.Event = new(AavePriceOracleAssetPriceUpdated)
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
func (it *AavePriceOracleAssetPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePriceOracleAssetPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePriceOracleAssetPriceUpdated represents a AssetPriceUpdated event raised by the AavePriceOracle contract.
type AavePriceOracleAssetPriceUpdated struct {
	Asset     common.Address
	Price     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAssetPriceUpdated is a free log retrieval operation binding the contract event 0xce6e0b57367bae95ca7198e1172f653ea64a645c16ab586b4cefa9237bfc2d92.
//
// Solidity: event AssetPriceUpdated(address _asset, uint256 _price, uint256 timestamp)
func (_AavePriceOracle *AavePriceOracleFilterer) FilterAssetPriceUpdated(opts *bind.FilterOpts) (*AavePriceOracleAssetPriceUpdatedIterator, error) {

	logs, sub, err := _AavePriceOracle.contract.FilterLogs(opts, "AssetPriceUpdated")
	if err != nil {
		return nil, err
	}
	return &AavePriceOracleAssetPriceUpdatedIterator{contract: _AavePriceOracle.contract, event: "AssetPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetPriceUpdated is a free log subscription operation binding the contract event 0xce6e0b57367bae95ca7198e1172f653ea64a645c16ab586b4cefa9237bfc2d92.
//
// Solidity: event AssetPriceUpdated(address _asset, uint256 _price, uint256 timestamp)
func (_AavePriceOracle *AavePriceOracleFilterer) WatchAssetPriceUpdated(opts *bind.WatchOpts, sink chan<- *AavePriceOracleAssetPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _AavePriceOracle.contract.WatchLogs(opts, "AssetPriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePriceOracleAssetPriceUpdated)
				if err := _AavePriceOracle.contract.UnpackLog(event, "AssetPriceUpdated", log); err != nil {
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

// ParseAssetPriceUpdated is a log parse operation binding the contract event 0xce6e0b57367bae95ca7198e1172f653ea64a645c16ab586b4cefa9237bfc2d92.
//
// Solidity: event AssetPriceUpdated(address _asset, uint256 _price, uint256 timestamp)
func (_AavePriceOracle *AavePriceOracleFilterer) ParseAssetPriceUpdated(log types.Log) (*AavePriceOracleAssetPriceUpdated, error) {
	event := new(AavePriceOracleAssetPriceUpdated)
	if err := _AavePriceOracle.contract.UnpackLog(event, "AssetPriceUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AavePriceOracleEthPriceUpdatedIterator is returned from FilterEthPriceUpdated and is used to iterate over the raw logs and unpacked data for EthPriceUpdated events raised by the AavePriceOracle contract.
type AavePriceOracleEthPriceUpdatedIterator struct {
	Event *AavePriceOracleEthPriceUpdated // Event containing the contract specifics and raw log

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
func (it *AavePriceOracleEthPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AavePriceOracleEthPriceUpdated)
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
		it.Event = new(AavePriceOracleEthPriceUpdated)
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
func (it *AavePriceOracleEthPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AavePriceOracleEthPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AavePriceOracleEthPriceUpdated represents a EthPriceUpdated event raised by the AavePriceOracle contract.
type AavePriceOracleEthPriceUpdated struct {
	Price     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthPriceUpdated is a free log retrieval operation binding the contract event 0xb4f35977939fa8b5ffe552d517a8ff5223046b1fdd3ee0068ae38d1e2b8d0016.
//
// Solidity: event EthPriceUpdated(uint256 _price, uint256 timestamp)
func (_AavePriceOracle *AavePriceOracleFilterer) FilterEthPriceUpdated(opts *bind.FilterOpts) (*AavePriceOracleEthPriceUpdatedIterator, error) {

	logs, sub, err := _AavePriceOracle.contract.FilterLogs(opts, "EthPriceUpdated")
	if err != nil {
		return nil, err
	}
	return &AavePriceOracleEthPriceUpdatedIterator{contract: _AavePriceOracle.contract, event: "EthPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchEthPriceUpdated is a free log subscription operation binding the contract event 0xb4f35977939fa8b5ffe552d517a8ff5223046b1fdd3ee0068ae38d1e2b8d0016.
//
// Solidity: event EthPriceUpdated(uint256 _price, uint256 timestamp)
func (_AavePriceOracle *AavePriceOracleFilterer) WatchEthPriceUpdated(opts *bind.WatchOpts, sink chan<- *AavePriceOracleEthPriceUpdated) (event.Subscription, error) {

	logs, sub, err := _AavePriceOracle.contract.WatchLogs(opts, "EthPriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AavePriceOracleEthPriceUpdated)
				if err := _AavePriceOracle.contract.UnpackLog(event, "EthPriceUpdated", log); err != nil {
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

// ParseEthPriceUpdated is a log parse operation binding the contract event 0xb4f35977939fa8b5ffe552d517a8ff5223046b1fdd3ee0068ae38d1e2b8d0016.
//
// Solidity: event EthPriceUpdated(uint256 _price, uint256 timestamp)
func (_AavePriceOracle *AavePriceOracleFilterer) ParseEthPriceUpdated(log types.Log) (*AavePriceOracleEthPriceUpdated, error) {
	event := new(AavePriceOracleEthPriceUpdated)
	if err := _AavePriceOracle.contract.UnpackLog(event, "EthPriceUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
