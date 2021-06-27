// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cerc20

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

// Cerc20ABI is the input ABI used to generate the binding from.
const Cerc20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cashPrior\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidateBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"oldComptroller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"NewComptroller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"oldInterestRateModel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewMarketInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReserveFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewReserveFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"RepayBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"benefactor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesReduced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"}],\"name\":\"_reduceReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"_setComptroller\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setReserveFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accrualBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isCToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying_\",\"type\":\"address\"},{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractCTokenInterface\",\"name\":\"cTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"}],\"name\":\"_addReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Cerc20Bin is the compiled bytecode used for deploying new contracts.
var Cerc20Bin = "0x608060405234801561001057600080fd5b50615227806100206000396000f3fe608060405234801561001057600080fd5b50600436106103205760003560e01c80638f840ddd116101a7578063c37f68e2116100ee578063f3fdb15a11610097578063f8f9da2811610071578063f8f9da2814610abd578063fca7820b14610ac5578063fe9c44ae14610ae257610320565b8063f3fdb15a14610a77578063f5e3c46214610a7f578063f851a44014610ab557610320565b8063dd62ed3e116100c8578063dd62ed3e14610a1b578063e9c714f214610a49578063f2b3abbd14610a5157610320565b8063c37f68e214610995578063c5ebeaec146109e1578063db006a75146109fe57610320565b8063a9059cbb11610150578063b2a02ff11161012a578063b2a02ff114610931578063b71d1a0c14610967578063bd6d894d1461098d57610320565b8063a9059cbb146108f5578063aa5af0fd14610921578063ae9d70b01461092957610320565b806399d8c1b41161018157806399d8c1b41461077e578063a0712d68146108d0578063a6afed95146108ed57610320565b80638f840ddd1461074857806395d89b411461075057806395dd91931461075857610320565b80633af9e6691161026b578063601a0bf11161021457806370a08231116101ee57806370a08231146106fd57806373acee9814610723578063852a12e31461072b57610320565b8063601a0bf1146106d05780636c540baf146106ed5780636f307dc3146106f557610320565b80634576b5db116102455780634576b5db1461069a57806347bd3718146106c05780635fe3b567146106c857610320565b80633af9e6691461064f5780633b1d21a2146106755780633e9410101461067d57610320565b8063182df0f5116102cd5780632608f818116102a75780632608f818146105e1578063267822471461060d578063313ce5671461063157610320565b8063182df0f5146104475780631a31d4651461044f57806323b872dd146105ab57610320565b8063173b9904116102fe578063173b99041461041157806317bfdfbc1461041957806318160ddd1461043f57610320565b806306fdde0314610325578063095ea7b3146103a25780630e752702146103e2575b600080fd5b61032d610aea565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561036757818101518382015260200161034f565b50505050905090810190601f1680156103945780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103ce600480360360408110156103b857600080fd5b506001600160a01b038135169060200135610b77565b604080519115158252519081900360200190f35b6103ff600480360360208110156103f857600080fd5b5035610be4565b60408051918252519081900360200190f35b6103ff610bfa565b6103ff6004803603602081101561042f57600080fd5b50356001600160a01b0316610c00565b6103ff610cc7565b6103ff610ccd565b6105a9600480360360e081101561046557600080fd5b6001600160a01b03823581169260208101358216926040820135909216916060820135919081019060a0810160808201356401000000008111156104a857600080fd5b8201836020820111156104ba57600080fd5b803590602001918460018302840111640100000000831117156104dc57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561052f57600080fd5b82018360208201111561054157600080fd5b8035906020019184600183028401116401000000008311171561056357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903560ff169150610d309050565b005b6103ce600480360360608110156105c157600080fd5b506001600160a01b03813581169160208101359091169060400135610de8565b6103ff600480360360408110156105f757600080fd5b506001600160a01b038135169060200135610e5a565b610615610e70565b604080516001600160a01b039092168252519081900360200190f35b610639610e7f565b6040805160ff9092168252519081900360200190f35b6103ff6004803603602081101561066557600080fd5b50356001600160a01b0316610e88565b6103ff610f3e565b6103ff6004803603602081101561069357600080fd5b5035610f4d565b6103ff600480360360208110156106b057600080fd5b50356001600160a01b0316610f58565b6103ff6110c6565b6106156110cc565b6103ff600480360360208110156106e657600080fd5b50356110db565b6103ff611176565b61061561117c565b6103ff6004803603602081101561071357600080fd5b50356001600160a01b031661118b565b6103ff6111a6565b6103ff6004803603602081101561074157600080fd5b5035611263565b6103ff61126e565b61032d611274565b6103ff6004803603602081101561076e57600080fd5b50356001600160a01b03166112cc565b6105a9600480360360c081101561079457600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156107cf57600080fd5b8201836020820111156107e157600080fd5b8035906020019184600183028401116401000000008311171561080357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561085657600080fd5b82018360208201111561086857600080fd5b8035906020019184600183028401116401000000008311171561088a57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903560ff1691506113299050565b6103ff600480360360208110156108e657600080fd5b5035611510565b6103ff61151c565b6103ce6004803603604081101561090b57600080fd5b506001600160a01b03813516906020013561188d565b6103ff6118fe565b6103ff611904565b6103ff6004803603606081101561094757600080fd5b506001600160a01b038135811691602081013590911690604001356119a3565b6103ff6004803603602081101561097d57600080fd5b50356001600160a01b0316611a14565b6103ff611aa0565b6109bb600480360360208110156109ab57600080fd5b50356001600160a01b0316611b63565b604080519485526020850193909352838301919091526060830152519081900360800190f35b6103ff600480360360208110156109f757600080fd5b5035611bf8565b6103ff60048036036020811015610a1457600080fd5b5035611c03565b6103ff60048036036040811015610a3157600080fd5b506001600160a01b0381358116916020013516611c0e565b6103ff611c39565b6103ff60048036036020811015610a6757600080fd5b50356001600160a01b0316611d53565b610615611d8d565b6103ff60048036036060811015610a9557600080fd5b506001600160a01b03813581169160208101359160409091013516611d9c565b610615611db4565b6103ff611dc8565b6103ff60048036036020811015610adb57600080fd5b5035611e2c565b6103ce611eaa565b60018054604080516020600284861615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610b6f5780601f10610b4457610100808354040283529160200191610b6f565b820191906000526020600020905b815481529060010190602001808311610b5257829003601f168201915b505050505081565b336000818152600f602090815260408083206001600160a01b03871680855290835281842086905581518681529151939493909284927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a360019150505b92915050565b600080610bf083611eaf565b509150505b919050565b60085481565b6000805460ff16610c45576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155610c5761151c565b14610ca9576040805162461bcd60e51b815260206004820152601660248201527f61636372756520696e746572657374206661696c656400000000000000000000604482015290519081900360640190fd5b610cb2826112cc565b90505b6000805460ff19166001179055919050565b600d5481565b6000806000610cda611f58565b90925090506000826003811115610ced57fe5b14610d295760405162461bcd60e51b815260040180806020018281038252603581526020018061513e6035913960400191505060405180910390fd5b9150505b90565b610d3e868686868686611329565b601180546001600160a01b0319166001600160a01b038981169190911791829055604080517f18160ddd000000000000000000000000000000000000000000000000000000008152905192909116916318160ddd91600480820192602092909190829003018186803b158015610db357600080fd5b505afa158015610dc7573d6000803e3d6000fd5b505050506040513d6020811015610ddd57600080fd5b505050505050505050565b6000805460ff16610e2d576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155610e4333868686612007565b1490506000805460ff191660011790559392505050565b600080610e678484612359565b50949350505050565b6004546001600160a01b031681565b60035460ff1681565b6000610e92614e4c565b6040518060200160405280610ea5611aa0565b90526001600160a01b0384166000908152600e6020526040812054919250908190610ed1908490612404565b90925090506000826003811115610ee457fe5b14610f36576040805162461bcd60e51b815260206004820152601f60248201527f62616c616e636520636f756c64206e6f742062652063616c63756c6174656400604482015290519081900360640190fd5b949350505050565b6000610f48612458565b905090565b6000610bde826124d8565b60035460009061010090046001600160a01b03163314610f8557610f7e6001603f61256c565b9050610bf5565b600554604080517e7e3dd200000000000000000000000000000000000000000000000000000000815290516001600160a01b0392831692851691627e3dd2916004808301926020929190829003018186803b158015610fe357600080fd5b505afa158015610ff7573d6000803e3d6000fd5b505050506040513d602081101561100d57600080fd5b5051611060576040805162461bcd60e51b815260206004820152601c60248201527f6d61726b6572206d6574686f642072657475726e65642066616c736500000000604482015290519081900360640190fd5b600580546001600160a01b0319166001600160a01b03858116918217909255604080519284168352602083019190915280517f7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d9281900390910190a160005b9392505050565b600b5481565b6005546001600160a01b031681565b6000805460ff16611120576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff1916815561113261151c565b905080156111585761115081601081111561114957fe5b603061256c565b915050610cb5565b611161836125d2565b9150506000805460ff19166001179055919050565b60095481565b6011546001600160a01b031681565b6001600160a01b03166000908152600e602052604090205490565b6000805460ff166111eb576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff191681556111fd61151c565b1461124f576040805162461bcd60e51b815260206004820152601660248201527f61636372756520696e746572657374206661696c656400000000000000000000604482015290519081900360640190fd5b50600b546000805460ff1916600117905590565b6000610bde82612705565b600c5481565b6002805460408051602060018416156101000260001901909316849004601f81018490048402820184019092528181529291830182828015610b6f5780601f10610b4457610100808354040283529160200191610b6f565b60008060006112da84612786565b909250905060008260038111156112ed57fe5b146110bf5760405162461bcd60e51b81526004018080602001828103825260378152602001806150696037913960400191505060405180910390fd5b60035461010090046001600160a01b031633146113775760405162461bcd60e51b8152600401808060200182810382526024815260200180614fa56024913960400191505060405180910390fd5b6009541580156113875750600a54155b6113c25760405162461bcd60e51b8152600401808060200182810382526023815260200180614fc96023913960400191505060405180910390fd5b6007849055836114035760405162461bcd60e51b8152600401808060200182810382526030815260200180614fec6030913960400191505060405180910390fd5b600061140e87610f58565b90508015611463576040805162461bcd60e51b815260206004820152601a60248201527f73657474696e6720636f6d7074726f6c6c6572206661696c6564000000000000604482015290519081900360640190fd5b61146b61283a565b600955670de0b6b3a7640000600a556114838661283e565b905080156114c25760405162461bcd60e51b815260040180806020018281038252602281526020018061501c6022913960400191505060405180910390fd5b83516114d5906001906020870190614e5f565b5082516114e9906002906020860190614e5f565b50506003805460ff90921660ff199283161790556000805490911660011790555050505050565b600080610bf0836129b3565b60008061152761283a565b6009549091508082141561154057600092505050610d2d565b600061154a612458565b600b54600c54600a54600654604080517f15f2405300000000000000000000000000000000000000000000000000000000815260048101879052602481018690526044810185905290519596509394929391926000926001600160a01b03909216916315f24053916064808301926020929190829003018186803b1580156115d157600080fd5b505afa1580156115e5573d6000803e3d6000fd5b505050506040513d60208110156115fb57600080fd5b5051905065048c2739500081111561165a576040805162461bcd60e51b815260206004820152601c60248201527f626f72726f772072617465206973206162737572646c79206869676800000000604482015290519081900360640190fd5b6000806116678989612a34565b9092509050600082600381111561167a57fe5b146116cc576040805162461bcd60e51b815260206004820152601f60248201527f636f756c64206e6f742063616c63756c61746520626c6f636b2064656c746100604482015290519081900360640190fd5b6116d4614e4c565b6000806000806116f260405180602001604052808a81525087612a57565b9097509450600087600381111561170557fe5b14611737576117226009600689600381111561171d57fe5b612abf565b9e505050505050505050505050505050610d2d565b611741858c612404565b9097509350600087600381111561175457fe5b1461176c576117226009600189600381111561171d57fe5b611776848c612b25565b9097509250600087600381111561178957fe5b146117a1576117226009600489600381111561171d57fe5b6117bc6040518060200160405280600854815250858c612b4b565b909750915060008760038111156117cf57fe5b146117e7576117226009600589600381111561171d57fe5b6117f2858a8b612b4b565b9097509050600087600381111561180557fe5b1461181d576117226009600389600381111561171d57fe5b60098e9055600a819055600b839055600c829055604080518d8152602081018690528082018390526060810185905290517f4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc049181900360800190a160009e50505050505050505050505050505090565b6000805460ff166118d2576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff191681556118e833338686612007565b1490506000805460ff1916600117905592915050565b600a5481565b6006546000906001600160a01b031663b8168816611920612458565b600b54600c546008546040518563ffffffff1660e01b81526004018085815260200184815260200183815260200182815260200194505050505060206040518083038186803b15801561197257600080fd5b505afa158015611986573d6000803e3d6000fd5b505050506040513d602081101561199c57600080fd5b5051905090565b6000805460ff166119e8576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff191690556119fe33858585612ba7565b90506000805460ff191660011790559392505050565b60035460009061010090046001600160a01b03163314611a3a57610f7e6001604561256c565b600480546001600160a01b038481166001600160a01b0319831681179093556040805191909216808252602082019390935281517fca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9929181900390910190a160006110bf565b6000805460ff16611ae5576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155611af761151c565b14611b49576040805162461bcd60e51b815260206004820152601660248201527f61636372756520696e746572657374206661696c656400000000000000000000604482015290519081900360640190fd5b611b51610ccd565b90506000805460ff1916600117905590565b6001600160a01b0381166000908152600e6020526040812054819081908190818080611b8e89612786565b935090506000816003811115611ba057fe5b14611bbe5760095b975060009650869550859450611bf19350505050565b611bc6611f58565b925090506000816003811115611bd857fe5b14611be4576009611ba8565b5060009650919450925090505b9193509193565b6000610bde82612e51565b6000610bde82612ed0565b6001600160a01b039182166000908152600f6020908152604080832093909416825291909152205490565b6004546000906001600160a01b031633141580611c54575033155b15611c6c57611c656001600061256c565b9050610d2d565b60038054600480546001600160a01b038181166101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff871617968790556001600160a01b031990931690935560408051948390048216808652929095041660208401528351909391927ff9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc92908290030190a1600454604080516001600160a01b038085168252909216602083015280517fca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a99281900390910190a160009250505090565b600080611d5e61151c565b90508015611d8457611d7c816010811115611d7557fe5b604061256c565b915050610bf5565b6110bf8361283e565b6006546001600160a01b031681565b600080611daa858585612f4a565b5095945050505050565b60035461010090046001600160a01b031681565b6006546000906001600160a01b03166315f24053611de4612458565b600b54600c546040518463ffffffff1660e01b815260040180848152602001838152602001828152602001935050505060206040518083038186803b15801561197257600080fd5b6000805460ff16611e71576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155611e8361151c565b90508015611ea157611150816010811115611e9a57fe5b604661256c565b6111618361307c565b600181565b60008054819060ff16611ef6576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155611f0861151c565b90508015611f3357611f26816010811115611f1f57fe5b603661256c565b925060009150611f449050565b611f3e333386613124565b92509250505b6000805460ff191660011790559092909150565b600d54600090819080611f7357505060075460009150612003565b6000611f7d612458565b90506000611f89614e4c565b6000611f9a84600b54600c5461353b565b935090506000816003811115611fac57fe5b14611fc1579550600094506120039350505050565b611fcb8386613579565b925090506000816003811115611fdd57fe5b14611ff2579550600094506120039350505050565b505160009550935061200392505050565b9091565b600554604080517fbdcdc2580000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b03868116602483015285811660448301526064820185905291516000938493169163bdcdc25891608480830192602092919082900301818787803b15801561208557600080fd5b505af1158015612099573d6000803e3d6000fd5b505050506040513d60208110156120af57600080fd5b5051905080156120ce576120c66003604a83612abf565b915050610f36565b836001600160a01b0316856001600160a01b031614156120f4576120c66002604b61256c565b60006001600160a01b038781169087161415612113575060001961213b565b506001600160a01b038086166000908152600f60209081526040808320938a16835292905220545b60008060008061214b8589612a34565b9094509250600084600381111561215e57fe5b1461217c5761216f6009604b61256c565b9650505050505050610f36565b6001600160a01b038a166000908152600e602052604090205461219f9089612a34565b909450915060008460038111156121b257fe5b146121c35761216f6009604c61256c565b6001600160a01b0389166000908152600e60205260409020546121e69089612b25565b909450905060008460038111156121f957fe5b1461220a5761216f6009604d61256c565b6001600160a01b03808b166000908152600e6020526040808220859055918b168152208190556000198514612262576001600160a01b03808b166000908152600f60209081526040808320938f168352929052208390555b886001600160a01b03168a6001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8a6040518082815260200191505060405180910390a3600554604080517f6a56947e0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038d811660248301528c81166044830152606482018c905291519190921691636a56947e91608480830192600092919082900301818387803b15801561232957600080fd5b505af115801561233d573d6000803e3d6000fd5b506000925061234a915050565b9b9a5050505050505050505050565b60008054819060ff166123a0576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff191681556123b261151c565b905080156123dd576123d08160108111156123c957fe5b603561256c565b9250600091506123ee9050565b6123e8338686613124565b92509250505b6000805460ff1916600117905590939092509050565b6000806000612411614e4c565b61241b8686612a57565b9092509050600082600381111561242e57fe5b1461243f5750915060009050612451565b600061244a82613629565b9350935050505b9250929050565b601154604080516370a0823160e01b815230600482015290516000926001600160a01b03169182916370a0823191602480820192602092909190829003018186803b1580156124a657600080fd5b505afa1580156124ba573d6000803e3d6000fd5b505050506040513d60208110156124d057600080fd5b505191505090565b6000805460ff1661251d576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff1916815561252f61151c565b9050801561254d5761115081601081111561254657fe5b604e61256c565b61255683613638565b509150506000805460ff19166001179055919050565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa083601081111561259b57fe5b8360508111156125a757fe5b604080519283526020830191909152600082820152519081900360600190a18260108111156110bf57fe5b600354600090819061010090046001600160a01b031633146125fa57611d7c6001603161256c565b61260261283a565b6009541461261657611d7c600a603361256c565b8261261f612458565b101561263157611d7c600e603261256c565b600c5483111561264757611d7c6002603461256c565b50600c548281039081111561268d5760405162461bcd60e51b81526004018080602001828103825260248152602001806151cf6024913960400191505060405180910390fd5b600c8190556003546126ad9061010090046001600160a01b031684613720565b600354604080516101009092046001600160a01b0316825260208201859052818101839052517f3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e916060908290030190a160006110bf565b6000805460ff1661274a576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff1916815561275c61151c565b9050801561277a5761115081601081111561277357fe5b602761256c565b61116133600085613830565b6001600160a01b0381166000908152601060205260408120805482918291829182916127bd57506000945084935061283592505050565b6127cd8160000154600a54613d3b565b909450925060008460038111156127e057fe5b146127f5575091935060009250612835915050565b612803838260010154613d7a565b9094509150600084600381111561281657fe5b1461282b575091935060009250612835915050565b5060009450925050505b915091565b4390565b600354600090819061010090046001600160a01b0316331461286657611d7c6001604261256c565b61286e61283a565b6009541461288257611d7c600a604161256c565b600660009054906101000a90046001600160a01b03169050826001600160a01b0316632191f92a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156128d357600080fd5b505afa1580156128e7573d6000803e3d6000fd5b505050506040513d60208110156128fd57600080fd5b5051612950576040805162461bcd60e51b815260206004820152601c60248201527f6d61726b6572206d6574686f642072657475726e65642066616c736500000000604482015290519081900360640190fd5b600680546001600160a01b0319166001600160a01b03858116918217909255604080519284168352602083019190915280517fedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f9269281900390910190a160006110bf565b60008054819060ff166129fa576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155612a0c61151c565b90508015612a2a57611f26816010811115612a2357fe5b601e61256c565b611f3e3385613da5565b600080838311612a4b575060009050818303612451565b50600390506000612451565b6000612a61614e4c565b600080612a72866000015186613d3b565b90925090506000826003811115612a8557fe5b14612aa457506040805160208101909152600081529092509050612451565b60408051602081019091529081526000969095509350505050565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0846010811115612aee57fe5b846050811115612afa57fe5b604080519283526020830191909152818101859052519081900360600190a1836010811115610f3657fe5b600080838301848110612b3d57600092509050612451565b506002915060009050612451565b6000806000612b58614e4c565b612b628787612a57565b90925090506000826003811115612b7557fe5b14612b865750915060009050612b9f565b612b98612b9282613629565b86612b25565b9350935050505b935093915050565b600554604080517fd02f73510000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038781166024830152868116604483015285811660648301526084820185905291516000938493169163d02f73519160a480830192602092919082900301818787803b158015612c2d57600080fd5b505af1158015612c41573d6000803e3d6000fd5b505050506040513d6020811015612c5757600080fd5b505190508015612c6e576120c66003601b83612abf565b846001600160a01b0316846001600160a01b03161415612c94576120c66006601c61256c565b6001600160a01b0384166000908152600e602052604081205481908190612cbb9087612a34565b90935091506000836003811115612cce57fe5b14612cf157612ce66009601a85600381111561171d57fe5b945050505050610f36565b6001600160a01b0388166000908152600e6020526040902054612d149087612b25565b90935090506000836003811115612d2757fe5b14612d3f57612ce66009601985600381111561171d57fe5b6001600160a01b038088166000818152600e60209081526040808320879055938c168083529184902085905583518a8152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a3600554604080517f6d35bf910000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038c811660248301528b811660448301528a81166064830152608482018a905291519190921691636d35bf919160a480830192600092919082900301818387803b158015612e2357600080fd5b505af1158015612e37573d6000803e3d6000fd5b5060009250612e44915050565b9998505050505050505050565b6000805460ff16612e96576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155612ea861151c565b90508015612ec657611150816010811115612ebf57fe5b600861256c565b6111613384614248565b6000805460ff16612f15576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155612f2761151c565b90508015612f3e5761115081601081111561277357fe5b61116133846000613830565b60008054819060ff16612f91576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155612fa361151c565b90508015612fce57612fc1816010811115612fba57fe5b600f61256c565b9250600091506130659050565b836001600160a01b031663a6afed956040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561300957600080fd5b505af115801561301d573d6000803e3d6000fd5b505050506040513d602081101561303357600080fd5b50519050801561305357612fc181601081111561304c57fe5b601061256c565b61305f33878787614588565b92509250505b6000805460ff191660011790559094909350915050565b60035460009061010090046001600160a01b031633146130a257610f7e6001604761256c565b6130aa61283a565b600954146130be57610f7e600a604861256c565b670de0b6b3a76400008211156130da57610f7e6002604961256c565b6008805490839055604080518281526020810185905281517faaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460929181900390910190a160006110bf565b600554604080517f24008a620000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b0386811660248301528581166044830152606482018590529151600093849384939116916324008a629160848082019260209290919082900301818787803b1580156131a657600080fd5b505af11580156131ba573d6000803e3d6000fd5b505050506040513d60208110156131d057600080fd5b5051905080156131f4576131e76003603883612abf565b925060009150612b9f9050565b6131fc61283a565b60095414613210576131e7600a603961256c565b613218614edd565b6001600160a01b038616600090815260106020526040902060010154606082015261324286612786565b608083018190526020830182600381111561325957fe5b600381111561326457fe5b905250600090508160200151600381111561327b57fe5b146132a557613297600960378360200151600381111561171d57fe5b935060009250612b9f915050565b6000198514156132be57608081015160408201526132c6565b604081018590525b6132d4878260400151614b78565b60e0820181905260808201516132e991612a34565b60a083018190526020830182600381111561330057fe5b600381111561330b57fe5b905250600090508160200151600381111561332257fe5b1461335e5760405162461bcd60e51b815260040180806020018281038252603a8152602001806150a0603a913960400191505060405180910390fd5b61336e600b548260e00151612a34565b60c083018190526020830182600381111561338557fe5b600381111561339057fe5b90525060009050816020015160038111156133a757fe5b146133e35760405162461bcd60e51b81526004018080602001828103825260318152602001806150da6031913960400191505060405180910390fd5b60a080820180516001600160a01b03808a16600081815260106020908152604091829020948555600a5460019095019490945560c0870151600b81905560e088015195518251948f16855294840192909252828101949094526060820192909252608081019190915290517f1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1929181900390910190a160055460e08201516060830151604080517f1ededc910000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038c811660248301528b8116604483015260648201949094526084810192909252519190921691631ededc919160a480830192600092919082900301818387803b15801561350757600080fd5b505af115801561351b573d6000803e3d6000fd5b5060009250613528915050565b8160e00151935093505050935093915050565b60008060008061354b8787612b25565b9092509050600082600381111561355e57fe5b1461356f5750915060009050612b9f565b612b988186612a34565b6000613583614e4c565b60008061359886670de0b6b3a7640000613d3b565b909250905060008260038111156135ab57fe5b146135ca57506040805160208101909152600081529092509050612451565b6000806135d78388613d7a565b909250905060008260038111156135ea57fe5b1461360c57506040805160208101909152600081529094509250612451915050565b604080516020810190915290815260009890975095505050505050565b51670de0b6b3a7640000900490565b60008060008061364661283a565b600954146136655761365a600a604f61256c565b935091506128359050565b61366f3386614b78565b905080600c54019150600c548210156136cf576040805162461bcd60e51b815260206004820181905260248201527f61646420726573657276657320756e6578706563746564206f766572666c6f77604482015290519081900360640190fd5b600c829055604080513381526020810183905280820184905290517fa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc59181900360600190a160009350915050915091565b601154604080517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301526024820185905291519190921691829163a9059cbb9160448082019260009290919082900301818387803b15801561379157600080fd5b505af11580156137a5573d6000803e3d6000fd5b5050505060003d600081146137c157602081146137cb57600080fd5b60001991506137d7565b60206000803e60005191505b508061382a576040805162461bcd60e51b815260206004820152601960248201527f544f4b454e5f5452414e534645525f4f55545f4641494c454400000000000000604482015290519081900360640190fd5b50505050565b600082158061383d575081155b6138785760405162461bcd60e51b815260040180806020018281038252603481526020018061519b6034913960400191505060405180910390fd5b613880614f23565b613888611f58565b604083018190526020830182600381111561389f57fe5b60038111156138aa57fe5b90525060009050816020015160038111156138c157fe5b146138e5576138dd6009602b8360200151600381111561171d57fe5b9150506110bf565b831561396657606081018490526040805160208101825290820151815261390c9085612404565b608083018190526020830182600381111561392357fe5b600381111561392e57fe5b905250600090508160200151600381111561394557fe5b14613961576138dd600960298360200151600381111561171d57fe5b6139df565b6139828360405180602001604052808460400151815250614ddb565b606083018190526020830182600381111561399957fe5b60038111156139a457fe5b90525060009050816020015160038111156139bb57fe5b146139d7576138dd6009602a8360200151600381111561171d57fe5b608081018390525b6005546060820151604080517feabe7d910000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b03898116602483015260448201939093529051600093929092169163eabe7d919160648082019260209290919082900301818787803b158015613a5d57600080fd5b505af1158015613a71573d6000803e3d6000fd5b505050506040513d6020811015613a8757600080fd5b505190508015613aa757613a9e6003602883612abf565b925050506110bf565b613aaf61283a565b60095414613ac357613a9e600a602c61256c565b613ad3600d548360600151612a34565b60a0840181905260208401826003811115613aea57fe5b6003811115613af557fe5b9052506000905082602001516003811115613b0c57fe5b14613b2857613a9e6009602e8460200151600381111561171d57fe5b6001600160a01b0386166000908152600e60205260409020546060830151613b509190612a34565b60c0840181905260208401826003811115613b6757fe5b6003811115613b7257fe5b9052506000905082602001516003811115613b8957fe5b14613ba557613a9e6009602d8460200151600381111561171d57fe5b8160800151613bb2612458565b1015613bc457613a9e600e602f61256c565b613bd2868360800151613720565b60a0820151600d5560c08201516001600160a01b0387166000818152600e60209081526040918290209390935560608501518151908152905130937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef928290030190a36080820151606080840151604080516001600160a01b038b168152602081019490945283810191909152517fe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a9299281900390910190a160055460808301516060840151604080517f51dff9890000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038b81166024830152604482019490945260648101929092525191909216916351dff98991608480830192600092919082900301818387803b158015613d1057600080fd5b505af1158015613d24573d6000803e3d6000fd5b5060009250613d31915050565b9695505050505050565b60008083613d4e57506000905080612451565b83830283858281613d5b57fe5b0414613d6f57506002915060009050612451565b600092509050612451565b60008082613d8e5750600190506000612451565b6000838581613d9957fe5b04915091509250929050565b600554604080517f4ef4c3e10000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b03858116602483015260448201859052915160009384938493911691634ef4c3e19160648082019260209290919082900301818787803b158015613e1f57600080fd5b505af1158015613e33573d6000803e3d6000fd5b505050506040513d6020811015613e4957600080fd5b505190508015613e6d57613e606003601f83612abf565b9250600091506124519050565b613e7561283a565b60095414613e8957613e60600a602261256c565b613e91614f23565b613e99611f58565b6040830181905260208301826003811115613eb057fe5b6003811115613ebb57fe5b9052506000905081602001516003811115613ed257fe5b14613efc57613eee600960218360200151600381111561171d57fe5b935060009250612451915050565b613f068686614b78565b60c0820181905260408051602081018252908301518152613f279190614ddb565b6060830181905260208301826003811115613f3e57fe5b6003811115613f4957fe5b9052506000905081602001516003811115613f6057fe5b14613fb2576040805162461bcd60e51b815260206004820181905260248201527f4d494e545f45584348414e47455f43414c43554c4154494f4e5f4641494c4544604482015290519081900360640190fd5b613fc2600d548260600151612b25565b6080830181905260208301826003811115613fd957fe5b6003811115613fe457fe5b9052506000905081602001516003811115613ffb57fe5b146140375760405162461bcd60e51b81526004018080602001828103825260288152602001806151736028913960400191505060405180910390fd5b6001600160a01b0386166000908152600e6020526040902054606082015161405f9190612b25565b60a083018190526020830182600381111561407657fe5b600381111561408157fe5b905250600090508160200151600381111561409857fe5b146140d45760405162461bcd60e51b815260040180806020018281038252602b81526020018061503e602b913960400191505060405180910390fd5b6080810151600d5560a08101516001600160a01b0387166000818152600e60209081526040918290209390935560c084015160608086015183519485529484019190915282820193909352517f4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f929181900390910190a1606081015160408051918252516001600160a01b0388169130917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a360055460c08201516060830151604080517f41c728b90000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038b81166024830152604482019490945260648101929092525191909216916341c728b991608480830192600092919082900301818387803b15801561421557600080fd5b505af1158015614229573d6000803e3d6000fd5b5060009250614236915050565b8160c001519350935050509250929050565b600554604080517fda3d454c0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b0385811660248301526044820185905291516000938493169163da3d454c91606480830192602092919082900301818787803b1580156142be57600080fd5b505af11580156142d2573d6000803e3d6000fd5b505050506040513d60208110156142e857600080fd5b505190508015614307576142ff6003600e83612abf565b915050610bde565b61430f61283a565b60095414614322576142ff600a8061256c565b8261432b612458565b101561433d576142ff600e600961256c565b614345614f61565b61434e85612786565b602083018190528282600381111561436257fe5b600381111561436d57fe5b905250600090508151600381111561438157fe5b146143a65761439d600960078360000151600381111561171d57fe5b92505050610bde565b6143b4816020015185612b25565b60408301819052828260038111156143c857fe5b60038111156143d357fe5b90525060009050815160038111156143e757fe5b146144035761439d6009600c8360000151600381111561171d57fe5b61440f600b5485612b25565b606083018190528282600381111561442357fe5b600381111561442e57fe5b905250600090508151600381111561444257fe5b1461445e5761439d6009600b8360000151600381111561171d57fe5b6144688585613720565b604080820180516001600160a01b03881660008181526010602090815290859020928355600a54600190930192909255606080860151600b81905593518551928352928201899052818501929092529081019190915290517f13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab809181900360800190a1600554604080517f5c7786050000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b0388811660248301526044820188905291519190921691635c77860591606480830192600092919082900301818387803b15801561455e57600080fd5b505af1158015614572573d6000803e3d6000fd5b506000925061457f915050565b95945050505050565b600554604080517f5fc7e71e0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b0384811660248301528781166044830152868116606483015260848201869052915160009384938493911691635fc7e71e9160a48082019260209290919082900301818787803b15801561461257600080fd5b505af1158015614626573d6000803e3d6000fd5b505050506040513d602081101561463c57600080fd5b505190508015614660576146536003601283612abf565b925060009150614b6f9050565b61466861283a565b6009541461467c57614653600a601661256c565b61468461283a565b846001600160a01b0316636c540baf6040518163ffffffff1660e01b815260040160206040518083038186803b1580156146bd57600080fd5b505afa1580156146d1573d6000803e3d6000fd5b505050506040513d60208110156146e757600080fd5b5051146146fa57614653600a601161256c565b866001600160a01b0316866001600160a01b03161415614720576146536006601761256c565b84614731576146536007601561256c565b600019851415614747576146536007601461256c565b600080614755898989613124565b909250905081156147855761477682601081111561476f57fe5b601861256c565b945060009350614b6f92505050565b600554604080517fc488847b0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038981166024830152604482018590528251600094859492169263c488847b926064808301939192829003018186803b1580156147f857600080fd5b505afa15801561480c573d6000803e3d6000fd5b505050506040513d604081101561482257600080fd5b5080516020909101519092509050811561486d5760405162461bcd60e51b815260040180806020018281038252603381526020018061510b6033913960400191505060405180910390fd5b80886001600160a01b03166370a082318c6040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156148c457600080fd5b505afa1580156148d8573d6000803e3d6000fd5b505050506040513d60208110156148ee57600080fd5b50511015614943576040805162461bcd60e51b815260206004820152601860248201527f4c49515549444154455f5345495a455f544f4f5f4d5543480000000000000000604482015290519081900360640190fd5b60006001600160a01b03891630141561496957614962308d8d85612ba7565b9050614a0c565b604080517fb2a02ff10000000000000000000000000000000000000000000000000000000081526001600160a01b038e811660048301528d81166024830152604482018590529151918b169163b2a02ff1916064808201926020929091908290030181600087803b1580156149dd57600080fd5b505af11580156149f1573d6000803e3d6000fd5b505050506040513d6020811015614a0757600080fd5b505190505b8015614a5f576040805162461bcd60e51b815260206004820152601460248201527f746f6b656e207365697a757265206661696c6564000000000000000000000000604482015290519081900360640190fd5b604080516001600160a01b03808f168252808e1660208301528183018790528b1660608201526080810184905290517f298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb529181900360a00190a1600554604080517f47ef3b3b0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038c811660248301528f811660448301528e811660648301526084820188905260a48201869052915191909216916347ef3b3b9160c480830192600092919082900301818387803b158015614b4357600080fd5b505af1158015614b57573d6000803e3d6000fd5b5060009250614b64915050565b975092955050505050505b94509492505050565b601154604080516370a0823160e01b815230600482015290516000926001600160a01b031691839183916370a08231916024808301926020929190829003018186803b158015614bc757600080fd5b505afa158015614bdb573d6000803e3d6000fd5b505050506040513d6020811015614bf157600080fd5b5051604080517f23b872dd0000000000000000000000000000000000000000000000000000000081526001600160a01b038881166004830152306024830152604482018890529151929350908416916323b872dd9160648082019260009290919082900301818387803b158015614c6757600080fd5b505af1158015614c7b573d6000803e3d6000fd5b5050505060003d60008114614c975760208114614ca157600080fd5b6000199150614cad565b60206000803e60005191505b5080614d00576040805162461bcd60e51b815260206004820152601860248201527f544f4b454e5f5452414e534645525f494e5f4641494c45440000000000000000604482015290519081900360640190fd5b601154604080516370a0823160e01b815230600482015290516000926001600160a01b0316916370a08231916024808301926020929190829003018186803b158015614d4b57600080fd5b505afa158015614d5f573d6000803e3d6000fd5b505050506040513d6020811015614d7557600080fd5b5051905082811015614dce576040805162461bcd60e51b815260206004820152601a60248201527f544f4b454e5f5452414e534645525f494e5f4f564552464c4f57000000000000604482015290519081900360640190fd5b9190910395945050505050565b6000806000614de8614e4c565b61241b86866000614df7614e4c565b600080614e0c670de0b6b3a764000087613d3b565b90925090506000826003811115614e1f57fe5b14614e3e57506040805160208101909152600081529092509050612451565b61244a818660000151613579565b6040518060200160405280600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10614ea057805160ff1916838001178555614ecd565b82800160010185558215614ecd579182015b82811115614ecd578251825591602001919060010190614eb2565b50614ed9929150614f8a565b5090565b6040805161010081019091528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b6040805160e0810190915280600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b604080516080810190915280600081526020016000815260200160008152602001600081525090565b610d2d91905b80821115614ed95760008155600101614f9056fe6f6e6c792061646d696e206d617920696e697469616c697a6520746865206d61726b65746d61726b6574206d6179206f6e6c7920626520696e697469616c697a6564206f6e6365696e697469616c2065786368616e67652072617465206d7573742062652067726561746572207468616e207a65726f2e73657474696e6720696e7465726573742072617465206d6f64656c206661696c65644d494e545f4e45575f4143434f554e545f42414c414e43455f43414c43554c4154494f4e5f4641494c4544626f72726f7742616c616e636553746f7265643a20626f72726f7742616c616e636553746f726564496e7465726e616c206661696c656452455041595f424f52524f575f4e45575f4143434f554e545f424f52524f575f42414c414e43455f43414c43554c4154494f4e5f4641494c454452455041595f424f52524f575f4e45575f544f54414c5f42414c414e43455f43414c43554c4154494f4e5f4641494c45444c49515549444154455f434f4d5054524f4c4c45525f43414c43554c4154455f414d4f554e545f5345495a455f4641494c454465786368616e67655261746553746f7265643a2065786368616e67655261746553746f726564496e7465726e616c206661696c65644d494e545f4e45575f544f54414c5f535550504c595f43414c43554c4154494f4e5f4641494c45446f6e65206f662072656465656d546f6b656e73496e206f722072656465656d416d6f756e74496e206d757374206265207a65726f72656475636520726573657276657320756e657870656374656420756e646572666c6f77a265627a7a72315820815296fff102257f63f102532004225c91abdb859bee75a72b7ef8dd087c1f0f64736f6c63430005100032"

// DeployCerc20 deploys a new Ethereum contract, binding an instance of Cerc20 to it.
func DeployCerc20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Cerc20, error) {
	parsed, err := abi.JSON(strings.NewReader(Cerc20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Cerc20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cerc20{Cerc20Caller: Cerc20Caller{contract: contract}, Cerc20Transactor: Cerc20Transactor{contract: contract}, Cerc20Filterer: Cerc20Filterer{contract: contract}}, nil
}

// Cerc20 is an auto generated Go binding around an Ethereum contract.
type Cerc20 struct {
	Cerc20Caller     // Read-only binding to the contract
	Cerc20Transactor // Write-only binding to the contract
	Cerc20Filterer   // Log filterer for contract events
}

// Cerc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Cerc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cerc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Cerc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cerc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Cerc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Cerc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Cerc20Session struct {
	Contract     *Cerc20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Cerc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Cerc20CallerSession struct {
	Contract *Cerc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Cerc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Cerc20TransactorSession struct {
	Contract     *Cerc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Cerc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Cerc20Raw struct {
	Contract *Cerc20 // Generic contract binding to access the raw methods on
}

// Cerc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Cerc20CallerRaw struct {
	Contract *Cerc20Caller // Generic read-only contract binding to access the raw methods on
}

// Cerc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Cerc20TransactorRaw struct {
	Contract *Cerc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCerc20 creates a new instance of Cerc20, bound to a specific deployed contract.
func NewCerc20(address common.Address, backend bind.ContractBackend) (*Cerc20, error) {
	contract, err := bindCerc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cerc20{Cerc20Caller: Cerc20Caller{contract: contract}, Cerc20Transactor: Cerc20Transactor{contract: contract}, Cerc20Filterer: Cerc20Filterer{contract: contract}}, nil
}

// NewCerc20Caller creates a new read-only instance of Cerc20, bound to a specific deployed contract.
func NewCerc20Caller(address common.Address, caller bind.ContractCaller) (*Cerc20Caller, error) {
	contract, err := bindCerc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Cerc20Caller{contract: contract}, nil
}

// NewCerc20Transactor creates a new write-only instance of Cerc20, bound to a specific deployed contract.
func NewCerc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Cerc20Transactor, error) {
	contract, err := bindCerc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Cerc20Transactor{contract: contract}, nil
}

// NewCerc20Filterer creates a new log filterer instance of Cerc20, bound to a specific deployed contract.
func NewCerc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Cerc20Filterer, error) {
	contract, err := bindCerc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Cerc20Filterer{contract: contract}, nil
}

// bindCerc20 binds a generic wrapper to an already deployed contract.
func bindCerc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Cerc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cerc20 *Cerc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cerc20.Contract.Cerc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cerc20 *Cerc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cerc20.Contract.Cerc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cerc20 *Cerc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cerc20.Contract.Cerc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cerc20 *Cerc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cerc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cerc20 *Cerc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cerc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cerc20 *Cerc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cerc20.Contract.contract.Transact(opts, method, params...)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Cerc20 *Cerc20Caller) AccrualBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "accrualBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Cerc20 *Cerc20Session) AccrualBlockNumber() (*big.Int, error) {
	return _Cerc20.Contract.AccrualBlockNumber(&_Cerc20.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Cerc20 *Cerc20CallerSession) AccrualBlockNumber() (*big.Int, error) {
	return _Cerc20.Contract.AccrualBlockNumber(&_Cerc20.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Cerc20 *Cerc20Caller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Cerc20 *Cerc20Session) Admin() (common.Address, error) {
	return _Cerc20.Contract.Admin(&_Cerc20.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Cerc20 *Cerc20CallerSession) Admin() (common.Address, error) {
	return _Cerc20.Contract.Admin(&_Cerc20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cerc20 *Cerc20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cerc20 *Cerc20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Cerc20.Contract.Allowance(&_Cerc20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cerc20 *Cerc20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Cerc20.Contract.Allowance(&_Cerc20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Cerc20 *Cerc20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Cerc20 *Cerc20Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Cerc20.Contract.BalanceOf(&_Cerc20.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Cerc20 *Cerc20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Cerc20.Contract.BalanceOf(&_Cerc20.CallOpts, owner)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Cerc20 *Cerc20Caller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "borrowBalanceStored", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Cerc20 *Cerc20Session) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Cerc20.Contract.BorrowBalanceStored(&_Cerc20.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Cerc20 *Cerc20CallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Cerc20.Contract.BorrowBalanceStored(&_Cerc20.CallOpts, account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Cerc20 *Cerc20Caller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "borrowIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Cerc20 *Cerc20Session) BorrowIndex() (*big.Int, error) {
	return _Cerc20.Contract.BorrowIndex(&_Cerc20.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Cerc20 *Cerc20CallerSession) BorrowIndex() (*big.Int, error) {
	return _Cerc20.Contract.BorrowIndex(&_Cerc20.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Cerc20 *Cerc20Caller) BorrowRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "borrowRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Cerc20 *Cerc20Session) BorrowRatePerBlock() (*big.Int, error) {
	return _Cerc20.Contract.BorrowRatePerBlock(&_Cerc20.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Cerc20 *Cerc20CallerSession) BorrowRatePerBlock() (*big.Int, error) {
	return _Cerc20.Contract.BorrowRatePerBlock(&_Cerc20.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Cerc20 *Cerc20Caller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Cerc20 *Cerc20Session) Comptroller() (common.Address, error) {
	return _Cerc20.Contract.Comptroller(&_Cerc20.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Cerc20 *Cerc20CallerSession) Comptroller() (common.Address, error) {
	return _Cerc20.Contract.Comptroller(&_Cerc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cerc20 *Cerc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cerc20 *Cerc20Session) Decimals() (uint8, error) {
	return _Cerc20.Contract.Decimals(&_Cerc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cerc20 *Cerc20CallerSession) Decimals() (uint8, error) {
	return _Cerc20.Contract.Decimals(&_Cerc20.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Cerc20 *Cerc20Caller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Cerc20 *Cerc20Session) ExchangeRateStored() (*big.Int, error) {
	return _Cerc20.Contract.ExchangeRateStored(&_Cerc20.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Cerc20 *Cerc20CallerSession) ExchangeRateStored() (*big.Int, error) {
	return _Cerc20.Contract.ExchangeRateStored(&_Cerc20.CallOpts)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Cerc20 *Cerc20Caller) GetAccountSnapshot(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "getAccountSnapshot", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Cerc20 *Cerc20Session) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Cerc20.Contract.GetAccountSnapshot(&_Cerc20.CallOpts, account)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Cerc20 *Cerc20CallerSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Cerc20.Contract.GetAccountSnapshot(&_Cerc20.CallOpts, account)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Cerc20 *Cerc20Caller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "getCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Cerc20 *Cerc20Session) GetCash() (*big.Int, error) {
	return _Cerc20.Contract.GetCash(&_Cerc20.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Cerc20 *Cerc20CallerSession) GetCash() (*big.Int, error) {
	return _Cerc20.Contract.GetCash(&_Cerc20.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Cerc20 *Cerc20Caller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Cerc20 *Cerc20Session) InterestRateModel() (common.Address, error) {
	return _Cerc20.Contract.InterestRateModel(&_Cerc20.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Cerc20 *Cerc20CallerSession) InterestRateModel() (common.Address, error) {
	return _Cerc20.Contract.InterestRateModel(&_Cerc20.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_Cerc20 *Cerc20Caller) IsCToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "isCToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_Cerc20 *Cerc20Session) IsCToken() (bool, error) {
	return _Cerc20.Contract.IsCToken(&_Cerc20.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_Cerc20 *Cerc20CallerSession) IsCToken() (bool, error) {
	return _Cerc20.Contract.IsCToken(&_Cerc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cerc20 *Cerc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cerc20 *Cerc20Session) Name() (string, error) {
	return _Cerc20.Contract.Name(&_Cerc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cerc20 *Cerc20CallerSession) Name() (string, error) {
	return _Cerc20.Contract.Name(&_Cerc20.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Cerc20 *Cerc20Caller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Cerc20 *Cerc20Session) PendingAdmin() (common.Address, error) {
	return _Cerc20.Contract.PendingAdmin(&_Cerc20.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Cerc20 *Cerc20CallerSession) PendingAdmin() (common.Address, error) {
	return _Cerc20.Contract.PendingAdmin(&_Cerc20.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Cerc20 *Cerc20Caller) ReserveFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "reserveFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Cerc20 *Cerc20Session) ReserveFactorMantissa() (*big.Int, error) {
	return _Cerc20.Contract.ReserveFactorMantissa(&_Cerc20.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Cerc20 *Cerc20CallerSession) ReserveFactorMantissa() (*big.Int, error) {
	return _Cerc20.Contract.ReserveFactorMantissa(&_Cerc20.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Cerc20 *Cerc20Caller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "supplyRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Cerc20 *Cerc20Session) SupplyRatePerBlock() (*big.Int, error) {
	return _Cerc20.Contract.SupplyRatePerBlock(&_Cerc20.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Cerc20 *Cerc20CallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _Cerc20.Contract.SupplyRatePerBlock(&_Cerc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cerc20 *Cerc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cerc20 *Cerc20Session) Symbol() (string, error) {
	return _Cerc20.Contract.Symbol(&_Cerc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cerc20 *Cerc20CallerSession) Symbol() (string, error) {
	return _Cerc20.Contract.Symbol(&_Cerc20.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Cerc20 *Cerc20Caller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Cerc20 *Cerc20Session) TotalBorrows() (*big.Int, error) {
	return _Cerc20.Contract.TotalBorrows(&_Cerc20.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Cerc20 *Cerc20CallerSession) TotalBorrows() (*big.Int, error) {
	return _Cerc20.Contract.TotalBorrows(&_Cerc20.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Cerc20 *Cerc20Caller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "totalReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Cerc20 *Cerc20Session) TotalReserves() (*big.Int, error) {
	return _Cerc20.Contract.TotalReserves(&_Cerc20.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Cerc20 *Cerc20CallerSession) TotalReserves() (*big.Int, error) {
	return _Cerc20.Contract.TotalReserves(&_Cerc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cerc20 *Cerc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cerc20 *Cerc20Session) TotalSupply() (*big.Int, error) {
	return _Cerc20.Contract.TotalSupply(&_Cerc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cerc20 *Cerc20CallerSession) TotalSupply() (*big.Int, error) {
	return _Cerc20.Contract.TotalSupply(&_Cerc20.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Cerc20 *Cerc20Caller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cerc20.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Cerc20 *Cerc20Session) Underlying() (common.Address, error) {
	return _Cerc20.Contract.Underlying(&_Cerc20.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Cerc20 *Cerc20CallerSession) Underlying() (common.Address, error) {
	return _Cerc20.Contract.Underlying(&_Cerc20.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Cerc20 *Cerc20Transactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Cerc20 *Cerc20Session) AcceptAdmin() (*types.Transaction, error) {
	return _Cerc20.Contract.AcceptAdmin(&_Cerc20.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _Cerc20.Contract.AcceptAdmin(&_Cerc20.TransactOpts)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_Cerc20 *Cerc20Transactor) AddReserves(opts *bind.TransactOpts, addAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "_addReserves", addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_Cerc20 *Cerc20Session) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.AddReserves(&_Cerc20.TransactOpts, addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.AddReserves(&_Cerc20.TransactOpts, addAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Cerc20 *Cerc20Transactor) ReduceReserves(opts *bind.TransactOpts, reduceAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "_reduceReserves", reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Cerc20 *Cerc20Session) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.ReduceReserves(&_Cerc20.TransactOpts, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.ReduceReserves(&_Cerc20.TransactOpts, reduceAmount)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Cerc20 *Cerc20Transactor) SetComptroller(opts *bind.TransactOpts, newComptroller common.Address) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "_setComptroller", newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Cerc20 *Cerc20Session) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _Cerc20.Contract.SetComptroller(&_Cerc20.TransactOpts, newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _Cerc20.Contract.SetComptroller(&_Cerc20.TransactOpts, newComptroller)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Cerc20 *Cerc20Transactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "_setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Cerc20 *Cerc20Session) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Cerc20.Contract.SetInterestRateModel(&_Cerc20.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Cerc20.Contract.SetInterestRateModel(&_Cerc20.TransactOpts, newInterestRateModel)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Cerc20 *Cerc20Transactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Cerc20 *Cerc20Session) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Cerc20.Contract.SetPendingAdmin(&_Cerc20.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Cerc20.Contract.SetPendingAdmin(&_Cerc20.TransactOpts, newPendingAdmin)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Cerc20 *Cerc20Transactor) SetReserveFactor(opts *bind.TransactOpts, newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "_setReserveFactor", newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Cerc20 *Cerc20Session) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.SetReserveFactor(&_Cerc20.TransactOpts, newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.SetReserveFactor(&_Cerc20.TransactOpts, newReserveFactorMantissa)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Cerc20 *Cerc20Transactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Cerc20 *Cerc20Session) AccrueInterest() (*types.Transaction, error) {
	return _Cerc20.Contract.AccrueInterest(&_Cerc20.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _Cerc20.Contract.AccrueInterest(&_Cerc20.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cerc20 *Cerc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cerc20 *Cerc20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.Approve(&_Cerc20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cerc20 *Cerc20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.Approve(&_Cerc20.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Cerc20 *Cerc20Transactor) BalanceOfUnderlying(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "balanceOfUnderlying", owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Cerc20 *Cerc20Session) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Cerc20.Contract.BalanceOfUnderlying(&_Cerc20.TransactOpts, owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Cerc20.Contract.BalanceOfUnderlying(&_Cerc20.TransactOpts, owner)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Cerc20 *Cerc20Transactor) Borrow(opts *bind.TransactOpts, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "borrow", borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Cerc20 *Cerc20Session) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.Borrow(&_Cerc20.TransactOpts, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.Borrow(&_Cerc20.TransactOpts, borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Cerc20 *Cerc20Transactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Cerc20 *Cerc20Session) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Cerc20.Contract.BorrowBalanceCurrent(&_Cerc20.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Cerc20.Contract.BorrowBalanceCurrent(&_Cerc20.TransactOpts, account)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Cerc20 *Cerc20Transactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Cerc20 *Cerc20Session) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Cerc20.Contract.ExchangeRateCurrent(&_Cerc20.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Cerc20.Contract.ExchangeRateCurrent(&_Cerc20.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x1a31d465.
//
// Solidity: function initialize(address underlying_, address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Cerc20 *Cerc20Transactor) Initialize(opts *bind.TransactOpts, underlying_ common.Address, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "initialize", underlying_, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1a31d465.
//
// Solidity: function initialize(address underlying_, address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Cerc20 *Cerc20Session) Initialize(underlying_ common.Address, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Cerc20.Contract.Initialize(&_Cerc20.TransactOpts, underlying_, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1a31d465.
//
// Solidity: function initialize(address underlying_, address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Cerc20 *Cerc20TransactorSession) Initialize(underlying_ common.Address, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Cerc20.Contract.Initialize(&_Cerc20.TransactOpts, underlying_, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Cerc20 *Cerc20Transactor) Initialize0(opts *bind.TransactOpts, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "initialize0", comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Cerc20 *Cerc20Session) Initialize0(comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Cerc20.Contract.Initialize0(&_Cerc20.TransactOpts, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Cerc20 *Cerc20TransactorSession) Initialize0(comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Cerc20.Contract.Initialize0(&_Cerc20.TransactOpts, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_Cerc20 *Cerc20Transactor) LiquidateBorrow(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "liquidateBorrow", borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_Cerc20 *Cerc20Session) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _Cerc20.Contract.LiquidateBorrow(&_Cerc20.TransactOpts, borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _Cerc20.Contract.LiquidateBorrow(&_Cerc20.TransactOpts, borrower, repayAmount, cTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_Cerc20 *Cerc20Transactor) Mint(opts *bind.TransactOpts, mintAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "mint", mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_Cerc20 *Cerc20Session) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.Mint(&_Cerc20.TransactOpts, mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.Mint(&_Cerc20.TransactOpts, mintAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Cerc20 *Cerc20Transactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Cerc20 *Cerc20Session) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.Redeem(&_Cerc20.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.Redeem(&_Cerc20.TransactOpts, redeemTokens)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Cerc20 *Cerc20Transactor) RedeemUnderlying(opts *bind.TransactOpts, redeemAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "redeemUnderlying", redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Cerc20 *Cerc20Session) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.RedeemUnderlying(&_Cerc20.TransactOpts, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.RedeemUnderlying(&_Cerc20.TransactOpts, redeemAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Cerc20 *Cerc20Transactor) RepayBorrow(opts *bind.TransactOpts, repayAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "repayBorrow", repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Cerc20 *Cerc20Session) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.RepayBorrow(&_Cerc20.TransactOpts, repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.RepayBorrow(&_Cerc20.TransactOpts, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Cerc20 *Cerc20Transactor) RepayBorrowBehalf(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "repayBorrowBehalf", borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Cerc20 *Cerc20Session) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.RepayBorrowBehalf(&_Cerc20.TransactOpts, borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.RepayBorrowBehalf(&_Cerc20.TransactOpts, borrower, repayAmount)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Cerc20 *Cerc20Transactor) Seize(opts *bind.TransactOpts, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "seize", liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Cerc20 *Cerc20Session) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.Seize(&_Cerc20.TransactOpts, liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.Seize(&_Cerc20.TransactOpts, liquidator, borrower, seizeTokens)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Cerc20 *Cerc20Transactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Cerc20 *Cerc20Session) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Cerc20.Contract.TotalBorrowsCurrent(&_Cerc20.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Cerc20 *Cerc20TransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Cerc20.Contract.TotalBorrowsCurrent(&_Cerc20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Cerc20 *Cerc20Transactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Cerc20 *Cerc20Session) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.Transfer(&_Cerc20.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Cerc20 *Cerc20TransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.Transfer(&_Cerc20.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Cerc20 *Cerc20Transactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cerc20.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Cerc20 *Cerc20Session) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.TransferFrom(&_Cerc20.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Cerc20 *Cerc20TransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cerc20.Contract.TransferFrom(&_Cerc20.TransactOpts, src, dst, amount)
}

// Cerc20AccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the Cerc20 contract.
type Cerc20AccrueInterestIterator struct {
	Event *Cerc20AccrueInterest // Event containing the contract specifics and raw log

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
func (it *Cerc20AccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cerc20AccrueInterest)
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
		it.Event = new(Cerc20AccrueInterest)
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
func (it *Cerc20AccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cerc20AccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cerc20AccrueInterest represents a AccrueInterest event raised by the Cerc20 contract.
type Cerc20AccrueInterest struct {
	CashPrior           *big.Int
	InterestAccumulated *big.Int
	BorrowIndex         *big.Int
	TotalBorrows        *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Cerc20 *Cerc20Filterer) FilterAccrueInterest(opts *bind.FilterOpts) (*Cerc20AccrueInterestIterator, error) {

	logs, sub, err := _Cerc20.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &Cerc20AccrueInterestIterator{contract: _Cerc20.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Cerc20 *Cerc20Filterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *Cerc20AccrueInterest) (event.Subscription, error) {

	logs, sub, err := _Cerc20.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cerc20AccrueInterest)
				if err := _Cerc20.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
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

// ParseAccrueInterest is a log parse operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Cerc20 *Cerc20Filterer) ParseAccrueInterest(log types.Log) (*Cerc20AccrueInterest, error) {
	event := new(Cerc20AccrueInterest)
	if err := _Cerc20.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Cerc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Cerc20 contract.
type Cerc20ApprovalIterator struct {
	Event *Cerc20Approval // Event containing the contract specifics and raw log

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
func (it *Cerc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cerc20Approval)
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
		it.Event = new(Cerc20Approval)
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
func (it *Cerc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cerc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cerc20Approval represents a Approval event raised by the Cerc20 contract.
type Cerc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Cerc20 *Cerc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Cerc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Cerc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Cerc20ApprovalIterator{contract: _Cerc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Cerc20 *Cerc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Cerc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Cerc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cerc20Approval)
				if err := _Cerc20.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Cerc20 *Cerc20Filterer) ParseApproval(log types.Log) (*Cerc20Approval, error) {
	event := new(Cerc20Approval)
	if err := _Cerc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Cerc20BorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Cerc20 contract.
type Cerc20BorrowIterator struct {
	Event *Cerc20Borrow // Event containing the contract specifics and raw log

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
func (it *Cerc20BorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cerc20Borrow)
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
		it.Event = new(Cerc20Borrow)
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
func (it *Cerc20BorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cerc20BorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cerc20Borrow represents a Borrow event raised by the Cerc20 contract.
type Cerc20Borrow struct {
	Borrower       common.Address
	BorrowAmount   *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Cerc20 *Cerc20Filterer) FilterBorrow(opts *bind.FilterOpts) (*Cerc20BorrowIterator, error) {

	logs, sub, err := _Cerc20.contract.FilterLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return &Cerc20BorrowIterator{contract: _Cerc20.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Cerc20 *Cerc20Filterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *Cerc20Borrow) (event.Subscription, error) {

	logs, sub, err := _Cerc20.contract.WatchLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cerc20Borrow)
				if err := _Cerc20.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Cerc20 *Cerc20Filterer) ParseBorrow(log types.Log) (*Cerc20Borrow, error) {
	event := new(Cerc20Borrow)
	if err := _Cerc20.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Cerc20FailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Cerc20 contract.
type Cerc20FailureIterator struct {
	Event *Cerc20Failure // Event containing the contract specifics and raw log

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
func (it *Cerc20FailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cerc20Failure)
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
		it.Event = new(Cerc20Failure)
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
func (it *Cerc20FailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cerc20FailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cerc20Failure represents a Failure event raised by the Cerc20 contract.
type Cerc20Failure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Cerc20 *Cerc20Filterer) FilterFailure(opts *bind.FilterOpts) (*Cerc20FailureIterator, error) {

	logs, sub, err := _Cerc20.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &Cerc20FailureIterator{contract: _Cerc20.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Cerc20 *Cerc20Filterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *Cerc20Failure) (event.Subscription, error) {

	logs, sub, err := _Cerc20.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cerc20Failure)
				if err := _Cerc20.contract.UnpackLog(event, "Failure", log); err != nil {
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

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Cerc20 *Cerc20Filterer) ParseFailure(log types.Log) (*Cerc20Failure, error) {
	event := new(Cerc20Failure)
	if err := _Cerc20.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Cerc20LiquidateBorrowIterator is returned from FilterLiquidateBorrow and is used to iterate over the raw logs and unpacked data for LiquidateBorrow events raised by the Cerc20 contract.
type Cerc20LiquidateBorrowIterator struct {
	Event *Cerc20LiquidateBorrow // Event containing the contract specifics and raw log

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
func (it *Cerc20LiquidateBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cerc20LiquidateBorrow)
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
		it.Event = new(Cerc20LiquidateBorrow)
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
func (it *Cerc20LiquidateBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cerc20LiquidateBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cerc20LiquidateBorrow represents a LiquidateBorrow event raised by the Cerc20 contract.
type Cerc20LiquidateBorrow struct {
	Liquidator       common.Address
	Borrower         common.Address
	RepayAmount      *big.Int
	CTokenCollateral common.Address
	SeizeTokens      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidateBorrow is a free log retrieval operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_Cerc20 *Cerc20Filterer) FilterLiquidateBorrow(opts *bind.FilterOpts) (*Cerc20LiquidateBorrowIterator, error) {

	logs, sub, err := _Cerc20.contract.FilterLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return &Cerc20LiquidateBorrowIterator{contract: _Cerc20.contract, event: "LiquidateBorrow", logs: logs, sub: sub}, nil
}

// WatchLiquidateBorrow is a free log subscription operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_Cerc20 *Cerc20Filterer) WatchLiquidateBorrow(opts *bind.WatchOpts, sink chan<- *Cerc20LiquidateBorrow) (event.Subscription, error) {

	logs, sub, err := _Cerc20.contract.WatchLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cerc20LiquidateBorrow)
				if err := _Cerc20.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
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

// ParseLiquidateBorrow is a log parse operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_Cerc20 *Cerc20Filterer) ParseLiquidateBorrow(log types.Log) (*Cerc20LiquidateBorrow, error) {
	event := new(Cerc20LiquidateBorrow)
	if err := _Cerc20.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Cerc20MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Cerc20 contract.
type Cerc20MintIterator struct {
	Event *Cerc20Mint // Event containing the contract specifics and raw log

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
func (it *Cerc20MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cerc20Mint)
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
		it.Event = new(Cerc20Mint)
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
func (it *Cerc20MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cerc20MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cerc20Mint represents a Mint event raised by the Cerc20 contract.
type Cerc20Mint struct {
	Minter     common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Cerc20 *Cerc20Filterer) FilterMint(opts *bind.FilterOpts) (*Cerc20MintIterator, error) {

	logs, sub, err := _Cerc20.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &Cerc20MintIterator{contract: _Cerc20.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Cerc20 *Cerc20Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *Cerc20Mint) (event.Subscription, error) {

	logs, sub, err := _Cerc20.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cerc20Mint)
				if err := _Cerc20.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Cerc20 *Cerc20Filterer) ParseMint(log types.Log) (*Cerc20Mint, error) {
	event := new(Cerc20Mint)
	if err := _Cerc20.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Cerc20NewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Cerc20 contract.
type Cerc20NewAdminIterator struct {
	Event *Cerc20NewAdmin // Event containing the contract specifics and raw log

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
func (it *Cerc20NewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cerc20NewAdmin)
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
		it.Event = new(Cerc20NewAdmin)
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
func (it *Cerc20NewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cerc20NewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cerc20NewAdmin represents a NewAdmin event raised by the Cerc20 contract.
type Cerc20NewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Cerc20 *Cerc20Filterer) FilterNewAdmin(opts *bind.FilterOpts) (*Cerc20NewAdminIterator, error) {

	logs, sub, err := _Cerc20.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &Cerc20NewAdminIterator{contract: _Cerc20.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Cerc20 *Cerc20Filterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *Cerc20NewAdmin) (event.Subscription, error) {

	logs, sub, err := _Cerc20.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cerc20NewAdmin)
				if err := _Cerc20.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Cerc20 *Cerc20Filterer) ParseNewAdmin(log types.Log) (*Cerc20NewAdmin, error) {
	event := new(Cerc20NewAdmin)
	if err := _Cerc20.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Cerc20NewComptrollerIterator is returned from FilterNewComptroller and is used to iterate over the raw logs and unpacked data for NewComptroller events raised by the Cerc20 contract.
type Cerc20NewComptrollerIterator struct {
	Event *Cerc20NewComptroller // Event containing the contract specifics and raw log

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
func (it *Cerc20NewComptrollerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cerc20NewComptroller)
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
		it.Event = new(Cerc20NewComptroller)
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
func (it *Cerc20NewComptrollerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cerc20NewComptrollerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cerc20NewComptroller represents a NewComptroller event raised by the Cerc20 contract.
type Cerc20NewComptroller struct {
	OldComptroller common.Address
	NewComptroller common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewComptroller is a free log retrieval operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Cerc20 *Cerc20Filterer) FilterNewComptroller(opts *bind.FilterOpts) (*Cerc20NewComptrollerIterator, error) {

	logs, sub, err := _Cerc20.contract.FilterLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return &Cerc20NewComptrollerIterator{contract: _Cerc20.contract, event: "NewComptroller", logs: logs, sub: sub}, nil
}

// WatchNewComptroller is a free log subscription operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Cerc20 *Cerc20Filterer) WatchNewComptroller(opts *bind.WatchOpts, sink chan<- *Cerc20NewComptroller) (event.Subscription, error) {

	logs, sub, err := _Cerc20.contract.WatchLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cerc20NewComptroller)
				if err := _Cerc20.contract.UnpackLog(event, "NewComptroller", log); err != nil {
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

// ParseNewComptroller is a log parse operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Cerc20 *Cerc20Filterer) ParseNewComptroller(log types.Log) (*Cerc20NewComptroller, error) {
	event := new(Cerc20NewComptroller)
	if err := _Cerc20.contract.UnpackLog(event, "NewComptroller", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Cerc20NewMarketInterestRateModelIterator is returned from FilterNewMarketInterestRateModel and is used to iterate over the raw logs and unpacked data for NewMarketInterestRateModel events raised by the Cerc20 contract.
type Cerc20NewMarketInterestRateModelIterator struct {
	Event *Cerc20NewMarketInterestRateModel // Event containing the contract specifics and raw log

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
func (it *Cerc20NewMarketInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cerc20NewMarketInterestRateModel)
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
		it.Event = new(Cerc20NewMarketInterestRateModel)
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
func (it *Cerc20NewMarketInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cerc20NewMarketInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cerc20NewMarketInterestRateModel represents a NewMarketInterestRateModel event raised by the Cerc20 contract.
type Cerc20NewMarketInterestRateModel struct {
	OldInterestRateModel common.Address
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewMarketInterestRateModel is a free log retrieval operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Cerc20 *Cerc20Filterer) FilterNewMarketInterestRateModel(opts *bind.FilterOpts) (*Cerc20NewMarketInterestRateModelIterator, error) {

	logs, sub, err := _Cerc20.contract.FilterLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &Cerc20NewMarketInterestRateModelIterator{contract: _Cerc20.contract, event: "NewMarketInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewMarketInterestRateModel is a free log subscription operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Cerc20 *Cerc20Filterer) WatchNewMarketInterestRateModel(opts *bind.WatchOpts, sink chan<- *Cerc20NewMarketInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _Cerc20.contract.WatchLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cerc20NewMarketInterestRateModel)
				if err := _Cerc20.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
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

// ParseNewMarketInterestRateModel is a log parse operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Cerc20 *Cerc20Filterer) ParseNewMarketInterestRateModel(log types.Log) (*Cerc20NewMarketInterestRateModel, error) {
	event := new(Cerc20NewMarketInterestRateModel)
	if err := _Cerc20.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Cerc20NewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the Cerc20 contract.
type Cerc20NewPendingAdminIterator struct {
	Event *Cerc20NewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *Cerc20NewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cerc20NewPendingAdmin)
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
		it.Event = new(Cerc20NewPendingAdmin)
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
func (it *Cerc20NewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cerc20NewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cerc20NewPendingAdmin represents a NewPendingAdmin event raised by the Cerc20 contract.
type Cerc20NewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Cerc20 *Cerc20Filterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*Cerc20NewPendingAdminIterator, error) {

	logs, sub, err := _Cerc20.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &Cerc20NewPendingAdminIterator{contract: _Cerc20.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Cerc20 *Cerc20Filterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *Cerc20NewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _Cerc20.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cerc20NewPendingAdmin)
				if err := _Cerc20.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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

// ParseNewPendingAdmin is a log parse operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Cerc20 *Cerc20Filterer) ParseNewPendingAdmin(log types.Log) (*Cerc20NewPendingAdmin, error) {
	event := new(Cerc20NewPendingAdmin)
	if err := _Cerc20.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Cerc20NewReserveFactorIterator is returned from FilterNewReserveFactor and is used to iterate over the raw logs and unpacked data for NewReserveFactor events raised by the Cerc20 contract.
type Cerc20NewReserveFactorIterator struct {
	Event *Cerc20NewReserveFactor // Event containing the contract specifics and raw log

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
func (it *Cerc20NewReserveFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cerc20NewReserveFactor)
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
		it.Event = new(Cerc20NewReserveFactor)
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
func (it *Cerc20NewReserveFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cerc20NewReserveFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cerc20NewReserveFactor represents a NewReserveFactor event raised by the Cerc20 contract.
type Cerc20NewReserveFactor struct {
	OldReserveFactorMantissa *big.Int
	NewReserveFactorMantissa *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewReserveFactor is a free log retrieval operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Cerc20 *Cerc20Filterer) FilterNewReserveFactor(opts *bind.FilterOpts) (*Cerc20NewReserveFactorIterator, error) {

	logs, sub, err := _Cerc20.contract.FilterLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return &Cerc20NewReserveFactorIterator{contract: _Cerc20.contract, event: "NewReserveFactor", logs: logs, sub: sub}, nil
}

// WatchNewReserveFactor is a free log subscription operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Cerc20 *Cerc20Filterer) WatchNewReserveFactor(opts *bind.WatchOpts, sink chan<- *Cerc20NewReserveFactor) (event.Subscription, error) {

	logs, sub, err := _Cerc20.contract.WatchLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cerc20NewReserveFactor)
				if err := _Cerc20.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
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

// ParseNewReserveFactor is a log parse operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Cerc20 *Cerc20Filterer) ParseNewReserveFactor(log types.Log) (*Cerc20NewReserveFactor, error) {
	event := new(Cerc20NewReserveFactor)
	if err := _Cerc20.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Cerc20RedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Cerc20 contract.
type Cerc20RedeemIterator struct {
	Event *Cerc20Redeem // Event containing the contract specifics and raw log

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
func (it *Cerc20RedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cerc20Redeem)
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
		it.Event = new(Cerc20Redeem)
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
func (it *Cerc20RedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cerc20RedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cerc20Redeem represents a Redeem event raised by the Cerc20 contract.
type Cerc20Redeem struct {
	Redeemer     common.Address
	RedeemAmount *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Cerc20 *Cerc20Filterer) FilterRedeem(opts *bind.FilterOpts) (*Cerc20RedeemIterator, error) {

	logs, sub, err := _Cerc20.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &Cerc20RedeemIterator{contract: _Cerc20.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Cerc20 *Cerc20Filterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *Cerc20Redeem) (event.Subscription, error) {

	logs, sub, err := _Cerc20.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cerc20Redeem)
				if err := _Cerc20.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Cerc20 *Cerc20Filterer) ParseRedeem(log types.Log) (*Cerc20Redeem, error) {
	event := new(Cerc20Redeem)
	if err := _Cerc20.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Cerc20RepayBorrowIterator is returned from FilterRepayBorrow and is used to iterate over the raw logs and unpacked data for RepayBorrow events raised by the Cerc20 contract.
type Cerc20RepayBorrowIterator struct {
	Event *Cerc20RepayBorrow // Event containing the contract specifics and raw log

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
func (it *Cerc20RepayBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cerc20RepayBorrow)
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
		it.Event = new(Cerc20RepayBorrow)
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
func (it *Cerc20RepayBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cerc20RepayBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cerc20RepayBorrow represents a RepayBorrow event raised by the Cerc20 contract.
type Cerc20RepayBorrow struct {
	Payer          common.Address
	Borrower       common.Address
	RepayAmount    *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepayBorrow is a free log retrieval operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Cerc20 *Cerc20Filterer) FilterRepayBorrow(opts *bind.FilterOpts) (*Cerc20RepayBorrowIterator, error) {

	logs, sub, err := _Cerc20.contract.FilterLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return &Cerc20RepayBorrowIterator{contract: _Cerc20.contract, event: "RepayBorrow", logs: logs, sub: sub}, nil
}

// WatchRepayBorrow is a free log subscription operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Cerc20 *Cerc20Filterer) WatchRepayBorrow(opts *bind.WatchOpts, sink chan<- *Cerc20RepayBorrow) (event.Subscription, error) {

	logs, sub, err := _Cerc20.contract.WatchLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cerc20RepayBorrow)
				if err := _Cerc20.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
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

// ParseRepayBorrow is a log parse operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Cerc20 *Cerc20Filterer) ParseRepayBorrow(log types.Log) (*Cerc20RepayBorrow, error) {
	event := new(Cerc20RepayBorrow)
	if err := _Cerc20.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Cerc20ReservesAddedIterator is returned from FilterReservesAdded and is used to iterate over the raw logs and unpacked data for ReservesAdded events raised by the Cerc20 contract.
type Cerc20ReservesAddedIterator struct {
	Event *Cerc20ReservesAdded // Event containing the contract specifics and raw log

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
func (it *Cerc20ReservesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cerc20ReservesAdded)
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
		it.Event = new(Cerc20ReservesAdded)
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
func (it *Cerc20ReservesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cerc20ReservesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cerc20ReservesAdded represents a ReservesAdded event raised by the Cerc20 contract.
type Cerc20ReservesAdded struct {
	Benefactor       common.Address
	AddAmount        *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesAdded is a free log retrieval operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Cerc20 *Cerc20Filterer) FilterReservesAdded(opts *bind.FilterOpts) (*Cerc20ReservesAddedIterator, error) {

	logs, sub, err := _Cerc20.contract.FilterLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return &Cerc20ReservesAddedIterator{contract: _Cerc20.contract, event: "ReservesAdded", logs: logs, sub: sub}, nil
}

// WatchReservesAdded is a free log subscription operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Cerc20 *Cerc20Filterer) WatchReservesAdded(opts *bind.WatchOpts, sink chan<- *Cerc20ReservesAdded) (event.Subscription, error) {

	logs, sub, err := _Cerc20.contract.WatchLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cerc20ReservesAdded)
				if err := _Cerc20.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
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

// ParseReservesAdded is a log parse operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Cerc20 *Cerc20Filterer) ParseReservesAdded(log types.Log) (*Cerc20ReservesAdded, error) {
	event := new(Cerc20ReservesAdded)
	if err := _Cerc20.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Cerc20ReservesReducedIterator is returned from FilterReservesReduced and is used to iterate over the raw logs and unpacked data for ReservesReduced events raised by the Cerc20 contract.
type Cerc20ReservesReducedIterator struct {
	Event *Cerc20ReservesReduced // Event containing the contract specifics and raw log

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
func (it *Cerc20ReservesReducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cerc20ReservesReduced)
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
		it.Event = new(Cerc20ReservesReduced)
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
func (it *Cerc20ReservesReducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cerc20ReservesReducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cerc20ReservesReduced represents a ReservesReduced event raised by the Cerc20 contract.
type Cerc20ReservesReduced struct {
	Admin            common.Address
	ReduceAmount     *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesReduced is a free log retrieval operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Cerc20 *Cerc20Filterer) FilterReservesReduced(opts *bind.FilterOpts) (*Cerc20ReservesReducedIterator, error) {

	logs, sub, err := _Cerc20.contract.FilterLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return &Cerc20ReservesReducedIterator{contract: _Cerc20.contract, event: "ReservesReduced", logs: logs, sub: sub}, nil
}

// WatchReservesReduced is a free log subscription operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Cerc20 *Cerc20Filterer) WatchReservesReduced(opts *bind.WatchOpts, sink chan<- *Cerc20ReservesReduced) (event.Subscription, error) {

	logs, sub, err := _Cerc20.contract.WatchLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cerc20ReservesReduced)
				if err := _Cerc20.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
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

// ParseReservesReduced is a log parse operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Cerc20 *Cerc20Filterer) ParseReservesReduced(log types.Log) (*Cerc20ReservesReduced, error) {
	event := new(Cerc20ReservesReduced)
	if err := _Cerc20.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Cerc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Cerc20 contract.
type Cerc20TransferIterator struct {
	Event *Cerc20Transfer // Event containing the contract specifics and raw log

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
func (it *Cerc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Cerc20Transfer)
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
		it.Event = new(Cerc20Transfer)
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
func (it *Cerc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Cerc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Cerc20Transfer represents a Transfer event raised by the Cerc20 contract.
type Cerc20Transfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Cerc20 *Cerc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Cerc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Cerc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Cerc20TransferIterator{contract: _Cerc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Cerc20 *Cerc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Cerc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Cerc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Cerc20Transfer)
				if err := _Cerc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Cerc20 *Cerc20Filterer) ParseTransfer(log types.Log) (*Cerc20Transfer, error) {
	event := new(Cerc20Transfer)
	if err := _Cerc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
