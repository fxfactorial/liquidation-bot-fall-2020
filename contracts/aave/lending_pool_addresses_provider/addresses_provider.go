// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lending_pool_addresses_provider

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

// LendingPoolAddressesProviderABI is the input ABI used to generate the binding from.
const LendingPoolAddressesProviderABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolCoreUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolParametersProviderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolConfiguratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolLiquidationManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolDataProviderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"EthereumAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PriceOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingRateOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"FeeProviderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"TokenDistributorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ProxyCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLendingPool\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"setLendingPoolImpl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLendingPoolCore\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lendingPoolCore\",\"type\":\"address\"}],\"name\":\"setLendingPoolCoreImpl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLendingPoolConfigurator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_configurator\",\"type\":\"address\"}],\"name\":\"setLendingPoolConfiguratorImpl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLendingPoolDataProvider\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"setLendingPoolDataProviderImpl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLendingPoolParametersProvider\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_parametersProvider\",\"type\":\"address\"}],\"name\":\"setLendingPoolParametersProviderImpl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFeeProvider\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_feeProvider\",\"type\":\"address\"}],\"name\":\"setFeeProviderImpl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLendingPoolLiquidationManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"setLendingPoolLiquidationManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLendingPoolManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lendingPoolManager\",\"type\":\"address\"}],\"name\":\"setLendingPoolManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPriceOracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_priceOracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLendingRateOracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lendingRateOracle\",\"type\":\"address\"}],\"name\":\"setLendingRateOracle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTokenDistributor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenDistributor\",\"type\":\"address\"}],\"name\":\"setTokenDistributor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// LendingPoolAddressesProviderBin is the compiled bytecode used for deploying new contracts.
var LendingPoolAddressesProviderBin = "0x608060405261001261005f60201b60201c565b600080546001600160a01b0319166001600160a01b03928316178082556040519216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3610063565b3390565b611d54806100726000396000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c80635aef021f116100f9578063bfedc10311610097578063ee89129611610071578063ee89129614610407578063f2fde38b1461040f578063fbeefc3c14610435578063fca513a81461043d576101b9565b8063bfedc103146103b3578063c12542df146103d9578063ed6ff760146103ff576101b9565b806385c858b1116100d357806385c858b1146103615780638da5cb5b146103695780638f32d59b14610371578063a5eface21461038d576101b9565b80635aef021f1461030d578063715018a614610333578063820d12741461033b576101b9565b806333128d591161016657806340fdcadc1161014057806340fdcadc1461029357806344ce375b146102b9578063530e784f146102df5780635834eb9a14610305576101b9565b806333128d591461025d5780633618abba1461026557806338280e6b1461026d576101b9565b806321f8a7211161019757806321f8a721146102125780632a62c6361461022f5780632f58b80d14610255576101b9565b80630261bf8b146101be57806304061d8e146101e25780631c827204146101ea575b600080fd5b6101c6610445565b604080516001600160a01b039092168252519081900360200190f35b6101c6610475565b6102106004803603602081101561020057600080fd5b50356001600160a01b03166104a0565b005b6101c66004803603602081101561022857600080fd5b503561054d565b6102106004803603602081101561024557600080fd5b50356001600160a01b0316610568565b6101c6610615565b6101c6610640565b6101c661066b565b6102106004803603602081101561028357600080fd5b50356001600160a01b0316610696565b610210600480360360208110156102a957600080fd5b50356001600160a01b0316610743565b610210600480360360208110156102cf57600080fd5b50356001600160a01b03166107f0565b610210600480360360208110156102f557600080fd5b50356001600160a01b031661089d565b6101c661094a565b6102106004803603602081101561032357600080fd5b50356001600160a01b0316610975565b610210610a22565b6102106004803603602081101561035157600080fd5b50356001600160a01b0316610ac5565b6101c6610b72565b6101c6610b9d565b610379610bac565b604080519115158252519081900360200190f35b610210600480360360208110156103a357600080fd5b50356001600160a01b0316610bd0565b610210600480360360208110156103c957600080fd5b50356001600160a01b0316610c7d565b610210600480360360208110156103ef57600080fd5b50356001600160a01b0316610d2a565b6101c6610dd7565b6101c6610e09565b6102106004803603602081101561042557600080fd5b50356001600160a01b0316610e34565b6101c6610e8c565b6101c6610eb7565b60006104707f4c454e44494e475f504f4f4c000000000000000000000000000000000000000061054d565b905090565b60006104707f504152414d45544552535f50524f56494445520000000000000000000000000061054d565b6104a8610bac565b15156104ec5760408051600160e51b62461bcd0281526020600482018190526024820152600080516020611d09833981519152604482015290519081900360640190fd5b6105167f4c454e44494e475f504f4f4c5f434f524500000000000000000000000000000082610ee2565b6040516001600160a01b038216907f71c226bb2879778ca1298196bf7cc1256baee9d05b496c31ee627d35a471b5b790600090a250565b6000908152600160205260409020546001600160a01b031690565b610570610bac565b15156105b45760408051600160e51b62461bcd0281526020600482018190526024820152600080516020611d09833981519152604482015290519081900360640190fd5b6105de7f4645455f50524f5649444552000000000000000000000000000000000000000082610ee2565b6040516001600160a01b038216907f18e1a55b8eff9c93921eecfa1462d6a8cbb80b3988db3eb14c039e43fdb2266190600090a250565b60006104707f444154415f50524f56494445520000000000000000000000000000000000000061054d565b60006104707f4c454e44494e475f504f4f4c5f4d414e4147455200000000000000000000000061054d565b60006104707f4c454e44494e475f524154455f4f5241434c450000000000000000000000000061054d565b61069e610bac565b15156106e25760408051600160e51b62461bcd0281526020600482018190526024820152600080516020611d09833981519152604482015290519081900360640190fd5b61070c7f544f4b454e5f4449535452494255544f52000000000000000000000000000000826111c4565b6040516001600160a01b038216907fa8b48a56ec01f81c3615a21ec43e16b925c26293e0801cf6330427f2a687f05390600090a250565b61074b610bac565b151561078f5760408051600160e51b62461bcd0281526020600482018190526024820152600080516020611d09833981519152604482015290519081900360640190fd5b6107b97f4c454e44494e475f504f4f4c5f4d414e41474552000000000000000000000000826111c4565b6040516001600160a01b038216907fd5280c51185a38d36f7a0f5e56cac6248312bb70d0974782fa5a595127e0e08e90600090a250565b6107f8610bac565b151561083c5760408051600160e51b62461bcd0281526020600482018190526024820152600080516020611d09833981519152604482015290519081900360640190fd5b6108667f4c49515549444154494f4e5f4d414e4147455200000000000000000000000000826111c4565b6040516001600160a01b038216907f1a76cb769b814bc038223da86932b099b20aae03473683e6d98f5c3554e2606490600090a250565b6108a5610bac565b15156108e95760408051600160e51b62461bcd0281526020600482018190526024820152600080516020611d09833981519152604482015290519081900360640190fd5b6109137f50524943455f4f5241434c450000000000000000000000000000000000000000826111c4565b6040516001600160a01b038216907fefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd90600090a250565b60006104707f4c49515549444154494f4e5f4d414e414745520000000000000000000000000061054d565b61097d610bac565b15156109c15760408051600160e51b62461bcd0281526020600482018190526024820152600080516020611d09833981519152604482015290519081900360640190fd5b6109eb7f4c454e44494e475f504f4f4c000000000000000000000000000000000000000082610ee2565b6040516001600160a01b038216907fc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa490600090a250565b610a2a610bac565b1515610a6e5760408051600160e51b62461bcd0281526020600482018190526024820152600080516020611d09833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36000805473ffffffffffffffffffffffffffffffffffffffff19169055565b610acd610bac565b1515610b115760408051600160e51b62461bcd0281526020600482018190526024820152600080516020611d09833981519152604482015290519081900360640190fd5b610b3b7f4c454e44494e475f524154455f4f5241434c4500000000000000000000000000826111c4565b6040516001600160a01b038216907f5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b590600090a250565b60006104707f4c454e44494e475f504f4f4c5f434f4e464947555241544f520000000000000061054d565b6000546001600160a01b031690565b600080546001600160a01b0316610bc16111ff565b6001600160a01b031614905090565b610bd8610bac565b1515610c1c5760408051600160e51b62461bcd0281526020600482018190526024820152600080516020611d09833981519152604482015290519081900360640190fd5b610c467f504152414d45544552535f50524f56494445520000000000000000000000000082610ee2565b6040516001600160a01b038216907fce16ea9b2fd7cadddd0f7359971028f50b5ddba33dfae1a9bdea956fabb1e28090600090a250565b610c85610bac565b1515610cc95760408051600160e51b62461bcd0281526020600482018190526024820152600080516020611d09833981519152604482015290519081900360640190fd5b610cf37f444154415f50524f56494445520000000000000000000000000000000000000082610ee2565b6040516001600160a01b038216907f07890d7d10db37434d76ee4f2928fb2dc66227c3b3391206aced4f7bcb59cdb090600090a250565b610d32610bac565b1515610d765760408051600160e51b62461bcd0281526020600482018190526024820152600080516020611d09833981519152604482015290519081900360640190fd5b610da07f4c454e44494e475f504f4f4c5f434f4e464947555241544f520000000000000082610ee2565b6040516001600160a01b038216907fdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae76372990600090a250565b600080610e037f4c454e44494e475f504f4f4c5f434f524500000000000000000000000000000061054d565b91505090565b60006104707f544f4b454e5f4449535452494255544f5200000000000000000000000000000061054d565b610e3c610bac565b1515610e805760408051600160e51b62461bcd0281526020600482018190526024820152600080516020611d09833981519152604482015290519081900360640190fd5b610e8981611203565b50565b60006104707f4645455f50524f5649444552000000000000000000000000000000000000000061054d565b60006104707f50524943455f4f5241434c45000000000000000000000000000000000000000061054d565b6000610eed8361054d565b60408051306024808301919091528251808303909101815260449091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fc4d66de80000000000000000000000000000000000000000000000000000000017905290915081906001600160a01b03821615156110d757604051610f76906112b5565b604051809103906000f080158015610f92573d6000803e3d6000fd5b509150816001600160a01b031663cf7a1d778530846040518463ffffffff1660e01b815260040180846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561102357818101518382015260200161100b565b50505050905090810190601f1680156110505780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b15801561107157600080fd5b505af1158015611085573d6000803e3d6000fd5b5050505061109385836111c4565b6040805186815290516001600160a01b038416917f1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438919081900360200190a26111bd565b604080517f4f1ef2860000000000000000000000000000000000000000000000000000000081526001600160a01b03868116600483019081526024830193845284516044840152845191861693634f1ef2869389938793929160640190602085019080838360005b8381101561115757818101518382015260200161113f565b50505050905090810190601f1680156111845780820380516001836020036101000a031916815260200191505b509350505050600060405180830381600087803b1580156111a457600080fd5b505af11580156111b8573d6000803e3d6000fd5b505050505b5050505050565b600091825260016020526040909120805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03909216919091179055565b3390565b6001600160a01b038116151561124d57604051600160e51b62461bcd028152600401808060200182810382526026815260200180611ce36026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b610a20806112c38339019056fe608060405234801561001057600080fd5b50610a00806100206000396000f3fe6080604052600436106100705760003560e01c80638f2839701161004e5780638f2839701461015e578063cf7a1d7714610191578063d1f5789414610250578063f851a4401461030657610070565b80633659cfe61461007a5780634f1ef286146100ad5780635c60da1b1461012d575b61007861031b565b005b34801561008657600080fd5b506100786004803603602081101561009d57600080fd5b50356001600160a01b0316610335565b610078600480360360408110156100c357600080fd5b6001600160a01b0382351691908101906040810160208201356401000000008111156100ee57600080fd5b82018360208201111561010057600080fd5b8035906020019184600183028401116401000000008311171561012257600080fd5b50909250905061036f565b34801561013957600080fd5b5061014261041e565b604080516001600160a01b039092168252519081900360200190f35b34801561016a57600080fd5b506100786004803603602081101561018157600080fd5b50356001600160a01b031661045b565b610078600480360360608110156101a757600080fd5b6001600160a01b0382358116926020810135909116918101906060810160408201356401000000008111156101db57600080fd5b8201836020820111156101ed57600080fd5b8035906020019184600183028401116401000000008311171561020f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610531945050505050565b6100786004803603604081101561026657600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561029157600080fd5b8201836020820111156102a357600080fd5b803590602001918460018302840111640100000000831117156102c557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506105c1945050505050565b34801561031257600080fd5b50610142610703565b61032361072e565b61033361032e61079f565b6107c4565b565b61033d6107e8565b6001600160a01b0316336001600160a01b031614156103645761035f8161080d565b61036c565b61036c61031b565b50565b6103776107e8565b6001600160a01b0316336001600160a01b03161415610411576103998361080d565b6000836001600160a01b031683836040518083838082843760405192019450600093509091505080830381855af49150503d80600081146103f6576040519150601f19603f3d011682016040523d82523d6000602084013e6103fb565b606091505b5050905080151561040b57600080fd5b50610419565b61041961031b565b505050565b60006104286107e8565b6001600160a01b0316336001600160a01b031614156104505761044961079f565b9050610458565b61045861031b565b90565b6104636107e8565b6001600160a01b0316336001600160a01b03161415610364576001600160a01b03811615156104dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001806109646036913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6105066107e8565b604080516001600160a01b03928316815291841660208301528051918290030190a161035f8161084d565b600061053b61079f565b6001600160a01b03161461054e57600080fd5b61055883826105c1565b604080517f656970313936372e70726f78792e61646d696e00000000000000000000000000815290519081900360130190207fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103600019909101146105b857fe5b6104198261084d565b60006105cb61079f565b6001600160a01b0316146105de57600080fd5b604080517f656970313936372e70726f78792e696d706c656d656e746174696f6e000000008152905190819003601c0190207f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6000199091011461063e57fe5b61064782610871565b8051156106ff576000826001600160a01b0316826040518082805190602001908083835b6020831061068a5780518252601f19909201916020918201910161066b565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146106ea576040519150601f19603f3d011682016040523d82523d6000602084013e6106ef565b606091505b5050905080151561041957600080fd5b5050565b600061070d6107e8565b6001600160a01b0316336001600160a01b03161415610450576104496107e8565b6107366107e8565b6001600160a01b0316331415610797576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001806109326032913960400191505060405180910390fd5b610333610333565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e8080156107e3573d6000f35b3d6000fd5b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b61081681610871565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b61087a816108f5565b15156108d1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b81526020018061099a603b913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47081158015906109295750808214155b94935050505056fe43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e2066726f6d207468652070726f78792061646d696e43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f787920746f20746865207a65726f206164647265737343616e6e6f742073657420612070726f787920696d706c656d656e746174696f6e20746f2061206e6f6e2d636f6e74726163742061646472657373a165627a7a72305820d487d975e5f3311a8ac0ada07858de7d0a7ebedea4657b56eb7294306768504900294f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a165627a7a723058208d686a822f621e1959a4c6bcd2d0cc0534923013d8787674473127bb97f61b090029"

// DeployLendingPoolAddressesProvider deploys a new Ethereum contract, binding an instance of LendingPoolAddressesProvider to it.
func DeployLendingPoolAddressesProvider(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LendingPoolAddressesProvider, error) {
	parsed, err := abi.JSON(strings.NewReader(LendingPoolAddressesProviderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LendingPoolAddressesProviderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LendingPoolAddressesProvider{LendingPoolAddressesProviderCaller: LendingPoolAddressesProviderCaller{contract: contract}, LendingPoolAddressesProviderTransactor: LendingPoolAddressesProviderTransactor{contract: contract}, LendingPoolAddressesProviderFilterer: LendingPoolAddressesProviderFilterer{contract: contract}}, nil
}

// LendingPoolAddressesProvider is an auto generated Go binding around an Ethereum contract.
type LendingPoolAddressesProvider struct {
	LendingPoolAddressesProviderCaller     // Read-only binding to the contract
	LendingPoolAddressesProviderTransactor // Write-only binding to the contract
	LendingPoolAddressesProviderFilterer   // Log filterer for contract events
}

// LendingPoolAddressesProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type LendingPoolAddressesProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolAddressesProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LendingPoolAddressesProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolAddressesProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LendingPoolAddressesProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolAddressesProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LendingPoolAddressesProviderSession struct {
	Contract     *LendingPoolAddressesProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// LendingPoolAddressesProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LendingPoolAddressesProviderCallerSession struct {
	Contract *LendingPoolAddressesProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// LendingPoolAddressesProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LendingPoolAddressesProviderTransactorSession struct {
	Contract     *LendingPoolAddressesProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// LendingPoolAddressesProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type LendingPoolAddressesProviderRaw struct {
	Contract *LendingPoolAddressesProvider // Generic contract binding to access the raw methods on
}

// LendingPoolAddressesProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LendingPoolAddressesProviderCallerRaw struct {
	Contract *LendingPoolAddressesProviderCaller // Generic read-only contract binding to access the raw methods on
}

// LendingPoolAddressesProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LendingPoolAddressesProviderTransactorRaw struct {
	Contract *LendingPoolAddressesProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLendingPoolAddressesProvider creates a new instance of LendingPoolAddressesProvider, bound to a specific deployed contract.
func NewLendingPoolAddressesProvider(address common.Address, backend bind.ContractBackend) (*LendingPoolAddressesProvider, error) {
	contract, err := bindLendingPoolAddressesProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProvider{LendingPoolAddressesProviderCaller: LendingPoolAddressesProviderCaller{contract: contract}, LendingPoolAddressesProviderTransactor: LendingPoolAddressesProviderTransactor{contract: contract}, LendingPoolAddressesProviderFilterer: LendingPoolAddressesProviderFilterer{contract: contract}}, nil
}

// NewLendingPoolAddressesProviderCaller creates a new read-only instance of LendingPoolAddressesProvider, bound to a specific deployed contract.
func NewLendingPoolAddressesProviderCaller(address common.Address, caller bind.ContractCaller) (*LendingPoolAddressesProviderCaller, error) {
	contract, err := bindLendingPoolAddressesProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderCaller{contract: contract}, nil
}

// NewLendingPoolAddressesProviderTransactor creates a new write-only instance of LendingPoolAddressesProvider, bound to a specific deployed contract.
func NewLendingPoolAddressesProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*LendingPoolAddressesProviderTransactor, error) {
	contract, err := bindLendingPoolAddressesProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderTransactor{contract: contract}, nil
}

// NewLendingPoolAddressesProviderFilterer creates a new log filterer instance of LendingPoolAddressesProvider, bound to a specific deployed contract.
func NewLendingPoolAddressesProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*LendingPoolAddressesProviderFilterer, error) {
	contract, err := bindLendingPoolAddressesProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderFilterer{contract: contract}, nil
}

// bindLendingPoolAddressesProvider binds a generic wrapper to an already deployed contract.
func bindLendingPoolAddressesProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LendingPoolAddressesProviderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingPoolAddressesProvider.Contract.LendingPoolAddressesProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.LendingPoolAddressesProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.LendingPoolAddressesProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingPoolAddressesProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _key) view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetAddress(opts *bind.CallOpts, _key [32]byte) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getAddress", _key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _key) view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetAddress(_key [32]byte) (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetAddress(&_LendingPoolAddressesProvider.CallOpts, _key)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 _key) view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetAddress(_key [32]byte) (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetAddress(&_LendingPoolAddressesProvider.CallOpts, _key)
}

// GetFeeProvider is a free data retrieval call binding the contract method 0xfbeefc3c.
//
// Solidity: function getFeeProvider() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetFeeProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getFeeProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeProvider is a free data retrieval call binding the contract method 0xfbeefc3c.
//
// Solidity: function getFeeProvider() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetFeeProvider() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetFeeProvider(&_LendingPoolAddressesProvider.CallOpts)
}

// GetFeeProvider is a free data retrieval call binding the contract method 0xfbeefc3c.
//
// Solidity: function getFeeProvider() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetFeeProvider() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetFeeProvider(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetLendingPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getLendingPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetLendingPool() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPool(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetLendingPool() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPool(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetLendingPoolConfigurator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getLendingPoolConfigurator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetLendingPoolConfigurator() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPoolConfigurator(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetLendingPoolConfigurator() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPoolConfigurator(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolCore is a free data retrieval call binding the contract method 0xed6ff760.
//
// Solidity: function getLendingPoolCore() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetLendingPoolCore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getLendingPoolCore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolCore is a free data retrieval call binding the contract method 0xed6ff760.
//
// Solidity: function getLendingPoolCore() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetLendingPoolCore() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPoolCore(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolCore is a free data retrieval call binding the contract method 0xed6ff760.
//
// Solidity: function getLendingPoolCore() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetLendingPoolCore() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPoolCore(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolDataProvider is a free data retrieval call binding the contract method 0x2f58b80d.
//
// Solidity: function getLendingPoolDataProvider() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetLendingPoolDataProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getLendingPoolDataProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolDataProvider is a free data retrieval call binding the contract method 0x2f58b80d.
//
// Solidity: function getLendingPoolDataProvider() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetLendingPoolDataProvider() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPoolDataProvider(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolDataProvider is a free data retrieval call binding the contract method 0x2f58b80d.
//
// Solidity: function getLendingPoolDataProvider() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetLendingPoolDataProvider() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPoolDataProvider(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolLiquidationManager is a free data retrieval call binding the contract method 0x5834eb9a.
//
// Solidity: function getLendingPoolLiquidationManager() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetLendingPoolLiquidationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getLendingPoolLiquidationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolLiquidationManager is a free data retrieval call binding the contract method 0x5834eb9a.
//
// Solidity: function getLendingPoolLiquidationManager() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetLendingPoolLiquidationManager() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPoolLiquidationManager(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolLiquidationManager is a free data retrieval call binding the contract method 0x5834eb9a.
//
// Solidity: function getLendingPoolLiquidationManager() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetLendingPoolLiquidationManager() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPoolLiquidationManager(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolManager is a free data retrieval call binding the contract method 0x33128d59.
//
// Solidity: function getLendingPoolManager() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetLendingPoolManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getLendingPoolManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolManager is a free data retrieval call binding the contract method 0x33128d59.
//
// Solidity: function getLendingPoolManager() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetLendingPoolManager() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPoolManager(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolManager is a free data retrieval call binding the contract method 0x33128d59.
//
// Solidity: function getLendingPoolManager() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetLendingPoolManager() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPoolManager(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolParametersProvider is a free data retrieval call binding the contract method 0x04061d8e.
//
// Solidity: function getLendingPoolParametersProvider() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetLendingPoolParametersProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getLendingPoolParametersProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolParametersProvider is a free data retrieval call binding the contract method 0x04061d8e.
//
// Solidity: function getLendingPoolParametersProvider() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetLendingPoolParametersProvider() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPoolParametersProvider(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolParametersProvider is a free data retrieval call binding the contract method 0x04061d8e.
//
// Solidity: function getLendingPoolParametersProvider() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetLendingPoolParametersProvider() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingPoolParametersProvider(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetLendingRateOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getLendingRateOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetLendingRateOracle() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingRateOracle(&_LendingPoolAddressesProvider.CallOpts)
}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetLendingRateOracle() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetLendingRateOracle(&_LendingPoolAddressesProvider.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetPriceOracle() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetPriceOracle(&_LendingPoolAddressesProvider.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetPriceOracle() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetPriceOracle(&_LendingPoolAddressesProvider.CallOpts)
}

// GetTokenDistributor is a free data retrieval call binding the contract method 0xee891296.
//
// Solidity: function getTokenDistributor() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) GetTokenDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "getTokenDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenDistributor is a free data retrieval call binding the contract method 0xee891296.
//
// Solidity: function getTokenDistributor() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) GetTokenDistributor() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetTokenDistributor(&_LendingPoolAddressesProvider.CallOpts)
}

// GetTokenDistributor is a free data retrieval call binding the contract method 0xee891296.
//
// Solidity: function getTokenDistributor() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) GetTokenDistributor() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.GetTokenDistributor(&_LendingPoolAddressesProvider.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) IsOwner() (bool, error) {
	return _LendingPoolAddressesProvider.Contract.IsOwner(&_LendingPoolAddressesProvider.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) IsOwner() (bool, error) {
	return _LendingPoolAddressesProvider.Contract.IsOwner(&_LendingPoolAddressesProvider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendingPoolAddressesProvider.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) Owner() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.Owner(&_LendingPoolAddressesProvider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderCallerSession) Owner() (common.Address, error) {
	return _LendingPoolAddressesProvider.Contract.Owner(&_LendingPoolAddressesProvider.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) RenounceOwnership() (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.RenounceOwnership(&_LendingPoolAddressesProvider.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.RenounceOwnership(&_LendingPoolAddressesProvider.TransactOpts)
}

// SetFeeProviderImpl is a paid mutator transaction binding the contract method 0x2a62c636.
//
// Solidity: function setFeeProviderImpl(address _feeProvider) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetFeeProviderImpl(opts *bind.TransactOpts, _feeProvider common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setFeeProviderImpl", _feeProvider)
}

// SetFeeProviderImpl is a paid mutator transaction binding the contract method 0x2a62c636.
//
// Solidity: function setFeeProviderImpl(address _feeProvider) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetFeeProviderImpl(_feeProvider common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetFeeProviderImpl(&_LendingPoolAddressesProvider.TransactOpts, _feeProvider)
}

// SetFeeProviderImpl is a paid mutator transaction binding the contract method 0x2a62c636.
//
// Solidity: function setFeeProviderImpl(address _feeProvider) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetFeeProviderImpl(_feeProvider common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetFeeProviderImpl(&_LendingPoolAddressesProvider.TransactOpts, _feeProvider)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address _configurator) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetLendingPoolConfiguratorImpl(opts *bind.TransactOpts, _configurator common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setLendingPoolConfiguratorImpl", _configurator)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address _configurator) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetLendingPoolConfiguratorImpl(_configurator common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolConfiguratorImpl(&_LendingPoolAddressesProvider.TransactOpts, _configurator)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address _configurator) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetLendingPoolConfiguratorImpl(_configurator common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolConfiguratorImpl(&_LendingPoolAddressesProvider.TransactOpts, _configurator)
}

// SetLendingPoolCoreImpl is a paid mutator transaction binding the contract method 0x1c827204.
//
// Solidity: function setLendingPoolCoreImpl(address _lendingPoolCore) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetLendingPoolCoreImpl(opts *bind.TransactOpts, _lendingPoolCore common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setLendingPoolCoreImpl", _lendingPoolCore)
}

// SetLendingPoolCoreImpl is a paid mutator transaction binding the contract method 0x1c827204.
//
// Solidity: function setLendingPoolCoreImpl(address _lendingPoolCore) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetLendingPoolCoreImpl(_lendingPoolCore common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolCoreImpl(&_LendingPoolAddressesProvider.TransactOpts, _lendingPoolCore)
}

// SetLendingPoolCoreImpl is a paid mutator transaction binding the contract method 0x1c827204.
//
// Solidity: function setLendingPoolCoreImpl(address _lendingPoolCore) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetLendingPoolCoreImpl(_lendingPoolCore common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolCoreImpl(&_LendingPoolAddressesProvider.TransactOpts, _lendingPoolCore)
}

// SetLendingPoolDataProviderImpl is a paid mutator transaction binding the contract method 0xbfedc103.
//
// Solidity: function setLendingPoolDataProviderImpl(address _provider) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetLendingPoolDataProviderImpl(opts *bind.TransactOpts, _provider common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setLendingPoolDataProviderImpl", _provider)
}

// SetLendingPoolDataProviderImpl is a paid mutator transaction binding the contract method 0xbfedc103.
//
// Solidity: function setLendingPoolDataProviderImpl(address _provider) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetLendingPoolDataProviderImpl(_provider common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolDataProviderImpl(&_LendingPoolAddressesProvider.TransactOpts, _provider)
}

// SetLendingPoolDataProviderImpl is a paid mutator transaction binding the contract method 0xbfedc103.
//
// Solidity: function setLendingPoolDataProviderImpl(address _provider) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetLendingPoolDataProviderImpl(_provider common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolDataProviderImpl(&_LendingPoolAddressesProvider.TransactOpts, _provider)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address _pool) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetLendingPoolImpl(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setLendingPoolImpl", _pool)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address _pool) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetLendingPoolImpl(_pool common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolImpl(&_LendingPoolAddressesProvider.TransactOpts, _pool)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address _pool) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetLendingPoolImpl(_pool common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolImpl(&_LendingPoolAddressesProvider.TransactOpts, _pool)
}

// SetLendingPoolLiquidationManager is a paid mutator transaction binding the contract method 0x44ce375b.
//
// Solidity: function setLendingPoolLiquidationManager(address _manager) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetLendingPoolLiquidationManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setLendingPoolLiquidationManager", _manager)
}

// SetLendingPoolLiquidationManager is a paid mutator transaction binding the contract method 0x44ce375b.
//
// Solidity: function setLendingPoolLiquidationManager(address _manager) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetLendingPoolLiquidationManager(_manager common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolLiquidationManager(&_LendingPoolAddressesProvider.TransactOpts, _manager)
}

// SetLendingPoolLiquidationManager is a paid mutator transaction binding the contract method 0x44ce375b.
//
// Solidity: function setLendingPoolLiquidationManager(address _manager) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetLendingPoolLiquidationManager(_manager common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolLiquidationManager(&_LendingPoolAddressesProvider.TransactOpts, _manager)
}

// SetLendingPoolManager is a paid mutator transaction binding the contract method 0x40fdcadc.
//
// Solidity: function setLendingPoolManager(address _lendingPoolManager) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetLendingPoolManager(opts *bind.TransactOpts, _lendingPoolManager common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setLendingPoolManager", _lendingPoolManager)
}

// SetLendingPoolManager is a paid mutator transaction binding the contract method 0x40fdcadc.
//
// Solidity: function setLendingPoolManager(address _lendingPoolManager) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetLendingPoolManager(_lendingPoolManager common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolManager(&_LendingPoolAddressesProvider.TransactOpts, _lendingPoolManager)
}

// SetLendingPoolManager is a paid mutator transaction binding the contract method 0x40fdcadc.
//
// Solidity: function setLendingPoolManager(address _lendingPoolManager) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetLendingPoolManager(_lendingPoolManager common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolManager(&_LendingPoolAddressesProvider.TransactOpts, _lendingPoolManager)
}

// SetLendingPoolParametersProviderImpl is a paid mutator transaction binding the contract method 0xa5eface2.
//
// Solidity: function setLendingPoolParametersProviderImpl(address _parametersProvider) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetLendingPoolParametersProviderImpl(opts *bind.TransactOpts, _parametersProvider common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setLendingPoolParametersProviderImpl", _parametersProvider)
}

// SetLendingPoolParametersProviderImpl is a paid mutator transaction binding the contract method 0xa5eface2.
//
// Solidity: function setLendingPoolParametersProviderImpl(address _parametersProvider) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetLendingPoolParametersProviderImpl(_parametersProvider common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolParametersProviderImpl(&_LendingPoolAddressesProvider.TransactOpts, _parametersProvider)
}

// SetLendingPoolParametersProviderImpl is a paid mutator transaction binding the contract method 0xa5eface2.
//
// Solidity: function setLendingPoolParametersProviderImpl(address _parametersProvider) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetLendingPoolParametersProviderImpl(_parametersProvider common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingPoolParametersProviderImpl(&_LendingPoolAddressesProvider.TransactOpts, _parametersProvider)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address _lendingRateOracle) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetLendingRateOracle(opts *bind.TransactOpts, _lendingRateOracle common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setLendingRateOracle", _lendingRateOracle)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address _lendingRateOracle) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetLendingRateOracle(_lendingRateOracle common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingRateOracle(&_LendingPoolAddressesProvider.TransactOpts, _lendingRateOracle)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address _lendingRateOracle) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetLendingRateOracle(_lendingRateOracle common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetLendingRateOracle(&_LendingPoolAddressesProvider.TransactOpts, _lendingRateOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _priceOracle) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetPriceOracle(opts *bind.TransactOpts, _priceOracle common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setPriceOracle", _priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _priceOracle) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetPriceOracle(_priceOracle common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetPriceOracle(&_LendingPoolAddressesProvider.TransactOpts, _priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address _priceOracle) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetPriceOracle(_priceOracle common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetPriceOracle(&_LendingPoolAddressesProvider.TransactOpts, _priceOracle)
}

// SetTokenDistributor is a paid mutator transaction binding the contract method 0x38280e6b.
//
// Solidity: function setTokenDistributor(address _tokenDistributor) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) SetTokenDistributor(opts *bind.TransactOpts, _tokenDistributor common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "setTokenDistributor", _tokenDistributor)
}

// SetTokenDistributor is a paid mutator transaction binding the contract method 0x38280e6b.
//
// Solidity: function setTokenDistributor(address _tokenDistributor) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) SetTokenDistributor(_tokenDistributor common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetTokenDistributor(&_LendingPoolAddressesProvider.TransactOpts, _tokenDistributor)
}

// SetTokenDistributor is a paid mutator transaction binding the contract method 0x38280e6b.
//
// Solidity: function setTokenDistributor(address _tokenDistributor) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) SetTokenDistributor(_tokenDistributor common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.SetTokenDistributor(&_LendingPoolAddressesProvider.TransactOpts, _tokenDistributor)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.TransferOwnership(&_LendingPoolAddressesProvider.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LendingPoolAddressesProvider.Contract.TransferOwnership(&_LendingPoolAddressesProvider.TransactOpts, newOwner)
}

// LendingPoolAddressesProviderEthereumAddressUpdatedIterator is returned from FilterEthereumAddressUpdated and is used to iterate over the raw logs and unpacked data for EthereumAddressUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderEthereumAddressUpdatedIterator struct {
	Event *LendingPoolAddressesProviderEthereumAddressUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderEthereumAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderEthereumAddressUpdated)
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
		it.Event = new(LendingPoolAddressesProviderEthereumAddressUpdated)
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
func (it *LendingPoolAddressesProviderEthereumAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderEthereumAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderEthereumAddressUpdated represents a EthereumAddressUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderEthereumAddressUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEthereumAddressUpdated is a free log retrieval operation binding the contract event 0x1941c4aa9ef07ed299ed253b4224020832574525e8db1f3f955c57f395aef829.
//
// Solidity: event EthereumAddressUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterEthereumAddressUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderEthereumAddressUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "EthereumAddressUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderEthereumAddressUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "EthereumAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchEthereumAddressUpdated is a free log subscription operation binding the contract event 0x1941c4aa9ef07ed299ed253b4224020832574525e8db1f3f955c57f395aef829.
//
// Solidity: event EthereumAddressUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchEthereumAddressUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderEthereumAddressUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "EthereumAddressUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderEthereumAddressUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "EthereumAddressUpdated", log); err != nil {
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

// ParseEthereumAddressUpdated is a log parse operation binding the contract event 0x1941c4aa9ef07ed299ed253b4224020832574525e8db1f3f955c57f395aef829.
//
// Solidity: event EthereumAddressUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseEthereumAddressUpdated(log types.Log) (*LendingPoolAddressesProviderEthereumAddressUpdated, error) {
	event := new(LendingPoolAddressesProviderEthereumAddressUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "EthereumAddressUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingPoolAddressesProviderFeeProviderUpdatedIterator is returned from FilterFeeProviderUpdated and is used to iterate over the raw logs and unpacked data for FeeProviderUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderFeeProviderUpdatedIterator struct {
	Event *LendingPoolAddressesProviderFeeProviderUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderFeeProviderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderFeeProviderUpdated)
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
		it.Event = new(LendingPoolAddressesProviderFeeProviderUpdated)
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
func (it *LendingPoolAddressesProviderFeeProviderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderFeeProviderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderFeeProviderUpdated represents a FeeProviderUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderFeeProviderUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeProviderUpdated is a free log retrieval operation binding the contract event 0x18e1a55b8eff9c93921eecfa1462d6a8cbb80b3988db3eb14c039e43fdb22661.
//
// Solidity: event FeeProviderUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterFeeProviderUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderFeeProviderUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "FeeProviderUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderFeeProviderUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "FeeProviderUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeProviderUpdated is a free log subscription operation binding the contract event 0x18e1a55b8eff9c93921eecfa1462d6a8cbb80b3988db3eb14c039e43fdb22661.
//
// Solidity: event FeeProviderUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchFeeProviderUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderFeeProviderUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "FeeProviderUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderFeeProviderUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "FeeProviderUpdated", log); err != nil {
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

// ParseFeeProviderUpdated is a log parse operation binding the contract event 0x18e1a55b8eff9c93921eecfa1462d6a8cbb80b3988db3eb14c039e43fdb22661.
//
// Solidity: event FeeProviderUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseFeeProviderUpdated(log types.Log) (*LendingPoolAddressesProviderFeeProviderUpdated, error) {
	event := new(LendingPoolAddressesProviderFeeProviderUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "FeeProviderUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator is returned from FilterLendingPoolConfiguratorUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolConfiguratorUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator struct {
	Event *LendingPoolAddressesProviderLendingPoolConfiguratorUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderLendingPoolConfiguratorUpdated)
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
		it.Event = new(LendingPoolAddressesProviderLendingPoolConfiguratorUpdated)
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
func (it *LendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderLendingPoolConfiguratorUpdated represents a LendingPoolConfiguratorUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolConfiguratorUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolConfiguratorUpdated is a free log retrieval operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterLendingPoolConfiguratorUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "LendingPoolConfiguratorUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "LendingPoolConfiguratorUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolConfiguratorUpdated is a free log subscription operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchLendingPoolConfiguratorUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderLendingPoolConfiguratorUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "LendingPoolConfiguratorUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderLendingPoolConfiguratorUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolConfiguratorUpdated", log); err != nil {
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

// ParseLendingPoolConfiguratorUpdated is a log parse operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseLendingPoolConfiguratorUpdated(log types.Log) (*LendingPoolAddressesProviderLendingPoolConfiguratorUpdated, error) {
	event := new(LendingPoolAddressesProviderLendingPoolConfiguratorUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolConfiguratorUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingPoolAddressesProviderLendingPoolCoreUpdatedIterator is returned from FilterLendingPoolCoreUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolCoreUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolCoreUpdatedIterator struct {
	Event *LendingPoolAddressesProviderLendingPoolCoreUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderLendingPoolCoreUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderLendingPoolCoreUpdated)
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
		it.Event = new(LendingPoolAddressesProviderLendingPoolCoreUpdated)
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
func (it *LendingPoolAddressesProviderLendingPoolCoreUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderLendingPoolCoreUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderLendingPoolCoreUpdated represents a LendingPoolCoreUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolCoreUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolCoreUpdated is a free log retrieval operation binding the contract event 0x71c226bb2879778ca1298196bf7cc1256baee9d05b496c31ee627d35a471b5b7.
//
// Solidity: event LendingPoolCoreUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterLendingPoolCoreUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderLendingPoolCoreUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "LendingPoolCoreUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderLendingPoolCoreUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "LendingPoolCoreUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolCoreUpdated is a free log subscription operation binding the contract event 0x71c226bb2879778ca1298196bf7cc1256baee9d05b496c31ee627d35a471b5b7.
//
// Solidity: event LendingPoolCoreUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchLendingPoolCoreUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderLendingPoolCoreUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "LendingPoolCoreUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderLendingPoolCoreUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolCoreUpdated", log); err != nil {
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

// ParseLendingPoolCoreUpdated is a log parse operation binding the contract event 0x71c226bb2879778ca1298196bf7cc1256baee9d05b496c31ee627d35a471b5b7.
//
// Solidity: event LendingPoolCoreUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseLendingPoolCoreUpdated(log types.Log) (*LendingPoolAddressesProviderLendingPoolCoreUpdated, error) {
	event := new(LendingPoolAddressesProviderLendingPoolCoreUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolCoreUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingPoolAddressesProviderLendingPoolDataProviderUpdatedIterator is returned from FilterLendingPoolDataProviderUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolDataProviderUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolDataProviderUpdatedIterator struct {
	Event *LendingPoolAddressesProviderLendingPoolDataProviderUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderLendingPoolDataProviderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderLendingPoolDataProviderUpdated)
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
		it.Event = new(LendingPoolAddressesProviderLendingPoolDataProviderUpdated)
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
func (it *LendingPoolAddressesProviderLendingPoolDataProviderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderLendingPoolDataProviderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderLendingPoolDataProviderUpdated represents a LendingPoolDataProviderUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolDataProviderUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolDataProviderUpdated is a free log retrieval operation binding the contract event 0x07890d7d10db37434d76ee4f2928fb2dc66227c3b3391206aced4f7bcb59cdb0.
//
// Solidity: event LendingPoolDataProviderUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterLendingPoolDataProviderUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderLendingPoolDataProviderUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "LendingPoolDataProviderUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderLendingPoolDataProviderUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "LendingPoolDataProviderUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolDataProviderUpdated is a free log subscription operation binding the contract event 0x07890d7d10db37434d76ee4f2928fb2dc66227c3b3391206aced4f7bcb59cdb0.
//
// Solidity: event LendingPoolDataProviderUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchLendingPoolDataProviderUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderLendingPoolDataProviderUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "LendingPoolDataProviderUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderLendingPoolDataProviderUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolDataProviderUpdated", log); err != nil {
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

// ParseLendingPoolDataProviderUpdated is a log parse operation binding the contract event 0x07890d7d10db37434d76ee4f2928fb2dc66227c3b3391206aced4f7bcb59cdb0.
//
// Solidity: event LendingPoolDataProviderUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseLendingPoolDataProviderUpdated(log types.Log) (*LendingPoolAddressesProviderLendingPoolDataProviderUpdated, error) {
	event := new(LendingPoolAddressesProviderLendingPoolDataProviderUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolDataProviderUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingPoolAddressesProviderLendingPoolLiquidationManagerUpdatedIterator is returned from FilterLendingPoolLiquidationManagerUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolLiquidationManagerUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolLiquidationManagerUpdatedIterator struct {
	Event *LendingPoolAddressesProviderLendingPoolLiquidationManagerUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderLendingPoolLiquidationManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderLendingPoolLiquidationManagerUpdated)
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
		it.Event = new(LendingPoolAddressesProviderLendingPoolLiquidationManagerUpdated)
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
func (it *LendingPoolAddressesProviderLendingPoolLiquidationManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderLendingPoolLiquidationManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderLendingPoolLiquidationManagerUpdated represents a LendingPoolLiquidationManagerUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolLiquidationManagerUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolLiquidationManagerUpdated is a free log retrieval operation binding the contract event 0x1a76cb769b814bc038223da86932b099b20aae03473683e6d98f5c3554e26064.
//
// Solidity: event LendingPoolLiquidationManagerUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterLendingPoolLiquidationManagerUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderLendingPoolLiquidationManagerUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "LendingPoolLiquidationManagerUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderLendingPoolLiquidationManagerUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "LendingPoolLiquidationManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolLiquidationManagerUpdated is a free log subscription operation binding the contract event 0x1a76cb769b814bc038223da86932b099b20aae03473683e6d98f5c3554e26064.
//
// Solidity: event LendingPoolLiquidationManagerUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchLendingPoolLiquidationManagerUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderLendingPoolLiquidationManagerUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "LendingPoolLiquidationManagerUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderLendingPoolLiquidationManagerUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolLiquidationManagerUpdated", log); err != nil {
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

// ParseLendingPoolLiquidationManagerUpdated is a log parse operation binding the contract event 0x1a76cb769b814bc038223da86932b099b20aae03473683e6d98f5c3554e26064.
//
// Solidity: event LendingPoolLiquidationManagerUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseLendingPoolLiquidationManagerUpdated(log types.Log) (*LendingPoolAddressesProviderLendingPoolLiquidationManagerUpdated, error) {
	event := new(LendingPoolAddressesProviderLendingPoolLiquidationManagerUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolLiquidationManagerUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingPoolAddressesProviderLendingPoolManagerUpdatedIterator is returned from FilterLendingPoolManagerUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolManagerUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolManagerUpdatedIterator struct {
	Event *LendingPoolAddressesProviderLendingPoolManagerUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderLendingPoolManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderLendingPoolManagerUpdated)
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
		it.Event = new(LendingPoolAddressesProviderLendingPoolManagerUpdated)
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
func (it *LendingPoolAddressesProviderLendingPoolManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderLendingPoolManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderLendingPoolManagerUpdated represents a LendingPoolManagerUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolManagerUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolManagerUpdated is a free log retrieval operation binding the contract event 0xd5280c51185a38d36f7a0f5e56cac6248312bb70d0974782fa5a595127e0e08e.
//
// Solidity: event LendingPoolManagerUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterLendingPoolManagerUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderLendingPoolManagerUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "LendingPoolManagerUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderLendingPoolManagerUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "LendingPoolManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolManagerUpdated is a free log subscription operation binding the contract event 0xd5280c51185a38d36f7a0f5e56cac6248312bb70d0974782fa5a595127e0e08e.
//
// Solidity: event LendingPoolManagerUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchLendingPoolManagerUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderLendingPoolManagerUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "LendingPoolManagerUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderLendingPoolManagerUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolManagerUpdated", log); err != nil {
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

// ParseLendingPoolManagerUpdated is a log parse operation binding the contract event 0xd5280c51185a38d36f7a0f5e56cac6248312bb70d0974782fa5a595127e0e08e.
//
// Solidity: event LendingPoolManagerUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseLendingPoolManagerUpdated(log types.Log) (*LendingPoolAddressesProviderLendingPoolManagerUpdated, error) {
	event := new(LendingPoolAddressesProviderLendingPoolManagerUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolManagerUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingPoolAddressesProviderLendingPoolParametersProviderUpdatedIterator is returned from FilterLendingPoolParametersProviderUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolParametersProviderUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolParametersProviderUpdatedIterator struct {
	Event *LendingPoolAddressesProviderLendingPoolParametersProviderUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderLendingPoolParametersProviderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderLendingPoolParametersProviderUpdated)
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
		it.Event = new(LendingPoolAddressesProviderLendingPoolParametersProviderUpdated)
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
func (it *LendingPoolAddressesProviderLendingPoolParametersProviderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderLendingPoolParametersProviderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderLendingPoolParametersProviderUpdated represents a LendingPoolParametersProviderUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolParametersProviderUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolParametersProviderUpdated is a free log retrieval operation binding the contract event 0xce16ea9b2fd7cadddd0f7359971028f50b5ddba33dfae1a9bdea956fabb1e280.
//
// Solidity: event LendingPoolParametersProviderUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterLendingPoolParametersProviderUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderLendingPoolParametersProviderUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "LendingPoolParametersProviderUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderLendingPoolParametersProviderUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "LendingPoolParametersProviderUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolParametersProviderUpdated is a free log subscription operation binding the contract event 0xce16ea9b2fd7cadddd0f7359971028f50b5ddba33dfae1a9bdea956fabb1e280.
//
// Solidity: event LendingPoolParametersProviderUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchLendingPoolParametersProviderUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderLendingPoolParametersProviderUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "LendingPoolParametersProviderUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderLendingPoolParametersProviderUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolParametersProviderUpdated", log); err != nil {
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

// ParseLendingPoolParametersProviderUpdated is a log parse operation binding the contract event 0xce16ea9b2fd7cadddd0f7359971028f50b5ddba33dfae1a9bdea956fabb1e280.
//
// Solidity: event LendingPoolParametersProviderUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseLendingPoolParametersProviderUpdated(log types.Log) (*LendingPoolAddressesProviderLendingPoolParametersProviderUpdated, error) {
	event := new(LendingPoolAddressesProviderLendingPoolParametersProviderUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolParametersProviderUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingPoolAddressesProviderLendingPoolUpdatedIterator is returned from FilterLendingPoolUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolUpdatedIterator struct {
	Event *LendingPoolAddressesProviderLendingPoolUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderLendingPoolUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderLendingPoolUpdated)
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
		it.Event = new(LendingPoolAddressesProviderLendingPoolUpdated)
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
func (it *LendingPoolAddressesProviderLendingPoolUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderLendingPoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderLendingPoolUpdated represents a LendingPoolUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingPoolUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolUpdated is a free log retrieval operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterLendingPoolUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderLendingPoolUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "LendingPoolUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderLendingPoolUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "LendingPoolUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolUpdated is a free log subscription operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchLendingPoolUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderLendingPoolUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "LendingPoolUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderLendingPoolUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolUpdated", log); err != nil {
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

// ParseLendingPoolUpdated is a log parse operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseLendingPoolUpdated(log types.Log) (*LendingPoolAddressesProviderLendingPoolUpdated, error) {
	event := new(LendingPoolAddressesProviderLendingPoolUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingPoolAddressesProviderLendingRateOracleUpdatedIterator is returned from FilterLendingRateOracleUpdated and is used to iterate over the raw logs and unpacked data for LendingRateOracleUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingRateOracleUpdatedIterator struct {
	Event *LendingPoolAddressesProviderLendingRateOracleUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderLendingRateOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderLendingRateOracleUpdated)
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
		it.Event = new(LendingPoolAddressesProviderLendingRateOracleUpdated)
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
func (it *LendingPoolAddressesProviderLendingRateOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderLendingRateOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderLendingRateOracleUpdated represents a LendingRateOracleUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderLendingRateOracleUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingRateOracleUpdated is a free log retrieval operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterLendingRateOracleUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderLendingRateOracleUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "LendingRateOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderLendingRateOracleUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "LendingRateOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingRateOracleUpdated is a free log subscription operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchLendingRateOracleUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderLendingRateOracleUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "LendingRateOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderLendingRateOracleUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingRateOracleUpdated", log); err != nil {
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

// ParseLendingRateOracleUpdated is a log parse operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseLendingRateOracleUpdated(log types.Log) (*LendingPoolAddressesProviderLendingRateOracleUpdated, error) {
	event := new(LendingPoolAddressesProviderLendingRateOracleUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "LendingRateOracleUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingPoolAddressesProviderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderOwnershipTransferredIterator struct {
	Event *LendingPoolAddressesProviderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderOwnershipTransferred)
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
		it.Event = new(LendingPoolAddressesProviderOwnershipTransferred)
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
func (it *LendingPoolAddressesProviderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderOwnershipTransferred represents a OwnershipTransferred event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LendingPoolAddressesProviderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderOwnershipTransferredIterator{contract: _LendingPoolAddressesProvider.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderOwnershipTransferred)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseOwnershipTransferred(log types.Log) (*LendingPoolAddressesProviderOwnershipTransferred, error) {
	event := new(LendingPoolAddressesProviderOwnershipTransferred)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingPoolAddressesProviderPriceOracleUpdatedIterator is returned from FilterPriceOracleUpdated and is used to iterate over the raw logs and unpacked data for PriceOracleUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderPriceOracleUpdatedIterator struct {
	Event *LendingPoolAddressesProviderPriceOracleUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderPriceOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderPriceOracleUpdated)
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
		it.Event = new(LendingPoolAddressesProviderPriceOracleUpdated)
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
func (it *LendingPoolAddressesProviderPriceOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderPriceOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderPriceOracleUpdated represents a PriceOracleUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderPriceOracleUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleUpdated is a free log retrieval operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterPriceOracleUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderPriceOracleUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "PriceOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderPriceOracleUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "PriceOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceOracleUpdated is a free log subscription operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchPriceOracleUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderPriceOracleUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "PriceOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderPriceOracleUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
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

// ParsePriceOracleUpdated is a log parse operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParsePriceOracleUpdated(log types.Log) (*LendingPoolAddressesProviderPriceOracleUpdated, error) {
	event := new(LendingPoolAddressesProviderPriceOracleUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingPoolAddressesProviderProxyCreatedIterator is returned from FilterProxyCreated and is used to iterate over the raw logs and unpacked data for ProxyCreated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderProxyCreatedIterator struct {
	Event *LendingPoolAddressesProviderProxyCreated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderProxyCreated)
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
		it.Event = new(LendingPoolAddressesProviderProxyCreated)
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
func (it *LendingPoolAddressesProviderProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderProxyCreated represents a ProxyCreated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderProxyCreated struct {
	Id         [32]byte
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProxyCreated is a free log retrieval operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterProxyCreated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderProxyCreatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "ProxyCreated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderProxyCreatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "ProxyCreated", logs: logs, sub: sub}, nil
}

// WatchProxyCreated is a free log subscription operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchProxyCreated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderProxyCreated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "ProxyCreated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderProxyCreated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
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

// ParseProxyCreated is a log parse operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseProxyCreated(log types.Log) (*LendingPoolAddressesProviderProxyCreated, error) {
	event := new(LendingPoolAddressesProviderProxyCreated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingPoolAddressesProviderTokenDistributorUpdatedIterator is returned from FilterTokenDistributorUpdated and is used to iterate over the raw logs and unpacked data for TokenDistributorUpdated events raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderTokenDistributorUpdatedIterator struct {
	Event *LendingPoolAddressesProviderTokenDistributorUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolAddressesProviderTokenDistributorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolAddressesProviderTokenDistributorUpdated)
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
		it.Event = new(LendingPoolAddressesProviderTokenDistributorUpdated)
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
func (it *LendingPoolAddressesProviderTokenDistributorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolAddressesProviderTokenDistributorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolAddressesProviderTokenDistributorUpdated represents a TokenDistributorUpdated event raised by the LendingPoolAddressesProvider contract.
type LendingPoolAddressesProviderTokenDistributorUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenDistributorUpdated is a free log retrieval operation binding the contract event 0xa8b48a56ec01f81c3615a21ec43e16b925c26293e0801cf6330427f2a687f053.
//
// Solidity: event TokenDistributorUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) FilterTokenDistributorUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*LendingPoolAddressesProviderTokenDistributorUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.FilterLogs(opts, "TokenDistributorUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolAddressesProviderTokenDistributorUpdatedIterator{contract: _LendingPoolAddressesProvider.contract, event: "TokenDistributorUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenDistributorUpdated is a free log subscription operation binding the contract event 0xa8b48a56ec01f81c3615a21ec43e16b925c26293e0801cf6330427f2a687f053.
//
// Solidity: event TokenDistributorUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) WatchTokenDistributorUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolAddressesProviderTokenDistributorUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _LendingPoolAddressesProvider.contract.WatchLogs(opts, "TokenDistributorUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolAddressesProviderTokenDistributorUpdated)
				if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "TokenDistributorUpdated", log); err != nil {
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

// ParseTokenDistributorUpdated is a log parse operation binding the contract event 0xa8b48a56ec01f81c3615a21ec43e16b925c26293e0801cf6330427f2a687f053.
//
// Solidity: event TokenDistributorUpdated(address indexed newAddress)
func (_LendingPoolAddressesProvider *LendingPoolAddressesProviderFilterer) ParseTokenDistributorUpdated(log types.Log) (*LendingPoolAddressesProviderTokenDistributorUpdated, error) {
	event := new(LendingPoolAddressesProviderTokenDistributorUpdated)
	if err := _LendingPoolAddressesProvider.contract.UnpackLog(event, "TokenDistributorUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
