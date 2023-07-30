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

// HyperlaneValidatorAnnounceMetaData contains all meta data concerning the HyperlaneValidatorAnnounce contract.
var HyperlaneValidatorAnnounceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mailbox\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"storageLocation\",\"type\":\"string\"}],\"name\":\"ValidatorAnnouncement\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_storageLocation\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"announce\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"}],\"name\":\"getAnnouncedStorageLocations\",\"outputs\":[{\"internalType\":\"string[][]\",\"name\":\"\",\"type\":\"string[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAnnouncedValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mailbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b5060405161138338038061138383398101604081905261002f9161004b565b6001600160a01b0390911660805263ffffffff1660a05261009a565b6000806040838503121561005e57600080fd5b82516001600160a01b038116811461007557600080fd5b602084015190925063ffffffff8116811461008f57600080fd5b809150509250929050565b60805160a0516112b76100cc6000396000818160ce015261029501526000818161010a015261027301526112b76000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c8063690cb78611610050578063690cb786146100b45780638d3638f4146100c9578063d5438eae1461010557600080fd5b806321f717811461006c57806351abe7cc14610094575b600080fd5b61007f61007a366004610cbf565b610151565b60405190151581526020015b60405180910390f35b6100a76100a2366004610d40565b61046c565b60405161008b9190610de5565b6100bc610639565b60405161008b9190610ee6565b6100f07f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161008b565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161008b565b60008086868660405160200161016993929190610f34565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600390935291205490915060ff161561021e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600660248201527f7265706c6179000000000000000000000000000000000000000000000000000060448201526064015b60405180910390fd5b600081815260036020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558051601f89018390048302810183019091528781526102d5917f0000000000000000000000000000000000000000000000000000000000000000917f0000000000000000000000000000000000000000000000000000000000000000918b908b908190840183828082843760009201919091525061064a92505050565b905060006103198287878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061077892505050565b90508873ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146103b0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600a60248201527f217369676e6174757265000000000000000000000000000000000000000000006044820152606401610215565b6103bb60008a610794565b6103cc576103ca60008a6107c6565b505b73ffffffffffffffffffffffffffffffffffffffff891660009081526002602090815260408220805460018101825590835291200161040c888a83611044565b508873ffffffffffffffffffffffffffffffffffffffff167f78066d8adb677a1353d1fc8be28cf03e2a8de7157bbab979953587d78076c11e898960405161045592919061115f565b60405180910390a250600198975050505050505050565b606060008267ffffffffffffffff81111561048957610489610f73565b6040519080825280602002602001820160405280156104bc57816020015b60608152602001906001900390816104a75790505b50905060005b8381101561062f57600260008686848181106104e0576104e06111ac565b90506020020160208101906104f591906111db565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156105fb57838290600052602060002001805461056e90610fa2565b80601f016020809104026020016040519081016040528092919081815260200182805461059a90610fa2565b80156105e75780601f106105bc576101008083540402835291602001916105e7565b820191906000526020600020905b8154815290600101906020018083116105ca57829003601f168201915b50505050508152602001906001019061054f565b50505050828281518110610611576106116111ac565b60200260200101819052508080610627906111f6565b9150506104c2565b5090505b92915050565b606061064560006107e8565b905090565b6000808373ffffffffffffffffffffffffffffffffffffffff861660405160e09290921b7fffffffff0000000000000000000000000000000000000000000000000000000016602083015260248201527f48595045524c414e455f414e4e4f554e43454d454e54000000000000000000006044820152605a0160405160208183030381529060405280519060200120905061076f81846040516020016106f1929190611255565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201207f19457468657265756d205369676e6564204d6573736167653a0a33320000000084830152603c8085019190915282518085039091018152605c909301909152815191012090565b95945050505050565b600080600061078785856107f5565b9150915061062f8161083a565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260018301602052604081205415155b9392505050565b60006107bf8373ffffffffffffffffffffffffffffffffffffffff8416610a91565b606060006107bf83610ae0565b600080825160410361082b5760208301516040840151606085015160001a61081f87828585610b3c565b94509450505050610833565b506000905060025b9250929050565b600081600481111561084e5761084e61127b565b036108565750565b600181600481111561086a5761086a61127b565b036108d1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610215565b60028160048111156108e5576108e561127b565b0361094c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610215565b60038160048111156109605761096061127b565b036109ed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610215565b6004816004811115610a0157610a0161127b565b03610a8e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610215565b50565b6000818152600183016020526040812054610ad857508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610633565b506000610633565b606081600001805480602002602001604051908101604052809291908181526020018280548015610b3057602002820191906000526020600020905b815481526020019060010190808311610b1c575b50505050509050919050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610b735750600090506003610c4b565b8460ff16601b14158015610b8b57508460ff16601c14155b15610b9c5750600090506004610c4b565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610bf0573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116610c4457600060019250925050610c4b565b9150600090505b94509492505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c7857600080fd5b919050565b60008083601f840112610c8f57600080fd5b50813567ffffffffffffffff811115610ca757600080fd5b60208301915083602082850101111561083357600080fd5b600080600080600060608688031215610cd757600080fd5b610ce086610c54565b9450602086013567ffffffffffffffff80821115610cfd57600080fd5b610d0989838a01610c7d565b90965094506040880135915080821115610d2257600080fd5b50610d2f88828901610c7d565b969995985093965092949392505050565b60008060208385031215610d5357600080fd5b823567ffffffffffffffff80821115610d6b57600080fd5b818501915085601f830112610d7f57600080fd5b813581811115610d8e57600080fd5b8660208260051b8501011115610da357600080fd5b60209290920196919550909350505050565b60005b83811015610dd0578181015183820152602001610db8565b83811115610ddf576000848401525b50505050565b60208152600060208201835180825260408401915060408160051b8501016020860160005b83811015610eda578683037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0018552815180518085526020918201918086019190600582901b87010160005b82811015610ec0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08089840301855285518051808552610e9e816020870160208501610db5565b60209788019796870196601f9190910190921693909301019150600101610e56565b506020988901989096509490940193505050600101610e0a565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b81811015610eda57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101610f02565b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008460601b168152818360148301376000910160140190815292915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600181811c90821680610fb657607f821691505b602082108103610fef577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f82111561103f57600081815260208120601f850160051c8101602086101561101c5750805b601f850160051c820191505b8181101561103b57828155600101611028565b5050505b505050565b67ffffffffffffffff83111561105c5761105c610f73565b6110708361106a8354610fa2565b83610ff5565b6000601f8411600181146110c2576000851561108c5750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b178355611158565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b8281101561111157868501358255602094850194600190920191016110f1565b508682101561114c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b60208152816020820152818360408301376000818301604090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000602082840312156111ed57600080fd5b6107bf82610c54565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361124e577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b8281526000825161126d816020850160208701610db5565b919091016020019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfea164736f6c634300080f000a",
}

// HyperlaneValidatorAnnounceABI is the input ABI used to generate the binding from.
// Deprecated: Use HyperlaneValidatorAnnounceMetaData.ABI instead.
var HyperlaneValidatorAnnounceABI = HyperlaneValidatorAnnounceMetaData.ABI

// HyperlaneValidatorAnnounceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HyperlaneValidatorAnnounceMetaData.Bin instead.
var HyperlaneValidatorAnnounceBin = HyperlaneValidatorAnnounceMetaData.Bin

// DeployHyperlaneValidatorAnnounce deploys a new Ethereum contract, binding an instance of HyperlaneValidatorAnnounce to it.
func DeployHyperlaneValidatorAnnounce(auth *bind.TransactOpts, backend bind.ContractBackend, _mailbox common.Address, _localDomain uint32) (common.Address, *types.Transaction, *HyperlaneValidatorAnnounce, error) {
	parsed, err := HyperlaneValidatorAnnounceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HyperlaneValidatorAnnounceBin), backend, _mailbox, _localDomain)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HyperlaneValidatorAnnounce{HyperlaneValidatorAnnounceCaller: HyperlaneValidatorAnnounceCaller{contract: contract}, HyperlaneValidatorAnnounceTransactor: HyperlaneValidatorAnnounceTransactor{contract: contract}, HyperlaneValidatorAnnounceFilterer: HyperlaneValidatorAnnounceFilterer{contract: contract}}, nil
}

// HyperlaneValidatorAnnounce is an auto generated Go binding around an Ethereum contract.
type HyperlaneValidatorAnnounce struct {
	HyperlaneValidatorAnnounceCaller     // Read-only binding to the contract
	HyperlaneValidatorAnnounceTransactor // Write-only binding to the contract
	HyperlaneValidatorAnnounceFilterer   // Log filterer for contract events
}

// HyperlaneValidatorAnnounceCaller is an auto generated read-only Go binding around an Ethereum contract.
type HyperlaneValidatorAnnounceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperlaneValidatorAnnounceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HyperlaneValidatorAnnounceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperlaneValidatorAnnounceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HyperlaneValidatorAnnounceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperlaneValidatorAnnounceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HyperlaneValidatorAnnounceSession struct {
	Contract     *HyperlaneValidatorAnnounce // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// HyperlaneValidatorAnnounceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HyperlaneValidatorAnnounceCallerSession struct {
	Contract *HyperlaneValidatorAnnounceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// HyperlaneValidatorAnnounceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HyperlaneValidatorAnnounceTransactorSession struct {
	Contract     *HyperlaneValidatorAnnounceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// HyperlaneValidatorAnnounceRaw is an auto generated low-level Go binding around an Ethereum contract.
type HyperlaneValidatorAnnounceRaw struct {
	Contract *HyperlaneValidatorAnnounce // Generic contract binding to access the raw methods on
}

// HyperlaneValidatorAnnounceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HyperlaneValidatorAnnounceCallerRaw struct {
	Contract *HyperlaneValidatorAnnounceCaller // Generic read-only contract binding to access the raw methods on
}

// HyperlaneValidatorAnnounceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HyperlaneValidatorAnnounceTransactorRaw struct {
	Contract *HyperlaneValidatorAnnounceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHyperlaneValidatorAnnounce creates a new instance of HyperlaneValidatorAnnounce, bound to a specific deployed contract.
func NewHyperlaneValidatorAnnounce(address common.Address, backend bind.ContractBackend) (*HyperlaneValidatorAnnounce, error) {
	contract, err := bindHyperlaneValidatorAnnounce(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HyperlaneValidatorAnnounce{HyperlaneValidatorAnnounceCaller: HyperlaneValidatorAnnounceCaller{contract: contract}, HyperlaneValidatorAnnounceTransactor: HyperlaneValidatorAnnounceTransactor{contract: contract}, HyperlaneValidatorAnnounceFilterer: HyperlaneValidatorAnnounceFilterer{contract: contract}}, nil
}

// NewHyperlaneValidatorAnnounceCaller creates a new read-only instance of HyperlaneValidatorAnnounce, bound to a specific deployed contract.
func NewHyperlaneValidatorAnnounceCaller(address common.Address, caller bind.ContractCaller) (*HyperlaneValidatorAnnounceCaller, error) {
	contract, err := bindHyperlaneValidatorAnnounce(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HyperlaneValidatorAnnounceCaller{contract: contract}, nil
}

// NewHyperlaneValidatorAnnounceTransactor creates a new write-only instance of HyperlaneValidatorAnnounce, bound to a specific deployed contract.
func NewHyperlaneValidatorAnnounceTransactor(address common.Address, transactor bind.ContractTransactor) (*HyperlaneValidatorAnnounceTransactor, error) {
	contract, err := bindHyperlaneValidatorAnnounce(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HyperlaneValidatorAnnounceTransactor{contract: contract}, nil
}

// NewHyperlaneValidatorAnnounceFilterer creates a new log filterer instance of HyperlaneValidatorAnnounce, bound to a specific deployed contract.
func NewHyperlaneValidatorAnnounceFilterer(address common.Address, filterer bind.ContractFilterer) (*HyperlaneValidatorAnnounceFilterer, error) {
	contract, err := bindHyperlaneValidatorAnnounce(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HyperlaneValidatorAnnounceFilterer{contract: contract}, nil
}

// bindHyperlaneValidatorAnnounce binds a generic wrapper to an already deployed contract.
func bindHyperlaneValidatorAnnounce(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HyperlaneValidatorAnnounceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HyperlaneValidatorAnnounce.Contract.HyperlaneValidatorAnnounceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperlaneValidatorAnnounce.Contract.HyperlaneValidatorAnnounceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HyperlaneValidatorAnnounce.Contract.HyperlaneValidatorAnnounceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HyperlaneValidatorAnnounce.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperlaneValidatorAnnounce.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HyperlaneValidatorAnnounce.Contract.contract.Transact(opts, method, params...)
}

// GetAnnouncedStorageLocations is a free data retrieval call binding the contract method 0x51abe7cc.
//
// Solidity: function getAnnouncedStorageLocations(address[] _validators) view returns(string[][])
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceCaller) GetAnnouncedStorageLocations(opts *bind.CallOpts, _validators []common.Address) ([][]string, error) {
	var out []interface{}
	err := _HyperlaneValidatorAnnounce.contract.Call(opts, &out, "getAnnouncedStorageLocations", _validators)

	if err != nil {
		return *new([][]string), err
	}

	out0 := *abi.ConvertType(out[0], new([][]string)).(*[][]string)

	return out0, err

}

// GetAnnouncedStorageLocations is a free data retrieval call binding the contract method 0x51abe7cc.
//
// Solidity: function getAnnouncedStorageLocations(address[] _validators) view returns(string[][])
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceSession) GetAnnouncedStorageLocations(_validators []common.Address) ([][]string, error) {
	return _HyperlaneValidatorAnnounce.Contract.GetAnnouncedStorageLocations(&_HyperlaneValidatorAnnounce.CallOpts, _validators)
}

// GetAnnouncedStorageLocations is a free data retrieval call binding the contract method 0x51abe7cc.
//
// Solidity: function getAnnouncedStorageLocations(address[] _validators) view returns(string[][])
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceCallerSession) GetAnnouncedStorageLocations(_validators []common.Address) ([][]string, error) {
	return _HyperlaneValidatorAnnounce.Contract.GetAnnouncedStorageLocations(&_HyperlaneValidatorAnnounce.CallOpts, _validators)
}

// GetAnnouncedValidators is a free data retrieval call binding the contract method 0x690cb786.
//
// Solidity: function getAnnouncedValidators() view returns(address[])
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceCaller) GetAnnouncedValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _HyperlaneValidatorAnnounce.contract.Call(opts, &out, "getAnnouncedValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAnnouncedValidators is a free data retrieval call binding the contract method 0x690cb786.
//
// Solidity: function getAnnouncedValidators() view returns(address[])
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceSession) GetAnnouncedValidators() ([]common.Address, error) {
	return _HyperlaneValidatorAnnounce.Contract.GetAnnouncedValidators(&_HyperlaneValidatorAnnounce.CallOpts)
}

// GetAnnouncedValidators is a free data retrieval call binding the contract method 0x690cb786.
//
// Solidity: function getAnnouncedValidators() view returns(address[])
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceCallerSession) GetAnnouncedValidators() ([]common.Address, error) {
	return _HyperlaneValidatorAnnounce.Contract.GetAnnouncedValidators(&_HyperlaneValidatorAnnounce.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _HyperlaneValidatorAnnounce.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceSession) LocalDomain() (uint32, error) {
	return _HyperlaneValidatorAnnounce.Contract.LocalDomain(&_HyperlaneValidatorAnnounce.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceCallerSession) LocalDomain() (uint32, error) {
	return _HyperlaneValidatorAnnounce.Contract.LocalDomain(&_HyperlaneValidatorAnnounce.CallOpts)
}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceCaller) Mailbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HyperlaneValidatorAnnounce.contract.Call(opts, &out, "mailbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceSession) Mailbox() (common.Address, error) {
	return _HyperlaneValidatorAnnounce.Contract.Mailbox(&_HyperlaneValidatorAnnounce.CallOpts)
}

// Mailbox is a free data retrieval call binding the contract method 0xd5438eae.
//
// Solidity: function mailbox() view returns(address)
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceCallerSession) Mailbox() (common.Address, error) {
	return _HyperlaneValidatorAnnounce.Contract.Mailbox(&_HyperlaneValidatorAnnounce.CallOpts)
}

// Announce is a paid mutator transaction binding the contract method 0x21f71781.
//
// Solidity: function announce(address _validator, string _storageLocation, bytes _signature) returns(bool)
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceTransactor) Announce(opts *bind.TransactOpts, _validator common.Address, _storageLocation string, _signature []byte) (*types.Transaction, error) {
	return _HyperlaneValidatorAnnounce.contract.Transact(opts, "announce", _validator, _storageLocation, _signature)
}

// Announce is a paid mutator transaction binding the contract method 0x21f71781.
//
// Solidity: function announce(address _validator, string _storageLocation, bytes _signature) returns(bool)
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceSession) Announce(_validator common.Address, _storageLocation string, _signature []byte) (*types.Transaction, error) {
	return _HyperlaneValidatorAnnounce.Contract.Announce(&_HyperlaneValidatorAnnounce.TransactOpts, _validator, _storageLocation, _signature)
}

// Announce is a paid mutator transaction binding the contract method 0x21f71781.
//
// Solidity: function announce(address _validator, string _storageLocation, bytes _signature) returns(bool)
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceTransactorSession) Announce(_validator common.Address, _storageLocation string, _signature []byte) (*types.Transaction, error) {
	return _HyperlaneValidatorAnnounce.Contract.Announce(&_HyperlaneValidatorAnnounce.TransactOpts, _validator, _storageLocation, _signature)
}

// HyperlaneValidatorAnnounceValidatorAnnouncementIterator is returned from FilterValidatorAnnouncement and is used to iterate over the raw logs and unpacked data for ValidatorAnnouncement events raised by the HyperlaneValidatorAnnounce contract.
type HyperlaneValidatorAnnounceValidatorAnnouncementIterator struct {
	Event *HyperlaneValidatorAnnounceValidatorAnnouncement // Event containing the contract specifics and raw log

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
func (it *HyperlaneValidatorAnnounceValidatorAnnouncementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperlaneValidatorAnnounceValidatorAnnouncement)
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
		it.Event = new(HyperlaneValidatorAnnounceValidatorAnnouncement)
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
func (it *HyperlaneValidatorAnnounceValidatorAnnouncementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperlaneValidatorAnnounceValidatorAnnouncementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperlaneValidatorAnnounceValidatorAnnouncement represents a ValidatorAnnouncement event raised by the HyperlaneValidatorAnnounce contract.
type HyperlaneValidatorAnnounceValidatorAnnouncement struct {
	Validator       common.Address
	StorageLocation string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorAnnouncement is a free log retrieval operation binding the contract event 0x78066d8adb677a1353d1fc8be28cf03e2a8de7157bbab979953587d78076c11e.
//
// Solidity: event ValidatorAnnouncement(address indexed validator, string storageLocation)
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceFilterer) FilterValidatorAnnouncement(opts *bind.FilterOpts, validator []common.Address) (*HyperlaneValidatorAnnounceValidatorAnnouncementIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _HyperlaneValidatorAnnounce.contract.FilterLogs(opts, "ValidatorAnnouncement", validatorRule)
	if err != nil {
		return nil, err
	}
	return &HyperlaneValidatorAnnounceValidatorAnnouncementIterator{contract: _HyperlaneValidatorAnnounce.contract, event: "ValidatorAnnouncement", logs: logs, sub: sub}, nil
}

// WatchValidatorAnnouncement is a free log subscription operation binding the contract event 0x78066d8adb677a1353d1fc8be28cf03e2a8de7157bbab979953587d78076c11e.
//
// Solidity: event ValidatorAnnouncement(address indexed validator, string storageLocation)
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceFilterer) WatchValidatorAnnouncement(opts *bind.WatchOpts, sink chan<- *HyperlaneValidatorAnnounceValidatorAnnouncement, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _HyperlaneValidatorAnnounce.contract.WatchLogs(opts, "ValidatorAnnouncement", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperlaneValidatorAnnounceValidatorAnnouncement)
				if err := _HyperlaneValidatorAnnounce.contract.UnpackLog(event, "ValidatorAnnouncement", log); err != nil {
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

// ParseValidatorAnnouncement is a log parse operation binding the contract event 0x78066d8adb677a1353d1fc8be28cf03e2a8de7157bbab979953587d78076c11e.
//
// Solidity: event ValidatorAnnouncement(address indexed validator, string storageLocation)
func (_HyperlaneValidatorAnnounce *HyperlaneValidatorAnnounceFilterer) ParseValidatorAnnouncement(log types.Log) (*HyperlaneValidatorAnnounceValidatorAnnouncement, error) {
	event := new(HyperlaneValidatorAnnounceValidatorAnnouncement)
	if err := _HyperlaneValidatorAnnounce.contract.UnpackLog(event, "ValidatorAnnouncement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
