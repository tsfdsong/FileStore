// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accountmanager

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

// AccountManagerABI is the input ABI used to generate the binding from.
const AccountManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"clear\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"accounts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"store\",\"type\":\"address\"}],\"name\":\"Record\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"DeleteRecord\",\"type\":\"event\"}]"

// AccountManagerBin is the compiled bytecode used for deploying new contracts.
const AccountManagerBin = `0x608060405234801561001057600080fd5b5060018054600160a060020a03191633179055610368806100326000396000f3fe608060405234801561001057600080fd5b5060043610610073577c010000000000000000000000000000000000000000000000000000000060003504632801617e81146100785780633d0a4061146100a05780635e5c06e2146100c65780638da5cb5b14610108578063c2bc2efc14610110575b600080fd5b61009e6004803603602081101561008e57600080fd5b5035600160a060020a0316610136565b005b61009e600480360360208110156100b657600080fd5b5035600160a060020a031661021f565b6100ec600480360360208110156100dc57600080fd5b5035600160a060020a03166102f4565b60408051600160a060020a039092168252519081900360200190f35b6100ec61030f565b6100ec6004803603602081101561012657600080fd5b5035600160a060020a031661031e565b33600090815260208190526040902054600160a060020a0316156101bb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f6163636f756e742073746f726520636f6e747261637420686173206578697374604482015290519081900360640190fd5b33600081815260208190526040808220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03861690811790915590519092917fb9fdd8bcd1f646683410d68c8129bde05a5fd7e30f5303f6e178dad58e97f6ba91a350565b600154600160a060020a0316331461029857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f796f7520617265206e6f7420616c6c6f77656420746f2064656c657465206974604482015290519081900360640190fd5b600160a060020a038116600081815260208190526040808220805473ffffffffffffffffffffffffffffffffffffffff19169055517ff861b8be4d57f617447f6a69001bf60f302658eeedec630e4c71e5a612ec2b799190a250565b600060208190529081526040902054600160a060020a031681565b600154600160a060020a031681565b600160a060020a03908116600090815260208190526040902054169056fea165627a7a72305820a45faff56e10ca737782591efceef5da8dbd87e31b0edff7d4ec24d054724fb80029`

// DeployAccountManager deploys a new Ethereum contract, binding an instance of AccountManager to it.
func DeployAccountManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccountManager, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccountManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccountManager{AccountManagerCaller: AccountManagerCaller{contract: contract}, AccountManagerTransactor: AccountManagerTransactor{contract: contract}, AccountManagerFilterer: AccountManagerFilterer{contract: contract}}, nil
}

// AccountManager is an auto generated Go binding around an Ethereum contract.
type AccountManager struct {
	AccountManagerCaller     // Read-only binding to the contract
	AccountManagerTransactor // Write-only binding to the contract
	AccountManagerFilterer   // Log filterer for contract events
}

// AccountManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountManagerSession struct {
	Contract     *AccountManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountManagerCallerSession struct {
	Contract *AccountManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AccountManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountManagerTransactorSession struct {
	Contract     *AccountManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AccountManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountManagerRaw struct {
	Contract *AccountManager // Generic contract binding to access the raw methods on
}

// AccountManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountManagerCallerRaw struct {
	Contract *AccountManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AccountManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountManagerTransactorRaw struct {
	Contract *AccountManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountManager creates a new instance of AccountManager, bound to a specific deployed contract.
func NewAccountManager(address common.Address, backend bind.ContractBackend) (*AccountManager, error) {
	contract, err := bindAccountManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountManager{AccountManagerCaller: AccountManagerCaller{contract: contract}, AccountManagerTransactor: AccountManagerTransactor{contract: contract}, AccountManagerFilterer: AccountManagerFilterer{contract: contract}}, nil
}

// NewAccountManagerCaller creates a new read-only instance of AccountManager, bound to a specific deployed contract.
func NewAccountManagerCaller(address common.Address, caller bind.ContractCaller) (*AccountManagerCaller, error) {
	contract, err := bindAccountManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerCaller{contract: contract}, nil
}

// NewAccountManagerTransactor creates a new write-only instance of AccountManager, bound to a specific deployed contract.
func NewAccountManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountManagerTransactor, error) {
	contract, err := bindAccountManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountManagerTransactor{contract: contract}, nil
}

// NewAccountManagerFilterer creates a new log filterer instance of AccountManager, bound to a specific deployed contract.
func NewAccountManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountManagerFilterer, error) {
	contract, err := bindAccountManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountManagerFilterer{contract: contract}, nil
}

// bindAccountManager binds a generic wrapper to an already deployed contract.
func bindAccountManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManager *AccountManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountManager.Contract.AccountManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManager *AccountManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManager.Contract.AccountManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManager *AccountManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountManager.Contract.AccountManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountManager *AccountManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountManager *AccountManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountManager *AccountManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountManager.Contract.contract.Transact(opts, method, params...)
}

// Accounts is a free data retrieval call binding the contract method 0x5e5c06e2.
//
// Solidity: function accounts(address ) constant returns(address)
func (_AccountManager *AccountManagerCaller) Accounts(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccountManager.contract.Call(opts, out, "accounts", arg0)
	return *ret0, err
}

// Accounts is a free data retrieval call binding the contract method 0x5e5c06e2.
//
// Solidity: function accounts(address ) constant returns(address)
func (_AccountManager *AccountManagerSession) Accounts(arg0 common.Address) (common.Address, error) {
	return _AccountManager.Contract.Accounts(&_AccountManager.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0x5e5c06e2.
//
// Solidity: function accounts(address ) constant returns(address)
func (_AccountManager *AccountManagerCallerSession) Accounts(arg0 common.Address) (common.Address, error) {
	return _AccountManager.Contract.Accounts(&_AccountManager.CallOpts, arg0)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address account) constant returns(address)
func (_AccountManager *AccountManagerCaller) Get(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccountManager.contract.Call(opts, out, "get", account)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address account) constant returns(address)
func (_AccountManager *AccountManagerSession) Get(account common.Address) (common.Address, error) {
	return _AccountManager.Contract.Get(&_AccountManager.CallOpts, account)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address account) constant returns(address)
func (_AccountManager *AccountManagerCallerSession) Get(account common.Address) (common.Address, error) {
	return _AccountManager.Contract.Get(&_AccountManager.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AccountManager *AccountManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccountManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AccountManager *AccountManagerSession) Owner() (common.Address, error) {
	return _AccountManager.Contract.Owner(&_AccountManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AccountManager *AccountManagerCallerSession) Owner() (common.Address, error) {
	return _AccountManager.Contract.Owner(&_AccountManager.CallOpts)
}

// Clear is a paid mutator transaction binding the contract method 0x3d0a4061.
//
// Solidity: function clear(address account) returns()
func (_AccountManager *AccountManagerTransactor) Clear(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "clear", account)
}

// Clear is a paid mutator transaction binding the contract method 0x3d0a4061.
//
// Solidity: function clear(address account) returns()
func (_AccountManager *AccountManagerSession) Clear(account common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.Clear(&_AccountManager.TransactOpts, account)
}

// Clear is a paid mutator transaction binding the contract method 0x3d0a4061.
//
// Solidity: function clear(address account) returns()
func (_AccountManager *AccountManagerTransactorSession) Clear(account common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.Clear(&_AccountManager.TransactOpts, account)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address contractAddr) returns()
func (_AccountManager *AccountManagerTransactor) Set(opts *bind.TransactOpts, contractAddr common.Address) (*types.Transaction, error) {
	return _AccountManager.contract.Transact(opts, "set", contractAddr)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address contractAddr) returns()
func (_AccountManager *AccountManagerSession) Set(contractAddr common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.Set(&_AccountManager.TransactOpts, contractAddr)
}

// Set is a paid mutator transaction binding the contract method 0x2801617e.
//
// Solidity: function set(address contractAddr) returns()
func (_AccountManager *AccountManagerTransactorSession) Set(contractAddr common.Address) (*types.Transaction, error) {
	return _AccountManager.Contract.Set(&_AccountManager.TransactOpts, contractAddr)
}

// AccountManagerDeleteRecordIterator is returned from FilterDeleteRecord and is used to iterate over the raw logs and unpacked data for DeleteRecord events raised by the AccountManager contract.
type AccountManagerDeleteRecordIterator struct {
	Event *AccountManagerDeleteRecord // Event containing the contract specifics and raw log

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
func (it *AccountManagerDeleteRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerDeleteRecord)
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
		it.Event = new(AccountManagerDeleteRecord)
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
func (it *AccountManagerDeleteRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerDeleteRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerDeleteRecord represents a DeleteRecord event raised by the AccountManager contract.
type AccountManagerDeleteRecord struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeleteRecord is a free log retrieval operation binding the contract event 0xf861b8be4d57f617447f6a69001bf60f302658eeedec630e4c71e5a612ec2b79.
//
// Solidity: event DeleteRecord(address indexed account)
func (_AccountManager *AccountManagerFilterer) FilterDeleteRecord(opts *bind.FilterOpts, account []common.Address) (*AccountManagerDeleteRecordIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "DeleteRecord", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountManagerDeleteRecordIterator{contract: _AccountManager.contract, event: "DeleteRecord", logs: logs, sub: sub}, nil
}

// WatchDeleteRecord is a free log subscription operation binding the contract event 0xf861b8be4d57f617447f6a69001bf60f302658eeedec630e4c71e5a612ec2b79.
//
// Solidity: event DeleteRecord(address indexed account)
func (_AccountManager *AccountManagerFilterer) WatchDeleteRecord(opts *bind.WatchOpts, sink chan<- *AccountManagerDeleteRecord, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "DeleteRecord", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerDeleteRecord)
				if err := _AccountManager.contract.UnpackLog(event, "DeleteRecord", log); err != nil {
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

// AccountManagerRecordIterator is returned from FilterRecord and is used to iterate over the raw logs and unpacked data for Record events raised by the AccountManager contract.
type AccountManagerRecordIterator struct {
	Event *AccountManagerRecord // Event containing the contract specifics and raw log

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
func (it *AccountManagerRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountManagerRecord)
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
		it.Event = new(AccountManagerRecord)
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
func (it *AccountManagerRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountManagerRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountManagerRecord represents a Record event raised by the AccountManager contract.
type AccountManagerRecord struct {
	Account common.Address
	Store   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRecord is a free log retrieval operation binding the contract event 0xb9fdd8bcd1f646683410d68c8129bde05a5fd7e30f5303f6e178dad58e97f6ba.
//
// Solidity: event Record(address indexed account, address indexed store)
func (_AccountManager *AccountManagerFilterer) FilterRecord(opts *bind.FilterOpts, account []common.Address, store []common.Address) (*AccountManagerRecordIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var storeRule []interface{}
	for _, storeItem := range store {
		storeRule = append(storeRule, storeItem)
	}

	logs, sub, err := _AccountManager.contract.FilterLogs(opts, "Record", accountRule, storeRule)
	if err != nil {
		return nil, err
	}
	return &AccountManagerRecordIterator{contract: _AccountManager.contract, event: "Record", logs: logs, sub: sub}, nil
}

// WatchRecord is a free log subscription operation binding the contract event 0xb9fdd8bcd1f646683410d68c8129bde05a5fd7e30f5303f6e178dad58e97f6ba.
//
// Solidity: event Record(address indexed account, address indexed store)
func (_AccountManager *AccountManagerFilterer) WatchRecord(opts *bind.WatchOpts, sink chan<- *AccountManagerRecord, account []common.Address, store []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var storeRule []interface{}
	for _, storeItem := range store {
		storeRule = append(storeRule, storeItem)
	}

	logs, sub, err := _AccountManager.contract.WatchLogs(opts, "Record", accountRule, storeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountManagerRecord)
				if err := _AccountManager.contract.UnpackLog(event, "Record", log); err != nil {
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
