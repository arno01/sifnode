// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MockUpgrade

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

// MockUpgradeMetaData contains all meta data concerning the MockUpgrade contract.
var MockUpgradeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MockUpgradeABI is the input ABI used to generate the binding from.
// Deprecated: Use MockUpgradeMetaData.ABI instead.
var MockUpgradeABI = MockUpgradeMetaData.ABI

// MockUpgrade is an auto generated Go binding around an Ethereum contract.
type MockUpgrade struct {
	MockUpgradeCaller     // Read-only binding to the contract
	MockUpgradeTransactor // Write-only binding to the contract
	MockUpgradeFilterer   // Log filterer for contract events
}

// MockUpgradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockUpgradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockUpgradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockUpgradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockUpgradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockUpgradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockUpgradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockUpgradeSession struct {
	Contract     *MockUpgrade      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockUpgradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockUpgradeCallerSession struct {
	Contract *MockUpgradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MockUpgradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockUpgradeTransactorSession struct {
	Contract     *MockUpgradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MockUpgradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockUpgradeRaw struct {
	Contract *MockUpgrade // Generic contract binding to access the raw methods on
}

// MockUpgradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockUpgradeCallerRaw struct {
	Contract *MockUpgradeCaller // Generic read-only contract binding to access the raw methods on
}

// MockUpgradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockUpgradeTransactorRaw struct {
	Contract *MockUpgradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockUpgrade creates a new instance of MockUpgrade, bound to a specific deployed contract.
func NewMockUpgrade(address common.Address, backend bind.ContractBackend) (*MockUpgrade, error) {
	contract, err := bindMockUpgrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockUpgrade{MockUpgradeCaller: MockUpgradeCaller{contract: contract}, MockUpgradeTransactor: MockUpgradeTransactor{contract: contract}, MockUpgradeFilterer: MockUpgradeFilterer{contract: contract}}, nil
}

// NewMockUpgradeCaller creates a new read-only instance of MockUpgrade, bound to a specific deployed contract.
func NewMockUpgradeCaller(address common.Address, caller bind.ContractCaller) (*MockUpgradeCaller, error) {
	contract, err := bindMockUpgrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockUpgradeCaller{contract: contract}, nil
}

// NewMockUpgradeTransactor creates a new write-only instance of MockUpgrade, bound to a specific deployed contract.
func NewMockUpgradeTransactor(address common.Address, transactor bind.ContractTransactor) (*MockUpgradeTransactor, error) {
	contract, err := bindMockUpgrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockUpgradeTransactor{contract: contract}, nil
}

// NewMockUpgradeFilterer creates a new log filterer instance of MockUpgrade, bound to a specific deployed contract.
func NewMockUpgradeFilterer(address common.Address, filterer bind.ContractFilterer) (*MockUpgradeFilterer, error) {
	contract, err := bindMockUpgrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockUpgradeFilterer{contract: contract}, nil
}

// bindMockUpgrade binds a generic wrapper to an already deployed contract.
func bindMockUpgrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockUpgradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockUpgrade *MockUpgradeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockUpgrade.Contract.MockUpgradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockUpgrade *MockUpgradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockUpgrade.Contract.MockUpgradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockUpgrade *MockUpgradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockUpgrade.Contract.MockUpgradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockUpgrade *MockUpgradeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockUpgrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockUpgrade *MockUpgradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockUpgrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockUpgrade *MockUpgradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockUpgrade.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockUpgrade *MockUpgradeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockUpgrade.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockUpgrade *MockUpgradeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MockUpgrade.Contract.BalanceOf(&_MockUpgrade.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockUpgrade *MockUpgradeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MockUpgrade.Contract.BalanceOf(&_MockUpgrade.CallOpts, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MockUpgrade *MockUpgradeTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUpgrade.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MockUpgrade *MockUpgradeSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUpgrade.Contract.Transfer(&_MockUpgrade.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MockUpgrade *MockUpgradeTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUpgrade.Contract.Transfer(&_MockUpgrade.TransactOpts, recipient, amount)
}
