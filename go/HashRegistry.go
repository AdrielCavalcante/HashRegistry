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

// HashregistryMetaData contains all meta data concerning the Hashregistry contract.
var HashregistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"HashRegistrado\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getAllHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getHashByIndex\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getHashDetails\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalHashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashList\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"registrarHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"verificarHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// HashregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use HashregistryMetaData.ABI instead.
var HashregistryABI = HashregistryMetaData.ABI

// Hashregistry is an auto generated Go binding around an Ethereum contract.
type Hashregistry struct {
	HashregistryCaller     // Read-only binding to the contract
	HashregistryTransactor // Write-only binding to the contract
	HashregistryFilterer   // Log filterer for contract events
}

// HashregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashregistrySession struct {
	Contract     *Hashregistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashregistryCallerSession struct {
	Contract *HashregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// HashregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashregistryTransactorSession struct {
	Contract     *HashregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// HashregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashregistryRaw struct {
	Contract *Hashregistry // Generic contract binding to access the raw methods on
}

// HashregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashregistryCallerRaw struct {
	Contract *HashregistryCaller // Generic read-only contract binding to access the raw methods on
}

// HashregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashregistryTransactorRaw struct {
	Contract *HashregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashregistry creates a new instance of Hashregistry, bound to a specific deployed contract.
func NewHashregistry(address common.Address, backend bind.ContractBackend) (*Hashregistry, error) {
	contract, err := bindHashregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hashregistry{HashregistryCaller: HashregistryCaller{contract: contract}, HashregistryTransactor: HashregistryTransactor{contract: contract}, HashregistryFilterer: HashregistryFilterer{contract: contract}}, nil
}

// NewHashregistryCaller creates a new read-only instance of Hashregistry, bound to a specific deployed contract.
func NewHashregistryCaller(address common.Address, caller bind.ContractCaller) (*HashregistryCaller, error) {
	contract, err := bindHashregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashregistryCaller{contract: contract}, nil
}

// NewHashregistryTransactor creates a new write-only instance of Hashregistry, bound to a specific deployed contract.
func NewHashregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*HashregistryTransactor, error) {
	contract, err := bindHashregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashregistryTransactor{contract: contract}, nil
}

// NewHashregistryFilterer creates a new log filterer instance of Hashregistry, bound to a specific deployed contract.
func NewHashregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*HashregistryFilterer, error) {
	contract, err := bindHashregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashregistryFilterer{contract: contract}, nil
}

// bindHashregistry binds a generic wrapper to an already deployed contract.
func bindHashregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HashregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hashregistry *HashregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hashregistry.Contract.HashregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hashregistry *HashregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hashregistry.Contract.HashregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hashregistry *HashregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hashregistry.Contract.HashregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hashregistry *HashregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hashregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hashregistry *HashregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hashregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hashregistry *HashregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hashregistry.Contract.contract.Transact(opts, method, params...)
}

// GetAllHashes is a free data retrieval call binding the contract method 0xfe99f742.
//
// Solidity: function getAllHashes() view returns(bytes32[])
func (_Hashregistry *HashregistryCaller) GetAllHashes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Hashregistry.contract.Call(opts, &out, "getAllHashes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetAllHashes is a free data retrieval call binding the contract method 0xfe99f742.
//
// Solidity: function getAllHashes() view returns(bytes32[])
func (_Hashregistry *HashregistrySession) GetAllHashes() ([][32]byte, error) {
	return _Hashregistry.Contract.GetAllHashes(&_Hashregistry.CallOpts)
}

// GetAllHashes is a free data retrieval call binding the contract method 0xfe99f742.
//
// Solidity: function getAllHashes() view returns(bytes32[])
func (_Hashregistry *HashregistryCallerSession) GetAllHashes() ([][32]byte, error) {
	return _Hashregistry.Contract.GetAllHashes(&_Hashregistry.CallOpts)
}

// GetHashByIndex is a free data retrieval call binding the contract method 0xec058186.
//
// Solidity: function getHashByIndex(uint256 index) view returns(bytes32)
func (_Hashregistry *HashregistryCaller) GetHashByIndex(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Hashregistry.contract.Call(opts, &out, "getHashByIndex", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHashByIndex is a free data retrieval call binding the contract method 0xec058186.
//
// Solidity: function getHashByIndex(uint256 index) view returns(bytes32)
func (_Hashregistry *HashregistrySession) GetHashByIndex(index *big.Int) ([32]byte, error) {
	return _Hashregistry.Contract.GetHashByIndex(&_Hashregistry.CallOpts, index)
}

// GetHashByIndex is a free data retrieval call binding the contract method 0xec058186.
//
// Solidity: function getHashByIndex(uint256 index) view returns(bytes32)
func (_Hashregistry *HashregistryCallerSession) GetHashByIndex(index *big.Int) ([32]byte, error) {
	return _Hashregistry.Contract.GetHashByIndex(&_Hashregistry.CallOpts, index)
}

// GetHashDetails is a free data retrieval call binding the contract method 0x1758c77a.
//
// Solidity: function getHashDetails(bytes32 hash) view returns(bool exists, uint256 timestamp, address owner)
func (_Hashregistry *HashregistryCaller) GetHashDetails(opts *bind.CallOpts, hash [32]byte) (struct {
	Exists    bool
	Timestamp *big.Int
	Owner     common.Address
}, error) {
	var out []interface{}
	err := _Hashregistry.contract.Call(opts, &out, "getHashDetails", hash)

	outstruct := new(struct {
		Exists    bool
		Timestamp *big.Int
		Owner     common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetHashDetails is a free data retrieval call binding the contract method 0x1758c77a.
//
// Solidity: function getHashDetails(bytes32 hash) view returns(bool exists, uint256 timestamp, address owner)
func (_Hashregistry *HashregistrySession) GetHashDetails(hash [32]byte) (struct {
	Exists    bool
	Timestamp *big.Int
	Owner     common.Address
}, error) {
	return _Hashregistry.Contract.GetHashDetails(&_Hashregistry.CallOpts, hash)
}

// GetHashDetails is a free data retrieval call binding the contract method 0x1758c77a.
//
// Solidity: function getHashDetails(bytes32 hash) view returns(bool exists, uint256 timestamp, address owner)
func (_Hashregistry *HashregistryCallerSession) GetHashDetails(hash [32]byte) (struct {
	Exists    bool
	Timestamp *big.Int
	Owner     common.Address
}, error) {
	return _Hashregistry.Contract.GetHashDetails(&_Hashregistry.CallOpts, hash)
}

// GetTotalHashes is a free data retrieval call binding the contract method 0xa4458dc4.
//
// Solidity: function getTotalHashes() view returns(uint256)
func (_Hashregistry *HashregistryCaller) GetTotalHashes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hashregistry.contract.Call(opts, &out, "getTotalHashes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalHashes is a free data retrieval call binding the contract method 0xa4458dc4.
//
// Solidity: function getTotalHashes() view returns(uint256)
func (_Hashregistry *HashregistrySession) GetTotalHashes() (*big.Int, error) {
	return _Hashregistry.Contract.GetTotalHashes(&_Hashregistry.CallOpts)
}

// GetTotalHashes is a free data retrieval call binding the contract method 0xa4458dc4.
//
// Solidity: function getTotalHashes() view returns(uint256)
func (_Hashregistry *HashregistryCallerSession) GetTotalHashes() (*big.Int, error) {
	return _Hashregistry.Contract.GetTotalHashes(&_Hashregistry.CallOpts)
}

// HashList is a free data retrieval call binding the contract method 0x2c8d41fd.
//
// Solidity: function hashList(uint256 ) view returns(bytes32)
func (_Hashregistry *HashregistryCaller) HashList(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Hashregistry.contract.Call(opts, &out, "hashList", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashList is a free data retrieval call binding the contract method 0x2c8d41fd.
//
// Solidity: function hashList(uint256 ) view returns(bytes32)
func (_Hashregistry *HashregistrySession) HashList(arg0 *big.Int) ([32]byte, error) {
	return _Hashregistry.Contract.HashList(&_Hashregistry.CallOpts, arg0)
}

// HashList is a free data retrieval call binding the contract method 0x2c8d41fd.
//
// Solidity: function hashList(uint256 ) view returns(bytes32)
func (_Hashregistry *HashregistryCallerSession) HashList(arg0 *big.Int) ([32]byte, error) {
	return _Hashregistry.Contract.HashList(&_Hashregistry.CallOpts, arg0)
}

// HashOwners is a free data retrieval call binding the contract method 0xf76c9394.
//
// Solidity: function hashOwners(bytes32 ) view returns(address)
func (_Hashregistry *HashregistryCaller) HashOwners(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Hashregistry.contract.Call(opts, &out, "hashOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HashOwners is a free data retrieval call binding the contract method 0xf76c9394.
//
// Solidity: function hashOwners(bytes32 ) view returns(address)
func (_Hashregistry *HashregistrySession) HashOwners(arg0 [32]byte) (common.Address, error) {
	return _Hashregistry.Contract.HashOwners(&_Hashregistry.CallOpts, arg0)
}

// HashOwners is a free data retrieval call binding the contract method 0xf76c9394.
//
// Solidity: function hashOwners(bytes32 ) view returns(address)
func (_Hashregistry *HashregistryCallerSession) HashOwners(arg0 [32]byte) (common.Address, error) {
	return _Hashregistry.Contract.HashOwners(&_Hashregistry.CallOpts, arg0)
}

// HashTimestamps is a free data retrieval call binding the contract method 0x3670430b.
//
// Solidity: function hashTimestamps(bytes32 ) view returns(uint256)
func (_Hashregistry *HashregistryCaller) HashTimestamps(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Hashregistry.contract.Call(opts, &out, "hashTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashTimestamps is a free data retrieval call binding the contract method 0x3670430b.
//
// Solidity: function hashTimestamps(bytes32 ) view returns(uint256)
func (_Hashregistry *HashregistrySession) HashTimestamps(arg0 [32]byte) (*big.Int, error) {
	return _Hashregistry.Contract.HashTimestamps(&_Hashregistry.CallOpts, arg0)
}

// HashTimestamps is a free data retrieval call binding the contract method 0x3670430b.
//
// Solidity: function hashTimestamps(bytes32 ) view returns(uint256)
func (_Hashregistry *HashregistryCallerSession) HashTimestamps(arg0 [32]byte) (*big.Int, error) {
	return _Hashregistry.Contract.HashTimestamps(&_Hashregistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0xd658d2e9.
//
// Solidity: function hashes(bytes32 ) view returns(bool)
func (_Hashregistry *HashregistryCaller) Hashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Hashregistry.contract.Call(opts, &out, "hashes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Hashes is a free data retrieval call binding the contract method 0xd658d2e9.
//
// Solidity: function hashes(bytes32 ) view returns(bool)
func (_Hashregistry *HashregistrySession) Hashes(arg0 [32]byte) (bool, error) {
	return _Hashregistry.Contract.Hashes(&_Hashregistry.CallOpts, arg0)
}

// Hashes is a free data retrieval call binding the contract method 0xd658d2e9.
//
// Solidity: function hashes(bytes32 ) view returns(bool)
func (_Hashregistry *HashregistryCallerSession) Hashes(arg0 [32]byte) (bool, error) {
	return _Hashregistry.Contract.Hashes(&_Hashregistry.CallOpts, arg0)
}

// VerificarHash is a free data retrieval call binding the contract method 0xdcfbc431.
//
// Solidity: function verificarHash(bytes32 hash) view returns(bool)
func (_Hashregistry *HashregistryCaller) VerificarHash(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _Hashregistry.contract.Call(opts, &out, "verificarHash", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerificarHash is a free data retrieval call binding the contract method 0xdcfbc431.
//
// Solidity: function verificarHash(bytes32 hash) view returns(bool)
func (_Hashregistry *HashregistrySession) VerificarHash(hash [32]byte) (bool, error) {
	return _Hashregistry.Contract.VerificarHash(&_Hashregistry.CallOpts, hash)
}

// VerificarHash is a free data retrieval call binding the contract method 0xdcfbc431.
//
// Solidity: function verificarHash(bytes32 hash) view returns(bool)
func (_Hashregistry *HashregistryCallerSession) VerificarHash(hash [32]byte) (bool, error) {
	return _Hashregistry.Contract.VerificarHash(&_Hashregistry.CallOpts, hash)
}

// RegistrarHash is a paid mutator transaction binding the contract method 0x7eb41ee8.
//
// Solidity: function registrarHash(bytes32 hash) returns()
func (_Hashregistry *HashregistryTransactor) RegistrarHash(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Hashregistry.contract.Transact(opts, "registrarHash", hash)
}

// RegistrarHash is a paid mutator transaction binding the contract method 0x7eb41ee8.
//
// Solidity: function registrarHash(bytes32 hash) returns()
func (_Hashregistry *HashregistrySession) RegistrarHash(hash [32]byte) (*types.Transaction, error) {
	return _Hashregistry.Contract.RegistrarHash(&_Hashregistry.TransactOpts, hash)
}

// RegistrarHash is a paid mutator transaction binding the contract method 0x7eb41ee8.
//
// Solidity: function registrarHash(bytes32 hash) returns()
func (_Hashregistry *HashregistryTransactorSession) RegistrarHash(hash [32]byte) (*types.Transaction, error) {
	return _Hashregistry.Contract.RegistrarHash(&_Hashregistry.TransactOpts, hash)
}

// HashregistryHashRegistradoIterator is returned from FilterHashRegistrado and is used to iterate over the raw logs and unpacked data for HashRegistrado events raised by the Hashregistry contract.
type HashregistryHashRegistradoIterator struct {
	Event *HashregistryHashRegistrado // Event containing the contract specifics and raw log

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
func (it *HashregistryHashRegistradoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashregistryHashRegistrado)
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
		it.Event = new(HashregistryHashRegistrado)
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
func (it *HashregistryHashRegistradoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashregistryHashRegistradoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashregistryHashRegistrado represents a HashRegistrado event raised by the Hashregistry contract.
type HashregistryHashRegistrado struct {
	Hash      [32]byte
	Owner     common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHashRegistrado is a free log retrieval operation binding the contract event 0x9f874d9673e1cbf2f13c1b5ad2e8dad7f871a3c62723c97a7eb230e757230e62.
//
// Solidity: event HashRegistrado(bytes32 hash, address owner, uint256 timestamp)
func (_Hashregistry *HashregistryFilterer) FilterHashRegistrado(opts *bind.FilterOpts) (*HashregistryHashRegistradoIterator, error) {

	logs, sub, err := _Hashregistry.contract.FilterLogs(opts, "HashRegistrado")
	if err != nil {
		return nil, err
	}
	return &HashregistryHashRegistradoIterator{contract: _Hashregistry.contract, event: "HashRegistrado", logs: logs, sub: sub}, nil
}

// WatchHashRegistrado is a free log subscription operation binding the contract event 0x9f874d9673e1cbf2f13c1b5ad2e8dad7f871a3c62723c97a7eb230e757230e62.
//
// Solidity: event HashRegistrado(bytes32 hash, address owner, uint256 timestamp)
func (_Hashregistry *HashregistryFilterer) WatchHashRegistrado(opts *bind.WatchOpts, sink chan<- *HashregistryHashRegistrado) (event.Subscription, error) {

	logs, sub, err := _Hashregistry.contract.WatchLogs(opts, "HashRegistrado")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashregistryHashRegistrado)
				if err := _Hashregistry.contract.UnpackLog(event, "HashRegistrado", log); err != nil {
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

// ParseHashRegistrado is a log parse operation binding the contract event 0x9f874d9673e1cbf2f13c1b5ad2e8dad7f871a3c62723c97a7eb230e757230e62.
//
// Solidity: event HashRegistrado(bytes32 hash, address owner, uint256 timestamp)
func (_Hashregistry *HashregistryFilterer) ParseHashRegistrado(log types.Log) (*HashregistryHashRegistrado, error) {
	event := new(HashregistryHashRegistrado)
	if err := _Hashregistry.contract.UnpackLog(event, "HashRegistrado", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
