// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ousd

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

// OusdABI is the input ABI used to generate the binding from.
const OusdABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimGovernance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGovernor\",\"type\":\"address\"}],\"name\":\"transferGovernance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalCredits\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creditsPerToken\",\"type\":\"uint256\"}],\"name\":\"TotalSupplyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousGovernor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"PendingGovernorshipTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousGovernor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"GovernorshipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nameArg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbolArg\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_vaultAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"creditsBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newTotalSupply\",\"type\":\"uint256\"}],\"name\":\"changeSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OusdBin is the compiled bytecode used for deploying new contracts.
var OusdBin = "0x60806040526000603e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610055336100c260201b60201c565b6100636100f160201b60201c565b73ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fc7c0c772add429241571afb3805861fb3cfa2af374534088b76cdb4325a87e9a60405160405180910390a3610122565b60007f7bea13895fa79d2831e0a9e28edede30099005a50d652d8957cf8a607ee6ca4a60001b90508181555050565b6000807f7bea13895fa79d2831e0a9e28edede30099005a50d652d8957cf8a607ee6ca4a60001b9050805491505090565b61286980620001326000396000f3fe608060405234801561001057600080fd5b50600436106101375760003560e01c8063430bf08a116100b8578063a457c2d71161007c578063a457c2d714610698578063a9059cbb146106fe578063c7af335214610764578063d38bfff414610786578063dd62ed3e146107ca578063f9854bfc1461084257610137565b8063430bf08a1461051b5780635d36b1901461056557806370a082311461056f57806395d89b41146105c75780639dc29fac1461064a57610137565b806323b872dd116100ff57806323b872dd1461037b578063313ce56714610401578063395093511461042557806339a7919f1461048b57806340c10f19146104cd57610137565b806306fdde031461013c578063077f224a146101bf578063095ea7b3146102ad5780630c340a241461031357806318160ddd1461035d575b600080fd5b61014461089a565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610184578082015181840152602081019050610169565b50505050905090810190601f1680156101b15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102ab600480360360608110156101d557600080fd5b81019080803590602001906401000000008111156101f257600080fd5b82018360208201111561020457600080fd5b8035906020019184600183028401116401000000008311171561022657600080fd5b90919293919293908035906020019064010000000081111561024757600080fd5b82018360208201111561025957600080fd5b8035906020019184600183028401116401000000008311171561027b57600080fd5b9091929391929390803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061093c565b005b6102f9600480360360408110156102c357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610bac565b604051808215151515815260200191505060405180910390f35b61031b610c9e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610365610cad565b6040518082815260200191505060405180910390f35b6103e76004803603606081101561039157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610cb7565b604051808215151515815260200191505060405180910390f35b610409610edc565b604051808260ff1660ff16815260200191505060405180910390f35b6104716004803603604081101561043b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610ef3565b604051808215151515815260200191505060405180910390f35b6104b7600480360360208110156104a157600080fd5b81019080803590602001909291905050506110ef565b6040518082815260200191505060405180910390f35b610519600480360360408110156104e357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611343565b005b610523611414565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61056d61143a565b005b6105b16004803603602081101561058557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506114d0565b6040518082815260200191505060405180910390f35b6105cf611541565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561060f5780820151818401526020810190506105f4565b50505050905090810190601f16801561063c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6106966004803603604081101561066057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506115e3565b005b6106e4600480360360408110156106ae57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506116b4565b604051808215151515815260200191505060405180910390f35b61074a6004803603604081101561071457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611944565b604051808215151515815260200191505060405180910390f35b61076c611a58565b604051808215151515815260200191505060405180910390f35b6107c86004803603602081101561079c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611a95565b005b61082c600480360360408110156107e057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611b7c565b6040518082815260200191505060405180910390f35b6108846004803603602081101561085857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611c03565b6040518082815260200191505060405180910390f35b606060368054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109325780601f1061090757610100808354040283529160200191610932565b820191906000526020600020905b81548152906001019060200180831161091557829003601f168201915b5050505050905090565b610944611a58565b6109b6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f43616c6c6572206973206e6f742074686520476f7665726e6f7200000000000081525060200191505060405180910390fd5b600060019054906101000a900460ff16806109d557506109d4611c4c565b5b806109ec57506000809054906101000a900460ff16155b610a41576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001806127d7602e913960400191505060405180910390fd5b60008060019054906101000a900460ff161590508015610a91576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b610b2386868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505085858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050611c63565b60006039819055506000603a81905550670de0b6b3a7640000603b8190555081603e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015610ba45760008060016101000a81548160ff0219169083151502179055505b505050505050565b600081603d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b6000610ca8611c73565b905090565b6000603954905090565b6000610d4882603d60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611ca490919063ffffffff16565b603d60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000610dd48584611cee565b9050610e2881603c60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611e7390919063ffffffff16565b603c60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040518082815260200191505060405180910390a360019150509392505050565b6000603860009054906101000a900460ff16905090565b6000610f8482603d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611e7390919063ffffffff16565b603d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925603d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040518082815260200191505060405180910390a36001905092915050565b60003373ffffffffffffffffffffffffffffffffffffffff16603e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146111b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f43616c6c6572206973206e6f7420746865205661756c7400000000000000000081525060200191505060405180910390fd5b60006039541161122c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f43616e6e6f7420696e637265617365203020737570706c79000000000000000081525060200191505060405180910390fd5b81603954141561128d577f99e56f783b536ffacf422d59183ea321dd80dcd6d23daa13023e8afea38c3df1603954603a54603b5460405180848152602001838152602001828152602001935050505060405180910390a1603954905061133e565b816039819055506000196fffffffffffffffffffffffffffffffff1660395411156112ce576000196fffffffffffffffffffffffffffffffff166039819055505b6112e5603954603a54611efb90919063ffffffff16565b603b819055507f99e56f783b536ffacf422d59183ea321dd80dcd6d23daa13023e8afea38c3df1603954603a54603b5460405180848152602001838152602001828152602001935050505060405180910390a160395490505b919050565b3373ffffffffffffffffffffffffffffffffffffffff16603e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611406576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f43616c6c6572206973206e6f7420746865205661756c7400000000000000000081525060200191505060405180910390fd5b6114108282611f37565b5050565b603e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b611442612129565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146114c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806128056030913960400191505060405180910390fd5b6114ce3361215a565b565b600080603b5414156114e5576000905061153c565b611539603b54603c60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611efb90919063ffffffff16565b90505b919050565b606060378054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156115d95780601f106115ae576101008083540402835291602001916115d9565b820191906000526020600020905b8154815290600101906020018083116115bc57829003601f168201915b5050505050905090565b3373ffffffffffffffffffffffffffffffffffffffff16603e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146116a6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f43616c6c6572206973206e6f7420746865205661756c7400000000000000000081525060200191505060405180910390fd5b6116b0828261226a565b5050565b600080603d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508083106117c4576000603d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611858565b6117d78382611ca490919063ffffffff16565b603d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b8373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925603d60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040518082815260200191505060405180910390a3600191505092915050565b6000806119513384611cee565b90506119a581603c60008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611e7390919063ffffffff16565b603c60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040518082815260200191505060405180910390a3600191505092915050565b6000611a62611c73565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b611a9d611a58565b611b0f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f43616c6c6572206973206e6f742074686520476f7665726e6f7200000000000081525060200191505060405180910390fd5b611b18816123bc565b8073ffffffffffffffffffffffffffffffffffffffff16611b37611c73565b73ffffffffffffffffffffffffffffffffffffffff167fa39cc5eb22d0f34d8beaefee8a3f17cc229c1a1d1ef87a5ad47313487b1c4f0d60405160405180910390a350565b6000603d60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000603c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000803090506000813b9050600081149250505090565b611c6f828260126123eb565b5050565b6000807f7bea13895fa79d2831e0a9e28edede30099005a50d652d8957cf8a607ee6ca4a60001b9050805491505090565b6000611ce683836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250612439565b905092915050565b6000611d05603b54836124f990919063ffffffff16565b90506000603c60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811480611d5b57508160018203145b15611daa576000603c60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611e6c565b81811115611dfd57818103603c60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611e6b565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f52656d6f766520657863656564732062616c616e63650000000000000000000081525060200191505060405180910390fd5b5b5092915050565b600080828401905083811015611ef1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600080611f19670de0b6b3a76400008561251690919063ffffffff16565b9050611f2e838261259c90919063ffffffff16565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611fda576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f4d696e7420746f20746865207a65726f2061646472657373000000000000000081525060200191505060405180910390fd5b611fef81603954611e7390919063ffffffff16565b603981905550600061200c603b54836124f990919063ffffffff16565b905061206081603c60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611e7390919063ffffffff16565b603c60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506120b881603a54611e7390919063ffffffff16565b603a819055508273ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3505050565b6000807f44c4d30b2eaad5130ad70c3ba6972730566f3e6359ab83e800d905c61b1c51db60001b9050805491505090565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156121fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f4e657720476f7665726e6f72206973206164647265737328302900000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1661221c611c73565b73ffffffffffffffffffffffffffffffffffffffff167fc7c0c772add429241571afb3805861fb3cfa2af374534088b76cdb4325a87e9a60405160405180910390a3612267816125e6565b50565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561230d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f4275726e2066726f6d20746865207a65726f206164647265737300000000000081525060200191505060405180910390fd5b61232281603954611ca490919063ffffffff16565b60398190555060006123348383611cee565b905061234b81603a54611ca490919063ffffffff16565b603a81905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3505050565b60007f44c4d30b2eaad5130ad70c3ba6972730566f3e6359ab83e800d905c61b1c51db60001b90508181555050565b8260369080519060200190612401929190612710565b508160379080519060200190612418929190612710565b5080603860006101000a81548160ff021916908360ff160217905550505050565b60008383111582906124e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156124ab578082015181840152602081019050612490565b50505050905090810190601f1680156124d85780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b600061250e8383670de0b6b3a7640000612615565b905092915050565b6000808314156125295760009050612596565b600082840290508284828161253a57fe5b0414612591576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806127b66021913960400191505060405180910390fd5b809150505b92915050565b60006125de83836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f00000000000081525061264a565b905092915050565b60007f7bea13895fa79d2831e0a9e28edede30099005a50d652d8957cf8a607ee6ca4a60001b90508181555050565b60008061262b848661251690919063ffffffff16565b9050612640838261259c90919063ffffffff16565b9150509392505050565b600080831182906126f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156126bb5780820151818401526020810190506126a0565b50505050905090810190601f1680156126e85780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858161270257fe5b049050809150509392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061275157805160ff191683800117855561277f565b8280016001018555821561277f579182015b8281111561277e578251825591602001919060010190612763565b5b50905061278c9190612790565b5090565b6127b291905b808211156127ae576000816000905550600101612796565b5090565b9056fe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a65644f6e6c79207468652070656e64696e6720476f7665726e6f722063616e20636f6d706c6574652074686520636c61696da265627a7a7231582098450d08bd32258ed2fb10c6f8aec97f0e5da019a8b3bbab7de398ca8948ab8a64736f6c634300050b0032"

// DeployOusd deploys a new Ethereum contract, binding an instance of Ousd to it.
func DeployOusd(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ousd, error) {
	parsed, err := abi.JSON(strings.NewReader(OusdABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OusdBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ousd{OusdCaller: OusdCaller{contract: contract}, OusdTransactor: OusdTransactor{contract: contract}, OusdFilterer: OusdFilterer{contract: contract}}, nil
}

// Ousd is an auto generated Go binding around an Ethereum contract.
type Ousd struct {
	OusdCaller     // Read-only binding to the contract
	OusdTransactor // Write-only binding to the contract
	OusdFilterer   // Log filterer for contract events
}

// OusdCaller is an auto generated read-only Go binding around an Ethereum contract.
type OusdCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OusdTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OusdTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OusdFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OusdFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OusdSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OusdSession struct {
	Contract     *Ousd             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OusdCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OusdCallerSession struct {
	Contract *OusdCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OusdTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OusdTransactorSession struct {
	Contract     *OusdTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OusdRaw is an auto generated low-level Go binding around an Ethereum contract.
type OusdRaw struct {
	Contract *Ousd // Generic contract binding to access the raw methods on
}

// OusdCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OusdCallerRaw struct {
	Contract *OusdCaller // Generic read-only contract binding to access the raw methods on
}

// OusdTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OusdTransactorRaw struct {
	Contract *OusdTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOusd creates a new instance of Ousd, bound to a specific deployed contract.
func NewOusd(address common.Address, backend bind.ContractBackend) (*Ousd, error) {
	contract, err := bindOusd(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ousd{OusdCaller: OusdCaller{contract: contract}, OusdTransactor: OusdTransactor{contract: contract}, OusdFilterer: OusdFilterer{contract: contract}}, nil
}

// NewOusdCaller creates a new read-only instance of Ousd, bound to a specific deployed contract.
func NewOusdCaller(address common.Address, caller bind.ContractCaller) (*OusdCaller, error) {
	contract, err := bindOusd(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OusdCaller{contract: contract}, nil
}

// NewOusdTransactor creates a new write-only instance of Ousd, bound to a specific deployed contract.
func NewOusdTransactor(address common.Address, transactor bind.ContractTransactor) (*OusdTransactor, error) {
	contract, err := bindOusd(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OusdTransactor{contract: contract}, nil
}

// NewOusdFilterer creates a new log filterer instance of Ousd, bound to a specific deployed contract.
func NewOusdFilterer(address common.Address, filterer bind.ContractFilterer) (*OusdFilterer, error) {
	contract, err := bindOusd(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OusdFilterer{contract: contract}, nil
}

// bindOusd binds a generic wrapper to an already deployed contract.
func bindOusd(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OusdABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ousd *OusdRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ousd.Contract.OusdCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ousd *OusdRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ousd.Contract.OusdTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ousd *OusdRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ousd.Contract.OusdTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ousd *OusdCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ousd.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ousd *OusdTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ousd.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ousd *OusdTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ousd.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Ousd *OusdCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ousd.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Ousd *OusdSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Ousd.Contract.Allowance(&_Ousd.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Ousd *OusdCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Ousd.Contract.Allowance(&_Ousd.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Ousd *OusdCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ousd.contract.Call(opts, out, "balanceOf", _account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Ousd *OusdSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Ousd.Contract.BalanceOf(&_Ousd.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Ousd *OusdCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Ousd.Contract.BalanceOf(&_Ousd.CallOpts, _account)
}

// CreditsBalanceOf is a free data retrieval call binding the contract method 0xf9854bfc.
//
// Solidity: function creditsBalanceOf(address _account) view returns(uint256)
func (_Ousd *OusdCaller) CreditsBalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ousd.contract.Call(opts, out, "creditsBalanceOf", _account)
	return *ret0, err
}

// CreditsBalanceOf is a free data retrieval call binding the contract method 0xf9854bfc.
//
// Solidity: function creditsBalanceOf(address _account) view returns(uint256)
func (_Ousd *OusdSession) CreditsBalanceOf(_account common.Address) (*big.Int, error) {
	return _Ousd.Contract.CreditsBalanceOf(&_Ousd.CallOpts, _account)
}

// CreditsBalanceOf is a free data retrieval call binding the contract method 0xf9854bfc.
//
// Solidity: function creditsBalanceOf(address _account) view returns(uint256)
func (_Ousd *OusdCallerSession) CreditsBalanceOf(_account common.Address) (*big.Int, error) {
	return _Ousd.Contract.CreditsBalanceOf(&_Ousd.CallOpts, _account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ousd *OusdCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Ousd.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ousd *OusdSession) Decimals() (uint8, error) {
	return _Ousd.Contract.Decimals(&_Ousd.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ousd *OusdCallerSession) Decimals() (uint8, error) {
	return _Ousd.Contract.Decimals(&_Ousd.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_Ousd *OusdCaller) Governor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ousd.contract.Call(opts, out, "governor")
	return *ret0, err
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_Ousd *OusdSession) Governor() (common.Address, error) {
	return _Ousd.Contract.Governor(&_Ousd.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_Ousd *OusdCallerSession) Governor() (common.Address, error) {
	return _Ousd.Contract.Governor(&_Ousd.CallOpts)
}

// IsGovernor is a free data retrieval call binding the contract method 0xc7af3352.
//
// Solidity: function isGovernor() view returns(bool)
func (_Ousd *OusdCaller) IsGovernor(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ousd.contract.Call(opts, out, "isGovernor")
	return *ret0, err
}

// IsGovernor is a free data retrieval call binding the contract method 0xc7af3352.
//
// Solidity: function isGovernor() view returns(bool)
func (_Ousd *OusdSession) IsGovernor() (bool, error) {
	return _Ousd.Contract.IsGovernor(&_Ousd.CallOpts)
}

// IsGovernor is a free data retrieval call binding the contract method 0xc7af3352.
//
// Solidity: function isGovernor() view returns(bool)
func (_Ousd *OusdCallerSession) IsGovernor() (bool, error) {
	return _Ousd.Contract.IsGovernor(&_Ousd.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ousd *OusdCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ousd.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ousd *OusdSession) Name() (string, error) {
	return _Ousd.Contract.Name(&_Ousd.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ousd *OusdCallerSession) Name() (string, error) {
	return _Ousd.Contract.Name(&_Ousd.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ousd *OusdCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ousd.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ousd *OusdSession) Symbol() (string, error) {
	return _Ousd.Contract.Symbol(&_Ousd.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ousd *OusdCallerSession) Symbol() (string, error) {
	return _Ousd.Contract.Symbol(&_Ousd.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ousd *OusdCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ousd.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ousd *OusdSession) TotalSupply() (*big.Int, error) {
	return _Ousd.Contract.TotalSupply(&_Ousd.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ousd *OusdCallerSession) TotalSupply() (*big.Int, error) {
	return _Ousd.Contract.TotalSupply(&_Ousd.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_Ousd *OusdCaller) VaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ousd.contract.Call(opts, out, "vaultAddress")
	return *ret0, err
}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_Ousd *OusdSession) VaultAddress() (common.Address, error) {
	return _Ousd.Contract.VaultAddress(&_Ousd.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_Ousd *OusdCallerSession) VaultAddress() (common.Address, error) {
	return _Ousd.Contract.VaultAddress(&_Ousd.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Ousd *OusdTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ousd.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Ousd *OusdSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ousd.Contract.Approve(&_Ousd.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Ousd *OusdTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ousd.Contract.Approve(&_Ousd.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Ousd *OusdTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ousd.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Ousd *OusdSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ousd.Contract.Burn(&_Ousd.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Ousd *OusdTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ousd.Contract.Burn(&_Ousd.TransactOpts, account, amount)
}

// ChangeSupply is a paid mutator transaction binding the contract method 0x39a7919f.
//
// Solidity: function changeSupply(uint256 _newTotalSupply) returns(uint256)
func (_Ousd *OusdTransactor) ChangeSupply(opts *bind.TransactOpts, _newTotalSupply *big.Int) (*types.Transaction, error) {
	return _Ousd.contract.Transact(opts, "changeSupply", _newTotalSupply)
}

// ChangeSupply is a paid mutator transaction binding the contract method 0x39a7919f.
//
// Solidity: function changeSupply(uint256 _newTotalSupply) returns(uint256)
func (_Ousd *OusdSession) ChangeSupply(_newTotalSupply *big.Int) (*types.Transaction, error) {
	return _Ousd.Contract.ChangeSupply(&_Ousd.TransactOpts, _newTotalSupply)
}

// ChangeSupply is a paid mutator transaction binding the contract method 0x39a7919f.
//
// Solidity: function changeSupply(uint256 _newTotalSupply) returns(uint256)
func (_Ousd *OusdTransactorSession) ChangeSupply(_newTotalSupply *big.Int) (*types.Transaction, error) {
	return _Ousd.Contract.ChangeSupply(&_Ousd.TransactOpts, _newTotalSupply)
}

// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
//
// Solidity: function claimGovernance() returns()
func (_Ousd *OusdTransactor) ClaimGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ousd.contract.Transact(opts, "claimGovernance")
}

// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
//
// Solidity: function claimGovernance() returns()
func (_Ousd *OusdSession) ClaimGovernance() (*types.Transaction, error) {
	return _Ousd.Contract.ClaimGovernance(&_Ousd.TransactOpts)
}

// ClaimGovernance is a paid mutator transaction binding the contract method 0x5d36b190.
//
// Solidity: function claimGovernance() returns()
func (_Ousd *OusdTransactorSession) ClaimGovernance() (*types.Transaction, error) {
	return _Ousd.Contract.ClaimGovernance(&_Ousd.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_Ousd *OusdTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Ousd.contract.Transact(opts, "decreaseAllowance", _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_Ousd *OusdSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Ousd.Contract.DecreaseAllowance(&_Ousd.TransactOpts, _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_Ousd *OusdTransactorSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Ousd.Contract.DecreaseAllowance(&_Ousd.TransactOpts, _spender, _subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_Ousd *OusdTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Ousd.contract.Transact(opts, "increaseAllowance", _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_Ousd *OusdSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Ousd.Contract.IncreaseAllowance(&_Ousd.TransactOpts, _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_Ousd *OusdTransactorSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Ousd.Contract.IncreaseAllowance(&_Ousd.TransactOpts, _spender, _addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _nameArg, string _symbolArg, address _vaultAddress) returns()
func (_Ousd *OusdTransactor) Initialize(opts *bind.TransactOpts, _nameArg string, _symbolArg string, _vaultAddress common.Address) (*types.Transaction, error) {
	return _Ousd.contract.Transact(opts, "initialize", _nameArg, _symbolArg, _vaultAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _nameArg, string _symbolArg, address _vaultAddress) returns()
func (_Ousd *OusdSession) Initialize(_nameArg string, _symbolArg string, _vaultAddress common.Address) (*types.Transaction, error) {
	return _Ousd.Contract.Initialize(&_Ousd.TransactOpts, _nameArg, _symbolArg, _vaultAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _nameArg, string _symbolArg, address _vaultAddress) returns()
func (_Ousd *OusdTransactorSession) Initialize(_nameArg string, _symbolArg string, _vaultAddress common.Address) (*types.Transaction, error) {
	return _Ousd.Contract.Initialize(&_Ousd.TransactOpts, _nameArg, _symbolArg, _vaultAddress)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_Ousd *OusdTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ousd.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_Ousd *OusdSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ousd.Contract.Mint(&_Ousd.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_Ousd *OusdTransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ousd.Contract.Mint(&_Ousd.TransactOpts, _account, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Ousd *OusdTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ousd.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Ousd *OusdSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ousd.Contract.Transfer(&_Ousd.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Ousd *OusdTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ousd.Contract.Transfer(&_Ousd.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Ousd *OusdTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ousd.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Ousd *OusdSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ousd.Contract.TransferFrom(&_Ousd.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Ousd *OusdTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Ousd.Contract.TransferFrom(&_Ousd.TransactOpts, _from, _to, _value)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _newGovernor) returns()
func (_Ousd *OusdTransactor) TransferGovernance(opts *bind.TransactOpts, _newGovernor common.Address) (*types.Transaction, error) {
	return _Ousd.contract.Transact(opts, "transferGovernance", _newGovernor)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _newGovernor) returns()
func (_Ousd *OusdSession) TransferGovernance(_newGovernor common.Address) (*types.Transaction, error) {
	return _Ousd.Contract.TransferGovernance(&_Ousd.TransactOpts, _newGovernor)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _newGovernor) returns()
func (_Ousd *OusdTransactorSession) TransferGovernance(_newGovernor common.Address) (*types.Transaction, error) {
	return _Ousd.Contract.TransferGovernance(&_Ousd.TransactOpts, _newGovernor)
}

// OusdApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Ousd contract.
type OusdApprovalIterator struct {
	Event *OusdApproval // Event containing the contract specifics and raw log

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
func (it *OusdApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OusdApproval)
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
		it.Event = new(OusdApproval)
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
func (it *OusdApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OusdApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OusdApproval represents a Approval event raised by the Ousd contract.
type OusdApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ousd *OusdFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*OusdApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ousd.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &OusdApprovalIterator{contract: _Ousd.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ousd *OusdFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OusdApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ousd.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OusdApproval)
				if err := _Ousd.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Ousd *OusdFilterer) ParseApproval(log types.Log) (*OusdApproval, error) {
	event := new(OusdApproval)
	if err := _Ousd.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OusdGovernorshipTransferredIterator is returned from FilterGovernorshipTransferred and is used to iterate over the raw logs and unpacked data for GovernorshipTransferred events raised by the Ousd contract.
type OusdGovernorshipTransferredIterator struct {
	Event *OusdGovernorshipTransferred // Event containing the contract specifics and raw log

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
func (it *OusdGovernorshipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OusdGovernorshipTransferred)
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
		it.Event = new(OusdGovernorshipTransferred)
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
func (it *OusdGovernorshipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OusdGovernorshipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OusdGovernorshipTransferred represents a GovernorshipTransferred event raised by the Ousd contract.
type OusdGovernorshipTransferred struct {
	PreviousGovernor common.Address
	NewGovernor      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterGovernorshipTransferred is a free log retrieval operation binding the contract event 0xc7c0c772add429241571afb3805861fb3cfa2af374534088b76cdb4325a87e9a.
//
// Solidity: event GovernorshipTransferred(address indexed previousGovernor, address indexed newGovernor)
func (_Ousd *OusdFilterer) FilterGovernorshipTransferred(opts *bind.FilterOpts, previousGovernor []common.Address, newGovernor []common.Address) (*OusdGovernorshipTransferredIterator, error) {

	var previousGovernorRule []interface{}
	for _, previousGovernorItem := range previousGovernor {
		previousGovernorRule = append(previousGovernorRule, previousGovernorItem)
	}
	var newGovernorRule []interface{}
	for _, newGovernorItem := range newGovernor {
		newGovernorRule = append(newGovernorRule, newGovernorItem)
	}

	logs, sub, err := _Ousd.contract.FilterLogs(opts, "GovernorshipTransferred", previousGovernorRule, newGovernorRule)
	if err != nil {
		return nil, err
	}
	return &OusdGovernorshipTransferredIterator{contract: _Ousd.contract, event: "GovernorshipTransferred", logs: logs, sub: sub}, nil
}

// WatchGovernorshipTransferred is a free log subscription operation binding the contract event 0xc7c0c772add429241571afb3805861fb3cfa2af374534088b76cdb4325a87e9a.
//
// Solidity: event GovernorshipTransferred(address indexed previousGovernor, address indexed newGovernor)
func (_Ousd *OusdFilterer) WatchGovernorshipTransferred(opts *bind.WatchOpts, sink chan<- *OusdGovernorshipTransferred, previousGovernor []common.Address, newGovernor []common.Address) (event.Subscription, error) {

	var previousGovernorRule []interface{}
	for _, previousGovernorItem := range previousGovernor {
		previousGovernorRule = append(previousGovernorRule, previousGovernorItem)
	}
	var newGovernorRule []interface{}
	for _, newGovernorItem := range newGovernor {
		newGovernorRule = append(newGovernorRule, newGovernorItem)
	}

	logs, sub, err := _Ousd.contract.WatchLogs(opts, "GovernorshipTransferred", previousGovernorRule, newGovernorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OusdGovernorshipTransferred)
				if err := _Ousd.contract.UnpackLog(event, "GovernorshipTransferred", log); err != nil {
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
func (_Ousd *OusdFilterer) ParseGovernorshipTransferred(log types.Log) (*OusdGovernorshipTransferred, error) {
	event := new(OusdGovernorshipTransferred)
	if err := _Ousd.contract.UnpackLog(event, "GovernorshipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OusdPendingGovernorshipTransferIterator is returned from FilterPendingGovernorshipTransfer and is used to iterate over the raw logs and unpacked data for PendingGovernorshipTransfer events raised by the Ousd contract.
type OusdPendingGovernorshipTransferIterator struct {
	Event *OusdPendingGovernorshipTransfer // Event containing the contract specifics and raw log

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
func (it *OusdPendingGovernorshipTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OusdPendingGovernorshipTransfer)
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
		it.Event = new(OusdPendingGovernorshipTransfer)
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
func (it *OusdPendingGovernorshipTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OusdPendingGovernorshipTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OusdPendingGovernorshipTransfer represents a PendingGovernorshipTransfer event raised by the Ousd contract.
type OusdPendingGovernorshipTransfer struct {
	PreviousGovernor common.Address
	NewGovernor      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPendingGovernorshipTransfer is a free log retrieval operation binding the contract event 0xa39cc5eb22d0f34d8beaefee8a3f17cc229c1a1d1ef87a5ad47313487b1c4f0d.
//
// Solidity: event PendingGovernorshipTransfer(address indexed previousGovernor, address indexed newGovernor)
func (_Ousd *OusdFilterer) FilterPendingGovernorshipTransfer(opts *bind.FilterOpts, previousGovernor []common.Address, newGovernor []common.Address) (*OusdPendingGovernorshipTransferIterator, error) {

	var previousGovernorRule []interface{}
	for _, previousGovernorItem := range previousGovernor {
		previousGovernorRule = append(previousGovernorRule, previousGovernorItem)
	}
	var newGovernorRule []interface{}
	for _, newGovernorItem := range newGovernor {
		newGovernorRule = append(newGovernorRule, newGovernorItem)
	}

	logs, sub, err := _Ousd.contract.FilterLogs(opts, "PendingGovernorshipTransfer", previousGovernorRule, newGovernorRule)
	if err != nil {
		return nil, err
	}
	return &OusdPendingGovernorshipTransferIterator{contract: _Ousd.contract, event: "PendingGovernorshipTransfer", logs: logs, sub: sub}, nil
}

// WatchPendingGovernorshipTransfer is a free log subscription operation binding the contract event 0xa39cc5eb22d0f34d8beaefee8a3f17cc229c1a1d1ef87a5ad47313487b1c4f0d.
//
// Solidity: event PendingGovernorshipTransfer(address indexed previousGovernor, address indexed newGovernor)
func (_Ousd *OusdFilterer) WatchPendingGovernorshipTransfer(opts *bind.WatchOpts, sink chan<- *OusdPendingGovernorshipTransfer, previousGovernor []common.Address, newGovernor []common.Address) (event.Subscription, error) {

	var previousGovernorRule []interface{}
	for _, previousGovernorItem := range previousGovernor {
		previousGovernorRule = append(previousGovernorRule, previousGovernorItem)
	}
	var newGovernorRule []interface{}
	for _, newGovernorItem := range newGovernor {
		newGovernorRule = append(newGovernorRule, newGovernorItem)
	}

	logs, sub, err := _Ousd.contract.WatchLogs(opts, "PendingGovernorshipTransfer", previousGovernorRule, newGovernorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OusdPendingGovernorshipTransfer)
				if err := _Ousd.contract.UnpackLog(event, "PendingGovernorshipTransfer", log); err != nil {
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
func (_Ousd *OusdFilterer) ParsePendingGovernorshipTransfer(log types.Log) (*OusdPendingGovernorshipTransfer, error) {
	event := new(OusdPendingGovernorshipTransfer)
	if err := _Ousd.contract.UnpackLog(event, "PendingGovernorshipTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OusdTotalSupplyUpdatedIterator is returned from FilterTotalSupplyUpdated and is used to iterate over the raw logs and unpacked data for TotalSupplyUpdated events raised by the Ousd contract.
type OusdTotalSupplyUpdatedIterator struct {
	Event *OusdTotalSupplyUpdated // Event containing the contract specifics and raw log

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
func (it *OusdTotalSupplyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OusdTotalSupplyUpdated)
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
		it.Event = new(OusdTotalSupplyUpdated)
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
func (it *OusdTotalSupplyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OusdTotalSupplyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OusdTotalSupplyUpdated represents a TotalSupplyUpdated event raised by the Ousd contract.
type OusdTotalSupplyUpdated struct {
	TotalSupply     *big.Int
	TotalCredits    *big.Int
	CreditsPerToken *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTotalSupplyUpdated is a free log retrieval operation binding the contract event 0x99e56f783b536ffacf422d59183ea321dd80dcd6d23daa13023e8afea38c3df1.
//
// Solidity: event TotalSupplyUpdated(uint256 totalSupply, uint256 totalCredits, uint256 creditsPerToken)
func (_Ousd *OusdFilterer) FilterTotalSupplyUpdated(opts *bind.FilterOpts) (*OusdTotalSupplyUpdatedIterator, error) {

	logs, sub, err := _Ousd.contract.FilterLogs(opts, "TotalSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return &OusdTotalSupplyUpdatedIterator{contract: _Ousd.contract, event: "TotalSupplyUpdated", logs: logs, sub: sub}, nil
}

// WatchTotalSupplyUpdated is a free log subscription operation binding the contract event 0x99e56f783b536ffacf422d59183ea321dd80dcd6d23daa13023e8afea38c3df1.
//
// Solidity: event TotalSupplyUpdated(uint256 totalSupply, uint256 totalCredits, uint256 creditsPerToken)
func (_Ousd *OusdFilterer) WatchTotalSupplyUpdated(opts *bind.WatchOpts, sink chan<- *OusdTotalSupplyUpdated) (event.Subscription, error) {

	logs, sub, err := _Ousd.contract.WatchLogs(opts, "TotalSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OusdTotalSupplyUpdated)
				if err := _Ousd.contract.UnpackLog(event, "TotalSupplyUpdated", log); err != nil {
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

// ParseTotalSupplyUpdated is a log parse operation binding the contract event 0x99e56f783b536ffacf422d59183ea321dd80dcd6d23daa13023e8afea38c3df1.
//
// Solidity: event TotalSupplyUpdated(uint256 totalSupply, uint256 totalCredits, uint256 creditsPerToken)
func (_Ousd *OusdFilterer) ParseTotalSupplyUpdated(log types.Log) (*OusdTotalSupplyUpdated, error) {
	event := new(OusdTotalSupplyUpdated)
	if err := _Ousd.contract.UnpackLog(event, "TotalSupplyUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OusdTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ousd contract.
type OusdTransferIterator struct {
	Event *OusdTransfer // Event containing the contract specifics and raw log

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
func (it *OusdTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OusdTransfer)
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
		it.Event = new(OusdTransfer)
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
func (it *OusdTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OusdTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OusdTransfer represents a Transfer event raised by the Ousd contract.
type OusdTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ousd *OusdFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OusdTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ousd.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OusdTransferIterator{contract: _Ousd.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ousd *OusdFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OusdTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ousd.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OusdTransfer)
				if err := _Ousd.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Ousd *OusdFilterer) ParseTransfer(log types.Log) (*OusdTransfer, error) {
	event := new(OusdTransfer)
	if err := _Ousd.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
