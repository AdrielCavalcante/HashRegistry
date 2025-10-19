// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hashregistry

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
	_ = abi.ConvertType
)

// HashRegistryMetaData contains all meta data concerning the HashRegistry contract.
var HashRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"HashRegistrado\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getAllHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getCID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getHashByIndex\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getHashDetails\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalHashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashCIDs\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashList\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"registrarHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"verificarHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// HashRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use HashRegistryMetaData.ABI instead.
var HashRegistryABI = HashRegistryMetaData.ABI

// HashRegistry is an auto generated Go binding around an Ethereum contract.
type HashRegistry struct {
	HashRegistryCaller     // Read-only binding to the contract
	HashRegistryTransactor // Write-only binding to the contract
	HashRegistryFilterer   // Log filterer for contract events
}

// HashRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashRegistrySession struct {
	Contract     *HashRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashRegistryCallerSession struct {
	Contract *HashRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// HashRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashRegistryTransactorSession struct {
	Contract     *HashRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// HashRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashRegistryRaw struct {
	Contract *HashRegistry // Generic contract binding to access the raw methods on
}

// HashRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashRegistryCallerRaw struct {
	Contract *HashRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// HashRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashRegistryTransactorRaw struct {
	Contract *HashRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashRegistry creates a new instance of HashRegistry, bound to a specific deployed contract.
func NewHashRegistry(address common.Address, backend bind.ContractBackend) (*HashRegistry, error) {
	contract, err := bindHashRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashRegistry{HashRegistryCaller: HashRegistryCaller{contract: contract}, HashRegistryTransactor: HashRegistryTransactor{contract: contract}, HashRegistryFilterer: HashRegistryFilterer{contract: contract}}, nil
}

// NewHashRegistryCaller creates a new read-only instance of HashRegistry, bound to a specific deployed contract.
func NewHashRegistryCaller(address common.Address, caller bind.ContractCaller) (*HashRegistryCaller, error) {
	contract, err := bindHashRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashRegistryCaller{contract: contract}, nil
}

// NewHashRegistryTransactor creates a new write-only instance of HashRegistry, bound to a specific deployed contract.
func NewHashRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*HashRegistryTransactor, error) {
	contract, err := bindHashRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashRegistryTransactor{contract: contract}, nil
}

// NewHashRegistryFilterer creates a new log filterer instance of HashRegistry, bound to a specific deployed contract.
func NewHashRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*HashRegistryFilterer, error) {
	contract, err := bindHashRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashRegistryFilterer{contract: contract}, nil
}

// bindHashRegistry binds a generic wrapper to an already deployed contract.
func bindHashRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HashRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashRegistry *HashRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashRegistry.Contract.HashRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashRegistry *HashRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashRegistry.Contract.HashRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashRegistry *HashRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashRegistry.Contract.HashRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashRegistry *HashRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashRegistry *HashRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashRegistry *HashRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetAllHashes is a free data retrieval call binding the contract method 0xfe99f742.
//
// Solidity: function getAllHashes() view returns(bytes32[])
func (_HashRegistry *HashRegistryCaller) GetAllHashes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _HashRegistry.contract.Call(opts, &out, "getAllHashes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetAllHashes is a free data retrieval call binding the contract method 0xfe99f742.
//
// Solidity: function getAllHashes() view returns(bytes32[])
func (_HashRegistry *HashRegistrySession) GetAllHashes() ([][32]byte, error) {
	return _HashRegistry.Contract.GetAllHashes(&_HashRegistry.CallOpts)
}

// GetAllHashes is a free data retrieval call binding the contract method 0xfe99f742.
//
// Solidity: function getAllHashes() view returns(bytes32[])
func (_HashRegistry *HashRegistryCallerSession) GetAllHashes() ([][32]byte, error) {
	return _HashRegistry.Contract.GetAllHashes(&_HashRegistry.CallOpts)
}

// GetCID is a free data retrieval call binding the contract method 0x39ecade2.
//
// Solidity: function getCID(bytes32 hash) view returns(string)
func (_HashRegistry *HashRegistryCaller) GetCID(opts *bind.CallOpts, hash [32]byte) (string, error) {
	var out []interface{}
	err := _HashRegistry.contract.Call(opts, &out, "getCID", hash)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCID is a free data retrieval call binding the contract method 0x39ecade2.
//
// Solidity: function getCID(bytes32 hash) view returns(string)
func (_HashRegistry *HashRegistrySession) GetCID(hash [32]byte) (string, error) {
	return _HashRegistry.Contract.GetCID(&_HashRegistry.CallOpts, hash)
}

// GetCID is a free data retrieval call binding the contract method 0x39ecade2.
//
// Solidity: function getCID(bytes32 hash) view returns(string)
func (_HashRegistry *HashRegistryCallerSession) GetCID(hash [32]byte) (string, error) {
	return _HashRegistry.Contract.GetCID(&_HashRegistry.CallOpts, hash)
}

// GetHashByIndex is a free data retrieval call binding the contract method 0xec058186.
//
// Solidity: function getHashByIndex(uint256 index) view returns(bytes32)
func (_HashRegistry *HashRegistryCaller) GetHashByIndex(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _HashRegistry.contract.Call(opts, &out, "getHashByIndex", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHashByIndex is a free data retrieval call binding the contract method 0xec058186.
//
// Solidity: function getHashByIndex(uint256 index) view returns(bytes32)
func (_HashRegistry *HashRegistrySession) GetHashByIndex(index *big.Int) ([32]byte, error) {
	return _HashRegistry.Contract.GetHashByIndex(&_HashRegistry.CallOpts, index)
}

// GetHashByIndex is a free data retrieval call binding the contract method 0xec058186.
//
// Solidity: function getHashByIndex(uint256 index) view returns(bytes32)
func (_HashRegistry *HashRegistryCallerSession) GetHashByIndex(index *big.Int) ([32]byte, error) {
	return _HashRegistry.Contract.GetHashByIndex(&_HashRegistry.CallOpts, index)
}

// GetHashDetails is a free data retrieval call binding the contract method 0x1758c77a.
//
// Solidity: function getHashDetails(bytes32 hash) view returns(bool exists, uint256 timestamp, address owner, string cid)
func (_HashRegistry *HashRegistryCaller) GetHashDetails(opts *bind.CallOpts, hash [32]byte) (struct {
	Exists    bool
	Timestamp *big.Int
	Owner     common.Address
	Cid       string
}, error) {
	var out []interface{}
	err := _HashRegistry.contract.Call(opts, &out, "getHashDetails", hash)

	outstruct := new(struct {
		Exists    bool
		Timestamp *big.Int
		Owner     common.Address
		Cid       string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Cid = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// GetHashDetails is a free data retrieval call binding the contract method 0x1758c77a.
//
// Solidity: function getHashDetails(bytes32 hash) view returns(bool exists, uint256 timestamp, address owner, string cid)
func (_HashRegistry *HashRegistrySession) GetHashDetails(hash [32]byte) (struct {
	Exists    bool
	Timestamp *big.Int
	Owner     common.Address
	Cid       string
}, error) {
	return _HashRegistry.Contract.GetHashDetails(&_HashRegistry.CallOpts, hash)
}

// GetHashDetails is a free data retrieval call binding the contract method 0x1758c77a.
//
// Solidity: function getHashDetails(bytes32 hash) view returns(bool exists, uint256 timestamp, address owner, string cid)
func (_HashRegistry *HashRegistryCallerSession) GetHashDetails(hash [32]byte) (struct {
	Exists    bool
	Timestamp *big.Int
	Owner     common.Address
	Cid       string
}, error) {
	return _HashRegistry.Contract.GetHashDetails(&_HashRegistry.CallOpts, hash)
}

// GetTotalHashes is a free data retrieval call binding the contract method 0xa4458dc4.
//
// Solidity: function getTotalHashes() view returns(uint256)
func (_HashRegistry *HashRegistryCaller) GetTotalHashes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HashRegistry.contract.Call(opts, &out, "getTotalHashes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalHashes is a free data retrieval call binding the contract method 0xa4458dc4.
//
// Solidity: function getTotalHashes() view returns(uint256)
func (_HashRegistry *HashRegistrySession) GetTotalHashes() (*big.Int, error) {
	return _HashRegistry.Contract.GetTotalHashes(&_HashRegistry.CallOpts)
}

// GetTotalHashes is a free data retrieval call binding the contract method 0xa4458dc4.
//
// Solidity: function getTotalHashes() view returns(uint256)
func (_HashRegistry *HashRegistryCallerSession) GetTotalHashes() (*big.Int, error) {
	return _HashRegistry.Contract.GetTotalHashes(&_HashRegistry.CallOpts)
}

// HashCIDs is a free data retrieval call binding the contract method 0xaaf9cc3b.
//
// Solidity: function hashCIDs(bytes32 ) view returns(string)
func (_HashRegistry *HashRegistryCaller) HashCIDs(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var out []interface{}
	err := _HashRegistry.contract.Call(opts, &out, "hashCIDs", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// HashCIDs is a free data retrieval call binding the contract method 0xaaf9cc3b.
//
// Solidity: function hashCIDs(bytes32 ) view returns(string)
func (_HashRegistry *HashRegistrySession) HashCIDs(arg0 [32]byte) (string, error) {
	return _HashRegistry.Contract.HashCIDs(&_HashRegistry.CallOpts, arg0)
}

// HashCIDs is a free data retrieval call binding the contract method 0xaaf9cc3b.
//
// Solidity: function hashCIDs(bytes32 ) view returns(string)
func (_HashRegistry *HashRegistryCallerSession) HashCIDs(arg0 [32]byte) (string, error) {
	return _HashRegistry.Contract.HashCIDs(&_HashRegistry.CallOpts, arg0)
}

// HashList is a free data retrieval call binding the contract method 0x2c8d41fd.
//
// Solidity: function hashList(uint256 ) view returns(bytes32)
func (_HashRegistry *HashRegistryCaller) HashList(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _HashRegistry.contract.Call(opts, &out, "hashList", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashList is a free data retrieval call binding the contract method 0x2c8d41fd.
//
// Solidity: function hashList(uint256 ) view returns(bytes32)
func (_HashRegistry *HashRegistrySession) HashList(arg0 *big.Int) ([32]byte, error) {
	return _HashRegistry.Contract.HashList(&_HashRegistry.CallOpts, arg0)
}

// HashList is a free data retrieval call binding the contract method 0x2c8d41fd.
//
// Solidity: function hashList(uint256 ) view returns(bytes32)
func (_HashRegistry *HashRegistryCallerSession) HashList(arg0 *big.Int) ([32]byte, error) {
	return _HashRegistry.Contract.HashList(&_HashRegistry.CallOpts, arg0)
}

// HashOwners is a free data retrieval call binding the contract method 0xf76c9394.
//
// Solidity: function hashOwners(bytes32 ) view returns(address)
func (_HashRegistry *HashRegistryCaller) HashOwners(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _HashRegistry.contract.Call(opts, &out, "hashOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HashOwners is a free data retrieval call binding the contract method 0xf76c9394.
//
// Solidity: function hashOwners(bytes32 ) view returns(address)
func (_HashRegistry *HashRegistrySession) HashOwners(arg0 [32]byte) (common.Address, error) {
	return _HashRegistry.Contract.HashOwners(&_HashRegistry.CallOpts, arg0)
}

// HashOwners is a free data retrieval call binding the contract method 0xf76c9394.
//
// Solidity: function hashOwners(bytes32 ) view returns(address)
func (_HashRegistry *HashRegistryCallerSession) HashOwners(arg0 [32]byte) (common.Address, error) {
	return _HashRegistry.Contract.HashOwners(&_HashRegistry.CallOpts, arg0)
}

// HashTimestamps is a free data retrieval call binding the contract method 0x3670430b.
//
// Solidity: function hashTimestamps(bytes32 ) view returns(uint256)
func (_HashRegistry *HashRegistryCaller) HashTimestamps(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _HashRegistry.contract.Call(opts, &out, "hashTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashTimestamps is a free data retrieval call binding the contract method 0x3670430b.
//
// Solidity: function hashTimestamps(bytes32 ) view returns(uint256)
func (_HashRegistry *HashRegistrySession) HashTimestamps(arg0 [32]byte) (*big.Int, error) {
	return _HashRegistry.Contract.HashTimestamps(&_HashRegistry.CallOpts, arg0)
}

// HashTimestamps is a free data retrieval call binding the contract method 0x3670430b.
//
// Solidity: function hashTimestamps(bytes32 ) view returns(uint256)
func (_HashRegistry *HashRegistryCallerSession) HashTimestamps(arg0 [32]byte) (*big.Int, error) {
	return _HashRegistry.Contract.HashTimestamps(&_HashRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0xd658d2e9.
//
// Solidity: function hashes(bytes32 ) view returns(bool)
func (_HashRegistry *HashRegistryCaller) Hashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _HashRegistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0xd658d2e9.
//
// Solidity: function hashes(bytes32 ) view returns(bool)
func (_HashRegistry *HashRegistrySession) Hashes(arg0 [32]byte) (bool, error) {
	return _HashRegistry.Contract.Hashes(&_HashRegistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0xd658d2e9.
//
// Solidity: function hashes(bytes32 ) view returns(bool)
func (_HashRegistry *HashRegistryCallerSession) Hashes(arg0 [32]byte) (bool, error) {
	return _HashRegistry.Contract.Hashes(&_HashRegistry.CallOpts, arg0)
}

// VerificarHash is a free data retrieval call binding the contract method 0xdcfbc431.
//
// Solidity: function verificarHash(bytes32 hash) view returns(bool)
func (_HashRegistry *HashRegistryCaller) VerificarHash(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _HashRegistry.contract.Call(opts, &out, "verificarHash", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerificarHash is a free data retrieval call binding the contract method 0xdcfbc431.
//
// Solidity: function verificarHash(bytes32 hash) view returns(bool)
func (_HashRegistry *HashRegistrySession) VerificarHash(hash [32]byte) (bool, error) {
	return _HashRegistry.Contract.VerificarHash(&_HashRegistry.CallOpts, hash)
}

// VerificarHash is a free data retrieval call binding the contract method 0xdcfbc431.
//
// Solidity: function verificarHash(bytes32 hash) view returns(bool)
func (_HashRegistry *HashRegistryCallerSession) VerificarHash(hash [32]byte) (bool, error) {
	return _HashRegistry.Contract.VerificarHash(&_HashRegistry.CallOpts, hash)
}

// RegistrarHash is a paid mutator transaction binding the contract method 0x99e61a6b.
//
// Solidity: function registrarHash(bytes32 hash, string cid) returns()
func (_HashRegistry *HashRegistryTransactor) RegistrarHash(opts *bind.TransactOpts, hash [32]byte, cid string) (*types.Transaction, error) {
	return _HashRegistry.contract.Transact(opts, "registrarHash", hash, cid)
}

// RegistrarHash is a paid mutator transaction binding the contract method 0x99e61a6b.
//
// Solidity: function registrarHash(bytes32 hash, string cid) returns()
func (_HashRegistry *HashRegistrySession) RegistrarHash(hash [32]byte, cid string) (*types.Transaction, error) {
	return _HashRegistry.Contract.RegistrarHash(&_HashRegistry.TransactOpts, hash, cid)
}

// RegistrarHash is a paid mutator transaction binding the contract method 0x99e61a6b.
//
// Solidity: function registrarHash(bytes32 hash, string cid) returns()
func (_HashRegistry *HashRegistryTransactorSession) RegistrarHash(hash [32]byte, cid string) (*types.Transaction, error) {
	return _HashRegistry.Contract.RegistrarHash(&_HashRegistry.TransactOpts, hash, cid)
}

// HashRegistryHashRegistradoIterator is returned from FilterHashRegistrado and is used to iterate over the raw logs and unpacked data for HashRegistrado events raised by the HashRegistry contract.
type HashRegistryHashRegistradoIterator struct {
	Event *HashRegistryHashRegistrado // Event containing the contract specifics and raw log

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
func (it *HashRegistryHashRegistradoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashRegistryHashRegistrado)
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
		it.Event = new(HashRegistryHashRegistrado)
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
func (it *HashRegistryHashRegistradoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashRegistryHashRegistradoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashRegistryHashRegistrado represents a HashRegistrado event raised by the HashRegistry contract.
type HashRegistryHashRegistrado struct {
	Hash      [32]byte
	Owner     common.Address
	Timestamp *big.Int
	Cid       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHashRegistrado is a free log retrieval operation binding the contract event 0xbed179fa938eaaafadded12e0319dfcb4101d2e491c92518e8042d4cb5b9c78b.
//
// Solidity: event HashRegistrado(bytes32 hash, address owner, uint256 timestamp, string cid)
func (_HashRegistry *HashRegistryFilterer) FilterHashRegistrado(opts *bind.FilterOpts) (*HashRegistryHashRegistradoIterator, error) {

	logs, sub, err := _HashRegistry.contract.FilterLogs(opts, "HashRegistrado")
	if err != nil {
		return nil, err
	}
	return &HashRegistryHashRegistradoIterator{contract: _HashRegistry.contract, event: "HashRegistrado", logs: logs, sub: sub}, nil
}

// WatchHashRegistrado is a free log subscription operation binding the contract event 0xbed179fa938eaaafadded12e0319dfcb4101d2e491c92518e8042d4cb5b9c78b.
//
// Solidity: event HashRegistrado(bytes32 hash, address owner, uint256 timestamp, string cid)
func (_HashRegistry *HashRegistryFilterer) WatchHashRegistrado(opts *bind.WatchOpts, sink chan<- *HashRegistryHashRegistrado) (event.Subscription, error) {

	logs, sub, err := _HashRegistry.contract.WatchLogs(opts, "HashRegistrado")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashRegistryHashRegistrado)
				if err := _HashRegistry.contract.UnpackLog(event, "HashRegistrado", log); err != nil {
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

// ParseHashRegistrado is a log parse operation binding the contract event 0xbed179fa938eaaafadded12e0319dfcb4101d2e491c92518e8042d4cb5b9c78b.
//
// Solidity: event HashRegistrado(bytes32 hash, address owner, uint256 timestamp, string cid)
func (_HashRegistry *HashRegistryFilterer) ParseHashRegistrado(log types.Log) (*HashRegistryHashRegistrado, error) {
	event := new(HashRegistryHashRegistrado)
	if err := _HashRegistry.contract.UnpackLog(event, "HashRegistrado", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
