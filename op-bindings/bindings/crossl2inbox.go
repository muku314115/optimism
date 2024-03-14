// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// ICrossL2InboxIdentifier is an auto generated low-level Go binding around an user-defined struct.
type ICrossL2InboxIdentifier struct {
	Origin      common.Address
	Blocknumber *big.Int
	LogIndex    *big.Int
	Timestamp   *big.Int
	ChainId     *big.Int
}

// CrossL2InboxMetaData contains all meta data concerning the CrossL2Inbox contract.
var CrossL2InboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"BLOCKNUMBER_SLOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CHAINID_SLOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LOG_INDEX_SLOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ORIGIN_SLOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TIMESTAMP_SLOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blocknumber\",\"inputs\":[],\"outputs\":[{\"name\":\"_blocknumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chainId\",\"inputs\":[],\"outputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeMessage\",\"inputs\":[{\"name\":\"_id\",\"type\":\"tuple\",\"internalType\":\"structICrossL2Inbox.Identifier\",\"components\":[{\"name\":\"origin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blocknumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"logIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_msg\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"logIndex\",\"inputs\":[],\"outputs\":[{\"name\":\"_logIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"origin\",\"inputs\":[],\"outputs\":[{\"name\":\"_origin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"timestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x608060405234801561000f575f80fd5b5061089d8061001d5f395ff3fe6080604052600436106100c3575f3560e01c80635984c53e116100715780639a8a05921161004c5780639a8a0592146102c3578063b80777ea146102f6578063da99f72914610329575f80fd5b80635984c53e1461022a57806379d626801461023f578063938b5f3214610272575f80fd5b8063122f8b66116100a1578063122f8b661461016f578063260e6413146101a257806354fd4d50146101d5575f80fd5b806305062247146100c757806307049933146101095780630f04cf1b1461013c575b5f80fd5b3480156100d2575f80fd5b507f5a1da0738b7fdc60047c07bb519beb02aa32a8619de57e6258da1f1c2e020ccc5c5b6040519081526020015b60405180910390f35b348015610114575f80fd5b506100f67f6e0446e8b5098b8c8193f964f1b567ec3a2bdaeba33d36acb85c1f1d3f92d31381565b348015610147575f80fd5b506100f67f5a1da0738b7fdc60047c07bb519beb02aa32a8619de57e6258da1f1c2e020ccc81565b34801561017a575f80fd5b506100f67f2e148a404a50bb94820b576997fd6450117132387be615e460fa8c5e11777e0281565b3480156101ad575f80fd5b506100f67fd2b7c5071ec59eb3ff0017d703a8ea513a7d0da4779b0dbefe845808c300c81581565b3480156101e0575f80fd5b5061021d6040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161010091906106c9565b61023d610238366004610760565b61035c565b005b34801561024a575f80fd5b506100f67fab8acc221aecea88a685fabca5b88bf3823b05f335b7b9f721ca7fe3ffb2c30d81565b34801561027d575f80fd5b5060405173ffffffffffffffffffffffffffffffffffffffff7fd2b7c5071ec59eb3ff0017d703a8ea513a7d0da4779b0dbefe845808c300c8155c168152602001610100565b3480156102ce575f80fd5b507f6e0446e8b5098b8c8193f964f1b567ec3a2bdaeba33d36acb85c1f1d3f92d3135c6100f6565b348015610301575f80fd5b507f2e148a404a50bb94820b576997fd6450117132387be615e460fa8c5e11777e025c6100f6565b348015610334575f80fd5b507fab8acc221aecea88a685fabca5b88bf3823b05f335b7b9f721ca7fe3ffb2c30d5c6100f6565b3332146103ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f43726f73734c32496e626f783a206e6f7420454f412073656e6465720000000060448201526064015b60405180910390fd5b428360600135111561045e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f43726f73734c32496e626f783a20696e76616c69642069642074696d6573746160448201527f6d7000000000000000000000000000000000000000000000000000000000000060648201526084016103c1565b6040517fe38bbc32000000000000000000000000000000000000000000000000000000008152608084013560048201527342000000000000000000000000000000000000159063e38bbc3290602401602060405180830381865afa1580156104c8573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104ec919061086a565b610578576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f43726f73734c32496e626f783a20696420636861696e206e6f7420696e20646560448201527f70656e64656e637920736574000000000000000000000000000000000000000060648201526084016103c1565b6105806105fa565b5f61058b83836106b5565b9050806105f4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f43726f73734c32496e626f783a207461726765742063616c6c206661696c656460448201526064016103c1565b50505050565b6004357fd2b7c5071ec59eb3ff0017d703a8ea513a7d0da4779b0dbefe845808c300c8155d6024357f5a1da0738b7fdc60047c07bb519beb02aa32a8619de57e6258da1f1c2e020ccc5d6044357fab8acc221aecea88a685fabca5b88bf3823b05f335b7b9f721ca7fe3ffb2c30d5d6064357f2e148a404a50bb94820b576997fd6450117132387be615e460fa8c5e11777e025d6084357f6e0446e8b5098b8c8193f964f1b567ec3a2bdaeba33d36acb85c1f1d3f92d3135d565b5f805f83516020850134875af19392505050565b5f602080835283518060208501525f5b818110156106f5578581018301518582016040015282016106d9565b505f6040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f805f83850360e0811215610773575f80fd5b60a0811215610780575f80fd5b5083925060a084013573ffffffffffffffffffffffffffffffffffffffff811681146107aa575f80fd5b915060c084013567ffffffffffffffff808211156107c6575f80fd5b818601915086601f8301126107d9575f80fd5b8135818111156107eb576107eb610733565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561083157610831610733565b81604052828152896020848701011115610849575f80fd5b826020860160208301375f6020848301015280955050505050509250925092565b5f6020828403121561087a575f80fd5b81518015158114610889575f80fd5b939250505056fea164736f6c6343000818000a",
}

// CrossL2InboxABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossL2InboxMetaData.ABI instead.
var CrossL2InboxABI = CrossL2InboxMetaData.ABI

// CrossL2InboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossL2InboxMetaData.Bin instead.
var CrossL2InboxBin = CrossL2InboxMetaData.Bin

// DeployCrossL2Inbox deploys a new Ethereum contract, binding an instance of CrossL2Inbox to it.
func DeployCrossL2Inbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CrossL2Inbox, error) {
	parsed, err := CrossL2InboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossL2InboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossL2Inbox{CrossL2InboxCaller: CrossL2InboxCaller{contract: contract}, CrossL2InboxTransactor: CrossL2InboxTransactor{contract: contract}, CrossL2InboxFilterer: CrossL2InboxFilterer{contract: contract}}, nil
}

// CrossL2Inbox is an auto generated Go binding around an Ethereum contract.
type CrossL2Inbox struct {
	CrossL2InboxCaller     // Read-only binding to the contract
	CrossL2InboxTransactor // Write-only binding to the contract
	CrossL2InboxFilterer   // Log filterer for contract events
}

// CrossL2InboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossL2InboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossL2InboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossL2InboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossL2InboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossL2InboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossL2InboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossL2InboxSession struct {
	Contract     *CrossL2Inbox     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossL2InboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossL2InboxCallerSession struct {
	Contract *CrossL2InboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CrossL2InboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossL2InboxTransactorSession struct {
	Contract     *CrossL2InboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CrossL2InboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossL2InboxRaw struct {
	Contract *CrossL2Inbox // Generic contract binding to access the raw methods on
}

// CrossL2InboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossL2InboxCallerRaw struct {
	Contract *CrossL2InboxCaller // Generic read-only contract binding to access the raw methods on
}

// CrossL2InboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossL2InboxTransactorRaw struct {
	Contract *CrossL2InboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossL2Inbox creates a new instance of CrossL2Inbox, bound to a specific deployed contract.
func NewCrossL2Inbox(address common.Address, backend bind.ContractBackend) (*CrossL2Inbox, error) {
	contract, err := bindCrossL2Inbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossL2Inbox{CrossL2InboxCaller: CrossL2InboxCaller{contract: contract}, CrossL2InboxTransactor: CrossL2InboxTransactor{contract: contract}, CrossL2InboxFilterer: CrossL2InboxFilterer{contract: contract}}, nil
}

// NewCrossL2InboxCaller creates a new read-only instance of CrossL2Inbox, bound to a specific deployed contract.
func NewCrossL2InboxCaller(address common.Address, caller bind.ContractCaller) (*CrossL2InboxCaller, error) {
	contract, err := bindCrossL2Inbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossL2InboxCaller{contract: contract}, nil
}

// NewCrossL2InboxTransactor creates a new write-only instance of CrossL2Inbox, bound to a specific deployed contract.
func NewCrossL2InboxTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossL2InboxTransactor, error) {
	contract, err := bindCrossL2Inbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossL2InboxTransactor{contract: contract}, nil
}

// NewCrossL2InboxFilterer creates a new log filterer instance of CrossL2Inbox, bound to a specific deployed contract.
func NewCrossL2InboxFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossL2InboxFilterer, error) {
	contract, err := bindCrossL2Inbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossL2InboxFilterer{contract: contract}, nil
}

// bindCrossL2Inbox binds a generic wrapper to an already deployed contract.
func bindCrossL2Inbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrossL2InboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossL2Inbox *CrossL2InboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossL2Inbox.Contract.CrossL2InboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossL2Inbox *CrossL2InboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossL2Inbox.Contract.CrossL2InboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossL2Inbox *CrossL2InboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossL2Inbox.Contract.CrossL2InboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossL2Inbox *CrossL2InboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossL2Inbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossL2Inbox *CrossL2InboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossL2Inbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossL2Inbox *CrossL2InboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossL2Inbox.Contract.contract.Transact(opts, method, params...)
}

// BLOCKNUMBERSLOT is a free data retrieval call binding the contract method 0x0f04cf1b.
//
// Solidity: function BLOCKNUMBER_SLOT() view returns(bytes32)
func (_CrossL2Inbox *CrossL2InboxCaller) BLOCKNUMBERSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossL2Inbox.contract.Call(opts, &out, "BLOCKNUMBER_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BLOCKNUMBERSLOT is a free data retrieval call binding the contract method 0x0f04cf1b.
//
// Solidity: function BLOCKNUMBER_SLOT() view returns(bytes32)
func (_CrossL2Inbox *CrossL2InboxSession) BLOCKNUMBERSLOT() ([32]byte, error) {
	return _CrossL2Inbox.Contract.BLOCKNUMBERSLOT(&_CrossL2Inbox.CallOpts)
}

// BLOCKNUMBERSLOT is a free data retrieval call binding the contract method 0x0f04cf1b.
//
// Solidity: function BLOCKNUMBER_SLOT() view returns(bytes32)
func (_CrossL2Inbox *CrossL2InboxCallerSession) BLOCKNUMBERSLOT() ([32]byte, error) {
	return _CrossL2Inbox.Contract.BLOCKNUMBERSLOT(&_CrossL2Inbox.CallOpts)
}

// CHAINIDSLOT is a free data retrieval call binding the contract method 0x07049933.
//
// Solidity: function CHAINID_SLOT() view returns(bytes32)
func (_CrossL2Inbox *CrossL2InboxCaller) CHAINIDSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossL2Inbox.contract.Call(opts, &out, "CHAINID_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CHAINIDSLOT is a free data retrieval call binding the contract method 0x07049933.
//
// Solidity: function CHAINID_SLOT() view returns(bytes32)
func (_CrossL2Inbox *CrossL2InboxSession) CHAINIDSLOT() ([32]byte, error) {
	return _CrossL2Inbox.Contract.CHAINIDSLOT(&_CrossL2Inbox.CallOpts)
}

// CHAINIDSLOT is a free data retrieval call binding the contract method 0x07049933.
//
// Solidity: function CHAINID_SLOT() view returns(bytes32)
func (_CrossL2Inbox *CrossL2InboxCallerSession) CHAINIDSLOT() ([32]byte, error) {
	return _CrossL2Inbox.Contract.CHAINIDSLOT(&_CrossL2Inbox.CallOpts)
}

// LOGINDEXSLOT is a free data retrieval call binding the contract method 0x79d62680.
//
// Solidity: function LOG_INDEX_SLOT() view returns(bytes32)
func (_CrossL2Inbox *CrossL2InboxCaller) LOGINDEXSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossL2Inbox.contract.Call(opts, &out, "LOG_INDEX_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LOGINDEXSLOT is a free data retrieval call binding the contract method 0x79d62680.
//
// Solidity: function LOG_INDEX_SLOT() view returns(bytes32)
func (_CrossL2Inbox *CrossL2InboxSession) LOGINDEXSLOT() ([32]byte, error) {
	return _CrossL2Inbox.Contract.LOGINDEXSLOT(&_CrossL2Inbox.CallOpts)
}

// LOGINDEXSLOT is a free data retrieval call binding the contract method 0x79d62680.
//
// Solidity: function LOG_INDEX_SLOT() view returns(bytes32)
func (_CrossL2Inbox *CrossL2InboxCallerSession) LOGINDEXSLOT() ([32]byte, error) {
	return _CrossL2Inbox.Contract.LOGINDEXSLOT(&_CrossL2Inbox.CallOpts)
}

// ORIGINSLOT is a free data retrieval call binding the contract method 0x260e6413.
//
// Solidity: function ORIGIN_SLOT() view returns(bytes32)
func (_CrossL2Inbox *CrossL2InboxCaller) ORIGINSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossL2Inbox.contract.Call(opts, &out, "ORIGIN_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORIGINSLOT is a free data retrieval call binding the contract method 0x260e6413.
//
// Solidity: function ORIGIN_SLOT() view returns(bytes32)
func (_CrossL2Inbox *CrossL2InboxSession) ORIGINSLOT() ([32]byte, error) {
	return _CrossL2Inbox.Contract.ORIGINSLOT(&_CrossL2Inbox.CallOpts)
}

// ORIGINSLOT is a free data retrieval call binding the contract method 0x260e6413.
//
// Solidity: function ORIGIN_SLOT() view returns(bytes32)
func (_CrossL2Inbox *CrossL2InboxCallerSession) ORIGINSLOT() ([32]byte, error) {
	return _CrossL2Inbox.Contract.ORIGINSLOT(&_CrossL2Inbox.CallOpts)
}

// TIMESTAMPSLOT is a free data retrieval call binding the contract method 0x122f8b66.
//
// Solidity: function TIMESTAMP_SLOT() view returns(bytes32)
func (_CrossL2Inbox *CrossL2InboxCaller) TIMESTAMPSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossL2Inbox.contract.Call(opts, &out, "TIMESTAMP_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TIMESTAMPSLOT is a free data retrieval call binding the contract method 0x122f8b66.
//
// Solidity: function TIMESTAMP_SLOT() view returns(bytes32)
func (_CrossL2Inbox *CrossL2InboxSession) TIMESTAMPSLOT() ([32]byte, error) {
	return _CrossL2Inbox.Contract.TIMESTAMPSLOT(&_CrossL2Inbox.CallOpts)
}

// TIMESTAMPSLOT is a free data retrieval call binding the contract method 0x122f8b66.
//
// Solidity: function TIMESTAMP_SLOT() view returns(bytes32)
func (_CrossL2Inbox *CrossL2InboxCallerSession) TIMESTAMPSLOT() ([32]byte, error) {
	return _CrossL2Inbox.Contract.TIMESTAMPSLOT(&_CrossL2Inbox.CallOpts)
}

// Blocknumber is a free data retrieval call binding the contract method 0x05062247.
//
// Solidity: function blocknumber() view returns(uint256 _blocknumber)
func (_CrossL2Inbox *CrossL2InboxCaller) Blocknumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossL2Inbox.contract.Call(opts, &out, "blocknumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Blocknumber is a free data retrieval call binding the contract method 0x05062247.
//
// Solidity: function blocknumber() view returns(uint256 _blocknumber)
func (_CrossL2Inbox *CrossL2InboxSession) Blocknumber() (*big.Int, error) {
	return _CrossL2Inbox.Contract.Blocknumber(&_CrossL2Inbox.CallOpts)
}

// Blocknumber is a free data retrieval call binding the contract method 0x05062247.
//
// Solidity: function blocknumber() view returns(uint256 _blocknumber)
func (_CrossL2Inbox *CrossL2InboxCallerSession) Blocknumber() (*big.Int, error) {
	return _CrossL2Inbox.Contract.Blocknumber(&_CrossL2Inbox.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256 _chainId)
func (_CrossL2Inbox *CrossL2InboxCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossL2Inbox.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256 _chainId)
func (_CrossL2Inbox *CrossL2InboxSession) ChainId() (*big.Int, error) {
	return _CrossL2Inbox.Contract.ChainId(&_CrossL2Inbox.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256 _chainId)
func (_CrossL2Inbox *CrossL2InboxCallerSession) ChainId() (*big.Int, error) {
	return _CrossL2Inbox.Contract.ChainId(&_CrossL2Inbox.CallOpts)
}

// LogIndex is a free data retrieval call binding the contract method 0xda99f729.
//
// Solidity: function logIndex() view returns(uint256 _logIndex)
func (_CrossL2Inbox *CrossL2InboxCaller) LogIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossL2Inbox.contract.Call(opts, &out, "logIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LogIndex is a free data retrieval call binding the contract method 0xda99f729.
//
// Solidity: function logIndex() view returns(uint256 _logIndex)
func (_CrossL2Inbox *CrossL2InboxSession) LogIndex() (*big.Int, error) {
	return _CrossL2Inbox.Contract.LogIndex(&_CrossL2Inbox.CallOpts)
}

// LogIndex is a free data retrieval call binding the contract method 0xda99f729.
//
// Solidity: function logIndex() view returns(uint256 _logIndex)
func (_CrossL2Inbox *CrossL2InboxCallerSession) LogIndex() (*big.Int, error) {
	return _CrossL2Inbox.Contract.LogIndex(&_CrossL2Inbox.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address _origin)
func (_CrossL2Inbox *CrossL2InboxCaller) Origin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossL2Inbox.contract.Call(opts, &out, "origin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address _origin)
func (_CrossL2Inbox *CrossL2InboxSession) Origin() (common.Address, error) {
	return _CrossL2Inbox.Contract.Origin(&_CrossL2Inbox.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address _origin)
func (_CrossL2Inbox *CrossL2InboxCallerSession) Origin() (common.Address, error) {
	return _CrossL2Inbox.Contract.Origin(&_CrossL2Inbox.CallOpts)
}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() view returns(uint256 _timestamp)
func (_CrossL2Inbox *CrossL2InboxCaller) Timestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrossL2Inbox.contract.Call(opts, &out, "timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() view returns(uint256 _timestamp)
func (_CrossL2Inbox *CrossL2InboxSession) Timestamp() (*big.Int, error) {
	return _CrossL2Inbox.Contract.Timestamp(&_CrossL2Inbox.CallOpts)
}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() view returns(uint256 _timestamp)
func (_CrossL2Inbox *CrossL2InboxCallerSession) Timestamp() (*big.Int, error) {
	return _CrossL2Inbox.Contract.Timestamp(&_CrossL2Inbox.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CrossL2Inbox *CrossL2InboxCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossL2Inbox.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CrossL2Inbox *CrossL2InboxSession) Version() (string, error) {
	return _CrossL2Inbox.Contract.Version(&_CrossL2Inbox.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CrossL2Inbox *CrossL2InboxCallerSession) Version() (string, error) {
	return _CrossL2Inbox.Contract.Version(&_CrossL2Inbox.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x5984c53e.
//
// Solidity: function executeMessage((address,uint256,uint256,uint256,uint256) _id, address _target, bytes _msg) payable returns()
func (_CrossL2Inbox *CrossL2InboxTransactor) ExecuteMessage(opts *bind.TransactOpts, _id ICrossL2InboxIdentifier, _target common.Address, _msg []byte) (*types.Transaction, error) {
	return _CrossL2Inbox.contract.Transact(opts, "executeMessage", _id, _target, _msg)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x5984c53e.
//
// Solidity: function executeMessage((address,uint256,uint256,uint256,uint256) _id, address _target, bytes _msg) payable returns()
func (_CrossL2Inbox *CrossL2InboxSession) ExecuteMessage(_id ICrossL2InboxIdentifier, _target common.Address, _msg []byte) (*types.Transaction, error) {
	return _CrossL2Inbox.Contract.ExecuteMessage(&_CrossL2Inbox.TransactOpts, _id, _target, _msg)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x5984c53e.
//
// Solidity: function executeMessage((address,uint256,uint256,uint256,uint256) _id, address _target, bytes _msg) payable returns()
func (_CrossL2Inbox *CrossL2InboxTransactorSession) ExecuteMessage(_id ICrossL2InboxIdentifier, _target common.Address, _msg []byte) (*types.Transaction, error) {
	return _CrossL2Inbox.Contract.ExecuteMessage(&_CrossL2Inbox.TransactOpts, _id, _target, _msg)
}
