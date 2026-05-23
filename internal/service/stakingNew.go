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

// StakingV1Order is an auto generated low-level Go binding around an user-defined struct.
type StakingV1Order struct {
	Account        common.Address
	Amount         *big.Int
	Cap            *big.Int
	Used           *big.Int
	LinePaid       *big.Int
	Created        *big.Int
	Start          *big.Int
	ClaimEffective *big.Int
	DaysCount      uint32
	Exited         bool
}

// StakingNewMetaData contains all meta data concerning the StakingNew contract.
var StakingNewMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usdt_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"market_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eco_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sink_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"init_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ST_BAD_PARAM\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ST_LIMIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ST_LOCKED\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ST_NOT_BOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ST_NO_AMOUNT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ST_NO_AUTH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ST_TOKEN_CALL\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ST_ZERO_ADDR\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"payU\",\"type\":\"bool\"}],\"name\":\"AntiDumpLinePaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"recovery\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"compBpsPerDay\",\"type\":\"uint16\"}],\"name\":\"CircuitBreakerParamsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"circuitBreakerTime\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"newHighTime\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentDropBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"action\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roundCount\",\"type\":\"uint256\"}],\"name\":\"CircuitBreakerStateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"liqBps\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"marketBps\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"ecoBps\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"teamMaxBps\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"teamDirectBps\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"smallFeeBps\",\"type\":\"uint16\"}],\"name\":\"EntryFeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldInit\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newInit\",\"type\":\"address\"}],\"name\":\"InitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"cutoff\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"done\",\"type\":\"uint256\"}],\"name\":\"LaunchQueueSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"grossU\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeU\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paidMs\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"msAmount\",\"type\":\"uint256\"}],\"name\":\"LineClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"lineFeeBps\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"lineEcoBps\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16[4]\",\"name\":\"lineLevelBps\",\"type\":\"uint16[4]\"}],\"name\":\"LineFeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"secondsAmount\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"beforeTime\",\"type\":\"uint40\"}],\"name\":\"LineSafeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"name\":\"LockedSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCap\",\"type\":\"uint256\"}],\"name\":\"OrderCapSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"plan\",\"type\":\"uint256\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"start\",\"type\":\"uint40\"}],\"name\":\"OrderEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"used\",\"type\":\"uint256\"}],\"name\":\"OrderExited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"daysCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"PlanSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"queueWait\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"autoQueueCount\",\"type\":\"uint256\"}],\"name\":\"QueueConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liqU\",\"type\":\"uint256\"}],\"name\":\"QueueDone\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liqU\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"queuedAt\",\"type\":\"uint40\"}],\"name\":\"Queued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usdt\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"SystemsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stored\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"directPaid\",\"type\":\"uint256\"}],\"name\":\"TeamBooked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gross\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"net\",\"type\":\"uint256\"}],\"name\":\"TeamClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"period\",\"type\":\"uint40\"}],\"name\":\"TeamClearConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"time\",\"type\":\"uint40\"}],\"name\":\"TeamCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TeamExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"eco\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sink\",\"type\":\"address\"}],\"name\":\"WalletsSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AMM_FEE_MUL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEAD\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DENOM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUTE1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUTE1_BPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUTE2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUTE2_BPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"antiDumpLinePayU\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approveAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"autoQueueCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"capOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"used\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"comp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cbCompBpsPerDay\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"circuitBreakerRecovery\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"circuitBreakerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"circuitBreakerTime\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"claimLine\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"grossU\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"paidMs\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxOrders\",\"type\":\"uint256\"}],\"name\":\"claimLineAll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalU\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"claimLineBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalU\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimTeam\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"net\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eco\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ecoBps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"cap\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"used\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"linePaid\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"created\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"start\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"claimEffective\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"daysCount\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exited\",\"type\":\"bool\"}],\"internalType\":\"structStakingV1.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getOrderById\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"cap\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"used\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"linePaid\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"created\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"start\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"claimEffective\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"daysCount\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exited\",\"type\":\"bool\"}],\"internalType\":\"structStakingV1.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usdt_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"router_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"}],\"name\":\"initSystems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"launchQueueAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"launchQueueCutoff\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"launchQueueDone\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"levelRewardUSDT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"lineClaimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"lineClaimableDetail\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"effective\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lineEcoBps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lineFeeBps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lineLevelBps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lineSafeBefore\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lineSafeSec\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liqBps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"market\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketBps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newHighTime\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOrderId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"orderCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"orderIdAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"cap\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"used\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"linePaid\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"created\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"start\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"claimEffective\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"daysCount\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exited\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"planLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"plans\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"minAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"outAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"daysCount\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxItems\",\"type\":\"uint256\"}],\"name\":\"processQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"done\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liqU\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"queuedAt\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueHead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueTail\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueWait\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"rList\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"start\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"end\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rListLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rounds\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"start\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"end\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"contractIRouterS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"payU\",\"type\":\"bool\"}],\"name\":\"setAntiDumpLinePayU\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"recovery\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"compBps\",\"type\":\"uint16\"}],\"name\":\"setCircuitBreakerParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"liq_\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"market_\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"eco_\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"team_\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"direct_\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"fee_\",\"type\":\"uint16\"}],\"name\":\"setEntryFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newInit\",\"type\":\"address\"}],\"name\":\"setInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"fee_\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"eco_\",\"type\":\"uint16\"},{\"internalType\":\"uint16[4]\",\"name\":\"lv_\",\"type\":\"uint16[4]\"}],\"name\":\"setLineFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"secondsAmount\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"beforeTime\",\"type\":\"uint40\"}],\"name\":\"setLineSafe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"values\",\"type\":\"bool[]\"}],\"name\":\"setLocked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"newCap\",\"type\":\"uint128\"}],\"name\":\"setOrderCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"minAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"maxAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"outAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"daysCount\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setPlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"wait_\",\"type\":\"uint40\"},{\"internalType\":\"uint256\",\"name\":\"autoCount_\",\"type\":\"uint256\"}],\"name\":\"setQueueConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"period\",\"type\":\"uint40\"}],\"name\":\"setTeamClearConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eco_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sink_\",\"type\":\"address\"}],\"name\":\"setWallets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shouldUpdateCircuitBreaker\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"needTrigger\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"needCountdown\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"needRecover\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"currentDropBps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sink\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"smallFeeBps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"plan\",\"type\":\"uint16\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"teamClaimedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"teamClearAt\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teamClearPeriod\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teamDirectBps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teamMaxBps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"teamU\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIMsTokenS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateCircuitBreaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdt\",\"outputs\":[{\"internalType\":\"contractIERC20S\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"user\",\"outputs\":[{\"internalType\":\"contractIUserS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakingNewABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingNewMetaData.ABI instead.
var StakingNewABI = StakingNewMetaData.ABI

// StakingNew is an auto generated Go binding around an Ethereum contract.
type StakingNew struct {
	StakingNewCaller     // Read-only binding to the contract
	StakingNewTransactor // Write-only binding to the contract
	StakingNewFilterer   // Log filterer for contract events
}

// StakingNewCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingNewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingNewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingNewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingNewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingNewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingNewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingNewSession struct {
	Contract     *StakingNew       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingNewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingNewCallerSession struct {
	Contract *StakingNewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StakingNewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingNewTransactorSession struct {
	Contract     *StakingNewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StakingNewRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingNewRaw struct {
	Contract *StakingNew // Generic contract binding to access the raw methods on
}

// StakingNewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingNewCallerRaw struct {
	Contract *StakingNewCaller // Generic read-only contract binding to access the raw methods on
}

// StakingNewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingNewTransactorRaw struct {
	Contract *StakingNewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingNew creates a new instance of StakingNew, bound to a specific deployed contract.
func NewStakingNew(address common.Address, backend bind.ContractBackend) (*StakingNew, error) {
	contract, err := bindStakingNew(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingNew{StakingNewCaller: StakingNewCaller{contract: contract}, StakingNewTransactor: StakingNewTransactor{contract: contract}, StakingNewFilterer: StakingNewFilterer{contract: contract}}, nil
}

// NewStakingNewCaller creates a new read-only instance of StakingNew, bound to a specific deployed contract.
func NewStakingNewCaller(address common.Address, caller bind.ContractCaller) (*StakingNewCaller, error) {
	contract, err := bindStakingNew(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingNewCaller{contract: contract}, nil
}

// NewStakingNewTransactor creates a new write-only instance of StakingNew, bound to a specific deployed contract.
func NewStakingNewTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingNewTransactor, error) {
	contract, err := bindStakingNew(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingNewTransactor{contract: contract}, nil
}

// NewStakingNewFilterer creates a new log filterer instance of StakingNew, bound to a specific deployed contract.
func NewStakingNewFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingNewFilterer, error) {
	contract, err := bindStakingNew(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingNewFilterer{contract: contract}, nil
}

// bindStakingNew binds a generic wrapper to an already deployed contract.
func bindStakingNew(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingNewABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingNew *StakingNewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingNew.Contract.StakingNewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingNew *StakingNewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingNew.Contract.StakingNewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingNew *StakingNewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingNew.Contract.StakingNewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingNew *StakingNewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingNew.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingNew *StakingNewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingNew.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingNew *StakingNewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingNew.Contract.contract.Transact(opts, method, params...)
}

// AMMFEEMUL is a free data retrieval call binding the contract method 0xbc12eee8.
//
// Solidity: function AMM_FEE_MUL() view returns(uint256)
func (_StakingNew *StakingNewCaller) AMMFEEMUL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "AMM_FEE_MUL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AMMFEEMUL is a free data retrieval call binding the contract method 0xbc12eee8.
//
// Solidity: function AMM_FEE_MUL() view returns(uint256)
func (_StakingNew *StakingNewSession) AMMFEEMUL() (*big.Int, error) {
	return _StakingNew.Contract.AMMFEEMUL(&_StakingNew.CallOpts)
}

// AMMFEEMUL is a free data retrieval call binding the contract method 0xbc12eee8.
//
// Solidity: function AMM_FEE_MUL() view returns(uint256)
func (_StakingNew *StakingNewCallerSession) AMMFEEMUL() (*big.Int, error) {
	return _StakingNew.Contract.AMMFEEMUL(&_StakingNew.CallOpts)
}

// DEAD is a free data retrieval call binding the contract method 0x03fd2a45.
//
// Solidity: function DEAD() view returns(address)
func (_StakingNew *StakingNewCaller) DEAD(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "DEAD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEAD is a free data retrieval call binding the contract method 0x03fd2a45.
//
// Solidity: function DEAD() view returns(address)
func (_StakingNew *StakingNewSession) DEAD() (common.Address, error) {
	return _StakingNew.Contract.DEAD(&_StakingNew.CallOpts)
}

// DEAD is a free data retrieval call binding the contract method 0x03fd2a45.
//
// Solidity: function DEAD() view returns(address)
func (_StakingNew *StakingNewCallerSession) DEAD() (common.Address, error) {
	return _StakingNew.Contract.DEAD(&_StakingNew.CallOpts)
}

// DENOM is a free data retrieval call binding the contract method 0x16343da4.
//
// Solidity: function DENOM() view returns(uint256)
func (_StakingNew *StakingNewCaller) DENOM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "DENOM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DENOM is a free data retrieval call binding the contract method 0x16343da4.
//
// Solidity: function DENOM() view returns(uint256)
func (_StakingNew *StakingNewSession) DENOM() (*big.Int, error) {
	return _StakingNew.Contract.DENOM(&_StakingNew.CallOpts)
}

// DENOM is a free data retrieval call binding the contract method 0x16343da4.
//
// Solidity: function DENOM() view returns(uint256)
func (_StakingNew *StakingNewCallerSession) DENOM() (*big.Int, error) {
	return _StakingNew.Contract.DENOM(&_StakingNew.CallOpts)
}

// ROUTE1 is a free data retrieval call binding the contract method 0x9d9ede2b.
//
// Solidity: function ROUTE1() view returns(address)
func (_StakingNew *StakingNewCaller) ROUTE1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "ROUTE1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ROUTE1 is a free data retrieval call binding the contract method 0x9d9ede2b.
//
// Solidity: function ROUTE1() view returns(address)
func (_StakingNew *StakingNewSession) ROUTE1() (common.Address, error) {
	return _StakingNew.Contract.ROUTE1(&_StakingNew.CallOpts)
}

// ROUTE1 is a free data retrieval call binding the contract method 0x9d9ede2b.
//
// Solidity: function ROUTE1() view returns(address)
func (_StakingNew *StakingNewCallerSession) ROUTE1() (common.Address, error) {
	return _StakingNew.Contract.ROUTE1(&_StakingNew.CallOpts)
}

// ROUTE1BPS is a free data retrieval call binding the contract method 0x443e6074.
//
// Solidity: function ROUTE1_BPS() view returns(uint16)
func (_StakingNew *StakingNewCaller) ROUTE1BPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "ROUTE1_BPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ROUTE1BPS is a free data retrieval call binding the contract method 0x443e6074.
//
// Solidity: function ROUTE1_BPS() view returns(uint16)
func (_StakingNew *StakingNewSession) ROUTE1BPS() (uint16, error) {
	return _StakingNew.Contract.ROUTE1BPS(&_StakingNew.CallOpts)
}

// ROUTE1BPS is a free data retrieval call binding the contract method 0x443e6074.
//
// Solidity: function ROUTE1_BPS() view returns(uint16)
func (_StakingNew *StakingNewCallerSession) ROUTE1BPS() (uint16, error) {
	return _StakingNew.Contract.ROUTE1BPS(&_StakingNew.CallOpts)
}

// ROUTE2 is a free data retrieval call binding the contract method 0xcb1cba05.
//
// Solidity: function ROUTE2() view returns(address)
func (_StakingNew *StakingNewCaller) ROUTE2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "ROUTE2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ROUTE2 is a free data retrieval call binding the contract method 0xcb1cba05.
//
// Solidity: function ROUTE2() view returns(address)
func (_StakingNew *StakingNewSession) ROUTE2() (common.Address, error) {
	return _StakingNew.Contract.ROUTE2(&_StakingNew.CallOpts)
}

// ROUTE2 is a free data retrieval call binding the contract method 0xcb1cba05.
//
// Solidity: function ROUTE2() view returns(address)
func (_StakingNew *StakingNewCallerSession) ROUTE2() (common.Address, error) {
	return _StakingNew.Contract.ROUTE2(&_StakingNew.CallOpts)
}

// ROUTE2BPS is a free data retrieval call binding the contract method 0xaff819c2.
//
// Solidity: function ROUTE2_BPS() view returns(uint16)
func (_StakingNew *StakingNewCaller) ROUTE2BPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "ROUTE2_BPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ROUTE2BPS is a free data retrieval call binding the contract method 0xaff819c2.
//
// Solidity: function ROUTE2_BPS() view returns(uint16)
func (_StakingNew *StakingNewSession) ROUTE2BPS() (uint16, error) {
	return _StakingNew.Contract.ROUTE2BPS(&_StakingNew.CallOpts)
}

// ROUTE2BPS is a free data retrieval call binding the contract method 0xaff819c2.
//
// Solidity: function ROUTE2_BPS() view returns(uint16)
func (_StakingNew *StakingNewCallerSession) ROUTE2BPS() (uint16, error) {
	return _StakingNew.Contract.ROUTE2BPS(&_StakingNew.CallOpts)
}

// AntiDumpLinePayU is a free data retrieval call binding the contract method 0x2e3dcd36.
//
// Solidity: function antiDumpLinePayU() view returns(bool)
func (_StakingNew *StakingNewCaller) AntiDumpLinePayU(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "antiDumpLinePayU")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AntiDumpLinePayU is a free data retrieval call binding the contract method 0x2e3dcd36.
//
// Solidity: function antiDumpLinePayU() view returns(bool)
func (_StakingNew *StakingNewSession) AntiDumpLinePayU() (bool, error) {
	return _StakingNew.Contract.AntiDumpLinePayU(&_StakingNew.CallOpts)
}

// AntiDumpLinePayU is a free data retrieval call binding the contract method 0x2e3dcd36.
//
// Solidity: function antiDumpLinePayU() view returns(bool)
func (_StakingNew *StakingNewCallerSession) AntiDumpLinePayU() (bool, error) {
	return _StakingNew.Contract.AntiDumpLinePayU(&_StakingNew.CallOpts)
}

// AutoQueueCount is a free data retrieval call binding the contract method 0xf82cc037.
//
// Solidity: function autoQueueCount() view returns(uint256)
func (_StakingNew *StakingNewCaller) AutoQueueCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "autoQueueCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AutoQueueCount is a free data retrieval call binding the contract method 0xf82cc037.
//
// Solidity: function autoQueueCount() view returns(uint256)
func (_StakingNew *StakingNewSession) AutoQueueCount() (*big.Int, error) {
	return _StakingNew.Contract.AutoQueueCount(&_StakingNew.CallOpts)
}

// AutoQueueCount is a free data retrieval call binding the contract method 0xf82cc037.
//
// Solidity: function autoQueueCount() view returns(uint256)
func (_StakingNew *StakingNewCallerSession) AutoQueueCount() (*big.Int, error) {
	return _StakingNew.Contract.AutoQueueCount(&_StakingNew.CallOpts)
}

// CapOf is a free data retrieval call binding the contract method 0x62aef832.
//
// Solidity: function capOf(address account, uint256 index) view returns(uint256 cap, uint256 used, uint256 left, uint256 comp)
func (_StakingNew *StakingNewCaller) CapOf(opts *bind.CallOpts, account common.Address, index *big.Int) (struct {
	Cap  *big.Int
	Used *big.Int
	Left *big.Int
	Comp *big.Int
}, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "capOf", account, index)

	outstruct := new(struct {
		Cap  *big.Int
		Used *big.Int
		Left *big.Int
		Comp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Cap = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Used = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Left = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Comp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CapOf is a free data retrieval call binding the contract method 0x62aef832.
//
// Solidity: function capOf(address account, uint256 index) view returns(uint256 cap, uint256 used, uint256 left, uint256 comp)
func (_StakingNew *StakingNewSession) CapOf(account common.Address, index *big.Int) (struct {
	Cap  *big.Int
	Used *big.Int
	Left *big.Int
	Comp *big.Int
}, error) {
	return _StakingNew.Contract.CapOf(&_StakingNew.CallOpts, account, index)
}

// CapOf is a free data retrieval call binding the contract method 0x62aef832.
//
// Solidity: function capOf(address account, uint256 index) view returns(uint256 cap, uint256 used, uint256 left, uint256 comp)
func (_StakingNew *StakingNewCallerSession) CapOf(account common.Address, index *big.Int) (struct {
	Cap  *big.Int
	Used *big.Int
	Left *big.Int
	Comp *big.Int
}, error) {
	return _StakingNew.Contract.CapOf(&_StakingNew.CallOpts, account, index)
}

// CbCompBpsPerDay is a free data retrieval call binding the contract method 0xd21103eb.
//
// Solidity: function cbCompBpsPerDay() view returns(uint16)
func (_StakingNew *StakingNewCaller) CbCompBpsPerDay(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "cbCompBpsPerDay")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// CbCompBpsPerDay is a free data retrieval call binding the contract method 0xd21103eb.
//
// Solidity: function cbCompBpsPerDay() view returns(uint16)
func (_StakingNew *StakingNewSession) CbCompBpsPerDay() (uint16, error) {
	return _StakingNew.Contract.CbCompBpsPerDay(&_StakingNew.CallOpts)
}

// CbCompBpsPerDay is a free data retrieval call binding the contract method 0xd21103eb.
//
// Solidity: function cbCompBpsPerDay() view returns(uint16)
func (_StakingNew *StakingNewCallerSession) CbCompBpsPerDay() (uint16, error) {
	return _StakingNew.Contract.CbCompBpsPerDay(&_StakingNew.CallOpts)
}

// CircuitBreakerRecovery is a free data retrieval call binding the contract method 0x86b54614.
//
// Solidity: function circuitBreakerRecovery() view returns(uint40)
func (_StakingNew *StakingNewCaller) CircuitBreakerRecovery(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "circuitBreakerRecovery")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CircuitBreakerRecovery is a free data retrieval call binding the contract method 0x86b54614.
//
// Solidity: function circuitBreakerRecovery() view returns(uint40)
func (_StakingNew *StakingNewSession) CircuitBreakerRecovery() (*big.Int, error) {
	return _StakingNew.Contract.CircuitBreakerRecovery(&_StakingNew.CallOpts)
}

// CircuitBreakerRecovery is a free data retrieval call binding the contract method 0x86b54614.
//
// Solidity: function circuitBreakerRecovery() view returns(uint40)
func (_StakingNew *StakingNewCallerSession) CircuitBreakerRecovery() (*big.Int, error) {
	return _StakingNew.Contract.CircuitBreakerRecovery(&_StakingNew.CallOpts)
}

// CircuitBreakerThreshold is a free data retrieval call binding the contract method 0x7d0b0a34.
//
// Solidity: function circuitBreakerThreshold() view returns(uint256)
func (_StakingNew *StakingNewCaller) CircuitBreakerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "circuitBreakerThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CircuitBreakerThreshold is a free data retrieval call binding the contract method 0x7d0b0a34.
//
// Solidity: function circuitBreakerThreshold() view returns(uint256)
func (_StakingNew *StakingNewSession) CircuitBreakerThreshold() (*big.Int, error) {
	return _StakingNew.Contract.CircuitBreakerThreshold(&_StakingNew.CallOpts)
}

// CircuitBreakerThreshold is a free data retrieval call binding the contract method 0x7d0b0a34.
//
// Solidity: function circuitBreakerThreshold() view returns(uint256)
func (_StakingNew *StakingNewCallerSession) CircuitBreakerThreshold() (*big.Int, error) {
	return _StakingNew.Contract.CircuitBreakerThreshold(&_StakingNew.CallOpts)
}

// CircuitBreakerTime is a free data retrieval call binding the contract method 0x7bd1bb14.
//
// Solidity: function circuitBreakerTime() view returns(uint40)
func (_StakingNew *StakingNewCaller) CircuitBreakerTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "circuitBreakerTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CircuitBreakerTime is a free data retrieval call binding the contract method 0x7bd1bb14.
//
// Solidity: function circuitBreakerTime() view returns(uint40)
func (_StakingNew *StakingNewSession) CircuitBreakerTime() (*big.Int, error) {
	return _StakingNew.Contract.CircuitBreakerTime(&_StakingNew.CallOpts)
}

// CircuitBreakerTime is a free data retrieval call binding the contract method 0x7bd1bb14.
//
// Solidity: function circuitBreakerTime() view returns(uint40)
func (_StakingNew *StakingNewCallerSession) CircuitBreakerTime() (*big.Int, error) {
	return _StakingNew.Contract.CircuitBreakerTime(&_StakingNew.CallOpts)
}

// Eco is a free data retrieval call binding the contract method 0x94e63e63.
//
// Solidity: function eco() view returns(address)
func (_StakingNew *StakingNewCaller) Eco(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "eco")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Eco is a free data retrieval call binding the contract method 0x94e63e63.
//
// Solidity: function eco() view returns(address)
func (_StakingNew *StakingNewSession) Eco() (common.Address, error) {
	return _StakingNew.Contract.Eco(&_StakingNew.CallOpts)
}

// Eco is a free data retrieval call binding the contract method 0x94e63e63.
//
// Solidity: function eco() view returns(address)
func (_StakingNew *StakingNewCallerSession) Eco() (common.Address, error) {
	return _StakingNew.Contract.Eco(&_StakingNew.CallOpts)
}

// EcoBps is a free data retrieval call binding the contract method 0xc458d0b7.
//
// Solidity: function ecoBps() view returns(uint16)
func (_StakingNew *StakingNewCaller) EcoBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "ecoBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// EcoBps is a free data retrieval call binding the contract method 0xc458d0b7.
//
// Solidity: function ecoBps() view returns(uint16)
func (_StakingNew *StakingNewSession) EcoBps() (uint16, error) {
	return _StakingNew.Contract.EcoBps(&_StakingNew.CallOpts)
}

// EcoBps is a free data retrieval call binding the contract method 0xc458d0b7.
//
// Solidity: function ecoBps() view returns(uint16)
func (_StakingNew *StakingNewCallerSession) EcoBps() (uint16, error) {
	return _StakingNew.Contract.EcoBps(&_StakingNew.CallOpts)
}

// GetOrder is a free data retrieval call binding the contract method 0xedb25841.
//
// Solidity: function getOrder(address account, uint256 index) view returns((address,uint128,uint128,uint128,uint128,uint40,uint40,uint40,uint32,bool))
func (_StakingNew *StakingNewCaller) GetOrder(opts *bind.CallOpts, account common.Address, index *big.Int) (StakingV1Order, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "getOrder", account, index)

	if err != nil {
		return *new(StakingV1Order), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingV1Order)).(*StakingV1Order)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0xedb25841.
//
// Solidity: function getOrder(address account, uint256 index) view returns((address,uint128,uint128,uint128,uint128,uint40,uint40,uint40,uint32,bool))
func (_StakingNew *StakingNewSession) GetOrder(account common.Address, index *big.Int) (StakingV1Order, error) {
	return _StakingNew.Contract.GetOrder(&_StakingNew.CallOpts, account, index)
}

// GetOrder is a free data retrieval call binding the contract method 0xedb25841.
//
// Solidity: function getOrder(address account, uint256 index) view returns((address,uint128,uint128,uint128,uint128,uint40,uint40,uint40,uint32,bool))
func (_StakingNew *StakingNewCallerSession) GetOrder(account common.Address, index *big.Int) (StakingV1Order, error) {
	return _StakingNew.Contract.GetOrder(&_StakingNew.CallOpts, account, index)
}

// GetOrderById is a free data retrieval call binding the contract method 0x42b50a7a.
//
// Solidity: function getOrderById(uint256 id) view returns((address,uint128,uint128,uint128,uint128,uint40,uint40,uint40,uint32,bool))
func (_StakingNew *StakingNewCaller) GetOrderById(opts *bind.CallOpts, id *big.Int) (StakingV1Order, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "getOrderById", id)

	if err != nil {
		return *new(StakingV1Order), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingV1Order)).(*StakingV1Order)

	return out0, err

}

// GetOrderById is a free data retrieval call binding the contract method 0x42b50a7a.
//
// Solidity: function getOrderById(uint256 id) view returns((address,uint128,uint128,uint128,uint128,uint40,uint40,uint40,uint32,bool))
func (_StakingNew *StakingNewSession) GetOrderById(id *big.Int) (StakingV1Order, error) {
	return _StakingNew.Contract.GetOrderById(&_StakingNew.CallOpts, id)
}

// GetOrderById is a free data retrieval call binding the contract method 0x42b50a7a.
//
// Solidity: function getOrderById(uint256 id) view returns((address,uint128,uint128,uint128,uint128,uint40,uint40,uint40,uint32,bool))
func (_StakingNew *StakingNewCallerSession) GetOrderById(id *big.Int) (StakingV1Order, error) {
	return _StakingNew.Contract.GetOrderById(&_StakingNew.CallOpts, id)
}

// Init is a free data retrieval call binding the contract method 0xe1c7392a.
//
// Solidity: function init() view returns(address)
func (_StakingNew *StakingNewCaller) Init(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "init")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Init is a free data retrieval call binding the contract method 0xe1c7392a.
//
// Solidity: function init() view returns(address)
func (_StakingNew *StakingNewSession) Init() (common.Address, error) {
	return _StakingNew.Contract.Init(&_StakingNew.CallOpts)
}

// Init is a free data retrieval call binding the contract method 0xe1c7392a.
//
// Solidity: function init() view returns(address)
func (_StakingNew *StakingNewCallerSession) Init() (common.Address, error) {
	return _StakingNew.Contract.Init(&_StakingNew.CallOpts)
}

// LaunchQueueAmount is a free data retrieval call binding the contract method 0x33762622.
//
// Solidity: function launchQueueAmount() view returns(uint256)
func (_StakingNew *StakingNewCaller) LaunchQueueAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "launchQueueAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LaunchQueueAmount is a free data retrieval call binding the contract method 0x33762622.
//
// Solidity: function launchQueueAmount() view returns(uint256)
func (_StakingNew *StakingNewSession) LaunchQueueAmount() (*big.Int, error) {
	return _StakingNew.Contract.LaunchQueueAmount(&_StakingNew.CallOpts)
}

// LaunchQueueAmount is a free data retrieval call binding the contract method 0x33762622.
//
// Solidity: function launchQueueAmount() view returns(uint256)
func (_StakingNew *StakingNewCallerSession) LaunchQueueAmount() (*big.Int, error) {
	return _StakingNew.Contract.LaunchQueueAmount(&_StakingNew.CallOpts)
}

// LaunchQueueCutoff is a free data retrieval call binding the contract method 0xf6069272.
//
// Solidity: function launchQueueCutoff() view returns(uint40)
func (_StakingNew *StakingNewCaller) LaunchQueueCutoff(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "launchQueueCutoff")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LaunchQueueCutoff is a free data retrieval call binding the contract method 0xf6069272.
//
// Solidity: function launchQueueCutoff() view returns(uint40)
func (_StakingNew *StakingNewSession) LaunchQueueCutoff() (*big.Int, error) {
	return _StakingNew.Contract.LaunchQueueCutoff(&_StakingNew.CallOpts)
}

// LaunchQueueCutoff is a free data retrieval call binding the contract method 0xf6069272.
//
// Solidity: function launchQueueCutoff() view returns(uint40)
func (_StakingNew *StakingNewCallerSession) LaunchQueueCutoff() (*big.Int, error) {
	return _StakingNew.Contract.LaunchQueueCutoff(&_StakingNew.CallOpts)
}

// LaunchQueueDone is a free data retrieval call binding the contract method 0x8717ba17.
//
// Solidity: function launchQueueDone() view returns(uint256)
func (_StakingNew *StakingNewCaller) LaunchQueueDone(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "launchQueueDone")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LaunchQueueDone is a free data retrieval call binding the contract method 0x8717ba17.
//
// Solidity: function launchQueueDone() view returns(uint256)
func (_StakingNew *StakingNewSession) LaunchQueueDone() (*big.Int, error) {
	return _StakingNew.Contract.LaunchQueueDone(&_StakingNew.CallOpts)
}

// LaunchQueueDone is a free data retrieval call binding the contract method 0x8717ba17.
//
// Solidity: function launchQueueDone() view returns(uint256)
func (_StakingNew *StakingNewCallerSession) LaunchQueueDone() (*big.Int, error) {
	return _StakingNew.Contract.LaunchQueueDone(&_StakingNew.CallOpts)
}

// LevelRewardUSDT is a free data retrieval call binding the contract method 0x78f2ea10.
//
// Solidity: function levelRewardUSDT(address ) view returns(uint256)
func (_StakingNew *StakingNewCaller) LevelRewardUSDT(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "levelRewardUSDT", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LevelRewardUSDT is a free data retrieval call binding the contract method 0x78f2ea10.
//
// Solidity: function levelRewardUSDT(address ) view returns(uint256)
func (_StakingNew *StakingNewSession) LevelRewardUSDT(arg0 common.Address) (*big.Int, error) {
	return _StakingNew.Contract.LevelRewardUSDT(&_StakingNew.CallOpts, arg0)
}

// LevelRewardUSDT is a free data retrieval call binding the contract method 0x78f2ea10.
//
// Solidity: function levelRewardUSDT(address ) view returns(uint256)
func (_StakingNew *StakingNewCallerSession) LevelRewardUSDT(arg0 common.Address) (*big.Int, error) {
	return _StakingNew.Contract.LevelRewardUSDT(&_StakingNew.CallOpts, arg0)
}

// LineClaimable is a free data retrieval call binding the contract method 0x1b0cac53.
//
// Solidity: function lineClaimable(address account, uint256 index) view returns(uint256 amount)
func (_StakingNew *StakingNewCaller) LineClaimable(opts *bind.CallOpts, account common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "lineClaimable", account, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LineClaimable is a free data retrieval call binding the contract method 0x1b0cac53.
//
// Solidity: function lineClaimable(address account, uint256 index) view returns(uint256 amount)
func (_StakingNew *StakingNewSession) LineClaimable(account common.Address, index *big.Int) (*big.Int, error) {
	return _StakingNew.Contract.LineClaimable(&_StakingNew.CallOpts, account, index)
}

// LineClaimable is a free data retrieval call binding the contract method 0x1b0cac53.
//
// Solidity: function lineClaimable(address account, uint256 index) view returns(uint256 amount)
func (_StakingNew *StakingNewCallerSession) LineClaimable(account common.Address, index *big.Int) (*big.Int, error) {
	return _StakingNew.Contract.LineClaimable(&_StakingNew.CallOpts, account, index)
}

// LineClaimableDetail is a free data retrieval call binding the contract method 0xbf6698f1.
//
// Solidity: function lineClaimableDetail(address account, uint256 index) view returns(uint256 amount, uint40 effective)
func (_StakingNew *StakingNewCaller) LineClaimableDetail(opts *bind.CallOpts, account common.Address, index *big.Int) (struct {
	Amount    *big.Int
	Effective *big.Int
}, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "lineClaimableDetail", account, index)

	outstruct := new(struct {
		Amount    *big.Int
		Effective *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Effective = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LineClaimableDetail is a free data retrieval call binding the contract method 0xbf6698f1.
//
// Solidity: function lineClaimableDetail(address account, uint256 index) view returns(uint256 amount, uint40 effective)
func (_StakingNew *StakingNewSession) LineClaimableDetail(account common.Address, index *big.Int) (struct {
	Amount    *big.Int
	Effective *big.Int
}, error) {
	return _StakingNew.Contract.LineClaimableDetail(&_StakingNew.CallOpts, account, index)
}

// LineClaimableDetail is a free data retrieval call binding the contract method 0xbf6698f1.
//
// Solidity: function lineClaimableDetail(address account, uint256 index) view returns(uint256 amount, uint40 effective)
func (_StakingNew *StakingNewCallerSession) LineClaimableDetail(account common.Address, index *big.Int) (struct {
	Amount    *big.Int
	Effective *big.Int
}, error) {
	return _StakingNew.Contract.LineClaimableDetail(&_StakingNew.CallOpts, account, index)
}

// LineEcoBps is a free data retrieval call binding the contract method 0x0c37ecb3.
//
// Solidity: function lineEcoBps() view returns(uint16)
func (_StakingNew *StakingNewCaller) LineEcoBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "lineEcoBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LineEcoBps is a free data retrieval call binding the contract method 0x0c37ecb3.
//
// Solidity: function lineEcoBps() view returns(uint16)
func (_StakingNew *StakingNewSession) LineEcoBps() (uint16, error) {
	return _StakingNew.Contract.LineEcoBps(&_StakingNew.CallOpts)
}

// LineEcoBps is a free data retrieval call binding the contract method 0x0c37ecb3.
//
// Solidity: function lineEcoBps() view returns(uint16)
func (_StakingNew *StakingNewCallerSession) LineEcoBps() (uint16, error) {
	return _StakingNew.Contract.LineEcoBps(&_StakingNew.CallOpts)
}

// LineFeeBps is a free data retrieval call binding the contract method 0x83f0ef8e.
//
// Solidity: function lineFeeBps() view returns(uint16)
func (_StakingNew *StakingNewCaller) LineFeeBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "lineFeeBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LineFeeBps is a free data retrieval call binding the contract method 0x83f0ef8e.
//
// Solidity: function lineFeeBps() view returns(uint16)
func (_StakingNew *StakingNewSession) LineFeeBps() (uint16, error) {
	return _StakingNew.Contract.LineFeeBps(&_StakingNew.CallOpts)
}

// LineFeeBps is a free data retrieval call binding the contract method 0x83f0ef8e.
//
// Solidity: function lineFeeBps() view returns(uint16)
func (_StakingNew *StakingNewCallerSession) LineFeeBps() (uint16, error) {
	return _StakingNew.Contract.LineFeeBps(&_StakingNew.CallOpts)
}

// LineLevelBps is a free data retrieval call binding the contract method 0x27a8c765.
//
// Solidity: function lineLevelBps(uint256 ) view returns(uint16)
func (_StakingNew *StakingNewCaller) LineLevelBps(opts *bind.CallOpts, arg0 *big.Int) (uint16, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "lineLevelBps", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LineLevelBps is a free data retrieval call binding the contract method 0x27a8c765.
//
// Solidity: function lineLevelBps(uint256 ) view returns(uint16)
func (_StakingNew *StakingNewSession) LineLevelBps(arg0 *big.Int) (uint16, error) {
	return _StakingNew.Contract.LineLevelBps(&_StakingNew.CallOpts, arg0)
}

// LineLevelBps is a free data retrieval call binding the contract method 0x27a8c765.
//
// Solidity: function lineLevelBps(uint256 ) view returns(uint16)
func (_StakingNew *StakingNewCallerSession) LineLevelBps(arg0 *big.Int) (uint16, error) {
	return _StakingNew.Contract.LineLevelBps(&_StakingNew.CallOpts, arg0)
}

// LineSafeBefore is a free data retrieval call binding the contract method 0xdef289b7.
//
// Solidity: function lineSafeBefore() view returns(uint40)
func (_StakingNew *StakingNewCaller) LineSafeBefore(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "lineSafeBefore")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LineSafeBefore is a free data retrieval call binding the contract method 0xdef289b7.
//
// Solidity: function lineSafeBefore() view returns(uint40)
func (_StakingNew *StakingNewSession) LineSafeBefore() (*big.Int, error) {
	return _StakingNew.Contract.LineSafeBefore(&_StakingNew.CallOpts)
}

// LineSafeBefore is a free data retrieval call binding the contract method 0xdef289b7.
//
// Solidity: function lineSafeBefore() view returns(uint40)
func (_StakingNew *StakingNewCallerSession) LineSafeBefore() (*big.Int, error) {
	return _StakingNew.Contract.LineSafeBefore(&_StakingNew.CallOpts)
}

// LineSafeSec is a free data retrieval call binding the contract method 0x6b526833.
//
// Solidity: function lineSafeSec() view returns(uint40)
func (_StakingNew *StakingNewCaller) LineSafeSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "lineSafeSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LineSafeSec is a free data retrieval call binding the contract method 0x6b526833.
//
// Solidity: function lineSafeSec() view returns(uint40)
func (_StakingNew *StakingNewSession) LineSafeSec() (*big.Int, error) {
	return _StakingNew.Contract.LineSafeSec(&_StakingNew.CallOpts)
}

// LineSafeSec is a free data retrieval call binding the contract method 0x6b526833.
//
// Solidity: function lineSafeSec() view returns(uint40)
func (_StakingNew *StakingNewCallerSession) LineSafeSec() (*big.Int, error) {
	return _StakingNew.Contract.LineSafeSec(&_StakingNew.CallOpts)
}

// LiqBps is a free data retrieval call binding the contract method 0xbcbc3c7e.
//
// Solidity: function liqBps() view returns(uint16)
func (_StakingNew *StakingNewCaller) LiqBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "liqBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LiqBps is a free data retrieval call binding the contract method 0xbcbc3c7e.
//
// Solidity: function liqBps() view returns(uint16)
func (_StakingNew *StakingNewSession) LiqBps() (uint16, error) {
	return _StakingNew.Contract.LiqBps(&_StakingNew.CallOpts)
}

// LiqBps is a free data retrieval call binding the contract method 0xbcbc3c7e.
//
// Solidity: function liqBps() view returns(uint16)
func (_StakingNew *StakingNewCallerSession) LiqBps() (uint16, error) {
	return _StakingNew.Contract.LiqBps(&_StakingNew.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address ) view returns(bool)
func (_StakingNew *StakingNewCaller) Locked(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "locked", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address ) view returns(bool)
func (_StakingNew *StakingNewSession) Locked(arg0 common.Address) (bool, error) {
	return _StakingNew.Contract.Locked(&_StakingNew.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address ) view returns(bool)
func (_StakingNew *StakingNewCallerSession) Locked(arg0 common.Address) (bool, error) {
	return _StakingNew.Contract.Locked(&_StakingNew.CallOpts, arg0)
}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_StakingNew *StakingNewCaller) Market(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "market")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_StakingNew *StakingNewSession) Market() (common.Address, error) {
	return _StakingNew.Contract.Market(&_StakingNew.CallOpts)
}

// Market is a free data retrieval call binding the contract method 0x80f55605.
//
// Solidity: function market() view returns(address)
func (_StakingNew *StakingNewCallerSession) Market() (common.Address, error) {
	return _StakingNew.Contract.Market(&_StakingNew.CallOpts)
}

// MarketBps is a free data retrieval call binding the contract method 0x7eaa3107.
//
// Solidity: function marketBps() view returns(uint16)
func (_StakingNew *StakingNewCaller) MarketBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "marketBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MarketBps is a free data retrieval call binding the contract method 0x7eaa3107.
//
// Solidity: function marketBps() view returns(uint16)
func (_StakingNew *StakingNewSession) MarketBps() (uint16, error) {
	return _StakingNew.Contract.MarketBps(&_StakingNew.CallOpts)
}

// MarketBps is a free data retrieval call binding the contract method 0x7eaa3107.
//
// Solidity: function marketBps() view returns(uint16)
func (_StakingNew *StakingNewCallerSession) MarketBps() (uint16, error) {
	return _StakingNew.Contract.MarketBps(&_StakingNew.CallOpts)
}

// NewHighTime is a free data retrieval call binding the contract method 0xdb872c3b.
//
// Solidity: function newHighTime() view returns(uint40)
func (_StakingNew *StakingNewCaller) NewHighTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "newHighTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewHighTime is a free data retrieval call binding the contract method 0xdb872c3b.
//
// Solidity: function newHighTime() view returns(uint40)
func (_StakingNew *StakingNewSession) NewHighTime() (*big.Int, error) {
	return _StakingNew.Contract.NewHighTime(&_StakingNew.CallOpts)
}

// NewHighTime is a free data retrieval call binding the contract method 0xdb872c3b.
//
// Solidity: function newHighTime() view returns(uint40)
func (_StakingNew *StakingNewCallerSession) NewHighTime() (*big.Int, error) {
	return _StakingNew.Contract.NewHighTime(&_StakingNew.CallOpts)
}

// NextOrderId is a free data retrieval call binding the contract method 0x2a58b330.
//
// Solidity: function nextOrderId() view returns(uint256)
func (_StakingNew *StakingNewCaller) NextOrderId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "nextOrderId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextOrderId is a free data retrieval call binding the contract method 0x2a58b330.
//
// Solidity: function nextOrderId() view returns(uint256)
func (_StakingNew *StakingNewSession) NextOrderId() (*big.Int, error) {
	return _StakingNew.Contract.NextOrderId(&_StakingNew.CallOpts)
}

// NextOrderId is a free data retrieval call binding the contract method 0x2a58b330.
//
// Solidity: function nextOrderId() view returns(uint256)
func (_StakingNew *StakingNewCallerSession) NextOrderId() (*big.Int, error) {
	return _StakingNew.Contract.NextOrderId(&_StakingNew.CallOpts)
}

// OrderCount is a free data retrieval call binding the contract method 0x713d856b.
//
// Solidity: function orderCount(address account) view returns(uint256)
func (_StakingNew *StakingNewCaller) OrderCount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "orderCount", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderCount is a free data retrieval call binding the contract method 0x713d856b.
//
// Solidity: function orderCount(address account) view returns(uint256)
func (_StakingNew *StakingNewSession) OrderCount(account common.Address) (*big.Int, error) {
	return _StakingNew.Contract.OrderCount(&_StakingNew.CallOpts, account)
}

// OrderCount is a free data retrieval call binding the contract method 0x713d856b.
//
// Solidity: function orderCount(address account) view returns(uint256)
func (_StakingNew *StakingNewCallerSession) OrderCount(account common.Address) (*big.Int, error) {
	return _StakingNew.Contract.OrderCount(&_StakingNew.CallOpts, account)
}

// OrderIdAt is a free data retrieval call binding the contract method 0x6decf45a.
//
// Solidity: function orderIdAt(address account, uint256 index) view returns(uint256)
func (_StakingNew *StakingNewCaller) OrderIdAt(opts *bind.CallOpts, account common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "orderIdAt", account, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderIdAt is a free data retrieval call binding the contract method 0x6decf45a.
//
// Solidity: function orderIdAt(address account, uint256 index) view returns(uint256)
func (_StakingNew *StakingNewSession) OrderIdAt(account common.Address, index *big.Int) (*big.Int, error) {
	return _StakingNew.Contract.OrderIdAt(&_StakingNew.CallOpts, account, index)
}

// OrderIdAt is a free data retrieval call binding the contract method 0x6decf45a.
//
// Solidity: function orderIdAt(address account, uint256 index) view returns(uint256)
func (_StakingNew *StakingNewCallerSession) OrderIdAt(account common.Address, index *big.Int) (*big.Int, error) {
	return _StakingNew.Contract.OrderIdAt(&_StakingNew.CallOpts, account, index)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address account, uint128 amount, uint128 cap, uint128 used, uint128 linePaid, uint40 created, uint40 start, uint40 claimEffective, uint32 daysCount, bool exited)
func (_StakingNew *StakingNewCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Account        common.Address
	Amount         *big.Int
	Cap            *big.Int
	Used           *big.Int
	LinePaid       *big.Int
	Created        *big.Int
	Start          *big.Int
	ClaimEffective *big.Int
	DaysCount      uint32
	Exited         bool
}, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		Account        common.Address
		Amount         *big.Int
		Cap            *big.Int
		Used           *big.Int
		LinePaid       *big.Int
		Created        *big.Int
		Start          *big.Int
		ClaimEffective *big.Int
		DaysCount      uint32
		Exited         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Account = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Cap = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Used = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LinePaid = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Created = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Start = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.ClaimEffective = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.DaysCount = *abi.ConvertType(out[8], new(uint32)).(*uint32)
	outstruct.Exited = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address account, uint128 amount, uint128 cap, uint128 used, uint128 linePaid, uint40 created, uint40 start, uint40 claimEffective, uint32 daysCount, bool exited)
func (_StakingNew *StakingNewSession) Orders(arg0 *big.Int) (struct {
	Account        common.Address
	Amount         *big.Int
	Cap            *big.Int
	Used           *big.Int
	LinePaid       *big.Int
	Created        *big.Int
	Start          *big.Int
	ClaimEffective *big.Int
	DaysCount      uint32
	Exited         bool
}, error) {
	return _StakingNew.Contract.Orders(&_StakingNew.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address account, uint128 amount, uint128 cap, uint128 used, uint128 linePaid, uint40 created, uint40 start, uint40 claimEffective, uint32 daysCount, bool exited)
func (_StakingNew *StakingNewCallerSession) Orders(arg0 *big.Int) (struct {
	Account        common.Address
	Amount         *big.Int
	Cap            *big.Int
	Used           *big.Int
	LinePaid       *big.Int
	Created        *big.Int
	Start          *big.Int
	ClaimEffective *big.Int
	DaysCount      uint32
	Exited         bool
}, error) {
	return _StakingNew.Contract.Orders(&_StakingNew.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingNew *StakingNewCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingNew *StakingNewSession) Owner() (common.Address, error) {
	return _StakingNew.Contract.Owner(&_StakingNew.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingNew *StakingNewCallerSession) Owner() (common.Address, error) {
	return _StakingNew.Contract.Owner(&_StakingNew.CallOpts)
}

// PlanLength is a free data retrieval call binding the contract method 0x153a6b0c.
//
// Solidity: function planLength() view returns(uint256)
func (_StakingNew *StakingNewCaller) PlanLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "planLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlanLength is a free data retrieval call binding the contract method 0x153a6b0c.
//
// Solidity: function planLength() view returns(uint256)
func (_StakingNew *StakingNewSession) PlanLength() (*big.Int, error) {
	return _StakingNew.Contract.PlanLength(&_StakingNew.CallOpts)
}

// PlanLength is a free data retrieval call binding the contract method 0x153a6b0c.
//
// Solidity: function planLength() view returns(uint256)
func (_StakingNew *StakingNewCallerSession) PlanLength() (*big.Int, error) {
	return _StakingNew.Contract.PlanLength(&_StakingNew.CallOpts)
}

// Plans is a free data retrieval call binding the contract method 0xb1620616.
//
// Solidity: function plans(uint256 ) view returns(uint128 minAmount, uint128 maxAmount, uint128 outAmount, uint32 daysCount, bool enabled)
func (_StakingNew *StakingNewCaller) Plans(opts *bind.CallOpts, arg0 *big.Int) (struct {
	MinAmount *big.Int
	MaxAmount *big.Int
	OutAmount *big.Int
	DaysCount uint32
	Enabled   bool
}, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "plans", arg0)

	outstruct := new(struct {
		MinAmount *big.Int
		MaxAmount *big.Int
		OutAmount *big.Int
		DaysCount uint32
		Enabled   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.OutAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DaysCount = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.Enabled = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Plans is a free data retrieval call binding the contract method 0xb1620616.
//
// Solidity: function plans(uint256 ) view returns(uint128 minAmount, uint128 maxAmount, uint128 outAmount, uint32 daysCount, bool enabled)
func (_StakingNew *StakingNewSession) Plans(arg0 *big.Int) (struct {
	MinAmount *big.Int
	MaxAmount *big.Int
	OutAmount *big.Int
	DaysCount uint32
	Enabled   bool
}, error) {
	return _StakingNew.Contract.Plans(&_StakingNew.CallOpts, arg0)
}

// Plans is a free data retrieval call binding the contract method 0xb1620616.
//
// Solidity: function plans(uint256 ) view returns(uint128 minAmount, uint128 maxAmount, uint128 outAmount, uint32 daysCount, bool enabled)
func (_StakingNew *StakingNewCallerSession) Plans(arg0 *big.Int) (struct {
	MinAmount *big.Int
	MaxAmount *big.Int
	OutAmount *big.Int
	DaysCount uint32
	Enabled   bool
}, error) {
	return _StakingNew.Contract.Plans(&_StakingNew.CallOpts, arg0)
}

// Queue is a free data retrieval call binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 ) view returns(uint256 orderId, uint128 liqU, uint40 queuedAt)
func (_StakingNew *StakingNewCaller) Queue(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OrderId  *big.Int
	LiqU     *big.Int
	QueuedAt *big.Int
}, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "queue", arg0)

	outstruct := new(struct {
		OrderId  *big.Int
		LiqU     *big.Int
		QueuedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiqU = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.QueuedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Queue is a free data retrieval call binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 ) view returns(uint256 orderId, uint128 liqU, uint40 queuedAt)
func (_StakingNew *StakingNewSession) Queue(arg0 *big.Int) (struct {
	OrderId  *big.Int
	LiqU     *big.Int
	QueuedAt *big.Int
}, error) {
	return _StakingNew.Contract.Queue(&_StakingNew.CallOpts, arg0)
}

// Queue is a free data retrieval call binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 ) view returns(uint256 orderId, uint128 liqU, uint40 queuedAt)
func (_StakingNew *StakingNewCallerSession) Queue(arg0 *big.Int) (struct {
	OrderId  *big.Int
	LiqU     *big.Int
	QueuedAt *big.Int
}, error) {
	return _StakingNew.Contract.Queue(&_StakingNew.CallOpts, arg0)
}

// QueueCount is a free data retrieval call binding the contract method 0x9ef36bd1.
//
// Solidity: function queueCount() view returns(uint256)
func (_StakingNew *StakingNewCaller) QueueCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "queueCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueueCount is a free data retrieval call binding the contract method 0x9ef36bd1.
//
// Solidity: function queueCount() view returns(uint256)
func (_StakingNew *StakingNewSession) QueueCount() (*big.Int, error) {
	return _StakingNew.Contract.QueueCount(&_StakingNew.CallOpts)
}

// QueueCount is a free data retrieval call binding the contract method 0x9ef36bd1.
//
// Solidity: function queueCount() view returns(uint256)
func (_StakingNew *StakingNewCallerSession) QueueCount() (*big.Int, error) {
	return _StakingNew.Contract.QueueCount(&_StakingNew.CallOpts)
}

// QueueHead is a free data retrieval call binding the contract method 0x96523644.
//
// Solidity: function queueHead() view returns(uint256)
func (_StakingNew *StakingNewCaller) QueueHead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "queueHead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueueHead is a free data retrieval call binding the contract method 0x96523644.
//
// Solidity: function queueHead() view returns(uint256)
func (_StakingNew *StakingNewSession) QueueHead() (*big.Int, error) {
	return _StakingNew.Contract.QueueHead(&_StakingNew.CallOpts)
}

// QueueHead is a free data retrieval call binding the contract method 0x96523644.
//
// Solidity: function queueHead() view returns(uint256)
func (_StakingNew *StakingNewCallerSession) QueueHead() (*big.Int, error) {
	return _StakingNew.Contract.QueueHead(&_StakingNew.CallOpts)
}

// QueueLength is a free data retrieval call binding the contract method 0xab91c7b0.
//
// Solidity: function queueLength() view returns(uint256)
func (_StakingNew *StakingNewCaller) QueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "queueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueueLength is a free data retrieval call binding the contract method 0xab91c7b0.
//
// Solidity: function queueLength() view returns(uint256)
func (_StakingNew *StakingNewSession) QueueLength() (*big.Int, error) {
	return _StakingNew.Contract.QueueLength(&_StakingNew.CallOpts)
}

// QueueLength is a free data retrieval call binding the contract method 0xab91c7b0.
//
// Solidity: function queueLength() view returns(uint256)
func (_StakingNew *StakingNewCallerSession) QueueLength() (*big.Int, error) {
	return _StakingNew.Contract.QueueLength(&_StakingNew.CallOpts)
}

// QueueTail is a free data retrieval call binding the contract method 0xcbafd3c2.
//
// Solidity: function queueTail() view returns(uint256)
func (_StakingNew *StakingNewCaller) QueueTail(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "queueTail")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueueTail is a free data retrieval call binding the contract method 0xcbafd3c2.
//
// Solidity: function queueTail() view returns(uint256)
func (_StakingNew *StakingNewSession) QueueTail() (*big.Int, error) {
	return _StakingNew.Contract.QueueTail(&_StakingNew.CallOpts)
}

// QueueTail is a free data retrieval call binding the contract method 0xcbafd3c2.
//
// Solidity: function queueTail() view returns(uint256)
func (_StakingNew *StakingNewCallerSession) QueueTail() (*big.Int, error) {
	return _StakingNew.Contract.QueueTail(&_StakingNew.CallOpts)
}

// QueueWait is a free data retrieval call binding the contract method 0x281d6501.
//
// Solidity: function queueWait() view returns(uint40)
func (_StakingNew *StakingNewCaller) QueueWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "queueWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueueWait is a free data retrieval call binding the contract method 0x281d6501.
//
// Solidity: function queueWait() view returns(uint40)
func (_StakingNew *StakingNewSession) QueueWait() (*big.Int, error) {
	return _StakingNew.Contract.QueueWait(&_StakingNew.CallOpts)
}

// QueueWait is a free data retrieval call binding the contract method 0x281d6501.
//
// Solidity: function queueWait() view returns(uint40)
func (_StakingNew *StakingNewCallerSession) QueueWait() (*big.Int, error) {
	return _StakingNew.Contract.QueueWait(&_StakingNew.CallOpts)
}

// RList is a free data retrieval call binding the contract method 0x54d3cd87.
//
// Solidity: function rList(uint256 index) view returns(uint40 start, uint40 end)
func (_StakingNew *StakingNewCaller) RList(opts *bind.CallOpts, index *big.Int) (struct {
	Start *big.Int
	End   *big.Int
}, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "rList", index)

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
// Solidity: function rList(uint256 index) view returns(uint40 start, uint40 end)
func (_StakingNew *StakingNewSession) RList(index *big.Int) (struct {
	Start *big.Int
	End   *big.Int
}, error) {
	return _StakingNew.Contract.RList(&_StakingNew.CallOpts, index)
}

// RList is a free data retrieval call binding the contract method 0x54d3cd87.
//
// Solidity: function rList(uint256 index) view returns(uint40 start, uint40 end)
func (_StakingNew *StakingNewCallerSession) RList(index *big.Int) (struct {
	Start *big.Int
	End   *big.Int
}, error) {
	return _StakingNew.Contract.RList(&_StakingNew.CallOpts, index)
}

// RListLength is a free data retrieval call binding the contract method 0x217959fc.
//
// Solidity: function rListLength() view returns(uint256)
func (_StakingNew *StakingNewCaller) RListLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "rListLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RListLength is a free data retrieval call binding the contract method 0x217959fc.
//
// Solidity: function rListLength() view returns(uint256)
func (_StakingNew *StakingNewSession) RListLength() (*big.Int, error) {
	return _StakingNew.Contract.RListLength(&_StakingNew.CallOpts)
}

// RListLength is a free data retrieval call binding the contract method 0x217959fc.
//
// Solidity: function rListLength() view returns(uint256)
func (_StakingNew *StakingNewCallerSession) RListLength() (*big.Int, error) {
	return _StakingNew.Contract.RListLength(&_StakingNew.CallOpts)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(uint40 start, uint40 end)
func (_StakingNew *StakingNewCaller) Rounds(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Start *big.Int
	End   *big.Int
}, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "rounds", arg0)

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

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(uint40 start, uint40 end)
func (_StakingNew *StakingNewSession) Rounds(arg0 *big.Int) (struct {
	Start *big.Int
	End   *big.Int
}, error) {
	return _StakingNew.Contract.Rounds(&_StakingNew.CallOpts, arg0)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(uint40 start, uint40 end)
func (_StakingNew *StakingNewCallerSession) Rounds(arg0 *big.Int) (struct {
	Start *big.Int
	End   *big.Int
}, error) {
	return _StakingNew.Contract.Rounds(&_StakingNew.CallOpts, arg0)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_StakingNew *StakingNewCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_StakingNew *StakingNewSession) Router() (common.Address, error) {
	return _StakingNew.Contract.Router(&_StakingNew.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_StakingNew *StakingNewCallerSession) Router() (common.Address, error) {
	return _StakingNew.Contract.Router(&_StakingNew.CallOpts)
}

// ShouldUpdateCircuitBreaker is a free data retrieval call binding the contract method 0xdcac1a16.
//
// Solidity: function shouldUpdateCircuitBreaker() view returns(bool needTrigger, bool needCountdown, bool needRecover, uint256 currentDropBps)
func (_StakingNew *StakingNewCaller) ShouldUpdateCircuitBreaker(opts *bind.CallOpts) (struct {
	NeedTrigger    bool
	NeedCountdown  bool
	NeedRecover    bool
	CurrentDropBps *big.Int
}, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "shouldUpdateCircuitBreaker")

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
func (_StakingNew *StakingNewSession) ShouldUpdateCircuitBreaker() (struct {
	NeedTrigger    bool
	NeedCountdown  bool
	NeedRecover    bool
	CurrentDropBps *big.Int
}, error) {
	return _StakingNew.Contract.ShouldUpdateCircuitBreaker(&_StakingNew.CallOpts)
}

// ShouldUpdateCircuitBreaker is a free data retrieval call binding the contract method 0xdcac1a16.
//
// Solidity: function shouldUpdateCircuitBreaker() view returns(bool needTrigger, bool needCountdown, bool needRecover, uint256 currentDropBps)
func (_StakingNew *StakingNewCallerSession) ShouldUpdateCircuitBreaker() (struct {
	NeedTrigger    bool
	NeedCountdown  bool
	NeedRecover    bool
	CurrentDropBps *big.Int
}, error) {
	return _StakingNew.Contract.ShouldUpdateCircuitBreaker(&_StakingNew.CallOpts)
}

// Sink is a free data retrieval call binding the contract method 0xc74e820e.
//
// Solidity: function sink() view returns(address)
func (_StakingNew *StakingNewCaller) Sink(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "sink")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sink is a free data retrieval call binding the contract method 0xc74e820e.
//
// Solidity: function sink() view returns(address)
func (_StakingNew *StakingNewSession) Sink() (common.Address, error) {
	return _StakingNew.Contract.Sink(&_StakingNew.CallOpts)
}

// Sink is a free data retrieval call binding the contract method 0xc74e820e.
//
// Solidity: function sink() view returns(address)
func (_StakingNew *StakingNewCallerSession) Sink() (common.Address, error) {
	return _StakingNew.Contract.Sink(&_StakingNew.CallOpts)
}

// SmallFeeBps is a free data retrieval call binding the contract method 0x8e901aa7.
//
// Solidity: function smallFeeBps() view returns(uint16)
func (_StakingNew *StakingNewCaller) SmallFeeBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "smallFeeBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// SmallFeeBps is a free data retrieval call binding the contract method 0x8e901aa7.
//
// Solidity: function smallFeeBps() view returns(uint16)
func (_StakingNew *StakingNewSession) SmallFeeBps() (uint16, error) {
	return _StakingNew.Contract.SmallFeeBps(&_StakingNew.CallOpts)
}

// SmallFeeBps is a free data retrieval call binding the contract method 0x8e901aa7.
//
// Solidity: function smallFeeBps() view returns(uint16)
func (_StakingNew *StakingNewCallerSession) SmallFeeBps() (uint16, error) {
	return _StakingNew.Contract.SmallFeeBps(&_StakingNew.CallOpts)
}

// TeamClaimedTotal is a free data retrieval call binding the contract method 0x2b44eb10.
//
// Solidity: function teamClaimedTotal(address ) view returns(uint256)
func (_StakingNew *StakingNewCaller) TeamClaimedTotal(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "teamClaimedTotal", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TeamClaimedTotal is a free data retrieval call binding the contract method 0x2b44eb10.
//
// Solidity: function teamClaimedTotal(address ) view returns(uint256)
func (_StakingNew *StakingNewSession) TeamClaimedTotal(arg0 common.Address) (*big.Int, error) {
	return _StakingNew.Contract.TeamClaimedTotal(&_StakingNew.CallOpts, arg0)
}

// TeamClaimedTotal is a free data retrieval call binding the contract method 0x2b44eb10.
//
// Solidity: function teamClaimedTotal(address ) view returns(uint256)
func (_StakingNew *StakingNewCallerSession) TeamClaimedTotal(arg0 common.Address) (*big.Int, error) {
	return _StakingNew.Contract.TeamClaimedTotal(&_StakingNew.CallOpts, arg0)
}

// TeamClearAt is a free data retrieval call binding the contract method 0xd2a2ec11.
//
// Solidity: function teamClearAt(address ) view returns(uint40)
func (_StakingNew *StakingNewCaller) TeamClearAt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "teamClearAt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TeamClearAt is a free data retrieval call binding the contract method 0xd2a2ec11.
//
// Solidity: function teamClearAt(address ) view returns(uint40)
func (_StakingNew *StakingNewSession) TeamClearAt(arg0 common.Address) (*big.Int, error) {
	return _StakingNew.Contract.TeamClearAt(&_StakingNew.CallOpts, arg0)
}

// TeamClearAt is a free data retrieval call binding the contract method 0xd2a2ec11.
//
// Solidity: function teamClearAt(address ) view returns(uint40)
func (_StakingNew *StakingNewCallerSession) TeamClearAt(arg0 common.Address) (*big.Int, error) {
	return _StakingNew.Contract.TeamClearAt(&_StakingNew.CallOpts, arg0)
}

// TeamClearPeriod is a free data retrieval call binding the contract method 0xb512f7c3.
//
// Solidity: function teamClearPeriod() view returns(uint40)
func (_StakingNew *StakingNewCaller) TeamClearPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "teamClearPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TeamClearPeriod is a free data retrieval call binding the contract method 0xb512f7c3.
//
// Solidity: function teamClearPeriod() view returns(uint40)
func (_StakingNew *StakingNewSession) TeamClearPeriod() (*big.Int, error) {
	return _StakingNew.Contract.TeamClearPeriod(&_StakingNew.CallOpts)
}

// TeamClearPeriod is a free data retrieval call binding the contract method 0xb512f7c3.
//
// Solidity: function teamClearPeriod() view returns(uint40)
func (_StakingNew *StakingNewCallerSession) TeamClearPeriod() (*big.Int, error) {
	return _StakingNew.Contract.TeamClearPeriod(&_StakingNew.CallOpts)
}

// TeamDirectBps is a free data retrieval call binding the contract method 0x81260575.
//
// Solidity: function teamDirectBps() view returns(uint16)
func (_StakingNew *StakingNewCaller) TeamDirectBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "teamDirectBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TeamDirectBps is a free data retrieval call binding the contract method 0x81260575.
//
// Solidity: function teamDirectBps() view returns(uint16)
func (_StakingNew *StakingNewSession) TeamDirectBps() (uint16, error) {
	return _StakingNew.Contract.TeamDirectBps(&_StakingNew.CallOpts)
}

// TeamDirectBps is a free data retrieval call binding the contract method 0x81260575.
//
// Solidity: function teamDirectBps() view returns(uint16)
func (_StakingNew *StakingNewCallerSession) TeamDirectBps() (uint16, error) {
	return _StakingNew.Contract.TeamDirectBps(&_StakingNew.CallOpts)
}

// TeamMaxBps is a free data retrieval call binding the contract method 0x1c8777ec.
//
// Solidity: function teamMaxBps() view returns(uint16)
func (_StakingNew *StakingNewCaller) TeamMaxBps(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "teamMaxBps")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TeamMaxBps is a free data retrieval call binding the contract method 0x1c8777ec.
//
// Solidity: function teamMaxBps() view returns(uint16)
func (_StakingNew *StakingNewSession) TeamMaxBps() (uint16, error) {
	return _StakingNew.Contract.TeamMaxBps(&_StakingNew.CallOpts)
}

// TeamMaxBps is a free data retrieval call binding the contract method 0x1c8777ec.
//
// Solidity: function teamMaxBps() view returns(uint16)
func (_StakingNew *StakingNewCallerSession) TeamMaxBps() (uint16, error) {
	return _StakingNew.Contract.TeamMaxBps(&_StakingNew.CallOpts)
}

// TeamU is a free data retrieval call binding the contract method 0xe24ae116.
//
// Solidity: function teamU(address ) view returns(uint256)
func (_StakingNew *StakingNewCaller) TeamU(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "teamU", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TeamU is a free data retrieval call binding the contract method 0xe24ae116.
//
// Solidity: function teamU(address ) view returns(uint256)
func (_StakingNew *StakingNewSession) TeamU(arg0 common.Address) (*big.Int, error) {
	return _StakingNew.Contract.TeamU(&_StakingNew.CallOpts, arg0)
}

// TeamU is a free data retrieval call binding the contract method 0xe24ae116.
//
// Solidity: function teamU(address ) view returns(uint256)
func (_StakingNew *StakingNewCallerSession) TeamU(arg0 common.Address) (*big.Int, error) {
	return _StakingNew.Contract.TeamU(&_StakingNew.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StakingNew *StakingNewCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StakingNew *StakingNewSession) Token() (common.Address, error) {
	return _StakingNew.Contract.Token(&_StakingNew.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StakingNew *StakingNewCallerSession) Token() (common.Address, error) {
	return _StakingNew.Contract.Token(&_StakingNew.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_StakingNew *StakingNewCaller) Usdt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "usdt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_StakingNew *StakingNewSession) Usdt() (common.Address, error) {
	return _StakingNew.Contract.Usdt(&_StakingNew.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_StakingNew *StakingNewCallerSession) Usdt() (common.Address, error) {
	return _StakingNew.Contract.Usdt(&_StakingNew.CallOpts)
}

// User is a free data retrieval call binding the contract method 0x4f8632ba.
//
// Solidity: function user() view returns(address)
func (_StakingNew *StakingNewCaller) User(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingNew.contract.Call(opts, &out, "user")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// User is a free data retrieval call binding the contract method 0x4f8632ba.
//
// Solidity: function user() view returns(address)
func (_StakingNew *StakingNewSession) User() (common.Address, error) {
	return _StakingNew.Contract.User(&_StakingNew.CallOpts)
}

// User is a free data retrieval call binding the contract method 0x4f8632ba.
//
// Solidity: function user() view returns(address)
func (_StakingNew *StakingNewCallerSession) User() (common.Address, error) {
	return _StakingNew.Contract.User(&_StakingNew.CallOpts)
}

// ApproveAll is a paid mutator transaction binding the contract method 0x380d0c08.
//
// Solidity: function approveAll() returns()
func (_StakingNew *StakingNewTransactor) ApproveAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "approveAll")
}

// ApproveAll is a paid mutator transaction binding the contract method 0x380d0c08.
//
// Solidity: function approveAll() returns()
func (_StakingNew *StakingNewSession) ApproveAll() (*types.Transaction, error) {
	return _StakingNew.Contract.ApproveAll(&_StakingNew.TransactOpts)
}

// ApproveAll is a paid mutator transaction binding the contract method 0x380d0c08.
//
// Solidity: function approveAll() returns()
func (_StakingNew *StakingNewTransactorSession) ApproveAll() (*types.Transaction, error) {
	return _StakingNew.Contract.ApproveAll(&_StakingNew.TransactOpts)
}

// ClaimLine is a paid mutator transaction binding the contract method 0x5af3d596.
//
// Solidity: function claimLine(uint256 index) returns(uint256 grossU, bool paidMs)
func (_StakingNew *StakingNewTransactor) ClaimLine(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "claimLine", index)
}

// ClaimLine is a paid mutator transaction binding the contract method 0x5af3d596.
//
// Solidity: function claimLine(uint256 index) returns(uint256 grossU, bool paidMs)
func (_StakingNew *StakingNewSession) ClaimLine(index *big.Int) (*types.Transaction, error) {
	return _StakingNew.Contract.ClaimLine(&_StakingNew.TransactOpts, index)
}

// ClaimLine is a paid mutator transaction binding the contract method 0x5af3d596.
//
// Solidity: function claimLine(uint256 index) returns(uint256 grossU, bool paidMs)
func (_StakingNew *StakingNewTransactorSession) ClaimLine(index *big.Int) (*types.Transaction, error) {
	return _StakingNew.Contract.ClaimLine(&_StakingNew.TransactOpts, index)
}

// ClaimLineAll is a paid mutator transaction binding the contract method 0xc16629b0.
//
// Solidity: function claimLineAll(uint256 maxOrders) returns(uint256 totalU)
func (_StakingNew *StakingNewTransactor) ClaimLineAll(opts *bind.TransactOpts, maxOrders *big.Int) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "claimLineAll", maxOrders)
}

// ClaimLineAll is a paid mutator transaction binding the contract method 0xc16629b0.
//
// Solidity: function claimLineAll(uint256 maxOrders) returns(uint256 totalU)
func (_StakingNew *StakingNewSession) ClaimLineAll(maxOrders *big.Int) (*types.Transaction, error) {
	return _StakingNew.Contract.ClaimLineAll(&_StakingNew.TransactOpts, maxOrders)
}

// ClaimLineAll is a paid mutator transaction binding the contract method 0xc16629b0.
//
// Solidity: function claimLineAll(uint256 maxOrders) returns(uint256 totalU)
func (_StakingNew *StakingNewTransactorSession) ClaimLineAll(maxOrders *big.Int) (*types.Transaction, error) {
	return _StakingNew.Contract.ClaimLineAll(&_StakingNew.TransactOpts, maxOrders)
}

// ClaimLineBatch is a paid mutator transaction binding the contract method 0x20797d7b.
//
// Solidity: function claimLineBatch(uint256[] indexes) returns(uint256 totalU)
func (_StakingNew *StakingNewTransactor) ClaimLineBatch(opts *bind.TransactOpts, indexes []*big.Int) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "claimLineBatch", indexes)
}

// ClaimLineBatch is a paid mutator transaction binding the contract method 0x20797d7b.
//
// Solidity: function claimLineBatch(uint256[] indexes) returns(uint256 totalU)
func (_StakingNew *StakingNewSession) ClaimLineBatch(indexes []*big.Int) (*types.Transaction, error) {
	return _StakingNew.Contract.ClaimLineBatch(&_StakingNew.TransactOpts, indexes)
}

// ClaimLineBatch is a paid mutator transaction binding the contract method 0x20797d7b.
//
// Solidity: function claimLineBatch(uint256[] indexes) returns(uint256 totalU)
func (_StakingNew *StakingNewTransactorSession) ClaimLineBatch(indexes []*big.Int) (*types.Transaction, error) {
	return _StakingNew.Contract.ClaimLineBatch(&_StakingNew.TransactOpts, indexes)
}

// ClaimTeam is a paid mutator transaction binding the contract method 0xa9c8822a.
//
// Solidity: function claimTeam(uint256 amount) returns(uint256 net)
func (_StakingNew *StakingNewTransactor) ClaimTeam(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "claimTeam", amount)
}

// ClaimTeam is a paid mutator transaction binding the contract method 0xa9c8822a.
//
// Solidity: function claimTeam(uint256 amount) returns(uint256 net)
func (_StakingNew *StakingNewSession) ClaimTeam(amount *big.Int) (*types.Transaction, error) {
	return _StakingNew.Contract.ClaimTeam(&_StakingNew.TransactOpts, amount)
}

// ClaimTeam is a paid mutator transaction binding the contract method 0xa9c8822a.
//
// Solidity: function claimTeam(uint256 amount) returns(uint256 net)
func (_StakingNew *StakingNewTransactorSession) ClaimTeam(amount *big.Int) (*types.Transaction, error) {
	return _StakingNew.Contract.ClaimTeam(&_StakingNew.TransactOpts, amount)
}

// InitSystems is a paid mutator transaction binding the contract method 0x08535a28.
//
// Solidity: function initSystems(address token_, address usdt_, address router_, address user_) returns()
func (_StakingNew *StakingNewTransactor) InitSystems(opts *bind.TransactOpts, token_ common.Address, usdt_ common.Address, router_ common.Address, user_ common.Address) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "initSystems", token_, usdt_, router_, user_)
}

// InitSystems is a paid mutator transaction binding the contract method 0x08535a28.
//
// Solidity: function initSystems(address token_, address usdt_, address router_, address user_) returns()
func (_StakingNew *StakingNewSession) InitSystems(token_ common.Address, usdt_ common.Address, router_ common.Address, user_ common.Address) (*types.Transaction, error) {
	return _StakingNew.Contract.InitSystems(&_StakingNew.TransactOpts, token_, usdt_, router_, user_)
}

// InitSystems is a paid mutator transaction binding the contract method 0x08535a28.
//
// Solidity: function initSystems(address token_, address usdt_, address router_, address user_) returns()
func (_StakingNew *StakingNewTransactorSession) InitSystems(token_ common.Address, usdt_ common.Address, router_ common.Address, user_ common.Address) (*types.Transaction, error) {
	return _StakingNew.Contract.InitSystems(&_StakingNew.TransactOpts, token_, usdt_, router_, user_)
}

// ProcessQueue is a paid mutator transaction binding the contract method 0xdc900765.
//
// Solidity: function processQueue(uint256 maxItems) returns(uint256 done)
func (_StakingNew *StakingNewTransactor) ProcessQueue(opts *bind.TransactOpts, maxItems *big.Int) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "processQueue", maxItems)
}

// ProcessQueue is a paid mutator transaction binding the contract method 0xdc900765.
//
// Solidity: function processQueue(uint256 maxItems) returns(uint256 done)
func (_StakingNew *StakingNewSession) ProcessQueue(maxItems *big.Int) (*types.Transaction, error) {
	return _StakingNew.Contract.ProcessQueue(&_StakingNew.TransactOpts, maxItems)
}

// ProcessQueue is a paid mutator transaction binding the contract method 0xdc900765.
//
// Solidity: function processQueue(uint256 maxItems) returns(uint256 done)
func (_StakingNew *StakingNewTransactorSession) ProcessQueue(maxItems *big.Int) (*types.Transaction, error) {
	return _StakingNew.Contract.ProcessQueue(&_StakingNew.TransactOpts, maxItems)
}

// SetAntiDumpLinePayU is a paid mutator transaction binding the contract method 0x41373eba.
//
// Solidity: function setAntiDumpLinePayU(bool payU) returns()
func (_StakingNew *StakingNewTransactor) SetAntiDumpLinePayU(opts *bind.TransactOpts, payU bool) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "setAntiDumpLinePayU", payU)
}

// SetAntiDumpLinePayU is a paid mutator transaction binding the contract method 0x41373eba.
//
// Solidity: function setAntiDumpLinePayU(bool payU) returns()
func (_StakingNew *StakingNewSession) SetAntiDumpLinePayU(payU bool) (*types.Transaction, error) {
	return _StakingNew.Contract.SetAntiDumpLinePayU(&_StakingNew.TransactOpts, payU)
}

// SetAntiDumpLinePayU is a paid mutator transaction binding the contract method 0x41373eba.
//
// Solidity: function setAntiDumpLinePayU(bool payU) returns()
func (_StakingNew *StakingNewTransactorSession) SetAntiDumpLinePayU(payU bool) (*types.Transaction, error) {
	return _StakingNew.Contract.SetAntiDumpLinePayU(&_StakingNew.TransactOpts, payU)
}

// SetCircuitBreakerParams is a paid mutator transaction binding the contract method 0x43114be5.
//
// Solidity: function setCircuitBreakerParams(uint256 threshold, uint40 recovery, uint16 compBps) returns()
func (_StakingNew *StakingNewTransactor) SetCircuitBreakerParams(opts *bind.TransactOpts, threshold *big.Int, recovery *big.Int, compBps uint16) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "setCircuitBreakerParams", threshold, recovery, compBps)
}

// SetCircuitBreakerParams is a paid mutator transaction binding the contract method 0x43114be5.
//
// Solidity: function setCircuitBreakerParams(uint256 threshold, uint40 recovery, uint16 compBps) returns()
func (_StakingNew *StakingNewSession) SetCircuitBreakerParams(threshold *big.Int, recovery *big.Int, compBps uint16) (*types.Transaction, error) {
	return _StakingNew.Contract.SetCircuitBreakerParams(&_StakingNew.TransactOpts, threshold, recovery, compBps)
}

// SetCircuitBreakerParams is a paid mutator transaction binding the contract method 0x43114be5.
//
// Solidity: function setCircuitBreakerParams(uint256 threshold, uint40 recovery, uint16 compBps) returns()
func (_StakingNew *StakingNewTransactorSession) SetCircuitBreakerParams(threshold *big.Int, recovery *big.Int, compBps uint16) (*types.Transaction, error) {
	return _StakingNew.Contract.SetCircuitBreakerParams(&_StakingNew.TransactOpts, threshold, recovery, compBps)
}

// SetEntryFees is a paid mutator transaction binding the contract method 0xf62a7f77.
//
// Solidity: function setEntryFees(uint16 liq_, uint16 market_, uint16 eco_, uint16 team_, uint16 direct_, uint16 fee_) returns()
func (_StakingNew *StakingNewTransactor) SetEntryFees(opts *bind.TransactOpts, liq_ uint16, market_ uint16, eco_ uint16, team_ uint16, direct_ uint16, fee_ uint16) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "setEntryFees", liq_, market_, eco_, team_, direct_, fee_)
}

// SetEntryFees is a paid mutator transaction binding the contract method 0xf62a7f77.
//
// Solidity: function setEntryFees(uint16 liq_, uint16 market_, uint16 eco_, uint16 team_, uint16 direct_, uint16 fee_) returns()
func (_StakingNew *StakingNewSession) SetEntryFees(liq_ uint16, market_ uint16, eco_ uint16, team_ uint16, direct_ uint16, fee_ uint16) (*types.Transaction, error) {
	return _StakingNew.Contract.SetEntryFees(&_StakingNew.TransactOpts, liq_, market_, eco_, team_, direct_, fee_)
}

// SetEntryFees is a paid mutator transaction binding the contract method 0xf62a7f77.
//
// Solidity: function setEntryFees(uint16 liq_, uint16 market_, uint16 eco_, uint16 team_, uint16 direct_, uint16 fee_) returns()
func (_StakingNew *StakingNewTransactorSession) SetEntryFees(liq_ uint16, market_ uint16, eco_ uint16, team_ uint16, direct_ uint16, fee_ uint16) (*types.Transaction, error) {
	return _StakingNew.Contract.SetEntryFees(&_StakingNew.TransactOpts, liq_, market_, eco_, team_, direct_, fee_)
}

// SetInit is a paid mutator transaction binding the contract method 0x9784f55b.
//
// Solidity: function setInit(address newInit) returns()
func (_StakingNew *StakingNewTransactor) SetInit(opts *bind.TransactOpts, newInit common.Address) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "setInit", newInit)
}

// SetInit is a paid mutator transaction binding the contract method 0x9784f55b.
//
// Solidity: function setInit(address newInit) returns()
func (_StakingNew *StakingNewSession) SetInit(newInit common.Address) (*types.Transaction, error) {
	return _StakingNew.Contract.SetInit(&_StakingNew.TransactOpts, newInit)
}

// SetInit is a paid mutator transaction binding the contract method 0x9784f55b.
//
// Solidity: function setInit(address newInit) returns()
func (_StakingNew *StakingNewTransactorSession) SetInit(newInit common.Address) (*types.Transaction, error) {
	return _StakingNew.Contract.SetInit(&_StakingNew.TransactOpts, newInit)
}

// SetLineFees is a paid mutator transaction binding the contract method 0x3c53b655.
//
// Solidity: function setLineFees(uint16 fee_, uint16 eco_, uint16[4] lv_) returns()
func (_StakingNew *StakingNewTransactor) SetLineFees(opts *bind.TransactOpts, fee_ uint16, eco_ uint16, lv_ [4]uint16) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "setLineFees", fee_, eco_, lv_)
}

// SetLineFees is a paid mutator transaction binding the contract method 0x3c53b655.
//
// Solidity: function setLineFees(uint16 fee_, uint16 eco_, uint16[4] lv_) returns()
func (_StakingNew *StakingNewSession) SetLineFees(fee_ uint16, eco_ uint16, lv_ [4]uint16) (*types.Transaction, error) {
	return _StakingNew.Contract.SetLineFees(&_StakingNew.TransactOpts, fee_, eco_, lv_)
}

// SetLineFees is a paid mutator transaction binding the contract method 0x3c53b655.
//
// Solidity: function setLineFees(uint16 fee_, uint16 eco_, uint16[4] lv_) returns()
func (_StakingNew *StakingNewTransactorSession) SetLineFees(fee_ uint16, eco_ uint16, lv_ [4]uint16) (*types.Transaction, error) {
	return _StakingNew.Contract.SetLineFees(&_StakingNew.TransactOpts, fee_, eco_, lv_)
}

// SetLineSafe is a paid mutator transaction binding the contract method 0x92849f24.
//
// Solidity: function setLineSafe(uint40 secondsAmount, uint40 beforeTime) returns()
func (_StakingNew *StakingNewTransactor) SetLineSafe(opts *bind.TransactOpts, secondsAmount *big.Int, beforeTime *big.Int) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "setLineSafe", secondsAmount, beforeTime)
}

// SetLineSafe is a paid mutator transaction binding the contract method 0x92849f24.
//
// Solidity: function setLineSafe(uint40 secondsAmount, uint40 beforeTime) returns()
func (_StakingNew *StakingNewSession) SetLineSafe(secondsAmount *big.Int, beforeTime *big.Int) (*types.Transaction, error) {
	return _StakingNew.Contract.SetLineSafe(&_StakingNew.TransactOpts, secondsAmount, beforeTime)
}

// SetLineSafe is a paid mutator transaction binding the contract method 0x92849f24.
//
// Solidity: function setLineSafe(uint40 secondsAmount, uint40 beforeTime) returns()
func (_StakingNew *StakingNewTransactorSession) SetLineSafe(secondsAmount *big.Int, beforeTime *big.Int) (*types.Transaction, error) {
	return _StakingNew.Contract.SetLineSafe(&_StakingNew.TransactOpts, secondsAmount, beforeTime)
}

// SetLocked is a paid mutator transaction binding the contract method 0xd12f1786.
//
// Solidity: function setLocked(address[] accounts, bool[] values) returns()
func (_StakingNew *StakingNewTransactor) SetLocked(opts *bind.TransactOpts, accounts []common.Address, values []bool) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "setLocked", accounts, values)
}

// SetLocked is a paid mutator transaction binding the contract method 0xd12f1786.
//
// Solidity: function setLocked(address[] accounts, bool[] values) returns()
func (_StakingNew *StakingNewSession) SetLocked(accounts []common.Address, values []bool) (*types.Transaction, error) {
	return _StakingNew.Contract.SetLocked(&_StakingNew.TransactOpts, accounts, values)
}

// SetLocked is a paid mutator transaction binding the contract method 0xd12f1786.
//
// Solidity: function setLocked(address[] accounts, bool[] values) returns()
func (_StakingNew *StakingNewTransactorSession) SetLocked(accounts []common.Address, values []bool) (*types.Transaction, error) {
	return _StakingNew.Contract.SetLocked(&_StakingNew.TransactOpts, accounts, values)
}

// SetOrderCap is a paid mutator transaction binding the contract method 0x2b4996ac.
//
// Solidity: function setOrderCap(address account, uint256 index, uint128 newCap) returns()
func (_StakingNew *StakingNewTransactor) SetOrderCap(opts *bind.TransactOpts, account common.Address, index *big.Int, newCap *big.Int) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "setOrderCap", account, index, newCap)
}

// SetOrderCap is a paid mutator transaction binding the contract method 0x2b4996ac.
//
// Solidity: function setOrderCap(address account, uint256 index, uint128 newCap) returns()
func (_StakingNew *StakingNewSession) SetOrderCap(account common.Address, index *big.Int, newCap *big.Int) (*types.Transaction, error) {
	return _StakingNew.Contract.SetOrderCap(&_StakingNew.TransactOpts, account, index, newCap)
}

// SetOrderCap is a paid mutator transaction binding the contract method 0x2b4996ac.
//
// Solidity: function setOrderCap(address account, uint256 index, uint128 newCap) returns()
func (_StakingNew *StakingNewTransactorSession) SetOrderCap(account common.Address, index *big.Int, newCap *big.Int) (*types.Transaction, error) {
	return _StakingNew.Contract.SetOrderCap(&_StakingNew.TransactOpts, account, index, newCap)
}

// SetPlan is a paid mutator transaction binding the contract method 0x4515b85b.
//
// Solidity: function setPlan(uint256 index, uint128 minAmount, uint128 maxAmount, uint128 outAmount, uint32 daysCount, bool enabled) returns()
func (_StakingNew *StakingNewTransactor) SetPlan(opts *bind.TransactOpts, index *big.Int, minAmount *big.Int, maxAmount *big.Int, outAmount *big.Int, daysCount uint32, enabled bool) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "setPlan", index, minAmount, maxAmount, outAmount, daysCount, enabled)
}

// SetPlan is a paid mutator transaction binding the contract method 0x4515b85b.
//
// Solidity: function setPlan(uint256 index, uint128 minAmount, uint128 maxAmount, uint128 outAmount, uint32 daysCount, bool enabled) returns()
func (_StakingNew *StakingNewSession) SetPlan(index *big.Int, minAmount *big.Int, maxAmount *big.Int, outAmount *big.Int, daysCount uint32, enabled bool) (*types.Transaction, error) {
	return _StakingNew.Contract.SetPlan(&_StakingNew.TransactOpts, index, minAmount, maxAmount, outAmount, daysCount, enabled)
}

// SetPlan is a paid mutator transaction binding the contract method 0x4515b85b.
//
// Solidity: function setPlan(uint256 index, uint128 minAmount, uint128 maxAmount, uint128 outAmount, uint32 daysCount, bool enabled) returns()
func (_StakingNew *StakingNewTransactorSession) SetPlan(index *big.Int, minAmount *big.Int, maxAmount *big.Int, outAmount *big.Int, daysCount uint32, enabled bool) (*types.Transaction, error) {
	return _StakingNew.Contract.SetPlan(&_StakingNew.TransactOpts, index, minAmount, maxAmount, outAmount, daysCount, enabled)
}

// SetQueueConfig is a paid mutator transaction binding the contract method 0xfd7fd35c.
//
// Solidity: function setQueueConfig(uint40 wait_, uint256 autoCount_) returns()
func (_StakingNew *StakingNewTransactor) SetQueueConfig(opts *bind.TransactOpts, wait_ *big.Int, autoCount_ *big.Int) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "setQueueConfig", wait_, autoCount_)
}

// SetQueueConfig is a paid mutator transaction binding the contract method 0xfd7fd35c.
//
// Solidity: function setQueueConfig(uint40 wait_, uint256 autoCount_) returns()
func (_StakingNew *StakingNewSession) SetQueueConfig(wait_ *big.Int, autoCount_ *big.Int) (*types.Transaction, error) {
	return _StakingNew.Contract.SetQueueConfig(&_StakingNew.TransactOpts, wait_, autoCount_)
}

// SetQueueConfig is a paid mutator transaction binding the contract method 0xfd7fd35c.
//
// Solidity: function setQueueConfig(uint40 wait_, uint256 autoCount_) returns()
func (_StakingNew *StakingNewTransactorSession) SetQueueConfig(wait_ *big.Int, autoCount_ *big.Int) (*types.Transaction, error) {
	return _StakingNew.Contract.SetQueueConfig(&_StakingNew.TransactOpts, wait_, autoCount_)
}

// SetTeamClearConfig is a paid mutator transaction binding the contract method 0xc84b0d7c.
//
// Solidity: function setTeamClearConfig(uint40 period) returns()
func (_StakingNew *StakingNewTransactor) SetTeamClearConfig(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "setTeamClearConfig", period)
}

// SetTeamClearConfig is a paid mutator transaction binding the contract method 0xc84b0d7c.
//
// Solidity: function setTeamClearConfig(uint40 period) returns()
func (_StakingNew *StakingNewSession) SetTeamClearConfig(period *big.Int) (*types.Transaction, error) {
	return _StakingNew.Contract.SetTeamClearConfig(&_StakingNew.TransactOpts, period)
}

// SetTeamClearConfig is a paid mutator transaction binding the contract method 0xc84b0d7c.
//
// Solidity: function setTeamClearConfig(uint40 period) returns()
func (_StakingNew *StakingNewTransactorSession) SetTeamClearConfig(period *big.Int) (*types.Transaction, error) {
	return _StakingNew.Contract.SetTeamClearConfig(&_StakingNew.TransactOpts, period)
}

// SetWallets is a paid mutator transaction binding the contract method 0x75cb1bd1.
//
// Solidity: function setWallets(address market_, address eco_, address sink_) returns()
func (_StakingNew *StakingNewTransactor) SetWallets(opts *bind.TransactOpts, market_ common.Address, eco_ common.Address, sink_ common.Address) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "setWallets", market_, eco_, sink_)
}

// SetWallets is a paid mutator transaction binding the contract method 0x75cb1bd1.
//
// Solidity: function setWallets(address market_, address eco_, address sink_) returns()
func (_StakingNew *StakingNewSession) SetWallets(market_ common.Address, eco_ common.Address, sink_ common.Address) (*types.Transaction, error) {
	return _StakingNew.Contract.SetWallets(&_StakingNew.TransactOpts, market_, eco_, sink_)
}

// SetWallets is a paid mutator transaction binding the contract method 0x75cb1bd1.
//
// Solidity: function setWallets(address market_, address eco_, address sink_) returns()
func (_StakingNew *StakingNewTransactorSession) SetWallets(market_ common.Address, eco_ common.Address, sink_ common.Address) (*types.Transaction, error) {
	return _StakingNew.Contract.SetWallets(&_StakingNew.TransactOpts, market_, eco_, sink_)
}

// Stake is a paid mutator transaction binding the contract method 0x42ea02c1.
//
// Solidity: function stake(uint256 amount, uint16 plan) returns(uint256 count)
func (_StakingNew *StakingNewTransactor) Stake(opts *bind.TransactOpts, amount *big.Int, plan uint16) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "stake", amount, plan)
}

// Stake is a paid mutator transaction binding the contract method 0x42ea02c1.
//
// Solidity: function stake(uint256 amount, uint16 plan) returns(uint256 count)
func (_StakingNew *StakingNewSession) Stake(amount *big.Int, plan uint16) (*types.Transaction, error) {
	return _StakingNew.Contract.Stake(&_StakingNew.TransactOpts, amount, plan)
}

// Stake is a paid mutator transaction binding the contract method 0x42ea02c1.
//
// Solidity: function stake(uint256 amount, uint16 plan) returns(uint256 count)
func (_StakingNew *StakingNewTransactorSession) Stake(amount *big.Int, plan uint16) (*types.Transaction, error) {
	return _StakingNew.Contract.Stake(&_StakingNew.TransactOpts, amount, plan)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingNew *StakingNewTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingNew *StakingNewSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingNew.Contract.TransferOwnership(&_StakingNew.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingNew *StakingNewTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingNew.Contract.TransferOwnership(&_StakingNew.TransactOpts, newOwner)
}

// UpdateCircuitBreaker is a paid mutator transaction binding the contract method 0x53378e09.
//
// Solidity: function updateCircuitBreaker() returns()
func (_StakingNew *StakingNewTransactor) UpdateCircuitBreaker(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingNew.contract.Transact(opts, "updateCircuitBreaker")
}

// UpdateCircuitBreaker is a paid mutator transaction binding the contract method 0x53378e09.
//
// Solidity: function updateCircuitBreaker() returns()
func (_StakingNew *StakingNewSession) UpdateCircuitBreaker() (*types.Transaction, error) {
	return _StakingNew.Contract.UpdateCircuitBreaker(&_StakingNew.TransactOpts)
}

// UpdateCircuitBreaker is a paid mutator transaction binding the contract method 0x53378e09.
//
// Solidity: function updateCircuitBreaker() returns()
func (_StakingNew *StakingNewTransactorSession) UpdateCircuitBreaker() (*types.Transaction, error) {
	return _StakingNew.Contract.UpdateCircuitBreaker(&_StakingNew.TransactOpts)
}

// StakingNewAntiDumpLinePaySetIterator is returned from FilterAntiDumpLinePaySet and is used to iterate over the raw logs and unpacked data for AntiDumpLinePaySet events raised by the StakingNew contract.
type StakingNewAntiDumpLinePaySetIterator struct {
	Event *StakingNewAntiDumpLinePaySet // Event containing the contract specifics and raw log

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
func (it *StakingNewAntiDumpLinePaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewAntiDumpLinePaySet)
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
		it.Event = new(StakingNewAntiDumpLinePaySet)
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
func (it *StakingNewAntiDumpLinePaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewAntiDumpLinePaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewAntiDumpLinePaySet represents a AntiDumpLinePaySet event raised by the StakingNew contract.
type StakingNewAntiDumpLinePaySet struct {
	PayU bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAntiDumpLinePaySet is a free log retrieval operation binding the contract event 0x1f712eaa709c72afb054dbd90d01141271ce76f090115599ed7d2173c9cf6b25.
//
// Solidity: event AntiDumpLinePaySet(bool payU)
func (_StakingNew *StakingNewFilterer) FilterAntiDumpLinePaySet(opts *bind.FilterOpts) (*StakingNewAntiDumpLinePaySetIterator, error) {

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "AntiDumpLinePaySet")
	if err != nil {
		return nil, err
	}
	return &StakingNewAntiDumpLinePaySetIterator{contract: _StakingNew.contract, event: "AntiDumpLinePaySet", logs: logs, sub: sub}, nil
}

// WatchAntiDumpLinePaySet is a free log subscription operation binding the contract event 0x1f712eaa709c72afb054dbd90d01141271ce76f090115599ed7d2173c9cf6b25.
//
// Solidity: event AntiDumpLinePaySet(bool payU)
func (_StakingNew *StakingNewFilterer) WatchAntiDumpLinePaySet(opts *bind.WatchOpts, sink chan<- *StakingNewAntiDumpLinePaySet) (event.Subscription, error) {

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "AntiDumpLinePaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewAntiDumpLinePaySet)
				if err := _StakingNew.contract.UnpackLog(event, "AntiDumpLinePaySet", log); err != nil {
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

// ParseAntiDumpLinePaySet is a log parse operation binding the contract event 0x1f712eaa709c72afb054dbd90d01141271ce76f090115599ed7d2173c9cf6b25.
//
// Solidity: event AntiDumpLinePaySet(bool payU)
func (_StakingNew *StakingNewFilterer) ParseAntiDumpLinePaySet(log types.Log) (*StakingNewAntiDumpLinePaySet, error) {
	event := new(StakingNewAntiDumpLinePaySet)
	if err := _StakingNew.contract.UnpackLog(event, "AntiDumpLinePaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewCircuitBreakerParamsSetIterator is returned from FilterCircuitBreakerParamsSet and is used to iterate over the raw logs and unpacked data for CircuitBreakerParamsSet events raised by the StakingNew contract.
type StakingNewCircuitBreakerParamsSetIterator struct {
	Event *StakingNewCircuitBreakerParamsSet // Event containing the contract specifics and raw log

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
func (it *StakingNewCircuitBreakerParamsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewCircuitBreakerParamsSet)
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
		it.Event = new(StakingNewCircuitBreakerParamsSet)
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
func (it *StakingNewCircuitBreakerParamsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewCircuitBreakerParamsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewCircuitBreakerParamsSet represents a CircuitBreakerParamsSet event raised by the StakingNew contract.
type StakingNewCircuitBreakerParamsSet struct {
	Threshold     *big.Int
	Recovery      *big.Int
	CompBpsPerDay uint16
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCircuitBreakerParamsSet is a free log retrieval operation binding the contract event 0x09d6ae02ab4434467669ef6eb5dc07682e228b38ed30f154d501369a01284332.
//
// Solidity: event CircuitBreakerParamsSet(uint256 threshold, uint40 recovery, uint16 compBpsPerDay)
func (_StakingNew *StakingNewFilterer) FilterCircuitBreakerParamsSet(opts *bind.FilterOpts) (*StakingNewCircuitBreakerParamsSetIterator, error) {

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "CircuitBreakerParamsSet")
	if err != nil {
		return nil, err
	}
	return &StakingNewCircuitBreakerParamsSetIterator{contract: _StakingNew.contract, event: "CircuitBreakerParamsSet", logs: logs, sub: sub}, nil
}

// WatchCircuitBreakerParamsSet is a free log subscription operation binding the contract event 0x09d6ae02ab4434467669ef6eb5dc07682e228b38ed30f154d501369a01284332.
//
// Solidity: event CircuitBreakerParamsSet(uint256 threshold, uint40 recovery, uint16 compBpsPerDay)
func (_StakingNew *StakingNewFilterer) WatchCircuitBreakerParamsSet(opts *bind.WatchOpts, sink chan<- *StakingNewCircuitBreakerParamsSet) (event.Subscription, error) {

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "CircuitBreakerParamsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewCircuitBreakerParamsSet)
				if err := _StakingNew.contract.UnpackLog(event, "CircuitBreakerParamsSet", log); err != nil {
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

// ParseCircuitBreakerParamsSet is a log parse operation binding the contract event 0x09d6ae02ab4434467669ef6eb5dc07682e228b38ed30f154d501369a01284332.
//
// Solidity: event CircuitBreakerParamsSet(uint256 threshold, uint40 recovery, uint16 compBpsPerDay)
func (_StakingNew *StakingNewFilterer) ParseCircuitBreakerParamsSet(log types.Log) (*StakingNewCircuitBreakerParamsSet, error) {
	event := new(StakingNewCircuitBreakerParamsSet)
	if err := _StakingNew.contract.UnpackLog(event, "CircuitBreakerParamsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewCircuitBreakerStateUpdatedIterator is returned from FilterCircuitBreakerStateUpdated and is used to iterate over the raw logs and unpacked data for CircuitBreakerStateUpdated events raised by the StakingNew contract.
type StakingNewCircuitBreakerStateUpdatedIterator struct {
	Event *StakingNewCircuitBreakerStateUpdated // Event containing the contract specifics and raw log

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
func (it *StakingNewCircuitBreakerStateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewCircuitBreakerStateUpdated)
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
		it.Event = new(StakingNewCircuitBreakerStateUpdated)
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
func (it *StakingNewCircuitBreakerStateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewCircuitBreakerStateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewCircuitBreakerStateUpdated represents a CircuitBreakerStateUpdated event raised by the StakingNew contract.
type StakingNewCircuitBreakerStateUpdated struct {
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
func (_StakingNew *StakingNewFilterer) FilterCircuitBreakerStateUpdated(opts *bind.FilterOpts) (*StakingNewCircuitBreakerStateUpdatedIterator, error) {

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "CircuitBreakerStateUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingNewCircuitBreakerStateUpdatedIterator{contract: _StakingNew.contract, event: "CircuitBreakerStateUpdated", logs: logs, sub: sub}, nil
}

// WatchCircuitBreakerStateUpdated is a free log subscription operation binding the contract event 0x5f796bb8756e027a2b866bace59149e48633b3e059078664ea6f1ccd3224f94f.
//
// Solidity: event CircuitBreakerStateUpdated(uint40 circuitBreakerTime, uint40 newHighTime, uint256 currentDropBps, uint8 action, uint256 roundCount)
func (_StakingNew *StakingNewFilterer) WatchCircuitBreakerStateUpdated(opts *bind.WatchOpts, sink chan<- *StakingNewCircuitBreakerStateUpdated) (event.Subscription, error) {

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "CircuitBreakerStateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewCircuitBreakerStateUpdated)
				if err := _StakingNew.contract.UnpackLog(event, "CircuitBreakerStateUpdated", log); err != nil {
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
func (_StakingNew *StakingNewFilterer) ParseCircuitBreakerStateUpdated(log types.Log) (*StakingNewCircuitBreakerStateUpdated, error) {
	event := new(StakingNewCircuitBreakerStateUpdated)
	if err := _StakingNew.contract.UnpackLog(event, "CircuitBreakerStateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewEntryFeesSetIterator is returned from FilterEntryFeesSet and is used to iterate over the raw logs and unpacked data for EntryFeesSet events raised by the StakingNew contract.
type StakingNewEntryFeesSetIterator struct {
	Event *StakingNewEntryFeesSet // Event containing the contract specifics and raw log

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
func (it *StakingNewEntryFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewEntryFeesSet)
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
		it.Event = new(StakingNewEntryFeesSet)
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
func (it *StakingNewEntryFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewEntryFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewEntryFeesSet represents a EntryFeesSet event raised by the StakingNew contract.
type StakingNewEntryFeesSet struct {
	LiqBps        uint16
	MarketBps     uint16
	EcoBps        uint16
	TeamMaxBps    uint16
	TeamDirectBps uint16
	SmallFeeBps   uint16
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEntryFeesSet is a free log retrieval operation binding the contract event 0x76a79d2857e8da3cadca0dce8d5adc18938908806fdb823e4758e83624e699ab.
//
// Solidity: event EntryFeesSet(uint16 liqBps, uint16 marketBps, uint16 ecoBps, uint16 teamMaxBps, uint16 teamDirectBps, uint16 smallFeeBps)
func (_StakingNew *StakingNewFilterer) FilterEntryFeesSet(opts *bind.FilterOpts) (*StakingNewEntryFeesSetIterator, error) {

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "EntryFeesSet")
	if err != nil {
		return nil, err
	}
	return &StakingNewEntryFeesSetIterator{contract: _StakingNew.contract, event: "EntryFeesSet", logs: logs, sub: sub}, nil
}

// WatchEntryFeesSet is a free log subscription operation binding the contract event 0x76a79d2857e8da3cadca0dce8d5adc18938908806fdb823e4758e83624e699ab.
//
// Solidity: event EntryFeesSet(uint16 liqBps, uint16 marketBps, uint16 ecoBps, uint16 teamMaxBps, uint16 teamDirectBps, uint16 smallFeeBps)
func (_StakingNew *StakingNewFilterer) WatchEntryFeesSet(opts *bind.WatchOpts, sink chan<- *StakingNewEntryFeesSet) (event.Subscription, error) {

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "EntryFeesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewEntryFeesSet)
				if err := _StakingNew.contract.UnpackLog(event, "EntryFeesSet", log); err != nil {
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

// ParseEntryFeesSet is a log parse operation binding the contract event 0x76a79d2857e8da3cadca0dce8d5adc18938908806fdb823e4758e83624e699ab.
//
// Solidity: event EntryFeesSet(uint16 liqBps, uint16 marketBps, uint16 ecoBps, uint16 teamMaxBps, uint16 teamDirectBps, uint16 smallFeeBps)
func (_StakingNew *StakingNewFilterer) ParseEntryFeesSet(log types.Log) (*StakingNewEntryFeesSet, error) {
	event := new(StakingNewEntryFeesSet)
	if err := _StakingNew.contract.UnpackLog(event, "EntryFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewInitChangedIterator is returned from FilterInitChanged and is used to iterate over the raw logs and unpacked data for InitChanged events raised by the StakingNew contract.
type StakingNewInitChangedIterator struct {
	Event *StakingNewInitChanged // Event containing the contract specifics and raw log

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
func (it *StakingNewInitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewInitChanged)
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
		it.Event = new(StakingNewInitChanged)
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
func (it *StakingNewInitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewInitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewInitChanged represents a InitChanged event raised by the StakingNew contract.
type StakingNewInitChanged struct {
	OldInit common.Address
	NewInit common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitChanged is a free log retrieval operation binding the contract event 0x3f17984892df4d730451c12071a40f5396f4470a024466ba1237f3163632185a.
//
// Solidity: event InitChanged(address indexed oldInit, address indexed newInit)
func (_StakingNew *StakingNewFilterer) FilterInitChanged(opts *bind.FilterOpts, oldInit []common.Address, newInit []common.Address) (*StakingNewInitChangedIterator, error) {

	var oldInitRule []interface{}
	for _, oldInitItem := range oldInit {
		oldInitRule = append(oldInitRule, oldInitItem)
	}
	var newInitRule []interface{}
	for _, newInitItem := range newInit {
		newInitRule = append(newInitRule, newInitItem)
	}

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "InitChanged", oldInitRule, newInitRule)
	if err != nil {
		return nil, err
	}
	return &StakingNewInitChangedIterator{contract: _StakingNew.contract, event: "InitChanged", logs: logs, sub: sub}, nil
}

// WatchInitChanged is a free log subscription operation binding the contract event 0x3f17984892df4d730451c12071a40f5396f4470a024466ba1237f3163632185a.
//
// Solidity: event InitChanged(address indexed oldInit, address indexed newInit)
func (_StakingNew *StakingNewFilterer) WatchInitChanged(opts *bind.WatchOpts, sink chan<- *StakingNewInitChanged, oldInit []common.Address, newInit []common.Address) (event.Subscription, error) {

	var oldInitRule []interface{}
	for _, oldInitItem := range oldInit {
		oldInitRule = append(oldInitRule, oldInitItem)
	}
	var newInitRule []interface{}
	for _, newInitItem := range newInit {
		newInitRule = append(newInitRule, newInitItem)
	}

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "InitChanged", oldInitRule, newInitRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewInitChanged)
				if err := _StakingNew.contract.UnpackLog(event, "InitChanged", log); err != nil {
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

// ParseInitChanged is a log parse operation binding the contract event 0x3f17984892df4d730451c12071a40f5396f4470a024466ba1237f3163632185a.
//
// Solidity: event InitChanged(address indexed oldInit, address indexed newInit)
func (_StakingNew *StakingNewFilterer) ParseInitChanged(log types.Log) (*StakingNewInitChanged, error) {
	event := new(StakingNewInitChanged)
	if err := _StakingNew.contract.UnpackLog(event, "InitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewLaunchQueueSetIterator is returned from FilterLaunchQueueSet and is used to iterate over the raw logs and unpacked data for LaunchQueueSet events raised by the StakingNew contract.
type StakingNewLaunchQueueSetIterator struct {
	Event *StakingNewLaunchQueueSet // Event containing the contract specifics and raw log

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
func (it *StakingNewLaunchQueueSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewLaunchQueueSet)
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
		it.Event = new(StakingNewLaunchQueueSet)
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
func (it *StakingNewLaunchQueueSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewLaunchQueueSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewLaunchQueueSet represents a LaunchQueueSet event raised by the StakingNew contract.
type StakingNewLaunchQueueSet struct {
	Cutoff *big.Int
	Amount *big.Int
	Done   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLaunchQueueSet is a free log retrieval operation binding the contract event 0xae6745b74a6c0e1bab149bf5f374e2f5f55efdd4de6c0991014c6ab70b6cf7eb.
//
// Solidity: event LaunchQueueSet(uint40 cutoff, uint256 amount, uint256 done)
func (_StakingNew *StakingNewFilterer) FilterLaunchQueueSet(opts *bind.FilterOpts) (*StakingNewLaunchQueueSetIterator, error) {

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "LaunchQueueSet")
	if err != nil {
		return nil, err
	}
	return &StakingNewLaunchQueueSetIterator{contract: _StakingNew.contract, event: "LaunchQueueSet", logs: logs, sub: sub}, nil
}

// WatchLaunchQueueSet is a free log subscription operation binding the contract event 0xae6745b74a6c0e1bab149bf5f374e2f5f55efdd4de6c0991014c6ab70b6cf7eb.
//
// Solidity: event LaunchQueueSet(uint40 cutoff, uint256 amount, uint256 done)
func (_StakingNew *StakingNewFilterer) WatchLaunchQueueSet(opts *bind.WatchOpts, sink chan<- *StakingNewLaunchQueueSet) (event.Subscription, error) {

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "LaunchQueueSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewLaunchQueueSet)
				if err := _StakingNew.contract.UnpackLog(event, "LaunchQueueSet", log); err != nil {
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

// ParseLaunchQueueSet is a log parse operation binding the contract event 0xae6745b74a6c0e1bab149bf5f374e2f5f55efdd4de6c0991014c6ab70b6cf7eb.
//
// Solidity: event LaunchQueueSet(uint40 cutoff, uint256 amount, uint256 done)
func (_StakingNew *StakingNewFilterer) ParseLaunchQueueSet(log types.Log) (*StakingNewLaunchQueueSet, error) {
	event := new(StakingNewLaunchQueueSet)
	if err := _StakingNew.contract.UnpackLog(event, "LaunchQueueSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewLineClaimedIterator is returned from FilterLineClaimed and is used to iterate over the raw logs and unpacked data for LineClaimed events raised by the StakingNew contract.
type StakingNewLineClaimedIterator struct {
	Event *StakingNewLineClaimed // Event containing the contract specifics and raw log

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
func (it *StakingNewLineClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewLineClaimed)
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
		it.Event = new(StakingNewLineClaimed)
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
func (it *StakingNewLineClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewLineClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewLineClaimed represents a LineClaimed event raised by the StakingNew contract.
type StakingNewLineClaimed struct {
	User     common.Address
	OrderId  *big.Int
	GrossU   *big.Int
	FeeU     *big.Int
	PaidMs   bool
	MsAmount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLineClaimed is a free log retrieval operation binding the contract event 0x9b323bf7a7ff7d9a2f07b56ab4f293d9a140b8e6dd6d39dbdc34788ff4e4ceaf.
//
// Solidity: event LineClaimed(address indexed user, uint256 orderId, uint256 grossU, uint256 feeU, bool paidMs, uint256 msAmount)
func (_StakingNew *StakingNewFilterer) FilterLineClaimed(opts *bind.FilterOpts, user []common.Address) (*StakingNewLineClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "LineClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingNewLineClaimedIterator{contract: _StakingNew.contract, event: "LineClaimed", logs: logs, sub: sub}, nil
}

// WatchLineClaimed is a free log subscription operation binding the contract event 0x9b323bf7a7ff7d9a2f07b56ab4f293d9a140b8e6dd6d39dbdc34788ff4e4ceaf.
//
// Solidity: event LineClaimed(address indexed user, uint256 orderId, uint256 grossU, uint256 feeU, bool paidMs, uint256 msAmount)
func (_StakingNew *StakingNewFilterer) WatchLineClaimed(opts *bind.WatchOpts, sink chan<- *StakingNewLineClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "LineClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewLineClaimed)
				if err := _StakingNew.contract.UnpackLog(event, "LineClaimed", log); err != nil {
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

// ParseLineClaimed is a log parse operation binding the contract event 0x9b323bf7a7ff7d9a2f07b56ab4f293d9a140b8e6dd6d39dbdc34788ff4e4ceaf.
//
// Solidity: event LineClaimed(address indexed user, uint256 orderId, uint256 grossU, uint256 feeU, bool paidMs, uint256 msAmount)
func (_StakingNew *StakingNewFilterer) ParseLineClaimed(log types.Log) (*StakingNewLineClaimed, error) {
	event := new(StakingNewLineClaimed)
	if err := _StakingNew.contract.UnpackLog(event, "LineClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewLineFeesSetIterator is returned from FilterLineFeesSet and is used to iterate over the raw logs and unpacked data for LineFeesSet events raised by the StakingNew contract.
type StakingNewLineFeesSetIterator struct {
	Event *StakingNewLineFeesSet // Event containing the contract specifics and raw log

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
func (it *StakingNewLineFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewLineFeesSet)
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
		it.Event = new(StakingNewLineFeesSet)
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
func (it *StakingNewLineFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewLineFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewLineFeesSet represents a LineFeesSet event raised by the StakingNew contract.
type StakingNewLineFeesSet struct {
	LineFeeBps   uint16
	LineEcoBps   uint16
	LineLevelBps [4]uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLineFeesSet is a free log retrieval operation binding the contract event 0x44077834aa18a430e258a029a980f3fa81c9f4f922a17786c4d047b87834c368.
//
// Solidity: event LineFeesSet(uint16 lineFeeBps, uint16 lineEcoBps, uint16[4] lineLevelBps)
func (_StakingNew *StakingNewFilterer) FilterLineFeesSet(opts *bind.FilterOpts) (*StakingNewLineFeesSetIterator, error) {

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "LineFeesSet")
	if err != nil {
		return nil, err
	}
	return &StakingNewLineFeesSetIterator{contract: _StakingNew.contract, event: "LineFeesSet", logs: logs, sub: sub}, nil
}

// WatchLineFeesSet is a free log subscription operation binding the contract event 0x44077834aa18a430e258a029a980f3fa81c9f4f922a17786c4d047b87834c368.
//
// Solidity: event LineFeesSet(uint16 lineFeeBps, uint16 lineEcoBps, uint16[4] lineLevelBps)
func (_StakingNew *StakingNewFilterer) WatchLineFeesSet(opts *bind.WatchOpts, sink chan<- *StakingNewLineFeesSet) (event.Subscription, error) {

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "LineFeesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewLineFeesSet)
				if err := _StakingNew.contract.UnpackLog(event, "LineFeesSet", log); err != nil {
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

// ParseLineFeesSet is a log parse operation binding the contract event 0x44077834aa18a430e258a029a980f3fa81c9f4f922a17786c4d047b87834c368.
//
// Solidity: event LineFeesSet(uint16 lineFeeBps, uint16 lineEcoBps, uint16[4] lineLevelBps)
func (_StakingNew *StakingNewFilterer) ParseLineFeesSet(log types.Log) (*StakingNewLineFeesSet, error) {
	event := new(StakingNewLineFeesSet)
	if err := _StakingNew.contract.UnpackLog(event, "LineFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewLineSafeSetIterator is returned from FilterLineSafeSet and is used to iterate over the raw logs and unpacked data for LineSafeSet events raised by the StakingNew contract.
type StakingNewLineSafeSetIterator struct {
	Event *StakingNewLineSafeSet // Event containing the contract specifics and raw log

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
func (it *StakingNewLineSafeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewLineSafeSet)
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
		it.Event = new(StakingNewLineSafeSet)
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
func (it *StakingNewLineSafeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewLineSafeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewLineSafeSet represents a LineSafeSet event raised by the StakingNew contract.
type StakingNewLineSafeSet struct {
	SecondsAmount *big.Int
	BeforeTime    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLineSafeSet is a free log retrieval operation binding the contract event 0x626b219e3adefa4821b3e1cb58781daf2a53f9c5df84083ac9d040a6d37a8e8b.
//
// Solidity: event LineSafeSet(uint40 secondsAmount, uint40 beforeTime)
func (_StakingNew *StakingNewFilterer) FilterLineSafeSet(opts *bind.FilterOpts) (*StakingNewLineSafeSetIterator, error) {

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "LineSafeSet")
	if err != nil {
		return nil, err
	}
	return &StakingNewLineSafeSetIterator{contract: _StakingNew.contract, event: "LineSafeSet", logs: logs, sub: sub}, nil
}

// WatchLineSafeSet is a free log subscription operation binding the contract event 0x626b219e3adefa4821b3e1cb58781daf2a53f9c5df84083ac9d040a6d37a8e8b.
//
// Solidity: event LineSafeSet(uint40 secondsAmount, uint40 beforeTime)
func (_StakingNew *StakingNewFilterer) WatchLineSafeSet(opts *bind.WatchOpts, sink chan<- *StakingNewLineSafeSet) (event.Subscription, error) {

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "LineSafeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewLineSafeSet)
				if err := _StakingNew.contract.UnpackLog(event, "LineSafeSet", log); err != nil {
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

// ParseLineSafeSet is a log parse operation binding the contract event 0x626b219e3adefa4821b3e1cb58781daf2a53f9c5df84083ac9d040a6d37a8e8b.
//
// Solidity: event LineSafeSet(uint40 secondsAmount, uint40 beforeTime)
func (_StakingNew *StakingNewFilterer) ParseLineSafeSet(log types.Log) (*StakingNewLineSafeSet, error) {
	event := new(StakingNewLineSafeSet)
	if err := _StakingNew.contract.UnpackLog(event, "LineSafeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewLockedSetIterator is returned from FilterLockedSet and is used to iterate over the raw logs and unpacked data for LockedSet events raised by the StakingNew contract.
type StakingNewLockedSetIterator struct {
	Event *StakingNewLockedSet // Event containing the contract specifics and raw log

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
func (it *StakingNewLockedSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewLockedSet)
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
		it.Event = new(StakingNewLockedSet)
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
func (it *StakingNewLockedSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewLockedSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewLockedSet represents a LockedSet event raised by the StakingNew contract.
type StakingNewLockedSet struct {
	Account common.Address
	Locked  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLockedSet is a free log retrieval operation binding the contract event 0xb2bfa61755f616ba3cfedb3ccae9f8172b523c43f77871c6aca0b52e8fd8c946.
//
// Solidity: event LockedSet(address indexed account, bool locked)
func (_StakingNew *StakingNewFilterer) FilterLockedSet(opts *bind.FilterOpts, account []common.Address) (*StakingNewLockedSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "LockedSet", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakingNewLockedSetIterator{contract: _StakingNew.contract, event: "LockedSet", logs: logs, sub: sub}, nil
}

// WatchLockedSet is a free log subscription operation binding the contract event 0xb2bfa61755f616ba3cfedb3ccae9f8172b523c43f77871c6aca0b52e8fd8c946.
//
// Solidity: event LockedSet(address indexed account, bool locked)
func (_StakingNew *StakingNewFilterer) WatchLockedSet(opts *bind.WatchOpts, sink chan<- *StakingNewLockedSet, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "LockedSet", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewLockedSet)
				if err := _StakingNew.contract.UnpackLog(event, "LockedSet", log); err != nil {
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

// ParseLockedSet is a log parse operation binding the contract event 0xb2bfa61755f616ba3cfedb3ccae9f8172b523c43f77871c6aca0b52e8fd8c946.
//
// Solidity: event LockedSet(address indexed account, bool locked)
func (_StakingNew *StakingNewFilterer) ParseLockedSet(log types.Log) (*StakingNewLockedSet, error) {
	event := new(StakingNewLockedSet)
	if err := _StakingNew.contract.UnpackLog(event, "LockedSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewOrderCapSetIterator is returned from FilterOrderCapSet and is used to iterate over the raw logs and unpacked data for OrderCapSet events raised by the StakingNew contract.
type StakingNewOrderCapSetIterator struct {
	Event *StakingNewOrderCapSet // Event containing the contract specifics and raw log

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
func (it *StakingNewOrderCapSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewOrderCapSet)
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
		it.Event = new(StakingNewOrderCapSet)
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
func (it *StakingNewOrderCapSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewOrderCapSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewOrderCapSet represents a OrderCapSet event raised by the StakingNew contract.
type StakingNewOrderCapSet struct {
	User    common.Address
	Index   *big.Int
	OrderId *big.Int
	OldCap  *big.Int
	NewCap  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderCapSet is a free log retrieval operation binding the contract event 0x6409fdae075f17dc54fe696f58828676d5fdec9bfa2fbb13bdc8aa5e662fd6e3.
//
// Solidity: event OrderCapSet(address indexed user, uint256 indexed index, uint256 indexed orderId, uint256 oldCap, uint256 newCap)
func (_StakingNew *StakingNewFilterer) FilterOrderCapSet(opts *bind.FilterOpts, user []common.Address, index []*big.Int, orderId []*big.Int) (*StakingNewOrderCapSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "OrderCapSet", userRule, indexRule, orderIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingNewOrderCapSetIterator{contract: _StakingNew.contract, event: "OrderCapSet", logs: logs, sub: sub}, nil
}

// WatchOrderCapSet is a free log subscription operation binding the contract event 0x6409fdae075f17dc54fe696f58828676d5fdec9bfa2fbb13bdc8aa5e662fd6e3.
//
// Solidity: event OrderCapSet(address indexed user, uint256 indexed index, uint256 indexed orderId, uint256 oldCap, uint256 newCap)
func (_StakingNew *StakingNewFilterer) WatchOrderCapSet(opts *bind.WatchOpts, sink chan<- *StakingNewOrderCapSet, user []common.Address, index []*big.Int, orderId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "OrderCapSet", userRule, indexRule, orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewOrderCapSet)
				if err := _StakingNew.contract.UnpackLog(event, "OrderCapSet", log); err != nil {
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

// ParseOrderCapSet is a log parse operation binding the contract event 0x6409fdae075f17dc54fe696f58828676d5fdec9bfa2fbb13bdc8aa5e662fd6e3.
//
// Solidity: event OrderCapSet(address indexed user, uint256 indexed index, uint256 indexed orderId, uint256 oldCap, uint256 newCap)
func (_StakingNew *StakingNewFilterer) ParseOrderCapSet(log types.Log) (*StakingNewOrderCapSet, error) {
	event := new(StakingNewOrderCapSet)
	if err := _StakingNew.contract.UnpackLog(event, "OrderCapSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the StakingNew contract.
type StakingNewOrderCreatedIterator struct {
	Event *StakingNewOrderCreated // Event containing the contract specifics and raw log

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
func (it *StakingNewOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewOrderCreated)
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
		it.Event = new(StakingNewOrderCreated)
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
func (it *StakingNewOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewOrderCreated represents a OrderCreated event raised by the StakingNew contract.
type StakingNewOrderCreated struct {
	OrderId *big.Int
	User    common.Address
	Amount  *big.Int
	Cap     *big.Int
	Plan    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0x92669f7499828c9972b3a5b6c0b107ca011f143600fd7b32517024d4a3d7ab28.
//
// Solidity: event OrderCreated(uint256 indexed orderId, address indexed user, uint256 amount, uint256 cap, uint256 plan)
func (_StakingNew *StakingNewFilterer) FilterOrderCreated(opts *bind.FilterOpts, orderId []*big.Int, user []common.Address) (*StakingNewOrderCreatedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "OrderCreated", orderIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &StakingNewOrderCreatedIterator{contract: _StakingNew.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x92669f7499828c9972b3a5b6c0b107ca011f143600fd7b32517024d4a3d7ab28.
//
// Solidity: event OrderCreated(uint256 indexed orderId, address indexed user, uint256 amount, uint256 cap, uint256 plan)
func (_StakingNew *StakingNewFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *StakingNewOrderCreated, orderId []*big.Int, user []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "OrderCreated", orderIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewOrderCreated)
				if err := _StakingNew.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0x92669f7499828c9972b3a5b6c0b107ca011f143600fd7b32517024d4a3d7ab28.
//
// Solidity: event OrderCreated(uint256 indexed orderId, address indexed user, uint256 amount, uint256 cap, uint256 plan)
func (_StakingNew *StakingNewFilterer) ParseOrderCreated(log types.Log) (*StakingNewOrderCreated, error) {
	event := new(StakingNewOrderCreated)
	if err := _StakingNew.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewOrderEnteredIterator is returned from FilterOrderEntered and is used to iterate over the raw logs and unpacked data for OrderEntered events raised by the StakingNew contract.
type StakingNewOrderEnteredIterator struct {
	Event *StakingNewOrderEntered // Event containing the contract specifics and raw log

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
func (it *StakingNewOrderEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewOrderEntered)
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
		it.Event = new(StakingNewOrderEntered)
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
func (it *StakingNewOrderEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewOrderEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewOrderEntered represents a OrderEntered event raised by the StakingNew contract.
type StakingNewOrderEntered struct {
	OrderId *big.Int
	User    common.Address
	Start   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderEntered is a free log retrieval operation binding the contract event 0x8e6880a634e5e4deb05e9fcba2112be19166e9721dd4946bc09091783c9a656f.
//
// Solidity: event OrderEntered(uint256 indexed orderId, address indexed user, uint40 start)
func (_StakingNew *StakingNewFilterer) FilterOrderEntered(opts *bind.FilterOpts, orderId []*big.Int, user []common.Address) (*StakingNewOrderEnteredIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "OrderEntered", orderIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &StakingNewOrderEnteredIterator{contract: _StakingNew.contract, event: "OrderEntered", logs: logs, sub: sub}, nil
}

// WatchOrderEntered is a free log subscription operation binding the contract event 0x8e6880a634e5e4deb05e9fcba2112be19166e9721dd4946bc09091783c9a656f.
//
// Solidity: event OrderEntered(uint256 indexed orderId, address indexed user, uint40 start)
func (_StakingNew *StakingNewFilterer) WatchOrderEntered(opts *bind.WatchOpts, sink chan<- *StakingNewOrderEntered, orderId []*big.Int, user []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "OrderEntered", orderIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewOrderEntered)
				if err := _StakingNew.contract.UnpackLog(event, "OrderEntered", log); err != nil {
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

// ParseOrderEntered is a log parse operation binding the contract event 0x8e6880a634e5e4deb05e9fcba2112be19166e9721dd4946bc09091783c9a656f.
//
// Solidity: event OrderEntered(uint256 indexed orderId, address indexed user, uint40 start)
func (_StakingNew *StakingNewFilterer) ParseOrderEntered(log types.Log) (*StakingNewOrderEntered, error) {
	event := new(StakingNewOrderEntered)
	if err := _StakingNew.contract.UnpackLog(event, "OrderEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewOrderExitedIterator is returned from FilterOrderExited and is used to iterate over the raw logs and unpacked data for OrderExited events raised by the StakingNew contract.
type StakingNewOrderExitedIterator struct {
	Event *StakingNewOrderExited // Event containing the contract specifics and raw log

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
func (it *StakingNewOrderExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewOrderExited)
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
		it.Event = new(StakingNewOrderExited)
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
func (it *StakingNewOrderExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewOrderExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewOrderExited represents a OrderExited event raised by the StakingNew contract.
type StakingNewOrderExited struct {
	OrderId *big.Int
	User    common.Address
	Amount  *big.Int
	Used    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderExited is a free log retrieval operation binding the contract event 0x53441419715315a4cc233346fb370d577a21d7fdff6d5f60860f24150a199cde.
//
// Solidity: event OrderExited(uint256 indexed orderId, address indexed user, uint256 amount, uint256 used)
func (_StakingNew *StakingNewFilterer) FilterOrderExited(opts *bind.FilterOpts, orderId []*big.Int, user []common.Address) (*StakingNewOrderExitedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "OrderExited", orderIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &StakingNewOrderExitedIterator{contract: _StakingNew.contract, event: "OrderExited", logs: logs, sub: sub}, nil
}

// WatchOrderExited is a free log subscription operation binding the contract event 0x53441419715315a4cc233346fb370d577a21d7fdff6d5f60860f24150a199cde.
//
// Solidity: event OrderExited(uint256 indexed orderId, address indexed user, uint256 amount, uint256 used)
func (_StakingNew *StakingNewFilterer) WatchOrderExited(opts *bind.WatchOpts, sink chan<- *StakingNewOrderExited, orderId []*big.Int, user []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "OrderExited", orderIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewOrderExited)
				if err := _StakingNew.contract.UnpackLog(event, "OrderExited", log); err != nil {
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

// ParseOrderExited is a log parse operation binding the contract event 0x53441419715315a4cc233346fb370d577a21d7fdff6d5f60860f24150a199cde.
//
// Solidity: event OrderExited(uint256 indexed orderId, address indexed user, uint256 amount, uint256 used)
func (_StakingNew *StakingNewFilterer) ParseOrderExited(log types.Log) (*StakingNewOrderExited, error) {
	event := new(StakingNewOrderExited)
	if err := _StakingNew.contract.UnpackLog(event, "OrderExited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingNew contract.
type StakingNewOwnershipTransferredIterator struct {
	Event *StakingNewOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingNewOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewOwnershipTransferred)
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
		it.Event = new(StakingNewOwnershipTransferred)
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
func (it *StakingNewOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewOwnershipTransferred represents a OwnershipTransferred event raised by the StakingNew contract.
type StakingNewOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_StakingNew *StakingNewFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*StakingNewOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingNewOwnershipTransferredIterator{contract: _StakingNew.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_StakingNew *StakingNewFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingNewOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewOwnershipTransferred)
				if err := _StakingNew.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_StakingNew *StakingNewFilterer) ParseOwnershipTransferred(log types.Log) (*StakingNewOwnershipTransferred, error) {
	event := new(StakingNewOwnershipTransferred)
	if err := _StakingNew.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewPlanSetIterator is returned from FilterPlanSet and is used to iterate over the raw logs and unpacked data for PlanSet events raised by the StakingNew contract.
type StakingNewPlanSetIterator struct {
	Event *StakingNewPlanSet // Event containing the contract specifics and raw log

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
func (it *StakingNewPlanSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewPlanSet)
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
		it.Event = new(StakingNewPlanSet)
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
func (it *StakingNewPlanSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewPlanSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewPlanSet represents a PlanSet event raised by the StakingNew contract.
type StakingNewPlanSet struct {
	Index     *big.Int
	MinAmount *big.Int
	MaxAmount *big.Int
	OutAmount *big.Int
	DaysCount uint32
	Enabled   bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPlanSet is a free log retrieval operation binding the contract event 0x7937dab559669144f9fd03274e00046917bb79f00ab85dc4cb7da1bb0b84227d.
//
// Solidity: event PlanSet(uint256 indexed index, uint256 minAmount, uint256 maxAmount, uint256 outAmount, uint32 daysCount, bool enabled)
func (_StakingNew *StakingNewFilterer) FilterPlanSet(opts *bind.FilterOpts, index []*big.Int) (*StakingNewPlanSetIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "PlanSet", indexRule)
	if err != nil {
		return nil, err
	}
	return &StakingNewPlanSetIterator{contract: _StakingNew.contract, event: "PlanSet", logs: logs, sub: sub}, nil
}

// WatchPlanSet is a free log subscription operation binding the contract event 0x7937dab559669144f9fd03274e00046917bb79f00ab85dc4cb7da1bb0b84227d.
//
// Solidity: event PlanSet(uint256 indexed index, uint256 minAmount, uint256 maxAmount, uint256 outAmount, uint32 daysCount, bool enabled)
func (_StakingNew *StakingNewFilterer) WatchPlanSet(opts *bind.WatchOpts, sink chan<- *StakingNewPlanSet, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "PlanSet", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewPlanSet)
				if err := _StakingNew.contract.UnpackLog(event, "PlanSet", log); err != nil {
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

// ParsePlanSet is a log parse operation binding the contract event 0x7937dab559669144f9fd03274e00046917bb79f00ab85dc4cb7da1bb0b84227d.
//
// Solidity: event PlanSet(uint256 indexed index, uint256 minAmount, uint256 maxAmount, uint256 outAmount, uint32 daysCount, bool enabled)
func (_StakingNew *StakingNewFilterer) ParsePlanSet(log types.Log) (*StakingNewPlanSet, error) {
	event := new(StakingNewPlanSet)
	if err := _StakingNew.contract.UnpackLog(event, "PlanSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewQueueConfigSetIterator is returned from FilterQueueConfigSet and is used to iterate over the raw logs and unpacked data for QueueConfigSet events raised by the StakingNew contract.
type StakingNewQueueConfigSetIterator struct {
	Event *StakingNewQueueConfigSet // Event containing the contract specifics and raw log

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
func (it *StakingNewQueueConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewQueueConfigSet)
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
		it.Event = new(StakingNewQueueConfigSet)
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
func (it *StakingNewQueueConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewQueueConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewQueueConfigSet represents a QueueConfigSet event raised by the StakingNew contract.
type StakingNewQueueConfigSet struct {
	QueueWait      *big.Int
	AutoQueueCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterQueueConfigSet is a free log retrieval operation binding the contract event 0x87353961a1c8681ab8f208a0f41986ae8adf72a6ae6a853acb2e4b7e60ca1a85.
//
// Solidity: event QueueConfigSet(uint40 queueWait, uint256 autoQueueCount)
func (_StakingNew *StakingNewFilterer) FilterQueueConfigSet(opts *bind.FilterOpts) (*StakingNewQueueConfigSetIterator, error) {

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "QueueConfigSet")
	if err != nil {
		return nil, err
	}
	return &StakingNewQueueConfigSetIterator{contract: _StakingNew.contract, event: "QueueConfigSet", logs: logs, sub: sub}, nil
}

// WatchQueueConfigSet is a free log subscription operation binding the contract event 0x87353961a1c8681ab8f208a0f41986ae8adf72a6ae6a853acb2e4b7e60ca1a85.
//
// Solidity: event QueueConfigSet(uint40 queueWait, uint256 autoQueueCount)
func (_StakingNew *StakingNewFilterer) WatchQueueConfigSet(opts *bind.WatchOpts, sink chan<- *StakingNewQueueConfigSet) (event.Subscription, error) {

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "QueueConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewQueueConfigSet)
				if err := _StakingNew.contract.UnpackLog(event, "QueueConfigSet", log); err != nil {
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

// ParseQueueConfigSet is a log parse operation binding the contract event 0x87353961a1c8681ab8f208a0f41986ae8adf72a6ae6a853acb2e4b7e60ca1a85.
//
// Solidity: event QueueConfigSet(uint40 queueWait, uint256 autoQueueCount)
func (_StakingNew *StakingNewFilterer) ParseQueueConfigSet(log types.Log) (*StakingNewQueueConfigSet, error) {
	event := new(StakingNewQueueConfigSet)
	if err := _StakingNew.contract.UnpackLog(event, "QueueConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewQueueDoneIterator is returned from FilterQueueDone and is used to iterate over the raw logs and unpacked data for QueueDone events raised by the StakingNew contract.
type StakingNewQueueDoneIterator struct {
	Event *StakingNewQueueDone // Event containing the contract specifics and raw log

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
func (it *StakingNewQueueDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewQueueDone)
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
		it.Event = new(StakingNewQueueDone)
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
func (it *StakingNewQueueDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewQueueDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewQueueDone represents a QueueDone event raised by the StakingNew contract.
type StakingNewQueueDone struct {
	Index   *big.Int
	OrderId *big.Int
	User    common.Address
	LiqU    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterQueueDone is a free log retrieval operation binding the contract event 0x7cb6d2f0f71ff6ec5c5549e81fadd5c1f8790b9d652b4c2d1ed4817998e4f4d6.
//
// Solidity: event QueueDone(uint256 indexed index, uint256 indexed orderId, address indexed user, uint256 liqU)
func (_StakingNew *StakingNewFilterer) FilterQueueDone(opts *bind.FilterOpts, index []*big.Int, orderId []*big.Int, user []common.Address) (*StakingNewQueueDoneIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "QueueDone", indexRule, orderIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &StakingNewQueueDoneIterator{contract: _StakingNew.contract, event: "QueueDone", logs: logs, sub: sub}, nil
}

// WatchQueueDone is a free log subscription operation binding the contract event 0x7cb6d2f0f71ff6ec5c5549e81fadd5c1f8790b9d652b4c2d1ed4817998e4f4d6.
//
// Solidity: event QueueDone(uint256 indexed index, uint256 indexed orderId, address indexed user, uint256 liqU)
func (_StakingNew *StakingNewFilterer) WatchQueueDone(opts *bind.WatchOpts, sink chan<- *StakingNewQueueDone, index []*big.Int, orderId []*big.Int, user []common.Address) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "QueueDone", indexRule, orderIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewQueueDone)
				if err := _StakingNew.contract.UnpackLog(event, "QueueDone", log); err != nil {
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

// ParseQueueDone is a log parse operation binding the contract event 0x7cb6d2f0f71ff6ec5c5549e81fadd5c1f8790b9d652b4c2d1ed4817998e4f4d6.
//
// Solidity: event QueueDone(uint256 indexed index, uint256 indexed orderId, address indexed user, uint256 liqU)
func (_StakingNew *StakingNewFilterer) ParseQueueDone(log types.Log) (*StakingNewQueueDone, error) {
	event := new(StakingNewQueueDone)
	if err := _StakingNew.contract.UnpackLog(event, "QueueDone", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewQueuedIterator is returned from FilterQueued and is used to iterate over the raw logs and unpacked data for Queued events raised by the StakingNew contract.
type StakingNewQueuedIterator struct {
	Event *StakingNewQueued // Event containing the contract specifics and raw log

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
func (it *StakingNewQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewQueued)
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
		it.Event = new(StakingNewQueued)
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
func (it *StakingNewQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewQueued represents a Queued event raised by the StakingNew contract.
type StakingNewQueued struct {
	Index    *big.Int
	OrderId  *big.Int
	User     common.Address
	LiqU     *big.Int
	QueuedAt *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterQueued is a free log retrieval operation binding the contract event 0x33ed2ea816789b9d12d676e2ed6d81ab2cccc83a440f9a2576d15858c810bea9.
//
// Solidity: event Queued(uint256 indexed index, uint256 indexed orderId, address indexed user, uint256 liqU, uint40 queuedAt)
func (_StakingNew *StakingNewFilterer) FilterQueued(opts *bind.FilterOpts, index []*big.Int, orderId []*big.Int, user []common.Address) (*StakingNewQueuedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "Queued", indexRule, orderIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &StakingNewQueuedIterator{contract: _StakingNew.contract, event: "Queued", logs: logs, sub: sub}, nil
}

// WatchQueued is a free log subscription operation binding the contract event 0x33ed2ea816789b9d12d676e2ed6d81ab2cccc83a440f9a2576d15858c810bea9.
//
// Solidity: event Queued(uint256 indexed index, uint256 indexed orderId, address indexed user, uint256 liqU, uint40 queuedAt)
func (_StakingNew *StakingNewFilterer) WatchQueued(opts *bind.WatchOpts, sink chan<- *StakingNewQueued, index []*big.Int, orderId []*big.Int, user []common.Address) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "Queued", indexRule, orderIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewQueued)
				if err := _StakingNew.contract.UnpackLog(event, "Queued", log); err != nil {
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

// ParseQueued is a log parse operation binding the contract event 0x33ed2ea816789b9d12d676e2ed6d81ab2cccc83a440f9a2576d15858c810bea9.
//
// Solidity: event Queued(uint256 indexed index, uint256 indexed orderId, address indexed user, uint256 liqU, uint40 queuedAt)
func (_StakingNew *StakingNewFilterer) ParseQueued(log types.Log) (*StakingNewQueued, error) {
	event := new(StakingNewQueued)
	if err := _StakingNew.contract.UnpackLog(event, "Queued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewSystemsSetIterator is returned from FilterSystemsSet and is used to iterate over the raw logs and unpacked data for SystemsSet events raised by the StakingNew contract.
type StakingNewSystemsSetIterator struct {
	Event *StakingNewSystemsSet // Event containing the contract specifics and raw log

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
func (it *StakingNewSystemsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewSystemsSet)
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
		it.Event = new(StakingNewSystemsSet)
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
func (it *StakingNewSystemsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewSystemsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewSystemsSet represents a SystemsSet event raised by the StakingNew contract.
type StakingNewSystemsSet struct {
	Token  common.Address
	Usdt   common.Address
	Router common.Address
	User   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSystemsSet is a free log retrieval operation binding the contract event 0xe4a8014e10055a7765bf639e07cb54784070ef60af6e9501c4e53b49e064e6f9.
//
// Solidity: event SystemsSet(address indexed token, address indexed usdt, address indexed router, address user)
func (_StakingNew *StakingNewFilterer) FilterSystemsSet(opts *bind.FilterOpts, token []common.Address, usdt []common.Address, router []common.Address) (*StakingNewSystemsSetIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var usdtRule []interface{}
	for _, usdtItem := range usdt {
		usdtRule = append(usdtRule, usdtItem)
	}
	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "SystemsSet", tokenRule, usdtRule, routerRule)
	if err != nil {
		return nil, err
	}
	return &StakingNewSystemsSetIterator{contract: _StakingNew.contract, event: "SystemsSet", logs: logs, sub: sub}, nil
}

// WatchSystemsSet is a free log subscription operation binding the contract event 0xe4a8014e10055a7765bf639e07cb54784070ef60af6e9501c4e53b49e064e6f9.
//
// Solidity: event SystemsSet(address indexed token, address indexed usdt, address indexed router, address user)
func (_StakingNew *StakingNewFilterer) WatchSystemsSet(opts *bind.WatchOpts, sink chan<- *StakingNewSystemsSet, token []common.Address, usdt []common.Address, router []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var usdtRule []interface{}
	for _, usdtItem := range usdt {
		usdtRule = append(usdtRule, usdtItem)
	}
	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "SystemsSet", tokenRule, usdtRule, routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewSystemsSet)
				if err := _StakingNew.contract.UnpackLog(event, "SystemsSet", log); err != nil {
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

// ParseSystemsSet is a log parse operation binding the contract event 0xe4a8014e10055a7765bf639e07cb54784070ef60af6e9501c4e53b49e064e6f9.
//
// Solidity: event SystemsSet(address indexed token, address indexed usdt, address indexed router, address user)
func (_StakingNew *StakingNewFilterer) ParseSystemsSet(log types.Log) (*StakingNewSystemsSet, error) {
	event := new(StakingNewSystemsSet)
	if err := _StakingNew.contract.UnpackLog(event, "SystemsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewTeamBookedIterator is returned from FilterTeamBooked and is used to iterate over the raw logs and unpacked data for TeamBooked events raised by the StakingNew contract.
type StakingNewTeamBookedIterator struct {
	Event *StakingNewTeamBooked // Event containing the contract specifics and raw log

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
func (it *StakingNewTeamBookedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewTeamBooked)
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
		it.Event = new(StakingNewTeamBooked)
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
func (it *StakingNewTeamBookedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewTeamBookedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewTeamBooked represents a TeamBooked event raised by the StakingNew contract.
type StakingNewTeamBooked struct {
	From       common.Address
	To         common.Address
	Amount     *big.Int
	Stored     *big.Int
	DirectPaid *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTeamBooked is a free log retrieval operation binding the contract event 0x01dff65ae78c0348becfc40252f464a2b20d1f8dc39f0e9bf3ac159be0530271.
//
// Solidity: event TeamBooked(address indexed from, address indexed to, uint256 amount, uint256 stored, uint256 directPaid)
func (_StakingNew *StakingNewFilterer) FilterTeamBooked(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StakingNewTeamBookedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "TeamBooked", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StakingNewTeamBookedIterator{contract: _StakingNew.contract, event: "TeamBooked", logs: logs, sub: sub}, nil
}

// WatchTeamBooked is a free log subscription operation binding the contract event 0x01dff65ae78c0348becfc40252f464a2b20d1f8dc39f0e9bf3ac159be0530271.
//
// Solidity: event TeamBooked(address indexed from, address indexed to, uint256 amount, uint256 stored, uint256 directPaid)
func (_StakingNew *StakingNewFilterer) WatchTeamBooked(opts *bind.WatchOpts, sink chan<- *StakingNewTeamBooked, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "TeamBooked", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewTeamBooked)
				if err := _StakingNew.contract.UnpackLog(event, "TeamBooked", log); err != nil {
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

// ParseTeamBooked is a log parse operation binding the contract event 0x01dff65ae78c0348becfc40252f464a2b20d1f8dc39f0e9bf3ac159be0530271.
//
// Solidity: event TeamBooked(address indexed from, address indexed to, uint256 amount, uint256 stored, uint256 directPaid)
func (_StakingNew *StakingNewFilterer) ParseTeamBooked(log types.Log) (*StakingNewTeamBooked, error) {
	event := new(StakingNewTeamBooked)
	if err := _StakingNew.contract.UnpackLog(event, "TeamBooked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewTeamClaimedIterator is returned from FilterTeamClaimed and is used to iterate over the raw logs and unpacked data for TeamClaimed events raised by the StakingNew contract.
type StakingNewTeamClaimedIterator struct {
	Event *StakingNewTeamClaimed // Event containing the contract specifics and raw log

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
func (it *StakingNewTeamClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewTeamClaimed)
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
		it.Event = new(StakingNewTeamClaimed)
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
func (it *StakingNewTeamClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewTeamClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewTeamClaimed represents a TeamClaimed event raised by the StakingNew contract.
type StakingNewTeamClaimed struct {
	User  common.Address
	Gross *big.Int
	Fee   *big.Int
	Net   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTeamClaimed is a free log retrieval operation binding the contract event 0x14ec2ddfba4ceb55e4e9ddd496f4328d615049242a5d308552b14af02df83cfc.
//
// Solidity: event TeamClaimed(address indexed user, uint256 gross, uint256 fee, uint256 net)
func (_StakingNew *StakingNewFilterer) FilterTeamClaimed(opts *bind.FilterOpts, user []common.Address) (*StakingNewTeamClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "TeamClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingNewTeamClaimedIterator{contract: _StakingNew.contract, event: "TeamClaimed", logs: logs, sub: sub}, nil
}

// WatchTeamClaimed is a free log subscription operation binding the contract event 0x14ec2ddfba4ceb55e4e9ddd496f4328d615049242a5d308552b14af02df83cfc.
//
// Solidity: event TeamClaimed(address indexed user, uint256 gross, uint256 fee, uint256 net)
func (_StakingNew *StakingNewFilterer) WatchTeamClaimed(opts *bind.WatchOpts, sink chan<- *StakingNewTeamClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "TeamClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewTeamClaimed)
				if err := _StakingNew.contract.UnpackLog(event, "TeamClaimed", log); err != nil {
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

// ParseTeamClaimed is a log parse operation binding the contract event 0x14ec2ddfba4ceb55e4e9ddd496f4328d615049242a5d308552b14af02df83cfc.
//
// Solidity: event TeamClaimed(address indexed user, uint256 gross, uint256 fee, uint256 net)
func (_StakingNew *StakingNewFilterer) ParseTeamClaimed(log types.Log) (*StakingNewTeamClaimed, error) {
	event := new(StakingNewTeamClaimed)
	if err := _StakingNew.contract.UnpackLog(event, "TeamClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewTeamClearConfigSetIterator is returned from FilterTeamClearConfigSet and is used to iterate over the raw logs and unpacked data for TeamClearConfigSet events raised by the StakingNew contract.
type StakingNewTeamClearConfigSetIterator struct {
	Event *StakingNewTeamClearConfigSet // Event containing the contract specifics and raw log

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
func (it *StakingNewTeamClearConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewTeamClearConfigSet)
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
		it.Event = new(StakingNewTeamClearConfigSet)
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
func (it *StakingNewTeamClearConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewTeamClearConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewTeamClearConfigSet represents a TeamClearConfigSet event raised by the StakingNew contract.
type StakingNewTeamClearConfigSet struct {
	Period *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTeamClearConfigSet is a free log retrieval operation binding the contract event 0x6d987cdd03d78990879151a08be0c14ffbb343036956b236bcdb146af7c50f5b.
//
// Solidity: event TeamClearConfigSet(uint40 period)
func (_StakingNew *StakingNewFilterer) FilterTeamClearConfigSet(opts *bind.FilterOpts) (*StakingNewTeamClearConfigSetIterator, error) {

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "TeamClearConfigSet")
	if err != nil {
		return nil, err
	}
	return &StakingNewTeamClearConfigSetIterator{contract: _StakingNew.contract, event: "TeamClearConfigSet", logs: logs, sub: sub}, nil
}

// WatchTeamClearConfigSet is a free log subscription operation binding the contract event 0x6d987cdd03d78990879151a08be0c14ffbb343036956b236bcdb146af7c50f5b.
//
// Solidity: event TeamClearConfigSet(uint40 period)
func (_StakingNew *StakingNewFilterer) WatchTeamClearConfigSet(opts *bind.WatchOpts, sink chan<- *StakingNewTeamClearConfigSet) (event.Subscription, error) {

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "TeamClearConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewTeamClearConfigSet)
				if err := _StakingNew.contract.UnpackLog(event, "TeamClearConfigSet", log); err != nil {
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

// ParseTeamClearConfigSet is a log parse operation binding the contract event 0x6d987cdd03d78990879151a08be0c14ffbb343036956b236bcdb146af7c50f5b.
//
// Solidity: event TeamClearConfigSet(uint40 period)
func (_StakingNew *StakingNewFilterer) ParseTeamClearConfigSet(log types.Log) (*StakingNewTeamClearConfigSet, error) {
	event := new(StakingNewTeamClearConfigSet)
	if err := _StakingNew.contract.UnpackLog(event, "TeamClearConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewTeamClearedIterator is returned from FilterTeamCleared and is used to iterate over the raw logs and unpacked data for TeamCleared events raised by the StakingNew contract.
type StakingNewTeamClearedIterator struct {
	Event *StakingNewTeamCleared // Event containing the contract specifics and raw log

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
func (it *StakingNewTeamClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewTeamCleared)
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
		it.Event = new(StakingNewTeamCleared)
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
func (it *StakingNewTeamClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewTeamClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewTeamCleared represents a TeamCleared event raised by the StakingNew contract.
type StakingNewTeamCleared struct {
	User common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTeamCleared is a free log retrieval operation binding the contract event 0x6cd33091c1b8b89fe065da03791b522fd297c913ce6ef0c1ede837ea02972402.
//
// Solidity: event TeamCleared(address indexed user, uint40 time)
func (_StakingNew *StakingNewFilterer) FilterTeamCleared(opts *bind.FilterOpts, user []common.Address) (*StakingNewTeamClearedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "TeamCleared", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingNewTeamClearedIterator{contract: _StakingNew.contract, event: "TeamCleared", logs: logs, sub: sub}, nil
}

// WatchTeamCleared is a free log subscription operation binding the contract event 0x6cd33091c1b8b89fe065da03791b522fd297c913ce6ef0c1ede837ea02972402.
//
// Solidity: event TeamCleared(address indexed user, uint40 time)
func (_StakingNew *StakingNewFilterer) WatchTeamCleared(opts *bind.WatchOpts, sink chan<- *StakingNewTeamCleared, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "TeamCleared", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewTeamCleared)
				if err := _StakingNew.contract.UnpackLog(event, "TeamCleared", log); err != nil {
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

// ParseTeamCleared is a log parse operation binding the contract event 0x6cd33091c1b8b89fe065da03791b522fd297c913ce6ef0c1ede837ea02972402.
//
// Solidity: event TeamCleared(address indexed user, uint40 time)
func (_StakingNew *StakingNewFilterer) ParseTeamCleared(log types.Log) (*StakingNewTeamCleared, error) {
	event := new(StakingNewTeamCleared)
	if err := _StakingNew.contract.UnpackLog(event, "TeamCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewTeamExpiredIterator is returned from FilterTeamExpired and is used to iterate over the raw logs and unpacked data for TeamExpired events raised by the StakingNew contract.
type StakingNewTeamExpiredIterator struct {
	Event *StakingNewTeamExpired // Event containing the contract specifics and raw log

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
func (it *StakingNewTeamExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewTeamExpired)
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
		it.Event = new(StakingNewTeamExpired)
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
func (it *StakingNewTeamExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewTeamExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewTeamExpired represents a TeamExpired event raised by the StakingNew contract.
type StakingNewTeamExpired struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTeamExpired is a free log retrieval operation binding the contract event 0x7511616f0998e2297cc1254ccaf6ec14b6223f54f36ec224cb7c2a6516ada922.
//
// Solidity: event TeamExpired(address indexed from, address indexed to, uint256 amount)
func (_StakingNew *StakingNewFilterer) FilterTeamExpired(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StakingNewTeamExpiredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "TeamExpired", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StakingNewTeamExpiredIterator{contract: _StakingNew.contract, event: "TeamExpired", logs: logs, sub: sub}, nil
}

// WatchTeamExpired is a free log subscription operation binding the contract event 0x7511616f0998e2297cc1254ccaf6ec14b6223f54f36ec224cb7c2a6516ada922.
//
// Solidity: event TeamExpired(address indexed from, address indexed to, uint256 amount)
func (_StakingNew *StakingNewFilterer) WatchTeamExpired(opts *bind.WatchOpts, sink chan<- *StakingNewTeamExpired, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingNew.contract.WatchLogs(opts, "TeamExpired", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingNewTeamExpired)
				if err := _StakingNew.contract.UnpackLog(event, "TeamExpired", log); err != nil {
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

// ParseTeamExpired is a log parse operation binding the contract event 0x7511616f0998e2297cc1254ccaf6ec14b6223f54f36ec224cb7c2a6516ada922.
//
// Solidity: event TeamExpired(address indexed from, address indexed to, uint256 amount)
func (_StakingNew *StakingNewFilterer) ParseTeamExpired(log types.Log) (*StakingNewTeamExpired, error) {
	event := new(StakingNewTeamExpired)
	if err := _StakingNew.contract.UnpackLog(event, "TeamExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingNewWalletsSetIterator is returned from FilterWalletsSet and is used to iterate over the raw logs and unpacked data for WalletsSet events raised by the StakingNew contract.
type StakingNewWalletsSetIterator struct {
	Event *StakingNewWalletsSet // Event containing the contract specifics and raw log

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
func (it *StakingNewWalletsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingNewWalletsSet)
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
		it.Event = new(StakingNewWalletsSet)
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
func (it *StakingNewWalletsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingNewWalletsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingNewWalletsSet represents a WalletsSet event raised by the StakingNew contract.
type StakingNewWalletsSet struct {
	Market common.Address
	Eco    common.Address
	Sink   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWalletsSet is a free log retrieval operation binding the contract event 0xb8ccc1ed0a6227eb93cf90b2b0005aa020bb991388d0be83ae37257fd10798fb.
//
// Solidity: event WalletsSet(address indexed market, address indexed eco, address indexed sink)
func (_StakingNew *StakingNewFilterer) FilterWalletsSet(opts *bind.FilterOpts, market []common.Address, eco []common.Address, sink []common.Address) (*StakingNewWalletsSetIterator, error) {

	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}
	var ecoRule []interface{}
	for _, ecoItem := range eco {
		ecoRule = append(ecoRule, ecoItem)
	}
	var sinkRule []interface{}
	for _, sinkItem := range sink {
		sinkRule = append(sinkRule, sinkItem)
	}

	logs, sub, err := _StakingNew.contract.FilterLogs(opts, "WalletsSet", marketRule, ecoRule, sinkRule)
	if err != nil {
		return nil, err
	}
	return &StakingNewWalletsSetIterator{contract: _StakingNew.contract, event: "WalletsSet", logs: logs, sub: sub}, nil
}
