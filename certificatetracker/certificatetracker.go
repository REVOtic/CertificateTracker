// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package certificatetracker

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

// CertificateTrackerABI is the input ABI used to generate the binding from.
const CertificateTrackerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"certificate\",\"type\":\"string\"}],\"name\":\"CertificateIssued\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_issuerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ownerAddress\",\"type\":\"address\"}],\"name\":\"getCertificateHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ownerAddress\",\"type\":\"address\"}],\"name\":\"getIssuedCertificateDetails\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_certificate\",\"type\":\"string\"}],\"name\":\"getOwnerDetailsByCertificateHash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ownerAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_certificate\",\"type\":\"string\"}],\"name\":\"setCertificateDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CertificateTracker is an auto generated Go binding around an Ethereum contract.
type CertificateTracker struct {
	CertificateTrackerCaller     // Read-only binding to the contract
	CertificateTrackerTransactor // Write-only binding to the contract
	CertificateTrackerFilterer   // Log filterer for contract events
}

// CertificateTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CertificateTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CertificateTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CertificateTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CertificateTrackerSession struct {
	Contract     *CertificateTracker // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CertificateTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CertificateTrackerCallerSession struct {
	Contract *CertificateTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// CertificateTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CertificateTrackerTransactorSession struct {
	Contract     *CertificateTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// CertificateTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CertificateTrackerRaw struct {
	Contract *CertificateTracker // Generic contract binding to access the raw methods on
}

// CertificateTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CertificateTrackerCallerRaw struct {
	Contract *CertificateTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// CertificateTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CertificateTrackerTransactorRaw struct {
	Contract *CertificateTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCertificateTracker creates a new instance of CertificateTracker, bound to a specific deployed contract.
func NewCertificateTracker(address common.Address, backend bind.ContractBackend) (*CertificateTracker, error) {
	contract, err := bindCertificateTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CertificateTracker{CertificateTrackerCaller: CertificateTrackerCaller{contract: contract}, CertificateTrackerTransactor: CertificateTrackerTransactor{contract: contract}, CertificateTrackerFilterer: CertificateTrackerFilterer{contract: contract}}, nil
}

// NewCertificateTrackerCaller creates a new read-only instance of CertificateTracker, bound to a specific deployed contract.
func NewCertificateTrackerCaller(address common.Address, caller bind.ContractCaller) (*CertificateTrackerCaller, error) {
	contract, err := bindCertificateTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateTrackerCaller{contract: contract}, nil
}

// NewCertificateTrackerTransactor creates a new write-only instance of CertificateTracker, bound to a specific deployed contract.
func NewCertificateTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*CertificateTrackerTransactor, error) {
	contract, err := bindCertificateTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateTrackerTransactor{contract: contract}, nil
}

// NewCertificateTrackerFilterer creates a new log filterer instance of CertificateTracker, bound to a specific deployed contract.
func NewCertificateTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*CertificateTrackerFilterer, error) {
	contract, err := bindCertificateTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertificateTrackerFilterer{contract: contract}, nil
}

// bindCertificateTracker binds a generic wrapper to an already deployed contract.
func bindCertificateTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CertificateTrackerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertificateTracker *CertificateTrackerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CertificateTracker.Contract.CertificateTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertificateTracker *CertificateTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertificateTracker.Contract.CertificateTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertificateTracker *CertificateTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CertificateTracker.Contract.CertificateTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertificateTracker *CertificateTrackerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CertificateTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertificateTracker *CertificateTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertificateTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertificateTracker *CertificateTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CertificateTracker.Contract.contract.Transact(opts, method, params...)
}

// GetCertificateHash is a free data retrieval call binding the contract method 0x37f7eb7e.
//
// Solidity: function getCertificateHash(address _issuerAddress, address _ownerAddress) constant returns(string)
func (_CertificateTracker *CertificateTrackerCaller) GetCertificateHash(opts *bind.CallOpts, _issuerAddress common.Address, _ownerAddress common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CertificateTracker.contract.Call(opts, out, "getCertificateHash", _issuerAddress, _ownerAddress)
	return *ret0, err
}

// GetCertificateHash is a free data retrieval call binding the contract method 0x37f7eb7e.
//
// Solidity: function getCertificateHash(address _issuerAddress, address _ownerAddress) constant returns(string)
func (_CertificateTracker *CertificateTrackerSession) GetCertificateHash(_issuerAddress common.Address, _ownerAddress common.Address) (string, error) {
	return _CertificateTracker.Contract.GetCertificateHash(&_CertificateTracker.CallOpts, _issuerAddress, _ownerAddress)
}

// GetCertificateHash is a free data retrieval call binding the contract method 0x37f7eb7e.
//
// Solidity: function getCertificateHash(address _issuerAddress, address _ownerAddress) constant returns(string)
func (_CertificateTracker *CertificateTrackerCallerSession) GetCertificateHash(_issuerAddress common.Address, _ownerAddress common.Address) (string, error) {
	return _CertificateTracker.Contract.GetCertificateHash(&_CertificateTracker.CallOpts, _issuerAddress, _ownerAddress)
}

// GetIssuedCertificateDetails is a free data retrieval call binding the contract method 0x748a43da.
//
// Solidity: function getIssuedCertificateDetails(address _ownerAddress) constant returns(address, string, string)
func (_CertificateTracker *CertificateTrackerCaller) GetIssuedCertificateDetails(opts *bind.CallOpts, _ownerAddress common.Address) (common.Address, string, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _CertificateTracker.contract.Call(opts, out, "getIssuedCertificateDetails", _ownerAddress)
	return *ret0, *ret1, *ret2, err
}

// GetIssuedCertificateDetails is a free data retrieval call binding the contract method 0x748a43da.
//
// Solidity: function getIssuedCertificateDetails(address _ownerAddress) constant returns(address, string, string)
func (_CertificateTracker *CertificateTrackerSession) GetIssuedCertificateDetails(_ownerAddress common.Address) (common.Address, string, string, error) {
	return _CertificateTracker.Contract.GetIssuedCertificateDetails(&_CertificateTracker.CallOpts, _ownerAddress)
}

// GetIssuedCertificateDetails is a free data retrieval call binding the contract method 0x748a43da.
//
// Solidity: function getIssuedCertificateDetails(address _ownerAddress) constant returns(address, string, string)
func (_CertificateTracker *CertificateTrackerCallerSession) GetIssuedCertificateDetails(_ownerAddress common.Address) (common.Address, string, string, error) {
	return _CertificateTracker.Contract.GetIssuedCertificateDetails(&_CertificateTracker.CallOpts, _ownerAddress)
}

// GetOwnerDetailsByCertificateHash is a free data retrieval call binding the contract method 0xca3bacbf.
//
// Solidity: function getOwnerDetailsByCertificateHash(string _certificate) constant returns(address, string, string)
func (_CertificateTracker *CertificateTrackerCaller) GetOwnerDetailsByCertificateHash(opts *bind.CallOpts, _certificate string) (common.Address, string, string, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(string)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _CertificateTracker.contract.Call(opts, out, "getOwnerDetailsByCertificateHash", _certificate)
	return *ret0, *ret1, *ret2, err
}

// GetOwnerDetailsByCertificateHash is a free data retrieval call binding the contract method 0xca3bacbf.
//
// Solidity: function getOwnerDetailsByCertificateHash(string _certificate) constant returns(address, string, string)
func (_CertificateTracker *CertificateTrackerSession) GetOwnerDetailsByCertificateHash(_certificate string) (common.Address, string, string, error) {
	return _CertificateTracker.Contract.GetOwnerDetailsByCertificateHash(&_CertificateTracker.CallOpts, _certificate)
}

// GetOwnerDetailsByCertificateHash is a free data retrieval call binding the contract method 0xca3bacbf.
//
// Solidity: function getOwnerDetailsByCertificateHash(string _certificate) constant returns(address, string, string)
func (_CertificateTracker *CertificateTrackerCallerSession) GetOwnerDetailsByCertificateHash(_certificate string) (common.Address, string, string, error) {
	return _CertificateTracker.Contract.GetOwnerDetailsByCertificateHash(&_CertificateTracker.CallOpts, _certificate)
}

// SetCertificateDetails is a paid mutator transaction binding the contract method 0xd0aec93e.
//
// Solidity: function setCertificateDetails(address _ownerAddress, string _name, string _email, string _certificate) returns()
func (_CertificateTracker *CertificateTrackerTransactor) SetCertificateDetails(opts *bind.TransactOpts, _ownerAddress common.Address, _name string, _email string, _certificate string) (*types.Transaction, error) {
	return _CertificateTracker.contract.Transact(opts, "setCertificateDetails", _ownerAddress, _name, _email, _certificate)
}

// SetCertificateDetails is a paid mutator transaction binding the contract method 0xd0aec93e.
//
// Solidity: function setCertificateDetails(address _ownerAddress, string _name, string _email, string _certificate) returns()
func (_CertificateTracker *CertificateTrackerSession) SetCertificateDetails(_ownerAddress common.Address, _name string, _email string, _certificate string) (*types.Transaction, error) {
	return _CertificateTracker.Contract.SetCertificateDetails(&_CertificateTracker.TransactOpts, _ownerAddress, _name, _email, _certificate)
}

// SetCertificateDetails is a paid mutator transaction binding the contract method 0xd0aec93e.
//
// Solidity: function setCertificateDetails(address _ownerAddress, string _name, string _email, string _certificate) returns()
func (_CertificateTracker *CertificateTrackerTransactorSession) SetCertificateDetails(_ownerAddress common.Address, _name string, _email string, _certificate string) (*types.Transaction, error) {
	return _CertificateTracker.Contract.SetCertificateDetails(&_CertificateTracker.TransactOpts, _ownerAddress, _name, _email, _certificate)
}

// CertificateTrackerCertificateIssuedIterator is returned from FilterCertificateIssued and is used to iterate over the raw logs and unpacked data for CertificateIssued events raised by the CertificateTracker contract.
type CertificateTrackerCertificateIssuedIterator struct {
	Event *CertificateTrackerCertificateIssued // Event containing the contract specifics and raw log

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
func (it *CertificateTrackerCertificateIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificateTrackerCertificateIssued)
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
		it.Event = new(CertificateTrackerCertificateIssued)
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
func (it *CertificateTrackerCertificateIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificateTrackerCertificateIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificateTrackerCertificateIssued represents a CertificateIssued event raised by the CertificateTracker contract.
type CertificateTrackerCertificateIssued struct {
	IssuerAddress common.Address
	OwnerAddress  common.Address
	Name          string
	Email         string
	Certificate   string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCertificateIssued is a free log retrieval operation binding the contract event 0xe3c11c61a04e90f454e77e350400fc783adfd085048dce972fb06773a489642e.
//
// Solidity: event CertificateIssued(address issuerAddress, address ownerAddress, string name, string email, string certificate)
func (_CertificateTracker *CertificateTrackerFilterer) FilterCertificateIssued(opts *bind.FilterOpts) (*CertificateTrackerCertificateIssuedIterator, error) {

	logs, sub, err := _CertificateTracker.contract.FilterLogs(opts, "CertificateIssued")
	if err != nil {
		return nil, err
	}
	return &CertificateTrackerCertificateIssuedIterator{contract: _CertificateTracker.contract, event: "CertificateIssued", logs: logs, sub: sub}, nil
}

// WatchCertificateIssued is a free log subscription operation binding the contract event 0xe3c11c61a04e90f454e77e350400fc783adfd085048dce972fb06773a489642e.
//
// Solidity: event CertificateIssued(address issuerAddress, address ownerAddress, string name, string email, string certificate)
func (_CertificateTracker *CertificateTrackerFilterer) WatchCertificateIssued(opts *bind.WatchOpts, sink chan<- *CertificateTrackerCertificateIssued) (event.Subscription, error) {

	logs, sub, err := _CertificateTracker.contract.WatchLogs(opts, "CertificateIssued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificateTrackerCertificateIssued)
				if err := _CertificateTracker.contract.UnpackLog(event, "CertificateIssued", log); err != nil {
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

// ParseCertificateIssued is a log parse operation binding the contract event 0xe3c11c61a04e90f454e77e350400fc783adfd085048dce972fb06773a489642e.
//
// Solidity: event CertificateIssued(address issuerAddress, address ownerAddress, string name, string email, string certificate)
func (_CertificateTracker *CertificateTrackerFilterer) ParseCertificateIssued(log types.Log) (*CertificateTrackerCertificateIssued, error) {
	event := new(CertificateTrackerCertificateIssued)
	if err := _CertificateTracker.contract.UnpackLog(event, "CertificateIssued", log); err != nil {
		return nil, err
	}
	return event, nil
}
