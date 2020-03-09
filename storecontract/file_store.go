// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storecontract

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FileStoreABI is the input ABI used to generate the binding from.
const FileStoreABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"codeTx\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"indexHash\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"updateHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashTx\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_codeTx\",\"type\":\"bytes32\"}],\"name\":\"setCode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"core\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"bool\"}],\"name\":\"LockEvent\",\"type\":\"event\"}]"

// FileStoreBin is the compiled bytecode used for deploying new contracts.
const FileStoreBin = `0x608060405260018054600160a860020a0319167396216849c49358b10257cb55b28ea603c874b05e17905560028054600160a060020a031916331790556004805460ff19169055610742806100556000396000f3fe608060405234801561001057600080fd5b50600436106100d1576000357c0100000000000000000000000000000000000000000000000000000000900480639e7487eb1161008e5780639e7487eb146102e7578063a69df4b514610303578063b9ef767f1461030b578063cf30901214610328578063f2f4eb2614610330578063f83d08ba14610338576100d1565b806304fb3167146100d65780633d7403a3146100f05780635cb8eca4146101985780637dc0d1d0146102155780638da5cb5b1461023957806392dea92214610241575b600080fd5b6100de610340565b60408051918252519081900360200190f35b6101966004803603602081101561010657600080fd5b81019060208101813564010000000081111561012157600080fd5b82018360208201111561013357600080fd5b8035906020019184600183028401116401000000008311171561015557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610346945050505050565b005b6101a06103a2565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101da5781810151838201526020016101c2565b50505050905090810190601f1680156102075780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61021d610439565b60408051600160a060020a039092168252519081900360200190f35b61021d610448565b6101966004803603602081101561025757600080fd5b81019060208101813564010000000081111561027257600080fd5b82018360208201111561028457600080fd5b803590602001918460018302840111640100000000831117156102a657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610457945050505050565b6102ef61047a565b604080519115158252519081900360200190f35b610196610483565b6101966004803603602081101561032157600080fd5b5035610514565b6102ef61054d565b6101a061055d565b6101966105eb565b60035481565b60015460a060020a900460ff161561037457600154600160a060020a0316331461036f57600080fd5b61038b565b600254600160a060020a0316331461038b57600080fd5b805161039e90600090602084019061067e565b5050565b60008054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561042e5780601f106104035761010080835404028352916020019161042e565b820191906000526020600020905b81548152906001019060200180831161041157829003601f168201915b505050505090505b90565b600154600160a060020a031681565b600254600160a060020a031681565b600254600160a060020a0316331461046e57600080fd5b61047781610346565b50565b60045460ff1681565b6001805460a060020a900460ff1615151461049d57600080fd5b600154600160a060020a031633146104b457600080fd5b6001805474ff00000000000000000000000000000000000000001916908190556040805160a060020a90920460ff1615158252517f553f3eb729050123008baa152b3cffa07e56302c061c734b190914c5b57334779181900360200190a1565b600254600160a060020a0316331461052b57600080fd5b60045460ff161561053b57600080fd5b6003556004805460ff19166001179055565b60015460a060020a900460ff1681565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156105e35780601f106105b8576101008083540402835291602001916105e3565b820191906000526020600020905b8154815290600101906020018083116105c657829003601f168201915b505050505081565b60015460a060020a900460ff161561060257600080fd5b600154600160a060020a0316331461061957600080fd5b6001805474ff0000000000000000000000000000000000000000191660a060020a90811791829055604080519190920460ff161515815290517f553f3eb729050123008baa152b3cffa07e56302c061c734b190914c5b57334779181900360200190a1565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106106bf57805160ff19168380011785556106ec565b828001600101855582156106ec579182015b828111156106ec5782518255916020019190600101906106d1565b506106f89291506106fc565b5090565b61043691905b808211156106f8576000815560010161070256fea165627a7a72305820f8b5e1d10105e49247da950e3bd726f150ee7adbe55725e5a1eed400c12bc2fd0029`

// DeployFileStore deploys a new Ethereum contract, binding an instance of FileStore to it.
func DeployFileStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FileStore, error) {
	parsed, err := abi.JSON(strings.NewReader(FileStoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FileStoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FileStore{FileStoreCaller: FileStoreCaller{contract: contract}, FileStoreTransactor: FileStoreTransactor{contract: contract}, FileStoreFilterer: FileStoreFilterer{contract: contract}}, nil
}

// FileStore is an auto generated Go binding around an Ethereum contract.
type FileStore struct {
	FileStoreCaller     // Read-only binding to the contract
	FileStoreTransactor // Write-only binding to the contract
	FileStoreFilterer   // Log filterer for contract events
}

// FileStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type FileStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FileStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FileStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FileStoreSession struct {
	Contract     *FileStore        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FileStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FileStoreCallerSession struct {
	Contract *FileStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FileStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FileStoreTransactorSession struct {
	Contract     *FileStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FileStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type FileStoreRaw struct {
	Contract *FileStore // Generic contract binding to access the raw methods on
}

// FileStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FileStoreCallerRaw struct {
	Contract *FileStoreCaller // Generic read-only contract binding to access the raw methods on
}

// FileStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FileStoreTransactorRaw struct {
	Contract *FileStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFileStore creates a new instance of FileStore, bound to a specific deployed contract.
func NewFileStore(address common.Address, backend bind.ContractBackend) (*FileStore, error) {
	contract, err := bindFileStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FileStore{FileStoreCaller: FileStoreCaller{contract: contract}, FileStoreTransactor: FileStoreTransactor{contract: contract}, FileStoreFilterer: FileStoreFilterer{contract: contract}}, nil
}

// NewFileStoreCaller creates a new read-only instance of FileStore, bound to a specific deployed contract.
func NewFileStoreCaller(address common.Address, caller bind.ContractCaller) (*FileStoreCaller, error) {
	contract, err := bindFileStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FileStoreCaller{contract: contract}, nil
}

// NewFileStoreTransactor creates a new write-only instance of FileStore, bound to a specific deployed contract.
func NewFileStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*FileStoreTransactor, error) {
	contract, err := bindFileStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FileStoreTransactor{contract: contract}, nil
}

// NewFileStoreFilterer creates a new log filterer instance of FileStore, bound to a specific deployed contract.
func NewFileStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*FileStoreFilterer, error) {
	contract, err := bindFileStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FileStoreFilterer{contract: contract}, nil
}

// bindFileStore binds a generic wrapper to an already deployed contract.
func bindFileStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FileStoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileStore *FileStoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FileStore.Contract.FileStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileStore *FileStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileStore.Contract.FileStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileStore *FileStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FileStore.Contract.FileStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileStore *FileStoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FileStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileStore *FileStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileStore *FileStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FileStore.Contract.contract.Transact(opts, method, params...)
}

// CodeTx is a free data retrieval call binding the contract method 0x04fb3167.
//
// Solidity: function codeTx() constant returns(bytes32)
func (_FileStore *FileStoreCaller) CodeTx(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _FileStore.contract.Call(opts, out, "codeTx")
	return *ret0, err
}

// CodeTx is a free data retrieval call binding the contract method 0x04fb3167.
//
// Solidity: function codeTx() constant returns(bytes32)
func (_FileStore *FileStoreSession) CodeTx() ([32]byte, error) {
	return _FileStore.Contract.CodeTx(&_FileStore.CallOpts)
}

// CodeTx is a free data retrieval call binding the contract method 0x04fb3167.
//
// Solidity: function codeTx() constant returns(bytes32)
func (_FileStore *FileStoreCallerSession) CodeTx() ([32]byte, error) {
	return _FileStore.Contract.CodeTx(&_FileStore.CallOpts)
}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() constant returns(string)
func (_FileStore *FileStoreCaller) Core(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FileStore.contract.Call(opts, out, "core")
	return *ret0, err
}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() constant returns(string)
func (_FileStore *FileStoreSession) Core() (string, error) {
	return _FileStore.Contract.Core(&_FileStore.CallOpts)
}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() constant returns(string)
func (_FileStore *FileStoreCallerSession) Core() (string, error) {
	return _FileStore.Contract.Core(&_FileStore.CallOpts)
}

// HashTx is a free data retrieval call binding the contract method 0x9e7487eb.
//
// Solidity: function hashTx() constant returns(bool)
func (_FileStore *FileStoreCaller) HashTx(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FileStore.contract.Call(opts, out, "hashTx")
	return *ret0, err
}

// HashTx is a free data retrieval call binding the contract method 0x9e7487eb.
//
// Solidity: function hashTx() constant returns(bool)
func (_FileStore *FileStoreSession) HashTx() (bool, error) {
	return _FileStore.Contract.HashTx(&_FileStore.CallOpts)
}

// HashTx is a free data retrieval call binding the contract method 0x9e7487eb.
//
// Solidity: function hashTx() constant returns(bool)
func (_FileStore *FileStoreCallerSession) HashTx() (bool, error) {
	return _FileStore.Contract.HashTx(&_FileStore.CallOpts)
}

// IndexHash is a free data retrieval call binding the contract method 0x5cb8eca4.
//
// Solidity: function indexHash() constant returns(string)
func (_FileStore *FileStoreCaller) IndexHash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FileStore.contract.Call(opts, out, "indexHash")
	return *ret0, err
}

// IndexHash is a free data retrieval call binding the contract method 0x5cb8eca4.
//
// Solidity: function indexHash() constant returns(string)
func (_FileStore *FileStoreSession) IndexHash() (string, error) {
	return _FileStore.Contract.IndexHash(&_FileStore.CallOpts)
}

// IndexHash is a free data retrieval call binding the contract method 0x5cb8eca4.
//
// Solidity: function indexHash() constant returns(string)
func (_FileStore *FileStoreCallerSession) IndexHash() (string, error) {
	return _FileStore.Contract.IndexHash(&_FileStore.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_FileStore *FileStoreCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FileStore.contract.Call(opts, out, "locked")
	return *ret0, err
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_FileStore *FileStoreSession) Locked() (bool, error) {
	return _FileStore.Contract.Locked(&_FileStore.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_FileStore *FileStoreCallerSession) Locked() (bool, error) {
	return _FileStore.Contract.Locked(&_FileStore.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_FileStore *FileStoreCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FileStore.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_FileStore *FileStoreSession) Oracle() (common.Address, error) {
	return _FileStore.Contract.Oracle(&_FileStore.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_FileStore *FileStoreCallerSession) Oracle() (common.Address, error) {
	return _FileStore.Contract.Oracle(&_FileStore.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FileStore *FileStoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FileStore.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FileStore *FileStoreSession) Owner() (common.Address, error) {
	return _FileStore.Contract.Owner(&_FileStore.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_FileStore *FileStoreCallerSession) Owner() (common.Address, error) {
	return _FileStore.Contract.Owner(&_FileStore.CallOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_FileStore *FileStoreTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileStore.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_FileStore *FileStoreSession) Lock() (*types.Transaction, error) {
	return _FileStore.Contract.Lock(&_FileStore.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_FileStore *FileStoreTransactorSession) Lock() (*types.Transaction, error) {
	return _FileStore.Contract.Lock(&_FileStore.TransactOpts)
}

// SetCode is a paid mutator transaction binding the contract method 0xb9ef767f.
//
// Solidity: function setCode(bytes32 _codeTx) returns()
func (_FileStore *FileStoreTransactor) SetCode(opts *bind.TransactOpts, _codeTx [32]byte) (*types.Transaction, error) {
	return _FileStore.contract.Transact(opts, "setCode", _codeTx)
}

// SetCode is a paid mutator transaction binding the contract method 0xb9ef767f.
//
// Solidity: function setCode(bytes32 _codeTx) returns()
func (_FileStore *FileStoreSession) SetCode(_codeTx [32]byte) (*types.Transaction, error) {
	return _FileStore.Contract.SetCode(&_FileStore.TransactOpts, _codeTx)
}

// SetCode is a paid mutator transaction binding the contract method 0xb9ef767f.
//
// Solidity: function setCode(bytes32 _codeTx) returns()
func (_FileStore *FileStoreTransactorSession) SetCode(_codeTx [32]byte) (*types.Transaction, error) {
	return _FileStore.Contract.SetCode(&_FileStore.TransactOpts, _codeTx)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_FileStore *FileStoreTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileStore.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_FileStore *FileStoreSession) Unlock() (*types.Transaction, error) {
	return _FileStore.Contract.Unlock(&_FileStore.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_FileStore *FileStoreTransactorSession) Unlock() (*types.Transaction, error) {
	return _FileStore.Contract.Unlock(&_FileStore.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0x3d7403a3.
//
// Solidity: function update(string _data) returns()
func (_FileStore *FileStoreTransactor) Update(opts *bind.TransactOpts, _data string) (*types.Transaction, error) {
	return _FileStore.contract.Transact(opts, "update", _data)
}

// Update is a paid mutator transaction binding the contract method 0x3d7403a3.
//
// Solidity: function update(string _data) returns()
func (_FileStore *FileStoreSession) Update(_data string) (*types.Transaction, error) {
	return _FileStore.Contract.Update(&_FileStore.TransactOpts, _data)
}

// Update is a paid mutator transaction binding the contract method 0x3d7403a3.
//
// Solidity: function update(string _data) returns()
func (_FileStore *FileStoreTransactorSession) Update(_data string) (*types.Transaction, error) {
	return _FileStore.Contract.Update(&_FileStore.TransactOpts, _data)
}

// UpdateHash is a paid mutator transaction binding the contract method 0x92dea922.
//
// Solidity: function updateHash(string hash) returns()
func (_FileStore *FileStoreTransactor) UpdateHash(opts *bind.TransactOpts, hash string) (*types.Transaction, error) {
	return _FileStore.contract.Transact(opts, "updateHash", hash)
}

// UpdateHash is a paid mutator transaction binding the contract method 0x92dea922.
//
// Solidity: function updateHash(string hash) returns()
func (_FileStore *FileStoreSession) UpdateHash(hash string) (*types.Transaction, error) {
	return _FileStore.Contract.UpdateHash(&_FileStore.TransactOpts, hash)
}

// UpdateHash is a paid mutator transaction binding the contract method 0x92dea922.
//
// Solidity: function updateHash(string hash) returns()
func (_FileStore *FileStoreTransactorSession) UpdateHash(hash string) (*types.Transaction, error) {
	return _FileStore.Contract.UpdateHash(&_FileStore.TransactOpts, hash)
}

// FileStoreLockEventIterator is returned from FilterLockEvent and is used to iterate over the raw logs and unpacked data for LockEvent events raised by the FileStore contract.
type FileStoreLockEventIterator struct {
	Event *FileStoreLockEvent // Event containing the contract specifics and raw log

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
func (it *FileStoreLockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileStoreLockEvent)
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
		it.Event = new(FileStoreLockEvent)
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
func (it *FileStoreLockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileStoreLockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileStoreLockEvent represents a LockEvent event raised by the FileStore contract.
type FileStoreLockEvent struct {
	bool
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLockEvent is a free log retrieval operation binding the contract event 0x553f3eb729050123008baa152b3cffa07e56302c061c734b190914c5b5733477.
//
// Solidity: event LockEvent(bool )
func (_FileStore *FileStoreFilterer) FilterLockEvent(opts *bind.FilterOpts) (*FileStoreLockEventIterator, error) {

	logs, sub, err := _FileStore.contract.FilterLogs(opts, "LockEvent")
	if err != nil {
		return nil, err
	}
	return &FileStoreLockEventIterator{contract: _FileStore.contract, event: "LockEvent", logs: logs, sub: sub}, nil
}

// WatchLockEvent is a free log subscription operation binding the contract event 0x553f3eb729050123008baa152b3cffa07e56302c061c734b190914c5b5733477.
//
// Solidity: event LockEvent(bool )
func (_FileStore *FileStoreFilterer) WatchLockEvent(opts *bind.WatchOpts, sink chan<- *FileStoreLockEvent) (event.Subscription, error) {

	logs, sub, err := _FileStore.contract.WatchLogs(opts, "LockEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileStoreLockEvent)
				if err := _FileStore.contract.UnpackLog(event, "LockEvent", log); err != nil {
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

// MetaDataABI is the input ABI used to generate the binding from.
const MetaDataABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"codeTx\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashTx\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_codeTx\",\"type\":\"bytes32\"}],\"name\":\"setCode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"core\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"bool\"}],\"name\":\"LockEvent\",\"type\":\"event\"}]"

// MetaDataBin is the compiled bytecode used for deploying new contracts.
const MetaDataBin = `0x608060405234801561001057600080fd5b5060018054600160a860020a0319167396216849c49358b10257cb55b28ea603c874b05e17905560028054600160a060020a031916331790556004805460ff191690556105c7806100626000396000f3fe608060405234801561001057600080fd5b50600436106100bb576000357c010000000000000000000000000000000000000000000000000000000090048063a69df4b511610083578063a69df4b5146101ca578063b9ef767f146101d2578063cf309012146101ef578063f2f4eb26146101f7578063f83d08ba14610274576100bb565b806304fb3167146100c05780633d7403a3146100da5780637dc0d1d0146101825780638da5cb5b146101a65780639e7487eb146101ae575b600080fd5b6100c861027c565b60408051918252519081900360200190f35b610180600480360360208110156100f057600080fd5b81019060208101813564010000000081111561010b57600080fd5b82018360208201111561011d57600080fd5b8035906020019184600183028401116401000000008311171561013f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610282945050505050565b005b61018a6102de565b60408051600160a060020a039092168252519081900360200190f35b61018a6102ed565b6101b66102fc565b604080519115158252519081900360200190f35b610180610305565b610180600480360360208110156101e857600080fd5b5035610396565b6101b66103cf565b6101ff6103df565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610239578181015183820152602001610221565b50505050905090810190601f1680156102665780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61018061046d565b60035481565b60015460a060020a900460ff16156102b057600154600160a060020a031633146102ab57600080fd5b6102c7565b600254600160a060020a031633146102c757600080fd5b80516102da906000906020840190610500565b5050565b600154600160a060020a031681565b600254600160a060020a031681565b60045460ff1681565b6001805460a060020a900460ff1615151461031f57600080fd5b600154600160a060020a0316331461033657600080fd5b6001805474ff00000000000000000000000000000000000000001916908190556040805160a060020a90920460ff1615158252517f553f3eb729050123008baa152b3cffa07e56302c061c734b190914c5b57334779181900360200190a1565b600254600160a060020a031633146103ad57600080fd5b60045460ff16156103bd57600080fd5b6003556004805460ff19166001179055565b60015460a060020a900460ff1681565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104655780601f1061043a57610100808354040283529160200191610465565b820191906000526020600020905b81548152906001019060200180831161044857829003601f168201915b505050505081565b60015460a060020a900460ff161561048457600080fd5b600154600160a060020a0316331461049b57600080fd5b6001805474ff0000000000000000000000000000000000000000191660a060020a90811791829055604080519190920460ff161515815290517f553f3eb729050123008baa152b3cffa07e56302c061c734b190914c5b57334779181900360200190a1565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061054157805160ff191683800117855561056e565b8280016001018555821561056e579182015b8281111561056e578251825591602001919060010190610553565b5061057a92915061057e565b5090565b61059891905b8082111561057a5760008155600101610584565b9056fea165627a7a72305820ec7d7f44fd5bdfac468942ce792918c36ae685c85440c9ddf99c9b249a9586580029`

// DeployMetaData deploys a new Ethereum contract, binding an instance of MetaData to it.
func DeployMetaData(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MetaData, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MetaDataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MetaData{MetaDataCaller: MetaDataCaller{contract: contract}, MetaDataTransactor: MetaDataTransactor{contract: contract}, MetaDataFilterer: MetaDataFilterer{contract: contract}}, nil
}

// MetaData is an auto generated Go binding around an Ethereum contract.
type MetaData struct {
	MetaDataCaller     // Read-only binding to the contract
	MetaDataTransactor // Write-only binding to the contract
	MetaDataFilterer   // Log filterer for contract events
}

// MetaDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetaDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetaDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetaDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetaDataSession struct {
	Contract     *MetaData         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetaDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetaDataCallerSession struct {
	Contract *MetaDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MetaDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetaDataTransactorSession struct {
	Contract     *MetaDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MetaDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetaDataRaw struct {
	Contract *MetaData // Generic contract binding to access the raw methods on
}

// MetaDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetaDataCallerRaw struct {
	Contract *MetaDataCaller // Generic read-only contract binding to access the raw methods on
}

// MetaDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetaDataTransactorRaw struct {
	Contract *MetaDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetaData creates a new instance of MetaData, bound to a specific deployed contract.
func NewMetaData(address common.Address, backend bind.ContractBackend) (*MetaData, error) {
	contract, err := bindMetaData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetaData{MetaDataCaller: MetaDataCaller{contract: contract}, MetaDataTransactor: MetaDataTransactor{contract: contract}, MetaDataFilterer: MetaDataFilterer{contract: contract}}, nil
}

// NewMetaDataCaller creates a new read-only instance of MetaData, bound to a specific deployed contract.
func NewMetaDataCaller(address common.Address, caller bind.ContractCaller) (*MetaDataCaller, error) {
	contract, err := bindMetaData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetaDataCaller{contract: contract}, nil
}

// NewMetaDataTransactor creates a new write-only instance of MetaData, bound to a specific deployed contract.
func NewMetaDataTransactor(address common.Address, transactor bind.ContractTransactor) (*MetaDataTransactor, error) {
	contract, err := bindMetaData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetaDataTransactor{contract: contract}, nil
}

// NewMetaDataFilterer creates a new log filterer instance of MetaData, bound to a specific deployed contract.
func NewMetaDataFilterer(address common.Address, filterer bind.ContractFilterer) (*MetaDataFilterer, error) {
	contract, err := bindMetaData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetaDataFilterer{contract: contract}, nil
}

// bindMetaData binds a generic wrapper to an already deployed contract.
func bindMetaData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaData *MetaDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MetaData.Contract.MetaDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaData *MetaDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaData.Contract.MetaDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaData *MetaDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaData.Contract.MetaDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaData *MetaDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MetaData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaData *MetaDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaData *MetaDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaData.Contract.contract.Transact(opts, method, params...)
}

// CodeTx is a free data retrieval call binding the contract method 0x04fb3167.
//
// Solidity: function codeTx() constant returns(bytes32)
func (_MetaData *MetaDataCaller) CodeTx(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MetaData.contract.Call(opts, out, "codeTx")
	return *ret0, err
}

// CodeTx is a free data retrieval call binding the contract method 0x04fb3167.
//
// Solidity: function codeTx() constant returns(bytes32)
func (_MetaData *MetaDataSession) CodeTx() ([32]byte, error) {
	return _MetaData.Contract.CodeTx(&_MetaData.CallOpts)
}

// CodeTx is a free data retrieval call binding the contract method 0x04fb3167.
//
// Solidity: function codeTx() constant returns(bytes32)
func (_MetaData *MetaDataCallerSession) CodeTx() ([32]byte, error) {
	return _MetaData.Contract.CodeTx(&_MetaData.CallOpts)
}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() constant returns(string)
func (_MetaData *MetaDataCaller) Core(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MetaData.contract.Call(opts, out, "core")
	return *ret0, err
}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() constant returns(string)
func (_MetaData *MetaDataSession) Core() (string, error) {
	return _MetaData.Contract.Core(&_MetaData.CallOpts)
}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() constant returns(string)
func (_MetaData *MetaDataCallerSession) Core() (string, error) {
	return _MetaData.Contract.Core(&_MetaData.CallOpts)
}

// HashTx is a free data retrieval call binding the contract method 0x9e7487eb.
//
// Solidity: function hashTx() constant returns(bool)
func (_MetaData *MetaDataCaller) HashTx(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MetaData.contract.Call(opts, out, "hashTx")
	return *ret0, err
}

// HashTx is a free data retrieval call binding the contract method 0x9e7487eb.
//
// Solidity: function hashTx() constant returns(bool)
func (_MetaData *MetaDataSession) HashTx() (bool, error) {
	return _MetaData.Contract.HashTx(&_MetaData.CallOpts)
}

// HashTx is a free data retrieval call binding the contract method 0x9e7487eb.
//
// Solidity: function hashTx() constant returns(bool)
func (_MetaData *MetaDataCallerSession) HashTx() (bool, error) {
	return _MetaData.Contract.HashTx(&_MetaData.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_MetaData *MetaDataCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MetaData.contract.Call(opts, out, "locked")
	return *ret0, err
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_MetaData *MetaDataSession) Locked() (bool, error) {
	return _MetaData.Contract.Locked(&_MetaData.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() constant returns(bool)
func (_MetaData *MetaDataCallerSession) Locked() (bool, error) {
	return _MetaData.Contract.Locked(&_MetaData.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_MetaData *MetaDataCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MetaData.contract.Call(opts, out, "oracle")
	return *ret0, err
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_MetaData *MetaDataSession) Oracle() (common.Address, error) {
	return _MetaData.Contract.Oracle(&_MetaData.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() constant returns(address)
func (_MetaData *MetaDataCallerSession) Oracle() (common.Address, error) {
	return _MetaData.Contract.Oracle(&_MetaData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MetaData *MetaDataCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MetaData.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MetaData *MetaDataSession) Owner() (common.Address, error) {
	return _MetaData.Contract.Owner(&_MetaData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MetaData *MetaDataCallerSession) Owner() (common.Address, error) {
	return _MetaData.Contract.Owner(&_MetaData.CallOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_MetaData *MetaDataTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaData.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_MetaData *MetaDataSession) Lock() (*types.Transaction, error) {
	return _MetaData.Contract.Lock(&_MetaData.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_MetaData *MetaDataTransactorSession) Lock() (*types.Transaction, error) {
	return _MetaData.Contract.Lock(&_MetaData.TransactOpts)
}

// SetCode is a paid mutator transaction binding the contract method 0xb9ef767f.
//
// Solidity: function setCode(bytes32 _codeTx) returns()
func (_MetaData *MetaDataTransactor) SetCode(opts *bind.TransactOpts, _codeTx [32]byte) (*types.Transaction, error) {
	return _MetaData.contract.Transact(opts, "setCode", _codeTx)
}

// SetCode is a paid mutator transaction binding the contract method 0xb9ef767f.
//
// Solidity: function setCode(bytes32 _codeTx) returns()
func (_MetaData *MetaDataSession) SetCode(_codeTx [32]byte) (*types.Transaction, error) {
	return _MetaData.Contract.SetCode(&_MetaData.TransactOpts, _codeTx)
}

// SetCode is a paid mutator transaction binding the contract method 0xb9ef767f.
//
// Solidity: function setCode(bytes32 _codeTx) returns()
func (_MetaData *MetaDataTransactorSession) SetCode(_codeTx [32]byte) (*types.Transaction, error) {
	return _MetaData.Contract.SetCode(&_MetaData.TransactOpts, _codeTx)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_MetaData *MetaDataTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaData.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_MetaData *MetaDataSession) Unlock() (*types.Transaction, error) {
	return _MetaData.Contract.Unlock(&_MetaData.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_MetaData *MetaDataTransactorSession) Unlock() (*types.Transaction, error) {
	return _MetaData.Contract.Unlock(&_MetaData.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0x3d7403a3.
//
// Solidity: function update(string _data) returns()
func (_MetaData *MetaDataTransactor) Update(opts *bind.TransactOpts, _data string) (*types.Transaction, error) {
	return _MetaData.contract.Transact(opts, "update", _data)
}

// Update is a paid mutator transaction binding the contract method 0x3d7403a3.
//
// Solidity: function update(string _data) returns()
func (_MetaData *MetaDataSession) Update(_data string) (*types.Transaction, error) {
	return _MetaData.Contract.Update(&_MetaData.TransactOpts, _data)
}

// Update is a paid mutator transaction binding the contract method 0x3d7403a3.
//
// Solidity: function update(string _data) returns()
func (_MetaData *MetaDataTransactorSession) Update(_data string) (*types.Transaction, error) {
	return _MetaData.Contract.Update(&_MetaData.TransactOpts, _data)
}

// MetaDataLockEventIterator is returned from FilterLockEvent and is used to iterate over the raw logs and unpacked data for LockEvent events raised by the MetaData contract.
type MetaDataLockEventIterator struct {
	Event *MetaDataLockEvent // Event containing the contract specifics and raw log

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
func (it *MetaDataLockEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaDataLockEvent)
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
		it.Event = new(MetaDataLockEvent)
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
func (it *MetaDataLockEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaDataLockEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaDataLockEvent represents a LockEvent event raised by the MetaData contract.
type MetaDataLockEvent struct {
	bool
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLockEvent is a free log retrieval operation binding the contract event 0x553f3eb729050123008baa152b3cffa07e56302c061c734b190914c5b5733477.
//
// Solidity: event LockEvent(bool )
func (_MetaData *MetaDataFilterer) FilterLockEvent(opts *bind.FilterOpts) (*MetaDataLockEventIterator, error) {

	logs, sub, err := _MetaData.contract.FilterLogs(opts, "LockEvent")
	if err != nil {
		return nil, err
	}
	return &MetaDataLockEventIterator{contract: _MetaData.contract, event: "LockEvent", logs: logs, sub: sub}, nil
}

// WatchLockEvent is a free log subscription operation binding the contract event 0x553f3eb729050123008baa152b3cffa07e56302c061c734b190914c5b5733477.
//
// Solidity: event LockEvent(bool )
func (_MetaData *MetaDataFilterer) WatchLockEvent(opts *bind.WatchOpts, sink chan<- *MetaDataLockEvent) (event.Subscription, error) {

	logs, sub, err := _MetaData.contract.WatchLogs(opts, "LockEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaDataLockEvent)
				if err := _MetaData.contract.UnpackLog(event, "LockEvent", log); err != nil {
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
