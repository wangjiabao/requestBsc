// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package service

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StakingV9Record is an auto generated low-level Go binding around an user-defined struct.
type StakingV9Record struct {
	StakeTime   *big.Int
	Amount      *big.Int
	StakeIndex  uint8
	UnstakeTime *big.Int
	Reward      *big.Int
	RestakeTime *big.Int
	CbNewOrder  bool
	Status      uint8
}

// StakeNewMetaData contains all meta data concerning the StakeNew contract.
var StakeNewMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"c\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"routeAddress1_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"routeAddress2_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initAccount_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"freeUsdtCollector_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dropBpsThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"recoveryDuration\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeRatePermille\",\"type\":\"uint256\"}],\"name\":\"CircuitBreakerParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"circuitBreakerTime\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"newHighTime\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentDropBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"action\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roundCount\",\"type\":\"uint256\"}],\"name\":\"CircuitBreakerStateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"stakeLow\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"stakeMid\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"stakeHigh\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"unstakeLow\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"unstakeMid\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"unstakeHigh\",\"type\":\"uint16\"}],\"name\":\"DailyLimitRatesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"DailyQueueCancelLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dividendSystem\",\"type\":\"address\"}],\"name\":\"DividendSystemUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousCollector\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCollector\",\"type\":\"address\"}],\"name\":\"FreeUsdtCollectorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FreeUsdtWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gateway\",\"type\":\"address\"}],\"name\":\"GatewaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousInitAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newInitAccount\",\"type\":\"address\"}],\"name\":\"InitAccountTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userSystem\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dividendSystem\",\"type\":\"address\"}],\"name\":\"InitSystemsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketingAddress\",\"type\":\"address\"}],\"name\":\"MarketingAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depth\",\"type\":\"uint256\"}],\"name\":\"MaxReferralDepthUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"queueIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"amount\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"stakeIndex\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"queuedAt\",\"type\":\"uint40\"}],\"name\":\"QueueAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"queueIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"amount\",\"type\":\"uint160\"}],\"name\":\"QueueCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"queueIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"amount\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"stakeIndex\",\"type\":\"uint8\"}],\"name\":\"QueueProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"queueIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"amount\",\"type\":\"uint160\"}],\"name\":\"QueueRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"Restaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"reward\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RewardBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"rewardEcoBpsTwo\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"rewardDirectBps\",\"type\":\"uint16\"}],\"name\":\"RewardCbNewOrderConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"feeTotalBps\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"ecoBps\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"gameBps\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"globalBps\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"s7Bps\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"teamRewardMaxBps\",\"type\":\"uint16\"}],\"name\":\"RewardFeeConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"reward\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ecoAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gameAddress\",\"type\":\"address\"}],\"name\":\"RewardRouteUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"StageThresholdsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"amount\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"duration\",\"type\":\"uint40\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"amount\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"reward\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"ttl\",\"type\":\"uint40\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEAD\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUTE_ADDR1_BPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUTE_ADDR2_BPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN\",\"outputs\":[{\"internalType\":\"contractIStakeToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDT\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USER\",\"outputs\":[{\"internalType\":\"contractIUserSystem\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"queueIndex\",\"type\":\"uint256\"}],\"name\":\"adminCancelQueuedStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"queueIndexes\",\"type\":\"uint256[]\"}],\"name\":\"adminCancelQueuedStakes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"burnExpiredReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"queueIndex\",\"type\":\"uint256\"}],\"name\":\"cancelQueuedStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCollector\",\"type\":\"address\"}],\"name\":\"changeFreeUsdtCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"circuitBreakerRecovery\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"circuitBreakerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"circuitBreakerTime\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"circuitBreakerUnstakeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"configs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"duration\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"ttl\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dailyCircuitBreakerUnstaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dailyQueueCancelLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dailyStakeRateHigh\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dailyStakeRateLow\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dailyStakeRateMid\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dailyTotalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dailyTotalUnstaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dailyUnstakeRateHigh\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dailyUnstakeRateLow\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dailyUnstakeRateMid\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dividendSystem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ecoAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"queueIndex\",\"type\":\"uint256\"}],\"name\":\"executeQueuedStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freeUsdtBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freeUsdtCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCircuitBreakerUnstakeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDailyStakeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDailyUnstakeLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getOrderRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint40\",\"name\":\"stakeTime\",\"type\":\"uint40\"},{\"internalType\":\"uint160\",\"name\":\"amount\",\"type\":\"uint160\"},{\"internalType\":\"uint8\",\"name\":\"stakeIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"unstakeTime\",\"type\":\"uint40\"},{\"internalType\":\"uint160\",\"name\":\"reward\",\"type\":\"uint160\"},{\"internalType\":\"uint40\",\"name\":\"restakeTime\",\"type\":\"uint40\"},{\"internalType\":\"bool\",\"name\":\"cbNewOrder\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structStakingV9.Record\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"userSystem_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dividendSystem_\",\"type\":\"address\"}],\"name\":\"initSystems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxReferralDepth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxStakeAmountValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newHighTime\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxItems\",\"type\":\"uint256\"}],\"name\":\"processStakeQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueCursor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queuedUsdtBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rList\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"start\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"end\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rListLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingCircuitBreakerUnstakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingUnstakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"amount\",\"type\":\"uint160\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"stakeIndex\",\"type\":\"uint8\"}],\"name\":\"restake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardDirectBps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardEcoBps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardEcoBpsTwo\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardFeeTotalBps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardGameBps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardGlobalBps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardS7Bps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardTeamMaxBps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routeAddress1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routeAddress2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"b\",\"type\":\"uint40\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"}],\"name\":\"setCircuitBreakerParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"a\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"b\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"c\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"d\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"e\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"f\",\"type\":\"uint16\"}],\"name\":\"setDailyLimitRates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"setDailyQueueCancelLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newInitAccount\",\"type\":\"address\"}],\"name\":\"setInitAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setMarketingAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"setMaxReferralDepth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"setRewardRoutes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"setStageThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"setStakeQueueHandleCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shouldUpdateCircuitBreaker\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"needTrigger\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"needCountdown\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"needRecover\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"currentDropBps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stage1Threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stage2Threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"amount\",\"type\":\"uint160\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"stakeIndex\",\"type\":\"uint8\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"stakeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeQueue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"amount\",\"type\":\"uint160\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"stakeIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"queuedAt\",\"type\":\"uint40\"},{\"internalType\":\"bool\",\"name\":\"canceled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeQueueHandleCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint160\",\"name\":\"amount\",\"type\":\"uint160\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"stakeIndex\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"parent\",\"type\":\"address\"}],\"name\":\"stakeWithInviter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateCircuitBreaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userStakeRecord\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"stakeTime\",\"type\":\"uint40\"},{\"internalType\":\"uint160\",\"name\":\"amount\",\"type\":\"uint160\"},{\"internalType\":\"uint8\",\"name\":\"stakeIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"unstakeTime\",\"type\":\"uint40\"},{\"internalType\":\"uint160\",\"name\":\"reward\",\"type\":\"uint160\"},{\"internalType\":\"uint40\",\"name\":\"restakeTime\",\"type\":\"uint40\"},{\"internalType\":\"bool\",\"name\":\"cbNewOrder\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAllFreeUsdt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFreeUsdt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakeNewABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeNewMetaData.ABI instead.
var StakeNewABI = StakeNewMetaData.ABI

// StakeNew is an auto generated Go binding around an Ethereum contract.
type StakeNew struct {
	StakeNewCaller     // Read-only binding to the contract
	StakeNewTransactor // Write-only binding to the contract
	StakeNewFilterer   // Log filterer for contract events
}

// StakeNewCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeNewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeNewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeNewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeNewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeNewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeNewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeNewSession struct {
	Contract     *StakeNew         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeNewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeNewCallerSession struct {
	Contract *StakeNewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StakeNewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeNewTransactorSession struct {
	Contract     *StakeNewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StakeNewRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeNewRaw struct {
	Contract *StakeNew // Generic contract binding to access the raw methods on
}

// StakeNewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeNewCallerRaw struct {
	Contract *StakeNewCaller // Generic read-only contract binding to access the raw methods on
}

// StakeNewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeNewTransactorRaw struct {
	Contract *StakeNewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakeNew creates a new instance of StakeNew, bound to a specific deployed contract.
func NewStakeNew(address common.Address, backend bind.ContractBackend) (*StakeNew, error) {
	contract, err := bindStakeNew(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakeNew{StakeNewCaller: StakeNewCaller{contract: contract}, StakeNewTransactor: StakeNewTransactor{contract: contract}, StakeNewFilterer: StakeNewFilterer{contract: contract}}, nil
}

// NewStakeNewCaller creates a new read-only instance of StakeNew, bound to a specific deployed contract.
func NewStakeNewCaller(address common.Address, caller bind.ContractCaller) (*StakeNewCaller, error) {
	contract, err := bindStakeNew(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeNewCaller{contract: contract}, nil
}

// NewStakeNewTransactor creates a new write-only instance of StakeNew, bound to a specific deployed contract.
func NewStakeNewTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeNewTransactor, error) {
	contract, err := bindStakeNew(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeNewTransactor{contract: contract}, nil
}

// NewStakeNewFilterer creates a new log filterer instance of StakeNew, bound to a specific deployed contract.
func NewStakeNewFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeNewFilterer, error) {
	contract, err := bindStakeNew(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeNewFilterer{contract: contract}, nil
}

// bindStakeNew binds a generic wrapper to an already deployed contract.
func bindStakeNew(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeNewABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeNew *StakeNewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeNew.Contract.StakeNewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeNew *StakeNewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeNew.Contract.StakeNewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeNew *StakeNewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeNew.Contract.StakeNewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeNew *StakeNewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeNew.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeNew *StakeNewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeNew.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeNew *StakeNewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeNew.Contract.contract.Transact(opts, method, params...)
}

// DEAD is a free data retrieval call binding the contract method 0x03fd2a45.
//
// Solidity: function DEAD() view returns(address)
func (_StakeNew *StakeNewCaller) DEAD(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "DEAD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEAD is a free data retrieval call binding the contract method 0x03fd2a45.
//
// Solidity: function DEAD() view returns(address)
func (_StakeNew *StakeNewSession) DEAD() (common.Address, error) {
	return _StakeNew.Contract.DEAD(&_StakeNew.CallOpts)
}

// DEAD is a free data retrieval call binding the contract method 0x03fd2a45.
//
// Solidity: function DEAD() view returns(address)
func (_StakeNew *StakeNewCallerSession) DEAD() (common.Address, error) {
	return _StakeNew.Contract.DEAD(&_StakeNew.CallOpts)
}

// ROUTEADDR1BPS is a free data retrieval call binding the contract method 0xc1bebb53.
//
// Solidity: function ROUTE_ADDR1_BPS() view returns(uint16)
func (_StakeNew *StakeNewCaller) ROUTEADDR1BPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "ROUTE_ADDR1_BPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ROUTEADDR1BPS is a free data retrieval call binding the contract method 0xc1bebb53.
//
// Solidity: function ROUTE_ADDR1_BPS() view returns(uint16)
func (_StakeNew *StakeNewSession) ROUTEADDR1BPS() (uint16, error) {
	return _StakeNew.Contract.ROUTEADDR1BPS(&_StakeNew.CallOpts)
}

// ROUTEADDR1BPS is a free data retrieval call binding the contract method 0xc1bebb53.
//
// Solidity: function ROUTE_ADDR1_BPS() view returns(uint16)
func (_StakeNew *StakeNewCallerSession) ROUTEADDR1BPS() (uint16, error) {
	return _StakeNew.Contract.ROUTEADDR1BPS(&_StakeNew.CallOpts)
}

// ROUTEADDR2BPS is a free data retrieval call binding the contract method 0xbe487788.
//
// Solidity: function ROUTE_ADDR2_BPS() view returns(uint16)
func (_StakeNew *StakeNewCaller) ROUTEADDR2BPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "ROUTE_ADDR2_BPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ROUTEADDR2BPS is a free data retrieval call binding the contract method 0xbe487788.
//
// Solidity: function ROUTE_ADDR2_BPS() view returns(uint16)
func (_StakeNew *StakeNewSession) ROUTEADDR2BPS() (uint16, error) {
	return _StakeNew.Contract.ROUTEADDR2BPS(&_StakeNew.CallOpts)
}

// ROUTEADDR2BPS is a free data retrieval call binding the contract method 0xbe487788.
//
// Solidity: function ROUTE_ADDR2_BPS() view returns(uint16)
func (_StakeNew *StakeNewCallerSession) ROUTEADDR2BPS() (uint16, error) {
	return _StakeNew.Contract.ROUTEADDR2BPS(&_StakeNew.CallOpts)
}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(address)
func (_StakeNew *StakeNewCaller) TOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(address)
func (_StakeNew *StakeNewSession) TOKEN() (common.Address, error) {
	return _StakeNew.Contract.TOKEN(&_StakeNew.CallOpts)
}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(address)
func (_StakeNew *StakeNewCallerSession) TOKEN() (common.Address, error) {
	return _StakeNew.Contract.TOKEN(&_StakeNew.CallOpts)
}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_StakeNew *StakeNewCaller) USDT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "USDT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_StakeNew *StakeNewSession) USDT() (common.Address, error) {
	return _StakeNew.Contract.USDT(&_StakeNew.CallOpts)
}

// USDT is a free data retrieval call binding the contract method 0xc54e44eb.
//
// Solidity: function USDT() view returns(address)
func (_StakeNew *StakeNewCallerSession) USDT() (common.Address, error) {
	return _StakeNew.Contract.USDT(&_StakeNew.CallOpts)
}

// USER is a free data retrieval call binding the contract method 0x81e167cf.
//
// Solidity: function USER() view returns(address)
func (_StakeNew *StakeNewCaller) USER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "USER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USER is a free data retrieval call binding the contract method 0x81e167cf.
//
// Solidity: function USER() view returns(address)
func (_StakeNew *StakeNewSession) USER() (common.Address, error) {
	return _StakeNew.Contract.USER(&_StakeNew.CallOpts)
}

// USER is a free data retrieval call binding the contract method 0x81e167cf.
//
// Solidity: function USER() view returns(address)
func (_StakeNew *StakeNewCallerSession) USER() (common.Address, error) {
	return _StakeNew.Contract.USER(&_StakeNew.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 balance)
func (_StakeNew *StakeNewCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 balance)
func (_StakeNew *StakeNewSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StakeNew.Contract.BalanceOf(&_StakeNew.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 balance)
func (_StakeNew *StakeNewCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StakeNew.Contract.BalanceOf(&_StakeNew.CallOpts, account)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_StakeNew *StakeNewCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_StakeNew *StakeNewSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _StakeNew.Contract.Balances(&_StakeNew.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_StakeNew *StakeNewCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _StakeNew.Contract.Balances(&_StakeNew.CallOpts, arg0)
}

// CircuitBreakerRecovery is a free data retrieval call binding the contract method 0x86b54614.
//
// Solidity: function circuitBreakerRecovery() view returns(uint40)
func (_StakeNew *StakeNewCaller) CircuitBreakerRecovery(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "circuitBreakerRecovery")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CircuitBreakerRecovery is a free data retrieval call binding the contract method 0x86b54614.
//
// Solidity: function circuitBreakerRecovery() view returns(uint40)
func (_StakeNew *StakeNewSession) CircuitBreakerRecovery() (*big.Int, error) {
	return _StakeNew.Contract.CircuitBreakerRecovery(&_StakeNew.CallOpts)
}

// CircuitBreakerRecovery is a free data retrieval call binding the contract method 0x86b54614.
//
// Solidity: function circuitBreakerRecovery() view returns(uint40)
func (_StakeNew *StakeNewCallerSession) CircuitBreakerRecovery() (*big.Int, error) {
	return _StakeNew.Contract.CircuitBreakerRecovery(&_StakeNew.CallOpts)
}

// CircuitBreakerThreshold is a free data retrieval call binding the contract method 0x7d0b0a34.
//
// Solidity: function circuitBreakerThreshold() view returns(uint256)
func (_StakeNew *StakeNewCaller) CircuitBreakerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "circuitBreakerThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CircuitBreakerThreshold is a free data retrieval call binding the contract method 0x7d0b0a34.
//
// Solidity: function circuitBreakerThreshold() view returns(uint256)
func (_StakeNew *StakeNewSession) CircuitBreakerThreshold() (*big.Int, error) {
	return _StakeNew.Contract.CircuitBreakerThreshold(&_StakeNew.CallOpts)
}

// CircuitBreakerThreshold is a free data retrieval call binding the contract method 0x7d0b0a34.
//
// Solidity: function circuitBreakerThreshold() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) CircuitBreakerThreshold() (*big.Int, error) {
	return _StakeNew.Contract.CircuitBreakerThreshold(&_StakeNew.CallOpts)
}

// CircuitBreakerTime is a free data retrieval call binding the contract method 0x7bd1bb14.
//
// Solidity: function circuitBreakerTime() view returns(uint40)
func (_StakeNew *StakeNewCaller) CircuitBreakerTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "circuitBreakerTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CircuitBreakerTime is a free data retrieval call binding the contract method 0x7bd1bb14.
//
// Solidity: function circuitBreakerTime() view returns(uint40)
func (_StakeNew *StakeNewSession) CircuitBreakerTime() (*big.Int, error) {
	return _StakeNew.Contract.CircuitBreakerTime(&_StakeNew.CallOpts)
}

// CircuitBreakerTime is a free data retrieval call binding the contract method 0x7bd1bb14.
//
// Solidity: function circuitBreakerTime() view returns(uint40)
func (_StakeNew *StakeNewCallerSession) CircuitBreakerTime() (*big.Int, error) {
	return _StakeNew.Contract.CircuitBreakerTime(&_StakeNew.CallOpts)
}

// CircuitBreakerUnstakeRate is a free data retrieval call binding the contract method 0xa596a432.
//
// Solidity: function circuitBreakerUnstakeRate() view returns(uint256)
func (_StakeNew *StakeNewCaller) CircuitBreakerUnstakeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "circuitBreakerUnstakeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CircuitBreakerUnstakeRate is a free data retrieval call binding the contract method 0xa596a432.
//
// Solidity: function circuitBreakerUnstakeRate() view returns(uint256)
func (_StakeNew *StakeNewSession) CircuitBreakerUnstakeRate() (*big.Int, error) {
	return _StakeNew.Contract.CircuitBreakerUnstakeRate(&_StakeNew.CallOpts)
}

// CircuitBreakerUnstakeRate is a free data retrieval call binding the contract method 0xa596a432.
//
// Solidity: function circuitBreakerUnstakeRate() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) CircuitBreakerUnstakeRate() (*big.Int, error) {
	return _StakeNew.Contract.CircuitBreakerUnstakeRate(&_StakeNew.CallOpts)
}

// Configs is a free data retrieval call binding the contract method 0x0098fa22.
//
// Solidity: function configs(uint256 ) view returns(uint256 rate, uint40 duration, uint40 ttl)
func (_StakeNew *StakeNewCaller) Configs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Rate     *big.Int
	Duration *big.Int
	Ttl      *big.Int
}, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "configs", arg0)

	outstruct := new(struct {
		Rate     *big.Int
		Duration *big.Int
		Ttl      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Rate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ttl = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Configs is a free data retrieval call binding the contract method 0x0098fa22.
//
// Solidity: function configs(uint256 ) view returns(uint256 rate, uint40 duration, uint40 ttl)
func (_StakeNew *StakeNewSession) Configs(arg0 *big.Int) (struct {
	Rate     *big.Int
	Duration *big.Int
	Ttl      *big.Int
}, error) {
	return _StakeNew.Contract.Configs(&_StakeNew.CallOpts, arg0)
}

// Configs is a free data retrieval call binding the contract method 0x0098fa22.
//
// Solidity: function configs(uint256 ) view returns(uint256 rate, uint40 duration, uint40 ttl)
func (_StakeNew *StakeNewCallerSession) Configs(arg0 *big.Int) (struct {
	Rate     *big.Int
	Duration *big.Int
	Ttl      *big.Int
}, error) {
	return _StakeNew.Contract.Configs(&_StakeNew.CallOpts, arg0)
}

// DailyCircuitBreakerUnstaked is a free data retrieval call binding the contract method 0x908d5def.
//
// Solidity: function dailyCircuitBreakerUnstaked(uint256 ) view returns(uint256)
func (_StakeNew *StakeNewCaller) DailyCircuitBreakerUnstaked(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "dailyCircuitBreakerUnstaked", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyCircuitBreakerUnstaked is a free data retrieval call binding the contract method 0x908d5def.
//
// Solidity: function dailyCircuitBreakerUnstaked(uint256 ) view returns(uint256)
func (_StakeNew *StakeNewSession) DailyCircuitBreakerUnstaked(arg0 *big.Int) (*big.Int, error) {
	return _StakeNew.Contract.DailyCircuitBreakerUnstaked(&_StakeNew.CallOpts, arg0)
}

// DailyCircuitBreakerUnstaked is a free data retrieval call binding the contract method 0x908d5def.
//
// Solidity: function dailyCircuitBreakerUnstaked(uint256 ) view returns(uint256)
func (_StakeNew *StakeNewCallerSession) DailyCircuitBreakerUnstaked(arg0 *big.Int) (*big.Int, error) {
	return _StakeNew.Contract.DailyCircuitBreakerUnstaked(&_StakeNew.CallOpts, arg0)
}

// DailyQueueCancelLimit is a free data retrieval call binding the contract method 0xa6e6dca4.
//
// Solidity: function dailyQueueCancelLimit() view returns(uint256)
func (_StakeNew *StakeNewCaller) DailyQueueCancelLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "dailyQueueCancelLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyQueueCancelLimit is a free data retrieval call binding the contract method 0xa6e6dca4.
//
// Solidity: function dailyQueueCancelLimit() view returns(uint256)
func (_StakeNew *StakeNewSession) DailyQueueCancelLimit() (*big.Int, error) {
	return _StakeNew.Contract.DailyQueueCancelLimit(&_StakeNew.CallOpts)
}

// DailyQueueCancelLimit is a free data retrieval call binding the contract method 0xa6e6dca4.
//
// Solidity: function dailyQueueCancelLimit() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) DailyQueueCancelLimit() (*big.Int, error) {
	return _StakeNew.Contract.DailyQueueCancelLimit(&_StakeNew.CallOpts)
}

// DailyStakeRateHigh is a free data retrieval call binding the contract method 0xeadd46d8.
//
// Solidity: function dailyStakeRateHigh() view returns(uint16)
func (_StakeNew *StakeNewCaller) DailyStakeRateHigh(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "dailyStakeRateHigh")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DailyStakeRateHigh is a free data retrieval call binding the contract method 0xeadd46d8.
//
// Solidity: function dailyStakeRateHigh() view returns(uint16)
func (_StakeNew *StakeNewSession) DailyStakeRateHigh() (uint16, error) {
	return _StakeNew.Contract.DailyStakeRateHigh(&_StakeNew.CallOpts)
}

// DailyStakeRateHigh is a free data retrieval call binding the contract method 0xeadd46d8.
//
// Solidity: function dailyStakeRateHigh() view returns(uint16)
func (_StakeNew *StakeNewCallerSession) DailyStakeRateHigh() (uint16, error) {
	return _StakeNew.Contract.DailyStakeRateHigh(&_StakeNew.CallOpts)
}

// DailyStakeRateLow is a free data retrieval call binding the contract method 0x2c4fc9b6.
//
// Solidity: function dailyStakeRateLow() view returns(uint16)
func (_StakeNew *StakeNewCaller) DailyStakeRateLow(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "dailyStakeRateLow")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DailyStakeRateLow is a free data retrieval call binding the contract method 0x2c4fc9b6.
//
// Solidity: function dailyStakeRateLow() view returns(uint16)
func (_StakeNew *StakeNewSession) DailyStakeRateLow() (uint16, error) {
	return _StakeNew.Contract.DailyStakeRateLow(&_StakeNew.CallOpts)
}

// DailyStakeRateLow is a free data retrieval call binding the contract method 0x2c4fc9b6.
//
// Solidity: function dailyStakeRateLow() view returns(uint16)
func (_StakeNew *StakeNewCallerSession) DailyStakeRateLow() (uint16, error) {
	return _StakeNew.Contract.DailyStakeRateLow(&_StakeNew.CallOpts)
}

// DailyStakeRateMid is a free data retrieval call binding the contract method 0x0ccc75eb.
//
// Solidity: function dailyStakeRateMid() view returns(uint16)
func (_StakeNew *StakeNewCaller) DailyStakeRateMid(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "dailyStakeRateMid")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DailyStakeRateMid is a free data retrieval call binding the contract method 0x0ccc75eb.
//
// Solidity: function dailyStakeRateMid() view returns(uint16)
func (_StakeNew *StakeNewSession) DailyStakeRateMid() (uint16, error) {
	return _StakeNew.Contract.DailyStakeRateMid(&_StakeNew.CallOpts)
}

// DailyStakeRateMid is a free data retrieval call binding the contract method 0x0ccc75eb.
//
// Solidity: function dailyStakeRateMid() view returns(uint16)
func (_StakeNew *StakeNewCallerSession) DailyStakeRateMid() (uint16, error) {
	return _StakeNew.Contract.DailyStakeRateMid(&_StakeNew.CallOpts)
}

// DailyTotalStaked is a free data retrieval call binding the contract method 0x77e9e339.
//
// Solidity: function dailyTotalStaked(uint256 ) view returns(uint256)
func (_StakeNew *StakeNewCaller) DailyTotalStaked(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "dailyTotalStaked", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyTotalStaked is a free data retrieval call binding the contract method 0x77e9e339.
//
// Solidity: function dailyTotalStaked(uint256 ) view returns(uint256)
func (_StakeNew *StakeNewSession) DailyTotalStaked(arg0 *big.Int) (*big.Int, error) {
	return _StakeNew.Contract.DailyTotalStaked(&_StakeNew.CallOpts, arg0)
}

// DailyTotalStaked is a free data retrieval call binding the contract method 0x77e9e339.
//
// Solidity: function dailyTotalStaked(uint256 ) view returns(uint256)
func (_StakeNew *StakeNewCallerSession) DailyTotalStaked(arg0 *big.Int) (*big.Int, error) {
	return _StakeNew.Contract.DailyTotalStaked(&_StakeNew.CallOpts, arg0)
}

// DailyTotalUnstaked is a free data retrieval call binding the contract method 0x8d7bdceb.
//
// Solidity: function dailyTotalUnstaked(uint256 ) view returns(uint256)
func (_StakeNew *StakeNewCaller) DailyTotalUnstaked(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "dailyTotalUnstaked", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyTotalUnstaked is a free data retrieval call binding the contract method 0x8d7bdceb.
//
// Solidity: function dailyTotalUnstaked(uint256 ) view returns(uint256)
func (_StakeNew *StakeNewSession) DailyTotalUnstaked(arg0 *big.Int) (*big.Int, error) {
	return _StakeNew.Contract.DailyTotalUnstaked(&_StakeNew.CallOpts, arg0)
}

// DailyTotalUnstaked is a free data retrieval call binding the contract method 0x8d7bdceb.
//
// Solidity: function dailyTotalUnstaked(uint256 ) view returns(uint256)
func (_StakeNew *StakeNewCallerSession) DailyTotalUnstaked(arg0 *big.Int) (*big.Int, error) {
	return _StakeNew.Contract.DailyTotalUnstaked(&_StakeNew.CallOpts, arg0)
}

// DailyUnstakeRateHigh is a free data retrieval call binding the contract method 0xfa154e2c.
//
// Solidity: function dailyUnstakeRateHigh() view returns(uint16)
func (_StakeNew *StakeNewCaller) DailyUnstakeRateHigh(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "dailyUnstakeRateHigh")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DailyUnstakeRateHigh is a free data retrieval call binding the contract method 0xfa154e2c.
//
// Solidity: function dailyUnstakeRateHigh() view returns(uint16)
func (_StakeNew *StakeNewSession) DailyUnstakeRateHigh() (uint16, error) {
	return _StakeNew.Contract.DailyUnstakeRateHigh(&_StakeNew.CallOpts)
}

// DailyUnstakeRateHigh is a free data retrieval call binding the contract method 0xfa154e2c.
//
// Solidity: function dailyUnstakeRateHigh() view returns(uint16)
func (_StakeNew *StakeNewCallerSession) DailyUnstakeRateHigh() (uint16, error) {
	return _StakeNew.Contract.DailyUnstakeRateHigh(&_StakeNew.CallOpts)
}

// DailyUnstakeRateLow is a free data retrieval call binding the contract method 0xae4b06c7.
//
// Solidity: function dailyUnstakeRateLow() view returns(uint16)
func (_StakeNew *StakeNewCaller) DailyUnstakeRateLow(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "dailyUnstakeRateLow")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DailyUnstakeRateLow is a free data retrieval call binding the contract method 0xae4b06c7.
//
// Solidity: function dailyUnstakeRateLow() view returns(uint16)
func (_StakeNew *StakeNewSession) DailyUnstakeRateLow() (uint16, error) {
	return _StakeNew.Contract.DailyUnstakeRateLow(&_StakeNew.CallOpts)
}

// DailyUnstakeRateLow is a free data retrieval call binding the contract method 0xae4b06c7.
//
// Solidity: function dailyUnstakeRateLow() view returns(uint16)
func (_StakeNew *StakeNewCallerSession) DailyUnstakeRateLow() (uint16, error) {
	return _StakeNew.Contract.DailyUnstakeRateLow(&_StakeNew.CallOpts)
}

// DailyUnstakeRateMid is a free data retrieval call binding the contract method 0x5aed1b34.
//
// Solidity: function dailyUnstakeRateMid() view returns(uint16)
func (_StakeNew *StakeNewCaller) DailyUnstakeRateMid(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "dailyUnstakeRateMid")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DailyUnstakeRateMid is a free data retrieval call binding the contract method 0x5aed1b34.
//
// Solidity: function dailyUnstakeRateMid() view returns(uint16)
func (_StakeNew *StakeNewSession) DailyUnstakeRateMid() (uint16, error) {
	return _StakeNew.Contract.DailyUnstakeRateMid(&_StakeNew.CallOpts)
}

// DailyUnstakeRateMid is a free data retrieval call binding the contract method 0x5aed1b34.
//
// Solidity: function dailyUnstakeRateMid() view returns(uint16)
func (_StakeNew *StakeNewCallerSession) DailyUnstakeRateMid() (uint16, error) {
	return _StakeNew.Contract.DailyUnstakeRateMid(&_StakeNew.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StakeNew *StakeNewCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StakeNew *StakeNewSession) Decimals() (uint8, error) {
	return _StakeNew.Contract.Decimals(&_StakeNew.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StakeNew *StakeNewCallerSession) Decimals() (uint8, error) {
	return _StakeNew.Contract.Decimals(&_StakeNew.CallOpts)
}

// DividendSystem is a free data retrieval call binding the contract method 0xc2511edb.
//
// Solidity: function dividendSystem() view returns(address)
func (_StakeNew *StakeNewCaller) DividendSystem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "dividendSystem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DividendSystem is a free data retrieval call binding the contract method 0xc2511edb.
//
// Solidity: function dividendSystem() view returns(address)
func (_StakeNew *StakeNewSession) DividendSystem() (common.Address, error) {
	return _StakeNew.Contract.DividendSystem(&_StakeNew.CallOpts)
}

// DividendSystem is a free data retrieval call binding the contract method 0xc2511edb.
//
// Solidity: function dividendSystem() view returns(address)
func (_StakeNew *StakeNewCallerSession) DividendSystem() (common.Address, error) {
	return _StakeNew.Contract.DividendSystem(&_StakeNew.CallOpts)
}

// EcoAddress is a free data retrieval call binding the contract method 0x68961ecb.
//
// Solidity: function ecoAddress() view returns(address)
func (_StakeNew *StakeNewCaller) EcoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "ecoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EcoAddress is a free data retrieval call binding the contract method 0x68961ecb.
//
// Solidity: function ecoAddress() view returns(address)
func (_StakeNew *StakeNewSession) EcoAddress() (common.Address, error) {
	return _StakeNew.Contract.EcoAddress(&_StakeNew.CallOpts)
}

// EcoAddress is a free data retrieval call binding the contract method 0x68961ecb.
//
// Solidity: function ecoAddress() view returns(address)
func (_StakeNew *StakeNewCallerSession) EcoAddress() (common.Address, error) {
	return _StakeNew.Contract.EcoAddress(&_StakeNew.CallOpts)
}

// FreeUsdtBalance is a free data retrieval call binding the contract method 0xac1b3379.
//
// Solidity: function freeUsdtBalance() view returns(uint256)
func (_StakeNew *StakeNewCaller) FreeUsdtBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "freeUsdtBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreeUsdtBalance is a free data retrieval call binding the contract method 0xac1b3379.
//
// Solidity: function freeUsdtBalance() view returns(uint256)
func (_StakeNew *StakeNewSession) FreeUsdtBalance() (*big.Int, error) {
	return _StakeNew.Contract.FreeUsdtBalance(&_StakeNew.CallOpts)
}

// FreeUsdtBalance is a free data retrieval call binding the contract method 0xac1b3379.
//
// Solidity: function freeUsdtBalance() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) FreeUsdtBalance() (*big.Int, error) {
	return _StakeNew.Contract.FreeUsdtBalance(&_StakeNew.CallOpts)
}

// FreeUsdtCollector is a free data retrieval call binding the contract method 0x583c6280.
//
// Solidity: function freeUsdtCollector() view returns(address)
func (_StakeNew *StakeNewCaller) FreeUsdtCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "freeUsdtCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FreeUsdtCollector is a free data retrieval call binding the contract method 0x583c6280.
//
// Solidity: function freeUsdtCollector() view returns(address)
func (_StakeNew *StakeNewSession) FreeUsdtCollector() (common.Address, error) {
	return _StakeNew.Contract.FreeUsdtCollector(&_StakeNew.CallOpts)
}

// FreeUsdtCollector is a free data retrieval call binding the contract method 0x583c6280.
//
// Solidity: function freeUsdtCollector() view returns(address)
func (_StakeNew *StakeNewCallerSession) FreeUsdtCollector() (common.Address, error) {
	return _StakeNew.Contract.FreeUsdtCollector(&_StakeNew.CallOpts)
}

// GameAddress is a free data retrieval call binding the contract method 0xa168d873.
//
// Solidity: function gameAddress() view returns(address)
func (_StakeNew *StakeNewCaller) GameAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "gameAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GameAddress is a free data retrieval call binding the contract method 0xa168d873.
//
// Solidity: function gameAddress() view returns(address)
func (_StakeNew *StakeNewSession) GameAddress() (common.Address, error) {
	return _StakeNew.Contract.GameAddress(&_StakeNew.CallOpts)
}

// GameAddress is a free data retrieval call binding the contract method 0xa168d873.
//
// Solidity: function gameAddress() view returns(address)
func (_StakeNew *StakeNewCallerSession) GameAddress() (common.Address, error) {
	return _StakeNew.Contract.GameAddress(&_StakeNew.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_StakeNew *StakeNewCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_StakeNew *StakeNewSession) Gateway() (common.Address, error) {
	return _StakeNew.Contract.Gateway(&_StakeNew.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_StakeNew *StakeNewCallerSession) Gateway() (common.Address, error) {
	return _StakeNew.Contract.Gateway(&_StakeNew.CallOpts)
}

// GetCircuitBreakerUnstakeLimit is a free data retrieval call binding the contract method 0x4e5f8ed9.
//
// Solidity: function getCircuitBreakerUnstakeLimit() view returns(uint256)
func (_StakeNew *StakeNewCaller) GetCircuitBreakerUnstakeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "getCircuitBreakerUnstakeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCircuitBreakerUnstakeLimit is a free data retrieval call binding the contract method 0x4e5f8ed9.
//
// Solidity: function getCircuitBreakerUnstakeLimit() view returns(uint256)
func (_StakeNew *StakeNewSession) GetCircuitBreakerUnstakeLimit() (*big.Int, error) {
	return _StakeNew.Contract.GetCircuitBreakerUnstakeLimit(&_StakeNew.CallOpts)
}

// GetCircuitBreakerUnstakeLimit is a free data retrieval call binding the contract method 0x4e5f8ed9.
//
// Solidity: function getCircuitBreakerUnstakeLimit() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) GetCircuitBreakerUnstakeLimit() (*big.Int, error) {
	return _StakeNew.Contract.GetCircuitBreakerUnstakeLimit(&_StakeNew.CallOpts)
}

// GetDailyStakeLimit is a free data retrieval call binding the contract method 0xcc34d99f.
//
// Solidity: function getDailyStakeLimit() view returns(uint256)
func (_StakeNew *StakeNewCaller) GetDailyStakeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "getDailyStakeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDailyStakeLimit is a free data retrieval call binding the contract method 0xcc34d99f.
//
// Solidity: function getDailyStakeLimit() view returns(uint256)
func (_StakeNew *StakeNewSession) GetDailyStakeLimit() (*big.Int, error) {
	return _StakeNew.Contract.GetDailyStakeLimit(&_StakeNew.CallOpts)
}

// GetDailyStakeLimit is a free data retrieval call binding the contract method 0xcc34d99f.
//
// Solidity: function getDailyStakeLimit() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) GetDailyStakeLimit() (*big.Int, error) {
	return _StakeNew.Contract.GetDailyStakeLimit(&_StakeNew.CallOpts)
}

// GetDailyUnstakeLimit is a free data retrieval call binding the contract method 0xfa40a1e5.
//
// Solidity: function getDailyUnstakeLimit() view returns(uint256)
func (_StakeNew *StakeNewCaller) GetDailyUnstakeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "getDailyUnstakeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDailyUnstakeLimit is a free data retrieval call binding the contract method 0xfa40a1e5.
//
// Solidity: function getDailyUnstakeLimit() view returns(uint256)
func (_StakeNew *StakeNewSession) GetDailyUnstakeLimit() (*big.Int, error) {
	return _StakeNew.Contract.GetDailyUnstakeLimit(&_StakeNew.CallOpts)
}

// GetDailyUnstakeLimit is a free data retrieval call binding the contract method 0xfa40a1e5.
//
// Solidity: function getDailyUnstakeLimit() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) GetDailyUnstakeLimit() (*big.Int, error) {
	return _StakeNew.Contract.GetDailyUnstakeLimit(&_StakeNew.CallOpts)
}

// GetOrderRecord is a free data retrieval call binding the contract method 0x20629568.
//
// Solidity: function getOrderRecord(address user, uint256 index) view returns((uint40,uint160,uint8,uint40,uint160,uint40,bool,uint8))
func (_StakeNew *StakeNewCaller) GetOrderRecord(opts *bind.CallOpts, user common.Address, index *big.Int) (StakingV9Record, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "getOrderRecord", user, index)

	if err != nil {
		return *new(StakingV9Record), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingV9Record)).(*StakingV9Record)

	return out0, err

}

// GetOrderRecord is a free data retrieval call binding the contract method 0x20629568.
//
// Solidity: function getOrderRecord(address user, uint256 index) view returns((uint40,uint160,uint8,uint40,uint160,uint40,bool,uint8))
func (_StakeNew *StakeNewSession) GetOrderRecord(user common.Address, index *big.Int) (StakingV9Record, error) {
	return _StakeNew.Contract.GetOrderRecord(&_StakeNew.CallOpts, user, index)
}

// GetOrderRecord is a free data retrieval call binding the contract method 0x20629568.
//
// Solidity: function getOrderRecord(address user, uint256 index) view returns((uint40,uint160,uint8,uint40,uint160,uint40,bool,uint8))
func (_StakeNew *StakeNewCallerSession) GetOrderRecord(user common.Address, index *big.Int) (StakingV9Record, error) {
	return _StakeNew.Contract.GetOrderRecord(&_StakeNew.CallOpts, user, index)
}

// InitAccount is a free data retrieval call binding the contract method 0x4ec1ce21.
//
// Solidity: function initAccount() view returns(address)
func (_StakeNew *StakeNewCaller) InitAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "initAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InitAccount is a free data retrieval call binding the contract method 0x4ec1ce21.
//
// Solidity: function initAccount() view returns(address)
func (_StakeNew *StakeNewSession) InitAccount() (common.Address, error) {
	return _StakeNew.Contract.InitAccount(&_StakeNew.CallOpts)
}

// InitAccount is a free data retrieval call binding the contract method 0x4ec1ce21.
//
// Solidity: function initAccount() view returns(address)
func (_StakeNew *StakeNewCallerSession) InitAccount() (common.Address, error) {
	return _StakeNew.Contract.InitAccount(&_StakeNew.CallOpts)
}

// MarketingAddress is a free data retrieval call binding the contract method 0xa5ece941.
//
// Solidity: function marketingAddress() view returns(address)
func (_StakeNew *StakeNewCaller) MarketingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "marketingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketingAddress is a free data retrieval call binding the contract method 0xa5ece941.
//
// Solidity: function marketingAddress() view returns(address)
func (_StakeNew *StakeNewSession) MarketingAddress() (common.Address, error) {
	return _StakeNew.Contract.MarketingAddress(&_StakeNew.CallOpts)
}

// MarketingAddress is a free data retrieval call binding the contract method 0xa5ece941.
//
// Solidity: function marketingAddress() view returns(address)
func (_StakeNew *StakeNewCallerSession) MarketingAddress() (common.Address, error) {
	return _StakeNew.Contract.MarketingAddress(&_StakeNew.CallOpts)
}

// MaxReferralDepth is a free data retrieval call binding the contract method 0x04189250.
//
// Solidity: function maxReferralDepth() view returns(uint256)
func (_StakeNew *StakeNewCaller) MaxReferralDepth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "maxReferralDepth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxReferralDepth is a free data retrieval call binding the contract method 0x04189250.
//
// Solidity: function maxReferralDepth() view returns(uint256)
func (_StakeNew *StakeNewSession) MaxReferralDepth() (*big.Int, error) {
	return _StakeNew.Contract.MaxReferralDepth(&_StakeNew.CallOpts)
}

// MaxReferralDepth is a free data retrieval call binding the contract method 0x04189250.
//
// Solidity: function maxReferralDepth() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) MaxReferralDepth() (*big.Int, error) {
	return _StakeNew.Contract.MaxReferralDepth(&_StakeNew.CallOpts)
}

// MaxStakeAmount is a free data retrieval call binding the contract method 0x5d80ca32.
//
// Solidity: function maxStakeAmount() view returns(uint256)
func (_StakeNew *StakeNewCaller) MaxStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "maxStakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxStakeAmount is a free data retrieval call binding the contract method 0x5d80ca32.
//
// Solidity: function maxStakeAmount() view returns(uint256)
func (_StakeNew *StakeNewSession) MaxStakeAmount() (*big.Int, error) {
	return _StakeNew.Contract.MaxStakeAmount(&_StakeNew.CallOpts)
}

// MaxStakeAmount is a free data retrieval call binding the contract method 0x5d80ca32.
//
// Solidity: function maxStakeAmount() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) MaxStakeAmount() (*big.Int, error) {
	return _StakeNew.Contract.MaxStakeAmount(&_StakeNew.CallOpts)
}

// MaxStakeAmountValue is a free data retrieval call binding the contract method 0x63f0f364.
//
// Solidity: function maxStakeAmountValue() view returns(uint256)
func (_StakeNew *StakeNewCaller) MaxStakeAmountValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "maxStakeAmountValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxStakeAmountValue is a free data retrieval call binding the contract method 0x63f0f364.
//
// Solidity: function maxStakeAmountValue() view returns(uint256)
func (_StakeNew *StakeNewSession) MaxStakeAmountValue() (*big.Int, error) {
	return _StakeNew.Contract.MaxStakeAmountValue(&_StakeNew.CallOpts)
}

// MaxStakeAmountValue is a free data retrieval call binding the contract method 0x63f0f364.
//
// Solidity: function maxStakeAmountValue() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) MaxStakeAmountValue() (*big.Int, error) {
	return _StakeNew.Contract.MaxStakeAmountValue(&_StakeNew.CallOpts)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (_StakeNew *StakeNewCaller) MinStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "minStakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (_StakeNew *StakeNewSession) MinStakeAmount() (*big.Int, error) {
	return _StakeNew.Contract.MinStakeAmount(&_StakeNew.CallOpts)
}

// MinStakeAmount is a free data retrieval call binding the contract method 0xf1887684.
//
// Solidity: function minStakeAmount() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) MinStakeAmount() (*big.Int, error) {
	return _StakeNew.Contract.MinStakeAmount(&_StakeNew.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakeNew *StakeNewCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakeNew *StakeNewSession) Name() (string, error) {
	return _StakeNew.Contract.Name(&_StakeNew.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakeNew *StakeNewCallerSession) Name() (string, error) {
	return _StakeNew.Contract.Name(&_StakeNew.CallOpts)
}

// NewHighTime is a free data retrieval call binding the contract method 0xdb872c3b.
//
// Solidity: function newHighTime() view returns(uint40)
func (_StakeNew *StakeNewCaller) NewHighTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "newHighTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewHighTime is a free data retrieval call binding the contract method 0xdb872c3b.
//
// Solidity: function newHighTime() view returns(uint40)
func (_StakeNew *StakeNewSession) NewHighTime() (*big.Int, error) {
	return _StakeNew.Contract.NewHighTime(&_StakeNew.CallOpts)
}

// NewHighTime is a free data retrieval call binding the contract method 0xdb872c3b.
//
// Solidity: function newHighTime() view returns(uint40)
func (_StakeNew *StakeNewCallerSession) NewHighTime() (*big.Int, error) {
	return _StakeNew.Contract.NewHighTime(&_StakeNew.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakeNew *StakeNewCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakeNew *StakeNewSession) Owner() (common.Address, error) {
	return _StakeNew.Contract.Owner(&_StakeNew.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakeNew *StakeNewCallerSession) Owner() (common.Address, error) {
	return _StakeNew.Contract.Owner(&_StakeNew.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakeNew *StakeNewCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakeNew *StakeNewSession) Paused() (bool, error) {
	return _StakeNew.Contract.Paused(&_StakeNew.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakeNew *StakeNewCallerSession) Paused() (bool, error) {
	return _StakeNew.Contract.Paused(&_StakeNew.CallOpts)
}

// QueueCursor is a free data retrieval call binding the contract method 0xf057a15e.
//
// Solidity: function queueCursor() view returns(uint256)
func (_StakeNew *StakeNewCaller) QueueCursor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "queueCursor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueueCursor is a free data retrieval call binding the contract method 0xf057a15e.
//
// Solidity: function queueCursor() view returns(uint256)
func (_StakeNew *StakeNewSession) QueueCursor() (*big.Int, error) {
	return _StakeNew.Contract.QueueCursor(&_StakeNew.CallOpts)
}

// QueueCursor is a free data retrieval call binding the contract method 0xf057a15e.
//
// Solidity: function queueCursor() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) QueueCursor() (*big.Int, error) {
	return _StakeNew.Contract.QueueCursor(&_StakeNew.CallOpts)
}

// QueueLength is a free data retrieval call binding the contract method 0xab91c7b0.
//
// Solidity: function queueLength() view returns(uint256)
func (_StakeNew *StakeNewCaller) QueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "queueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueueLength is a free data retrieval call binding the contract method 0xab91c7b0.
//
// Solidity: function queueLength() view returns(uint256)
func (_StakeNew *StakeNewSession) QueueLength() (*big.Int, error) {
	return _StakeNew.Contract.QueueLength(&_StakeNew.CallOpts)
}

// QueueLength is a free data retrieval call binding the contract method 0xab91c7b0.
//
// Solidity: function queueLength() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) QueueLength() (*big.Int, error) {
	return _StakeNew.Contract.QueueLength(&_StakeNew.CallOpts)
}

// QueuedUsdtBalance is a free data retrieval call binding the contract method 0xe2b4141e.
//
// Solidity: function queuedUsdtBalance() view returns(uint256)
func (_StakeNew *StakeNewCaller) QueuedUsdtBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "queuedUsdtBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueuedUsdtBalance is a free data retrieval call binding the contract method 0xe2b4141e.
//
// Solidity: function queuedUsdtBalance() view returns(uint256)
func (_StakeNew *StakeNewSession) QueuedUsdtBalance() (*big.Int, error) {
	return _StakeNew.Contract.QueuedUsdtBalance(&_StakeNew.CallOpts)
}

// QueuedUsdtBalance is a free data retrieval call binding the contract method 0xe2b4141e.
//
// Solidity: function queuedUsdtBalance() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) QueuedUsdtBalance() (*big.Int, error) {
	return _StakeNew.Contract.QueuedUsdtBalance(&_StakeNew.CallOpts)
}

// RList is a free data retrieval call binding the contract method 0x54d3cd87.
//
// Solidity: function rList(uint256 ) view returns(uint40 start, uint40 end)
func (_StakeNew *StakeNewCaller) RList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Start *big.Int
	End   *big.Int
}, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "rList", arg0)

	outstruct := new(struct {
		Start *big.Int
		End   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Start = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RList is a free data retrieval call binding the contract method 0x54d3cd87.
//
// Solidity: function rList(uint256 ) view returns(uint40 start, uint40 end)
func (_StakeNew *StakeNewSession) RList(arg0 *big.Int) (struct {
	Start *big.Int
	End   *big.Int
}, error) {
	return _StakeNew.Contract.RList(&_StakeNew.CallOpts, arg0)
}

// RList is a free data retrieval call binding the contract method 0x54d3cd87.
//
// Solidity: function rList(uint256 ) view returns(uint40 start, uint40 end)
func (_StakeNew *StakeNewCallerSession) RList(arg0 *big.Int) (struct {
	Start *big.Int
	End   *big.Int
}, error) {
	return _StakeNew.Contract.RList(&_StakeNew.CallOpts, arg0)
}

// RListLength is a free data retrieval call binding the contract method 0x217959fc.
//
// Solidity: function rListLength() view returns(uint256)
func (_StakeNew *StakeNewCaller) RListLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "rListLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RListLength is a free data retrieval call binding the contract method 0x217959fc.
//
// Solidity: function rListLength() view returns(uint256)
func (_StakeNew *StakeNewSession) RListLength() (*big.Int, error) {
	return _StakeNew.Contract.RListLength(&_StakeNew.CallOpts)
}

// RListLength is a free data retrieval call binding the contract method 0x217959fc.
//
// Solidity: function rListLength() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) RListLength() (*big.Int, error) {
	return _StakeNew.Contract.RListLength(&_StakeNew.CallOpts)
}

// RemainingCircuitBreakerUnstakeAmount is a free data retrieval call binding the contract method 0x702c370a.
//
// Solidity: function remainingCircuitBreakerUnstakeAmount() view returns(uint256)
func (_StakeNew *StakeNewCaller) RemainingCircuitBreakerUnstakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "remainingCircuitBreakerUnstakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingCircuitBreakerUnstakeAmount is a free data retrieval call binding the contract method 0x702c370a.
//
// Solidity: function remainingCircuitBreakerUnstakeAmount() view returns(uint256)
func (_StakeNew *StakeNewSession) RemainingCircuitBreakerUnstakeAmount() (*big.Int, error) {
	return _StakeNew.Contract.RemainingCircuitBreakerUnstakeAmount(&_StakeNew.CallOpts)
}

// RemainingCircuitBreakerUnstakeAmount is a free data retrieval call binding the contract method 0x702c370a.
//
// Solidity: function remainingCircuitBreakerUnstakeAmount() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) RemainingCircuitBreakerUnstakeAmount() (*big.Int, error) {
	return _StakeNew.Contract.RemainingCircuitBreakerUnstakeAmount(&_StakeNew.CallOpts)
}

// RemainingUnstakeAmount is a free data retrieval call binding the contract method 0xe6ee3a71.
//
// Solidity: function remainingUnstakeAmount() view returns(uint256)
func (_StakeNew *StakeNewCaller) RemainingUnstakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "remainingUnstakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingUnstakeAmount is a free data retrieval call binding the contract method 0xe6ee3a71.
//
// Solidity: function remainingUnstakeAmount() view returns(uint256)
func (_StakeNew *StakeNewSession) RemainingUnstakeAmount() (*big.Int, error) {
	return _StakeNew.Contract.RemainingUnstakeAmount(&_StakeNew.CallOpts)
}

// RemainingUnstakeAmount is a free data retrieval call binding the contract method 0xe6ee3a71.
//
// Solidity: function remainingUnstakeAmount() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) RemainingUnstakeAmount() (*big.Int, error) {
	return _StakeNew.Contract.RemainingUnstakeAmount(&_StakeNew.CallOpts)
}

// RewardDirectBps is a free data retrieval call binding the contract method 0x22d5f3f6.
//
// Solidity: function rewardDirectBps() view returns(uint16)
func (_StakeNew *StakeNewCaller) RewardDirectBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "rewardDirectBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RewardDirectBps is a free data retrieval call binding the contract method 0x22d5f3f6.
//
// Solidity: function rewardDirectBps() view returns(uint16)
func (_StakeNew *StakeNewSession) RewardDirectBps() (uint16, error) {
	return _StakeNew.Contract.RewardDirectBps(&_StakeNew.CallOpts)
}

// RewardDirectBps is a free data retrieval call binding the contract method 0x22d5f3f6.
//
// Solidity: function rewardDirectBps() view returns(uint16)
func (_StakeNew *StakeNewCallerSession) RewardDirectBps() (uint16, error) {
	return _StakeNew.Contract.RewardDirectBps(&_StakeNew.CallOpts)
}

// RewardEcoBps is a free data retrieval call binding the contract method 0xc3afa111.
//
// Solidity: function rewardEcoBps() view returns(uint16)
func (_StakeNew *StakeNewCaller) RewardEcoBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "rewardEcoBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RewardEcoBps is a free data retrieval call binding the contract method 0xc3afa111.
//
// Solidity: function rewardEcoBps() view returns(uint16)
func (_StakeNew *StakeNewSession) RewardEcoBps() (uint16, error) {
	return _StakeNew.Contract.RewardEcoBps(&_StakeNew.CallOpts)
}

// RewardEcoBps is a free data retrieval call binding the contract method 0xc3afa111.
//
// Solidity: function rewardEcoBps() view returns(uint16)
func (_StakeNew *StakeNewCallerSession) RewardEcoBps() (uint16, error) {
	return _StakeNew.Contract.RewardEcoBps(&_StakeNew.CallOpts)
}

// RewardEcoBpsTwo is a free data retrieval call binding the contract method 0x49e86e3f.
//
// Solidity: function rewardEcoBpsTwo() view returns(uint16)
func (_StakeNew *StakeNewCaller) RewardEcoBpsTwo(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "rewardEcoBpsTwo")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RewardEcoBpsTwo is a free data retrieval call binding the contract method 0x49e86e3f.
//
// Solidity: function rewardEcoBpsTwo() view returns(uint16)
func (_StakeNew *StakeNewSession) RewardEcoBpsTwo() (uint16, error) {
	return _StakeNew.Contract.RewardEcoBpsTwo(&_StakeNew.CallOpts)
}

// RewardEcoBpsTwo is a free data retrieval call binding the contract method 0x49e86e3f.
//
// Solidity: function rewardEcoBpsTwo() view returns(uint16)
func (_StakeNew *StakeNewCallerSession) RewardEcoBpsTwo() (uint16, error) {
	return _StakeNew.Contract.RewardEcoBpsTwo(&_StakeNew.CallOpts)
}

// RewardFeeTotalBps is a free data retrieval call binding the contract method 0xf37db2ed.
//
// Solidity: function rewardFeeTotalBps() view returns(uint16)
func (_StakeNew *StakeNewCaller) RewardFeeTotalBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "rewardFeeTotalBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RewardFeeTotalBps is a free data retrieval call binding the contract method 0xf37db2ed.
//
// Solidity: function rewardFeeTotalBps() view returns(uint16)
func (_StakeNew *StakeNewSession) RewardFeeTotalBps() (uint16, error) {
	return _StakeNew.Contract.RewardFeeTotalBps(&_StakeNew.CallOpts)
}

// RewardFeeTotalBps is a free data retrieval call binding the contract method 0xf37db2ed.
//
// Solidity: function rewardFeeTotalBps() view returns(uint16)
func (_StakeNew *StakeNewCallerSession) RewardFeeTotalBps() (uint16, error) {
	return _StakeNew.Contract.RewardFeeTotalBps(&_StakeNew.CallOpts)
}

// RewardGameBps is a free data retrieval call binding the contract method 0xeac7da0e.
//
// Solidity: function rewardGameBps() view returns(uint16)
func (_StakeNew *StakeNewCaller) RewardGameBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "rewardGameBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RewardGameBps is a free data retrieval call binding the contract method 0xeac7da0e.
//
// Solidity: function rewardGameBps() view returns(uint16)
func (_StakeNew *StakeNewSession) RewardGameBps() (uint16, error) {
	return _StakeNew.Contract.RewardGameBps(&_StakeNew.CallOpts)
}

// RewardGameBps is a free data retrieval call binding the contract method 0xeac7da0e.
//
// Solidity: function rewardGameBps() view returns(uint16)
func (_StakeNew *StakeNewCallerSession) RewardGameBps() (uint16, error) {
	return _StakeNew.Contract.RewardGameBps(&_StakeNew.CallOpts)
}

// RewardGlobalBps is a free data retrieval call binding the contract method 0xce765062.
//
// Solidity: function rewardGlobalBps() view returns(uint16)
func (_StakeNew *StakeNewCaller) RewardGlobalBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "rewardGlobalBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RewardGlobalBps is a free data retrieval call binding the contract method 0xce765062.
//
// Solidity: function rewardGlobalBps() view returns(uint16)
func (_StakeNew *StakeNewSession) RewardGlobalBps() (uint16, error) {
	return _StakeNew.Contract.RewardGlobalBps(&_StakeNew.CallOpts)
}

// RewardGlobalBps is a free data retrieval call binding the contract method 0xce765062.
//
// Solidity: function rewardGlobalBps() view returns(uint16)
func (_StakeNew *StakeNewCallerSession) RewardGlobalBps() (uint16, error) {
	return _StakeNew.Contract.RewardGlobalBps(&_StakeNew.CallOpts)
}

// RewardS7Bps is a free data retrieval call binding the contract method 0x2c00552f.
//
// Solidity: function rewardS7Bps() view returns(uint16)
func (_StakeNew *StakeNewCaller) RewardS7Bps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "rewardS7Bps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RewardS7Bps is a free data retrieval call binding the contract method 0x2c00552f.
//
// Solidity: function rewardS7Bps() view returns(uint16)
func (_StakeNew *StakeNewSession) RewardS7Bps() (uint16, error) {
	return _StakeNew.Contract.RewardS7Bps(&_StakeNew.CallOpts)
}

// RewardS7Bps is a free data retrieval call binding the contract method 0x2c00552f.
//
// Solidity: function rewardS7Bps() view returns(uint16)
func (_StakeNew *StakeNewCallerSession) RewardS7Bps() (uint16, error) {
	return _StakeNew.Contract.RewardS7Bps(&_StakeNew.CallOpts)
}

// RewardTeamMaxBps is a free data retrieval call binding the contract method 0xb1897132.
//
// Solidity: function rewardTeamMaxBps() view returns(uint16)
func (_StakeNew *StakeNewCaller) RewardTeamMaxBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "rewardTeamMaxBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RewardTeamMaxBps is a free data retrieval call binding the contract method 0xb1897132.
//
// Solidity: function rewardTeamMaxBps() view returns(uint16)
func (_StakeNew *StakeNewSession) RewardTeamMaxBps() (uint16, error) {
	return _StakeNew.Contract.RewardTeamMaxBps(&_StakeNew.CallOpts)
}

// RewardTeamMaxBps is a free data retrieval call binding the contract method 0xb1897132.
//
// Solidity: function rewardTeamMaxBps() view returns(uint16)
func (_StakeNew *StakeNewCallerSession) RewardTeamMaxBps() (uint16, error) {
	return _StakeNew.Contract.RewardTeamMaxBps(&_StakeNew.CallOpts)
}

// RouteAddress1 is a free data retrieval call binding the contract method 0xfd0dab4f.
//
// Solidity: function routeAddress1() view returns(address)
func (_StakeNew *StakeNewCaller) RouteAddress1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "routeAddress1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RouteAddress1 is a free data retrieval call binding the contract method 0xfd0dab4f.
//
// Solidity: function routeAddress1() view returns(address)
func (_StakeNew *StakeNewSession) RouteAddress1() (common.Address, error) {
	return _StakeNew.Contract.RouteAddress1(&_StakeNew.CallOpts)
}

// RouteAddress1 is a free data retrieval call binding the contract method 0xfd0dab4f.
//
// Solidity: function routeAddress1() view returns(address)
func (_StakeNew *StakeNewCallerSession) RouteAddress1() (common.Address, error) {
	return _StakeNew.Contract.RouteAddress1(&_StakeNew.CallOpts)
}

// RouteAddress2 is a free data retrieval call binding the contract method 0x26504cb9.
//
// Solidity: function routeAddress2() view returns(address)
func (_StakeNew *StakeNewCaller) RouteAddress2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "routeAddress2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RouteAddress2 is a free data retrieval call binding the contract method 0x26504cb9.
//
// Solidity: function routeAddress2() view returns(address)
func (_StakeNew *StakeNewSession) RouteAddress2() (common.Address, error) {
	return _StakeNew.Contract.RouteAddress2(&_StakeNew.CallOpts)
}

// RouteAddress2 is a free data retrieval call binding the contract method 0x26504cb9.
//
// Solidity: function routeAddress2() view returns(address)
func (_StakeNew *StakeNewCallerSession) RouteAddress2() (common.Address, error) {
	return _StakeNew.Contract.RouteAddress2(&_StakeNew.CallOpts)
}

// ShouldUpdateCircuitBreaker is a free data retrieval call binding the contract method 0xdcac1a16.
//
// Solidity: function shouldUpdateCircuitBreaker() view returns(bool needTrigger, bool needCountdown, bool needRecover, uint256 currentDropBps)
func (_StakeNew *StakeNewCaller) ShouldUpdateCircuitBreaker(opts *bind.CallOpts) (struct {
	NeedTrigger    bool
	NeedCountdown  bool
	NeedRecover    bool
	CurrentDropBps *big.Int
}, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "shouldUpdateCircuitBreaker")

	outstruct := new(struct {
		NeedTrigger    bool
		NeedCountdown  bool
		NeedRecover    bool
		CurrentDropBps *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NeedTrigger = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.NeedCountdown = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.NeedRecover = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.CurrentDropBps = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ShouldUpdateCircuitBreaker is a free data retrieval call binding the contract method 0xdcac1a16.
//
// Solidity: function shouldUpdateCircuitBreaker() view returns(bool needTrigger, bool needCountdown, bool needRecover, uint256 currentDropBps)
func (_StakeNew *StakeNewSession) ShouldUpdateCircuitBreaker() (struct {
	NeedTrigger    bool
	NeedCountdown  bool
	NeedRecover    bool
	CurrentDropBps *big.Int
}, error) {
	return _StakeNew.Contract.ShouldUpdateCircuitBreaker(&_StakeNew.CallOpts)
}

// ShouldUpdateCircuitBreaker is a free data retrieval call binding the contract method 0xdcac1a16.
//
// Solidity: function shouldUpdateCircuitBreaker() view returns(bool needTrigger, bool needCountdown, bool needRecover, uint256 currentDropBps)
func (_StakeNew *StakeNewCallerSession) ShouldUpdateCircuitBreaker() (struct {
	NeedTrigger    bool
	NeedCountdown  bool
	NeedRecover    bool
	CurrentDropBps *big.Int
}, error) {
	return _StakeNew.Contract.ShouldUpdateCircuitBreaker(&_StakeNew.CallOpts)
}

// Stage1Threshold is a free data retrieval call binding the contract method 0x2f13e5d9.
//
// Solidity: function stage1Threshold() view returns(uint256)
func (_StakeNew *StakeNewCaller) Stage1Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "stage1Threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Stage1Threshold is a free data retrieval call binding the contract method 0x2f13e5d9.
//
// Solidity: function stage1Threshold() view returns(uint256)
func (_StakeNew *StakeNewSession) Stage1Threshold() (*big.Int, error) {
	return _StakeNew.Contract.Stage1Threshold(&_StakeNew.CallOpts)
}

// Stage1Threshold is a free data retrieval call binding the contract method 0x2f13e5d9.
//
// Solidity: function stage1Threshold() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) Stage1Threshold() (*big.Int, error) {
	return _StakeNew.Contract.Stage1Threshold(&_StakeNew.CallOpts)
}

// Stage2Threshold is a free data retrieval call binding the contract method 0x54682443.
//
// Solidity: function stage2Threshold() view returns(uint256)
func (_StakeNew *StakeNewCaller) Stage2Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "stage2Threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Stage2Threshold is a free data retrieval call binding the contract method 0x54682443.
//
// Solidity: function stage2Threshold() view returns(uint256)
func (_StakeNew *StakeNewSession) Stage2Threshold() (*big.Int, error) {
	return _StakeNew.Contract.Stage2Threshold(&_StakeNew.CallOpts)
}

// Stage2Threshold is a free data retrieval call binding the contract method 0x54682443.
//
// Solidity: function stage2Threshold() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) Stage2Threshold() (*big.Int, error) {
	return _StakeNew.Contract.Stage2Threshold(&_StakeNew.CallOpts)
}

// StakeCount is a free data retrieval call binding the contract method 0x33060d90.
//
// Solidity: function stakeCount(address user) view returns(uint256)
func (_StakeNew *StakeNewCaller) StakeCount(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "stakeCount", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeCount is a free data retrieval call binding the contract method 0x33060d90.
//
// Solidity: function stakeCount(address user) view returns(uint256)
func (_StakeNew *StakeNewSession) StakeCount(user common.Address) (*big.Int, error) {
	return _StakeNew.Contract.StakeCount(&_StakeNew.CallOpts, user)
}

// StakeCount is a free data retrieval call binding the contract method 0x33060d90.
//
// Solidity: function stakeCount(address user) view returns(uint256)
func (_StakeNew *StakeNewCallerSession) StakeCount(user common.Address) (*big.Int, error) {
	return _StakeNew.Contract.StakeCount(&_StakeNew.CallOpts, user)
}

// StakeQueue is a free data retrieval call binding the contract method 0x0177d287.
//
// Solidity: function stakeQueue(uint256 ) view returns(address user, uint160 amount, uint256 amountOutMin, uint8 stakeIndex, uint40 queuedAt, bool canceled)
func (_StakeNew *StakeNewCaller) StakeQueue(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User         common.Address
	Amount       *big.Int
	AmountOutMin *big.Int
	StakeIndex   uint8
	QueuedAt     *big.Int
	Canceled     bool
}, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "stakeQueue", arg0)

	outstruct := new(struct {
		User         common.Address
		Amount       *big.Int
		AmountOutMin *big.Int
		StakeIndex   uint8
		QueuedAt     *big.Int
		Canceled     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AmountOutMin = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StakeIndex = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.QueuedAt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Canceled = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// StakeQueue is a free data retrieval call binding the contract method 0x0177d287.
//
// Solidity: function stakeQueue(uint256 ) view returns(address user, uint160 amount, uint256 amountOutMin, uint8 stakeIndex, uint40 queuedAt, bool canceled)
func (_StakeNew *StakeNewSession) StakeQueue(arg0 *big.Int) (struct {
	User         common.Address
	Amount       *big.Int
	AmountOutMin *big.Int
	StakeIndex   uint8
	QueuedAt     *big.Int
	Canceled     bool
}, error) {
	return _StakeNew.Contract.StakeQueue(&_StakeNew.CallOpts, arg0)
}

// StakeQueue is a free data retrieval call binding the contract method 0x0177d287.
//
// Solidity: function stakeQueue(uint256 ) view returns(address user, uint160 amount, uint256 amountOutMin, uint8 stakeIndex, uint40 queuedAt, bool canceled)
func (_StakeNew *StakeNewCallerSession) StakeQueue(arg0 *big.Int) (struct {
	User         common.Address
	Amount       *big.Int
	AmountOutMin *big.Int
	StakeIndex   uint8
	QueuedAt     *big.Int
	Canceled     bool
}, error) {
	return _StakeNew.Contract.StakeQueue(&_StakeNew.CallOpts, arg0)
}

// StakeQueueHandleCount is a free data retrieval call binding the contract method 0x161ca758.
//
// Solidity: function stakeQueueHandleCount() view returns(uint256)
func (_StakeNew *StakeNewCaller) StakeQueueHandleCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "stakeQueueHandleCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeQueueHandleCount is a free data retrieval call binding the contract method 0x161ca758.
//
// Solidity: function stakeQueueHandleCount() view returns(uint256)
func (_StakeNew *StakeNewSession) StakeQueueHandleCount() (*big.Int, error) {
	return _StakeNew.Contract.StakeQueueHandleCount(&_StakeNew.CallOpts)
}

// StakeQueueHandleCount is a free data retrieval call binding the contract method 0x161ca758.
//
// Solidity: function stakeQueueHandleCount() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) StakeQueueHandleCount() (*big.Int, error) {
	return _StakeNew.Contract.StakeQueueHandleCount(&_StakeNew.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StakeNew *StakeNewCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StakeNew *StakeNewSession) Symbol() (string, error) {
	return _StakeNew.Contract.Symbol(&_StakeNew.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StakeNew *StakeNewCallerSession) Symbol() (string, error) {
	return _StakeNew.Contract.Symbol(&_StakeNew.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakeNew *StakeNewCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakeNew *StakeNewSession) TotalSupply() (*big.Int, error) {
	return _StakeNew.Contract.TotalSupply(&_StakeNew.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakeNew *StakeNewCallerSession) TotalSupply() (*big.Int, error) {
	return _StakeNew.Contract.TotalSupply(&_StakeNew.CallOpts)
}

// UserStakeRecord is a free data retrieval call binding the contract method 0x08adb4be.
//
// Solidity: function userStakeRecord(address , uint256 ) view returns(uint40 stakeTime, uint160 amount, uint8 stakeIndex, uint40 unstakeTime, uint160 reward, uint40 restakeTime, bool cbNewOrder, uint8 status)
func (_StakeNew *StakeNewCaller) UserStakeRecord(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StakeTime   *big.Int
	Amount      *big.Int
	StakeIndex  uint8
	UnstakeTime *big.Int
	Reward      *big.Int
	RestakeTime *big.Int
	CbNewOrder  bool
	Status      uint8
}, error) {
	var out []interface{}
	err := _StakeNew.contract.Call(opts, &out, "userStakeRecord", arg0, arg1)

	outstruct := new(struct {
		StakeTime   *big.Int
		Amount      *big.Int
		StakeIndex  uint8
		UnstakeTime *big.Int
		Reward      *big.Int
		RestakeTime *big.Int
		CbNewOrder  bool
		Status      uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakeTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StakeIndex = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.UnstakeTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Reward = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RestakeTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.CbNewOrder = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)

	return *outstruct, err

}

// UserStakeRecord is a free data retrieval call binding the contract method 0x08adb4be.
//
// Solidity: function userStakeRecord(address , uint256 ) view returns(uint40 stakeTime, uint160 amount, uint8 stakeIndex, uint40 unstakeTime, uint160 reward, uint40 restakeTime, bool cbNewOrder, uint8 status)
func (_StakeNew *StakeNewSession) UserStakeRecord(arg0 common.Address, arg1 *big.Int) (struct {
	StakeTime   *big.Int
	Amount      *big.Int
	StakeIndex  uint8
	UnstakeTime *big.Int
	Reward      *big.Int
	RestakeTime *big.Int
	CbNewOrder  bool
	Status      uint8
}, error) {
	return _StakeNew.Contract.UserStakeRecord(&_StakeNew.CallOpts, arg0, arg1)
}

// UserStakeRecord is a free data retrieval call binding the contract method 0x08adb4be.
//
// Solidity: function userStakeRecord(address , uint256 ) view returns(uint40 stakeTime, uint160 amount, uint8 stakeIndex, uint40 unstakeTime, uint160 reward, uint40 restakeTime, bool cbNewOrder, uint8 status)
func (_StakeNew *StakeNewCallerSession) UserStakeRecord(arg0 common.Address, arg1 *big.Int) (struct {
	StakeTime   *big.Int
	Amount      *big.Int
	StakeIndex  uint8
	UnstakeTime *big.Int
	Reward      *big.Int
	RestakeTime *big.Int
	CbNewOrder  bool
	Status      uint8
}, error) {
	return _StakeNew.Contract.UserStakeRecord(&_StakeNew.CallOpts, arg0, arg1)
}

// AdminCancelQueuedStake is a paid mutator transaction binding the contract method 0x95ec5770.
//
// Solidity: function adminCancelQueuedStake(uint256 queueIndex) returns()
func (_StakeNew *StakeNewTransactor) AdminCancelQueuedStake(opts *bind.TransactOpts, queueIndex *big.Int) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "adminCancelQueuedStake", queueIndex)
}

// AdminCancelQueuedStake is a paid mutator transaction binding the contract method 0x95ec5770.
//
// Solidity: function adminCancelQueuedStake(uint256 queueIndex) returns()
func (_StakeNew *StakeNewSession) AdminCancelQueuedStake(queueIndex *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.AdminCancelQueuedStake(&_StakeNew.TransactOpts, queueIndex)
}

// AdminCancelQueuedStake is a paid mutator transaction binding the contract method 0x95ec5770.
//
// Solidity: function adminCancelQueuedStake(uint256 queueIndex) returns()
func (_StakeNew *StakeNewTransactorSession) AdminCancelQueuedStake(queueIndex *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.AdminCancelQueuedStake(&_StakeNew.TransactOpts, queueIndex)
}

// AdminCancelQueuedStakes is a paid mutator transaction binding the contract method 0x403b1db8.
//
// Solidity: function adminCancelQueuedStakes(uint256[] queueIndexes) returns()
func (_StakeNew *StakeNewTransactor) AdminCancelQueuedStakes(opts *bind.TransactOpts, queueIndexes []*big.Int) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "adminCancelQueuedStakes", queueIndexes)
}

// AdminCancelQueuedStakes is a paid mutator transaction binding the contract method 0x403b1db8.
//
// Solidity: function adminCancelQueuedStakes(uint256[] queueIndexes) returns()
func (_StakeNew *StakeNewSession) AdminCancelQueuedStakes(queueIndexes []*big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.AdminCancelQueuedStakes(&_StakeNew.TransactOpts, queueIndexes)
}

// AdminCancelQueuedStakes is a paid mutator transaction binding the contract method 0x403b1db8.
//
// Solidity: function adminCancelQueuedStakes(uint256[] queueIndexes) returns()
func (_StakeNew *StakeNewTransactorSession) AdminCancelQueuedStakes(queueIndexes []*big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.AdminCancelQueuedStakes(&_StakeNew.TransactOpts, queueIndexes)
}

// BurnExpiredReward is a paid mutator transaction binding the contract method 0x51d2376a.
//
// Solidity: function burnExpiredReward(address user, uint256 index) returns()
func (_StakeNew *StakeNewTransactor) BurnExpiredReward(opts *bind.TransactOpts, user common.Address, index *big.Int) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "burnExpiredReward", user, index)
}

// BurnExpiredReward is a paid mutator transaction binding the contract method 0x51d2376a.
//
// Solidity: function burnExpiredReward(address user, uint256 index) returns()
func (_StakeNew *StakeNewSession) BurnExpiredReward(user common.Address, index *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.BurnExpiredReward(&_StakeNew.TransactOpts, user, index)
}

// BurnExpiredReward is a paid mutator transaction binding the contract method 0x51d2376a.
//
// Solidity: function burnExpiredReward(address user, uint256 index) returns()
func (_StakeNew *StakeNewTransactorSession) BurnExpiredReward(user common.Address, index *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.BurnExpiredReward(&_StakeNew.TransactOpts, user, index)
}

// CancelQueuedStake is a paid mutator transaction binding the contract method 0xc9f5a5d8.
//
// Solidity: function cancelQueuedStake(uint256 queueIndex) returns()
func (_StakeNew *StakeNewTransactor) CancelQueuedStake(opts *bind.TransactOpts, queueIndex *big.Int) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "cancelQueuedStake", queueIndex)
}

// CancelQueuedStake is a paid mutator transaction binding the contract method 0xc9f5a5d8.
//
// Solidity: function cancelQueuedStake(uint256 queueIndex) returns()
func (_StakeNew *StakeNewSession) CancelQueuedStake(queueIndex *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.CancelQueuedStake(&_StakeNew.TransactOpts, queueIndex)
}

// CancelQueuedStake is a paid mutator transaction binding the contract method 0xc9f5a5d8.
//
// Solidity: function cancelQueuedStake(uint256 queueIndex) returns()
func (_StakeNew *StakeNewTransactorSession) CancelQueuedStake(queueIndex *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.CancelQueuedStake(&_StakeNew.TransactOpts, queueIndex)
}

// ChangeFreeUsdtCollector is a paid mutator transaction binding the contract method 0x0aa258e8.
//
// Solidity: function changeFreeUsdtCollector(address newCollector) returns()
func (_StakeNew *StakeNewTransactor) ChangeFreeUsdtCollector(opts *bind.TransactOpts, newCollector common.Address) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "changeFreeUsdtCollector", newCollector)
}

// ChangeFreeUsdtCollector is a paid mutator transaction binding the contract method 0x0aa258e8.
//
// Solidity: function changeFreeUsdtCollector(address newCollector) returns()
func (_StakeNew *StakeNewSession) ChangeFreeUsdtCollector(newCollector common.Address) (*types.Transaction, error) {
	return _StakeNew.Contract.ChangeFreeUsdtCollector(&_StakeNew.TransactOpts, newCollector)
}

// ChangeFreeUsdtCollector is a paid mutator transaction binding the contract method 0x0aa258e8.
//
// Solidity: function changeFreeUsdtCollector(address newCollector) returns()
func (_StakeNew *StakeNewTransactorSession) ChangeFreeUsdtCollector(newCollector common.Address) (*types.Transaction, error) {
	return _StakeNew.Contract.ChangeFreeUsdtCollector(&_StakeNew.TransactOpts, newCollector)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address user, uint256 index) returns()
func (_StakeNew *StakeNewTransactor) Claim(opts *bind.TransactOpts, user common.Address, index *big.Int) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "claim", user, index)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address user, uint256 index) returns()
func (_StakeNew *StakeNewSession) Claim(user common.Address, index *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.Claim(&_StakeNew.TransactOpts, user, index)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address user, uint256 index) returns()
func (_StakeNew *StakeNewTransactorSession) Claim(user common.Address, index *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.Claim(&_StakeNew.TransactOpts, user, index)
}

// ExecuteQueuedStake is a paid mutator transaction binding the contract method 0xb74470cd.
//
// Solidity: function executeQueuedStake(uint256 queueIndex) returns()
func (_StakeNew *StakeNewTransactor) ExecuteQueuedStake(opts *bind.TransactOpts, queueIndex *big.Int) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "executeQueuedStake", queueIndex)
}

// ExecuteQueuedStake is a paid mutator transaction binding the contract method 0xb74470cd.
//
// Solidity: function executeQueuedStake(uint256 queueIndex) returns()
func (_StakeNew *StakeNewSession) ExecuteQueuedStake(queueIndex *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.ExecuteQueuedStake(&_StakeNew.TransactOpts, queueIndex)
}

// ExecuteQueuedStake is a paid mutator transaction binding the contract method 0xb74470cd.
//
// Solidity: function executeQueuedStake(uint256 queueIndex) returns()
func (_StakeNew *StakeNewTransactorSession) ExecuteQueuedStake(queueIndex *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.ExecuteQueuedStake(&_StakeNew.TransactOpts, queueIndex)
}

// InitSystems is a paid mutator transaction binding the contract method 0x974c9792.
//
// Solidity: function initSystems(address token_, address userSystem_, address dividendSystem_) returns()
func (_StakeNew *StakeNewTransactor) InitSystems(opts *bind.TransactOpts, token_ common.Address, userSystem_ common.Address, dividendSystem_ common.Address) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "initSystems", token_, userSystem_, dividendSystem_)
}

// InitSystems is a paid mutator transaction binding the contract method 0x974c9792.
//
// Solidity: function initSystems(address token_, address userSystem_, address dividendSystem_) returns()
func (_StakeNew *StakeNewSession) InitSystems(token_ common.Address, userSystem_ common.Address, dividendSystem_ common.Address) (*types.Transaction, error) {
	return _StakeNew.Contract.InitSystems(&_StakeNew.TransactOpts, token_, userSystem_, dividendSystem_)
}

// InitSystems is a paid mutator transaction binding the contract method 0x974c9792.
//
// Solidity: function initSystems(address token_, address userSystem_, address dividendSystem_) returns()
func (_StakeNew *StakeNewTransactorSession) InitSystems(token_ common.Address, userSystem_ common.Address, dividendSystem_ common.Address) (*types.Transaction, error) {
	return _StakeNew.Contract.InitSystems(&_StakeNew.TransactOpts, token_, userSystem_, dividendSystem_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakeNew *StakeNewTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakeNew *StakeNewSession) Pause() (*types.Transaction, error) {
	return _StakeNew.Contract.Pause(&_StakeNew.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakeNew *StakeNewTransactorSession) Pause() (*types.Transaction, error) {
	return _StakeNew.Contract.Pause(&_StakeNew.TransactOpts)
}

// ProcessStakeQueue is a paid mutator transaction binding the contract method 0xfbf75b1e.
//
// Solidity: function processStakeQueue(uint256 maxItems) returns()
func (_StakeNew *StakeNewTransactor) ProcessStakeQueue(opts *bind.TransactOpts, maxItems *big.Int) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "processStakeQueue", maxItems)
}

// ProcessStakeQueue is a paid mutator transaction binding the contract method 0xfbf75b1e.
//
// Solidity: function processStakeQueue(uint256 maxItems) returns()
func (_StakeNew *StakeNewSession) ProcessStakeQueue(maxItems *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.ProcessStakeQueue(&_StakeNew.TransactOpts, maxItems)
}

// ProcessStakeQueue is a paid mutator transaction binding the contract method 0xfbf75b1e.
//
// Solidity: function processStakeQueue(uint256 maxItems) returns()
func (_StakeNew *StakeNewTransactorSession) ProcessStakeQueue(maxItems *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.ProcessStakeQueue(&_StakeNew.TransactOpts, maxItems)
}

// Restake is a paid mutator transaction binding the contract method 0x049686bf.
//
// Solidity: function restake(address user, uint256 index, uint160 amount, uint256 amountOutMin, uint8 stakeIndex) returns()
func (_StakeNew *StakeNewTransactor) Restake(opts *bind.TransactOpts, user common.Address, index *big.Int, amount *big.Int, amountOutMin *big.Int, stakeIndex uint8) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "restake", user, index, amount, amountOutMin, stakeIndex)
}

// Restake is a paid mutator transaction binding the contract method 0x049686bf.
//
// Solidity: function restake(address user, uint256 index, uint160 amount, uint256 amountOutMin, uint8 stakeIndex) returns()
func (_StakeNew *StakeNewSession) Restake(user common.Address, index *big.Int, amount *big.Int, amountOutMin *big.Int, stakeIndex uint8) (*types.Transaction, error) {
	return _StakeNew.Contract.Restake(&_StakeNew.TransactOpts, user, index, amount, amountOutMin, stakeIndex)
}

// Restake is a paid mutator transaction binding the contract method 0x049686bf.
//
// Solidity: function restake(address user, uint256 index, uint160 amount, uint256 amountOutMin, uint8 stakeIndex) returns()
func (_StakeNew *StakeNewTransactorSession) Restake(user common.Address, index *big.Int, amount *big.Int, amountOutMin *big.Int, stakeIndex uint8) (*types.Transaction, error) {
	return _StakeNew.Contract.Restake(&_StakeNew.TransactOpts, user, index, amount, amountOutMin, stakeIndex)
}

// SetCircuitBreakerParams is a paid mutator transaction binding the contract method 0xcbca8f47.
//
// Solidity: function setCircuitBreakerParams(uint256 a, uint40 b, uint256 c) returns()
func (_StakeNew *StakeNewTransactor) SetCircuitBreakerParams(opts *bind.TransactOpts, a *big.Int, b *big.Int, c *big.Int) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "setCircuitBreakerParams", a, b, c)
}

// SetCircuitBreakerParams is a paid mutator transaction binding the contract method 0xcbca8f47.
//
// Solidity: function setCircuitBreakerParams(uint256 a, uint40 b, uint256 c) returns()
func (_StakeNew *StakeNewSession) SetCircuitBreakerParams(a *big.Int, b *big.Int, c *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.SetCircuitBreakerParams(&_StakeNew.TransactOpts, a, b, c)
}

// SetCircuitBreakerParams is a paid mutator transaction binding the contract method 0xcbca8f47.
//
// Solidity: function setCircuitBreakerParams(uint256 a, uint40 b, uint256 c) returns()
func (_StakeNew *StakeNewTransactorSession) SetCircuitBreakerParams(a *big.Int, b *big.Int, c *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.SetCircuitBreakerParams(&_StakeNew.TransactOpts, a, b, c)
}

// SetDailyLimitRates is a paid mutator transaction binding the contract method 0xa61236f7.
//
// Solidity: function setDailyLimitRates(uint16 a, uint16 b, uint16 c, uint16 d, uint16 e, uint16 f) returns()
func (_StakeNew *StakeNewTransactor) SetDailyLimitRates(opts *bind.TransactOpts, a uint16, b uint16, c uint16, d uint16, e uint16, f uint16) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "setDailyLimitRates", a, b, c, d, e, f)
}

// SetDailyLimitRates is a paid mutator transaction binding the contract method 0xa61236f7.
//
// Solidity: function setDailyLimitRates(uint16 a, uint16 b, uint16 c, uint16 d, uint16 e, uint16 f) returns()
func (_StakeNew *StakeNewSession) SetDailyLimitRates(a uint16, b uint16, c uint16, d uint16, e uint16, f uint16) (*types.Transaction, error) {
	return _StakeNew.Contract.SetDailyLimitRates(&_StakeNew.TransactOpts, a, b, c, d, e, f)
}

// SetDailyLimitRates is a paid mutator transaction binding the contract method 0xa61236f7.
//
// Solidity: function setDailyLimitRates(uint16 a, uint16 b, uint16 c, uint16 d, uint16 e, uint16 f) returns()
func (_StakeNew *StakeNewTransactorSession) SetDailyLimitRates(a uint16, b uint16, c uint16, d uint16, e uint16, f uint16) (*types.Transaction, error) {
	return _StakeNew.Contract.SetDailyLimitRates(&_StakeNew.TransactOpts, a, b, c, d, e, f)
}

// SetDailyQueueCancelLimit is a paid mutator transaction binding the contract method 0x6c095f75.
//
// Solidity: function setDailyQueueCancelLimit(uint256 a) returns()
func (_StakeNew *StakeNewTransactor) SetDailyQueueCancelLimit(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "setDailyQueueCancelLimit", a)
}

// SetDailyQueueCancelLimit is a paid mutator transaction binding the contract method 0x6c095f75.
//
// Solidity: function setDailyQueueCancelLimit(uint256 a) returns()
func (_StakeNew *StakeNewSession) SetDailyQueueCancelLimit(a *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.SetDailyQueueCancelLimit(&_StakeNew.TransactOpts, a)
}

// SetDailyQueueCancelLimit is a paid mutator transaction binding the contract method 0x6c095f75.
//
// Solidity: function setDailyQueueCancelLimit(uint256 a) returns()
func (_StakeNew *StakeNewTransactorSession) SetDailyQueueCancelLimit(a *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.SetDailyQueueCancelLimit(&_StakeNew.TransactOpts, a)
}

// SetGateway is a paid mutator transaction binding the contract method 0x90646b4a.
//
// Solidity: function setGateway(address a) returns()
func (_StakeNew *StakeNewTransactor) SetGateway(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "setGateway", a)
}

// SetGateway is a paid mutator transaction binding the contract method 0x90646b4a.
//
// Solidity: function setGateway(address a) returns()
func (_StakeNew *StakeNewSession) SetGateway(a common.Address) (*types.Transaction, error) {
	return _StakeNew.Contract.SetGateway(&_StakeNew.TransactOpts, a)
}

// SetGateway is a paid mutator transaction binding the contract method 0x90646b4a.
//
// Solidity: function setGateway(address a) returns()
func (_StakeNew *StakeNewTransactorSession) SetGateway(a common.Address) (*types.Transaction, error) {
	return _StakeNew.Contract.SetGateway(&_StakeNew.TransactOpts, a)
}

// SetInitAccount is a paid mutator transaction binding the contract method 0x42b34815.
//
// Solidity: function setInitAccount(address newInitAccount) returns()
func (_StakeNew *StakeNewTransactor) SetInitAccount(opts *bind.TransactOpts, newInitAccount common.Address) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "setInitAccount", newInitAccount)
}

// SetInitAccount is a paid mutator transaction binding the contract method 0x42b34815.
//
// Solidity: function setInitAccount(address newInitAccount) returns()
func (_StakeNew *StakeNewSession) SetInitAccount(newInitAccount common.Address) (*types.Transaction, error) {
	return _StakeNew.Contract.SetInitAccount(&_StakeNew.TransactOpts, newInitAccount)
}

// SetInitAccount is a paid mutator transaction binding the contract method 0x42b34815.
//
// Solidity: function setInitAccount(address newInitAccount) returns()
func (_StakeNew *StakeNewTransactorSession) SetInitAccount(newInitAccount common.Address) (*types.Transaction, error) {
	return _StakeNew.Contract.SetInitAccount(&_StakeNew.TransactOpts, newInitAccount)
}

// SetMarketingAddress is a paid mutator transaction binding the contract method 0x906e9dd0.
//
// Solidity: function setMarketingAddress(address a) returns()
func (_StakeNew *StakeNewTransactor) SetMarketingAddress(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "setMarketingAddress", a)
}

// SetMarketingAddress is a paid mutator transaction binding the contract method 0x906e9dd0.
//
// Solidity: function setMarketingAddress(address a) returns()
func (_StakeNew *StakeNewSession) SetMarketingAddress(a common.Address) (*types.Transaction, error) {
	return _StakeNew.Contract.SetMarketingAddress(&_StakeNew.TransactOpts, a)
}

// SetMarketingAddress is a paid mutator transaction binding the contract method 0x906e9dd0.
//
// Solidity: function setMarketingAddress(address a) returns()
func (_StakeNew *StakeNewTransactorSession) SetMarketingAddress(a common.Address) (*types.Transaction, error) {
	return _StakeNew.Contract.SetMarketingAddress(&_StakeNew.TransactOpts, a)
}

// SetMaxReferralDepth is a paid mutator transaction binding the contract method 0xb006aed2.
//
// Solidity: function setMaxReferralDepth(uint256 a) returns()
func (_StakeNew *StakeNewTransactor) SetMaxReferralDepth(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "setMaxReferralDepth", a)
}

// SetMaxReferralDepth is a paid mutator transaction binding the contract method 0xb006aed2.
//
// Solidity: function setMaxReferralDepth(uint256 a) returns()
func (_StakeNew *StakeNewSession) SetMaxReferralDepth(a *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.SetMaxReferralDepth(&_StakeNew.TransactOpts, a)
}

// SetMaxReferralDepth is a paid mutator transaction binding the contract method 0xb006aed2.
//
// Solidity: function setMaxReferralDepth(uint256 a) returns()
func (_StakeNew *StakeNewTransactorSession) SetMaxReferralDepth(a *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.SetMaxReferralDepth(&_StakeNew.TransactOpts, a)
}

// SetRewardRoutes is a paid mutator transaction binding the contract method 0x84d946f0.
//
// Solidity: function setRewardRoutes(address a, address b) returns()
func (_StakeNew *StakeNewTransactor) SetRewardRoutes(opts *bind.TransactOpts, a common.Address, b common.Address) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "setRewardRoutes", a, b)
}

// SetRewardRoutes is a paid mutator transaction binding the contract method 0x84d946f0.
//
// Solidity: function setRewardRoutes(address a, address b) returns()
func (_StakeNew *StakeNewSession) SetRewardRoutes(a common.Address, b common.Address) (*types.Transaction, error) {
	return _StakeNew.Contract.SetRewardRoutes(&_StakeNew.TransactOpts, a, b)
}

// SetRewardRoutes is a paid mutator transaction binding the contract method 0x84d946f0.
//
// Solidity: function setRewardRoutes(address a, address b) returns()
func (_StakeNew *StakeNewTransactorSession) SetRewardRoutes(a common.Address, b common.Address) (*types.Transaction, error) {
	return _StakeNew.Contract.SetRewardRoutes(&_StakeNew.TransactOpts, a, b)
}

// SetStageThresholds is a paid mutator transaction binding the contract method 0x167e3a12.
//
// Solidity: function setStageThresholds(uint256 a, uint256 b) returns()
func (_StakeNew *StakeNewTransactor) SetStageThresholds(opts *bind.TransactOpts, a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "setStageThresholds", a, b)
}

// SetStageThresholds is a paid mutator transaction binding the contract method 0x167e3a12.
//
// Solidity: function setStageThresholds(uint256 a, uint256 b) returns()
func (_StakeNew *StakeNewSession) SetStageThresholds(a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.SetStageThresholds(&_StakeNew.TransactOpts, a, b)
}

// SetStageThresholds is a paid mutator transaction binding the contract method 0x167e3a12.
//
// Solidity: function setStageThresholds(uint256 a, uint256 b) returns()
func (_StakeNew *StakeNewTransactorSession) SetStageThresholds(a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.SetStageThresholds(&_StakeNew.TransactOpts, a, b)
}

// SetStakeQueueHandleCount is a paid mutator transaction binding the contract method 0x25bd6184.
//
// Solidity: function setStakeQueueHandleCount(uint256 a) returns()
func (_StakeNew *StakeNewTransactor) SetStakeQueueHandleCount(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "setStakeQueueHandleCount", a)
}

// SetStakeQueueHandleCount is a paid mutator transaction binding the contract method 0x25bd6184.
//
// Solidity: function setStakeQueueHandleCount(uint256 a) returns()
func (_StakeNew *StakeNewSession) SetStakeQueueHandleCount(a *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.SetStakeQueueHandleCount(&_StakeNew.TransactOpts, a)
}

// SetStakeQueueHandleCount is a paid mutator transaction binding the contract method 0x25bd6184.
//
// Solidity: function setStakeQueueHandleCount(uint256 a) returns()
func (_StakeNew *StakeNewTransactorSession) SetStakeQueueHandleCount(a *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.SetStakeQueueHandleCount(&_StakeNew.TransactOpts, a)
}

// Stake is a paid mutator transaction binding the contract method 0x7fce8b1a.
//
// Solidity: function stake(address user, uint160 amount, uint256 amountOutMin, uint8 stakeIndex) returns()
func (_StakeNew *StakeNewTransactor) Stake(opts *bind.TransactOpts, user common.Address, amount *big.Int, amountOutMin *big.Int, stakeIndex uint8) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "stake", user, amount, amountOutMin, stakeIndex)
}

// Stake is a paid mutator transaction binding the contract method 0x7fce8b1a.
//
// Solidity: function stake(address user, uint160 amount, uint256 amountOutMin, uint8 stakeIndex) returns()
func (_StakeNew *StakeNewSession) Stake(user common.Address, amount *big.Int, amountOutMin *big.Int, stakeIndex uint8) (*types.Transaction, error) {
	return _StakeNew.Contract.Stake(&_StakeNew.TransactOpts, user, amount, amountOutMin, stakeIndex)
}

// Stake is a paid mutator transaction binding the contract method 0x7fce8b1a.
//
// Solidity: function stake(address user, uint160 amount, uint256 amountOutMin, uint8 stakeIndex) returns()
func (_StakeNew *StakeNewTransactorSession) Stake(user common.Address, amount *big.Int, amountOutMin *big.Int, stakeIndex uint8) (*types.Transaction, error) {
	return _StakeNew.Contract.Stake(&_StakeNew.TransactOpts, user, amount, amountOutMin, stakeIndex)
}

// StakeWithInviter is a paid mutator transaction binding the contract method 0xc82b6d1c.
//
// Solidity: function stakeWithInviter(address user, uint160 amount, uint256 amountOutMin, uint8 stakeIndex, address parent) returns()
func (_StakeNew *StakeNewTransactor) StakeWithInviter(opts *bind.TransactOpts, user common.Address, amount *big.Int, amountOutMin *big.Int, stakeIndex uint8, parent common.Address) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "stakeWithInviter", user, amount, amountOutMin, stakeIndex, parent)
}

// StakeWithInviter is a paid mutator transaction binding the contract method 0xc82b6d1c.
//
// Solidity: function stakeWithInviter(address user, uint160 amount, uint256 amountOutMin, uint8 stakeIndex, address parent) returns()
func (_StakeNew *StakeNewSession) StakeWithInviter(user common.Address, amount *big.Int, amountOutMin *big.Int, stakeIndex uint8, parent common.Address) (*types.Transaction, error) {
	return _StakeNew.Contract.StakeWithInviter(&_StakeNew.TransactOpts, user, amount, amountOutMin, stakeIndex, parent)
}

// StakeWithInviter is a paid mutator transaction binding the contract method 0xc82b6d1c.
//
// Solidity: function stakeWithInviter(address user, uint160 amount, uint256 amountOutMin, uint8 stakeIndex, address parent) returns()
func (_StakeNew *StakeNewTransactorSession) StakeWithInviter(user common.Address, amount *big.Int, amountOutMin *big.Int, stakeIndex uint8, parent common.Address) (*types.Transaction, error) {
	return _StakeNew.Contract.StakeWithInviter(&_StakeNew.TransactOpts, user, amount, amountOutMin, stakeIndex, parent)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakeNew *StakeNewTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakeNew *StakeNewSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakeNew.Contract.TransferOwnership(&_StakeNew.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakeNew *StakeNewTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakeNew.Contract.TransferOwnership(&_StakeNew.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakeNew *StakeNewTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakeNew *StakeNewSession) Unpause() (*types.Transaction, error) {
	return _StakeNew.Contract.Unpause(&_StakeNew.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakeNew *StakeNewTransactorSession) Unpause() (*types.Transaction, error) {
	return _StakeNew.Contract.Unpause(&_StakeNew.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address user, uint256 index) returns(uint256)
func (_StakeNew *StakeNewTransactor) Unstake(opts *bind.TransactOpts, user common.Address, index *big.Int) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "unstake", user, index)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address user, uint256 index) returns(uint256)
func (_StakeNew *StakeNewSession) Unstake(user common.Address, index *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.Unstake(&_StakeNew.TransactOpts, user, index)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address user, uint256 index) returns(uint256)
func (_StakeNew *StakeNewTransactorSession) Unstake(user common.Address, index *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.Unstake(&_StakeNew.TransactOpts, user, index)
}

// UpdateCircuitBreaker is a paid mutator transaction binding the contract method 0x53378e09.
//
// Solidity: function updateCircuitBreaker() returns()
func (_StakeNew *StakeNewTransactor) UpdateCircuitBreaker(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "updateCircuitBreaker")
}

// UpdateCircuitBreaker is a paid mutator transaction binding the contract method 0x53378e09.
//
// Solidity: function updateCircuitBreaker() returns()
func (_StakeNew *StakeNewSession) UpdateCircuitBreaker() (*types.Transaction, error) {
	return _StakeNew.Contract.UpdateCircuitBreaker(&_StakeNew.TransactOpts)
}

// UpdateCircuitBreaker is a paid mutator transaction binding the contract method 0x53378e09.
//
// Solidity: function updateCircuitBreaker() returns()
func (_StakeNew *StakeNewTransactorSession) UpdateCircuitBreaker() (*types.Transaction, error) {
	return _StakeNew.Contract.UpdateCircuitBreaker(&_StakeNew.TransactOpts)
}

// WithdrawAllFreeUsdt is a paid mutator transaction binding the contract method 0x50612b78.
//
// Solidity: function withdrawAllFreeUsdt() returns(uint256 amount)
func (_StakeNew *StakeNewTransactor) WithdrawAllFreeUsdt(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "withdrawAllFreeUsdt")
}

// WithdrawAllFreeUsdt is a paid mutator transaction binding the contract method 0x50612b78.
//
// Solidity: function withdrawAllFreeUsdt() returns(uint256 amount)
func (_StakeNew *StakeNewSession) WithdrawAllFreeUsdt() (*types.Transaction, error) {
	return _StakeNew.Contract.WithdrawAllFreeUsdt(&_StakeNew.TransactOpts)
}

// WithdrawAllFreeUsdt is a paid mutator transaction binding the contract method 0x50612b78.
//
// Solidity: function withdrawAllFreeUsdt() returns(uint256 amount)
func (_StakeNew *StakeNewTransactorSession) WithdrawAllFreeUsdt() (*types.Transaction, error) {
	return _StakeNew.Contract.WithdrawAllFreeUsdt(&_StakeNew.TransactOpts)
}

// WithdrawFreeUsdt is a paid mutator transaction binding the contract method 0x4aad6388.
//
// Solidity: function withdrawFreeUsdt(uint256 amount) returns(uint256)
func (_StakeNew *StakeNewTransactor) WithdrawFreeUsdt(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakeNew.contract.Transact(opts, "withdrawFreeUsdt", amount)
}

// WithdrawFreeUsdt is a paid mutator transaction binding the contract method 0x4aad6388.
//
// Solidity: function withdrawFreeUsdt(uint256 amount) returns(uint256)
func (_StakeNew *StakeNewSession) WithdrawFreeUsdt(amount *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.WithdrawFreeUsdt(&_StakeNew.TransactOpts, amount)
}

// WithdrawFreeUsdt is a paid mutator transaction binding the contract method 0x4aad6388.
//
// Solidity: function withdrawFreeUsdt(uint256 amount) returns(uint256)
func (_StakeNew *StakeNewTransactorSession) WithdrawFreeUsdt(amount *big.Int) (*types.Transaction, error) {
	return _StakeNew.Contract.WithdrawFreeUsdt(&_StakeNew.TransactOpts, amount)
}

// StakeNewCircuitBreakerParamsUpdatedIterator is returned from FilterCircuitBreakerParamsUpdated and is used to iterate over the raw logs and unpacked data for CircuitBreakerParamsUpdated events raised by the StakeNew contract.
type StakeNewCircuitBreakerParamsUpdatedIterator struct {
	Event *StakeNewCircuitBreakerParamsUpdated // Event containing the contract specifics and raw log

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
func (it *StakeNewCircuitBreakerParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewCircuitBreakerParamsUpdated)
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
		it.Event = new(StakeNewCircuitBreakerParamsUpdated)
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
func (it *StakeNewCircuitBreakerParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewCircuitBreakerParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewCircuitBreakerParamsUpdated represents a CircuitBreakerParamsUpdated event raised by the StakeNew contract.
type StakeNewCircuitBreakerParamsUpdated struct {
	DropBpsThreshold    *big.Int
	RecoveryDuration    *big.Int
	UnstakeRatePermille *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCircuitBreakerParamsUpdated is a free log retrieval operation binding the contract event 0xd31117f1181a295b1fcb8eef26accf397015787bbcd593a666e5f603943f71cd.
//
// Solidity: event CircuitBreakerParamsUpdated(uint256 dropBpsThreshold, uint40 recoveryDuration, uint256 unstakeRatePermille)
func (_StakeNew *StakeNewFilterer) FilterCircuitBreakerParamsUpdated(opts *bind.FilterOpts) (*StakeNewCircuitBreakerParamsUpdatedIterator, error) {

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "CircuitBreakerParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &StakeNewCircuitBreakerParamsUpdatedIterator{contract: _StakeNew.contract, event: "CircuitBreakerParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchCircuitBreakerParamsUpdated is a free log subscription operation binding the contract event 0xd31117f1181a295b1fcb8eef26accf397015787bbcd593a666e5f603943f71cd.
//
// Solidity: event CircuitBreakerParamsUpdated(uint256 dropBpsThreshold, uint40 recoveryDuration, uint256 unstakeRatePermille)
func (_StakeNew *StakeNewFilterer) WatchCircuitBreakerParamsUpdated(opts *bind.WatchOpts, sink chan<- *StakeNewCircuitBreakerParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "CircuitBreakerParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewCircuitBreakerParamsUpdated)
				if err := _StakeNew.contract.UnpackLog(event, "CircuitBreakerParamsUpdated", log); err != nil {
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

// ParseCircuitBreakerParamsUpdated is a log parse operation binding the contract event 0xd31117f1181a295b1fcb8eef26accf397015787bbcd593a666e5f603943f71cd.
//
// Solidity: event CircuitBreakerParamsUpdated(uint256 dropBpsThreshold, uint40 recoveryDuration, uint256 unstakeRatePermille)
func (_StakeNew *StakeNewFilterer) ParseCircuitBreakerParamsUpdated(log types.Log) (*StakeNewCircuitBreakerParamsUpdated, error) {
	event := new(StakeNewCircuitBreakerParamsUpdated)
	if err := _StakeNew.contract.UnpackLog(event, "CircuitBreakerParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewCircuitBreakerStateUpdatedIterator is returned from FilterCircuitBreakerStateUpdated and is used to iterate over the raw logs and unpacked data for CircuitBreakerStateUpdated events raised by the StakeNew contract.
type StakeNewCircuitBreakerStateUpdatedIterator struct {
	Event *StakeNewCircuitBreakerStateUpdated // Event containing the contract specifics and raw log

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
func (it *StakeNewCircuitBreakerStateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewCircuitBreakerStateUpdated)
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
		it.Event = new(StakeNewCircuitBreakerStateUpdated)
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
func (it *StakeNewCircuitBreakerStateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewCircuitBreakerStateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewCircuitBreakerStateUpdated represents a CircuitBreakerStateUpdated event raised by the StakeNew contract.
type StakeNewCircuitBreakerStateUpdated struct {
	CircuitBreakerTime *big.Int
	NewHighTime        *big.Int
	CurrentDropBps     *big.Int
	Action             uint8
	RoundCount         *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCircuitBreakerStateUpdated is a free log retrieval operation binding the contract event 0x5f796bb8756e027a2b866bace59149e48633b3e059078664ea6f1ccd3224f94f.
//
// Solidity: event CircuitBreakerStateUpdated(uint40 circuitBreakerTime, uint40 newHighTime, uint256 currentDropBps, uint8 action, uint256 roundCount)
func (_StakeNew *StakeNewFilterer) FilterCircuitBreakerStateUpdated(opts *bind.FilterOpts) (*StakeNewCircuitBreakerStateUpdatedIterator, error) {

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "CircuitBreakerStateUpdated")
	if err != nil {
		return nil, err
	}
	return &StakeNewCircuitBreakerStateUpdatedIterator{contract: _StakeNew.contract, event: "CircuitBreakerStateUpdated", logs: logs, sub: sub}, nil
}

// WatchCircuitBreakerStateUpdated is a free log subscription operation binding the contract event 0x5f796bb8756e027a2b866bace59149e48633b3e059078664ea6f1ccd3224f94f.
//
// Solidity: event CircuitBreakerStateUpdated(uint40 circuitBreakerTime, uint40 newHighTime, uint256 currentDropBps, uint8 action, uint256 roundCount)
func (_StakeNew *StakeNewFilterer) WatchCircuitBreakerStateUpdated(opts *bind.WatchOpts, sink chan<- *StakeNewCircuitBreakerStateUpdated) (event.Subscription, error) {

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "CircuitBreakerStateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewCircuitBreakerStateUpdated)
				if err := _StakeNew.contract.UnpackLog(event, "CircuitBreakerStateUpdated", log); err != nil {
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

// ParseCircuitBreakerStateUpdated is a log parse operation binding the contract event 0x5f796bb8756e027a2b866bace59149e48633b3e059078664ea6f1ccd3224f94f.
//
// Solidity: event CircuitBreakerStateUpdated(uint40 circuitBreakerTime, uint40 newHighTime, uint256 currentDropBps, uint8 action, uint256 roundCount)
func (_StakeNew *StakeNewFilterer) ParseCircuitBreakerStateUpdated(log types.Log) (*StakeNewCircuitBreakerStateUpdated, error) {
	event := new(StakeNewCircuitBreakerStateUpdated)
	if err := _StakeNew.contract.UnpackLog(event, "CircuitBreakerStateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewDailyLimitRatesUpdatedIterator is returned from FilterDailyLimitRatesUpdated and is used to iterate over the raw logs and unpacked data for DailyLimitRatesUpdated events raised by the StakeNew contract.
type StakeNewDailyLimitRatesUpdatedIterator struct {
	Event *StakeNewDailyLimitRatesUpdated // Event containing the contract specifics and raw log

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
func (it *StakeNewDailyLimitRatesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewDailyLimitRatesUpdated)
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
		it.Event = new(StakeNewDailyLimitRatesUpdated)
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
func (it *StakeNewDailyLimitRatesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewDailyLimitRatesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewDailyLimitRatesUpdated represents a DailyLimitRatesUpdated event raised by the StakeNew contract.
type StakeNewDailyLimitRatesUpdated struct {
	StakeLow    uint16
	StakeMid    uint16
	StakeHigh   uint16
	UnstakeLow  uint16
	UnstakeMid  uint16
	UnstakeHigh uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDailyLimitRatesUpdated is a free log retrieval operation binding the contract event 0x4f1604a0ed530fb8afbf5ecc0c927ac8d1e90078b469b6412cd0335e842fd993.
//
// Solidity: event DailyLimitRatesUpdated(uint16 stakeLow, uint16 stakeMid, uint16 stakeHigh, uint16 unstakeLow, uint16 unstakeMid, uint16 unstakeHigh)
func (_StakeNew *StakeNewFilterer) FilterDailyLimitRatesUpdated(opts *bind.FilterOpts) (*StakeNewDailyLimitRatesUpdatedIterator, error) {

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "DailyLimitRatesUpdated")
	if err != nil {
		return nil, err
	}
	return &StakeNewDailyLimitRatesUpdatedIterator{contract: _StakeNew.contract, event: "DailyLimitRatesUpdated", logs: logs, sub: sub}, nil
}

// WatchDailyLimitRatesUpdated is a free log subscription operation binding the contract event 0x4f1604a0ed530fb8afbf5ecc0c927ac8d1e90078b469b6412cd0335e842fd993.
//
// Solidity: event DailyLimitRatesUpdated(uint16 stakeLow, uint16 stakeMid, uint16 stakeHigh, uint16 unstakeLow, uint16 unstakeMid, uint16 unstakeHigh)
func (_StakeNew *StakeNewFilterer) WatchDailyLimitRatesUpdated(opts *bind.WatchOpts, sink chan<- *StakeNewDailyLimitRatesUpdated) (event.Subscription, error) {

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "DailyLimitRatesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewDailyLimitRatesUpdated)
				if err := _StakeNew.contract.UnpackLog(event, "DailyLimitRatesUpdated", log); err != nil {
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

// ParseDailyLimitRatesUpdated is a log parse operation binding the contract event 0x4f1604a0ed530fb8afbf5ecc0c927ac8d1e90078b469b6412cd0335e842fd993.
//
// Solidity: event DailyLimitRatesUpdated(uint16 stakeLow, uint16 stakeMid, uint16 stakeHigh, uint16 unstakeLow, uint16 unstakeMid, uint16 unstakeHigh)
func (_StakeNew *StakeNewFilterer) ParseDailyLimitRatesUpdated(log types.Log) (*StakeNewDailyLimitRatesUpdated, error) {
	event := new(StakeNewDailyLimitRatesUpdated)
	if err := _StakeNew.contract.UnpackLog(event, "DailyLimitRatesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewDailyQueueCancelLimitUpdatedIterator is returned from FilterDailyQueueCancelLimitUpdated and is used to iterate over the raw logs and unpacked data for DailyQueueCancelLimitUpdated events raised by the StakeNew contract.
type StakeNewDailyQueueCancelLimitUpdatedIterator struct {
	Event *StakeNewDailyQueueCancelLimitUpdated // Event containing the contract specifics and raw log

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
func (it *StakeNewDailyQueueCancelLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewDailyQueueCancelLimitUpdated)
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
		it.Event = new(StakeNewDailyQueueCancelLimitUpdated)
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
func (it *StakeNewDailyQueueCancelLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewDailyQueueCancelLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewDailyQueueCancelLimitUpdated represents a DailyQueueCancelLimitUpdated event raised by the StakeNew contract.
type StakeNewDailyQueueCancelLimitUpdated struct {
	Limit *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDailyQueueCancelLimitUpdated is a free log retrieval operation binding the contract event 0x5cf97520ddba1c430c95f9202e2b857492e2dae80589051982dc79536b05afea.
//
// Solidity: event DailyQueueCancelLimitUpdated(uint256 limit)
func (_StakeNew *StakeNewFilterer) FilterDailyQueueCancelLimitUpdated(opts *bind.FilterOpts) (*StakeNewDailyQueueCancelLimitUpdatedIterator, error) {

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "DailyQueueCancelLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &StakeNewDailyQueueCancelLimitUpdatedIterator{contract: _StakeNew.contract, event: "DailyQueueCancelLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchDailyQueueCancelLimitUpdated is a free log subscription operation binding the contract event 0x5cf97520ddba1c430c95f9202e2b857492e2dae80589051982dc79536b05afea.
//
// Solidity: event DailyQueueCancelLimitUpdated(uint256 limit)
func (_StakeNew *StakeNewFilterer) WatchDailyQueueCancelLimitUpdated(opts *bind.WatchOpts, sink chan<- *StakeNewDailyQueueCancelLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "DailyQueueCancelLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewDailyQueueCancelLimitUpdated)
				if err := _StakeNew.contract.UnpackLog(event, "DailyQueueCancelLimitUpdated", log); err != nil {
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

// ParseDailyQueueCancelLimitUpdated is a log parse operation binding the contract event 0x5cf97520ddba1c430c95f9202e2b857492e2dae80589051982dc79536b05afea.
//
// Solidity: event DailyQueueCancelLimitUpdated(uint256 limit)
func (_StakeNew *StakeNewFilterer) ParseDailyQueueCancelLimitUpdated(log types.Log) (*StakeNewDailyQueueCancelLimitUpdated, error) {
	event := new(StakeNewDailyQueueCancelLimitUpdated)
	if err := _StakeNew.contract.UnpackLog(event, "DailyQueueCancelLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewDividendSystemUpdatedIterator is returned from FilterDividendSystemUpdated and is used to iterate over the raw logs and unpacked data for DividendSystemUpdated events raised by the StakeNew contract.
type StakeNewDividendSystemUpdatedIterator struct {
	Event *StakeNewDividendSystemUpdated // Event containing the contract specifics and raw log

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
func (it *StakeNewDividendSystemUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewDividendSystemUpdated)
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
		it.Event = new(StakeNewDividendSystemUpdated)
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
func (it *StakeNewDividendSystemUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewDividendSystemUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewDividendSystemUpdated represents a DividendSystemUpdated event raised by the StakeNew contract.
type StakeNewDividendSystemUpdated struct {
	DividendSystem common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDividendSystemUpdated is a free log retrieval operation binding the contract event 0xb7fa8611c311f61b6eb236ee9ebe0e2c6c4d175233290fa1ad63dd6e265469f0.
//
// Solidity: event DividendSystemUpdated(address indexed dividendSystem)
func (_StakeNew *StakeNewFilterer) FilterDividendSystemUpdated(opts *bind.FilterOpts, dividendSystem []common.Address) (*StakeNewDividendSystemUpdatedIterator, error) {

	var dividendSystemRule []interface{}
	for _, dividendSystemItem := range dividendSystem {
		dividendSystemRule = append(dividendSystemRule, dividendSystemItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "DividendSystemUpdated", dividendSystemRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewDividendSystemUpdatedIterator{contract: _StakeNew.contract, event: "DividendSystemUpdated", logs: logs, sub: sub}, nil
}

// WatchDividendSystemUpdated is a free log subscription operation binding the contract event 0xb7fa8611c311f61b6eb236ee9ebe0e2c6c4d175233290fa1ad63dd6e265469f0.
//
// Solidity: event DividendSystemUpdated(address indexed dividendSystem)
func (_StakeNew *StakeNewFilterer) WatchDividendSystemUpdated(opts *bind.WatchOpts, sink chan<- *StakeNewDividendSystemUpdated, dividendSystem []common.Address) (event.Subscription, error) {

	var dividendSystemRule []interface{}
	for _, dividendSystemItem := range dividendSystem {
		dividendSystemRule = append(dividendSystemRule, dividendSystemItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "DividendSystemUpdated", dividendSystemRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewDividendSystemUpdated)
				if err := _StakeNew.contract.UnpackLog(event, "DividendSystemUpdated", log); err != nil {
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

// ParseDividendSystemUpdated is a log parse operation binding the contract event 0xb7fa8611c311f61b6eb236ee9ebe0e2c6c4d175233290fa1ad63dd6e265469f0.
//
// Solidity: event DividendSystemUpdated(address indexed dividendSystem)
func (_StakeNew *StakeNewFilterer) ParseDividendSystemUpdated(log types.Log) (*StakeNewDividendSystemUpdated, error) {
	event := new(StakeNewDividendSystemUpdated)
	if err := _StakeNew.contract.UnpackLog(event, "DividendSystemUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewFreeUsdtCollectorUpdatedIterator is returned from FilterFreeUsdtCollectorUpdated and is used to iterate over the raw logs and unpacked data for FreeUsdtCollectorUpdated events raised by the StakeNew contract.
type StakeNewFreeUsdtCollectorUpdatedIterator struct {
	Event *StakeNewFreeUsdtCollectorUpdated // Event containing the contract specifics and raw log

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
func (it *StakeNewFreeUsdtCollectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewFreeUsdtCollectorUpdated)
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
		it.Event = new(StakeNewFreeUsdtCollectorUpdated)
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
func (it *StakeNewFreeUsdtCollectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewFreeUsdtCollectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewFreeUsdtCollectorUpdated represents a FreeUsdtCollectorUpdated event raised by the StakeNew contract.
type StakeNewFreeUsdtCollectorUpdated struct {
	PreviousCollector common.Address
	NewCollector      common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFreeUsdtCollectorUpdated is a free log retrieval operation binding the contract event 0xd9d498bdb44b545eb075adbab7c0db98b171ef7191ba9f4e00543c700ef25189.
//
// Solidity: event FreeUsdtCollectorUpdated(address indexed previousCollector, address indexed newCollector)
func (_StakeNew *StakeNewFilterer) FilterFreeUsdtCollectorUpdated(opts *bind.FilterOpts, previousCollector []common.Address, newCollector []common.Address) (*StakeNewFreeUsdtCollectorUpdatedIterator, error) {

	var previousCollectorRule []interface{}
	for _, previousCollectorItem := range previousCollector {
		previousCollectorRule = append(previousCollectorRule, previousCollectorItem)
	}
	var newCollectorRule []interface{}
	for _, newCollectorItem := range newCollector {
		newCollectorRule = append(newCollectorRule, newCollectorItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "FreeUsdtCollectorUpdated", previousCollectorRule, newCollectorRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewFreeUsdtCollectorUpdatedIterator{contract: _StakeNew.contract, event: "FreeUsdtCollectorUpdated", logs: logs, sub: sub}, nil
}

// WatchFreeUsdtCollectorUpdated is a free log subscription operation binding the contract event 0xd9d498bdb44b545eb075adbab7c0db98b171ef7191ba9f4e00543c700ef25189.
//
// Solidity: event FreeUsdtCollectorUpdated(address indexed previousCollector, address indexed newCollector)
func (_StakeNew *StakeNewFilterer) WatchFreeUsdtCollectorUpdated(opts *bind.WatchOpts, sink chan<- *StakeNewFreeUsdtCollectorUpdated, previousCollector []common.Address, newCollector []common.Address) (event.Subscription, error) {

	var previousCollectorRule []interface{}
	for _, previousCollectorItem := range previousCollector {
		previousCollectorRule = append(previousCollectorRule, previousCollectorItem)
	}
	var newCollectorRule []interface{}
	for _, newCollectorItem := range newCollector {
		newCollectorRule = append(newCollectorRule, newCollectorItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "FreeUsdtCollectorUpdated", previousCollectorRule, newCollectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewFreeUsdtCollectorUpdated)
				if err := _StakeNew.contract.UnpackLog(event, "FreeUsdtCollectorUpdated", log); err != nil {
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

// ParseFreeUsdtCollectorUpdated is a log parse operation binding the contract event 0xd9d498bdb44b545eb075adbab7c0db98b171ef7191ba9f4e00543c700ef25189.
//
// Solidity: event FreeUsdtCollectorUpdated(address indexed previousCollector, address indexed newCollector)
func (_StakeNew *StakeNewFilterer) ParseFreeUsdtCollectorUpdated(log types.Log) (*StakeNewFreeUsdtCollectorUpdated, error) {
	event := new(StakeNewFreeUsdtCollectorUpdated)
	if err := _StakeNew.contract.UnpackLog(event, "FreeUsdtCollectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewFreeUsdtWithdrawnIterator is returned from FilterFreeUsdtWithdrawn and is used to iterate over the raw logs and unpacked data for FreeUsdtWithdrawn events raised by the StakeNew contract.
type StakeNewFreeUsdtWithdrawnIterator struct {
	Event *StakeNewFreeUsdtWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakeNewFreeUsdtWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewFreeUsdtWithdrawn)
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
		it.Event = new(StakeNewFreeUsdtWithdrawn)
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
func (it *StakeNewFreeUsdtWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewFreeUsdtWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewFreeUsdtWithdrawn represents a FreeUsdtWithdrawn event raised by the StakeNew contract.
type StakeNewFreeUsdtWithdrawn struct {
	Collector common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFreeUsdtWithdrawn is a free log retrieval operation binding the contract event 0x47a313a951ef6b1f8b61ebcfef31fba7fe7f3f606ebf5a8d481ccb852cebdf5a.
//
// Solidity: event FreeUsdtWithdrawn(address indexed collector, uint256 amount)
func (_StakeNew *StakeNewFilterer) FilterFreeUsdtWithdrawn(opts *bind.FilterOpts, collector []common.Address) (*StakeNewFreeUsdtWithdrawnIterator, error) {

	var collectorRule []interface{}
	for _, collectorItem := range collector {
		collectorRule = append(collectorRule, collectorItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "FreeUsdtWithdrawn", collectorRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewFreeUsdtWithdrawnIterator{contract: _StakeNew.contract, event: "FreeUsdtWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFreeUsdtWithdrawn is a free log subscription operation binding the contract event 0x47a313a951ef6b1f8b61ebcfef31fba7fe7f3f606ebf5a8d481ccb852cebdf5a.
//
// Solidity: event FreeUsdtWithdrawn(address indexed collector, uint256 amount)
func (_StakeNew *StakeNewFilterer) WatchFreeUsdtWithdrawn(opts *bind.WatchOpts, sink chan<- *StakeNewFreeUsdtWithdrawn, collector []common.Address) (event.Subscription, error) {

	var collectorRule []interface{}
	for _, collectorItem := range collector {
		collectorRule = append(collectorRule, collectorItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "FreeUsdtWithdrawn", collectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewFreeUsdtWithdrawn)
				if err := _StakeNew.contract.UnpackLog(event, "FreeUsdtWithdrawn", log); err != nil {
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

// ParseFreeUsdtWithdrawn is a log parse operation binding the contract event 0x47a313a951ef6b1f8b61ebcfef31fba7fe7f3f606ebf5a8d481ccb852cebdf5a.
//
// Solidity: event FreeUsdtWithdrawn(address indexed collector, uint256 amount)
func (_StakeNew *StakeNewFilterer) ParseFreeUsdtWithdrawn(log types.Log) (*StakeNewFreeUsdtWithdrawn, error) {
	event := new(StakeNewFreeUsdtWithdrawn)
	if err := _StakeNew.contract.UnpackLog(event, "FreeUsdtWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewGatewaySetIterator is returned from FilterGatewaySet and is used to iterate over the raw logs and unpacked data for GatewaySet events raised by the StakeNew contract.
type StakeNewGatewaySetIterator struct {
	Event *StakeNewGatewaySet // Event containing the contract specifics and raw log

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
func (it *StakeNewGatewaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewGatewaySet)
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
		it.Event = new(StakeNewGatewaySet)
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
func (it *StakeNewGatewaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewGatewaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewGatewaySet represents a GatewaySet event raised by the StakeNew contract.
type StakeNewGatewaySet struct {
	Gateway common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGatewaySet is a free log retrieval operation binding the contract event 0x5317fa585931182194fed99f2ea5f2efd38af9cff9724273704c8501c521e34b.
//
// Solidity: event GatewaySet(address indexed gateway)
func (_StakeNew *StakeNewFilterer) FilterGatewaySet(opts *bind.FilterOpts, gateway []common.Address) (*StakeNewGatewaySetIterator, error) {

	var gatewayRule []interface{}
	for _, gatewayItem := range gateway {
		gatewayRule = append(gatewayRule, gatewayItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "GatewaySet", gatewayRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewGatewaySetIterator{contract: _StakeNew.contract, event: "GatewaySet", logs: logs, sub: sub}, nil
}

// WatchGatewaySet is a free log subscription operation binding the contract event 0x5317fa585931182194fed99f2ea5f2efd38af9cff9724273704c8501c521e34b.
//
// Solidity: event GatewaySet(address indexed gateway)
func (_StakeNew *StakeNewFilterer) WatchGatewaySet(opts *bind.WatchOpts, sink chan<- *StakeNewGatewaySet, gateway []common.Address) (event.Subscription, error) {

	var gatewayRule []interface{}
	for _, gatewayItem := range gateway {
		gatewayRule = append(gatewayRule, gatewayItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "GatewaySet", gatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewGatewaySet)
				if err := _StakeNew.contract.UnpackLog(event, "GatewaySet", log); err != nil {
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

// ParseGatewaySet is a log parse operation binding the contract event 0x5317fa585931182194fed99f2ea5f2efd38af9cff9724273704c8501c521e34b.
//
// Solidity: event GatewaySet(address indexed gateway)
func (_StakeNew *StakeNewFilterer) ParseGatewaySet(log types.Log) (*StakeNewGatewaySet, error) {
	event := new(StakeNewGatewaySet)
	if err := _StakeNew.contract.UnpackLog(event, "GatewaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewInitAccountTransferredIterator is returned from FilterInitAccountTransferred and is used to iterate over the raw logs and unpacked data for InitAccountTransferred events raised by the StakeNew contract.
type StakeNewInitAccountTransferredIterator struct {
	Event *StakeNewInitAccountTransferred // Event containing the contract specifics and raw log

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
func (it *StakeNewInitAccountTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewInitAccountTransferred)
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
		it.Event = new(StakeNewInitAccountTransferred)
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
func (it *StakeNewInitAccountTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewInitAccountTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewInitAccountTransferred represents a InitAccountTransferred event raised by the StakeNew contract.
type StakeNewInitAccountTransferred struct {
	PreviousInitAccount common.Address
	NewInitAccount      common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterInitAccountTransferred is a free log retrieval operation binding the contract event 0xdbdef5214b96ccc548368e20ee0423776be2da50a3db91a024c5b43802150380.
//
// Solidity: event InitAccountTransferred(address indexed previousInitAccount, address indexed newInitAccount)
func (_StakeNew *StakeNewFilterer) FilterInitAccountTransferred(opts *bind.FilterOpts, previousInitAccount []common.Address, newInitAccount []common.Address) (*StakeNewInitAccountTransferredIterator, error) {

	var previousInitAccountRule []interface{}
	for _, previousInitAccountItem := range previousInitAccount {
		previousInitAccountRule = append(previousInitAccountRule, previousInitAccountItem)
	}
	var newInitAccountRule []interface{}
	for _, newInitAccountItem := range newInitAccount {
		newInitAccountRule = append(newInitAccountRule, newInitAccountItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "InitAccountTransferred", previousInitAccountRule, newInitAccountRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewInitAccountTransferredIterator{contract: _StakeNew.contract, event: "InitAccountTransferred", logs: logs, sub: sub}, nil
}

// WatchInitAccountTransferred is a free log subscription operation binding the contract event 0xdbdef5214b96ccc548368e20ee0423776be2da50a3db91a024c5b43802150380.
//
// Solidity: event InitAccountTransferred(address indexed previousInitAccount, address indexed newInitAccount)
func (_StakeNew *StakeNewFilterer) WatchInitAccountTransferred(opts *bind.WatchOpts, sink chan<- *StakeNewInitAccountTransferred, previousInitAccount []common.Address, newInitAccount []common.Address) (event.Subscription, error) {

	var previousInitAccountRule []interface{}
	for _, previousInitAccountItem := range previousInitAccount {
		previousInitAccountRule = append(previousInitAccountRule, previousInitAccountItem)
	}
	var newInitAccountRule []interface{}
	for _, newInitAccountItem := range newInitAccount {
		newInitAccountRule = append(newInitAccountRule, newInitAccountItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "InitAccountTransferred", previousInitAccountRule, newInitAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewInitAccountTransferred)
				if err := _StakeNew.contract.UnpackLog(event, "InitAccountTransferred", log); err != nil {
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

// ParseInitAccountTransferred is a log parse operation binding the contract event 0xdbdef5214b96ccc548368e20ee0423776be2da50a3db91a024c5b43802150380.
//
// Solidity: event InitAccountTransferred(address indexed previousInitAccount, address indexed newInitAccount)
func (_StakeNew *StakeNewFilterer) ParseInitAccountTransferred(log types.Log) (*StakeNewInitAccountTransferred, error) {
	event := new(StakeNewInitAccountTransferred)
	if err := _StakeNew.contract.UnpackLog(event, "InitAccountTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewInitSystemsUpdatedIterator is returned from FilterInitSystemsUpdated and is used to iterate over the raw logs and unpacked data for InitSystemsUpdated events raised by the StakeNew contract.
type StakeNewInitSystemsUpdatedIterator struct {
	Event *StakeNewInitSystemsUpdated // Event containing the contract specifics and raw log

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
func (it *StakeNewInitSystemsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewInitSystemsUpdated)
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
		it.Event = new(StakeNewInitSystemsUpdated)
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
func (it *StakeNewInitSystemsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewInitSystemsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewInitSystemsUpdated represents a InitSystemsUpdated event raised by the StakeNew contract.
type StakeNewInitSystemsUpdated struct {
	Token          common.Address
	UserSystem     common.Address
	DividendSystem common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitSystemsUpdated is a free log retrieval operation binding the contract event 0x28e97779985d9ea743e9760a0bbac8aa8fee829fa807ca850617a5597ccd54a7.
//
// Solidity: event InitSystemsUpdated(address indexed token, address indexed userSystem, address indexed dividendSystem)
func (_StakeNew *StakeNewFilterer) FilterInitSystemsUpdated(opts *bind.FilterOpts, token []common.Address, userSystem []common.Address, dividendSystem []common.Address) (*StakeNewInitSystemsUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userSystemRule []interface{}
	for _, userSystemItem := range userSystem {
		userSystemRule = append(userSystemRule, userSystemItem)
	}
	var dividendSystemRule []interface{}
	for _, dividendSystemItem := range dividendSystem {
		dividendSystemRule = append(dividendSystemRule, dividendSystemItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "InitSystemsUpdated", tokenRule, userSystemRule, dividendSystemRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewInitSystemsUpdatedIterator{contract: _StakeNew.contract, event: "InitSystemsUpdated", logs: logs, sub: sub}, nil
}

// WatchInitSystemsUpdated is a free log subscription operation binding the contract event 0x28e97779985d9ea743e9760a0bbac8aa8fee829fa807ca850617a5597ccd54a7.
//
// Solidity: event InitSystemsUpdated(address indexed token, address indexed userSystem, address indexed dividendSystem)
func (_StakeNew *StakeNewFilterer) WatchInitSystemsUpdated(opts *bind.WatchOpts, sink chan<- *StakeNewInitSystemsUpdated, token []common.Address, userSystem []common.Address, dividendSystem []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var userSystemRule []interface{}
	for _, userSystemItem := range userSystem {
		userSystemRule = append(userSystemRule, userSystemItem)
	}
	var dividendSystemRule []interface{}
	for _, dividendSystemItem := range dividendSystem {
		dividendSystemRule = append(dividendSystemRule, dividendSystemItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "InitSystemsUpdated", tokenRule, userSystemRule, dividendSystemRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewInitSystemsUpdated)
				if err := _StakeNew.contract.UnpackLog(event, "InitSystemsUpdated", log); err != nil {
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

// ParseInitSystemsUpdated is a log parse operation binding the contract event 0x28e97779985d9ea743e9760a0bbac8aa8fee829fa807ca850617a5597ccd54a7.
//
// Solidity: event InitSystemsUpdated(address indexed token, address indexed userSystem, address indexed dividendSystem)
func (_StakeNew *StakeNewFilterer) ParseInitSystemsUpdated(log types.Log) (*StakeNewInitSystemsUpdated, error) {
	event := new(StakeNewInitSystemsUpdated)
	if err := _StakeNew.contract.UnpackLog(event, "InitSystemsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewMarketingAddressUpdatedIterator is returned from FilterMarketingAddressUpdated and is used to iterate over the raw logs and unpacked data for MarketingAddressUpdated events raised by the StakeNew contract.
type StakeNewMarketingAddressUpdatedIterator struct {
	Event *StakeNewMarketingAddressUpdated // Event containing the contract specifics and raw log

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
func (it *StakeNewMarketingAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewMarketingAddressUpdated)
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
		it.Event = new(StakeNewMarketingAddressUpdated)
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
func (it *StakeNewMarketingAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewMarketingAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewMarketingAddressUpdated represents a MarketingAddressUpdated event raised by the StakeNew contract.
type StakeNewMarketingAddressUpdated struct {
	MarketingAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMarketingAddressUpdated is a free log retrieval operation binding the contract event 0x0c5dca4975009b82f989b2f70aef27e88d40a772c4661bf9c76767eebdd4ec75.
//
// Solidity: event MarketingAddressUpdated(address indexed marketingAddress)
func (_StakeNew *StakeNewFilterer) FilterMarketingAddressUpdated(opts *bind.FilterOpts, marketingAddress []common.Address) (*StakeNewMarketingAddressUpdatedIterator, error) {

	var marketingAddressRule []interface{}
	for _, marketingAddressItem := range marketingAddress {
		marketingAddressRule = append(marketingAddressRule, marketingAddressItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "MarketingAddressUpdated", marketingAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewMarketingAddressUpdatedIterator{contract: _StakeNew.contract, event: "MarketingAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchMarketingAddressUpdated is a free log subscription operation binding the contract event 0x0c5dca4975009b82f989b2f70aef27e88d40a772c4661bf9c76767eebdd4ec75.
//
// Solidity: event MarketingAddressUpdated(address indexed marketingAddress)
func (_StakeNew *StakeNewFilterer) WatchMarketingAddressUpdated(opts *bind.WatchOpts, sink chan<- *StakeNewMarketingAddressUpdated, marketingAddress []common.Address) (event.Subscription, error) {

	var marketingAddressRule []interface{}
	for _, marketingAddressItem := range marketingAddress {
		marketingAddressRule = append(marketingAddressRule, marketingAddressItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "MarketingAddressUpdated", marketingAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewMarketingAddressUpdated)
				if err := _StakeNew.contract.UnpackLog(event, "MarketingAddressUpdated", log); err != nil {
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

// ParseMarketingAddressUpdated is a log parse operation binding the contract event 0x0c5dca4975009b82f989b2f70aef27e88d40a772c4661bf9c76767eebdd4ec75.
//
// Solidity: event MarketingAddressUpdated(address indexed marketingAddress)
func (_StakeNew *StakeNewFilterer) ParseMarketingAddressUpdated(log types.Log) (*StakeNewMarketingAddressUpdated, error) {
	event := new(StakeNewMarketingAddressUpdated)
	if err := _StakeNew.contract.UnpackLog(event, "MarketingAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewMaxReferralDepthUpdatedIterator is returned from FilterMaxReferralDepthUpdated and is used to iterate over the raw logs and unpacked data for MaxReferralDepthUpdated events raised by the StakeNew contract.
type StakeNewMaxReferralDepthUpdatedIterator struct {
	Event *StakeNewMaxReferralDepthUpdated // Event containing the contract specifics and raw log

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
func (it *StakeNewMaxReferralDepthUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewMaxReferralDepthUpdated)
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
		it.Event = new(StakeNewMaxReferralDepthUpdated)
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
func (it *StakeNewMaxReferralDepthUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewMaxReferralDepthUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewMaxReferralDepthUpdated represents a MaxReferralDepthUpdated event raised by the StakeNew contract.
type StakeNewMaxReferralDepthUpdated struct {
	Depth *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMaxReferralDepthUpdated is a free log retrieval operation binding the contract event 0x8d1c470a9259851b52b07a09f8a63d34282689fab0ad39115523957e2866d09e.
//
// Solidity: event MaxReferralDepthUpdated(uint256 depth)
func (_StakeNew *StakeNewFilterer) FilterMaxReferralDepthUpdated(opts *bind.FilterOpts) (*StakeNewMaxReferralDepthUpdatedIterator, error) {

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "MaxReferralDepthUpdated")
	if err != nil {
		return nil, err
	}
	return &StakeNewMaxReferralDepthUpdatedIterator{contract: _StakeNew.contract, event: "MaxReferralDepthUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxReferralDepthUpdated is a free log subscription operation binding the contract event 0x8d1c470a9259851b52b07a09f8a63d34282689fab0ad39115523957e2866d09e.
//
// Solidity: event MaxReferralDepthUpdated(uint256 depth)
func (_StakeNew *StakeNewFilterer) WatchMaxReferralDepthUpdated(opts *bind.WatchOpts, sink chan<- *StakeNewMaxReferralDepthUpdated) (event.Subscription, error) {

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "MaxReferralDepthUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewMaxReferralDepthUpdated)
				if err := _StakeNew.contract.UnpackLog(event, "MaxReferralDepthUpdated", log); err != nil {
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

// ParseMaxReferralDepthUpdated is a log parse operation binding the contract event 0x8d1c470a9259851b52b07a09f8a63d34282689fab0ad39115523957e2866d09e.
//
// Solidity: event MaxReferralDepthUpdated(uint256 depth)
func (_StakeNew *StakeNewFilterer) ParseMaxReferralDepthUpdated(log types.Log) (*StakeNewMaxReferralDepthUpdated, error) {
	event := new(StakeNewMaxReferralDepthUpdated)
	if err := _StakeNew.contract.UnpackLog(event, "MaxReferralDepthUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakeNew contract.
type StakeNewOwnershipTransferredIterator struct {
	Event *StakeNewOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakeNewOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewOwnershipTransferred)
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
		it.Event = new(StakeNewOwnershipTransferred)
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
func (it *StakeNewOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewOwnershipTransferred represents a OwnershipTransferred event raised by the StakeNew contract.
type StakeNewOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakeNew *StakeNewFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakeNewOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewOwnershipTransferredIterator{contract: _StakeNew.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakeNew *StakeNewFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakeNewOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewOwnershipTransferred)
				if err := _StakeNew.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakeNew *StakeNewFilterer) ParseOwnershipTransferred(log types.Log) (*StakeNewOwnershipTransferred, error) {
	event := new(StakeNewOwnershipTransferred)
	if err := _StakeNew.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StakeNew contract.
type StakeNewPausedIterator struct {
	Event *StakeNewPaused // Event containing the contract specifics and raw log

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
func (it *StakeNewPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewPaused)
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
		it.Event = new(StakeNewPaused)
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
func (it *StakeNewPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewPaused represents a Paused event raised by the StakeNew contract.
type StakeNewPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed account)
func (_StakeNew *StakeNewFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*StakeNewPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewPausedIterator{contract: _StakeNew.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed account)
func (_StakeNew *StakeNewFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakeNewPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewPaused)
				if err := _StakeNew.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address indexed account)
func (_StakeNew *StakeNewFilterer) ParsePaused(log types.Log) (*StakeNewPaused, error) {
	event := new(StakeNewPaused)
	if err := _StakeNew.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewQueueAddedIterator is returned from FilterQueueAdded and is used to iterate over the raw logs and unpacked data for QueueAdded events raised by the StakeNew contract.
type StakeNewQueueAddedIterator struct {
	Event *StakeNewQueueAdded // Event containing the contract specifics and raw log

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
func (it *StakeNewQueueAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewQueueAdded)
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
		it.Event = new(StakeNewQueueAdded)
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
func (it *StakeNewQueueAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewQueueAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewQueueAdded represents a QueueAdded event raised by the StakeNew contract.
type StakeNewQueueAdded struct {
	QueueIndex *big.Int
	User       common.Address
	Amount     *big.Int
	StakeIndex uint8
	QueuedAt   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQueueAdded is a free log retrieval operation binding the contract event 0x6903c90bdb72983604ebe83fad97483fa1aa92db6557b0712d8ed75cafd7d9db.
//
// Solidity: event QueueAdded(uint256 indexed queueIndex, address indexed user, uint160 amount, uint8 stakeIndex, uint40 queuedAt)
func (_StakeNew *StakeNewFilterer) FilterQueueAdded(opts *bind.FilterOpts, queueIndex []*big.Int, user []common.Address) (*StakeNewQueueAddedIterator, error) {

	var queueIndexRule []interface{}
	for _, queueIndexItem := range queueIndex {
		queueIndexRule = append(queueIndexRule, queueIndexItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "QueueAdded", queueIndexRule, userRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewQueueAddedIterator{contract: _StakeNew.contract, event: "QueueAdded", logs: logs, sub: sub}, nil
}

// WatchQueueAdded is a free log subscription operation binding the contract event 0x6903c90bdb72983604ebe83fad97483fa1aa92db6557b0712d8ed75cafd7d9db.
//
// Solidity: event QueueAdded(uint256 indexed queueIndex, address indexed user, uint160 amount, uint8 stakeIndex, uint40 queuedAt)
func (_StakeNew *StakeNewFilterer) WatchQueueAdded(opts *bind.WatchOpts, sink chan<- *StakeNewQueueAdded, queueIndex []*big.Int, user []common.Address) (event.Subscription, error) {

	var queueIndexRule []interface{}
	for _, queueIndexItem := range queueIndex {
		queueIndexRule = append(queueIndexRule, queueIndexItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "QueueAdded", queueIndexRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewQueueAdded)
				if err := _StakeNew.contract.UnpackLog(event, "QueueAdded", log); err != nil {
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

// ParseQueueAdded is a log parse operation binding the contract event 0x6903c90bdb72983604ebe83fad97483fa1aa92db6557b0712d8ed75cafd7d9db.
//
// Solidity: event QueueAdded(uint256 indexed queueIndex, address indexed user, uint160 amount, uint8 stakeIndex, uint40 queuedAt)
func (_StakeNew *StakeNewFilterer) ParseQueueAdded(log types.Log) (*StakeNewQueueAdded, error) {
	event := new(StakeNewQueueAdded)
	if err := _StakeNew.contract.UnpackLog(event, "QueueAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewQueueCancelledIterator is returned from FilterQueueCancelled and is used to iterate over the raw logs and unpacked data for QueueCancelled events raised by the StakeNew contract.
type StakeNewQueueCancelledIterator struct {
	Event *StakeNewQueueCancelled // Event containing the contract specifics and raw log

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
func (it *StakeNewQueueCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewQueueCancelled)
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
		it.Event = new(StakeNewQueueCancelled)
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
func (it *StakeNewQueueCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewQueueCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewQueueCancelled represents a QueueCancelled event raised by the StakeNew contract.
type StakeNewQueueCancelled struct {
	QueueIndex *big.Int
	User       common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQueueCancelled is a free log retrieval operation binding the contract event 0xa388301c2cd248e005b812868e13c479f50a4cde783815b3d6d42a917209816c.
//
// Solidity: event QueueCancelled(uint256 indexed queueIndex, address indexed user, uint160 amount)
func (_StakeNew *StakeNewFilterer) FilterQueueCancelled(opts *bind.FilterOpts, queueIndex []*big.Int, user []common.Address) (*StakeNewQueueCancelledIterator, error) {

	var queueIndexRule []interface{}
	for _, queueIndexItem := range queueIndex {
		queueIndexRule = append(queueIndexRule, queueIndexItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "QueueCancelled", queueIndexRule, userRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewQueueCancelledIterator{contract: _StakeNew.contract, event: "QueueCancelled", logs: logs, sub: sub}, nil
}

// WatchQueueCancelled is a free log subscription operation binding the contract event 0xa388301c2cd248e005b812868e13c479f50a4cde783815b3d6d42a917209816c.
//
// Solidity: event QueueCancelled(uint256 indexed queueIndex, address indexed user, uint160 amount)
func (_StakeNew *StakeNewFilterer) WatchQueueCancelled(opts *bind.WatchOpts, sink chan<- *StakeNewQueueCancelled, queueIndex []*big.Int, user []common.Address) (event.Subscription, error) {

	var queueIndexRule []interface{}
	for _, queueIndexItem := range queueIndex {
		queueIndexRule = append(queueIndexRule, queueIndexItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "QueueCancelled", queueIndexRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewQueueCancelled)
				if err := _StakeNew.contract.UnpackLog(event, "QueueCancelled", log); err != nil {
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

// ParseQueueCancelled is a log parse operation binding the contract event 0xa388301c2cd248e005b812868e13c479f50a4cde783815b3d6d42a917209816c.
//
// Solidity: event QueueCancelled(uint256 indexed queueIndex, address indexed user, uint160 amount)
func (_StakeNew *StakeNewFilterer) ParseQueueCancelled(log types.Log) (*StakeNewQueueCancelled, error) {
	event := new(StakeNewQueueCancelled)
	if err := _StakeNew.contract.UnpackLog(event, "QueueCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewQueueProcessedIterator is returned from FilterQueueProcessed and is used to iterate over the raw logs and unpacked data for QueueProcessed events raised by the StakeNew contract.
type StakeNewQueueProcessedIterator struct {
	Event *StakeNewQueueProcessed // Event containing the contract specifics and raw log

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
func (it *StakeNewQueueProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewQueueProcessed)
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
		it.Event = new(StakeNewQueueProcessed)
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
func (it *StakeNewQueueProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewQueueProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewQueueProcessed represents a QueueProcessed event raised by the StakeNew contract.
type StakeNewQueueProcessed struct {
	QueueIndex *big.Int
	User       common.Address
	Amount     *big.Int
	StakeIndex uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQueueProcessed is a free log retrieval operation binding the contract event 0xf1a03c67438669ead639767cdce590abaefd1a7ff67aa39a78eaf34b42e23119.
//
// Solidity: event QueueProcessed(uint256 indexed queueIndex, address indexed user, uint160 amount, uint8 stakeIndex)
func (_StakeNew *StakeNewFilterer) FilterQueueProcessed(opts *bind.FilterOpts, queueIndex []*big.Int, user []common.Address) (*StakeNewQueueProcessedIterator, error) {

	var queueIndexRule []interface{}
	for _, queueIndexItem := range queueIndex {
		queueIndexRule = append(queueIndexRule, queueIndexItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "QueueProcessed", queueIndexRule, userRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewQueueProcessedIterator{contract: _StakeNew.contract, event: "QueueProcessed", logs: logs, sub: sub}, nil
}

// WatchQueueProcessed is a free log subscription operation binding the contract event 0xf1a03c67438669ead639767cdce590abaefd1a7ff67aa39a78eaf34b42e23119.
//
// Solidity: event QueueProcessed(uint256 indexed queueIndex, address indexed user, uint160 amount, uint8 stakeIndex)
func (_StakeNew *StakeNewFilterer) WatchQueueProcessed(opts *bind.WatchOpts, sink chan<- *StakeNewQueueProcessed, queueIndex []*big.Int, user []common.Address) (event.Subscription, error) {

	var queueIndexRule []interface{}
	for _, queueIndexItem := range queueIndex {
		queueIndexRule = append(queueIndexRule, queueIndexItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "QueueProcessed", queueIndexRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewQueueProcessed)
				if err := _StakeNew.contract.UnpackLog(event, "QueueProcessed", log); err != nil {
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

// ParseQueueProcessed is a log parse operation binding the contract event 0xf1a03c67438669ead639767cdce590abaefd1a7ff67aa39a78eaf34b42e23119.
//
// Solidity: event QueueProcessed(uint256 indexed queueIndex, address indexed user, uint160 amount, uint8 stakeIndex)
func (_StakeNew *StakeNewFilterer) ParseQueueProcessed(log types.Log) (*StakeNewQueueProcessed, error) {
	event := new(StakeNewQueueProcessed)
	if err := _StakeNew.contract.UnpackLog(event, "QueueProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewQueueRefundedIterator is returned from FilterQueueRefunded and is used to iterate over the raw logs and unpacked data for QueueRefunded events raised by the StakeNew contract.
type StakeNewQueueRefundedIterator struct {
	Event *StakeNewQueueRefunded // Event containing the contract specifics and raw log

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
func (it *StakeNewQueueRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewQueueRefunded)
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
		it.Event = new(StakeNewQueueRefunded)
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
func (it *StakeNewQueueRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewQueueRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewQueueRefunded represents a QueueRefunded event raised by the StakeNew contract.
type StakeNewQueueRefunded struct {
	QueueIndex *big.Int
	User       common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQueueRefunded is a free log retrieval operation binding the contract event 0xe5d2790f79d5c0f8c4f8535a45b0480cb7613fd990e3df63eac15cc86dc8ff25.
//
// Solidity: event QueueRefunded(uint256 indexed queueIndex, address indexed user, uint160 amount)
func (_StakeNew *StakeNewFilterer) FilterQueueRefunded(opts *bind.FilterOpts, queueIndex []*big.Int, user []common.Address) (*StakeNewQueueRefundedIterator, error) {

	var queueIndexRule []interface{}
	for _, queueIndexItem := range queueIndex {
		queueIndexRule = append(queueIndexRule, queueIndexItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "QueueRefunded", queueIndexRule, userRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewQueueRefundedIterator{contract: _StakeNew.contract, event: "QueueRefunded", logs: logs, sub: sub}, nil
}

// WatchQueueRefunded is a free log subscription operation binding the contract event 0xe5d2790f79d5c0f8c4f8535a45b0480cb7613fd990e3df63eac15cc86dc8ff25.
//
// Solidity: event QueueRefunded(uint256 indexed queueIndex, address indexed user, uint160 amount)
func (_StakeNew *StakeNewFilterer) WatchQueueRefunded(opts *bind.WatchOpts, sink chan<- *StakeNewQueueRefunded, queueIndex []*big.Int, user []common.Address) (event.Subscription, error) {

	var queueIndexRule []interface{}
	for _, queueIndexItem := range queueIndex {
		queueIndexRule = append(queueIndexRule, queueIndexItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "QueueRefunded", queueIndexRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewQueueRefunded)
				if err := _StakeNew.contract.UnpackLog(event, "QueueRefunded", log); err != nil {
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

// ParseQueueRefunded is a log parse operation binding the contract event 0xe5d2790f79d5c0f8c4f8535a45b0480cb7613fd990e3df63eac15cc86dc8ff25.
//
// Solidity: event QueueRefunded(uint256 indexed queueIndex, address indexed user, uint160 amount)
func (_StakeNew *StakeNewFilterer) ParseQueueRefunded(log types.Log) (*StakeNewQueueRefunded, error) {
	event := new(StakeNewQueueRefunded)
	if err := _StakeNew.contract.UnpackLog(event, "QueueRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewRestakedIterator is returned from FilterRestaked and is used to iterate over the raw logs and unpacked data for Restaked events raised by the StakeNew contract.
type StakeNewRestakedIterator struct {
	Event *StakeNewRestaked // Event containing the contract specifics and raw log

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
func (it *StakeNewRestakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewRestaked)
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
		it.Event = new(StakeNewRestaked)
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
func (it *StakeNewRestakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewRestakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewRestaked represents a Restaked event raised by the StakeNew contract.
type StakeNewRestaked struct {
	User      common.Address
	Timestamp *big.Int
	Index     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRestaked is a free log retrieval operation binding the contract event 0x22ea4d12074016dab9862b37c4e634b11084494dbeb71ef9af7c5b33b5511c1f.
//
// Solidity: event Restaked(address indexed user, uint40 timestamp, uint256 index)
func (_StakeNew *StakeNewFilterer) FilterRestaked(opts *bind.FilterOpts, user []common.Address) (*StakeNewRestakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "Restaked", userRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewRestakedIterator{contract: _StakeNew.contract, event: "Restaked", logs: logs, sub: sub}, nil
}

// WatchRestaked is a free log subscription operation binding the contract event 0x22ea4d12074016dab9862b37c4e634b11084494dbeb71ef9af7c5b33b5511c1f.
//
// Solidity: event Restaked(address indexed user, uint40 timestamp, uint256 index)
func (_StakeNew *StakeNewFilterer) WatchRestaked(opts *bind.WatchOpts, sink chan<- *StakeNewRestaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "Restaked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewRestaked)
				if err := _StakeNew.contract.UnpackLog(event, "Restaked", log); err != nil {
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

// ParseRestaked is a log parse operation binding the contract event 0x22ea4d12074016dab9862b37c4e634b11084494dbeb71ef9af7c5b33b5511c1f.
//
// Solidity: event Restaked(address indexed user, uint40 timestamp, uint256 index)
func (_StakeNew *StakeNewFilterer) ParseRestaked(log types.Log) (*StakeNewRestaked, error) {
	event := new(StakeNewRestaked)
	if err := _StakeNew.contract.UnpackLog(event, "Restaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewRewardBurnedIterator is returned from FilterRewardBurned and is used to iterate over the raw logs and unpacked data for RewardBurned events raised by the StakeNew contract.
type StakeNewRewardBurnedIterator struct {
	Event *StakeNewRewardBurned // Event containing the contract specifics and raw log

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
func (it *StakeNewRewardBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewRewardBurned)
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
		it.Event = new(StakeNewRewardBurned)
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
func (it *StakeNewRewardBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewRewardBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewRewardBurned represents a RewardBurned event raised by the StakeNew contract.
type StakeNewRewardBurned struct {
	User      common.Address
	Reward    *big.Int
	Timestamp *big.Int
	Index     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardBurned is a free log retrieval operation binding the contract event 0xa33fe9c91dd1a30125904f6c6b34e7816d74746a2b895d8224e1babd711e9cda.
//
// Solidity: event RewardBurned(address indexed user, uint160 reward, uint40 timestamp, uint256 index)
func (_StakeNew *StakeNewFilterer) FilterRewardBurned(opts *bind.FilterOpts, user []common.Address) (*StakeNewRewardBurnedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "RewardBurned", userRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewRewardBurnedIterator{contract: _StakeNew.contract, event: "RewardBurned", logs: logs, sub: sub}, nil
}

// WatchRewardBurned is a free log subscription operation binding the contract event 0xa33fe9c91dd1a30125904f6c6b34e7816d74746a2b895d8224e1babd711e9cda.
//
// Solidity: event RewardBurned(address indexed user, uint160 reward, uint40 timestamp, uint256 index)
func (_StakeNew *StakeNewFilterer) WatchRewardBurned(opts *bind.WatchOpts, sink chan<- *StakeNewRewardBurned, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "RewardBurned", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewRewardBurned)
				if err := _StakeNew.contract.UnpackLog(event, "RewardBurned", log); err != nil {
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

// ParseRewardBurned is a log parse operation binding the contract event 0xa33fe9c91dd1a30125904f6c6b34e7816d74746a2b895d8224e1babd711e9cda.
//
// Solidity: event RewardBurned(address indexed user, uint160 reward, uint40 timestamp, uint256 index)
func (_StakeNew *StakeNewFilterer) ParseRewardBurned(log types.Log) (*StakeNewRewardBurned, error) {
	event := new(StakeNewRewardBurned)
	if err := _StakeNew.contract.UnpackLog(event, "RewardBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewRewardCbNewOrderConfigUpdatedIterator is returned from FilterRewardCbNewOrderConfigUpdated and is used to iterate over the raw logs and unpacked data for RewardCbNewOrderConfigUpdated events raised by the StakeNew contract.
type StakeNewRewardCbNewOrderConfigUpdatedIterator struct {
	Event *StakeNewRewardCbNewOrderConfigUpdated // Event containing the contract specifics and raw log

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
func (it *StakeNewRewardCbNewOrderConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewRewardCbNewOrderConfigUpdated)
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
		it.Event = new(StakeNewRewardCbNewOrderConfigUpdated)
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
func (it *StakeNewRewardCbNewOrderConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewRewardCbNewOrderConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewRewardCbNewOrderConfigUpdated represents a RewardCbNewOrderConfigUpdated event raised by the StakeNew contract.
type StakeNewRewardCbNewOrderConfigUpdated struct {
	RewardEcoBpsTwo uint16
	RewardDirectBps uint16
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRewardCbNewOrderConfigUpdated is a free log retrieval operation binding the contract event 0x587e1e9050cb676838b874e90b6e0edf150348d8aad2622ec77d5105eaf4f37b.
//
// Solidity: event RewardCbNewOrderConfigUpdated(uint16 rewardEcoBpsTwo, uint16 rewardDirectBps)
func (_StakeNew *StakeNewFilterer) FilterRewardCbNewOrderConfigUpdated(opts *bind.FilterOpts) (*StakeNewRewardCbNewOrderConfigUpdatedIterator, error) {

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "RewardCbNewOrderConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &StakeNewRewardCbNewOrderConfigUpdatedIterator{contract: _StakeNew.contract, event: "RewardCbNewOrderConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardCbNewOrderConfigUpdated is a free log subscription operation binding the contract event 0x587e1e9050cb676838b874e90b6e0edf150348d8aad2622ec77d5105eaf4f37b.
//
// Solidity: event RewardCbNewOrderConfigUpdated(uint16 rewardEcoBpsTwo, uint16 rewardDirectBps)
func (_StakeNew *StakeNewFilterer) WatchRewardCbNewOrderConfigUpdated(opts *bind.WatchOpts, sink chan<- *StakeNewRewardCbNewOrderConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "RewardCbNewOrderConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewRewardCbNewOrderConfigUpdated)
				if err := _StakeNew.contract.UnpackLog(event, "RewardCbNewOrderConfigUpdated", log); err != nil {
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

// ParseRewardCbNewOrderConfigUpdated is a log parse operation binding the contract event 0x587e1e9050cb676838b874e90b6e0edf150348d8aad2622ec77d5105eaf4f37b.
//
// Solidity: event RewardCbNewOrderConfigUpdated(uint16 rewardEcoBpsTwo, uint16 rewardDirectBps)
func (_StakeNew *StakeNewFilterer) ParseRewardCbNewOrderConfigUpdated(log types.Log) (*StakeNewRewardCbNewOrderConfigUpdated, error) {
	event := new(StakeNewRewardCbNewOrderConfigUpdated)
	if err := _StakeNew.contract.UnpackLog(event, "RewardCbNewOrderConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewRewardFeeConfigUpdatedIterator is returned from FilterRewardFeeConfigUpdated and is used to iterate over the raw logs and unpacked data for RewardFeeConfigUpdated events raised by the StakeNew contract.
type StakeNewRewardFeeConfigUpdatedIterator struct {
	Event *StakeNewRewardFeeConfigUpdated // Event containing the contract specifics and raw log

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
func (it *StakeNewRewardFeeConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewRewardFeeConfigUpdated)
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
		it.Event = new(StakeNewRewardFeeConfigUpdated)
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
func (it *StakeNewRewardFeeConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewRewardFeeConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewRewardFeeConfigUpdated represents a RewardFeeConfigUpdated event raised by the StakeNew contract.
type StakeNewRewardFeeConfigUpdated struct {
	FeeTotalBps      uint16
	EcoBps           uint16
	GameBps          uint16
	GlobalBps        uint16
	S7Bps            uint16
	TeamRewardMaxBps uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRewardFeeConfigUpdated is a free log retrieval operation binding the contract event 0x46fa74113e8fc3a6da1fb647330b9b9aadc33cad953ad2ed26161f48480bf5d3.
//
// Solidity: event RewardFeeConfigUpdated(uint16 feeTotalBps, uint16 ecoBps, uint16 gameBps, uint16 globalBps, uint16 s7Bps, uint16 teamRewardMaxBps)
func (_StakeNew *StakeNewFilterer) FilterRewardFeeConfigUpdated(opts *bind.FilterOpts) (*StakeNewRewardFeeConfigUpdatedIterator, error) {

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "RewardFeeConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &StakeNewRewardFeeConfigUpdatedIterator{contract: _StakeNew.contract, event: "RewardFeeConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardFeeConfigUpdated is a free log subscription operation binding the contract event 0x46fa74113e8fc3a6da1fb647330b9b9aadc33cad953ad2ed26161f48480bf5d3.
//
// Solidity: event RewardFeeConfigUpdated(uint16 feeTotalBps, uint16 ecoBps, uint16 gameBps, uint16 globalBps, uint16 s7Bps, uint16 teamRewardMaxBps)
func (_StakeNew *StakeNewFilterer) WatchRewardFeeConfigUpdated(opts *bind.WatchOpts, sink chan<- *StakeNewRewardFeeConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "RewardFeeConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewRewardFeeConfigUpdated)
				if err := _StakeNew.contract.UnpackLog(event, "RewardFeeConfigUpdated", log); err != nil {
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

// ParseRewardFeeConfigUpdated is a log parse operation binding the contract event 0x46fa74113e8fc3a6da1fb647330b9b9aadc33cad953ad2ed26161f48480bf5d3.
//
// Solidity: event RewardFeeConfigUpdated(uint16 feeTotalBps, uint16 ecoBps, uint16 gameBps, uint16 globalBps, uint16 s7Bps, uint16 teamRewardMaxBps)
func (_StakeNew *StakeNewFilterer) ParseRewardFeeConfigUpdated(log types.Log) (*StakeNewRewardFeeConfigUpdated, error) {
	event := new(StakeNewRewardFeeConfigUpdated)
	if err := _StakeNew.contract.UnpackLog(event, "RewardFeeConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the StakeNew contract.
type StakeNewRewardPaidIterator struct {
	Event *StakeNewRewardPaid // Event containing the contract specifics and raw log

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
func (it *StakeNewRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewRewardPaid)
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
		it.Event = new(StakeNewRewardPaid)
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
func (it *StakeNewRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewRewardPaid represents a RewardPaid event raised by the StakeNew contract.
type StakeNewRewardPaid struct {
	User      common.Address
	Reward    *big.Int
	Timestamp *big.Int
	Index     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0x2b3ca68148d94d2c4f4e619c0125be1a33866e18aa56e0195aa4793107950f11.
//
// Solidity: event RewardPaid(address indexed user, uint160 reward, uint40 timestamp, uint256 index)
func (_StakeNew *StakeNewFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*StakeNewRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewRewardPaidIterator{contract: _StakeNew.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0x2b3ca68148d94d2c4f4e619c0125be1a33866e18aa56e0195aa4793107950f11.
//
// Solidity: event RewardPaid(address indexed user, uint160 reward, uint40 timestamp, uint256 index)
func (_StakeNew *StakeNewFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *StakeNewRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewRewardPaid)
				if err := _StakeNew.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0x2b3ca68148d94d2c4f4e619c0125be1a33866e18aa56e0195aa4793107950f11.
//
// Solidity: event RewardPaid(address indexed user, uint160 reward, uint40 timestamp, uint256 index)
func (_StakeNew *StakeNewFilterer) ParseRewardPaid(log types.Log) (*StakeNewRewardPaid, error) {
	event := new(StakeNewRewardPaid)
	if err := _StakeNew.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewRewardRouteUpdatedIterator is returned from FilterRewardRouteUpdated and is used to iterate over the raw logs and unpacked data for RewardRouteUpdated events raised by the StakeNew contract.
type StakeNewRewardRouteUpdatedIterator struct {
	Event *StakeNewRewardRouteUpdated // Event containing the contract specifics and raw log

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
func (it *StakeNewRewardRouteUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewRewardRouteUpdated)
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
		it.Event = new(StakeNewRewardRouteUpdated)
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
func (it *StakeNewRewardRouteUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewRewardRouteUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewRewardRouteUpdated represents a RewardRouteUpdated event raised by the StakeNew contract.
type StakeNewRewardRouteUpdated struct {
	EcoAddress  common.Address
	GameAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardRouteUpdated is a free log retrieval operation binding the contract event 0x55e0c6627408118aff29d16ddfa828a8589446f0301741f53621d777d3143b13.
//
// Solidity: event RewardRouteUpdated(address indexed ecoAddress, address indexed gameAddress)
func (_StakeNew *StakeNewFilterer) FilterRewardRouteUpdated(opts *bind.FilterOpts, ecoAddress []common.Address, gameAddress []common.Address) (*StakeNewRewardRouteUpdatedIterator, error) {

	var ecoAddressRule []interface{}
	for _, ecoAddressItem := range ecoAddress {
		ecoAddressRule = append(ecoAddressRule, ecoAddressItem)
	}
	var gameAddressRule []interface{}
	for _, gameAddressItem := range gameAddress {
		gameAddressRule = append(gameAddressRule, gameAddressItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "RewardRouteUpdated", ecoAddressRule, gameAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewRewardRouteUpdatedIterator{contract: _StakeNew.contract, event: "RewardRouteUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardRouteUpdated is a free log subscription operation binding the contract event 0x55e0c6627408118aff29d16ddfa828a8589446f0301741f53621d777d3143b13.
//
// Solidity: event RewardRouteUpdated(address indexed ecoAddress, address indexed gameAddress)
func (_StakeNew *StakeNewFilterer) WatchRewardRouteUpdated(opts *bind.WatchOpts, sink chan<- *StakeNewRewardRouteUpdated, ecoAddress []common.Address, gameAddress []common.Address) (event.Subscription, error) {

	var ecoAddressRule []interface{}
	for _, ecoAddressItem := range ecoAddress {
		ecoAddressRule = append(ecoAddressRule, ecoAddressItem)
	}
	var gameAddressRule []interface{}
	for _, gameAddressItem := range gameAddress {
		gameAddressRule = append(gameAddressRule, gameAddressItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "RewardRouteUpdated", ecoAddressRule, gameAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewRewardRouteUpdated)
				if err := _StakeNew.contract.UnpackLog(event, "RewardRouteUpdated", log); err != nil {
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

// ParseRewardRouteUpdated is a log parse operation binding the contract event 0x55e0c6627408118aff29d16ddfa828a8589446f0301741f53621d777d3143b13.
//
// Solidity: event RewardRouteUpdated(address indexed ecoAddress, address indexed gameAddress)
func (_StakeNew *StakeNewFilterer) ParseRewardRouteUpdated(log types.Log) (*StakeNewRewardRouteUpdated, error) {
	event := new(StakeNewRewardRouteUpdated)
	if err := _StakeNew.contract.UnpackLog(event, "RewardRouteUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewStageThresholdsUpdatedIterator is returned from FilterStageThresholdsUpdated and is used to iterate over the raw logs and unpacked data for StageThresholdsUpdated events raised by the StakeNew contract.
type StakeNewStageThresholdsUpdatedIterator struct {
	Event *StakeNewStageThresholdsUpdated // Event containing the contract specifics and raw log

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
func (it *StakeNewStageThresholdsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewStageThresholdsUpdated)
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
		it.Event = new(StakeNewStageThresholdsUpdated)
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
func (it *StakeNewStageThresholdsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewStageThresholdsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewStageThresholdsUpdated represents a StageThresholdsUpdated event raised by the StakeNew contract.
type StakeNewStageThresholdsUpdated struct {
	A   *big.Int
	B   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStageThresholdsUpdated is a free log retrieval operation binding the contract event 0x0915d81517168fd994e8bc88b5148625e36180e59b2bbe6965671119326d202b.
//
// Solidity: event StageThresholdsUpdated(uint256 a, uint256 b)
func (_StakeNew *StakeNewFilterer) FilterStageThresholdsUpdated(opts *bind.FilterOpts) (*StakeNewStageThresholdsUpdatedIterator, error) {

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "StageThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return &StakeNewStageThresholdsUpdatedIterator{contract: _StakeNew.contract, event: "StageThresholdsUpdated", logs: logs, sub: sub}, nil
}

// WatchStageThresholdsUpdated is a free log subscription operation binding the contract event 0x0915d81517168fd994e8bc88b5148625e36180e59b2bbe6965671119326d202b.
//
// Solidity: event StageThresholdsUpdated(uint256 a, uint256 b)
func (_StakeNew *StakeNewFilterer) WatchStageThresholdsUpdated(opts *bind.WatchOpts, sink chan<- *StakeNewStageThresholdsUpdated) (event.Subscription, error) {

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "StageThresholdsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewStageThresholdsUpdated)
				if err := _StakeNew.contract.UnpackLog(event, "StageThresholdsUpdated", log); err != nil {
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

// ParseStageThresholdsUpdated is a log parse operation binding the contract event 0x0915d81517168fd994e8bc88b5148625e36180e59b2bbe6965671119326d202b.
//
// Solidity: event StageThresholdsUpdated(uint256 a, uint256 b)
func (_StakeNew *StakeNewFilterer) ParseStageThresholdsUpdated(log types.Log) (*StakeNewStageThresholdsUpdated, error) {
	event := new(StakeNewStageThresholdsUpdated)
	if err := _StakeNew.contract.UnpackLog(event, "StageThresholdsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the StakeNew contract.
type StakeNewStakedIterator struct {
	Event *StakeNewStaked // Event containing the contract specifics and raw log

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
func (it *StakeNewStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewStaked)
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
		it.Event = new(StakeNewStaked)
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
func (it *StakeNewStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewStaked represents a Staked event raised by the StakeNew contract.
type StakeNewStaked struct {
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Index     *big.Int
	Duration  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x3a9398000906a7b3c6ab07ec3d849739a537132436f8db617248c858f632c56d.
//
// Solidity: event Staked(address indexed user, uint160 amount, uint40 timestamp, uint256 index, uint40 duration)
func (_StakeNew *StakeNewFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*StakeNewStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewStakedIterator{contract: _StakeNew.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x3a9398000906a7b3c6ab07ec3d849739a537132436f8db617248c858f632c56d.
//
// Solidity: event Staked(address indexed user, uint160 amount, uint40 timestamp, uint256 index, uint40 duration)
func (_StakeNew *StakeNewFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakeNewStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewStaked)
				if err := _StakeNew.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x3a9398000906a7b3c6ab07ec3d849739a537132436f8db617248c858f632c56d.
//
// Solidity: event Staked(address indexed user, uint160 amount, uint40 timestamp, uint256 index, uint40 duration)
func (_StakeNew *StakeNewFilterer) ParseStaked(log types.Log) (*StakeNewStaked, error) {
	event := new(StakeNewStaked)
	if err := _StakeNew.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StakeNew contract.
type StakeNewTransferIterator struct {
	Event *StakeNewTransfer // Event containing the contract specifics and raw log

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
func (it *StakeNewTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewTransfer)
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
		it.Event = new(StakeNewTransfer)
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
func (it *StakeNewTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewTransfer represents a Transfer event raised by the StakeNew contract.
type StakeNewTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_StakeNew *StakeNewFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StakeNewTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewTransferIterator{contract: _StakeNew.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_StakeNew *StakeNewFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StakeNewTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewTransfer)
				if err := _StakeNew.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_StakeNew *StakeNewFilterer) ParseTransfer(log types.Log) (*StakeNewTransfer, error) {
	event := new(StakeNewTransfer)
	if err := _StakeNew.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StakeNew contract.
type StakeNewUnpausedIterator struct {
	Event *StakeNewUnpaused // Event containing the contract specifics and raw log

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
func (it *StakeNewUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewUnpaused)
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
		it.Event = new(StakeNewUnpaused)
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
func (it *StakeNewUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewUnpaused represents a Unpaused event raised by the StakeNew contract.
type StakeNewUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address indexed account)
func (_StakeNew *StakeNewFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*StakeNewUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewUnpausedIterator{contract: _StakeNew.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address indexed account)
func (_StakeNew *StakeNewFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakeNewUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewUnpaused)
				if err := _StakeNew.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address indexed account)
func (_StakeNew *StakeNewFilterer) ParseUnpaused(log types.Log) (*StakeNewUnpaused, error) {
	event := new(StakeNewUnpaused)
	if err := _StakeNew.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeNewUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the StakeNew contract.
type StakeNewUnstakedIterator struct {
	Event *StakeNewUnstaked // Event containing the contract specifics and raw log

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
func (it *StakeNewUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeNewUnstaked)
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
		it.Event = new(StakeNewUnstaked)
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
func (it *StakeNewUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeNewUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeNewUnstaked represents a Unstaked event raised by the StakeNew contract.
type StakeNewUnstaked struct {
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Index     *big.Int
	Reward    *big.Int
	Ttl       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0xb0b25333efd643d6b20b95dfa2e50b21744cf2810221c85dc17334d32e35802b.
//
// Solidity: event Unstaked(address indexed user, uint160 amount, uint40 timestamp, uint256 index, uint160 reward, uint40 ttl)
func (_StakeNew *StakeNewFilterer) FilterUnstaked(opts *bind.FilterOpts, user []common.Address) (*StakeNewUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNew.contract.FilterLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return &StakeNewUnstakedIterator{contract: _StakeNew.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0xb0b25333efd643d6b20b95dfa2e50b21744cf2810221c85dc17334d32e35802b.
//
// Solidity: event Unstaked(address indexed user, uint160 amount, uint40 timestamp, uint256 index, uint160 reward, uint40 ttl)
func (_StakeNew *StakeNewFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *StakeNewUnstaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakeNew.contract.WatchLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeNewUnstaked)
				if err := _StakeNew.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0xb0b25333efd643d6b20b95dfa2e50b21744cf2810221c85dc17334d32e35802b.
//
// Solidity: event Unstaked(address indexed user, uint160 amount, uint40 timestamp, uint256 index, uint160 reward, uint40 ttl)
func (_StakeNew *StakeNewFilterer) ParseUnstaked(log types.Log) (*StakeNewUnstaked, error) {
	event := new(StakeNewUnstaked)
	if err := _StakeNew.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
