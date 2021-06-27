// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package comptroller_v4_storage

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

// ComptrollerV4StorageABI is the input ABI used to generate the binding from.
const ComptrollerV4StorageABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountAssets\",\"outputs\":[{\"internalType\":\"contractCToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"contractCToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowCapGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closeFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compBorrowState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compBorrowerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compSupplierIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compSupplyState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"liquidationIncentiveMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isComped\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pauseGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingComptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seizeGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transferGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ComptrollerV4StorageBin is the compiled bytecode used for deploying new contracts.
var ComptrollerV4StorageBin = "0x608060405234801561001057600080fd5b506106de806100206000396000f3fe608060405234801561001057600080fd5b50600436106101a35760003560e01c80638c57804e116100ee578063bb82aa5e11610097578063dce1544911610071578063dce154491461042d578063dcfbc0c714610459578063e875544614610461578063f851a44014610469576101a3565b8063bb82aa5e146103d1578063ca0af043146103d9578063cc7ebdc414610407576101a3565b8063aa900754116100c8578063aa90075414610393578063ac0b0bb71461039b578063b21be7fd146103a3576101a3565b80638c57804e1461031d5780638e8f294b1461034357806394b2294b1461038b576101a3565b806352d84d1e11610150578063731f0c2b1161012a578063731f0c2b146102e75780637dc0d1d01461030d57806387f7630314610315576101a3565b806352d84d1e146102425780636b79c38d1461025f5780636d154ea5146102ad576101a3565b80632678224711610181578063267822471461020c5780634a584432146102145780634ada90af1461023a576101a3565b80631d7b33d7146101a857806321af4569146101e057806324a3d62214610204575b600080fd5b6101ce600480360360208110156101be57600080fd5b50356001600160a01b0316610471565b60408051918252519081900360200190f35b6101e8610483565b604080516001600160a01b039092168252519081900360200190f35b6101e8610492565b6101e86104a1565b6101ce6004803603602081101561022a57600080fd5b50356001600160a01b03166104b0565b6101ce6104c2565b6101e86004803603602081101561025857600080fd5b50356104c8565b6102856004803603602081101561027557600080fd5b50356001600160a01b03166104ef565b604080516001600160e01b03909316835263ffffffff90911660208301528051918290030190f35b6102d3600480360360208110156102c357600080fd5b50356001600160a01b0316610519565b604080519115158252519081900360200190f35b6102d3600480360360208110156102fd57600080fd5b50356001600160a01b031661052e565b6101e8610543565b6102d3610552565b6102856004803603602081101561033357600080fd5b50356001600160a01b0316610575565b6103696004803603602081101561035957600080fd5b50356001600160a01b031661059f565b6040805193151584526020840192909252151582820152519081900360600190f35b6101ce6105c5565b6101ce6105cb565b6102d36105d1565b6101ce600480360360408110156103b957600080fd5b506001600160a01b03813581169160200135166105f5565b6101e8610612565b6101ce600480360360408110156103ef57600080fd5b506001600160a01b0381358116916020013516610621565b6101ce6004803603602081101561041d57600080fd5b50356001600160a01b031661063e565b6101e86004803603604081101561044357600080fd5b506001600160a01b038135169060200135610650565b6101e8610685565b6101ce610694565b6101e861069a565b600f6020526000908152604090205481565b6015546001600160a01b031681565b600a546001600160a01b031681565b6001546001600160a01b031681565b60166020526000908152604090205481565b60065481565b600d81815481106104d557fe5b6000918252602090912001546001600160a01b0316905081565b6010602052600090815260409020546001600160e01b03811690600160e01b900463ffffffff1682565b600c6020526000908152604090205460ff1681565b600b6020526000908152604090205460ff1681565b6004546001600160a01b031681565b600a54760100000000000000000000000000000000000000000000900460ff1681565b6011602052600090815260409020546001600160e01b03811690600160e01b900463ffffffff1682565b60096020526000908152604090208054600182015460039092015460ff91821692911683565b60075481565b600e5481565b600a5477010000000000000000000000000000000000000000000000900460ff1681565b601260209081526000928352604080842090915290825290205481565b6002546001600160a01b031681565b601360209081526000928352604080842090915290825290205481565b60146020526000908152604090205481565b6008602052816000526040600020818154811061066957fe5b6000918252602090912001546001600160a01b03169150829050565b6003546001600160a01b031681565b60055481565b6000546001600160a01b03168156fea265627a7a72315820ea50f053e004b2cd51449f8f5d332ae874c405605f9dfb2ed83f5248862ecc4164736f6c63430005100032"

// DeployComptrollerV4Storage deploys a new Ethereum contract, binding an instance of ComptrollerV4Storage to it.
func DeployComptrollerV4Storage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ComptrollerV4Storage, error) {
	parsed, err := abi.JSON(strings.NewReader(ComptrollerV4StorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ComptrollerV4StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ComptrollerV4Storage{ComptrollerV4StorageCaller: ComptrollerV4StorageCaller{contract: contract}, ComptrollerV4StorageTransactor: ComptrollerV4StorageTransactor{contract: contract}, ComptrollerV4StorageFilterer: ComptrollerV4StorageFilterer{contract: contract}}, nil
}

// ComptrollerV4Storage is an auto generated Go binding around an Ethereum contract.
type ComptrollerV4Storage struct {
	ComptrollerV4StorageCaller     // Read-only binding to the contract
	ComptrollerV4StorageTransactor // Write-only binding to the contract
	ComptrollerV4StorageFilterer   // Log filterer for contract events
}

// ComptrollerV4StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComptrollerV4StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerV4StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComptrollerV4StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerV4StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComptrollerV4StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerV4StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComptrollerV4StorageSession struct {
	Contract     *ComptrollerV4Storage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ComptrollerV4StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComptrollerV4StorageCallerSession struct {
	Contract *ComptrollerV4StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ComptrollerV4StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComptrollerV4StorageTransactorSession struct {
	Contract     *ComptrollerV4StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ComptrollerV4StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComptrollerV4StorageRaw struct {
	Contract *ComptrollerV4Storage // Generic contract binding to access the raw methods on
}

// ComptrollerV4StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComptrollerV4StorageCallerRaw struct {
	Contract *ComptrollerV4StorageCaller // Generic read-only contract binding to access the raw methods on
}

// ComptrollerV4StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComptrollerV4StorageTransactorRaw struct {
	Contract *ComptrollerV4StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComptrollerV4Storage creates a new instance of ComptrollerV4Storage, bound to a specific deployed contract.
func NewComptrollerV4Storage(address common.Address, backend bind.ContractBackend) (*ComptrollerV4Storage, error) {
	contract, err := bindComptrollerV4Storage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ComptrollerV4Storage{ComptrollerV4StorageCaller: ComptrollerV4StorageCaller{contract: contract}, ComptrollerV4StorageTransactor: ComptrollerV4StorageTransactor{contract: contract}, ComptrollerV4StorageFilterer: ComptrollerV4StorageFilterer{contract: contract}}, nil
}

// NewComptrollerV4StorageCaller creates a new read-only instance of ComptrollerV4Storage, bound to a specific deployed contract.
func NewComptrollerV4StorageCaller(address common.Address, caller bind.ContractCaller) (*ComptrollerV4StorageCaller, error) {
	contract, err := bindComptrollerV4Storage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComptrollerV4StorageCaller{contract: contract}, nil
}

// NewComptrollerV4StorageTransactor creates a new write-only instance of ComptrollerV4Storage, bound to a specific deployed contract.
func NewComptrollerV4StorageTransactor(address common.Address, transactor bind.ContractTransactor) (*ComptrollerV4StorageTransactor, error) {
	contract, err := bindComptrollerV4Storage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComptrollerV4StorageTransactor{contract: contract}, nil
}

// NewComptrollerV4StorageFilterer creates a new log filterer instance of ComptrollerV4Storage, bound to a specific deployed contract.
func NewComptrollerV4StorageFilterer(address common.Address, filterer bind.ContractFilterer) (*ComptrollerV4StorageFilterer, error) {
	contract, err := bindComptrollerV4Storage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComptrollerV4StorageFilterer{contract: contract}, nil
}

// bindComptrollerV4Storage binds a generic wrapper to an already deployed contract.
func bindComptrollerV4Storage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ComptrollerV4StorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComptrollerV4Storage *ComptrollerV4StorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ComptrollerV4Storage.Contract.ComptrollerV4StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComptrollerV4Storage *ComptrollerV4StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComptrollerV4Storage.Contract.ComptrollerV4StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComptrollerV4Storage *ComptrollerV4StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComptrollerV4Storage.Contract.ComptrollerV4StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ComptrollerV4Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComptrollerV4Storage *ComptrollerV4StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComptrollerV4Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComptrollerV4Storage *ComptrollerV4StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComptrollerV4Storage.Contract.contract.Transact(opts, method, params...)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) AccountAssets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "accountAssets", arg0, arg1)
	return *ret0, err
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _ComptrollerV4Storage.Contract.AccountAssets(&_ComptrollerV4Storage.CallOpts, arg0, arg1)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _ComptrollerV4Storage.Contract.AccountAssets(&_ComptrollerV4Storage.CallOpts, arg0, arg1)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) Admin() (common.Address, error) {
	return _ComptrollerV4Storage.Contract.Admin(&_ComptrollerV4Storage.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) Admin() (common.Address, error) {
	return _ComptrollerV4Storage.Contract.Admin(&_ComptrollerV4Storage.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) AllMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "allMarkets", arg0)
	return *ret0, err
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _ComptrollerV4Storage.Contract.AllMarkets(&_ComptrollerV4Storage.CallOpts, arg0)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _ComptrollerV4Storage.Contract.AllMarkets(&_ComptrollerV4Storage.CallOpts, arg0)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) BorrowCapGuardian(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "borrowCapGuardian")
	return *ret0, err
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) BorrowCapGuardian() (common.Address, error) {
	return _ComptrollerV4Storage.Contract.BorrowCapGuardian(&_ComptrollerV4Storage.CallOpts)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) BorrowCapGuardian() (common.Address, error) {
	return _ComptrollerV4Storage.Contract.BorrowCapGuardian(&_ComptrollerV4Storage.CallOpts)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) BorrowCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "borrowCaps", arg0)
	return *ret0, err
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _ComptrollerV4Storage.Contract.BorrowCaps(&_ComptrollerV4Storage.CallOpts, arg0)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _ComptrollerV4Storage.Contract.BorrowCaps(&_ComptrollerV4Storage.CallOpts, arg0)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) BorrowGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "borrowGuardianPaused", arg0)
	return *ret0, err
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) BorrowGuardianPaused(arg0 common.Address) (bool, error) {
	return _ComptrollerV4Storage.Contract.BorrowGuardianPaused(&_ComptrollerV4Storage.CallOpts, arg0)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) BorrowGuardianPaused(arg0 common.Address) (bool, error) {
	return _ComptrollerV4Storage.Contract.BorrowGuardianPaused(&_ComptrollerV4Storage.CallOpts, arg0)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) CloseFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "closeFactorMantissa")
	return *ret0, err
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) CloseFactorMantissa() (*big.Int, error) {
	return _ComptrollerV4Storage.Contract.CloseFactorMantissa(&_ComptrollerV4Storage.CallOpts)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) CloseFactorMantissa() (*big.Int, error) {
	return _ComptrollerV4Storage.Contract.CloseFactorMantissa(&_ComptrollerV4Storage.CallOpts)
}

// CompAccrued is a free data retrieval call binding the contract method 0xcc7ebdc4.
//
// Solidity: function compAccrued(address ) view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) CompAccrued(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "compAccrued", arg0)
	return *ret0, err
}

// CompAccrued is a free data retrieval call binding the contract method 0xcc7ebdc4.
//
// Solidity: function compAccrued(address ) view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) CompAccrued(arg0 common.Address) (*big.Int, error) {
	return _ComptrollerV4Storage.Contract.CompAccrued(&_ComptrollerV4Storage.CallOpts, arg0)
}

// CompAccrued is a free data retrieval call binding the contract method 0xcc7ebdc4.
//
// Solidity: function compAccrued(address ) view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) CompAccrued(arg0 common.Address) (*big.Int, error) {
	return _ComptrollerV4Storage.Contract.CompAccrued(&_ComptrollerV4Storage.CallOpts, arg0)
}

// CompBorrowState is a free data retrieval call binding the contract method 0x8c57804e.
//
// Solidity: function compBorrowState(address ) view returns(uint224 index, uint32 block)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) CompBorrowState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	ret := new(struct {
		Index *big.Int
		Block uint32
	})
	out := ret
	err := _ComptrollerV4Storage.contract.Call(opts, out, "compBorrowState", arg0)
	return *ret, err
}

// CompBorrowState is a free data retrieval call binding the contract method 0x8c57804e.
//
// Solidity: function compBorrowState(address ) view returns(uint224 index, uint32 block)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) CompBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _ComptrollerV4Storage.Contract.CompBorrowState(&_ComptrollerV4Storage.CallOpts, arg0)
}

// CompBorrowState is a free data retrieval call binding the contract method 0x8c57804e.
//
// Solidity: function compBorrowState(address ) view returns(uint224 index, uint32 block)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) CompBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _ComptrollerV4Storage.Contract.CompBorrowState(&_ComptrollerV4Storage.CallOpts, arg0)
}

// CompBorrowerIndex is a free data retrieval call binding the contract method 0xca0af043.
//
// Solidity: function compBorrowerIndex(address , address ) view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) CompBorrowerIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "compBorrowerIndex", arg0, arg1)
	return *ret0, err
}

// CompBorrowerIndex is a free data retrieval call binding the contract method 0xca0af043.
//
// Solidity: function compBorrowerIndex(address , address ) view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) CompBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ComptrollerV4Storage.Contract.CompBorrowerIndex(&_ComptrollerV4Storage.CallOpts, arg0, arg1)
}

// CompBorrowerIndex is a free data retrieval call binding the contract method 0xca0af043.
//
// Solidity: function compBorrowerIndex(address , address ) view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) CompBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ComptrollerV4Storage.Contract.CompBorrowerIndex(&_ComptrollerV4Storage.CallOpts, arg0, arg1)
}

// CompRate is a free data retrieval call binding the contract method 0xaa900754.
//
// Solidity: function compRate() view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) CompRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "compRate")
	return *ret0, err
}

// CompRate is a free data retrieval call binding the contract method 0xaa900754.
//
// Solidity: function compRate() view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) CompRate() (*big.Int, error) {
	return _ComptrollerV4Storage.Contract.CompRate(&_ComptrollerV4Storage.CallOpts)
}

// CompRate is a free data retrieval call binding the contract method 0xaa900754.
//
// Solidity: function compRate() view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) CompRate() (*big.Int, error) {
	return _ComptrollerV4Storage.Contract.CompRate(&_ComptrollerV4Storage.CallOpts)
}

// CompSpeeds is a free data retrieval call binding the contract method 0x1d7b33d7.
//
// Solidity: function compSpeeds(address ) view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) CompSpeeds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "compSpeeds", arg0)
	return *ret0, err
}

// CompSpeeds is a free data retrieval call binding the contract method 0x1d7b33d7.
//
// Solidity: function compSpeeds(address ) view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) CompSpeeds(arg0 common.Address) (*big.Int, error) {
	return _ComptrollerV4Storage.Contract.CompSpeeds(&_ComptrollerV4Storage.CallOpts, arg0)
}

// CompSpeeds is a free data retrieval call binding the contract method 0x1d7b33d7.
//
// Solidity: function compSpeeds(address ) view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) CompSpeeds(arg0 common.Address) (*big.Int, error) {
	return _ComptrollerV4Storage.Contract.CompSpeeds(&_ComptrollerV4Storage.CallOpts, arg0)
}

// CompSupplierIndex is a free data retrieval call binding the contract method 0xb21be7fd.
//
// Solidity: function compSupplierIndex(address , address ) view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) CompSupplierIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "compSupplierIndex", arg0, arg1)
	return *ret0, err
}

// CompSupplierIndex is a free data retrieval call binding the contract method 0xb21be7fd.
//
// Solidity: function compSupplierIndex(address , address ) view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) CompSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ComptrollerV4Storage.Contract.CompSupplierIndex(&_ComptrollerV4Storage.CallOpts, arg0, arg1)
}

// CompSupplierIndex is a free data retrieval call binding the contract method 0xb21be7fd.
//
// Solidity: function compSupplierIndex(address , address ) view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) CompSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ComptrollerV4Storage.Contract.CompSupplierIndex(&_ComptrollerV4Storage.CallOpts, arg0, arg1)
}

// CompSupplyState is a free data retrieval call binding the contract method 0x6b79c38d.
//
// Solidity: function compSupplyState(address ) view returns(uint224 index, uint32 block)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) CompSupplyState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	ret := new(struct {
		Index *big.Int
		Block uint32
	})
	out := ret
	err := _ComptrollerV4Storage.contract.Call(opts, out, "compSupplyState", arg0)
	return *ret, err
}

// CompSupplyState is a free data retrieval call binding the contract method 0x6b79c38d.
//
// Solidity: function compSupplyState(address ) view returns(uint224 index, uint32 block)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) CompSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _ComptrollerV4Storage.Contract.CompSupplyState(&_ComptrollerV4Storage.CallOpts, arg0)
}

// CompSupplyState is a free data retrieval call binding the contract method 0x6b79c38d.
//
// Solidity: function compSupplyState(address ) view returns(uint224 index, uint32 block)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) CompSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _ComptrollerV4Storage.Contract.CompSupplyState(&_ComptrollerV4Storage.CallOpts, arg0)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) ComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "comptrollerImplementation")
	return *ret0, err
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) ComptrollerImplementation() (common.Address, error) {
	return _ComptrollerV4Storage.Contract.ComptrollerImplementation(&_ComptrollerV4Storage.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) ComptrollerImplementation() (common.Address, error) {
	return _ComptrollerV4Storage.Contract.ComptrollerImplementation(&_ComptrollerV4Storage.CallOpts)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) LiquidationIncentiveMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "liquidationIncentiveMantissa")
	return *ret0, err
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _ComptrollerV4Storage.Contract.LiquidationIncentiveMantissa(&_ComptrollerV4Storage.CallOpts)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _ComptrollerV4Storage.Contract.LiquidationIncentiveMantissa(&_ComptrollerV4Storage.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isComped)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) Markets(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsComped                 bool
}, error) {
	ret := new(struct {
		IsListed                 bool
		CollateralFactorMantissa *big.Int
		IsComped                 bool
	})
	out := ret
	err := _ComptrollerV4Storage.contract.Call(opts, out, "markets", arg0)
	return *ret, err
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isComped)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsComped                 bool
}, error) {
	return _ComptrollerV4Storage.Contract.Markets(&_ComptrollerV4Storage.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isComped)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsComped                 bool
}, error) {
	return _ComptrollerV4Storage.Contract.Markets(&_ComptrollerV4Storage.CallOpts, arg0)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) MaxAssets(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "maxAssets")
	return *ret0, err
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) MaxAssets() (*big.Int, error) {
	return _ComptrollerV4Storage.Contract.MaxAssets(&_ComptrollerV4Storage.CallOpts)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) MaxAssets() (*big.Int, error) {
	return _ComptrollerV4Storage.Contract.MaxAssets(&_ComptrollerV4Storage.CallOpts)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) MintGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "mintGuardianPaused", arg0)
	return *ret0, err
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _ComptrollerV4Storage.Contract.MintGuardianPaused(&_ComptrollerV4Storage.CallOpts, arg0)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _ComptrollerV4Storage.Contract.MintGuardianPaused(&_ComptrollerV4Storage.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) Oracle() (common.Address, error) {
	return _ComptrollerV4Storage.Contract.Oracle(&_ComptrollerV4Storage.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) Oracle() (common.Address, error) {
	return _ComptrollerV4Storage.Contract.Oracle(&_ComptrollerV4Storage.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) PauseGuardian(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "pauseGuardian")
	return *ret0, err
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) PauseGuardian() (common.Address, error) {
	return _ComptrollerV4Storage.Contract.PauseGuardian(&_ComptrollerV4Storage.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) PauseGuardian() (common.Address, error) {
	return _ComptrollerV4Storage.Contract.PauseGuardian(&_ComptrollerV4Storage.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "pendingAdmin")
	return *ret0, err
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) PendingAdmin() (common.Address, error) {
	return _ComptrollerV4Storage.Contract.PendingAdmin(&_ComptrollerV4Storage.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) PendingAdmin() (common.Address, error) {
	return _ComptrollerV4Storage.Contract.PendingAdmin(&_ComptrollerV4Storage.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) PendingComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "pendingComptrollerImplementation")
	return *ret0, err
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) PendingComptrollerImplementation() (common.Address, error) {
	return _ComptrollerV4Storage.Contract.PendingComptrollerImplementation(&_ComptrollerV4Storage.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _ComptrollerV4Storage.Contract.PendingComptrollerImplementation(&_ComptrollerV4Storage.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) SeizeGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "seizeGuardianPaused")
	return *ret0, err
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) SeizeGuardianPaused() (bool, error) {
	return _ComptrollerV4Storage.Contract.SeizeGuardianPaused(&_ComptrollerV4Storage.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) SeizeGuardianPaused() (bool, error) {
	return _ComptrollerV4Storage.Contract.SeizeGuardianPaused(&_ComptrollerV4Storage.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_ComptrollerV4Storage *ComptrollerV4StorageCaller) TransferGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ComptrollerV4Storage.contract.Call(opts, out, "transferGuardianPaused")
	return *ret0, err
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_ComptrollerV4Storage *ComptrollerV4StorageSession) TransferGuardianPaused() (bool, error) {
	return _ComptrollerV4Storage.Contract.TransferGuardianPaused(&_ComptrollerV4Storage.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_ComptrollerV4Storage *ComptrollerV4StorageCallerSession) TransferGuardianPaused() (bool, error) {
	return _ComptrollerV4Storage.Contract.TransferGuardianPaused(&_ComptrollerV4Storage.CallOpts)
}
