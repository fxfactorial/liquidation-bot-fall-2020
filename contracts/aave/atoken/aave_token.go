// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package atoken

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

// AtokenABI is the input ABI used to generate the binding from.
const AtokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"underlyingAssetAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"UINT_MAX_VALUE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_addressesProvider\",\"type\":\"address\"},{\"name\":\"_underlyingAsset\",\"type\":\"address\"},{\"name\":\"_underlyingAssetDecimals\",\"type\":\"uint8\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_fromBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_fromIndex\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_fromBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_fromIndex\",\"type\":\"uint256\"}],\"name\":\"MintOnDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_fromBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_fromIndex\",\"type\":\"uint256\"}],\"name\":\"BurnOnLiquidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_fromBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_toBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_fromIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_toIndex\",\"type\":\"uint256\"}],\"name\":\"BalanceTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_redirectedBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_fromBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_fromIndex\",\"type\":\"uint256\"}],\"name\":\"InterestStreamRedirected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_targetAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_targetBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_targetIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_redirectedBalanceAdded\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_redirectedBalanceRemoved\",\"type\":\"uint256\"}],\"name\":\"RedirectedBalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"InterestRedirectionAllowanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"redirectInterestStream\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"redirectInterestStreamOf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"allowInterestRedirectionTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mintOnDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burnOnLiquidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferOnLiquidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"principalBalanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"isTransferAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getInterestRedirectionAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getRedirectedBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AtokenBin is the compiled bytecode used for deploying new contracts.
var AtokenBin = "0x60806040523480156200001157600080fd5b50604051620026cf380380620026cf833981018060405260a08110156200003757600080fd5b81516020830151604084015160608501805193959294919391830192916401000000008111156200006757600080fd5b820160208101848111156200007b57600080fd5b81516401000000008111828201871017156200009657600080fd5b50509291906020018051640100000000811115620000b357600080fd5b82016020810184811115620000c757600080fd5b8151640100000000811182820187101715620000e257600080fd5b505084519093508492508391508590620001049060039060208601906200035c565b5081516200011a9060049060208501906200035c565b506005805460ff191660ff929092169190911790555050600a80546001600160a01b0319166001600160a01b038781169190911791829055604080517fed6ff7600000000000000000000000000000000000000000000000000000000081529051929091169163ed6ff76091600480820192602092909190829003018186803b158015620001a757600080fd5b505afa158015620001bc573d6000803e3d6000fd5b505050506040513d6020811015620001d357600080fd5b5051600b80546001600160a01b0319166001600160a01b03928316179055600a54604080517f0261bf8b00000000000000000000000000000000000000000000000000000000815290519190921691630261bf8b916004808301926020929190829003018186803b1580156200024857600080fd5b505afa1580156200025d573d6000803e3d6000fd5b505050506040513d60208110156200027457600080fd5b5051600c80546001600160a01b0319166001600160a01b03928316179055600a54604080517f2f58b80d00000000000000000000000000000000000000000000000000000000815290519190921691632f58b80d916004808301926020929190829003018186803b158015620002e957600080fd5b505afa158015620002fe573d6000803e3d6000fd5b505050506040513d60208110156200031557600080fd5b5051600d80546001600160a01b0319166001600160a01b0392831617905560058054610100600160a81b031916610100969092169590950217909355506200040192505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200039f57805160ff1916838001178555620003cf565b82800160010185558215620003cf579182015b82811115620003cf578251825591602001919060010190620003b2565b50620003dd929150620003e1565b5090565b620003fe91905b80821115620003dd5760008155600101620003e8565b90565b6122be80620004116000396000f3fe608060405234801561001057600080fd5b50600436106101a35760003560e01c80635eae177c116100ee578063a9059cbb11610097578063db006a7511610071578063db006a7514610523578063dd62ed3e14610540578063ee9907a41461056e578063f866c31914610594576101a3565b8063a9059cbb146104c9578063c634dfaa146104f5578063d0fc81d21461051b576101a3565b806394362e8b116100c857806394362e8b1461046957806395d89b4114610495578063a457c2d71461049d576101a3565b80635eae177c1461040f57806370a082311461043b57806389d1a0fc14610461576101a3565b806323b872dd11610150578063395093511161012a57806339509351146103755780633edb7cb8146103a1578063445e8010146103cd576101a3565b806323b872dd146102f3578063313ce56714610329578063325a9b1314610347576101a3565b806312c87c2d1161018157806312c87c2d1461028d57806318160ddd146102b35780631d51e7cf146102cd576101a3565b806306fdde03146101a8578063095ea7b3146102255780630e49072d14610265575b600080fd5b6101b06105ca565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101ea5781810151838201526020016101d2565b50505050905090810190601f1680156102175780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102516004803603604081101561023b57600080fd5b506001600160a01b038135169060200135610661565b604080519115158252519081900360200190f35b61028b6004803603602081101561027b57600080fd5b50356001600160a01b031661067f565b005b61028b600480360360208110156102a357600080fd5b50356001600160a01b031661068c565b6102bb61073b565b60408051918252519081900360200190f35b6102bb600480360360208110156102e357600080fd5b50356001600160a01b0316610818565b6102516004803603606081101561030957600080fd5b506001600160a01b03813581169160208101359091169060400135610837565b6103316108c4565b6040805160ff9092168252519081900360200190f35b61028b6004803603604081101561035d57600080fd5b506001600160a01b03813581169160200135166108cd565b6102516004803603604081101561038b57600080fd5b506001600160a01b038135169060200135610936565b61028b600480360360408110156103b757600080fd5b506001600160a01b03813516906020013561098a565b6103f3600480360360208110156103e357600080fd5b50356001600160a01b0316610a84565b604080516001600160a01b039092168252519081900360200190f35b6102516004803603604081101561042557600080fd5b506001600160a01b038135169060200135610aa2565b6102bb6004803603602081101561045157600080fd5b50356001600160a01b0316610b56565b6103f3610c0b565b61028b6004803603604081101561047f57600080fd5b506001600160a01b038135169060200135610c1f565b6101b0610cf7565b610251600480360360408110156104b357600080fd5b506001600160a01b038135169060200135610d58565b610251600480360360408110156104df57600080fd5b506001600160a01b038135169060200135610dc6565b6102bb6004803603602081101561050b57600080fd5b50356001600160a01b0316610dda565b6102bb610de5565b61028b6004803603602081101561053957600080fd5b5035610deb565b6102bb6004803603604081101561055657600080fd5b506001600160a01b0381358116916020013516611031565b6102bb6004803603602081101561058457600080fd5b50356001600160a01b031661105c565b61028b600480360360608110156105aa57600080fd5b506001600160a01b03813581169160208101359091169060400135611077565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156106565780601f1061062b57610100808354040283529160200191610656565b820191906000526020600020905b81548152906001019060200180831161063957829003601f168201915b505050505090505b90565b600061067561066e6110d3565b84846110d7565b5060015b92915050565b61068933826111cd565b50565b6001600160a01b0381163314156106d757604051600160e51b62461bcd02815260040180806020018281038252602581526020018061208e6025913960400191505060405180910390fd5b33600081815260096020526040808220805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03861690811790915590519092917fc2d6a42a9d5273283f73009a07aacfb043f2f91173a8d9d33b504afe898db08b91a350565b6000806107466113de565b905080151561075957600091505061065e565b600b54600554604080517fd15e00530000000000000000000000000000000000000000000000000000000081526001600160a01b036101009093048316600482015290516108129361080d93169163d15e0053916024808301926020929190829003018186803b1580156107cc57600080fd5b505afa1580156107e0573d6000803e3d6000fd5b505050506040513d60208110156107f657600080fd5b5051610801846113e4565b9063ffffffff6113fa16565b611445565b91505090565b6001600160a01b0381166000908152600860205260409020545b919050565b600061084484848461145e565b6108ba846108506110d3565b6108b5856040518060600160405280602881526020016120d4602891396001600160a01b038a1660009081526001602052604081209061088e6110d3565b6001600160a01b03168152602081019190915260400160002054919063ffffffff6114d216565b6110d7565b5060019392505050565b60055460ff1690565b6001600160a01b0382811660009081526009602052604090205416331461092857604051600160e51b62461bcd02815260040180806020018281038252603a815260200180612259603a913960400191505060405180910390fd5b61093282826111cd565b5050565b60006106756109436110d3565b846108b585600160006109546110d3565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff61156c16565b600c546001600160a01b031633146109d657604051600160e51b62461bcd0281526004018080602001828103825260328152602001806122026032913960400191505060405180910390fd5b60008060006109e4856115c9565b935093509350506109f68583866116d3565b610a00858561180c565b6000610a12848663ffffffff61190d16565b1515610a2457610a218661194f565b90505b856001600160a01b03167f90e5d3d68ec162d9c7de393037a3ede04dd44f68e051bf2ace4a73c299dbc4db868584610a5c5785610a5f565b60005b60408051938452602084019290925282820152519081900360600190a2505050505050565b6001600160a01b039081166000908152600760205260409020541690565b600d54600554604080517f76e9d6150000000000000000000000000000000000000000000000000000000081526101009092046001600160a01b039081166004840152858116602484015260448301859052905160009391909116916376e9d615916064808301926020929190829003018186803b158015610b2357600080fd5b505afa158015610b37573d6000803e3d6000fd5b505050506040513d6020811015610b4d57600080fd5b50519392505050565b600080610b6283611a10565b6001600160a01b03841660009081526008602052604090205490915081158015610b8a575080155b15610b9a57600092505050610832565b6001600160a01b03848116600090815260076020526040902054161515610bed57610be481610bd886610bd3868463ffffffff61156c16565b611a2b565b9063ffffffff61190d16565b92505050610832565b610be4610bfe82610bd88785611a2b565b839063ffffffff61156c16565b60055461010090046001600160a01b031681565b600c546001600160a01b03163314610c6b57604051600160e51b62461bcd0281526004018080602001828103825260328152602001806122026032913960400191505060405180910390fd5b600080610c77846115c9565b935093505050610c9b84610c94858561156c90919063ffffffff16565b60006116d3565b610ca58484611af9565b604080518481526020810184905280820183905290516001600160a01b038616917fbe7799898ca2d813ff902b487c1b434ab45b47edd8f38b77ad5e99aae8341b7a919081900360600190a250505050565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156106565780601f1061062b57610100808354040283529160200191610656565b6000610675610d656110d3565b846108b5856040518060600160405280602581526020016122346025913960016000610d8f6110d3565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff6114d216565b6000610675610dd36110d3565b848461145e565b600061067982611a10565b60001981565b801515610e425760408051600160e51b62461bcd02815260206004820181905260248201527f416d6f756e7420746f2072656465656d206e6565647320746f206265203e2030604482015290519081900360640190fd5b6000806000610e50336115c9565b91955093509150849050600019811415610e675750825b83811115610ea957604051600160e51b62461bcd0281526004018080602001828103825260328152602001806120146032913960400191505060405180910390fd5b610eb33382610aa2565b1515610f095760408051600160e51b62461bcd02815260206004820152601b60248201527f5472616e736665722063616e6e6f7420626520616c6c6f7765642e0000000000604482015290519081900360640190fd5b610f143384836116d3565b610f1e338261180c565b6000610f30858363ffffffff61190d16565b1515610f4257610f3f3361194f565b90505b600c546005546001600160a01b0391821691639895e3d8916101009004163385610f728a8263ffffffff61190d16565b6040518563ffffffff1660e01b815260040180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b03168152602001838152602001828152602001945050505050600060405180830381600087803b158015610fe157600080fd5b505af1158015610ff5573d6000803e3d6000fd5b50505050336001600160a01b03167fbd5034ffbd47e4e72a94baa2cdb74c6fad73cb3bcdc13036b72ec8306f5a7646838684610a5c5786610a5f565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b031660009081526006602052604090205490565b600c546001600160a01b031633146110c357604051600160e51b62461bcd0281526004018080602001828103825260328152602001806122026032913960400191505060405180910390fd5b6110ce838383611bee565b505050565b3390565b6001600160a01b038316151561112157604051600160e51b62461bcd0281526004018080602001828103825260248152602001806121de6024913960400191505060405180910390fd5b6001600160a01b038216151561116b57604051600160e51b62461bcd0281526004018080602001828103825260228152602001806120466022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b03808316600090815260076020526040902054811690821681141561122d57604051600160e51b62461bcd02815260040180806020018281038252602a8152602001806120fc602a913960400191505060405180910390fd5b60008060008061123c876115c9565b935093509350935060008311151561128857604051600160e51b62461bcd0281526004018080602001828103825260428152602001806121266042913960600191505060405180910390fd5b6001600160a01b038516156112a3576112a3876000866116d3565b866001600160a01b0316866001600160a01b0316141561133c576001600160a01b0387166000818152600760209081526040808320805473ffffffffffffffffffffffffffffffffffffffff191690558051878152918201869052818101859052519192917f5e3cad45b1fe24159d1cb39788d82d0f69cc15770aa96fba1d3d1a73487355949181900360600190a35050505050610932565b6001600160a01b038781166000908152600760205260408120805473ffffffffffffffffffffffffffffffffffffffff19169289169290921790915561138590889085906116d3565b604080518481526020810184905280820183905290516001600160a01b0380891692908a16917f5e3cad45b1fe24159d1cb39788d82d0f69cc15770aa96fba1d3d1a73487355949181900360600190a350505050505050565b60025490565b600061067982633b9aca0063ffffffff611d2116565b600061143e6b033b2e3c9fd0803ce800000061143261141f868663ffffffff611d2116565b6b019d971e4fe8401e740000009061156c565b9063ffffffff611d8116565b9392505050565b6000631dcd650061143e633b9aca00611432838661156c565b828161146a8282610aa2565b15156114c05760408051600160e51b62461bcd02815260206004820152601b60248201527f5472616e736665722063616e6e6f7420626520616c6c6f7765642e0000000000604482015290519081900360640190fd5b6114cb858585611bee565b5050505050565b6000818484111561156457604051600160e51b62461bcd0281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611529578181015183820152602001611511565b50505050905090810190601f1680156115565780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b60008282018381101561143e5760408051600160e51b62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b60008060008060006115da86611a10565b905060006115eb82610bd889610b56565b90506115f78782611af9565b600b54600554604080517fd15e00530000000000000000000000000000000000000000000000000000000081526101009092046001600160a01b0390811660048401529051600093919091169163d15e0053916024808301926020929190829003018186803b15801561166957600080fd5b505afa15801561167d573d6000803e3d6000fd5b505050506040513d602081101561169357600080fd5b50516001600160a01b03891660009081526006602052604090208190559050826116c3818463ffffffff61156c16565b9099909850919650945092505050565b6001600160a01b03808416600090815260076020526040902054168015156116fb57506110ce565b600080611707836115c9565b6001600160a01b038716600090815260086020526040902054919550935061173d9250869150610bd8908863ffffffff61156c16565b6001600160a01b038085166000908152600860209081526040808320949094556007905291909120541680156117b0576001600160a01b038116600090815260086020526040902054611796908463ffffffff61156c16565b6001600160a01b0382166000908152600860205260409020555b60408051848152602081018490528082018890526060810187905290516001600160a01b038616917f70ff8cf632603e2bfd1afb7e4061acce53d95356b1be9726b99fa22ba733b01f919081900360800190a250505050505050565b6001600160a01b038216151561185657604051600160e51b62461bcd0281526004018080602001828103825260218152602001806121986021913960400191505060405180910390fd5b61189981604051806060016040528060228152602001611ff2602291396001600160a01b038516600090815260208190526040902054919063ffffffff6114d216565b6001600160a01b0383166000908152602081905260409020556002546118c5908263ffffffff61190d16565b6002556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b600061143e83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506114d2565b6001600160a01b0381166000818152600760209081526040808320805473ffffffffffffffffffffffffffffffffffffffff191690558051838152918201839052818101839052519192839290917f5e3cad45b1fe24159d1cb39788d82d0f69cc15770aa96fba1d3d1a7348735594919081900360600190a36001600160a01b0382166000908152600860205260409020541515611a0857506001600160a01b0381166000908152600660205260408120556001610832565b506000610832565b6001600160a01b031660009081526020819052604090205490565b6001600160a01b03808316600090815260066020908152604080832054600b5460055483517fd15e0053000000000000000000000000000000000000000000000000000000008152610100909104871660048201529251949561143e9561080d959394611aed94939092169263d15e00539260248082019391829003018186803b158015611ab857600080fd5b505afa158015611acc573d6000803e3d6000fd5b505050506040513d6020811015611ae257600080fd5b5051610801876113e4565b9063ffffffff611dc316565b6001600160a01b0382161515611b595760408051600160e51b62461bcd02815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b600254611b6c908263ffffffff61156c16565b6002556001600160a01b038216600090815260208190526040902054611b98908263ffffffff61156c16565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b801515611c2f57604051600160e51b62461bcd0281526004018080602001828103825260308152602001806121686030913960400191505060405180910390fd5b6000806000611c3d866115c9565b93509350935050600080611c50876115c9565b935093505050611c618885886116d3565b611c7587610c94848963ffffffff61156c16565b611c80888888611dfb565b6000611c92868863ffffffff61190d16565b1515611ca457611ca18961194f565b90505b876001600160a01b0316896001600160a01b03167f89a178eaa27e0cd201bd795ca8ff716ac0df9618494510ca79771cfc66ffcde889888786611ce75789611cea565b60005b60408051948552602085019390935283830191909152606083015260808201879052519081900360a00190a3505050505050505050565b6000821515611d3257506000610679565b828202828482811515611d4157fe5b041461143e57604051600160e51b62461bcd0281526004018080602001828103825260218152602001806120b36021913960400191505060405180910390fd5b600061143e83836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250611f61565b600060028204611df383611432611de6876b033b2e3c9fd0803ce8000000611d21565b849063ffffffff61156c16565b949350505050565b6001600160a01b0383161515611e4557604051600160e51b62461bcd0281526004018080602001828103825260258152602001806121b96025913960400191505060405180910390fd5b6001600160a01b0382161515611e8f57604051600160e51b62461bcd028152600401808060200182810382526023815260200180611fcf6023913960400191505060405180910390fd5b611ed281604051806060016040528060268152602001612068602691396001600160a01b038616600090815260208190526040902054919063ffffffff6114d216565b6001600160a01b038085166000908152602081905260408082209390935590841681522054611f07908263ffffffff61156c16565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600081831515611fb657604051600160e51b62461bcd02815260040180806020018281038252838181518152602001915080519060200190808383600083811015611529578181015183820152602001611511565b5060008385811515611fc457fe5b049594505050505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e6365557365722063616e6e6f742072656465656d206d6f7265207468616e2074686520617661696c61626c652062616c616e636545524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365557365722063616e6e6f74206769766520616c6c6f77616e636520746f2068696d73656c66536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f7745524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e6365496e74657265737420697320616c7265616479207265646972656374656420746f207468652075736572496e7465726573742073747265616d2063616e206f6e6c79206265207265646972656374656420696620746865726520697320612076616c69642062616c616e63655472616e7366657272656420616d6f756e74206e6565647320746f2062652067726561746572207468616e207a65726f45524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f20616464726573735468652063616c6c6572206f6620746869732066756e6374696f6e206d7573742062652061206c656e64696e6720706f6f6c45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726f43616c6c6572206973206e6f7420616c6c6f77656420746f2072656469726563742074686520696e746572657374206f66207468652075736572a165627a7a72305820d8fa4c06edadfcaae644e68713a3e2f15c40e84cb44cac09b1fde3962e9e16690029"

// DeployAtoken deploys a new Ethereum contract, binding an instance of Atoken to it.
func DeployAtoken(auth *bind.TransactOpts, backend bind.ContractBackend, _addressesProvider common.Address, _underlyingAsset common.Address, _underlyingAssetDecimals uint8, _name string, _symbol string) (common.Address, *types.Transaction, *Atoken, error) {
	parsed, err := abi.JSON(strings.NewReader(AtokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AtokenBin), backend, _addressesProvider, _underlyingAsset, _underlyingAssetDecimals, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Atoken{AtokenCaller: AtokenCaller{contract: contract}, AtokenTransactor: AtokenTransactor{contract: contract}, AtokenFilterer: AtokenFilterer{contract: contract}}, nil
}

// Atoken is an auto generated Go binding around an Ethereum contract.
type Atoken struct {
	AtokenCaller     // Read-only binding to the contract
	AtokenTransactor // Write-only binding to the contract
	AtokenFilterer   // Log filterer for contract events
}

// AtokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type AtokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AtokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AtokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AtokenSession struct {
	Contract     *Atoken           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AtokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AtokenCallerSession struct {
	Contract *AtokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AtokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AtokenTransactorSession struct {
	Contract     *AtokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AtokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type AtokenRaw struct {
	Contract *Atoken // Generic contract binding to access the raw methods on
}

// AtokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AtokenCallerRaw struct {
	Contract *AtokenCaller // Generic read-only contract binding to access the raw methods on
}

// AtokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AtokenTransactorRaw struct {
	Contract *AtokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAtoken creates a new instance of Atoken, bound to a specific deployed contract.
func NewAtoken(address common.Address, backend bind.ContractBackend) (*Atoken, error) {
	contract, err := bindAtoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Atoken{AtokenCaller: AtokenCaller{contract: contract}, AtokenTransactor: AtokenTransactor{contract: contract}, AtokenFilterer: AtokenFilterer{contract: contract}}, nil
}

// NewAtokenCaller creates a new read-only instance of Atoken, bound to a specific deployed contract.
func NewAtokenCaller(address common.Address, caller bind.ContractCaller) (*AtokenCaller, error) {
	contract, err := bindAtoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AtokenCaller{contract: contract}, nil
}

// NewAtokenTransactor creates a new write-only instance of Atoken, bound to a specific deployed contract.
func NewAtokenTransactor(address common.Address, transactor bind.ContractTransactor) (*AtokenTransactor, error) {
	contract, err := bindAtoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AtokenTransactor{contract: contract}, nil
}

// NewAtokenFilterer creates a new log filterer instance of Atoken, bound to a specific deployed contract.
func NewAtokenFilterer(address common.Address, filterer bind.ContractFilterer) (*AtokenFilterer, error) {
	contract, err := bindAtoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AtokenFilterer{contract: contract}, nil
}

// bindAtoken binds a generic wrapper to an already deployed contract.
func bindAtoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AtokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Atoken *AtokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Atoken.Contract.AtokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Atoken *AtokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atoken.Contract.AtokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Atoken *AtokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Atoken.Contract.AtokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Atoken *AtokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Atoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Atoken *AtokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Atoken *AtokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Atoken.Contract.contract.Transact(opts, method, params...)
}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_Atoken *AtokenCaller) UINTMAXVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Atoken.contract.Call(opts, out, "UINT_MAX_VALUE")
	return *ret0, err
}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_Atoken *AtokenSession) UINTMAXVALUE() (*big.Int, error) {
	return _Atoken.Contract.UINTMAXVALUE(&_Atoken.CallOpts)
}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_Atoken *AtokenCallerSession) UINTMAXVALUE() (*big.Int, error) {
	return _Atoken.Contract.UINTMAXVALUE(&_Atoken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Atoken *AtokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Atoken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Atoken *AtokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Atoken.Contract.Allowance(&_Atoken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Atoken *AtokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Atoken.Contract.Allowance(&_Atoken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_Atoken *AtokenCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Atoken.contract.Call(opts, out, "balanceOf", _user)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_Atoken *AtokenSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _Atoken.Contract.BalanceOf(&_Atoken.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256)
func (_Atoken *AtokenCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _Atoken.Contract.BalanceOf(&_Atoken.CallOpts, _user)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Atoken *AtokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Atoken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Atoken *AtokenSession) Decimals() (uint8, error) {
	return _Atoken.Contract.Decimals(&_Atoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Atoken *AtokenCallerSession) Decimals() (uint8, error) {
	return _Atoken.Contract.Decimals(&_Atoken.CallOpts)
}

// GetInterestRedirectionAddress is a free data retrieval call binding the contract method 0x445e8010.
//
// Solidity: function getInterestRedirectionAddress(address _user) view returns(address)
func (_Atoken *AtokenCaller) GetInterestRedirectionAddress(opts *bind.CallOpts, _user common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Atoken.contract.Call(opts, out, "getInterestRedirectionAddress", _user)
	return *ret0, err
}

// GetInterestRedirectionAddress is a free data retrieval call binding the contract method 0x445e8010.
//
// Solidity: function getInterestRedirectionAddress(address _user) view returns(address)
func (_Atoken *AtokenSession) GetInterestRedirectionAddress(_user common.Address) (common.Address, error) {
	return _Atoken.Contract.GetInterestRedirectionAddress(&_Atoken.CallOpts, _user)
}

// GetInterestRedirectionAddress is a free data retrieval call binding the contract method 0x445e8010.
//
// Solidity: function getInterestRedirectionAddress(address _user) view returns(address)
func (_Atoken *AtokenCallerSession) GetInterestRedirectionAddress(_user common.Address) (common.Address, error) {
	return _Atoken.Contract.GetInterestRedirectionAddress(&_Atoken.CallOpts, _user)
}

// GetRedirectedBalance is a free data retrieval call binding the contract method 0x1d51e7cf.
//
// Solidity: function getRedirectedBalance(address _user) view returns(uint256)
func (_Atoken *AtokenCaller) GetRedirectedBalance(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Atoken.contract.Call(opts, out, "getRedirectedBalance", _user)
	return *ret0, err
}

// GetRedirectedBalance is a free data retrieval call binding the contract method 0x1d51e7cf.
//
// Solidity: function getRedirectedBalance(address _user) view returns(uint256)
func (_Atoken *AtokenSession) GetRedirectedBalance(_user common.Address) (*big.Int, error) {
	return _Atoken.Contract.GetRedirectedBalance(&_Atoken.CallOpts, _user)
}

// GetRedirectedBalance is a free data retrieval call binding the contract method 0x1d51e7cf.
//
// Solidity: function getRedirectedBalance(address _user) view returns(uint256)
func (_Atoken *AtokenCallerSession) GetRedirectedBalance(_user common.Address) (*big.Int, error) {
	return _Atoken.Contract.GetRedirectedBalance(&_Atoken.CallOpts, _user)
}

// GetUserIndex is a free data retrieval call binding the contract method 0xee9907a4.
//
// Solidity: function getUserIndex(address _user) view returns(uint256)
func (_Atoken *AtokenCaller) GetUserIndex(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Atoken.contract.Call(opts, out, "getUserIndex", _user)
	return *ret0, err
}

// GetUserIndex is a free data retrieval call binding the contract method 0xee9907a4.
//
// Solidity: function getUserIndex(address _user) view returns(uint256)
func (_Atoken *AtokenSession) GetUserIndex(_user common.Address) (*big.Int, error) {
	return _Atoken.Contract.GetUserIndex(&_Atoken.CallOpts, _user)
}

// GetUserIndex is a free data retrieval call binding the contract method 0xee9907a4.
//
// Solidity: function getUserIndex(address _user) view returns(uint256)
func (_Atoken *AtokenCallerSession) GetUserIndex(_user common.Address) (*big.Int, error) {
	return _Atoken.Contract.GetUserIndex(&_Atoken.CallOpts, _user)
}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x5eae177c.
//
// Solidity: function isTransferAllowed(address _user, uint256 _amount) view returns(bool)
func (_Atoken *AtokenCaller) IsTransferAllowed(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Atoken.contract.Call(opts, out, "isTransferAllowed", _user, _amount)
	return *ret0, err
}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x5eae177c.
//
// Solidity: function isTransferAllowed(address _user, uint256 _amount) view returns(bool)
func (_Atoken *AtokenSession) IsTransferAllowed(_user common.Address, _amount *big.Int) (bool, error) {
	return _Atoken.Contract.IsTransferAllowed(&_Atoken.CallOpts, _user, _amount)
}

// IsTransferAllowed is a free data retrieval call binding the contract method 0x5eae177c.
//
// Solidity: function isTransferAllowed(address _user, uint256 _amount) view returns(bool)
func (_Atoken *AtokenCallerSession) IsTransferAllowed(_user common.Address, _amount *big.Int) (bool, error) {
	return _Atoken.Contract.IsTransferAllowed(&_Atoken.CallOpts, _user, _amount)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Atoken *AtokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Atoken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Atoken *AtokenSession) Name() (string, error) {
	return _Atoken.Contract.Name(&_Atoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Atoken *AtokenCallerSession) Name() (string, error) {
	return _Atoken.Contract.Name(&_Atoken.CallOpts)
}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address _user) view returns(uint256)
func (_Atoken *AtokenCaller) PrincipalBalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Atoken.contract.Call(opts, out, "principalBalanceOf", _user)
	return *ret0, err
}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address _user) view returns(uint256)
func (_Atoken *AtokenSession) PrincipalBalanceOf(_user common.Address) (*big.Int, error) {
	return _Atoken.Contract.PrincipalBalanceOf(&_Atoken.CallOpts, _user)
}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address _user) view returns(uint256)
func (_Atoken *AtokenCallerSession) PrincipalBalanceOf(_user common.Address) (*big.Int, error) {
	return _Atoken.Contract.PrincipalBalanceOf(&_Atoken.CallOpts, _user)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Atoken *AtokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Atoken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Atoken *AtokenSession) Symbol() (string, error) {
	return _Atoken.Contract.Symbol(&_Atoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Atoken *AtokenCallerSession) Symbol() (string, error) {
	return _Atoken.Contract.Symbol(&_Atoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Atoken *AtokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Atoken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Atoken *AtokenSession) TotalSupply() (*big.Int, error) {
	return _Atoken.Contract.TotalSupply(&_Atoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Atoken *AtokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Atoken.Contract.TotalSupply(&_Atoken.CallOpts)
}

// UnderlyingAssetAddress is a free data retrieval call binding the contract method 0x89d1a0fc.
//
// Solidity: function underlyingAssetAddress() view returns(address)
func (_Atoken *AtokenCaller) UnderlyingAssetAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Atoken.contract.Call(opts, out, "underlyingAssetAddress")
	return *ret0, err
}

// UnderlyingAssetAddress is a free data retrieval call binding the contract method 0x89d1a0fc.
//
// Solidity: function underlyingAssetAddress() view returns(address)
func (_Atoken *AtokenSession) UnderlyingAssetAddress() (common.Address, error) {
	return _Atoken.Contract.UnderlyingAssetAddress(&_Atoken.CallOpts)
}

// UnderlyingAssetAddress is a free data retrieval call binding the contract method 0x89d1a0fc.
//
// Solidity: function underlyingAssetAddress() view returns(address)
func (_Atoken *AtokenCallerSession) UnderlyingAssetAddress() (common.Address, error) {
	return _Atoken.Contract.UnderlyingAssetAddress(&_Atoken.CallOpts)
}

// AllowInterestRedirectionTo is a paid mutator transaction binding the contract method 0x12c87c2d.
//
// Solidity: function allowInterestRedirectionTo(address _to) returns()
func (_Atoken *AtokenTransactor) AllowInterestRedirectionTo(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "allowInterestRedirectionTo", _to)
}

// AllowInterestRedirectionTo is a paid mutator transaction binding the contract method 0x12c87c2d.
//
// Solidity: function allowInterestRedirectionTo(address _to) returns()
func (_Atoken *AtokenSession) AllowInterestRedirectionTo(_to common.Address) (*types.Transaction, error) {
	return _Atoken.Contract.AllowInterestRedirectionTo(&_Atoken.TransactOpts, _to)
}

// AllowInterestRedirectionTo is a paid mutator transaction binding the contract method 0x12c87c2d.
//
// Solidity: function allowInterestRedirectionTo(address _to) returns()
func (_Atoken *AtokenTransactorSession) AllowInterestRedirectionTo(_to common.Address) (*types.Transaction, error) {
	return _Atoken.Contract.AllowInterestRedirectionTo(&_Atoken.TransactOpts, _to)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Atoken *AtokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Atoken *AtokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.Approve(&_Atoken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Atoken *AtokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.Approve(&_Atoken.TransactOpts, spender, amount)
}

// BurnOnLiquidation is a paid mutator transaction binding the contract method 0x3edb7cb8.
//
// Solidity: function burnOnLiquidation(address _account, uint256 _value) returns()
func (_Atoken *AtokenTransactor) BurnOnLiquidation(opts *bind.TransactOpts, _account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "burnOnLiquidation", _account, _value)
}

// BurnOnLiquidation is a paid mutator transaction binding the contract method 0x3edb7cb8.
//
// Solidity: function burnOnLiquidation(address _account, uint256 _value) returns()
func (_Atoken *AtokenSession) BurnOnLiquidation(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.BurnOnLiquidation(&_Atoken.TransactOpts, _account, _value)
}

// BurnOnLiquidation is a paid mutator transaction binding the contract method 0x3edb7cb8.
//
// Solidity: function burnOnLiquidation(address _account, uint256 _value) returns()
func (_Atoken *AtokenTransactorSession) BurnOnLiquidation(_account common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.BurnOnLiquidation(&_Atoken.TransactOpts, _account, _value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Atoken *AtokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Atoken *AtokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.DecreaseAllowance(&_Atoken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Atoken *AtokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.DecreaseAllowance(&_Atoken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Atoken *AtokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Atoken *AtokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.IncreaseAllowance(&_Atoken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Atoken *AtokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.IncreaseAllowance(&_Atoken.TransactOpts, spender, addedValue)
}

// MintOnDeposit is a paid mutator transaction binding the contract method 0x94362e8b.
//
// Solidity: function mintOnDeposit(address _account, uint256 _amount) returns()
func (_Atoken *AtokenTransactor) MintOnDeposit(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "mintOnDeposit", _account, _amount)
}

// MintOnDeposit is a paid mutator transaction binding the contract method 0x94362e8b.
//
// Solidity: function mintOnDeposit(address _account, uint256 _amount) returns()
func (_Atoken *AtokenSession) MintOnDeposit(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.MintOnDeposit(&_Atoken.TransactOpts, _account, _amount)
}

// MintOnDeposit is a paid mutator transaction binding the contract method 0x94362e8b.
//
// Solidity: function mintOnDeposit(address _account, uint256 _amount) returns()
func (_Atoken *AtokenTransactorSession) MintOnDeposit(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.MintOnDeposit(&_Atoken.TransactOpts, _account, _amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 _amount) returns()
func (_Atoken *AtokenTransactor) Redeem(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "redeem", _amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 _amount) returns()
func (_Atoken *AtokenSession) Redeem(_amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.Redeem(&_Atoken.TransactOpts, _amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 _amount) returns()
func (_Atoken *AtokenTransactorSession) Redeem(_amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.Redeem(&_Atoken.TransactOpts, _amount)
}

// RedirectInterestStream is a paid mutator transaction binding the contract method 0x0e49072d.
//
// Solidity: function redirectInterestStream(address _to) returns()
func (_Atoken *AtokenTransactor) RedirectInterestStream(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "redirectInterestStream", _to)
}

// RedirectInterestStream is a paid mutator transaction binding the contract method 0x0e49072d.
//
// Solidity: function redirectInterestStream(address _to) returns()
func (_Atoken *AtokenSession) RedirectInterestStream(_to common.Address) (*types.Transaction, error) {
	return _Atoken.Contract.RedirectInterestStream(&_Atoken.TransactOpts, _to)
}

// RedirectInterestStream is a paid mutator transaction binding the contract method 0x0e49072d.
//
// Solidity: function redirectInterestStream(address _to) returns()
func (_Atoken *AtokenTransactorSession) RedirectInterestStream(_to common.Address) (*types.Transaction, error) {
	return _Atoken.Contract.RedirectInterestStream(&_Atoken.TransactOpts, _to)
}

// RedirectInterestStreamOf is a paid mutator transaction binding the contract method 0x325a9b13.
//
// Solidity: function redirectInterestStreamOf(address _from, address _to) returns()
func (_Atoken *AtokenTransactor) RedirectInterestStreamOf(opts *bind.TransactOpts, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "redirectInterestStreamOf", _from, _to)
}

// RedirectInterestStreamOf is a paid mutator transaction binding the contract method 0x325a9b13.
//
// Solidity: function redirectInterestStreamOf(address _from, address _to) returns()
func (_Atoken *AtokenSession) RedirectInterestStreamOf(_from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Atoken.Contract.RedirectInterestStreamOf(&_Atoken.TransactOpts, _from, _to)
}

// RedirectInterestStreamOf is a paid mutator transaction binding the contract method 0x325a9b13.
//
// Solidity: function redirectInterestStreamOf(address _from, address _to) returns()
func (_Atoken *AtokenTransactorSession) RedirectInterestStreamOf(_from common.Address, _to common.Address) (*types.Transaction, error) {
	return _Atoken.Contract.RedirectInterestStreamOf(&_Atoken.TransactOpts, _from, _to)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Atoken *AtokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Atoken *AtokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.Transfer(&_Atoken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Atoken *AtokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.Transfer(&_Atoken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Atoken *AtokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Atoken *AtokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.TransferFrom(&_Atoken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Atoken *AtokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.TransferFrom(&_Atoken.TransactOpts, sender, recipient, amount)
}

// TransferOnLiquidation is a paid mutator transaction binding the contract method 0xf866c319.
//
// Solidity: function transferOnLiquidation(address _from, address _to, uint256 _value) returns()
func (_Atoken *AtokenTransactor) TransferOnLiquidation(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Atoken.contract.Transact(opts, "transferOnLiquidation", _from, _to, _value)
}

// TransferOnLiquidation is a paid mutator transaction binding the contract method 0xf866c319.
//
// Solidity: function transferOnLiquidation(address _from, address _to, uint256 _value) returns()
func (_Atoken *AtokenSession) TransferOnLiquidation(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.TransferOnLiquidation(&_Atoken.TransactOpts, _from, _to, _value)
}

// TransferOnLiquidation is a paid mutator transaction binding the contract method 0xf866c319.
//
// Solidity: function transferOnLiquidation(address _from, address _to, uint256 _value) returns()
func (_Atoken *AtokenTransactorSession) TransferOnLiquidation(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Atoken.Contract.TransferOnLiquidation(&_Atoken.TransactOpts, _from, _to, _value)
}

// AtokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Atoken contract.
type AtokenApprovalIterator struct {
	Event *AtokenApproval // Event containing the contract specifics and raw log

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
func (it *AtokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtokenApproval)
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
		it.Event = new(AtokenApproval)
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
func (it *AtokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtokenApproval represents a Approval event raised by the Atoken contract.
type AtokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Atoken *AtokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AtokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Atoken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AtokenApprovalIterator{contract: _Atoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Atoken *AtokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AtokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Atoken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtokenApproval)
				if err := _Atoken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Atoken *AtokenFilterer) ParseApproval(log types.Log) (*AtokenApproval, error) {
	event := new(AtokenApproval)
	if err := _Atoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AtokenBalanceTransferIterator is returned from FilterBalanceTransfer and is used to iterate over the raw logs and unpacked data for BalanceTransfer events raised by the Atoken contract.
type AtokenBalanceTransferIterator struct {
	Event *AtokenBalanceTransfer // Event containing the contract specifics and raw log

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
func (it *AtokenBalanceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtokenBalanceTransfer)
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
		it.Event = new(AtokenBalanceTransfer)
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
func (it *AtokenBalanceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtokenBalanceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtokenBalanceTransfer represents a BalanceTransfer event raised by the Atoken contract.
type AtokenBalanceTransfer struct {
	From                common.Address
	To                  common.Address
	Value               *big.Int
	FromBalanceIncrease *big.Int
	ToBalanceIncrease   *big.Int
	FromIndex           *big.Int
	ToIndex             *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterBalanceTransfer is a free log retrieval operation binding the contract event 0x89a178eaa27e0cd201bd795ca8ff716ac0df9618494510ca79771cfc66ffcde8.
//
// Solidity: event BalanceTransfer(address indexed _from, address indexed _to, uint256 _value, uint256 _fromBalanceIncrease, uint256 _toBalanceIncrease, uint256 _fromIndex, uint256 _toIndex)
func (_Atoken *AtokenFilterer) FilterBalanceTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*AtokenBalanceTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Atoken.contract.FilterLogs(opts, "BalanceTransfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &AtokenBalanceTransferIterator{contract: _Atoken.contract, event: "BalanceTransfer", logs: logs, sub: sub}, nil
}

// WatchBalanceTransfer is a free log subscription operation binding the contract event 0x89a178eaa27e0cd201bd795ca8ff716ac0df9618494510ca79771cfc66ffcde8.
//
// Solidity: event BalanceTransfer(address indexed _from, address indexed _to, uint256 _value, uint256 _fromBalanceIncrease, uint256 _toBalanceIncrease, uint256 _fromIndex, uint256 _toIndex)
func (_Atoken *AtokenFilterer) WatchBalanceTransfer(opts *bind.WatchOpts, sink chan<- *AtokenBalanceTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Atoken.contract.WatchLogs(opts, "BalanceTransfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtokenBalanceTransfer)
				if err := _Atoken.contract.UnpackLog(event, "BalanceTransfer", log); err != nil {
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

// ParseBalanceTransfer is a log parse operation binding the contract event 0x89a178eaa27e0cd201bd795ca8ff716ac0df9618494510ca79771cfc66ffcde8.
//
// Solidity: event BalanceTransfer(address indexed _from, address indexed _to, uint256 _value, uint256 _fromBalanceIncrease, uint256 _toBalanceIncrease, uint256 _fromIndex, uint256 _toIndex)
func (_Atoken *AtokenFilterer) ParseBalanceTransfer(log types.Log) (*AtokenBalanceTransfer, error) {
	event := new(AtokenBalanceTransfer)
	if err := _Atoken.contract.UnpackLog(event, "BalanceTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AtokenBurnOnLiquidationIterator is returned from FilterBurnOnLiquidation and is used to iterate over the raw logs and unpacked data for BurnOnLiquidation events raised by the Atoken contract.
type AtokenBurnOnLiquidationIterator struct {
	Event *AtokenBurnOnLiquidation // Event containing the contract specifics and raw log

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
func (it *AtokenBurnOnLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtokenBurnOnLiquidation)
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
		it.Event = new(AtokenBurnOnLiquidation)
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
func (it *AtokenBurnOnLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtokenBurnOnLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtokenBurnOnLiquidation represents a BurnOnLiquidation event raised by the Atoken contract.
type AtokenBurnOnLiquidation struct {
	From                common.Address
	Value               *big.Int
	FromBalanceIncrease *big.Int
	FromIndex           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterBurnOnLiquidation is a free log retrieval operation binding the contract event 0x90e5d3d68ec162d9c7de393037a3ede04dd44f68e051bf2ace4a73c299dbc4db.
//
// Solidity: event BurnOnLiquidation(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_Atoken *AtokenFilterer) FilterBurnOnLiquidation(opts *bind.FilterOpts, _from []common.Address) (*AtokenBurnOnLiquidationIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Atoken.contract.FilterLogs(opts, "BurnOnLiquidation", _fromRule)
	if err != nil {
		return nil, err
	}
	return &AtokenBurnOnLiquidationIterator{contract: _Atoken.contract, event: "BurnOnLiquidation", logs: logs, sub: sub}, nil
}

// WatchBurnOnLiquidation is a free log subscription operation binding the contract event 0x90e5d3d68ec162d9c7de393037a3ede04dd44f68e051bf2ace4a73c299dbc4db.
//
// Solidity: event BurnOnLiquidation(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_Atoken *AtokenFilterer) WatchBurnOnLiquidation(opts *bind.WatchOpts, sink chan<- *AtokenBurnOnLiquidation, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Atoken.contract.WatchLogs(opts, "BurnOnLiquidation", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtokenBurnOnLiquidation)
				if err := _Atoken.contract.UnpackLog(event, "BurnOnLiquidation", log); err != nil {
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

// ParseBurnOnLiquidation is a log parse operation binding the contract event 0x90e5d3d68ec162d9c7de393037a3ede04dd44f68e051bf2ace4a73c299dbc4db.
//
// Solidity: event BurnOnLiquidation(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_Atoken *AtokenFilterer) ParseBurnOnLiquidation(log types.Log) (*AtokenBurnOnLiquidation, error) {
	event := new(AtokenBurnOnLiquidation)
	if err := _Atoken.contract.UnpackLog(event, "BurnOnLiquidation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AtokenInterestRedirectionAllowanceChangedIterator is returned from FilterInterestRedirectionAllowanceChanged and is used to iterate over the raw logs and unpacked data for InterestRedirectionAllowanceChanged events raised by the Atoken contract.
type AtokenInterestRedirectionAllowanceChangedIterator struct {
	Event *AtokenInterestRedirectionAllowanceChanged // Event containing the contract specifics and raw log

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
func (it *AtokenInterestRedirectionAllowanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtokenInterestRedirectionAllowanceChanged)
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
		it.Event = new(AtokenInterestRedirectionAllowanceChanged)
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
func (it *AtokenInterestRedirectionAllowanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtokenInterestRedirectionAllowanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtokenInterestRedirectionAllowanceChanged represents a InterestRedirectionAllowanceChanged event raised by the Atoken contract.
type AtokenInterestRedirectionAllowanceChanged struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterInterestRedirectionAllowanceChanged is a free log retrieval operation binding the contract event 0xc2d6a42a9d5273283f73009a07aacfb043f2f91173a8d9d33b504afe898db08b.
//
// Solidity: event InterestRedirectionAllowanceChanged(address indexed _from, address indexed _to)
func (_Atoken *AtokenFilterer) FilterInterestRedirectionAllowanceChanged(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*AtokenInterestRedirectionAllowanceChangedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Atoken.contract.FilterLogs(opts, "InterestRedirectionAllowanceChanged", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &AtokenInterestRedirectionAllowanceChangedIterator{contract: _Atoken.contract, event: "InterestRedirectionAllowanceChanged", logs: logs, sub: sub}, nil
}

// WatchInterestRedirectionAllowanceChanged is a free log subscription operation binding the contract event 0xc2d6a42a9d5273283f73009a07aacfb043f2f91173a8d9d33b504afe898db08b.
//
// Solidity: event InterestRedirectionAllowanceChanged(address indexed _from, address indexed _to)
func (_Atoken *AtokenFilterer) WatchInterestRedirectionAllowanceChanged(opts *bind.WatchOpts, sink chan<- *AtokenInterestRedirectionAllowanceChanged, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Atoken.contract.WatchLogs(opts, "InterestRedirectionAllowanceChanged", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtokenInterestRedirectionAllowanceChanged)
				if err := _Atoken.contract.UnpackLog(event, "InterestRedirectionAllowanceChanged", log); err != nil {
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

// ParseInterestRedirectionAllowanceChanged is a log parse operation binding the contract event 0xc2d6a42a9d5273283f73009a07aacfb043f2f91173a8d9d33b504afe898db08b.
//
// Solidity: event InterestRedirectionAllowanceChanged(address indexed _from, address indexed _to)
func (_Atoken *AtokenFilterer) ParseInterestRedirectionAllowanceChanged(log types.Log) (*AtokenInterestRedirectionAllowanceChanged, error) {
	event := new(AtokenInterestRedirectionAllowanceChanged)
	if err := _Atoken.contract.UnpackLog(event, "InterestRedirectionAllowanceChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AtokenInterestStreamRedirectedIterator is returned from FilterInterestStreamRedirected and is used to iterate over the raw logs and unpacked data for InterestStreamRedirected events raised by the Atoken contract.
type AtokenInterestStreamRedirectedIterator struct {
	Event *AtokenInterestStreamRedirected // Event containing the contract specifics and raw log

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
func (it *AtokenInterestStreamRedirectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtokenInterestStreamRedirected)
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
		it.Event = new(AtokenInterestStreamRedirected)
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
func (it *AtokenInterestStreamRedirectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtokenInterestStreamRedirectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtokenInterestStreamRedirected represents a InterestStreamRedirected event raised by the Atoken contract.
type AtokenInterestStreamRedirected struct {
	From                common.Address
	To                  common.Address
	RedirectedBalance   *big.Int
	FromBalanceIncrease *big.Int
	FromIndex           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterInterestStreamRedirected is a free log retrieval operation binding the contract event 0x5e3cad45b1fe24159d1cb39788d82d0f69cc15770aa96fba1d3d1a7348735594.
//
// Solidity: event InterestStreamRedirected(address indexed _from, address indexed _to, uint256 _redirectedBalance, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_Atoken *AtokenFilterer) FilterInterestStreamRedirected(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*AtokenInterestStreamRedirectedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Atoken.contract.FilterLogs(opts, "InterestStreamRedirected", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &AtokenInterestStreamRedirectedIterator{contract: _Atoken.contract, event: "InterestStreamRedirected", logs: logs, sub: sub}, nil
}

// WatchInterestStreamRedirected is a free log subscription operation binding the contract event 0x5e3cad45b1fe24159d1cb39788d82d0f69cc15770aa96fba1d3d1a7348735594.
//
// Solidity: event InterestStreamRedirected(address indexed _from, address indexed _to, uint256 _redirectedBalance, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_Atoken *AtokenFilterer) WatchInterestStreamRedirected(opts *bind.WatchOpts, sink chan<- *AtokenInterestStreamRedirected, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Atoken.contract.WatchLogs(opts, "InterestStreamRedirected", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtokenInterestStreamRedirected)
				if err := _Atoken.contract.UnpackLog(event, "InterestStreamRedirected", log); err != nil {
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

// ParseInterestStreamRedirected is a log parse operation binding the contract event 0x5e3cad45b1fe24159d1cb39788d82d0f69cc15770aa96fba1d3d1a7348735594.
//
// Solidity: event InterestStreamRedirected(address indexed _from, address indexed _to, uint256 _redirectedBalance, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_Atoken *AtokenFilterer) ParseInterestStreamRedirected(log types.Log) (*AtokenInterestStreamRedirected, error) {
	event := new(AtokenInterestStreamRedirected)
	if err := _Atoken.contract.UnpackLog(event, "InterestStreamRedirected", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AtokenMintOnDepositIterator is returned from FilterMintOnDeposit and is used to iterate over the raw logs and unpacked data for MintOnDeposit events raised by the Atoken contract.
type AtokenMintOnDepositIterator struct {
	Event *AtokenMintOnDeposit // Event containing the contract specifics and raw log

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
func (it *AtokenMintOnDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtokenMintOnDeposit)
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
		it.Event = new(AtokenMintOnDeposit)
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
func (it *AtokenMintOnDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtokenMintOnDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtokenMintOnDeposit represents a MintOnDeposit event raised by the Atoken contract.
type AtokenMintOnDeposit struct {
	From                common.Address
	Value               *big.Int
	FromBalanceIncrease *big.Int
	FromIndex           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMintOnDeposit is a free log retrieval operation binding the contract event 0xbe7799898ca2d813ff902b487c1b434ab45b47edd8f38b77ad5e99aae8341b7a.
//
// Solidity: event MintOnDeposit(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_Atoken *AtokenFilterer) FilterMintOnDeposit(opts *bind.FilterOpts, _from []common.Address) (*AtokenMintOnDepositIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Atoken.contract.FilterLogs(opts, "MintOnDeposit", _fromRule)
	if err != nil {
		return nil, err
	}
	return &AtokenMintOnDepositIterator{contract: _Atoken.contract, event: "MintOnDeposit", logs: logs, sub: sub}, nil
}

// WatchMintOnDeposit is a free log subscription operation binding the contract event 0xbe7799898ca2d813ff902b487c1b434ab45b47edd8f38b77ad5e99aae8341b7a.
//
// Solidity: event MintOnDeposit(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_Atoken *AtokenFilterer) WatchMintOnDeposit(opts *bind.WatchOpts, sink chan<- *AtokenMintOnDeposit, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Atoken.contract.WatchLogs(opts, "MintOnDeposit", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtokenMintOnDeposit)
				if err := _Atoken.contract.UnpackLog(event, "MintOnDeposit", log); err != nil {
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

// ParseMintOnDeposit is a log parse operation binding the contract event 0xbe7799898ca2d813ff902b487c1b434ab45b47edd8f38b77ad5e99aae8341b7a.
//
// Solidity: event MintOnDeposit(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_Atoken *AtokenFilterer) ParseMintOnDeposit(log types.Log) (*AtokenMintOnDeposit, error) {
	event := new(AtokenMintOnDeposit)
	if err := _Atoken.contract.UnpackLog(event, "MintOnDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AtokenRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Atoken contract.
type AtokenRedeemIterator struct {
	Event *AtokenRedeem // Event containing the contract specifics and raw log

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
func (it *AtokenRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtokenRedeem)
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
		it.Event = new(AtokenRedeem)
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
func (it *AtokenRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtokenRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtokenRedeem represents a Redeem event raised by the Atoken contract.
type AtokenRedeem struct {
	From                common.Address
	Value               *big.Int
	FromBalanceIncrease *big.Int
	FromIndex           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xbd5034ffbd47e4e72a94baa2cdb74c6fad73cb3bcdc13036b72ec8306f5a7646.
//
// Solidity: event Redeem(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_Atoken *AtokenFilterer) FilterRedeem(opts *bind.FilterOpts, _from []common.Address) (*AtokenRedeemIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Atoken.contract.FilterLogs(opts, "Redeem", _fromRule)
	if err != nil {
		return nil, err
	}
	return &AtokenRedeemIterator{contract: _Atoken.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xbd5034ffbd47e4e72a94baa2cdb74c6fad73cb3bcdc13036b72ec8306f5a7646.
//
// Solidity: event Redeem(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_Atoken *AtokenFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *AtokenRedeem, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Atoken.contract.WatchLogs(opts, "Redeem", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtokenRedeem)
				if err := _Atoken.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xbd5034ffbd47e4e72a94baa2cdb74c6fad73cb3bcdc13036b72ec8306f5a7646.
//
// Solidity: event Redeem(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_Atoken *AtokenFilterer) ParseRedeem(log types.Log) (*AtokenRedeem, error) {
	event := new(AtokenRedeem)
	if err := _Atoken.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AtokenRedirectedBalanceUpdatedIterator is returned from FilterRedirectedBalanceUpdated and is used to iterate over the raw logs and unpacked data for RedirectedBalanceUpdated events raised by the Atoken contract.
type AtokenRedirectedBalanceUpdatedIterator struct {
	Event *AtokenRedirectedBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *AtokenRedirectedBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtokenRedirectedBalanceUpdated)
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
		it.Event = new(AtokenRedirectedBalanceUpdated)
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
func (it *AtokenRedirectedBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtokenRedirectedBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtokenRedirectedBalanceUpdated represents a RedirectedBalanceUpdated event raised by the Atoken contract.
type AtokenRedirectedBalanceUpdated struct {
	TargetAddress            common.Address
	TargetBalanceIncrease    *big.Int
	TargetIndex              *big.Int
	RedirectedBalanceAdded   *big.Int
	RedirectedBalanceRemoved *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterRedirectedBalanceUpdated is a free log retrieval operation binding the contract event 0x70ff8cf632603e2bfd1afb7e4061acce53d95356b1be9726b99fa22ba733b01f.
//
// Solidity: event RedirectedBalanceUpdated(address indexed _targetAddress, uint256 _targetBalanceIncrease, uint256 _targetIndex, uint256 _redirectedBalanceAdded, uint256 _redirectedBalanceRemoved)
func (_Atoken *AtokenFilterer) FilterRedirectedBalanceUpdated(opts *bind.FilterOpts, _targetAddress []common.Address) (*AtokenRedirectedBalanceUpdatedIterator, error) {

	var _targetAddressRule []interface{}
	for _, _targetAddressItem := range _targetAddress {
		_targetAddressRule = append(_targetAddressRule, _targetAddressItem)
	}

	logs, sub, err := _Atoken.contract.FilterLogs(opts, "RedirectedBalanceUpdated", _targetAddressRule)
	if err != nil {
		return nil, err
	}
	return &AtokenRedirectedBalanceUpdatedIterator{contract: _Atoken.contract, event: "RedirectedBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchRedirectedBalanceUpdated is a free log subscription operation binding the contract event 0x70ff8cf632603e2bfd1afb7e4061acce53d95356b1be9726b99fa22ba733b01f.
//
// Solidity: event RedirectedBalanceUpdated(address indexed _targetAddress, uint256 _targetBalanceIncrease, uint256 _targetIndex, uint256 _redirectedBalanceAdded, uint256 _redirectedBalanceRemoved)
func (_Atoken *AtokenFilterer) WatchRedirectedBalanceUpdated(opts *bind.WatchOpts, sink chan<- *AtokenRedirectedBalanceUpdated, _targetAddress []common.Address) (event.Subscription, error) {

	var _targetAddressRule []interface{}
	for _, _targetAddressItem := range _targetAddress {
		_targetAddressRule = append(_targetAddressRule, _targetAddressItem)
	}

	logs, sub, err := _Atoken.contract.WatchLogs(opts, "RedirectedBalanceUpdated", _targetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtokenRedirectedBalanceUpdated)
				if err := _Atoken.contract.UnpackLog(event, "RedirectedBalanceUpdated", log); err != nil {
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

// ParseRedirectedBalanceUpdated is a log parse operation binding the contract event 0x70ff8cf632603e2bfd1afb7e4061acce53d95356b1be9726b99fa22ba733b01f.
//
// Solidity: event RedirectedBalanceUpdated(address indexed _targetAddress, uint256 _targetBalanceIncrease, uint256 _targetIndex, uint256 _redirectedBalanceAdded, uint256 _redirectedBalanceRemoved)
func (_Atoken *AtokenFilterer) ParseRedirectedBalanceUpdated(log types.Log) (*AtokenRedirectedBalanceUpdated, error) {
	event := new(AtokenRedirectedBalanceUpdated)
	if err := _Atoken.contract.UnpackLog(event, "RedirectedBalanceUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AtokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Atoken contract.
type AtokenTransferIterator struct {
	Event *AtokenTransfer // Event containing the contract specifics and raw log

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
func (it *AtokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtokenTransfer)
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
		it.Event = new(AtokenTransfer)
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
func (it *AtokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtokenTransfer represents a Transfer event raised by the Atoken contract.
type AtokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Atoken *AtokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AtokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Atoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AtokenTransferIterator{contract: _Atoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Atoken *AtokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AtokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Atoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtokenTransfer)
				if err := _Atoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Atoken *AtokenFilterer) ParseTransfer(log types.Log) (*AtokenTransfer, error) {
	event := new(AtokenTransfer)
	if err := _Atoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
