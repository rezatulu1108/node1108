// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zevmswap

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

// ZEVMSwapAppMetaData contains all meta data concerning the ZEVMSwapApp contract.
var ZEVMSwapAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router02_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"systemContract_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LowAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"decodeMemo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetZRC20\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"encodeMemo\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"zrc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"onCrossChainCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router02\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040516200121e3803806200121e833981810160405281019062000037919062000111565b8173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1681525050505062000158565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000d982620000ac565b9050919050565b620000eb81620000cc565b8114620000f757600080fd5b50565b6000815190506200010b81620000e0565b92915050565b600080604083850312156200012b576200012a620000a7565b5b60006200013b85828601620000fa565b92505060206200014e85828601620000fa565b9150509250929050565b60805160a05161108b62000193600039600081816101b201526101fa0152600081816101d60152818161039d0152610422015261108b6000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063a06ea8bc1461005c578063bb88b7691461008d578063bd00c9c4146100ab578063c8522691146100c9578063df73044e146100e5575b600080fd5b610076600480360381019061007191906107b1565b610115565b6040516100849291906108cf565b60405180910390f35b6100956101b0565b6040516100a291906108ff565b60405180910390f35b6100b36101d4565b6040516100c091906108ff565b60405180910390f35b6100e360048036038101906100de919061097c565b6101f8565b005b6100ff60048036038101906100fa91906109f0565b610709565b60405161010c9190610a50565b60405180910390f35b600060608060008585905090506000868660009060149261013893929190610a7c565b906101439190610afb565b60601c90508686601490809261015b93929190610a7c565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505092508083945094505050509250929050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461027d576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000606061028b8484610115565b80925081935050506060600267ffffffffffffffff8111156102b0576102af610b5a565b5b6040519080825280602002602001820160405280156102de5781602001602082028036833780820191505090505b50905086816000815181106102f6576102f5610b89565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050828160018151811061034557610344610b89565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508673ffffffffffffffffffffffffffffffffffffffff1663095ea7b37f0000000000000000000000000000000000000000000000000000000000000000886040518363ffffffff1660e01b81526004016103da929190610bc7565b6020604051808303816000875af11580156103f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061041d9190610c28565b5060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166338ed17398860008530680100000000000000006040518663ffffffff1660e01b815260040161048b959493929190610d58565b6000604051808303816000875af11580156104aa573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906104d39190610ed6565b90506000808573ffffffffffffffffffffffffffffffffffffffff1663d9eeebed6040518163ffffffff1660e01b81526004016040805180830381865afa158015610522573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105469190610f34565b915091508173ffffffffffffffffffffffffffffffffffffffff1663095ea7b387836040518363ffffffff1660e01b8152600401610585929190610bc7565b6020604051808303816000875af11580156105a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c89190610c28565b508573ffffffffffffffffffffffffffffffffffffffff1663095ea7b387856001815181106105fa576105f9610b89565b5b60200260200101516040518363ffffffff1660e01b815260040161061f929190610bc7565b6020604051808303816000875af115801561063e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106629190610c28565b508573ffffffffffffffffffffffffffffffffffffffff1663c7012626868560018151811061069457610693610b89565b5b60200260200101516040518363ffffffff1660e01b81526004016106b9929190610f74565b6020604051808303816000875af11580156106d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106fc9190610c28565b5050505050505050505050565b60608383836040516020016107209392919061102b565b60405160208183030381529060405290509392505050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f8401126107715761077061074c565b5b8235905067ffffffffffffffff81111561078e5761078d610751565b5b6020830191508360018202830111156107aa576107a9610756565b5b9250929050565b600080602083850312156107c8576107c7610742565b5b600083013567ffffffffffffffff8111156107e6576107e5610747565b5b6107f28582860161075b565b92509250509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610829826107fe565b9050919050565b6108398161081e565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561087957808201518184015260208101905061085e565b60008484015250505050565b6000601f19601f8301169050919050565b60006108a18261083f565b6108ab818561084a565b93506108bb81856020860161085b565b6108c481610885565b840191505092915050565b60006040820190506108e46000830185610830565b81810360208301526108f68184610896565b90509392505050565b60006020820190506109146000830184610830565b92915050565b6109238161081e565b811461092e57600080fd5b50565b6000813590506109408161091a565b92915050565b6000819050919050565b61095981610946565b811461096457600080fd5b50565b60008135905061097681610950565b92915050565b6000806000806060858703121561099657610995610742565b5b60006109a487828801610931565b94505060206109b587828801610967565b935050604085013567ffffffffffffffff8111156109d6576109d5610747565b5b6109e28782880161075b565b925092505092959194509250565b600080600060408486031215610a0957610a08610742565b5b6000610a1786828701610931565b935050602084013567ffffffffffffffff811115610a3857610a37610747565b5b610a448682870161075b565b92509250509250925092565b60006020820190508181036000830152610a6a8184610896565b905092915050565b600080fd5b600080fd5b60008085851115610a9057610a8f610a72565b5b83861115610aa157610aa0610a77565b5b6001850283019150848603905094509492505050565b600082905092915050565b60007fffffffffffffffffffffffffffffffffffffffff00000000000000000000000082169050919050565b600082821b905092915050565b6000610b078383610ab7565b82610b128135610ac2565b92506014821015610b5257610b4d7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000083601403600802610aee565b831692505b505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b610bc181610946565b82525050565b6000604082019050610bdc6000830185610830565b610be96020830184610bb8565b9392505050565b60008115159050919050565b610c0581610bf0565b8114610c1057600080fd5b50565b600081519050610c2281610bfc565b92915050565b600060208284031215610c3e57610c3d610742565b5b6000610c4c84828501610c13565b91505092915050565b6000819050919050565b6000819050919050565b6000610c84610c7f610c7a84610c55565b610c5f565b610946565b9050919050565b610c9481610c69565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610ccf8161081e565b82525050565b6000610ce18383610cc6565b60208301905092915050565b6000602082019050919050565b6000610d0582610c9a565b610d0f8185610ca5565b9350610d1a83610cb6565b8060005b83811015610d4b578151610d328882610cd5565b9750610d3d83610ced565b925050600181019050610d1e565b5085935050505092915050565b600060a082019050610d6d6000830188610bb8565b610d7a6020830187610c8b565b8181036040830152610d8c8186610cfa565b9050610d9b6060830185610830565b610da86080830184610bb8565b9695505050505050565b610dbb82610885565b810181811067ffffffffffffffff82111715610dda57610dd9610b5a565b5b80604052505050565b6000610ded610738565b9050610df98282610db2565b919050565b600067ffffffffffffffff821115610e1957610e18610b5a565b5b602082029050602081019050919050565b600081519050610e3981610950565b92915050565b6000610e52610e4d84610dfe565b610de3565b90508083825260208201905060208402830185811115610e7557610e74610756565b5b835b81811015610e9e5780610e8a8882610e2a565b845260208401935050602081019050610e77565b5050509392505050565b600082601f830112610ebd57610ebc61074c565b5b8151610ecd848260208601610e3f565b91505092915050565b600060208284031215610eec57610eeb610742565b5b600082015167ffffffffffffffff811115610f0a57610f09610747565b5b610f1684828501610ea8565b91505092915050565b600081519050610f2e8161091a565b92915050565b60008060408385031215610f4b57610f4a610742565b5b6000610f5985828601610f1f565b9250506020610f6a85828601610e2a565b9150509250929050565b60006040820190508181036000830152610f8e8185610896565b9050610f9d6020830184610bb8565b9392505050565b60008160601b9050919050565b6000610fbc82610fa4565b9050919050565b6000610fce82610fb1565b9050919050565b610fe6610fe18261081e565b610fc3565b82525050565b600081905092915050565b82818337600083830152505050565b60006110128385610fec565b935061101f838584610ff7565b82840190509392505050565b60006110378286610fd5565b601482019150611048828486611006565b915081905094935050505056fea2646970667358221220d849874c97d6127643ba0363fe1dec561e63a7a801178715667509df9a3dcef364736f6c63430008120033",
}

// ZEVMSwapAppABI is the input ABI used to generate the binding from.
// Deprecated: Use ZEVMSwapAppMetaData.ABI instead.
var ZEVMSwapAppABI = ZEVMSwapAppMetaData.ABI

// ZEVMSwapAppBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZEVMSwapAppMetaData.Bin instead.
var ZEVMSwapAppBin = ZEVMSwapAppMetaData.Bin

// DeployZEVMSwapApp deploys a new Ethereum contract, binding an instance of ZEVMSwapApp to it.
func DeployZEVMSwapApp(auth *bind.TransactOpts, backend bind.ContractBackend, router02_ common.Address, systemContract_ common.Address) (common.Address, *types.Transaction, *ZEVMSwapApp, error) {
	parsed, err := ZEVMSwapAppMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZEVMSwapAppBin), backend, router02_, systemContract_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZEVMSwapApp{ZEVMSwapAppCaller: ZEVMSwapAppCaller{contract: contract}, ZEVMSwapAppTransactor: ZEVMSwapAppTransactor{contract: contract}, ZEVMSwapAppFilterer: ZEVMSwapAppFilterer{contract: contract}}, nil
}

// ZEVMSwapApp is an auto generated Go binding around an Ethereum contract.
type ZEVMSwapApp struct {
	ZEVMSwapAppCaller     // Read-only binding to the contract
	ZEVMSwapAppTransactor // Write-only binding to the contract
	ZEVMSwapAppFilterer   // Log filterer for contract events
}

// ZEVMSwapAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZEVMSwapAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZEVMSwapAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZEVMSwapAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZEVMSwapAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZEVMSwapAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZEVMSwapAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZEVMSwapAppSession struct {
	Contract     *ZEVMSwapApp      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZEVMSwapAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZEVMSwapAppCallerSession struct {
	Contract *ZEVMSwapAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ZEVMSwapAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZEVMSwapAppTransactorSession struct {
	Contract     *ZEVMSwapAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ZEVMSwapAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZEVMSwapAppRaw struct {
	Contract *ZEVMSwapApp // Generic contract binding to access the raw methods on
}

// ZEVMSwapAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZEVMSwapAppCallerRaw struct {
	Contract *ZEVMSwapAppCaller // Generic read-only contract binding to access the raw methods on
}

// ZEVMSwapAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZEVMSwapAppTransactorRaw struct {
	Contract *ZEVMSwapAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZEVMSwapApp creates a new instance of ZEVMSwapApp, bound to a specific deployed contract.
func NewZEVMSwapApp(address common.Address, backend bind.ContractBackend) (*ZEVMSwapApp, error) {
	contract, err := bindZEVMSwapApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZEVMSwapApp{ZEVMSwapAppCaller: ZEVMSwapAppCaller{contract: contract}, ZEVMSwapAppTransactor: ZEVMSwapAppTransactor{contract: contract}, ZEVMSwapAppFilterer: ZEVMSwapAppFilterer{contract: contract}}, nil
}

// NewZEVMSwapAppCaller creates a new read-only instance of ZEVMSwapApp, bound to a specific deployed contract.
func NewZEVMSwapAppCaller(address common.Address, caller bind.ContractCaller) (*ZEVMSwapAppCaller, error) {
	contract, err := bindZEVMSwapApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZEVMSwapAppCaller{contract: contract}, nil
}

// NewZEVMSwapAppTransactor creates a new write-only instance of ZEVMSwapApp, bound to a specific deployed contract.
func NewZEVMSwapAppTransactor(address common.Address, transactor bind.ContractTransactor) (*ZEVMSwapAppTransactor, error) {
	contract, err := bindZEVMSwapApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZEVMSwapAppTransactor{contract: contract}, nil
}

// NewZEVMSwapAppFilterer creates a new log filterer instance of ZEVMSwapApp, bound to a specific deployed contract.
func NewZEVMSwapAppFilterer(address common.Address, filterer bind.ContractFilterer) (*ZEVMSwapAppFilterer, error) {
	contract, err := bindZEVMSwapApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZEVMSwapAppFilterer{contract: contract}, nil
}

// bindZEVMSwapApp binds a generic wrapper to an already deployed contract.
func bindZEVMSwapApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZEVMSwapAppMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZEVMSwapApp *ZEVMSwapAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZEVMSwapApp.Contract.ZEVMSwapAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZEVMSwapApp *ZEVMSwapAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZEVMSwapApp.Contract.ZEVMSwapAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZEVMSwapApp *ZEVMSwapAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZEVMSwapApp.Contract.ZEVMSwapAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZEVMSwapApp *ZEVMSwapAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZEVMSwapApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZEVMSwapApp *ZEVMSwapAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZEVMSwapApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZEVMSwapApp *ZEVMSwapAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZEVMSwapApp.Contract.contract.Transact(opts, method, params...)
}

// DecodeMemo is a free data retrieval call binding the contract method 0xa06ea8bc.
//
// Solidity: function decodeMemo(bytes data) pure returns(address, bytes)
func (_ZEVMSwapApp *ZEVMSwapAppCaller) DecodeMemo(opts *bind.CallOpts, data []byte) (common.Address, []byte, error) {
	var out []interface{}
	err := _ZEVMSwapApp.contract.Call(opts, &out, "decodeMemo", data)

	if err != nil {
		return *new(common.Address), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// DecodeMemo is a free data retrieval call binding the contract method 0xa06ea8bc.
//
// Solidity: function decodeMemo(bytes data) pure returns(address, bytes)
func (_ZEVMSwapApp *ZEVMSwapAppSession) DecodeMemo(data []byte) (common.Address, []byte, error) {
	return _ZEVMSwapApp.Contract.DecodeMemo(&_ZEVMSwapApp.CallOpts, data)
}

// DecodeMemo is a free data retrieval call binding the contract method 0xa06ea8bc.
//
// Solidity: function decodeMemo(bytes data) pure returns(address, bytes)
func (_ZEVMSwapApp *ZEVMSwapAppCallerSession) DecodeMemo(data []byte) (common.Address, []byte, error) {
	return _ZEVMSwapApp.Contract.DecodeMemo(&_ZEVMSwapApp.CallOpts, data)
}

// EncodeMemo is a free data retrieval call binding the contract method 0xdf73044e.
//
// Solidity: function encodeMemo(address targetZRC20, bytes recipient) pure returns(bytes)
func (_ZEVMSwapApp *ZEVMSwapAppCaller) EncodeMemo(opts *bind.CallOpts, targetZRC20 common.Address, recipient []byte) ([]byte, error) {
	var out []interface{}
	err := _ZEVMSwapApp.contract.Call(opts, &out, "encodeMemo", targetZRC20, recipient)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeMemo is a free data retrieval call binding the contract method 0xdf73044e.
//
// Solidity: function encodeMemo(address targetZRC20, bytes recipient) pure returns(bytes)
func (_ZEVMSwapApp *ZEVMSwapAppSession) EncodeMemo(targetZRC20 common.Address, recipient []byte) ([]byte, error) {
	return _ZEVMSwapApp.Contract.EncodeMemo(&_ZEVMSwapApp.CallOpts, targetZRC20, recipient)
}

// EncodeMemo is a free data retrieval call binding the contract method 0xdf73044e.
//
// Solidity: function encodeMemo(address targetZRC20, bytes recipient) pure returns(bytes)
func (_ZEVMSwapApp *ZEVMSwapAppCallerSession) EncodeMemo(targetZRC20 common.Address, recipient []byte) ([]byte, error) {
	return _ZEVMSwapApp.Contract.EncodeMemo(&_ZEVMSwapApp.CallOpts, targetZRC20, recipient)
}

// Router02 is a free data retrieval call binding the contract method 0xbd00c9c4.
//
// Solidity: function router02() view returns(address)
func (_ZEVMSwapApp *ZEVMSwapAppCaller) Router02(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZEVMSwapApp.contract.Call(opts, &out, "router02")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router02 is a free data retrieval call binding the contract method 0xbd00c9c4.
//
// Solidity: function router02() view returns(address)
func (_ZEVMSwapApp *ZEVMSwapAppSession) Router02() (common.Address, error) {
	return _ZEVMSwapApp.Contract.Router02(&_ZEVMSwapApp.CallOpts)
}

// Router02 is a free data retrieval call binding the contract method 0xbd00c9c4.
//
// Solidity: function router02() view returns(address)
func (_ZEVMSwapApp *ZEVMSwapAppCallerSession) Router02() (common.Address, error) {
	return _ZEVMSwapApp.Contract.Router02(&_ZEVMSwapApp.CallOpts)
}

// SystemContract is a free data retrieval call binding the contract method 0xbb88b769.
//
// Solidity: function systemContract() view returns(address)
func (_ZEVMSwapApp *ZEVMSwapAppCaller) SystemContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZEVMSwapApp.contract.Call(opts, &out, "systemContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemContract is a free data retrieval call binding the contract method 0xbb88b769.
//
// Solidity: function systemContract() view returns(address)
func (_ZEVMSwapApp *ZEVMSwapAppSession) SystemContract() (common.Address, error) {
	return _ZEVMSwapApp.Contract.SystemContract(&_ZEVMSwapApp.CallOpts)
}

// SystemContract is a free data retrieval call binding the contract method 0xbb88b769.
//
// Solidity: function systemContract() view returns(address)
func (_ZEVMSwapApp *ZEVMSwapAppCallerSession) SystemContract() (common.Address, error) {
	return _ZEVMSwapApp.Contract.SystemContract(&_ZEVMSwapApp.CallOpts)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xc8522691.
//
// Solidity: function onCrossChainCall(address zrc20, uint256 amount, bytes message) returns()
func (_ZEVMSwapApp *ZEVMSwapAppTransactor) OnCrossChainCall(opts *bind.TransactOpts, zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ZEVMSwapApp.contract.Transact(opts, "onCrossChainCall", zrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xc8522691.
//
// Solidity: function onCrossChainCall(address zrc20, uint256 amount, bytes message) returns()
func (_ZEVMSwapApp *ZEVMSwapAppSession) OnCrossChainCall(zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ZEVMSwapApp.Contract.OnCrossChainCall(&_ZEVMSwapApp.TransactOpts, zrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xc8522691.
//
// Solidity: function onCrossChainCall(address zrc20, uint256 amount, bytes message) returns()
func (_ZEVMSwapApp *ZEVMSwapAppTransactorSession) OnCrossChainCall(zrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _ZEVMSwapApp.Contract.OnCrossChainCall(&_ZEVMSwapApp.TransactOpts, zrc20, amount, message)
}
