// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lending_pool_core

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

// LendingPoolCoreABI is the input ABI used to generate the binding from.
const LendingPoolCoreABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"CORE_REVISION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reservesList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressesProvider\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lendingPoolAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"ReserveUpdated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addressesProvider\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_isFirstDeposit\",\"type\":\"bool\"}],\"name\":\"updateStateOnDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_amountRedeemed\",\"type\":\"uint256\"},{\"name\":\"_userRedeemedEverything\",\"type\":\"bool\"}],\"name\":\"updateStateOnRedeem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_availableLiquidityBefore\",\"type\":\"uint256\"},{\"name\":\"_income\",\"type\":\"uint256\"},{\"name\":\"_protocolFee\",\"type\":\"uint256\"}],\"name\":\"updateStateOnFlashLoan\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_amountBorrowed\",\"type\":\"uint256\"},{\"name\":\"_borrowFee\",\"type\":\"uint256\"},{\"name\":\"_rateMode\",\"type\":\"uint8\"}],\"name\":\"updateStateOnBorrow\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_paybackAmountMinusFees\",\"type\":\"uint256\"},{\"name\":\"_originationFeeRepaid\",\"type\":\"uint256\"},{\"name\":\"_balanceIncrease\",\"type\":\"uint256\"},{\"name\":\"_repaidWholeLoan\",\"type\":\"bool\"}],\"name\":\"updateStateOnRepay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_principalBorrowBalance\",\"type\":\"uint256\"},{\"name\":\"_compoundedBorrowBalance\",\"type\":\"uint256\"},{\"name\":\"_balanceIncrease\",\"type\":\"uint256\"},{\"name\":\"_currentRateMode\",\"type\":\"uint8\"}],\"name\":\"updateStateOnSwapRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_principalReserve\",\"type\":\"address\"},{\"name\":\"_collateralReserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_amountToLiquidate\",\"type\":\"uint256\"},{\"name\":\"_collateralToLiquidate\",\"type\":\"uint256\"},{\"name\":\"_feeLiquidated\",\"type\":\"uint256\"},{\"name\":\"_liquidatedCollateralForFee\",\"type\":\"uint256\"},{\"name\":\"_balanceIncrease\",\"type\":\"uint256\"},{\"name\":\"_liquidatorReceivesAToken\",\"type\":\"bool\"}],\"name\":\"updateStateOnLiquidation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_balanceIncrease\",\"type\":\"uint256\"}],\"name\":\"updateStateOnRebalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_useAsCollateral\",\"type\":\"bool\"}],\"name\":\"setUserUseReserveAsCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferToUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"transferToFeeCollectionAddress\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"liquidateFee\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferToReserve\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserBasicReserveData\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"isUserAllowedToBorrowAtStable\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserUnderlyingAssetBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveInterestRateStrategyAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveATokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveAvailableLiquidity\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveTotalLiquidity\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedIncome\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveTotalBorrows\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveTotalBorrowsStable\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveTotalBorrowsVariable\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveLiquidationThreshold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveLiquidationBonus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveCurrentVariableBorrowRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveCurrentStableBorrowRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveCurrentAverageStableBorrowRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveCurrentLiquidityRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveLiquidityCumulativeIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveVariableBorrowsCumulativeIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveConfiguration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveDecimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"isReserveBorrowingEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"isReserveUsageAsCollateralEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveIsStableBorrowRateEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveIsActive\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveIsFreezed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveLastUpdate\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint40\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveUtilizationRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isUserUseReserveAsCollateralEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserOriginationFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserCurrentBorrowRateMode\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserCurrentStableBorrowRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserBorrowBalances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserVariableBorrowCumulativeIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserLastUpdate\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refreshConfiguration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_aTokenAddress\",\"type\":\"address\"},{\"name\":\"_decimals\",\"type\":\"uint256\"},{\"name\":\"_interestRateStrategyAddress\",\"type\":\"address\"}],\"name\":\"initReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserveToRemove\",\"type\":\"address\"}],\"name\":\"removeLastAddedReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_rateStrategyAddress\",\"type\":\"address\"}],\"name\":\"setReserveInterestRateStrategyAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_stableBorrowRateEnabled\",\"type\":\"bool\"}],\"name\":\"enableBorrowingOnReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"disableBorrowingOnReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_baseLTVasCollateral\",\"type\":\"uint256\"},{\"name\":\"_liquidationThreshold\",\"type\":\"uint256\"},{\"name\":\"_liquidationBonus\",\"type\":\"uint256\"}],\"name\":\"enableReserveAsCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"disableReserveAsCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"enableReserveStableBorrowRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"disableReserveStableBorrowRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"activateReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"deactivateReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"freezeReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"unfreezeReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_ltv\",\"type\":\"uint256\"}],\"name\":\"setReserveBaseLTVasCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setReserveLiquidationThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_bonus\",\"type\":\"uint256\"}],\"name\":\"setReserveLiquidationBonus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserve\",\"type\":\"address\"},{\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"setReserveDecimals\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// LendingPoolCoreBin is the compiled bytecode used for deploying new contracts.
var LendingPoolCoreBin = "0x60806040526000805534801561001457600080fd5b50615bfb80620000256000396000f3fe60806040526004361061049b5760003560e01c8063a2353fdc1161025e578063d15e005311610143578063e8ae2f5b116100bb578063f61483111161008a578063fa51854c1161006f578063fa51854c14611503578063fa93b2a514611548578063feab31ac1461158b5761049b565b8063f614831114611444578063f6ea8d76146114c85761049b565b8063e8ae2f5b14611370578063eede87c1146113a3578063ef1f9373146113de578063f054ab46146114115761049b565b8063dae4c4e711610112578063e2174d86116100f7578063e2174d86146112c7578063e24030191461130a578063e6d181901461133d5761049b565b8063dae4c4e714611259578063e10076ad1461128c5761049b565b8063d15e005314611181578063d3ae26b3146111b4578063d466016f146111c9578063da12d96f146112025761049b565b8063bd7fd79a116101d6578063c540148e116101a5578063c76a6c9c1161018a578063c76a6c9c146110df578063c7d1423714611112578063d06e2ec11461114e5761049b565b8063c540148e14611097578063c72c4d10146110ca5761049b565b8063bd7fd79a14610fcb578063bfacad8414610ffe578063c33cfd9014611031578063c4d66de8146110645761049b565b8063afcdbea31161022d578063b75d6f3411610212578063b75d6f3414610f1a578063b8c0f74514610f4d578063bcd6ffa414610f805761049b565b8063afcdbea314610e9c578063b701d09314610ee75761049b565b8063a2353fdc14610dae578063a5bc826c14610de1578063a8dc0f4514610e26578063af825b0714610e595761049b565b80635cf2e656116103845780637aca76eb116102fc578063906c0a41116102cb5780639e3c4f3b116102b05780639e3c4f3b14610ce75780639e67484814610d225780639fb8afcd14610d555761049b565b8063906c0a4114610c8157806398bd473714610cb45761049b565b80637aca76eb14610bb25780637f90fec514610be557806388079d8814610c185780638f385c2214610c4b5761049b565b806366d103f3116103535780636ae14416116103385780636ae1441614610b0b5780636fffab0c14610b3e57806370fb84f414610b795761049b565b806366d103f314610a6557806368beb4d614610aa05761049b565b80635cf2e656146109895780635fc526ff146109bc5780636468108314610a1757806366bbd92814610a2c5761049b565b80633443a14b1161041757806345330a40116103e65780634a08accb116103cb5780634a08accb146108df5780634f144609146109125780634fe7a6e51461095f5761049b565b806345330a401461086357806346bc0f28146108ac5761049b565b80633443a14b1461073c57806334b3beee1461077557806337ac6fe4146107c45780633e72a454146108305761049b565b806318a4dbca1161046e5780631ca19f19116104535780631ca19f191461066c5780631d2118f9146106cb57806328fcf4d3146107065761049b565b806318a4dbca146105fe57806318f9bbae146106395761049b565b806305075d6e146104e65780630902f1ac1461052d57806309ac2953146105925780630c7de4e9146105d7575b6104a4336115c6565b15156104e457604051600160e51b62461bcd0281526004018080602001828103825260368152602001806159586036913960400191505060405180910390fd5b005b3480156104f257600080fd5b506105196004803603602081101561050957600080fd5b50356001600160a01b0316611604565b604080519115158252519081900360200190f35b34801561053957600080fd5b5061054261162c565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561057e578181015183820152602001610566565b505050509050019250505060405180910390f35b34801561059e57600080fd5b506104e4600480360360808110156105b557600080fd5b506001600160a01b03813516906020810135906040810135906060013561168e565b3480156105e357600080fd5b506105ec61175f565b60408051918252519081900360200190f35b34801561060a57600080fd5b506105ec6004803603604081101561062157600080fd5b506001600160a01b0381358116916020013516611764565b34801561064557600080fd5b506105196004803603602081101561065c57600080fd5b50356001600160a01b0316611810565b34801561067857600080fd5b506106a76004803603604081101561068f57600080fd5b506001600160a01b0381358116916020013516611838565b604051808260028111156106b757fe5b60ff16815260200191505060405180910390f35b3480156106d757600080fd5b506104e4600480360360408110156106ee57600080fd5b506001600160a01b038135811691602001351661188f565b6104e46004803603606081101561071c57600080fd5b506001600160a01b0381358116916020810135909116906040013561197b565b34801561074857600080fd5b506104e46004803603604081101561075f57600080fd5b506001600160a01b038135169060200135611b55565b34801561078157600080fd5b506107a86004803603602081101561079857600080fd5b50356001600160a01b0316611c2f565b604080516001600160a01b039092168252519081900360200190f35b3480156107d057600080fd5b50610817600480360360a08110156107e757600080fd5b5080356001600160a01b03908116916020810135909116906040810135906060810135906080013560ff16611c50565b6040805192835260208301919091528051918290030190f35b34801561083c57600080fd5b506104e46004803603602081101561085357600080fd5b50356001600160a01b0316611cf3565b34801561086f57600080fd5b506104e46004803603608081101561088657600080fd5b506001600160a01b03813581169160208101358216916040820135916060013516611df0565b3480156108b857600080fd5b506105ec600480360360208110156108cf57600080fd5b50356001600160a01b0316611f69565b3480156108eb57600080fd5b506105196004803603602081101561090257600080fd5b50356001600160a01b0316611f87565b34801561091e57600080fd5b506109456004803603602081101561093557600080fd5b50356001600160a01b0316611faf565b6040805164ffffffffff9092168252519081900360200190f35b34801561096b57600080fd5b506107a86004803603602081101561098257600080fd5b5035611fdb565b34801561099557600080fd5b50610519600480360360208110156109ac57600080fd5b50356001600160a01b0316612003565b3480156109c857600080fd5b506109ef600480360360208110156109df57600080fd5b50356001600160a01b0316612041565b6040805194855260208501939093528383019190915215156060830152519081900360800190f35b348015610a2357600080fd5b506104e461207f565b348015610a3857600080fd5b506104e460048036036040811015610a4f57600080fd5b506001600160a01b038135169060200135612144565b348015610a7157600080fd5b506105ec60048036036040811015610a8857600080fd5b506001600160a01b038135811691602001351661221e565b348015610aac57600080fd5b506104e46004803603610120811015610ac457600080fd5b506001600160a01b03813581169160208101358216916040820135169060608101359060808101359060a08101359060c08101359060e08101359061010001351515612254565b348015610b1757600080fd5b506107a860048036036020811015610b2e57600080fd5b50356001600160a01b03166122fb565b348015610b4a57600080fd5b506105ec60048036036040811015610b6157600080fd5b506001600160a01b038135811691602001351661231c565b348015610b8557600080fd5b506104e460048036036040811015610b9c57600080fd5b506001600160a01b03813516906020013561234b565b348015610bbe57600080fd5b506104e460048036036020811015610bd557600080fd5b50356001600160a01b0316612425565b348015610bf157600080fd5b506105ec60048036036020811015610c0857600080fd5b50356001600160a01b0316612528565b348015610c2457600080fd5b506105ec60048036036020811015610c3b57600080fd5b50356001600160a01b0316612546565b6104e460048036036060811015610c6157600080fd5b506001600160a01b03813581169160208101359160409091013516612689565b348015610c8d57600080fd5b506105ec60048036036020811015610ca457600080fd5b50356001600160a01b031661279f565b348015610cc057600080fd5b506105ec60048036036020811015610cd757600080fd5b50356001600160a01b0316612851565b348015610cf357600080fd5b5061051960048036036040811015610d0a57600080fd5b506001600160a01b038135811691602001351661286f565b348015610d2e57600080fd5b5061051960048036036020811015610d4557600080fd5b50356001600160a01b03166128aa565b348015610d6157600080fd5b50610d9060048036036040811015610d7857600080fd5b506001600160a01b03813581169160200135166128d2565b60408051938452602084019290925282820152519081900360600190f35b348015610dba57600080fd5b506105ec60048036036020811015610dd157600080fd5b50356001600160a01b031661295f565b348015610ded57600080fd5b506104e460048036036080811015610e0457600080fd5b506001600160a01b03813516906020810135906040810135906060013561297d565b348015610e3257600080fd5b506104e460048036036020811015610e4957600080fd5b50356001600160a01b0316612aea565b348015610e6557600080fd5b506105ec60048036036060811015610e7c57600080fd5b506001600160a01b03813581169160208101359091169060400135612c38565b348015610ea857600080fd5b506104e460048036036080811015610ebf57600080fd5b506001600160a01b038135811691602081013590911690604081013590606001351515612cdb565b348015610ef357600080fd5b506105ec60048036036020811015610f0a57600080fd5b50356001600160a01b0316612d66565b348015610f2657600080fd5b506104e460048036036020811015610f3d57600080fd5b50356001600160a01b0316612d84565b348015610f5957600080fd5b506104e460048036036020811015610f7057600080fd5b50356001600160a01b0316612edd565b348015610f8c57600080fd5b506104e460048036036080811015610fa357600080fd5b506001600160a01b038135811691602081013590911690604081013590606001351515612fda565b348015610fd757600080fd5b506105ec60048036036020811015610fee57600080fd5b50356001600160a01b0316613065565b34801561100a57600080fd5b506105ec6004803603602081101561102157600080fd5b50356001600160a01b0316613080565b34801561103d57600080fd5b506105ec6004803603602081101561105457600080fd5b50356001600160a01b03166130ec565b34801561107057600080fd5b506104e46004803603602081101561108757600080fd5b50356001600160a01b0316613125565b3480156110a357600080fd5b506105ec600480360360208110156110ba57600080fd5b50356001600160a01b03166131e9565b3480156110d657600080fd5b506107a8613207565b3480156110eb57600080fd5b506105ec6004803603602081101561110257600080fd5b50356001600160a01b0316613216565b6104e46004803603608081101561112857600080fd5b506001600160a01b03813581169160208101358216916040820135916060013516613234565b34801561115a57600080fd5b506104e46004803603602081101561117157600080fd5b50356001600160a01b03166133f5565b34801561118d57600080fd5b506105ec600480360360208110156111a457600080fd5b50356001600160a01b0316613618565b3480156111c057600080fd5b506107a8613639565b3480156111d557600080fd5b506104e4600480360360408110156111ec57600080fd5b506001600160a01b038135169060200135613648565b34801561120e57600080fd5b506104e4600480360360c081101561122557600080fd5b506001600160a01b03813581169160208101359091169060408101359060608101359060808101359060a001351515613722565b34801561126557600080fd5b506104e46004803603602081101561127c57600080fd5b50356001600160a01b0316613794565b34801561129857600080fd5b506109ef600480360360408110156112af57600080fd5b506001600160a01b0381358116916020013516613897565b3480156112d357600080fd5b50610519600480360360608110156112ea57600080fd5b506001600160a01b0381358116916020810135909116906040013561394a565b34801561131657600080fd5b506105ec6004803603602081101561132d57600080fd5b50356001600160a01b03166139e4565b34801561134957600080fd5b506105ec6004803603602081101561136057600080fd5b50356001600160a01b0316613aa2565b34801561137c57600080fd5b506104e46004803603602081101561139357600080fd5b50356001600160a01b0316613ac3565b3480156113af57600080fd5b506104e4600480360360408110156113c657600080fd5b506001600160a01b0381351690602001351515613bfd565b3480156113ea57600080fd5b506104e46004803603602081101561140157600080fd5b50356001600160a01b0316613d53565b34801561141d57600080fd5b506105ec6004803603602081101561143457600080fd5b50356001600160a01b0316613e50565b34801561145057600080fd5b5061149d600480360360c081101561146757600080fd5b5080356001600160a01b039081169160208101359091169060408101359060608101359060808101359060a0013560ff16613e6e565b604051808360028111156114ad57fe5b60ff1681526020018281526020019250505060405180910390f35b3480156114d457600080fd5b506105ec600480360360408110156114eb57600080fd5b506001600160a01b0381358116916020013516613f03565b34801561150f57600080fd5b506104e46004803603606081101561152657600080fd5b506001600160a01b038135811691602081013590911690604001351515613f32565b34801561155457600080fd5b506104e46004803603606081101561156b57600080fd5b506001600160a01b03813581169160208101359091169060400135613fc9565b34801561159757600080fd5b506105ec600480360360408110156115ae57600080fd5b506001600160a01b03813581169160200135166140f9565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47081158015906115fa5750808214155b925050505b919050565b6001600160a01b03166000908152603660205260409020600d0154600160e01b900460ff1690565b6060603880548060200260200160405190810160405280929190818152602001828054801561168457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611666575b5050505050905090565b6034546001600160a01b031633146116da57604051600160e51b62461bcd02815260040180806020018281038252602a815260200180615b59602a913960400191505060405180910390fd5b6116e48482614128565b6001600160a01b03841660009081526036602052604090206117059061421e565b600061172061171386613aa2565b859063ffffffff6142b616565b6001600160a01b038616600090815260366020526040902090915061174c90828563ffffffff61431316565b61175885846000614371565b5050505050565b600481565b6001600160a01b038083166000908152603660209081526040808320600c015481517f70a08231000000000000000000000000000000000000000000000000000000008152868616600482015291519394169283926370a082319260248082019391829003018186803b1580156117da57600080fd5b505afa1580156117ee573d6000803e3d6000fd5b505050506040513d602081101561180457600080fd5b50519150505b92915050565b6001600160a01b03166000908152603660205260409020600d0154600160d01b900460ff1690565b6001600160a01b03808216600090815260376020908152604080832093861683529290529081208054151561187157600091505061180a565b6000816003015411611884576002611887565b60015b949350505050565b60355460408051600160e01b6385c858b1028152905133926001600160a01b0316916385c858b1916004808301926020929190829003018186803b1580156118d657600080fd5b505afa1580156118ea573d6000803e3d6000fd5b505050506040513d602081101561190057600080fd5b50516001600160a01b03161461194a57604051600160e51b62461bcd028152600401808060200182810382526037815260200180615aaf6037913960400191505060405180910390fd5b6001600160a01b039182166000908152603660205260409020600d0180546001600160a01b03191691909216179055565b6034546001600160a01b031633146119c757604051600160e51b62461bcd02815260040180806020018281038252602a815260200180615b59602a913960400191505060405180910390fd5b6119cf614521565b6001600160a01b03848116911614611a41573415611a2157604051600160e51b62461bcd0281526004018080602001828103825260328152602001806159ee6032913960400191505060405180910390fd5b611a3c6001600160a01b03841683308463ffffffff61453916565b611b50565b34811115611a8357604051600160e51b62461bcd0281526004018080602001828103825260358152602001806159236035913960400191505060405180910390fd5b80341115611b50576000611a9d348363ffffffff6145c116565b6040519091506000906001600160a01b0385169061c35090849084818181858888f193505050503d8060008114611af0576040519150601f19603f3d011682016040523d82523d6000602084013e611af5565b606091505b505090508015156117585760408051600160e51b62461bcd02815260206004820152601660248201527f5472616e73666572206f6620455448206661696c656400000000000000000000604482015290519081900360640190fd5b505050565b60355460408051600160e01b6385c858b1028152905133926001600160a01b0316916385c858b1916004808301926020929190829003018186803b158015611b9c57600080fd5b505afa158015611bb0573d6000803e3d6000fd5b505050506040513d6020811015611bc657600080fd5b50516001600160a01b031614611c1057604051600160e51b62461bcd028152600401808060200182810382526037815260200180615aaf6037913960400191505060405180910390fd5b6001600160a01b03909116600090815260366020526040902060090155565b6001600160a01b039081166000908152603660205260409020600c01541690565b60345460009081906001600160a01b03163314611ca157604051600160e51b62461bcd02815260040180806020018281038252602a815260200180615b59602a913960400191505060405180910390fd5b600080611cae89896128d2565b9250509150611cc1898984848b8a614603565b611ccf898989848a8a614632565b611cdb89600089614371565b611ce5898961475a565b999098509650505050505050565b60355460408051600160e01b6385c858b1028152905133926001600160a01b0316916385c858b1916004808301926020929190829003018186803b158015611d3a57600080fd5b505afa158015611d4e573d6000803e3d6000fd5b505050506040513d6020811015611d6457600080fd5b50516001600160a01b031614611dae57604051600160e51b62461bcd028152600401808060200182810382526037815260200180615aaf6037913960400191505060405180910390fd5b6001600160a01b03166000908152603660205260409020600d0180547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff169055565b60355460408051600160e01b6385c858b1028152905133926001600160a01b0316916385c858b1916004808301926020929190829003018186803b158015611e3757600080fd5b505afa158015611e4b573d6000803e3d6000fd5b505050506040513d6020811015611e6157600080fd5b50516001600160a01b031614611eab57604051600160e51b62461bcd028152600401808060200182810382526037815260200180615aaf6037913960400191505060405180910390fd5b6001600160a01b0380851660009081526036602052604080822081517f13769cd4000000000000000000000000000000000000000000000000000000008152600481019190915286841660248201526044810186905292841660648401525173__CoreLibrary___________________________926313769cd4926084808301939192829003018186803b158015611f4257600080fd5b505af4158015611f56573d6000803e3d6000fd5b50505050611f63846147ec565b50505050565b6001600160a01b031660009081526036602052604090206006015490565b6001600160a01b03166000908152603660205260409020600d0154600160e81b900460ff1690565b6001600160a01b03166000908152603660205260409020600d0154600160a01b900464ffffffffff1690565b6038805482908110611fe957fe5b6000918252602090912001546001600160a01b0316905081565b6001600160a01b03166000908152603660205260409020600d0154790100000000000000000000000000000000000000000000000000900460ff1690565b6001600160a01b03166000908152603660205260409020600b81015460088201546009830154600d909301549193909291600160d01b900460ff1690565b60355460408051600160e01b6385c858b1028152905133926001600160a01b0316916385c858b1916004808301926020929190829003018186803b1580156120c657600080fd5b505afa1580156120da573d6000803e3d6000fd5b505050506040513d60208110156120f057600080fd5b50516001600160a01b03161461213a57604051600160e51b62461bcd028152600401808060200182810382526037815260200180615aaf6037913960400191505060405180910390fd5b612142614894565b565b60355460408051600160e01b6385c858b1028152905133926001600160a01b0316916385c858b1916004808301926020929190829003018186803b15801561218b57600080fd5b505afa15801561219f573d6000803e3d6000fd5b505050506040513d60208110156121b557600080fd5b50516001600160a01b0316146121ff57604051600160e51b62461bcd028152600401808060200182810382526037815260200180615aaf6037913960400191505060405180910390fd5b6001600160a01b039091166000908152603660205260409020600b0155565b6001600160a01b0390811660009081526037602090815260408083209490931682529290925290206004015464ffffffffff1690565b6034546001600160a01b031633146122a057604051600160e51b62461bcd02815260040180806020018281038252602a815260200180615b59602a913960400191505060405180910390fd5b6122ac8988888561492f565b6122b5886149e8565b6122c28988888786614a0c565b6122ce89876000614371565b8015156122f0576122f08860006122eb888763ffffffff6142b616565b614371565b505050505050505050565b6001600160a01b039081166000908152603660205260409020600d01541690565b6001600160a01b0380821660009081526037602090815260408083209386168352929052206003015492915050565b60355460408051600160e01b6385c858b1028152905133926001600160a01b0316916385c858b1916004808301926020929190829003018186803b15801561239257600080fd5b505afa1580156123a6573d6000803e3d6000fd5b505050506040513d60208110156123bc57600080fd5b50516001600160a01b03161461240657604051600160e51b62461bcd028152600401808060200182810382526037815260200180615aaf6037913960400191505060405180910390fd5b6001600160a01b039091166000908152603660205260409020600a0155565b60355460408051600160e01b6385c858b1028152905133926001600160a01b0316916385c858b1916004808301926020929190829003018186803b15801561246c57600080fd5b505afa158015612480573d6000803e3d6000fd5b505050506040513d602081101561249657600080fd5b50516001600160a01b0316146124e057604051600160e51b62461bcd028152600401808060200182810382526037815260200180615aaf6037913960400191505060405180910390fd5b6001600160a01b03166000908152603660205260409020600d0180547fffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff16600160e81b179055565b6001600160a01b031660009081526036602052604090206002015490565b6001600160a01b03808216600090815260366020908152604080832060355482517f3618abba00000000000000000000000000000000000000000000000000000000815292519495919486949190921692633618abba9260048083019392829003018186803b1580156125b857600080fd5b505afa1580156125cc573d6000803e3d6000fd5b505050506040513d60208110156125e257600080fd5b50516005830154909150151561267e57806001600160a01b031663bb85c0bb856040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561264857600080fd5b505afa15801561265c573d6000803e3d6000fd5b505050506040513d602081101561267257600080fd5b505192506115ff915050565b506005015492915050565b6034546001600160a01b031633146126d557604051600160e51b62461bcd02815260040180806020018281038252602a815260200180615b59602a913960400191505060405180910390fd5b80341561271657604051600160e51b62461bcd02815260040180806020018281038252603681526020018061589d6036913960400191505060405180910390fd5b61271e614521565b6001600160a01b0385811691161461274f5761274a6001600160a01b038516828563ffffffff614abe16565b611f63565b6040516000906001600160a01b0383169061c35090869084818181858888f193505050503d8060008114611af0576040519150601f19603f3d011682016040523d82523d6000602084013e611af5565b6001600160a01b0381166000908152603660205260408120600481015415156128475780600d0160009054906101000a90046001600160a01b03166001600160a01b03166334762ca56040518163ffffffff1660e01b815260040160206040518083038186803b15801561281257600080fd5b505afa158015612826573d6000803e3d6000fd5b505050506040513d602081101561283c57600080fd5b505191506115ff9050565b6004015492915050565b6001600160a01b031660009081526036602052604090206003015490565b6001600160a01b0390811660009081526037602090815260408083209490931682529290925290206004015465010000000000900460ff1690565b6001600160a01b03166000908152603660205260409020600d0154600160d81b900460ff1690565b6001600160a01b03808216600090815260376020908152604080832093861683529290529081208054829182911515612915575060009250829150819050612958565b80546001600160a01b038716600090815260366020526040812061293a908490614b3e565b9050818161294e818363ffffffff6145c116565b9550955095505050505b9250925092565b6001600160a01b03166000908152603660205260409020600b015490565b60355460408051600160e01b6385c858b1028152905133926001600160a01b0316916385c858b1916004808301926020929190829003018186803b1580156129c457600080fd5b505afa1580156129d8573d6000803e3d6000fd5b505050506040513d60208110156129ee57600080fd5b50516001600160a01b031614612a3857604051600160e51b62461bcd028152600401808060200182810382526037815260200180615aaf6037913960400191505060405180910390fd5b6001600160a01b03841660009081526036602052604080822081517f24ddc4e60000000000000000000000000000000000000000000000000000000081526004810191909152602481018690526044810185905260648101849052905173__CoreLibrary___________________________926324ddc4e69260848082019391829003018186803b158015612acc57600080fd5b505af4158015612ae0573d6000803e3d6000fd5b5050505050505050565b60355460408051600160e01b6385c858b1028152905133926001600160a01b0316916385c858b1916004808301926020929190829003018186803b158015612b3157600080fd5b505afa158015612b45573d6000803e3d6000fd5b505050506040513d6020811015612b5b57600080fd5b50516001600160a01b031614612ba557604051600160e51b62461bcd028152600401808060200182810382526037815260200180615aaf6037913960400191505060405180910390fd5b6001600160a01b03811660009081526036602052604080822081517fe5df56a60000000000000000000000000000000000000000000000000000000081526004810191909152905173__CoreLibrary___________________________9263e5df56a69260248082019391829003018186803b158015612c2457600080fd5b505af4158015611758573d6000803e3d6000fd5b6034546000906001600160a01b03163314612c8757604051600160e51b62461bcd02815260040180806020018281038252602a815260200180615b59602a913960400191505060405180910390fd5b612c92848484614c32565b612c9d848484614c82565b612ca984600080614371565b506001600160a01b038083166000908152603760209081526040808320938716835292905220600301545b9392505050565b6034546001600160a01b03163314612d2757604051600160e51b62461bcd02815260040180806020018281038252602a815260200180615b59602a913960400191505060405180910390fd5b6001600160a01b0384166000908152603660205260409020612d489061421e565b612d5484600084614371565b8015611f6357611f6384846000613f32565b6001600160a01b031660009081526036602052604090206007015490565b60355460408051600160e01b6385c858b1028152905133926001600160a01b0316916385c858b1916004808301926020929190829003018186803b158015612dcb57600080fd5b505afa158015612ddf573d6000803e3d6000fd5b505050506040513d6020811015612df557600080fd5b50516001600160a01b031614612e3f57604051600160e51b62461bcd028152600401808060200182810382526037815260200180615aaf6037913960400191505060405180910390fd5b6001600160a01b0381166000908152603660205260409020805415801590612e6b575060008160070154115b1515612eab57604051600160e51b62461bcd028152600401808060200182810382526024815260200180615ae66024913960400191505060405180910390fd5b600d0180547fffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffff16600160e01b17905550565b60355460408051600160e01b6385c858b1028152905133926001600160a01b0316916385c858b1916004808301926020929190829003018186803b158015612f2457600080fd5b505afa158015612f38573d6000803e3d6000fd5b505050506040513d6020811015612f4e57600080fd5b50516001600160a01b031614612f9857604051600160e51b62461bcd028152600401808060200182810382526037815260200180615aaf6037913960400191505060405180910390fd5b6001600160a01b03166000908152603660205260409020600d0180547fffffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffff169055565b6034546001600160a01b0316331461302657604051600160e51b62461bcd02815260040180806020018281038252602a815260200180615b59602a913960400191505060405180910390fd5b6001600160a01b03841660009081526036602052604090206130479061421e565b61305384836000614371565b8015611f6357611f6384846001613f32565b6001600160a01b031660009081526036602052604090205490565b6001600160a01b0381166000908152603660205260408120816130a282614ceb565b90508015156130b6576000925050506115ff565b60006130c1856139e4565b90506130e36130d6828463ffffffff6142b616565b839063ffffffff614d0816565b95945050505050565b6001600160a01b0381166000908152603660205260408120612cd461311082614ceb565b613119856139e4565b9063ffffffff6142b616565b600061312f614d44565b60015490915060ff16806131465750613146614d49565b80613152575060005481115b151561319257604051600160e51b62461bcd02815260040180806020018281038252602e815260200180615b2b602e913960400191505060405180910390fd5b60015460ff161580156131b1576001805460ff19168117905560008290555b603580546001600160a01b0319166001600160a01b0385161790556131d4614894565b8015611b50576001805460ff19169055505050565b6001600160a01b031660009081526036602052604090206001015490565b6035546001600160a01b031681565b6001600160a01b03166000908152603660205260409020600a015490565b6034546001600160a01b0316331461328057604051600160e51b62461bcd02815260040180806020018281038252602a815260200180615b59602a913960400191505060405180910390fd5b80613289614521565b6001600160a01b038681169116146132fb5734156132db57604051600160e51b62461bcd02815260040180806020018281038252605f815260200180615a20605f913960600191505060405180910390fd5b6132f66001600160a01b03861685838663ffffffff61453916565b611758565b3483111561333d57604051600160e51b62461bcd0281526004018080602001828103825260358152602001806159236035913960400191505060405180910390fd5b6040516000906001600160a01b0383169061c35090869084818181858888f193505050503d806000811461338d576040519150601f19603f3d011682016040523d82523d6000602084013e613392565b606091505b505090508015156133ed5760408051600160e51b62461bcd02815260206004820152601660248201527f5472616e73666572206f6620455448206661696c656400000000000000000000604482015290519081900360640190fd5b505050505050565b60355460408051600160e01b6385c858b1028152905133926001600160a01b0316916385c858b1916004808301926020929190829003018186803b15801561343c57600080fd5b505afa158015613450573d6000803e3d6000fd5b505050506040513d602081101561346657600080fd5b50516001600160a01b0316146134b057604051600160e51b62461bcd028152600401808060200182810382526037815260200180615aaf6037913960400191505060405180910390fd5b603880546000919060001981019081106134c657fe5b6000918252602090912001546001600160a01b0390811691508216811461352157604051600160e51b62461bcd02815260040180806020018281038252603d8152602001806159b1603d913960400191505060405180910390fd5b61352a81613aa2565b1561356957604051600160e51b62461bcd028152600401808060200182810382526030815260200180615a7f6030913960400191505060405180910390fd5b6001600160a01b0381166000908152603660205260408120600d81018054600c830180546001600160a01b0319169055600b8301849055838355600783018490556008830184905560098301849055600a909201929092557fffffff00ff0000ffffffffff0000000000000000000000000000000000000000169055603880548015156135f257fe5b600082815260209020810160001990810180546001600160a01b03191690550190555050565b6001600160a01b0381166000908152603660205260408120612cd481614d4f565b6034546001600160a01b031681565b60355460408051600160e01b6385c858b1028152905133926001600160a01b0316916385c858b1916004808301926020929190829003018186803b15801561368f57600080fd5b505afa1580156136a3573d6000803e3d6000fd5b505050506040513d60208110156136b957600080fd5b50516001600160a01b03161461370357604051600160e51b62461bcd028152600401808060200182810382526037815260200180615aaf6037913960400191505060405180910390fd5b6001600160a01b03909116600090815260366020526040902060080155565b6034546001600160a01b0316331461376e57604051600160e51b62461bcd02815260040180806020018281038252602a815260200180615b59602a913960400191505060405180910390fd5b61377a86868685614d7d565b613788868686868686614de1565b6133ed86856000614371565b60355460408051600160e01b6385c858b1028152905133926001600160a01b0316916385c858b1916004808301926020929190829003018186803b1580156137db57600080fd5b505afa1580156137ef573d6000803e3d6000fd5b505050506040513d602081101561380557600080fd5b50516001600160a01b03161461384f57604051600160e51b62461bcd028152600401808060200182810382526037815260200180615aaf6037913960400191505060405180910390fd5b6001600160a01b03166000908152603660205260409020600d0180547fffffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffff16600160d81b179055565b6001600160a01b0380831660008181526036602090815260408083209486168352603782528083209383529290529081209091829182918291826138db8989611764565b8254909150151561390c576004909101549095506000945084935060ff650100000000009091041691506139419050565b8061391d838563ffffffff614b3e16565b600284015460049094015491985096509194505065010000000000900460ff169150505b92959194509250565b6001600160a01b038381166000818152603660209081526040808320948716835260378252808320938352929052908120600d830154919291600160d81b900460ff16151561399e57600092505050612cd4565b600481015465010000000000900460ff1615806139c75750600d820154600160d01b900460ff16155b806139da57506139d78686611764565b84115b9695505050505050565b6000806139ef614521565b6001600160a01b0316836001600160a01b03161415613a105750303161180a565b604080517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290516001600160a01b038516916370a08231916024808301926020929190829003018186803b158015613a6f57600080fd5b505afa158015613a83573d6000803e3d6000fd5b505050506040513d6020811015613a9957600080fd5b50519392505050565b6001600160a01b038116600090815260366020526040812061180a90614ceb565b60355460408051600160e01b6385c858b1028152905133926001600160a01b0316916385c858b1916004808301926020929190829003018186803b158015613b0a57600080fd5b505afa158015613b1e573d6000803e3d6000fd5b505050506040513d6020811015613b3457600080fd5b50516001600160a01b031614613b7e57604051600160e51b62461bcd028152600401808060200182810382526037815260200180615aaf6037913960400191505060405180910390fd5b6001600160a01b03811660009081526036602052604080822081517f83c165a00000000000000000000000000000000000000000000000000000000081526004810191909152905173__CoreLibrary___________________________926383c165a09260248082019391829003018186803b158015612c2457600080fd5b60355460408051600160e01b6385c858b1028152905133926001600160a01b0316916385c858b1916004808301926020929190829003018186803b158015613c4457600080fd5b505afa158015613c58573d6000803e3d6000fd5b505050506040513d6020811015613c6e57600080fd5b50516001600160a01b031614613cb857604051600160e51b62461bcd028152600401808060200182810382526037815260200180615aaf6037913960400191505060405180910390fd5b6001600160a01b03821660009081526036602052604080822081517ff63babbe00000000000000000000000000000000000000000000000000000000815260048101919091528315156024820152905173__CoreLibrary___________________________9263f63babbe9260448082019391829003018186803b158015613d3f57600080fd5b505af41580156133ed573d6000803e3d6000fd5b60355460408051600160e01b6385c858b1028152905133926001600160a01b0316916385c858b1916004808301926020929190829003018186803b158015613d9a57600080fd5b505afa158015613dae573d6000803e3d6000fd5b505050506040513d6020811015613dc457600080fd5b50516001600160a01b031614613e0e57604051600160e51b62461bcd028152600401808060200182810382526037815260200180615aaf6037913960400191505060405180910390fd5b6001600160a01b03166000908152603660205260409020600d0180547fffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffff169055565b6001600160a01b031660009081526036602052604090206009015490565b60345460009081906001600160a01b03163314613ebf57604051600160e51b62461bcd02815260040180806020018281038252602a815260200180615b59602a913960400191505060405180910390fd5b613ecc8888888887614e5e565b6000613eda89898787614f60565b9050613ee889600080614371565b80613ef38a8a61475a565b9250925050965096945050505050565b6001600160a01b0380821660009081526037602090815260408083209386168352929052206001015492915050565b6034546001600160a01b03163314613f7e57604051600160e51b62461bcd02815260040180806020018281038252602a815260200180615b59602a913960400191505060405180910390fd5b6001600160a01b0391821660009081526037602090815260408083209590941682529390935291206004018054911515650100000000000265ff000000000019909216919091179055565b6034546001600160a01b0316331461401557604051600160e51b62461bcd02815260040180806020018281038252602a815260200180615b59602a913960400191505060405180910390fd5b61401d614521565b6001600160a01b0384811691161461404957611a3c6001600160a01b038416838363ffffffff614abe16565b6040516000906001600160a01b0384169061c35090849084818181858888f193505050503d8060008114614099576040519150601f19603f3d011682016040523d82523d6000602084013e61409e565b606091505b50509050801515611f635760408051600160e51b62461bcd02815260206004820152601660248201527f5472616e73666572206f6620455448206661696c656400000000000000000000604482015290519081900360640190fd5b6001600160a01b0380821660009081526037602090815260408083209386168352929052206002015492915050565b603554604080517fee89129600000000000000000000000000000000000000000000000000000000815290516000926001600160a01b03169163ee891296916004808301926020929190829003018186803b15801561418657600080fd5b505afa15801561419a573d6000803e3d6000fd5b505050506040513d60208110156141b057600080fd5b505190506141bc614521565b6001600160a01b038481169116146141e857611a3c6001600160a01b038416828463ffffffff614abe16565b6040516001600160a01b0382169083156108fc029084906000818181858888f19350505050158015611f63573d6000803e3d6000fd5b600061422982614ceb565b905080156142b2576001820154600d83015460009161425591600160a01b900464ffffffffff16615061565b835490915061426b90829063ffffffff6150ac16565b83556004830154600d84015460009161429191600160a01b900464ffffffffff166150e4565b90506142aa8460070154826150ac90919063ffffffff16565b600785015550505b5050565b600082820183811015612cd45760408051600160e51b62461bcd02815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b600061433661432184615133565b61432a84615133565b9063ffffffff614d0816565b90506000614352614345615149565b839063ffffffff6142b616565b855490915061436890829063ffffffff6150ac16565b90945550505050565b6001600160a01b038084166000908152603660205260408120600d810154909282918291166357e37af0886143b9886143ad8b613119856139e4565b9063ffffffff6145c116565b8760020154886003015489600601546040518663ffffffff1660e01b815260040180866001600160a01b03166001600160a01b031681526020018581526020018481526020018381526020018281526020019550505050505060606040518083038186803b15801561442a57600080fd5b505afa15801561443e573d6000803e3d6000fd5b505050506040513d606081101561445457600080fd5b508051602080830151604093840151600189018490556005890182905560048901819055600d8901805464ffffffffff4216600160a01b027fffffffffffffff0000000000ffffffffffffffffffffffffffffffffffffffff909116179055885460078a0154865186815294850184905284870183905260608501919091526080840152935192965094509192506001600160a01b038916917f04e4f521f16fcfd987978b05504262c2a2db223844977aab000e5accedb2d725919081900360a00190a250505050505050565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee90565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052611f63908590615159565b6000612cd483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250615320565b6001600160a01b03861660009081526036602052604090206146249061421e565b6133ed8686868686866153ba565b6001600160a01b038087166000818152603660209081526040808320948a1683526037825280832093835292905220600183600281111561466f57fe5b141561468b576005820154600382015560006001820155614705565b600283600281111561469957fe5b14156146b5576000600382015560078201546001820155614705565b60408051600160e51b62461bcd02815260206004820152601860248201527f496e76616c696420626f72726f772072617465206d6f64650000000000000000604482015290519081900360640190fd5b805461471d908690613119908963ffffffff6142b616565b81556002810154614734908563ffffffff6142b616565b6002820155600401805464ffffffffff19164264ffffffffff1617905550505050505050565b6000806147678484611838565b9050600081600281111561477757fe5b141561478757600091505061180a565b600181600281111561479557fe5b146147bb576001600160a01b038416600090815260366020526040902060040154611887565b50506001600160a01b0390811660009081526037602090815260408083209490931682529290925290206003015490565b6000805b60385481101561483b57826001600160a01b031660388281548110151561481357fe5b6000918252602090912001546001600160a01b0316141561483357600191505b6001016147f0565b508015156142b257603880546001810182556000919091527f38395c5dceade9603479b177b68959049485df8aa97b39f3533039af5f4561990180546001600160a01b0384166001600160a01b03199091161790555050565b603560009054906101000a90046001600160a01b03166001600160a01b0316630261bf8b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156148e257600080fd5b505afa1580156148f6573d6000803e3d6000fd5b505050506040513d602081101561490c57600080fd5b5051603480546001600160a01b0319166001600160a01b03909216919091179055565b6001600160a01b038085166000818152603660209081526040808320948816835260378252808320938352929052206149678261421e565b60006149738787611838565b905060015b81600281111561498457fe5b14156149bf5760038201546149a2908490869063ffffffff61551916565b60038201546149ba908490879063ffffffff61558916565b6149df565b6149cf838563ffffffff61569c16565b6149df838663ffffffff6156bd16565b50505050505050565b6001600160a01b0381166000908152603660205260409020614a099061421e565b50565b6001600160a01b038085166000908152603760209081526040808320938916835292815282822060369091529190208154614a539086906143ad908663ffffffff6142b616565b82556002614a618888611838565b6002811115614a6c57fe5b1415614a7d57600781015460018301555b8315614a9e576002820154614a98908563ffffffff6145c116565b60028301555b50600401805464ffffffffff19164264ffffffffff161790555050505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052611b50908490615159565b81546000901515614b515750600061180a565b6000614b608460000154615133565b6003850154909150600090819015614b945760038601546004870154614b8d919064ffffffffff166150e4565b9050614bd6565b614bd3866001015461432a8760070154614bc789600401548a600d0160149054906101000a900464ffffffffff166150e4565b9063ffffffff6150ac16565b90505b614bee614be9848363ffffffff6150ac16565b615718565b8654909250821415614c2957600486015464ffffffffff164214614c29578554614c1f90600163ffffffff6142b616565b935050505061180a565b50949350505050565b6001600160a01b03808416600081815260366020908152604080832094871683526037825280832093835292905220614c6a8261421e565b6003810154611758908390859063ffffffff61551916565b6001600160a01b038083166000908152603760209081526040808320938716835292815282822060369091529190208154614cc3908463ffffffff6142b616565b8255600501546003820155600401805464ffffffffff19164264ffffffffff16179055505050565b600061180a826003015483600201546142b690919063ffffffff16565b60006002820461188783614d38614d2b876b033b2e3c9fd0803ce8000000615731565b849063ffffffff6142b616565b9063ffffffff61579116565b600490565b303b1590565b600080612cd48360000154614bc7856001015486600d0160149054906101000a900464ffffffffff16615061565b6001600160a01b03808516600090815260366020908152604080832060378352818420948816845293909152812090614db68787611838565b6001600160a01b0388166000908152603660205260409020909150614dda9061421e565b6001614978565b6001600160a01b038087166000818152603660209081526040808320948a16835260378252808320938352929052208054614e289087906143ad908763ffffffff6142b616565b8155600782015460018201558215614e495760006003820181905560018201555b6002810154614734908663ffffffff6145c116565b6001600160a01b03808616600081815260366020908152604080832094891683526037825280832093835292905220614e968261421e565b6001836002811115614ea457fe5b1415614ed6576003810154614ec083878363ffffffff61558916565b614ed0838663ffffffff61569c16565b506149df565b6002836002811115614ee457fe5b1415614f10576005820154614eff838763ffffffff6156bd16565b614ed083868363ffffffff61551916565b60408051600160e51b62461bcd02815260206004820152601a60248201527f496e76616c69642072617465206d6f6465207265636569766564000000000000604482015290519081900360640190fd5b6001600160a01b03808416600090815260376020908152604080832093881683529281528282206036909152918120909190826002856002811115614fa157fe5b1415614fc25750600581015460038301556000600183810191909155615029565b6001856002811115614fd057fe5b1415614fef575060006003830155600781015460018301556002615029565b604051600160e51b62461bcd028152600401808060200182810382526023815260200180615b836023913960400191505060405180910390fd5b825461503b908763ffffffff6142b616565b83556004909201805464ffffffffff19164264ffffffffff161790555095945050505050565b60008061507b4264ffffffffff851663ffffffff6145c116565b9050600061508f6143216301e13380615133565b90506130e361509c615149565b613119878463ffffffff6150ac16565b6000612cd46b033b2e3c9fd0803ce8000000614d386150d1868663ffffffff61573116565b6b019d971e4fe8401e74000000906142b6565b6000806150fe4264ffffffffff851663ffffffff6145c116565b90506000615116856301e1338063ffffffff61579116565b90506130e382615127614d2b615149565b9063ffffffff6157d316565b600061180a82633b9aca0063ffffffff61573116565b6b033b2e3c9fd0803ce800000090565b61516b826001600160a01b03166115c6565b15156151c15760408051600160e51b62461bcd02815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b602083106151ff5780518252601f1990920191602091820191016151e0565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114615261576040519150601f19603f3d011682016040523d82523d6000602084013e615266565b606091505b50915091508115156152c25760408051600160e51b62461bcd02815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115611f63578080602001905160208110156152de57600080fd5b50511515611f6357604051600160e51b62461bcd02815260040180806020018281038252602a815260200180615ba6602a913960400191505060405180910390fd5b600081848411156153b257604051600160e51b62461bcd0281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561537757818101518382015260200161535f565b50505050905090810190601f1680156153a45780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b60006153c68787611838565b6001600160a01b038816600090815260366020526040902090915060018260028111156153ef57fe5b1415615438576001600160a01b038088166000908152603760209081526040808320938c168352929052206003810154615432908390899063ffffffff61558916565b5061545c565b600282600281111561544657fe5b141561545c5761545c818763ffffffff6156bd16565b600061547285613119898963ffffffff6142b616565b9050600184600281111561548257fe5b14156154a55760058201546154a0908390839063ffffffff61551916565b6122f0565b60028460028111156154b357fe5b14156154c9576154a0828263ffffffff61569c16565b60408051600160e51b62461bcd02815260206004820152601c60248201527f496e76616c6964206e657720626f72726f772072617465206d6f646500000000604482015290519081900360640190fd5b600283015461552e818463ffffffff6142b616565b6002850155600061554283614bc786615133565b905060006155578660060154614bc785615133565b90506155796155698760020154615133565b61432a848463ffffffff6142b616565b8660060181905550505050505050565b60028301548211156155e55760408051600160e51b62461bcd02815260206004820152601a60248201527f496e76616c696420616d6f756e7420746f206465637265617365000000000000604482015290519081900360640190fd5b60028301546155fa818463ffffffff6145c116565b600285018190551515615614575060006006840155611b50565b600061562383614bc786615133565b905060006156388660060154614bc785615133565b90508181101561567c57604051600160e51b62461bcd02815260040180806020018281038252602381526020018061598e6023913960400191505060405180910390fd5b61557961568c8760020154615133565b61432a838563ffffffff6145c116565b60038201546156b1908263ffffffff6142b616565b82600301819055505050565b600382015481111561570357604051600160e51b62461bcd0281526004018080602001828103825260508152602001806158d36050913960600191505060405180910390fd5b60038201546156b1908263ffffffff6145c116565b6000631dcd6500612cd4633b9aca00614d3883866142b6565b60008215156157425750600061180a565b82820282848281151561575157fe5b0414612cd457604051600160e51b62461bcd028152600401808060200182810382526021815260200180615b0a6021913960400191505060405180910390fd5b6000612cd483836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f00000000000081525061582f565b60006002820615156157f1576b033b2e3c9fd0803ce80000006157f3565b825b90506002820491505b811561180a5761580c83846150ac565b925060028206156158245761582181846150ac565b90505b6002820491506157fc565b60008183151561588457604051600160e51b62461bcd0281526004018080602001828103825283818151815260200191508051906020019080838360008381101561537757818101518382015260200161535f565b506000838581151561589257fe5b049594505050505056fe466565206c69717569646174696f6e20646f6573206e6f74207265717569726520616e79207472616e73666572206f662076616c756554686520616d6f756e742074686174206973206265696e6720737562747261637465642066726f6d20746865207661726961626c6520746f74616c20626f72726f777320697320696e636f727265637454686520616d6f756e7420616e64207468652076616c75652073656e7420746f206465706f73697420646f206e6f74206d617463684f6e6c7920636f6e7472616374732063616e2073656e6420657468657220746f20746865204c656e64696e6720706f6f6c20636f726554686520616d6f756e747320746f20737562747261637420646f6e2774206d6174636852657365727665206265696e672072656d6f76656420697320646966666572656e74207468616e20746865207265736572766520726571756573746564557365722069732073656e64696e672045544820616c6f6e67207769746820746865204552433230207472616e736665722e557365722069732073656e64696e672045544820616c6f6e67207769746820746865204552433230207472616e736665722e20436865636b207468652076616c756520617474726962757465206f6620746865207472616e73616374696f6e43616e6e6f742072656d6f7665206120726573657276652077697468206c6971756964697479206465706f73697465645468652063616c6c6572206d7573742062652061206c656e64696e6720706f6f6c20636f6e666967757261746f7220636f6e74726163745265736572766520686173206e6f74206265656e20696e697469616c697a656420796574536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a65645468652063616c6c6572206d7573742062652061206c656e64696e6720706f6f6c20636f6e7472616374496e76616c696420696e7465726573742072617465206d6f64652072656365697665645361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a165627a7a72305820a1b8d940061c7adb8ddc7746859b130da74c2fc47d21976ed23ed9ef954de0220029"

// DeployLendingPoolCore deploys a new Ethereum contract, binding an instance of LendingPoolCore to it.
func DeployLendingPoolCore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LendingPoolCore, error) {
	parsed, err := abi.JSON(strings.NewReader(LendingPoolCoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LendingPoolCoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LendingPoolCore{LendingPoolCoreCaller: LendingPoolCoreCaller{contract: contract}, LendingPoolCoreTransactor: LendingPoolCoreTransactor{contract: contract}, LendingPoolCoreFilterer: LendingPoolCoreFilterer{contract: contract}}, nil
}

// LendingPoolCore is an auto generated Go binding around an Ethereum contract.
type LendingPoolCore struct {
	LendingPoolCoreCaller     // Read-only binding to the contract
	LendingPoolCoreTransactor // Write-only binding to the contract
	LendingPoolCoreFilterer   // Log filterer for contract events
}

// LendingPoolCoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type LendingPoolCoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolCoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LendingPoolCoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolCoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LendingPoolCoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolCoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LendingPoolCoreSession struct {
	Contract     *LendingPoolCore  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LendingPoolCoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LendingPoolCoreCallerSession struct {
	Contract *LendingPoolCoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// LendingPoolCoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LendingPoolCoreTransactorSession struct {
	Contract     *LendingPoolCoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// LendingPoolCoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type LendingPoolCoreRaw struct {
	Contract *LendingPoolCore // Generic contract binding to access the raw methods on
}

// LendingPoolCoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LendingPoolCoreCallerRaw struct {
	Contract *LendingPoolCoreCaller // Generic read-only contract binding to access the raw methods on
}

// LendingPoolCoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LendingPoolCoreTransactorRaw struct {
	Contract *LendingPoolCoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLendingPoolCore creates a new instance of LendingPoolCore, bound to a specific deployed contract.
func NewLendingPoolCore(address common.Address, backend bind.ContractBackend) (*LendingPoolCore, error) {
	contract, err := bindLendingPoolCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LendingPoolCore{LendingPoolCoreCaller: LendingPoolCoreCaller{contract: contract}, LendingPoolCoreTransactor: LendingPoolCoreTransactor{contract: contract}, LendingPoolCoreFilterer: LendingPoolCoreFilterer{contract: contract}}, nil
}

// NewLendingPoolCoreCaller creates a new read-only instance of LendingPoolCore, bound to a specific deployed contract.
func NewLendingPoolCoreCaller(address common.Address, caller bind.ContractCaller) (*LendingPoolCoreCaller, error) {
	contract, err := bindLendingPoolCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LendingPoolCoreCaller{contract: contract}, nil
}

// NewLendingPoolCoreTransactor creates a new write-only instance of LendingPoolCore, bound to a specific deployed contract.
func NewLendingPoolCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*LendingPoolCoreTransactor, error) {
	contract, err := bindLendingPoolCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LendingPoolCoreTransactor{contract: contract}, nil
}

// NewLendingPoolCoreFilterer creates a new log filterer instance of LendingPoolCore, bound to a specific deployed contract.
func NewLendingPoolCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*LendingPoolCoreFilterer, error) {
	contract, err := bindLendingPoolCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LendingPoolCoreFilterer{contract: contract}, nil
}

// bindLendingPoolCore binds a generic wrapper to an already deployed contract.
func bindLendingPoolCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LendingPoolCoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingPoolCore *LendingPoolCoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LendingPoolCore.Contract.LendingPoolCoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingPoolCore *LendingPoolCoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.LendingPoolCoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingPoolCore *LendingPoolCoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.LendingPoolCoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingPoolCore *LendingPoolCoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LendingPoolCore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingPoolCore *LendingPoolCoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingPoolCore *LendingPoolCoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.contract.Transact(opts, method, params...)
}

// COREREVISION is a free data retrieval call binding the contract method 0x0c7de4e9.
//
// Solidity: function CORE_REVISION() view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) COREREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "CORE_REVISION")
	return *ret0, err
}

// COREREVISION is a free data retrieval call binding the contract method 0x0c7de4e9.
//
// Solidity: function CORE_REVISION() view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) COREREVISION() (*big.Int, error) {
	return _LendingPoolCore.Contract.COREREVISION(&_LendingPoolCore.CallOpts)
}

// COREREVISION is a free data retrieval call binding the contract method 0x0c7de4e9.
//
// Solidity: function CORE_REVISION() view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) COREREVISION() (*big.Int, error) {
	return _LendingPoolCore.Contract.COREREVISION(&_LendingPoolCore.CallOpts)
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_LendingPoolCore *LendingPoolCoreCaller) AddressesProvider(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "addressesProvider")
	return *ret0, err
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_LendingPoolCore *LendingPoolCoreSession) AddressesProvider() (common.Address, error) {
	return _LendingPoolCore.Contract.AddressesProvider(&_LendingPoolCore.CallOpts)
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_LendingPoolCore *LendingPoolCoreCallerSession) AddressesProvider() (common.Address, error) {
	return _LendingPoolCore.Contract.AddressesProvider(&_LendingPoolCore.CallOpts)
}

// GetReserveATokenAddress is a free data retrieval call binding the contract method 0x34b3beee.
//
// Solidity: function getReserveATokenAddress(address _reserve) view returns(address)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveATokenAddress(opts *bind.CallOpts, _reserve common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveATokenAddress", _reserve)
	return *ret0, err
}

// GetReserveATokenAddress is a free data retrieval call binding the contract method 0x34b3beee.
//
// Solidity: function getReserveATokenAddress(address _reserve) view returns(address)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveATokenAddress(_reserve common.Address) (common.Address, error) {
	return _LendingPoolCore.Contract.GetReserveATokenAddress(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveATokenAddress is a free data retrieval call binding the contract method 0x34b3beee.
//
// Solidity: function getReserveATokenAddress(address _reserve) view returns(address)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveATokenAddress(_reserve common.Address) (common.Address, error) {
	return _LendingPoolCore.Contract.GetReserveATokenAddress(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveAvailableLiquidity is a free data retrieval call binding the contract method 0xe2403019.
//
// Solidity: function getReserveAvailableLiquidity(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveAvailableLiquidity(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveAvailableLiquidity", _reserve)
	return *ret0, err
}

// GetReserveAvailableLiquidity is a free data retrieval call binding the contract method 0xe2403019.
//
// Solidity: function getReserveAvailableLiquidity(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveAvailableLiquidity(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveAvailableLiquidity(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveAvailableLiquidity is a free data retrieval call binding the contract method 0xe2403019.
//
// Solidity: function getReserveAvailableLiquidity(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveAvailableLiquidity(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveAvailableLiquidity(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveConfiguration is a free data retrieval call binding the contract method 0x5fc526ff.
//
// Solidity: function getReserveConfiguration(address _reserve) view returns(uint256, uint256, uint256, bool)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveConfiguration(opts *bind.CallOpts, _reserve common.Address) (*big.Int, *big.Int, *big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveConfiguration", _reserve)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetReserveConfiguration is a free data retrieval call binding the contract method 0x5fc526ff.
//
// Solidity: function getReserveConfiguration(address _reserve) view returns(uint256, uint256, uint256, bool)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveConfiguration(_reserve common.Address) (*big.Int, *big.Int, *big.Int, bool, error) {
	return _LendingPoolCore.Contract.GetReserveConfiguration(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveConfiguration is a free data retrieval call binding the contract method 0x5fc526ff.
//
// Solidity: function getReserveConfiguration(address _reserve) view returns(uint256, uint256, uint256, bool)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveConfiguration(_reserve common.Address) (*big.Int, *big.Int, *big.Int, bool, error) {
	return _LendingPoolCore.Contract.GetReserveConfiguration(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveCurrentAverageStableBorrowRate is a free data retrieval call binding the contract method 0x46bc0f28.
//
// Solidity: function getReserveCurrentAverageStableBorrowRate(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveCurrentAverageStableBorrowRate(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveCurrentAverageStableBorrowRate", _reserve)
	return *ret0, err
}

// GetReserveCurrentAverageStableBorrowRate is a free data retrieval call binding the contract method 0x46bc0f28.
//
// Solidity: function getReserveCurrentAverageStableBorrowRate(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveCurrentAverageStableBorrowRate(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveCurrentAverageStableBorrowRate(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveCurrentAverageStableBorrowRate is a free data retrieval call binding the contract method 0x46bc0f28.
//
// Solidity: function getReserveCurrentAverageStableBorrowRate(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveCurrentAverageStableBorrowRate(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveCurrentAverageStableBorrowRate(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveCurrentLiquidityRate is a free data retrieval call binding the contract method 0xc540148e.
//
// Solidity: function getReserveCurrentLiquidityRate(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveCurrentLiquidityRate(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveCurrentLiquidityRate", _reserve)
	return *ret0, err
}

// GetReserveCurrentLiquidityRate is a free data retrieval call binding the contract method 0xc540148e.
//
// Solidity: function getReserveCurrentLiquidityRate(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveCurrentLiquidityRate(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveCurrentLiquidityRate(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveCurrentLiquidityRate is a free data retrieval call binding the contract method 0xc540148e.
//
// Solidity: function getReserveCurrentLiquidityRate(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveCurrentLiquidityRate(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveCurrentLiquidityRate(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveCurrentStableBorrowRate is a free data retrieval call binding the contract method 0x88079d88.
//
// Solidity: function getReserveCurrentStableBorrowRate(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveCurrentStableBorrowRate(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveCurrentStableBorrowRate", _reserve)
	return *ret0, err
}

// GetReserveCurrentStableBorrowRate is a free data retrieval call binding the contract method 0x88079d88.
//
// Solidity: function getReserveCurrentStableBorrowRate(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveCurrentStableBorrowRate(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveCurrentStableBorrowRate(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveCurrentStableBorrowRate is a free data retrieval call binding the contract method 0x88079d88.
//
// Solidity: function getReserveCurrentStableBorrowRate(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveCurrentStableBorrowRate(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveCurrentStableBorrowRate(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveCurrentVariableBorrowRate is a free data retrieval call binding the contract method 0x906c0a41.
//
// Solidity: function getReserveCurrentVariableBorrowRate(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveCurrentVariableBorrowRate(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveCurrentVariableBorrowRate", _reserve)
	return *ret0, err
}

// GetReserveCurrentVariableBorrowRate is a free data retrieval call binding the contract method 0x906c0a41.
//
// Solidity: function getReserveCurrentVariableBorrowRate(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveCurrentVariableBorrowRate(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveCurrentVariableBorrowRate(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveCurrentVariableBorrowRate is a free data retrieval call binding the contract method 0x906c0a41.
//
// Solidity: function getReserveCurrentVariableBorrowRate(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveCurrentVariableBorrowRate(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveCurrentVariableBorrowRate(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveDecimals is a free data retrieval call binding the contract method 0xa2353fdc.
//
// Solidity: function getReserveDecimals(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveDecimals(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveDecimals", _reserve)
	return *ret0, err
}

// GetReserveDecimals is a free data retrieval call binding the contract method 0xa2353fdc.
//
// Solidity: function getReserveDecimals(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveDecimals(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveDecimals(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveDecimals is a free data retrieval call binding the contract method 0xa2353fdc.
//
// Solidity: function getReserveDecimals(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveDecimals(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveDecimals(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveInterestRateStrategyAddress is a free data retrieval call binding the contract method 0x6ae14416.
//
// Solidity: function getReserveInterestRateStrategyAddress(address _reserve) view returns(address)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveInterestRateStrategyAddress(opts *bind.CallOpts, _reserve common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveInterestRateStrategyAddress", _reserve)
	return *ret0, err
}

// GetReserveInterestRateStrategyAddress is a free data retrieval call binding the contract method 0x6ae14416.
//
// Solidity: function getReserveInterestRateStrategyAddress(address _reserve) view returns(address)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveInterestRateStrategyAddress(_reserve common.Address) (common.Address, error) {
	return _LendingPoolCore.Contract.GetReserveInterestRateStrategyAddress(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveInterestRateStrategyAddress is a free data retrieval call binding the contract method 0x6ae14416.
//
// Solidity: function getReserveInterestRateStrategyAddress(address _reserve) view returns(address)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveInterestRateStrategyAddress(_reserve common.Address) (common.Address, error) {
	return _LendingPoolCore.Contract.GetReserveInterestRateStrategyAddress(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveIsActive is a free data retrieval call binding the contract method 0x05075d6e.
//
// Solidity: function getReserveIsActive(address _reserve) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveIsActive(opts *bind.CallOpts, _reserve common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveIsActive", _reserve)
	return *ret0, err
}

// GetReserveIsActive is a free data retrieval call binding the contract method 0x05075d6e.
//
// Solidity: function getReserveIsActive(address _reserve) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveIsActive(_reserve common.Address) (bool, error) {
	return _LendingPoolCore.Contract.GetReserveIsActive(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveIsActive is a free data retrieval call binding the contract method 0x05075d6e.
//
// Solidity: function getReserveIsActive(address _reserve) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveIsActive(_reserve common.Address) (bool, error) {
	return _LendingPoolCore.Contract.GetReserveIsActive(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveIsFreezed is a free data retrieval call binding the contract method 0x4a08accb.
//
// Solidity: function getReserveIsFreezed(address _reserve) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveIsFreezed(opts *bind.CallOpts, _reserve common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveIsFreezed", _reserve)
	return *ret0, err
}

// GetReserveIsFreezed is a free data retrieval call binding the contract method 0x4a08accb.
//
// Solidity: function getReserveIsFreezed(address _reserve) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveIsFreezed(_reserve common.Address) (bool, error) {
	return _LendingPoolCore.Contract.GetReserveIsFreezed(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveIsFreezed is a free data retrieval call binding the contract method 0x4a08accb.
//
// Solidity: function getReserveIsFreezed(address _reserve) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveIsFreezed(_reserve common.Address) (bool, error) {
	return _LendingPoolCore.Contract.GetReserveIsFreezed(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveIsStableBorrowRateEnabled is a free data retrieval call binding the contract method 0x9e674848.
//
// Solidity: function getReserveIsStableBorrowRateEnabled(address _reserve) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveIsStableBorrowRateEnabled(opts *bind.CallOpts, _reserve common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveIsStableBorrowRateEnabled", _reserve)
	return *ret0, err
}

// GetReserveIsStableBorrowRateEnabled is a free data retrieval call binding the contract method 0x9e674848.
//
// Solidity: function getReserveIsStableBorrowRateEnabled(address _reserve) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveIsStableBorrowRateEnabled(_reserve common.Address) (bool, error) {
	return _LendingPoolCore.Contract.GetReserveIsStableBorrowRateEnabled(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveIsStableBorrowRateEnabled is a free data retrieval call binding the contract method 0x9e674848.
//
// Solidity: function getReserveIsStableBorrowRateEnabled(address _reserve) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveIsStableBorrowRateEnabled(_reserve common.Address) (bool, error) {
	return _LendingPoolCore.Contract.GetReserveIsStableBorrowRateEnabled(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveLastUpdate is a free data retrieval call binding the contract method 0x4f144609.
//
// Solidity: function getReserveLastUpdate(address _reserve) view returns(uint40 timestamp)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveLastUpdate(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveLastUpdate", _reserve)
	return *ret0, err
}

// GetReserveLastUpdate is a free data retrieval call binding the contract method 0x4f144609.
//
// Solidity: function getReserveLastUpdate(address _reserve) view returns(uint40 timestamp)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveLastUpdate(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveLastUpdate(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveLastUpdate is a free data retrieval call binding the contract method 0x4f144609.
//
// Solidity: function getReserveLastUpdate(address _reserve) view returns(uint40 timestamp)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveLastUpdate(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveLastUpdate(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveLiquidationBonus is a free data retrieval call binding the contract method 0xc76a6c9c.
//
// Solidity: function getReserveLiquidationBonus(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveLiquidationBonus(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveLiquidationBonus", _reserve)
	return *ret0, err
}

// GetReserveLiquidationBonus is a free data retrieval call binding the contract method 0xc76a6c9c.
//
// Solidity: function getReserveLiquidationBonus(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveLiquidationBonus(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveLiquidationBonus(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveLiquidationBonus is a free data retrieval call binding the contract method 0xc76a6c9c.
//
// Solidity: function getReserveLiquidationBonus(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveLiquidationBonus(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveLiquidationBonus(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveLiquidationThreshold is a free data retrieval call binding the contract method 0xf054ab46.
//
// Solidity: function getReserveLiquidationThreshold(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveLiquidationThreshold(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveLiquidationThreshold", _reserve)
	return *ret0, err
}

// GetReserveLiquidationThreshold is a free data retrieval call binding the contract method 0xf054ab46.
//
// Solidity: function getReserveLiquidationThreshold(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveLiquidationThreshold(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveLiquidationThreshold(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveLiquidationThreshold is a free data retrieval call binding the contract method 0xf054ab46.
//
// Solidity: function getReserveLiquidationThreshold(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveLiquidationThreshold(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveLiquidationThreshold(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveLiquidityCumulativeIndex is a free data retrieval call binding the contract method 0xbd7fd79a.
//
// Solidity: function getReserveLiquidityCumulativeIndex(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveLiquidityCumulativeIndex(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveLiquidityCumulativeIndex", _reserve)
	return *ret0, err
}

// GetReserveLiquidityCumulativeIndex is a free data retrieval call binding the contract method 0xbd7fd79a.
//
// Solidity: function getReserveLiquidityCumulativeIndex(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveLiquidityCumulativeIndex(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveLiquidityCumulativeIndex(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveLiquidityCumulativeIndex is a free data retrieval call binding the contract method 0xbd7fd79a.
//
// Solidity: function getReserveLiquidityCumulativeIndex(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveLiquidityCumulativeIndex(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveLiquidityCumulativeIndex(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveNormalizedIncome(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveNormalizedIncome", _reserve)
	return *ret0, err
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveNormalizedIncome(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveNormalizedIncome(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveNormalizedIncome(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveNormalizedIncome(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveTotalBorrows is a free data retrieval call binding the contract method 0xe6d18190.
//
// Solidity: function getReserveTotalBorrows(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveTotalBorrows(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveTotalBorrows", _reserve)
	return *ret0, err
}

// GetReserveTotalBorrows is a free data retrieval call binding the contract method 0xe6d18190.
//
// Solidity: function getReserveTotalBorrows(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveTotalBorrows(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveTotalBorrows(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveTotalBorrows is a free data retrieval call binding the contract method 0xe6d18190.
//
// Solidity: function getReserveTotalBorrows(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveTotalBorrows(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveTotalBorrows(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveTotalBorrowsStable is a free data retrieval call binding the contract method 0x7f90fec5.
//
// Solidity: function getReserveTotalBorrowsStable(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveTotalBorrowsStable(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveTotalBorrowsStable", _reserve)
	return *ret0, err
}

// GetReserveTotalBorrowsStable is a free data retrieval call binding the contract method 0x7f90fec5.
//
// Solidity: function getReserveTotalBorrowsStable(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveTotalBorrowsStable(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveTotalBorrowsStable(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveTotalBorrowsStable is a free data retrieval call binding the contract method 0x7f90fec5.
//
// Solidity: function getReserveTotalBorrowsStable(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveTotalBorrowsStable(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveTotalBorrowsStable(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveTotalBorrowsVariable is a free data retrieval call binding the contract method 0x98bd4737.
//
// Solidity: function getReserveTotalBorrowsVariable(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveTotalBorrowsVariable(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveTotalBorrowsVariable", _reserve)
	return *ret0, err
}

// GetReserveTotalBorrowsVariable is a free data retrieval call binding the contract method 0x98bd4737.
//
// Solidity: function getReserveTotalBorrowsVariable(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveTotalBorrowsVariable(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveTotalBorrowsVariable(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveTotalBorrowsVariable is a free data retrieval call binding the contract method 0x98bd4737.
//
// Solidity: function getReserveTotalBorrowsVariable(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveTotalBorrowsVariable(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveTotalBorrowsVariable(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveTotalLiquidity is a free data retrieval call binding the contract method 0xc33cfd90.
//
// Solidity: function getReserveTotalLiquidity(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveTotalLiquidity(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveTotalLiquidity", _reserve)
	return *ret0, err
}

// GetReserveTotalLiquidity is a free data retrieval call binding the contract method 0xc33cfd90.
//
// Solidity: function getReserveTotalLiquidity(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveTotalLiquidity(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveTotalLiquidity(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveTotalLiquidity is a free data retrieval call binding the contract method 0xc33cfd90.
//
// Solidity: function getReserveTotalLiquidity(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveTotalLiquidity(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveTotalLiquidity(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveUtilizationRate is a free data retrieval call binding the contract method 0xbfacad84.
//
// Solidity: function getReserveUtilizationRate(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveUtilizationRate(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveUtilizationRate", _reserve)
	return *ret0, err
}

// GetReserveUtilizationRate is a free data retrieval call binding the contract method 0xbfacad84.
//
// Solidity: function getReserveUtilizationRate(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveUtilizationRate(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveUtilizationRate(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveUtilizationRate is a free data retrieval call binding the contract method 0xbfacad84.
//
// Solidity: function getReserveUtilizationRate(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveUtilizationRate(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveUtilizationRate(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveVariableBorrowsCumulativeIndex is a free data retrieval call binding the contract method 0xb701d093.
//
// Solidity: function getReserveVariableBorrowsCumulativeIndex(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserveVariableBorrowsCumulativeIndex(opts *bind.CallOpts, _reserve common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserveVariableBorrowsCumulativeIndex", _reserve)
	return *ret0, err
}

// GetReserveVariableBorrowsCumulativeIndex is a free data retrieval call binding the contract method 0xb701d093.
//
// Solidity: function getReserveVariableBorrowsCumulativeIndex(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetReserveVariableBorrowsCumulativeIndex(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveVariableBorrowsCumulativeIndex(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserveVariableBorrowsCumulativeIndex is a free data retrieval call binding the contract method 0xb701d093.
//
// Solidity: function getReserveVariableBorrowsCumulativeIndex(address _reserve) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserveVariableBorrowsCumulativeIndex(_reserve common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetReserveVariableBorrowsCumulativeIndex(&_LendingPoolCore.CallOpts, _reserve)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(address[])
func (_LendingPoolCore *LendingPoolCoreCaller) GetReserves(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getReserves")
	return *ret0, err
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(address[])
func (_LendingPoolCore *LendingPoolCoreSession) GetReserves() ([]common.Address, error) {
	return _LendingPoolCore.Contract.GetReserves(&_LendingPoolCore.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(address[])
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetReserves() ([]common.Address, error) {
	return _LendingPoolCore.Contract.GetReserves(&_LendingPoolCore.CallOpts)
}

// GetUserBasicReserveData is a free data retrieval call binding the contract method 0xe10076ad.
//
// Solidity: function getUserBasicReserveData(address _reserve, address _user) view returns(uint256, uint256, uint256, bool)
func (_LendingPoolCore *LendingPoolCoreCaller) GetUserBasicReserveData(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (*big.Int, *big.Int, *big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _LendingPoolCore.contract.Call(opts, out, "getUserBasicReserveData", _reserve, _user)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetUserBasicReserveData is a free data retrieval call binding the contract method 0xe10076ad.
//
// Solidity: function getUserBasicReserveData(address _reserve, address _user) view returns(uint256, uint256, uint256, bool)
func (_LendingPoolCore *LendingPoolCoreSession) GetUserBasicReserveData(_reserve common.Address, _user common.Address) (*big.Int, *big.Int, *big.Int, bool, error) {
	return _LendingPoolCore.Contract.GetUserBasicReserveData(&_LendingPoolCore.CallOpts, _reserve, _user)
}

// GetUserBasicReserveData is a free data retrieval call binding the contract method 0xe10076ad.
//
// Solidity: function getUserBasicReserveData(address _reserve, address _user) view returns(uint256, uint256, uint256, bool)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetUserBasicReserveData(_reserve common.Address, _user common.Address) (*big.Int, *big.Int, *big.Int, bool, error) {
	return _LendingPoolCore.Contract.GetUserBasicReserveData(&_LendingPoolCore.CallOpts, _reserve, _user)
}

// GetUserBorrowBalances is a free data retrieval call binding the contract method 0x9fb8afcd.
//
// Solidity: function getUserBorrowBalances(address _reserve, address _user) view returns(uint256, uint256, uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetUserBorrowBalances(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _LendingPoolCore.contract.Call(opts, out, "getUserBorrowBalances", _reserve, _user)
	return *ret0, *ret1, *ret2, err
}

// GetUserBorrowBalances is a free data retrieval call binding the contract method 0x9fb8afcd.
//
// Solidity: function getUserBorrowBalances(address _reserve, address _user) view returns(uint256, uint256, uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetUserBorrowBalances(_reserve common.Address, _user common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _LendingPoolCore.Contract.GetUserBorrowBalances(&_LendingPoolCore.CallOpts, _reserve, _user)
}

// GetUserBorrowBalances is a free data retrieval call binding the contract method 0x9fb8afcd.
//
// Solidity: function getUserBorrowBalances(address _reserve, address _user) view returns(uint256, uint256, uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetUserBorrowBalances(_reserve common.Address, _user common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _LendingPoolCore.Contract.GetUserBorrowBalances(&_LendingPoolCore.CallOpts, _reserve, _user)
}

// GetUserCurrentBorrowRateMode is a free data retrieval call binding the contract method 0x1ca19f19.
//
// Solidity: function getUserCurrentBorrowRateMode(address _reserve, address _user) view returns(uint8)
func (_LendingPoolCore *LendingPoolCoreCaller) GetUserCurrentBorrowRateMode(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getUserCurrentBorrowRateMode", _reserve, _user)
	return *ret0, err
}

// GetUserCurrentBorrowRateMode is a free data retrieval call binding the contract method 0x1ca19f19.
//
// Solidity: function getUserCurrentBorrowRateMode(address _reserve, address _user) view returns(uint8)
func (_LendingPoolCore *LendingPoolCoreSession) GetUserCurrentBorrowRateMode(_reserve common.Address, _user common.Address) (uint8, error) {
	return _LendingPoolCore.Contract.GetUserCurrentBorrowRateMode(&_LendingPoolCore.CallOpts, _reserve, _user)
}

// GetUserCurrentBorrowRateMode is a free data retrieval call binding the contract method 0x1ca19f19.
//
// Solidity: function getUserCurrentBorrowRateMode(address _reserve, address _user) view returns(uint8)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetUserCurrentBorrowRateMode(_reserve common.Address, _user common.Address) (uint8, error) {
	return _LendingPoolCore.Contract.GetUserCurrentBorrowRateMode(&_LendingPoolCore.CallOpts, _reserve, _user)
}

// GetUserCurrentStableBorrowRate is a free data retrieval call binding the contract method 0x6fffab0c.
//
// Solidity: function getUserCurrentStableBorrowRate(address _reserve, address _user) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetUserCurrentStableBorrowRate(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getUserCurrentStableBorrowRate", _reserve, _user)
	return *ret0, err
}

// GetUserCurrentStableBorrowRate is a free data retrieval call binding the contract method 0x6fffab0c.
//
// Solidity: function getUserCurrentStableBorrowRate(address _reserve, address _user) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetUserCurrentStableBorrowRate(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetUserCurrentStableBorrowRate(&_LendingPoolCore.CallOpts, _reserve, _user)
}

// GetUserCurrentStableBorrowRate is a free data retrieval call binding the contract method 0x6fffab0c.
//
// Solidity: function getUserCurrentStableBorrowRate(address _reserve, address _user) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetUserCurrentStableBorrowRate(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetUserCurrentStableBorrowRate(&_LendingPoolCore.CallOpts, _reserve, _user)
}

// GetUserLastUpdate is a free data retrieval call binding the contract method 0x66d103f3.
//
// Solidity: function getUserLastUpdate(address _reserve, address _user) view returns(uint256 timestamp)
func (_LendingPoolCore *LendingPoolCoreCaller) GetUserLastUpdate(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getUserLastUpdate", _reserve, _user)
	return *ret0, err
}

// GetUserLastUpdate is a free data retrieval call binding the contract method 0x66d103f3.
//
// Solidity: function getUserLastUpdate(address _reserve, address _user) view returns(uint256 timestamp)
func (_LendingPoolCore *LendingPoolCoreSession) GetUserLastUpdate(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetUserLastUpdate(&_LendingPoolCore.CallOpts, _reserve, _user)
}

// GetUserLastUpdate is a free data retrieval call binding the contract method 0x66d103f3.
//
// Solidity: function getUserLastUpdate(address _reserve, address _user) view returns(uint256 timestamp)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetUserLastUpdate(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetUserLastUpdate(&_LendingPoolCore.CallOpts, _reserve, _user)
}

// GetUserOriginationFee is a free data retrieval call binding the contract method 0xfeab31ac.
//
// Solidity: function getUserOriginationFee(address _reserve, address _user) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetUserOriginationFee(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getUserOriginationFee", _reserve, _user)
	return *ret0, err
}

// GetUserOriginationFee is a free data retrieval call binding the contract method 0xfeab31ac.
//
// Solidity: function getUserOriginationFee(address _reserve, address _user) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetUserOriginationFee(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetUserOriginationFee(&_LendingPoolCore.CallOpts, _reserve, _user)
}

// GetUserOriginationFee is a free data retrieval call binding the contract method 0xfeab31ac.
//
// Solidity: function getUserOriginationFee(address _reserve, address _user) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetUserOriginationFee(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetUserOriginationFee(&_LendingPoolCore.CallOpts, _reserve, _user)
}

// GetUserUnderlyingAssetBalance is a free data retrieval call binding the contract method 0x18a4dbca.
//
// Solidity: function getUserUnderlyingAssetBalance(address _reserve, address _user) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetUserUnderlyingAssetBalance(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getUserUnderlyingAssetBalance", _reserve, _user)
	return *ret0, err
}

// GetUserUnderlyingAssetBalance is a free data retrieval call binding the contract method 0x18a4dbca.
//
// Solidity: function getUserUnderlyingAssetBalance(address _reserve, address _user) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetUserUnderlyingAssetBalance(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetUserUnderlyingAssetBalance(&_LendingPoolCore.CallOpts, _reserve, _user)
}

// GetUserUnderlyingAssetBalance is a free data retrieval call binding the contract method 0x18a4dbca.
//
// Solidity: function getUserUnderlyingAssetBalance(address _reserve, address _user) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetUserUnderlyingAssetBalance(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetUserUnderlyingAssetBalance(&_LendingPoolCore.CallOpts, _reserve, _user)
}

// GetUserVariableBorrowCumulativeIndex is a free data retrieval call binding the contract method 0xf6ea8d76.
//
// Solidity: function getUserVariableBorrowCumulativeIndex(address _reserve, address _user) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCaller) GetUserVariableBorrowCumulativeIndex(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "getUserVariableBorrowCumulativeIndex", _reserve, _user)
	return *ret0, err
}

// GetUserVariableBorrowCumulativeIndex is a free data retrieval call binding the contract method 0xf6ea8d76.
//
// Solidity: function getUserVariableBorrowCumulativeIndex(address _reserve, address _user) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) GetUserVariableBorrowCumulativeIndex(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetUserVariableBorrowCumulativeIndex(&_LendingPoolCore.CallOpts, _reserve, _user)
}

// GetUserVariableBorrowCumulativeIndex is a free data retrieval call binding the contract method 0xf6ea8d76.
//
// Solidity: function getUserVariableBorrowCumulativeIndex(address _reserve, address _user) view returns(uint256)
func (_LendingPoolCore *LendingPoolCoreCallerSession) GetUserVariableBorrowCumulativeIndex(_reserve common.Address, _user common.Address) (*big.Int, error) {
	return _LendingPoolCore.Contract.GetUserVariableBorrowCumulativeIndex(&_LendingPoolCore.CallOpts, _reserve, _user)
}

// IsReserveBorrowingEnabled is a free data retrieval call binding the contract method 0x5cf2e656.
//
// Solidity: function isReserveBorrowingEnabled(address _reserve) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreCaller) IsReserveBorrowingEnabled(opts *bind.CallOpts, _reserve common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "isReserveBorrowingEnabled", _reserve)
	return *ret0, err
}

// IsReserveBorrowingEnabled is a free data retrieval call binding the contract method 0x5cf2e656.
//
// Solidity: function isReserveBorrowingEnabled(address _reserve) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreSession) IsReserveBorrowingEnabled(_reserve common.Address) (bool, error) {
	return _LendingPoolCore.Contract.IsReserveBorrowingEnabled(&_LendingPoolCore.CallOpts, _reserve)
}

// IsReserveBorrowingEnabled is a free data retrieval call binding the contract method 0x5cf2e656.
//
// Solidity: function isReserveBorrowingEnabled(address _reserve) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreCallerSession) IsReserveBorrowingEnabled(_reserve common.Address) (bool, error) {
	return _LendingPoolCore.Contract.IsReserveBorrowingEnabled(&_LendingPoolCore.CallOpts, _reserve)
}

// IsReserveUsageAsCollateralEnabled is a free data retrieval call binding the contract method 0x18f9bbae.
//
// Solidity: function isReserveUsageAsCollateralEnabled(address _reserve) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreCaller) IsReserveUsageAsCollateralEnabled(opts *bind.CallOpts, _reserve common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "isReserveUsageAsCollateralEnabled", _reserve)
	return *ret0, err
}

// IsReserveUsageAsCollateralEnabled is a free data retrieval call binding the contract method 0x18f9bbae.
//
// Solidity: function isReserveUsageAsCollateralEnabled(address _reserve) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreSession) IsReserveUsageAsCollateralEnabled(_reserve common.Address) (bool, error) {
	return _LendingPoolCore.Contract.IsReserveUsageAsCollateralEnabled(&_LendingPoolCore.CallOpts, _reserve)
}

// IsReserveUsageAsCollateralEnabled is a free data retrieval call binding the contract method 0x18f9bbae.
//
// Solidity: function isReserveUsageAsCollateralEnabled(address _reserve) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreCallerSession) IsReserveUsageAsCollateralEnabled(_reserve common.Address) (bool, error) {
	return _LendingPoolCore.Contract.IsReserveUsageAsCollateralEnabled(&_LendingPoolCore.CallOpts, _reserve)
}

// IsUserAllowedToBorrowAtStable is a free data retrieval call binding the contract method 0xe2174d86.
//
// Solidity: function isUserAllowedToBorrowAtStable(address _reserve, address _user, uint256 _amount) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreCaller) IsUserAllowedToBorrowAtStable(opts *bind.CallOpts, _reserve common.Address, _user common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "isUserAllowedToBorrowAtStable", _reserve, _user, _amount)
	return *ret0, err
}

// IsUserAllowedToBorrowAtStable is a free data retrieval call binding the contract method 0xe2174d86.
//
// Solidity: function isUserAllowedToBorrowAtStable(address _reserve, address _user, uint256 _amount) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreSession) IsUserAllowedToBorrowAtStable(_reserve common.Address, _user common.Address, _amount *big.Int) (bool, error) {
	return _LendingPoolCore.Contract.IsUserAllowedToBorrowAtStable(&_LendingPoolCore.CallOpts, _reserve, _user, _amount)
}

// IsUserAllowedToBorrowAtStable is a free data retrieval call binding the contract method 0xe2174d86.
//
// Solidity: function isUserAllowedToBorrowAtStable(address _reserve, address _user, uint256 _amount) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreCallerSession) IsUserAllowedToBorrowAtStable(_reserve common.Address, _user common.Address, _amount *big.Int) (bool, error) {
	return _LendingPoolCore.Contract.IsUserAllowedToBorrowAtStable(&_LendingPoolCore.CallOpts, _reserve, _user, _amount)
}

// IsUserUseReserveAsCollateralEnabled is a free data retrieval call binding the contract method 0x9e3c4f3b.
//
// Solidity: function isUserUseReserveAsCollateralEnabled(address _reserve, address _user) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreCaller) IsUserUseReserveAsCollateralEnabled(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "isUserUseReserveAsCollateralEnabled", _reserve, _user)
	return *ret0, err
}

// IsUserUseReserveAsCollateralEnabled is a free data retrieval call binding the contract method 0x9e3c4f3b.
//
// Solidity: function isUserUseReserveAsCollateralEnabled(address _reserve, address _user) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreSession) IsUserUseReserveAsCollateralEnabled(_reserve common.Address, _user common.Address) (bool, error) {
	return _LendingPoolCore.Contract.IsUserUseReserveAsCollateralEnabled(&_LendingPoolCore.CallOpts, _reserve, _user)
}

// IsUserUseReserveAsCollateralEnabled is a free data retrieval call binding the contract method 0x9e3c4f3b.
//
// Solidity: function isUserUseReserveAsCollateralEnabled(address _reserve, address _user) view returns(bool)
func (_LendingPoolCore *LendingPoolCoreCallerSession) IsUserUseReserveAsCollateralEnabled(_reserve common.Address, _user common.Address) (bool, error) {
	return _LendingPoolCore.Contract.IsUserUseReserveAsCollateralEnabled(&_LendingPoolCore.CallOpts, _reserve, _user)
}

// LendingPoolAddress is a free data retrieval call binding the contract method 0xd3ae26b3.
//
// Solidity: function lendingPoolAddress() view returns(address)
func (_LendingPoolCore *LendingPoolCoreCaller) LendingPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "lendingPoolAddress")
	return *ret0, err
}

// LendingPoolAddress is a free data retrieval call binding the contract method 0xd3ae26b3.
//
// Solidity: function lendingPoolAddress() view returns(address)
func (_LendingPoolCore *LendingPoolCoreSession) LendingPoolAddress() (common.Address, error) {
	return _LendingPoolCore.Contract.LendingPoolAddress(&_LendingPoolCore.CallOpts)
}

// LendingPoolAddress is a free data retrieval call binding the contract method 0xd3ae26b3.
//
// Solidity: function lendingPoolAddress() view returns(address)
func (_LendingPoolCore *LendingPoolCoreCallerSession) LendingPoolAddress() (common.Address, error) {
	return _LendingPoolCore.Contract.LendingPoolAddress(&_LendingPoolCore.CallOpts)
}

// ReservesList is a free data retrieval call binding the contract method 0x4fe7a6e5.
//
// Solidity: function reservesList(uint256 ) view returns(address)
func (_LendingPoolCore *LendingPoolCoreCaller) ReservesList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LendingPoolCore.contract.Call(opts, out, "reservesList", arg0)
	return *ret0, err
}

// ReservesList is a free data retrieval call binding the contract method 0x4fe7a6e5.
//
// Solidity: function reservesList(uint256 ) view returns(address)
func (_LendingPoolCore *LendingPoolCoreSession) ReservesList(arg0 *big.Int) (common.Address, error) {
	return _LendingPoolCore.Contract.ReservesList(&_LendingPoolCore.CallOpts, arg0)
}

// ReservesList is a free data retrieval call binding the contract method 0x4fe7a6e5.
//
// Solidity: function reservesList(uint256 ) view returns(address)
func (_LendingPoolCore *LendingPoolCoreCallerSession) ReservesList(arg0 *big.Int) (common.Address, error) {
	return _LendingPoolCore.Contract.ReservesList(&_LendingPoolCore.CallOpts, arg0)
}

// ActivateReserve is a paid mutator transaction binding the contract method 0xb75d6f34.
//
// Solidity: function activateReserve(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) ActivateReserve(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "activateReserve", _reserve)
}

// ActivateReserve is a paid mutator transaction binding the contract method 0xb75d6f34.
//
// Solidity: function activateReserve(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreSession) ActivateReserve(_reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.ActivateReserve(&_LendingPoolCore.TransactOpts, _reserve)
}

// ActivateReserve is a paid mutator transaction binding the contract method 0xb75d6f34.
//
// Solidity: function activateReserve(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) ActivateReserve(_reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.ActivateReserve(&_LendingPoolCore.TransactOpts, _reserve)
}

// DeactivateReserve is a paid mutator transaction binding the contract method 0x3e72a454.
//
// Solidity: function deactivateReserve(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) DeactivateReserve(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "deactivateReserve", _reserve)
}

// DeactivateReserve is a paid mutator transaction binding the contract method 0x3e72a454.
//
// Solidity: function deactivateReserve(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreSession) DeactivateReserve(_reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.DeactivateReserve(&_LendingPoolCore.TransactOpts, _reserve)
}

// DeactivateReserve is a paid mutator transaction binding the contract method 0x3e72a454.
//
// Solidity: function deactivateReserve(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) DeactivateReserve(_reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.DeactivateReserve(&_LendingPoolCore.TransactOpts, _reserve)
}

// DisableBorrowingOnReserve is a paid mutator transaction binding the contract method 0xa8dc0f45.
//
// Solidity: function disableBorrowingOnReserve(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) DisableBorrowingOnReserve(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "disableBorrowingOnReserve", _reserve)
}

// DisableBorrowingOnReserve is a paid mutator transaction binding the contract method 0xa8dc0f45.
//
// Solidity: function disableBorrowingOnReserve(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreSession) DisableBorrowingOnReserve(_reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.DisableBorrowingOnReserve(&_LendingPoolCore.TransactOpts, _reserve)
}

// DisableBorrowingOnReserve is a paid mutator transaction binding the contract method 0xa8dc0f45.
//
// Solidity: function disableBorrowingOnReserve(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) DisableBorrowingOnReserve(_reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.DisableBorrowingOnReserve(&_LendingPoolCore.TransactOpts, _reserve)
}

// DisableReserveAsCollateral is a paid mutator transaction binding the contract method 0xe8ae2f5b.
//
// Solidity: function disableReserveAsCollateral(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) DisableReserveAsCollateral(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "disableReserveAsCollateral", _reserve)
}

// DisableReserveAsCollateral is a paid mutator transaction binding the contract method 0xe8ae2f5b.
//
// Solidity: function disableReserveAsCollateral(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreSession) DisableReserveAsCollateral(_reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.DisableReserveAsCollateral(&_LendingPoolCore.TransactOpts, _reserve)
}

// DisableReserveAsCollateral is a paid mutator transaction binding the contract method 0xe8ae2f5b.
//
// Solidity: function disableReserveAsCollateral(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) DisableReserveAsCollateral(_reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.DisableReserveAsCollateral(&_LendingPoolCore.TransactOpts, _reserve)
}

// DisableReserveStableBorrowRate is a paid mutator transaction binding the contract method 0xb8c0f745.
//
// Solidity: function disableReserveStableBorrowRate(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) DisableReserveStableBorrowRate(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "disableReserveStableBorrowRate", _reserve)
}

// DisableReserveStableBorrowRate is a paid mutator transaction binding the contract method 0xb8c0f745.
//
// Solidity: function disableReserveStableBorrowRate(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreSession) DisableReserveStableBorrowRate(_reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.DisableReserveStableBorrowRate(&_LendingPoolCore.TransactOpts, _reserve)
}

// DisableReserveStableBorrowRate is a paid mutator transaction binding the contract method 0xb8c0f745.
//
// Solidity: function disableReserveStableBorrowRate(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) DisableReserveStableBorrowRate(_reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.DisableReserveStableBorrowRate(&_LendingPoolCore.TransactOpts, _reserve)
}

// EnableBorrowingOnReserve is a paid mutator transaction binding the contract method 0xeede87c1.
//
// Solidity: function enableBorrowingOnReserve(address _reserve, bool _stableBorrowRateEnabled) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) EnableBorrowingOnReserve(opts *bind.TransactOpts, _reserve common.Address, _stableBorrowRateEnabled bool) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "enableBorrowingOnReserve", _reserve, _stableBorrowRateEnabled)
}

// EnableBorrowingOnReserve is a paid mutator transaction binding the contract method 0xeede87c1.
//
// Solidity: function enableBorrowingOnReserve(address _reserve, bool _stableBorrowRateEnabled) returns()
func (_LendingPoolCore *LendingPoolCoreSession) EnableBorrowingOnReserve(_reserve common.Address, _stableBorrowRateEnabled bool) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.EnableBorrowingOnReserve(&_LendingPoolCore.TransactOpts, _reserve, _stableBorrowRateEnabled)
}

// EnableBorrowingOnReserve is a paid mutator transaction binding the contract method 0xeede87c1.
//
// Solidity: function enableBorrowingOnReserve(address _reserve, bool _stableBorrowRateEnabled) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) EnableBorrowingOnReserve(_reserve common.Address, _stableBorrowRateEnabled bool) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.EnableBorrowingOnReserve(&_LendingPoolCore.TransactOpts, _reserve, _stableBorrowRateEnabled)
}

// EnableReserveAsCollateral is a paid mutator transaction binding the contract method 0xa5bc826c.
//
// Solidity: function enableReserveAsCollateral(address _reserve, uint256 _baseLTVasCollateral, uint256 _liquidationThreshold, uint256 _liquidationBonus) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) EnableReserveAsCollateral(opts *bind.TransactOpts, _reserve common.Address, _baseLTVasCollateral *big.Int, _liquidationThreshold *big.Int, _liquidationBonus *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "enableReserveAsCollateral", _reserve, _baseLTVasCollateral, _liquidationThreshold, _liquidationBonus)
}

// EnableReserveAsCollateral is a paid mutator transaction binding the contract method 0xa5bc826c.
//
// Solidity: function enableReserveAsCollateral(address _reserve, uint256 _baseLTVasCollateral, uint256 _liquidationThreshold, uint256 _liquidationBonus) returns()
func (_LendingPoolCore *LendingPoolCoreSession) EnableReserveAsCollateral(_reserve common.Address, _baseLTVasCollateral *big.Int, _liquidationThreshold *big.Int, _liquidationBonus *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.EnableReserveAsCollateral(&_LendingPoolCore.TransactOpts, _reserve, _baseLTVasCollateral, _liquidationThreshold, _liquidationBonus)
}

// EnableReserveAsCollateral is a paid mutator transaction binding the contract method 0xa5bc826c.
//
// Solidity: function enableReserveAsCollateral(address _reserve, uint256 _baseLTVasCollateral, uint256 _liquidationThreshold, uint256 _liquidationBonus) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) EnableReserveAsCollateral(_reserve common.Address, _baseLTVasCollateral *big.Int, _liquidationThreshold *big.Int, _liquidationBonus *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.EnableReserveAsCollateral(&_LendingPoolCore.TransactOpts, _reserve, _baseLTVasCollateral, _liquidationThreshold, _liquidationBonus)
}

// EnableReserveStableBorrowRate is a paid mutator transaction binding the contract method 0xdae4c4e7.
//
// Solidity: function enableReserveStableBorrowRate(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) EnableReserveStableBorrowRate(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "enableReserveStableBorrowRate", _reserve)
}

// EnableReserveStableBorrowRate is a paid mutator transaction binding the contract method 0xdae4c4e7.
//
// Solidity: function enableReserveStableBorrowRate(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreSession) EnableReserveStableBorrowRate(_reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.EnableReserveStableBorrowRate(&_LendingPoolCore.TransactOpts, _reserve)
}

// EnableReserveStableBorrowRate is a paid mutator transaction binding the contract method 0xdae4c4e7.
//
// Solidity: function enableReserveStableBorrowRate(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) EnableReserveStableBorrowRate(_reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.EnableReserveStableBorrowRate(&_LendingPoolCore.TransactOpts, _reserve)
}

// FreezeReserve is a paid mutator transaction binding the contract method 0x7aca76eb.
//
// Solidity: function freezeReserve(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) FreezeReserve(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "freezeReserve", _reserve)
}

// FreezeReserve is a paid mutator transaction binding the contract method 0x7aca76eb.
//
// Solidity: function freezeReserve(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreSession) FreezeReserve(_reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.FreezeReserve(&_LendingPoolCore.TransactOpts, _reserve)
}

// FreezeReserve is a paid mutator transaction binding the contract method 0x7aca76eb.
//
// Solidity: function freezeReserve(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) FreezeReserve(_reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.FreezeReserve(&_LendingPoolCore.TransactOpts, _reserve)
}

// InitReserve is a paid mutator transaction binding the contract method 0x45330a40.
//
// Solidity: function initReserve(address _reserve, address _aTokenAddress, uint256 _decimals, address _interestRateStrategyAddress) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) InitReserve(opts *bind.TransactOpts, _reserve common.Address, _aTokenAddress common.Address, _decimals *big.Int, _interestRateStrategyAddress common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "initReserve", _reserve, _aTokenAddress, _decimals, _interestRateStrategyAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x45330a40.
//
// Solidity: function initReserve(address _reserve, address _aTokenAddress, uint256 _decimals, address _interestRateStrategyAddress) returns()
func (_LendingPoolCore *LendingPoolCoreSession) InitReserve(_reserve common.Address, _aTokenAddress common.Address, _decimals *big.Int, _interestRateStrategyAddress common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.InitReserve(&_LendingPoolCore.TransactOpts, _reserve, _aTokenAddress, _decimals, _interestRateStrategyAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x45330a40.
//
// Solidity: function initReserve(address _reserve, address _aTokenAddress, uint256 _decimals, address _interestRateStrategyAddress) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) InitReserve(_reserve common.Address, _aTokenAddress common.Address, _decimals *big.Int, _interestRateStrategyAddress common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.InitReserve(&_LendingPoolCore.TransactOpts, _reserve, _aTokenAddress, _decimals, _interestRateStrategyAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _addressesProvider) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) Initialize(opts *bind.TransactOpts, _addressesProvider common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "initialize", _addressesProvider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _addressesProvider) returns()
func (_LendingPoolCore *LendingPoolCoreSession) Initialize(_addressesProvider common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.Initialize(&_LendingPoolCore.TransactOpts, _addressesProvider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _addressesProvider) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) Initialize(_addressesProvider common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.Initialize(&_LendingPoolCore.TransactOpts, _addressesProvider)
}

// LiquidateFee is a paid mutator transaction binding the contract method 0x8f385c22.
//
// Solidity: function liquidateFee(address _token, uint256 _amount, address _destination) payable returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) LiquidateFee(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _destination common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "liquidateFee", _token, _amount, _destination)
}

// LiquidateFee is a paid mutator transaction binding the contract method 0x8f385c22.
//
// Solidity: function liquidateFee(address _token, uint256 _amount, address _destination) payable returns()
func (_LendingPoolCore *LendingPoolCoreSession) LiquidateFee(_token common.Address, _amount *big.Int, _destination common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.LiquidateFee(&_LendingPoolCore.TransactOpts, _token, _amount, _destination)
}

// LiquidateFee is a paid mutator transaction binding the contract method 0x8f385c22.
//
// Solidity: function liquidateFee(address _token, uint256 _amount, address _destination) payable returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) LiquidateFee(_token common.Address, _amount *big.Int, _destination common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.LiquidateFee(&_LendingPoolCore.TransactOpts, _token, _amount, _destination)
}

// RefreshConfiguration is a paid mutator transaction binding the contract method 0x64681083.
//
// Solidity: function refreshConfiguration() returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) RefreshConfiguration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "refreshConfiguration")
}

// RefreshConfiguration is a paid mutator transaction binding the contract method 0x64681083.
//
// Solidity: function refreshConfiguration() returns()
func (_LendingPoolCore *LendingPoolCoreSession) RefreshConfiguration() (*types.Transaction, error) {
	return _LendingPoolCore.Contract.RefreshConfiguration(&_LendingPoolCore.TransactOpts)
}

// RefreshConfiguration is a paid mutator transaction binding the contract method 0x64681083.
//
// Solidity: function refreshConfiguration() returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) RefreshConfiguration() (*types.Transaction, error) {
	return _LendingPoolCore.Contract.RefreshConfiguration(&_LendingPoolCore.TransactOpts)
}

// RemoveLastAddedReserve is a paid mutator transaction binding the contract method 0xd06e2ec1.
//
// Solidity: function removeLastAddedReserve(address _reserveToRemove) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) RemoveLastAddedReserve(opts *bind.TransactOpts, _reserveToRemove common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "removeLastAddedReserve", _reserveToRemove)
}

// RemoveLastAddedReserve is a paid mutator transaction binding the contract method 0xd06e2ec1.
//
// Solidity: function removeLastAddedReserve(address _reserveToRemove) returns()
func (_LendingPoolCore *LendingPoolCoreSession) RemoveLastAddedReserve(_reserveToRemove common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.RemoveLastAddedReserve(&_LendingPoolCore.TransactOpts, _reserveToRemove)
}

// RemoveLastAddedReserve is a paid mutator transaction binding the contract method 0xd06e2ec1.
//
// Solidity: function removeLastAddedReserve(address _reserveToRemove) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) RemoveLastAddedReserve(_reserveToRemove common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.RemoveLastAddedReserve(&_LendingPoolCore.TransactOpts, _reserveToRemove)
}

// SetReserveBaseLTVasCollateral is a paid mutator transaction binding the contract method 0xd466016f.
//
// Solidity: function setReserveBaseLTVasCollateral(address _reserve, uint256 _ltv) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) SetReserveBaseLTVasCollateral(opts *bind.TransactOpts, _reserve common.Address, _ltv *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "setReserveBaseLTVasCollateral", _reserve, _ltv)
}

// SetReserveBaseLTVasCollateral is a paid mutator transaction binding the contract method 0xd466016f.
//
// Solidity: function setReserveBaseLTVasCollateral(address _reserve, uint256 _ltv) returns()
func (_LendingPoolCore *LendingPoolCoreSession) SetReserveBaseLTVasCollateral(_reserve common.Address, _ltv *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.SetReserveBaseLTVasCollateral(&_LendingPoolCore.TransactOpts, _reserve, _ltv)
}

// SetReserveBaseLTVasCollateral is a paid mutator transaction binding the contract method 0xd466016f.
//
// Solidity: function setReserveBaseLTVasCollateral(address _reserve, uint256 _ltv) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) SetReserveBaseLTVasCollateral(_reserve common.Address, _ltv *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.SetReserveBaseLTVasCollateral(&_LendingPoolCore.TransactOpts, _reserve, _ltv)
}

// SetReserveDecimals is a paid mutator transaction binding the contract method 0x66bbd928.
//
// Solidity: function setReserveDecimals(address _reserve, uint256 _decimals) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) SetReserveDecimals(opts *bind.TransactOpts, _reserve common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "setReserveDecimals", _reserve, _decimals)
}

// SetReserveDecimals is a paid mutator transaction binding the contract method 0x66bbd928.
//
// Solidity: function setReserveDecimals(address _reserve, uint256 _decimals) returns()
func (_LendingPoolCore *LendingPoolCoreSession) SetReserveDecimals(_reserve common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.SetReserveDecimals(&_LendingPoolCore.TransactOpts, _reserve, _decimals)
}

// SetReserveDecimals is a paid mutator transaction binding the contract method 0x66bbd928.
//
// Solidity: function setReserveDecimals(address _reserve, uint256 _decimals) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) SetReserveDecimals(_reserve common.Address, _decimals *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.SetReserveDecimals(&_LendingPoolCore.TransactOpts, _reserve, _decimals)
}

// SetReserveInterestRateStrategyAddress is a paid mutator transaction binding the contract method 0x1d2118f9.
//
// Solidity: function setReserveInterestRateStrategyAddress(address _reserve, address _rateStrategyAddress) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) SetReserveInterestRateStrategyAddress(opts *bind.TransactOpts, _reserve common.Address, _rateStrategyAddress common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "setReserveInterestRateStrategyAddress", _reserve, _rateStrategyAddress)
}

// SetReserveInterestRateStrategyAddress is a paid mutator transaction binding the contract method 0x1d2118f9.
//
// Solidity: function setReserveInterestRateStrategyAddress(address _reserve, address _rateStrategyAddress) returns()
func (_LendingPoolCore *LendingPoolCoreSession) SetReserveInterestRateStrategyAddress(_reserve common.Address, _rateStrategyAddress common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.SetReserveInterestRateStrategyAddress(&_LendingPoolCore.TransactOpts, _reserve, _rateStrategyAddress)
}

// SetReserveInterestRateStrategyAddress is a paid mutator transaction binding the contract method 0x1d2118f9.
//
// Solidity: function setReserveInterestRateStrategyAddress(address _reserve, address _rateStrategyAddress) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) SetReserveInterestRateStrategyAddress(_reserve common.Address, _rateStrategyAddress common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.SetReserveInterestRateStrategyAddress(&_LendingPoolCore.TransactOpts, _reserve, _rateStrategyAddress)
}

// SetReserveLiquidationBonus is a paid mutator transaction binding the contract method 0x70fb84f4.
//
// Solidity: function setReserveLiquidationBonus(address _reserve, uint256 _bonus) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) SetReserveLiquidationBonus(opts *bind.TransactOpts, _reserve common.Address, _bonus *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "setReserveLiquidationBonus", _reserve, _bonus)
}

// SetReserveLiquidationBonus is a paid mutator transaction binding the contract method 0x70fb84f4.
//
// Solidity: function setReserveLiquidationBonus(address _reserve, uint256 _bonus) returns()
func (_LendingPoolCore *LendingPoolCoreSession) SetReserveLiquidationBonus(_reserve common.Address, _bonus *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.SetReserveLiquidationBonus(&_LendingPoolCore.TransactOpts, _reserve, _bonus)
}

// SetReserveLiquidationBonus is a paid mutator transaction binding the contract method 0x70fb84f4.
//
// Solidity: function setReserveLiquidationBonus(address _reserve, uint256 _bonus) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) SetReserveLiquidationBonus(_reserve common.Address, _bonus *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.SetReserveLiquidationBonus(&_LendingPoolCore.TransactOpts, _reserve, _bonus)
}

// SetReserveLiquidationThreshold is a paid mutator transaction binding the contract method 0x3443a14b.
//
// Solidity: function setReserveLiquidationThreshold(address _reserve, uint256 _threshold) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) SetReserveLiquidationThreshold(opts *bind.TransactOpts, _reserve common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "setReserveLiquidationThreshold", _reserve, _threshold)
}

// SetReserveLiquidationThreshold is a paid mutator transaction binding the contract method 0x3443a14b.
//
// Solidity: function setReserveLiquidationThreshold(address _reserve, uint256 _threshold) returns()
func (_LendingPoolCore *LendingPoolCoreSession) SetReserveLiquidationThreshold(_reserve common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.SetReserveLiquidationThreshold(&_LendingPoolCore.TransactOpts, _reserve, _threshold)
}

// SetReserveLiquidationThreshold is a paid mutator transaction binding the contract method 0x3443a14b.
//
// Solidity: function setReserveLiquidationThreshold(address _reserve, uint256 _threshold) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) SetReserveLiquidationThreshold(_reserve common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.SetReserveLiquidationThreshold(&_LendingPoolCore.TransactOpts, _reserve, _threshold)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0xfa51854c.
//
// Solidity: function setUserUseReserveAsCollateral(address _reserve, address _user, bool _useAsCollateral) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) SetUserUseReserveAsCollateral(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "setUserUseReserveAsCollateral", _reserve, _user, _useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0xfa51854c.
//
// Solidity: function setUserUseReserveAsCollateral(address _reserve, address _user, bool _useAsCollateral) returns()
func (_LendingPoolCore *LendingPoolCoreSession) SetUserUseReserveAsCollateral(_reserve common.Address, _user common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.SetUserUseReserveAsCollateral(&_LendingPoolCore.TransactOpts, _reserve, _user, _useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0xfa51854c.
//
// Solidity: function setUserUseReserveAsCollateral(address _reserve, address _user, bool _useAsCollateral) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) SetUserUseReserveAsCollateral(_reserve common.Address, _user common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.SetUserUseReserveAsCollateral(&_LendingPoolCore.TransactOpts, _reserve, _user, _useAsCollateral)
}

// TransferToFeeCollectionAddress is a paid mutator transaction binding the contract method 0xc7d14237.
//
// Solidity: function transferToFeeCollectionAddress(address _token, address _user, uint256 _amount, address _destination) payable returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) TransferToFeeCollectionAddress(opts *bind.TransactOpts, _token common.Address, _user common.Address, _amount *big.Int, _destination common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "transferToFeeCollectionAddress", _token, _user, _amount, _destination)
}

// TransferToFeeCollectionAddress is a paid mutator transaction binding the contract method 0xc7d14237.
//
// Solidity: function transferToFeeCollectionAddress(address _token, address _user, uint256 _amount, address _destination) payable returns()
func (_LendingPoolCore *LendingPoolCoreSession) TransferToFeeCollectionAddress(_token common.Address, _user common.Address, _amount *big.Int, _destination common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.TransferToFeeCollectionAddress(&_LendingPoolCore.TransactOpts, _token, _user, _amount, _destination)
}

// TransferToFeeCollectionAddress is a paid mutator transaction binding the contract method 0xc7d14237.
//
// Solidity: function transferToFeeCollectionAddress(address _token, address _user, uint256 _amount, address _destination) payable returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) TransferToFeeCollectionAddress(_token common.Address, _user common.Address, _amount *big.Int, _destination common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.TransferToFeeCollectionAddress(&_LendingPoolCore.TransactOpts, _token, _user, _amount, _destination)
}

// TransferToReserve is a paid mutator transaction binding the contract method 0x28fcf4d3.
//
// Solidity: function transferToReserve(address _reserve, address _user, uint256 _amount) payable returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) TransferToReserve(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "transferToReserve", _reserve, _user, _amount)
}

// TransferToReserve is a paid mutator transaction binding the contract method 0x28fcf4d3.
//
// Solidity: function transferToReserve(address _reserve, address _user, uint256 _amount) payable returns()
func (_LendingPoolCore *LendingPoolCoreSession) TransferToReserve(_reserve common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.TransferToReserve(&_LendingPoolCore.TransactOpts, _reserve, _user, _amount)
}

// TransferToReserve is a paid mutator transaction binding the contract method 0x28fcf4d3.
//
// Solidity: function transferToReserve(address _reserve, address _user, uint256 _amount) payable returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) TransferToReserve(_reserve common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.TransferToReserve(&_LendingPoolCore.TransactOpts, _reserve, _user, _amount)
}

// TransferToUser is a paid mutator transaction binding the contract method 0xfa93b2a5.
//
// Solidity: function transferToUser(address _reserve, address _user, uint256 _amount) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) TransferToUser(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "transferToUser", _reserve, _user, _amount)
}

// TransferToUser is a paid mutator transaction binding the contract method 0xfa93b2a5.
//
// Solidity: function transferToUser(address _reserve, address _user, uint256 _amount) returns()
func (_LendingPoolCore *LendingPoolCoreSession) TransferToUser(_reserve common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.TransferToUser(&_LendingPoolCore.TransactOpts, _reserve, _user, _amount)
}

// TransferToUser is a paid mutator transaction binding the contract method 0xfa93b2a5.
//
// Solidity: function transferToUser(address _reserve, address _user, uint256 _amount) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) TransferToUser(_reserve common.Address, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.TransferToUser(&_LendingPoolCore.TransactOpts, _reserve, _user, _amount)
}

// UnfreezeReserve is a paid mutator transaction binding the contract method 0xef1f9373.
//
// Solidity: function unfreezeReserve(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) UnfreezeReserve(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "unfreezeReserve", _reserve)
}

// UnfreezeReserve is a paid mutator transaction binding the contract method 0xef1f9373.
//
// Solidity: function unfreezeReserve(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreSession) UnfreezeReserve(_reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.UnfreezeReserve(&_LendingPoolCore.TransactOpts, _reserve)
}

// UnfreezeReserve is a paid mutator transaction binding the contract method 0xef1f9373.
//
// Solidity: function unfreezeReserve(address _reserve) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) UnfreezeReserve(_reserve common.Address) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.UnfreezeReserve(&_LendingPoolCore.TransactOpts, _reserve)
}

// UpdateStateOnBorrow is a paid mutator transaction binding the contract method 0x37ac6fe4.
//
// Solidity: function updateStateOnBorrow(address _reserve, address _user, uint256 _amountBorrowed, uint256 _borrowFee, uint8 _rateMode) returns(uint256, uint256)
func (_LendingPoolCore *LendingPoolCoreTransactor) UpdateStateOnBorrow(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _amountBorrowed *big.Int, _borrowFee *big.Int, _rateMode uint8) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "updateStateOnBorrow", _reserve, _user, _amountBorrowed, _borrowFee, _rateMode)
}

// UpdateStateOnBorrow is a paid mutator transaction binding the contract method 0x37ac6fe4.
//
// Solidity: function updateStateOnBorrow(address _reserve, address _user, uint256 _amountBorrowed, uint256 _borrowFee, uint8 _rateMode) returns(uint256, uint256)
func (_LendingPoolCore *LendingPoolCoreSession) UpdateStateOnBorrow(_reserve common.Address, _user common.Address, _amountBorrowed *big.Int, _borrowFee *big.Int, _rateMode uint8) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.UpdateStateOnBorrow(&_LendingPoolCore.TransactOpts, _reserve, _user, _amountBorrowed, _borrowFee, _rateMode)
}

// UpdateStateOnBorrow is a paid mutator transaction binding the contract method 0x37ac6fe4.
//
// Solidity: function updateStateOnBorrow(address _reserve, address _user, uint256 _amountBorrowed, uint256 _borrowFee, uint8 _rateMode) returns(uint256, uint256)
func (_LendingPoolCore *LendingPoolCoreTransactorSession) UpdateStateOnBorrow(_reserve common.Address, _user common.Address, _amountBorrowed *big.Int, _borrowFee *big.Int, _rateMode uint8) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.UpdateStateOnBorrow(&_LendingPoolCore.TransactOpts, _reserve, _user, _amountBorrowed, _borrowFee, _rateMode)
}

// UpdateStateOnDeposit is a paid mutator transaction binding the contract method 0xbcd6ffa4.
//
// Solidity: function updateStateOnDeposit(address _reserve, address _user, uint256 _amount, bool _isFirstDeposit) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) UpdateStateOnDeposit(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _amount *big.Int, _isFirstDeposit bool) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "updateStateOnDeposit", _reserve, _user, _amount, _isFirstDeposit)
}

// UpdateStateOnDeposit is a paid mutator transaction binding the contract method 0xbcd6ffa4.
//
// Solidity: function updateStateOnDeposit(address _reserve, address _user, uint256 _amount, bool _isFirstDeposit) returns()
func (_LendingPoolCore *LendingPoolCoreSession) UpdateStateOnDeposit(_reserve common.Address, _user common.Address, _amount *big.Int, _isFirstDeposit bool) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.UpdateStateOnDeposit(&_LendingPoolCore.TransactOpts, _reserve, _user, _amount, _isFirstDeposit)
}

// UpdateStateOnDeposit is a paid mutator transaction binding the contract method 0xbcd6ffa4.
//
// Solidity: function updateStateOnDeposit(address _reserve, address _user, uint256 _amount, bool _isFirstDeposit) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) UpdateStateOnDeposit(_reserve common.Address, _user common.Address, _amount *big.Int, _isFirstDeposit bool) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.UpdateStateOnDeposit(&_LendingPoolCore.TransactOpts, _reserve, _user, _amount, _isFirstDeposit)
}

// UpdateStateOnFlashLoan is a paid mutator transaction binding the contract method 0x09ac2953.
//
// Solidity: function updateStateOnFlashLoan(address _reserve, uint256 _availableLiquidityBefore, uint256 _income, uint256 _protocolFee) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) UpdateStateOnFlashLoan(opts *bind.TransactOpts, _reserve common.Address, _availableLiquidityBefore *big.Int, _income *big.Int, _protocolFee *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "updateStateOnFlashLoan", _reserve, _availableLiquidityBefore, _income, _protocolFee)
}

// UpdateStateOnFlashLoan is a paid mutator transaction binding the contract method 0x09ac2953.
//
// Solidity: function updateStateOnFlashLoan(address _reserve, uint256 _availableLiquidityBefore, uint256 _income, uint256 _protocolFee) returns()
func (_LendingPoolCore *LendingPoolCoreSession) UpdateStateOnFlashLoan(_reserve common.Address, _availableLiquidityBefore *big.Int, _income *big.Int, _protocolFee *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.UpdateStateOnFlashLoan(&_LendingPoolCore.TransactOpts, _reserve, _availableLiquidityBefore, _income, _protocolFee)
}

// UpdateStateOnFlashLoan is a paid mutator transaction binding the contract method 0x09ac2953.
//
// Solidity: function updateStateOnFlashLoan(address _reserve, uint256 _availableLiquidityBefore, uint256 _income, uint256 _protocolFee) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) UpdateStateOnFlashLoan(_reserve common.Address, _availableLiquidityBefore *big.Int, _income *big.Int, _protocolFee *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.UpdateStateOnFlashLoan(&_LendingPoolCore.TransactOpts, _reserve, _availableLiquidityBefore, _income, _protocolFee)
}

// UpdateStateOnLiquidation is a paid mutator transaction binding the contract method 0x68beb4d6.
//
// Solidity: function updateStateOnLiquidation(address _principalReserve, address _collateralReserve, address _user, uint256 _amountToLiquidate, uint256 _collateralToLiquidate, uint256 _feeLiquidated, uint256 _liquidatedCollateralForFee, uint256 _balanceIncrease, bool _liquidatorReceivesAToken) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) UpdateStateOnLiquidation(opts *bind.TransactOpts, _principalReserve common.Address, _collateralReserve common.Address, _user common.Address, _amountToLiquidate *big.Int, _collateralToLiquidate *big.Int, _feeLiquidated *big.Int, _liquidatedCollateralForFee *big.Int, _balanceIncrease *big.Int, _liquidatorReceivesAToken bool) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "updateStateOnLiquidation", _principalReserve, _collateralReserve, _user, _amountToLiquidate, _collateralToLiquidate, _feeLiquidated, _liquidatedCollateralForFee, _balanceIncrease, _liquidatorReceivesAToken)
}

// UpdateStateOnLiquidation is a paid mutator transaction binding the contract method 0x68beb4d6.
//
// Solidity: function updateStateOnLiquidation(address _principalReserve, address _collateralReserve, address _user, uint256 _amountToLiquidate, uint256 _collateralToLiquidate, uint256 _feeLiquidated, uint256 _liquidatedCollateralForFee, uint256 _balanceIncrease, bool _liquidatorReceivesAToken) returns()
func (_LendingPoolCore *LendingPoolCoreSession) UpdateStateOnLiquidation(_principalReserve common.Address, _collateralReserve common.Address, _user common.Address, _amountToLiquidate *big.Int, _collateralToLiquidate *big.Int, _feeLiquidated *big.Int, _liquidatedCollateralForFee *big.Int, _balanceIncrease *big.Int, _liquidatorReceivesAToken bool) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.UpdateStateOnLiquidation(&_LendingPoolCore.TransactOpts, _principalReserve, _collateralReserve, _user, _amountToLiquidate, _collateralToLiquidate, _feeLiquidated, _liquidatedCollateralForFee, _balanceIncrease, _liquidatorReceivesAToken)
}

// UpdateStateOnLiquidation is a paid mutator transaction binding the contract method 0x68beb4d6.
//
// Solidity: function updateStateOnLiquidation(address _principalReserve, address _collateralReserve, address _user, uint256 _amountToLiquidate, uint256 _collateralToLiquidate, uint256 _feeLiquidated, uint256 _liquidatedCollateralForFee, uint256 _balanceIncrease, bool _liquidatorReceivesAToken) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) UpdateStateOnLiquidation(_principalReserve common.Address, _collateralReserve common.Address, _user common.Address, _amountToLiquidate *big.Int, _collateralToLiquidate *big.Int, _feeLiquidated *big.Int, _liquidatedCollateralForFee *big.Int, _balanceIncrease *big.Int, _liquidatorReceivesAToken bool) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.UpdateStateOnLiquidation(&_LendingPoolCore.TransactOpts, _principalReserve, _collateralReserve, _user, _amountToLiquidate, _collateralToLiquidate, _feeLiquidated, _liquidatedCollateralForFee, _balanceIncrease, _liquidatorReceivesAToken)
}

// UpdateStateOnRebalance is a paid mutator transaction binding the contract method 0xaf825b07.
//
// Solidity: function updateStateOnRebalance(address _reserve, address _user, uint256 _balanceIncrease) returns(uint256)
func (_LendingPoolCore *LendingPoolCoreTransactor) UpdateStateOnRebalance(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _balanceIncrease *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "updateStateOnRebalance", _reserve, _user, _balanceIncrease)
}

// UpdateStateOnRebalance is a paid mutator transaction binding the contract method 0xaf825b07.
//
// Solidity: function updateStateOnRebalance(address _reserve, address _user, uint256 _balanceIncrease) returns(uint256)
func (_LendingPoolCore *LendingPoolCoreSession) UpdateStateOnRebalance(_reserve common.Address, _user common.Address, _balanceIncrease *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.UpdateStateOnRebalance(&_LendingPoolCore.TransactOpts, _reserve, _user, _balanceIncrease)
}

// UpdateStateOnRebalance is a paid mutator transaction binding the contract method 0xaf825b07.
//
// Solidity: function updateStateOnRebalance(address _reserve, address _user, uint256 _balanceIncrease) returns(uint256)
func (_LendingPoolCore *LendingPoolCoreTransactorSession) UpdateStateOnRebalance(_reserve common.Address, _user common.Address, _balanceIncrease *big.Int) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.UpdateStateOnRebalance(&_LendingPoolCore.TransactOpts, _reserve, _user, _balanceIncrease)
}

// UpdateStateOnRedeem is a paid mutator transaction binding the contract method 0xafcdbea3.
//
// Solidity: function updateStateOnRedeem(address _reserve, address _user, uint256 _amountRedeemed, bool _userRedeemedEverything) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) UpdateStateOnRedeem(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _amountRedeemed *big.Int, _userRedeemedEverything bool) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "updateStateOnRedeem", _reserve, _user, _amountRedeemed, _userRedeemedEverything)
}

// UpdateStateOnRedeem is a paid mutator transaction binding the contract method 0xafcdbea3.
//
// Solidity: function updateStateOnRedeem(address _reserve, address _user, uint256 _amountRedeemed, bool _userRedeemedEverything) returns()
func (_LendingPoolCore *LendingPoolCoreSession) UpdateStateOnRedeem(_reserve common.Address, _user common.Address, _amountRedeemed *big.Int, _userRedeemedEverything bool) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.UpdateStateOnRedeem(&_LendingPoolCore.TransactOpts, _reserve, _user, _amountRedeemed, _userRedeemedEverything)
}

// UpdateStateOnRedeem is a paid mutator transaction binding the contract method 0xafcdbea3.
//
// Solidity: function updateStateOnRedeem(address _reserve, address _user, uint256 _amountRedeemed, bool _userRedeemedEverything) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) UpdateStateOnRedeem(_reserve common.Address, _user common.Address, _amountRedeemed *big.Int, _userRedeemedEverything bool) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.UpdateStateOnRedeem(&_LendingPoolCore.TransactOpts, _reserve, _user, _amountRedeemed, _userRedeemedEverything)
}

// UpdateStateOnRepay is a paid mutator transaction binding the contract method 0xda12d96f.
//
// Solidity: function updateStateOnRepay(address _reserve, address _user, uint256 _paybackAmountMinusFees, uint256 _originationFeeRepaid, uint256 _balanceIncrease, bool _repaidWholeLoan) returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) UpdateStateOnRepay(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _paybackAmountMinusFees *big.Int, _originationFeeRepaid *big.Int, _balanceIncrease *big.Int, _repaidWholeLoan bool) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "updateStateOnRepay", _reserve, _user, _paybackAmountMinusFees, _originationFeeRepaid, _balanceIncrease, _repaidWholeLoan)
}

// UpdateStateOnRepay is a paid mutator transaction binding the contract method 0xda12d96f.
//
// Solidity: function updateStateOnRepay(address _reserve, address _user, uint256 _paybackAmountMinusFees, uint256 _originationFeeRepaid, uint256 _balanceIncrease, bool _repaidWholeLoan) returns()
func (_LendingPoolCore *LendingPoolCoreSession) UpdateStateOnRepay(_reserve common.Address, _user common.Address, _paybackAmountMinusFees *big.Int, _originationFeeRepaid *big.Int, _balanceIncrease *big.Int, _repaidWholeLoan bool) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.UpdateStateOnRepay(&_LendingPoolCore.TransactOpts, _reserve, _user, _paybackAmountMinusFees, _originationFeeRepaid, _balanceIncrease, _repaidWholeLoan)
}

// UpdateStateOnRepay is a paid mutator transaction binding the contract method 0xda12d96f.
//
// Solidity: function updateStateOnRepay(address _reserve, address _user, uint256 _paybackAmountMinusFees, uint256 _originationFeeRepaid, uint256 _balanceIncrease, bool _repaidWholeLoan) returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) UpdateStateOnRepay(_reserve common.Address, _user common.Address, _paybackAmountMinusFees *big.Int, _originationFeeRepaid *big.Int, _balanceIncrease *big.Int, _repaidWholeLoan bool) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.UpdateStateOnRepay(&_LendingPoolCore.TransactOpts, _reserve, _user, _paybackAmountMinusFees, _originationFeeRepaid, _balanceIncrease, _repaidWholeLoan)
}

// UpdateStateOnSwapRate is a paid mutator transaction binding the contract method 0xf6148311.
//
// Solidity: function updateStateOnSwapRate(address _reserve, address _user, uint256 _principalBorrowBalance, uint256 _compoundedBorrowBalance, uint256 _balanceIncrease, uint8 _currentRateMode) returns(uint8, uint256)
func (_LendingPoolCore *LendingPoolCoreTransactor) UpdateStateOnSwapRate(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _principalBorrowBalance *big.Int, _compoundedBorrowBalance *big.Int, _balanceIncrease *big.Int, _currentRateMode uint8) (*types.Transaction, error) {
	return _LendingPoolCore.contract.Transact(opts, "updateStateOnSwapRate", _reserve, _user, _principalBorrowBalance, _compoundedBorrowBalance, _balanceIncrease, _currentRateMode)
}

// UpdateStateOnSwapRate is a paid mutator transaction binding the contract method 0xf6148311.
//
// Solidity: function updateStateOnSwapRate(address _reserve, address _user, uint256 _principalBorrowBalance, uint256 _compoundedBorrowBalance, uint256 _balanceIncrease, uint8 _currentRateMode) returns(uint8, uint256)
func (_LendingPoolCore *LendingPoolCoreSession) UpdateStateOnSwapRate(_reserve common.Address, _user common.Address, _principalBorrowBalance *big.Int, _compoundedBorrowBalance *big.Int, _balanceIncrease *big.Int, _currentRateMode uint8) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.UpdateStateOnSwapRate(&_LendingPoolCore.TransactOpts, _reserve, _user, _principalBorrowBalance, _compoundedBorrowBalance, _balanceIncrease, _currentRateMode)
}

// UpdateStateOnSwapRate is a paid mutator transaction binding the contract method 0xf6148311.
//
// Solidity: function updateStateOnSwapRate(address _reserve, address _user, uint256 _principalBorrowBalance, uint256 _compoundedBorrowBalance, uint256 _balanceIncrease, uint8 _currentRateMode) returns(uint8, uint256)
func (_LendingPoolCore *LendingPoolCoreTransactorSession) UpdateStateOnSwapRate(_reserve common.Address, _user common.Address, _principalBorrowBalance *big.Int, _compoundedBorrowBalance *big.Int, _balanceIncrease *big.Int, _currentRateMode uint8) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.UpdateStateOnSwapRate(&_LendingPoolCore.TransactOpts, _reserve, _user, _principalBorrowBalance, _compoundedBorrowBalance, _balanceIncrease, _currentRateMode)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_LendingPoolCore *LendingPoolCoreTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _LendingPoolCore.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_LendingPoolCore *LendingPoolCoreSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.Fallback(&_LendingPoolCore.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_LendingPoolCore *LendingPoolCoreTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _LendingPoolCore.Contract.Fallback(&_LendingPoolCore.TransactOpts, calldata)
}

// LendingPoolCoreReserveUpdatedIterator is returned from FilterReserveUpdated and is used to iterate over the raw logs and unpacked data for ReserveUpdated events raised by the LendingPoolCore contract.
type LendingPoolCoreReserveUpdatedIterator struct {
	Event *LendingPoolCoreReserveUpdated // Event containing the contract specifics and raw log

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
func (it *LendingPoolCoreReserveUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolCoreReserveUpdated)
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
		it.Event = new(LendingPoolCoreReserveUpdated)
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
func (it *LendingPoolCoreReserveUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolCoreReserveUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolCoreReserveUpdated represents a ReserveUpdated event raised by the LendingPoolCore contract.
type LendingPoolCoreReserveUpdated struct {
	Reserve             common.Address
	LiquidityRate       *big.Int
	StableBorrowRate    *big.Int
	VariableBorrowRate  *big.Int
	LiquidityIndex      *big.Int
	VariableBorrowIndex *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReserveUpdated is a free log retrieval operation binding the contract event 0x04e4f521f16fcfd987978b05504262c2a2db223844977aab000e5accedb2d725.
//
// Solidity: event ReserveUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_LendingPoolCore *LendingPoolCoreFilterer) FilterReserveUpdated(opts *bind.FilterOpts, reserve []common.Address) (*LendingPoolCoreReserveUpdatedIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _LendingPoolCore.contract.FilterLogs(opts, "ReserveUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolCoreReserveUpdatedIterator{contract: _LendingPoolCore.contract, event: "ReserveUpdated", logs: logs, sub: sub}, nil
}

// WatchReserveUpdated is a free log subscription operation binding the contract event 0x04e4f521f16fcfd987978b05504262c2a2db223844977aab000e5accedb2d725.
//
// Solidity: event ReserveUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_LendingPoolCore *LendingPoolCoreFilterer) WatchReserveUpdated(opts *bind.WatchOpts, sink chan<- *LendingPoolCoreReserveUpdated, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _LendingPoolCore.contract.WatchLogs(opts, "ReserveUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolCoreReserveUpdated)
				if err := _LendingPoolCore.contract.UnpackLog(event, "ReserveUpdated", log); err != nil {
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

// ParseReserveUpdated is a log parse operation binding the contract event 0x04e4f521f16fcfd987978b05504262c2a2db223844977aab000e5accedb2d725.
//
// Solidity: event ReserveUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_LendingPoolCore *LendingPoolCoreFilterer) ParseReserveUpdated(log types.Log) (*LendingPoolCoreReserveUpdated, error) {
	event := new(LendingPoolCoreReserveUpdated)
	if err := _LendingPoolCore.contract.UnpackLog(event, "ReserveUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
