// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package baas721

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Baas721ABI is the input ABI used to generate the binding from.
const Baas721ABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"memoAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"property\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenUri\",\"type\":\"string\"}],\"name\":\"tokenMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"transactionMemo\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"}],\"name\":\"addMemo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getProperty\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_property\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenUri\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"property\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_property\",\"type\":\"string\"}],\"name\":\"setProperty\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_memo\",\"type\":\"string\"}],\"name\":\"transferWithMemo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Baas721 is an auto generated Go binding around an Ethereum contract.
type Baas721 struct {
	Baas721Caller     // Read-only binding to the contract
	Baas721Transactor // Write-only binding to the contract
	Baas721Filterer   // Log filterer for contract events
}

// Baas721Caller is an auto generated read-only Go binding around an Ethereum contract.
type Baas721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Baas721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Baas721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Baas721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Baas721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Baas721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Baas721Session struct {
	Contract     *Baas721          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Baas721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Baas721CallerSession struct {
	Contract *Baas721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Baas721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Baas721TransactorSession struct {
	Contract     *Baas721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Baas721Raw is an auto generated low-level Go binding around an Ethereum contract.
type Baas721Raw struct {
	Contract *Baas721 // Generic contract binding to access the raw methods on
}

// Baas721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Baas721CallerRaw struct {
	Contract *Baas721Caller // Generic read-only contract binding to access the raw methods on
}

// Baas721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Baas721TransactorRaw struct {
	Contract *Baas721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBaas721 creates a new instance of Baas721, bound to a specific deployed contract.
func NewBaas721(address common.Address, backend bind.ContractBackend) (*Baas721, error) {
	contract, err := bindBaas721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Baas721{Baas721Caller: Baas721Caller{contract: contract}, Baas721Transactor: Baas721Transactor{contract: contract}, Baas721Filterer: Baas721Filterer{contract: contract}}, nil
}

// NewBaas721Caller creates a new read-only instance of Baas721, bound to a specific deployed contract.
func NewBaas721Caller(address common.Address, caller bind.ContractCaller) (*Baas721Caller, error) {
	contract, err := bindBaas721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Baas721Caller{contract: contract}, nil
}

// NewBaas721Transactor creates a new write-only instance of Baas721, bound to a specific deployed contract.
func NewBaas721Transactor(address common.Address, transactor bind.ContractTransactor) (*Baas721Transactor, error) {
	contract, err := bindBaas721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Baas721Transactor{contract: contract}, nil
}

// NewBaas721Filterer creates a new log filterer instance of Baas721, bound to a specific deployed contract.
func NewBaas721Filterer(address common.Address, filterer bind.ContractFilterer) (*Baas721Filterer, error) {
	contract, err := bindBaas721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Baas721Filterer{contract: contract}, nil
}

// bindBaas721 binds a generic wrapper to an already deployed contract.
func bindBaas721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Baas721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Baas721 *Baas721Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Baas721.Contract.Baas721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Baas721 *Baas721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Baas721.Contract.Baas721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Baas721 *Baas721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Baas721.Contract.Baas721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Baas721 *Baas721CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Baas721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Baas721 *Baas721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Baas721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Baas721 *Baas721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Baas721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Baas721 *Baas721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Baas721.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Baas721 *Baas721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Baas721.Contract.BalanceOf(&_Baas721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Baas721 *Baas721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Baas721.Contract.BalanceOf(&_Baas721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Baas721 *Baas721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Baas721.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Baas721 *Baas721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Baas721.Contract.GetApproved(&_Baas721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Baas721 *Baas721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Baas721.Contract.GetApproved(&_Baas721.CallOpts, tokenId)
}

// GetProperty is a free data retrieval call binding the contract method 0x32665ffb.
//
// Solidity: function getProperty(uint256 _tokenId) view returns(string)
func (_Baas721 *Baas721Caller) GetProperty(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Baas721.contract.Call(opts, out, "getProperty", _tokenId)
	return *ret0, err
}

// GetProperty is a free data retrieval call binding the contract method 0x32665ffb.
//
// Solidity: function getProperty(uint256 _tokenId) view returns(string)
func (_Baas721 *Baas721Session) GetProperty(_tokenId *big.Int) (string, error) {
	return _Baas721.Contract.GetProperty(&_Baas721.CallOpts, _tokenId)
}

// GetProperty is a free data retrieval call binding the contract method 0x32665ffb.
//
// Solidity: function getProperty(uint256 _tokenId) view returns(string)
func (_Baas721 *Baas721CallerSession) GetProperty(_tokenId *big.Int) (string, error) {
	return _Baas721.Contract.GetProperty(&_Baas721.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Baas721 *Baas721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Baas721.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Baas721 *Baas721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Baas721.Contract.IsApprovedForAll(&_Baas721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Baas721 *Baas721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Baas721.Contract.IsApprovedForAll(&_Baas721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Baas721 *Baas721Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Baas721.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Baas721 *Baas721Session) Name() (string, error) {
	return _Baas721.Contract.Name(&_Baas721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Baas721 *Baas721CallerSession) Name() (string, error) {
	return _Baas721.Contract.Name(&_Baas721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Baas721 *Baas721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Baas721.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Baas721 *Baas721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Baas721.Contract.OwnerOf(&_Baas721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Baas721 *Baas721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Baas721.Contract.OwnerOf(&_Baas721.CallOpts, tokenId)
}

// Property is a free data retrieval call binding the contract method 0x30e1ad32.
//
// Solidity: function property(uint256 ) view returns(string)
func (_Baas721 *Baas721Caller) Property(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Baas721.contract.Call(opts, out, "property", arg0)
	return *ret0, err
}

// Property is a free data retrieval call binding the contract method 0x30e1ad32.
//
// Solidity: function property(uint256 ) view returns(string)
func (_Baas721 *Baas721Session) Property(arg0 *big.Int) (string, error) {
	return _Baas721.Contract.Property(&_Baas721.CallOpts, arg0)
}

// Property is a free data retrieval call binding the contract method 0x30e1ad32.
//
// Solidity: function property(uint256 ) view returns(string)
func (_Baas721 *Baas721CallerSession) Property(arg0 *big.Int) (string, error) {
	return _Baas721.Contract.Property(&_Baas721.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Baas721 *Baas721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Baas721.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Baas721 *Baas721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Baas721.Contract.SupportsInterface(&_Baas721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Baas721 *Baas721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Baas721.Contract.SupportsInterface(&_Baas721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Baas721 *Baas721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Baas721.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Baas721 *Baas721Session) Symbol() (string, error) {
	return _Baas721.Contract.Symbol(&_Baas721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Baas721 *Baas721CallerSession) Symbol() (string, error) {
	return _Baas721.Contract.Symbol(&_Baas721.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Baas721 *Baas721Caller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Baas721.contract.Call(opts, out, "tokenByIndex", index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Baas721 *Baas721Session) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Baas721.Contract.TokenByIndex(&_Baas721.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Baas721 *Baas721CallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Baas721.Contract.TokenByIndex(&_Baas721.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Baas721 *Baas721Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Baas721.contract.Call(opts, out, "tokenOfOwnerByIndex", owner, index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Baas721 *Baas721Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Baas721.Contract.TokenOfOwnerByIndex(&_Baas721.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Baas721 *Baas721CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Baas721.Contract.TokenOfOwnerByIndex(&_Baas721.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Baas721 *Baas721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Baas721.contract.Call(opts, out, "tokenURI", tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Baas721 *Baas721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _Baas721.Contract.TokenURI(&_Baas721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Baas721 *Baas721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Baas721.Contract.TokenURI(&_Baas721.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Baas721 *Baas721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Baas721.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Baas721 *Baas721Session) TotalSupply() (*big.Int, error) {
	return _Baas721.Contract.TotalSupply(&_Baas721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Baas721 *Baas721CallerSession) TotalSupply() (*big.Int, error) {
	return _Baas721.Contract.TotalSupply(&_Baas721.CallOpts)
}

// AddMemo is a paid mutator transaction binding the contract method 0xf93e4ec8.
//
// Solidity: function addMemo(uint256 _tokenId, string _memo) returns()
func (_Baas721 *Baas721Transactor) AddMemo(opts *bind.TransactOpts, _tokenId *big.Int, _memo string) (*types.Transaction, error) {
	return _Baas721.contract.Transact(opts, "addMemo", _tokenId, _memo)
}

// AddMemo is a paid mutator transaction binding the contract method 0xf93e4ec8.
//
// Solidity: function addMemo(uint256 _tokenId, string _memo) returns()
func (_Baas721 *Baas721Session) AddMemo(_tokenId *big.Int, _memo string) (*types.Transaction, error) {
	return _Baas721.Contract.AddMemo(&_Baas721.TransactOpts, _tokenId, _memo)
}

// AddMemo is a paid mutator transaction binding the contract method 0xf93e4ec8.
//
// Solidity: function addMemo(uint256 _tokenId, string _memo) returns()
func (_Baas721 *Baas721TransactorSession) AddMemo(_tokenId *big.Int, _memo string) (*types.Transaction, error) {
	return _Baas721.Contract.AddMemo(&_Baas721.TransactOpts, _tokenId, _memo)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Baas721 *Baas721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Baas721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Baas721 *Baas721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Baas721.Contract.Approve(&_Baas721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Baas721 *Baas721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Baas721.Contract.Approve(&_Baas721.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x2825df7b.
//
// Solidity: function mint(uint256 _tokenId, string _property, string _tokenUri) returns()
func (_Baas721 *Baas721Transactor) Mint(opts *bind.TransactOpts, _tokenId *big.Int, _property string, _tokenUri string) (*types.Transaction, error) {
	return _Baas721.contract.Transact(opts, "mint", _tokenId, _property, _tokenUri)
}

// Mint is a paid mutator transaction binding the contract method 0x2825df7b.
//
// Solidity: function mint(uint256 _tokenId, string _property, string _tokenUri) returns()
func (_Baas721 *Baas721Session) Mint(_tokenId *big.Int, _property string, _tokenUri string) (*types.Transaction, error) {
	return _Baas721.Contract.Mint(&_Baas721.TransactOpts, _tokenId, _property, _tokenUri)
}

// Mint is a paid mutator transaction binding the contract method 0x2825df7b.
//
// Solidity: function mint(uint256 _tokenId, string _property, string _tokenUri) returns()
func (_Baas721 *Baas721TransactorSession) Mint(_tokenId *big.Int, _property string, _tokenUri string) (*types.Transaction, error) {
	return _Baas721.Contract.Mint(&_Baas721.TransactOpts, _tokenId, _property, _tokenUri)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Baas721 *Baas721Transactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Baas721.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Baas721 *Baas721Session) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Baas721.Contract.OnERC721Received(&_Baas721.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Baas721 *Baas721TransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Baas721.Contract.OnERC721Received(&_Baas721.TransactOpts, arg0, arg1, arg2, arg3)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Baas721 *Baas721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Baas721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Baas721 *Baas721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Baas721.Contract.SafeTransferFrom(&_Baas721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Baas721 *Baas721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Baas721.Contract.SafeTransferFrom(&_Baas721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Baas721 *Baas721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Baas721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Baas721 *Baas721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Baas721.Contract.SafeTransferFrom0(&_Baas721.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Baas721 *Baas721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Baas721.Contract.SafeTransferFrom0(&_Baas721.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_Baas721 *Baas721Transactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _Baas721.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_Baas721 *Baas721Session) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _Baas721.Contract.SetApprovalForAll(&_Baas721.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_Baas721 *Baas721TransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _Baas721.Contract.SetApprovalForAll(&_Baas721.TransactOpts, to, approved)
}

// SetProperty is a paid mutator transaction binding the contract method 0x077419f3.
//
// Solidity: function setProperty(uint256 _tokenId, string _property) returns()
func (_Baas721 *Baas721Transactor) SetProperty(opts *bind.TransactOpts, _tokenId *big.Int, _property string) (*types.Transaction, error) {
	return _Baas721.contract.Transact(opts, "setProperty", _tokenId, _property)
}

// SetProperty is a paid mutator transaction binding the contract method 0x077419f3.
//
// Solidity: function setProperty(uint256 _tokenId, string _property) returns()
func (_Baas721 *Baas721Session) SetProperty(_tokenId *big.Int, _property string) (*types.Transaction, error) {
	return _Baas721.Contract.SetProperty(&_Baas721.TransactOpts, _tokenId, _property)
}

// SetProperty is a paid mutator transaction binding the contract method 0x077419f3.
//
// Solidity: function setProperty(uint256 _tokenId, string _property) returns()
func (_Baas721 *Baas721TransactorSession) SetProperty(_tokenId *big.Int, _property string) (*types.Transaction, error) {
	return _Baas721.Contract.SetProperty(&_Baas721.TransactOpts, _tokenId, _property)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Baas721 *Baas721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Baas721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Baas721 *Baas721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Baas721.Contract.TransferFrom(&_Baas721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Baas721 *Baas721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Baas721.Contract.TransferFrom(&_Baas721.TransactOpts, from, to, tokenId)
}

// TransferWithMemo is a paid mutator transaction binding the contract method 0xdcb49c73.
//
// Solidity: function transferWithMemo(address _from, address _to, uint256 _tokenId, string _memo) returns()
func (_Baas721 *Baas721Transactor) TransferWithMemo(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _memo string) (*types.Transaction, error) {
	return _Baas721.contract.Transact(opts, "transferWithMemo", _from, _to, _tokenId, _memo)
}

// TransferWithMemo is a paid mutator transaction binding the contract method 0xdcb49c73.
//
// Solidity: function transferWithMemo(address _from, address _to, uint256 _tokenId, string _memo) returns()
func (_Baas721 *Baas721Session) TransferWithMemo(_from common.Address, _to common.Address, _tokenId *big.Int, _memo string) (*types.Transaction, error) {
	return _Baas721.Contract.TransferWithMemo(&_Baas721.TransactOpts, _from, _to, _tokenId, _memo)
}

// TransferWithMemo is a paid mutator transaction binding the contract method 0xdcb49c73.
//
// Solidity: function transferWithMemo(address _from, address _to, uint256 _tokenId, string _memo) returns()
func (_Baas721 *Baas721TransactorSession) TransferWithMemo(_from common.Address, _to common.Address, _tokenId *big.Int, _memo string) (*types.Transaction, error) {
	return _Baas721.Contract.TransferWithMemo(&_Baas721.TransactOpts, _from, _to, _tokenId, _memo)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Baas721 *Baas721Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Baas721.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Baas721 *Baas721Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Baas721.Contract.Fallback(&_Baas721.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Baas721 *Baas721TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Baas721.Contract.Fallback(&_Baas721.TransactOpts, calldata)
}

// Baas721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Baas721 contract.
type Baas721ApprovalIterator struct {
	Event *Baas721Approval // Event containing the contract specifics and raw log

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
func (it *Baas721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Baas721Approval)
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
		it.Event = new(Baas721Approval)
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
func (it *Baas721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Baas721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Baas721Approval represents a Approval event raised by the Baas721 contract.
type Baas721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Baas721 *Baas721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Baas721ApprovalIterator, error) {

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

	logs, sub, err := _Baas721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Baas721ApprovalIterator{contract: _Baas721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Baas721 *Baas721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Baas721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Baas721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Baas721Approval)
				if err := _Baas721.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Baas721 *Baas721Filterer) ParseApproval(log types.Log) (*Baas721Approval, error) {
	event := new(Baas721Approval)
	if err := _Baas721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Baas721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Baas721 contract.
type Baas721ApprovalForAllIterator struct {
	Event *Baas721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Baas721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Baas721ApprovalForAll)
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
		it.Event = new(Baas721ApprovalForAll)
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
func (it *Baas721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Baas721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Baas721ApprovalForAll represents a ApprovalForAll event raised by the Baas721 contract.
type Baas721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Baas721 *Baas721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Baas721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Baas721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Baas721ApprovalForAllIterator{contract: _Baas721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Baas721 *Baas721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Baas721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Baas721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Baas721ApprovalForAll)
				if err := _Baas721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Baas721 *Baas721Filterer) ParseApprovalForAll(log types.Log) (*Baas721ApprovalForAll, error) {
	event := new(Baas721ApprovalForAll)
	if err := _Baas721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Baas721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Baas721 contract.
type Baas721TransferIterator struct {
	Event *Baas721Transfer // Event containing the contract specifics and raw log

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
func (it *Baas721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Baas721Transfer)
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
		it.Event = new(Baas721Transfer)
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
func (it *Baas721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Baas721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Baas721Transfer represents a Transfer event raised by the Baas721 contract.
type Baas721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Baas721 *Baas721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Baas721TransferIterator, error) {

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

	logs, sub, err := _Baas721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Baas721TransferIterator{contract: _Baas721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Baas721 *Baas721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Baas721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Baas721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Baas721Transfer)
				if err := _Baas721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Baas721 *Baas721Filterer) ParseTransfer(log types.Log) (*Baas721Transfer, error) {
	event := new(Baas721Transfer)
	if err := _Baas721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Baas721MemoAddedIterator is returned from FilterMemoAdded and is used to iterate over the raw logs and unpacked data for MemoAdded events raised by the Baas721 contract.
type Baas721MemoAddedIterator struct {
	Event *Baas721MemoAdded // Event containing the contract specifics and raw log

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
func (it *Baas721MemoAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Baas721MemoAdded)
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
		it.Event = new(Baas721MemoAdded)
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
func (it *Baas721MemoAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Baas721MemoAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Baas721MemoAdded represents a MemoAdded event raised by the Baas721 contract.
type Baas721MemoAdded struct {
	TokenId *big.Int
	Memo    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMemoAdded is a free log retrieval operation binding the contract event 0x4a62fa304d9a5993f36546283e9d483bc1cb4909a24a0bb0104fb47505caea20.
//
// Solidity: event memoAdded(uint256 indexed tokenId, string memo)
func (_Baas721 *Baas721Filterer) FilterMemoAdded(opts *bind.FilterOpts, tokenId []*big.Int) (*Baas721MemoAddedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Baas721.contract.FilterLogs(opts, "memoAdded", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Baas721MemoAddedIterator{contract: _Baas721.contract, event: "memoAdded", logs: logs, sub: sub}, nil
}

// WatchMemoAdded is a free log subscription operation binding the contract event 0x4a62fa304d9a5993f36546283e9d483bc1cb4909a24a0bb0104fb47505caea20.
//
// Solidity: event memoAdded(uint256 indexed tokenId, string memo)
func (_Baas721 *Baas721Filterer) WatchMemoAdded(opts *bind.WatchOpts, sink chan<- *Baas721MemoAdded, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Baas721.contract.WatchLogs(opts, "memoAdded", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Baas721MemoAdded)
				if err := _Baas721.contract.UnpackLog(event, "memoAdded", log); err != nil {
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

// ParseMemoAdded is a log parse operation binding the contract event 0x4a62fa304d9a5993f36546283e9d483bc1cb4909a24a0bb0104fb47505caea20.
//
// Solidity: event memoAdded(uint256 indexed tokenId, string memo)
func (_Baas721 *Baas721Filterer) ParseMemoAdded(log types.Log) (*Baas721MemoAdded, error) {
	event := new(Baas721MemoAdded)
	if err := _Baas721.contract.UnpackLog(event, "memoAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Baas721TokenMintedIterator is returned from FilterTokenMinted and is used to iterate over the raw logs and unpacked data for TokenMinted events raised by the Baas721 contract.
type Baas721TokenMintedIterator struct {
	Event *Baas721TokenMinted // Event containing the contract specifics and raw log

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
func (it *Baas721TokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Baas721TokenMinted)
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
		it.Event = new(Baas721TokenMinted)
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
func (it *Baas721TokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Baas721TokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Baas721TokenMinted represents a TokenMinted event raised by the Baas721 contract.
type Baas721TokenMinted struct {
	TokenId  *big.Int
	Property string
	TokenUri string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenMinted is a free log retrieval operation binding the contract event 0x20dc370201d8359018b0db6be114390d5374d7af7458399ed9cba5ffee738ec6.
//
// Solidity: event tokenMinted(uint256 indexed tokenId, string property, string tokenUri)
func (_Baas721 *Baas721Filterer) FilterTokenMinted(opts *bind.FilterOpts, tokenId []*big.Int) (*Baas721TokenMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Baas721.contract.FilterLogs(opts, "tokenMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Baas721TokenMintedIterator{contract: _Baas721.contract, event: "tokenMinted", logs: logs, sub: sub}, nil
}

// WatchTokenMinted is a free log subscription operation binding the contract event 0x20dc370201d8359018b0db6be114390d5374d7af7458399ed9cba5ffee738ec6.
//
// Solidity: event tokenMinted(uint256 indexed tokenId, string property, string tokenUri)
func (_Baas721 *Baas721Filterer) WatchTokenMinted(opts *bind.WatchOpts, sink chan<- *Baas721TokenMinted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Baas721.contract.WatchLogs(opts, "tokenMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Baas721TokenMinted)
				if err := _Baas721.contract.UnpackLog(event, "tokenMinted", log); err != nil {
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

// ParseTokenMinted is a log parse operation binding the contract event 0x20dc370201d8359018b0db6be114390d5374d7af7458399ed9cba5ffee738ec6.
//
// Solidity: event tokenMinted(uint256 indexed tokenId, string property, string tokenUri)
func (_Baas721 *Baas721Filterer) ParseTokenMinted(log types.Log) (*Baas721TokenMinted, error) {
	event := new(Baas721TokenMinted)
	if err := _Baas721.contract.UnpackLog(event, "tokenMinted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// Baas721TransactionMemoIterator is returned from FilterTransactionMemo and is used to iterate over the raw logs and unpacked data for TransactionMemo events raised by the Baas721 contract.
type Baas721TransactionMemoIterator struct {
	Event *Baas721TransactionMemo // Event containing the contract specifics and raw log

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
func (it *Baas721TransactionMemoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Baas721TransactionMemo)
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
		it.Event = new(Baas721TransactionMemo)
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
func (it *Baas721TransactionMemoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Baas721TransactionMemoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Baas721TransactionMemo represents a TransactionMemo event raised by the Baas721 contract.
type Baas721TransactionMemo struct {
	TokenId *big.Int
	Memo    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransactionMemo is a free log retrieval operation binding the contract event 0x6bb79e977ebcae518f2d094cc741d1a77f05ae1d52cd8e23e5b7118572a438ce.
//
// Solidity: event transactionMemo(uint256 indexed tokenId, string memo)
func (_Baas721 *Baas721Filterer) FilterTransactionMemo(opts *bind.FilterOpts, tokenId []*big.Int) (*Baas721TransactionMemoIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Baas721.contract.FilterLogs(opts, "transactionMemo", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Baas721TransactionMemoIterator{contract: _Baas721.contract, event: "transactionMemo", logs: logs, sub: sub}, nil
}

// WatchTransactionMemo is a free log subscription operation binding the contract event 0x6bb79e977ebcae518f2d094cc741d1a77f05ae1d52cd8e23e5b7118572a438ce.
//
// Solidity: event transactionMemo(uint256 indexed tokenId, string memo)
func (_Baas721 *Baas721Filterer) WatchTransactionMemo(opts *bind.WatchOpts, sink chan<- *Baas721TransactionMemo, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Baas721.contract.WatchLogs(opts, "transactionMemo", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Baas721TransactionMemo)
				if err := _Baas721.contract.UnpackLog(event, "transactionMemo", log); err != nil {
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

// ParseTransactionMemo is a log parse operation binding the contract event 0x6bb79e977ebcae518f2d094cc741d1a77f05ae1d52cd8e23e5b7118572a438ce.
//
// Solidity: event transactionMemo(uint256 indexed tokenId, string memo)
func (_Baas721 *Baas721Filterer) ParseTransactionMemo(log types.Log) (*Baas721TransactionMemo, error) {
	event := new(Baas721TransactionMemo)
	if err := _Baas721.contract.UnpackLog(event, "transactionMemo", log); err != nil {
		return nil, err
	}
	return event, nil
}
