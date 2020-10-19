Fatal: Failed to generate ABI binding: 5:9: expected 'IDENT', found 721

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package 721

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





	// AddressABI is the input ABI used to generate the binding from.
	const AddressABI = "[]"

	

	
		// AddressBin is the compiled bytecode used for deploying new contracts.
		var AddressBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820a2f70e877c817cb9a594e485624560e6b37f38469deb4acddae0ce6710c0e9f664736f6c63430005110032"

		// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
		func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend ) (common.Address, *types.Transaction, *Address, error) {
		  parsed, err := abi.JSON(strings.NewReader(AddressABI))
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  
		  address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend )
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  return address, tx, &Address{ AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract} }, nil
		}
	

	// Address is an auto generated Go binding around an Ethereum contract.
	type Address struct {
	  AddressCaller     // Read-only binding to the contract
	  AddressTransactor // Write-only binding to the contract
	  AddressFilterer   // Log filterer for contract events
	}

	// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
	type AddressCaller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
	type AddressTransactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type AddressFilterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// AddressSession is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type AddressSession struct {
	  Contract     *Address        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type AddressCallerSession struct {
	  Contract *AddressCaller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type AddressTransactorSession struct {
	  Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
	type AddressRaw struct {
	  Contract *Address // Generic contract binding to access the raw methods on
	}

	// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type AddressCallerRaw struct {
		Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
	}

	// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type AddressTransactorRaw struct {
		Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
	}

	// NewAddress creates a new instance of Address, bound to a specific deployed contract.
	func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	  contract, err := bindAddress(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &Address{ AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract} }, nil
	}

	// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
	func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	  contract, err := bindAddress(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &AddressCaller{contract: contract}, nil
	}

	// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
	func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	  contract, err := bindAddress(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &AddressTransactor{contract: contract}, nil
	}

	// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
 	func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
 	  contract, err := bindAddress(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &AddressFilterer{contract: contract}, nil
 	}

	// bindAddress binds a generic wrapper to an already deployed contract.
	func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(AddressABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_Address *AddressRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _Address.Contract.AddressTransactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _Address.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _Address.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _Address.Contract.contract.Transact(opts, method, params...)
	}

	

	

	

	

	

	// ContextABI is the input ABI used to generate the binding from.
	const ContextABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

	

	

	// Context is an auto generated Go binding around an Ethereum contract.
	type Context struct {
	  ContextCaller     // Read-only binding to the contract
	  ContextTransactor // Write-only binding to the contract
	  ContextFilterer   // Log filterer for contract events
	}

	// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
	type ContextCaller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
	type ContextTransactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type ContextFilterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ContextSession is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type ContextSession struct {
	  Contract     *Context        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type ContextCallerSession struct {
	  Contract *ContextCaller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type ContextTransactorSession struct {
	  Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
	type ContextRaw struct {
	  Contract *Context // Generic contract binding to access the raw methods on
	}

	// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type ContextCallerRaw struct {
		Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
	}

	// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type ContextTransactorRaw struct {
		Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
	}

	// NewContext creates a new instance of Context, bound to a specific deployed contract.
	func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	  contract, err := bindContext(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &Context{ ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract} }, nil
	}

	// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
	func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	  contract, err := bindContext(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &ContextCaller{contract: contract}, nil
	}

	// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
	func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	  contract, err := bindContext(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &ContextTransactor{contract: contract}, nil
	}

	// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
 	func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
 	  contract, err := bindContext(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &ContextFilterer{contract: contract}, nil
 	}

	// bindContext binds a generic wrapper to an already deployed contract.
	func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(ContextABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_Context *ContextRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _Context.Contract.ContextTransactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _Context.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _Context.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _Context.Contract.contract.Transact(opts, method, params...)
	}

	

	

	

	

	

	// CountersABI is the input ABI used to generate the binding from.
	const CountersABI = "[]"

	

	
		// CountersBin is the compiled bytecode used for deploying new contracts.
		var CountersBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820d688472adf5c1a989af0deab98484b639e1d334286e76660a9d5d37e37ea649b64736f6c63430005110032"

		// DeployCounters deploys a new Ethereum contract, binding an instance of Counters to it.
		func DeployCounters(auth *bind.TransactOpts, backend bind.ContractBackend ) (common.Address, *types.Transaction, *Counters, error) {
		  parsed, err := abi.JSON(strings.NewReader(CountersABI))
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  
		  address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CountersBin), backend )
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  return address, tx, &Counters{ CountersCaller: CountersCaller{contract: contract}, CountersTransactor: CountersTransactor{contract: contract}, CountersFilterer: CountersFilterer{contract: contract} }, nil
		}
	

	// Counters is an auto generated Go binding around an Ethereum contract.
	type Counters struct {
	  CountersCaller     // Read-only binding to the contract
	  CountersTransactor // Write-only binding to the contract
	  CountersFilterer   // Log filterer for contract events
	}

	// CountersCaller is an auto generated read-only Go binding around an Ethereum contract.
	type CountersCaller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// CountersTransactor is an auto generated write-only Go binding around an Ethereum contract.
	type CountersTransactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// CountersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type CountersFilterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// CountersSession is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type CountersSession struct {
	  Contract     *Counters        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// CountersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type CountersCallerSession struct {
	  Contract *CountersCaller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// CountersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type CountersTransactorSession struct {
	  Contract     *CountersTransactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// CountersRaw is an auto generated low-level Go binding around an Ethereum contract.
	type CountersRaw struct {
	  Contract *Counters // Generic contract binding to access the raw methods on
	}

	// CountersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type CountersCallerRaw struct {
		Contract *CountersCaller // Generic read-only contract binding to access the raw methods on
	}

	// CountersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type CountersTransactorRaw struct {
		Contract *CountersTransactor // Generic write-only contract binding to access the raw methods on
	}

	// NewCounters creates a new instance of Counters, bound to a specific deployed contract.
	func NewCounters(address common.Address, backend bind.ContractBackend) (*Counters, error) {
	  contract, err := bindCounters(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &Counters{ CountersCaller: CountersCaller{contract: contract}, CountersTransactor: CountersTransactor{contract: contract}, CountersFilterer: CountersFilterer{contract: contract} }, nil
	}

	// NewCountersCaller creates a new read-only instance of Counters, bound to a specific deployed contract.
	func NewCountersCaller(address common.Address, caller bind.ContractCaller) (*CountersCaller, error) {
	  contract, err := bindCounters(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &CountersCaller{contract: contract}, nil
	}

	// NewCountersTransactor creates a new write-only instance of Counters, bound to a specific deployed contract.
	func NewCountersTransactor(address common.Address, transactor bind.ContractTransactor) (*CountersTransactor, error) {
	  contract, err := bindCounters(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &CountersTransactor{contract: contract}, nil
	}

	// NewCountersFilterer creates a new log filterer instance of Counters, bound to a specific deployed contract.
 	func NewCountersFilterer(address common.Address, filterer bind.ContractFilterer) (*CountersFilterer, error) {
 	  contract, err := bindCounters(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &CountersFilterer{contract: contract}, nil
 	}

	// bindCounters binds a generic wrapper to an already deployed contract.
	func bindCounters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(CountersABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_Counters *CountersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _Counters.Contract.CountersCaller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_Counters *CountersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _Counters.Contract.CountersTransactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_Counters *CountersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _Counters.Contract.CountersTransactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_Counters *CountersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _Counters.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_Counters *CountersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _Counters.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_Counters *CountersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _Counters.Contract.contract.Transact(opts, method, params...)
	}

	

	

	

	

	

	// ERC165ABI is the input ABI used to generate the binding from.
	const ERC165ABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

	
		// ERC165FuncSigs maps the 4-byte function signature to its string representation.
		var ERC165FuncSigs = map[string]string{
			"01ffc9a7": "supportsInterface(bytes4)",
			
		}
	

	

	// ERC165 is an auto generated Go binding around an Ethereum contract.
	type ERC165 struct {
	  ERC165Caller     // Read-only binding to the contract
	  ERC165Transactor // Write-only binding to the contract
	  ERC165Filterer   // Log filterer for contract events
	}

	// ERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
	type ERC165Caller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
	type ERC165Transactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type ERC165Filterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC165Session is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type ERC165Session struct {
	  Contract     *ERC165        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// ERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type ERC165CallerSession struct {
	  Contract *ERC165Caller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// ERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type ERC165TransactorSession struct {
	  Contract     *ERC165Transactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// ERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
	type ERC165Raw struct {
	  Contract *ERC165 // Generic contract binding to access the raw methods on
	}

	// ERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type ERC165CallerRaw struct {
		Contract *ERC165Caller // Generic read-only contract binding to access the raw methods on
	}

	// ERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type ERC165TransactorRaw struct {
		Contract *ERC165Transactor // Generic write-only contract binding to access the raw methods on
	}

	// NewERC165 creates a new instance of ERC165, bound to a specific deployed contract.
	func NewERC165(address common.Address, backend bind.ContractBackend) (*ERC165, error) {
	  contract, err := bindERC165(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC165{ ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract} }, nil
	}

	// NewERC165Caller creates a new read-only instance of ERC165, bound to a specific deployed contract.
	func NewERC165Caller(address common.Address, caller bind.ContractCaller) (*ERC165Caller, error) {
	  contract, err := bindERC165(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC165Caller{contract: contract}, nil
	}

	// NewERC165Transactor creates a new write-only instance of ERC165, bound to a specific deployed contract.
	func NewERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC165Transactor, error) {
	  contract, err := bindERC165(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC165Transactor{contract: contract}, nil
	}

	// NewERC165Filterer creates a new log filterer instance of ERC165, bound to a specific deployed contract.
 	func NewERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC165Filterer, error) {
 	  contract, err := bindERC165(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &ERC165Filterer{contract: contract}, nil
 	}

	// bindERC165 binds a generic wrapper to an already deployed contract.
	func bindERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(ERC165ABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_ERC165 *ERC165Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _ERC165.Contract.ERC165Caller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_ERC165 *ERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _ERC165.Contract.ERC165Transactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_ERC165 *ERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _ERC165.Contract.ERC165Transactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_ERC165 *ERC165CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _ERC165.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_ERC165 *ERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _ERC165.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_ERC165 *ERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _ERC165.Contract.contract.Transact(opts, method, params...)
	}

	
		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_ERC165 *ERC165Caller) SupportsInterface(opts *bind.CallOpts , interfaceId [4]byte ) (bool, error) {
			var (
				ret0 = new(bool)
				
			)
			out := ret0
			err := _ERC165.contract.Call(opts, out, "supportsInterface" , interfaceId)
			return *ret0, err
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_ERC165 *ERC165Session) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts , interfaceId)
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_ERC165 *ERC165CallerSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts , interfaceId)
		}
	

	

	

	

	

	// ERC721ABI is the input ABI used to generate the binding from.
	const ERC721ABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

	
		// ERC721FuncSigs maps the 4-byte function signature to its string representation.
		var ERC721FuncSigs = map[string]string{
			"095ea7b3": "approve(address,uint256)",
			"70a08231": "balanceOf(address)",
			"081812fc": "getApproved(uint256)",
			"e985e9c5": "isApprovedForAll(address,address)",
			"6352211e": "ownerOf(uint256)",
			"42842e0e": "safeTransferFrom(address,address,uint256)",
			"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
			"a22cb465": "setApprovalForAll(address,bool)",
			"01ffc9a7": "supportsInterface(bytes4)",
			"23b872dd": "transferFrom(address,address,uint256)",
			
		}
	

	
		// ERC721Bin is the compiled bytecode used for deploying new contracts.
		var ERC721Bin = "0x608060405234801561001057600080fd5b5061002a6301ffc9a760e01b6001600160e01b0361004816565b6100436380ac58cd60e01b6001600160e01b0361004816565b6100cc565b6001600160e01b031980821614156100a7576040805162461bcd60e51b815260206004820152601c60248201527f4552433136353a20696e76616c696420696e7465726661636520696400000000604482015290519081900360640190fd5b6001600160e01b0319166000908152602081905260409020805460ff19166001179055565b610f94806100db6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80636352211e116100665780636352211e146101b157806370a08231146101ce578063a22cb46514610206578063b88d4fde14610234578063e985e9c5146102fa5761009e565b806301ffc9a7146100a3578063081812fc146100de578063095ea7b31461011757806323b872dd1461014557806342842e0e1461017b575b600080fd5b6100ca600480360360208110156100b957600080fd5b50356001600160e01b031916610328565b604080519115158252519081900360200190f35b6100fb600480360360208110156100f457600080fd5b5035610347565b604080516001600160a01b039092168252519081900360200190f35b6101436004803603604081101561012d57600080fd5b506001600160a01b0381351690602001356103a9565b005b6101436004803603606081101561015b57600080fd5b506001600160a01b038135811691602081013590911690604001356104d1565b6101436004803603606081101561019157600080fd5b506001600160a01b0381358116916020810135909116906040013561052d565b6100fb600480360360208110156101c757600080fd5b5035610548565b6101f4600480360360208110156101e457600080fd5b50356001600160a01b03166105a2565b60408051918252519081900360200190f35b6101436004803603604081101561021c57600080fd5b506001600160a01b038135169060200135151561060a565b6101436004803603608081101561024a57600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561028557600080fd5b82018360208201111561029757600080fd5b803590602001918460018302840111640100000000831117156102b957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061070f945050505050565b6100ca6004803603604081101561031057600080fd5b506001600160a01b038135811691602001351661076d565b6001600160e01b03191660009081526020819052604090205460ff1690565b60006103528261079b565b61038d5760405162461bcd60e51b815260040180806020018281038252602c815260200180610eb9602c913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b60006103b482610548565b9050806001600160a01b0316836001600160a01b031614156104075760405162461bcd60e51b8152600401808060200182810382526021815260200180610f0e6021913960400191505060405180910390fd5b806001600160a01b03166104196107b8565b6001600160a01b0316148061043a575061043a816104356107b8565b61076d565b6104755760405162461bcd60e51b8152600401808060200182810382526038815260200180610e2e6038913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b6104e26104dc6107b8565b826107bc565b61051d5760405162461bcd60e51b8152600401808060200182810382526031815260200180610f2f6031913960400191505060405180910390fd5b610528838383610860565b505050565b6105288383836040518060200160405280600081525061070f565b6000818152600160205260408120546001600160a01b03168061059c5760405162461bcd60e51b8152600401808060200182810382526029815260200180610e906029913960400191505060405180910390fd5b92915050565b60006001600160a01b0382166105e95760405162461bcd60e51b815260040180806020018281038252602a815260200180610e66602a913960400191505060405180910390fd5b6001600160a01b038216600090815260036020526040902061059c906109a4565b6106126107b8565b6001600160a01b0316826001600160a01b03161415610678576040805162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015290519081900360640190fd5b80600460006106856107b8565b6001600160a01b03908116825260208083019390935260409182016000908120918716808252919093529120805460ff1916921515929092179091556106c96107b8565b60408051841515815290516001600160a01b0392909216917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c319181900360200190a35050565b61072061071a6107b8565b836107bc565b61075b5760405162461bcd60e51b8152600401808060200182810382526031815260200180610f2f6031913960400191505060405180910390fd5b610767848484846109a8565b50505050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b3390565b60006107c78261079b565b6108025760405162461bcd60e51b815260040180806020018281038252602c815260200180610e02602c913960400191505060405180910390fd5b600061080d83610548565b9050806001600160a01b0316846001600160a01b031614806108485750836001600160a01b031661083d84610347565b6001600160a01b0316145b806108585750610858818561076d565b949350505050565b826001600160a01b031661087382610548565b6001600160a01b0316146108b85760405162461bcd60e51b8152600401808060200182810382526029815260200180610ee56029913960400191505060405180910390fd5b6001600160a01b0382166108fd5760405162461bcd60e51b8152600401808060200182810382526024815260200180610dde6024913960400191505060405180910390fd5b610906816109fa565b6001600160a01b038316600090815260036020526040902061092790610a37565b6001600160a01b038216600090815260036020526040902061094890610a4e565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b5490565b6109b3848484610860565b6109bf84848484610a57565b6107675760405162461bcd60e51b8152600401808060200182810382526032815260200180610dac6032913960400191505060405180910390fd5b6000818152600260205260409020546001600160a01b031615610a3457600081815260026020526040902080546001600160a01b03191690555b50565b8054610a4a90600163ffffffff610c9216565b9055565b80546001019055565b6000610a6b846001600160a01b0316610cdb565b610a7757506001610858565b600060606001600160a01b038616630a85bd0160e11b610a956107b8565b89888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610b0e578181015183820152602001610af6565b50505050905090810190601f168015610b3b5780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610ba35780518252601f199092019160209182019101610b84565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610c05576040519150601f19603f3d011682016040523d82523d6000602084013e610c0a565b606091505b509150915081610c5b57805115610c245780518082602001fd5b60405162461bcd60e51b8152600401808060200182810382526032815260200180610dac6032913960400191505060405180910390fd5b6000818060200190516020811015610c7257600080fd5b50516001600160e01b031916630a85bd0160e11b14935061085892505050565b6000610cd483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610d14565b9392505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470818114801590610858575050151592915050565b60008184841115610da35760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610d68578181015183820152602001610d50565b50505050905090810190601f168015610d955780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50505090039056fe4552433732313a207472616e7366657220746f206e6f6e20455243373231526563656976657220696d706c656d656e7465724552433732313a207472616e7366657220746f20746865207a65726f20616464726573734552433732313a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4552433732313a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734552433732313a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4552433732313a20617070726f76616c20746f2063757272656e74206f776e65724552433732313a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564a265627a7a7231582014721f2b865d6b7fc3772957431d0140319e7610c04a9326d4d206188a9b4b8164736f6c63430005110032"

		// DeployERC721 deploys a new Ethereum contract, binding an instance of ERC721 to it.
		func DeployERC721(auth *bind.TransactOpts, backend bind.ContractBackend ) (common.Address, *types.Transaction, *ERC721, error) {
		  parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  
		  address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721Bin), backend )
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  return address, tx, &ERC721{ ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract} }, nil
		}
	

	// ERC721 is an auto generated Go binding around an Ethereum contract.
	type ERC721 struct {
	  ERC721Caller     // Read-only binding to the contract
	  ERC721Transactor // Write-only binding to the contract
	  ERC721Filterer   // Log filterer for contract events
	}

	// ERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
	type ERC721Caller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
	type ERC721Transactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type ERC721Filterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC721Session is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type ERC721Session struct {
	  Contract     *ERC721        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// ERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type ERC721CallerSession struct {
	  Contract *ERC721Caller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// ERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type ERC721TransactorSession struct {
	  Contract     *ERC721Transactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// ERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
	type ERC721Raw struct {
	  Contract *ERC721 // Generic contract binding to access the raw methods on
	}

	// ERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type ERC721CallerRaw struct {
		Contract *ERC721Caller // Generic read-only contract binding to access the raw methods on
	}

	// ERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type ERC721TransactorRaw struct {
		Contract *ERC721Transactor // Generic write-only contract binding to access the raw methods on
	}

	// NewERC721 creates a new instance of ERC721, bound to a specific deployed contract.
	func NewERC721(address common.Address, backend bind.ContractBackend) (*ERC721, error) {
	  contract, err := bindERC721(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC721{ ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract} }, nil
	}

	// NewERC721Caller creates a new read-only instance of ERC721, bound to a specific deployed contract.
	func NewERC721Caller(address common.Address, caller bind.ContractCaller) (*ERC721Caller, error) {
	  contract, err := bindERC721(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC721Caller{contract: contract}, nil
	}

	// NewERC721Transactor creates a new write-only instance of ERC721, bound to a specific deployed contract.
	func NewERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC721Transactor, error) {
	  contract, err := bindERC721(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC721Transactor{contract: contract}, nil
	}

	// NewERC721Filterer creates a new log filterer instance of ERC721, bound to a specific deployed contract.
 	func NewERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC721Filterer, error) {
 	  contract, err := bindERC721(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &ERC721Filterer{contract: contract}, nil
 	}

	// bindERC721 binds a generic wrapper to an already deployed contract.
	func bindERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_ERC721 *ERC721Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _ERC721.Contract.ERC721Caller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_ERC721 *ERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _ERC721.Contract.ERC721Transactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_ERC721 *ERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _ERC721.Contract.ERC721Transactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_ERC721 *ERC721CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _ERC721.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_ERC721 *ERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _ERC721.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_ERC721 *ERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _ERC721.Contract.contract.Transact(opts, method, params...)
	}

	
		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256)
		func (_ERC721 *ERC721Caller) BalanceOf(opts *bind.CallOpts , owner common.Address ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _ERC721.contract.Call(opts, out, "balanceOf" , owner)
			return *ret0, err
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256)
		func (_ERC721 *ERC721Session) BalanceOf( owner common.Address ) ( *big.Int,  error) {
		  return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts , owner)
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256)
		func (_ERC721 *ERC721CallerSession) BalanceOf( owner common.Address ) ( *big.Int,  error) {
		  return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts , owner)
		}
	
		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address)
		func (_ERC721 *ERC721Caller) GetApproved(opts *bind.CallOpts , tokenId *big.Int ) (common.Address, error) {
			var (
				ret0 = new(common.Address)
				
			)
			out := ret0
			err := _ERC721.contract.Call(opts, out, "getApproved" , tokenId)
			return *ret0, err
		}

		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address)
		func (_ERC721 *ERC721Session) GetApproved( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721.Contract.GetApproved(&_ERC721.CallOpts , tokenId)
		}

		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address)
		func (_ERC721 *ERC721CallerSession) GetApproved( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721.Contract.GetApproved(&_ERC721.CallOpts , tokenId)
		}
	
		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_ERC721 *ERC721Caller) IsApprovedForAll(opts *bind.CallOpts , owner common.Address , operator common.Address ) (bool, error) {
			var (
				ret0 = new(bool)
				
			)
			out := ret0
			err := _ERC721.contract.Call(opts, out, "isApprovedForAll" , owner, operator)
			return *ret0, err
		}

		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_ERC721 *ERC721Session) IsApprovedForAll( owner common.Address , operator common.Address ) ( bool,  error) {
		  return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts , owner, operator)
		}

		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_ERC721 *ERC721CallerSession) IsApprovedForAll( owner common.Address , operator common.Address ) ( bool,  error) {
		  return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts , owner, operator)
		}
	
		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address)
		func (_ERC721 *ERC721Caller) OwnerOf(opts *bind.CallOpts , tokenId *big.Int ) (common.Address, error) {
			var (
				ret0 = new(common.Address)
				
			)
			out := ret0
			err := _ERC721.contract.Call(opts, out, "ownerOf" , tokenId)
			return *ret0, err
		}

		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address)
		func (_ERC721 *ERC721Session) OwnerOf( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts , tokenId)
		}

		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address)
		func (_ERC721 *ERC721CallerSession) OwnerOf( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts , tokenId)
		}
	
		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_ERC721 *ERC721Caller) SupportsInterface(opts *bind.CallOpts , interfaceId [4]byte ) (bool, error) {
			var (
				ret0 = new(bool)
				
			)
			out := ret0
			err := _ERC721.contract.Call(opts, out, "supportsInterface" , interfaceId)
			return *ret0, err
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_ERC721 *ERC721Session) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts , interfaceId)
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_ERC721 *ERC721CallerSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts , interfaceId)
		}
	

	
		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_ERC721 *ERC721Transactor) Approve(opts *bind.TransactOpts , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _ERC721.contract.Transact(opts, "approve" , to, tokenId)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_ERC721 *ERC721Session) Approve( to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721.Contract.Approve(&_ERC721.TransactOpts , to, tokenId)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_ERC721 *ERC721TransactorSession) Approve( to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721.Contract.Approve(&_ERC721.TransactOpts , to, tokenId)
		}
	
		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721 *ERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _ERC721.contract.Transact(opts, "safeTransferFrom" , from, to, tokenId)
		}

		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721 *ERC721Session) SafeTransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts , from, to, tokenId)
		}

		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721 *ERC721TransactorSession) SafeTransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts , from, to, tokenId)
		}
	
		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
		func (_ERC721 *ERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int , _data []byte ) (*types.Transaction, error) {
			return _ERC721.contract.Transact(opts, "safeTransferFrom0" , from, to, tokenId, _data)
		}

		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
		func (_ERC721 *ERC721Session) SafeTransferFrom0( from common.Address , to common.Address , tokenId *big.Int , _data []byte ) (*types.Transaction, error) {
		  return _ERC721.Contract.SafeTransferFrom0(&_ERC721.TransactOpts , from, to, tokenId, _data)
		}

		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
		func (_ERC721 *ERC721TransactorSession) SafeTransferFrom0( from common.Address , to common.Address , tokenId *big.Int , _data []byte ) (*types.Transaction, error) {
		  return _ERC721.Contract.SafeTransferFrom0(&_ERC721.TransactOpts , from, to, tokenId, _data)
		}
	
		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address to, bool approved) returns()
		func (_ERC721 *ERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts , to common.Address , approved bool ) (*types.Transaction, error) {
			return _ERC721.contract.Transact(opts, "setApprovalForAll" , to, approved)
		}

		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address to, bool approved) returns()
		func (_ERC721 *ERC721Session) SetApprovalForAll( to common.Address , approved bool ) (*types.Transaction, error) {
		  return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts , to, approved)
		}

		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address to, bool approved) returns()
		func (_ERC721 *ERC721TransactorSession) SetApprovalForAll( to common.Address , approved bool ) (*types.Transaction, error) {
		  return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts , to, approved)
		}
	
		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721 *ERC721Transactor) TransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _ERC721.contract.Transact(opts, "transferFrom" , from, to, tokenId)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721 *ERC721Session) TransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts , from, to, tokenId)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721 *ERC721TransactorSession) TransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts , from, to, tokenId)
		}
	

	

	

	
		// ERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721 contract.
		type ERC721ApprovalIterator struct {
			Event *ERC721Approval // Event containing the contract specifics and raw log

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
		func (it *ERC721ApprovalIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ERC721Approval)
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
				it.Event = new(ERC721Approval)
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
		func (it *ERC721ApprovalIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ERC721ApprovalIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ERC721Approval represents a Approval event raised by the ERC721 contract.
		type ERC721Approval struct { 
			Owner common.Address; 
			Approved common.Address; 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
 		func (_ERC721 *ERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721ApprovalIterator, error) {
			
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

			logs, sub, err := _ERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &ERC721ApprovalIterator{contract: _ERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
 		}

		// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
		func (_ERC721 *ERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {
			
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

			logs, sub, err := _ERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ERC721Approval)
						if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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
		func (_ERC721 *ERC721Filterer) ParseApproval(log types.Log) (*ERC721Approval, error) {
			event := new(ERC721Approval)
			if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// ERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721 contract.
		type ERC721ApprovalForAllIterator struct {
			Event *ERC721ApprovalForAll // Event containing the contract specifics and raw log

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
		func (it *ERC721ApprovalForAllIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ERC721ApprovalForAll)
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
				it.Event = new(ERC721ApprovalForAll)
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
		func (it *ERC721ApprovalForAllIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ERC721ApprovalForAllIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ERC721ApprovalForAll represents a ApprovalForAll event raised by the ERC721 contract.
		type ERC721ApprovalForAll struct { 
			Owner common.Address; 
			Operator common.Address; 
			Approved bool; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
 		func (_ERC721 *ERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721ApprovalForAllIterator, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var operatorRule []interface{}
			for _, operatorItem := range operator {
				operatorRule = append(operatorRule, operatorItem)
			}
			

			logs, sub, err := _ERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
			if err != nil {
				return nil, err
			}
			return &ERC721ApprovalForAllIterator{contract: _ERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
 		}

		// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
		func (_ERC721 *ERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var operatorRule []interface{}
			for _, operatorItem := range operator {
				operatorRule = append(operatorRule, operatorItem)
			}
			

			logs, sub, err := _ERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ERC721ApprovalForAll)
						if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
		func (_ERC721 *ERC721Filterer) ParseApprovalForAll(log types.Log) (*ERC721ApprovalForAll, error) {
			event := new(ERC721ApprovalForAll)
			if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// ERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721 contract.
		type ERC721TransferIterator struct {
			Event *ERC721Transfer // Event containing the contract specifics and raw log

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
		func (it *ERC721TransferIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ERC721Transfer)
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
				it.Event = new(ERC721Transfer)
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
		func (it *ERC721TransferIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ERC721TransferIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ERC721Transfer represents a Transfer event raised by the ERC721 contract.
		type ERC721Transfer struct { 
			From common.Address; 
			To common.Address; 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
 		func (_ERC721 *ERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721TransferIterator, error) {
			
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

			logs, sub, err := _ERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &ERC721TransferIterator{contract: _ERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
 		}

		// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
		func (_ERC721 *ERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {
			
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

			logs, sub, err := _ERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ERC721Transfer)
						if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
		func (_ERC721 *ERC721Filterer) ParseTransfer(log types.Log) (*ERC721Transfer, error) {
			event := new(ERC721Transfer)
			if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	

	// ERC721BaseABI is the input ABI used to generate the binding from.
	const ERC721BaseABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"memoAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenProperty\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenUri\",\"type\":\"string\"}],\"name\":\"tokenMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"transactionMemo\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"addMemo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getProperty\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenProperty\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenUri\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"property\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenProperty\",\"type\":\"string\"}],\"name\":\"setProperty\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"memo\",\"type\":\"string\"}],\"name\":\"transferWithMemo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

	
		// ERC721BaseFuncSigs maps the 4-byte function signature to its string representation.
		var ERC721BaseFuncSigs = map[string]string{
			"f93e4ec8": "addMemo(uint256,string)",
			"095ea7b3": "approve(address,uint256)",
			"70a08231": "balanceOf(address)",
			"6c0360eb": "baseURI()",
			"081812fc": "getApproved(uint256)",
			"32665ffb": "getProperty(uint256)",
			"e985e9c5": "isApprovedForAll(address,address)",
			"2825df7b": "mint(uint256,string,string)",
			"06fdde03": "name()",
			"150b7a02": "onERC721Received(address,address,uint256,bytes)",
			"6352211e": "ownerOf(uint256)",
			"30e1ad32": "property(uint256)",
			"42842e0e": "safeTransferFrom(address,address,uint256)",
			"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
			"a22cb465": "setApprovalForAll(address,bool)",
			"077419f3": "setProperty(uint256,string)",
			"01ffc9a7": "supportsInterface(bytes4)",
			"95d89b41": "symbol()",
			"4f6ccce7": "tokenByIndex(uint256)",
			"2f745c59": "tokenOfOwnerByIndex(address,uint256)",
			"c87b56dd": "tokenURI(uint256)",
			"18160ddd": "totalSupply()",
			"23b872dd": "transferFrom(address,address,uint256)",
			"dcb49c73": "transferWithMemo(address,address,uint256,string)",
			
		}
	

	
		// ERC721BaseBin is the compiled bytecode used for deploying new contracts.
		var ERC721BaseBin = "0x60806040523480156200001157600080fd5b506040516200276c3803806200276c833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b50604052508391508290508181620001c06301ffc9a760e01b6001600160e01b036200024916565b620001db6380ac58cd60e01b6001600160e01b036200024916565b620001f663780e9d6360e01b6001600160e01b036200024916565b81516200020b906009906020850190620002ce565b5080516200022190600a906020840190620002ce565b506200023d635b5e139f60e01b6001600160e01b036200024916565b50505050505062000373565b6001600160e01b03198082161415620002a9576040805162461bcd60e51b815260206004820152601c60248201527f4552433136353a20696e76616c696420696e7465726661636520696400000000604482015290519081900360640190fd5b6001600160e01b0319166000908152602081905260409020805460ff19166001179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200031157805160ff191683800117855562000341565b8280016001018555821562000341579182015b828111156200034157825182559160200191906001019062000324565b506200034f92915062000353565b5090565b6200037091905b808211156200034f57600081556001016200035a565b90565b6123e980620003836000396000f3fe60806040526004361061014b5760003560e01c806342842e0e116100b6578063a22cb4651161006f578063a22cb4651461076c578063b88d4fde146107a7578063c87b56dd14610878578063dcb49c73146108a2578063e985e9c514610973578063f93e4ec8146109ae5761014b565b806342842e0e146106785780634f6ccce7146106bb5780636352211e146106e55780636c0360eb1461070f57806370a082311461072457806395d89b41146107575761014b565b806318160ddd1161010857806318160ddd1461044457806323b872dd1461046b5780632825df7b146104ae5780632f745c59146105eb57806330e1ad321461062457806332665ffb1461064e5761014b565b806301ffc9a71461014d57806306fdde0314610195578063077419f31461021f578063081812fc146102d7578063095ea7b31461031d578063150b7a0214610356575b005b34801561015957600080fd5b506101816004803603602081101561017057600080fd5b50356001600160e01b031916610a66565b604080519115158252519081900360200190f35b3480156101a157600080fd5b506101aa610a89565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101e45781810151838201526020016101cc565b50505050905090810190601f1680156102115780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561022b57600080fd5b5061014b6004803603604081101561024257600080fd5b81359190810190604081016020820135600160201b81111561026357600080fd5b82018360208201111561027557600080fd5b803590602001918460018302840111600160201b8311171561029657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610b20945050505050565b3480156102e357600080fd5b50610301600480360360208110156102fa57600080fd5b5035610ba5565b604080516001600160a01b039092168252519081900360200190f35b34801561032957600080fd5b5061014b6004803603604081101561034057600080fd5b506001600160a01b038135169060200135610c07565b34801561036257600080fd5b506104276004803603608081101561037957600080fd5b6001600160a01b03823581169260208101359091169160408201359190810190608081016060820135600160201b8111156103b357600080fd5b8201836020820111156103c557600080fd5b803590602001918460018302840111600160201b831117156103e657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610d2f945050505050565b604080516001600160e01b03199092168252519081900360200190f35b34801561045057600080fd5b50610459610d40565b60408051918252519081900360200190f35b34801561047757600080fd5b5061014b6004803603606081101561048e57600080fd5b506001600160a01b03813581169160208101359091169060400135610d46565b3480156104ba57600080fd5b5061014b600480360360608110156104d157600080fd5b81359190810190604081016020820135600160201b8111156104f257600080fd5b82018360208201111561050457600080fd5b803590602001918460018302840111600160201b8311171561052557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561057757600080fd5b82018360208201111561058957600080fd5b803590602001918460018302840111600160201b831117156105aa57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610da2945050505050565b3480156105f757600080fd5b506104596004803603604081101561060e57600080fd5b506001600160a01b038135169060200135610edc565b34801561063057600080fd5b506101aa6004803603602081101561064757600080fd5b5035610f5b565b34801561065a57600080fd5b506101aa6004803603602081101561067157600080fd5b5035610ff6565b34801561068457600080fd5b5061014b6004803603606081101561069b57600080fd5b506001600160a01b03813581169160208101359091169060400135611097565b3480156106c757600080fd5b50610459600480360360208110156106de57600080fd5b50356110b2565b3480156106f157600080fd5b506103016004803603602081101561070857600080fd5b5035611118565b34801561071b57600080fd5b506101aa611172565b34801561073057600080fd5b506104596004803603602081101561074757600080fd5b50356001600160a01b03166111d3565b34801561076357600080fd5b506101aa61123b565b34801561077857600080fd5b5061014b6004803603604081101561078f57600080fd5b506001600160a01b038135169060200135151561129c565b3480156107b357600080fd5b5061014b600480360360808110156107ca57600080fd5b6001600160a01b03823581169260208101359091169160408201359190810190608081016060820135600160201b81111561080457600080fd5b82018360208201111561081657600080fd5b803590602001918460018302840111600160201b8311171561083757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506113a1945050505050565b34801561088457600080fd5b506101aa6004803603602081101561089b57600080fd5b50356113f9565b3480156108ae57600080fd5b5061014b600480360360808110156108c557600080fd5b6001600160a01b03823581169260208101359091169160408201359190810190608081016060820135600160201b8111156108ff57600080fd5b82018360208201111561091157600080fd5b803590602001918460018302840111600160201b8311171561093257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506115c5945050505050565b34801561097f57600080fd5b506101816004803603604081101561099657600080fd5b506001600160a01b0381358116916020013516611677565b3480156109ba57600080fd5b5061014b600480360360408110156109d157600080fd5b81359190810190604081016020820135600160201b8111156109f257600080fd5b820183602082011115610a0457600080fd5b803590602001918460018302840111600160201b83111715610a2557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506116a5945050505050565b6001600160e01b0319811660009081526020819052604090205460ff165b919050565b60098054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610b155780601f10610aea57610100808354040283529160200191610b15565b820191906000526020600020905b815481529060010190602001808311610af857829003601f168201915b505050505090505b90565b6000610b2b83611118565b9050336001600160a01b03821614610b80576040805162461bcd60e51b815260206004820152601360248201527237b7363c9031b7b73a3930b1ba1037bbb732b960691b604482015290519081900360640190fd5b6000838152600d602090815260409091208351610b9f92850190612096565b50505050565b6000610bb082611743565b610beb5760405162461bcd60e51b815260040180806020018281038252602c815260200180612287602c913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b6000610c1282611118565b9050806001600160a01b0316836001600160a01b03161415610c655760405162461bcd60e51b81526004018080602001828103825260218152602001806123376021913960400191505060405180910390fd5b806001600160a01b0316610c77611760565b6001600160a01b03161480610c985750610c9881610c93611760565b611677565b610cd35760405162461bcd60e51b81526004018080602001828103825260388152602001806121fc6038913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b630a85bd0160e11b5b949350505050565b60075490565b610d57610d51611760565b82611764565b610d925760405162461bcd60e51b81526004018080602001828103825260318152602001806123586031913960400191505060405180910390fd5b610d9d838383611800565b505050565b610dac338461181f565b6000838152600d602090815260409091208351610dcb92850190612096565b50610dd68382611840565b827f20dc370201d8359018b0db6be114390d5374d7af7458399ed9cba5ffee738ec68383604051808060200180602001838103835285818151815260200191508051906020019080838360005b83811015610e3b578181015183820152602001610e23565b50505050905090810190601f168015610e685780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610e9b578181015183820152602001610e83565b50505050905090810190601f168015610ec85780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a2505050565b6000610ee7836111d3565b8210610f245760405162461bcd60e51b815260040180806020018281038252602b81526020018061214f602b913960400191505060405180910390fd5b6001600160a01b0383166000908152600560205260409020805483908110610f4857fe5b9060005260206000200154905092915050565b600d6020908152600091825260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015610fee5780601f10610fc357610100808354040283529160200191610fee565b820191906000526020600020905b815481529060010190602001808311610fd157829003601f168201915b505050505081565b6000818152600d602090815260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084526060939283018282801561108b5780601f106110605761010080835404028352916020019161108b565b820191906000526020600020905b81548152906001019060200180831161106e57829003601f168201915b50505050509050919050565b610d9d838383604051806020016040528060008152506113a1565b60006110bc610d40565b82106110f95760405162461bcd60e51b815260040180806020018281038252602c815260200180612389602c913960400191505060405180910390fd5b6007828154811061110657fe5b90600052602060002001549050919050565b6000818152600160205260408120546001600160a01b03168061116c5760405162461bcd60e51b815260040180806020018281038252602981526020018061225e6029913960400191505060405180910390fd5b92915050565b600b8054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610b155780601f10610aea57610100808354040283529160200191610b15565b60006001600160a01b03821661121a5760405162461bcd60e51b815260040180806020018281038252602a815260200180612234602a913960400191505060405180910390fd5b6001600160a01b038216600090815260036020526040902061116c906118a3565b600a8054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610b155780601f10610aea57610100808354040283529160200191610b15565b6112a4611760565b6001600160a01b0316826001600160a01b0316141561130a576040805162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015290519081900360640190fd5b8060046000611317611760565b6001600160a01b03908116825260208083019390935260409182016000908120918716808252919093529120805460ff19169215159290921790915561135b611760565b60408051841515815290516001600160a01b0392909216917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c319181900360200190a35050565b6113b26113ac611760565b83611764565b6113ed5760405162461bcd60e51b81526004018080602001828103825260318152602001806123586031913960400191505060405180910390fd5b610b9f848484846118a7565b606061140482611743565b61143f5760405162461bcd60e51b815260040180806020018281038252602f815260200180612308602f913960400191505060405180910390fd5b6000828152600c602090815260409182902080548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452606093928301828280156114d45780601f106114a9576101008083540402835291602001916114d4565b820191906000526020600020905b8154815290600101906020018083116114b757829003601f168201915b505050505090508051600014156114fb575050604080516020810190915260008152610a84565b600b81604051602001808380546001816001161561010002031660029004801561155c5780601f1061153a57610100808354040283529182019161155c565b820191906000526020600020905b815481529060010190602001808311611548575b5050825160208401908083835b602083106115885780518252601f199092019160209182019101611569565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052915050610a84565b6115d0848484610d46565b805115610b9f57817f6bb79e977ebcae518f2d094cc741d1a77f05ae1d52cd8e23e5b7118572a438ce826040518080602001828103825283818151815260200191508051906020019080838360005b8381101561163757818101518382015260200161161f565b50505050905090810190601f1680156116645780820380516001836020036101000a031916815260200191505b509250505060405180910390a250505050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b817f4a62fa304d9a5993f36546283e9d483bc1cb4909a24a0bb0104fb47505caea20826040518080602001828103825283818151815260200191508051906020019080838360005b838110156117055781810151838201526020016116ed565b50505050905090810190601f1680156117325780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b6000908152600160205260409020546001600160a01b0316151590565b3390565b600061176f82611743565b6117aa5760405162461bcd60e51b815260040180806020018281038252602c8152602001806121d0602c913960400191505060405180910390fd5b60006117b583611118565b9050806001600160a01b0316846001600160a01b031614806117f05750836001600160a01b03166117e584610ba5565b6001600160a01b0316145b80610d385750610d388185611677565b61180b8383836118f9565b6118158382611a3d565b610d9d8282611b32565b6118298282611b70565b6118338282611b32565b61183c81611ca1565b5050565b61184982611743565b6118845760405162461bcd60e51b815260040180806020018281038252602c8152602001806122b3602c913960400191505060405180910390fd5b6000828152600c602090815260409091208251610d9d92840190612096565b5490565b6118b2848484611800565b6118be84848484611ce5565b610b9f5760405162461bcd60e51b815260040180806020018281038252603281526020018061217a6032913960400191505060405180910390fd5b826001600160a01b031661190c82611118565b6001600160a01b0316146119515760405162461bcd60e51b81526004018080602001828103825260298152602001806122df6029913960400191505060405180910390fd5b6001600160a01b0382166119965760405162461bcd60e51b81526004018080602001828103825260248152602001806121ac6024913960400191505060405180910390fd5b61199f81611f20565b6001600160a01b03831660009081526003602052604090206119c090611f5d565b6001600160a01b03821660009081526003602052604090206119e190611f74565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6001600160a01b038216600090815260056020526040812054611a6790600163ffffffff611f7d16565b600083815260066020526040902054909150808214611b02576001600160a01b0384166000908152600560205260408120805484908110611aa457fe5b906000526020600020015490508060056000876001600160a01b03166001600160a01b031681526020019081526020016000208381548110611ae257fe5b600091825260208083209091019290925591825260069052604090208190555b6001600160a01b0384166000908152600560205260409020805490611b2b906000198301612114565b5050505050565b6001600160a01b0390911660009081526005602081815260408084208054868652600684529185208290559282526001810183559183529091200155565b6001600160a01b038216611bcb576040805162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f2061646472657373604482015290519081900360640190fd5b611bd481611743565b15611c26576040805162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000604482015290519081900360640190fd5b600081815260016020908152604080832080546001600160a01b0319166001600160a01b038716908117909155835260039091529020611c6590611f74565b60405181906001600160a01b038416906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b600780546000838152600860205260408120829055600182018355919091527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880155565b6000611cf9846001600160a01b0316611fc6565b611d0557506001610d38565b600060606001600160a01b038616630a85bd0160e11b611d23611760565b89888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611d9c578181015183820152602001611d84565b50505050905090810190601f168015611dc95780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310611e315780518252601f199092019160209182019101611e12565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611e93576040519150601f19603f3d011682016040523d82523d6000602084013e611e98565b606091505b509150915081611ee957805115611eb25780518082602001fd5b60405162461bcd60e51b815260040180806020018281038252603281526020018061217a6032913960400191505060405180910390fd5b6000818060200190516020811015611f0057600080fd5b50516001600160e01b031916630a85bd0160e11b149350610d3892505050565b6000818152600260205260409020546001600160a01b031615611f5a57600081815260026020526040902080546001600160a01b03191690555b50565b8054611f7090600163ffffffff611f7d16565b9055565b80546001019055565b6000611fbf83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611fff565b9392505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470818114801590610d38575050151592915050565b6000818484111561208e5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561205357818101518382015260200161203b565b50505050905090810190601f1680156120805780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106120d757805160ff1916838001178555612104565b82800160010185558215612104579182015b828111156121045782518255916020019190600101906120e9565b50612110929150612134565b5090565b815481835581811115610d9d57600083815260209020610d9d9181019083015b610b1d91905b80821115612110576000815560010161213a56fe455243373231456e756d657261626c653a206f776e657220696e646578206f7574206f6620626f756e64734552433732313a207472616e7366657220746f206e6f6e20455243373231526563656976657220696d706c656d656e7465724552433732313a207472616e7366657220746f20746865207a65726f20616464726573734552433732313a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4552433732313a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734552433732313a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732314d657461646174613a2055524920736574206f66206e6f6e6578697374656e7420746f6b656e4552433732313a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4552433732314d657461646174613a2055524920717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76616c20746f2063757272656e74206f776e65724552433732313a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564455243373231456e756d657261626c653a20676c6f62616c20696e646578206f7574206f6620626f756e6473a265627a7a72315820aa65ab1cb5a8e959249c68836bfa609d8bb99d95230db4a62e1d89679eefaab764736f6c63430005110032"

		// DeployERC721Base deploys a new Ethereum contract, binding an instance of ERC721Base to it.
		func DeployERC721Base(auth *bind.TransactOpts, backend bind.ContractBackend , name string, symbol string) (common.Address, *types.Transaction, *ERC721Base, error) {
		  parsed, err := abi.JSON(strings.NewReader(ERC721BaseABI))
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  
		  address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721BaseBin), backend , name, symbol)
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  return address, tx, &ERC721Base{ ERC721BaseCaller: ERC721BaseCaller{contract: contract}, ERC721BaseTransactor: ERC721BaseTransactor{contract: contract}, ERC721BaseFilterer: ERC721BaseFilterer{contract: contract} }, nil
		}
	

	// ERC721Base is an auto generated Go binding around an Ethereum contract.
	type ERC721Base struct {
	  ERC721BaseCaller     // Read-only binding to the contract
	  ERC721BaseTransactor // Write-only binding to the contract
	  ERC721BaseFilterer   // Log filterer for contract events
	}

	// ERC721BaseCaller is an auto generated read-only Go binding around an Ethereum contract.
	type ERC721BaseCaller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC721BaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
	type ERC721BaseTransactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC721BaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type ERC721BaseFilterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC721BaseSession is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type ERC721BaseSession struct {
	  Contract     *ERC721Base        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// ERC721BaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type ERC721BaseCallerSession struct {
	  Contract *ERC721BaseCaller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// ERC721BaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type ERC721BaseTransactorSession struct {
	  Contract     *ERC721BaseTransactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// ERC721BaseRaw is an auto generated low-level Go binding around an Ethereum contract.
	type ERC721BaseRaw struct {
	  Contract *ERC721Base // Generic contract binding to access the raw methods on
	}

	// ERC721BaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type ERC721BaseCallerRaw struct {
		Contract *ERC721BaseCaller // Generic read-only contract binding to access the raw methods on
	}

	// ERC721BaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type ERC721BaseTransactorRaw struct {
		Contract *ERC721BaseTransactor // Generic write-only contract binding to access the raw methods on
	}

	// NewERC721Base creates a new instance of ERC721Base, bound to a specific deployed contract.
	func NewERC721Base(address common.Address, backend bind.ContractBackend) (*ERC721Base, error) {
	  contract, err := bindERC721Base(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC721Base{ ERC721BaseCaller: ERC721BaseCaller{contract: contract}, ERC721BaseTransactor: ERC721BaseTransactor{contract: contract}, ERC721BaseFilterer: ERC721BaseFilterer{contract: contract} }, nil
	}

	// NewERC721BaseCaller creates a new read-only instance of ERC721Base, bound to a specific deployed contract.
	func NewERC721BaseCaller(address common.Address, caller bind.ContractCaller) (*ERC721BaseCaller, error) {
	  contract, err := bindERC721Base(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC721BaseCaller{contract: contract}, nil
	}

	// NewERC721BaseTransactor creates a new write-only instance of ERC721Base, bound to a specific deployed contract.
	func NewERC721BaseTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721BaseTransactor, error) {
	  contract, err := bindERC721Base(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC721BaseTransactor{contract: contract}, nil
	}

	// NewERC721BaseFilterer creates a new log filterer instance of ERC721Base, bound to a specific deployed contract.
 	func NewERC721BaseFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721BaseFilterer, error) {
 	  contract, err := bindERC721Base(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &ERC721BaseFilterer{contract: contract}, nil
 	}

	// bindERC721Base binds a generic wrapper to an already deployed contract.
	func bindERC721Base(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(ERC721BaseABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_ERC721Base *ERC721BaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _ERC721Base.Contract.ERC721BaseCaller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_ERC721Base *ERC721BaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _ERC721Base.Contract.ERC721BaseTransactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_ERC721Base *ERC721BaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _ERC721Base.Contract.ERC721BaseTransactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_ERC721Base *ERC721BaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _ERC721Base.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_ERC721Base *ERC721BaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _ERC721Base.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_ERC721Base *ERC721BaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _ERC721Base.Contract.contract.Transact(opts, method, params...)
	}

	
		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256)
		func (_ERC721Base *ERC721BaseCaller) BalanceOf(opts *bind.CallOpts , owner common.Address ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _ERC721Base.contract.Call(opts, out, "balanceOf" , owner)
			return *ret0, err
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256)
		func (_ERC721Base *ERC721BaseSession) BalanceOf( owner common.Address ) ( *big.Int,  error) {
		  return _ERC721Base.Contract.BalanceOf(&_ERC721Base.CallOpts , owner)
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256)
		func (_ERC721Base *ERC721BaseCallerSession) BalanceOf( owner common.Address ) ( *big.Int,  error) {
		  return _ERC721Base.Contract.BalanceOf(&_ERC721Base.CallOpts , owner)
		}
	
		// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
		//
		// Solidity: function baseURI() view returns(string)
		func (_ERC721Base *ERC721BaseCaller) BaseURI(opts *bind.CallOpts ) (string, error) {
			var (
				ret0 = new(string)
				
			)
			out := ret0
			err := _ERC721Base.contract.Call(opts, out, "baseURI" )
			return *ret0, err
		}

		// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
		//
		// Solidity: function baseURI() view returns(string)
		func (_ERC721Base *ERC721BaseSession) BaseURI() ( string,  error) {
		  return _ERC721Base.Contract.BaseURI(&_ERC721Base.CallOpts )
		}

		// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
		//
		// Solidity: function baseURI() view returns(string)
		func (_ERC721Base *ERC721BaseCallerSession) BaseURI() ( string,  error) {
		  return _ERC721Base.Contract.BaseURI(&_ERC721Base.CallOpts )
		}
	
		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address)
		func (_ERC721Base *ERC721BaseCaller) GetApproved(opts *bind.CallOpts , tokenId *big.Int ) (common.Address, error) {
			var (
				ret0 = new(common.Address)
				
			)
			out := ret0
			err := _ERC721Base.contract.Call(opts, out, "getApproved" , tokenId)
			return *ret0, err
		}

		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address)
		func (_ERC721Base *ERC721BaseSession) GetApproved( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721Base.Contract.GetApproved(&_ERC721Base.CallOpts , tokenId)
		}

		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address)
		func (_ERC721Base *ERC721BaseCallerSession) GetApproved( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721Base.Contract.GetApproved(&_ERC721Base.CallOpts , tokenId)
		}
	
		// GetProperty is a free data retrieval call binding the contract method 0x32665ffb.
		//
		// Solidity: function getProperty(uint256 tokenId) view returns(string)
		func (_ERC721Base *ERC721BaseCaller) GetProperty(opts *bind.CallOpts , tokenId *big.Int ) (string, error) {
			var (
				ret0 = new(string)
				
			)
			out := ret0
			err := _ERC721Base.contract.Call(opts, out, "getProperty" , tokenId)
			return *ret0, err
		}

		// GetProperty is a free data retrieval call binding the contract method 0x32665ffb.
		//
		// Solidity: function getProperty(uint256 tokenId) view returns(string)
		func (_ERC721Base *ERC721BaseSession) GetProperty( tokenId *big.Int ) ( string,  error) {
		  return _ERC721Base.Contract.GetProperty(&_ERC721Base.CallOpts , tokenId)
		}

		// GetProperty is a free data retrieval call binding the contract method 0x32665ffb.
		//
		// Solidity: function getProperty(uint256 tokenId) view returns(string)
		func (_ERC721Base *ERC721BaseCallerSession) GetProperty( tokenId *big.Int ) ( string,  error) {
		  return _ERC721Base.Contract.GetProperty(&_ERC721Base.CallOpts , tokenId)
		}
	
		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_ERC721Base *ERC721BaseCaller) IsApprovedForAll(opts *bind.CallOpts , owner common.Address , operator common.Address ) (bool, error) {
			var (
				ret0 = new(bool)
				
			)
			out := ret0
			err := _ERC721Base.contract.Call(opts, out, "isApprovedForAll" , owner, operator)
			return *ret0, err
		}

		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_ERC721Base *ERC721BaseSession) IsApprovedForAll( owner common.Address , operator common.Address ) ( bool,  error) {
		  return _ERC721Base.Contract.IsApprovedForAll(&_ERC721Base.CallOpts , owner, operator)
		}

		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_ERC721Base *ERC721BaseCallerSession) IsApprovedForAll( owner common.Address , operator common.Address ) ( bool,  error) {
		  return _ERC721Base.Contract.IsApprovedForAll(&_ERC721Base.CallOpts , owner, operator)
		}
	
		// Name is a free data retrieval call binding the contract method 0x06fdde03.
		//
		// Solidity: function name() view returns(string)
		func (_ERC721Base *ERC721BaseCaller) Name(opts *bind.CallOpts ) (string, error) {
			var (
				ret0 = new(string)
				
			)
			out := ret0
			err := _ERC721Base.contract.Call(opts, out, "name" )
			return *ret0, err
		}

		// Name is a free data retrieval call binding the contract method 0x06fdde03.
		//
		// Solidity: function name() view returns(string)
		func (_ERC721Base *ERC721BaseSession) Name() ( string,  error) {
		  return _ERC721Base.Contract.Name(&_ERC721Base.CallOpts )
		}

		// Name is a free data retrieval call binding the contract method 0x06fdde03.
		//
		// Solidity: function name() view returns(string)
		func (_ERC721Base *ERC721BaseCallerSession) Name() ( string,  error) {
		  return _ERC721Base.Contract.Name(&_ERC721Base.CallOpts )
		}
	
		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address)
		func (_ERC721Base *ERC721BaseCaller) OwnerOf(opts *bind.CallOpts , tokenId *big.Int ) (common.Address, error) {
			var (
				ret0 = new(common.Address)
				
			)
			out := ret0
			err := _ERC721Base.contract.Call(opts, out, "ownerOf" , tokenId)
			return *ret0, err
		}

		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address)
		func (_ERC721Base *ERC721BaseSession) OwnerOf( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721Base.Contract.OwnerOf(&_ERC721Base.CallOpts , tokenId)
		}

		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address)
		func (_ERC721Base *ERC721BaseCallerSession) OwnerOf( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721Base.Contract.OwnerOf(&_ERC721Base.CallOpts , tokenId)
		}
	
		// Property is a free data retrieval call binding the contract method 0x30e1ad32.
		//
		// Solidity: function property(uint256 ) view returns(string)
		func (_ERC721Base *ERC721BaseCaller) Property(opts *bind.CallOpts , arg0 *big.Int ) (string, error) {
			var (
				ret0 = new(string)
				
			)
			out := ret0
			err := _ERC721Base.contract.Call(opts, out, "property" , arg0)
			return *ret0, err
		}

		// Property is a free data retrieval call binding the contract method 0x30e1ad32.
		//
		// Solidity: function property(uint256 ) view returns(string)
		func (_ERC721Base *ERC721BaseSession) Property( arg0 *big.Int ) ( string,  error) {
		  return _ERC721Base.Contract.Property(&_ERC721Base.CallOpts , arg0)
		}

		// Property is a free data retrieval call binding the contract method 0x30e1ad32.
		//
		// Solidity: function property(uint256 ) view returns(string)
		func (_ERC721Base *ERC721BaseCallerSession) Property( arg0 *big.Int ) ( string,  error) {
		  return _ERC721Base.Contract.Property(&_ERC721Base.CallOpts , arg0)
		}
	
		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_ERC721Base *ERC721BaseCaller) SupportsInterface(opts *bind.CallOpts , interfaceId [4]byte ) (bool, error) {
			var (
				ret0 = new(bool)
				
			)
			out := ret0
			err := _ERC721Base.contract.Call(opts, out, "supportsInterface" , interfaceId)
			return *ret0, err
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_ERC721Base *ERC721BaseSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _ERC721Base.Contract.SupportsInterface(&_ERC721Base.CallOpts , interfaceId)
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_ERC721Base *ERC721BaseCallerSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _ERC721Base.Contract.SupportsInterface(&_ERC721Base.CallOpts , interfaceId)
		}
	
		// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
		//
		// Solidity: function symbol() view returns(string)
		func (_ERC721Base *ERC721BaseCaller) Symbol(opts *bind.CallOpts ) (string, error) {
			var (
				ret0 = new(string)
				
			)
			out := ret0
			err := _ERC721Base.contract.Call(opts, out, "symbol" )
			return *ret0, err
		}

		// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
		//
		// Solidity: function symbol() view returns(string)
		func (_ERC721Base *ERC721BaseSession) Symbol() ( string,  error) {
		  return _ERC721Base.Contract.Symbol(&_ERC721Base.CallOpts )
		}

		// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
		//
		// Solidity: function symbol() view returns(string)
		func (_ERC721Base *ERC721BaseCallerSession) Symbol() ( string,  error) {
		  return _ERC721Base.Contract.Symbol(&_ERC721Base.CallOpts )
		}
	
		// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
		//
		// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
		func (_ERC721Base *ERC721BaseCaller) TokenByIndex(opts *bind.CallOpts , index *big.Int ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _ERC721Base.contract.Call(opts, out, "tokenByIndex" , index)
			return *ret0, err
		}

		// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
		//
		// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
		func (_ERC721Base *ERC721BaseSession) TokenByIndex( index *big.Int ) ( *big.Int,  error) {
		  return _ERC721Base.Contract.TokenByIndex(&_ERC721Base.CallOpts , index)
		}

		// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
		//
		// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
		func (_ERC721Base *ERC721BaseCallerSession) TokenByIndex( index *big.Int ) ( *big.Int,  error) {
		  return _ERC721Base.Contract.TokenByIndex(&_ERC721Base.CallOpts , index)
		}
	
		// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
		//
		// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
		func (_ERC721Base *ERC721BaseCaller) TokenOfOwnerByIndex(opts *bind.CallOpts , owner common.Address , index *big.Int ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _ERC721Base.contract.Call(opts, out, "tokenOfOwnerByIndex" , owner, index)
			return *ret0, err
		}

		// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
		//
		// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
		func (_ERC721Base *ERC721BaseSession) TokenOfOwnerByIndex( owner common.Address , index *big.Int ) ( *big.Int,  error) {
		  return _ERC721Base.Contract.TokenOfOwnerByIndex(&_ERC721Base.CallOpts , owner, index)
		}

		// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
		//
		// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
		func (_ERC721Base *ERC721BaseCallerSession) TokenOfOwnerByIndex( owner common.Address , index *big.Int ) ( *big.Int,  error) {
		  return _ERC721Base.Contract.TokenOfOwnerByIndex(&_ERC721Base.CallOpts , owner, index)
		}
	
		// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
		//
		// Solidity: function tokenURI(uint256 tokenId) view returns(string)
		func (_ERC721Base *ERC721BaseCaller) TokenURI(opts *bind.CallOpts , tokenId *big.Int ) (string, error) {
			var (
				ret0 = new(string)
				
			)
			out := ret0
			err := _ERC721Base.contract.Call(opts, out, "tokenURI" , tokenId)
			return *ret0, err
		}

		// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
		//
		// Solidity: function tokenURI(uint256 tokenId) view returns(string)
		func (_ERC721Base *ERC721BaseSession) TokenURI( tokenId *big.Int ) ( string,  error) {
		  return _ERC721Base.Contract.TokenURI(&_ERC721Base.CallOpts , tokenId)
		}

		// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
		//
		// Solidity: function tokenURI(uint256 tokenId) view returns(string)
		func (_ERC721Base *ERC721BaseCallerSession) TokenURI( tokenId *big.Int ) ( string,  error) {
		  return _ERC721Base.Contract.TokenURI(&_ERC721Base.CallOpts , tokenId)
		}
	
		// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
		//
		// Solidity: function totalSupply() view returns(uint256)
		func (_ERC721Base *ERC721BaseCaller) TotalSupply(opts *bind.CallOpts ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _ERC721Base.contract.Call(opts, out, "totalSupply" )
			return *ret0, err
		}

		// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
		//
		// Solidity: function totalSupply() view returns(uint256)
		func (_ERC721Base *ERC721BaseSession) TotalSupply() ( *big.Int,  error) {
		  return _ERC721Base.Contract.TotalSupply(&_ERC721Base.CallOpts )
		}

		// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
		//
		// Solidity: function totalSupply() view returns(uint256)
		func (_ERC721Base *ERC721BaseCallerSession) TotalSupply() ( *big.Int,  error) {
		  return _ERC721Base.Contract.TotalSupply(&_ERC721Base.CallOpts )
		}
	

	
		// AddMemo is a paid mutator transaction binding the contract method 0xf93e4ec8.
		//
		// Solidity: function addMemo(uint256 tokenId, string memo) returns()
		func (_ERC721Base *ERC721BaseTransactor) AddMemo(opts *bind.TransactOpts , tokenId *big.Int , memo string ) (*types.Transaction, error) {
			return _ERC721Base.contract.Transact(opts, "addMemo" , tokenId, memo)
		}

		// AddMemo is a paid mutator transaction binding the contract method 0xf93e4ec8.
		//
		// Solidity: function addMemo(uint256 tokenId, string memo) returns()
		func (_ERC721Base *ERC721BaseSession) AddMemo( tokenId *big.Int , memo string ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.AddMemo(&_ERC721Base.TransactOpts , tokenId, memo)
		}

		// AddMemo is a paid mutator transaction binding the contract method 0xf93e4ec8.
		//
		// Solidity: function addMemo(uint256 tokenId, string memo) returns()
		func (_ERC721Base *ERC721BaseTransactorSession) AddMemo( tokenId *big.Int , memo string ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.AddMemo(&_ERC721Base.TransactOpts , tokenId, memo)
		}
	
		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_ERC721Base *ERC721BaseTransactor) Approve(opts *bind.TransactOpts , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _ERC721Base.contract.Transact(opts, "approve" , to, tokenId)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_ERC721Base *ERC721BaseSession) Approve( to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.Approve(&_ERC721Base.TransactOpts , to, tokenId)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_ERC721Base *ERC721BaseTransactorSession) Approve( to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.Approve(&_ERC721Base.TransactOpts , to, tokenId)
		}
	
		// Mint is a paid mutator transaction binding the contract method 0x2825df7b.
		//
		// Solidity: function mint(uint256 tokenId, string tokenProperty, string tokenUri) returns()
		func (_ERC721Base *ERC721BaseTransactor) Mint(opts *bind.TransactOpts , tokenId *big.Int , tokenProperty string , tokenUri string ) (*types.Transaction, error) {
			return _ERC721Base.contract.Transact(opts, "mint" , tokenId, tokenProperty, tokenUri)
		}

		// Mint is a paid mutator transaction binding the contract method 0x2825df7b.
		//
		// Solidity: function mint(uint256 tokenId, string tokenProperty, string tokenUri) returns()
		func (_ERC721Base *ERC721BaseSession) Mint( tokenId *big.Int , tokenProperty string , tokenUri string ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.Mint(&_ERC721Base.TransactOpts , tokenId, tokenProperty, tokenUri)
		}

		// Mint is a paid mutator transaction binding the contract method 0x2825df7b.
		//
		// Solidity: function mint(uint256 tokenId, string tokenProperty, string tokenUri) returns()
		func (_ERC721Base *ERC721BaseTransactorSession) Mint( tokenId *big.Int , tokenProperty string , tokenUri string ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.Mint(&_ERC721Base.TransactOpts , tokenId, tokenProperty, tokenUri)
		}
	
		// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
		//
		// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
		func (_ERC721Base *ERC721BaseTransactor) OnERC721Received(opts *bind.TransactOpts , arg0 common.Address , arg1 common.Address , arg2 *big.Int , arg3 []byte ) (*types.Transaction, error) {
			return _ERC721Base.contract.Transact(opts, "onERC721Received" , arg0, arg1, arg2, arg3)
		}

		// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
		//
		// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
		func (_ERC721Base *ERC721BaseSession) OnERC721Received( arg0 common.Address , arg1 common.Address , arg2 *big.Int , arg3 []byte ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.OnERC721Received(&_ERC721Base.TransactOpts , arg0, arg1, arg2, arg3)
		}

		// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
		//
		// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
		func (_ERC721Base *ERC721BaseTransactorSession) OnERC721Received( arg0 common.Address , arg1 common.Address , arg2 *big.Int , arg3 []byte ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.OnERC721Received(&_ERC721Base.TransactOpts , arg0, arg1, arg2, arg3)
		}
	
		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Base *ERC721BaseTransactor) SafeTransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _ERC721Base.contract.Transact(opts, "safeTransferFrom" , from, to, tokenId)
		}

		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Base *ERC721BaseSession) SafeTransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.SafeTransferFrom(&_ERC721Base.TransactOpts , from, to, tokenId)
		}

		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Base *ERC721BaseTransactorSession) SafeTransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.SafeTransferFrom(&_ERC721Base.TransactOpts , from, to, tokenId)
		}
	
		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
		func (_ERC721Base *ERC721BaseTransactor) SafeTransferFrom0(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int , _data []byte ) (*types.Transaction, error) {
			return _ERC721Base.contract.Transact(opts, "safeTransferFrom0" , from, to, tokenId, _data)
		}

		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
		func (_ERC721Base *ERC721BaseSession) SafeTransferFrom0( from common.Address , to common.Address , tokenId *big.Int , _data []byte ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.SafeTransferFrom0(&_ERC721Base.TransactOpts , from, to, tokenId, _data)
		}

		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
		func (_ERC721Base *ERC721BaseTransactorSession) SafeTransferFrom0( from common.Address , to common.Address , tokenId *big.Int , _data []byte ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.SafeTransferFrom0(&_ERC721Base.TransactOpts , from, to, tokenId, _data)
		}
	
		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address to, bool approved) returns()
		func (_ERC721Base *ERC721BaseTransactor) SetApprovalForAll(opts *bind.TransactOpts , to common.Address , approved bool ) (*types.Transaction, error) {
			return _ERC721Base.contract.Transact(opts, "setApprovalForAll" , to, approved)
		}

		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address to, bool approved) returns()
		func (_ERC721Base *ERC721BaseSession) SetApprovalForAll( to common.Address , approved bool ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.SetApprovalForAll(&_ERC721Base.TransactOpts , to, approved)
		}

		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address to, bool approved) returns()
		func (_ERC721Base *ERC721BaseTransactorSession) SetApprovalForAll( to common.Address , approved bool ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.SetApprovalForAll(&_ERC721Base.TransactOpts , to, approved)
		}
	
		// SetProperty is a paid mutator transaction binding the contract method 0x077419f3.
		//
		// Solidity: function setProperty(uint256 tokenId, string tokenProperty) returns()
		func (_ERC721Base *ERC721BaseTransactor) SetProperty(opts *bind.TransactOpts , tokenId *big.Int , tokenProperty string ) (*types.Transaction, error) {
			return _ERC721Base.contract.Transact(opts, "setProperty" , tokenId, tokenProperty)
		}

		// SetProperty is a paid mutator transaction binding the contract method 0x077419f3.
		//
		// Solidity: function setProperty(uint256 tokenId, string tokenProperty) returns()
		func (_ERC721Base *ERC721BaseSession) SetProperty( tokenId *big.Int , tokenProperty string ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.SetProperty(&_ERC721Base.TransactOpts , tokenId, tokenProperty)
		}

		// SetProperty is a paid mutator transaction binding the contract method 0x077419f3.
		//
		// Solidity: function setProperty(uint256 tokenId, string tokenProperty) returns()
		func (_ERC721Base *ERC721BaseTransactorSession) SetProperty( tokenId *big.Int , tokenProperty string ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.SetProperty(&_ERC721Base.TransactOpts , tokenId, tokenProperty)
		}
	
		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Base *ERC721BaseTransactor) TransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _ERC721Base.contract.Transact(opts, "transferFrom" , from, to, tokenId)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Base *ERC721BaseSession) TransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.TransferFrom(&_ERC721Base.TransactOpts , from, to, tokenId)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Base *ERC721BaseTransactorSession) TransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.TransferFrom(&_ERC721Base.TransactOpts , from, to, tokenId)
		}
	
		// TransferWithMemo is a paid mutator transaction binding the contract method 0xdcb49c73.
		//
		// Solidity: function transferWithMemo(address from, address to, uint256 tokenId, string memo) returns()
		func (_ERC721Base *ERC721BaseTransactor) TransferWithMemo(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int , memo string ) (*types.Transaction, error) {
			return _ERC721Base.contract.Transact(opts, "transferWithMemo" , from, to, tokenId, memo)
		}

		// TransferWithMemo is a paid mutator transaction binding the contract method 0xdcb49c73.
		//
		// Solidity: function transferWithMemo(address from, address to, uint256 tokenId, string memo) returns()
		func (_ERC721Base *ERC721BaseSession) TransferWithMemo( from common.Address , to common.Address , tokenId *big.Int , memo string ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.TransferWithMemo(&_ERC721Base.TransactOpts , from, to, tokenId, memo)
		}

		// TransferWithMemo is a paid mutator transaction binding the contract method 0xdcb49c73.
		//
		// Solidity: function transferWithMemo(address from, address to, uint256 tokenId, string memo) returns()
		func (_ERC721Base *ERC721BaseTransactorSession) TransferWithMemo( from common.Address , to common.Address , tokenId *big.Int , memo string ) (*types.Transaction, error) {
		  return _ERC721Base.Contract.TransferWithMemo(&_ERC721Base.TransactOpts , from, to, tokenId, memo)
		}
	

	 
		// Fallback is a paid mutator transaction binding the contract fallback function.
		//
		// Solidity: fallback() payable returns()
		func (_ERC721Base *ERC721BaseTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
			return _ERC721Base.contract.RawTransact(opts, calldata)
		}

		// Fallback is a paid mutator transaction binding the contract fallback function.
		//
		// Solidity: fallback() payable returns()
		func (_ERC721Base *ERC721BaseSession) Fallback(calldata []byte) (*types.Transaction, error) {
		  return _ERC721Base.Contract.Fallback(&_ERC721Base.TransactOpts, calldata)
		}
	
		// Fallback is a paid mutator transaction binding the contract fallback function.
		// 
		// Solidity: fallback() payable returns()
		func (_ERC721Base *ERC721BaseTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
		  return _ERC721Base.Contract.Fallback(&_ERC721Base.TransactOpts, calldata)
		}
	

	

	
		// ERC721BaseApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Base contract.
		type ERC721BaseApprovalIterator struct {
			Event *ERC721BaseApproval // Event containing the contract specifics and raw log

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
		func (it *ERC721BaseApprovalIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ERC721BaseApproval)
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
				it.Event = new(ERC721BaseApproval)
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
		func (it *ERC721BaseApprovalIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ERC721BaseApprovalIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ERC721BaseApproval represents a Approval event raised by the ERC721Base contract.
		type ERC721BaseApproval struct { 
			Owner common.Address; 
			Approved common.Address; 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
 		func (_ERC721Base *ERC721BaseFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721BaseApprovalIterator, error) {
			
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

			logs, sub, err := _ERC721Base.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &ERC721BaseApprovalIterator{contract: _ERC721Base.contract, event: "Approval", logs: logs, sub: sub}, nil
 		}

		// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
		func (_ERC721Base *ERC721BaseFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721BaseApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {
			
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

			logs, sub, err := _ERC721Base.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ERC721BaseApproval)
						if err := _ERC721Base.contract.UnpackLog(event, "Approval", log); err != nil {
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
		func (_ERC721Base *ERC721BaseFilterer) ParseApproval(log types.Log) (*ERC721BaseApproval, error) {
			event := new(ERC721BaseApproval)
			if err := _ERC721Base.contract.UnpackLog(event, "Approval", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// ERC721BaseApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Base contract.
		type ERC721BaseApprovalForAllIterator struct {
			Event *ERC721BaseApprovalForAll // Event containing the contract specifics and raw log

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
		func (it *ERC721BaseApprovalForAllIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ERC721BaseApprovalForAll)
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
				it.Event = new(ERC721BaseApprovalForAll)
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
		func (it *ERC721BaseApprovalForAllIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ERC721BaseApprovalForAllIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ERC721BaseApprovalForAll represents a ApprovalForAll event raised by the ERC721Base contract.
		type ERC721BaseApprovalForAll struct { 
			Owner common.Address; 
			Operator common.Address; 
			Approved bool; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
 		func (_ERC721Base *ERC721BaseFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721BaseApprovalForAllIterator, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var operatorRule []interface{}
			for _, operatorItem := range operator {
				operatorRule = append(operatorRule, operatorItem)
			}
			

			logs, sub, err := _ERC721Base.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
			if err != nil {
				return nil, err
			}
			return &ERC721BaseApprovalForAllIterator{contract: _ERC721Base.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
 		}

		// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
		func (_ERC721Base *ERC721BaseFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721BaseApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var operatorRule []interface{}
			for _, operatorItem := range operator {
				operatorRule = append(operatorRule, operatorItem)
			}
			

			logs, sub, err := _ERC721Base.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ERC721BaseApprovalForAll)
						if err := _ERC721Base.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
		func (_ERC721Base *ERC721BaseFilterer) ParseApprovalForAll(log types.Log) (*ERC721BaseApprovalForAll, error) {
			event := new(ERC721BaseApprovalForAll)
			if err := _ERC721Base.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// ERC721BaseTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Base contract.
		type ERC721BaseTransferIterator struct {
			Event *ERC721BaseTransfer // Event containing the contract specifics and raw log

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
		func (it *ERC721BaseTransferIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ERC721BaseTransfer)
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
				it.Event = new(ERC721BaseTransfer)
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
		func (it *ERC721BaseTransferIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ERC721BaseTransferIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ERC721BaseTransfer represents a Transfer event raised by the ERC721Base contract.
		type ERC721BaseTransfer struct { 
			From common.Address; 
			To common.Address; 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
 		func (_ERC721Base *ERC721BaseFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721BaseTransferIterator, error) {
			
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

			logs, sub, err := _ERC721Base.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &ERC721BaseTransferIterator{contract: _ERC721Base.contract, event: "Transfer", logs: logs, sub: sub}, nil
 		}

		// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
		func (_ERC721Base *ERC721BaseFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721BaseTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {
			
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

			logs, sub, err := _ERC721Base.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ERC721BaseTransfer)
						if err := _ERC721Base.contract.UnpackLog(event, "Transfer", log); err != nil {
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
		func (_ERC721Base *ERC721BaseFilterer) ParseTransfer(log types.Log) (*ERC721BaseTransfer, error) {
			event := new(ERC721BaseTransfer)
			if err := _ERC721Base.contract.UnpackLog(event, "Transfer", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// ERC721BaseMemoAddedIterator is returned from FilterMemoAdded and is used to iterate over the raw logs and unpacked data for MemoAdded events raised by the ERC721Base contract.
		type ERC721BaseMemoAddedIterator struct {
			Event *ERC721BaseMemoAdded // Event containing the contract specifics and raw log

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
		func (it *ERC721BaseMemoAddedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ERC721BaseMemoAdded)
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
				it.Event = new(ERC721BaseMemoAdded)
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
		func (it *ERC721BaseMemoAddedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ERC721BaseMemoAddedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ERC721BaseMemoAdded represents a MemoAdded event raised by the ERC721Base contract.
		type ERC721BaseMemoAdded struct { 
			TokenId *big.Int; 
			Memo string; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterMemoAdded is a free log retrieval operation binding the contract event 0x4a62fa304d9a5993f36546283e9d483bc1cb4909a24a0bb0104fb47505caea20.
		//
		// Solidity: event memoAdded(uint256 indexed tokenId, string memo)
 		func (_ERC721Base *ERC721BaseFilterer) FilterMemoAdded(opts *bind.FilterOpts, tokenId []*big.Int) (*ERC721BaseMemoAddedIterator, error) {
			
			var tokenIdRule []interface{}
			for _, tokenIdItem := range tokenId {
				tokenIdRule = append(tokenIdRule, tokenIdItem)
			}
			

			logs, sub, err := _ERC721Base.contract.FilterLogs(opts, "memoAdded", tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &ERC721BaseMemoAddedIterator{contract: _ERC721Base.contract, event: "memoAdded", logs: logs, sub: sub}, nil
 		}

		// WatchMemoAdded is a free log subscription operation binding the contract event 0x4a62fa304d9a5993f36546283e9d483bc1cb4909a24a0bb0104fb47505caea20.
		//
		// Solidity: event memoAdded(uint256 indexed tokenId, string memo)
		func (_ERC721Base *ERC721BaseFilterer) WatchMemoAdded(opts *bind.WatchOpts, sink chan<- *ERC721BaseMemoAdded, tokenId []*big.Int) (event.Subscription, error) {
			
			var tokenIdRule []interface{}
			for _, tokenIdItem := range tokenId {
				tokenIdRule = append(tokenIdRule, tokenIdItem)
			}
			

			logs, sub, err := _ERC721Base.contract.WatchLogs(opts, "memoAdded", tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ERC721BaseMemoAdded)
						if err := _ERC721Base.contract.UnpackLog(event, "memoAdded", log); err != nil {
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
		func (_ERC721Base *ERC721BaseFilterer) ParseMemoAdded(log types.Log) (*ERC721BaseMemoAdded, error) {
			event := new(ERC721BaseMemoAdded)
			if err := _ERC721Base.contract.UnpackLog(event, "memoAdded", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// ERC721BaseTokenMintedIterator is returned from FilterTokenMinted and is used to iterate over the raw logs and unpacked data for TokenMinted events raised by the ERC721Base contract.
		type ERC721BaseTokenMintedIterator struct {
			Event *ERC721BaseTokenMinted // Event containing the contract specifics and raw log

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
		func (it *ERC721BaseTokenMintedIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ERC721BaseTokenMinted)
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
				it.Event = new(ERC721BaseTokenMinted)
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
		func (it *ERC721BaseTokenMintedIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ERC721BaseTokenMintedIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ERC721BaseTokenMinted represents a TokenMinted event raised by the ERC721Base contract.
		type ERC721BaseTokenMinted struct { 
			TokenId *big.Int; 
			TokenProperty string; 
			TokenUri string; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterTokenMinted is a free log retrieval operation binding the contract event 0x20dc370201d8359018b0db6be114390d5374d7af7458399ed9cba5ffee738ec6.
		//
		// Solidity: event tokenMinted(uint256 indexed tokenId, string tokenProperty, string tokenUri)
 		func (_ERC721Base *ERC721BaseFilterer) FilterTokenMinted(opts *bind.FilterOpts, tokenId []*big.Int) (*ERC721BaseTokenMintedIterator, error) {
			
			var tokenIdRule []interface{}
			for _, tokenIdItem := range tokenId {
				tokenIdRule = append(tokenIdRule, tokenIdItem)
			}
			
			

			logs, sub, err := _ERC721Base.contract.FilterLogs(opts, "tokenMinted", tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &ERC721BaseTokenMintedIterator{contract: _ERC721Base.contract, event: "tokenMinted", logs: logs, sub: sub}, nil
 		}

		// WatchTokenMinted is a free log subscription operation binding the contract event 0x20dc370201d8359018b0db6be114390d5374d7af7458399ed9cba5ffee738ec6.
		//
		// Solidity: event tokenMinted(uint256 indexed tokenId, string tokenProperty, string tokenUri)
		func (_ERC721Base *ERC721BaseFilterer) WatchTokenMinted(opts *bind.WatchOpts, sink chan<- *ERC721BaseTokenMinted, tokenId []*big.Int) (event.Subscription, error) {
			
			var tokenIdRule []interface{}
			for _, tokenIdItem := range tokenId {
				tokenIdRule = append(tokenIdRule, tokenIdItem)
			}
			
			

			logs, sub, err := _ERC721Base.contract.WatchLogs(opts, "tokenMinted", tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ERC721BaseTokenMinted)
						if err := _ERC721Base.contract.UnpackLog(event, "tokenMinted", log); err != nil {
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
		// Solidity: event tokenMinted(uint256 indexed tokenId, string tokenProperty, string tokenUri)
		func (_ERC721Base *ERC721BaseFilterer) ParseTokenMinted(log types.Log) (*ERC721BaseTokenMinted, error) {
			event := new(ERC721BaseTokenMinted)
			if err := _ERC721Base.contract.UnpackLog(event, "tokenMinted", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// ERC721BaseTransactionMemoIterator is returned from FilterTransactionMemo and is used to iterate over the raw logs and unpacked data for TransactionMemo events raised by the ERC721Base contract.
		type ERC721BaseTransactionMemoIterator struct {
			Event *ERC721BaseTransactionMemo // Event containing the contract specifics and raw log

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
		func (it *ERC721BaseTransactionMemoIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ERC721BaseTransactionMemo)
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
				it.Event = new(ERC721BaseTransactionMemo)
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
		func (it *ERC721BaseTransactionMemoIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ERC721BaseTransactionMemoIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ERC721BaseTransactionMemo represents a TransactionMemo event raised by the ERC721Base contract.
		type ERC721BaseTransactionMemo struct { 
			TokenId *big.Int; 
			Memo string; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterTransactionMemo is a free log retrieval operation binding the contract event 0x6bb79e977ebcae518f2d094cc741d1a77f05ae1d52cd8e23e5b7118572a438ce.
		//
		// Solidity: event transactionMemo(uint256 indexed tokenId, string memo)
 		func (_ERC721Base *ERC721BaseFilterer) FilterTransactionMemo(opts *bind.FilterOpts, tokenId []*big.Int) (*ERC721BaseTransactionMemoIterator, error) {
			
			var tokenIdRule []interface{}
			for _, tokenIdItem := range tokenId {
				tokenIdRule = append(tokenIdRule, tokenIdItem)
			}
			

			logs, sub, err := _ERC721Base.contract.FilterLogs(opts, "transactionMemo", tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &ERC721BaseTransactionMemoIterator{contract: _ERC721Base.contract, event: "transactionMemo", logs: logs, sub: sub}, nil
 		}

		// WatchTransactionMemo is a free log subscription operation binding the contract event 0x6bb79e977ebcae518f2d094cc741d1a77f05ae1d52cd8e23e5b7118572a438ce.
		//
		// Solidity: event transactionMemo(uint256 indexed tokenId, string memo)
		func (_ERC721Base *ERC721BaseFilterer) WatchTransactionMemo(opts *bind.WatchOpts, sink chan<- *ERC721BaseTransactionMemo, tokenId []*big.Int) (event.Subscription, error) {
			
			var tokenIdRule []interface{}
			for _, tokenIdItem := range tokenId {
				tokenIdRule = append(tokenIdRule, tokenIdItem)
			}
			

			logs, sub, err := _ERC721Base.contract.WatchLogs(opts, "transactionMemo", tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ERC721BaseTransactionMemo)
						if err := _ERC721Base.contract.UnpackLog(event, "transactionMemo", log); err != nil {
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
		func (_ERC721Base *ERC721BaseFilterer) ParseTransactionMemo(log types.Log) (*ERC721BaseTransactionMemo, error) {
			event := new(ERC721BaseTransactionMemo)
			if err := _ERC721Base.contract.UnpackLog(event, "transactionMemo", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	

	// ERC721EnumerableABI is the input ABI used to generate the binding from.
	const ERC721EnumerableABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

	
		// ERC721EnumerableFuncSigs maps the 4-byte function signature to its string representation.
		var ERC721EnumerableFuncSigs = map[string]string{
			"095ea7b3": "approve(address,uint256)",
			"70a08231": "balanceOf(address)",
			"081812fc": "getApproved(uint256)",
			"e985e9c5": "isApprovedForAll(address,address)",
			"6352211e": "ownerOf(uint256)",
			"42842e0e": "safeTransferFrom(address,address,uint256)",
			"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
			"a22cb465": "setApprovalForAll(address,bool)",
			"01ffc9a7": "supportsInterface(bytes4)",
			"4f6ccce7": "tokenByIndex(uint256)",
			"2f745c59": "tokenOfOwnerByIndex(address,uint256)",
			"18160ddd": "totalSupply()",
			"23b872dd": "transferFrom(address,address,uint256)",
			
		}
	

	
		// ERC721EnumerableBin is the compiled bytecode used for deploying new contracts.
		var ERC721EnumerableBin = "0x608060405234801561001057600080fd5b5061002a6301ffc9a760e01b6001600160e01b0361006116565b6100436380ac58cd60e01b6001600160e01b0361006116565b61005c63780e9d6360e01b6001600160e01b0361006116565b6100e5565b6001600160e01b031980821614156100c0576040805162461bcd60e51b815260206004820152601c60248201527f4552433136353a20696e76616c696420696e7465726661636520696400000000604482015290519081900360640190fd5b6001600160e01b0319166000908152602081905260409020805460ff19166001179055565b6112e8806100f46000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806342842e0e1161008c57806370a082311161006657806370a0823114610262578063a22cb46514610288578063b88d4fde146102b6578063e985e9c51461037c576100cf565b806342842e0e146101f25780634f6ccce7146102285780636352211e14610245576100cf565b806301ffc9a7146100d4578063081812fc1461010f578063095ea7b31461014857806318160ddd1461017657806323b872dd146101905780632f745c59146101c6575b600080fd5b6100fb600480360360208110156100ea57600080fd5b50356001600160e01b0319166103aa565b604080519115158252519081900360200190f35b61012c6004803603602081101561012557600080fd5b50356103c9565b604080516001600160a01b039092168252519081900360200190f35b6101746004803603604081101561015e57600080fd5b506001600160a01b03813516906020013561042b565b005b61017e610553565b60408051918252519081900360200190f35b610174600480360360608110156101a657600080fd5b506001600160a01b0381358116916020810135909116906040013561055a565b61017e600480360360408110156101dc57600080fd5b506001600160a01b0381351690602001356105b6565b6101746004803603606081101561020857600080fd5b506001600160a01b03813581169160208101359091169060400135610635565b61017e6004803603602081101561023e57600080fd5b5035610650565b61012c6004803603602081101561025b57600080fd5b50356106b6565b61017e6004803603602081101561027857600080fd5b50356001600160a01b0316610710565b6101746004803603604081101561029e57600080fd5b506001600160a01b0381351690602001351515610778565b610174600480360360808110156102cc57600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561030757600080fd5b82018360208201111561031957600080fd5b8035906020019184600183028401116401000000008311171561033b57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061087d945050505050565b6100fb6004803603604081101561039257600080fd5b506001600160a01b03813581169160200135166108db565b6001600160e01b03191660009081526020819052604090205460ff1690565b60006103d482610909565b61040f5760405162461bcd60e51b815260040180806020018281038252602c8152602001806111e1602c913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b6000610436826106b6565b9050806001600160a01b0316836001600160a01b031614156104895760405162461bcd60e51b81526004018080602001828103825260218152602001806112366021913960400191505060405180910390fd5b806001600160a01b031661049b610926565b6001600160a01b031614806104bc57506104bc816104b7610926565b6108db565b6104f75760405162461bcd60e51b81526004018080602001828103825260388152602001806111566038913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b6007545b90565b61056b610565610926565b8261092a565b6105a65760405162461bcd60e51b81526004018080602001828103825260318152602001806112576031913960400191505060405180910390fd5b6105b18383836109ce565b505050565b60006105c183610710565b82106105fe5760405162461bcd60e51b815260040180806020018281038252602b8152602001806110a9602b913960400191505060405180910390fd5b6001600160a01b038316600090815260056020526040902080548390811061062257fe5b9060005260206000200154905092915050565b6105b18383836040518060200160405280600081525061087d565b600061065a610553565b82106106975760405162461bcd60e51b815260040180806020018281038252602c815260200180611288602c913960400191505060405180910390fd5b600782815481106106a457fe5b90600052602060002001549050919050565b6000818152600160205260408120546001600160a01b03168061070a5760405162461bcd60e51b81526004018080602001828103825260298152602001806111b86029913960400191505060405180910390fd5b92915050565b60006001600160a01b0382166107575760405162461bcd60e51b815260040180806020018281038252602a81526020018061118e602a913960400191505060405180910390fd5b6001600160a01b038216600090815260036020526040902061070a906109ed565b610780610926565b6001600160a01b0316826001600160a01b031614156107e6576040805162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015290519081900360640190fd5b80600460006107f3610926565b6001600160a01b03908116825260208083019390935260409182016000908120918716808252919093529120805460ff191692151592909217909155610837610926565b60408051841515815290516001600160a01b0392909216917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c319181900360200190a35050565b61088e610888610926565b8361092a565b6108c95760405162461bcd60e51b81526004018080602001828103825260318152602001806112576031913960400191505060405180910390fd5b6108d5848484846109f1565b50505050565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b3390565b600061093582610909565b6109705760405162461bcd60e51b815260040180806020018281038252602c81526020018061112a602c913960400191505060405180910390fd5b600061097b836106b6565b9050806001600160a01b0316846001600160a01b031614806109b65750836001600160a01b03166109ab846103c9565b6001600160a01b0316145b806109c657506109c681856108db565b949350505050565b6109d9838383610a43565b6109e38382610b87565b6105b18282610c7c565b5490565b6109fc8484846109ce565b610a0884848484610cba565b6108d55760405162461bcd60e51b81526004018080602001828103825260328152602001806110d46032913960400191505060405180910390fd5b826001600160a01b0316610a56826106b6565b6001600160a01b031614610a9b5760405162461bcd60e51b815260040180806020018281038252602981526020018061120d6029913960400191505060405180910390fd5b6001600160a01b038216610ae05760405162461bcd60e51b81526004018080602001828103825260248152602001806111066024913960400191505060405180910390fd5b610ae981610ef5565b6001600160a01b0383166000908152600360205260409020610b0a90610f32565b6001600160a01b0382166000908152600360205260409020610b2b90610f49565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6001600160a01b038216600090815260056020526040812054610bb190600163ffffffff610f5216565b600083815260066020526040902054909150808214610c4c576001600160a01b0384166000908152600560205260408120805484908110610bee57fe5b906000526020600020015490508060056000876001600160a01b03166001600160a01b031681526020019081526020016000208381548110610c2c57fe5b600091825260208083209091019290925591825260069052604090208190555b6001600160a01b0384166000908152600560205260409020805490610c7590600019830161106b565b5050505050565b6001600160a01b0390911660009081526005602081815260408084208054868652600684529185208290559282526001810183559183529091200155565b6000610cce846001600160a01b0316610f9b565b610cda575060016109c6565b600060606001600160a01b038616630a85bd0160e11b610cf8610926565b89888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610d71578181015183820152602001610d59565b50505050905090810190601f168015610d9e5780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610e065780518252601f199092019160209182019101610de7565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610e68576040519150601f19603f3d011682016040523d82523d6000602084013e610e6d565b606091505b509150915081610ebe57805115610e875780518082602001fd5b60405162461bcd60e51b81526004018080602001828103825260328152602001806110d46032913960400191505060405180910390fd5b6000818060200190516020811015610ed557600080fd5b50516001600160e01b031916630a85bd0160e11b1493506109c692505050565b6000818152600260205260409020546001600160a01b031615610f2f57600081815260026020526040902080546001600160a01b03191690555b50565b8054610f4590600163ffffffff610f5216565b9055565b80546001019055565b6000610f9483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610fd4565b9392505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4708181148015906109c6575050151592915050565b600081848411156110635760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611028578181015183820152602001611010565b50505050905090810190601f1680156110555780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b8154818355818111156105b1576000838152602090206105b191810190830161055791905b808211156110a45760008155600101611090565b509056fe455243373231456e756d657261626c653a206f776e657220696e646578206f7574206f6620626f756e64734552433732313a207472616e7366657220746f206e6f6e20455243373231526563656976657220696d706c656d656e7465724552433732313a207472616e7366657220746f20746865207a65726f20616464726573734552433732313a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4552433732313a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734552433732313a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4552433732313a20617070726f76616c20746f2063757272656e74206f776e65724552433732313a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564455243373231456e756d657261626c653a20676c6f62616c20696e646578206f7574206f6620626f756e6473a265627a7a72315820d3115f6747957aa60096ea2dbcf21056f55632c1f68266412624c2711ae2bd9264736f6c63430005110032"

		// DeployERC721Enumerable deploys a new Ethereum contract, binding an instance of ERC721Enumerable to it.
		func DeployERC721Enumerable(auth *bind.TransactOpts, backend bind.ContractBackend ) (common.Address, *types.Transaction, *ERC721Enumerable, error) {
		  parsed, err := abi.JSON(strings.NewReader(ERC721EnumerableABI))
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  
		  address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721EnumerableBin), backend )
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  return address, tx, &ERC721Enumerable{ ERC721EnumerableCaller: ERC721EnumerableCaller{contract: contract}, ERC721EnumerableTransactor: ERC721EnumerableTransactor{contract: contract}, ERC721EnumerableFilterer: ERC721EnumerableFilterer{contract: contract} }, nil
		}
	

	// ERC721Enumerable is an auto generated Go binding around an Ethereum contract.
	type ERC721Enumerable struct {
	  ERC721EnumerableCaller     // Read-only binding to the contract
	  ERC721EnumerableTransactor // Write-only binding to the contract
	  ERC721EnumerableFilterer   // Log filterer for contract events
	}

	// ERC721EnumerableCaller is an auto generated read-only Go binding around an Ethereum contract.
	type ERC721EnumerableCaller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC721EnumerableTransactor is an auto generated write-only Go binding around an Ethereum contract.
	type ERC721EnumerableTransactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC721EnumerableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type ERC721EnumerableFilterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC721EnumerableSession is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type ERC721EnumerableSession struct {
	  Contract     *ERC721Enumerable        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// ERC721EnumerableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type ERC721EnumerableCallerSession struct {
	  Contract *ERC721EnumerableCaller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// ERC721EnumerableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type ERC721EnumerableTransactorSession struct {
	  Contract     *ERC721EnumerableTransactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// ERC721EnumerableRaw is an auto generated low-level Go binding around an Ethereum contract.
	type ERC721EnumerableRaw struct {
	  Contract *ERC721Enumerable // Generic contract binding to access the raw methods on
	}

	// ERC721EnumerableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type ERC721EnumerableCallerRaw struct {
		Contract *ERC721EnumerableCaller // Generic read-only contract binding to access the raw methods on
	}

	// ERC721EnumerableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type ERC721EnumerableTransactorRaw struct {
		Contract *ERC721EnumerableTransactor // Generic write-only contract binding to access the raw methods on
	}

	// NewERC721Enumerable creates a new instance of ERC721Enumerable, bound to a specific deployed contract.
	func NewERC721Enumerable(address common.Address, backend bind.ContractBackend) (*ERC721Enumerable, error) {
	  contract, err := bindERC721Enumerable(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC721Enumerable{ ERC721EnumerableCaller: ERC721EnumerableCaller{contract: contract}, ERC721EnumerableTransactor: ERC721EnumerableTransactor{contract: contract}, ERC721EnumerableFilterer: ERC721EnumerableFilterer{contract: contract} }, nil
	}

	// NewERC721EnumerableCaller creates a new read-only instance of ERC721Enumerable, bound to a specific deployed contract.
	func NewERC721EnumerableCaller(address common.Address, caller bind.ContractCaller) (*ERC721EnumerableCaller, error) {
	  contract, err := bindERC721Enumerable(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC721EnumerableCaller{contract: contract}, nil
	}

	// NewERC721EnumerableTransactor creates a new write-only instance of ERC721Enumerable, bound to a specific deployed contract.
	func NewERC721EnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721EnumerableTransactor, error) {
	  contract, err := bindERC721Enumerable(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC721EnumerableTransactor{contract: contract}, nil
	}

	// NewERC721EnumerableFilterer creates a new log filterer instance of ERC721Enumerable, bound to a specific deployed contract.
 	func NewERC721EnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721EnumerableFilterer, error) {
 	  contract, err := bindERC721Enumerable(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &ERC721EnumerableFilterer{contract: contract}, nil
 	}

	// bindERC721Enumerable binds a generic wrapper to an already deployed contract.
	func bindERC721Enumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(ERC721EnumerableABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_ERC721Enumerable *ERC721EnumerableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _ERC721Enumerable.Contract.ERC721EnumerableCaller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_ERC721Enumerable *ERC721EnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _ERC721Enumerable.Contract.ERC721EnumerableTransactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_ERC721Enumerable *ERC721EnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _ERC721Enumerable.Contract.ERC721EnumerableTransactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_ERC721Enumerable *ERC721EnumerableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _ERC721Enumerable.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_ERC721Enumerable *ERC721EnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _ERC721Enumerable.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_ERC721Enumerable *ERC721EnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _ERC721Enumerable.Contract.contract.Transact(opts, method, params...)
	}

	
		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256)
		func (_ERC721Enumerable *ERC721EnumerableCaller) BalanceOf(opts *bind.CallOpts , owner common.Address ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _ERC721Enumerable.contract.Call(opts, out, "balanceOf" , owner)
			return *ret0, err
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256)
		func (_ERC721Enumerable *ERC721EnumerableSession) BalanceOf( owner common.Address ) ( *big.Int,  error) {
		  return _ERC721Enumerable.Contract.BalanceOf(&_ERC721Enumerable.CallOpts , owner)
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256)
		func (_ERC721Enumerable *ERC721EnumerableCallerSession) BalanceOf( owner common.Address ) ( *big.Int,  error) {
		  return _ERC721Enumerable.Contract.BalanceOf(&_ERC721Enumerable.CallOpts , owner)
		}
	
		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address)
		func (_ERC721Enumerable *ERC721EnumerableCaller) GetApproved(opts *bind.CallOpts , tokenId *big.Int ) (common.Address, error) {
			var (
				ret0 = new(common.Address)
				
			)
			out := ret0
			err := _ERC721Enumerable.contract.Call(opts, out, "getApproved" , tokenId)
			return *ret0, err
		}

		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address)
		func (_ERC721Enumerable *ERC721EnumerableSession) GetApproved( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721Enumerable.Contract.GetApproved(&_ERC721Enumerable.CallOpts , tokenId)
		}

		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address)
		func (_ERC721Enumerable *ERC721EnumerableCallerSession) GetApproved( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721Enumerable.Contract.GetApproved(&_ERC721Enumerable.CallOpts , tokenId)
		}
	
		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_ERC721Enumerable *ERC721EnumerableCaller) IsApprovedForAll(opts *bind.CallOpts , owner common.Address , operator common.Address ) (bool, error) {
			var (
				ret0 = new(bool)
				
			)
			out := ret0
			err := _ERC721Enumerable.contract.Call(opts, out, "isApprovedForAll" , owner, operator)
			return *ret0, err
		}

		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_ERC721Enumerable *ERC721EnumerableSession) IsApprovedForAll( owner common.Address , operator common.Address ) ( bool,  error) {
		  return _ERC721Enumerable.Contract.IsApprovedForAll(&_ERC721Enumerable.CallOpts , owner, operator)
		}

		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_ERC721Enumerable *ERC721EnumerableCallerSession) IsApprovedForAll( owner common.Address , operator common.Address ) ( bool,  error) {
		  return _ERC721Enumerable.Contract.IsApprovedForAll(&_ERC721Enumerable.CallOpts , owner, operator)
		}
	
		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address)
		func (_ERC721Enumerable *ERC721EnumerableCaller) OwnerOf(opts *bind.CallOpts , tokenId *big.Int ) (common.Address, error) {
			var (
				ret0 = new(common.Address)
				
			)
			out := ret0
			err := _ERC721Enumerable.contract.Call(opts, out, "ownerOf" , tokenId)
			return *ret0, err
		}

		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address)
		func (_ERC721Enumerable *ERC721EnumerableSession) OwnerOf( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721Enumerable.Contract.OwnerOf(&_ERC721Enumerable.CallOpts , tokenId)
		}

		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address)
		func (_ERC721Enumerable *ERC721EnumerableCallerSession) OwnerOf( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721Enumerable.Contract.OwnerOf(&_ERC721Enumerable.CallOpts , tokenId)
		}
	
		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_ERC721Enumerable *ERC721EnumerableCaller) SupportsInterface(opts *bind.CallOpts , interfaceId [4]byte ) (bool, error) {
			var (
				ret0 = new(bool)
				
			)
			out := ret0
			err := _ERC721Enumerable.contract.Call(opts, out, "supportsInterface" , interfaceId)
			return *ret0, err
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_ERC721Enumerable *ERC721EnumerableSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _ERC721Enumerable.Contract.SupportsInterface(&_ERC721Enumerable.CallOpts , interfaceId)
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_ERC721Enumerable *ERC721EnumerableCallerSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _ERC721Enumerable.Contract.SupportsInterface(&_ERC721Enumerable.CallOpts , interfaceId)
		}
	
		// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
		//
		// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
		func (_ERC721Enumerable *ERC721EnumerableCaller) TokenByIndex(opts *bind.CallOpts , index *big.Int ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _ERC721Enumerable.contract.Call(opts, out, "tokenByIndex" , index)
			return *ret0, err
		}

		// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
		//
		// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
		func (_ERC721Enumerable *ERC721EnumerableSession) TokenByIndex( index *big.Int ) ( *big.Int,  error) {
		  return _ERC721Enumerable.Contract.TokenByIndex(&_ERC721Enumerable.CallOpts , index)
		}

		// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
		//
		// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
		func (_ERC721Enumerable *ERC721EnumerableCallerSession) TokenByIndex( index *big.Int ) ( *big.Int,  error) {
		  return _ERC721Enumerable.Contract.TokenByIndex(&_ERC721Enumerable.CallOpts , index)
		}
	
		// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
		//
		// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
		func (_ERC721Enumerable *ERC721EnumerableCaller) TokenOfOwnerByIndex(opts *bind.CallOpts , owner common.Address , index *big.Int ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _ERC721Enumerable.contract.Call(opts, out, "tokenOfOwnerByIndex" , owner, index)
			return *ret0, err
		}

		// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
		//
		// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
		func (_ERC721Enumerable *ERC721EnumerableSession) TokenOfOwnerByIndex( owner common.Address , index *big.Int ) ( *big.Int,  error) {
		  return _ERC721Enumerable.Contract.TokenOfOwnerByIndex(&_ERC721Enumerable.CallOpts , owner, index)
		}

		// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
		//
		// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
		func (_ERC721Enumerable *ERC721EnumerableCallerSession) TokenOfOwnerByIndex( owner common.Address , index *big.Int ) ( *big.Int,  error) {
		  return _ERC721Enumerable.Contract.TokenOfOwnerByIndex(&_ERC721Enumerable.CallOpts , owner, index)
		}
	
		// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
		//
		// Solidity: function totalSupply() view returns(uint256)
		func (_ERC721Enumerable *ERC721EnumerableCaller) TotalSupply(opts *bind.CallOpts ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _ERC721Enumerable.contract.Call(opts, out, "totalSupply" )
			return *ret0, err
		}

		// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
		//
		// Solidity: function totalSupply() view returns(uint256)
		func (_ERC721Enumerable *ERC721EnumerableSession) TotalSupply() ( *big.Int,  error) {
		  return _ERC721Enumerable.Contract.TotalSupply(&_ERC721Enumerable.CallOpts )
		}

		// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
		//
		// Solidity: function totalSupply() view returns(uint256)
		func (_ERC721Enumerable *ERC721EnumerableCallerSession) TotalSupply() ( *big.Int,  error) {
		  return _ERC721Enumerable.Contract.TotalSupply(&_ERC721Enumerable.CallOpts )
		}
	

	
		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_ERC721Enumerable *ERC721EnumerableTransactor) Approve(opts *bind.TransactOpts , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _ERC721Enumerable.contract.Transact(opts, "approve" , to, tokenId)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_ERC721Enumerable *ERC721EnumerableSession) Approve( to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Enumerable.Contract.Approve(&_ERC721Enumerable.TransactOpts , to, tokenId)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_ERC721Enumerable *ERC721EnumerableTransactorSession) Approve( to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Enumerable.Contract.Approve(&_ERC721Enumerable.TransactOpts , to, tokenId)
		}
	
		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Enumerable *ERC721EnumerableTransactor) SafeTransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _ERC721Enumerable.contract.Transact(opts, "safeTransferFrom" , from, to, tokenId)
		}

		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Enumerable *ERC721EnumerableSession) SafeTransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Enumerable.Contract.SafeTransferFrom(&_ERC721Enumerable.TransactOpts , from, to, tokenId)
		}

		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Enumerable *ERC721EnumerableTransactorSession) SafeTransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Enumerable.Contract.SafeTransferFrom(&_ERC721Enumerable.TransactOpts , from, to, tokenId)
		}
	
		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
		func (_ERC721Enumerable *ERC721EnumerableTransactor) SafeTransferFrom0(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int , _data []byte ) (*types.Transaction, error) {
			return _ERC721Enumerable.contract.Transact(opts, "safeTransferFrom0" , from, to, tokenId, _data)
		}

		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
		func (_ERC721Enumerable *ERC721EnumerableSession) SafeTransferFrom0( from common.Address , to common.Address , tokenId *big.Int , _data []byte ) (*types.Transaction, error) {
		  return _ERC721Enumerable.Contract.SafeTransferFrom0(&_ERC721Enumerable.TransactOpts , from, to, tokenId, _data)
		}

		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
		func (_ERC721Enumerable *ERC721EnumerableTransactorSession) SafeTransferFrom0( from common.Address , to common.Address , tokenId *big.Int , _data []byte ) (*types.Transaction, error) {
		  return _ERC721Enumerable.Contract.SafeTransferFrom0(&_ERC721Enumerable.TransactOpts , from, to, tokenId, _data)
		}
	
		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address to, bool approved) returns()
		func (_ERC721Enumerable *ERC721EnumerableTransactor) SetApprovalForAll(opts *bind.TransactOpts , to common.Address , approved bool ) (*types.Transaction, error) {
			return _ERC721Enumerable.contract.Transact(opts, "setApprovalForAll" , to, approved)
		}

		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address to, bool approved) returns()
		func (_ERC721Enumerable *ERC721EnumerableSession) SetApprovalForAll( to common.Address , approved bool ) (*types.Transaction, error) {
		  return _ERC721Enumerable.Contract.SetApprovalForAll(&_ERC721Enumerable.TransactOpts , to, approved)
		}

		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address to, bool approved) returns()
		func (_ERC721Enumerable *ERC721EnumerableTransactorSession) SetApprovalForAll( to common.Address , approved bool ) (*types.Transaction, error) {
		  return _ERC721Enumerable.Contract.SetApprovalForAll(&_ERC721Enumerable.TransactOpts , to, approved)
		}
	
		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Enumerable *ERC721EnumerableTransactor) TransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _ERC721Enumerable.contract.Transact(opts, "transferFrom" , from, to, tokenId)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Enumerable *ERC721EnumerableSession) TransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Enumerable.Contract.TransferFrom(&_ERC721Enumerable.TransactOpts , from, to, tokenId)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Enumerable *ERC721EnumerableTransactorSession) TransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Enumerable.Contract.TransferFrom(&_ERC721Enumerable.TransactOpts , from, to, tokenId)
		}
	

	

	

	
		// ERC721EnumerableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Enumerable contract.
		type ERC721EnumerableApprovalIterator struct {
			Event *ERC721EnumerableApproval // Event containing the contract specifics and raw log

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
		func (it *ERC721EnumerableApprovalIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ERC721EnumerableApproval)
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
				it.Event = new(ERC721EnumerableApproval)
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
		func (it *ERC721EnumerableApprovalIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ERC721EnumerableApprovalIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ERC721EnumerableApproval represents a Approval event raised by the ERC721Enumerable contract.
		type ERC721EnumerableApproval struct { 
			Owner common.Address; 
			Approved common.Address; 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
 		func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721EnumerableApprovalIterator, error) {
			
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

			logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &ERC721EnumerableApprovalIterator{contract: _ERC721Enumerable.contract, event: "Approval", logs: logs, sub: sub}, nil
 		}

		// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
		func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {
			
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

			logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ERC721EnumerableApproval)
						if err := _ERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
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
		func (_ERC721Enumerable *ERC721EnumerableFilterer) ParseApproval(log types.Log) (*ERC721EnumerableApproval, error) {
			event := new(ERC721EnumerableApproval)
			if err := _ERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// ERC721EnumerableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Enumerable contract.
		type ERC721EnumerableApprovalForAllIterator struct {
			Event *ERC721EnumerableApprovalForAll // Event containing the contract specifics and raw log

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
		func (it *ERC721EnumerableApprovalForAllIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ERC721EnumerableApprovalForAll)
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
				it.Event = new(ERC721EnumerableApprovalForAll)
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
		func (it *ERC721EnumerableApprovalForAllIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ERC721EnumerableApprovalForAllIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ERC721EnumerableApprovalForAll represents a ApprovalForAll event raised by the ERC721Enumerable contract.
		type ERC721EnumerableApprovalForAll struct { 
			Owner common.Address; 
			Operator common.Address; 
			Approved bool; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
 		func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721EnumerableApprovalForAllIterator, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var operatorRule []interface{}
			for _, operatorItem := range operator {
				operatorRule = append(operatorRule, operatorItem)
			}
			

			logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
			if err != nil {
				return nil, err
			}
			return &ERC721EnumerableApprovalForAllIterator{contract: _ERC721Enumerable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
 		}

		// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
		func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var operatorRule []interface{}
			for _, operatorItem := range operator {
				operatorRule = append(operatorRule, operatorItem)
			}
			

			logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ERC721EnumerableApprovalForAll)
						if err := _ERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
		func (_ERC721Enumerable *ERC721EnumerableFilterer) ParseApprovalForAll(log types.Log) (*ERC721EnumerableApprovalForAll, error) {
			event := new(ERC721EnumerableApprovalForAll)
			if err := _ERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// ERC721EnumerableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Enumerable contract.
		type ERC721EnumerableTransferIterator struct {
			Event *ERC721EnumerableTransfer // Event containing the contract specifics and raw log

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
		func (it *ERC721EnumerableTransferIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ERC721EnumerableTransfer)
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
				it.Event = new(ERC721EnumerableTransfer)
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
		func (it *ERC721EnumerableTransferIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ERC721EnumerableTransferIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ERC721EnumerableTransfer represents a Transfer event raised by the ERC721Enumerable contract.
		type ERC721EnumerableTransfer struct { 
			From common.Address; 
			To common.Address; 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
 		func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721EnumerableTransferIterator, error) {
			
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

			logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &ERC721EnumerableTransferIterator{contract: _ERC721Enumerable.contract, event: "Transfer", logs: logs, sub: sub}, nil
 		}

		// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
		func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {
			
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

			logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ERC721EnumerableTransfer)
						if err := _ERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
		func (_ERC721Enumerable *ERC721EnumerableFilterer) ParseTransfer(log types.Log) (*ERC721EnumerableTransfer, error) {
			event := new(ERC721EnumerableTransfer)
			if err := _ERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	

	// ERC721FullABI is the input ABI used to generate the binding from.
	const ERC721FullABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

	
		// ERC721FullFuncSigs maps the 4-byte function signature to its string representation.
		var ERC721FullFuncSigs = map[string]string{
			"095ea7b3": "approve(address,uint256)",
			"70a08231": "balanceOf(address)",
			"6c0360eb": "baseURI()",
			"081812fc": "getApproved(uint256)",
			"e985e9c5": "isApprovedForAll(address,address)",
			"06fdde03": "name()",
			"6352211e": "ownerOf(uint256)",
			"42842e0e": "safeTransferFrom(address,address,uint256)",
			"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
			"a22cb465": "setApprovalForAll(address,bool)",
			"01ffc9a7": "supportsInterface(bytes4)",
			"95d89b41": "symbol()",
			"4f6ccce7": "tokenByIndex(uint256)",
			"2f745c59": "tokenOfOwnerByIndex(address,uint256)",
			"c87b56dd": "tokenURI(uint256)",
			"18160ddd": "totalSupply()",
			"23b872dd": "transferFrom(address,address,uint256)",
			
		}
	

	
		// ERC721FullBin is the compiled bytecode used for deploying new contracts.
		var ERC721FullBin = "0x60806040523480156200001157600080fd5b5060405162001aa438038062001aa4833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b5060405250839150829050620001be6301ffc9a760e01b6001600160e01b036200024516565b620001d96380ac58cd60e01b6001600160e01b036200024516565b620001f463780e9d6360e01b6001600160e01b036200024516565b815162000209906009906020850190620002ca565b5080516200021f90600a906020840190620002ca565b506200023b635b5e139f60e01b6001600160e01b036200024516565b505050506200036f565b6001600160e01b03198082161415620002a5576040805162461bcd60e51b815260206004820152601c60248201527f4552433136353a20696e76616c696420696e7465726661636520696400000000604482015290519081900360640190fd5b6001600160e01b0319166000908152602081905260409020805460ff19166001179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200030d57805160ff19168380011785556200033d565b828001600101855582156200033d579182015b828111156200033d57825182559160200191906001019062000320565b506200034b9291506200034f565b5090565b6200036c91905b808211156200034b576000815560010162000356565b90565b611725806200037f6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80634f6ccce7116100a257806395d89b411161007157806395d89b4114610349578063a22cb46514610351578063b88d4fde1461037f578063c87b56dd14610445578063e985e9c5146104625761010b565b80634f6ccce7146102e15780636352211e146102fe5780636c0360eb1461031b57806370a08231146103235761010b565b806318160ddd116100de57806318160ddd1461022f57806323b872dd146102495780632f745c591461027f57806342842e0e146102ab5761010b565b806301ffc9a71461011057806306fdde031461014b578063081812fc146101c8578063095ea7b314610201575b600080fd5b6101376004803603602081101561012657600080fd5b50356001600160e01b031916610490565b604080519115158252519081900360200190f35b6101536104b3565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561018d578181015183820152602001610175565b50505050905090810190601f1680156101ba5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101e5600480360360208110156101de57600080fd5b503561054a565b604080516001600160a01b039092168252519081900360200190f35b61022d6004803603604081101561021757600080fd5b506001600160a01b0381351690602001356105ac565b005b6102376106d4565b60408051918252519081900360200190f35b61022d6004803603606081101561025f57600080fd5b506001600160a01b038135811691602081013590911690604001356106da565b6102376004803603604081101561029557600080fd5b506001600160a01b038135169060200135610736565b61022d600480360360608110156102c157600080fd5b506001600160a01b038135811691602081013590911690604001356107b5565b610237600480360360208110156102f757600080fd5b50356107d0565b6101e56004803603602081101561031457600080fd5b5035610836565b610153610890565b6102376004803603602081101561033957600080fd5b50356001600160a01b03166108f1565b610153610959565b61022d6004803603604081101561036757600080fd5b506001600160a01b03813516906020013515156109ba565b61022d6004803603608081101561039557600080fd5b6001600160a01b038235811692602081013590911691604082013591908101906080810160608201356401000000008111156103d057600080fd5b8201836020820111156103e257600080fd5b8035906020019184600183028401116401000000008311171561040457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610abf945050505050565b6101536004803603602081101561045b57600080fd5b5035610b1d565b6101376004803603604081101561047857600080fd5b506001600160a01b0381358116916020013516610ce9565b6001600160e01b0319811660009081526020819052604090205460ff165b919050565b60098054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561053f5780601f106105145761010080835404028352916020019161053f565b820191906000526020600020905b81548152906001019060200180831161052257829003601f168201915b505050505090505b90565b600061055582610d17565b6105905760405162461bcd60e51b815260040180806020018281038252602c8152602001806115ef602c913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b60006105b782610836565b9050806001600160a01b0316836001600160a01b0316141561060a5760405162461bcd60e51b81526004018080602001828103825260218152602001806116736021913960400191505060405180910390fd5b806001600160a01b031661061c610d34565b6001600160a01b0316148061063d575061063d81610638610d34565b610ce9565b6106785760405162461bcd60e51b81526004018080602001828103825260388152602001806115646038913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b60075490565b6106eb6106e5610d34565b82610d38565b6107265760405162461bcd60e51b81526004018080602001828103825260318152602001806116946031913960400191505060405180910390fd5b610731838383610ddc565b505050565b6000610741836108f1565b821061077e5760405162461bcd60e51b815260040180806020018281038252602b8152602001806114b7602b913960400191505060405180910390fd5b6001600160a01b03831660009081526005602052604090208054839081106107a257fe5b9060005260206000200154905092915050565b61073183838360405180602001604052806000815250610abf565b60006107da6106d4565b82106108175760405162461bcd60e51b815260040180806020018281038252602c8152602001806116c5602c913960400191505060405180910390fd5b6007828154811061082457fe5b90600052602060002001549050919050565b6000818152600160205260408120546001600160a01b03168061088a5760405162461bcd60e51b81526004018080602001828103825260298152602001806115c66029913960400191505060405180910390fd5b92915050565b600b8054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561053f5780601f106105145761010080835404028352916020019161053f565b60006001600160a01b0382166109385760405162461bcd60e51b815260040180806020018281038252602a81526020018061159c602a913960400191505060405180910390fd5b6001600160a01b038216600090815260036020526040902061088a90610dfb565b600a8054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561053f5780601f106105145761010080835404028352916020019161053f565b6109c2610d34565b6001600160a01b0316826001600160a01b03161415610a28576040805162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015290519081900360640190fd5b8060046000610a35610d34565b6001600160a01b03908116825260208083019390935260409182016000908120918716808252919093529120805460ff191692151592909217909155610a79610d34565b60408051841515815290516001600160a01b0392909216917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c319181900360200190a35050565b610ad0610aca610d34565b83610d38565b610b0b5760405162461bcd60e51b81526004018080602001828103825260318152602001806116946031913960400191505060405180910390fd5b610b1784848484610dff565b50505050565b6060610b2882610d17565b610b635760405162461bcd60e51b815260040180806020018281038252602f815260200180611644602f913960400191505060405180910390fd5b6000828152600c602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845260609392830182828015610bf85780601f10610bcd57610100808354040283529160200191610bf8565b820191906000526020600020905b815481529060010190602001808311610bdb57829003601f168201915b50505050509050805160001415610c1f5750506040805160208101909152600081526104ae565b600b816040516020018083805460018160011615610100020316600290048015610c805780601f10610c5e576101008083540402835291820191610c80565b820191906000526020600020905b815481529060010190602001808311610c6c575b5050825160208401908083835b60208310610cac5780518252601f199092019160209182019101610c8d565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529150506104ae565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b3390565b6000610d4382610d17565b610d7e5760405162461bcd60e51b815260040180806020018281038252602c815260200180611538602c913960400191505060405180910390fd5b6000610d8983610836565b9050806001600160a01b0316846001600160a01b03161480610dc45750836001600160a01b0316610db98461054a565b6001600160a01b0316145b80610dd45750610dd48185610ce9565b949350505050565b610de7838383610e51565b610df18382610f95565b610731828261108a565b5490565b610e0a848484610ddc565b610e16848484846110c8565b610b175760405162461bcd60e51b81526004018080602001828103825260328152602001806114e26032913960400191505060405180910390fd5b826001600160a01b0316610e6482610836565b6001600160a01b031614610ea95760405162461bcd60e51b815260040180806020018281038252602981526020018061161b6029913960400191505060405180910390fd5b6001600160a01b038216610eee5760405162461bcd60e51b81526004018080602001828103825260248152602001806115146024913960400191505060405180910390fd5b610ef781611303565b6001600160a01b0383166000908152600360205260409020610f1890611340565b6001600160a01b0382166000908152600360205260409020610f3990611357565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b6001600160a01b038216600090815260056020526040812054610fbf90600163ffffffff61136016565b60008381526006602052604090205490915080821461105a576001600160a01b0384166000908152600560205260408120805484908110610ffc57fe5b906000526020600020015490508060056000876001600160a01b03166001600160a01b03168152602001908152602001600020838154811061103a57fe5b600091825260208083209091019290925591825260069052604090208190555b6001600160a01b0384166000908152600560205260409020805490611083906000198301611479565b5050505050565b6001600160a01b0390911660009081526005602081815260408084208054868652600684529185208290559282526001810183559183529091200155565b60006110dc846001600160a01b03166113a9565b6110e857506001610dd4565b600060606001600160a01b038616630a85bd0160e11b611106610d34565b89888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561117f578181015183820152602001611167565b50505050905090810190601f1680156111ac5780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b602083106112145780518252601f1990920191602091820191016111f5565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611276576040519150601f19603f3d011682016040523d82523d6000602084013e61127b565b606091505b5091509150816112cc578051156112955780518082602001fd5b60405162461bcd60e51b81526004018080602001828103825260328152602001806114e26032913960400191505060405180910390fd5b60008180602001905160208110156112e357600080fd5b50516001600160e01b031916630a85bd0160e11b149350610dd492505050565b6000818152600260205260409020546001600160a01b03161561133d57600081815260026020526040902080546001600160a01b03191690555b50565b805461135390600163ffffffff61136016565b9055565b80546001019055565b60006113a283836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506113e2565b9392505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470818114801590610dd4575050151592915050565b600081848411156114715760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561143657818101518382015260200161141e565b50505050905090810190601f1680156114635780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b8154818355818111156107315760008381526020902061073191810190830161054791905b808211156114b2576000815560010161149e565b509056fe455243373231456e756d657261626c653a206f776e657220696e646578206f7574206f6620626f756e64734552433732313a207472616e7366657220746f206e6f6e20455243373231526563656976657220696d706c656d656e7465724552433732313a207472616e7366657220746f20746865207a65726f20616464726573734552433732313a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4552433732313a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734552433732313a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4552433732314d657461646174613a2055524920717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76616c20746f2063757272656e74206f776e65724552433732313a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564455243373231456e756d657261626c653a20676c6f62616c20696e646578206f7574206f6620626f756e6473a265627a7a72315820b57200e623190b5d1a62d250b04b0cec970ed4d1d282c0b3bd2b925b9c4cbb5764736f6c63430005110032"

		// DeployERC721Full deploys a new Ethereum contract, binding an instance of ERC721Full to it.
		func DeployERC721Full(auth *bind.TransactOpts, backend bind.ContractBackend , name string, symbol string) (common.Address, *types.Transaction, *ERC721Full, error) {
		  parsed, err := abi.JSON(strings.NewReader(ERC721FullABI))
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  
		  address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721FullBin), backend , name, symbol)
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  return address, tx, &ERC721Full{ ERC721FullCaller: ERC721FullCaller{contract: contract}, ERC721FullTransactor: ERC721FullTransactor{contract: contract}, ERC721FullFilterer: ERC721FullFilterer{contract: contract} }, nil
		}
	

	// ERC721Full is an auto generated Go binding around an Ethereum contract.
	type ERC721Full struct {
	  ERC721FullCaller     // Read-only binding to the contract
	  ERC721FullTransactor // Write-only binding to the contract
	  ERC721FullFilterer   // Log filterer for contract events
	}

	// ERC721FullCaller is an auto generated read-only Go binding around an Ethereum contract.
	type ERC721FullCaller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC721FullTransactor is an auto generated write-only Go binding around an Ethereum contract.
	type ERC721FullTransactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC721FullFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type ERC721FullFilterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC721FullSession is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type ERC721FullSession struct {
	  Contract     *ERC721Full        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// ERC721FullCallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type ERC721FullCallerSession struct {
	  Contract *ERC721FullCaller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// ERC721FullTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type ERC721FullTransactorSession struct {
	  Contract     *ERC721FullTransactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// ERC721FullRaw is an auto generated low-level Go binding around an Ethereum contract.
	type ERC721FullRaw struct {
	  Contract *ERC721Full // Generic contract binding to access the raw methods on
	}

	// ERC721FullCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type ERC721FullCallerRaw struct {
		Contract *ERC721FullCaller // Generic read-only contract binding to access the raw methods on
	}

	// ERC721FullTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type ERC721FullTransactorRaw struct {
		Contract *ERC721FullTransactor // Generic write-only contract binding to access the raw methods on
	}

	// NewERC721Full creates a new instance of ERC721Full, bound to a specific deployed contract.
	func NewERC721Full(address common.Address, backend bind.ContractBackend) (*ERC721Full, error) {
	  contract, err := bindERC721Full(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC721Full{ ERC721FullCaller: ERC721FullCaller{contract: contract}, ERC721FullTransactor: ERC721FullTransactor{contract: contract}, ERC721FullFilterer: ERC721FullFilterer{contract: contract} }, nil
	}

	// NewERC721FullCaller creates a new read-only instance of ERC721Full, bound to a specific deployed contract.
	func NewERC721FullCaller(address common.Address, caller bind.ContractCaller) (*ERC721FullCaller, error) {
	  contract, err := bindERC721Full(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC721FullCaller{contract: contract}, nil
	}

	// NewERC721FullTransactor creates a new write-only instance of ERC721Full, bound to a specific deployed contract.
	func NewERC721FullTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721FullTransactor, error) {
	  contract, err := bindERC721Full(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC721FullTransactor{contract: contract}, nil
	}

	// NewERC721FullFilterer creates a new log filterer instance of ERC721Full, bound to a specific deployed contract.
 	func NewERC721FullFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721FullFilterer, error) {
 	  contract, err := bindERC721Full(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &ERC721FullFilterer{contract: contract}, nil
 	}

	// bindERC721Full binds a generic wrapper to an already deployed contract.
	func bindERC721Full(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(ERC721FullABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_ERC721Full *ERC721FullRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _ERC721Full.Contract.ERC721FullCaller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_ERC721Full *ERC721FullRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _ERC721Full.Contract.ERC721FullTransactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_ERC721Full *ERC721FullRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _ERC721Full.Contract.ERC721FullTransactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_ERC721Full *ERC721FullCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _ERC721Full.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_ERC721Full *ERC721FullTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _ERC721Full.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_ERC721Full *ERC721FullTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _ERC721Full.Contract.contract.Transact(opts, method, params...)
	}

	
		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256)
		func (_ERC721Full *ERC721FullCaller) BalanceOf(opts *bind.CallOpts , owner common.Address ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _ERC721Full.contract.Call(opts, out, "balanceOf" , owner)
			return *ret0, err
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256)
		func (_ERC721Full *ERC721FullSession) BalanceOf( owner common.Address ) ( *big.Int,  error) {
		  return _ERC721Full.Contract.BalanceOf(&_ERC721Full.CallOpts , owner)
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256)
		func (_ERC721Full *ERC721FullCallerSession) BalanceOf( owner common.Address ) ( *big.Int,  error) {
		  return _ERC721Full.Contract.BalanceOf(&_ERC721Full.CallOpts , owner)
		}
	
		// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
		//
		// Solidity: function baseURI() view returns(string)
		func (_ERC721Full *ERC721FullCaller) BaseURI(opts *bind.CallOpts ) (string, error) {
			var (
				ret0 = new(string)
				
			)
			out := ret0
			err := _ERC721Full.contract.Call(opts, out, "baseURI" )
			return *ret0, err
		}

		// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
		//
		// Solidity: function baseURI() view returns(string)
		func (_ERC721Full *ERC721FullSession) BaseURI() ( string,  error) {
		  return _ERC721Full.Contract.BaseURI(&_ERC721Full.CallOpts )
		}

		// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
		//
		// Solidity: function baseURI() view returns(string)
		func (_ERC721Full *ERC721FullCallerSession) BaseURI() ( string,  error) {
		  return _ERC721Full.Contract.BaseURI(&_ERC721Full.CallOpts )
		}
	
		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address)
		func (_ERC721Full *ERC721FullCaller) GetApproved(opts *bind.CallOpts , tokenId *big.Int ) (common.Address, error) {
			var (
				ret0 = new(common.Address)
				
			)
			out := ret0
			err := _ERC721Full.contract.Call(opts, out, "getApproved" , tokenId)
			return *ret0, err
		}

		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address)
		func (_ERC721Full *ERC721FullSession) GetApproved( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721Full.Contract.GetApproved(&_ERC721Full.CallOpts , tokenId)
		}

		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address)
		func (_ERC721Full *ERC721FullCallerSession) GetApproved( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721Full.Contract.GetApproved(&_ERC721Full.CallOpts , tokenId)
		}
	
		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_ERC721Full *ERC721FullCaller) IsApprovedForAll(opts *bind.CallOpts , owner common.Address , operator common.Address ) (bool, error) {
			var (
				ret0 = new(bool)
				
			)
			out := ret0
			err := _ERC721Full.contract.Call(opts, out, "isApprovedForAll" , owner, operator)
			return *ret0, err
		}

		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_ERC721Full *ERC721FullSession) IsApprovedForAll( owner common.Address , operator common.Address ) ( bool,  error) {
		  return _ERC721Full.Contract.IsApprovedForAll(&_ERC721Full.CallOpts , owner, operator)
		}

		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_ERC721Full *ERC721FullCallerSession) IsApprovedForAll( owner common.Address , operator common.Address ) ( bool,  error) {
		  return _ERC721Full.Contract.IsApprovedForAll(&_ERC721Full.CallOpts , owner, operator)
		}
	
		// Name is a free data retrieval call binding the contract method 0x06fdde03.
		//
		// Solidity: function name() view returns(string)
		func (_ERC721Full *ERC721FullCaller) Name(opts *bind.CallOpts ) (string, error) {
			var (
				ret0 = new(string)
				
			)
			out := ret0
			err := _ERC721Full.contract.Call(opts, out, "name" )
			return *ret0, err
		}

		// Name is a free data retrieval call binding the contract method 0x06fdde03.
		//
		// Solidity: function name() view returns(string)
		func (_ERC721Full *ERC721FullSession) Name() ( string,  error) {
		  return _ERC721Full.Contract.Name(&_ERC721Full.CallOpts )
		}

		// Name is a free data retrieval call binding the contract method 0x06fdde03.
		//
		// Solidity: function name() view returns(string)
		func (_ERC721Full *ERC721FullCallerSession) Name() ( string,  error) {
		  return _ERC721Full.Contract.Name(&_ERC721Full.CallOpts )
		}
	
		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address)
		func (_ERC721Full *ERC721FullCaller) OwnerOf(opts *bind.CallOpts , tokenId *big.Int ) (common.Address, error) {
			var (
				ret0 = new(common.Address)
				
			)
			out := ret0
			err := _ERC721Full.contract.Call(opts, out, "ownerOf" , tokenId)
			return *ret0, err
		}

		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address)
		func (_ERC721Full *ERC721FullSession) OwnerOf( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721Full.Contract.OwnerOf(&_ERC721Full.CallOpts , tokenId)
		}

		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address)
		func (_ERC721Full *ERC721FullCallerSession) OwnerOf( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721Full.Contract.OwnerOf(&_ERC721Full.CallOpts , tokenId)
		}
	
		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_ERC721Full *ERC721FullCaller) SupportsInterface(opts *bind.CallOpts , interfaceId [4]byte ) (bool, error) {
			var (
				ret0 = new(bool)
				
			)
			out := ret0
			err := _ERC721Full.contract.Call(opts, out, "supportsInterface" , interfaceId)
			return *ret0, err
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_ERC721Full *ERC721FullSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _ERC721Full.Contract.SupportsInterface(&_ERC721Full.CallOpts , interfaceId)
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_ERC721Full *ERC721FullCallerSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _ERC721Full.Contract.SupportsInterface(&_ERC721Full.CallOpts , interfaceId)
		}
	
		// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
		//
		// Solidity: function symbol() view returns(string)
		func (_ERC721Full *ERC721FullCaller) Symbol(opts *bind.CallOpts ) (string, error) {
			var (
				ret0 = new(string)
				
			)
			out := ret0
			err := _ERC721Full.contract.Call(opts, out, "symbol" )
			return *ret0, err
		}

		// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
		//
		// Solidity: function symbol() view returns(string)
		func (_ERC721Full *ERC721FullSession) Symbol() ( string,  error) {
		  return _ERC721Full.Contract.Symbol(&_ERC721Full.CallOpts )
		}

		// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
		//
		// Solidity: function symbol() view returns(string)
		func (_ERC721Full *ERC721FullCallerSession) Symbol() ( string,  error) {
		  return _ERC721Full.Contract.Symbol(&_ERC721Full.CallOpts )
		}
	
		// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
		//
		// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
		func (_ERC721Full *ERC721FullCaller) TokenByIndex(opts *bind.CallOpts , index *big.Int ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _ERC721Full.contract.Call(opts, out, "tokenByIndex" , index)
			return *ret0, err
		}

		// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
		//
		// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
		func (_ERC721Full *ERC721FullSession) TokenByIndex( index *big.Int ) ( *big.Int,  error) {
		  return _ERC721Full.Contract.TokenByIndex(&_ERC721Full.CallOpts , index)
		}

		// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
		//
		// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
		func (_ERC721Full *ERC721FullCallerSession) TokenByIndex( index *big.Int ) ( *big.Int,  error) {
		  return _ERC721Full.Contract.TokenByIndex(&_ERC721Full.CallOpts , index)
		}
	
		// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
		//
		// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
		func (_ERC721Full *ERC721FullCaller) TokenOfOwnerByIndex(opts *bind.CallOpts , owner common.Address , index *big.Int ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _ERC721Full.contract.Call(opts, out, "tokenOfOwnerByIndex" , owner, index)
			return *ret0, err
		}

		// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
		//
		// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
		func (_ERC721Full *ERC721FullSession) TokenOfOwnerByIndex( owner common.Address , index *big.Int ) ( *big.Int,  error) {
		  return _ERC721Full.Contract.TokenOfOwnerByIndex(&_ERC721Full.CallOpts , owner, index)
		}

		// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
		//
		// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
		func (_ERC721Full *ERC721FullCallerSession) TokenOfOwnerByIndex( owner common.Address , index *big.Int ) ( *big.Int,  error) {
		  return _ERC721Full.Contract.TokenOfOwnerByIndex(&_ERC721Full.CallOpts , owner, index)
		}
	
		// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
		//
		// Solidity: function tokenURI(uint256 tokenId) view returns(string)
		func (_ERC721Full *ERC721FullCaller) TokenURI(opts *bind.CallOpts , tokenId *big.Int ) (string, error) {
			var (
				ret0 = new(string)
				
			)
			out := ret0
			err := _ERC721Full.contract.Call(opts, out, "tokenURI" , tokenId)
			return *ret0, err
		}

		// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
		//
		// Solidity: function tokenURI(uint256 tokenId) view returns(string)
		func (_ERC721Full *ERC721FullSession) TokenURI( tokenId *big.Int ) ( string,  error) {
		  return _ERC721Full.Contract.TokenURI(&_ERC721Full.CallOpts , tokenId)
		}

		// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
		//
		// Solidity: function tokenURI(uint256 tokenId) view returns(string)
		func (_ERC721Full *ERC721FullCallerSession) TokenURI( tokenId *big.Int ) ( string,  error) {
		  return _ERC721Full.Contract.TokenURI(&_ERC721Full.CallOpts , tokenId)
		}
	
		// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
		//
		// Solidity: function totalSupply() view returns(uint256)
		func (_ERC721Full *ERC721FullCaller) TotalSupply(opts *bind.CallOpts ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _ERC721Full.contract.Call(opts, out, "totalSupply" )
			return *ret0, err
		}

		// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
		//
		// Solidity: function totalSupply() view returns(uint256)
		func (_ERC721Full *ERC721FullSession) TotalSupply() ( *big.Int,  error) {
		  return _ERC721Full.Contract.TotalSupply(&_ERC721Full.CallOpts )
		}

		// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
		//
		// Solidity: function totalSupply() view returns(uint256)
		func (_ERC721Full *ERC721FullCallerSession) TotalSupply() ( *big.Int,  error) {
		  return _ERC721Full.Contract.TotalSupply(&_ERC721Full.CallOpts )
		}
	

	
		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_ERC721Full *ERC721FullTransactor) Approve(opts *bind.TransactOpts , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _ERC721Full.contract.Transact(opts, "approve" , to, tokenId)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_ERC721Full *ERC721FullSession) Approve( to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Full.Contract.Approve(&_ERC721Full.TransactOpts , to, tokenId)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_ERC721Full *ERC721FullTransactorSession) Approve( to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Full.Contract.Approve(&_ERC721Full.TransactOpts , to, tokenId)
		}
	
		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Full *ERC721FullTransactor) SafeTransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _ERC721Full.contract.Transact(opts, "safeTransferFrom" , from, to, tokenId)
		}

		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Full *ERC721FullSession) SafeTransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Full.Contract.SafeTransferFrom(&_ERC721Full.TransactOpts , from, to, tokenId)
		}

		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Full *ERC721FullTransactorSession) SafeTransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Full.Contract.SafeTransferFrom(&_ERC721Full.TransactOpts , from, to, tokenId)
		}
	
		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
		func (_ERC721Full *ERC721FullTransactor) SafeTransferFrom0(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int , _data []byte ) (*types.Transaction, error) {
			return _ERC721Full.contract.Transact(opts, "safeTransferFrom0" , from, to, tokenId, _data)
		}

		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
		func (_ERC721Full *ERC721FullSession) SafeTransferFrom0( from common.Address , to common.Address , tokenId *big.Int , _data []byte ) (*types.Transaction, error) {
		  return _ERC721Full.Contract.SafeTransferFrom0(&_ERC721Full.TransactOpts , from, to, tokenId, _data)
		}

		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
		func (_ERC721Full *ERC721FullTransactorSession) SafeTransferFrom0( from common.Address , to common.Address , tokenId *big.Int , _data []byte ) (*types.Transaction, error) {
		  return _ERC721Full.Contract.SafeTransferFrom0(&_ERC721Full.TransactOpts , from, to, tokenId, _data)
		}
	
		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address to, bool approved) returns()
		func (_ERC721Full *ERC721FullTransactor) SetApprovalForAll(opts *bind.TransactOpts , to common.Address , approved bool ) (*types.Transaction, error) {
			return _ERC721Full.contract.Transact(opts, "setApprovalForAll" , to, approved)
		}

		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address to, bool approved) returns()
		func (_ERC721Full *ERC721FullSession) SetApprovalForAll( to common.Address , approved bool ) (*types.Transaction, error) {
		  return _ERC721Full.Contract.SetApprovalForAll(&_ERC721Full.TransactOpts , to, approved)
		}

		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address to, bool approved) returns()
		func (_ERC721Full *ERC721FullTransactorSession) SetApprovalForAll( to common.Address , approved bool ) (*types.Transaction, error) {
		  return _ERC721Full.Contract.SetApprovalForAll(&_ERC721Full.TransactOpts , to, approved)
		}
	
		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Full *ERC721FullTransactor) TransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _ERC721Full.contract.Transact(opts, "transferFrom" , from, to, tokenId)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Full *ERC721FullSession) TransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Full.Contract.TransferFrom(&_ERC721Full.TransactOpts , from, to, tokenId)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Full *ERC721FullTransactorSession) TransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Full.Contract.TransferFrom(&_ERC721Full.TransactOpts , from, to, tokenId)
		}
	

	

	

	
		// ERC721FullApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Full contract.
		type ERC721FullApprovalIterator struct {
			Event *ERC721FullApproval // Event containing the contract specifics and raw log

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
		func (it *ERC721FullApprovalIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ERC721FullApproval)
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
				it.Event = new(ERC721FullApproval)
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
		func (it *ERC721FullApprovalIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ERC721FullApprovalIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ERC721FullApproval represents a Approval event raised by the ERC721Full contract.
		type ERC721FullApproval struct { 
			Owner common.Address; 
			Approved common.Address; 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
 		func (_ERC721Full *ERC721FullFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721FullApprovalIterator, error) {
			
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

			logs, sub, err := _ERC721Full.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &ERC721FullApprovalIterator{contract: _ERC721Full.contract, event: "Approval", logs: logs, sub: sub}, nil
 		}

		// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
		func (_ERC721Full *ERC721FullFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721FullApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {
			
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

			logs, sub, err := _ERC721Full.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ERC721FullApproval)
						if err := _ERC721Full.contract.UnpackLog(event, "Approval", log); err != nil {
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
		func (_ERC721Full *ERC721FullFilterer) ParseApproval(log types.Log) (*ERC721FullApproval, error) {
			event := new(ERC721FullApproval)
			if err := _ERC721Full.contract.UnpackLog(event, "Approval", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// ERC721FullApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Full contract.
		type ERC721FullApprovalForAllIterator struct {
			Event *ERC721FullApprovalForAll // Event containing the contract specifics and raw log

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
		func (it *ERC721FullApprovalForAllIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ERC721FullApprovalForAll)
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
				it.Event = new(ERC721FullApprovalForAll)
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
		func (it *ERC721FullApprovalForAllIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ERC721FullApprovalForAllIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ERC721FullApprovalForAll represents a ApprovalForAll event raised by the ERC721Full contract.
		type ERC721FullApprovalForAll struct { 
			Owner common.Address; 
			Operator common.Address; 
			Approved bool; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
 		func (_ERC721Full *ERC721FullFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721FullApprovalForAllIterator, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var operatorRule []interface{}
			for _, operatorItem := range operator {
				operatorRule = append(operatorRule, operatorItem)
			}
			

			logs, sub, err := _ERC721Full.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
			if err != nil {
				return nil, err
			}
			return &ERC721FullApprovalForAllIterator{contract: _ERC721Full.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
 		}

		// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
		func (_ERC721Full *ERC721FullFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721FullApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var operatorRule []interface{}
			for _, operatorItem := range operator {
				operatorRule = append(operatorRule, operatorItem)
			}
			

			logs, sub, err := _ERC721Full.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ERC721FullApprovalForAll)
						if err := _ERC721Full.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
		func (_ERC721Full *ERC721FullFilterer) ParseApprovalForAll(log types.Log) (*ERC721FullApprovalForAll, error) {
			event := new(ERC721FullApprovalForAll)
			if err := _ERC721Full.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// ERC721FullTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Full contract.
		type ERC721FullTransferIterator struct {
			Event *ERC721FullTransfer // Event containing the contract specifics and raw log

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
		func (it *ERC721FullTransferIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ERC721FullTransfer)
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
				it.Event = new(ERC721FullTransfer)
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
		func (it *ERC721FullTransferIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ERC721FullTransferIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ERC721FullTransfer represents a Transfer event raised by the ERC721Full contract.
		type ERC721FullTransfer struct { 
			From common.Address; 
			To common.Address; 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
 		func (_ERC721Full *ERC721FullFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721FullTransferIterator, error) {
			
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

			logs, sub, err := _ERC721Full.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &ERC721FullTransferIterator{contract: _ERC721Full.contract, event: "Transfer", logs: logs, sub: sub}, nil
 		}

		// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
		func (_ERC721Full *ERC721FullFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721FullTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {
			
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

			logs, sub, err := _ERC721Full.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ERC721FullTransfer)
						if err := _ERC721Full.contract.UnpackLog(event, "Transfer", log); err != nil {
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
		func (_ERC721Full *ERC721FullFilterer) ParseTransfer(log types.Log) (*ERC721FullTransfer, error) {
			event := new(ERC721FullTransfer)
			if err := _ERC721Full.contract.UnpackLog(event, "Transfer", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	

	// ERC721HolderABI is the input ABI used to generate the binding from.
	const ERC721HolderABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

	
		// ERC721HolderFuncSigs maps the 4-byte function signature to its string representation.
		var ERC721HolderFuncSigs = map[string]string{
			"150b7a02": "onERC721Received(address,address,uint256,bytes)",
			
		}
	

	
		// ERC721HolderBin is the compiled bytecode used for deploying new contracts.
		var ERC721HolderBin = "0x608060405234801561001057600080fd5b50610158806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063150b7a0214610030575b600080fd5b6100f66004803603608081101561004657600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561008157600080fd5b82018360208201111561009357600080fd5b803590602001918460018302840111640100000000831117156100b557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610113945050505050565b604080516001600160e01b03199092168252519081900360200190f35b630a85bd0160e11b94935050505056fea265627a7a72315820013da66d312370aa9c13e1a7188e545d286c74f28a5786f1d5874cd8f73e7df964736f6c63430005110032"

		// DeployERC721Holder deploys a new Ethereum contract, binding an instance of ERC721Holder to it.
		func DeployERC721Holder(auth *bind.TransactOpts, backend bind.ContractBackend ) (common.Address, *types.Transaction, *ERC721Holder, error) {
		  parsed, err := abi.JSON(strings.NewReader(ERC721HolderABI))
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  
		  address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721HolderBin), backend )
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  return address, tx, &ERC721Holder{ ERC721HolderCaller: ERC721HolderCaller{contract: contract}, ERC721HolderTransactor: ERC721HolderTransactor{contract: contract}, ERC721HolderFilterer: ERC721HolderFilterer{contract: contract} }, nil
		}
	

	// ERC721Holder is an auto generated Go binding around an Ethereum contract.
	type ERC721Holder struct {
	  ERC721HolderCaller     // Read-only binding to the contract
	  ERC721HolderTransactor // Write-only binding to the contract
	  ERC721HolderFilterer   // Log filterer for contract events
	}

	// ERC721HolderCaller is an auto generated read-only Go binding around an Ethereum contract.
	type ERC721HolderCaller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC721HolderTransactor is an auto generated write-only Go binding around an Ethereum contract.
	type ERC721HolderTransactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC721HolderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type ERC721HolderFilterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC721HolderSession is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type ERC721HolderSession struct {
	  Contract     *ERC721Holder        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// ERC721HolderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type ERC721HolderCallerSession struct {
	  Contract *ERC721HolderCaller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// ERC721HolderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type ERC721HolderTransactorSession struct {
	  Contract     *ERC721HolderTransactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// ERC721HolderRaw is an auto generated low-level Go binding around an Ethereum contract.
	type ERC721HolderRaw struct {
	  Contract *ERC721Holder // Generic contract binding to access the raw methods on
	}

	// ERC721HolderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type ERC721HolderCallerRaw struct {
		Contract *ERC721HolderCaller // Generic read-only contract binding to access the raw methods on
	}

	// ERC721HolderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type ERC721HolderTransactorRaw struct {
		Contract *ERC721HolderTransactor // Generic write-only contract binding to access the raw methods on
	}

	// NewERC721Holder creates a new instance of ERC721Holder, bound to a specific deployed contract.
	func NewERC721Holder(address common.Address, backend bind.ContractBackend) (*ERC721Holder, error) {
	  contract, err := bindERC721Holder(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC721Holder{ ERC721HolderCaller: ERC721HolderCaller{contract: contract}, ERC721HolderTransactor: ERC721HolderTransactor{contract: contract}, ERC721HolderFilterer: ERC721HolderFilterer{contract: contract} }, nil
	}

	// NewERC721HolderCaller creates a new read-only instance of ERC721Holder, bound to a specific deployed contract.
	func NewERC721HolderCaller(address common.Address, caller bind.ContractCaller) (*ERC721HolderCaller, error) {
	  contract, err := bindERC721Holder(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC721HolderCaller{contract: contract}, nil
	}

	// NewERC721HolderTransactor creates a new write-only instance of ERC721Holder, bound to a specific deployed contract.
	func NewERC721HolderTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721HolderTransactor, error) {
	  contract, err := bindERC721Holder(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC721HolderTransactor{contract: contract}, nil
	}

	// NewERC721HolderFilterer creates a new log filterer instance of ERC721Holder, bound to a specific deployed contract.
 	func NewERC721HolderFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721HolderFilterer, error) {
 	  contract, err := bindERC721Holder(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &ERC721HolderFilterer{contract: contract}, nil
 	}

	// bindERC721Holder binds a generic wrapper to an already deployed contract.
	func bindERC721Holder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(ERC721HolderABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_ERC721Holder *ERC721HolderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _ERC721Holder.Contract.ERC721HolderCaller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_ERC721Holder *ERC721HolderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _ERC721Holder.Contract.ERC721HolderTransactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_ERC721Holder *ERC721HolderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _ERC721Holder.Contract.ERC721HolderTransactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_ERC721Holder *ERC721HolderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _ERC721Holder.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_ERC721Holder *ERC721HolderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _ERC721Holder.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_ERC721Holder *ERC721HolderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _ERC721Holder.Contract.contract.Transact(opts, method, params...)
	}

	

	
		// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
		//
		// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
		func (_ERC721Holder *ERC721HolderTransactor) OnERC721Received(opts *bind.TransactOpts , arg0 common.Address , arg1 common.Address , arg2 *big.Int , arg3 []byte ) (*types.Transaction, error) {
			return _ERC721Holder.contract.Transact(opts, "onERC721Received" , arg0, arg1, arg2, arg3)
		}

		// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
		//
		// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
		func (_ERC721Holder *ERC721HolderSession) OnERC721Received( arg0 common.Address , arg1 common.Address , arg2 *big.Int , arg3 []byte ) (*types.Transaction, error) {
		  return _ERC721Holder.Contract.OnERC721Received(&_ERC721Holder.TransactOpts , arg0, arg1, arg2, arg3)
		}

		// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
		//
		// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
		func (_ERC721Holder *ERC721HolderTransactorSession) OnERC721Received( arg0 common.Address , arg1 common.Address , arg2 *big.Int , arg3 []byte ) (*types.Transaction, error) {
		  return _ERC721Holder.Contract.OnERC721Received(&_ERC721Holder.TransactOpts , arg0, arg1, arg2, arg3)
		}
	

	

	

	

	// ERC721MetadataABI is the input ABI used to generate the binding from.
	const ERC721MetadataABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

	
		// ERC721MetadataFuncSigs maps the 4-byte function signature to its string representation.
		var ERC721MetadataFuncSigs = map[string]string{
			"095ea7b3": "approve(address,uint256)",
			"70a08231": "balanceOf(address)",
			"6c0360eb": "baseURI()",
			"081812fc": "getApproved(uint256)",
			"e985e9c5": "isApprovedForAll(address,address)",
			"06fdde03": "name()",
			"6352211e": "ownerOf(uint256)",
			"42842e0e": "safeTransferFrom(address,address,uint256)",
			"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
			"a22cb465": "setApprovalForAll(address,bool)",
			"01ffc9a7": "supportsInterface(bytes4)",
			"95d89b41": "symbol()",
			"c87b56dd": "tokenURI(uint256)",
			"23b872dd": "transferFrom(address,address,uint256)",
			
		}
	

	
		// ERC721MetadataBin is the compiled bytecode used for deploying new contracts.
		var ERC721MetadataBin = "0x60806040523480156200001157600080fd5b506040516200174138038062001741833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b5060405250620001bc91506301ffc9a760e01b90506001600160e01b036200022616565b620001d76380ac58cd60e01b6001600160e01b036200022616565b8151620001ec906005906020850190620002ab565b50805162000202906006906020840190620002ab565b506200021e635b5e139f60e01b6001600160e01b036200022616565b505062000350565b6001600160e01b0319808216141562000286576040805162461bcd60e51b815260206004820152601c60248201527f4552433136353a20696e76616c696420696e7465726661636520696400000000604482015290519081900360640190fd5b6001600160e01b0319166000908152602081905260409020805460ff19166001179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002ee57805160ff19168380011785556200031e565b828001600101855582156200031e579182015b828111156200031e57825182559160200191906001019062000301565b506200032c92915062000330565b5090565b6200034d91905b808211156200032c576000815560010162000337565b90565b6113e180620003606000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80636c0360eb1161008c578063a22cb46511610066578063a22cb465146102df578063b88d4fde1461030d578063c87b56dd146103d3578063e985e9c5146103f0576100ea565b80636c0360eb1461029757806370a082311461029f57806395d89b41146102d7576100ea565b8063095ea7b3116100c8578063095ea7b3146101e057806323b872dd1461020e57806342842e0e146102445780636352211e1461027a576100ea565b806301ffc9a7146100ef57806306fdde031461012a578063081812fc146101a7575b600080fd5b6101166004803603602081101561010557600080fd5b50356001600160e01b03191661041e565b604080519115158252519081900360200190f35b610132610441565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561016c578181015183820152602001610154565b50505050905090810190601f1680156101995780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101c4600480360360208110156101bd57600080fd5b50356104d7565b604080516001600160a01b039092168252519081900360200190f35b61020c600480360360408110156101f657600080fd5b506001600160a01b038135169060200135610539565b005b61020c6004803603606081101561022457600080fd5b506001600160a01b03813581169160208101359091169060400135610661565b61020c6004803603606081101561025a57600080fd5b506001600160a01b038135811691602081013590911690604001356106bd565b6101c46004803603602081101561029057600080fd5b50356106d8565b610132610732565b6102c5600480360360208110156102b557600080fd5b50356001600160a01b0316610793565b60408051918252519081900360200190f35b6101326107fb565b61020c600480360360408110156102f557600080fd5b506001600160a01b038135169060200135151561085c565b61020c6004803603608081101561032357600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561035e57600080fd5b82018360208201111561037057600080fd5b8035906020019184600183028401116401000000008311171561039257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610961945050505050565b610132600480360360208110156103e957600080fd5b50356109bf565b6101166004803603604081101561040657600080fd5b506001600160a01b0381358116916020013516610b8b565b6001600160e01b0319811660009081526020819052604090205460ff165b919050565b60058054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104cd5780601f106104a2576101008083540402835291602001916104cd565b820191906000526020600020905b8154815290600101906020018083116104b057829003601f168201915b5050505050905090565b60006104e282610bb9565b61051d5760405162461bcd60e51b815260040180806020018281038252602c8152602001806112d7602c913960400191505060405180910390fd5b506000908152600260205260409020546001600160a01b031690565b6000610544826106d8565b9050806001600160a01b0316836001600160a01b031614156105975760405162461bcd60e51b815260040180806020018281038252602181526020018061135b6021913960400191505060405180910390fd5b806001600160a01b03166105a9610bd6565b6001600160a01b031614806105ca57506105ca816105c5610bd6565b610b8b565b6106055760405162461bcd60e51b815260040180806020018281038252603881526020018061124c6038913960400191505060405180910390fd5b60008281526002602052604080822080546001600160a01b0319166001600160a01b0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b61067261066c610bd6565b82610bda565b6106ad5760405162461bcd60e51b815260040180806020018281038252603181526020018061137c6031913960400191505060405180910390fd5b6106b8838383610c7e565b505050565b6106b883838360405180602001604052806000815250610961565b6000818152600160205260408120546001600160a01b03168061072c5760405162461bcd60e51b81526004018080602001828103825260298152602001806112ae6029913960400191505060405180910390fd5b92915050565b60078054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104cd5780601f106104a2576101008083540402835291602001916104cd565b60006001600160a01b0382166107da5760405162461bcd60e51b815260040180806020018281038252602a815260200180611284602a913960400191505060405180910390fd5b6001600160a01b038216600090815260036020526040902061072c90610dc2565b60068054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104cd5780601f106104a2576101008083540402835291602001916104cd565b610864610bd6565b6001600160a01b0316826001600160a01b031614156108ca576040805162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c657200000000000000604482015290519081900360640190fd5b80600460006108d7610bd6565b6001600160a01b03908116825260208083019390935260409182016000908120918716808252919093529120805460ff19169215159290921790915561091b610bd6565b60408051841515815290516001600160a01b0392909216917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c319181900360200190a35050565b61097261096c610bd6565b83610bda565b6109ad5760405162461bcd60e51b815260040180806020018281038252603181526020018061137c6031913960400191505060405180910390fd5b6109b984848484610dc6565b50505050565b60606109ca82610bb9565b610a055760405162461bcd60e51b815260040180806020018281038252602f81526020018061132c602f913960400191505060405180910390fd5b60008281526008602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845260609392830182828015610a9a5780601f10610a6f57610100808354040283529160200191610a9a565b820191906000526020600020905b815481529060010190602001808311610a7d57829003601f168201915b50505050509050805160001415610ac157505060408051602081019091526000815261043c565b6007816040516020018083805460018160011615610100020316600290048015610b225780601f10610b00576101008083540402835291820191610b22565b820191906000526020600020905b815481529060010190602001808311610b0e575b5050825160208401908083835b60208310610b4e5780518252601f199092019160209182019101610b2f565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405291505061043c565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000908152600160205260409020546001600160a01b0316151590565b3390565b6000610be582610bb9565b610c205760405162461bcd60e51b815260040180806020018281038252602c815260200180611220602c913960400191505060405180910390fd5b6000610c2b836106d8565b9050806001600160a01b0316846001600160a01b03161480610c665750836001600160a01b0316610c5b846104d7565b6001600160a01b0316145b80610c765750610c768185610b8b565b949350505050565b826001600160a01b0316610c91826106d8565b6001600160a01b031614610cd65760405162461bcd60e51b81526004018080602001828103825260298152602001806113036029913960400191505060405180910390fd5b6001600160a01b038216610d1b5760405162461bcd60e51b81526004018080602001828103825260248152602001806111fc6024913960400191505060405180910390fd5b610d2481610e18565b6001600160a01b0383166000908152600360205260409020610d4590610e55565b6001600160a01b0382166000908152600360205260409020610d6690610e6c565b60008181526001602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b5490565b610dd1848484610c7e565b610ddd84848484610e75565b6109b95760405162461bcd60e51b81526004018080602001828103825260328152602001806111ca6032913960400191505060405180910390fd5b6000818152600260205260409020546001600160a01b031615610e5257600081815260026020526040902080546001600160a01b03191690555b50565b8054610e6890600163ffffffff6110b016565b9055565b80546001019055565b6000610e89846001600160a01b03166110f9565b610e9557506001610c76565b600060606001600160a01b038616630a85bd0160e11b610eb3610bd6565b89888860405160240180856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610f2c578181015183820152602001610f14565b50505050905090810190601f168015610f595780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909a16999099178952518151919890975087965094509250829150849050835b60208310610fc15780518252601f199092019160209182019101610fa2565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611023576040519150601f19603f3d011682016040523d82523d6000602084013e611028565b606091505b509150915081611079578051156110425780518082602001fd5b60405162461bcd60e51b81526004018080602001828103825260328152602001806111ca6032913960400191505060405180910390fd5b600081806020019051602081101561109057600080fd5b50516001600160e01b031916630a85bd0160e11b149350610c7692505050565b60006110f283836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611132565b9392505050565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470818114801590610c76575050151592915050565b600081848411156111c15760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561118657818101518382015260200161116e565b50505050905090810190601f1680156111b35780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50505090039056fe4552433732313a207472616e7366657220746f206e6f6e20455243373231526563656976657220696d706c656d656e7465724552433732313a207472616e7366657220746f20746865207a65726f20616464726573734552433732313a206f70657261746f7220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76652063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c4552433732313a2062616c616e636520717565727920666f7220746865207a65726f20616464726573734552433732313a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76656420717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a207472616e73666572206f6620746f6b656e2074686174206973206e6f74206f776e4552433732314d657461646174613a2055524920717565727920666f72206e6f6e6578697374656e7420746f6b656e4552433732313a20617070726f76616c20746f2063757272656e74206f776e65724552433732313a207472616e736665722063616c6c6572206973206e6f74206f776e6572206e6f7220617070726f766564a265627a7a72315820c5f72d148dd541a46706ce435f70c50f5270db86750b1e2c0a36bd4d2e3bc36b64736f6c63430005110032"

		// DeployERC721Metadata deploys a new Ethereum contract, binding an instance of ERC721Metadata to it.
		func DeployERC721Metadata(auth *bind.TransactOpts, backend bind.ContractBackend , name string, symbol string) (common.Address, *types.Transaction, *ERC721Metadata, error) {
		  parsed, err := abi.JSON(strings.NewReader(ERC721MetadataABI))
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  
		  address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721MetadataBin), backend , name, symbol)
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  return address, tx, &ERC721Metadata{ ERC721MetadataCaller: ERC721MetadataCaller{contract: contract}, ERC721MetadataTransactor: ERC721MetadataTransactor{contract: contract}, ERC721MetadataFilterer: ERC721MetadataFilterer{contract: contract} }, nil
		}
	

	// ERC721Metadata is an auto generated Go binding around an Ethereum contract.
	type ERC721Metadata struct {
	  ERC721MetadataCaller     // Read-only binding to the contract
	  ERC721MetadataTransactor // Write-only binding to the contract
	  ERC721MetadataFilterer   // Log filterer for contract events
	}

	// ERC721MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
	type ERC721MetadataCaller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC721MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
	type ERC721MetadataTransactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC721MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type ERC721MetadataFilterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// ERC721MetadataSession is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type ERC721MetadataSession struct {
	  Contract     *ERC721Metadata        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// ERC721MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type ERC721MetadataCallerSession struct {
	  Contract *ERC721MetadataCaller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// ERC721MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type ERC721MetadataTransactorSession struct {
	  Contract     *ERC721MetadataTransactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// ERC721MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
	type ERC721MetadataRaw struct {
	  Contract *ERC721Metadata // Generic contract binding to access the raw methods on
	}

	// ERC721MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type ERC721MetadataCallerRaw struct {
		Contract *ERC721MetadataCaller // Generic read-only contract binding to access the raw methods on
	}

	// ERC721MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type ERC721MetadataTransactorRaw struct {
		Contract *ERC721MetadataTransactor // Generic write-only contract binding to access the raw methods on
	}

	// NewERC721Metadata creates a new instance of ERC721Metadata, bound to a specific deployed contract.
	func NewERC721Metadata(address common.Address, backend bind.ContractBackend) (*ERC721Metadata, error) {
	  contract, err := bindERC721Metadata(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC721Metadata{ ERC721MetadataCaller: ERC721MetadataCaller{contract: contract}, ERC721MetadataTransactor: ERC721MetadataTransactor{contract: contract}, ERC721MetadataFilterer: ERC721MetadataFilterer{contract: contract} }, nil
	}

	// NewERC721MetadataCaller creates a new read-only instance of ERC721Metadata, bound to a specific deployed contract.
	func NewERC721MetadataCaller(address common.Address, caller bind.ContractCaller) (*ERC721MetadataCaller, error) {
	  contract, err := bindERC721Metadata(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC721MetadataCaller{contract: contract}, nil
	}

	// NewERC721MetadataTransactor creates a new write-only instance of ERC721Metadata, bound to a specific deployed contract.
	func NewERC721MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721MetadataTransactor, error) {
	  contract, err := bindERC721Metadata(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &ERC721MetadataTransactor{contract: contract}, nil
	}

	// NewERC721MetadataFilterer creates a new log filterer instance of ERC721Metadata, bound to a specific deployed contract.
 	func NewERC721MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721MetadataFilterer, error) {
 	  contract, err := bindERC721Metadata(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &ERC721MetadataFilterer{contract: contract}, nil
 	}

	// bindERC721Metadata binds a generic wrapper to an already deployed contract.
	func bindERC721Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(ERC721MetadataABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_ERC721Metadata *ERC721MetadataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _ERC721Metadata.Contract.ERC721MetadataCaller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_ERC721Metadata *ERC721MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _ERC721Metadata.Contract.ERC721MetadataTransactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_ERC721Metadata *ERC721MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _ERC721Metadata.Contract.ERC721MetadataTransactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_ERC721Metadata *ERC721MetadataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _ERC721Metadata.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_ERC721Metadata *ERC721MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _ERC721Metadata.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_ERC721Metadata *ERC721MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _ERC721Metadata.Contract.contract.Transact(opts, method, params...)
	}

	
		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256)
		func (_ERC721Metadata *ERC721MetadataCaller) BalanceOf(opts *bind.CallOpts , owner common.Address ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _ERC721Metadata.contract.Call(opts, out, "balanceOf" , owner)
			return *ret0, err
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256)
		func (_ERC721Metadata *ERC721MetadataSession) BalanceOf( owner common.Address ) ( *big.Int,  error) {
		  return _ERC721Metadata.Contract.BalanceOf(&_ERC721Metadata.CallOpts , owner)
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256)
		func (_ERC721Metadata *ERC721MetadataCallerSession) BalanceOf( owner common.Address ) ( *big.Int,  error) {
		  return _ERC721Metadata.Contract.BalanceOf(&_ERC721Metadata.CallOpts , owner)
		}
	
		// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
		//
		// Solidity: function baseURI() view returns(string)
		func (_ERC721Metadata *ERC721MetadataCaller) BaseURI(opts *bind.CallOpts ) (string, error) {
			var (
				ret0 = new(string)
				
			)
			out := ret0
			err := _ERC721Metadata.contract.Call(opts, out, "baseURI" )
			return *ret0, err
		}

		// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
		//
		// Solidity: function baseURI() view returns(string)
		func (_ERC721Metadata *ERC721MetadataSession) BaseURI() ( string,  error) {
		  return _ERC721Metadata.Contract.BaseURI(&_ERC721Metadata.CallOpts )
		}

		// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
		//
		// Solidity: function baseURI() view returns(string)
		func (_ERC721Metadata *ERC721MetadataCallerSession) BaseURI() ( string,  error) {
		  return _ERC721Metadata.Contract.BaseURI(&_ERC721Metadata.CallOpts )
		}
	
		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address)
		func (_ERC721Metadata *ERC721MetadataCaller) GetApproved(opts *bind.CallOpts , tokenId *big.Int ) (common.Address, error) {
			var (
				ret0 = new(common.Address)
				
			)
			out := ret0
			err := _ERC721Metadata.contract.Call(opts, out, "getApproved" , tokenId)
			return *ret0, err
		}

		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address)
		func (_ERC721Metadata *ERC721MetadataSession) GetApproved( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721Metadata.Contract.GetApproved(&_ERC721Metadata.CallOpts , tokenId)
		}

		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address)
		func (_ERC721Metadata *ERC721MetadataCallerSession) GetApproved( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721Metadata.Contract.GetApproved(&_ERC721Metadata.CallOpts , tokenId)
		}
	
		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_ERC721Metadata *ERC721MetadataCaller) IsApprovedForAll(opts *bind.CallOpts , owner common.Address , operator common.Address ) (bool, error) {
			var (
				ret0 = new(bool)
				
			)
			out := ret0
			err := _ERC721Metadata.contract.Call(opts, out, "isApprovedForAll" , owner, operator)
			return *ret0, err
		}

		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_ERC721Metadata *ERC721MetadataSession) IsApprovedForAll( owner common.Address , operator common.Address ) ( bool,  error) {
		  return _ERC721Metadata.Contract.IsApprovedForAll(&_ERC721Metadata.CallOpts , owner, operator)
		}

		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_ERC721Metadata *ERC721MetadataCallerSession) IsApprovedForAll( owner common.Address , operator common.Address ) ( bool,  error) {
		  return _ERC721Metadata.Contract.IsApprovedForAll(&_ERC721Metadata.CallOpts , owner, operator)
		}
	
		// Name is a free data retrieval call binding the contract method 0x06fdde03.
		//
		// Solidity: function name() view returns(string)
		func (_ERC721Metadata *ERC721MetadataCaller) Name(opts *bind.CallOpts ) (string, error) {
			var (
				ret0 = new(string)
				
			)
			out := ret0
			err := _ERC721Metadata.contract.Call(opts, out, "name" )
			return *ret0, err
		}

		// Name is a free data retrieval call binding the contract method 0x06fdde03.
		//
		// Solidity: function name() view returns(string)
		func (_ERC721Metadata *ERC721MetadataSession) Name() ( string,  error) {
		  return _ERC721Metadata.Contract.Name(&_ERC721Metadata.CallOpts )
		}

		// Name is a free data retrieval call binding the contract method 0x06fdde03.
		//
		// Solidity: function name() view returns(string)
		func (_ERC721Metadata *ERC721MetadataCallerSession) Name() ( string,  error) {
		  return _ERC721Metadata.Contract.Name(&_ERC721Metadata.CallOpts )
		}
	
		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address)
		func (_ERC721Metadata *ERC721MetadataCaller) OwnerOf(opts *bind.CallOpts , tokenId *big.Int ) (common.Address, error) {
			var (
				ret0 = new(common.Address)
				
			)
			out := ret0
			err := _ERC721Metadata.contract.Call(opts, out, "ownerOf" , tokenId)
			return *ret0, err
		}

		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address)
		func (_ERC721Metadata *ERC721MetadataSession) OwnerOf( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721Metadata.Contract.OwnerOf(&_ERC721Metadata.CallOpts , tokenId)
		}

		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address)
		func (_ERC721Metadata *ERC721MetadataCallerSession) OwnerOf( tokenId *big.Int ) ( common.Address,  error) {
		  return _ERC721Metadata.Contract.OwnerOf(&_ERC721Metadata.CallOpts , tokenId)
		}
	
		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_ERC721Metadata *ERC721MetadataCaller) SupportsInterface(opts *bind.CallOpts , interfaceId [4]byte ) (bool, error) {
			var (
				ret0 = new(bool)
				
			)
			out := ret0
			err := _ERC721Metadata.contract.Call(opts, out, "supportsInterface" , interfaceId)
			return *ret0, err
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_ERC721Metadata *ERC721MetadataSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _ERC721Metadata.Contract.SupportsInterface(&_ERC721Metadata.CallOpts , interfaceId)
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_ERC721Metadata *ERC721MetadataCallerSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _ERC721Metadata.Contract.SupportsInterface(&_ERC721Metadata.CallOpts , interfaceId)
		}
	
		// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
		//
		// Solidity: function symbol() view returns(string)
		func (_ERC721Metadata *ERC721MetadataCaller) Symbol(opts *bind.CallOpts ) (string, error) {
			var (
				ret0 = new(string)
				
			)
			out := ret0
			err := _ERC721Metadata.contract.Call(opts, out, "symbol" )
			return *ret0, err
		}

		// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
		//
		// Solidity: function symbol() view returns(string)
		func (_ERC721Metadata *ERC721MetadataSession) Symbol() ( string,  error) {
		  return _ERC721Metadata.Contract.Symbol(&_ERC721Metadata.CallOpts )
		}

		// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
		//
		// Solidity: function symbol() view returns(string)
		func (_ERC721Metadata *ERC721MetadataCallerSession) Symbol() ( string,  error) {
		  return _ERC721Metadata.Contract.Symbol(&_ERC721Metadata.CallOpts )
		}
	
		// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
		//
		// Solidity: function tokenURI(uint256 tokenId) view returns(string)
		func (_ERC721Metadata *ERC721MetadataCaller) TokenURI(opts *bind.CallOpts , tokenId *big.Int ) (string, error) {
			var (
				ret0 = new(string)
				
			)
			out := ret0
			err := _ERC721Metadata.contract.Call(opts, out, "tokenURI" , tokenId)
			return *ret0, err
		}

		// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
		//
		// Solidity: function tokenURI(uint256 tokenId) view returns(string)
		func (_ERC721Metadata *ERC721MetadataSession) TokenURI( tokenId *big.Int ) ( string,  error) {
		  return _ERC721Metadata.Contract.TokenURI(&_ERC721Metadata.CallOpts , tokenId)
		}

		// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
		//
		// Solidity: function tokenURI(uint256 tokenId) view returns(string)
		func (_ERC721Metadata *ERC721MetadataCallerSession) TokenURI( tokenId *big.Int ) ( string,  error) {
		  return _ERC721Metadata.Contract.TokenURI(&_ERC721Metadata.CallOpts , tokenId)
		}
	

	
		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_ERC721Metadata *ERC721MetadataTransactor) Approve(opts *bind.TransactOpts , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _ERC721Metadata.contract.Transact(opts, "approve" , to, tokenId)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_ERC721Metadata *ERC721MetadataSession) Approve( to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Metadata.Contract.Approve(&_ERC721Metadata.TransactOpts , to, tokenId)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_ERC721Metadata *ERC721MetadataTransactorSession) Approve( to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Metadata.Contract.Approve(&_ERC721Metadata.TransactOpts , to, tokenId)
		}
	
		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Metadata *ERC721MetadataTransactor) SafeTransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _ERC721Metadata.contract.Transact(opts, "safeTransferFrom" , from, to, tokenId)
		}

		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Metadata *ERC721MetadataSession) SafeTransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Metadata.Contract.SafeTransferFrom(&_ERC721Metadata.TransactOpts , from, to, tokenId)
		}

		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Metadata *ERC721MetadataTransactorSession) SafeTransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Metadata.Contract.SafeTransferFrom(&_ERC721Metadata.TransactOpts , from, to, tokenId)
		}
	
		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
		func (_ERC721Metadata *ERC721MetadataTransactor) SafeTransferFrom0(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int , _data []byte ) (*types.Transaction, error) {
			return _ERC721Metadata.contract.Transact(opts, "safeTransferFrom0" , from, to, tokenId, _data)
		}

		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
		func (_ERC721Metadata *ERC721MetadataSession) SafeTransferFrom0( from common.Address , to common.Address , tokenId *big.Int , _data []byte ) (*types.Transaction, error) {
		  return _ERC721Metadata.Contract.SafeTransferFrom0(&_ERC721Metadata.TransactOpts , from, to, tokenId, _data)
		}

		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
		func (_ERC721Metadata *ERC721MetadataTransactorSession) SafeTransferFrom0( from common.Address , to common.Address , tokenId *big.Int , _data []byte ) (*types.Transaction, error) {
		  return _ERC721Metadata.Contract.SafeTransferFrom0(&_ERC721Metadata.TransactOpts , from, to, tokenId, _data)
		}
	
		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address to, bool approved) returns()
		func (_ERC721Metadata *ERC721MetadataTransactor) SetApprovalForAll(opts *bind.TransactOpts , to common.Address , approved bool ) (*types.Transaction, error) {
			return _ERC721Metadata.contract.Transact(opts, "setApprovalForAll" , to, approved)
		}

		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address to, bool approved) returns()
		func (_ERC721Metadata *ERC721MetadataSession) SetApprovalForAll( to common.Address , approved bool ) (*types.Transaction, error) {
		  return _ERC721Metadata.Contract.SetApprovalForAll(&_ERC721Metadata.TransactOpts , to, approved)
		}

		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address to, bool approved) returns()
		func (_ERC721Metadata *ERC721MetadataTransactorSession) SetApprovalForAll( to common.Address , approved bool ) (*types.Transaction, error) {
		  return _ERC721Metadata.Contract.SetApprovalForAll(&_ERC721Metadata.TransactOpts , to, approved)
		}
	
		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Metadata *ERC721MetadataTransactor) TransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _ERC721Metadata.contract.Transact(opts, "transferFrom" , from, to, tokenId)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Metadata *ERC721MetadataSession) TransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Metadata.Contract.TransferFrom(&_ERC721Metadata.TransactOpts , from, to, tokenId)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_ERC721Metadata *ERC721MetadataTransactorSession) TransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _ERC721Metadata.Contract.TransferFrom(&_ERC721Metadata.TransactOpts , from, to, tokenId)
		}
	

	

	

	
		// ERC721MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Metadata contract.
		type ERC721MetadataApprovalIterator struct {
			Event *ERC721MetadataApproval // Event containing the contract specifics and raw log

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
		func (it *ERC721MetadataApprovalIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ERC721MetadataApproval)
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
				it.Event = new(ERC721MetadataApproval)
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
		func (it *ERC721MetadataApprovalIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ERC721MetadataApprovalIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ERC721MetadataApproval represents a Approval event raised by the ERC721Metadata contract.
		type ERC721MetadataApproval struct { 
			Owner common.Address; 
			Approved common.Address; 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
 		func (_ERC721Metadata *ERC721MetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ERC721MetadataApprovalIterator, error) {
			
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

			logs, sub, err := _ERC721Metadata.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &ERC721MetadataApprovalIterator{contract: _ERC721Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
 		}

		// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
		func (_ERC721Metadata *ERC721MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721MetadataApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {
			
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

			logs, sub, err := _ERC721Metadata.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ERC721MetadataApproval)
						if err := _ERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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
		func (_ERC721Metadata *ERC721MetadataFilterer) ParseApproval(log types.Log) (*ERC721MetadataApproval, error) {
			event := new(ERC721MetadataApproval)
			if err := _ERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// ERC721MetadataApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Metadata contract.
		type ERC721MetadataApprovalForAllIterator struct {
			Event *ERC721MetadataApprovalForAll // Event containing the contract specifics and raw log

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
		func (it *ERC721MetadataApprovalForAllIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ERC721MetadataApprovalForAll)
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
				it.Event = new(ERC721MetadataApprovalForAll)
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
		func (it *ERC721MetadataApprovalForAllIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ERC721MetadataApprovalForAllIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ERC721MetadataApprovalForAll represents a ApprovalForAll event raised by the ERC721Metadata contract.
		type ERC721MetadataApprovalForAll struct { 
			Owner common.Address; 
			Operator common.Address; 
			Approved bool; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
 		func (_ERC721Metadata *ERC721MetadataFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ERC721MetadataApprovalForAllIterator, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var operatorRule []interface{}
			for _, operatorItem := range operator {
				operatorRule = append(operatorRule, operatorItem)
			}
			

			logs, sub, err := _ERC721Metadata.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
			if err != nil {
				return nil, err
			}
			return &ERC721MetadataApprovalForAllIterator{contract: _ERC721Metadata.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
 		}

		// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
		func (_ERC721Metadata *ERC721MetadataFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721MetadataApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var operatorRule []interface{}
			for _, operatorItem := range operator {
				operatorRule = append(operatorRule, operatorItem)
			}
			

			logs, sub, err := _ERC721Metadata.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ERC721MetadataApprovalForAll)
						if err := _ERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
		func (_ERC721Metadata *ERC721MetadataFilterer) ParseApprovalForAll(log types.Log) (*ERC721MetadataApprovalForAll, error) {
			event := new(ERC721MetadataApprovalForAll)
			if err := _ERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// ERC721MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Metadata contract.
		type ERC721MetadataTransferIterator struct {
			Event *ERC721MetadataTransfer // Event containing the contract specifics and raw log

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
		func (it *ERC721MetadataTransferIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(ERC721MetadataTransfer)
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
				it.Event = new(ERC721MetadataTransfer)
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
		func (it *ERC721MetadataTransferIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *ERC721MetadataTransferIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// ERC721MetadataTransfer represents a Transfer event raised by the ERC721Metadata contract.
		type ERC721MetadataTransfer struct { 
			From common.Address; 
			To common.Address; 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
 		func (_ERC721Metadata *ERC721MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721MetadataTransferIterator, error) {
			
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

			logs, sub, err := _ERC721Metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &ERC721MetadataTransferIterator{contract: _ERC721Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
 		}

		// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
		func (_ERC721Metadata *ERC721MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721MetadataTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {
			
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

			logs, sub, err := _ERC721Metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(ERC721MetadataTransfer)
						if err := _ERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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
		func (_ERC721Metadata *ERC721MetadataFilterer) ParseTransfer(log types.Log) (*ERC721MetadataTransfer, error) {
			event := new(ERC721MetadataTransfer)
			if err := _ERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	

	// IERC165ABI is the input ABI used to generate the binding from.
	const IERC165ABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

	
		// IERC165FuncSigs maps the 4-byte function signature to its string representation.
		var IERC165FuncSigs = map[string]string{
			"01ffc9a7": "supportsInterface(bytes4)",
			
		}
	

	

	// IERC165 is an auto generated Go binding around an Ethereum contract.
	type IERC165 struct {
	  IERC165Caller     // Read-only binding to the contract
	  IERC165Transactor // Write-only binding to the contract
	  IERC165Filterer   // Log filterer for contract events
	}

	// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
	type IERC165Caller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
	type IERC165Transactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type IERC165Filterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// IERC165Session is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type IERC165Session struct {
	  Contract     *IERC165        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type IERC165CallerSession struct {
	  Contract *IERC165Caller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type IERC165TransactorSession struct {
	  Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
	type IERC165Raw struct {
	  Contract *IERC165 // Generic contract binding to access the raw methods on
	}

	// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type IERC165CallerRaw struct {
		Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
	}

	// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type IERC165TransactorRaw struct {
		Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
	}

	// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
	func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	  contract, err := bindIERC165(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &IERC165{ IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract} }, nil
	}

	// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
	func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	  contract, err := bindIERC165(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &IERC165Caller{contract: contract}, nil
	}

	// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
	func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	  contract, err := bindIERC165(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &IERC165Transactor{contract: contract}, nil
	}

	// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
 	func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
 	  contract, err := bindIERC165(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &IERC165Filterer{contract: contract}, nil
 	}

	// bindIERC165 binds a generic wrapper to an already deployed contract.
	func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _IERC165.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _IERC165.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _IERC165.Contract.contract.Transact(opts, method, params...)
	}

	
		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts , interfaceId [4]byte ) (bool, error) {
			var (
				ret0 = new(bool)
				
			)
			out := ret0
			err := _IERC165.contract.Call(opts, out, "supportsInterface" , interfaceId)
			return *ret0, err
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_IERC165 *IERC165Session) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts , interfaceId)
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_IERC165 *IERC165CallerSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts , interfaceId)
		}
	

	

	

	

	

	// IERC721ABI is the input ABI used to generate the binding from.
	const IERC721ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

	
		// IERC721FuncSigs maps the 4-byte function signature to its string representation.
		var IERC721FuncSigs = map[string]string{
			"095ea7b3": "approve(address,uint256)",
			"70a08231": "balanceOf(address)",
			"081812fc": "getApproved(uint256)",
			"e985e9c5": "isApprovedForAll(address,address)",
			"6352211e": "ownerOf(uint256)",
			"42842e0e": "safeTransferFrom(address,address,uint256)",
			"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
			"a22cb465": "setApprovalForAll(address,bool)",
			"01ffc9a7": "supportsInterface(bytes4)",
			"23b872dd": "transferFrom(address,address,uint256)",
			
		}
	

	

	// IERC721 is an auto generated Go binding around an Ethereum contract.
	type IERC721 struct {
	  IERC721Caller     // Read-only binding to the contract
	  IERC721Transactor // Write-only binding to the contract
	  IERC721Filterer   // Log filterer for contract events
	}

	// IERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
	type IERC721Caller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// IERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
	type IERC721Transactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// IERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type IERC721Filterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// IERC721Session is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type IERC721Session struct {
	  Contract     *IERC721        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// IERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type IERC721CallerSession struct {
	  Contract *IERC721Caller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// IERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type IERC721TransactorSession struct {
	  Contract     *IERC721Transactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// IERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
	type IERC721Raw struct {
	  Contract *IERC721 // Generic contract binding to access the raw methods on
	}

	// IERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type IERC721CallerRaw struct {
		Contract *IERC721Caller // Generic read-only contract binding to access the raw methods on
	}

	// IERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type IERC721TransactorRaw struct {
		Contract *IERC721Transactor // Generic write-only contract binding to access the raw methods on
	}

	// NewIERC721 creates a new instance of IERC721, bound to a specific deployed contract.
	func NewIERC721(address common.Address, backend bind.ContractBackend) (*IERC721, error) {
	  contract, err := bindIERC721(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &IERC721{ IERC721Caller: IERC721Caller{contract: contract}, IERC721Transactor: IERC721Transactor{contract: contract}, IERC721Filterer: IERC721Filterer{contract: contract} }, nil
	}

	// NewIERC721Caller creates a new read-only instance of IERC721, bound to a specific deployed contract.
	func NewIERC721Caller(address common.Address, caller bind.ContractCaller) (*IERC721Caller, error) {
	  contract, err := bindIERC721(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &IERC721Caller{contract: contract}, nil
	}

	// NewIERC721Transactor creates a new write-only instance of IERC721, bound to a specific deployed contract.
	func NewIERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC721Transactor, error) {
	  contract, err := bindIERC721(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &IERC721Transactor{contract: contract}, nil
	}

	// NewIERC721Filterer creates a new log filterer instance of IERC721, bound to a specific deployed contract.
 	func NewIERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC721Filterer, error) {
 	  contract, err := bindIERC721(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &IERC721Filterer{contract: contract}, nil
 	}

	// bindIERC721 binds a generic wrapper to an already deployed contract.
	func bindIERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(IERC721ABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_IERC721 *IERC721Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _IERC721.Contract.IERC721Caller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_IERC721 *IERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _IERC721.Contract.IERC721Transactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_IERC721 *IERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _IERC721.Contract.IERC721Transactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_IERC721 *IERC721CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _IERC721.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_IERC721 *IERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _IERC721.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_IERC721 *IERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _IERC721.Contract.contract.Transact(opts, method, params...)
	}

	
		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256 balance)
		func (_IERC721 *IERC721Caller) BalanceOf(opts *bind.CallOpts , owner common.Address ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _IERC721.contract.Call(opts, out, "balanceOf" , owner)
			return *ret0, err
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256 balance)
		func (_IERC721 *IERC721Session) BalanceOf( owner common.Address ) ( *big.Int,  error) {
		  return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts , owner)
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256 balance)
		func (_IERC721 *IERC721CallerSession) BalanceOf( owner common.Address ) ( *big.Int,  error) {
		  return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts , owner)
		}
	
		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
		func (_IERC721 *IERC721Caller) GetApproved(opts *bind.CallOpts , tokenId *big.Int ) (common.Address, error) {
			var (
				ret0 = new(common.Address)
				
			)
			out := ret0
			err := _IERC721.contract.Call(opts, out, "getApproved" , tokenId)
			return *ret0, err
		}

		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
		func (_IERC721 *IERC721Session) GetApproved( tokenId *big.Int ) ( common.Address,  error) {
		  return _IERC721.Contract.GetApproved(&_IERC721.CallOpts , tokenId)
		}

		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
		func (_IERC721 *IERC721CallerSession) GetApproved( tokenId *big.Int ) ( common.Address,  error) {
		  return _IERC721.Contract.GetApproved(&_IERC721.CallOpts , tokenId)
		}
	
		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_IERC721 *IERC721Caller) IsApprovedForAll(opts *bind.CallOpts , owner common.Address , operator common.Address ) (bool, error) {
			var (
				ret0 = new(bool)
				
			)
			out := ret0
			err := _IERC721.contract.Call(opts, out, "isApprovedForAll" , owner, operator)
			return *ret0, err
		}

		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_IERC721 *IERC721Session) IsApprovedForAll( owner common.Address , operator common.Address ) ( bool,  error) {
		  return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts , owner, operator)
		}

		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_IERC721 *IERC721CallerSession) IsApprovedForAll( owner common.Address , operator common.Address ) ( bool,  error) {
		  return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts , owner, operator)
		}
	
		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
		func (_IERC721 *IERC721Caller) OwnerOf(opts *bind.CallOpts , tokenId *big.Int ) (common.Address, error) {
			var (
				ret0 = new(common.Address)
				
			)
			out := ret0
			err := _IERC721.contract.Call(opts, out, "ownerOf" , tokenId)
			return *ret0, err
		}

		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
		func (_IERC721 *IERC721Session) OwnerOf( tokenId *big.Int ) ( common.Address,  error) {
		  return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts , tokenId)
		}

		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
		func (_IERC721 *IERC721CallerSession) OwnerOf( tokenId *big.Int ) ( common.Address,  error) {
		  return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts , tokenId)
		}
	
		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_IERC721 *IERC721Caller) SupportsInterface(opts *bind.CallOpts , interfaceId [4]byte ) (bool, error) {
			var (
				ret0 = new(bool)
				
			)
			out := ret0
			err := _IERC721.contract.Call(opts, out, "supportsInterface" , interfaceId)
			return *ret0, err
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_IERC721 *IERC721Session) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts , interfaceId)
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_IERC721 *IERC721CallerSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts , interfaceId)
		}
	

	
		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_IERC721 *IERC721Transactor) Approve(opts *bind.TransactOpts , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _IERC721.contract.Transact(opts, "approve" , to, tokenId)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_IERC721 *IERC721Session) Approve( to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _IERC721.Contract.Approve(&_IERC721.TransactOpts , to, tokenId)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_IERC721 *IERC721TransactorSession) Approve( to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _IERC721.Contract.Approve(&_IERC721.TransactOpts , to, tokenId)
		}
	
		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_IERC721 *IERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _IERC721.contract.Transact(opts, "safeTransferFrom" , from, to, tokenId)
		}

		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_IERC721 *IERC721Session) SafeTransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts , from, to, tokenId)
		}

		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_IERC721 *IERC721TransactorSession) SafeTransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts , from, to, tokenId)
		}
	
		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
		func (_IERC721 *IERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int , data []byte ) (*types.Transaction, error) {
			return _IERC721.contract.Transact(opts, "safeTransferFrom0" , from, to, tokenId, data)
		}

		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
		func (_IERC721 *IERC721Session) SafeTransferFrom0( from common.Address , to common.Address , tokenId *big.Int , data []byte ) (*types.Transaction, error) {
		  return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts , from, to, tokenId, data)
		}

		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
		func (_IERC721 *IERC721TransactorSession) SafeTransferFrom0( from common.Address , to common.Address , tokenId *big.Int , data []byte ) (*types.Transaction, error) {
		  return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts , from, to, tokenId, data)
		}
	
		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
		func (_IERC721 *IERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts , operator common.Address , _approved bool ) (*types.Transaction, error) {
			return _IERC721.contract.Transact(opts, "setApprovalForAll" , operator, _approved)
		}

		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
		func (_IERC721 *IERC721Session) SetApprovalForAll( operator common.Address , _approved bool ) (*types.Transaction, error) {
		  return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts , operator, _approved)
		}

		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
		func (_IERC721 *IERC721TransactorSession) SetApprovalForAll( operator common.Address , _approved bool ) (*types.Transaction, error) {
		  return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts , operator, _approved)
		}
	
		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_IERC721 *IERC721Transactor) TransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _IERC721.contract.Transact(opts, "transferFrom" , from, to, tokenId)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_IERC721 *IERC721Session) TransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts , from, to, tokenId)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_IERC721 *IERC721TransactorSession) TransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts , from, to, tokenId)
		}
	

	

	

	
		// IERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721 contract.
		type IERC721ApprovalIterator struct {
			Event *IERC721Approval // Event containing the contract specifics and raw log

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
		func (it *IERC721ApprovalIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(IERC721Approval)
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
				it.Event = new(IERC721Approval)
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
		func (it *IERC721ApprovalIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *IERC721ApprovalIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// IERC721Approval represents a Approval event raised by the IERC721 contract.
		type IERC721Approval struct { 
			Owner common.Address; 
			Approved common.Address; 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
 		func (_IERC721 *IERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721ApprovalIterator, error) {
			
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

			logs, sub, err := _IERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &IERC721ApprovalIterator{contract: _IERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
 		}

		// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
		func (_IERC721 *IERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {
			
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

			logs, sub, err := _IERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(IERC721Approval)
						if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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
		func (_IERC721 *IERC721Filterer) ParseApproval(log types.Log) (*IERC721Approval, error) {
			event := new(IERC721Approval)
			if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// IERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721 contract.
		type IERC721ApprovalForAllIterator struct {
			Event *IERC721ApprovalForAll // Event containing the contract specifics and raw log

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
		func (it *IERC721ApprovalForAllIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(IERC721ApprovalForAll)
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
				it.Event = new(IERC721ApprovalForAll)
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
		func (it *IERC721ApprovalForAllIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *IERC721ApprovalForAllIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// IERC721ApprovalForAll represents a ApprovalForAll event raised by the IERC721 contract.
		type IERC721ApprovalForAll struct { 
			Owner common.Address; 
			Operator common.Address; 
			Approved bool; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
 		func (_IERC721 *IERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721ApprovalForAllIterator, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var operatorRule []interface{}
			for _, operatorItem := range operator {
				operatorRule = append(operatorRule, operatorItem)
			}
			

			logs, sub, err := _IERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
			if err != nil {
				return nil, err
			}
			return &IERC721ApprovalForAllIterator{contract: _IERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
 		}

		// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
		func (_IERC721 *IERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var operatorRule []interface{}
			for _, operatorItem := range operator {
				operatorRule = append(operatorRule, operatorItem)
			}
			

			logs, sub, err := _IERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(IERC721ApprovalForAll)
						if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
		func (_IERC721 *IERC721Filterer) ParseApprovalForAll(log types.Log) (*IERC721ApprovalForAll, error) {
			event := new(IERC721ApprovalForAll)
			if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// IERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721 contract.
		type IERC721TransferIterator struct {
			Event *IERC721Transfer // Event containing the contract specifics and raw log

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
		func (it *IERC721TransferIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(IERC721Transfer)
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
				it.Event = new(IERC721Transfer)
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
		func (it *IERC721TransferIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *IERC721TransferIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// IERC721Transfer represents a Transfer event raised by the IERC721 contract.
		type IERC721Transfer struct { 
			From common.Address; 
			To common.Address; 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
 		func (_IERC721 *IERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721TransferIterator, error) {
			
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

			logs, sub, err := _IERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &IERC721TransferIterator{contract: _IERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
 		}

		// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
		func (_IERC721 *IERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {
			
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

			logs, sub, err := _IERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(IERC721Transfer)
						if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
		func (_IERC721 *IERC721Filterer) ParseTransfer(log types.Log) (*IERC721Transfer, error) {
			event := new(IERC721Transfer)
			if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	

	// IERC721EnumerableABI is the input ABI used to generate the binding from.
	const IERC721EnumerableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

	
		// IERC721EnumerableFuncSigs maps the 4-byte function signature to its string representation.
		var IERC721EnumerableFuncSigs = map[string]string{
			"095ea7b3": "approve(address,uint256)",
			"70a08231": "balanceOf(address)",
			"081812fc": "getApproved(uint256)",
			"e985e9c5": "isApprovedForAll(address,address)",
			"6352211e": "ownerOf(uint256)",
			"42842e0e": "safeTransferFrom(address,address,uint256)",
			"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
			"a22cb465": "setApprovalForAll(address,bool)",
			"01ffc9a7": "supportsInterface(bytes4)",
			"4f6ccce7": "tokenByIndex(uint256)",
			"2f745c59": "tokenOfOwnerByIndex(address,uint256)",
			"18160ddd": "totalSupply()",
			"23b872dd": "transferFrom(address,address,uint256)",
			
		}
	

	

	// IERC721Enumerable is an auto generated Go binding around an Ethereum contract.
	type IERC721Enumerable struct {
	  IERC721EnumerableCaller     // Read-only binding to the contract
	  IERC721EnumerableTransactor // Write-only binding to the contract
	  IERC721EnumerableFilterer   // Log filterer for contract events
	}

	// IERC721EnumerableCaller is an auto generated read-only Go binding around an Ethereum contract.
	type IERC721EnumerableCaller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// IERC721EnumerableTransactor is an auto generated write-only Go binding around an Ethereum contract.
	type IERC721EnumerableTransactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// IERC721EnumerableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type IERC721EnumerableFilterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// IERC721EnumerableSession is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type IERC721EnumerableSession struct {
	  Contract     *IERC721Enumerable        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// IERC721EnumerableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type IERC721EnumerableCallerSession struct {
	  Contract *IERC721EnumerableCaller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// IERC721EnumerableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type IERC721EnumerableTransactorSession struct {
	  Contract     *IERC721EnumerableTransactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// IERC721EnumerableRaw is an auto generated low-level Go binding around an Ethereum contract.
	type IERC721EnumerableRaw struct {
	  Contract *IERC721Enumerable // Generic contract binding to access the raw methods on
	}

	// IERC721EnumerableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type IERC721EnumerableCallerRaw struct {
		Contract *IERC721EnumerableCaller // Generic read-only contract binding to access the raw methods on
	}

	// IERC721EnumerableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type IERC721EnumerableTransactorRaw struct {
		Contract *IERC721EnumerableTransactor // Generic write-only contract binding to access the raw methods on
	}

	// NewIERC721Enumerable creates a new instance of IERC721Enumerable, bound to a specific deployed contract.
	func NewIERC721Enumerable(address common.Address, backend bind.ContractBackend) (*IERC721Enumerable, error) {
	  contract, err := bindIERC721Enumerable(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &IERC721Enumerable{ IERC721EnumerableCaller: IERC721EnumerableCaller{contract: contract}, IERC721EnumerableTransactor: IERC721EnumerableTransactor{contract: contract}, IERC721EnumerableFilterer: IERC721EnumerableFilterer{contract: contract} }, nil
	}

	// NewIERC721EnumerableCaller creates a new read-only instance of IERC721Enumerable, bound to a specific deployed contract.
	func NewIERC721EnumerableCaller(address common.Address, caller bind.ContractCaller) (*IERC721EnumerableCaller, error) {
	  contract, err := bindIERC721Enumerable(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &IERC721EnumerableCaller{contract: contract}, nil
	}

	// NewIERC721EnumerableTransactor creates a new write-only instance of IERC721Enumerable, bound to a specific deployed contract.
	func NewIERC721EnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721EnumerableTransactor, error) {
	  contract, err := bindIERC721Enumerable(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &IERC721EnumerableTransactor{contract: contract}, nil
	}

	// NewIERC721EnumerableFilterer creates a new log filterer instance of IERC721Enumerable, bound to a specific deployed contract.
 	func NewIERC721EnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721EnumerableFilterer, error) {
 	  contract, err := bindIERC721Enumerable(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &IERC721EnumerableFilterer{contract: contract}, nil
 	}

	// bindIERC721Enumerable binds a generic wrapper to an already deployed contract.
	func bindIERC721Enumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(IERC721EnumerableABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_IERC721Enumerable *IERC721EnumerableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _IERC721Enumerable.Contract.IERC721EnumerableCaller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_IERC721Enumerable *IERC721EnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _IERC721Enumerable.Contract.IERC721EnumerableTransactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_IERC721Enumerable *IERC721EnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _IERC721Enumerable.Contract.IERC721EnumerableTransactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_IERC721Enumerable *IERC721EnumerableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _IERC721Enumerable.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_IERC721Enumerable *IERC721EnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _IERC721Enumerable.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_IERC721Enumerable *IERC721EnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _IERC721Enumerable.Contract.contract.Transact(opts, method, params...)
	}

	
		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256 balance)
		func (_IERC721Enumerable *IERC721EnumerableCaller) BalanceOf(opts *bind.CallOpts , owner common.Address ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _IERC721Enumerable.contract.Call(opts, out, "balanceOf" , owner)
			return *ret0, err
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256 balance)
		func (_IERC721Enumerable *IERC721EnumerableSession) BalanceOf( owner common.Address ) ( *big.Int,  error) {
		  return _IERC721Enumerable.Contract.BalanceOf(&_IERC721Enumerable.CallOpts , owner)
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256 balance)
		func (_IERC721Enumerable *IERC721EnumerableCallerSession) BalanceOf( owner common.Address ) ( *big.Int,  error) {
		  return _IERC721Enumerable.Contract.BalanceOf(&_IERC721Enumerable.CallOpts , owner)
		}
	
		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
		func (_IERC721Enumerable *IERC721EnumerableCaller) GetApproved(opts *bind.CallOpts , tokenId *big.Int ) (common.Address, error) {
			var (
				ret0 = new(common.Address)
				
			)
			out := ret0
			err := _IERC721Enumerable.contract.Call(opts, out, "getApproved" , tokenId)
			return *ret0, err
		}

		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
		func (_IERC721Enumerable *IERC721EnumerableSession) GetApproved( tokenId *big.Int ) ( common.Address,  error) {
		  return _IERC721Enumerable.Contract.GetApproved(&_IERC721Enumerable.CallOpts , tokenId)
		}

		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
		func (_IERC721Enumerable *IERC721EnumerableCallerSession) GetApproved( tokenId *big.Int ) ( common.Address,  error) {
		  return _IERC721Enumerable.Contract.GetApproved(&_IERC721Enumerable.CallOpts , tokenId)
		}
	
		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_IERC721Enumerable *IERC721EnumerableCaller) IsApprovedForAll(opts *bind.CallOpts , owner common.Address , operator common.Address ) (bool, error) {
			var (
				ret0 = new(bool)
				
			)
			out := ret0
			err := _IERC721Enumerable.contract.Call(opts, out, "isApprovedForAll" , owner, operator)
			return *ret0, err
		}

		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_IERC721Enumerable *IERC721EnumerableSession) IsApprovedForAll( owner common.Address , operator common.Address ) ( bool,  error) {
		  return _IERC721Enumerable.Contract.IsApprovedForAll(&_IERC721Enumerable.CallOpts , owner, operator)
		}

		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_IERC721Enumerable *IERC721EnumerableCallerSession) IsApprovedForAll( owner common.Address , operator common.Address ) ( bool,  error) {
		  return _IERC721Enumerable.Contract.IsApprovedForAll(&_IERC721Enumerable.CallOpts , owner, operator)
		}
	
		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
		func (_IERC721Enumerable *IERC721EnumerableCaller) OwnerOf(opts *bind.CallOpts , tokenId *big.Int ) (common.Address, error) {
			var (
				ret0 = new(common.Address)
				
			)
			out := ret0
			err := _IERC721Enumerable.contract.Call(opts, out, "ownerOf" , tokenId)
			return *ret0, err
		}

		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
		func (_IERC721Enumerable *IERC721EnumerableSession) OwnerOf( tokenId *big.Int ) ( common.Address,  error) {
		  return _IERC721Enumerable.Contract.OwnerOf(&_IERC721Enumerable.CallOpts , tokenId)
		}

		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
		func (_IERC721Enumerable *IERC721EnumerableCallerSession) OwnerOf( tokenId *big.Int ) ( common.Address,  error) {
		  return _IERC721Enumerable.Contract.OwnerOf(&_IERC721Enumerable.CallOpts , tokenId)
		}
	
		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_IERC721Enumerable *IERC721EnumerableCaller) SupportsInterface(opts *bind.CallOpts , interfaceId [4]byte ) (bool, error) {
			var (
				ret0 = new(bool)
				
			)
			out := ret0
			err := _IERC721Enumerable.contract.Call(opts, out, "supportsInterface" , interfaceId)
			return *ret0, err
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_IERC721Enumerable *IERC721EnumerableSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _IERC721Enumerable.Contract.SupportsInterface(&_IERC721Enumerable.CallOpts , interfaceId)
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_IERC721Enumerable *IERC721EnumerableCallerSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _IERC721Enumerable.Contract.SupportsInterface(&_IERC721Enumerable.CallOpts , interfaceId)
		}
	
		// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
		//
		// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
		func (_IERC721Enumerable *IERC721EnumerableCaller) TokenByIndex(opts *bind.CallOpts , index *big.Int ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _IERC721Enumerable.contract.Call(opts, out, "tokenByIndex" , index)
			return *ret0, err
		}

		// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
		//
		// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
		func (_IERC721Enumerable *IERC721EnumerableSession) TokenByIndex( index *big.Int ) ( *big.Int,  error) {
		  return _IERC721Enumerable.Contract.TokenByIndex(&_IERC721Enumerable.CallOpts , index)
		}

		// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
		//
		// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
		func (_IERC721Enumerable *IERC721EnumerableCallerSession) TokenByIndex( index *big.Int ) ( *big.Int,  error) {
		  return _IERC721Enumerable.Contract.TokenByIndex(&_IERC721Enumerable.CallOpts , index)
		}
	
		// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
		//
		// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
		func (_IERC721Enumerable *IERC721EnumerableCaller) TokenOfOwnerByIndex(opts *bind.CallOpts , owner common.Address , index *big.Int ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _IERC721Enumerable.contract.Call(opts, out, "tokenOfOwnerByIndex" , owner, index)
			return *ret0, err
		}

		// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
		//
		// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
		func (_IERC721Enumerable *IERC721EnumerableSession) TokenOfOwnerByIndex( owner common.Address , index *big.Int ) ( *big.Int,  error) {
		  return _IERC721Enumerable.Contract.TokenOfOwnerByIndex(&_IERC721Enumerable.CallOpts , owner, index)
		}

		// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
		//
		// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256 tokenId)
		func (_IERC721Enumerable *IERC721EnumerableCallerSession) TokenOfOwnerByIndex( owner common.Address , index *big.Int ) ( *big.Int,  error) {
		  return _IERC721Enumerable.Contract.TokenOfOwnerByIndex(&_IERC721Enumerable.CallOpts , owner, index)
		}
	
		// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
		//
		// Solidity: function totalSupply() view returns(uint256)
		func (_IERC721Enumerable *IERC721EnumerableCaller) TotalSupply(opts *bind.CallOpts ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _IERC721Enumerable.contract.Call(opts, out, "totalSupply" )
			return *ret0, err
		}

		// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
		//
		// Solidity: function totalSupply() view returns(uint256)
		func (_IERC721Enumerable *IERC721EnumerableSession) TotalSupply() ( *big.Int,  error) {
		  return _IERC721Enumerable.Contract.TotalSupply(&_IERC721Enumerable.CallOpts )
		}

		// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
		//
		// Solidity: function totalSupply() view returns(uint256)
		func (_IERC721Enumerable *IERC721EnumerableCallerSession) TotalSupply() ( *big.Int,  error) {
		  return _IERC721Enumerable.Contract.TotalSupply(&_IERC721Enumerable.CallOpts )
		}
	

	
		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_IERC721Enumerable *IERC721EnumerableTransactor) Approve(opts *bind.TransactOpts , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _IERC721Enumerable.contract.Transact(opts, "approve" , to, tokenId)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_IERC721Enumerable *IERC721EnumerableSession) Approve( to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _IERC721Enumerable.Contract.Approve(&_IERC721Enumerable.TransactOpts , to, tokenId)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_IERC721Enumerable *IERC721EnumerableTransactorSession) Approve( to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _IERC721Enumerable.Contract.Approve(&_IERC721Enumerable.TransactOpts , to, tokenId)
		}
	
		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_IERC721Enumerable *IERC721EnumerableTransactor) SafeTransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _IERC721Enumerable.contract.Transact(opts, "safeTransferFrom" , from, to, tokenId)
		}

		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_IERC721Enumerable *IERC721EnumerableSession) SafeTransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _IERC721Enumerable.Contract.SafeTransferFrom(&_IERC721Enumerable.TransactOpts , from, to, tokenId)
		}

		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_IERC721Enumerable *IERC721EnumerableTransactorSession) SafeTransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _IERC721Enumerable.Contract.SafeTransferFrom(&_IERC721Enumerable.TransactOpts , from, to, tokenId)
		}
	
		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
		func (_IERC721Enumerable *IERC721EnumerableTransactor) SafeTransferFrom0(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int , data []byte ) (*types.Transaction, error) {
			return _IERC721Enumerable.contract.Transact(opts, "safeTransferFrom0" , from, to, tokenId, data)
		}

		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
		func (_IERC721Enumerable *IERC721EnumerableSession) SafeTransferFrom0( from common.Address , to common.Address , tokenId *big.Int , data []byte ) (*types.Transaction, error) {
		  return _IERC721Enumerable.Contract.SafeTransferFrom0(&_IERC721Enumerable.TransactOpts , from, to, tokenId, data)
		}

		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
		func (_IERC721Enumerable *IERC721EnumerableTransactorSession) SafeTransferFrom0( from common.Address , to common.Address , tokenId *big.Int , data []byte ) (*types.Transaction, error) {
		  return _IERC721Enumerable.Contract.SafeTransferFrom0(&_IERC721Enumerable.TransactOpts , from, to, tokenId, data)
		}
	
		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
		func (_IERC721Enumerable *IERC721EnumerableTransactor) SetApprovalForAll(opts *bind.TransactOpts , operator common.Address , _approved bool ) (*types.Transaction, error) {
			return _IERC721Enumerable.contract.Transact(opts, "setApprovalForAll" , operator, _approved)
		}

		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
		func (_IERC721Enumerable *IERC721EnumerableSession) SetApprovalForAll( operator common.Address , _approved bool ) (*types.Transaction, error) {
		  return _IERC721Enumerable.Contract.SetApprovalForAll(&_IERC721Enumerable.TransactOpts , operator, _approved)
		}

		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
		func (_IERC721Enumerable *IERC721EnumerableTransactorSession) SetApprovalForAll( operator common.Address , _approved bool ) (*types.Transaction, error) {
		  return _IERC721Enumerable.Contract.SetApprovalForAll(&_IERC721Enumerable.TransactOpts , operator, _approved)
		}
	
		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_IERC721Enumerable *IERC721EnumerableTransactor) TransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _IERC721Enumerable.contract.Transact(opts, "transferFrom" , from, to, tokenId)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_IERC721Enumerable *IERC721EnumerableSession) TransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _IERC721Enumerable.Contract.TransferFrom(&_IERC721Enumerable.TransactOpts , from, to, tokenId)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_IERC721Enumerable *IERC721EnumerableTransactorSession) TransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _IERC721Enumerable.Contract.TransferFrom(&_IERC721Enumerable.TransactOpts , from, to, tokenId)
		}
	

	

	

	
		// IERC721EnumerableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721Enumerable contract.
		type IERC721EnumerableApprovalIterator struct {
			Event *IERC721EnumerableApproval // Event containing the contract specifics and raw log

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
		func (it *IERC721EnumerableApprovalIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(IERC721EnumerableApproval)
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
				it.Event = new(IERC721EnumerableApproval)
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
		func (it *IERC721EnumerableApprovalIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *IERC721EnumerableApprovalIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// IERC721EnumerableApproval represents a Approval event raised by the IERC721Enumerable contract.
		type IERC721EnumerableApproval struct { 
			Owner common.Address; 
			Approved common.Address; 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
 		func (_IERC721Enumerable *IERC721EnumerableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721EnumerableApprovalIterator, error) {
			
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

			logs, sub, err := _IERC721Enumerable.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &IERC721EnumerableApprovalIterator{contract: _IERC721Enumerable.contract, event: "Approval", logs: logs, sub: sub}, nil
 		}

		// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
		func (_IERC721Enumerable *IERC721EnumerableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721EnumerableApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {
			
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

			logs, sub, err := _IERC721Enumerable.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(IERC721EnumerableApproval)
						if err := _IERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
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
		func (_IERC721Enumerable *IERC721EnumerableFilterer) ParseApproval(log types.Log) (*IERC721EnumerableApproval, error) {
			event := new(IERC721EnumerableApproval)
			if err := _IERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// IERC721EnumerableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721Enumerable contract.
		type IERC721EnumerableApprovalForAllIterator struct {
			Event *IERC721EnumerableApprovalForAll // Event containing the contract specifics and raw log

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
		func (it *IERC721EnumerableApprovalForAllIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(IERC721EnumerableApprovalForAll)
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
				it.Event = new(IERC721EnumerableApprovalForAll)
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
		func (it *IERC721EnumerableApprovalForAllIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *IERC721EnumerableApprovalForAllIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// IERC721EnumerableApprovalForAll represents a ApprovalForAll event raised by the IERC721Enumerable contract.
		type IERC721EnumerableApprovalForAll struct { 
			Owner common.Address; 
			Operator common.Address; 
			Approved bool; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
 		func (_IERC721Enumerable *IERC721EnumerableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721EnumerableApprovalForAllIterator, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var operatorRule []interface{}
			for _, operatorItem := range operator {
				operatorRule = append(operatorRule, operatorItem)
			}
			

			logs, sub, err := _IERC721Enumerable.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
			if err != nil {
				return nil, err
			}
			return &IERC721EnumerableApprovalForAllIterator{contract: _IERC721Enumerable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
 		}

		// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
		func (_IERC721Enumerable *IERC721EnumerableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721EnumerableApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var operatorRule []interface{}
			for _, operatorItem := range operator {
				operatorRule = append(operatorRule, operatorItem)
			}
			

			logs, sub, err := _IERC721Enumerable.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(IERC721EnumerableApprovalForAll)
						if err := _IERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
		func (_IERC721Enumerable *IERC721EnumerableFilterer) ParseApprovalForAll(log types.Log) (*IERC721EnumerableApprovalForAll, error) {
			event := new(IERC721EnumerableApprovalForAll)
			if err := _IERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// IERC721EnumerableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721Enumerable contract.
		type IERC721EnumerableTransferIterator struct {
			Event *IERC721EnumerableTransfer // Event containing the contract specifics and raw log

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
		func (it *IERC721EnumerableTransferIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(IERC721EnumerableTransfer)
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
				it.Event = new(IERC721EnumerableTransfer)
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
		func (it *IERC721EnumerableTransferIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *IERC721EnumerableTransferIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// IERC721EnumerableTransfer represents a Transfer event raised by the IERC721Enumerable contract.
		type IERC721EnumerableTransfer struct { 
			From common.Address; 
			To common.Address; 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
 		func (_IERC721Enumerable *IERC721EnumerableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721EnumerableTransferIterator, error) {
			
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

			logs, sub, err := _IERC721Enumerable.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &IERC721EnumerableTransferIterator{contract: _IERC721Enumerable.contract, event: "Transfer", logs: logs, sub: sub}, nil
 		}

		// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
		func (_IERC721Enumerable *IERC721EnumerableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721EnumerableTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {
			
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

			logs, sub, err := _IERC721Enumerable.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(IERC721EnumerableTransfer)
						if err := _IERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
		func (_IERC721Enumerable *IERC721EnumerableFilterer) ParseTransfer(log types.Log) (*IERC721EnumerableTransfer, error) {
			event := new(IERC721EnumerableTransfer)
			if err := _IERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	

	// IERC721MetadataABI is the input ABI used to generate the binding from.
	const IERC721MetadataABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

	
		// IERC721MetadataFuncSigs maps the 4-byte function signature to its string representation.
		var IERC721MetadataFuncSigs = map[string]string{
			"095ea7b3": "approve(address,uint256)",
			"70a08231": "balanceOf(address)",
			"081812fc": "getApproved(uint256)",
			"e985e9c5": "isApprovedForAll(address,address)",
			"06fdde03": "name()",
			"6352211e": "ownerOf(uint256)",
			"42842e0e": "safeTransferFrom(address,address,uint256)",
			"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
			"a22cb465": "setApprovalForAll(address,bool)",
			"01ffc9a7": "supportsInterface(bytes4)",
			"95d89b41": "symbol()",
			"c87b56dd": "tokenURI(uint256)",
			"23b872dd": "transferFrom(address,address,uint256)",
			
		}
	

	

	// IERC721Metadata is an auto generated Go binding around an Ethereum contract.
	type IERC721Metadata struct {
	  IERC721MetadataCaller     // Read-only binding to the contract
	  IERC721MetadataTransactor // Write-only binding to the contract
	  IERC721MetadataFilterer   // Log filterer for contract events
	}

	// IERC721MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
	type IERC721MetadataCaller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// IERC721MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
	type IERC721MetadataTransactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// IERC721MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type IERC721MetadataFilterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// IERC721MetadataSession is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type IERC721MetadataSession struct {
	  Contract     *IERC721Metadata        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// IERC721MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type IERC721MetadataCallerSession struct {
	  Contract *IERC721MetadataCaller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// IERC721MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type IERC721MetadataTransactorSession struct {
	  Contract     *IERC721MetadataTransactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// IERC721MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
	type IERC721MetadataRaw struct {
	  Contract *IERC721Metadata // Generic contract binding to access the raw methods on
	}

	// IERC721MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type IERC721MetadataCallerRaw struct {
		Contract *IERC721MetadataCaller // Generic read-only contract binding to access the raw methods on
	}

	// IERC721MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type IERC721MetadataTransactorRaw struct {
		Contract *IERC721MetadataTransactor // Generic write-only contract binding to access the raw methods on
	}

	// NewIERC721Metadata creates a new instance of IERC721Metadata, bound to a specific deployed contract.
	func NewIERC721Metadata(address common.Address, backend bind.ContractBackend) (*IERC721Metadata, error) {
	  contract, err := bindIERC721Metadata(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &IERC721Metadata{ IERC721MetadataCaller: IERC721MetadataCaller{contract: contract}, IERC721MetadataTransactor: IERC721MetadataTransactor{contract: contract}, IERC721MetadataFilterer: IERC721MetadataFilterer{contract: contract} }, nil
	}

	// NewIERC721MetadataCaller creates a new read-only instance of IERC721Metadata, bound to a specific deployed contract.
	func NewIERC721MetadataCaller(address common.Address, caller bind.ContractCaller) (*IERC721MetadataCaller, error) {
	  contract, err := bindIERC721Metadata(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &IERC721MetadataCaller{contract: contract}, nil
	}

	// NewIERC721MetadataTransactor creates a new write-only instance of IERC721Metadata, bound to a specific deployed contract.
	func NewIERC721MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721MetadataTransactor, error) {
	  contract, err := bindIERC721Metadata(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &IERC721MetadataTransactor{contract: contract}, nil
	}

	// NewIERC721MetadataFilterer creates a new log filterer instance of IERC721Metadata, bound to a specific deployed contract.
 	func NewIERC721MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721MetadataFilterer, error) {
 	  contract, err := bindIERC721Metadata(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &IERC721MetadataFilterer{contract: contract}, nil
 	}

	// bindIERC721Metadata binds a generic wrapper to an already deployed contract.
	func bindIERC721Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(IERC721MetadataABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_IERC721Metadata *IERC721MetadataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _IERC721Metadata.Contract.IERC721MetadataCaller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_IERC721Metadata *IERC721MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _IERC721Metadata.Contract.IERC721MetadataTransactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_IERC721Metadata *IERC721MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _IERC721Metadata.Contract.IERC721MetadataTransactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_IERC721Metadata *IERC721MetadataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _IERC721Metadata.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_IERC721Metadata *IERC721MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _IERC721Metadata.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_IERC721Metadata *IERC721MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _IERC721Metadata.Contract.contract.Transact(opts, method, params...)
	}

	
		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256 balance)
		func (_IERC721Metadata *IERC721MetadataCaller) BalanceOf(opts *bind.CallOpts , owner common.Address ) (*big.Int, error) {
			var (
				ret0 = new(*big.Int)
				
			)
			out := ret0
			err := _IERC721Metadata.contract.Call(opts, out, "balanceOf" , owner)
			return *ret0, err
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256 balance)
		func (_IERC721Metadata *IERC721MetadataSession) BalanceOf( owner common.Address ) ( *big.Int,  error) {
		  return _IERC721Metadata.Contract.BalanceOf(&_IERC721Metadata.CallOpts , owner)
		}

		// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
		//
		// Solidity: function balanceOf(address owner) view returns(uint256 balance)
		func (_IERC721Metadata *IERC721MetadataCallerSession) BalanceOf( owner common.Address ) ( *big.Int,  error) {
		  return _IERC721Metadata.Contract.BalanceOf(&_IERC721Metadata.CallOpts , owner)
		}
	
		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
		func (_IERC721Metadata *IERC721MetadataCaller) GetApproved(opts *bind.CallOpts , tokenId *big.Int ) (common.Address, error) {
			var (
				ret0 = new(common.Address)
				
			)
			out := ret0
			err := _IERC721Metadata.contract.Call(opts, out, "getApproved" , tokenId)
			return *ret0, err
		}

		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
		func (_IERC721Metadata *IERC721MetadataSession) GetApproved( tokenId *big.Int ) ( common.Address,  error) {
		  return _IERC721Metadata.Contract.GetApproved(&_IERC721Metadata.CallOpts , tokenId)
		}

		// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
		//
		// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
		func (_IERC721Metadata *IERC721MetadataCallerSession) GetApproved( tokenId *big.Int ) ( common.Address,  error) {
		  return _IERC721Metadata.Contract.GetApproved(&_IERC721Metadata.CallOpts , tokenId)
		}
	
		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_IERC721Metadata *IERC721MetadataCaller) IsApprovedForAll(opts *bind.CallOpts , owner common.Address , operator common.Address ) (bool, error) {
			var (
				ret0 = new(bool)
				
			)
			out := ret0
			err := _IERC721Metadata.contract.Call(opts, out, "isApprovedForAll" , owner, operator)
			return *ret0, err
		}

		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_IERC721Metadata *IERC721MetadataSession) IsApprovedForAll( owner common.Address , operator common.Address ) ( bool,  error) {
		  return _IERC721Metadata.Contract.IsApprovedForAll(&_IERC721Metadata.CallOpts , owner, operator)
		}

		// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
		//
		// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
		func (_IERC721Metadata *IERC721MetadataCallerSession) IsApprovedForAll( owner common.Address , operator common.Address ) ( bool,  error) {
		  return _IERC721Metadata.Contract.IsApprovedForAll(&_IERC721Metadata.CallOpts , owner, operator)
		}
	
		// Name is a free data retrieval call binding the contract method 0x06fdde03.
		//
		// Solidity: function name() view returns(string)
		func (_IERC721Metadata *IERC721MetadataCaller) Name(opts *bind.CallOpts ) (string, error) {
			var (
				ret0 = new(string)
				
			)
			out := ret0
			err := _IERC721Metadata.contract.Call(opts, out, "name" )
			return *ret0, err
		}

		// Name is a free data retrieval call binding the contract method 0x06fdde03.
		//
		// Solidity: function name() view returns(string)
		func (_IERC721Metadata *IERC721MetadataSession) Name() ( string,  error) {
		  return _IERC721Metadata.Contract.Name(&_IERC721Metadata.CallOpts )
		}

		// Name is a free data retrieval call binding the contract method 0x06fdde03.
		//
		// Solidity: function name() view returns(string)
		func (_IERC721Metadata *IERC721MetadataCallerSession) Name() ( string,  error) {
		  return _IERC721Metadata.Contract.Name(&_IERC721Metadata.CallOpts )
		}
	
		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
		func (_IERC721Metadata *IERC721MetadataCaller) OwnerOf(opts *bind.CallOpts , tokenId *big.Int ) (common.Address, error) {
			var (
				ret0 = new(common.Address)
				
			)
			out := ret0
			err := _IERC721Metadata.contract.Call(opts, out, "ownerOf" , tokenId)
			return *ret0, err
		}

		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
		func (_IERC721Metadata *IERC721MetadataSession) OwnerOf( tokenId *big.Int ) ( common.Address,  error) {
		  return _IERC721Metadata.Contract.OwnerOf(&_IERC721Metadata.CallOpts , tokenId)
		}

		// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
		//
		// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
		func (_IERC721Metadata *IERC721MetadataCallerSession) OwnerOf( tokenId *big.Int ) ( common.Address,  error) {
		  return _IERC721Metadata.Contract.OwnerOf(&_IERC721Metadata.CallOpts , tokenId)
		}
	
		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_IERC721Metadata *IERC721MetadataCaller) SupportsInterface(opts *bind.CallOpts , interfaceId [4]byte ) (bool, error) {
			var (
				ret0 = new(bool)
				
			)
			out := ret0
			err := _IERC721Metadata.contract.Call(opts, out, "supportsInterface" , interfaceId)
			return *ret0, err
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_IERC721Metadata *IERC721MetadataSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _IERC721Metadata.Contract.SupportsInterface(&_IERC721Metadata.CallOpts , interfaceId)
		}

		// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
		//
		// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
		func (_IERC721Metadata *IERC721MetadataCallerSession) SupportsInterface( interfaceId [4]byte ) ( bool,  error) {
		  return _IERC721Metadata.Contract.SupportsInterface(&_IERC721Metadata.CallOpts , interfaceId)
		}
	
		// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
		//
		// Solidity: function symbol() view returns(string)
		func (_IERC721Metadata *IERC721MetadataCaller) Symbol(opts *bind.CallOpts ) (string, error) {
			var (
				ret0 = new(string)
				
			)
			out := ret0
			err := _IERC721Metadata.contract.Call(opts, out, "symbol" )
			return *ret0, err
		}

		// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
		//
		// Solidity: function symbol() view returns(string)
		func (_IERC721Metadata *IERC721MetadataSession) Symbol() ( string,  error) {
		  return _IERC721Metadata.Contract.Symbol(&_IERC721Metadata.CallOpts )
		}

		// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
		//
		// Solidity: function symbol() view returns(string)
		func (_IERC721Metadata *IERC721MetadataCallerSession) Symbol() ( string,  error) {
		  return _IERC721Metadata.Contract.Symbol(&_IERC721Metadata.CallOpts )
		}
	
		// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
		//
		// Solidity: function tokenURI(uint256 tokenId) view returns(string)
		func (_IERC721Metadata *IERC721MetadataCaller) TokenURI(opts *bind.CallOpts , tokenId *big.Int ) (string, error) {
			var (
				ret0 = new(string)
				
			)
			out := ret0
			err := _IERC721Metadata.contract.Call(opts, out, "tokenURI" , tokenId)
			return *ret0, err
		}

		// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
		//
		// Solidity: function tokenURI(uint256 tokenId) view returns(string)
		func (_IERC721Metadata *IERC721MetadataSession) TokenURI( tokenId *big.Int ) ( string,  error) {
		  return _IERC721Metadata.Contract.TokenURI(&_IERC721Metadata.CallOpts , tokenId)
		}

		// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
		//
		// Solidity: function tokenURI(uint256 tokenId) view returns(string)
		func (_IERC721Metadata *IERC721MetadataCallerSession) TokenURI( tokenId *big.Int ) ( string,  error) {
		  return _IERC721Metadata.Contract.TokenURI(&_IERC721Metadata.CallOpts , tokenId)
		}
	

	
		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_IERC721Metadata *IERC721MetadataTransactor) Approve(opts *bind.TransactOpts , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _IERC721Metadata.contract.Transact(opts, "approve" , to, tokenId)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_IERC721Metadata *IERC721MetadataSession) Approve( to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _IERC721Metadata.Contract.Approve(&_IERC721Metadata.TransactOpts , to, tokenId)
		}

		// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
		//
		// Solidity: function approve(address to, uint256 tokenId) returns()
		func (_IERC721Metadata *IERC721MetadataTransactorSession) Approve( to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _IERC721Metadata.Contract.Approve(&_IERC721Metadata.TransactOpts , to, tokenId)
		}
	
		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_IERC721Metadata *IERC721MetadataTransactor) SafeTransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _IERC721Metadata.contract.Transact(opts, "safeTransferFrom" , from, to, tokenId)
		}

		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_IERC721Metadata *IERC721MetadataSession) SafeTransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _IERC721Metadata.Contract.SafeTransferFrom(&_IERC721Metadata.TransactOpts , from, to, tokenId)
		}

		// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
		func (_IERC721Metadata *IERC721MetadataTransactorSession) SafeTransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _IERC721Metadata.Contract.SafeTransferFrom(&_IERC721Metadata.TransactOpts , from, to, tokenId)
		}
	
		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
		func (_IERC721Metadata *IERC721MetadataTransactor) SafeTransferFrom0(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int , data []byte ) (*types.Transaction, error) {
			return _IERC721Metadata.contract.Transact(opts, "safeTransferFrom0" , from, to, tokenId, data)
		}

		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
		func (_IERC721Metadata *IERC721MetadataSession) SafeTransferFrom0( from common.Address , to common.Address , tokenId *big.Int , data []byte ) (*types.Transaction, error) {
		  return _IERC721Metadata.Contract.SafeTransferFrom0(&_IERC721Metadata.TransactOpts , from, to, tokenId, data)
		}

		// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
		//
		// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
		func (_IERC721Metadata *IERC721MetadataTransactorSession) SafeTransferFrom0( from common.Address , to common.Address , tokenId *big.Int , data []byte ) (*types.Transaction, error) {
		  return _IERC721Metadata.Contract.SafeTransferFrom0(&_IERC721Metadata.TransactOpts , from, to, tokenId, data)
		}
	
		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
		func (_IERC721Metadata *IERC721MetadataTransactor) SetApprovalForAll(opts *bind.TransactOpts , operator common.Address , _approved bool ) (*types.Transaction, error) {
			return _IERC721Metadata.contract.Transact(opts, "setApprovalForAll" , operator, _approved)
		}

		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
		func (_IERC721Metadata *IERC721MetadataSession) SetApprovalForAll( operator common.Address , _approved bool ) (*types.Transaction, error) {
		  return _IERC721Metadata.Contract.SetApprovalForAll(&_IERC721Metadata.TransactOpts , operator, _approved)
		}

		// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
		//
		// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
		func (_IERC721Metadata *IERC721MetadataTransactorSession) SetApprovalForAll( operator common.Address , _approved bool ) (*types.Transaction, error) {
		  return _IERC721Metadata.Contract.SetApprovalForAll(&_IERC721Metadata.TransactOpts , operator, _approved)
		}
	
		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_IERC721Metadata *IERC721MetadataTransactor) TransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
			return _IERC721Metadata.contract.Transact(opts, "transferFrom" , from, to, tokenId)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_IERC721Metadata *IERC721MetadataSession) TransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _IERC721Metadata.Contract.TransferFrom(&_IERC721Metadata.TransactOpts , from, to, tokenId)
		}

		// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
		//
		// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
		func (_IERC721Metadata *IERC721MetadataTransactorSession) TransferFrom( from common.Address , to common.Address , tokenId *big.Int ) (*types.Transaction, error) {
		  return _IERC721Metadata.Contract.TransferFrom(&_IERC721Metadata.TransactOpts , from, to, tokenId)
		}
	

	

	

	
		// IERC721MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721Metadata contract.
		type IERC721MetadataApprovalIterator struct {
			Event *IERC721MetadataApproval // Event containing the contract specifics and raw log

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
		func (it *IERC721MetadataApprovalIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(IERC721MetadataApproval)
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
				it.Event = new(IERC721MetadataApproval)
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
		func (it *IERC721MetadataApprovalIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *IERC721MetadataApprovalIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// IERC721MetadataApproval represents a Approval event raised by the IERC721Metadata contract.
		type IERC721MetadataApproval struct { 
			Owner common.Address; 
			Approved common.Address; 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
 		func (_IERC721Metadata *IERC721MetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721MetadataApprovalIterator, error) {
			
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

			logs, sub, err := _IERC721Metadata.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &IERC721MetadataApprovalIterator{contract: _IERC721Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
 		}

		// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
		//
		// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
		func (_IERC721Metadata *IERC721MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721MetadataApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {
			
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

			logs, sub, err := _IERC721Metadata.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(IERC721MetadataApproval)
						if err := _IERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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
		func (_IERC721Metadata *IERC721MetadataFilterer) ParseApproval(log types.Log) (*IERC721MetadataApproval, error) {
			event := new(IERC721MetadataApproval)
			if err := _IERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// IERC721MetadataApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721Metadata contract.
		type IERC721MetadataApprovalForAllIterator struct {
			Event *IERC721MetadataApprovalForAll // Event containing the contract specifics and raw log

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
		func (it *IERC721MetadataApprovalForAllIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(IERC721MetadataApprovalForAll)
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
				it.Event = new(IERC721MetadataApprovalForAll)
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
		func (it *IERC721MetadataApprovalForAllIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *IERC721MetadataApprovalForAllIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// IERC721MetadataApprovalForAll represents a ApprovalForAll event raised by the IERC721Metadata contract.
		type IERC721MetadataApprovalForAll struct { 
			Owner common.Address; 
			Operator common.Address; 
			Approved bool; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
 		func (_IERC721Metadata *IERC721MetadataFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721MetadataApprovalForAllIterator, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var operatorRule []interface{}
			for _, operatorItem := range operator {
				operatorRule = append(operatorRule, operatorItem)
			}
			

			logs, sub, err := _IERC721Metadata.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
			if err != nil {
				return nil, err
			}
			return &IERC721MetadataApprovalForAllIterator{contract: _IERC721Metadata.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
 		}

		// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
		//
		// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
		func (_IERC721Metadata *IERC721MetadataFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721MetadataApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {
			
			var ownerRule []interface{}
			for _, ownerItem := range owner {
				ownerRule = append(ownerRule, ownerItem)
			}
			var operatorRule []interface{}
			for _, operatorItem := range operator {
				operatorRule = append(operatorRule, operatorItem)
			}
			

			logs, sub, err := _IERC721Metadata.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(IERC721MetadataApprovalForAll)
						if err := _IERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
		func (_IERC721Metadata *IERC721MetadataFilterer) ParseApprovalForAll(log types.Log) (*IERC721MetadataApprovalForAll, error) {
			event := new(IERC721MetadataApprovalForAll)
			if err := _IERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	
		// IERC721MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721Metadata contract.
		type IERC721MetadataTransferIterator struct {
			Event *IERC721MetadataTransfer // Event containing the contract specifics and raw log

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
		func (it *IERC721MetadataTransferIterator) Next() bool {
			// If the iterator failed, stop iterating
			if (it.fail != nil) {
				return false
			}
			// If the iterator completed, deliver directly whatever's available
			if (it.done) {
				select {
				case log := <-it.logs:
					it.Event = new(IERC721MetadataTransfer)
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
				it.Event = new(IERC721MetadataTransfer)
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
		func (it *IERC721MetadataTransferIterator) Error() error {
			return it.fail
		}
		// Close terminates the iteration process, releasing any pending underlying
		// resources.
		func (it *IERC721MetadataTransferIterator) Close() error {
			it.sub.Unsubscribe()
			return nil
		}

		// IERC721MetadataTransfer represents a Transfer event raised by the IERC721Metadata contract.
		type IERC721MetadataTransfer struct { 
			From common.Address; 
			To common.Address; 
			TokenId *big.Int; 
			Raw types.Log // Blockchain specific contextual infos
		}

		// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
 		func (_IERC721Metadata *IERC721MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721MetadataTransferIterator, error) {
			
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

			logs, sub, err := _IERC721Metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return &IERC721MetadataTransferIterator{contract: _IERC721Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
 		}

		// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
		//
		// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
		func (_IERC721Metadata *IERC721MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721MetadataTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {
			
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

			logs, sub, err := _IERC721Metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
			if err != nil {
				return nil, err
			}
			return event.NewSubscription(func(quit <-chan struct{}) error {
				defer sub.Unsubscribe()
				for {
					select {
					case log := <-logs:
						// New log arrived, parse the event and forward to the user
						event := new(IERC721MetadataTransfer)
						if err := _IERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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
		func (_IERC721Metadata *IERC721MetadataFilterer) ParseTransfer(log types.Log) (*IERC721MetadataTransfer, error) {
			event := new(IERC721MetadataTransfer)
			if err := _IERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
				return nil, err
			}
			return event, nil
		}

 	

	// IERC721ReceiverABI is the input ABI used to generate the binding from.
	const IERC721ReceiverABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

	
		// IERC721ReceiverFuncSigs maps the 4-byte function signature to its string representation.
		var IERC721ReceiverFuncSigs = map[string]string{
			"150b7a02": "onERC721Received(address,address,uint256,bytes)",
			
		}
	

	

	// IERC721Receiver is an auto generated Go binding around an Ethereum contract.
	type IERC721Receiver struct {
	  IERC721ReceiverCaller     // Read-only binding to the contract
	  IERC721ReceiverTransactor // Write-only binding to the contract
	  IERC721ReceiverFilterer   // Log filterer for contract events
	}

	// IERC721ReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
	type IERC721ReceiverCaller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// IERC721ReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
	type IERC721ReceiverTransactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// IERC721ReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type IERC721ReceiverFilterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// IERC721ReceiverSession is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type IERC721ReceiverSession struct {
	  Contract     *IERC721Receiver        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// IERC721ReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type IERC721ReceiverCallerSession struct {
	  Contract *IERC721ReceiverCaller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// IERC721ReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type IERC721ReceiverTransactorSession struct {
	  Contract     *IERC721ReceiverTransactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// IERC721ReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
	type IERC721ReceiverRaw struct {
	  Contract *IERC721Receiver // Generic contract binding to access the raw methods on
	}

	// IERC721ReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type IERC721ReceiverCallerRaw struct {
		Contract *IERC721ReceiverCaller // Generic read-only contract binding to access the raw methods on
	}

	// IERC721ReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type IERC721ReceiverTransactorRaw struct {
		Contract *IERC721ReceiverTransactor // Generic write-only contract binding to access the raw methods on
	}

	// NewIERC721Receiver creates a new instance of IERC721Receiver, bound to a specific deployed contract.
	func NewIERC721Receiver(address common.Address, backend bind.ContractBackend) (*IERC721Receiver, error) {
	  contract, err := bindIERC721Receiver(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &IERC721Receiver{ IERC721ReceiverCaller: IERC721ReceiverCaller{contract: contract}, IERC721ReceiverTransactor: IERC721ReceiverTransactor{contract: contract}, IERC721ReceiverFilterer: IERC721ReceiverFilterer{contract: contract} }, nil
	}

	// NewIERC721ReceiverCaller creates a new read-only instance of IERC721Receiver, bound to a specific deployed contract.
	func NewIERC721ReceiverCaller(address common.Address, caller bind.ContractCaller) (*IERC721ReceiverCaller, error) {
	  contract, err := bindIERC721Receiver(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &IERC721ReceiverCaller{contract: contract}, nil
	}

	// NewIERC721ReceiverTransactor creates a new write-only instance of IERC721Receiver, bound to a specific deployed contract.
	func NewIERC721ReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC721ReceiverTransactor, error) {
	  contract, err := bindIERC721Receiver(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &IERC721ReceiverTransactor{contract: contract}, nil
	}

	// NewIERC721ReceiverFilterer creates a new log filterer instance of IERC721Receiver, bound to a specific deployed contract.
 	func NewIERC721ReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC721ReceiverFilterer, error) {
 	  contract, err := bindIERC721Receiver(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &IERC721ReceiverFilterer{contract: contract}, nil
 	}

	// bindIERC721Receiver binds a generic wrapper to an already deployed contract.
	func bindIERC721Receiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(IERC721ReceiverABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_IERC721Receiver *IERC721ReceiverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _IERC721Receiver.Contract.IERC721ReceiverCaller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_IERC721Receiver *IERC721ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _IERC721Receiver.Contract.IERC721ReceiverTransactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_IERC721Receiver *IERC721ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _IERC721Receiver.Contract.IERC721ReceiverTransactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_IERC721Receiver *IERC721ReceiverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _IERC721Receiver.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_IERC721Receiver *IERC721ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _IERC721Receiver.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_IERC721Receiver *IERC721ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _IERC721Receiver.Contract.contract.Transact(opts, method, params...)
	}

	

	
		// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
		//
		// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
		func (_IERC721Receiver *IERC721ReceiverTransactor) OnERC721Received(opts *bind.TransactOpts , operator common.Address , from common.Address , tokenId *big.Int , data []byte ) (*types.Transaction, error) {
			return _IERC721Receiver.contract.Transact(opts, "onERC721Received" , operator, from, tokenId, data)
		}

		// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
		//
		// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
		func (_IERC721Receiver *IERC721ReceiverSession) OnERC721Received( operator common.Address , from common.Address , tokenId *big.Int , data []byte ) (*types.Transaction, error) {
		  return _IERC721Receiver.Contract.OnERC721Received(&_IERC721Receiver.TransactOpts , operator, from, tokenId, data)
		}

		// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
		//
		// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
		func (_IERC721Receiver *IERC721ReceiverTransactorSession) OnERC721Received( operator common.Address , from common.Address , tokenId *big.Int , data []byte ) (*types.Transaction, error) {
		  return _IERC721Receiver.Contract.OnERC721Received(&_IERC721Receiver.TransactOpts , operator, from, tokenId, data)
		}
	

	

	

	

	// SafeMathABI is the input ABI used to generate the binding from.
	const SafeMathABI = "[]"

	

	
		// SafeMathBin is the compiled bytecode used for deploying new contracts.
		var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582058824308d64b99c9d1edc8f56f2c7195596f80db47da09b1af2e7f89c9b6a4e264736f6c63430005110032"

		// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
		func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend ) (common.Address, *types.Transaction, *SafeMath, error) {
		  parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  
		  address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend )
		  if err != nil {
		    return common.Address{}, nil, nil, err
		  }
		  return address, tx, &SafeMath{ SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract} }, nil
		}
	

	// SafeMath is an auto generated Go binding around an Ethereum contract.
	type SafeMath struct {
	  SafeMathCaller     // Read-only binding to the contract
	  SafeMathTransactor // Write-only binding to the contract
	  SafeMathFilterer   // Log filterer for contract events
	}

	// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
	type SafeMathCaller struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
	type SafeMathTransactor struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
	type SafeMathFilterer struct {
	  contract *bind.BoundContract // Generic contract wrapper for the low level calls
	}

	// SafeMathSession is an auto generated Go binding around an Ethereum contract,
	// with pre-set call and transact options.
	type SafeMathSession struct {
	  Contract     *SafeMath        // Generic contract binding to set the session for
	  CallOpts     bind.CallOpts     // Call options to use throughout this session
	  TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
	}

	// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
	// with pre-set call options.
	type SafeMathCallerSession struct {
	  Contract *SafeMathCaller // Generic contract caller binding to set the session for
	  CallOpts bind.CallOpts    // Call options to use throughout this session
	}

	// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
	// with pre-set transact options.
	type SafeMathTransactorSession struct {
	  Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	  TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
	}

	// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
	type SafeMathRaw struct {
	  Contract *SafeMath // Generic contract binding to access the raw methods on
	}

	// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
	type SafeMathCallerRaw struct {
		Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
	}

	// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
	type SafeMathTransactorRaw struct {
		Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
	}

	// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
	func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	  contract, err := bindSafeMath(address, backend, backend, backend)
	  if err != nil {
	    return nil, err
	  }
	  return &SafeMath{ SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract} }, nil
	}

	// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
	func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	  contract, err := bindSafeMath(address, caller, nil, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &SafeMathCaller{contract: contract}, nil
	}

	// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
	func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	  contract, err := bindSafeMath(address, nil, transactor, nil)
	  if err != nil {
	    return nil, err
	  }
	  return &SafeMathTransactor{contract: contract}, nil
	}

	// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
 	func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
 	  contract, err := bindSafeMath(address, nil, nil, filterer)
 	  if err != nil {
 	    return nil, err
 	  }
 	  return &SafeMathFilterer{contract: contract}, nil
 	}

	// bindSafeMath binds a generic wrapper to an already deployed contract.
	func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	  parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	  if err != nil {
	    return nil, err
	  }
	  return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
		return _SafeMath.Contract.contract.Call(opts, result, method, params...)
	}

	// Transfer initiates a plain transaction to move funds to the contract, calling
	// its default method if one is available.
	func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
		return _SafeMath.Contract.contract.Transfer(opts)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
		return _SafeMath.Contract.contract.Transact(opts, method, params...)
	}

	

	

	

	

	


