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

// UserMetaData contains all meta data concerning the User contract.
var UserMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"root_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"init_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ALREADY_BOUND\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BAD_DEPTH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BAD_INDEX\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BAD_LEN\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BAD_PARENT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NO_AUTH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZERO_ADDR\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"parent\",\"type\":\"address\"}],\"name\":\"Bound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depth\",\"type\":\"uint256\"}],\"name\":\"DepthChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extraPerf\",\"type\":\"uint256\"}],\"name\":\"ExtraChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldInit\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newInit\",\"type\":\"address\"}],\"name\":\"InitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"perf\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"name\":\"RuleChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"StakeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stake\",\"type\":\"address\"}],\"name\":\"StakeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"UsersInited\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"basePerf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"baseStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"p\",\"type\":\"address\"}],\"name\":\"bind\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"childrenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"childrenOf\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"out\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"extraPerf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"parents\",\"type\":\"address[]\"}],\"name\":\"initBinds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stake_\",\"type\":\"address\"}],\"name\":\"initSystems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"parents\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"extraPerfs\",\"type\":\"uint256[]\"}],\"name\":\"initUsers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"levelOf\",\"outputs\":[{\"internalType\":\"int8\",\"name\":\"\",\"type\":\"int8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"onStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"onUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"parent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"perfOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"rateOf\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rules\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"perf\",\"type\":\"uint128\"},{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setDepth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"perfValue\",\"type\":\"uint256\"}],\"name\":\"setExtra\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newInit\",\"type\":\"address\"}],\"name\":\"setInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"perfNeed\",\"type\":\"uint128\"},{\"internalType\":\"uint16\",\"name\":\"rateBps\",\"type\":\"uint16\"}],\"name\":\"setRule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"stakeOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"uplines\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"out\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UserABI is the input ABI used to generate the binding from.
// Deprecated: Use UserMetaData.ABI instead.
var UserABI = UserMetaData.ABI

// User is an auto generated Go binding around an Ethereum contract.
type User struct {
	UserCaller     // Read-only binding to the contract
	UserTransactor // Write-only binding to the contract
	UserFilterer   // Log filterer for contract events
}

// UserCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserSession struct {
	Contract     *User             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserCallerSession struct {
	Contract *UserCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserTransactorSession struct {
	Contract     *UserTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserRaw struct {
	Contract *User // Generic contract binding to access the raw methods on
}

// UserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserCallerRaw struct {
	Contract *UserCaller // Generic read-only contract binding to access the raw methods on
}

// UserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserTransactorRaw struct {
	Contract *UserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUser creates a new instance of User, bound to a specific deployed contract.
func NewUser(address common.Address, backend bind.ContractBackend) (*User, error) {
	contract, err := bindUser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &User{UserCaller: UserCaller{contract: contract}, UserTransactor: UserTransactor{contract: contract}, UserFilterer: UserFilterer{contract: contract}}, nil
}

// NewUserCaller creates a new read-only instance of User, bound to a specific deployed contract.
func NewUserCaller(address common.Address, caller bind.ContractCaller) (*UserCaller, error) {
	contract, err := bindUser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserCaller{contract: contract}, nil
}

// NewUserTransactor creates a new write-only instance of User, bound to a specific deployed contract.
func NewUserTransactor(address common.Address, transactor bind.ContractTransactor) (*UserTransactor, error) {
	contract, err := bindUser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserTransactor{contract: contract}, nil
}

// NewUserFilterer creates a new log filterer instance of User, bound to a specific deployed contract.
func NewUserFilterer(address common.Address, filterer bind.ContractFilterer) (*UserFilterer, error) {
	contract, err := bindUser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserFilterer{contract: contract}, nil
}

// bindUser binds a generic wrapper to an already deployed contract.
func bindUser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _User.Contract.UserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _User.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.contract.Transact(opts, method, params...)
}

// BasePerf is a free data retrieval call binding the contract method 0xec8bc659.
//
// Solidity: function basePerf(address ) view returns(uint256)
func (_User *UserCaller) BasePerf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "basePerf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasePerf is a free data retrieval call binding the contract method 0xec8bc659.
//
// Solidity: function basePerf(address ) view returns(uint256)
func (_User *UserSession) BasePerf(arg0 common.Address) (*big.Int, error) {
	return _User.Contract.BasePerf(&_User.CallOpts, arg0)
}

// BasePerf is a free data retrieval call binding the contract method 0xec8bc659.
//
// Solidity: function basePerf(address ) view returns(uint256)
func (_User *UserCallerSession) BasePerf(arg0 common.Address) (*big.Int, error) {
	return _User.Contract.BasePerf(&_User.CallOpts, arg0)
}

// BaseStake is a free data retrieval call binding the contract method 0x873807de.
//
// Solidity: function baseStake(address ) view returns(uint256)
func (_User *UserCaller) BaseStake(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "baseStake", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseStake is a free data retrieval call binding the contract method 0x873807de.
//
// Solidity: function baseStake(address ) view returns(uint256)
func (_User *UserSession) BaseStake(arg0 common.Address) (*big.Int, error) {
	return _User.Contract.BaseStake(&_User.CallOpts, arg0)
}

// BaseStake is a free data retrieval call binding the contract method 0x873807de.
//
// Solidity: function baseStake(address ) view returns(uint256)
func (_User *UserCallerSession) BaseStake(arg0 common.Address) (*big.Int, error) {
	return _User.Contract.BaseStake(&_User.CallOpts, arg0)
}

// Bound is a free data retrieval call binding the contract method 0xea291942.
//
// Solidity: function bound(address ) view returns(bool)
func (_User *UserCaller) Bound(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "bound", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Bound is a free data retrieval call binding the contract method 0xea291942.
//
// Solidity: function bound(address ) view returns(bool)
func (_User *UserSession) Bound(arg0 common.Address) (bool, error) {
	return _User.Contract.Bound(&_User.CallOpts, arg0)
}

// Bound is a free data retrieval call binding the contract method 0xea291942.
//
// Solidity: function bound(address ) view returns(bool)
func (_User *UserCallerSession) Bound(arg0 common.Address) (bool, error) {
	return _User.Contract.Bound(&_User.CallOpts, arg0)
}

// ChildrenCount is a free data retrieval call binding the contract method 0x960b2969.
//
// Solidity: function childrenCount(address account) view returns(uint256)
func (_User *UserCaller) ChildrenCount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "childrenCount", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChildrenCount is a free data retrieval call binding the contract method 0x960b2969.
//
// Solidity: function childrenCount(address account) view returns(uint256)
func (_User *UserSession) ChildrenCount(account common.Address) (*big.Int, error) {
	return _User.Contract.ChildrenCount(&_User.CallOpts, account)
}

// ChildrenCount is a free data retrieval call binding the contract method 0x960b2969.
//
// Solidity: function childrenCount(address account) view returns(uint256)
func (_User *UserCallerSession) ChildrenCount(account common.Address) (*big.Int, error) {
	return _User.Contract.ChildrenCount(&_User.CallOpts, account)
}

// ChildrenOf is a free data retrieval call binding the contract method 0x2880c66b.
//
// Solidity: function childrenOf(address account, uint256 offset, uint256 limit) view returns(address[] out)
func (_User *UserCaller) ChildrenOf(opts *bind.CallOpts, account common.Address, offset *big.Int, limit *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "childrenOf", account, offset, limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ChildrenOf is a free data retrieval call binding the contract method 0x2880c66b.
//
// Solidity: function childrenOf(address account, uint256 offset, uint256 limit) view returns(address[] out)
func (_User *UserSession) ChildrenOf(account common.Address, offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _User.Contract.ChildrenOf(&_User.CallOpts, account, offset, limit)
}

// ChildrenOf is a free data retrieval call binding the contract method 0x2880c66b.
//
// Solidity: function childrenOf(address account, uint256 offset, uint256 limit) view returns(address[] out)
func (_User *UserCallerSession) ChildrenOf(account common.Address, offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _User.Contract.ChildrenOf(&_User.CallOpts, account, offset, limit)
}

// Depth is a free data retrieval call binding the contract method 0x631c56ef.
//
// Solidity: function depth() view returns(uint256)
func (_User *UserCaller) Depth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "depth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Depth is a free data retrieval call binding the contract method 0x631c56ef.
//
// Solidity: function depth() view returns(uint256)
func (_User *UserSession) Depth() (*big.Int, error) {
	return _User.Contract.Depth(&_User.CallOpts)
}

// Depth is a free data retrieval call binding the contract method 0x631c56ef.
//
// Solidity: function depth() view returns(uint256)
func (_User *UserCallerSession) Depth() (*big.Int, error) {
	return _User.Contract.Depth(&_User.CallOpts)
}

// ExtraPerf is a free data retrieval call binding the contract method 0xd60041b0.
//
// Solidity: function extraPerf(address ) view returns(uint256)
func (_User *UserCaller) ExtraPerf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "extraPerf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtraPerf is a free data retrieval call binding the contract method 0xd60041b0.
//
// Solidity: function extraPerf(address ) view returns(uint256)
func (_User *UserSession) ExtraPerf(arg0 common.Address) (*big.Int, error) {
	return _User.Contract.ExtraPerf(&_User.CallOpts, arg0)
}

// ExtraPerf is a free data retrieval call binding the contract method 0xd60041b0.
//
// Solidity: function extraPerf(address ) view returns(uint256)
func (_User *UserCallerSession) ExtraPerf(arg0 common.Address) (*big.Int, error) {
	return _User.Contract.ExtraPerf(&_User.CallOpts, arg0)
}

// Init is a free data retrieval call binding the contract method 0xe1c7392a.
//
// Solidity: function init() view returns(address)
func (_User *UserCaller) Init(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "init")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Init is a free data retrieval call binding the contract method 0xe1c7392a.
//
// Solidity: function init() view returns(address)
func (_User *UserSession) Init() (common.Address, error) {
	return _User.Contract.Init(&_User.CallOpts)
}

// Init is a free data retrieval call binding the contract method 0xe1c7392a.
//
// Solidity: function init() view returns(address)
func (_User *UserCallerSession) Init() (common.Address, error) {
	return _User.Contract.Init(&_User.CallOpts)
}

// LevelOf is a free data retrieval call binding the contract method 0x5c138c9d.
//
// Solidity: function levelOf(address account) view returns(int8)
func (_User *UserCaller) LevelOf(opts *bind.CallOpts, account common.Address) (int8, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "levelOf", account)

	if err != nil {
		return *new(int8), err
	}

	out0 := *abi.ConvertType(out[0], new(int8)).(*int8)

	return out0, err

}

// LevelOf is a free data retrieval call binding the contract method 0x5c138c9d.
//
// Solidity: function levelOf(address account) view returns(int8)
func (_User *UserSession) LevelOf(account common.Address) (int8, error) {
	return _User.Contract.LevelOf(&_User.CallOpts, account)
}

// LevelOf is a free data retrieval call binding the contract method 0x5c138c9d.
//
// Solidity: function levelOf(address account) view returns(int8)
func (_User *UserCallerSession) LevelOf(account common.Address) (int8, error) {
	return _User.Contract.LevelOf(&_User.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_User *UserCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_User *UserSession) Owner() (common.Address, error) {
	return _User.Contract.Owner(&_User.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_User *UserCallerSession) Owner() (common.Address, error) {
	return _User.Contract.Owner(&_User.CallOpts)
}

// Parent is a free data retrieval call binding the contract method 0xf1f9d8c9.
//
// Solidity: function parent(address ) view returns(address)
func (_User *UserCaller) Parent(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "parent", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Parent is a free data retrieval call binding the contract method 0xf1f9d8c9.
//
// Solidity: function parent(address ) view returns(address)
func (_User *UserSession) Parent(arg0 common.Address) (common.Address, error) {
	return _User.Contract.Parent(&_User.CallOpts, arg0)
}

// Parent is a free data retrieval call binding the contract method 0xf1f9d8c9.
//
// Solidity: function parent(address ) view returns(address)
func (_User *UserCallerSession) Parent(arg0 common.Address) (common.Address, error) {
	return _User.Contract.Parent(&_User.CallOpts, arg0)
}

// PerfOf is a free data retrieval call binding the contract method 0x97f0bec5.
//
// Solidity: function perfOf(address account) view returns(uint256)
func (_User *UserCaller) PerfOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "perfOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerfOf is a free data retrieval call binding the contract method 0x97f0bec5.
//
// Solidity: function perfOf(address account) view returns(uint256)
func (_User *UserSession) PerfOf(account common.Address) (*big.Int, error) {
	return _User.Contract.PerfOf(&_User.CallOpts, account)
}

// PerfOf is a free data retrieval call binding the contract method 0x97f0bec5.
//
// Solidity: function perfOf(address account) view returns(uint256)
func (_User *UserCallerSession) PerfOf(account common.Address) (*big.Int, error) {
	return _User.Contract.PerfOf(&_User.CallOpts, account)
}

// RateOf is a free data retrieval call binding the contract method 0xcbf09802.
//
// Solidity: function rateOf(address account) view returns(uint16)
func (_User *UserCaller) RateOf(opts *bind.CallOpts, account common.Address) (uint16, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "rateOf", account)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RateOf is a free data retrieval call binding the contract method 0xcbf09802.
//
// Solidity: function rateOf(address account) view returns(uint16)
func (_User *UserSession) RateOf(account common.Address) (uint16, error) {
	return _User.Contract.RateOf(&_User.CallOpts, account)
}

// RateOf is a free data retrieval call binding the contract method 0xcbf09802.
//
// Solidity: function rateOf(address account) view returns(uint16)
func (_User *UserCallerSession) RateOf(account common.Address) (uint16, error) {
	return _User.Contract.RateOf(&_User.CallOpts, account)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(address)
func (_User *UserCaller) Root(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "root")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(address)
func (_User *UserSession) Root() (common.Address, error) {
	return _User.Contract.Root(&_User.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(address)
func (_User *UserCallerSession) Root() (common.Address, error) {
	return _User.Contract.Root(&_User.CallOpts)
}

// Rules is a free data retrieval call binding the contract method 0x04d6ded4.
//
// Solidity: function rules(uint256 ) view returns(uint128 perf, uint16 rate)
func (_User *UserCaller) Rules(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Perf *big.Int
	Rate uint16
}, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "rules", arg0)

	outstruct := new(struct {
		Perf *big.Int
		Rate uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Perf = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Rate = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// Rules is a free data retrieval call binding the contract method 0x04d6ded4.
//
// Solidity: function rules(uint256 ) view returns(uint128 perf, uint16 rate)
func (_User *UserSession) Rules(arg0 *big.Int) (struct {
	Perf *big.Int
	Rate uint16
}, error) {
	return _User.Contract.Rules(&_User.CallOpts, arg0)
}

// Rules is a free data retrieval call binding the contract method 0x04d6ded4.
//
// Solidity: function rules(uint256 ) view returns(uint128 perf, uint16 rate)
func (_User *UserCallerSession) Rules(arg0 *big.Int) (struct {
	Perf *big.Int
	Rate uint16
}, error) {
	return _User.Contract.Rules(&_User.CallOpts, arg0)
}

// Stake is a free data retrieval call binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() view returns(address)
func (_User *UserCaller) Stake(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "stake")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stake is a free data retrieval call binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() view returns(address)
func (_User *UserSession) Stake() (common.Address, error) {
	return _User.Contract.Stake(&_User.CallOpts)
}

// Stake is a free data retrieval call binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() view returns(address)
func (_User *UserCallerSession) Stake() (common.Address, error) {
	return _User.Contract.Stake(&_User.CallOpts)
}

// StakeOf is a free data retrieval call binding the contract method 0x42623360.
//
// Solidity: function stakeOf(address account) view returns(uint256)
func (_User *UserCaller) StakeOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "stakeOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeOf is a free data retrieval call binding the contract method 0x42623360.
//
// Solidity: function stakeOf(address account) view returns(uint256)
func (_User *UserSession) StakeOf(account common.Address) (*big.Int, error) {
	return _User.Contract.StakeOf(&_User.CallOpts, account)
}

// StakeOf is a free data retrieval call binding the contract method 0x42623360.
//
// Solidity: function stakeOf(address account) view returns(uint256)
func (_User *UserCallerSession) StakeOf(account common.Address) (*big.Int, error) {
	return _User.Contract.StakeOf(&_User.CallOpts, account)
}

// Uplines is a free data retrieval call binding the contract method 0x7d186766.
//
// Solidity: function uplines(address account, uint256 limit) view returns(address[] out)
func (_User *UserCaller) Uplines(opts *bind.CallOpts, account common.Address, limit *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "uplines", account, limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Uplines is a free data retrieval call binding the contract method 0x7d186766.
//
// Solidity: function uplines(address account, uint256 limit) view returns(address[] out)
func (_User *UserSession) Uplines(account common.Address, limit *big.Int) ([]common.Address, error) {
	return _User.Contract.Uplines(&_User.CallOpts, account, limit)
}

// Uplines is a free data retrieval call binding the contract method 0x7d186766.
//
// Solidity: function uplines(address account, uint256 limit) view returns(address[] out)
func (_User *UserCallerSession) Uplines(account common.Address, limit *big.Int) ([]common.Address, error) {
	return _User.Contract.Uplines(&_User.CallOpts, account, limit)
}

// Bind is a paid mutator transaction binding the contract method 0x81bac14f.
//
// Solidity: function bind(address p) returns()
func (_User *UserTransactor) Bind(opts *bind.TransactOpts, p common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "bind", p)
}

// Bind is a paid mutator transaction binding the contract method 0x81bac14f.
//
// Solidity: function bind(address p) returns()
func (_User *UserSession) Bind(p common.Address) (*types.Transaction, error) {
	return _User.Contract.Bind(&_User.TransactOpts, p)
}

// Bind is a paid mutator transaction binding the contract method 0x81bac14f.
//
// Solidity: function bind(address p) returns()
func (_User *UserTransactorSession) Bind(p common.Address) (*types.Transaction, error) {
	return _User.Contract.Bind(&_User.TransactOpts, p)
}

// InitBinds is a paid mutator transaction binding the contract method 0xd884ba08.
//
// Solidity: function initBinds(address[] users, address[] parents) returns()
func (_User *UserTransactor) InitBinds(opts *bind.TransactOpts, users []common.Address, parents []common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "initBinds", users, parents)
}

// InitBinds is a paid mutator transaction binding the contract method 0xd884ba08.
//
// Solidity: function initBinds(address[] users, address[] parents) returns()
func (_User *UserSession) InitBinds(users []common.Address, parents []common.Address) (*types.Transaction, error) {
	return _User.Contract.InitBinds(&_User.TransactOpts, users, parents)
}

// InitBinds is a paid mutator transaction binding the contract method 0xd884ba08.
//
// Solidity: function initBinds(address[] users, address[] parents) returns()
func (_User *UserTransactorSession) InitBinds(users []common.Address, parents []common.Address) (*types.Transaction, error) {
	return _User.Contract.InitBinds(&_User.TransactOpts, users, parents)
}

// InitSystems is a paid mutator transaction binding the contract method 0xbea36441.
//
// Solidity: function initSystems(address stake_) returns()
func (_User *UserTransactor) InitSystems(opts *bind.TransactOpts, stake_ common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "initSystems", stake_)
}

// InitSystems is a paid mutator transaction binding the contract method 0xbea36441.
//
// Solidity: function initSystems(address stake_) returns()
func (_User *UserSession) InitSystems(stake_ common.Address) (*types.Transaction, error) {
	return _User.Contract.InitSystems(&_User.TransactOpts, stake_)
}

// InitSystems is a paid mutator transaction binding the contract method 0xbea36441.
//
// Solidity: function initSystems(address stake_) returns()
func (_User *UserTransactorSession) InitSystems(stake_ common.Address) (*types.Transaction, error) {
	return _User.Contract.InitSystems(&_User.TransactOpts, stake_)
}

// InitUsers is a paid mutator transaction binding the contract method 0xb666c0e6.
//
// Solidity: function initUsers(address[] users, address[] parents, uint256[] extraPerfs) returns()
func (_User *UserTransactor) InitUsers(opts *bind.TransactOpts, users []common.Address, parents []common.Address, extraPerfs []*big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "initUsers", users, parents, extraPerfs)
}

// InitUsers is a paid mutator transaction binding the contract method 0xb666c0e6.
//
// Solidity: function initUsers(address[] users, address[] parents, uint256[] extraPerfs) returns()
func (_User *UserSession) InitUsers(users []common.Address, parents []common.Address, extraPerfs []*big.Int) (*types.Transaction, error) {
	return _User.Contract.InitUsers(&_User.TransactOpts, users, parents, extraPerfs)
}

// InitUsers is a paid mutator transaction binding the contract method 0xb666c0e6.
//
// Solidity: function initUsers(address[] users, address[] parents, uint256[] extraPerfs) returns()
func (_User *UserTransactorSession) InitUsers(users []common.Address, parents []common.Address, extraPerfs []*big.Int) (*types.Transaction, error) {
	return _User.Contract.InitUsers(&_User.TransactOpts, users, parents, extraPerfs)
}

// OnStake is a paid mutator transaction binding the contract method 0x8c087b1c.
//
// Solidity: function onStake(address account, uint256 amount) returns()
func (_User *UserTransactor) OnStake(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "onStake", account, amount)
}

// OnStake is a paid mutator transaction binding the contract method 0x8c087b1c.
//
// Solidity: function onStake(address account, uint256 amount) returns()
func (_User *UserSession) OnStake(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _User.Contract.OnStake(&_User.TransactOpts, account, amount)
}

// OnStake is a paid mutator transaction binding the contract method 0x8c087b1c.
//
// Solidity: function onStake(address account, uint256 amount) returns()
func (_User *UserTransactorSession) OnStake(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _User.Contract.OnStake(&_User.TransactOpts, account, amount)
}

// OnUnstake is a paid mutator transaction binding the contract method 0x7fb6d3a3.
//
// Solidity: function onUnstake(address account, uint256 amount) returns()
func (_User *UserTransactor) OnUnstake(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "onUnstake", account, amount)
}

// OnUnstake is a paid mutator transaction binding the contract method 0x7fb6d3a3.
//
// Solidity: function onUnstake(address account, uint256 amount) returns()
func (_User *UserSession) OnUnstake(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _User.Contract.OnUnstake(&_User.TransactOpts, account, amount)
}

// OnUnstake is a paid mutator transaction binding the contract method 0x7fb6d3a3.
//
// Solidity: function onUnstake(address account, uint256 amount) returns()
func (_User *UserTransactorSession) OnUnstake(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _User.Contract.OnUnstake(&_User.TransactOpts, account, amount)
}

// SetDepth is a paid mutator transaction binding the contract method 0x641db30e.
//
// Solidity: function setDepth(uint256 value) returns()
func (_User *UserTransactor) SetDepth(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "setDepth", value)
}

// SetDepth is a paid mutator transaction binding the contract method 0x641db30e.
//
// Solidity: function setDepth(uint256 value) returns()
func (_User *UserSession) SetDepth(value *big.Int) (*types.Transaction, error) {
	return _User.Contract.SetDepth(&_User.TransactOpts, value)
}

// SetDepth is a paid mutator transaction binding the contract method 0x641db30e.
//
// Solidity: function setDepth(uint256 value) returns()
func (_User *UserTransactorSession) SetDepth(value *big.Int) (*types.Transaction, error) {
	return _User.Contract.SetDepth(&_User.TransactOpts, value)
}

// SetExtra is a paid mutator transaction binding the contract method 0x83eb5d70.
//
// Solidity: function setExtra(address account, uint256 perfValue) returns()
func (_User *UserTransactor) SetExtra(opts *bind.TransactOpts, account common.Address, perfValue *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "setExtra", account, perfValue)
}

// SetExtra is a paid mutator transaction binding the contract method 0x83eb5d70.
//
// Solidity: function setExtra(address account, uint256 perfValue) returns()
func (_User *UserSession) SetExtra(account common.Address, perfValue *big.Int) (*types.Transaction, error) {
	return _User.Contract.SetExtra(&_User.TransactOpts, account, perfValue)
}

// SetExtra is a paid mutator transaction binding the contract method 0x83eb5d70.
//
// Solidity: function setExtra(address account, uint256 perfValue) returns()
func (_User *UserTransactorSession) SetExtra(account common.Address, perfValue *big.Int) (*types.Transaction, error) {
	return _User.Contract.SetExtra(&_User.TransactOpts, account, perfValue)
}

// SetInit is a paid mutator transaction binding the contract method 0x9784f55b.
//
// Solidity: function setInit(address newInit) returns()
func (_User *UserTransactor) SetInit(opts *bind.TransactOpts, newInit common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "setInit", newInit)
}

// SetInit is a paid mutator transaction binding the contract method 0x9784f55b.
//
// Solidity: function setInit(address newInit) returns()
func (_User *UserSession) SetInit(newInit common.Address) (*types.Transaction, error) {
	return _User.Contract.SetInit(&_User.TransactOpts, newInit)
}

// SetInit is a paid mutator transaction binding the contract method 0x9784f55b.
//
// Solidity: function setInit(address newInit) returns()
func (_User *UserTransactorSession) SetInit(newInit common.Address) (*types.Transaction, error) {
	return _User.Contract.SetInit(&_User.TransactOpts, newInit)
}

// SetRule is a paid mutator transaction binding the contract method 0x304ab654.
//
// Solidity: function setRule(uint8 index, uint128 perfNeed, uint16 rateBps) returns()
func (_User *UserTransactor) SetRule(opts *bind.TransactOpts, index uint8, perfNeed *big.Int, rateBps uint16) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "setRule", index, perfNeed, rateBps)
}

// SetRule is a paid mutator transaction binding the contract method 0x304ab654.
//
// Solidity: function setRule(uint8 index, uint128 perfNeed, uint16 rateBps) returns()
func (_User *UserSession) SetRule(index uint8, perfNeed *big.Int, rateBps uint16) (*types.Transaction, error) {
	return _User.Contract.SetRule(&_User.TransactOpts, index, perfNeed, rateBps)
}

// SetRule is a paid mutator transaction binding the contract method 0x304ab654.
//
// Solidity: function setRule(uint8 index, uint128 perfNeed, uint16 rateBps) returns()
func (_User *UserTransactorSession) SetRule(index uint8, perfNeed *big.Int, rateBps uint16) (*types.Transaction, error) {
	return _User.Contract.SetRule(&_User.TransactOpts, index, perfNeed, rateBps)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_User *UserTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_User *UserSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _User.Contract.TransferOwnership(&_User.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_User *UserTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _User.Contract.TransferOwnership(&_User.TransactOpts, newOwner)
}

// UserBoundIterator is returned from FilterBound and is used to iterate over the raw logs and unpacked data for Bound events raised by the User contract.
type UserBoundIterator struct {
	Event *UserBound // Event containing the contract specifics and raw log

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
func (it *UserBoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserBound)
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
		it.Event = new(UserBound)
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
func (it *UserBoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserBoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserBound represents a Bound event raised by the User contract.
type UserBound struct {
	User   common.Address
	Parent common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBound is a free log retrieval operation binding the contract event 0x0d128562eaa47ab89086803e64a0f96847c0ed3cc63c26251f29ba1aede09d4e.
//
// Solidity: event Bound(address indexed user, address indexed parent)
func (_User *UserFilterer) FilterBound(opts *bind.FilterOpts, user []common.Address, parent []common.Address) (*UserBoundIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var parentRule []interface{}
	for _, parentItem := range parent {
		parentRule = append(parentRule, parentItem)
	}

	logs, sub, err := _User.contract.FilterLogs(opts, "Bound", userRule, parentRule)
	if err != nil {
		return nil, err
	}
	return &UserBoundIterator{contract: _User.contract, event: "Bound", logs: logs, sub: sub}, nil
}

// WatchBound is a free log subscription operation binding the contract event 0x0d128562eaa47ab89086803e64a0f96847c0ed3cc63c26251f29ba1aede09d4e.
//
// Solidity: event Bound(address indexed user, address indexed parent)
func (_User *UserFilterer) WatchBound(opts *bind.WatchOpts, sink chan<- *UserBound, user []common.Address, parent []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var parentRule []interface{}
	for _, parentItem := range parent {
		parentRule = append(parentRule, parentItem)
	}

	logs, sub, err := _User.contract.WatchLogs(opts, "Bound", userRule, parentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserBound)
				if err := _User.contract.UnpackLog(event, "Bound", log); err != nil {
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

// ParseBound is a log parse operation binding the contract event 0x0d128562eaa47ab89086803e64a0f96847c0ed3cc63c26251f29ba1aede09d4e.
//
// Solidity: event Bound(address indexed user, address indexed parent)
func (_User *UserFilterer) ParseBound(log types.Log) (*UserBound, error) {
	event := new(UserBound)
	if err := _User.contract.UnpackLog(event, "Bound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserDepthChangedIterator is returned from FilterDepthChanged and is used to iterate over the raw logs and unpacked data for DepthChanged events raised by the User contract.
type UserDepthChangedIterator struct {
	Event *UserDepthChanged // Event containing the contract specifics and raw log

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
func (it *UserDepthChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserDepthChanged)
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
		it.Event = new(UserDepthChanged)
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
func (it *UserDepthChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserDepthChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserDepthChanged represents a DepthChanged event raised by the User contract.
type UserDepthChanged struct {
	Depth *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDepthChanged is a free log retrieval operation binding the contract event 0x4db838d00597b0867cd09a46f02cf67889495d7e63f476add3d5f81cf9d4b6ee.
//
// Solidity: event DepthChanged(uint256 depth)
func (_User *UserFilterer) FilterDepthChanged(opts *bind.FilterOpts) (*UserDepthChangedIterator, error) {

	logs, sub, err := _User.contract.FilterLogs(opts, "DepthChanged")
	if err != nil {
		return nil, err
	}
	return &UserDepthChangedIterator{contract: _User.contract, event: "DepthChanged", logs: logs, sub: sub}, nil
}

// WatchDepthChanged is a free log subscription operation binding the contract event 0x4db838d00597b0867cd09a46f02cf67889495d7e63f476add3d5f81cf9d4b6ee.
//
// Solidity: event DepthChanged(uint256 depth)
func (_User *UserFilterer) WatchDepthChanged(opts *bind.WatchOpts, sink chan<- *UserDepthChanged) (event.Subscription, error) {

	logs, sub, err := _User.contract.WatchLogs(opts, "DepthChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserDepthChanged)
				if err := _User.contract.UnpackLog(event, "DepthChanged", log); err != nil {
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

// ParseDepthChanged is a log parse operation binding the contract event 0x4db838d00597b0867cd09a46f02cf67889495d7e63f476add3d5f81cf9d4b6ee.
//
// Solidity: event DepthChanged(uint256 depth)
func (_User *UserFilterer) ParseDepthChanged(log types.Log) (*UserDepthChanged, error) {
	event := new(UserDepthChanged)
	if err := _User.contract.UnpackLog(event, "DepthChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserExtraChangedIterator is returned from FilterExtraChanged and is used to iterate over the raw logs and unpacked data for ExtraChanged events raised by the User contract.
type UserExtraChangedIterator struct {
	Event *UserExtraChanged // Event containing the contract specifics and raw log

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
func (it *UserExtraChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserExtraChanged)
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
		it.Event = new(UserExtraChanged)
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
func (it *UserExtraChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserExtraChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserExtraChanged represents a ExtraChanged event raised by the User contract.
type UserExtraChanged struct {
	Account   common.Address
	ExtraPerf *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExtraChanged is a free log retrieval operation binding the contract event 0xa1416bb110754a8cd5905ddb97926b43030460b0be1d031aba6830300d145a65.
//
// Solidity: event ExtraChanged(address indexed account, uint256 extraPerf)
func (_User *UserFilterer) FilterExtraChanged(opts *bind.FilterOpts, account []common.Address) (*UserExtraChangedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _User.contract.FilterLogs(opts, "ExtraChanged", accountRule)
	if err != nil {
		return nil, err
	}
	return &UserExtraChangedIterator{contract: _User.contract, event: "ExtraChanged", logs: logs, sub: sub}, nil
}

// WatchExtraChanged is a free log subscription operation binding the contract event 0xa1416bb110754a8cd5905ddb97926b43030460b0be1d031aba6830300d145a65.
//
// Solidity: event ExtraChanged(address indexed account, uint256 extraPerf)
func (_User *UserFilterer) WatchExtraChanged(opts *bind.WatchOpts, sink chan<- *UserExtraChanged, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _User.contract.WatchLogs(opts, "ExtraChanged", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserExtraChanged)
				if err := _User.contract.UnpackLog(event, "ExtraChanged", log); err != nil {
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

// ParseExtraChanged is a log parse operation binding the contract event 0xa1416bb110754a8cd5905ddb97926b43030460b0be1d031aba6830300d145a65.
//
// Solidity: event ExtraChanged(address indexed account, uint256 extraPerf)
func (_User *UserFilterer) ParseExtraChanged(log types.Log) (*UserExtraChanged, error) {
	event := new(UserExtraChanged)
	if err := _User.contract.UnpackLog(event, "ExtraChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserInitChangedIterator is returned from FilterInitChanged and is used to iterate over the raw logs and unpacked data for InitChanged events raised by the User contract.
type UserInitChangedIterator struct {
	Event *UserInitChanged // Event containing the contract specifics and raw log

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
func (it *UserInitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserInitChanged)
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
		it.Event = new(UserInitChanged)
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
func (it *UserInitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserInitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserInitChanged represents a InitChanged event raised by the User contract.
type UserInitChanged struct {
	OldInit common.Address
	NewInit common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitChanged is a free log retrieval operation binding the contract event 0x3f17984892df4d730451c12071a40f5396f4470a024466ba1237f3163632185a.
//
// Solidity: event InitChanged(address indexed oldInit, address indexed newInit)
func (_User *UserFilterer) FilterInitChanged(opts *bind.FilterOpts, oldInit []common.Address, newInit []common.Address) (*UserInitChangedIterator, error) {

	var oldInitRule []interface{}
	for _, oldInitItem := range oldInit {
		oldInitRule = append(oldInitRule, oldInitItem)
	}
	var newInitRule []interface{}
	for _, newInitItem := range newInit {
		newInitRule = append(newInitRule, newInitItem)
	}

	logs, sub, err := _User.contract.FilterLogs(opts, "InitChanged", oldInitRule, newInitRule)
	if err != nil {
		return nil, err
	}
	return &UserInitChangedIterator{contract: _User.contract, event: "InitChanged", logs: logs, sub: sub}, nil
}

// WatchInitChanged is a free log subscription operation binding the contract event 0x3f17984892df4d730451c12071a40f5396f4470a024466ba1237f3163632185a.
//
// Solidity: event InitChanged(address indexed oldInit, address indexed newInit)
func (_User *UserFilterer) WatchInitChanged(opts *bind.WatchOpts, sink chan<- *UserInitChanged, oldInit []common.Address, newInit []common.Address) (event.Subscription, error) {

	var oldInitRule []interface{}
	for _, oldInitItem := range oldInit {
		oldInitRule = append(oldInitRule, oldInitItem)
	}
	var newInitRule []interface{}
	for _, newInitItem := range newInit {
		newInitRule = append(newInitRule, newInitItem)
	}

	logs, sub, err := _User.contract.WatchLogs(opts, "InitChanged", oldInitRule, newInitRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserInitChanged)
				if err := _User.contract.UnpackLog(event, "InitChanged", log); err != nil {
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
func (_User *UserFilterer) ParseInitChanged(log types.Log) (*UserInitChanged, error) {
	event := new(UserInitChanged)
	if err := _User.contract.UnpackLog(event, "InitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the User contract.
type UserOwnershipTransferredIterator struct {
	Event *UserOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UserOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserOwnershipTransferred)
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
		it.Event = new(UserOwnershipTransferred)
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
func (it *UserOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserOwnershipTransferred represents a OwnershipTransferred event raised by the User contract.
type UserOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_User *UserFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*UserOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _User.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UserOwnershipTransferredIterator{contract: _User.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_User *UserFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UserOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _User.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserOwnershipTransferred)
				if err := _User.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_User *UserFilterer) ParseOwnershipTransferred(log types.Log) (*UserOwnershipTransferred, error) {
	event := new(UserOwnershipTransferred)
	if err := _User.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserRuleChangedIterator is returned from FilterRuleChanged and is used to iterate over the raw logs and unpacked data for RuleChanged events raised by the User contract.
type UserRuleChangedIterator struct {
	Event *UserRuleChanged // Event containing the contract specifics and raw log

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
func (it *UserRuleChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserRuleChanged)
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
		it.Event = new(UserRuleChanged)
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
func (it *UserRuleChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserRuleChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserRuleChanged represents a RuleChanged event raised by the User contract.
type UserRuleChanged struct {
	Index uint8
	Perf  *big.Int
	Rate  uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRuleChanged is a free log retrieval operation binding the contract event 0x0c4a2357ed5a611f3985b5887bf2a1f7bbb85f288b07fb4c2f18ab07f70eed7a.
//
// Solidity: event RuleChanged(uint8 indexed index, uint128 perf, uint16 rate)
func (_User *UserFilterer) FilterRuleChanged(opts *bind.FilterOpts, index []uint8) (*UserRuleChangedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _User.contract.FilterLogs(opts, "RuleChanged", indexRule)
	if err != nil {
		return nil, err
	}
	return &UserRuleChangedIterator{contract: _User.contract, event: "RuleChanged", logs: logs, sub: sub}, nil
}

// WatchRuleChanged is a free log subscription operation binding the contract event 0x0c4a2357ed5a611f3985b5887bf2a1f7bbb85f288b07fb4c2f18ab07f70eed7a.
//
// Solidity: event RuleChanged(uint8 indexed index, uint128 perf, uint16 rate)
func (_User *UserFilterer) WatchRuleChanged(opts *bind.WatchOpts, sink chan<- *UserRuleChanged, index []uint8) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _User.contract.WatchLogs(opts, "RuleChanged", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserRuleChanged)
				if err := _User.contract.UnpackLog(event, "RuleChanged", log); err != nil {
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

// ParseRuleChanged is a log parse operation binding the contract event 0x0c4a2357ed5a611f3985b5887bf2a1f7bbb85f288b07fb4c2f18ab07f70eed7a.
//
// Solidity: event RuleChanged(uint8 indexed index, uint128 perf, uint16 rate)
func (_User *UserFilterer) ParseRuleChanged(log types.Log) (*UserRuleChanged, error) {
	event := new(UserRuleChanged)
	if err := _User.contract.UnpackLog(event, "RuleChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserStakeChangedIterator is returned from FilterStakeChanged and is used to iterate over the raw logs and unpacked data for StakeChanged events raised by the User contract.
type UserStakeChangedIterator struct {
	Event *UserStakeChanged // Event containing the contract specifics and raw log

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
func (it *UserStakeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserStakeChanged)
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
		it.Event = new(UserStakeChanged)
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
func (it *UserStakeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserStakeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserStakeChanged represents a StakeChanged event raised by the User contract.
type UserStakeChanged struct {
	User   common.Address
	Amount *big.Int
	Add    bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeChanged is a free log retrieval operation binding the contract event 0x5c2b1af9817829e3745bfe7eb801e744beec39c89a0215f43f7304b3276f6120.
//
// Solidity: event StakeChanged(address indexed user, uint256 amount, bool add)
func (_User *UserFilterer) FilterStakeChanged(opts *bind.FilterOpts, user []common.Address) (*UserStakeChangedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _User.contract.FilterLogs(opts, "StakeChanged", userRule)
	if err != nil {
		return nil, err
	}
	return &UserStakeChangedIterator{contract: _User.contract, event: "StakeChanged", logs: logs, sub: sub}, nil
}

// WatchStakeChanged is a free log subscription operation binding the contract event 0x5c2b1af9817829e3745bfe7eb801e744beec39c89a0215f43f7304b3276f6120.
//
// Solidity: event StakeChanged(address indexed user, uint256 amount, bool add)
func (_User *UserFilterer) WatchStakeChanged(opts *bind.WatchOpts, sink chan<- *UserStakeChanged, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _User.contract.WatchLogs(opts, "StakeChanged", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserStakeChanged)
				if err := _User.contract.UnpackLog(event, "StakeChanged", log); err != nil {
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

// ParseStakeChanged is a log parse operation binding the contract event 0x5c2b1af9817829e3745bfe7eb801e744beec39c89a0215f43f7304b3276f6120.
//
// Solidity: event StakeChanged(address indexed user, uint256 amount, bool add)
func (_User *UserFilterer) ParseStakeChanged(log types.Log) (*UserStakeChanged, error) {
	event := new(UserStakeChanged)
	if err := _User.contract.UnpackLog(event, "StakeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserStakeSetIterator is returned from FilterStakeSet and is used to iterate over the raw logs and unpacked data for StakeSet events raised by the User contract.
type UserStakeSetIterator struct {
	Event *UserStakeSet // Event containing the contract specifics and raw log

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
func (it *UserStakeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserStakeSet)
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
		it.Event = new(UserStakeSet)
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
func (it *UserStakeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserStakeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserStakeSet represents a StakeSet event raised by the User contract.
type UserStakeSet struct {
	Stake common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStakeSet is a free log retrieval operation binding the contract event 0x5ba1cdab2363f96401f6ecb909a002f65cfcbb5ee23f743aa584a70b630b9052.
//
// Solidity: event StakeSet(address indexed stake)
func (_User *UserFilterer) FilterStakeSet(opts *bind.FilterOpts, stake []common.Address) (*UserStakeSetIterator, error) {

	var stakeRule []interface{}
	for _, stakeItem := range stake {
		stakeRule = append(stakeRule, stakeItem)
	}

	logs, sub, err := _User.contract.FilterLogs(opts, "StakeSet", stakeRule)
	if err != nil {
		return nil, err
	}
	return &UserStakeSetIterator{contract: _User.contract, event: "StakeSet", logs: logs, sub: sub}, nil
}

// WatchStakeSet is a free log subscription operation binding the contract event 0x5ba1cdab2363f96401f6ecb909a002f65cfcbb5ee23f743aa584a70b630b9052.
//
// Solidity: event StakeSet(address indexed stake)
func (_User *UserFilterer) WatchStakeSet(opts *bind.WatchOpts, sink chan<- *UserStakeSet, stake []common.Address) (event.Subscription, error) {

	var stakeRule []interface{}
	for _, stakeItem := range stake {
		stakeRule = append(stakeRule, stakeItem)
	}

	logs, sub, err := _User.contract.WatchLogs(opts, "StakeSet", stakeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserStakeSet)
				if err := _User.contract.UnpackLog(event, "StakeSet", log); err != nil {
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

// ParseStakeSet is a log parse operation binding the contract event 0x5ba1cdab2363f96401f6ecb909a002f65cfcbb5ee23f743aa584a70b630b9052.
//
// Solidity: event StakeSet(address indexed stake)
func (_User *UserFilterer) ParseStakeSet(log types.Log) (*UserStakeSet, error) {
	event := new(UserStakeSet)
	if err := _User.contract.UnpackLog(event, "StakeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserUsersInitedIterator is returned from FilterUsersInited and is used to iterate over the raw logs and unpacked data for UsersInited events raised by the User contract.
type UserUsersInitedIterator struct {
	Event *UserUsersInited // Event containing the contract specifics and raw log

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
func (it *UserUsersInitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserUsersInited)
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
		it.Event = new(UserUsersInited)
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
func (it *UserUsersInitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserUsersInitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserUsersInited represents a UsersInited event raised by the User contract.
type UserUsersInited struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUsersInited is a free log retrieval operation binding the contract event 0xdf10ac64b32d39a0a97a71786c64ee86bd7962478b433bebbb10ba5a80d3f6ee.
//
// Solidity: event UsersInited(uint256 count)
func (_User *UserFilterer) FilterUsersInited(opts *bind.FilterOpts) (*UserUsersInitedIterator, error) {

	logs, sub, err := _User.contract.FilterLogs(opts, "UsersInited")
	if err != nil {
		return nil, err
	}
	return &UserUsersInitedIterator{contract: _User.contract, event: "UsersInited", logs: logs, sub: sub}, nil
}

// WatchUsersInited is a free log subscription operation binding the contract event 0xdf10ac64b32d39a0a97a71786c64ee86bd7962478b433bebbb10ba5a80d3f6ee.
//
// Solidity: event UsersInited(uint256 count)
func (_User *UserFilterer) WatchUsersInited(opts *bind.WatchOpts, sink chan<- *UserUsersInited) (event.Subscription, error) {

	logs, sub, err := _User.contract.WatchLogs(opts, "UsersInited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserUsersInited)
				if err := _User.contract.UnpackLog(event, "UsersInited", log); err != nil {
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

// ParseUsersInited is a log parse operation binding the contract event 0xdf10ac64b32d39a0a97a71786c64ee86bd7962478b433bebbb10ba5a80d3f6ee.
//
// Solidity: event UsersInited(uint256 count)
func (_User *UserFilterer) ParseUsersInited(log types.Log) (*UserUsersInited, error) {
	event := new(UserUsersInited)
	if err := _User.contract.UnpackLog(event, "UsersInited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
