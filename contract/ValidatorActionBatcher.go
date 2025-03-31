// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"CONSOLIDATION_SYSTEM_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_SYSTEM_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"feeLimit\",\"type\":\"uint256\"}],\"name\":\"batchConsolidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"feeLimit\",\"type\":\"uint256\"}],\"name\":\"batchWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsolidationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawalFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506107628061001c5f395ff3fe608060405260043610610055575f3560e01c8063030cac801461005e57806343bcdb0614610085578063446ae019146100a457806348b40f74146100c3578063ca612a0d14610100578063d8e159f81461012557005b3661005c57005b005b348015610069575f80fd5b50610072610139565b6040519081526020015b60405180910390f35b348015610090575f80fd5b5061005c61009f36600461060a565b6101c2565b3480156100af575f80fd5b5061005c6100be36600461060a565b6103b6565b3480156100ce575f80fd5b506100e871bbddc7ce488642fb579f8b00f3a59000725181565b6040516001600160a01b03909116815260200161007c565b34801561010b575f80fd5b506100e8710961ef480eb55e80d19ad83579a64c00700281565b348015610130575f80fd5b506100726105a1565b5f805f71bbddc7ce488642fb579f8b00f3a5900072516001600160a01b03166040515b5f60405180830381855afa9150503d805f8114610194576040519150601f19603f3d011682016040523d82523d5f602084013e610199565b606091505b5091509150816101a7575f80fd5b808060200190518101906101bb919061067f565b9250505090565b6101ca6105c8565b5f6101d36105a1565b90508115806101e25750818111155b6102225760405162461bcd60e51b815260206004820152600c60248201526b0cccaca40e8dede40d0d2ced60a31b60448201526064015b60405180910390fd5b61022c8382610696565b47101561026d5760405162461bcd60e51b815260206004820152600f60248201526e62616c616e636520746f6f206c6f7760881b6044820152606401610219565b5f5b838110156103af57848482818110610289576102896106bf565b905060200281019061029b91906106d3565b90506038146102e15760405162461bcd60e51b81526020600482015260126024820152711a5b9d985b1a59081dda5d1a191c985dd85b60721b6044820152606401610219565b5f710961ef480eb55e80d19ad83579a64c00700283878785818110610308576103086106bf565b905060200281019061031a91906106d3565b60405161032892919061071d565b5f6040518083038185875af1925050503d805f8114610362576040519150601f19603f3d011682016040523d82523d5f602084013e610367565b606091505b50509050806103a65760405162461bcd60e51b815260206004820152600b60248201526a18d85b1b0819985a5b195960aa1b6044820152606401610219565b5060010161026f565b5050505050565b6103be6105c8565b5f6103c7610139565b90508115806103d65750818111155b6104115760405162461bcd60e51b815260206004820152600c60248201526b0cccaca40e8dede40d0d2ced60a31b6044820152606401610219565b61041b8382610696565b47101561045c5760405162461bcd60e51b815260206004820152600f60248201526e62616c616e636520746f6f206c6f7760881b6044820152606401610219565b5f5b838110156103af57848482818110610478576104786106bf565b905060200281019061048a91906106d3565b90506060146104d35760405162461bcd60e51b815260206004820152601560248201527434b73b30b634b21031b7b739b7b634b230ba34b7b760591b6044820152606401610219565b5f71bbddc7ce488642fb579f8b00f3a590007251838787858181106104fa576104fa6106bf565b905060200281019061050c91906106d3565b60405161051a92919061071d565b5f6040518083038185875af1925050503d805f8114610554576040519150601f19603f3d011682016040523d82523d5f602084013e610559565b606091505b50509050806105985760405162461bcd60e51b815260206004820152600b60248201526a18d85b1b0819985a5b195960aa1b6044820152606401610219565b5060010161045e565b5f805f710961ef480eb55e80d19ad83579a64c0070026001600160a01b031660405161015c565b3330146106085760405162461bcd60e51b815260206004820152600e60248201526d73656c662063616c6c206f6e6c7960901b6044820152606401610219565b565b5f805f6040848603121561061c575f80fd5b833567ffffffffffffffff811115610632575f80fd5b8401601f81018613610642575f80fd5b803567ffffffffffffffff811115610658575f80fd5b8660208260051b840101111561066c575f80fd5b6020918201979096509401359392505050565b5f6020828403121561068f575f80fd5b5051919050565b80820281158282048414176106b957634e487b7160e01b5f52601160045260245ffd5b92915050565b634e487b7160e01b5f52603260045260245ffd5b5f808335601e198436030181126106e8575f80fd5b83018035915067ffffffffffffffff821115610702575f80fd5b602001915036819003821315610716575f80fd5b9250929050565b818382375f910190815291905056fea2646970667358221220d4c00cc9f8bffa7c2fd80eacf81d38bbd3abd23dde80808e8d76b1d47a9e52a364736f6c634300081a0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// CONSOLIDATIONSYSTEMCONTRACT is a free data retrieval call binding the contract method 0x48b40f74.
//
// Solidity: function CONSOLIDATION_SYSTEM_CONTRACT() view returns(address)
func (_Contract *ContractCaller) CONSOLIDATIONSYSTEMCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "CONSOLIDATION_SYSTEM_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CONSOLIDATIONSYSTEMCONTRACT is a free data retrieval call binding the contract method 0x48b40f74.
//
// Solidity: function CONSOLIDATION_SYSTEM_CONTRACT() view returns(address)
func (_Contract *ContractSession) CONSOLIDATIONSYSTEMCONTRACT() (common.Address, error) {
	return _Contract.Contract.CONSOLIDATIONSYSTEMCONTRACT(&_Contract.CallOpts)
}

// CONSOLIDATIONSYSTEMCONTRACT is a free data retrieval call binding the contract method 0x48b40f74.
//
// Solidity: function CONSOLIDATION_SYSTEM_CONTRACT() view returns(address)
func (_Contract *ContractCallerSession) CONSOLIDATIONSYSTEMCONTRACT() (common.Address, error) {
	return _Contract.Contract.CONSOLIDATIONSYSTEMCONTRACT(&_Contract.CallOpts)
}

// WITHDRAWALSYSTEMCONTRACT is a free data retrieval call binding the contract method 0xca612a0d.
//
// Solidity: function WITHDRAWAL_SYSTEM_CONTRACT() view returns(address)
func (_Contract *ContractCaller) WITHDRAWALSYSTEMCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "WITHDRAWAL_SYSTEM_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WITHDRAWALSYSTEMCONTRACT is a free data retrieval call binding the contract method 0xca612a0d.
//
// Solidity: function WITHDRAWAL_SYSTEM_CONTRACT() view returns(address)
func (_Contract *ContractSession) WITHDRAWALSYSTEMCONTRACT() (common.Address, error) {
	return _Contract.Contract.WITHDRAWALSYSTEMCONTRACT(&_Contract.CallOpts)
}

// WITHDRAWALSYSTEMCONTRACT is a free data retrieval call binding the contract method 0xca612a0d.
//
// Solidity: function WITHDRAWAL_SYSTEM_CONTRACT() view returns(address)
func (_Contract *ContractCallerSession) WITHDRAWALSYSTEMCONTRACT() (common.Address, error) {
	return _Contract.Contract.WITHDRAWALSYSTEMCONTRACT(&_Contract.CallOpts)
}

// GetConsolidationFee is a free data retrieval call binding the contract method 0x030cac80.
//
// Solidity: function getConsolidationFee() view returns(uint256 fee)
func (_Contract *ContractCaller) GetConsolidationFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getConsolidationFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConsolidationFee is a free data retrieval call binding the contract method 0x030cac80.
//
// Solidity: function getConsolidationFee() view returns(uint256 fee)
func (_Contract *ContractSession) GetConsolidationFee() (*big.Int, error) {
	return _Contract.Contract.GetConsolidationFee(&_Contract.CallOpts)
}

// GetConsolidationFee is a free data retrieval call binding the contract method 0x030cac80.
//
// Solidity: function getConsolidationFee() view returns(uint256 fee)
func (_Contract *ContractCallerSession) GetConsolidationFee() (*big.Int, error) {
	return _Contract.Contract.GetConsolidationFee(&_Contract.CallOpts)
}

// GetWithdrawalFee is a free data retrieval call binding the contract method 0xd8e159f8.
//
// Solidity: function getWithdrawalFee() view returns(uint256 fee)
func (_Contract *ContractCaller) GetWithdrawalFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getWithdrawalFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawalFee is a free data retrieval call binding the contract method 0xd8e159f8.
//
// Solidity: function getWithdrawalFee() view returns(uint256 fee)
func (_Contract *ContractSession) GetWithdrawalFee() (*big.Int, error) {
	return _Contract.Contract.GetWithdrawalFee(&_Contract.CallOpts)
}

// GetWithdrawalFee is a free data retrieval call binding the contract method 0xd8e159f8.
//
// Solidity: function getWithdrawalFee() view returns(uint256 fee)
func (_Contract *ContractCallerSession) GetWithdrawalFee() (*big.Int, error) {
	return _Contract.Contract.GetWithdrawalFee(&_Contract.CallOpts)
}

// BatchConsolidate is a paid mutator transaction binding the contract method 0x446ae019.
//
// Solidity: function batchConsolidate(bytes[] data, uint256 feeLimit) returns()
func (_Contract *ContractTransactor) BatchConsolidate(opts *bind.TransactOpts, data [][]byte, feeLimit *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "batchConsolidate", data, feeLimit)
}

// BatchConsolidate is a paid mutator transaction binding the contract method 0x446ae019.
//
// Solidity: function batchConsolidate(bytes[] data, uint256 feeLimit) returns()
func (_Contract *ContractSession) BatchConsolidate(data [][]byte, feeLimit *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BatchConsolidate(&_Contract.TransactOpts, data, feeLimit)
}

// BatchConsolidate is a paid mutator transaction binding the contract method 0x446ae019.
//
// Solidity: function batchConsolidate(bytes[] data, uint256 feeLimit) returns()
func (_Contract *ContractTransactorSession) BatchConsolidate(data [][]byte, feeLimit *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BatchConsolidate(&_Contract.TransactOpts, data, feeLimit)
}

// BatchWithdraw is a paid mutator transaction binding the contract method 0x43bcdb06.
//
// Solidity: function batchWithdraw(bytes[] data, uint256 feeLimit) returns()
func (_Contract *ContractTransactor) BatchWithdraw(opts *bind.TransactOpts, data [][]byte, feeLimit *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "batchWithdraw", data, feeLimit)
}

// BatchWithdraw is a paid mutator transaction binding the contract method 0x43bcdb06.
//
// Solidity: function batchWithdraw(bytes[] data, uint256 feeLimit) returns()
func (_Contract *ContractSession) BatchWithdraw(data [][]byte, feeLimit *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BatchWithdraw(&_Contract.TransactOpts, data, feeLimit)
}

// BatchWithdraw is a paid mutator transaction binding the contract method 0x43bcdb06.
//
// Solidity: function batchWithdraw(bytes[] data, uint256 feeLimit) returns()
func (_Contract *ContractTransactorSession) BatchWithdraw(data [][]byte, feeLimit *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BatchWithdraw(&_Contract.TransactOpts, data, feeLimit)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Contract *ContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Contract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Contract *ContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Contract.Contract.Fallback(&_Contract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Contract *ContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Contract.Contract.Fallback(&_Contract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactorSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}
