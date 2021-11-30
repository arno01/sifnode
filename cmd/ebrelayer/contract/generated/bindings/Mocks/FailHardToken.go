// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Mocks

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

// MocksMetaData contains all meta data concerning the Mocks contract.
var MocksMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountToMint\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cosmosDenom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"setDenom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MocksABI is the input ABI used to generate the binding from.
// Deprecated: Use MocksMetaData.ABI instead.
var MocksABI = MocksMetaData.ABI

// Mocks is an auto generated Go binding around an Ethereum contract.
type Mocks struct {
	MocksCaller     // Read-only binding to the contract
	MocksTransactor // Write-only binding to the contract
	MocksFilterer   // Log filterer for contract events
}

// MocksCaller is an auto generated read-only Go binding around an Ethereum contract.
type MocksCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MocksTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MocksTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MocksFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MocksFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MocksSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MocksSession struct {
	Contract     *Mocks            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MocksCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MocksCallerSession struct {
	Contract *MocksCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MocksTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MocksTransactorSession struct {
	Contract     *MocksTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MocksRaw is an auto generated low-level Go binding around an Ethereum contract.
type MocksRaw struct {
	Contract *Mocks // Generic contract binding to access the raw methods on
}

// MocksCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MocksCallerRaw struct {
	Contract *MocksCaller // Generic read-only contract binding to access the raw methods on
}

// MocksTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MocksTransactorRaw struct {
	Contract *MocksTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMocks creates a new instance of Mocks, bound to a specific deployed contract.
func NewMocks(address common.Address, backend bind.ContractBackend) (*Mocks, error) {
	contract, err := bindMocks(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mocks{MocksCaller: MocksCaller{contract: contract}, MocksTransactor: MocksTransactor{contract: contract}, MocksFilterer: MocksFilterer{contract: contract}}, nil
}

// NewMocksCaller creates a new read-only instance of Mocks, bound to a specific deployed contract.
func NewMocksCaller(address common.Address, caller bind.ContractCaller) (*MocksCaller, error) {
	contract, err := bindMocks(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MocksCaller{contract: contract}, nil
}

// NewMocksTransactor creates a new write-only instance of Mocks, bound to a specific deployed contract.
func NewMocksTransactor(address common.Address, transactor bind.ContractTransactor) (*MocksTransactor, error) {
	contract, err := bindMocks(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MocksTransactor{contract: contract}, nil
}

// NewMocksFilterer creates a new log filterer instance of Mocks, bound to a specific deployed contract.
func NewMocksFilterer(address common.Address, filterer bind.ContractFilterer) (*MocksFilterer, error) {
	contract, err := bindMocks(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MocksFilterer{contract: contract}, nil
}

// bindMocks binds a generic wrapper to an already deployed contract.
func bindMocks(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MocksABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mocks *MocksRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mocks.Contract.MocksCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mocks *MocksRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mocks.Contract.MocksTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mocks *MocksRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mocks.Contract.MocksTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mocks *MocksCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mocks.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mocks *MocksTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mocks.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mocks *MocksTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mocks.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mocks *MocksCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mocks.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mocks *MocksSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Mocks.Contract.Allowance(&_Mocks.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mocks *MocksCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Mocks.Contract.Allowance(&_Mocks.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mocks *MocksCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mocks.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mocks *MocksSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Mocks.Contract.BalanceOf(&_Mocks.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mocks *MocksCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Mocks.Contract.BalanceOf(&_Mocks.CallOpts, account)
}

// CosmosDenom is a free data retrieval call binding the contract method 0xde17fcc7.
//
// Solidity: function cosmosDenom() view returns(string)
func (_Mocks *MocksCaller) CosmosDenom(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mocks.contract.Call(opts, &out, "cosmosDenom")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CosmosDenom is a free data retrieval call binding the contract method 0xde17fcc7.
//
// Solidity: function cosmosDenom() view returns(string)
func (_Mocks *MocksSession) CosmosDenom() (string, error) {
	return _Mocks.Contract.CosmosDenom(&_Mocks.CallOpts)
}

// CosmosDenom is a free data retrieval call binding the contract method 0xde17fcc7.
//
// Solidity: function cosmosDenom() view returns(string)
func (_Mocks *MocksCallerSession) CosmosDenom() (string, error) {
	return _Mocks.Contract.CosmosDenom(&_Mocks.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mocks *MocksCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Mocks.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mocks *MocksSession) Decimals() (uint8, error) {
	return _Mocks.Contract.Decimals(&_Mocks.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mocks *MocksCallerSession) Decimals() (uint8, error) {
	return _Mocks.Contract.Decimals(&_Mocks.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mocks *MocksCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mocks.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mocks *MocksSession) Name() (string, error) {
	return _Mocks.Contract.Name(&_Mocks.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mocks *MocksCallerSession) Name() (string, error) {
	return _Mocks.Contract.Name(&_Mocks.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mocks *MocksCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mocks.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mocks *MocksSession) Symbol() (string, error) {
	return _Mocks.Contract.Symbol(&_Mocks.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mocks *MocksCallerSession) Symbol() (string, error) {
	return _Mocks.Contract.Symbol(&_Mocks.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mocks *MocksCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mocks.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mocks *MocksSession) TotalSupply() (*big.Int, error) {
	return _Mocks.Contract.TotalSupply(&_Mocks.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mocks *MocksCallerSession) TotalSupply() (*big.Int, error) {
	return _Mocks.Contract.TotalSupply(&_Mocks.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mocks *MocksTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mocks.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mocks *MocksSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.Approve(&_Mocks.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mocks *MocksTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.Approve(&_Mocks.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Mocks *MocksTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Mocks.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Mocks *MocksSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.Burn(&_Mocks.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Mocks *MocksTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.Burn(&_Mocks.TransactOpts, amount)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns(bool)
func (_Mocks *MocksTransactor) Burn0(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mocks.contract.Transact(opts, "burn0", user, amount)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns(bool)
func (_Mocks *MocksSession) Burn0(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.Burn0(&_Mocks.TransactOpts, user, amount)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns(bool)
func (_Mocks *MocksTransactorSession) Burn0(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.Burn0(&_Mocks.TransactOpts, user, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address user, uint256 amount) returns()
func (_Mocks *MocksTransactor) BurnFrom(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mocks.contract.Transact(opts, "burnFrom", user, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address user, uint256 amount) returns()
func (_Mocks *MocksSession) BurnFrom(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.BurnFrom(&_Mocks.TransactOpts, user, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address user, uint256 amount) returns()
func (_Mocks *MocksTransactorSession) BurnFrom(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.BurnFrom(&_Mocks.TransactOpts, user, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Mocks *MocksTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Mocks.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Mocks *MocksSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.DecreaseAllowance(&_Mocks.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Mocks *MocksTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.DecreaseAllowance(&_Mocks.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Mocks *MocksTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Mocks.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Mocks *MocksSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.IncreaseAllowance(&_Mocks.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Mocks *MocksTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.IncreaseAllowance(&_Mocks.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address user, uint256 amount) returns(bool)
func (_Mocks *MocksTransactor) Mint(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mocks.contract.Transact(opts, "mint", user, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address user, uint256 amount) returns(bool)
func (_Mocks *MocksSession) Mint(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.Mint(&_Mocks.TransactOpts, user, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address user, uint256 amount) returns(bool)
func (_Mocks *MocksTransactorSession) Mint(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.Mint(&_Mocks.TransactOpts, user, amount)
}

// SetDenom is a paid mutator transaction binding the contract method 0xa2a1c9fe.
//
// Solidity: function setDenom(string denom) returns(bool)
func (_Mocks *MocksTransactor) SetDenom(opts *bind.TransactOpts, denom string) (*types.Transaction, error) {
	return _Mocks.contract.Transact(opts, "setDenom", denom)
}

// SetDenom is a paid mutator transaction binding the contract method 0xa2a1c9fe.
//
// Solidity: function setDenom(string denom) returns(bool)
func (_Mocks *MocksSession) SetDenom(denom string) (*types.Transaction, error) {
	return _Mocks.Contract.SetDenom(&_Mocks.TransactOpts, denom)
}

// SetDenom is a paid mutator transaction binding the contract method 0xa2a1c9fe.
//
// Solidity: function setDenom(string denom) returns(bool)
func (_Mocks *MocksTransactorSession) SetDenom(denom string) (*types.Transaction, error) {
	return _Mocks.Contract.SetDenom(&_Mocks.TransactOpts, denom)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Mocks *MocksTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mocks.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Mocks *MocksSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.Transfer(&_Mocks.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Mocks *MocksTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.Transfer(&_Mocks.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Mocks *MocksTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mocks.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Mocks *MocksSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.TransferFrom(&_Mocks.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Mocks *MocksTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Mocks.Contract.TransferFrom(&_Mocks.TransactOpts, from, to, value)
}

// MocksApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Mocks contract.
type MocksApprovalIterator struct {
	Event *MocksApproval // Event containing the contract specifics and raw log

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
func (it *MocksApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MocksApproval)
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
		it.Event = new(MocksApproval)
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
func (it *MocksApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MocksApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MocksApproval represents a Approval event raised by the Mocks contract.
type MocksApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mocks *MocksFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MocksApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Mocks.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MocksApprovalIterator{contract: _Mocks.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mocks *MocksFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MocksApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Mocks.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MocksApproval)
				if err := _Mocks.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Mocks *MocksFilterer) ParseApproval(log types.Log) (*MocksApproval, error) {
	event := new(MocksApproval)
	if err := _Mocks.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MocksTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Mocks contract.
type MocksTransferIterator struct {
	Event *MocksTransfer // Event containing the contract specifics and raw log

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
func (it *MocksTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MocksTransfer)
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
		it.Event = new(MocksTransfer)
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
func (it *MocksTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MocksTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MocksTransfer represents a Transfer event raised by the Mocks contract.
type MocksTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mocks *MocksFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MocksTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mocks.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MocksTransferIterator{contract: _Mocks.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mocks *MocksFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MocksTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mocks.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MocksTransfer)
				if err := _Mocks.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Mocks *MocksFilterer) ParseTransfer(log types.Log) (*MocksTransfer, error) {
	event := new(MocksTransfer)
	if err := _Mocks.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
