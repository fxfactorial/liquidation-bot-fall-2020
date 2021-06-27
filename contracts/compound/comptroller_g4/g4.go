// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package comptroller_g4

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

// ComptrollerG4ABI is the input ABI used to generate the binding from.
const ComptrollerG4ABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSpeed\",\"type\":\"uint256\"}],\"name\":\"CompSpeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedBorrowerComp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compSupplyIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedSupplierComp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isComped\",\"type\":\"bool\"}],\"name\":\"MarketComped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketExited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"MarketListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCloseFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCloseFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCollateralFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCollateralFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCompRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCompRate\",\"type\":\"uint256\"}],\"name\":\"NewCompRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLiquidationIncentiveMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"NewLiquidationIncentive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMaxAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxAssets\",\"type\":\"uint256\"}],\"name\":\"NewMaxAssets\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPauseGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"NewPauseGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"oldPriceOracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"NewPriceOracle\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountAssets\",\"outputs\":[{\"internalType\":\"contractCToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"contractCToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closeFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compBorrowState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compBorrowerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compClaimThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compInitialIndex\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compSupplierIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compSupplyState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isComptroller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"liquidationIncentiveMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isComped\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pauseGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingComptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seizeGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transferGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAssetsIn\",\"outputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"checkMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cTokens\",\"type\":\"address[]\"}],\"name\":\"enterMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenAddress\",\"type\":\"address\"}],\"name\":\"exitMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualMintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"mintVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowerIndex\",\"type\":\"uint256\"}],\"name\":\"repayBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenModify\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"getHypotheticalAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateCalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"_setPriceOracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCloseFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCollateralFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxAssets\",\"type\":\"uint256\"}],\"name\":\"_setMaxAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"_setLiquidationIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"_supportMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"_setPauseGuardian\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setMintPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setBorrowPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setTransferPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setSeizePaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractUnitroller\",\"name\":\"unitroller\",\"type\":\"address\"}],\"name\":\"_become\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refreshCompSpeeds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"}],\"name\":\"claimComp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"holders\",\"type\":\"address[]\"},{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"borrowers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"suppliers\",\"type\":\"bool\"}],\"name\":\"claimComp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"claimComp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"compRate_\",\"type\":\"uint256\"}],\"name\":\"_setCompRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cTokens\",\"type\":\"address[]\"}],\"name\":\"_addCompMarkets\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"_dropCompMarket\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllMarkets\",\"outputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCompAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ComptrollerG4Bin is the compiled bytecode used for deploying new contracts.
var ComptrollerG4Bin = "0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055615bfe80620000336000396000f3fe608060405234801561001057600080fd5b50600436106104565760003560e01c8063731f0c2b11610250578063bb82aa5e11610150578063da3d454c116100c8578063e875544611610097578063eabe7d911161007c578063eabe7d911461112c578063ede4edd014611162578063f851a4401461118857610456565b8063e8755446146110fe578063e9af02921461110657610456565b8063da3d454c14611068578063dce154491461109e578063dcfbc0c7146110ca578063e4028eee146110d257610456565b8063ca0af0431161011f578063ce485c5e11610104578063ce485c5e14610f62578063d02f735114611005578063d9226ced1461104b57610456565b8063ca0af04314610f0e578063cc7ebdc414610f3c57610456565b8063bb82aa5e14610dd8578063bdcdc25814610de0578063c299823814610e1c578063c488847b14610ebf57610456565b806394b2294b116101e3578063aa900754116101b2578063ac0b0bb711610197578063ac0b0bb714610d9a578063b0772d0b14610da2578063b21be7fd14610daa57610456565b8063aa90075414610d1c578063abfceffc14610d2457610456565b806394b2294b14610cc25780639d1b5a0a14610cca578063a76b3fda14610cd2578063a7f0e23114610cf857610456565b80638c57804e1161021f5780638c57804e14610c075780638e8f294b14610c2d5780638ebf636414610c75578063929fe9a114610c9457610456565b8063731f0c2b14610bc9578063747026c914610bef5780637dc0d1d014610bf757806387f7630314610bff57610456565b80634d8e50371161035b5780635ec88c79116102ee5780636a491112116102bd5780636b79c38d116102a25780636b79c38d14610b0f5780636d154ea514610b5d5780636d35bf9114610b8357610456565b80636a49111214610ab65780636a56947e14610ad357610456565b80635ec88c79146108f45780635f5af1aa1461091a5780635fc7e71e146109405780636810dfa61461098657610456565b806351dff9891161032a57806351dff9891461083f57806352d84d1e1461087b57806355ee1fe1146108985780635c778605146108be57610456565b80634d8e50371461078a5780634e79238f146107925780634ef4c3e1146107ec5780634fd42e171461082257610456565b806326782247116103ee5780633bcf7ec1116103bd57806342cbb15c116103a257806342cbb15c1461072e57806347ef3b3b146107365780634ada90af1461078257610456565b80633bcf7ec1146106c457806341c728b9146106f257610456565b8063267822471461065a5780632d70db7814610662578063317b0b77146106815780633aa729b41461069e57610456565b80631d7b33d71161042a5780631d7b33d7146105805780631ededc91146105b857806324008a62146105fa57806324a3d6221461063657610456565b80627e3dd21461045b57806318c882a5146104775780631c3db2e0146104a55780631d504dc61461055a575b600080fd5b610463611190565b604080519115158252519081900360200190f35b6104636004803603604081101561048d57600080fd5b506001600160a01b0381351690602001351515611195565b610558600480360360408110156104bb57600080fd5b6001600160a01b0382351691908101906040810160208201356401000000008111156104e657600080fd5b8201836020820111156104f857600080fd5b8035906020019184602083028401116401000000008311171561051a57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611353945050505050565b005b6105586004803603602081101561057057600080fd5b50356001600160a01b03166113b5565b6105a66004803603602081101561059657600080fd5b50356001600160a01b031661151c565b60408051918252519081900360200190f35b610558600480360360a08110156105ce57600080fd5b506001600160a01b0381358116916020810135821691604082013516906060810135906080013561152e565b6105a66004803603608081101561061057600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135611535565b61063e6115fe565b604080516001600160a01b039092168252519081900360200190f35b61063e61160d565b6104636004803603602081101561067857600080fd5b5035151561161c565b6105a66004803603602081101561069757600080fd5b5035611790565b610558600480360360208110156106b457600080fd5b50356001600160a01b03166118a1565b610463600480360360408110156106da57600080fd5b506001600160a01b03813516906020013515156119d2565b6105586004803603608081101561070857600080fd5b506001600160a01b03813581169160208101359091169060408101359060600135611b8d565b6105a6611b93565b610558600480360360c081101561074c57600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101359091169060808101359060a00135611b98565b6105a6611ba0565b610558611ba6565b6107ce600480360360808110156107a857600080fd5b506001600160a01b03813581169160208101359091169060408101359060600135611bee565b60408051938452602084019290925282820152519081900360600190f35b6105a66004803603606081101561080257600080fd5b506001600160a01b03813581169160208101359091169060400135611c28565b6105a66004803603602081101561083857600080fd5b5035611ce2565b6105586004803603608081101561085557600080fd5b506001600160a01b03813581169160208101359091169060408101359060600135611dd6565b61063e6004803603602081101561089157600080fd5b5035611e36565b6105a6600480360360208110156108ae57600080fd5b50356001600160a01b0316611e5d565b610558600480360360608110156108d457600080fd5b506001600160a01b03813581169160208101359091169060400135611ee4565b6107ce6004803603602081101561090a57600080fd5b50356001600160a01b0316611ee9565b6105a66004803603602081101561093057600080fd5b50356001600160a01b0316611f1e565b6105a6600480360360a081101561095657600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101359091169060800135611fa2565b6105586004803603608081101561099c57600080fd5b8101906020810181356401000000008111156109b757600080fd5b8201836020820111156109c957600080fd5b803590602001918460208302840111640100000000831117156109eb57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050640100000000811115610a3b57600080fd5b820183602082011115610a4d57600080fd5b80359060200191846020830284011164010000000083111715610a6f57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550505050803515159150602001351515612129565b61055860048036036020811015610acc57600080fd5b50356122da565b61055860048036036080811015610ae957600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135611b8d565b610b3560048036036020811015610b2557600080fd5b50356001600160a01b031661237e565b604080516001600160e01b03909316835263ffffffff90911660208301528051918290030190f35b61046360048036036020811015610b7357600080fd5b50356001600160a01b03166123a8565b610558600480360360a0811015610b9957600080fd5b506001600160a01b038135811691602081013582169160408201358116916060810135909116906080013561152e565b61046360048036036020811015610bdf57600080fd5b50356001600160a01b03166123bd565b6105a66123d2565b61063e6123dd565b6104636123ec565b610b3560048036036020811015610c1d57600080fd5b50356001600160a01b03166123fc565b610c5360048036036020811015610c4357600080fd5b50356001600160a01b0316612426565b6040805193151584526020840192909252151582820152519081900360600190f35b61046360048036036020811015610c8b57600080fd5b5035151561244c565b61046360048036036040811015610caa57600080fd5b506001600160a01b03813581169160200135166125bc565b6105a66125ef565b61063e6125f5565b6105a660048036036020811015610ce857600080fd5b50356001600160a01b031661260d565b610d0061276a565b604080516001600160e01b039092168252519081900360200190f35b6105a661277d565b610d4a60048036036020811015610d3a57600080fd5b50356001600160a01b0316612783565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610d86578181015183820152602001610d6e565b505050509050019250505060405180910390f35b61046361280c565b610d4a61281c565b6105a660048036036040811015610dc057600080fd5b506001600160a01b038135811691602001351661287e565b61063e61289b565b6105a660048036036080811015610df657600080fd5b506001600160a01b038135811691602081013582169160408201351690606001356128aa565b610d4a60048036036020811015610e3257600080fd5b810190602081018135640100000000811115610e4d57600080fd5b820183602082011115610e5f57600080fd5b80359060200191846020830284011164010000000083111715610e8157600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550612949945050505050565b610ef560048036036060811015610ed557600080fd5b506001600160a01b038135811691602081013590911690604001356129e0565b6040805192835260208301919091528051918290030190f35b6105a660048036036040811015610f2457600080fd5b506001600160a01b0381358116916020013516612c55565b6105a660048036036020811015610f5257600080fd5b50356001600160a01b0316612c72565b61055860048036036020811015610f7857600080fd5b810190602081018135640100000000811115610f9357600080fd5b820183602082011115610fa557600080fd5b80359060200191846020830284011164010000000083111715610fc757600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550612c84945050505050565b6105a6600480360360a081101561101b57600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101359091169060800135612d16565b6105a66004803603602081101561106157600080fd5b5035612ef5565b6105a66004803603606081101561107e57600080fd5b506001600160a01b03813581169160208101359091169060400135612f5e565b61063e600480360360408110156110b457600080fd5b506001600160a01b038135169060200135613260565b61063e613295565b6105a6600480360360408110156110e857600080fd5b506001600160a01b0381351690602001356132a4565b6105a6613454565b6105586004803603602081101561111c57600080fd5b50356001600160a01b031661345a565b6105a66004803603606081101561114257600080fd5b506001600160a01b038135811691602081013590911690604001356134be565b6105a66004803603602081101561117857600080fd5b50356001600160a01b03166134fb565b61063e61380e565b600181565b6001600160a01b03821660009081526009602052604081205460ff166111ec5760405162461bcd60e51b8152600401808060200182810382526028815260200180615afe6028913960400191505060405180910390fd5b600a546001600160a01b031633148061120f57506000546001600160a01b031633145b61124a5760405162461bcd60e51b8152600401808060200182810382526027815260200180615b576027913960400191505060405180910390fd5b6000546001600160a01b031633148061126557506001821515145b6112b6576040805162461bcd60e51b815260206004820152601660248201527f6f6e6c792061646d696e2063616e20756e706175736500000000000000000000604482015290519081900360640190fd5b6001600160a01b0383166000818152600c6020908152604091829020805486151560ff19909116811790915582519384528383015260609083018190526006908301527f426f72726f7700000000000000000000000000000000000000000000000000006080830152517f71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b09181900360a00190a150805b92915050565b60408051600180825281830190925260609160208083019080388339019050509050828160008151811061138357fe5b60200260200101906001600160a01b031690816001600160a01b0316815250506113b08183600180612129565b505050565b806001600160a01b031663f851a4406040518163ffffffff1660e01b815260040160206040518083038186803b1580156113ee57600080fd5b505afa158015611402573d6000803e3d6000fd5b505050506040513d602081101561141857600080fd5b50516001600160a01b031633146114605760405162461bcd60e51b8152600401808060200182810382526027815260200180615ba36027913960400191505060405180910390fd5b806001600160a01b031663c1e803346040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561149b57600080fd5b505af11580156114af573d6000803e3d6000fd5b505050506040513d60208110156114c557600080fd5b505115611519576040805162461bcd60e51b815260206004820152601560248201527f6368616e6765206e6f7420617574686f72697a65640000000000000000000000604482015290519081900360640190fd5b50565b600f6020526000908152604090205481565b5050505050565b6001600160a01b03841660009081526009602052604081205460ff1661155d575060096115f6565b611565615a3e565b6040518060200160405280876001600160a01b031663aa5af0fd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156115a957600080fd5b505afa1580156115bd573d6000803e3d6000fd5b505050506040513d60208110156115d357600080fd5b5051905290506115e3868261381d565b6115f08685836000613ac9565b60009150505b949350505050565b600a546001600160a01b031681565b6001546001600160a01b031681565b600a546000906001600160a01b031633148061164257506000546001600160a01b031633145b61167d5760405162461bcd60e51b8152600401808060200182810382526027815260200180615b576027913960400191505060405180910390fd5b6000546001600160a01b031633148061169857506001821515145b6116e9576040805162461bcd60e51b815260206004820152601660248201527f6f6e6c792061646d696e2063616e20756e706175736500000000000000000000604482015290519081900360640190fd5b600a8054831515600160b81b81027fffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffff909216919091179091556040805160208101929092528082526005828201527f5365697a650000000000000000000000000000000000000000000000000000006060830152517fef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de09181900360800190a150805b919050565b600080546001600160a01b031633146117b6576117af60016004613cb4565b905061178b565b6117be615a3e565b5060408051602081019091528281526117d5615a3e565b50604080516020810190915266b1a2bc2ec5000081526117f58282613d1a565b1561180e57611805600580613cb4565b9250505061178b565b611816615a3e565b506040805160208101909152670c7d713b49da000081526118378184613d22565b1561185157611847600580613cb4565b935050505061178b565b6005805490869055604080518281526020810188905281517f3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9929181900390910190a160005b9695505050505050565b6000546001600160a01b03163314611900576040805162461bcd60e51b815260206004820152601f60248201527f6f6e6c792061646d696e2063616e2064726f7020636f6d70206d61726b657400604482015290519081900360640190fd5b6001600160a01b0381166000908152600960205260409020600381015460ff161515600114611976576040805162461bcd60e51b815260206004820152601b60248201527f6d61726b6574206973206e6f74206120636f6d70206d61726b65740000000000604482015290519081900360640190fd5b60038101805460ff19169055604080516001600160a01b03841681526000602082015281517f93c1f3e36ed71139f466a4ce8c9751790e2e33f5afb2df0dcfb3aeabe55d5aa2929181900390910190a16119ce613d29565b5050565b6001600160a01b03821660009081526009602052604081205460ff16611a295760405162461bcd60e51b8152600401808060200182810382526028815260200180615afe6028913960400191505060405180910390fd5b600a546001600160a01b0316331480611a4c57506000546001600160a01b031633145b611a875760405162461bcd60e51b8152600401808060200182810382526027815260200180615b576027913960400191505060405180910390fd5b6000546001600160a01b0316331480611aa257506001821515145b611af3576040805162461bcd60e51b815260206004820152601660248201527f6f6e6c792061646d696e2063616e20756e706175736500000000000000000000604482015290519081900360640190fd5b6001600160a01b0383166000818152600b6020908152604091829020805486151560ff19909116811790915582519384528383015260609083018190526004908301527f4d696e74000000000000000000000000000000000000000000000000000000006080830152517f71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b09181900360a00190a150919050565b50505050565b435b90565b505050505050565b60065481565b333214611be45760405162461bcd60e51b8152600401808060200182810382526031815260200180615b266031913960400191505060405180910390fd5b611bec613d29565b565b600080600080600080611c038a8a8a8a6140ee565b925092509250826011811115611c1557fe5b95509093509150505b9450945094915050565b6001600160a01b0383166000908152600b602052604081205460ff1615611c96576040805162461bcd60e51b815260206004820152600e60248201527f6d696e7420697320706175736564000000000000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b03841660009081526009602052604090205460ff16611cc05760095b9050611cdb565b611cc984614509565b611cd5848460006147ab565b60005b90505b9392505050565b600080546001600160a01b03163314611d01576117af6001600b613cb4565b611d09615a3e565b506040805160208101909152828152611d20615a3e565b506040805160208101909152670de0b6b3a76400008152611d418282613d22565b15611d52576118056007600c613cb4565b611d5a615a3e565b5060408051602081019091526714d1120d7b1600008152611d7b8184613d22565b15611d8c576118476007600c613cb4565b6006805490869055604080518281526020810188905281517faeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316929181900390910190a16000611897565b80158015611de45750600082115b15611b8d576040805162461bcd60e51b815260206004820152601160248201527f72656465656d546f6b656e73207a65726f000000000000000000000000000000604482015290519081900360640190fd5b600d8181548110611e4357fe5b6000918252602090912001546001600160a01b0316905081565b600080546001600160a01b03163314611e7c576117af60016010613cb4565b600480546001600160a01b038481166001600160a01b0319831681179093556040805191909216808252602082019390935281517fd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22929181900390910190a160009392505050565b6113b0565b600080600080600080611f008760008060006140ee565b925092509250826011811115611f1257fe5b97919650945092505050565b600080546001600160a01b03163314611f3d576117af60016013613cb4565b600a80546001600160a01b038481166001600160a01b0319831617928390556040805192821680845293909116602083015280517f0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e9281900390910190a16000611cdb565b6001600160a01b03851660009081526009602052604081205460ff161580611fe357506001600160a01b03851660009081526009602052604090205460ff16155b15611ff25760095b9050612120565b600080611ffe856149a3565b9193509091506000905082601181111561201457fe5b1461202e5781601181111561202557fe5b92505050612120565b8061203a576003612025565b6000886001600160a01b03166395dd9193876040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561209257600080fd5b505afa1580156120a6573d6000803e3d6000fd5b505050506040513d60208110156120bc57600080fd5b50516040805160208101909152600554815290915060009081906120e090846149c3565b909250905060008260038111156120f357fe5b1461210757600b5b95505050505050612120565b808711156121165760116120fb565b6000955050505050505b95945050505050565b60005b835181101561152e57600084828151811061214357fe5b6020908102919091018101516001600160a01b0381166000908152600990925260409091205490915060ff166121c0576040805162461bcd60e51b815260206004820152601560248201527f6d61726b6574206d757374206265206c69737465640000000000000000000000604482015290519081900360640190fd5b60018415151415612288576121d3615a3e565b6040518060200160405280836001600160a01b031663aa5af0fd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561221757600080fd5b505afa15801561222b573d6000803e3d6000fd5b505050506040513d602081101561224157600080fd5b505190529050612251828261381d565b60005b87518110156122855761227d8389838151811061226d57fe5b6020026020010151846001613ac9565b600101612254565b50505b600183151514156122d15761229c81614509565b60005b86518110156122cf576122c7828883815181106122b857fe5b602002602001015160016147ab565b60010161229f565b505b5060010161212c565b6122e2614a17565b612333576040805162461bcd60e51b815260206004820152601f60248201527f6f6e6c792061646d696e2063616e206368616e676520636f6d70207261746500604482015290519081900360640190fd5b600e805490829055604080518281526020810184905281517fc227c9272633c3a307d9845bf2bc2509cefb20d655b5f3c1002d8e1e3f22c8b0929181900390910190a16119ce613d29565b6010602052600090815260409020546001600160e01b03811690600160e01b900463ffffffff1682565b600c6020526000908152604090205460ff1681565b600b6020526000908152604090205460ff1681565b66038d7ea4c6800081565b6004546001600160a01b031681565b600a54600160b01b900460ff1681565b6011602052600090815260409020546001600160e01b03811690600160e01b900463ffffffff1682565b60096020526000908152604090208054600182015460039092015460ff91821692911683565b600a546000906001600160a01b031633148061247257506000546001600160a01b031633145b6124ad5760405162461bcd60e51b8152600401808060200182810382526027815260200180615b576027913960400191505060405180910390fd5b6000546001600160a01b03163314806124c857506001821515145b612519576040805162461bcd60e51b815260206004820152601660248201527f6f6e6c792061646d696e2063616e20756e706175736500000000000000000000604482015290519081900360640190fd5b600a8054831515600160b01b81027fffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffff909216919091179091556040805160208101929092528082526008828201527f5472616e736665720000000000000000000000000000000000000000000000006060830152517fef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de09181900360800190a15090565b6001600160a01b038082166000908152600960209081526040808320938616835260029093019052205460ff1692915050565b60075481565b73c00e94cb662c3520282e6f5717214004a7f2688890565b600080546001600160a01b0316331461262c576117af60016012613cb4565b6001600160a01b03821660009081526009602052604090205460ff1615612659576117af600a6011613cb4565b816001600160a01b031663fe9c44ae6040518163ffffffff1660e01b815260040160206040518083038186803b15801561269257600080fd5b505afa1580156126a6573d6000803e3d6000fd5b505050506040513d60208110156126bc57600080fd5b5050604080516060810182526001808252600060208381018281528486018381526001600160a01b03891684526009909252949091209251835490151560ff1991821617845593519183019190915551600390910180549115159190921617905561272682614a40565b604080516001600160a01b038416815290517fcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f9181900360200190a1600092915050565b6ec097ce7bc90715b34b9f100000000081565b600e5481565b60608060086000846001600160a01b03166001600160a01b031681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156127ff57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116127e1575b5093979650505050505050565b600a54600160b81b900460ff1681565b6060600d80548060200260200160405190810160405280929190818152602001828054801561287457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612856575b5050505050905090565b601260209081526000928352604080842090915290825290205481565b6002546001600160a01b031681565b600a54600090600160b01b900460ff161561290c576040805162461bcd60e51b815260206004820152601260248201527f7472616e73666572206973207061757365640000000000000000000000000000604482015290519081900360640190fd5b6000612919868685614b27565b905080156129285790506115f6565b61293186614509565b61293d868660006147ab565b6115f0868560006147ab565b606060008251905060608160405190808252806020026020018201604052801561297d578160200160208202803883390190505b50905060005b828110156129d857600085828151811061299957fe5b602002602001015190506129ad8133614bca565b60118111156129b857fe5b8383815181106129c457fe5b602090810291909101015250600101612983565b509392505050565b600480546040805163fc57d4df60e01b81526001600160a01b038781169482019490945290516000938493849391169163fc57d4df91602480820192602092909190829003018186803b158015612a3657600080fd5b505afa158015612a4a573d6000803e3d6000fd5b505050506040513d6020811015612a6057600080fd5b5051600480546040805163fc57d4df60e01b81526001600160a01b038a8116948201949094529051939450600093929091169163fc57d4df91602480820192602092909190829003018186803b158015612ab957600080fd5b505afa158015612acd573d6000803e3d6000fd5b505050506040513d6020811015612ae357600080fd5b50519050811580612af2575080155b15612b0757600d935060009250612c4d915050565b6000866001600160a01b031663182df0f56040518163ffffffff1660e01b815260040160206040518083038186803b158015612b4257600080fd5b505afa158015612b56573d6000803e3d6000fd5b505050506040513d6020811015612b6c57600080fd5b505190506000612b7a615a3e565b612b82615a3e565b612b8a615a3e565b6000612b9860065489614ceb565b945090506000816003811115612baa57fe5b14612bc657600b5b995060009850612c4d975050505050505050565b612bd08787614ceb565b935090506000816003811115612be257fe5b14612bee57600b612bb2565b612bf88484614d26565b925090506000816003811115612c0a57fe5b14612c1657600b612bb2565b612c20828c6149c3565b955090506000816003811115612c3257fe5b14612c3e57600b612bb2565b60009950939750505050505050505b935093915050565b601360209081526000928352604080842090915290825290205481565b60146020526000908152604090205481565b612c8c614a17565b612cdd576040805162461bcd60e51b815260206004820152601e60248201527f6f6e6c792061646d696e2063616e2061646420636f6d70206d61726b65740000604482015290519081900360640190fd5b60005b8151811015612d0d57612d05828281518110612cf857fe5b6020026020010151614d3e565b600101612ce0565b50611519613d29565b600a54600090600160b81b900460ff1615612d78576040805162461bcd60e51b815260206004820152600f60248201527f7365697a65206973207061757365640000000000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b03861660009081526009602052604090205460ff161580612db957506001600160a01b03851660009081526009602052604090205460ff16155b15612dc5576009611feb565b846001600160a01b0316635fe3b5676040518163ffffffff1660e01b815260040160206040518083038186803b158015612dfe57600080fd5b505afa158015612e12573d6000803e3d6000fd5b505050506040513d6020811015612e2857600080fd5b5051604080517f5fe3b56700000000000000000000000000000000000000000000000000000000815290516001600160a01b0392831692891691635fe3b567916004808301926020929190829003018186803b158015612e8757600080fd5b505afa158015612e9b573d6000803e3d6000fd5b505050506040513d6020811015612eb157600080fd5b50516001600160a01b031614612ec8576002611feb565b612ed186614509565b612edd868460006147ab565b612ee9868560006147ab565b60009695505050505050565b600080546001600160a01b03163314612f14576117af6001600d613cb4565b6007805490839055604080518281526020810185905281517f7093cf1eb653f749c3ff531d6df7f92764536a7fa0d13530cd26e070780c32ea929181900390910190a16000611cdb565b6001600160a01b0383166000908152600c602052604081205460ff1615612fcc576040805162461bcd60e51b815260206004820152601060248201527f626f72726f772069732070617573656400000000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b03841660009081526009602052604090205460ff16612ff3576009611cb9565b6001600160a01b038085166000908152600960209081526040808320938716835260029093019052205460ff166130eb57336001600160a01b03851614613081576040805162461bcd60e51b815260206004820152601560248201527f73656e646572206d7573742062652063546f6b656e0000000000000000000000604482015290519081900360640190fd5b600061308d3385614bca565b9050600081601181111561309d57fe5b146130b6578060118111156130ae57fe5b915050611cdb565b6001600160a01b038086166000908152600960209081526040808320938816835260029093019052205460ff166130e957fe5b505b600480546040805163fc57d4df60e01b81526001600160a01b03888116948201949094529051929091169163fc57d4df91602480820192602092909190829003018186803b15801561313c57600080fd5b505afa158015613150573d6000803e3d6000fd5b505050506040513d602081101561316657600080fd5b505161317357600d611cb9565b60008061318385876000876140ee565b9193509091506000905082601181111561319957fe5b146131b3578160118111156131aa57fe5b92505050611cdb565b80156131c05760046131aa565b6131c8615a3e565b6040518060200160405280886001600160a01b031663aa5af0fd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561320c57600080fd5b505afa158015613220573d6000803e3d6000fd5b505050506040513d602081101561323657600080fd5b505190529050613246878261381d565b6132538787836000613ac9565b6000979650505050505050565b6008602052816000526040600020818154811061327957fe5b6000918252602090912001546001600160a01b03169150829050565b6003546001600160a01b031681565b600080546001600160a01b031633146132ca576132c360016006613cb4565b905061134d565b6001600160a01b0383166000908152600960205260409020805460ff166132ff576132f760096007613cb4565b91505061134d565b613307615a3e565b50604080516020810190915283815261331e615a3e565b506040805160208101909152670c7d713b49da0000815261333f8183613d22565b1561335a5761335060066008613cb4565b935050505061134d565b84158015906133e35750600480546040805163fc57d4df60e01b81526001600160a01b038a8116948201949094529051929091169163fc57d4df91602480820192602092909190829003018186803b1580156133b557600080fd5b505afa1580156133c9573d6000803e3d6000fd5b505050506040513d60208110156133df57600080fd5b5051155b156133f457613350600d6009613cb4565b60018301805490869055604080516001600160a01b03891681526020810183905280820188905290517f70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc59181900360600190a16000979650505050505050565b60055481565b61151981600d8054806020026020016040519081016040528092919081815260200182805480156134b457602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613496575b5050505050611353565b6000806134cc858585614b27565b905080156134db579050611cdb565b6134e485614509565b6134f0858560006147ab565b600095945050505050565b6000808290506000806000836001600160a01b031663c37f68e2336040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060806040518083038186803b15801561355c57600080fd5b505afa158015613570573d6000803e3d6000fd5b505050506040513d608081101561358657600080fd5b5080516020820151604090920151909450909250905082156135d95760405162461bcd60e51b8152600401808060200182810382526025815260200180615b7e6025913960400191505060405180910390fd5b80156135f6576135eb600c6002613cb4565b94505050505061178b565b6000613603873385614b27565b9050801561362457613618600e60038361506b565b9550505050505061178b565b6001600160a01b0385166000908152600960209081526040808320338452600281019092529091205460ff16613663576000965050505050505061178b565b3360009081526002820160209081526040808320805460ff1916905560088252918290208054835181840281018401909452808452606093928301828280156136d557602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116136b7575b5050835193945083925060009150505b8281101561372a57896001600160a01b031684828151811061370357fe5b60200260200101516001600160a01b031614156137225780915061372a565b6001016136e5565b5081811061373457fe5b33600090815260086020526040902080548190600019810190811061375557fe5b9060005260206000200160009054906101000a90046001600160a01b031681838154811061377f57fe5b600091825260209091200180546001600160a01b0319166001600160a01b039290921691909117905580546137b8826000198301615a51565b50604080516001600160a01b038c16815233602082015281517fe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d929181900390910190a160009c9b505050505050505050505050565b6000546001600160a01b031681565b6001600160a01b0382166000908152601160209081526040808320600f909252822054909161384a611b93565b835490915060009061386a908390600160e01b900463ffffffff166150d1565b905060008111801561387c5750600083115b15613a5d5760006138f1876001600160a01b03166347bd37186040518163ffffffff1660e01b815260040160206040518083038186803b1580156138bf57600080fd5b505afa1580156138d3573d6000803e3d6000fd5b505050506040513d60208110156138e957600080fd5b505187615113565b905060006138ff8386615131565b9050613909615a3e565b600083116139265760405180602001604052806000815250613930565b6139308284615173565b905061393a615a3e565b604080516020810190915288546001600160e01b0316815261395c90836151b1565b905060405180604001604052806139ac83600001516040518060400160405280601a81526020017f6e657720696e64657820657863656564732032323420626974730000000000008152506151d6565b6001600160e01b031681526020016139f9886040518060400160405280601c81526020017f626c6f636b206e756d6265722065786365656473203332206269747300000000815250615270565b63ffffffff9081169091526001600160a01b038c166000908152601160209081526040909120835181549490920151909216600160e01b026001600160e01b039182166001600160e01b0319909416939093171691909117905550611b9892505050565b8015611b9857613aa2826040518060400160405280601c81526020017f626c6f636b206e756d6265722065786365656473203332206269747300000000815250615270565b845463ffffffff91909116600160e01b026001600160e01b03909116178455505050505050565b6001600160a01b0384166000908152601160205260409020613ae9615a3e565b50604080516020810190915281546001600160e01b03168152613b0a615a3e565b5060408051602080820183526001600160a01b03808a16600090815260138352848120918a1680825282845294812080548552865195909152915291909155805115613cab57613b58615a3e565b613b6283836152c6565b90506000613bf1896001600160a01b03166395dd91938a6040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015613bbf57600080fd5b505afa158015613bd3573d6000803e3d6000fd5b505050506040513d6020811015613be957600080fd5b505188615113565b90506000613bff82846152eb565b6001600160a01b038a1660009081526014602052604081205491925090613c26908361531a565b9050613c478a828a613c3f5766038d7ea4c68000613c42565b60005b61535c565b6001600160a01b03808c1660008181526014602090815260409182902094909455895181518781529485015280519193928f16927f1fc3ecc087d8d2d15e23d0032af5a47059c3892d003d8e139fdcb6bb327c99a6929081900390910190a3505050505b50505050505050565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0836011811115613ce357fe5b836013811115613cef57fe5b604080519283526020830191909152600082820152519081900360600190a1826011811115611cdb57fe5b519051111590565b5190511090565b6060600d805480602002602001604051908101604052809291908181526020018280548015613d8157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613d63575b50939450600093505050505b8151811015613e47576000828281518110613da457fe5b60200260200101519050613db6615a3e565b6040518060200160405280836001600160a01b031663aa5af0fd6040518163ffffffff1660e01b815260040160206040518083038186803b158015613dfa57600080fd5b505afa158015613e0e573d6000803e3d6000fd5b505050506040513d6020811015613e2457600080fd5b505190529050613e3382614509565b613e3d828261381d565b5050600101613d8d565b50613e50615a3e565b60405180602001604052806000815250905060608251604051908082528060200260200182016040528015613e9f57816020015b613e8c615a3e565b815260200190600190039081613e845790505b50905060005b8351811015614025576000848281518110613ebc57fe5b6020908102919091018101516001600160a01b0381166000908152600990925260409091206003015490915060ff161561401c57613ef8615a3e565b60408051602080820180845260045463fc57d4df60e01b9091526001600160a01b03868116602485015293519293849391169163fc57d4df916044808601929190818703018186803b158015613f4d57600080fd5b505afa158015613f61573d6000803e3d6000fd5b505050506040513d6020811015613f7757600080fd5b505190529050613f85615a3e565b613ff382846001600160a01b03166347bd37186040518163ffffffff1660e01b815260040160206040518083038186803b158015613fc257600080fd5b505afa158015613fd6573d6000803e3d6000fd5b505050506040513d6020811015613fec57600080fd5b50516154ba565b90508085858151811061400257fe5b602002602001018190525061401786826151b1565b955050505b50600101613ea5565b5060005b8351811015611b8d576000600d828154811061404157fe5b600091825260208220015485516001600160a01b03909116925061406657600061408e565b61408e600e5461408986868151811061407b57fe5b6020026020010151886154db565b61550e565b6001600160a01b0383166000818152600f60209081526040918290208490558151848152915193945091927f2ab93f65628379309f36cb125e90d7c902454a545c4f8b8cb0794af75c24b807929181900390910190a25050600101614029565b60008060006140fb615a75565b6001600160a01b03881660009081526008602090815260408083208054825181850281018501909352808352849360609392919083018282801561416857602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161414a575b50939450600093505050505b81518110156144c457600082828151811061418b57fe5b60200260200101519050806001600160a01b031663c37f68e28e6040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060806040518083038186803b1580156141eb57600080fd5b505afa1580156141ff573d6000803e3d6000fd5b505050506040513d608081101561421557600080fd5b508051602082015160408084015160609485015160808c0152938a0193909352918801919091529450841561425b5750600f975060009650869550611c1e945050505050565b60408051602080820183526001600160a01b0380851660008181526009845285902060010154845260c08b01939093528351808301855260808b0151815260e08b015260048054855163fc57d4df60e01b815291820194909452935192169263fc57d4df9260248083019392829003018186803b1580156142db57600080fd5b505afa1580156142ef573d6000803e3d6000fd5b505050506040513d602081101561430557600080fd5b505160a087018190526143295750600d975060009650869550611c1e945050505050565b604080516020810190915260a08701518152610100870181905260c087015160e088015161435692615527565b6101208801529350600084600381111561436c57fe5b146143885750600b975060009650869550611c1e945050505050565b6143a08661012001518760400151886000015161557f565b8752935060008460038111156143b257fe5b146143ce5750600b975060009650869550611c1e945050505050565b6143e68661010001518760600151886020015161557f565b6020880152935060008460038111156143fb57fe5b146144175750600b975060009650869550611c1e945050505050565b8b6001600160a01b0316816001600160a01b031614156144bb576144458661012001518c886020015161557f565b60208801529350600084600381111561445a57fe5b146144765750600b975060009650869550611c1e945050505050565b61448a8661010001518b886020015161557f565b60208801529350600084600381111561449f57fe5b146144bb5750600b975060009650869550611c1e945050505050565b50600101614174565b506020840151845111156144eb575050506020810151905160009450039150829050611c1e565b5050815160209092015160009550859450919091039150611c1e9050565b6001600160a01b0381166000908152601060209081526040808320600f9092528220549091614536611b93565b8354909150600090614556908390600160e01b900463ffffffff166150d1565b90506000811180156145685750600083115b15614740576000856001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156145a857600080fd5b505afa1580156145bc573d6000803e3d6000fd5b505050506040513d60208110156145d257600080fd5b5051905060006145e28386615131565b90506145ec615a3e565b600083116146095760405180602001604052806000815250614613565b6146138284615173565b905061461d615a3e565b604080516020810190915288546001600160e01b0316815261463f90836151b1565b9050604051806040016040528061468f83600001516040518060400160405280601a81526020017f6e657720696e64657820657863656564732032323420626974730000000000008152506151d6565b6001600160e01b031681526020016146dc886040518060400160405280601c81526020017f626c6f636b206e756d6265722065786365656473203332206269747300000000815250615270565b63ffffffff9081169091526001600160a01b038b166000908152601060209081526040909120835181549490920151909216600160e01b026001600160e01b039182166001600160e01b031990941693909317169190911790555061152e92505050565b801561152e57614785826040518060400160405280601c81526020017f626c6f636b206e756d6265722065786365656473203332206269747300000000815250615270565b845463ffffffff91909116600160e01b026001600160e01b039091161784555050505050565b6001600160a01b03831660009081526010602052604090206147cb615a3e565b50604080516020810190915281546001600160e01b031681526147ec615a3e565b5060408051602080820183526001600160a01b0380891660009081526012835284812091891680825282845294812080548552865195909152915291909155805115801561483a5750815115155b15614852576ec097ce7bc90715b34b9f100000000081525b61485a615a3e565b61486483836152c6565b90506000876001600160a01b03166370a08231886040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156148be57600080fd5b505afa1580156148d2573d6000803e3d6000fd5b505050506040513d60208110156148e857600080fd5b5051905060006148f882846152eb565b6001600160a01b0389166000908152601460205260408120549192509061491f908361531a565b905061493889828a613c3f5766038d7ea4c68000613c42565b6001600160a01b03808b1660008181526014602090815260409182902094909455895181518781529485015280519193928e16927f2caecd17d02f56fa897705dcc740da2d237c373f70686f4e0d9bd3bf0400ea7a929081900390910190a350505050505050505050565b60008060006149b68460008060006140ee565b9250925092509193909250565b60008060006149d0615a3e565b6149da86866155cc565b909250905060008260038111156149ed57fe5b146149fe5750915060009050614a10565b6000614a0982615634565b9350935050505b9250929050565b600080546001600160a01b0316331480614a3b57506002546001600160a01b031633145b905090565b60005b600d54811015614ad457816001600160a01b0316600d8281548110614a6457fe5b6000918252602090912001546001600160a01b03161415614acc576040805162461bcd60e51b815260206004820152601460248201527f6d61726b657420616c7265616479206164646564000000000000000000000000604482015290519081900360640190fd5b600101614a43565b50600d80546001810182556000919091527fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb50180546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b03831660009081526009602052604081205460ff16614b4e576009611cb9565b6001600160a01b038085166000908152600960209081526040808320938716835260029093019052205460ff16614b86576000611cb9565b600080614b9685878660006140ee565b91935090915060009050826011811115614bac57fe5b14614bbd578160118111156131aa57fe5b8015612ee95760046131aa565b6001600160a01b0382166000908152600960205260408120805460ff16614bf557600991505061134d565b6001600160a01b038316600090815260028201602052604090205460ff16151560011415614c2757600091505061134d565b6007546001600160a01b03841660009081526008602052604090205410614c5257601091505061134d565b6001600160a01b0380841660008181526002840160209081526040808320805460ff19166001908117909155600883528184208054918201815584529282902090920180549489166001600160a01b031990951685179055815193845283019190915280517f3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a59281900390910190a15060009392505050565b6000614cf5615a3e565b614d1b604051806020016040528086815250604051806020016040528086815250615643565b915091509250929050565b6000614d30615a3e565b83518351614d1b919061572c565b6001600160a01b0381166000908152600960205260409020805460ff161515600114614db1576040805162461bcd60e51b815260206004820152601960248201527f636f6d70206d61726b6574206973206e6f74206c697374656400000000000000604482015290519081900360640190fd5b600381015460ff1615614e0b576040805162461bcd60e51b815260206004820152601960248201527f636f6d70206d61726b657420616c726561647920616464656400000000000000604482015290519081900360640190fd5b60038101805460ff19166001908117909155604080516001600160a01b0385168152602081019290925280517f93c1f3e36ed71139f466a4ce8c9751790e2e33f5afb2df0dcfb3aeabe55d5aa29281900390910190a16001600160a01b0382166000908152601060205260409020546001600160e01b0316158015614eb357506001600160a01b038216600090815260106020526040902054600160e01b900463ffffffff16155b15614f825760405180604001604052806ec097ce7bc90715b34b9f10000000006001600160e01b03168152602001614f27614eec611b93565b6040518060400160405280601c81526020017f626c6f636b206e756d6265722065786365656473203332206269747300000000815250615270565b63ffffffff9081169091526001600160a01b0384166000908152601060209081526040909120835181549490920151909216600160e01b026001600160e01b039182166001600160e01b031990941693909317169190911790555b6001600160a01b0382166000908152601160205260409020546001600160e01b0316158015614fd457506001600160a01b038216600090815260116020526040902054600160e01b900463ffffffff16155b156119ce5760405180604001604052806ec097ce7bc90715b34b9f10000000006001600160e01b0316815260200161500d614eec611b93565b63ffffffff9081169091526001600160a01b0384166000908152601160209081526040909120835181549490920151909216600160e01b026001600160e01b039182166001600160e01b031990941693909317169190911790555050565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa084601181111561509a57fe5b8460138111156150a657fe5b604080519283526020830191909152818101859052519081900360600190a1836011811115611cd857fe5b6000611cdb83836040518060400160405280601581526020017f7375627472616374696f6e20756e646572666c6f7700000000000000000000008152506157dc565b6000611cdb61512a84670de0b6b3a7640000615131565b8351615836565b6000611cdb83836040518060400160405280601781526020017f6d756c7469706c69636174696f6e206f766572666c6f77000000000000000000815250615878565b61517b615a3e565b60405180602001604052806151a86151a2866ec097ce7bc90715b34b9f1000000000615131565b85615836565b90529392505050565b6151b9615a3e565b60405180602001604052806151a88560000151856000015161531a565b600081600160e01b84106152685760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561522d578181015183820152602001615215565b50505050905090810190601f16801561525a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b509192915050565b60008164010000000084106152685760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561522d578181015183820152602001615215565b6152ce615a3e565b60405180602001604052806151a8856000015185600001516150d1565b60006ec097ce7bc90715b34b9f100000000061530b848460000151615131565b8161531257fe5b049392505050565b6000611cdb83836040518060400160405280601181526020017f6164646974696f6e206f766572666c6f770000000000000000000000000000008152506158f7565b600081831015801561536e5750600083115b156154b257600061537d6125f5565b604080517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290519192506000916001600160a01b038416916370a08231916024808301926020929190829003018186803b1580156153e257600080fd5b505afa1580156153f6573d6000803e3d6000fd5b505050506040513d602081101561540c57600080fd5b505190508085116154af57816001600160a01b031663a9059cbb87876040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561547757600080fd5b505af115801561548b573d6000803e3d6000fd5b505050506040513d60208110156154a157600080fd5b5060009350611cdb92505050565b50505b509092915050565b6154c2615a3e565b60405180602001604052806151a8856000015185615131565b6154e3615a3e565b60405180602001604052806151a86155078660000151670de0b6b3a7640000615131565b8551615836565b6000670de0b6b3a764000061530b848460000151615131565b6000615531615a3e565b600061553b615a3e565b6155458787615643565b9092509050600082600381111561555857fe5b14615567579092509050612c4d565b6155718186615643565b935093505050935093915050565b600080600061558c615a3e565b61559687876155cc565b909250905060008260038111156155a957fe5b146155ba5750915060009050612c4d565b6155716155c682615634565b8661594c565b60006155d6615a3e565b6000806155e7866000015186615972565b909250905060008260038111156155fa57fe5b1461561957506040805160208101909152600081529092509050614a10565b60408051602081019091529081526000969095509350505050565b51670de0b6b3a7640000900490565b600061564d615a3e565b60008061566286600001518660000151615972565b9092509050600082600381111561567557fe5b1461569457506040805160208101909152600081529092509050614a10565b6000806156a96706f05b59d3b200008461594c565b909250905060008260038111156156bc57fe5b146156de57506040805160208101909152600081529094509250614a10915050565b6000806156f383670de0b6b3a76400006159b1565b9092509050600082600381111561570657fe5b1461570d57fe5b604080516020810190915290815260009a909950975050505050505050565b6000615736615a3e565b60008061574b86670de0b6b3a7640000615972565b9092509050600082600381111561575e57fe5b1461577d57506040805160208101909152600081529092509050614a10565b60008061578a83886159b1565b9092509050600082600381111561579d57fe5b146157bf57506040805160208101909152600081529094509250614a10915050565b604080516020810190915290815260009890975095505050505050565b6000818484111561582e5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561522d578181015183820152602001615215565b505050900390565b6000611cdb83836040518060400160405280600e81526020017f646976696465206279207a65726f0000000000000000000000000000000000008152506159dc565b6000831580615885575082155b1561589257506000611cdb565b8383028385828161589f57fe5b041483906158ee5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561522d578181015183820152602001615215565b50949350505050565b600083830182858210156158ee5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561522d578181015183820152602001615215565b60008083830184811061596457600092509050614a10565b506002915060009050614a10565b6000808361598557506000905080614a10565b8383028385828161599257fe5b04146159a657506002915060009050614a10565b600092509050614a10565b600080826159c55750600190506000614a10565b60008385816159d057fe5b04915091509250929050565b60008183615a2b5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561522d578181015183820152602001615215565b50828481615a3557fe5b04949350505050565b6040518060200160405280600081525090565b8154818355818111156113b0576000838152602090206113b0918101908301615adf565b604051806101400160405280600081526020016000815260200160008152602001600081526020016000815260200160008152602001615ab3615a3e565b8152602001615ac0615a3e565b8152602001615acd615a3e565b8152602001615ada615a3e565b905290565b611b9591905b80821115615af95760008155600101615ae5565b509056fe63616e6e6f742070617573652061206d61726b65742074686174206973206e6f74206c69737465646f6e6c792065787465726e616c6c79206f776e6564206163636f756e7473206d61792072656672657368207370656564736f6e6c7920706175736520677561726469616e20616e642061646d696e2063616e207061757365657869744d61726b65743a206765744163636f756e74536e617073686f74206661696c65646f6e6c7920756e6974726f6c6c65722061646d696e2063616e206368616e676520627261696e73a265627a7a72315820101ffe2f0199e79c88bc2efeabcd53a1bfa93797f47ae15200670b2add70d63e64736f6c63430005100032"

// DeployComptrollerG4 deploys a new Ethereum contract, binding an instance of ComptrollerG4 to it.
func DeployComptrollerG4(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ComptrollerG4, error) {
	parsed, err := abi.JSON(strings.NewReader(ComptrollerG4ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ComptrollerG4Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ComptrollerG4{ComptrollerG4Caller: ComptrollerG4Caller{contract: contract}, ComptrollerG4Transactor: ComptrollerG4Transactor{contract: contract}, ComptrollerG4Filterer: ComptrollerG4Filterer{contract: contract}}, nil
}

// ComptrollerG4 is an auto generated Go binding around an Ethereum contract.
type ComptrollerG4 struct {
	ComptrollerG4Caller     // Read-only binding to the contract
	ComptrollerG4Transactor // Write-only binding to the contract
	ComptrollerG4Filterer   // Log filterer for contract events
}

// ComptrollerG4Caller is an auto generated read-only Go binding around an Ethereum contract.
type ComptrollerG4Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerG4Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ComptrollerG4Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerG4Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComptrollerG4Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerG4Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComptrollerG4Session struct {
	Contract     *ComptrollerG4    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ComptrollerG4CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComptrollerG4CallerSession struct {
	Contract *ComptrollerG4Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ComptrollerG4TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComptrollerG4TransactorSession struct {
	Contract     *ComptrollerG4Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ComptrollerG4Raw is an auto generated low-level Go binding around an Ethereum contract.
type ComptrollerG4Raw struct {
	Contract *ComptrollerG4 // Generic contract binding to access the raw methods on
}

// ComptrollerG4CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComptrollerG4CallerRaw struct {
	Contract *ComptrollerG4Caller // Generic read-only contract binding to access the raw methods on
}

// ComptrollerG4TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComptrollerG4TransactorRaw struct {
	Contract *ComptrollerG4Transactor // Generic write-only contract binding to access the raw methods on
}

// NewComptrollerG4 creates a new instance of ComptrollerG4, bound to a specific deployed contract.
func NewComptrollerG4(address common.Address, backend bind.ContractBackend) (*ComptrollerG4, error) {
	contract, err := bindComptrollerG4(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4{ComptrollerG4Caller: ComptrollerG4Caller{contract: contract}, ComptrollerG4Transactor: ComptrollerG4Transactor{contract: contract}, ComptrollerG4Filterer: ComptrollerG4Filterer{contract: contract}}, nil
}

// NewComptrollerG4Caller creates a new read-only instance of ComptrollerG4, bound to a specific deployed contract.
func NewComptrollerG4Caller(address common.Address, caller bind.ContractCaller) (*ComptrollerG4Caller, error) {
	contract, err := bindComptrollerG4(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4Caller{contract: contract}, nil
}

// NewComptrollerG4Transactor creates a new write-only instance of ComptrollerG4, bound to a specific deployed contract.
func NewComptrollerG4Transactor(address common.Address, transactor bind.ContractTransactor) (*ComptrollerG4Transactor, error) {
	contract, err := bindComptrollerG4(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4Transactor{contract: contract}, nil
}

// NewComptrollerG4Filterer creates a new log filterer instance of ComptrollerG4, bound to a specific deployed contract.
func NewComptrollerG4Filterer(address common.Address, filterer bind.ContractFilterer) (*ComptrollerG4Filterer, error) {
	contract, err := bindComptrollerG4(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4Filterer{contract: contract}, nil
}

// bindComptrollerG4 binds a generic wrapper to an already deployed contract.
func bindComptrollerG4(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ComptrollerG4ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComptrollerG4 *ComptrollerG4Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComptrollerG4.Contract.ComptrollerG4Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComptrollerG4 *ComptrollerG4Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.ComptrollerG4Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComptrollerG4 *ComptrollerG4Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.ComptrollerG4Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComptrollerG4 *ComptrollerG4CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComptrollerG4.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComptrollerG4 *ComptrollerG4TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComptrollerG4 *ComptrollerG4TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.contract.Transact(opts, method, params...)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_ComptrollerG4 *ComptrollerG4Caller) AccountAssets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "accountAssets", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_ComptrollerG4 *ComptrollerG4Session) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _ComptrollerG4.Contract.AccountAssets(&_ComptrollerG4.CallOpts, arg0, arg1)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_ComptrollerG4 *ComptrollerG4CallerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _ComptrollerG4.Contract.AccountAssets(&_ComptrollerG4.CallOpts, arg0, arg1)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ComptrollerG4 *ComptrollerG4Caller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ComptrollerG4 *ComptrollerG4Session) Admin() (common.Address, error) {
	return _ComptrollerG4.Contract.Admin(&_ComptrollerG4.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_ComptrollerG4 *ComptrollerG4CallerSession) Admin() (common.Address, error) {
	return _ComptrollerG4.Contract.Admin(&_ComptrollerG4.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_ComptrollerG4 *ComptrollerG4Caller) AllMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "allMarkets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_ComptrollerG4 *ComptrollerG4Session) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _ComptrollerG4.Contract.AllMarkets(&_ComptrollerG4.CallOpts, arg0)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_ComptrollerG4 *ComptrollerG4CallerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _ComptrollerG4.Contract.AllMarkets(&_ComptrollerG4.CallOpts, arg0)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_ComptrollerG4 *ComptrollerG4Caller) BorrowGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "borrowGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_ComptrollerG4 *ComptrollerG4Session) BorrowGuardianPaused(arg0 common.Address) (bool, error) {
	return _ComptrollerG4.Contract.BorrowGuardianPaused(&_ComptrollerG4.CallOpts, arg0)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_ComptrollerG4 *ComptrollerG4CallerSession) BorrowGuardianPaused(arg0 common.Address) (bool, error) {
	return _ComptrollerG4.Contract.BorrowGuardianPaused(&_ComptrollerG4.CallOpts, arg0)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address cToken) view returns(bool)
func (_ComptrollerG4 *ComptrollerG4Caller) CheckMembership(opts *bind.CallOpts, account common.Address, cToken common.Address) (bool, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "checkMembership", account, cToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address cToken) view returns(bool)
func (_ComptrollerG4 *ComptrollerG4Session) CheckMembership(account common.Address, cToken common.Address) (bool, error) {
	return _ComptrollerG4.Contract.CheckMembership(&_ComptrollerG4.CallOpts, account, cToken)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address cToken) view returns(bool)
func (_ComptrollerG4 *ComptrollerG4CallerSession) CheckMembership(account common.Address, cToken common.Address) (bool, error) {
	return _ComptrollerG4.Contract.CheckMembership(&_ComptrollerG4.CallOpts, account, cToken)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Caller) CloseFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "closeFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) CloseFactorMantissa() (*big.Int, error) {
	return _ComptrollerG4.Contract.CloseFactorMantissa(&_ComptrollerG4.CallOpts)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4CallerSession) CloseFactorMantissa() (*big.Int, error) {
	return _ComptrollerG4.Contract.CloseFactorMantissa(&_ComptrollerG4.CallOpts)
}

// CompAccrued is a free data retrieval call binding the contract method 0xcc7ebdc4.
//
// Solidity: function compAccrued(address ) view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Caller) CompAccrued(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "compAccrued", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompAccrued is a free data retrieval call binding the contract method 0xcc7ebdc4.
//
// Solidity: function compAccrued(address ) view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) CompAccrued(arg0 common.Address) (*big.Int, error) {
	return _ComptrollerG4.Contract.CompAccrued(&_ComptrollerG4.CallOpts, arg0)
}

// CompAccrued is a free data retrieval call binding the contract method 0xcc7ebdc4.
//
// Solidity: function compAccrued(address ) view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4CallerSession) CompAccrued(arg0 common.Address) (*big.Int, error) {
	return _ComptrollerG4.Contract.CompAccrued(&_ComptrollerG4.CallOpts, arg0)
}

// CompBorrowState is a free data retrieval call binding the contract method 0x8c57804e.
//
// Solidity: function compBorrowState(address ) view returns(uint224 index, uint32 block)
func (_ComptrollerG4 *ComptrollerG4Caller) CompBorrowState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "compBorrowState", arg0)

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})

	outstruct.Index = out[0].(*big.Int)
	outstruct.Block = out[1].(uint32)

	return *outstruct, err

}

// CompBorrowState is a free data retrieval call binding the contract method 0x8c57804e.
//
// Solidity: function compBorrowState(address ) view returns(uint224 index, uint32 block)
func (_ComptrollerG4 *ComptrollerG4Session) CompBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _ComptrollerG4.Contract.CompBorrowState(&_ComptrollerG4.CallOpts, arg0)
}

// CompBorrowState is a free data retrieval call binding the contract method 0x8c57804e.
//
// Solidity: function compBorrowState(address ) view returns(uint224 index, uint32 block)
func (_ComptrollerG4 *ComptrollerG4CallerSession) CompBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _ComptrollerG4.Contract.CompBorrowState(&_ComptrollerG4.CallOpts, arg0)
}

// CompBorrowerIndex is a free data retrieval call binding the contract method 0xca0af043.
//
// Solidity: function compBorrowerIndex(address , address ) view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Caller) CompBorrowerIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "compBorrowerIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompBorrowerIndex is a free data retrieval call binding the contract method 0xca0af043.
//
// Solidity: function compBorrowerIndex(address , address ) view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) CompBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ComptrollerG4.Contract.CompBorrowerIndex(&_ComptrollerG4.CallOpts, arg0, arg1)
}

// CompBorrowerIndex is a free data retrieval call binding the contract method 0xca0af043.
//
// Solidity: function compBorrowerIndex(address , address ) view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4CallerSession) CompBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ComptrollerG4.Contract.CompBorrowerIndex(&_ComptrollerG4.CallOpts, arg0, arg1)
}

// CompClaimThreshold is a free data retrieval call binding the contract method 0x747026c9.
//
// Solidity: function compClaimThreshold() view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Caller) CompClaimThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "compClaimThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompClaimThreshold is a free data retrieval call binding the contract method 0x747026c9.
//
// Solidity: function compClaimThreshold() view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) CompClaimThreshold() (*big.Int, error) {
	return _ComptrollerG4.Contract.CompClaimThreshold(&_ComptrollerG4.CallOpts)
}

// CompClaimThreshold is a free data retrieval call binding the contract method 0x747026c9.
//
// Solidity: function compClaimThreshold() view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4CallerSession) CompClaimThreshold() (*big.Int, error) {
	return _ComptrollerG4.Contract.CompClaimThreshold(&_ComptrollerG4.CallOpts)
}

// CompInitialIndex is a free data retrieval call binding the contract method 0xa7f0e231.
//
// Solidity: function compInitialIndex() view returns(uint224)
func (_ComptrollerG4 *ComptrollerG4Caller) CompInitialIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "compInitialIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompInitialIndex is a free data retrieval call binding the contract method 0xa7f0e231.
//
// Solidity: function compInitialIndex() view returns(uint224)
func (_ComptrollerG4 *ComptrollerG4Session) CompInitialIndex() (*big.Int, error) {
	return _ComptrollerG4.Contract.CompInitialIndex(&_ComptrollerG4.CallOpts)
}

// CompInitialIndex is a free data retrieval call binding the contract method 0xa7f0e231.
//
// Solidity: function compInitialIndex() view returns(uint224)
func (_ComptrollerG4 *ComptrollerG4CallerSession) CompInitialIndex() (*big.Int, error) {
	return _ComptrollerG4.Contract.CompInitialIndex(&_ComptrollerG4.CallOpts)
}

// CompRate is a free data retrieval call binding the contract method 0xaa900754.
//
// Solidity: function compRate() view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Caller) CompRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "compRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompRate is a free data retrieval call binding the contract method 0xaa900754.
//
// Solidity: function compRate() view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) CompRate() (*big.Int, error) {
	return _ComptrollerG4.Contract.CompRate(&_ComptrollerG4.CallOpts)
}

// CompRate is a free data retrieval call binding the contract method 0xaa900754.
//
// Solidity: function compRate() view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4CallerSession) CompRate() (*big.Int, error) {
	return _ComptrollerG4.Contract.CompRate(&_ComptrollerG4.CallOpts)
}

// CompSpeeds is a free data retrieval call binding the contract method 0x1d7b33d7.
//
// Solidity: function compSpeeds(address ) view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Caller) CompSpeeds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "compSpeeds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompSpeeds is a free data retrieval call binding the contract method 0x1d7b33d7.
//
// Solidity: function compSpeeds(address ) view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) CompSpeeds(arg0 common.Address) (*big.Int, error) {
	return _ComptrollerG4.Contract.CompSpeeds(&_ComptrollerG4.CallOpts, arg0)
}

// CompSpeeds is a free data retrieval call binding the contract method 0x1d7b33d7.
//
// Solidity: function compSpeeds(address ) view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4CallerSession) CompSpeeds(arg0 common.Address) (*big.Int, error) {
	return _ComptrollerG4.Contract.CompSpeeds(&_ComptrollerG4.CallOpts, arg0)
}

// CompSupplierIndex is a free data retrieval call binding the contract method 0xb21be7fd.
//
// Solidity: function compSupplierIndex(address , address ) view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Caller) CompSupplierIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "compSupplierIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompSupplierIndex is a free data retrieval call binding the contract method 0xb21be7fd.
//
// Solidity: function compSupplierIndex(address , address ) view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) CompSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ComptrollerG4.Contract.CompSupplierIndex(&_ComptrollerG4.CallOpts, arg0, arg1)
}

// CompSupplierIndex is a free data retrieval call binding the contract method 0xb21be7fd.
//
// Solidity: function compSupplierIndex(address , address ) view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4CallerSession) CompSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ComptrollerG4.Contract.CompSupplierIndex(&_ComptrollerG4.CallOpts, arg0, arg1)
}

// CompSupplyState is a free data retrieval call binding the contract method 0x6b79c38d.
//
// Solidity: function compSupplyState(address ) view returns(uint224 index, uint32 block)
func (_ComptrollerG4 *ComptrollerG4Caller) CompSupplyState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "compSupplyState", arg0)

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})

	outstruct.Index = out[0].(*big.Int)
	outstruct.Block = out[1].(uint32)

	return *outstruct, err

}

// CompSupplyState is a free data retrieval call binding the contract method 0x6b79c38d.
//
// Solidity: function compSupplyState(address ) view returns(uint224 index, uint32 block)
func (_ComptrollerG4 *ComptrollerG4Session) CompSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _ComptrollerG4.Contract.CompSupplyState(&_ComptrollerG4.CallOpts, arg0)
}

// CompSupplyState is a free data retrieval call binding the contract method 0x6b79c38d.
//
// Solidity: function compSupplyState(address ) view returns(uint224 index, uint32 block)
func (_ComptrollerG4 *ComptrollerG4CallerSession) CompSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _ComptrollerG4.Contract.CompSupplyState(&_ComptrollerG4.CallOpts, arg0)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_ComptrollerG4 *ComptrollerG4Caller) ComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "comptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_ComptrollerG4 *ComptrollerG4Session) ComptrollerImplementation() (common.Address, error) {
	return _ComptrollerG4.Contract.ComptrollerImplementation(&_ComptrollerG4.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_ComptrollerG4 *ComptrollerG4CallerSession) ComptrollerImplementation() (common.Address, error) {
	return _ComptrollerG4.Contract.ComptrollerImplementation(&_ComptrollerG4.CallOpts)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_ComptrollerG4 *ComptrollerG4Caller) GetAccountLiquidity(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "getAccountLiquidity", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_ComptrollerG4 *ComptrollerG4Session) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _ComptrollerG4.Contract.GetAccountLiquidity(&_ComptrollerG4.CallOpts, account)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_ComptrollerG4 *ComptrollerG4CallerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _ComptrollerG4.Contract.GetAccountLiquidity(&_ComptrollerG4.CallOpts, account)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_ComptrollerG4 *ComptrollerG4Caller) GetAllMarkets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "getAllMarkets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_ComptrollerG4 *ComptrollerG4Session) GetAllMarkets() ([]common.Address, error) {
	return _ComptrollerG4.Contract.GetAllMarkets(&_ComptrollerG4.CallOpts)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_ComptrollerG4 *ComptrollerG4CallerSession) GetAllMarkets() ([]common.Address, error) {
	return _ComptrollerG4.Contract.GetAllMarkets(&_ComptrollerG4.CallOpts)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_ComptrollerG4 *ComptrollerG4Caller) GetAssetsIn(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "getAssetsIn", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_ComptrollerG4 *ComptrollerG4Session) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _ComptrollerG4.Contract.GetAssetsIn(&_ComptrollerG4.CallOpts, account)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_ComptrollerG4 *ComptrollerG4CallerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _ComptrollerG4.Contract.GetAssetsIn(&_ComptrollerG4.CallOpts, account)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Caller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) GetBlockNumber() (*big.Int, error) {
	return _ComptrollerG4.Contract.GetBlockNumber(&_ComptrollerG4.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4CallerSession) GetBlockNumber() (*big.Int, error) {
	return _ComptrollerG4.Contract.GetBlockNumber(&_ComptrollerG4.CallOpts)
}

// GetCompAddress is a free data retrieval call binding the contract method 0x9d1b5a0a.
//
// Solidity: function getCompAddress() view returns(address)
func (_ComptrollerG4 *ComptrollerG4Caller) GetCompAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "getCompAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCompAddress is a free data retrieval call binding the contract method 0x9d1b5a0a.
//
// Solidity: function getCompAddress() view returns(address)
func (_ComptrollerG4 *ComptrollerG4Session) GetCompAddress() (common.Address, error) {
	return _ComptrollerG4.Contract.GetCompAddress(&_ComptrollerG4.CallOpts)
}

// GetCompAddress is a free data retrieval call binding the contract method 0x9d1b5a0a.
//
// Solidity: function getCompAddress() view returns(address)
func (_ComptrollerG4 *ComptrollerG4CallerSession) GetCompAddress() (common.Address, error) {
	return _ComptrollerG4.Contract.GetCompAddress(&_ComptrollerG4.CallOpts)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address cTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_ComptrollerG4 *ComptrollerG4Caller) GetHypotheticalAccountLiquidity(opts *bind.CallOpts, account common.Address, cTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "getHypotheticalAccountLiquidity", account, cTokenModify, redeemTokens, borrowAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address cTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_ComptrollerG4 *ComptrollerG4Session) GetHypotheticalAccountLiquidity(account common.Address, cTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _ComptrollerG4.Contract.GetHypotheticalAccountLiquidity(&_ComptrollerG4.CallOpts, account, cTokenModify, redeemTokens, borrowAmount)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address cTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_ComptrollerG4 *ComptrollerG4CallerSession) GetHypotheticalAccountLiquidity(account common.Address, cTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _ComptrollerG4.Contract.GetHypotheticalAccountLiquidity(&_ComptrollerG4.CallOpts, account, cTokenModify, redeemTokens, borrowAmount)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_ComptrollerG4 *ComptrollerG4Caller) IsComptroller(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "isComptroller")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_ComptrollerG4 *ComptrollerG4Session) IsComptroller() (bool, error) {
	return _ComptrollerG4.Contract.IsComptroller(&_ComptrollerG4.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_ComptrollerG4 *ComptrollerG4CallerSession) IsComptroller() (bool, error) {
	return _ComptrollerG4.Contract.IsComptroller(&_ComptrollerG4.CallOpts)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_ComptrollerG4 *ComptrollerG4Caller) LiquidateCalculateSeizeTokens(opts *bind.CallOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "liquidateCalculateSeizeTokens", cTokenBorrowed, cTokenCollateral, actualRepayAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_ComptrollerG4 *ComptrollerG4Session) LiquidateCalculateSeizeTokens(cTokenBorrowed common.Address, cTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _ComptrollerG4.Contract.LiquidateCalculateSeizeTokens(&_ComptrollerG4.CallOpts, cTokenBorrowed, cTokenCollateral, actualRepayAmount)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_ComptrollerG4 *ComptrollerG4CallerSession) LiquidateCalculateSeizeTokens(cTokenBorrowed common.Address, cTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _ComptrollerG4.Contract.LiquidateCalculateSeizeTokens(&_ComptrollerG4.CallOpts, cTokenBorrowed, cTokenCollateral, actualRepayAmount)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Caller) LiquidationIncentiveMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "liquidationIncentiveMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _ComptrollerG4.Contract.LiquidationIncentiveMantissa(&_ComptrollerG4.CallOpts)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4CallerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _ComptrollerG4.Contract.LiquidationIncentiveMantissa(&_ComptrollerG4.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isComped)
func (_ComptrollerG4 *ComptrollerG4Caller) Markets(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsComped                 bool
}, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "markets", arg0)

	outstruct := new(struct {
		IsListed                 bool
		CollateralFactorMantissa *big.Int
		IsComped                 bool
	})

	outstruct.IsListed = out[0].(bool)
	outstruct.CollateralFactorMantissa = out[1].(*big.Int)
	outstruct.IsComped = out[2].(bool)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isComped)
func (_ComptrollerG4 *ComptrollerG4Session) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsComped                 bool
}, error) {
	return _ComptrollerG4.Contract.Markets(&_ComptrollerG4.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isComped)
func (_ComptrollerG4 *ComptrollerG4CallerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsComped                 bool
}, error) {
	return _ComptrollerG4.Contract.Markets(&_ComptrollerG4.CallOpts, arg0)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Caller) MaxAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "maxAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) MaxAssets() (*big.Int, error) {
	return _ComptrollerG4.Contract.MaxAssets(&_ComptrollerG4.CallOpts)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_ComptrollerG4 *ComptrollerG4CallerSession) MaxAssets() (*big.Int, error) {
	return _ComptrollerG4.Contract.MaxAssets(&_ComptrollerG4.CallOpts)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_ComptrollerG4 *ComptrollerG4Caller) MintGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "mintGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_ComptrollerG4 *ComptrollerG4Session) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _ComptrollerG4.Contract.MintGuardianPaused(&_ComptrollerG4.CallOpts, arg0)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_ComptrollerG4 *ComptrollerG4CallerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _ComptrollerG4.Contract.MintGuardianPaused(&_ComptrollerG4.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_ComptrollerG4 *ComptrollerG4Caller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_ComptrollerG4 *ComptrollerG4Session) Oracle() (common.Address, error) {
	return _ComptrollerG4.Contract.Oracle(&_ComptrollerG4.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_ComptrollerG4 *ComptrollerG4CallerSession) Oracle() (common.Address, error) {
	return _ComptrollerG4.Contract.Oracle(&_ComptrollerG4.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_ComptrollerG4 *ComptrollerG4Caller) PauseGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "pauseGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_ComptrollerG4 *ComptrollerG4Session) PauseGuardian() (common.Address, error) {
	return _ComptrollerG4.Contract.PauseGuardian(&_ComptrollerG4.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_ComptrollerG4 *ComptrollerG4CallerSession) PauseGuardian() (common.Address, error) {
	return _ComptrollerG4.Contract.PauseGuardian(&_ComptrollerG4.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_ComptrollerG4 *ComptrollerG4Caller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_ComptrollerG4 *ComptrollerG4Session) PendingAdmin() (common.Address, error) {
	return _ComptrollerG4.Contract.PendingAdmin(&_ComptrollerG4.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_ComptrollerG4 *ComptrollerG4CallerSession) PendingAdmin() (common.Address, error) {
	return _ComptrollerG4.Contract.PendingAdmin(&_ComptrollerG4.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_ComptrollerG4 *ComptrollerG4Caller) PendingComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "pendingComptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_ComptrollerG4 *ComptrollerG4Session) PendingComptrollerImplementation() (common.Address, error) {
	return _ComptrollerG4.Contract.PendingComptrollerImplementation(&_ComptrollerG4.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_ComptrollerG4 *ComptrollerG4CallerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _ComptrollerG4.Contract.PendingComptrollerImplementation(&_ComptrollerG4.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_ComptrollerG4 *ComptrollerG4Caller) SeizeGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "seizeGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_ComptrollerG4 *ComptrollerG4Session) SeizeGuardianPaused() (bool, error) {
	return _ComptrollerG4.Contract.SeizeGuardianPaused(&_ComptrollerG4.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_ComptrollerG4 *ComptrollerG4CallerSession) SeizeGuardianPaused() (bool, error) {
	return _ComptrollerG4.Contract.SeizeGuardianPaused(&_ComptrollerG4.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_ComptrollerG4 *ComptrollerG4Caller) TransferGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ComptrollerG4.contract.Call(opts, &out, "transferGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_ComptrollerG4 *ComptrollerG4Session) TransferGuardianPaused() (bool, error) {
	return _ComptrollerG4.Contract.TransferGuardianPaused(&_ComptrollerG4.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_ComptrollerG4 *ComptrollerG4CallerSession) TransferGuardianPaused() (bool, error) {
	return _ComptrollerG4.Contract.TransferGuardianPaused(&_ComptrollerG4.CallOpts)
}

// AddCompMarkets is a paid mutator transaction binding the contract method 0xce485c5e.
//
// Solidity: function _addCompMarkets(address[] cTokens) returns()
func (_ComptrollerG4 *ComptrollerG4Transactor) AddCompMarkets(opts *bind.TransactOpts, cTokens []common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "_addCompMarkets", cTokens)
}

// AddCompMarkets is a paid mutator transaction binding the contract method 0xce485c5e.
//
// Solidity: function _addCompMarkets(address[] cTokens) returns()
func (_ComptrollerG4 *ComptrollerG4Session) AddCompMarkets(cTokens []common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.AddCompMarkets(&_ComptrollerG4.TransactOpts, cTokens)
}

// AddCompMarkets is a paid mutator transaction binding the contract method 0xce485c5e.
//
// Solidity: function _addCompMarkets(address[] cTokens) returns()
func (_ComptrollerG4 *ComptrollerG4TransactorSession) AddCompMarkets(cTokens []common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.AddCompMarkets(&_ComptrollerG4.TransactOpts, cTokens)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_ComptrollerG4 *ComptrollerG4Transactor) Become(opts *bind.TransactOpts, unitroller common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "_become", unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_ComptrollerG4 *ComptrollerG4Session) Become(unitroller common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.Become(&_ComptrollerG4.TransactOpts, unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_ComptrollerG4 *ComptrollerG4TransactorSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.Become(&_ComptrollerG4.TransactOpts, unitroller)
}

// DropCompMarket is a paid mutator transaction binding the contract method 0x3aa729b4.
//
// Solidity: function _dropCompMarket(address cToken) returns()
func (_ComptrollerG4 *ComptrollerG4Transactor) DropCompMarket(opts *bind.TransactOpts, cToken common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "_dropCompMarket", cToken)
}

// DropCompMarket is a paid mutator transaction binding the contract method 0x3aa729b4.
//
// Solidity: function _dropCompMarket(address cToken) returns()
func (_ComptrollerG4 *ComptrollerG4Session) DropCompMarket(cToken common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.DropCompMarket(&_ComptrollerG4.TransactOpts, cToken)
}

// DropCompMarket is a paid mutator transaction binding the contract method 0x3aa729b4.
//
// Solidity: function _dropCompMarket(address cToken) returns()
func (_ComptrollerG4 *ComptrollerG4TransactorSession) DropCompMarket(cToken common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.DropCompMarket(&_ComptrollerG4.TransactOpts, cToken)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x18c882a5.
//
// Solidity: function _setBorrowPaused(address cToken, bool state) returns(bool)
func (_ComptrollerG4 *ComptrollerG4Transactor) SetBorrowPaused(opts *bind.TransactOpts, cToken common.Address, state bool) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "_setBorrowPaused", cToken, state)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x18c882a5.
//
// Solidity: function _setBorrowPaused(address cToken, bool state) returns(bool)
func (_ComptrollerG4 *ComptrollerG4Session) SetBorrowPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetBorrowPaused(&_ComptrollerG4.TransactOpts, cToken, state)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x18c882a5.
//
// Solidity: function _setBorrowPaused(address cToken, bool state) returns(bool)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) SetBorrowPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetBorrowPaused(&_ComptrollerG4.TransactOpts, cToken, state)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Transactor) SetCloseFactor(opts *bind.TransactOpts, newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "_setCloseFactor", newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetCloseFactor(&_ComptrollerG4.TransactOpts, newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetCloseFactor(&_ComptrollerG4.TransactOpts, newCloseFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address cToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Transactor) SetCollateralFactor(opts *bind.TransactOpts, cToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "_setCollateralFactor", cToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address cToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) SetCollateralFactor(cToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetCollateralFactor(&_ComptrollerG4.TransactOpts, cToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address cToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) SetCollateralFactor(cToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetCollateralFactor(&_ComptrollerG4.TransactOpts, cToken, newCollateralFactorMantissa)
}

// SetCompRate is a paid mutator transaction binding the contract method 0x6a491112.
//
// Solidity: function _setCompRate(uint256 compRate_) returns()
func (_ComptrollerG4 *ComptrollerG4Transactor) SetCompRate(opts *bind.TransactOpts, compRate_ *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "_setCompRate", compRate_)
}

// SetCompRate is a paid mutator transaction binding the contract method 0x6a491112.
//
// Solidity: function _setCompRate(uint256 compRate_) returns()
func (_ComptrollerG4 *ComptrollerG4Session) SetCompRate(compRate_ *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetCompRate(&_ComptrollerG4.TransactOpts, compRate_)
}

// SetCompRate is a paid mutator transaction binding the contract method 0x6a491112.
//
// Solidity: function _setCompRate(uint256 compRate_) returns()
func (_ComptrollerG4 *ComptrollerG4TransactorSession) SetCompRate(compRate_ *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetCompRate(&_ComptrollerG4.TransactOpts, compRate_)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Transactor) SetLiquidationIncentive(opts *bind.TransactOpts, newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "_setLiquidationIncentive", newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetLiquidationIncentive(&_ComptrollerG4.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetLiquidationIncentive(&_ComptrollerG4.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetMaxAssets is a paid mutator transaction binding the contract method 0xd9226ced.
//
// Solidity: function _setMaxAssets(uint256 newMaxAssets) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Transactor) SetMaxAssets(opts *bind.TransactOpts, newMaxAssets *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "_setMaxAssets", newMaxAssets)
}

// SetMaxAssets is a paid mutator transaction binding the contract method 0xd9226ced.
//
// Solidity: function _setMaxAssets(uint256 newMaxAssets) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) SetMaxAssets(newMaxAssets *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetMaxAssets(&_ComptrollerG4.TransactOpts, newMaxAssets)
}

// SetMaxAssets is a paid mutator transaction binding the contract method 0xd9226ced.
//
// Solidity: function _setMaxAssets(uint256 newMaxAssets) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) SetMaxAssets(newMaxAssets *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetMaxAssets(&_ComptrollerG4.TransactOpts, newMaxAssets)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x3bcf7ec1.
//
// Solidity: function _setMintPaused(address cToken, bool state) returns(bool)
func (_ComptrollerG4 *ComptrollerG4Transactor) SetMintPaused(opts *bind.TransactOpts, cToken common.Address, state bool) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "_setMintPaused", cToken, state)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x3bcf7ec1.
//
// Solidity: function _setMintPaused(address cToken, bool state) returns(bool)
func (_ComptrollerG4 *ComptrollerG4Session) SetMintPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetMintPaused(&_ComptrollerG4.TransactOpts, cToken, state)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x3bcf7ec1.
//
// Solidity: function _setMintPaused(address cToken, bool state) returns(bool)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) SetMintPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetMintPaused(&_ComptrollerG4.TransactOpts, cToken, state)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Transactor) SetPauseGuardian(opts *bind.TransactOpts, newPauseGuardian common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "_setPauseGuardian", newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetPauseGuardian(&_ComptrollerG4.TransactOpts, newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetPauseGuardian(&_ComptrollerG4.TransactOpts, newPauseGuardian)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Transactor) SetPriceOracle(opts *bind.TransactOpts, newOracle common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "_setPriceOracle", newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetPriceOracle(&_ComptrollerG4.TransactOpts, newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetPriceOracle(&_ComptrollerG4.TransactOpts, newOracle)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool state) returns(bool)
func (_ComptrollerG4 *ComptrollerG4Transactor) SetSeizePaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "_setSeizePaused", state)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool state) returns(bool)
func (_ComptrollerG4 *ComptrollerG4Session) SetSeizePaused(state bool) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetSeizePaused(&_ComptrollerG4.TransactOpts, state)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool state) returns(bool)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) SetSeizePaused(state bool) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetSeizePaused(&_ComptrollerG4.TransactOpts, state)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool state) returns(bool)
func (_ComptrollerG4 *ComptrollerG4Transactor) SetTransferPaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "_setTransferPaused", state)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool state) returns(bool)
func (_ComptrollerG4 *ComptrollerG4Session) SetTransferPaused(state bool) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetTransferPaused(&_ComptrollerG4.TransactOpts, state)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool state) returns(bool)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) SetTransferPaused(state bool) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SetTransferPaused(&_ComptrollerG4.TransactOpts, state)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address cToken) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Transactor) SupportMarket(opts *bind.TransactOpts, cToken common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "_supportMarket", cToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address cToken) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) SupportMarket(cToken common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SupportMarket(&_ComptrollerG4.TransactOpts, cToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address cToken) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) SupportMarket(cToken common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SupportMarket(&_ComptrollerG4.TransactOpts, cToken)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Transactor) BorrowAllowed(opts *bind.TransactOpts, cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "borrowAllowed", cToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) BorrowAllowed(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.BorrowAllowed(&_ComptrollerG4.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) BorrowAllowed(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.BorrowAllowed(&_ComptrollerG4.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_ComptrollerG4 *ComptrollerG4Transactor) BorrowVerify(opts *bind.TransactOpts, cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "borrowVerify", cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_ComptrollerG4 *ComptrollerG4Session) BorrowVerify(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.BorrowVerify(&_ComptrollerG4.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_ComptrollerG4 *ComptrollerG4TransactorSession) BorrowVerify(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.BorrowVerify(&_ComptrollerG4.TransactOpts, cToken, borrower, borrowAmount)
}

// ClaimComp is a paid mutator transaction binding the contract method 0x1c3db2e0.
//
// Solidity: function claimComp(address holder, address[] cTokens) returns()
func (_ComptrollerG4 *ComptrollerG4Transactor) ClaimComp(opts *bind.TransactOpts, holder common.Address, cTokens []common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "claimComp", holder, cTokens)
}

// ClaimComp is a paid mutator transaction binding the contract method 0x1c3db2e0.
//
// Solidity: function claimComp(address holder, address[] cTokens) returns()
func (_ComptrollerG4 *ComptrollerG4Session) ClaimComp(holder common.Address, cTokens []common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.ClaimComp(&_ComptrollerG4.TransactOpts, holder, cTokens)
}

// ClaimComp is a paid mutator transaction binding the contract method 0x1c3db2e0.
//
// Solidity: function claimComp(address holder, address[] cTokens) returns()
func (_ComptrollerG4 *ComptrollerG4TransactorSession) ClaimComp(holder common.Address, cTokens []common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.ClaimComp(&_ComptrollerG4.TransactOpts, holder, cTokens)
}

// ClaimComp0 is a paid mutator transaction binding the contract method 0x6810dfa6.
//
// Solidity: function claimComp(address[] holders, address[] cTokens, bool borrowers, bool suppliers) returns()
func (_ComptrollerG4 *ComptrollerG4Transactor) ClaimComp0(opts *bind.TransactOpts, holders []common.Address, cTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "claimComp0", holders, cTokens, borrowers, suppliers)
}

// ClaimComp0 is a paid mutator transaction binding the contract method 0x6810dfa6.
//
// Solidity: function claimComp(address[] holders, address[] cTokens, bool borrowers, bool suppliers) returns()
func (_ComptrollerG4 *ComptrollerG4Session) ClaimComp0(holders []common.Address, cTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.ClaimComp0(&_ComptrollerG4.TransactOpts, holders, cTokens, borrowers, suppliers)
}

// ClaimComp0 is a paid mutator transaction binding the contract method 0x6810dfa6.
//
// Solidity: function claimComp(address[] holders, address[] cTokens, bool borrowers, bool suppliers) returns()
func (_ComptrollerG4 *ComptrollerG4TransactorSession) ClaimComp0(holders []common.Address, cTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.ClaimComp0(&_ComptrollerG4.TransactOpts, holders, cTokens, borrowers, suppliers)
}

// ClaimComp1 is a paid mutator transaction binding the contract method 0xe9af0292.
//
// Solidity: function claimComp(address holder) returns()
func (_ComptrollerG4 *ComptrollerG4Transactor) ClaimComp1(opts *bind.TransactOpts, holder common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "claimComp1", holder)
}

// ClaimComp1 is a paid mutator transaction binding the contract method 0xe9af0292.
//
// Solidity: function claimComp(address holder) returns()
func (_ComptrollerG4 *ComptrollerG4Session) ClaimComp1(holder common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.ClaimComp1(&_ComptrollerG4.TransactOpts, holder)
}

// ClaimComp1 is a paid mutator transaction binding the contract method 0xe9af0292.
//
// Solidity: function claimComp(address holder) returns()
func (_ComptrollerG4 *ComptrollerG4TransactorSession) ClaimComp1(holder common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.ClaimComp1(&_ComptrollerG4.TransactOpts, holder)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_ComptrollerG4 *ComptrollerG4Transactor) EnterMarkets(opts *bind.TransactOpts, cTokens []common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "enterMarkets", cTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_ComptrollerG4 *ComptrollerG4Session) EnterMarkets(cTokens []common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.EnterMarkets(&_ComptrollerG4.TransactOpts, cTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_ComptrollerG4 *ComptrollerG4TransactorSession) EnterMarkets(cTokens []common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.EnterMarkets(&_ComptrollerG4.TransactOpts, cTokens)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cTokenAddress) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Transactor) ExitMarket(opts *bind.TransactOpts, cTokenAddress common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "exitMarket", cTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cTokenAddress) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) ExitMarket(cTokenAddress common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.ExitMarket(&_ComptrollerG4.TransactOpts, cTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cTokenAddress) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) ExitMarket(cTokenAddress common.Address) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.ExitMarket(&_ComptrollerG4.TransactOpts, cTokenAddress)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Transactor) LiquidateBorrowAllowed(opts *bind.TransactOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "liquidateBorrowAllowed", cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) LiquidateBorrowAllowed(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.LiquidateBorrowAllowed(&_ComptrollerG4.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) LiquidateBorrowAllowed(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.LiquidateBorrowAllowed(&_ComptrollerG4.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_ComptrollerG4 *ComptrollerG4Transactor) LiquidateBorrowVerify(opts *bind.TransactOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "liquidateBorrowVerify", cTokenBorrowed, cTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_ComptrollerG4 *ComptrollerG4Session) LiquidateBorrowVerify(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.LiquidateBorrowVerify(&_ComptrollerG4.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_ComptrollerG4 *ComptrollerG4TransactorSession) LiquidateBorrowVerify(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.LiquidateBorrowVerify(&_ComptrollerG4.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Transactor) MintAllowed(opts *bind.TransactOpts, cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "mintAllowed", cToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) MintAllowed(cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.MintAllowed(&_ComptrollerG4.TransactOpts, cToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) MintAllowed(cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.MintAllowed(&_ComptrollerG4.TransactOpts, cToken, minter, mintAmount)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_ComptrollerG4 *ComptrollerG4Transactor) MintVerify(opts *bind.TransactOpts, cToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "mintVerify", cToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_ComptrollerG4 *ComptrollerG4Session) MintVerify(cToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.MintVerify(&_ComptrollerG4.TransactOpts, cToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_ComptrollerG4 *ComptrollerG4TransactorSession) MintVerify(cToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.MintVerify(&_ComptrollerG4.TransactOpts, cToken, minter, actualMintAmount, mintTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Transactor) RedeemAllowed(opts *bind.TransactOpts, cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "redeemAllowed", cToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) RedeemAllowed(cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.RedeemAllowed(&_ComptrollerG4.TransactOpts, cToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) RedeemAllowed(cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.RedeemAllowed(&_ComptrollerG4.TransactOpts, cToken, redeemer, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_ComptrollerG4 *ComptrollerG4Transactor) RedeemVerify(opts *bind.TransactOpts, cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "redeemVerify", cToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_ComptrollerG4 *ComptrollerG4Session) RedeemVerify(cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.RedeemVerify(&_ComptrollerG4.TransactOpts, cToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_ComptrollerG4 *ComptrollerG4TransactorSession) RedeemVerify(cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.RedeemVerify(&_ComptrollerG4.TransactOpts, cToken, redeemer, redeemAmount, redeemTokens)
}

// RefreshCompSpeeds is a paid mutator transaction binding the contract method 0x4d8e5037.
//
// Solidity: function refreshCompSpeeds() returns()
func (_ComptrollerG4 *ComptrollerG4Transactor) RefreshCompSpeeds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "refreshCompSpeeds")
}

// RefreshCompSpeeds is a paid mutator transaction binding the contract method 0x4d8e5037.
//
// Solidity: function refreshCompSpeeds() returns()
func (_ComptrollerG4 *ComptrollerG4Session) RefreshCompSpeeds() (*types.Transaction, error) {
	return _ComptrollerG4.Contract.RefreshCompSpeeds(&_ComptrollerG4.TransactOpts)
}

// RefreshCompSpeeds is a paid mutator transaction binding the contract method 0x4d8e5037.
//
// Solidity: function refreshCompSpeeds() returns()
func (_ComptrollerG4 *ComptrollerG4TransactorSession) RefreshCompSpeeds() (*types.Transaction, error) {
	return _ComptrollerG4.Contract.RefreshCompSpeeds(&_ComptrollerG4.TransactOpts)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Transactor) RepayBorrowAllowed(opts *bind.TransactOpts, cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "repayBorrowAllowed", cToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) RepayBorrowAllowed(cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.RepayBorrowAllowed(&_ComptrollerG4.TransactOpts, cToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) RepayBorrowAllowed(cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.RepayBorrowAllowed(&_ComptrollerG4.TransactOpts, cToken, payer, borrower, repayAmount)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_ComptrollerG4 *ComptrollerG4Transactor) RepayBorrowVerify(opts *bind.TransactOpts, cToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "repayBorrowVerify", cToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_ComptrollerG4 *ComptrollerG4Session) RepayBorrowVerify(cToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.RepayBorrowVerify(&_ComptrollerG4.TransactOpts, cToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_ComptrollerG4 *ComptrollerG4TransactorSession) RepayBorrowVerify(cToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.RepayBorrowVerify(&_ComptrollerG4.TransactOpts, cToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Transactor) SeizeAllowed(opts *bind.TransactOpts, cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "seizeAllowed", cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) SeizeAllowed(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SeizeAllowed(&_ComptrollerG4.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) SeizeAllowed(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SeizeAllowed(&_ComptrollerG4.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_ComptrollerG4 *ComptrollerG4Transactor) SeizeVerify(opts *bind.TransactOpts, cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "seizeVerify", cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_ComptrollerG4 *ComptrollerG4Session) SeizeVerify(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SeizeVerify(&_ComptrollerG4.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_ComptrollerG4 *ComptrollerG4TransactorSession) SeizeVerify(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.SeizeVerify(&_ComptrollerG4.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Transactor) TransferAllowed(opts *bind.TransactOpts, cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "transferAllowed", cToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4Session) TransferAllowed(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.TransferAllowed(&_ComptrollerG4.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_ComptrollerG4 *ComptrollerG4TransactorSession) TransferAllowed(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.TransferAllowed(&_ComptrollerG4.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_ComptrollerG4 *ComptrollerG4Transactor) TransferVerify(opts *bind.TransactOpts, cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.contract.Transact(opts, "transferVerify", cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_ComptrollerG4 *ComptrollerG4Session) TransferVerify(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.TransferVerify(&_ComptrollerG4.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_ComptrollerG4 *ComptrollerG4TransactorSession) TransferVerify(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerG4.Contract.TransferVerify(&_ComptrollerG4.TransactOpts, cToken, src, dst, transferTokens)
}

// ComptrollerG4ActionPausedIterator is returned from FilterActionPaused and is used to iterate over the raw logs and unpacked data for ActionPaused events raised by the ComptrollerG4 contract.
type ComptrollerG4ActionPausedIterator struct {
	Event *ComptrollerG4ActionPaused // Event containing the contract specifics and raw log

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
func (it *ComptrollerG4ActionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerG4ActionPaused)
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
		it.Event = new(ComptrollerG4ActionPaused)
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
func (it *ComptrollerG4ActionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerG4ActionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerG4ActionPaused represents a ActionPaused event raised by the ComptrollerG4 contract.
type ComptrollerG4ActionPaused struct {
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused is a free log retrieval operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_ComptrollerG4 *ComptrollerG4Filterer) FilterActionPaused(opts *bind.FilterOpts) (*ComptrollerG4ActionPausedIterator, error) {

	logs, sub, err := _ComptrollerG4.contract.FilterLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4ActionPausedIterator{contract: _ComptrollerG4.contract, event: "ActionPaused", logs: logs, sub: sub}, nil
}

// WatchActionPaused is a free log subscription operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_ComptrollerG4 *ComptrollerG4Filterer) WatchActionPaused(opts *bind.WatchOpts, sink chan<- *ComptrollerG4ActionPaused) (event.Subscription, error) {

	logs, sub, err := _ComptrollerG4.contract.WatchLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerG4ActionPaused)
				if err := _ComptrollerG4.contract.UnpackLog(event, "ActionPaused", log); err != nil {
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

// ParseActionPaused is a log parse operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_ComptrollerG4 *ComptrollerG4Filterer) ParseActionPaused(log types.Log) (*ComptrollerG4ActionPaused, error) {
	event := new(ComptrollerG4ActionPaused)
	if err := _ComptrollerG4.contract.UnpackLog(event, "ActionPaused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerG4ActionPaused0Iterator is returned from FilterActionPaused0 and is used to iterate over the raw logs and unpacked data for ActionPaused0 events raised by the ComptrollerG4 contract.
type ComptrollerG4ActionPaused0Iterator struct {
	Event *ComptrollerG4ActionPaused0 // Event containing the contract specifics and raw log

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
func (it *ComptrollerG4ActionPaused0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerG4ActionPaused0)
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
		it.Event = new(ComptrollerG4ActionPaused0)
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
func (it *ComptrollerG4ActionPaused0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerG4ActionPaused0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerG4ActionPaused0 represents a ActionPaused0 event raised by the ComptrollerG4 contract.
type ComptrollerG4ActionPaused0 struct {
	CToken     common.Address
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused0 is a free log retrieval operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address cToken, string action, bool pauseState)
func (_ComptrollerG4 *ComptrollerG4Filterer) FilterActionPaused0(opts *bind.FilterOpts) (*ComptrollerG4ActionPaused0Iterator, error) {

	logs, sub, err := _ComptrollerG4.contract.FilterLogs(opts, "ActionPaused0")
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4ActionPaused0Iterator{contract: _ComptrollerG4.contract, event: "ActionPaused0", logs: logs, sub: sub}, nil
}

// WatchActionPaused0 is a free log subscription operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address cToken, string action, bool pauseState)
func (_ComptrollerG4 *ComptrollerG4Filterer) WatchActionPaused0(opts *bind.WatchOpts, sink chan<- *ComptrollerG4ActionPaused0) (event.Subscription, error) {

	logs, sub, err := _ComptrollerG4.contract.WatchLogs(opts, "ActionPaused0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerG4ActionPaused0)
				if err := _ComptrollerG4.contract.UnpackLog(event, "ActionPaused0", log); err != nil {
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

// ParseActionPaused0 is a log parse operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address cToken, string action, bool pauseState)
func (_ComptrollerG4 *ComptrollerG4Filterer) ParseActionPaused0(log types.Log) (*ComptrollerG4ActionPaused0, error) {
	event := new(ComptrollerG4ActionPaused0)
	if err := _ComptrollerG4.contract.UnpackLog(event, "ActionPaused0", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerG4CompSpeedUpdatedIterator is returned from FilterCompSpeedUpdated and is used to iterate over the raw logs and unpacked data for CompSpeedUpdated events raised by the ComptrollerG4 contract.
type ComptrollerG4CompSpeedUpdatedIterator struct {
	Event *ComptrollerG4CompSpeedUpdated // Event containing the contract specifics and raw log

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
func (it *ComptrollerG4CompSpeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerG4CompSpeedUpdated)
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
		it.Event = new(ComptrollerG4CompSpeedUpdated)
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
func (it *ComptrollerG4CompSpeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerG4CompSpeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerG4CompSpeedUpdated represents a CompSpeedUpdated event raised by the ComptrollerG4 contract.
type ComptrollerG4CompSpeedUpdated struct {
	CToken   common.Address
	NewSpeed *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCompSpeedUpdated is a free log retrieval operation binding the contract event 0x2ab93f65628379309f36cb125e90d7c902454a545c4f8b8cb0794af75c24b807.
//
// Solidity: event CompSpeedUpdated(address indexed cToken, uint256 newSpeed)
func (_ComptrollerG4 *ComptrollerG4Filterer) FilterCompSpeedUpdated(opts *bind.FilterOpts, cToken []common.Address) (*ComptrollerG4CompSpeedUpdatedIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _ComptrollerG4.contract.FilterLogs(opts, "CompSpeedUpdated", cTokenRule)
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4CompSpeedUpdatedIterator{contract: _ComptrollerG4.contract, event: "CompSpeedUpdated", logs: logs, sub: sub}, nil
}

// WatchCompSpeedUpdated is a free log subscription operation binding the contract event 0x2ab93f65628379309f36cb125e90d7c902454a545c4f8b8cb0794af75c24b807.
//
// Solidity: event CompSpeedUpdated(address indexed cToken, uint256 newSpeed)
func (_ComptrollerG4 *ComptrollerG4Filterer) WatchCompSpeedUpdated(opts *bind.WatchOpts, sink chan<- *ComptrollerG4CompSpeedUpdated, cToken []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _ComptrollerG4.contract.WatchLogs(opts, "CompSpeedUpdated", cTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerG4CompSpeedUpdated)
				if err := _ComptrollerG4.contract.UnpackLog(event, "CompSpeedUpdated", log); err != nil {
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

// ParseCompSpeedUpdated is a log parse operation binding the contract event 0x2ab93f65628379309f36cb125e90d7c902454a545c4f8b8cb0794af75c24b807.
//
// Solidity: event CompSpeedUpdated(address indexed cToken, uint256 newSpeed)
func (_ComptrollerG4 *ComptrollerG4Filterer) ParseCompSpeedUpdated(log types.Log) (*ComptrollerG4CompSpeedUpdated, error) {
	event := new(ComptrollerG4CompSpeedUpdated)
	if err := _ComptrollerG4.contract.UnpackLog(event, "CompSpeedUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerG4DistributedBorrowerCompIterator is returned from FilterDistributedBorrowerComp and is used to iterate over the raw logs and unpacked data for DistributedBorrowerComp events raised by the ComptrollerG4 contract.
type ComptrollerG4DistributedBorrowerCompIterator struct {
	Event *ComptrollerG4DistributedBorrowerComp // Event containing the contract specifics and raw log

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
func (it *ComptrollerG4DistributedBorrowerCompIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerG4DistributedBorrowerComp)
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
		it.Event = new(ComptrollerG4DistributedBorrowerComp)
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
func (it *ComptrollerG4DistributedBorrowerCompIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerG4DistributedBorrowerCompIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerG4DistributedBorrowerComp represents a DistributedBorrowerComp event raised by the ComptrollerG4 contract.
type ComptrollerG4DistributedBorrowerComp struct {
	CToken          common.Address
	Borrower        common.Address
	CompDelta       *big.Int
	CompBorrowIndex *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributedBorrowerComp is a free log retrieval operation binding the contract event 0x1fc3ecc087d8d2d15e23d0032af5a47059c3892d003d8e139fdcb6bb327c99a6.
//
// Solidity: event DistributedBorrowerComp(address indexed cToken, address indexed borrower, uint256 compDelta, uint256 compBorrowIndex)
func (_ComptrollerG4 *ComptrollerG4Filterer) FilterDistributedBorrowerComp(opts *bind.FilterOpts, cToken []common.Address, borrower []common.Address) (*ComptrollerG4DistributedBorrowerCompIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _ComptrollerG4.contract.FilterLogs(opts, "DistributedBorrowerComp", cTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4DistributedBorrowerCompIterator{contract: _ComptrollerG4.contract, event: "DistributedBorrowerComp", logs: logs, sub: sub}, nil
}

// WatchDistributedBorrowerComp is a free log subscription operation binding the contract event 0x1fc3ecc087d8d2d15e23d0032af5a47059c3892d003d8e139fdcb6bb327c99a6.
//
// Solidity: event DistributedBorrowerComp(address indexed cToken, address indexed borrower, uint256 compDelta, uint256 compBorrowIndex)
func (_ComptrollerG4 *ComptrollerG4Filterer) WatchDistributedBorrowerComp(opts *bind.WatchOpts, sink chan<- *ComptrollerG4DistributedBorrowerComp, cToken []common.Address, borrower []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _ComptrollerG4.contract.WatchLogs(opts, "DistributedBorrowerComp", cTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerG4DistributedBorrowerComp)
				if err := _ComptrollerG4.contract.UnpackLog(event, "DistributedBorrowerComp", log); err != nil {
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

// ParseDistributedBorrowerComp is a log parse operation binding the contract event 0x1fc3ecc087d8d2d15e23d0032af5a47059c3892d003d8e139fdcb6bb327c99a6.
//
// Solidity: event DistributedBorrowerComp(address indexed cToken, address indexed borrower, uint256 compDelta, uint256 compBorrowIndex)
func (_ComptrollerG4 *ComptrollerG4Filterer) ParseDistributedBorrowerComp(log types.Log) (*ComptrollerG4DistributedBorrowerComp, error) {
	event := new(ComptrollerG4DistributedBorrowerComp)
	if err := _ComptrollerG4.contract.UnpackLog(event, "DistributedBorrowerComp", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerG4DistributedSupplierCompIterator is returned from FilterDistributedSupplierComp and is used to iterate over the raw logs and unpacked data for DistributedSupplierComp events raised by the ComptrollerG4 contract.
type ComptrollerG4DistributedSupplierCompIterator struct {
	Event *ComptrollerG4DistributedSupplierComp // Event containing the contract specifics and raw log

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
func (it *ComptrollerG4DistributedSupplierCompIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerG4DistributedSupplierComp)
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
		it.Event = new(ComptrollerG4DistributedSupplierComp)
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
func (it *ComptrollerG4DistributedSupplierCompIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerG4DistributedSupplierCompIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerG4DistributedSupplierComp represents a DistributedSupplierComp event raised by the ComptrollerG4 contract.
type ComptrollerG4DistributedSupplierComp struct {
	CToken          common.Address
	Supplier        common.Address
	CompDelta       *big.Int
	CompSupplyIndex *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributedSupplierComp is a free log retrieval operation binding the contract event 0x2caecd17d02f56fa897705dcc740da2d237c373f70686f4e0d9bd3bf0400ea7a.
//
// Solidity: event DistributedSupplierComp(address indexed cToken, address indexed supplier, uint256 compDelta, uint256 compSupplyIndex)
func (_ComptrollerG4 *ComptrollerG4Filterer) FilterDistributedSupplierComp(opts *bind.FilterOpts, cToken []common.Address, supplier []common.Address) (*ComptrollerG4DistributedSupplierCompIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _ComptrollerG4.contract.FilterLogs(opts, "DistributedSupplierComp", cTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4DistributedSupplierCompIterator{contract: _ComptrollerG4.contract, event: "DistributedSupplierComp", logs: logs, sub: sub}, nil
}

// WatchDistributedSupplierComp is a free log subscription operation binding the contract event 0x2caecd17d02f56fa897705dcc740da2d237c373f70686f4e0d9bd3bf0400ea7a.
//
// Solidity: event DistributedSupplierComp(address indexed cToken, address indexed supplier, uint256 compDelta, uint256 compSupplyIndex)
func (_ComptrollerG4 *ComptrollerG4Filterer) WatchDistributedSupplierComp(opts *bind.WatchOpts, sink chan<- *ComptrollerG4DistributedSupplierComp, cToken []common.Address, supplier []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _ComptrollerG4.contract.WatchLogs(opts, "DistributedSupplierComp", cTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerG4DistributedSupplierComp)
				if err := _ComptrollerG4.contract.UnpackLog(event, "DistributedSupplierComp", log); err != nil {
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

// ParseDistributedSupplierComp is a log parse operation binding the contract event 0x2caecd17d02f56fa897705dcc740da2d237c373f70686f4e0d9bd3bf0400ea7a.
//
// Solidity: event DistributedSupplierComp(address indexed cToken, address indexed supplier, uint256 compDelta, uint256 compSupplyIndex)
func (_ComptrollerG4 *ComptrollerG4Filterer) ParseDistributedSupplierComp(log types.Log) (*ComptrollerG4DistributedSupplierComp, error) {
	event := new(ComptrollerG4DistributedSupplierComp)
	if err := _ComptrollerG4.contract.UnpackLog(event, "DistributedSupplierComp", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerG4FailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the ComptrollerG4 contract.
type ComptrollerG4FailureIterator struct {
	Event *ComptrollerG4Failure // Event containing the contract specifics and raw log

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
func (it *ComptrollerG4FailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerG4Failure)
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
		it.Event = new(ComptrollerG4Failure)
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
func (it *ComptrollerG4FailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerG4FailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerG4Failure represents a Failure event raised by the ComptrollerG4 contract.
type ComptrollerG4Failure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_ComptrollerG4 *ComptrollerG4Filterer) FilterFailure(opts *bind.FilterOpts) (*ComptrollerG4FailureIterator, error) {

	logs, sub, err := _ComptrollerG4.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4FailureIterator{contract: _ComptrollerG4.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_ComptrollerG4 *ComptrollerG4Filterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *ComptrollerG4Failure) (event.Subscription, error) {

	logs, sub, err := _ComptrollerG4.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerG4Failure)
				if err := _ComptrollerG4.contract.UnpackLog(event, "Failure", log); err != nil {
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
func (_ComptrollerG4 *ComptrollerG4Filterer) ParseFailure(log types.Log) (*ComptrollerG4Failure, error) {
	event := new(ComptrollerG4Failure)
	if err := _ComptrollerG4.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerG4MarketCompedIterator is returned from FilterMarketComped and is used to iterate over the raw logs and unpacked data for MarketComped events raised by the ComptrollerG4 contract.
type ComptrollerG4MarketCompedIterator struct {
	Event *ComptrollerG4MarketComped // Event containing the contract specifics and raw log

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
func (it *ComptrollerG4MarketCompedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerG4MarketComped)
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
		it.Event = new(ComptrollerG4MarketComped)
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
func (it *ComptrollerG4MarketCompedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerG4MarketCompedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerG4MarketComped represents a MarketComped event raised by the ComptrollerG4 contract.
type ComptrollerG4MarketComped struct {
	CToken   common.Address
	IsComped bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketComped is a free log retrieval operation binding the contract event 0x93c1f3e36ed71139f466a4ce8c9751790e2e33f5afb2df0dcfb3aeabe55d5aa2.
//
// Solidity: event MarketComped(address cToken, bool isComped)
func (_ComptrollerG4 *ComptrollerG4Filterer) FilterMarketComped(opts *bind.FilterOpts) (*ComptrollerG4MarketCompedIterator, error) {

	logs, sub, err := _ComptrollerG4.contract.FilterLogs(opts, "MarketComped")
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4MarketCompedIterator{contract: _ComptrollerG4.contract, event: "MarketComped", logs: logs, sub: sub}, nil
}

// WatchMarketComped is a free log subscription operation binding the contract event 0x93c1f3e36ed71139f466a4ce8c9751790e2e33f5afb2df0dcfb3aeabe55d5aa2.
//
// Solidity: event MarketComped(address cToken, bool isComped)
func (_ComptrollerG4 *ComptrollerG4Filterer) WatchMarketComped(opts *bind.WatchOpts, sink chan<- *ComptrollerG4MarketComped) (event.Subscription, error) {

	logs, sub, err := _ComptrollerG4.contract.WatchLogs(opts, "MarketComped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerG4MarketComped)
				if err := _ComptrollerG4.contract.UnpackLog(event, "MarketComped", log); err != nil {
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

// ParseMarketComped is a log parse operation binding the contract event 0x93c1f3e36ed71139f466a4ce8c9751790e2e33f5afb2df0dcfb3aeabe55d5aa2.
//
// Solidity: event MarketComped(address cToken, bool isComped)
func (_ComptrollerG4 *ComptrollerG4Filterer) ParseMarketComped(log types.Log) (*ComptrollerG4MarketComped, error) {
	event := new(ComptrollerG4MarketComped)
	if err := _ComptrollerG4.contract.UnpackLog(event, "MarketComped", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerG4MarketEnteredIterator is returned from FilterMarketEntered and is used to iterate over the raw logs and unpacked data for MarketEntered events raised by the ComptrollerG4 contract.
type ComptrollerG4MarketEnteredIterator struct {
	Event *ComptrollerG4MarketEntered // Event containing the contract specifics and raw log

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
func (it *ComptrollerG4MarketEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerG4MarketEntered)
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
		it.Event = new(ComptrollerG4MarketEntered)
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
func (it *ComptrollerG4MarketEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerG4MarketEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerG4MarketEntered represents a MarketEntered event raised by the ComptrollerG4 contract.
type ComptrollerG4MarketEntered struct {
	CToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketEntered is a free log retrieval operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address cToken, address account)
func (_ComptrollerG4 *ComptrollerG4Filterer) FilterMarketEntered(opts *bind.FilterOpts) (*ComptrollerG4MarketEnteredIterator, error) {

	logs, sub, err := _ComptrollerG4.contract.FilterLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4MarketEnteredIterator{contract: _ComptrollerG4.contract, event: "MarketEntered", logs: logs, sub: sub}, nil
}

// WatchMarketEntered is a free log subscription operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address cToken, address account)
func (_ComptrollerG4 *ComptrollerG4Filterer) WatchMarketEntered(opts *bind.WatchOpts, sink chan<- *ComptrollerG4MarketEntered) (event.Subscription, error) {

	logs, sub, err := _ComptrollerG4.contract.WatchLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerG4MarketEntered)
				if err := _ComptrollerG4.contract.UnpackLog(event, "MarketEntered", log); err != nil {
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

// ParseMarketEntered is a log parse operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address cToken, address account)
func (_ComptrollerG4 *ComptrollerG4Filterer) ParseMarketEntered(log types.Log) (*ComptrollerG4MarketEntered, error) {
	event := new(ComptrollerG4MarketEntered)
	if err := _ComptrollerG4.contract.UnpackLog(event, "MarketEntered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerG4MarketExitedIterator is returned from FilterMarketExited and is used to iterate over the raw logs and unpacked data for MarketExited events raised by the ComptrollerG4 contract.
type ComptrollerG4MarketExitedIterator struct {
	Event *ComptrollerG4MarketExited // Event containing the contract specifics and raw log

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
func (it *ComptrollerG4MarketExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerG4MarketExited)
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
		it.Event = new(ComptrollerG4MarketExited)
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
func (it *ComptrollerG4MarketExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerG4MarketExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerG4MarketExited represents a MarketExited event raised by the ComptrollerG4 contract.
type ComptrollerG4MarketExited struct {
	CToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketExited is a free log retrieval operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address cToken, address account)
func (_ComptrollerG4 *ComptrollerG4Filterer) FilterMarketExited(opts *bind.FilterOpts) (*ComptrollerG4MarketExitedIterator, error) {

	logs, sub, err := _ComptrollerG4.contract.FilterLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4MarketExitedIterator{contract: _ComptrollerG4.contract, event: "MarketExited", logs: logs, sub: sub}, nil
}

// WatchMarketExited is a free log subscription operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address cToken, address account)
func (_ComptrollerG4 *ComptrollerG4Filterer) WatchMarketExited(opts *bind.WatchOpts, sink chan<- *ComptrollerG4MarketExited) (event.Subscription, error) {

	logs, sub, err := _ComptrollerG4.contract.WatchLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerG4MarketExited)
				if err := _ComptrollerG4.contract.UnpackLog(event, "MarketExited", log); err != nil {
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

// ParseMarketExited is a log parse operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address cToken, address account)
func (_ComptrollerG4 *ComptrollerG4Filterer) ParseMarketExited(log types.Log) (*ComptrollerG4MarketExited, error) {
	event := new(ComptrollerG4MarketExited)
	if err := _ComptrollerG4.contract.UnpackLog(event, "MarketExited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerG4MarketListedIterator is returned from FilterMarketListed and is used to iterate over the raw logs and unpacked data for MarketListed events raised by the ComptrollerG4 contract.
type ComptrollerG4MarketListedIterator struct {
	Event *ComptrollerG4MarketListed // Event containing the contract specifics and raw log

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
func (it *ComptrollerG4MarketListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerG4MarketListed)
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
		it.Event = new(ComptrollerG4MarketListed)
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
func (it *ComptrollerG4MarketListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerG4MarketListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerG4MarketListed represents a MarketListed event raised by the ComptrollerG4 contract.
type ComptrollerG4MarketListed struct {
	CToken common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketListed is a free log retrieval operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address cToken)
func (_ComptrollerG4 *ComptrollerG4Filterer) FilterMarketListed(opts *bind.FilterOpts) (*ComptrollerG4MarketListedIterator, error) {

	logs, sub, err := _ComptrollerG4.contract.FilterLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4MarketListedIterator{contract: _ComptrollerG4.contract, event: "MarketListed", logs: logs, sub: sub}, nil
}

// WatchMarketListed is a free log subscription operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address cToken)
func (_ComptrollerG4 *ComptrollerG4Filterer) WatchMarketListed(opts *bind.WatchOpts, sink chan<- *ComptrollerG4MarketListed) (event.Subscription, error) {

	logs, sub, err := _ComptrollerG4.contract.WatchLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerG4MarketListed)
				if err := _ComptrollerG4.contract.UnpackLog(event, "MarketListed", log); err != nil {
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

// ParseMarketListed is a log parse operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address cToken)
func (_ComptrollerG4 *ComptrollerG4Filterer) ParseMarketListed(log types.Log) (*ComptrollerG4MarketListed, error) {
	event := new(ComptrollerG4MarketListed)
	if err := _ComptrollerG4.contract.UnpackLog(event, "MarketListed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerG4NewCloseFactorIterator is returned from FilterNewCloseFactor and is used to iterate over the raw logs and unpacked data for NewCloseFactor events raised by the ComptrollerG4 contract.
type ComptrollerG4NewCloseFactorIterator struct {
	Event *ComptrollerG4NewCloseFactor // Event containing the contract specifics and raw log

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
func (it *ComptrollerG4NewCloseFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerG4NewCloseFactor)
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
		it.Event = new(ComptrollerG4NewCloseFactor)
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
func (it *ComptrollerG4NewCloseFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerG4NewCloseFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerG4NewCloseFactor represents a NewCloseFactor event raised by the ComptrollerG4 contract.
type ComptrollerG4NewCloseFactor struct {
	OldCloseFactorMantissa *big.Int
	NewCloseFactorMantissa *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewCloseFactor is a free log retrieval operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_ComptrollerG4 *ComptrollerG4Filterer) FilterNewCloseFactor(opts *bind.FilterOpts) (*ComptrollerG4NewCloseFactorIterator, error) {

	logs, sub, err := _ComptrollerG4.contract.FilterLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4NewCloseFactorIterator{contract: _ComptrollerG4.contract, event: "NewCloseFactor", logs: logs, sub: sub}, nil
}

// WatchNewCloseFactor is a free log subscription operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_ComptrollerG4 *ComptrollerG4Filterer) WatchNewCloseFactor(opts *bind.WatchOpts, sink chan<- *ComptrollerG4NewCloseFactor) (event.Subscription, error) {

	logs, sub, err := _ComptrollerG4.contract.WatchLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerG4NewCloseFactor)
				if err := _ComptrollerG4.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
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

// ParseNewCloseFactor is a log parse operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_ComptrollerG4 *ComptrollerG4Filterer) ParseNewCloseFactor(log types.Log) (*ComptrollerG4NewCloseFactor, error) {
	event := new(ComptrollerG4NewCloseFactor)
	if err := _ComptrollerG4.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerG4NewCollateralFactorIterator is returned from FilterNewCollateralFactor and is used to iterate over the raw logs and unpacked data for NewCollateralFactor events raised by the ComptrollerG4 contract.
type ComptrollerG4NewCollateralFactorIterator struct {
	Event *ComptrollerG4NewCollateralFactor // Event containing the contract specifics and raw log

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
func (it *ComptrollerG4NewCollateralFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerG4NewCollateralFactor)
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
		it.Event = new(ComptrollerG4NewCollateralFactor)
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
func (it *ComptrollerG4NewCollateralFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerG4NewCollateralFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerG4NewCollateralFactor represents a NewCollateralFactor event raised by the ComptrollerG4 contract.
type ComptrollerG4NewCollateralFactor struct {
	CToken                      common.Address
	OldCollateralFactorMantissa *big.Int
	NewCollateralFactorMantissa *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterNewCollateralFactor is a free log retrieval operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address cToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_ComptrollerG4 *ComptrollerG4Filterer) FilterNewCollateralFactor(opts *bind.FilterOpts) (*ComptrollerG4NewCollateralFactorIterator, error) {

	logs, sub, err := _ComptrollerG4.contract.FilterLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4NewCollateralFactorIterator{contract: _ComptrollerG4.contract, event: "NewCollateralFactor", logs: logs, sub: sub}, nil
}

// WatchNewCollateralFactor is a free log subscription operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address cToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_ComptrollerG4 *ComptrollerG4Filterer) WatchNewCollateralFactor(opts *bind.WatchOpts, sink chan<- *ComptrollerG4NewCollateralFactor) (event.Subscription, error) {

	logs, sub, err := _ComptrollerG4.contract.WatchLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerG4NewCollateralFactor)
				if err := _ComptrollerG4.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
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

// ParseNewCollateralFactor is a log parse operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address cToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_ComptrollerG4 *ComptrollerG4Filterer) ParseNewCollateralFactor(log types.Log) (*ComptrollerG4NewCollateralFactor, error) {
	event := new(ComptrollerG4NewCollateralFactor)
	if err := _ComptrollerG4.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerG4NewCompRateIterator is returned from FilterNewCompRate and is used to iterate over the raw logs and unpacked data for NewCompRate events raised by the ComptrollerG4 contract.
type ComptrollerG4NewCompRateIterator struct {
	Event *ComptrollerG4NewCompRate // Event containing the contract specifics and raw log

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
func (it *ComptrollerG4NewCompRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerG4NewCompRate)
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
		it.Event = new(ComptrollerG4NewCompRate)
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
func (it *ComptrollerG4NewCompRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerG4NewCompRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerG4NewCompRate represents a NewCompRate event raised by the ComptrollerG4 contract.
type ComptrollerG4NewCompRate struct {
	OldCompRate *big.Int
	NewCompRate *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewCompRate is a free log retrieval operation binding the contract event 0xc227c9272633c3a307d9845bf2bc2509cefb20d655b5f3c1002d8e1e3f22c8b0.
//
// Solidity: event NewCompRate(uint256 oldCompRate, uint256 newCompRate)
func (_ComptrollerG4 *ComptrollerG4Filterer) FilterNewCompRate(opts *bind.FilterOpts) (*ComptrollerG4NewCompRateIterator, error) {

	logs, sub, err := _ComptrollerG4.contract.FilterLogs(opts, "NewCompRate")
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4NewCompRateIterator{contract: _ComptrollerG4.contract, event: "NewCompRate", logs: logs, sub: sub}, nil
}

// WatchNewCompRate is a free log subscription operation binding the contract event 0xc227c9272633c3a307d9845bf2bc2509cefb20d655b5f3c1002d8e1e3f22c8b0.
//
// Solidity: event NewCompRate(uint256 oldCompRate, uint256 newCompRate)
func (_ComptrollerG4 *ComptrollerG4Filterer) WatchNewCompRate(opts *bind.WatchOpts, sink chan<- *ComptrollerG4NewCompRate) (event.Subscription, error) {

	logs, sub, err := _ComptrollerG4.contract.WatchLogs(opts, "NewCompRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerG4NewCompRate)
				if err := _ComptrollerG4.contract.UnpackLog(event, "NewCompRate", log); err != nil {
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

// ParseNewCompRate is a log parse operation binding the contract event 0xc227c9272633c3a307d9845bf2bc2509cefb20d655b5f3c1002d8e1e3f22c8b0.
//
// Solidity: event NewCompRate(uint256 oldCompRate, uint256 newCompRate)
func (_ComptrollerG4 *ComptrollerG4Filterer) ParseNewCompRate(log types.Log) (*ComptrollerG4NewCompRate, error) {
	event := new(ComptrollerG4NewCompRate)
	if err := _ComptrollerG4.contract.UnpackLog(event, "NewCompRate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerG4NewLiquidationIncentiveIterator is returned from FilterNewLiquidationIncentive and is used to iterate over the raw logs and unpacked data for NewLiquidationIncentive events raised by the ComptrollerG4 contract.
type ComptrollerG4NewLiquidationIncentiveIterator struct {
	Event *ComptrollerG4NewLiquidationIncentive // Event containing the contract specifics and raw log

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
func (it *ComptrollerG4NewLiquidationIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerG4NewLiquidationIncentive)
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
		it.Event = new(ComptrollerG4NewLiquidationIncentive)
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
func (it *ComptrollerG4NewLiquidationIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerG4NewLiquidationIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerG4NewLiquidationIncentive represents a NewLiquidationIncentive event raised by the ComptrollerG4 contract.
type ComptrollerG4NewLiquidationIncentive struct {
	OldLiquidationIncentiveMantissa *big.Int
	NewLiquidationIncentiveMantissa *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterNewLiquidationIncentive is a free log retrieval operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_ComptrollerG4 *ComptrollerG4Filterer) FilterNewLiquidationIncentive(opts *bind.FilterOpts) (*ComptrollerG4NewLiquidationIncentiveIterator, error) {

	logs, sub, err := _ComptrollerG4.contract.FilterLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4NewLiquidationIncentiveIterator{contract: _ComptrollerG4.contract, event: "NewLiquidationIncentive", logs: logs, sub: sub}, nil
}

// WatchNewLiquidationIncentive is a free log subscription operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_ComptrollerG4 *ComptrollerG4Filterer) WatchNewLiquidationIncentive(opts *bind.WatchOpts, sink chan<- *ComptrollerG4NewLiquidationIncentive) (event.Subscription, error) {

	logs, sub, err := _ComptrollerG4.contract.WatchLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerG4NewLiquidationIncentive)
				if err := _ComptrollerG4.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
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

// ParseNewLiquidationIncentive is a log parse operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_ComptrollerG4 *ComptrollerG4Filterer) ParseNewLiquidationIncentive(log types.Log) (*ComptrollerG4NewLiquidationIncentive, error) {
	event := new(ComptrollerG4NewLiquidationIncentive)
	if err := _ComptrollerG4.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerG4NewMaxAssetsIterator is returned from FilterNewMaxAssets and is used to iterate over the raw logs and unpacked data for NewMaxAssets events raised by the ComptrollerG4 contract.
type ComptrollerG4NewMaxAssetsIterator struct {
	Event *ComptrollerG4NewMaxAssets // Event containing the contract specifics and raw log

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
func (it *ComptrollerG4NewMaxAssetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerG4NewMaxAssets)
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
		it.Event = new(ComptrollerG4NewMaxAssets)
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
func (it *ComptrollerG4NewMaxAssetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerG4NewMaxAssetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerG4NewMaxAssets represents a NewMaxAssets event raised by the ComptrollerG4 contract.
type ComptrollerG4NewMaxAssets struct {
	OldMaxAssets *big.Int
	NewMaxAssets *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewMaxAssets is a free log retrieval operation binding the contract event 0x7093cf1eb653f749c3ff531d6df7f92764536a7fa0d13530cd26e070780c32ea.
//
// Solidity: event NewMaxAssets(uint256 oldMaxAssets, uint256 newMaxAssets)
func (_ComptrollerG4 *ComptrollerG4Filterer) FilterNewMaxAssets(opts *bind.FilterOpts) (*ComptrollerG4NewMaxAssetsIterator, error) {

	logs, sub, err := _ComptrollerG4.contract.FilterLogs(opts, "NewMaxAssets")
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4NewMaxAssetsIterator{contract: _ComptrollerG4.contract, event: "NewMaxAssets", logs: logs, sub: sub}, nil
}

// WatchNewMaxAssets is a free log subscription operation binding the contract event 0x7093cf1eb653f749c3ff531d6df7f92764536a7fa0d13530cd26e070780c32ea.
//
// Solidity: event NewMaxAssets(uint256 oldMaxAssets, uint256 newMaxAssets)
func (_ComptrollerG4 *ComptrollerG4Filterer) WatchNewMaxAssets(opts *bind.WatchOpts, sink chan<- *ComptrollerG4NewMaxAssets) (event.Subscription, error) {

	logs, sub, err := _ComptrollerG4.contract.WatchLogs(opts, "NewMaxAssets")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerG4NewMaxAssets)
				if err := _ComptrollerG4.contract.UnpackLog(event, "NewMaxAssets", log); err != nil {
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

// ParseNewMaxAssets is a log parse operation binding the contract event 0x7093cf1eb653f749c3ff531d6df7f92764536a7fa0d13530cd26e070780c32ea.
//
// Solidity: event NewMaxAssets(uint256 oldMaxAssets, uint256 newMaxAssets)
func (_ComptrollerG4 *ComptrollerG4Filterer) ParseNewMaxAssets(log types.Log) (*ComptrollerG4NewMaxAssets, error) {
	event := new(ComptrollerG4NewMaxAssets)
	if err := _ComptrollerG4.contract.UnpackLog(event, "NewMaxAssets", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerG4NewPauseGuardianIterator is returned from FilterNewPauseGuardian and is used to iterate over the raw logs and unpacked data for NewPauseGuardian events raised by the ComptrollerG4 contract.
type ComptrollerG4NewPauseGuardianIterator struct {
	Event *ComptrollerG4NewPauseGuardian // Event containing the contract specifics and raw log

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
func (it *ComptrollerG4NewPauseGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerG4NewPauseGuardian)
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
		it.Event = new(ComptrollerG4NewPauseGuardian)
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
func (it *ComptrollerG4NewPauseGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerG4NewPauseGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerG4NewPauseGuardian represents a NewPauseGuardian event raised by the ComptrollerG4 contract.
type ComptrollerG4NewPauseGuardian struct {
	OldPauseGuardian common.Address
	NewPauseGuardian common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPauseGuardian is a free log retrieval operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_ComptrollerG4 *ComptrollerG4Filterer) FilterNewPauseGuardian(opts *bind.FilterOpts) (*ComptrollerG4NewPauseGuardianIterator, error) {

	logs, sub, err := _ComptrollerG4.contract.FilterLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4NewPauseGuardianIterator{contract: _ComptrollerG4.contract, event: "NewPauseGuardian", logs: logs, sub: sub}, nil
}

// WatchNewPauseGuardian is a free log subscription operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_ComptrollerG4 *ComptrollerG4Filterer) WatchNewPauseGuardian(opts *bind.WatchOpts, sink chan<- *ComptrollerG4NewPauseGuardian) (event.Subscription, error) {

	logs, sub, err := _ComptrollerG4.contract.WatchLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerG4NewPauseGuardian)
				if err := _ComptrollerG4.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
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

// ParseNewPauseGuardian is a log parse operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_ComptrollerG4 *ComptrollerG4Filterer) ParseNewPauseGuardian(log types.Log) (*ComptrollerG4NewPauseGuardian, error) {
	event := new(ComptrollerG4NewPauseGuardian)
	if err := _ComptrollerG4.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComptrollerG4NewPriceOracleIterator is returned from FilterNewPriceOracle and is used to iterate over the raw logs and unpacked data for NewPriceOracle events raised by the ComptrollerG4 contract.
type ComptrollerG4NewPriceOracleIterator struct {
	Event *ComptrollerG4NewPriceOracle // Event containing the contract specifics and raw log

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
func (it *ComptrollerG4NewPriceOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerG4NewPriceOracle)
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
		it.Event = new(ComptrollerG4NewPriceOracle)
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
func (it *ComptrollerG4NewPriceOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerG4NewPriceOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerG4NewPriceOracle represents a NewPriceOracle event raised by the ComptrollerG4 contract.
type ComptrollerG4NewPriceOracle struct {
	OldPriceOracle common.Address
	NewPriceOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewPriceOracle is a free log retrieval operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_ComptrollerG4 *ComptrollerG4Filterer) FilterNewPriceOracle(opts *bind.FilterOpts) (*ComptrollerG4NewPriceOracleIterator, error) {

	logs, sub, err := _ComptrollerG4.contract.FilterLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return &ComptrollerG4NewPriceOracleIterator{contract: _ComptrollerG4.contract, event: "NewPriceOracle", logs: logs, sub: sub}, nil
}

// WatchNewPriceOracle is a free log subscription operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_ComptrollerG4 *ComptrollerG4Filterer) WatchNewPriceOracle(opts *bind.WatchOpts, sink chan<- *ComptrollerG4NewPriceOracle) (event.Subscription, error) {

	logs, sub, err := _ComptrollerG4.contract.WatchLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerG4NewPriceOracle)
				if err := _ComptrollerG4.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
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

// ParseNewPriceOracle is a log parse operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_ComptrollerG4 *ComptrollerG4Filterer) ParseNewPriceOracle(log types.Log) (*ComptrollerG4NewPriceOracle, error) {
	event := new(ComptrollerG4NewPriceOracle)
	if err := _ComptrollerG4.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
		return nil, err
	}
	return event, nil
}
