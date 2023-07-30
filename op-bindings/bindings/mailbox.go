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

// MailboxMetaData contains all meta data concerning the Mailbox contract.
var MailboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"DefaultIsmSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"DispatchId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"Process\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"ProcessId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_MESSAGE_BODY_BYTES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_status\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultIsm\",\"outputs\":[{\"internalType\":\"contractIInterchainSecurityModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"delivered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipientAddress\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"dispatch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_defaultIsm\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_metadata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"process\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"recipientIsm\",\"outputs\":[{\"internalType\":\"contractIInterchainSecurityModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_module\",\"type\":\"address\"}],\"name\":\"setDefaultIsm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405162002225380380620022258339810160408190526100319161003f565b63ffffffff1660805261006c565b60006020828403121561005157600080fd5b815163ffffffff8116811461006557600080fd5b9392505050565b60805161218f6200009660003960008181610244015281816106440152610d19015261218f6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80638da5cb5b116100d8578063ebf0c7171161008c578063fa31de0111610066578063fa31de0114610322578063fd54b22814610335578063ffa1ad741461033f57600080fd5b8063ebf0c717146102f4578063f2fde38b146102fc578063f794687a1461030f57600080fd5b8063b187bd26116100bd578063b187bd26146102a6578063e495f1d4146102be578063e70f48ac146102e157600080fd5b80638da5cb5b14610266578063907c0f921461028457600080fd5b80636e5f516e1161012f5780637c39d130116101145780637c39d130146102245780638456cb59146102375780638d3638f41461023f57600080fd5b80636e5f516e146101d7578063715018a61461021c57600080fd5b80633f4ba83a116101605780633f4ba83a146101b1578063485cc955146101bb578063522ae002146101ce57600080fd5b806306661abd1461017c5780630fb3844c1461019a575b600080fd5b60b8545b60405163ffffffff90911681526020015b60405180910390f35b6101a360655481565b604051908152602001610191565b6101b9610359565b005b6101b96101c9366004611aa7565b610394565b6101a361080081565b6097546101f79073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610191565b6101b9610547565b6101b9610232366004611b22565b61055b565b6101b9610a0e565b6101807f000000000000000000000000000000000000000000000000000000000000000081565b60335473ffffffffffffffffffffffffffffffffffffffff166101f7565b61028c610a49565b6040805192835263ffffffff909116602083015201610191565b6102ae610a71565b6040519015158152602001610191565b6102ae6102cc366004611b8e565b60b96020526000908152604090205460ff1681565b6101f76102ef366004611ba7565b610a84565b6101a3610b52565b6101b961030a366004611ba7565b610b5e565b6101b961031d366004611ba7565b610c15565b6101a3610330366004611bc4565b610c26565b60b8546101a39081565b610347600081565b60405160ff9091168152602001610191565b610361610de2565b610369610e63565b6040517fa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d1693390600090a1565b600054610100900460ff16158080156103b45750600054600160ff909116105b806103ce5750303b1580156103ce575060005460ff166001145b61045f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156104bd57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6104c5610ed6565b6104cd610f6d565b6104d683610b5e565b6104df8261100c565b801561054257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b61054f610de2565b61055960006110f9565b565b6001606554146105c7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f7265656e7472616e742063616c6c20286f7220706175736564290000000000006044820152606401610456565b600060658190556105d88383611170565b60ff1614610642576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600860248201527f2176657273696f6e0000000000000000000000000000000000000000000000006044820152606401610456565b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff166106738383611194565b63ffffffff16146106e0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f2164657374696e6174696f6e00000000000000000000000000000000000000006044820152606401610456565b600061072183838080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506111b792505050565b600081815260b9602052604090205490915060ff161561079d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f64656c69766572656400000000000000000000000000000000000000000000006044820152606401610456565b600081815260b96020526040812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556107e16102ef85856111c2565b6040517ff7e83aee00000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff82169063f7e83aee9061083c908990899089908990600401611c66565b6020604051808303816000875af115801561085b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061087f9190611c98565b6108e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f216d6f64756c65000000000000000000000000000000000000000000000000006044820152606401610456565b60006108f185856111db565b905060006108ff86866111eb565b9050600061090d87876111c2565b90508073ffffffffffffffffffffffffffffffffffffffff166356d5d47584846109378b8b611204565b6040518563ffffffff1660e01b81526004016109569493929190611cba565b600060405180830381600087803b15801561097057600080fd5b505af1158015610984573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff16828463ffffffff167f0d381c2a574ae8f04e213db7cfb4df8df712cdbd427d9868ffef380660ca657460405160405180910390a460405185907f1cae38cdd3d3919489272725a5ae62a4f48b2989b0dae843d3c279fee18073a990600090a25050600160655550505050505050565b610a16610de2565b610a1e611220565b6040517f9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e75290600090a1565b600080610a54610b52565b6001610a5f60b85490565b610a699190611d19565b915091509091565b6000610a7f60655460021490565b905090565b60008173ffffffffffffffffffffffffffffffffffffffff1663de523cf36040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610b0b575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252610b0891810190611d3e565b60015b15610b345773ffffffffffffffffffffffffffffffffffffffff811615610b325792915050565b505b505060975473ffffffffffffffffffffffffffffffffffffffff1690565b6000610a7f6098611293565b610b66610de2565b73ffffffffffffffffffffffffffffffffffffffff8116610c09576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610456565b610c12816110f9565b50565b610c1d610de2565b610c128161100c565b6000610c3460655460021490565b15610c9b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600660248201527f70617573656400000000000000000000000000000000000000000000000000006044820152606401610456565b610800821115610d07576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f6d736720746f6f206c6f6e6700000000000000000000000000000000000000006044820152606401610456565b6000610d426000610d1760b85490565b7f0000000000000000000000000000000000000000000000000000000000000000338a8a8a8a6112a6565b80516020820120909150610d576098826112e4565b858763ffffffff163373ffffffffffffffffffffffffffffffffffffffff167f769f711d20c679153d382254f59892613b58a97cc876b249134ac25c80f9c81485604051610da59190611d5b565b60405180910390a460405181907f788dbc1b7152732178210e7f4d9d010ef016f9eafbe66786bd7169f56e0c353a90600090a29695505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610559576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610456565b606554600214610ecf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f21706175736564000000000000000000000000000000000000000000000000006044820152606401610456565b6001606555565b600054610100900460ff16610ecf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610456565b600054610100900460ff16611004576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610456565b610559611423565b73ffffffffffffffffffffffffffffffffffffffff81163b61108a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f21636f6e747261637400000000000000000000000000000000000000000000006044820152606401610456565b609780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517fa76ad0adbf45318f8633aa0210f711273d50fbb6fef76ed95bbae97082c75daa90600090a250565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600061117f6001828486611dce565b61118891611df8565b60f81c90505b92915050565b60006111a4602d60298486611dce565b6111ad91611e3e565b60e01c9392505050565b805160209091012090565b60006111d46111d184846114c3565b90565b9392505050565b60006111a4600960058486611dce565b60006111fb602960098486611dce565b6111d491611e84565b36600061121483604d8187611dce565b915091505b9250929050565b60655460020361128c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600660248201527f70617573656400000000000000000000000000000000000000000000000000006044820152606401610456565b6002606555565b600061118e826112a16114d3565b611994565b606088888888888888886040516020016112c7989796959493929190611ec0565b604051602081830303815290604052905098975050505050505050565b60016112f260206002612076565b6112fc9190612082565b826020015410611368576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f6d65726b6c6520747265652066756c6c000000000000000000000000000000006044820152606401610456565b600182602001600082825461137d9190612099565b9091555050602082015460005b602081101561141a57816001166001036113b957828482602081106113b1576113b16120b1565b015550505050565b8381602081106113cb576113cb6120b1565b0154604080516020810192909252810184905260600160405160208183030381529060405280519060200120925060028261140691906120e0565b9150806114128161211b565b91505061138a565b50610542612153565b600054610100900460ff166114ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610456565b610559336110f9565b60006111fb604d602d8486611dce565b6114db611a66565b600081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560208201527fb4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d3060408201527f21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba8560608201527fe58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a1934460808201527f0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d60a08201527f887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a196860c08201527fffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f8360e08201527f9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af6101008201527fcefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e06101208201527ff9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a56101408201527ff8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf8926101608201527f3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c6101808201527fc1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb6101a08201527f5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc6101c08201527fda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d26101e08201527f2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f6102008201527fe1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a6102208201527f5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a06102408201527fb46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa06102608201527fc65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e26102808201527ff4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd96102a08201527f5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e3776102c08201527f4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee6526102e08201527fcdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef6103008201527f0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d6103208201527fb8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d06103408201527f838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e6103608201527f662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e6103808201527f388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea3226103a08201527f93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d7356103c08201527f8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a96103e082015290565b6020820154600090815b6020811015611a5e57600182821c1660008683602081106119c1576119c16120b1565b01549050816001036119fe576040805160208101839052908101869052606001604051602081830303815290604052805190602001209450611a49565b84868460208110611a1157611a116120b1565b6020020151604051602001611a30929190918252602082015260400190565b6040516020818303038152906040528051906020012094505b50508080611a569061211b565b91505061199e565b505092915050565b6040518061040001604052806020906020820280368337509192915050565b73ffffffffffffffffffffffffffffffffffffffff81168114610c1257600080fd5b60008060408385031215611aba57600080fd5b8235611ac581611a85565b91506020830135611ad581611a85565b809150509250929050565b60008083601f840112611af257600080fd5b50813567ffffffffffffffff811115611b0a57600080fd5b60208301915083602082850101111561121957600080fd5b60008060008060408587031215611b3857600080fd5b843567ffffffffffffffff80821115611b5057600080fd5b611b5c88838901611ae0565b90965094506020870135915080821115611b7557600080fd5b50611b8287828801611ae0565b95989497509550505050565b600060208284031215611ba057600080fd5b5035919050565b600060208284031215611bb957600080fd5b81356111d481611a85565b60008060008060608587031215611bda57600080fd5b843563ffffffff81168114611bee57600080fd5b935060208501359250604085013567ffffffffffffffff811115611c1157600080fd5b611b8287828801611ae0565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000611c7a604083018688611c1d565b8281036020840152611c8d818587611c1d565b979650505050505050565b600060208284031215611caa57600080fd5b815180151581146111d457600080fd5b63ffffffff85168152836020820152606060408201526000611ce0606083018486611c1d565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600063ffffffff83811690831681811015611d3657611d36611cea565b039392505050565b600060208284031215611d5057600080fd5b81516111d481611a85565b600060208083528351808285015260005b81811015611d8857858101830151858201604001528201611d6c565b81811115611d9a576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60008085851115611dde57600080fd5b83861115611deb57600080fd5b5050820193919092039150565b7fff000000000000000000000000000000000000000000000000000000000000008135818116916001851015611a5e5760019490940360031b84901b1690921692915050565b7fffffffff000000000000000000000000000000000000000000000000000000008135818116916004851015611a5e5760049490940360031b84901b1690921692915050565b8035602083101561118e577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff602084900360031b1b1692915050565b7fff000000000000000000000000000000000000000000000000000000000000008960f81b16815260007fffffffff00000000000000000000000000000000000000000000000000000000808a60e01b166001840152808960e01b166005840152876009840152808760e01b1660298401525084602d8301528284604d8401375060009101604d01908152979650505050505050565b600181815b80851115611faf57817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115611f9557611f95611cea565b80851615611fa257918102915b93841c9390800290611f5b565b509250929050565b600082611fc65750600161118e565b81611fd35750600061118e565b8160018114611fe95760028114611ff35761200f565b600191505061118e565b60ff84111561200457612004611cea565b50506001821b61118e565b5060208310610133831016604e8410600b8410161715612032575081810a61118e565b61203c8383611f56565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0482111561206e5761206e611cea565b029392505050565b60006111d48383611fb7565b60008282101561209457612094611cea565b500390565b600082198211156120ac576120ac611cea565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082612116577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361214c5761214c611cea565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fdfea164736f6c634300080f000a",
}

// MailboxABI is the input ABI used to generate the binding from.
// Deprecated: Use MailboxMetaData.ABI instead.
var MailboxABI = MailboxMetaData.ABI

// MailboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MailboxMetaData.Bin instead.
var MailboxBin = MailboxMetaData.Bin

// DeployMailbox deploys a new Ethereum contract, binding an instance of Mailbox to it.
func DeployMailbox(auth *bind.TransactOpts, backend bind.ContractBackend, _localDomain uint32) (common.Address, *types.Transaction, *Mailbox, error) {
	parsed, err := MailboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MailboxBin), backend, _localDomain)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mailbox{MailboxCaller: MailboxCaller{contract: contract}, MailboxTransactor: MailboxTransactor{contract: contract}, MailboxFilterer: MailboxFilterer{contract: contract}}, nil
}

// Mailbox is an auto generated Go binding around an Ethereum contract.
type Mailbox struct {
	MailboxCaller     // Read-only binding to the contract
	MailboxTransactor // Write-only binding to the contract
	MailboxFilterer   // Log filterer for contract events
}

// MailboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type MailboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MailboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MailboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MailboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MailboxSession struct {
	Contract     *Mailbox          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MailboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MailboxCallerSession struct {
	Contract *MailboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MailboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MailboxTransactorSession struct {
	Contract     *MailboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MailboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type MailboxRaw struct {
	Contract *Mailbox // Generic contract binding to access the raw methods on
}

// MailboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MailboxCallerRaw struct {
	Contract *MailboxCaller // Generic read-only contract binding to access the raw methods on
}

// MailboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MailboxTransactorRaw struct {
	Contract *MailboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMailbox creates a new instance of Mailbox, bound to a specific deployed contract.
func NewMailbox(address common.Address, backend bind.ContractBackend) (*Mailbox, error) {
	contract, err := bindMailbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mailbox{MailboxCaller: MailboxCaller{contract: contract}, MailboxTransactor: MailboxTransactor{contract: contract}, MailboxFilterer: MailboxFilterer{contract: contract}}, nil
}

// NewMailboxCaller creates a new read-only instance of Mailbox, bound to a specific deployed contract.
func NewMailboxCaller(address common.Address, caller bind.ContractCaller) (*MailboxCaller, error) {
	contract, err := bindMailbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MailboxCaller{contract: contract}, nil
}

// NewMailboxTransactor creates a new write-only instance of Mailbox, bound to a specific deployed contract.
func NewMailboxTransactor(address common.Address, transactor bind.ContractTransactor) (*MailboxTransactor, error) {
	contract, err := bindMailbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MailboxTransactor{contract: contract}, nil
}

// NewMailboxFilterer creates a new log filterer instance of Mailbox, bound to a specific deployed contract.
func NewMailboxFilterer(address common.Address, filterer bind.ContractFilterer) (*MailboxFilterer, error) {
	contract, err := bindMailbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MailboxFilterer{contract: contract}, nil
}

// bindMailbox binds a generic wrapper to an already deployed contract.
func bindMailbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MailboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mailbox *MailboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mailbox.Contract.MailboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mailbox *MailboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.Contract.MailboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mailbox *MailboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mailbox.Contract.MailboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mailbox *MailboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mailbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mailbox *MailboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mailbox *MailboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mailbox.Contract.contract.Transact(opts, method, params...)
}

// MAXMESSAGEBODYBYTES is a free data retrieval call binding the contract method 0x522ae002.
//
// Solidity: function MAX_MESSAGE_BODY_BYTES() view returns(uint256)
func (_Mailbox *MailboxCaller) MAXMESSAGEBODYBYTES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "MAX_MESSAGE_BODY_BYTES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMESSAGEBODYBYTES is a free data retrieval call binding the contract method 0x522ae002.
//
// Solidity: function MAX_MESSAGE_BODY_BYTES() view returns(uint256)
func (_Mailbox *MailboxSession) MAXMESSAGEBODYBYTES() (*big.Int, error) {
	return _Mailbox.Contract.MAXMESSAGEBODYBYTES(&_Mailbox.CallOpts)
}

// MAXMESSAGEBODYBYTES is a free data retrieval call binding the contract method 0x522ae002.
//
// Solidity: function MAX_MESSAGE_BODY_BYTES() view returns(uint256)
func (_Mailbox *MailboxCallerSession) MAXMESSAGEBODYBYTES() (*big.Int, error) {
	return _Mailbox.Contract.MAXMESSAGEBODYBYTES(&_Mailbox.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Mailbox *MailboxCaller) VERSION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Mailbox *MailboxSession) VERSION() (uint8, error) {
	return _Mailbox.Contract.VERSION(&_Mailbox.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Mailbox *MailboxCallerSession) VERSION() (uint8, error) {
	return _Mailbox.Contract.VERSION(&_Mailbox.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x0fb3844c.
//
// Solidity: function _status() view returns(uint256)
func (_Mailbox *MailboxCaller) Status(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "_status")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x0fb3844c.
//
// Solidity: function _status() view returns(uint256)
func (_Mailbox *MailboxSession) Status() (*big.Int, error) {
	return _Mailbox.Contract.Status(&_Mailbox.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x0fb3844c.
//
// Solidity: function _status() view returns(uint256)
func (_Mailbox *MailboxCallerSession) Status() (*big.Int, error) {
	return _Mailbox.Contract.Status(&_Mailbox.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint32)
func (_Mailbox *MailboxCaller) Count(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint32)
func (_Mailbox *MailboxSession) Count() (uint32, error) {
	return _Mailbox.Contract.Count(&_Mailbox.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint32)
func (_Mailbox *MailboxCallerSession) Count() (uint32, error) {
	return _Mailbox.Contract.Count(&_Mailbox.CallOpts)
}

// DefaultIsm is a free data retrieval call binding the contract method 0x6e5f516e.
//
// Solidity: function defaultIsm() view returns(address)
func (_Mailbox *MailboxCaller) DefaultIsm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "defaultIsm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultIsm is a free data retrieval call binding the contract method 0x6e5f516e.
//
// Solidity: function defaultIsm() view returns(address)
func (_Mailbox *MailboxSession) DefaultIsm() (common.Address, error) {
	return _Mailbox.Contract.DefaultIsm(&_Mailbox.CallOpts)
}

// DefaultIsm is a free data retrieval call binding the contract method 0x6e5f516e.
//
// Solidity: function defaultIsm() view returns(address)
func (_Mailbox *MailboxCallerSession) DefaultIsm() (common.Address, error) {
	return _Mailbox.Contract.DefaultIsm(&_Mailbox.CallOpts)
}

// Delivered is a free data retrieval call binding the contract method 0xe495f1d4.
//
// Solidity: function delivered(bytes32 ) view returns(bool)
func (_Mailbox *MailboxCaller) Delivered(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "delivered", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Delivered is a free data retrieval call binding the contract method 0xe495f1d4.
//
// Solidity: function delivered(bytes32 ) view returns(bool)
func (_Mailbox *MailboxSession) Delivered(arg0 [32]byte) (bool, error) {
	return _Mailbox.Contract.Delivered(&_Mailbox.CallOpts, arg0)
}

// Delivered is a free data retrieval call binding the contract method 0xe495f1d4.
//
// Solidity: function delivered(bytes32 ) view returns(bool)
func (_Mailbox *MailboxCallerSession) Delivered(arg0 [32]byte) (bool, error) {
	return _Mailbox.Contract.Delivered(&_Mailbox.CallOpts, arg0)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Mailbox *MailboxCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Mailbox *MailboxSession) IsPaused() (bool, error) {
	return _Mailbox.Contract.IsPaused(&_Mailbox.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Mailbox *MailboxCallerSession) IsPaused() (bool, error) {
	return _Mailbox.Contract.IsPaused(&_Mailbox.CallOpts)
}

// LatestCheckpoint is a free data retrieval call binding the contract method 0x907c0f92.
//
// Solidity: function latestCheckpoint() view returns(bytes32, uint32)
func (_Mailbox *MailboxCaller) LatestCheckpoint(opts *bind.CallOpts) ([32]byte, uint32, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "latestCheckpoint")

	if err != nil {
		return *new([32]byte), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// LatestCheckpoint is a free data retrieval call binding the contract method 0x907c0f92.
//
// Solidity: function latestCheckpoint() view returns(bytes32, uint32)
func (_Mailbox *MailboxSession) LatestCheckpoint() ([32]byte, uint32, error) {
	return _Mailbox.Contract.LatestCheckpoint(&_Mailbox.CallOpts)
}

// LatestCheckpoint is a free data retrieval call binding the contract method 0x907c0f92.
//
// Solidity: function latestCheckpoint() view returns(bytes32, uint32)
func (_Mailbox *MailboxCallerSession) LatestCheckpoint() ([32]byte, uint32, error) {
	return _Mailbox.Contract.LatestCheckpoint(&_Mailbox.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Mailbox *MailboxCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Mailbox *MailboxSession) LocalDomain() (uint32, error) {
	return _Mailbox.Contract.LocalDomain(&_Mailbox.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Mailbox *MailboxCallerSession) LocalDomain() (uint32, error) {
	return _Mailbox.Contract.LocalDomain(&_Mailbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mailbox *MailboxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mailbox *MailboxSession) Owner() (common.Address, error) {
	return _Mailbox.Contract.Owner(&_Mailbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mailbox *MailboxCallerSession) Owner() (common.Address, error) {
	return _Mailbox.Contract.Owner(&_Mailbox.CallOpts)
}

// RecipientIsm is a free data retrieval call binding the contract method 0xe70f48ac.
//
// Solidity: function recipientIsm(address _recipient) view returns(address)
func (_Mailbox *MailboxCaller) RecipientIsm(opts *bind.CallOpts, _recipient common.Address) (common.Address, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "recipientIsm", _recipient)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecipientIsm is a free data retrieval call binding the contract method 0xe70f48ac.
//
// Solidity: function recipientIsm(address _recipient) view returns(address)
func (_Mailbox *MailboxSession) RecipientIsm(_recipient common.Address) (common.Address, error) {
	return _Mailbox.Contract.RecipientIsm(&_Mailbox.CallOpts, _recipient)
}

// RecipientIsm is a free data retrieval call binding the contract method 0xe70f48ac.
//
// Solidity: function recipientIsm(address _recipient) view returns(address)
func (_Mailbox *MailboxCallerSession) RecipientIsm(_recipient common.Address) (common.Address, error) {
	return _Mailbox.Contract.RecipientIsm(&_Mailbox.CallOpts, _recipient)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Mailbox *MailboxCaller) Root(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Mailbox *MailboxSession) Root() ([32]byte, error) {
	return _Mailbox.Contract.Root(&_Mailbox.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() view returns(bytes32)
func (_Mailbox *MailboxCallerSession) Root() ([32]byte, error) {
	return _Mailbox.Contract.Root(&_Mailbox.CallOpts)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(uint256 count)
func (_Mailbox *MailboxCaller) Tree(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mailbox.contract.Call(opts, &out, "tree")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(uint256 count)
func (_Mailbox *MailboxSession) Tree() (*big.Int, error) {
	return _Mailbox.Contract.Tree(&_Mailbox.CallOpts)
}

// Tree is a free data retrieval call binding the contract method 0xfd54b228.
//
// Solidity: function tree() view returns(uint256 count)
func (_Mailbox *MailboxCallerSession) Tree() (*big.Int, error) {
	return _Mailbox.Contract.Tree(&_Mailbox.CallOpts)
}

// Dispatch is a paid mutator transaction binding the contract method 0xfa31de01.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, bytes _messageBody) returns(bytes32)
func (_Mailbox *MailboxTransactor) Dispatch(opts *bind.TransactOpts, _destinationDomain uint32, _recipientAddress [32]byte, _messageBody []byte) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "dispatch", _destinationDomain, _recipientAddress, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xfa31de01.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, bytes _messageBody) returns(bytes32)
func (_Mailbox *MailboxSession) Dispatch(_destinationDomain uint32, _recipientAddress [32]byte, _messageBody []byte) (*types.Transaction, error) {
	return _Mailbox.Contract.Dispatch(&_Mailbox.TransactOpts, _destinationDomain, _recipientAddress, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xfa31de01.
//
// Solidity: function dispatch(uint32 _destinationDomain, bytes32 _recipientAddress, bytes _messageBody) returns(bytes32)
func (_Mailbox *MailboxTransactorSession) Dispatch(_destinationDomain uint32, _recipientAddress [32]byte, _messageBody []byte) (*types.Transaction, error) {
	return _Mailbox.Contract.Dispatch(&_Mailbox.TransactOpts, _destinationDomain, _recipientAddress, _messageBody)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _defaultIsm) returns()
func (_Mailbox *MailboxTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _defaultIsm common.Address) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "initialize", _owner, _defaultIsm)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _defaultIsm) returns()
func (_Mailbox *MailboxSession) Initialize(_owner common.Address, _defaultIsm common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.Initialize(&_Mailbox.TransactOpts, _owner, _defaultIsm)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _defaultIsm) returns()
func (_Mailbox *MailboxTransactorSession) Initialize(_owner common.Address, _defaultIsm common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.Initialize(&_Mailbox.TransactOpts, _owner, _defaultIsm)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Mailbox *MailboxTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Mailbox *MailboxSession) Pause() (*types.Transaction, error) {
	return _Mailbox.Contract.Pause(&_Mailbox.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Mailbox *MailboxTransactorSession) Pause() (*types.Transaction, error) {
	return _Mailbox.Contract.Pause(&_Mailbox.TransactOpts)
}

// Process is a paid mutator transaction binding the contract method 0x7c39d130.
//
// Solidity: function process(bytes _metadata, bytes _message) returns()
func (_Mailbox *MailboxTransactor) Process(opts *bind.TransactOpts, _metadata []byte, _message []byte) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "process", _metadata, _message)
}

// Process is a paid mutator transaction binding the contract method 0x7c39d130.
//
// Solidity: function process(bytes _metadata, bytes _message) returns()
func (_Mailbox *MailboxSession) Process(_metadata []byte, _message []byte) (*types.Transaction, error) {
	return _Mailbox.Contract.Process(&_Mailbox.TransactOpts, _metadata, _message)
}

// Process is a paid mutator transaction binding the contract method 0x7c39d130.
//
// Solidity: function process(bytes _metadata, bytes _message) returns()
func (_Mailbox *MailboxTransactorSession) Process(_metadata []byte, _message []byte) (*types.Transaction, error) {
	return _Mailbox.Contract.Process(&_Mailbox.TransactOpts, _metadata, _message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mailbox *MailboxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mailbox *MailboxSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mailbox.Contract.RenounceOwnership(&_Mailbox.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mailbox *MailboxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mailbox.Contract.RenounceOwnership(&_Mailbox.TransactOpts)
}

// SetDefaultIsm is a paid mutator transaction binding the contract method 0xf794687a.
//
// Solidity: function setDefaultIsm(address _module) returns()
func (_Mailbox *MailboxTransactor) SetDefaultIsm(opts *bind.TransactOpts, _module common.Address) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "setDefaultIsm", _module)
}

// SetDefaultIsm is a paid mutator transaction binding the contract method 0xf794687a.
//
// Solidity: function setDefaultIsm(address _module) returns()
func (_Mailbox *MailboxSession) SetDefaultIsm(_module common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.SetDefaultIsm(&_Mailbox.TransactOpts, _module)
}

// SetDefaultIsm is a paid mutator transaction binding the contract method 0xf794687a.
//
// Solidity: function setDefaultIsm(address _module) returns()
func (_Mailbox *MailboxTransactorSession) SetDefaultIsm(_module common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.SetDefaultIsm(&_Mailbox.TransactOpts, _module)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mailbox *MailboxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mailbox *MailboxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.TransferOwnership(&_Mailbox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mailbox *MailboxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mailbox.Contract.TransferOwnership(&_Mailbox.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Mailbox *MailboxTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mailbox.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Mailbox *MailboxSession) Unpause() (*types.Transaction, error) {
	return _Mailbox.Contract.Unpause(&_Mailbox.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Mailbox *MailboxTransactorSession) Unpause() (*types.Transaction, error) {
	return _Mailbox.Contract.Unpause(&_Mailbox.TransactOpts)
}

// MailboxDefaultIsmSetIterator is returned from FilterDefaultIsmSet and is used to iterate over the raw logs and unpacked data for DefaultIsmSet events raised by the Mailbox contract.
type MailboxDefaultIsmSetIterator struct {
	Event *MailboxDefaultIsmSet // Event containing the contract specifics and raw log

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
func (it *MailboxDefaultIsmSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxDefaultIsmSet)
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
		it.Event = new(MailboxDefaultIsmSet)
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
func (it *MailboxDefaultIsmSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxDefaultIsmSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxDefaultIsmSet represents a DefaultIsmSet event raised by the Mailbox contract.
type MailboxDefaultIsmSet struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDefaultIsmSet is a free log retrieval operation binding the contract event 0xa76ad0adbf45318f8633aa0210f711273d50fbb6fef76ed95bbae97082c75daa.
//
// Solidity: event DefaultIsmSet(address indexed module)
func (_Mailbox *MailboxFilterer) FilterDefaultIsmSet(opts *bind.FilterOpts, module []common.Address) (*MailboxDefaultIsmSetIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "DefaultIsmSet", moduleRule)
	if err != nil {
		return nil, err
	}
	return &MailboxDefaultIsmSetIterator{contract: _Mailbox.contract, event: "DefaultIsmSet", logs: logs, sub: sub}, nil
}

// WatchDefaultIsmSet is a free log subscription operation binding the contract event 0xa76ad0adbf45318f8633aa0210f711273d50fbb6fef76ed95bbae97082c75daa.
//
// Solidity: event DefaultIsmSet(address indexed module)
func (_Mailbox *MailboxFilterer) WatchDefaultIsmSet(opts *bind.WatchOpts, sink chan<- *MailboxDefaultIsmSet, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "DefaultIsmSet", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxDefaultIsmSet)
				if err := _Mailbox.contract.UnpackLog(event, "DefaultIsmSet", log); err != nil {
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

// ParseDefaultIsmSet is a log parse operation binding the contract event 0xa76ad0adbf45318f8633aa0210f711273d50fbb6fef76ed95bbae97082c75daa.
//
// Solidity: event DefaultIsmSet(address indexed module)
func (_Mailbox *MailboxFilterer) ParseDefaultIsmSet(log types.Log) (*MailboxDefaultIsmSet, error) {
	event := new(MailboxDefaultIsmSet)
	if err := _Mailbox.contract.UnpackLog(event, "DefaultIsmSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxDispatchIterator is returned from FilterDispatch and is used to iterate over the raw logs and unpacked data for Dispatch events raised by the Mailbox contract.
type MailboxDispatchIterator struct {
	Event *MailboxDispatch // Event containing the contract specifics and raw log

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
func (it *MailboxDispatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxDispatch)
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
		it.Event = new(MailboxDispatch)
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
func (it *MailboxDispatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxDispatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxDispatch represents a Dispatch event raised by the Mailbox contract.
type MailboxDispatch struct {
	Sender      common.Address
	Destination uint32
	Recipient   [32]byte
	Message     []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDispatch is a free log retrieval operation binding the contract event 0x769f711d20c679153d382254f59892613b58a97cc876b249134ac25c80f9c814.
//
// Solidity: event Dispatch(address indexed sender, uint32 indexed destination, bytes32 indexed recipient, bytes message)
func (_Mailbox *MailboxFilterer) FilterDispatch(opts *bind.FilterOpts, sender []common.Address, destination []uint32, recipient [][32]byte) (*MailboxDispatchIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "Dispatch", senderRule, destinationRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &MailboxDispatchIterator{contract: _Mailbox.contract, event: "Dispatch", logs: logs, sub: sub}, nil
}

// WatchDispatch is a free log subscription operation binding the contract event 0x769f711d20c679153d382254f59892613b58a97cc876b249134ac25c80f9c814.
//
// Solidity: event Dispatch(address indexed sender, uint32 indexed destination, bytes32 indexed recipient, bytes message)
func (_Mailbox *MailboxFilterer) WatchDispatch(opts *bind.WatchOpts, sink chan<- *MailboxDispatch, sender []common.Address, destination []uint32, recipient [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "Dispatch", senderRule, destinationRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxDispatch)
				if err := _Mailbox.contract.UnpackLog(event, "Dispatch", log); err != nil {
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

// ParseDispatch is a log parse operation binding the contract event 0x769f711d20c679153d382254f59892613b58a97cc876b249134ac25c80f9c814.
//
// Solidity: event Dispatch(address indexed sender, uint32 indexed destination, bytes32 indexed recipient, bytes message)
func (_Mailbox *MailboxFilterer) ParseDispatch(log types.Log) (*MailboxDispatch, error) {
	event := new(MailboxDispatch)
	if err := _Mailbox.contract.UnpackLog(event, "Dispatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxDispatchIdIterator is returned from FilterDispatchId and is used to iterate over the raw logs and unpacked data for DispatchId events raised by the Mailbox contract.
type MailboxDispatchIdIterator struct {
	Event *MailboxDispatchId // Event containing the contract specifics and raw log

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
func (it *MailboxDispatchIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxDispatchId)
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
		it.Event = new(MailboxDispatchId)
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
func (it *MailboxDispatchIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxDispatchIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxDispatchId represents a DispatchId event raised by the Mailbox contract.
type MailboxDispatchId struct {
	MessageId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDispatchId is a free log retrieval operation binding the contract event 0x788dbc1b7152732178210e7f4d9d010ef016f9eafbe66786bd7169f56e0c353a.
//
// Solidity: event DispatchId(bytes32 indexed messageId)
func (_Mailbox *MailboxFilterer) FilterDispatchId(opts *bind.FilterOpts, messageId [][32]byte) (*MailboxDispatchIdIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "DispatchId", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &MailboxDispatchIdIterator{contract: _Mailbox.contract, event: "DispatchId", logs: logs, sub: sub}, nil
}

// WatchDispatchId is a free log subscription operation binding the contract event 0x788dbc1b7152732178210e7f4d9d010ef016f9eafbe66786bd7169f56e0c353a.
//
// Solidity: event DispatchId(bytes32 indexed messageId)
func (_Mailbox *MailboxFilterer) WatchDispatchId(opts *bind.WatchOpts, sink chan<- *MailboxDispatchId, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "DispatchId", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxDispatchId)
				if err := _Mailbox.contract.UnpackLog(event, "DispatchId", log); err != nil {
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

// ParseDispatchId is a log parse operation binding the contract event 0x788dbc1b7152732178210e7f4d9d010ef016f9eafbe66786bd7169f56e0c353a.
//
// Solidity: event DispatchId(bytes32 indexed messageId)
func (_Mailbox *MailboxFilterer) ParseDispatchId(log types.Log) (*MailboxDispatchId, error) {
	event := new(MailboxDispatchId)
	if err := _Mailbox.contract.UnpackLog(event, "DispatchId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Mailbox contract.
type MailboxInitializedIterator struct {
	Event *MailboxInitialized // Event containing the contract specifics and raw log

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
func (it *MailboxInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxInitialized)
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
		it.Event = new(MailboxInitialized)
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
func (it *MailboxInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxInitialized represents a Initialized event raised by the Mailbox contract.
type MailboxInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mailbox *MailboxFilterer) FilterInitialized(opts *bind.FilterOpts) (*MailboxInitializedIterator, error) {

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MailboxInitializedIterator{contract: _Mailbox.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mailbox *MailboxFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MailboxInitialized) (event.Subscription, error) {

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxInitialized)
				if err := _Mailbox.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mailbox *MailboxFilterer) ParseInitialized(log types.Log) (*MailboxInitialized, error) {
	event := new(MailboxInitialized)
	if err := _Mailbox.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mailbox contract.
type MailboxOwnershipTransferredIterator struct {
	Event *MailboxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MailboxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxOwnershipTransferred)
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
		it.Event = new(MailboxOwnershipTransferred)
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
func (it *MailboxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxOwnershipTransferred represents a OwnershipTransferred event raised by the Mailbox contract.
type MailboxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mailbox *MailboxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MailboxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MailboxOwnershipTransferredIterator{contract: _Mailbox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mailbox *MailboxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MailboxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxOwnershipTransferred)
				if err := _Mailbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Mailbox *MailboxFilterer) ParseOwnershipTransferred(log types.Log) (*MailboxOwnershipTransferred, error) {
	event := new(MailboxOwnershipTransferred)
	if err := _Mailbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Mailbox contract.
type MailboxPausedIterator struct {
	Event *MailboxPaused // Event containing the contract specifics and raw log

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
func (it *MailboxPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxPaused)
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
		it.Event = new(MailboxPaused)
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
func (it *MailboxPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxPaused represents a Paused event raised by the Mailbox contract.
type MailboxPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Mailbox *MailboxFilterer) FilterPaused(opts *bind.FilterOpts) (*MailboxPausedIterator, error) {

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MailboxPausedIterator{contract: _Mailbox.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Mailbox *MailboxFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MailboxPaused) (event.Subscription, error) {

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxPaused)
				if err := _Mailbox.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Mailbox *MailboxFilterer) ParsePaused(log types.Log) (*MailboxPaused, error) {
	event := new(MailboxPaused)
	if err := _Mailbox.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxProcessIterator is returned from FilterProcess and is used to iterate over the raw logs and unpacked data for Process events raised by the Mailbox contract.
type MailboxProcessIterator struct {
	Event *MailboxProcess // Event containing the contract specifics and raw log

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
func (it *MailboxProcessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxProcess)
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
		it.Event = new(MailboxProcess)
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
func (it *MailboxProcessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxProcessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxProcess represents a Process event raised by the Mailbox contract.
type MailboxProcess struct {
	Origin    uint32
	Sender    [32]byte
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcess is a free log retrieval operation binding the contract event 0x0d381c2a574ae8f04e213db7cfb4df8df712cdbd427d9868ffef380660ca6574.
//
// Solidity: event Process(uint32 indexed origin, bytes32 indexed sender, address indexed recipient)
func (_Mailbox *MailboxFilterer) FilterProcess(opts *bind.FilterOpts, origin []uint32, sender [][32]byte, recipient []common.Address) (*MailboxProcessIterator, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "Process", originRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &MailboxProcessIterator{contract: _Mailbox.contract, event: "Process", logs: logs, sub: sub}, nil
}

// WatchProcess is a free log subscription operation binding the contract event 0x0d381c2a574ae8f04e213db7cfb4df8df712cdbd427d9868ffef380660ca6574.
//
// Solidity: event Process(uint32 indexed origin, bytes32 indexed sender, address indexed recipient)
func (_Mailbox *MailboxFilterer) WatchProcess(opts *bind.WatchOpts, sink chan<- *MailboxProcess, origin []uint32, sender [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "Process", originRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxProcess)
				if err := _Mailbox.contract.UnpackLog(event, "Process", log); err != nil {
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

// ParseProcess is a log parse operation binding the contract event 0x0d381c2a574ae8f04e213db7cfb4df8df712cdbd427d9868ffef380660ca6574.
//
// Solidity: event Process(uint32 indexed origin, bytes32 indexed sender, address indexed recipient)
func (_Mailbox *MailboxFilterer) ParseProcess(log types.Log) (*MailboxProcess, error) {
	event := new(MailboxProcess)
	if err := _Mailbox.contract.UnpackLog(event, "Process", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxProcessIdIterator is returned from FilterProcessId and is used to iterate over the raw logs and unpacked data for ProcessId events raised by the Mailbox contract.
type MailboxProcessIdIterator struct {
	Event *MailboxProcessId // Event containing the contract specifics and raw log

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
func (it *MailboxProcessIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxProcessId)
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
		it.Event = new(MailboxProcessId)
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
func (it *MailboxProcessIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxProcessIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxProcessId represents a ProcessId event raised by the Mailbox contract.
type MailboxProcessId struct {
	MessageId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProcessId is a free log retrieval operation binding the contract event 0x1cae38cdd3d3919489272725a5ae62a4f48b2989b0dae843d3c279fee18073a9.
//
// Solidity: event ProcessId(bytes32 indexed messageId)
func (_Mailbox *MailboxFilterer) FilterProcessId(opts *bind.FilterOpts, messageId [][32]byte) (*MailboxProcessIdIterator, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "ProcessId", messageIdRule)
	if err != nil {
		return nil, err
	}
	return &MailboxProcessIdIterator{contract: _Mailbox.contract, event: "ProcessId", logs: logs, sub: sub}, nil
}

// WatchProcessId is a free log subscription operation binding the contract event 0x1cae38cdd3d3919489272725a5ae62a4f48b2989b0dae843d3c279fee18073a9.
//
// Solidity: event ProcessId(bytes32 indexed messageId)
func (_Mailbox *MailboxFilterer) WatchProcessId(opts *bind.WatchOpts, sink chan<- *MailboxProcessId, messageId [][32]byte) (event.Subscription, error) {

	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "ProcessId", messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxProcessId)
				if err := _Mailbox.contract.UnpackLog(event, "ProcessId", log); err != nil {
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

// ParseProcessId is a log parse operation binding the contract event 0x1cae38cdd3d3919489272725a5ae62a4f48b2989b0dae843d3c279fee18073a9.
//
// Solidity: event ProcessId(bytes32 indexed messageId)
func (_Mailbox *MailboxFilterer) ParseProcessId(log types.Log) (*MailboxProcessId, error) {
	event := new(MailboxProcessId)
	if err := _Mailbox.contract.UnpackLog(event, "ProcessId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MailboxUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Mailbox contract.
type MailboxUnpausedIterator struct {
	Event *MailboxUnpaused // Event containing the contract specifics and raw log

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
func (it *MailboxUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MailboxUnpaused)
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
		it.Event = new(MailboxUnpaused)
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
func (it *MailboxUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MailboxUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MailboxUnpaused represents a Unpaused event raised by the Mailbox contract.
type MailboxUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Mailbox *MailboxFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MailboxUnpausedIterator, error) {

	logs, sub, err := _Mailbox.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MailboxUnpausedIterator{contract: _Mailbox.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Mailbox *MailboxFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MailboxUnpaused) (event.Subscription, error) {

	logs, sub, err := _Mailbox.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MailboxUnpaused)
				if err := _Mailbox.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Mailbox *MailboxFilterer) ParseUnpaused(log types.Log) (*MailboxUnpaused, error) {
	event := new(MailboxUnpaused)
	if err := _Mailbox.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
