// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testnft

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

// TestnftMetaData contains all meta data concerning the Testnft contract.
var TestnftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TOKEN_URI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintNFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260006006553480156200001657600080fd5b506040518060400160405280600e81526020017f5465737420417661696c204e46540000000000000000000000000000000000008152506040518060400160405280600381526020017f54414e0000000000000000000000000000000000000000000000000000000000815250816000908162000094919062000329565b508060019081620000a6919062000329565b50505062000410565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200013157607f821691505b602082108103620001475762000146620000e9565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620001b17fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000172565b620001bd868362000172565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200020a62000204620001fe84620001d5565b620001df565b620001d5565b9050919050565b6000819050919050565b6200022683620001e9565b6200023e620002358262000211565b8484546200017f565b825550505050565b600090565b6200025562000246565b620002628184846200021b565b505050565b5b818110156200028a576200027e6000826200024b565b60018101905062000268565b5050565b601f821115620002d957620002a3816200014d565b620002ae8462000162565b81016020851015620002be578190505b620002d6620002cd8562000162565b83018262000267565b50505b505050565b600082821c905092915050565b6000620002fe60001984600802620002de565b1980831691505092915050565b6000620003198383620002eb565b9150826002028217905092915050565b6200033482620000af565b67ffffffffffffffff81111562000350576200034f620000ba565b5b6200035c825462000118565b620003698282856200028e565b600060209050601f831160018114620003a157600084156200038c578287015190505b6200039885826200030b565b86555062000408565b601f198416620003b1866200014d565b60005b82811015620003db57848901518255600182019150602085019450602081019050620003b4565b86831015620003fb5784890151620003f7601f891682620002eb565b8355505b6001600288020188555050505b505050505050565b6124b680620004206000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80636e02007d11610097578063a22cb46511610066578063a22cb465146102af578063b88d4fde146102cb578063c87b56dd146102e7578063e985e9c51461031757610100565b80636e02007d1461022557806370a082311461024357806378ce90351461027357806395d89b411461029157610100565b806314f710fe116100d357806314f710fe1461019f57806323b872dd146101bd57806342842e0e146101d95780636352211e146101f557610100565b806301ffc9a71461010557806306fdde0314610135578063081812fc14610153578063095ea7b314610183575b600080fd5b61011f600480360381019061011a919061170b565b610347565b60405161012c9190611753565b60405180910390f35b61013d610429565b60405161014a91906117fe565b60405180910390f35b61016d60048036038101906101689190611856565b6104bb565b60405161017a91906118c4565b60405180910390f35b61019d6004803603810190610198919061190b565b610501565b005b6101a7610618565b6040516101b4919061195a565b60405180910390f35b6101d760048036038101906101d29190611975565b610646565b005b6101f360048036038101906101ee9190611975565b6106a6565b005b61020f600480360381019061020a9190611856565b6106c6565b60405161021c91906118c4565b60405180910390f35b61022d61074c565b60405161023a919061195a565b60405180910390f35b61025d600480360381019061025891906119c8565b610756565b60405161026a919061195a565b60405180910390f35b61027b61080d565b60405161028891906117fe565b60405180910390f35b610299610829565b6040516102a691906117fe565b60405180910390f35b6102c960048036038101906102c49190611a21565b6108bb565b005b6102e560048036038101906102e09190611b96565b6108d1565b005b61030160048036038101906102fc9190611856565b610933565b60405161030e91906117fe565b60405180910390f35b610331600480360381019061032c9190611c19565b610955565b60405161033e9190611753565b60405180910390f35b60007f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061041257507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806104225750610421826109e9565b5b9050919050565b60606000805461043890611c88565b80601f016020809104026020016040519081016040528092919081815260200182805461046490611c88565b80156104b15780601f10610486576101008083540402835291602001916104b1565b820191906000526020600020905b81548152906001019060200180831161049457829003601f168201915b5050505050905090565b60006104c682610a53565b6004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b600061050c826106c6565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361057c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161057390611d2b565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1661059b610a9e565b73ffffffffffffffffffffffffffffffffffffffff1614806105ca57506105c9816105c4610a9e565b610955565b5b610609576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060090611dbd565b60405180910390fd5b6106138383610aa6565b505050565b600061062633600654610b5f565b6006600081548092919061063990611e0c565b9190505550600654905090565b610657610651610a9e565b82610b7d565b610696576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161068d90611ec6565b60405180910390fd5b6106a1838383610c12565b505050565b6106c1838383604051806020016040528060008152506108d1565b505050565b6000806106d283610f0b565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610743576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073a90611f32565b60405180910390fd5b80915050919050565b6000600654905090565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036107c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107bd90611fc4565b60405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60405180608001604052806057815260200161242a6057913981565b60606001805461083890611c88565b80601f016020809104026020016040519081016040528092919081815260200182805461086490611c88565b80156108b15780601f10610886576101008083540402835291602001916108b1565b820191906000526020600020905b81548152906001019060200180831161089457829003601f168201915b5050505050905090565b6108cd6108c6610a9e565b8383610f48565b5050565b6108e26108dc610a9e565b83610b7d565b610921576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091890611ec6565b60405180910390fd5b61092d848484846110b4565b50505050565b606060405180608001604052806057815260200161242a605791399050919050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b610a5c81611110565b610a9b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9290611f32565b60405180910390fd5b50565b600033905090565b816004600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff16610b19836106c6565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b610b79828260405180602001604052806000815250611151565b5050565b600080610b89836106c6565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480610bcb5750610bca8185610955565b5b80610c0957508373ffffffffffffffffffffffffffffffffffffffff16610bf1846104bb565b73ffffffffffffffffffffffffffffffffffffffff16145b91505092915050565b8273ffffffffffffffffffffffffffffffffffffffff16610c32826106c6565b73ffffffffffffffffffffffffffffffffffffffff1614610c88576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7f90612056565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610cf7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cee906120e8565b60405180910390fd5b610d0483838360016111ac565b8273ffffffffffffffffffffffffffffffffffffffff16610d24826106c6565b73ffffffffffffffffffffffffffffffffffffffff1614610d7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d7190612056565b60405180910390fd5b6004600082815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055506001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4610f0683838360016112d2565b505050565b60006002600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610fb6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fad90612154565b60405180910390fd5b80600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516110a79190611753565b60405180910390a3505050565b6110bf848484610c12565b6110cb848484846112d8565b61110a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611101906121e6565b60405180910390fd5b50505050565b60008073ffffffffffffffffffffffffffffffffffffffff1661113283610f0b565b73ffffffffffffffffffffffffffffffffffffffff1614159050919050565b61115b838361145f565b61116860008484846112d8565b6111a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161119e906121e6565b60405180910390fd5b505050565b60018111156112cc57600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16146112405780600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546112389190612206565b925050819055505b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146112cb5780600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546112c3919061223a565b925050819055505b5b50505050565b50505050565b60006112f98473ffffffffffffffffffffffffffffffffffffffff1661167c565b15611452578373ffffffffffffffffffffffffffffffffffffffff1663150b7a02611322610a9e565b8786866040518563ffffffff1660e01b815260040161134494939291906122c3565b6020604051808303816000875af192505050801561138057506040513d601f19601f8201168201806040525081019061137d9190612324565b60015b611402573d80600081146113b0576040519150601f19603f3d011682016040523d82523d6000602084013e6113b5565b606091505b5060008151036113fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113f1906121e6565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614915050611457565b600190505b949350505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036114ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114c59061239d565b60405180910390fd5b6114d781611110565b15611517576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161150e90612409565b60405180910390fd5b6115256000838360016111ac565b61152e81611110565b1561156e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161156590612409565b60405180910390fd5b6001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a46116786000838360016112d2565b5050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6116e8816116b3565b81146116f357600080fd5b50565b600081359050611705816116df565b92915050565b600060208284031215611721576117206116a9565b5b600061172f848285016116f6565b91505092915050565b60008115159050919050565b61174d81611738565b82525050565b60006020820190506117686000830184611744565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156117a857808201518184015260208101905061178d565b60008484015250505050565b6000601f19601f8301169050919050565b60006117d08261176e565b6117da8185611779565b93506117ea81856020860161178a565b6117f3816117b4565b840191505092915050565b6000602082019050818103600083015261181881846117c5565b905092915050565b6000819050919050565b61183381611820565b811461183e57600080fd5b50565b6000813590506118508161182a565b92915050565b60006020828403121561186c5761186b6116a9565b5b600061187a84828501611841565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006118ae82611883565b9050919050565b6118be816118a3565b82525050565b60006020820190506118d960008301846118b5565b92915050565b6118e8816118a3565b81146118f357600080fd5b50565b600081359050611905816118df565b92915050565b60008060408385031215611922576119216116a9565b5b6000611930858286016118f6565b925050602061194185828601611841565b9150509250929050565b61195481611820565b82525050565b600060208201905061196f600083018461194b565b92915050565b60008060006060848603121561198e5761198d6116a9565b5b600061199c868287016118f6565b93505060206119ad868287016118f6565b92505060406119be86828701611841565b9150509250925092565b6000602082840312156119de576119dd6116a9565b5b60006119ec848285016118f6565b91505092915050565b6119fe81611738565b8114611a0957600080fd5b50565b600081359050611a1b816119f5565b92915050565b60008060408385031215611a3857611a376116a9565b5b6000611a46858286016118f6565b9250506020611a5785828601611a0c565b9150509250929050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611aa3826117b4565b810181811067ffffffffffffffff82111715611ac257611ac1611a6b565b5b80604052505050565b6000611ad561169f565b9050611ae18282611a9a565b919050565b600067ffffffffffffffff821115611b0157611b00611a6b565b5b611b0a826117b4565b9050602081019050919050565b82818337600083830152505050565b6000611b39611b3484611ae6565b611acb565b905082815260208101848484011115611b5557611b54611a66565b5b611b60848285611b17565b509392505050565b600082601f830112611b7d57611b7c611a61565b5b8135611b8d848260208601611b26565b91505092915050565b60008060008060808587031215611bb057611baf6116a9565b5b6000611bbe878288016118f6565b9450506020611bcf878288016118f6565b9350506040611be087828801611841565b925050606085013567ffffffffffffffff811115611c0157611c006116ae565b5b611c0d87828801611b68565b91505092959194509250565b60008060408385031215611c3057611c2f6116a9565b5b6000611c3e858286016118f6565b9250506020611c4f858286016118f6565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611ca057607f821691505b602082108103611cb357611cb2611c59565b5b50919050565b7f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b6000611d15602183611779565b9150611d2082611cb9565b604082019050919050565b60006020820190508181036000830152611d4481611d08565b9050919050565b7f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60008201527f6b656e206f776e6572206f7220617070726f76656420666f7220616c6c000000602082015250565b6000611da7603d83611779565b9150611db282611d4b565b604082019050919050565b60006020820190508181036000830152611dd681611d9a565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611e1782611820565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611e4957611e48611ddd565b5b600182019050919050565b7f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560008201527f72206f7220617070726f76656400000000000000000000000000000000000000602082015250565b6000611eb0602d83611779565b9150611ebb82611e54565b604082019050919050565b60006020820190508181036000830152611edf81611ea3565b9050919050565b7f4552433732313a20696e76616c696420746f6b656e2049440000000000000000600082015250565b6000611f1c601883611779565b9150611f2782611ee6565b602082019050919050565b60006020820190508181036000830152611f4b81611f0f565b9050919050565b7f4552433732313a2061646472657373207a65726f206973206e6f74206120766160008201527f6c6964206f776e65720000000000000000000000000000000000000000000000602082015250565b6000611fae602983611779565b9150611fb982611f52565b604082019050919050565b60006020820190508181036000830152611fdd81611fa1565b9050919050565b7f4552433732313a207472616e736665722066726f6d20696e636f72726563742060008201527f6f776e6572000000000000000000000000000000000000000000000000000000602082015250565b6000612040602583611779565b915061204b82611fe4565b604082019050919050565b6000602082019050818103600083015261206f81612033565b9050919050565b7f4552433732313a207472616e7366657220746f20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b60006120d2602483611779565b91506120dd82612076565b604082019050919050565b60006020820190508181036000830152612101816120c5565b9050919050565b7f4552433732313a20617070726f766520746f2063616c6c657200000000000000600082015250565b600061213e601983611779565b915061214982612108565b602082019050919050565b6000602082019050818103600083015261216d81612131565b9050919050565b7f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560008201527f63656976657220696d706c656d656e7465720000000000000000000000000000602082015250565b60006121d0603283611779565b91506121db82612174565b604082019050919050565b600060208201905081810360008301526121ff816121c3565b9050919050565b600061221182611820565b915061221c83611820565b925082820390508181111561223457612233611ddd565b5b92915050565b600061224582611820565b915061225083611820565b925082820190508082111561226857612267611ddd565b5b92915050565b600081519050919050565b600082825260208201905092915050565b60006122958261226e565b61229f8185612279565b93506122af81856020860161178a565b6122b8816117b4565b840191505092915050565b60006080820190506122d860008301876118b5565b6122e560208301866118b5565b6122f2604083018561194b565b8181036060830152612304818461228a565b905095945050505050565b60008151905061231e816116df565b92915050565b60006020828403121561233a576123396116a9565b5b60006123488482850161230f565b91505092915050565b7f4552433732313a206d696e7420746f20746865207a65726f2061646472657373600082015250565b6000612387602083611779565b915061239282612351565b602082019050919050565b600060208201905081810360008301526123b68161237a565b9050919050565b7f4552433732313a20746f6b656e20616c7265616479206d696e74656400000000600082015250565b60006123f3601c83611779565b91506123fe826123bd565b602082019050919050565b60006020820190508181036000830152612422816123e6565b905091905056fe697066733a2f2f62616679626569673337696f6972373673376d67356f6f6265746e636f6a636d3363336878617379643472766964346a71687934676b61686567342f3f66696c656e616d653d302d5055472e6a736f6ea264697066735822122074bca39b250b140cfa84ae51d4f3943d8098c9a0780448dea10be95e0717172764736f6c63430008130033",
}

// TestnftABI is the input ABI used to generate the binding from.
// Deprecated: Use TestnftMetaData.ABI instead.
var TestnftABI = TestnftMetaData.ABI

// TestnftBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestnftMetaData.Bin instead.
var TestnftBin = TestnftMetaData.Bin

// DeployTestnft deploys a new Ethereum contract, binding an instance of Testnft to it.
func DeployTestnft(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Testnft, error) {
	parsed, err := TestnftMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestnftBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Testnft{TestnftCaller: TestnftCaller{contract: contract}, TestnftTransactor: TestnftTransactor{contract: contract}, TestnftFilterer: TestnftFilterer{contract: contract}}, nil
}

// Testnft is an auto generated Go binding around an Ethereum contract.
type Testnft struct {
	TestnftCaller     // Read-only binding to the contract
	TestnftTransactor // Write-only binding to the contract
	TestnftFilterer   // Log filterer for contract events
}

// TestnftCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestnftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestnftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestnftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestnftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestnftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestnftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestnftSession struct {
	Contract     *Testnft          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestnftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestnftCallerSession struct {
	Contract *TestnftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TestnftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestnftTransactorSession struct {
	Contract     *TestnftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TestnftRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestnftRaw struct {
	Contract *Testnft // Generic contract binding to access the raw methods on
}

// TestnftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestnftCallerRaw struct {
	Contract *TestnftCaller // Generic read-only contract binding to access the raw methods on
}

// TestnftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestnftTransactorRaw struct {
	Contract *TestnftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestnft creates a new instance of Testnft, bound to a specific deployed contract.
func NewTestnft(address common.Address, backend bind.ContractBackend) (*Testnft, error) {
	contract, err := bindTestnft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Testnft{TestnftCaller: TestnftCaller{contract: contract}, TestnftTransactor: TestnftTransactor{contract: contract}, TestnftFilterer: TestnftFilterer{contract: contract}}, nil
}

// NewTestnftCaller creates a new read-only instance of Testnft, bound to a specific deployed contract.
func NewTestnftCaller(address common.Address, caller bind.ContractCaller) (*TestnftCaller, error) {
	contract, err := bindTestnft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestnftCaller{contract: contract}, nil
}

// NewTestnftTransactor creates a new write-only instance of Testnft, bound to a specific deployed contract.
func NewTestnftTransactor(address common.Address, transactor bind.ContractTransactor) (*TestnftTransactor, error) {
	contract, err := bindTestnft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestnftTransactor{contract: contract}, nil
}

// NewTestnftFilterer creates a new log filterer instance of Testnft, bound to a specific deployed contract.
func NewTestnftFilterer(address common.Address, filterer bind.ContractFilterer) (*TestnftFilterer, error) {
	contract, err := bindTestnft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestnftFilterer{contract: contract}, nil
}

// bindTestnft binds a generic wrapper to an already deployed contract.
func bindTestnft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestnftABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testnft *TestnftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Testnft.Contract.TestnftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testnft *TestnftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testnft.Contract.TestnftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testnft *TestnftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Testnft.Contract.TestnftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testnft *TestnftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Testnft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testnft *TestnftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testnft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testnft *TestnftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Testnft.Contract.contract.Transact(opts, method, params...)
}

// TOKENURI is a free data retrieval call binding the contract method 0x78ce9035.
//
// Solidity: function TOKEN_URI() view returns(string)
func (_Testnft *TestnftCaller) TOKENURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Testnft.contract.Call(opts, &out, "TOKEN_URI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TOKENURI is a free data retrieval call binding the contract method 0x78ce9035.
//
// Solidity: function TOKEN_URI() view returns(string)
func (_Testnft *TestnftSession) TOKENURI() (string, error) {
	return _Testnft.Contract.TOKENURI(&_Testnft.CallOpts)
}

// TOKENURI is a free data retrieval call binding the contract method 0x78ce9035.
//
// Solidity: function TOKEN_URI() view returns(string)
func (_Testnft *TestnftCallerSession) TOKENURI() (string, error) {
	return _Testnft.Contract.TOKENURI(&_Testnft.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Testnft *TestnftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Testnft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Testnft *TestnftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Testnft.Contract.BalanceOf(&_Testnft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Testnft *TestnftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Testnft.Contract.BalanceOf(&_Testnft.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Testnft *TestnftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Testnft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Testnft *TestnftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Testnft.Contract.GetApproved(&_Testnft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Testnft *TestnftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Testnft.Contract.GetApproved(&_Testnft.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Testnft *TestnftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Testnft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Testnft *TestnftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Testnft.Contract.IsApprovedForAll(&_Testnft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Testnft *TestnftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Testnft.Contract.IsApprovedForAll(&_Testnft.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Testnft *TestnftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Testnft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Testnft *TestnftSession) Name() (string, error) {
	return _Testnft.Contract.Name(&_Testnft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Testnft *TestnftCallerSession) Name() (string, error) {
	return _Testnft.Contract.Name(&_Testnft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Testnft *TestnftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Testnft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Testnft *TestnftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Testnft.Contract.OwnerOf(&_Testnft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Testnft *TestnftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Testnft.Contract.OwnerOf(&_Testnft.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Testnft *TestnftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Testnft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Testnft *TestnftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Testnft.Contract.SupportsInterface(&_Testnft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Testnft *TestnftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Testnft.Contract.SupportsInterface(&_Testnft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Testnft *TestnftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Testnft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Testnft *TestnftSession) Symbol() (string, error) {
	return _Testnft.Contract.Symbol(&_Testnft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Testnft *TestnftCallerSession) Symbol() (string, error) {
	return _Testnft.Contract.Symbol(&_Testnft.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Testnft *TestnftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Testnft.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Testnft *TestnftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Testnft.Contract.TokenURI(&_Testnft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Testnft *TestnftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Testnft.Contract.TokenURI(&_Testnft.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Testnft *TestnftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Testnft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Testnft *TestnftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Testnft.Contract.Approve(&_Testnft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Testnft *TestnftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Testnft.Contract.Approve(&_Testnft.TransactOpts, to, tokenId)
}

// GetTokenCounter is a paid mutator transaction binding the contract method 0x6e02007d.
//
// Solidity: function getTokenCounter() returns(uint256)
func (_Testnft *TestnftTransactor) GetTokenCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testnft.contract.Transact(opts, "getTokenCounter")
}

// GetTokenCounter is a paid mutator transaction binding the contract method 0x6e02007d.
//
// Solidity: function getTokenCounter() returns(uint256)
func (_Testnft *TestnftSession) GetTokenCounter() (*types.Transaction, error) {
	return _Testnft.Contract.GetTokenCounter(&_Testnft.TransactOpts)
}

// GetTokenCounter is a paid mutator transaction binding the contract method 0x6e02007d.
//
// Solidity: function getTokenCounter() returns(uint256)
func (_Testnft *TestnftTransactorSession) GetTokenCounter() (*types.Transaction, error) {
	return _Testnft.Contract.GetTokenCounter(&_Testnft.TransactOpts)
}

// MintNFT is a paid mutator transaction binding the contract method 0x14f710fe.
//
// Solidity: function mintNFT() returns(uint256)
func (_Testnft *TestnftTransactor) MintNFT(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testnft.contract.Transact(opts, "mintNFT")
}

// MintNFT is a paid mutator transaction binding the contract method 0x14f710fe.
//
// Solidity: function mintNFT() returns(uint256)
func (_Testnft *TestnftSession) MintNFT() (*types.Transaction, error) {
	return _Testnft.Contract.MintNFT(&_Testnft.TransactOpts)
}

// MintNFT is a paid mutator transaction binding the contract method 0x14f710fe.
//
// Solidity: function mintNFT() returns(uint256)
func (_Testnft *TestnftTransactorSession) MintNFT() (*types.Transaction, error) {
	return _Testnft.Contract.MintNFT(&_Testnft.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Testnft *TestnftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Testnft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Testnft *TestnftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Testnft.Contract.SafeTransferFrom(&_Testnft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Testnft *TestnftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Testnft.Contract.SafeTransferFrom(&_Testnft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Testnft *TestnftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Testnft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Testnft *TestnftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Testnft.Contract.SafeTransferFrom0(&_Testnft.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Testnft *TestnftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Testnft.Contract.SafeTransferFrom0(&_Testnft.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Testnft *TestnftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Testnft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Testnft *TestnftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Testnft.Contract.SetApprovalForAll(&_Testnft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Testnft *TestnftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Testnft.Contract.SetApprovalForAll(&_Testnft.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Testnft *TestnftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Testnft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Testnft *TestnftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Testnft.Contract.TransferFrom(&_Testnft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Testnft *TestnftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Testnft.Contract.TransferFrom(&_Testnft.TransactOpts, from, to, tokenId)
}

// TestnftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Testnft contract.
type TestnftApprovalIterator struct {
	Event *TestnftApproval // Event containing the contract specifics and raw log

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
func (it *TestnftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestnftApproval)
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
		it.Event = new(TestnftApproval)
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
func (it *TestnftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestnftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestnftApproval represents a Approval event raised by the Testnft contract.
type TestnftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Testnft *TestnftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*TestnftApprovalIterator, error) {

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

	logs, sub, err := _Testnft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TestnftApprovalIterator{contract: _Testnft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Testnft *TestnftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TestnftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Testnft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestnftApproval)
				if err := _Testnft.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Testnft *TestnftFilterer) ParseApproval(log types.Log) (*TestnftApproval, error) {
	event := new(TestnftApproval)
	if err := _Testnft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestnftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Testnft contract.
type TestnftApprovalForAllIterator struct {
	Event *TestnftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TestnftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestnftApprovalForAll)
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
		it.Event = new(TestnftApprovalForAll)
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
func (it *TestnftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestnftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestnftApprovalForAll represents a ApprovalForAll event raised by the Testnft contract.
type TestnftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Testnft *TestnftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TestnftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Testnft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TestnftApprovalForAllIterator{contract: _Testnft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Testnft *TestnftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TestnftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Testnft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestnftApprovalForAll)
				if err := _Testnft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Testnft *TestnftFilterer) ParseApprovalForAll(log types.Log) (*TestnftApprovalForAll, error) {
	event := new(TestnftApprovalForAll)
	if err := _Testnft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestnftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Testnft contract.
type TestnftTransferIterator struct {
	Event *TestnftTransfer // Event containing the contract specifics and raw log

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
func (it *TestnftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestnftTransfer)
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
		it.Event = new(TestnftTransfer)
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
func (it *TestnftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestnftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestnftTransfer represents a Transfer event raised by the Testnft contract.
type TestnftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Testnft *TestnftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*TestnftTransferIterator, error) {

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

	logs, sub, err := _Testnft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TestnftTransferIterator{contract: _Testnft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Testnft *TestnftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TestnftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Testnft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestnftTransfer)
				if err := _Testnft.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Testnft *TestnftFilterer) ParseTransfer(log types.Log) (*TestnftTransfer, error) {
	event := new(TestnftTransfer)
	if err := _Testnft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
