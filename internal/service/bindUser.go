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

// BindUserMetaData contains all meta data concerning the BindUser contract.
var BindUserMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"root_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initAccount_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"parent\",\"type\":\"address\"}],\"name\":\"BindReferral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dividendSystem\",\"type\":\"address\"}],\"name\":\"DividendSystemUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extraTeamPerformance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extraStakeAmount\",\"type\":\"uint256\"}],\"name\":\"ExtraPerformanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousInitAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newInitAccount\",\"type\":\"address\"}],\"name\":\"InitAccountTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakeSystem\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dividendSystem\",\"type\":\"address\"}],\"name\":\"InitSystemsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Level7Added\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Level7Removed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"rateBps\",\"type\":\"uint16\"}],\"name\":\"LevelRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rootAddress\",\"type\":\"address\"}],\"name\":\"RootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakeSystem\",\"type\":\"address\"}],\"name\":\"StakeSystemUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"basePerfValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"basePerfValueSon\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"baseStakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"parent\",\"type\":\"address\"}],\"name\":\"bindMyReferral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"parent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"bindReferral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"childrenCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"childrenOf\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"out\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dividendSystem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"extraPerfValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"extraStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"extraTeamPerfValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"firstS7UplineOrRoot\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getLevel7Slice\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"out\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getReferral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getReferralCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"depth\",\"type\":\"uint256\"}],\"name\":\"getReferrals\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"result\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getReferrers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"direct\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"indirect\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRootAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"parents\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"initBinds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakeSystem_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dividendSystem_\",\"type\":\"address\"}],\"name\":\"initSystems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isBindReferral\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"level7Count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"levelOf\",\"outputs\":[{\"internalType\":\"int8\",\"name\":\"level\",\"type\":\"int8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"levelRateBps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"levelRules\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"perfReq\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeReq\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxUplineDepth\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"onStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"onUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"parentOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"perfValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"perfValueSon\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"rateBpsOf\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"parent\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"extraTeamPerformance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extraPersonalStake\",\"type\":\"uint256\"}],\"name\":\"setExtraPerformance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"extraTeamPerformance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"extraPersonalStake\",\"type\":\"uint256\"}],\"name\":\"setExtraPerformanceFromDividend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newInitAccount\",\"type\":\"address\"}],\"name\":\"setInitAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"depth\",\"type\":\"uint16\"}],\"name\":\"setMaxUplineDepth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeSystem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"stakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BindUserABI is the input ABI used to generate the binding from.
// Deprecated: Use BindUserMetaData.ABI instead.
var BindUserABI = BindUserMetaData.ABI

// BindUser is an auto generated Go binding around an Ethereum contract.
type BindUser struct {
	BindUserCaller     // Read-only binding to the contract
	BindUserTransactor // Write-only binding to the contract
	BindUserFilterer   // Log filterer for contract events
}

// BindUserCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindUserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindUserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindUserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindUserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindUserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindUserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindUserSession struct {
	Contract     *BindUser         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindUserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindUserCallerSession struct {
	Contract *BindUserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindUserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindUserTransactorSession struct {
	Contract     *BindUserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindUserRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindUserRaw struct {
	Contract *BindUser // Generic contract binding to access the raw methods on
}

// BindUserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindUserCallerRaw struct {
	Contract *BindUserCaller // Generic read-only contract binding to access the raw methods on
}

// BindUserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindUserTransactorRaw struct {
	Contract *BindUserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindUser creates a new instance of BindUser, bound to a specific deployed contract.
func NewBindUser(address common.Address, backend bind.ContractBackend) (*BindUser, error) {
	contract, err := bindBindUser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BindUser{BindUserCaller: BindUserCaller{contract: contract}, BindUserTransactor: BindUserTransactor{contract: contract}, BindUserFilterer: BindUserFilterer{contract: contract}}, nil
}

// NewBindUserCaller creates a new read-only instance of BindUser, bound to a specific deployed contract.
func NewBindUserCaller(address common.Address, caller bind.ContractCaller) (*BindUserCaller, error) {
	contract, err := bindBindUser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindUserCaller{contract: contract}, nil
}

// NewBindUserTransactor creates a new write-only instance of BindUser, bound to a specific deployed contract.
func NewBindUserTransactor(address common.Address, transactor bind.ContractTransactor) (*BindUserTransactor, error) {
	contract, err := bindBindUser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindUserTransactor{contract: contract}, nil
}

// NewBindUserFilterer creates a new log filterer instance of BindUser, bound to a specific deployed contract.
func NewBindUserFilterer(address common.Address, filterer bind.ContractFilterer) (*BindUserFilterer, error) {
	contract, err := bindBindUser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindUserFilterer{contract: contract}, nil
}

// bindBindUser binds a generic wrapper to an already deployed contract.
func bindBindUser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BindUserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindUser *BindUserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindUser.Contract.BindUserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindUser *BindUserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindUser.Contract.BindUserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindUser *BindUserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindUser.Contract.BindUserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BindUser *BindUserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BindUser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BindUser *BindUserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BindUser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BindUser *BindUserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BindUser.Contract.contract.Transact(opts, method, params...)
}

// BasePerfValue is a free data retrieval call binding the contract method 0xf55417bd.
//
// Solidity: function basePerfValue(address account) view returns(uint256)
func (_BindUser *BindUserCaller) BasePerfValue(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "basePerfValue", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasePerfValue is a free data retrieval call binding the contract method 0xf55417bd.
//
// Solidity: function basePerfValue(address account) view returns(uint256)
func (_BindUser *BindUserSession) BasePerfValue(account common.Address) (*big.Int, error) {
	return _BindUser.Contract.BasePerfValue(&_BindUser.CallOpts, account)
}

// BasePerfValue is a free data retrieval call binding the contract method 0xf55417bd.
//
// Solidity: function basePerfValue(address account) view returns(uint256)
func (_BindUser *BindUserCallerSession) BasePerfValue(account common.Address) (*big.Int, error) {
	return _BindUser.Contract.BasePerfValue(&_BindUser.CallOpts, account)
}

// BasePerfValueSon is a free data retrieval call binding the contract method 0x0a1cb4fd.
//
// Solidity: function basePerfValueSon(address account) view returns(uint256)
func (_BindUser *BindUserCaller) BasePerfValueSon(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "basePerfValueSon", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasePerfValueSon is a free data retrieval call binding the contract method 0x0a1cb4fd.
//
// Solidity: function basePerfValueSon(address account) view returns(uint256)
func (_BindUser *BindUserSession) BasePerfValueSon(account common.Address) (*big.Int, error) {
	return _BindUser.Contract.BasePerfValueSon(&_BindUser.CallOpts, account)
}

// BasePerfValueSon is a free data retrieval call binding the contract method 0x0a1cb4fd.
//
// Solidity: function basePerfValueSon(address account) view returns(uint256)
func (_BindUser *BindUserCallerSession) BasePerfValueSon(account common.Address) (*big.Int, error) {
	return _BindUser.Contract.BasePerfValueSon(&_BindUser.CallOpts, account)
}

// BaseStakedAmount is a free data retrieval call binding the contract method 0x673d9494.
//
// Solidity: function baseStakedAmount(address account) view returns(uint256)
func (_BindUser *BindUserCaller) BaseStakedAmount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "baseStakedAmount", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseStakedAmount is a free data retrieval call binding the contract method 0x673d9494.
//
// Solidity: function baseStakedAmount(address account) view returns(uint256)
func (_BindUser *BindUserSession) BaseStakedAmount(account common.Address) (*big.Int, error) {
	return _BindUser.Contract.BaseStakedAmount(&_BindUser.CallOpts, account)
}

// BaseStakedAmount is a free data retrieval call binding the contract method 0x673d9494.
//
// Solidity: function baseStakedAmount(address account) view returns(uint256)
func (_BindUser *BindUserCallerSession) BaseStakedAmount(account common.Address) (*big.Int, error) {
	return _BindUser.Contract.BaseStakedAmount(&_BindUser.CallOpts, account)
}

// ChildrenCount is a free data retrieval call binding the contract method 0x960b2969.
//
// Solidity: function childrenCount(address account) view returns(uint256)
func (_BindUser *BindUserCaller) ChildrenCount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "childrenCount", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChildrenCount is a free data retrieval call binding the contract method 0x960b2969.
//
// Solidity: function childrenCount(address account) view returns(uint256)
func (_BindUser *BindUserSession) ChildrenCount(account common.Address) (*big.Int, error) {
	return _BindUser.Contract.ChildrenCount(&_BindUser.CallOpts, account)
}

// ChildrenCount is a free data retrieval call binding the contract method 0x960b2969.
//
// Solidity: function childrenCount(address account) view returns(uint256)
func (_BindUser *BindUserCallerSession) ChildrenCount(account common.Address) (*big.Int, error) {
	return _BindUser.Contract.ChildrenCount(&_BindUser.CallOpts, account)
}

// ChildrenOf is a free data retrieval call binding the contract method 0x2880c66b.
//
// Solidity: function childrenOf(address account, uint256 offset, uint256 limit) view returns(address[] out)
func (_BindUser *BindUserCaller) ChildrenOf(opts *bind.CallOpts, account common.Address, offset *big.Int, limit *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "childrenOf", account, offset, limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ChildrenOf is a free data retrieval call binding the contract method 0x2880c66b.
//
// Solidity: function childrenOf(address account, uint256 offset, uint256 limit) view returns(address[] out)
func (_BindUser *BindUserSession) ChildrenOf(account common.Address, offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _BindUser.Contract.ChildrenOf(&_BindUser.CallOpts, account, offset, limit)
}

// ChildrenOf is a free data retrieval call binding the contract method 0x2880c66b.
//
// Solidity: function childrenOf(address account, uint256 offset, uint256 limit) view returns(address[] out)
func (_BindUser *BindUserCallerSession) ChildrenOf(account common.Address, offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _BindUser.Contract.ChildrenOf(&_BindUser.CallOpts, account, offset, limit)
}

// DividendSystem is a free data retrieval call binding the contract method 0xc2511edb.
//
// Solidity: function dividendSystem() view returns(address)
func (_BindUser *BindUserCaller) DividendSystem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "dividendSystem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DividendSystem is a free data retrieval call binding the contract method 0xc2511edb.
//
// Solidity: function dividendSystem() view returns(address)
func (_BindUser *BindUserSession) DividendSystem() (common.Address, error) {
	return _BindUser.Contract.DividendSystem(&_BindUser.CallOpts)
}

// DividendSystem is a free data retrieval call binding the contract method 0xc2511edb.
//
// Solidity: function dividendSystem() view returns(address)
func (_BindUser *BindUserCallerSession) DividendSystem() (common.Address, error) {
	return _BindUser.Contract.DividendSystem(&_BindUser.CallOpts)
}

// ExtraPerfValue is a free data retrieval call binding the contract method 0xea4ba401.
//
// Solidity: function extraPerfValue(address account) view returns(uint256)
func (_BindUser *BindUserCaller) ExtraPerfValue(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "extraPerfValue", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtraPerfValue is a free data retrieval call binding the contract method 0xea4ba401.
//
// Solidity: function extraPerfValue(address account) view returns(uint256)
func (_BindUser *BindUserSession) ExtraPerfValue(account common.Address) (*big.Int, error) {
	return _BindUser.Contract.ExtraPerfValue(&_BindUser.CallOpts, account)
}

// ExtraPerfValue is a free data retrieval call binding the contract method 0xea4ba401.
//
// Solidity: function extraPerfValue(address account) view returns(uint256)
func (_BindUser *BindUserCallerSession) ExtraPerfValue(account common.Address) (*big.Int, error) {
	return _BindUser.Contract.ExtraPerfValue(&_BindUser.CallOpts, account)
}

// ExtraStakeAmount is a free data retrieval call binding the contract method 0x6135b815.
//
// Solidity: function extraStakeAmount(address ) view returns(uint256)
func (_BindUser *BindUserCaller) ExtraStakeAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "extraStakeAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtraStakeAmount is a free data retrieval call binding the contract method 0x6135b815.
//
// Solidity: function extraStakeAmount(address ) view returns(uint256)
func (_BindUser *BindUserSession) ExtraStakeAmount(arg0 common.Address) (*big.Int, error) {
	return _BindUser.Contract.ExtraStakeAmount(&_BindUser.CallOpts, arg0)
}

// ExtraStakeAmount is a free data retrieval call binding the contract method 0x6135b815.
//
// Solidity: function extraStakeAmount(address ) view returns(uint256)
func (_BindUser *BindUserCallerSession) ExtraStakeAmount(arg0 common.Address) (*big.Int, error) {
	return _BindUser.Contract.ExtraStakeAmount(&_BindUser.CallOpts, arg0)
}

// ExtraTeamPerfValue is a free data retrieval call binding the contract method 0x6f0d6400.
//
// Solidity: function extraTeamPerfValue(address ) view returns(uint256)
func (_BindUser *BindUserCaller) ExtraTeamPerfValue(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "extraTeamPerfValue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtraTeamPerfValue is a free data retrieval call binding the contract method 0x6f0d6400.
//
// Solidity: function extraTeamPerfValue(address ) view returns(uint256)
func (_BindUser *BindUserSession) ExtraTeamPerfValue(arg0 common.Address) (*big.Int, error) {
	return _BindUser.Contract.ExtraTeamPerfValue(&_BindUser.CallOpts, arg0)
}

// ExtraTeamPerfValue is a free data retrieval call binding the contract method 0x6f0d6400.
//
// Solidity: function extraTeamPerfValue(address ) view returns(uint256)
func (_BindUser *BindUserCallerSession) ExtraTeamPerfValue(arg0 common.Address) (*big.Int, error) {
	return _BindUser.Contract.ExtraTeamPerfValue(&_BindUser.CallOpts, arg0)
}

// FirstS7UplineOrRoot is a free data retrieval call binding the contract method 0x988775d0.
//
// Solidity: function firstS7UplineOrRoot(address user) view returns(address)
func (_BindUser *BindUserCaller) FirstS7UplineOrRoot(opts *bind.CallOpts, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "firstS7UplineOrRoot", user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FirstS7UplineOrRoot is a free data retrieval call binding the contract method 0x988775d0.
//
// Solidity: function firstS7UplineOrRoot(address user) view returns(address)
func (_BindUser *BindUserSession) FirstS7UplineOrRoot(user common.Address) (common.Address, error) {
	return _BindUser.Contract.FirstS7UplineOrRoot(&_BindUser.CallOpts, user)
}

// FirstS7UplineOrRoot is a free data retrieval call binding the contract method 0x988775d0.
//
// Solidity: function firstS7UplineOrRoot(address user) view returns(address)
func (_BindUser *BindUserCallerSession) FirstS7UplineOrRoot(user common.Address) (common.Address, error) {
	return _BindUser.Contract.FirstS7UplineOrRoot(&_BindUser.CallOpts, user)
}

// GetLevel7Slice is a free data retrieval call binding the contract method 0x9f341f05.
//
// Solidity: function getLevel7Slice(uint256 offset, uint256 limit) view returns(address[] out)
func (_BindUser *BindUserCaller) GetLevel7Slice(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "getLevel7Slice", offset, limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetLevel7Slice is a free data retrieval call binding the contract method 0x9f341f05.
//
// Solidity: function getLevel7Slice(uint256 offset, uint256 limit) view returns(address[] out)
func (_BindUser *BindUserSession) GetLevel7Slice(offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _BindUser.Contract.GetLevel7Slice(&_BindUser.CallOpts, offset, limit)
}

// GetLevel7Slice is a free data retrieval call binding the contract method 0x9f341f05.
//
// Solidity: function getLevel7Slice(uint256 offset, uint256 limit) view returns(address[] out)
func (_BindUser *BindUserCallerSession) GetLevel7Slice(offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _BindUser.Contract.GetLevel7Slice(&_BindUser.CallOpts, offset, limit)
}

// GetReferral is a free data retrieval call binding the contract method 0x3b0f0f2f.
//
// Solidity: function getReferral(address account) view returns(address)
func (_BindUser *BindUserCaller) GetReferral(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "getReferral", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReferral is a free data retrieval call binding the contract method 0x3b0f0f2f.
//
// Solidity: function getReferral(address account) view returns(address)
func (_BindUser *BindUserSession) GetReferral(account common.Address) (common.Address, error) {
	return _BindUser.Contract.GetReferral(&_BindUser.CallOpts, account)
}

// GetReferral is a free data retrieval call binding the contract method 0x3b0f0f2f.
//
// Solidity: function getReferral(address account) view returns(address)
func (_BindUser *BindUserCallerSession) GetReferral(account common.Address) (common.Address, error) {
	return _BindUser.Contract.GetReferral(&_BindUser.CallOpts, account)
}

// GetReferralCount is a free data retrieval call binding the contract method 0x24acbd69.
//
// Solidity: function getReferralCount(address account) view returns(uint256)
func (_BindUser *BindUserCaller) GetReferralCount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "getReferralCount", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReferralCount is a free data retrieval call binding the contract method 0x24acbd69.
//
// Solidity: function getReferralCount(address account) view returns(uint256)
func (_BindUser *BindUserSession) GetReferralCount(account common.Address) (*big.Int, error) {
	return _BindUser.Contract.GetReferralCount(&_BindUser.CallOpts, account)
}

// GetReferralCount is a free data retrieval call binding the contract method 0x24acbd69.
//
// Solidity: function getReferralCount(address account) view returns(uint256)
func (_BindUser *BindUserCallerSession) GetReferralCount(account common.Address) (*big.Int, error) {
	return _BindUser.Contract.GetReferralCount(&_BindUser.CallOpts, account)
}

// GetReferrals is a free data retrieval call binding the contract method 0x32a00d8e.
//
// Solidity: function getReferrals(address account, uint256 depth) view returns(address[] result)
func (_BindUser *BindUserCaller) GetReferrals(opts *bind.CallOpts, account common.Address, depth *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "getReferrals", account, depth)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReferrals is a free data retrieval call binding the contract method 0x32a00d8e.
//
// Solidity: function getReferrals(address account, uint256 depth) view returns(address[] result)
func (_BindUser *BindUserSession) GetReferrals(account common.Address, depth *big.Int) ([]common.Address, error) {
	return _BindUser.Contract.GetReferrals(&_BindUser.CallOpts, account, depth)
}

// GetReferrals is a free data retrieval call binding the contract method 0x32a00d8e.
//
// Solidity: function getReferrals(address account, uint256 depth) view returns(address[] result)
func (_BindUser *BindUserCallerSession) GetReferrals(account common.Address, depth *big.Int) ([]common.Address, error) {
	return _BindUser.Contract.GetReferrals(&_BindUser.CallOpts, account, depth)
}

// GetReferrers is a free data retrieval call binding the contract method 0xdf65845a.
//
// Solidity: function getReferrers(address user) view returns(address direct, address indirect)
func (_BindUser *BindUserCaller) GetReferrers(opts *bind.CallOpts, user common.Address) (struct {
	Direct   common.Address
	Indirect common.Address
}, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "getReferrers", user)

	outstruct := new(struct {
		Direct   common.Address
		Indirect common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Direct = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Indirect = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetReferrers is a free data retrieval call binding the contract method 0xdf65845a.
//
// Solidity: function getReferrers(address user) view returns(address direct, address indirect)
func (_BindUser *BindUserSession) GetReferrers(user common.Address) (struct {
	Direct   common.Address
	Indirect common.Address
}, error) {
	return _BindUser.Contract.GetReferrers(&_BindUser.CallOpts, user)
}

// GetReferrers is a free data retrieval call binding the contract method 0xdf65845a.
//
// Solidity: function getReferrers(address user) view returns(address direct, address indirect)
func (_BindUser *BindUserCallerSession) GetReferrers(user common.Address) (struct {
	Direct   common.Address
	Indirect common.Address
}, error) {
	return _BindUser.Contract.GetReferrers(&_BindUser.CallOpts, user)
}

// GetRootAddress is a free data retrieval call binding the contract method 0xf6d5660f.
//
// Solidity: function getRootAddress() view returns(address)
func (_BindUser *BindUserCaller) GetRootAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "getRootAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRootAddress is a free data retrieval call binding the contract method 0xf6d5660f.
//
// Solidity: function getRootAddress() view returns(address)
func (_BindUser *BindUserSession) GetRootAddress() (common.Address, error) {
	return _BindUser.Contract.GetRootAddress(&_BindUser.CallOpts)
}

// GetRootAddress is a free data retrieval call binding the contract method 0xf6d5660f.
//
// Solidity: function getRootAddress() view returns(address)
func (_BindUser *BindUserCallerSession) GetRootAddress() (common.Address, error) {
	return _BindUser.Contract.GetRootAddress(&_BindUser.CallOpts)
}

// InitAccount is a free data retrieval call binding the contract method 0x4ec1ce21.
//
// Solidity: function initAccount() view returns(address)
func (_BindUser *BindUserCaller) InitAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "initAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InitAccount is a free data retrieval call binding the contract method 0x4ec1ce21.
//
// Solidity: function initAccount() view returns(address)
func (_BindUser *BindUserSession) InitAccount() (common.Address, error) {
	return _BindUser.Contract.InitAccount(&_BindUser.CallOpts)
}

// InitAccount is a free data retrieval call binding the contract method 0x4ec1ce21.
//
// Solidity: function initAccount() view returns(address)
func (_BindUser *BindUserCallerSession) InitAccount() (common.Address, error) {
	return _BindUser.Contract.InitAccount(&_BindUser.CallOpts)
}

// IsBindReferral is a free data retrieval call binding the contract method 0xeec76e04.
//
// Solidity: function isBindReferral(address account) view returns(bool)
func (_BindUser *BindUserCaller) IsBindReferral(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "isBindReferral", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBindReferral is a free data retrieval call binding the contract method 0xeec76e04.
//
// Solidity: function isBindReferral(address account) view returns(bool)
func (_BindUser *BindUserSession) IsBindReferral(account common.Address) (bool, error) {
	return _BindUser.Contract.IsBindReferral(&_BindUser.CallOpts, account)
}

// IsBindReferral is a free data retrieval call binding the contract method 0xeec76e04.
//
// Solidity: function isBindReferral(address account) view returns(bool)
func (_BindUser *BindUserCallerSession) IsBindReferral(account common.Address) (bool, error) {
	return _BindUser.Contract.IsBindReferral(&_BindUser.CallOpts, account)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_BindUser *BindUserCaller) IsRegistered(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "isRegistered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_BindUser *BindUserSession) IsRegistered(arg0 common.Address) (bool, error) {
	return _BindUser.Contract.IsRegistered(&_BindUser.CallOpts, arg0)
}

// IsRegistered is a free data retrieval call binding the contract method 0xc3c5a547.
//
// Solidity: function isRegistered(address ) view returns(bool)
func (_BindUser *BindUserCallerSession) IsRegistered(arg0 common.Address) (bool, error) {
	return _BindUser.Contract.IsRegistered(&_BindUser.CallOpts, arg0)
}

// Level7Count is a free data retrieval call binding the contract method 0xef8a791f.
//
// Solidity: function level7Count() view returns(uint256)
func (_BindUser *BindUserCaller) Level7Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "level7Count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Level7Count is a free data retrieval call binding the contract method 0xef8a791f.
//
// Solidity: function level7Count() view returns(uint256)
func (_BindUser *BindUserSession) Level7Count() (*big.Int, error) {
	return _BindUser.Contract.Level7Count(&_BindUser.CallOpts)
}

// Level7Count is a free data retrieval call binding the contract method 0xef8a791f.
//
// Solidity: function level7Count() view returns(uint256)
func (_BindUser *BindUserCallerSession) Level7Count() (*big.Int, error) {
	return _BindUser.Contract.Level7Count(&_BindUser.CallOpts)
}

// LevelOf is a free data retrieval call binding the contract method 0x5c138c9d.
//
// Solidity: function levelOf(address account) view returns(int8 level)
func (_BindUser *BindUserCaller) LevelOf(opts *bind.CallOpts, account common.Address) (int8, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "levelOf", account)

	if err != nil {
		return *new(int8), err
	}

	out0 := *abi.ConvertType(out[0], new(int8)).(*int8)

	return out0, err

}

// LevelOf is a free data retrieval call binding the contract method 0x5c138c9d.
//
// Solidity: function levelOf(address account) view returns(int8 level)
func (_BindUser *BindUserSession) LevelOf(account common.Address) (int8, error) {
	return _BindUser.Contract.LevelOf(&_BindUser.CallOpts, account)
}

// LevelOf is a free data retrieval call binding the contract method 0x5c138c9d.
//
// Solidity: function levelOf(address account) view returns(int8 level)
func (_BindUser *BindUserCallerSession) LevelOf(account common.Address) (int8, error) {
	return _BindUser.Contract.LevelOf(&_BindUser.CallOpts, account)
}

// LevelRateBps is a free data retrieval call binding the contract method 0x9d44431a.
//
// Solidity: function levelRateBps(uint256 ) view returns(uint16)
func (_BindUser *BindUserCaller) LevelRateBps(opts *bind.CallOpts, arg0 *big.Int) (uint16, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "levelRateBps", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LevelRateBps is a free data retrieval call binding the contract method 0x9d44431a.
//
// Solidity: function levelRateBps(uint256 ) view returns(uint16)
func (_BindUser *BindUserSession) LevelRateBps(arg0 *big.Int) (uint16, error) {
	return _BindUser.Contract.LevelRateBps(&_BindUser.CallOpts, arg0)
}

// LevelRateBps is a free data retrieval call binding the contract method 0x9d44431a.
//
// Solidity: function levelRateBps(uint256 ) view returns(uint16)
func (_BindUser *BindUserCallerSession) LevelRateBps(arg0 *big.Int) (uint16, error) {
	return _BindUser.Contract.LevelRateBps(&_BindUser.CallOpts, arg0)
}

// LevelRules is a free data retrieval call binding the contract method 0xa48471b0.
//
// Solidity: function levelRules(uint256 ) view returns(uint256 perfReq, uint256 stakeReq)
func (_BindUser *BindUserCaller) LevelRules(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PerfReq  *big.Int
	StakeReq *big.Int
}, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "levelRules", arg0)

	outstruct := new(struct {
		PerfReq  *big.Int
		StakeReq *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PerfReq = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StakeReq = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LevelRules is a free data retrieval call binding the contract method 0xa48471b0.
//
// Solidity: function levelRules(uint256 ) view returns(uint256 perfReq, uint256 stakeReq)
func (_BindUser *BindUserSession) LevelRules(arg0 *big.Int) (struct {
	PerfReq  *big.Int
	StakeReq *big.Int
}, error) {
	return _BindUser.Contract.LevelRules(&_BindUser.CallOpts, arg0)
}

// LevelRules is a free data retrieval call binding the contract method 0xa48471b0.
//
// Solidity: function levelRules(uint256 ) view returns(uint256 perfReq, uint256 stakeReq)
func (_BindUser *BindUserCallerSession) LevelRules(arg0 *big.Int) (struct {
	PerfReq  *big.Int
	StakeReq *big.Int
}, error) {
	return _BindUser.Contract.LevelRules(&_BindUser.CallOpts, arg0)
}

// MaxUplineDepth is a free data retrieval call binding the contract method 0xf1690537.
//
// Solidity: function maxUplineDepth() view returns(uint16)
func (_BindUser *BindUserCaller) MaxUplineDepth(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "maxUplineDepth")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaxUplineDepth is a free data retrieval call binding the contract method 0xf1690537.
//
// Solidity: function maxUplineDepth() view returns(uint16)
func (_BindUser *BindUserSession) MaxUplineDepth() (uint16, error) {
	return _BindUser.Contract.MaxUplineDepth(&_BindUser.CallOpts)
}

// MaxUplineDepth is a free data retrieval call binding the contract method 0xf1690537.
//
// Solidity: function maxUplineDepth() view returns(uint16)
func (_BindUser *BindUserCallerSession) MaxUplineDepth() (uint16, error) {
	return _BindUser.Contract.MaxUplineDepth(&_BindUser.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BindUser *BindUserCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BindUser *BindUserSession) Owner() (common.Address, error) {
	return _BindUser.Contract.Owner(&_BindUser.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BindUser *BindUserCallerSession) Owner() (common.Address, error) {
	return _BindUser.Contract.Owner(&_BindUser.CallOpts)
}

// ParentOf is a free data retrieval call binding the contract method 0xee08388e.
//
// Solidity: function parentOf(address ) view returns(address)
func (_BindUser *BindUserCaller) ParentOf(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "parentOf", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParentOf is a free data retrieval call binding the contract method 0xee08388e.
//
// Solidity: function parentOf(address ) view returns(address)
func (_BindUser *BindUserSession) ParentOf(arg0 common.Address) (common.Address, error) {
	return _BindUser.Contract.ParentOf(&_BindUser.CallOpts, arg0)
}

// ParentOf is a free data retrieval call binding the contract method 0xee08388e.
//
// Solidity: function parentOf(address ) view returns(address)
func (_BindUser *BindUserCallerSession) ParentOf(arg0 common.Address) (common.Address, error) {
	return _BindUser.Contract.ParentOf(&_BindUser.CallOpts, arg0)
}

// PerfValue is a free data retrieval call binding the contract method 0xfaebd33d.
//
// Solidity: function perfValue(address account) view returns(uint256)
func (_BindUser *BindUserCaller) PerfValue(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "perfValue", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerfValue is a free data retrieval call binding the contract method 0xfaebd33d.
//
// Solidity: function perfValue(address account) view returns(uint256)
func (_BindUser *BindUserSession) PerfValue(account common.Address) (*big.Int, error) {
	return _BindUser.Contract.PerfValue(&_BindUser.CallOpts, account)
}

// PerfValue is a free data retrieval call binding the contract method 0xfaebd33d.
//
// Solidity: function perfValue(address account) view returns(uint256)
func (_BindUser *BindUserCallerSession) PerfValue(account common.Address) (*big.Int, error) {
	return _BindUser.Contract.PerfValue(&_BindUser.CallOpts, account)
}

// PerfValueSon is a free data retrieval call binding the contract method 0xf844f7ac.
//
// Solidity: function perfValueSon(address account) view returns(uint256)
func (_BindUser *BindUserCaller) PerfValueSon(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "perfValueSon", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerfValueSon is a free data retrieval call binding the contract method 0xf844f7ac.
//
// Solidity: function perfValueSon(address account) view returns(uint256)
func (_BindUser *BindUserSession) PerfValueSon(account common.Address) (*big.Int, error) {
	return _BindUser.Contract.PerfValueSon(&_BindUser.CallOpts, account)
}

// PerfValueSon is a free data retrieval call binding the contract method 0xf844f7ac.
//
// Solidity: function perfValueSon(address account) view returns(uint256)
func (_BindUser *BindUserCallerSession) PerfValueSon(account common.Address) (*big.Int, error) {
	return _BindUser.Contract.PerfValueSon(&_BindUser.CallOpts, account)
}

// RateBpsOf is a free data retrieval call binding the contract method 0xced76b95.
//
// Solidity: function rateBpsOf(address account) view returns(uint16)
func (_BindUser *BindUserCaller) RateBpsOf(opts *bind.CallOpts, account common.Address) (uint16, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "rateBpsOf", account)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RateBpsOf is a free data retrieval call binding the contract method 0xced76b95.
//
// Solidity: function rateBpsOf(address account) view returns(uint16)
func (_BindUser *BindUserSession) RateBpsOf(account common.Address) (uint16, error) {
	return _BindUser.Contract.RateBpsOf(&_BindUser.CallOpts, account)
}

// RateBpsOf is a free data retrieval call binding the contract method 0xced76b95.
//
// Solidity: function rateBpsOf(address account) view returns(uint16)
func (_BindUser *BindUserCallerSession) RateBpsOf(account common.Address) (uint16, error) {
	return _BindUser.Contract.RateBpsOf(&_BindUser.CallOpts, account)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(address)
func (_BindUser *BindUserCaller) Root(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "root")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(address)
func (_BindUser *BindUserSession) Root() (common.Address, error) {
	return _BindUser.Contract.Root(&_BindUser.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(address)
func (_BindUser *BindUserCallerSession) Root() (common.Address, error) {
	return _BindUser.Contract.Root(&_BindUser.CallOpts)
}

// StakeSystem is a free data retrieval call binding the contract method 0x32821765.
//
// Solidity: function stakeSystem() view returns(address)
func (_BindUser *BindUserCaller) StakeSystem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "stakeSystem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeSystem is a free data retrieval call binding the contract method 0x32821765.
//
// Solidity: function stakeSystem() view returns(address)
func (_BindUser *BindUserSession) StakeSystem() (common.Address, error) {
	return _BindUser.Contract.StakeSystem(&_BindUser.CallOpts)
}

// StakeSystem is a free data retrieval call binding the contract method 0x32821765.
//
// Solidity: function stakeSystem() view returns(address)
func (_BindUser *BindUserCallerSession) StakeSystem() (common.Address, error) {
	return _BindUser.Contract.StakeSystem(&_BindUser.CallOpts)
}

// StakedAmount is a free data retrieval call binding the contract method 0xf9931855.
//
// Solidity: function stakedAmount(address account) view returns(uint256)
func (_BindUser *BindUserCaller) StakedAmount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BindUser.contract.Call(opts, &out, "stakedAmount", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedAmount is a free data retrieval call binding the contract method 0xf9931855.
//
// Solidity: function stakedAmount(address account) view returns(uint256)
func (_BindUser *BindUserSession) StakedAmount(account common.Address) (*big.Int, error) {
	return _BindUser.Contract.StakedAmount(&_BindUser.CallOpts, account)
}

// StakedAmount is a free data retrieval call binding the contract method 0xf9931855.
//
// Solidity: function stakedAmount(address account) view returns(uint256)
func (_BindUser *BindUserCallerSession) StakedAmount(account common.Address) (*big.Int, error) {
	return _BindUser.Contract.StakedAmount(&_BindUser.CallOpts, account)
}

// BindMyReferral is a paid mutator transaction binding the contract method 0x60a627a1.
//
// Solidity: function bindMyReferral(address parent) returns()
func (_BindUser *BindUserTransactor) BindMyReferral(opts *bind.TransactOpts, parent common.Address) (*types.Transaction, error) {
	return _BindUser.contract.Transact(opts, "bindMyReferral", parent)
}

// BindMyReferral is a paid mutator transaction binding the contract method 0x60a627a1.
//
// Solidity: function bindMyReferral(address parent) returns()
func (_BindUser *BindUserSession) BindMyReferral(parent common.Address) (*types.Transaction, error) {
	return _BindUser.Contract.BindMyReferral(&_BindUser.TransactOpts, parent)
}

// BindMyReferral is a paid mutator transaction binding the contract method 0x60a627a1.
//
// Solidity: function bindMyReferral(address parent) returns()
func (_BindUser *BindUserTransactorSession) BindMyReferral(parent common.Address) (*types.Transaction, error) {
	return _BindUser.Contract.BindMyReferral(&_BindUser.TransactOpts, parent)
}

// BindReferral is a paid mutator transaction binding the contract method 0x7a9a433c.
//
// Solidity: function bindReferral(address parent, address user) returns()
func (_BindUser *BindUserTransactor) BindReferral(opts *bind.TransactOpts, parent common.Address, user common.Address) (*types.Transaction, error) {
	return _BindUser.contract.Transact(opts, "bindReferral", parent, user)
}

// BindReferral is a paid mutator transaction binding the contract method 0x7a9a433c.
//
// Solidity: function bindReferral(address parent, address user) returns()
func (_BindUser *BindUserSession) BindReferral(parent common.Address, user common.Address) (*types.Transaction, error) {
	return _BindUser.Contract.BindReferral(&_BindUser.TransactOpts, parent, user)
}

// BindReferral is a paid mutator transaction binding the contract method 0x7a9a433c.
//
// Solidity: function bindReferral(address parent, address user) returns()
func (_BindUser *BindUserTransactorSession) BindReferral(parent common.Address, user common.Address) (*types.Transaction, error) {
	return _BindUser.Contract.BindReferral(&_BindUser.TransactOpts, parent, user)
}

// InitBinds is a paid mutator transaction binding the contract method 0xd884ba08.
//
// Solidity: function initBinds(address[] parents, address[] users) returns()
func (_BindUser *BindUserTransactor) InitBinds(opts *bind.TransactOpts, parents []common.Address, users []common.Address) (*types.Transaction, error) {
	return _BindUser.contract.Transact(opts, "initBinds", parents, users)
}

// InitBinds is a paid mutator transaction binding the contract method 0xd884ba08.
//
// Solidity: function initBinds(address[] parents, address[] users) returns()
func (_BindUser *BindUserSession) InitBinds(parents []common.Address, users []common.Address) (*types.Transaction, error) {
	return _BindUser.Contract.InitBinds(&_BindUser.TransactOpts, parents, users)
}

// InitBinds is a paid mutator transaction binding the contract method 0xd884ba08.
//
// Solidity: function initBinds(address[] parents, address[] users) returns()
func (_BindUser *BindUserTransactorSession) InitBinds(parents []common.Address, users []common.Address) (*types.Transaction, error) {
	return _BindUser.Contract.InitBinds(&_BindUser.TransactOpts, parents, users)
}

// InitSystems is a paid mutator transaction binding the contract method 0x154f56a5.
//
// Solidity: function initSystems(address stakeSystem_, address dividendSystem_) returns()
func (_BindUser *BindUserTransactor) InitSystems(opts *bind.TransactOpts, stakeSystem_ common.Address, dividendSystem_ common.Address) (*types.Transaction, error) {
	return _BindUser.contract.Transact(opts, "initSystems", stakeSystem_, dividendSystem_)
}

// InitSystems is a paid mutator transaction binding the contract method 0x154f56a5.
//
// Solidity: function initSystems(address stakeSystem_, address dividendSystem_) returns()
func (_BindUser *BindUserSession) InitSystems(stakeSystem_ common.Address, dividendSystem_ common.Address) (*types.Transaction, error) {
	return _BindUser.Contract.InitSystems(&_BindUser.TransactOpts, stakeSystem_, dividendSystem_)
}

// InitSystems is a paid mutator transaction binding the contract method 0x154f56a5.
//
// Solidity: function initSystems(address stakeSystem_, address dividendSystem_) returns()
func (_BindUser *BindUserTransactorSession) InitSystems(stakeSystem_ common.Address, dividendSystem_ common.Address) (*types.Transaction, error) {
	return _BindUser.Contract.InitSystems(&_BindUser.TransactOpts, stakeSystem_, dividendSystem_)
}

// OnStake is a paid mutator transaction binding the contract method 0x8c087b1c.
//
// Solidity: function onStake(address user, uint256 amount) returns()
func (_BindUser *BindUserTransactor) OnStake(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BindUser.contract.Transact(opts, "onStake", user, amount)
}

// OnStake is a paid mutator transaction binding the contract method 0x8c087b1c.
//
// Solidity: function onStake(address user, uint256 amount) returns()
func (_BindUser *BindUserSession) OnStake(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BindUser.Contract.OnStake(&_BindUser.TransactOpts, user, amount)
}

// OnStake is a paid mutator transaction binding the contract method 0x8c087b1c.
//
// Solidity: function onStake(address user, uint256 amount) returns()
func (_BindUser *BindUserTransactorSession) OnStake(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BindUser.Contract.OnStake(&_BindUser.TransactOpts, user, amount)
}

// OnUnstake is a paid mutator transaction binding the contract method 0x7fb6d3a3.
//
// Solidity: function onUnstake(address user, uint256 amount) returns()
func (_BindUser *BindUserTransactor) OnUnstake(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BindUser.contract.Transact(opts, "onUnstake", user, amount)
}

// OnUnstake is a paid mutator transaction binding the contract method 0x7fb6d3a3.
//
// Solidity: function onUnstake(address user, uint256 amount) returns()
func (_BindUser *BindUserSession) OnUnstake(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BindUser.Contract.OnUnstake(&_BindUser.TransactOpts, user, amount)
}

// OnUnstake is a paid mutator transaction binding the contract method 0x7fb6d3a3.
//
// Solidity: function onUnstake(address user, uint256 amount) returns()
func (_BindUser *BindUserTransactorSession) OnUnstake(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BindUser.Contract.OnUnstake(&_BindUser.TransactOpts, user, amount)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address parent) returns()
func (_BindUser *BindUserTransactor) Register(opts *bind.TransactOpts, parent common.Address) (*types.Transaction, error) {
	return _BindUser.contract.Transact(opts, "register", parent)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address parent) returns()
func (_BindUser *BindUserSession) Register(parent common.Address) (*types.Transaction, error) {
	return _BindUser.Contract.Register(&_BindUser.TransactOpts, parent)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address parent) returns()
func (_BindUser *BindUserTransactorSession) Register(parent common.Address) (*types.Transaction, error) {
	return _BindUser.Contract.Register(&_BindUser.TransactOpts, parent)
}

// SetExtraPerformance is a paid mutator transaction binding the contract method 0xc59ea26a.
//
// Solidity: function setExtraPerformance(address account, uint256 extraTeamPerformance, uint256 extraPersonalStake) returns()
func (_BindUser *BindUserTransactor) SetExtraPerformance(opts *bind.TransactOpts, account common.Address, extraTeamPerformance *big.Int, extraPersonalStake *big.Int) (*types.Transaction, error) {
	return _BindUser.contract.Transact(opts, "setExtraPerformance", account, extraTeamPerformance, extraPersonalStake)
}

// SetExtraPerformance is a paid mutator transaction binding the contract method 0xc59ea26a.
//
// Solidity: function setExtraPerformance(address account, uint256 extraTeamPerformance, uint256 extraPersonalStake) returns()
func (_BindUser *BindUserSession) SetExtraPerformance(account common.Address, extraTeamPerformance *big.Int, extraPersonalStake *big.Int) (*types.Transaction, error) {
	return _BindUser.Contract.SetExtraPerformance(&_BindUser.TransactOpts, account, extraTeamPerformance, extraPersonalStake)
}

// SetExtraPerformance is a paid mutator transaction binding the contract method 0xc59ea26a.
//
// Solidity: function setExtraPerformance(address account, uint256 extraTeamPerformance, uint256 extraPersonalStake) returns()
func (_BindUser *BindUserTransactorSession) SetExtraPerformance(account common.Address, extraTeamPerformance *big.Int, extraPersonalStake *big.Int) (*types.Transaction, error) {
	return _BindUser.Contract.SetExtraPerformance(&_BindUser.TransactOpts, account, extraTeamPerformance, extraPersonalStake)
}

// SetExtraPerformanceFromDividend is a paid mutator transaction binding the contract method 0x3a532fd0.
//
// Solidity: function setExtraPerformanceFromDividend(address account, uint256 extraTeamPerformance, uint256 extraPersonalStake) returns()
func (_BindUser *BindUserTransactor) SetExtraPerformanceFromDividend(opts *bind.TransactOpts, account common.Address, extraTeamPerformance *big.Int, extraPersonalStake *big.Int) (*types.Transaction, error) {
	return _BindUser.contract.Transact(opts, "setExtraPerformanceFromDividend", account, extraTeamPerformance, extraPersonalStake)
}

// SetExtraPerformanceFromDividend is a paid mutator transaction binding the contract method 0x3a532fd0.
//
// Solidity: function setExtraPerformanceFromDividend(address account, uint256 extraTeamPerformance, uint256 extraPersonalStake) returns()
func (_BindUser *BindUserSession) SetExtraPerformanceFromDividend(account common.Address, extraTeamPerformance *big.Int, extraPersonalStake *big.Int) (*types.Transaction, error) {
	return _BindUser.Contract.SetExtraPerformanceFromDividend(&_BindUser.TransactOpts, account, extraTeamPerformance, extraPersonalStake)
}

// SetExtraPerformanceFromDividend is a paid mutator transaction binding the contract method 0x3a532fd0.
//
// Solidity: function setExtraPerformanceFromDividend(address account, uint256 extraTeamPerformance, uint256 extraPersonalStake) returns()
func (_BindUser *BindUserTransactorSession) SetExtraPerformanceFromDividend(account common.Address, extraTeamPerformance *big.Int, extraPersonalStake *big.Int) (*types.Transaction, error) {
	return _BindUser.Contract.SetExtraPerformanceFromDividend(&_BindUser.TransactOpts, account, extraTeamPerformance, extraPersonalStake)
}

// SetInitAccount is a paid mutator transaction binding the contract method 0x42b34815.
//
// Solidity: function setInitAccount(address newInitAccount) returns()
func (_BindUser *BindUserTransactor) SetInitAccount(opts *bind.TransactOpts, newInitAccount common.Address) (*types.Transaction, error) {
	return _BindUser.contract.Transact(opts, "setInitAccount", newInitAccount)
}

// SetInitAccount is a paid mutator transaction binding the contract method 0x42b34815.
//
// Solidity: function setInitAccount(address newInitAccount) returns()
func (_BindUser *BindUserSession) SetInitAccount(newInitAccount common.Address) (*types.Transaction, error) {
	return _BindUser.Contract.SetInitAccount(&_BindUser.TransactOpts, newInitAccount)
}

// SetInitAccount is a paid mutator transaction binding the contract method 0x42b34815.
//
// Solidity: function setInitAccount(address newInitAccount) returns()
func (_BindUser *BindUserTransactorSession) SetInitAccount(newInitAccount common.Address) (*types.Transaction, error) {
	return _BindUser.Contract.SetInitAccount(&_BindUser.TransactOpts, newInitAccount)
}

// SetMaxUplineDepth is a paid mutator transaction binding the contract method 0x7518c44a.
//
// Solidity: function setMaxUplineDepth(uint16 depth) returns()
func (_BindUser *BindUserTransactor) SetMaxUplineDepth(opts *bind.TransactOpts, depth uint16) (*types.Transaction, error) {
	return _BindUser.contract.Transact(opts, "setMaxUplineDepth", depth)
}

// SetMaxUplineDepth is a paid mutator transaction binding the contract method 0x7518c44a.
//
// Solidity: function setMaxUplineDepth(uint16 depth) returns()
func (_BindUser *BindUserSession) SetMaxUplineDepth(depth uint16) (*types.Transaction, error) {
	return _BindUser.Contract.SetMaxUplineDepth(&_BindUser.TransactOpts, depth)
}

// SetMaxUplineDepth is a paid mutator transaction binding the contract method 0x7518c44a.
//
// Solidity: function setMaxUplineDepth(uint16 depth) returns()
func (_BindUser *BindUserTransactorSession) SetMaxUplineDepth(depth uint16) (*types.Transaction, error) {
	return _BindUser.Contract.SetMaxUplineDepth(&_BindUser.TransactOpts, depth)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BindUser *BindUserTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BindUser.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BindUser *BindUserSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BindUser.Contract.TransferOwnership(&_BindUser.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BindUser *BindUserTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BindUser.Contract.TransferOwnership(&_BindUser.TransactOpts, newOwner)
}

// BindUserBindReferralIterator is returned from FilterBindReferral and is used to iterate over the raw logs and unpacked data for BindReferral events raised by the BindUser contract.
type BindUserBindReferralIterator struct {
	Event *BindUserBindReferral // Event containing the contract specifics and raw log

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
func (it *BindUserBindReferralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindUserBindReferral)
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
		it.Event = new(BindUserBindReferral)
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
func (it *BindUserBindReferralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindUserBindReferralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindUserBindReferral represents a BindReferral event raised by the BindUser contract.
type BindUserBindReferral struct {
	User   common.Address
	Parent common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBindReferral is a free log retrieval operation binding the contract event 0xcb282389824b9ba09f76f40ee522be8104006b13f0a7a8d78c2198a1cd1da468.
//
// Solidity: event BindReferral(address indexed user, address indexed parent)
func (_BindUser *BindUserFilterer) FilterBindReferral(opts *bind.FilterOpts, user []common.Address, parent []common.Address) (*BindUserBindReferralIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var parentRule []interface{}
	for _, parentItem := range parent {
		parentRule = append(parentRule, parentItem)
	}

	logs, sub, err := _BindUser.contract.FilterLogs(opts, "BindReferral", userRule, parentRule)
	if err != nil {
		return nil, err
	}
	return &BindUserBindReferralIterator{contract: _BindUser.contract, event: "BindReferral", logs: logs, sub: sub}, nil
}

// WatchBindReferral is a free log subscription operation binding the contract event 0xcb282389824b9ba09f76f40ee522be8104006b13f0a7a8d78c2198a1cd1da468.
//
// Solidity: event BindReferral(address indexed user, address indexed parent)
func (_BindUser *BindUserFilterer) WatchBindReferral(opts *bind.WatchOpts, sink chan<- *BindUserBindReferral, user []common.Address, parent []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var parentRule []interface{}
	for _, parentItem := range parent {
		parentRule = append(parentRule, parentItem)
	}

	logs, sub, err := _BindUser.contract.WatchLogs(opts, "BindReferral", userRule, parentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindUserBindReferral)
				if err := _BindUser.contract.UnpackLog(event, "BindReferral", log); err != nil {
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

// ParseBindReferral is a log parse operation binding the contract event 0xcb282389824b9ba09f76f40ee522be8104006b13f0a7a8d78c2198a1cd1da468.
//
// Solidity: event BindReferral(address indexed user, address indexed parent)
func (_BindUser *BindUserFilterer) ParseBindReferral(log types.Log) (*BindUserBindReferral, error) {
	event := new(BindUserBindReferral)
	if err := _BindUser.contract.UnpackLog(event, "BindReferral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindUserDividendSystemUpdatedIterator is returned from FilterDividendSystemUpdated and is used to iterate over the raw logs and unpacked data for DividendSystemUpdated events raised by the BindUser contract.
type BindUserDividendSystemUpdatedIterator struct {
	Event *BindUserDividendSystemUpdated // Event containing the contract specifics and raw log

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
func (it *BindUserDividendSystemUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindUserDividendSystemUpdated)
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
		it.Event = new(BindUserDividendSystemUpdated)
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
func (it *BindUserDividendSystemUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindUserDividendSystemUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindUserDividendSystemUpdated represents a DividendSystemUpdated event raised by the BindUser contract.
type BindUserDividendSystemUpdated struct {
	DividendSystem common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDividendSystemUpdated is a free log retrieval operation binding the contract event 0xb7fa8611c311f61b6eb236ee9ebe0e2c6c4d175233290fa1ad63dd6e265469f0.
//
// Solidity: event DividendSystemUpdated(address indexed dividendSystem)
func (_BindUser *BindUserFilterer) FilterDividendSystemUpdated(opts *bind.FilterOpts, dividendSystem []common.Address) (*BindUserDividendSystemUpdatedIterator, error) {

	var dividendSystemRule []interface{}
	for _, dividendSystemItem := range dividendSystem {
		dividendSystemRule = append(dividendSystemRule, dividendSystemItem)
	}

	logs, sub, err := _BindUser.contract.FilterLogs(opts, "DividendSystemUpdated", dividendSystemRule)
	if err != nil {
		return nil, err
	}
	return &BindUserDividendSystemUpdatedIterator{contract: _BindUser.contract, event: "DividendSystemUpdated", logs: logs, sub: sub}, nil
}

// WatchDividendSystemUpdated is a free log subscription operation binding the contract event 0xb7fa8611c311f61b6eb236ee9ebe0e2c6c4d175233290fa1ad63dd6e265469f0.
//
// Solidity: event DividendSystemUpdated(address indexed dividendSystem)
func (_BindUser *BindUserFilterer) WatchDividendSystemUpdated(opts *bind.WatchOpts, sink chan<- *BindUserDividendSystemUpdated, dividendSystem []common.Address) (event.Subscription, error) {

	var dividendSystemRule []interface{}
	for _, dividendSystemItem := range dividendSystem {
		dividendSystemRule = append(dividendSystemRule, dividendSystemItem)
	}

	logs, sub, err := _BindUser.contract.WatchLogs(opts, "DividendSystemUpdated", dividendSystemRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindUserDividendSystemUpdated)
				if err := _BindUser.contract.UnpackLog(event, "DividendSystemUpdated", log); err != nil {
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
func (_BindUser *BindUserFilterer) ParseDividendSystemUpdated(log types.Log) (*BindUserDividendSystemUpdated, error) {
	event := new(BindUserDividendSystemUpdated)
	if err := _BindUser.contract.UnpackLog(event, "DividendSystemUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindUserExtraPerformanceUpdatedIterator is returned from FilterExtraPerformanceUpdated and is used to iterate over the raw logs and unpacked data for ExtraPerformanceUpdated events raised by the BindUser contract.
type BindUserExtraPerformanceUpdatedIterator struct {
	Event *BindUserExtraPerformanceUpdated // Event containing the contract specifics and raw log

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
func (it *BindUserExtraPerformanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindUserExtraPerformanceUpdated)
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
		it.Event = new(BindUserExtraPerformanceUpdated)
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
func (it *BindUserExtraPerformanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindUserExtraPerformanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindUserExtraPerformanceUpdated represents a ExtraPerformanceUpdated event raised by the BindUser contract.
type BindUserExtraPerformanceUpdated struct {
	Account              common.Address
	ExtraTeamPerformance *big.Int
	ExtraStakeAmount     *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterExtraPerformanceUpdated is a free log retrieval operation binding the contract event 0x7204e217511b6b6dfc3aa5d5c286642d2875af307c400dc28398b060b3f07dc9.
//
// Solidity: event ExtraPerformanceUpdated(address indexed account, uint256 extraTeamPerformance, uint256 extraStakeAmount)
func (_BindUser *BindUserFilterer) FilterExtraPerformanceUpdated(opts *bind.FilterOpts, account []common.Address) (*BindUserExtraPerformanceUpdatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BindUser.contract.FilterLogs(opts, "ExtraPerformanceUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return &BindUserExtraPerformanceUpdatedIterator{contract: _BindUser.contract, event: "ExtraPerformanceUpdated", logs: logs, sub: sub}, nil
}

// WatchExtraPerformanceUpdated is a free log subscription operation binding the contract event 0x7204e217511b6b6dfc3aa5d5c286642d2875af307c400dc28398b060b3f07dc9.
//
// Solidity: event ExtraPerformanceUpdated(address indexed account, uint256 extraTeamPerformance, uint256 extraStakeAmount)
func (_BindUser *BindUserFilterer) WatchExtraPerformanceUpdated(opts *bind.WatchOpts, sink chan<- *BindUserExtraPerformanceUpdated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BindUser.contract.WatchLogs(opts, "ExtraPerformanceUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindUserExtraPerformanceUpdated)
				if err := _BindUser.contract.UnpackLog(event, "ExtraPerformanceUpdated", log); err != nil {
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

// ParseExtraPerformanceUpdated is a log parse operation binding the contract event 0x7204e217511b6b6dfc3aa5d5c286642d2875af307c400dc28398b060b3f07dc9.
//
// Solidity: event ExtraPerformanceUpdated(address indexed account, uint256 extraTeamPerformance, uint256 extraStakeAmount)
func (_BindUser *BindUserFilterer) ParseExtraPerformanceUpdated(log types.Log) (*BindUserExtraPerformanceUpdated, error) {
	event := new(BindUserExtraPerformanceUpdated)
	if err := _BindUser.contract.UnpackLog(event, "ExtraPerformanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindUserInitAccountTransferredIterator is returned from FilterInitAccountTransferred and is used to iterate over the raw logs and unpacked data for InitAccountTransferred events raised by the BindUser contract.
type BindUserInitAccountTransferredIterator struct {
	Event *BindUserInitAccountTransferred // Event containing the contract specifics and raw log

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
func (it *BindUserInitAccountTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindUserInitAccountTransferred)
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
		it.Event = new(BindUserInitAccountTransferred)
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
func (it *BindUserInitAccountTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindUserInitAccountTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindUserInitAccountTransferred represents a InitAccountTransferred event raised by the BindUser contract.
type BindUserInitAccountTransferred struct {
	PreviousInitAccount common.Address
	NewInitAccount      common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterInitAccountTransferred is a free log retrieval operation binding the contract event 0xdbdef5214b96ccc548368e20ee0423776be2da50a3db91a024c5b43802150380.
//
// Solidity: event InitAccountTransferred(address indexed previousInitAccount, address indexed newInitAccount)
func (_BindUser *BindUserFilterer) FilterInitAccountTransferred(opts *bind.FilterOpts, previousInitAccount []common.Address, newInitAccount []common.Address) (*BindUserInitAccountTransferredIterator, error) {

	var previousInitAccountRule []interface{}
	for _, previousInitAccountItem := range previousInitAccount {
		previousInitAccountRule = append(previousInitAccountRule, previousInitAccountItem)
	}
	var newInitAccountRule []interface{}
	for _, newInitAccountItem := range newInitAccount {
		newInitAccountRule = append(newInitAccountRule, newInitAccountItem)
	}

	logs, sub, err := _BindUser.contract.FilterLogs(opts, "InitAccountTransferred", previousInitAccountRule, newInitAccountRule)
	if err != nil {
		return nil, err
	}
	return &BindUserInitAccountTransferredIterator{contract: _BindUser.contract, event: "InitAccountTransferred", logs: logs, sub: sub}, nil
}

// WatchInitAccountTransferred is a free log subscription operation binding the contract event 0xdbdef5214b96ccc548368e20ee0423776be2da50a3db91a024c5b43802150380.
//
// Solidity: event InitAccountTransferred(address indexed previousInitAccount, address indexed newInitAccount)
func (_BindUser *BindUserFilterer) WatchInitAccountTransferred(opts *bind.WatchOpts, sink chan<- *BindUserInitAccountTransferred, previousInitAccount []common.Address, newInitAccount []common.Address) (event.Subscription, error) {

	var previousInitAccountRule []interface{}
	for _, previousInitAccountItem := range previousInitAccount {
		previousInitAccountRule = append(previousInitAccountRule, previousInitAccountItem)
	}
	var newInitAccountRule []interface{}
	for _, newInitAccountItem := range newInitAccount {
		newInitAccountRule = append(newInitAccountRule, newInitAccountItem)
	}

	logs, sub, err := _BindUser.contract.WatchLogs(opts, "InitAccountTransferred", previousInitAccountRule, newInitAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindUserInitAccountTransferred)
				if err := _BindUser.contract.UnpackLog(event, "InitAccountTransferred", log); err != nil {
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
func (_BindUser *BindUserFilterer) ParseInitAccountTransferred(log types.Log) (*BindUserInitAccountTransferred, error) {
	event := new(BindUserInitAccountTransferred)
	if err := _BindUser.contract.UnpackLog(event, "InitAccountTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindUserInitSystemsUpdatedIterator is returned from FilterInitSystemsUpdated and is used to iterate over the raw logs and unpacked data for InitSystemsUpdated events raised by the BindUser contract.
type BindUserInitSystemsUpdatedIterator struct {
	Event *BindUserInitSystemsUpdated // Event containing the contract specifics and raw log

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
func (it *BindUserInitSystemsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindUserInitSystemsUpdated)
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
		it.Event = new(BindUserInitSystemsUpdated)
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
func (it *BindUserInitSystemsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindUserInitSystemsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindUserInitSystemsUpdated represents a InitSystemsUpdated event raised by the BindUser contract.
type BindUserInitSystemsUpdated struct {
	StakeSystem    common.Address
	DividendSystem common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitSystemsUpdated is a free log retrieval operation binding the contract event 0xbf44382949ee7563a72894d0bc28ab735a831324baa0ca9989b05bb81c147fd2.
//
// Solidity: event InitSystemsUpdated(address indexed stakeSystem, address indexed dividendSystem)
func (_BindUser *BindUserFilterer) FilterInitSystemsUpdated(opts *bind.FilterOpts, stakeSystem []common.Address, dividendSystem []common.Address) (*BindUserInitSystemsUpdatedIterator, error) {

	var stakeSystemRule []interface{}
	for _, stakeSystemItem := range stakeSystem {
		stakeSystemRule = append(stakeSystemRule, stakeSystemItem)
	}
	var dividendSystemRule []interface{}
	for _, dividendSystemItem := range dividendSystem {
		dividendSystemRule = append(dividendSystemRule, dividendSystemItem)
	}

	logs, sub, err := _BindUser.contract.FilterLogs(opts, "InitSystemsUpdated", stakeSystemRule, dividendSystemRule)
	if err != nil {
		return nil, err
	}
	return &BindUserInitSystemsUpdatedIterator{contract: _BindUser.contract, event: "InitSystemsUpdated", logs: logs, sub: sub}, nil
}

// WatchInitSystemsUpdated is a free log subscription operation binding the contract event 0xbf44382949ee7563a72894d0bc28ab735a831324baa0ca9989b05bb81c147fd2.
//
// Solidity: event InitSystemsUpdated(address indexed stakeSystem, address indexed dividendSystem)
func (_BindUser *BindUserFilterer) WatchInitSystemsUpdated(opts *bind.WatchOpts, sink chan<- *BindUserInitSystemsUpdated, stakeSystem []common.Address, dividendSystem []common.Address) (event.Subscription, error) {

	var stakeSystemRule []interface{}
	for _, stakeSystemItem := range stakeSystem {
		stakeSystemRule = append(stakeSystemRule, stakeSystemItem)
	}
	var dividendSystemRule []interface{}
	for _, dividendSystemItem := range dividendSystem {
		dividendSystemRule = append(dividendSystemRule, dividendSystemItem)
	}

	logs, sub, err := _BindUser.contract.WatchLogs(opts, "InitSystemsUpdated", stakeSystemRule, dividendSystemRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindUserInitSystemsUpdated)
				if err := _BindUser.contract.UnpackLog(event, "InitSystemsUpdated", log); err != nil {
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

// ParseInitSystemsUpdated is a log parse operation binding the contract event 0xbf44382949ee7563a72894d0bc28ab735a831324baa0ca9989b05bb81c147fd2.
//
// Solidity: event InitSystemsUpdated(address indexed stakeSystem, address indexed dividendSystem)
func (_BindUser *BindUserFilterer) ParseInitSystemsUpdated(log types.Log) (*BindUserInitSystemsUpdated, error) {
	event := new(BindUserInitSystemsUpdated)
	if err := _BindUser.contract.UnpackLog(event, "InitSystemsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindUserLevel7AddedIterator is returned from FilterLevel7Added and is used to iterate over the raw logs and unpacked data for Level7Added events raised by the BindUser contract.
type BindUserLevel7AddedIterator struct {
	Event *BindUserLevel7Added // Event containing the contract specifics and raw log

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
func (it *BindUserLevel7AddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindUserLevel7Added)
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
		it.Event = new(BindUserLevel7Added)
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
func (it *BindUserLevel7AddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindUserLevel7AddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindUserLevel7Added represents a Level7Added event raised by the BindUser contract.
type BindUserLevel7Added struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLevel7Added is a free log retrieval operation binding the contract event 0x2658a9ef3bb3a975dc49d647a2e70feac4377824a960e7b304e94cff5cb7b23d.
//
// Solidity: event Level7Added(address indexed account)
func (_BindUser *BindUserFilterer) FilterLevel7Added(opts *bind.FilterOpts, account []common.Address) (*BindUserLevel7AddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BindUser.contract.FilterLogs(opts, "Level7Added", accountRule)
	if err != nil {
		return nil, err
	}
	return &BindUserLevel7AddedIterator{contract: _BindUser.contract, event: "Level7Added", logs: logs, sub: sub}, nil
}

// WatchLevel7Added is a free log subscription operation binding the contract event 0x2658a9ef3bb3a975dc49d647a2e70feac4377824a960e7b304e94cff5cb7b23d.
//
// Solidity: event Level7Added(address indexed account)
func (_BindUser *BindUserFilterer) WatchLevel7Added(opts *bind.WatchOpts, sink chan<- *BindUserLevel7Added, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BindUser.contract.WatchLogs(opts, "Level7Added", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindUserLevel7Added)
				if err := _BindUser.contract.UnpackLog(event, "Level7Added", log); err != nil {
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

// ParseLevel7Added is a log parse operation binding the contract event 0x2658a9ef3bb3a975dc49d647a2e70feac4377824a960e7b304e94cff5cb7b23d.
//
// Solidity: event Level7Added(address indexed account)
func (_BindUser *BindUserFilterer) ParseLevel7Added(log types.Log) (*BindUserLevel7Added, error) {
	event := new(BindUserLevel7Added)
	if err := _BindUser.contract.UnpackLog(event, "Level7Added", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindUserLevel7RemovedIterator is returned from FilterLevel7Removed and is used to iterate over the raw logs and unpacked data for Level7Removed events raised by the BindUser contract.
type BindUserLevel7RemovedIterator struct {
	Event *BindUserLevel7Removed // Event containing the contract specifics and raw log

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
func (it *BindUserLevel7RemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindUserLevel7Removed)
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
		it.Event = new(BindUserLevel7Removed)
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
func (it *BindUserLevel7RemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindUserLevel7RemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindUserLevel7Removed represents a Level7Removed event raised by the BindUser contract.
type BindUserLevel7Removed struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLevel7Removed is a free log retrieval operation binding the contract event 0xfb59e24951e53bf9e88899a624110852317e2abed639f8319c1564c110fe6de2.
//
// Solidity: event Level7Removed(address indexed account)
func (_BindUser *BindUserFilterer) FilterLevel7Removed(opts *bind.FilterOpts, account []common.Address) (*BindUserLevel7RemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BindUser.contract.FilterLogs(opts, "Level7Removed", accountRule)
	if err != nil {
		return nil, err
	}
	return &BindUserLevel7RemovedIterator{contract: _BindUser.contract, event: "Level7Removed", logs: logs, sub: sub}, nil
}

// WatchLevel7Removed is a free log subscription operation binding the contract event 0xfb59e24951e53bf9e88899a624110852317e2abed639f8319c1564c110fe6de2.
//
// Solidity: event Level7Removed(address indexed account)
func (_BindUser *BindUserFilterer) WatchLevel7Removed(opts *bind.WatchOpts, sink chan<- *BindUserLevel7Removed, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BindUser.contract.WatchLogs(opts, "Level7Removed", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindUserLevel7Removed)
				if err := _BindUser.contract.UnpackLog(event, "Level7Removed", log); err != nil {
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

// ParseLevel7Removed is a log parse operation binding the contract event 0xfb59e24951e53bf9e88899a624110852317e2abed639f8319c1564c110fe6de2.
//
// Solidity: event Level7Removed(address indexed account)
func (_BindUser *BindUserFilterer) ParseLevel7Removed(log types.Log) (*BindUserLevel7Removed, error) {
	event := new(BindUserLevel7Removed)
	if err := _BindUser.contract.UnpackLog(event, "Level7Removed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindUserLevelRateUpdatedIterator is returned from FilterLevelRateUpdated and is used to iterate over the raw logs and unpacked data for LevelRateUpdated events raised by the BindUser contract.
type BindUserLevelRateUpdatedIterator struct {
	Event *BindUserLevelRateUpdated // Event containing the contract specifics and raw log

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
func (it *BindUserLevelRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindUserLevelRateUpdated)
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
		it.Event = new(BindUserLevelRateUpdated)
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
func (it *BindUserLevelRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindUserLevelRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindUserLevelRateUpdated represents a LevelRateUpdated event raised by the BindUser contract.
type BindUserLevelRateUpdated struct {
	Index   uint8
	RateBps uint16
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLevelRateUpdated is a free log retrieval operation binding the contract event 0xdf2f057c2f374e9dc10a16cc469485d0825a9b1bc9392ead32679be6ded75a9c.
//
// Solidity: event LevelRateUpdated(uint8 indexed index, uint16 rateBps)
func (_BindUser *BindUserFilterer) FilterLevelRateUpdated(opts *bind.FilterOpts, index []uint8) (*BindUserLevelRateUpdatedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BindUser.contract.FilterLogs(opts, "LevelRateUpdated", indexRule)
	if err != nil {
		return nil, err
	}
	return &BindUserLevelRateUpdatedIterator{contract: _BindUser.contract, event: "LevelRateUpdated", logs: logs, sub: sub}, nil
}

// WatchLevelRateUpdated is a free log subscription operation binding the contract event 0xdf2f057c2f374e9dc10a16cc469485d0825a9b1bc9392ead32679be6ded75a9c.
//
// Solidity: event LevelRateUpdated(uint8 indexed index, uint16 rateBps)
func (_BindUser *BindUserFilterer) WatchLevelRateUpdated(opts *bind.WatchOpts, sink chan<- *BindUserLevelRateUpdated, index []uint8) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _BindUser.contract.WatchLogs(opts, "LevelRateUpdated", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindUserLevelRateUpdated)
				if err := _BindUser.contract.UnpackLog(event, "LevelRateUpdated", log); err != nil {
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

// ParseLevelRateUpdated is a log parse operation binding the contract event 0xdf2f057c2f374e9dc10a16cc469485d0825a9b1bc9392ead32679be6ded75a9c.
//
// Solidity: event LevelRateUpdated(uint8 indexed index, uint16 rateBps)
func (_BindUser *BindUserFilterer) ParseLevelRateUpdated(log types.Log) (*BindUserLevelRateUpdated, error) {
	event := new(BindUserLevelRateUpdated)
	if err := _BindUser.contract.UnpackLog(event, "LevelRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindUserOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BindUser contract.
type BindUserOwnershipTransferredIterator struct {
	Event *BindUserOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BindUserOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindUserOwnershipTransferred)
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
		it.Event = new(BindUserOwnershipTransferred)
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
func (it *BindUserOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindUserOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindUserOwnershipTransferred represents a OwnershipTransferred event raised by the BindUser contract.
type BindUserOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BindUser *BindUserFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BindUserOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BindUser.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BindUserOwnershipTransferredIterator{contract: _BindUser.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BindUser *BindUserFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BindUserOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BindUser.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindUserOwnershipTransferred)
				if err := _BindUser.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BindUser *BindUserFilterer) ParseOwnershipTransferred(log types.Log) (*BindUserOwnershipTransferred, error) {
	event := new(BindUserOwnershipTransferred)
	if err := _BindUser.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindUserRootUpdatedIterator is returned from FilterRootUpdated and is used to iterate over the raw logs and unpacked data for RootUpdated events raised by the BindUser contract.
type BindUserRootUpdatedIterator struct {
	Event *BindUserRootUpdated // Event containing the contract specifics and raw log

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
func (it *BindUserRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindUserRootUpdated)
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
		it.Event = new(BindUserRootUpdated)
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
func (it *BindUserRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindUserRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindUserRootUpdated represents a RootUpdated event raised by the BindUser contract.
type BindUserRootUpdated struct {
	RootAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRootUpdated is a free log retrieval operation binding the contract event 0x2e3d44d6f64f37f940e282810e892fed0e8c19a49df0e6dfd25ac4c42f7507f9.
//
// Solidity: event RootUpdated(address indexed rootAddress)
func (_BindUser *BindUserFilterer) FilterRootUpdated(opts *bind.FilterOpts, rootAddress []common.Address) (*BindUserRootUpdatedIterator, error) {

	var rootAddressRule []interface{}
	for _, rootAddressItem := range rootAddress {
		rootAddressRule = append(rootAddressRule, rootAddressItem)
	}

	logs, sub, err := _BindUser.contract.FilterLogs(opts, "RootUpdated", rootAddressRule)
	if err != nil {
		return nil, err
	}
	return &BindUserRootUpdatedIterator{contract: _BindUser.contract, event: "RootUpdated", logs: logs, sub: sub}, nil
}

// WatchRootUpdated is a free log subscription operation binding the contract event 0x2e3d44d6f64f37f940e282810e892fed0e8c19a49df0e6dfd25ac4c42f7507f9.
//
// Solidity: event RootUpdated(address indexed rootAddress)
func (_BindUser *BindUserFilterer) WatchRootUpdated(opts *bind.WatchOpts, sink chan<- *BindUserRootUpdated, rootAddress []common.Address) (event.Subscription, error) {

	var rootAddressRule []interface{}
	for _, rootAddressItem := range rootAddress {
		rootAddressRule = append(rootAddressRule, rootAddressItem)
	}

	logs, sub, err := _BindUser.contract.WatchLogs(opts, "RootUpdated", rootAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindUserRootUpdated)
				if err := _BindUser.contract.UnpackLog(event, "RootUpdated", log); err != nil {
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

// ParseRootUpdated is a log parse operation binding the contract event 0x2e3d44d6f64f37f940e282810e892fed0e8c19a49df0e6dfd25ac4c42f7507f9.
//
// Solidity: event RootUpdated(address indexed rootAddress)
func (_BindUser *BindUserFilterer) ParseRootUpdated(log types.Log) (*BindUserRootUpdated, error) {
	event := new(BindUserRootUpdated)
	if err := _BindUser.contract.UnpackLog(event, "RootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindUserStakeSystemUpdatedIterator is returned from FilterStakeSystemUpdated and is used to iterate over the raw logs and unpacked data for StakeSystemUpdated events raised by the BindUser contract.
type BindUserStakeSystemUpdatedIterator struct {
	Event *BindUserStakeSystemUpdated // Event containing the contract specifics and raw log

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
func (it *BindUserStakeSystemUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindUserStakeSystemUpdated)
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
		it.Event = new(BindUserStakeSystemUpdated)
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
func (it *BindUserStakeSystemUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindUserStakeSystemUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindUserStakeSystemUpdated represents a StakeSystemUpdated event raised by the BindUser contract.
type BindUserStakeSystemUpdated struct {
	StakeSystem common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakeSystemUpdated is a free log retrieval operation binding the contract event 0x67ca87fce21096af1c404473ac6493c255ca444be43696a508d198428b0b12a6.
//
// Solidity: event StakeSystemUpdated(address indexed stakeSystem)
func (_BindUser *BindUserFilterer) FilterStakeSystemUpdated(opts *bind.FilterOpts, stakeSystem []common.Address) (*BindUserStakeSystemUpdatedIterator, error) {

	var stakeSystemRule []interface{}
	for _, stakeSystemItem := range stakeSystem {
		stakeSystemRule = append(stakeSystemRule, stakeSystemItem)
	}

	logs, sub, err := _BindUser.contract.FilterLogs(opts, "StakeSystemUpdated", stakeSystemRule)
	if err != nil {
		return nil, err
	}
	return &BindUserStakeSystemUpdatedIterator{contract: _BindUser.contract, event: "StakeSystemUpdated", logs: logs, sub: sub}, nil
}

// WatchStakeSystemUpdated is a free log subscription operation binding the contract event 0x67ca87fce21096af1c404473ac6493c255ca444be43696a508d198428b0b12a6.
//
// Solidity: event StakeSystemUpdated(address indexed stakeSystem)
func (_BindUser *BindUserFilterer) WatchStakeSystemUpdated(opts *bind.WatchOpts, sink chan<- *BindUserStakeSystemUpdated, stakeSystem []common.Address) (event.Subscription, error) {

	var stakeSystemRule []interface{}
	for _, stakeSystemItem := range stakeSystem {
		stakeSystemRule = append(stakeSystemRule, stakeSystemItem)
	}

	logs, sub, err := _BindUser.contract.WatchLogs(opts, "StakeSystemUpdated", stakeSystemRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindUserStakeSystemUpdated)
				if err := _BindUser.contract.UnpackLog(event, "StakeSystemUpdated", log); err != nil {
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

// ParseStakeSystemUpdated is a log parse operation binding the contract event 0x67ca87fce21096af1c404473ac6493c255ca444be43696a508d198428b0b12a6.
//
// Solidity: event StakeSystemUpdated(address indexed stakeSystem)
func (_BindUser *BindUserFilterer) ParseStakeSystemUpdated(log types.Log) (*BindUserStakeSystemUpdated, error) {
	event := new(BindUserStakeSystemUpdated)
	if err := _BindUser.contract.UnpackLog(event, "StakeSystemUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
