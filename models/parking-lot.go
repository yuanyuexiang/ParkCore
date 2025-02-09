// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package models

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

// ParkingLotParkingSpot is an auto generated low-level Go binding around an user-defined struct.
type ParkingLotParkingSpot struct {
	Id          *big.Int
	Location    string
	IsRented    bool
	Renter      common.Address
	RentEndTime *big.Int
	RentPrice   *big.Int
	Latitude    *big.Int
	Longitude   *big.Int
}

// ParkingLotMetaData contains all meta data concerning the ParkingLot contract.
var ParkingLotMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"checkRentalStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllParkingSpots\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isRented\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"renter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rentEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rentPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"latitude\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"longitude\",\"type\":\"int256\"}],\"internalType\":\"structParkingLot.ParkingSpot[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rentPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"latitude\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"longitude\",\"type\":\"int256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"parkingSpots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isRented\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"renter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rentEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rentPrice\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"latitude\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"longitude\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"rent\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"revokeParkingSpot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"terminateRental\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ParkingLotABI is the input ABI used to generate the binding from.
// Deprecated: Use ParkingLotMetaData.ABI instead.
var ParkingLotABI = ParkingLotMetaData.ABI

// ParkingLot is an auto generated Go binding around an Ethereum contract.
type ParkingLot struct {
	ParkingLotCaller     // Read-only binding to the contract
	ParkingLotTransactor // Write-only binding to the contract
	ParkingLotFilterer   // Log filterer for contract events
}

// ParkingLotCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParkingLotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParkingLotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParkingLotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParkingLotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParkingLotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParkingLotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParkingLotSession struct {
	Contract     *ParkingLot       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParkingLotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParkingLotCallerSession struct {
	Contract *ParkingLotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ParkingLotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParkingLotTransactorSession struct {
	Contract     *ParkingLotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ParkingLotRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParkingLotRaw struct {
	Contract *ParkingLot // Generic contract binding to access the raw methods on
}

// ParkingLotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParkingLotCallerRaw struct {
	Contract *ParkingLotCaller // Generic read-only contract binding to access the raw methods on
}

// ParkingLotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParkingLotTransactorRaw struct {
	Contract *ParkingLotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParkingLot creates a new instance of ParkingLot, bound to a specific deployed contract.
func NewParkingLot(address common.Address, backend bind.ContractBackend) (*ParkingLot, error) {
	contract, err := bindParkingLot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ParkingLot{ParkingLotCaller: ParkingLotCaller{contract: contract}, ParkingLotTransactor: ParkingLotTransactor{contract: contract}, ParkingLotFilterer: ParkingLotFilterer{contract: contract}}, nil
}

// NewParkingLotCaller creates a new read-only instance of ParkingLot, bound to a specific deployed contract.
func NewParkingLotCaller(address common.Address, caller bind.ContractCaller) (*ParkingLotCaller, error) {
	contract, err := bindParkingLot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParkingLotCaller{contract: contract}, nil
}

// NewParkingLotTransactor creates a new write-only instance of ParkingLot, bound to a specific deployed contract.
func NewParkingLotTransactor(address common.Address, transactor bind.ContractTransactor) (*ParkingLotTransactor, error) {
	contract, err := bindParkingLot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParkingLotTransactor{contract: contract}, nil
}

// NewParkingLotFilterer creates a new log filterer instance of ParkingLot, bound to a specific deployed contract.
func NewParkingLotFilterer(address common.Address, filterer bind.ContractFilterer) (*ParkingLotFilterer, error) {
	contract, err := bindParkingLot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParkingLotFilterer{contract: contract}, nil
}

// bindParkingLot binds a generic wrapper to an already deployed contract.
func bindParkingLot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParkingLotMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParkingLot *ParkingLotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ParkingLot.Contract.ParkingLotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParkingLot *ParkingLotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParkingLot.Contract.ParkingLotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParkingLot *ParkingLotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParkingLot.Contract.ParkingLotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParkingLot *ParkingLotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ParkingLot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParkingLot *ParkingLotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParkingLot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParkingLot *ParkingLotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParkingLot.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ParkingLot *ParkingLotCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ParkingLot.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ParkingLot *ParkingLotSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ParkingLot.Contract.BalanceOf(&_ParkingLot.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ParkingLot *ParkingLotCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ParkingLot.Contract.BalanceOf(&_ParkingLot.CallOpts, owner)
}

// CheckRentalStatus is a free data retrieval call binding the contract method 0xcfc9262d.
//
// Solidity: function checkRentalStatus(uint256 tokenId) view returns(bool)
func (_ParkingLot *ParkingLotCaller) CheckRentalStatus(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _ParkingLot.contract.Call(opts, &out, "checkRentalStatus", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckRentalStatus is a free data retrieval call binding the contract method 0xcfc9262d.
//
// Solidity: function checkRentalStatus(uint256 tokenId) view returns(bool)
func (_ParkingLot *ParkingLotSession) CheckRentalStatus(tokenId *big.Int) (bool, error) {
	return _ParkingLot.Contract.CheckRentalStatus(&_ParkingLot.CallOpts, tokenId)
}

// CheckRentalStatus is a free data retrieval call binding the contract method 0xcfc9262d.
//
// Solidity: function checkRentalStatus(uint256 tokenId) view returns(bool)
func (_ParkingLot *ParkingLotCallerSession) CheckRentalStatus(tokenId *big.Int) (bool, error) {
	return _ParkingLot.Contract.CheckRentalStatus(&_ParkingLot.CallOpts, tokenId)
}

// GetAllParkingSpots is a free data retrieval call binding the contract method 0xf432bf73.
//
// Solidity: function getAllParkingSpots() view returns((uint256,string,bool,address,uint256,uint256,int256,int256)[])
func (_ParkingLot *ParkingLotCaller) GetAllParkingSpots(opts *bind.CallOpts) ([]ParkingLotParkingSpot, error) {
	var out []interface{}
	err := _ParkingLot.contract.Call(opts, &out, "getAllParkingSpots")

	if err != nil {
		return *new([]ParkingLotParkingSpot), err
	}

	out0 := *abi.ConvertType(out[0], new([]ParkingLotParkingSpot)).(*[]ParkingLotParkingSpot)

	return out0, err

}

// GetAllParkingSpots is a free data retrieval call binding the contract method 0xf432bf73.
//
// Solidity: function getAllParkingSpots() view returns((uint256,string,bool,address,uint256,uint256,int256,int256)[])
func (_ParkingLot *ParkingLotSession) GetAllParkingSpots() ([]ParkingLotParkingSpot, error) {
	return _ParkingLot.Contract.GetAllParkingSpots(&_ParkingLot.CallOpts)
}

// GetAllParkingSpots is a free data retrieval call binding the contract method 0xf432bf73.
//
// Solidity: function getAllParkingSpots() view returns((uint256,string,bool,address,uint256,uint256,int256,int256)[])
func (_ParkingLot *ParkingLotCallerSession) GetAllParkingSpots() ([]ParkingLotParkingSpot, error) {
	return _ParkingLot.Contract.GetAllParkingSpots(&_ParkingLot.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ParkingLot *ParkingLotCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ParkingLot.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ParkingLot *ParkingLotSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ParkingLot.Contract.GetApproved(&_ParkingLot.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ParkingLot *ParkingLotCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ParkingLot.Contract.GetApproved(&_ParkingLot.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ParkingLot *ParkingLotCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ParkingLot.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ParkingLot *ParkingLotSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ParkingLot.Contract.IsApprovedForAll(&_ParkingLot.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ParkingLot *ParkingLotCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ParkingLot.Contract.IsApprovedForAll(&_ParkingLot.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ParkingLot *ParkingLotCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ParkingLot.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ParkingLot *ParkingLotSession) Name() (string, error) {
	return _ParkingLot.Contract.Name(&_ParkingLot.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ParkingLot *ParkingLotCallerSession) Name() (string, error) {
	return _ParkingLot.Contract.Name(&_ParkingLot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ParkingLot *ParkingLotCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ParkingLot.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ParkingLot *ParkingLotSession) Owner() (common.Address, error) {
	return _ParkingLot.Contract.Owner(&_ParkingLot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ParkingLot *ParkingLotCallerSession) Owner() (common.Address, error) {
	return _ParkingLot.Contract.Owner(&_ParkingLot.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ParkingLot *ParkingLotCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ParkingLot.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ParkingLot *ParkingLotSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ParkingLot.Contract.OwnerOf(&_ParkingLot.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ParkingLot *ParkingLotCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ParkingLot.Contract.OwnerOf(&_ParkingLot.CallOpts, tokenId)
}

// ParkingSpots is a free data retrieval call binding the contract method 0xf5defcd3.
//
// Solidity: function parkingSpots(uint256 ) view returns(uint256 id, string location, bool isRented, address renter, uint256 rentEndTime, uint256 rentPrice, int256 latitude, int256 longitude)
func (_ParkingLot *ParkingLotCaller) ParkingSpots(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id          *big.Int
	Location    string
	IsRented    bool
	Renter      common.Address
	RentEndTime *big.Int
	RentPrice   *big.Int
	Latitude    *big.Int
	Longitude   *big.Int
}, error) {
	var out []interface{}
	err := _ParkingLot.contract.Call(opts, &out, "parkingSpots", arg0)

	outstruct := new(struct {
		Id          *big.Int
		Location    string
		IsRented    bool
		Renter      common.Address
		RentEndTime *big.Int
		RentPrice   *big.Int
		Latitude    *big.Int
		Longitude   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Location = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.IsRented = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Renter = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.RentEndTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RentPrice = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Latitude = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Longitude = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ParkingSpots is a free data retrieval call binding the contract method 0xf5defcd3.
//
// Solidity: function parkingSpots(uint256 ) view returns(uint256 id, string location, bool isRented, address renter, uint256 rentEndTime, uint256 rentPrice, int256 latitude, int256 longitude)
func (_ParkingLot *ParkingLotSession) ParkingSpots(arg0 *big.Int) (struct {
	Id          *big.Int
	Location    string
	IsRented    bool
	Renter      common.Address
	RentEndTime *big.Int
	RentPrice   *big.Int
	Latitude    *big.Int
	Longitude   *big.Int
}, error) {
	return _ParkingLot.Contract.ParkingSpots(&_ParkingLot.CallOpts, arg0)
}

// ParkingSpots is a free data retrieval call binding the contract method 0xf5defcd3.
//
// Solidity: function parkingSpots(uint256 ) view returns(uint256 id, string location, bool isRented, address renter, uint256 rentEndTime, uint256 rentPrice, int256 latitude, int256 longitude)
func (_ParkingLot *ParkingLotCallerSession) ParkingSpots(arg0 *big.Int) (struct {
	Id          *big.Int
	Location    string
	IsRented    bool
	Renter      common.Address
	RentEndTime *big.Int
	RentPrice   *big.Int
	Latitude    *big.Int
	Longitude   *big.Int
}, error) {
	return _ParkingLot.Contract.ParkingSpots(&_ParkingLot.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ParkingLot *ParkingLotCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ParkingLot.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ParkingLot *ParkingLotSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ParkingLot.Contract.SupportsInterface(&_ParkingLot.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ParkingLot *ParkingLotCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ParkingLot.Contract.SupportsInterface(&_ParkingLot.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ParkingLot *ParkingLotCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ParkingLot.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ParkingLot *ParkingLotSession) Symbol() (string, error) {
	return _ParkingLot.Contract.Symbol(&_ParkingLot.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ParkingLot *ParkingLotCallerSession) Symbol() (string, error) {
	return _ParkingLot.Contract.Symbol(&_ParkingLot.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ParkingLot *ParkingLotCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ParkingLot.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ParkingLot *ParkingLotSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ParkingLot.Contract.TokenURI(&_ParkingLot.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ParkingLot *ParkingLotCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ParkingLot.Contract.TokenURI(&_ParkingLot.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ParkingLot *ParkingLotCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ParkingLot.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ParkingLot *ParkingLotSession) TotalSupply() (*big.Int, error) {
	return _ParkingLot.Contract.TotalSupply(&_ParkingLot.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ParkingLot *ParkingLotCallerSession) TotalSupply() (*big.Int, error) {
	return _ParkingLot.Contract.TotalSupply(&_ParkingLot.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ParkingLot *ParkingLotTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParkingLot.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ParkingLot *ParkingLotSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParkingLot.Contract.Approve(&_ParkingLot.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ParkingLot *ParkingLotTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParkingLot.Contract.Approve(&_ParkingLot.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ParkingLot *ParkingLotTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ParkingLot.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ParkingLot *ParkingLotSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _ParkingLot.Contract.Burn(&_ParkingLot.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ParkingLot *ParkingLotTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _ParkingLot.Contract.Burn(&_ParkingLot.TransactOpts, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x42a1497e.
//
// Solidity: function mint(address to, string location, uint256 rentPrice, int256 latitude, int256 longitude) returns()
func (_ParkingLot *ParkingLotTransactor) Mint(opts *bind.TransactOpts, to common.Address, location string, rentPrice *big.Int, latitude *big.Int, longitude *big.Int) (*types.Transaction, error) {
	return _ParkingLot.contract.Transact(opts, "mint", to, location, rentPrice, latitude, longitude)
}

// Mint is a paid mutator transaction binding the contract method 0x42a1497e.
//
// Solidity: function mint(address to, string location, uint256 rentPrice, int256 latitude, int256 longitude) returns()
func (_ParkingLot *ParkingLotSession) Mint(to common.Address, location string, rentPrice *big.Int, latitude *big.Int, longitude *big.Int) (*types.Transaction, error) {
	return _ParkingLot.Contract.Mint(&_ParkingLot.TransactOpts, to, location, rentPrice, latitude, longitude)
}

// Mint is a paid mutator transaction binding the contract method 0x42a1497e.
//
// Solidity: function mint(address to, string location, uint256 rentPrice, int256 latitude, int256 longitude) returns()
func (_ParkingLot *ParkingLotTransactorSession) Mint(to common.Address, location string, rentPrice *big.Int, latitude *big.Int, longitude *big.Int) (*types.Transaction, error) {
	return _ParkingLot.Contract.Mint(&_ParkingLot.TransactOpts, to, location, rentPrice, latitude, longitude)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ParkingLot *ParkingLotTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParkingLot.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ParkingLot *ParkingLotSession) RenounceOwnership() (*types.Transaction, error) {
	return _ParkingLot.Contract.RenounceOwnership(&_ParkingLot.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ParkingLot *ParkingLotTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ParkingLot.Contract.RenounceOwnership(&_ParkingLot.TransactOpts)
}

// Rent is a paid mutator transaction binding the contract method 0x783a112b.
//
// Solidity: function rent(uint256 tokenId, uint256 duration) payable returns()
func (_ParkingLot *ParkingLotTransactor) Rent(opts *bind.TransactOpts, tokenId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _ParkingLot.contract.Transact(opts, "rent", tokenId, duration)
}

// Rent is a paid mutator transaction binding the contract method 0x783a112b.
//
// Solidity: function rent(uint256 tokenId, uint256 duration) payable returns()
func (_ParkingLot *ParkingLotSession) Rent(tokenId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _ParkingLot.Contract.Rent(&_ParkingLot.TransactOpts, tokenId, duration)
}

// Rent is a paid mutator transaction binding the contract method 0x783a112b.
//
// Solidity: function rent(uint256 tokenId, uint256 duration) payable returns()
func (_ParkingLot *ParkingLotTransactorSession) Rent(tokenId *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _ParkingLot.Contract.Rent(&_ParkingLot.TransactOpts, tokenId, duration)
}

// RevokeParkingSpot is a paid mutator transaction binding the contract method 0xd69ed0cc.
//
// Solidity: function revokeParkingSpot(uint256 tokenId) returns()
func (_ParkingLot *ParkingLotTransactor) RevokeParkingSpot(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ParkingLot.contract.Transact(opts, "revokeParkingSpot", tokenId)
}

// RevokeParkingSpot is a paid mutator transaction binding the contract method 0xd69ed0cc.
//
// Solidity: function revokeParkingSpot(uint256 tokenId) returns()
func (_ParkingLot *ParkingLotSession) RevokeParkingSpot(tokenId *big.Int) (*types.Transaction, error) {
	return _ParkingLot.Contract.RevokeParkingSpot(&_ParkingLot.TransactOpts, tokenId)
}

// RevokeParkingSpot is a paid mutator transaction binding the contract method 0xd69ed0cc.
//
// Solidity: function revokeParkingSpot(uint256 tokenId) returns()
func (_ParkingLot *ParkingLotTransactorSession) RevokeParkingSpot(tokenId *big.Int) (*types.Transaction, error) {
	return _ParkingLot.Contract.RevokeParkingSpot(&_ParkingLot.TransactOpts, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ParkingLot *ParkingLotTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParkingLot.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ParkingLot *ParkingLotSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParkingLot.Contract.SafeTransferFrom(&_ParkingLot.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ParkingLot *ParkingLotTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParkingLot.Contract.SafeTransferFrom(&_ParkingLot.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ParkingLot *ParkingLotTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ParkingLot.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ParkingLot *ParkingLotSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ParkingLot.Contract.SafeTransferFrom0(&_ParkingLot.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_ParkingLot *ParkingLotTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _ParkingLot.Contract.SafeTransferFrom0(&_ParkingLot.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ParkingLot *ParkingLotTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ParkingLot.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ParkingLot *ParkingLotSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ParkingLot.Contract.SetApprovalForAll(&_ParkingLot.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ParkingLot *ParkingLotTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ParkingLot.Contract.SetApprovalForAll(&_ParkingLot.TransactOpts, operator, approved)
}

// TerminateRental is a paid mutator transaction binding the contract method 0xf9217419.
//
// Solidity: function terminateRental(uint256 tokenId) returns()
func (_ParkingLot *ParkingLotTransactor) TerminateRental(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ParkingLot.contract.Transact(opts, "terminateRental", tokenId)
}

// TerminateRental is a paid mutator transaction binding the contract method 0xf9217419.
//
// Solidity: function terminateRental(uint256 tokenId) returns()
func (_ParkingLot *ParkingLotSession) TerminateRental(tokenId *big.Int) (*types.Transaction, error) {
	return _ParkingLot.Contract.TerminateRental(&_ParkingLot.TransactOpts, tokenId)
}

// TerminateRental is a paid mutator transaction binding the contract method 0xf9217419.
//
// Solidity: function terminateRental(uint256 tokenId) returns()
func (_ParkingLot *ParkingLotTransactorSession) TerminateRental(tokenId *big.Int) (*types.Transaction, error) {
	return _ParkingLot.Contract.TerminateRental(&_ParkingLot.TransactOpts, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ParkingLot *ParkingLotTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParkingLot.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ParkingLot *ParkingLotSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParkingLot.Contract.TransferFrom(&_ParkingLot.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ParkingLot *ParkingLotTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ParkingLot.Contract.TransferFrom(&_ParkingLot.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ParkingLot *ParkingLotTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ParkingLot.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ParkingLot *ParkingLotSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ParkingLot.Contract.TransferOwnership(&_ParkingLot.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ParkingLot *ParkingLotTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ParkingLot.Contract.TransferOwnership(&_ParkingLot.TransactOpts, newOwner)
}

// ParkingLotApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ParkingLot contract.
type ParkingLotApprovalIterator struct {
	Event *ParkingLotApproval // Event containing the contract specifics and raw log

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
func (it *ParkingLotApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParkingLotApproval)
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
		it.Event = new(ParkingLotApproval)
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
func (it *ParkingLotApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParkingLotApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParkingLotApproval represents a Approval event raised by the ParkingLot contract.
type ParkingLotApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ParkingLot *ParkingLotFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ParkingLotApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ParkingLot.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ParkingLotApprovalIterator{contract: _ParkingLot.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ParkingLot *ParkingLotFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ParkingLotApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ParkingLot.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParkingLotApproval)
				if err := _ParkingLot.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ParkingLot *ParkingLotFilterer) ParseApproval(log types.Log) (*ParkingLotApproval, error) {
	event := new(ParkingLotApproval)
	if err := _ParkingLot.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParkingLotApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ParkingLot contract.
type ParkingLotApprovalForAllIterator struct {
	Event *ParkingLotApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ParkingLotApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParkingLotApprovalForAll)
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
		it.Event = new(ParkingLotApprovalForAll)
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
func (it *ParkingLotApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParkingLotApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParkingLotApprovalForAll represents a ApprovalForAll event raised by the ParkingLot contract.
type ParkingLotApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ParkingLot *ParkingLotFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ParkingLotApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ParkingLot.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ParkingLotApprovalForAllIterator{contract: _ParkingLot.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ParkingLot *ParkingLotFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ParkingLotApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ParkingLot.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParkingLotApprovalForAll)
				if err := _ParkingLot.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ParkingLot *ParkingLotFilterer) ParseApprovalForAll(log types.Log) (*ParkingLotApprovalForAll, error) {
	event := new(ParkingLotApprovalForAll)
	if err := _ParkingLot.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParkingLotOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ParkingLot contract.
type ParkingLotOwnershipTransferredIterator struct {
	Event *ParkingLotOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ParkingLotOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParkingLotOwnershipTransferred)
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
		it.Event = new(ParkingLotOwnershipTransferred)
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
func (it *ParkingLotOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParkingLotOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParkingLotOwnershipTransferred represents a OwnershipTransferred event raised by the ParkingLot contract.
type ParkingLotOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ParkingLot *ParkingLotFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ParkingLotOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ParkingLot.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ParkingLotOwnershipTransferredIterator{contract: _ParkingLot.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ParkingLot *ParkingLotFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ParkingLotOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ParkingLot.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParkingLotOwnershipTransferred)
				if err := _ParkingLot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ParkingLot *ParkingLotFilterer) ParseOwnershipTransferred(log types.Log) (*ParkingLotOwnershipTransferred, error) {
	event := new(ParkingLotOwnershipTransferred)
	if err := _ParkingLot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ParkingLotTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ParkingLot contract.
type ParkingLotTransferIterator struct {
	Event *ParkingLotTransfer // Event containing the contract specifics and raw log

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
func (it *ParkingLotTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParkingLotTransfer)
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
		it.Event = new(ParkingLotTransfer)
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
func (it *ParkingLotTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParkingLotTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParkingLotTransfer represents a Transfer event raised by the ParkingLot contract.
type ParkingLotTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ParkingLot *ParkingLotFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ParkingLotTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ParkingLot.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ParkingLotTransferIterator{contract: _ParkingLot.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ParkingLot *ParkingLotFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ParkingLotTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ParkingLot.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParkingLotTransfer)
				if err := _ParkingLot.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ParkingLot *ParkingLotFilterer) ParseTransfer(log types.Log) (*ParkingLotTransfer, error) {
	event := new(ParkingLotTransfer)
	if err := _ParkingLot.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
