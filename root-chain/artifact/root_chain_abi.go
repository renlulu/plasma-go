// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package root_chain

import (
	"math/big"
	"strings"

	ethereum "github.com/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// RootChainABI is the input ABI used to generate the binding from.
const RootChainABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_utxoPos\",\"type\":\"uint256\"},{\"name\":\"_txBytes\",\"type\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes\"},{\"name\":\"_sigs\",\"type\":\"bytes\"}],\"name\":\"startExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cUtxoPos\",\"type\":\"uint256\"},{\"name\":\"_eUtxoIndex\",\"type\":\"uint256\"},{\"name\":\"_txBytes\",\"type\":\"bytes\"},{\"name\":\"_proof\",\"type\":\"bytes\"},{\"name\":\"_sigs\",\"type\":\"bytes\"},{\"name\":\"_confirmationSig\",\"type\":\"bytes\"}],\"name\":\"challengeExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exits\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"startFeeExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentChildBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"finalizeExits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getChildChain\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getNextExit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CHILD_BLOCK_INTERVAL\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentDepositBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"submitBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentFeeExit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"exitsQueues\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_utxoPos\",\"type\":\"uint256\"}],\"name\":\"getExit\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_depositPos\",\"type\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"startDepositExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"childChain\",\"outputs\":[{\"name\":\"root\",\"type\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"depositBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"exitor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"utxoPos\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExitStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BlockSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"}]"

// RootChainBin is the compiled bytecode used for deploying new contracts.
const RootChainBin = `608060405234801561001057600080fd5b5060008054600160a060020a031916331790556103e86001908155600281905560035561003b6100a5565b604051809103906000f080158015610057573d6000803e3d6000fd5b506000805260066020527f54cdd369e4e8a8515e52ca72ec816c2101831ad1f18bf44102ed171459c9b4f88054600160a060020a031916600160a060020a03929092169190911790556100b5565b60405161067f80611a4783390190565b611983806100c46000396000f3006080604052600436106101065763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631c91a6b9811461010b57806332773ba3146101e7578063342de1791461030157806340b8d53a14610343578063570ca735146103675780637a95f1e8146103985780637d7b2051146103bf57806385444de3146103e05780638bfe0aec14610411578063a831fa0714610432578063a98c7f2c14610447578063baa476941461045c578063bcd5926114610474578063cfe7d85514610489578063d0e30db01461049e578063d11f045c146104a6578063e60f1ff1146104c7578063f82f7a4b146104df578063f95643b114610506575b600080fd5b34801561011757600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101e595833595369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975061051e9650505050505050565b005b3480156101f357600080fd5b50604080516020600460443581810135601f81018490048402850184019095528484526101e594823594602480359536959460649492019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975061073e9650505050505050565b34801561030d57600080fd5b506103196004356108ec565b60408051600160a060020a0394851681529290931660208301528183015290519081900360600190f35b34801561034f57600080fd5b506101e5600160a060020a036004351660243561091a565b34801561037357600080fd5b5061037c61095e565b60408051600160a060020a039092168252519081900360200190f35b3480156103a457600080fd5b506103ad61096d565b60408051918252519081900360200190f35b3480156103cb57600080fd5b506101e5600160a060020a0360043516610973565b3480156103ec57600080fd5b506103f8600435610bd2565b6040805192835260208301919091528051918290030190f35b34801561041d57600080fd5b506103f8600160a060020a0360043516610bec565b34801561043e57600080fd5b506103ad610cbf565b34801561045357600080fd5b506103ad610cc5565b34801561046857600080fd5b506101e5600435610ccb565b34801561048057600080fd5b506103ad610d69565b34801561049557600080fd5b506103ad610d99565b6101e5610d9f565b3480156104b257600080fd5b5061037c600160a060020a0360043516610e77565b3480156104d357600080fd5b50610319600435610e92565b3480156104eb57600080fd5b506101e5600435600160a060020a0360243516604435610ec1565b34801561051257600080fd5b506103f8600435610f5d565b600080600061052b6118f9565b600080633b9aca008a049550612710633b9aca008b0681151561054a57fe5b0494506127108502633b9aca0087028b0303935061056e898563ffffffff610f7616565b8051909350600160a060020a0316331461058757600080fd5b60046000878152602001908152602001600020600001549150886040518082805190602001908083835b602083106105d05780518252601f1990920191602091820191016105b1565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902061060a886000608261104d565b6040518281528151602080830191908401908083835b6020831061063f5780518252601f199092019160209182019101610620565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051809103902090506106db896040518082805190602001908083835b602083106106a25780518252601f199092019160209182019101610683565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390208385606001518a6110ce565b15156106e657600080fd5b6106f88186848b63ffffffff6111d916565b151561070357600080fd5b6107328a846000015185602001518660400151600460008c815260200190815260200160002060010154611265565b50505050505050505050565b60008080808080806107568b8d63ffffffff61147f16565b9650612710633b9aca008e0604955060046000633b9aca008f0481526020019081526020016000206000015494508a6040518082805190602001908083835b602083106107b45780518252601f199092019160209182019101610795565b51815160001960209485036101000a01908116901991909116179052604080519490920184900384208085528482018c905282519485900390920184208285528f51929a5098508995508e945083810192508401908083835b6020831061082c5780518252601f19909201916020918201910161080d565b51815160209384036101000a6000190180199092169116179052604080519290940182900390912060008e815260059092529290205491965050600160a060020a0316935061088192508591508a90506114de565b600160a060020a0382811691161461089857600080fd5b6108aa8287878d63ffffffff6111d916565b15156108b557600080fd5b5050506000938452505060056020525060409020805473ffffffffffffffffffffffffffffffffffffffff19169055505050505050565b600560205260009081526040902080546001820154600290920154600160a060020a03918216929091169083565b600054600160a060020a0316331461093157600080fd5b61094360035433848442600101611265565b60035461095790600163ffffffff6115b316565b6003555050565b600054600160a060020a031681565b60015481565b60008061097e611920565b600061098985610bec565b600082815260056020908152604080832081516060810183528154600160a060020a0390811682526001830154811682860152600290920154818401528b8216855260069093529220549397509195509093501690505b42831015610bcb5760008481526005602090815260409182902082516060810184528154600160a060020a039081168252600183015481169382019390935260029091015492810192909252909250851615610a3b57600080fd5b81516040808401519051600160a060020a039092169181156108fc0291906000818181858888f19350505050158015610a78573d6000803e3d6000fd5b5080600160a060020a031663b07576ac6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015610ad057600080fd5b505af1158015610ae4573d6000803e3d6000fd5b505050506040513d6020811015610afa57600080fd5b50506000848152600560209081526040808320805473ffffffffffffffffffffffffffffffffffffffff1916905580517fbda1504b0000000000000000000000000000000000000000000000000000000081529051600160a060020a0385169263bda1504b926004808201939182900301818787803b158015610b7c57600080fd5b505af1158015610b90573d6000803e3d6000fd5b505050506040513d6020811015610ba657600080fd5b50511115610bc157610bb785610bec565b9094509250610bc6565b610bcb565b6109e0565b5050505050565b600090815260046020526040902080546001909101549091565b600160a060020a0380821660009081526006602090815260408083205481517fd6362e970000000000000000000000000000000000000000000000000000000081529151939485948594859485949093169263d6362e97926004808301939282900301818787803b158015610c6057600080fd5b505af1158015610c74573d6000803e3d6000fd5b505050506040513d6020811015610c8a57600080fd5b50516fffffffffffffffffffffffffffffffff8116977001000000000000000000000000000000009091049650945050505050565b6103e881565b60025481565b600054600160a060020a03163314610ce257600080fd5b604080518082018252828152426020808301918252600180546000908152600490925293902091518255519082015554610d24906103e863ffffffff6115b316565b60019081556002556040805182815242602082015281517f3be20e672590051e92360aa9dd9247e52541ad35d7a0297ff9e0f94e993295e9929181900390910190a150565b6000610d94600254610d886103e86001546115c990919063ffffffff16565b9063ffffffff6115b316565b905090565b60035481565b6000806103e8600254101515610db457600080fd5b604080516c01000000000000000000000000330281526000601482015234602882015290519081900360480190209150610dec610d69565b6040805180820182528481524260208083019182526000858152600490915292909220905181559051600191820155600254919250610e31919063ffffffff6115b316565b60025560408051600081523460208201528151839233927fd2f8022f659fd9c8c558f30c00fd5ee7038f7cb56da45095c3e0e7d48b3e0c4b929081900390910190a35050565b600660205260009081526040902054600160a060020a031681565b600090815260056020526040902080546001820154600290920154600160a060020a0391821693919092169190565b633b9aca0083046000806103e883061515610edb57600080fd5b5050600081815260046020526040908190205481516c010000000000000000000000003381028252600160a060020a03871602601482015260288101859052915191829003604801909120808214610f3257600080fd5b610f55863387876004600089815260200190815260200160002060010154611265565b505050505050565b6004602052600090815260409020805460019091015482565b610f7e6118f9565b6060610f91610f8c856115db565b61162c565b9050608060405190810160405280610fc58386600202600701815181101515610fb657fe5b906020019060200201516116ff565b600160a060020a03168152602001610fe5836006815181101515610fb657fe5b600160a060020a03168152602001611019838660020260080181518110151561100a57fe5b9060200190602002015161173b565b815260200161103083600381518110151561100a57fe5b61104284600081518110151561100a57fe5b029052949350505050565b60608082840185511015151561106257600080fd5b8215801561107b576040519150602082016040526110c5565b6040519150601f8416801560200281840101858101878315602002848b0101015b818310156110b457805183526020928301920161109c565b5050858452601f01601f1916604052505b50949350505050565b6000606080606060008060006060604189518115156110e957fe5b061580156110fa5750610104895111155b151561110557600080fd5b611112896000604161104d565b96506111208960418061104d565b955061112f896082604161104d565b604080518e8152602081018e9052815190819003909101902090955093506001925082915061115e84866114de565b600160a060020a03166111718d896114de565b600160a060020a031614925060008a11156111bf576111938960c3604161104d565b905061119f84826114de565b600160a060020a03166111b28d886114de565b600160a060020a03161491505b8280156111c95750815b9c9b505050505050505050505050565b60008060008084516102001415156111f057600080fd5b5086905060205b61020081116112575784810151925060028706151561122e5760408051928352602083018490528051928390030190912090611249565b60408051848152602081019390935280519283900301909120905b6002870496506020016111f7565b509390931495945050505050565b600160a060020a0383811660009081526006602052604081205490918291829116151561129157600080fd5b6112a58462127500014262093a8001611771565b9250700100000000000000000000000000000000830288179150600085116112cc57600080fd5b600088815260056020526040902060020154156112e857600080fd5b50600160a060020a038086166000908152600660205260408082205481517f90b5561d00000000000000000000000000000000000000000000000000000000815260048101869052915193169283926390b5561d926024808201939182900301818387803b15801561135957600080fd5b505af115801561136d573d6000803e3d6000fd5b5050505060606040519081016040528088600160a060020a0316815260200187600160a060020a0316815260200186815250600560008a815260200190815260200160002060008201518160000160006101000a815481600160a060020a030219169083600160a060020a0316021790555060208201518160010160006101000a815481600160a060020a030219169083600160a060020a03160217905550604082015181600201559050508733600160a060020a03167ff0537507c0bb9b823a1e4f5522ed9f25512618d4d169c9eb70a5ee2e9fb8bd5788886040518083600160a060020a0316600160a060020a031681526020018281526020019250505060405180910390a35050505050505050565b600060606000611491610f8c866115db565b91508360030290506114ad828260020181518110151561100a57fe5b6114c1838360010181518110151561100a57fe5b83516114d39085908590811061100a57fe5b010195945050505050565b600080600080845160411415156114f857600093506115aa565b50505060208201516040830151606084015160001a601b60ff8216101561151d57601b015b8060ff16601b1415801561153557508060ff16601c14155b1561154357600093506115aa565b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af115801561159d573d6000803e3d6000fd5b5050506020604051035193505b50505092915050565b6000828201838110156115c257fe5b9392505050565b6000828211156115d557fe5b50900390565b6115e3611940565b600082516000141561160a5760408051808201909152600080825260208201529150611626565b5050604080518082019091528151815260208281019082018190525b50919050565b606060008060008061163d86611788565b151561164857600080fd5b611651866117b5565b93508360405190808252806020026020018201604052801561168d57816020015b61167a611940565b8152602001906001900390816116725790505b50945061169d8660200151611805565b8660200151019250600090505b838110156116f6576116bb83611868565b915060408051908101604052808381526020018481525085828151811015156116e057fe5b60209081029091010152918101916001016116aa565b50505050919050565b80516000908190819060151461171457600080fd5b505050602081015160010180516c01000000000000000000000000900490815b5050919050565b60008060008060006117508660200151611805565b86516020978801518201519190039096036101000a90950495945050505050565b60008183101561178157816115c2565b5090919050565b6020810151805160009190821a9060c060ff831610156117ab5760009250611734565b5060019392505050565b600080600080600092506117cc8560200151611805565b6020860151865191810193500190505b808210156117fc576117ed82611868565b600190930192909101906117dc565b50909392505050565b8051600090811a608081101561181e5760009150611626565b60b8811080611839575060c08110801590611839575060f881105b156118475760019150611626565b60c081101561185c5760b51981019150611626565b60f51981019150611626565b8051600090811a60808110156118815760019150611626565b60b881101561189657607e1981019150611626565b60c08110156118c35760b78103600184019350806020036101000a84510460018201810193505050611626565b60f88110156118d85760be1981019150611626565b60019290920151602083900360f7016101000a900490910160f51901919050565b60408051608081018252600080825260208201819052918101829052606081019190915290565b604080516060810182526000808252602082018190529181019190915290565b6040805180820190915260008082526020820152905600a165627a7a723058201669ff99443c803b68235948ce658514eba2ee27bf3f117272f6efe4337377260029608060405234801561001057600080fd5b5060008054600160a060020a03191633178155604080516020810190915290815261003e9060019081610049565b5060006002556100b6565b828054828255906000526020600020908101928215610089579160200282015b82811115610089578251829060ff16905591602001919060010190610069565b50610095929150610099565b5090565b6100b391905b80821115610095576000815560010161009f565b90565b6105ba806100c56000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632dcdcd0c811461007157806390b5561d1461009b578063b07576ac146100b5578063bda1504b146100ca578063d6362e97146100df575b600080fd5b34801561007d57600080fd5b506100896004356100f4565b60408051918252519081900360200190f35b3480156100a757600080fd5b506100b36004356101c3565b005b3480156100c157600080fd5b5061008961023d565b3480156100d657600080fd5b5061008961031c565b3480156100eb57600080fd5b50610089610322565b600060025461011e600161011260028661034490919063ffffffff16565b9063ffffffff61037a16565b111561013c5761013582600263ffffffff61034416565b90506101be565b60016101538161011285600263ffffffff61034416565b8154811061015d57fe5b600091825260209091200154600161017c84600263ffffffff61034416565b8154811061018657fe5b906000526020600020015410156101a85761013582600263ffffffff61034416565b610135600161011284600263ffffffff61034416565b919050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146101e757600080fd5b60018054808201825560008290527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60182905560025461022c9163ffffffff61037a16565b600281905561023a90610389565b50565b60008054819073ffffffffffffffffffffffffffffffffffffffff16331461026457600080fd5b600180548190811061027257fe5b90600052602060002001549050600160025481548110151561029057fe5b90600052602060002001546001808154811015156102aa57fe5b906000526020600020018190555060016002548154811015156102c957fe5b60009182526020822001556002546102e890600163ffffffff61045d16565b6002556102f5600161046f565b600180546103089163ffffffff61045d16565b610313600182610555565b508091505b5090565b60025481565b600060018081548110151561033357fe5b906000526020600020015490505b90565b6000808315156103575760009150610373565b5082820282848281151561036757fe5b041461036f57fe5b8091505b5092915050565b60008282018381101561036f57fe5b6001805482916000918390811061039c57fe5b906000526020600020015490505b60016103bd84600263ffffffff61053e16565b815481106103c757fe5b90600052602060002001548110156104345760016103ec84600263ffffffff61053e16565b815481106103f657fe5b906000526020600020015460018481548110151561041057fe5b60009182526020909120015561042d83600263ffffffff61053e16565b92506103aa565b828214610458578060018481548110151561044b57fe5b6000918252602090912001555b505050565b60008282111561046957fe5b50900390565b600080600083925060018481548110151561048657fe5b9060005260206000200154915061049c846100f4565b90505b60025481111580156104c8575060018054829081106104ba57fe5b906000526020600020015482115b156105145760018054829081106104db57fe5b90600052602060002001546001858154811015156104f557fe5b60009182526020909120015592508261050d816100f4565b905061049f565b838314610538578160018581548110151561052b57fe5b6000918252602090912001555b50505050565b600080828481151561054c57fe5b04949350505050565b8154818355818111156104585760008381526020902061045891810190830161034191905b80821115610318576000815560010161057a5600a165627a7a7230582041034e9038a5c4721bc87b5349fd6a537d733be1bf1d6b261cf58c2e2a00870d0029`

// DeployRootChain deploys a new Ethereum contract, binding an instance of RootChain to it.
func DeployRootChain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RootChain, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootChainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RootChain{RootChainCaller: RootChainCaller{contract: contract}, RootChainTransactor: RootChainTransactor{contract: contract}, RootChainFilterer: RootChainFilterer{contract: contract}}, nil
}

// RootChain is an auto generated Go binding around an Ethereum contract.
type RootChain struct {
	RootChainCaller     // Read-only binding to the contract
	RootChainTransactor // Write-only binding to the contract
	RootChainFilterer   // Log filterer for contract events
}

// RootChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootChainSession struct {
	Contract     *RootChain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootChainCallerSession struct {
	Contract *RootChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RootChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootChainTransactorSession struct {
	Contract     *RootChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RootChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootChainRaw struct {
	Contract *RootChain // Generic contract binding to access the raw methods on
}

// RootChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootChainCallerRaw struct {
	Contract *RootChainCaller // Generic read-only contract binding to access the raw methods on
}

// RootChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootChainTransactorRaw struct {
	Contract *RootChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootChain creates a new instance of RootChain, bound to a specific deployed contract.
func NewRootChain(address common.Address, backend bind.ContractBackend) (*RootChain, error) {
	contract, err := bindRootChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootChain{RootChainCaller: RootChainCaller{contract: contract}, RootChainTransactor: RootChainTransactor{contract: contract}, RootChainFilterer: RootChainFilterer{contract: contract}}, nil
}

// NewRootChainCaller creates a new read-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainCaller(address common.Address, caller bind.ContractCaller) (*RootChainCaller, error) {
	contract, err := bindRootChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainCaller{contract: contract}, nil
}

// NewRootChainTransactor creates a new write-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainTransactor(address common.Address, transactor bind.ContractTransactor) (*RootChainTransactor, error) {
	contract, err := bindRootChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainTransactor{contract: contract}, nil
}

// NewRootChainFilterer creates a new log filterer instance of RootChain, bound to a specific deployed contract.
func NewRootChainFilterer(address common.Address, filterer bind.ContractFilterer) (*RootChainFilterer, error) {
	contract, err := bindRootChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootChainFilterer{contract: contract}, nil
}

// bindRootChain binds a generic wrapper to an already deployed contract.
func bindRootChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChain *RootChainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChain.Contract.RootChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChain *RootChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.Contract.RootChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChain *RootChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChain.Contract.RootChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChain *RootChainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChain *RootChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChain *RootChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChain.Contract.contract.Transact(opts, method, params...)
}

// CHILDBLOCKINTERVAL is a free data retrieval call binding the contract method 0xa831fa07.
//
// Solidity: function CHILD_BLOCK_INTERVAL() constant returns(uint256)
func (_RootChain *RootChainCaller) CHILDBLOCKINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "CHILD_BLOCK_INTERVAL")
	return *ret0, err
}

// CHILDBLOCKINTERVAL is a free data retrieval call binding the contract method 0xa831fa07.
//
// Solidity: function CHILD_BLOCK_INTERVAL() constant returns(uint256)
func (_RootChain *RootChainSession) CHILDBLOCKINTERVAL() (*big.Int, error) {
	return _RootChain.Contract.CHILDBLOCKINTERVAL(&_RootChain.CallOpts)
}

// CHILDBLOCKINTERVAL is a free data retrieval call binding the contract method 0xa831fa07.
//
// Solidity: function CHILD_BLOCK_INTERVAL() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CHILDBLOCKINTERVAL() (*big.Int, error) {
	return _RootChain.Contract.CHILDBLOCKINTERVAL(&_RootChain.CallOpts)
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(root bytes32, timestamp uint256)
func (_RootChain *RootChainCaller) ChildChain(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Root      [32]byte
	Timestamp *big.Int
}, error) {
	ret := new(struct {
		Root      [32]byte
		Timestamp *big.Int
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "childChain", arg0)
	return *ret, err
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(root bytes32, timestamp uint256)
func (_RootChain *RootChainSession) ChildChain(arg0 *big.Int) (struct {
	Root      [32]byte
	Timestamp *big.Int
}, error) {
	return _RootChain.Contract.ChildChain(&_RootChain.CallOpts, arg0)
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(root bytes32, timestamp uint256)
func (_RootChain *RootChainCallerSession) ChildChain(arg0 *big.Int) (struct {
	Root      [32]byte
	Timestamp *big.Int
}, error) {
	return _RootChain.Contract.ChildChain(&_RootChain.CallOpts, arg0)
}

// CurrentChildBlock is a free data retrieval call binding the contract method 0x7a95f1e8.
//
// Solidity: function currentChildBlock() constant returns(uint256)
func (_RootChain *RootChainCaller) CurrentChildBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "currentChildBlock")
	return *ret0, err
}

// CurrentChildBlock is a free data retrieval call binding the contract method 0x7a95f1e8.
//
// Solidity: function currentChildBlock() constant returns(uint256)
func (_RootChain *RootChainSession) CurrentChildBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentChildBlock(&_RootChain.CallOpts)
}

// CurrentChildBlock is a free data retrieval call binding the contract method 0x7a95f1e8.
//
// Solidity: function currentChildBlock() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CurrentChildBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentChildBlock(&_RootChain.CallOpts)
}

// CurrentDepositBlock is a free data retrieval call binding the contract method 0xa98c7f2c.
//
// Solidity: function currentDepositBlock() constant returns(uint256)
func (_RootChain *RootChainCaller) CurrentDepositBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "currentDepositBlock")
	return *ret0, err
}

// CurrentDepositBlock is a free data retrieval call binding the contract method 0xa98c7f2c.
//
// Solidity: function currentDepositBlock() constant returns(uint256)
func (_RootChain *RootChainSession) CurrentDepositBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentDepositBlock(&_RootChain.CallOpts)
}

// CurrentDepositBlock is a free data retrieval call binding the contract method 0xa98c7f2c.
//
// Solidity: function currentDepositBlock() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CurrentDepositBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentDepositBlock(&_RootChain.CallOpts)
}

// CurrentFeeExit is a free data retrieval call binding the contract method 0xcfe7d855.
//
// Solidity: function currentFeeExit() constant returns(uint256)
func (_RootChain *RootChainCaller) CurrentFeeExit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "currentFeeExit")
	return *ret0, err
}

// CurrentFeeExit is a free data retrieval call binding the contract method 0xcfe7d855.
//
// Solidity: function currentFeeExit() constant returns(uint256)
func (_RootChain *RootChainSession) CurrentFeeExit() (*big.Int, error) {
	return _RootChain.Contract.CurrentFeeExit(&_RootChain.CallOpts)
}

// CurrentFeeExit is a free data retrieval call binding the contract method 0xcfe7d855.
//
// Solidity: function currentFeeExit() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CurrentFeeExit() (*big.Int, error) {
	return _RootChain.Contract.CurrentFeeExit(&_RootChain.CallOpts)
}

// Exits is a free data retrieval call binding the contract method 0x342de179.
//
// Solidity: function exits( uint256) constant returns(owner address, token address, amount uint256)
func (_RootChain *RootChainCaller) Exits(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner  common.Address
	Token  common.Address
	Amount *big.Int
}, error) {
	ret := new(struct {
		Owner  common.Address
		Token  common.Address
		Amount *big.Int
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "exits", arg0)
	return *ret, err
}

// Exits is a free data retrieval call binding the contract method 0x342de179.
//
// Solidity: function exits( uint256) constant returns(owner address, token address, amount uint256)
func (_RootChain *RootChainSession) Exits(arg0 *big.Int) (struct {
	Owner  common.Address
	Token  common.Address
	Amount *big.Int
}, error) {
	return _RootChain.Contract.Exits(&_RootChain.CallOpts, arg0)
}

// Exits is a free data retrieval call binding the contract method 0x342de179.
//
// Solidity: function exits( uint256) constant returns(owner address, token address, amount uint256)
func (_RootChain *RootChainCallerSession) Exits(arg0 *big.Int) (struct {
	Owner  common.Address
	Token  common.Address
	Amount *big.Int
}, error) {
	return _RootChain.Contract.Exits(&_RootChain.CallOpts, arg0)
}

// ExitsQueues is a free data retrieval call binding the contract method 0xd11f045c.
//
// Solidity: function exitsQueues( address) constant returns(address)
func (_RootChain *RootChainCaller) ExitsQueues(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "exitsQueues", arg0)
	return *ret0, err
}

// ExitsQueues is a free data retrieval call binding the contract method 0xd11f045c.
//
// Solidity: function exitsQueues( address) constant returns(address)
func (_RootChain *RootChainSession) ExitsQueues(arg0 common.Address) (common.Address, error) {
	return _RootChain.Contract.ExitsQueues(&_RootChain.CallOpts, arg0)
}

// ExitsQueues is a free data retrieval call binding the contract method 0xd11f045c.
//
// Solidity: function exitsQueues( address) constant returns(address)
func (_RootChain *RootChainCallerSession) ExitsQueues(arg0 common.Address) (common.Address, error) {
	return _RootChain.Contract.ExitsQueues(&_RootChain.CallOpts, arg0)
}

// GetChildChain is a free data retrieval call binding the contract method 0x85444de3.
//
// Solidity: function getChildChain(_blockNumber uint256) constant returns(bytes32, uint256)
func (_RootChain *RootChainCaller) GetChildChain(opts *bind.CallOpts, _blockNumber *big.Int) ([32]byte, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _RootChain.contract.Call(opts, out, "getChildChain", _blockNumber)
	return *ret0, *ret1, err
}

// GetChildChain is a free data retrieval call binding the contract method 0x85444de3.
//
// Solidity: function getChildChain(_blockNumber uint256) constant returns(bytes32, uint256)
func (_RootChain *RootChainSession) GetChildChain(_blockNumber *big.Int) ([32]byte, *big.Int, error) {
	return _RootChain.Contract.GetChildChain(&_RootChain.CallOpts, _blockNumber)
}

// GetChildChain is a free data retrieval call binding the contract method 0x85444de3.
//
// Solidity: function getChildChain(_blockNumber uint256) constant returns(bytes32, uint256)
func (_RootChain *RootChainCallerSession) GetChildChain(_blockNumber *big.Int) ([32]byte, *big.Int, error) {
	return _RootChain.Contract.GetChildChain(&_RootChain.CallOpts, _blockNumber)
}

// GetDepositBlock is a free data retrieval call binding the contract method 0xbcd59261.
//
// Solidity: function getDepositBlock() constant returns(uint256)
func (_RootChain *RootChainCaller) GetDepositBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getDepositBlock")
	return *ret0, err
}

// GetDepositBlock is a free data retrieval call binding the contract method 0xbcd59261.
//
// Solidity: function getDepositBlock() constant returns(uint256)
func (_RootChain *RootChainSession) GetDepositBlock() (*big.Int, error) {
	return _RootChain.Contract.GetDepositBlock(&_RootChain.CallOpts)
}

// GetDepositBlock is a free data retrieval call binding the contract method 0xbcd59261.
//
// Solidity: function getDepositBlock() constant returns(uint256)
func (_RootChain *RootChainCallerSession) GetDepositBlock() (*big.Int, error) {
	return _RootChain.Contract.GetDepositBlock(&_RootChain.CallOpts)
}

// GetExit is a free data retrieval call binding the contract method 0xe60f1ff1.
//
// Solidity: function getExit(_utxoPos uint256) constant returns(address, address, uint256)
func (_RootChain *RootChainCaller) GetExit(opts *bind.CallOpts, _utxoPos *big.Int) (common.Address, common.Address, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _RootChain.contract.Call(opts, out, "getExit", _utxoPos)
	return *ret0, *ret1, *ret2, err
}

// GetExit is a free data retrieval call binding the contract method 0xe60f1ff1.
//
// Solidity: function getExit(_utxoPos uint256) constant returns(address, address, uint256)
func (_RootChain *RootChainSession) GetExit(_utxoPos *big.Int) (common.Address, common.Address, *big.Int, error) {
	return _RootChain.Contract.GetExit(&_RootChain.CallOpts, _utxoPos)
}

// GetExit is a free data retrieval call binding the contract method 0xe60f1ff1.
//
// Solidity: function getExit(_utxoPos uint256) constant returns(address, address, uint256)
func (_RootChain *RootChainCallerSession) GetExit(_utxoPos *big.Int) (common.Address, common.Address, *big.Int, error) {
	return _RootChain.Contract.GetExit(&_RootChain.CallOpts, _utxoPos)
}

// GetNextExit is a free data retrieval call binding the contract method 0x8bfe0aec.
//
// Solidity: function getNextExit(_token address) constant returns(uint256, uint256)
func (_RootChain *RootChainCaller) GetNextExit(opts *bind.CallOpts, _token common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _RootChain.contract.Call(opts, out, "getNextExit", _token)
	return *ret0, *ret1, err
}

// GetNextExit is a free data retrieval call binding the contract method 0x8bfe0aec.
//
// Solidity: function getNextExit(_token address) constant returns(uint256, uint256)
func (_RootChain *RootChainSession) GetNextExit(_token common.Address) (*big.Int, *big.Int, error) {
	return _RootChain.Contract.GetNextExit(&_RootChain.CallOpts, _token)
}

// GetNextExit is a free data retrieval call binding the contract method 0x8bfe0aec.
//
// Solidity: function getNextExit(_token address) constant returns(uint256, uint256)
func (_RootChain *RootChainCallerSession) GetNextExit(_token common.Address) (*big.Int, *big.Int, error) {
	return _RootChain.Contract.GetNextExit(&_RootChain.CallOpts, _token)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChain *RootChainCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "operator")
	return *ret0, err
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChain *RootChainSession) Operator() (common.Address, error) {
	return _RootChain.Contract.Operator(&_RootChain.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() constant returns(address)
func (_RootChain *RootChainCallerSession) Operator() (common.Address, error) {
	return _RootChain.Contract.Operator(&_RootChain.CallOpts)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x32773ba3.
//
// Solidity: function challengeExit(_cUtxoPos uint256, _eUtxoIndex uint256, _txBytes bytes, _proof bytes, _sigs bytes, _confirmationSig bytes) returns()
func (_RootChain *RootChainTransactor) ChallengeExit(opts *bind.TransactOpts, _cUtxoPos *big.Int, _eUtxoIndex *big.Int, _txBytes []byte, _proof []byte, _sigs []byte, _confirmationSig []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "challengeExit", _cUtxoPos, _eUtxoIndex, _txBytes, _proof, _sigs, _confirmationSig)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x32773ba3.
//
// Solidity: function challengeExit(_cUtxoPos uint256, _eUtxoIndex uint256, _txBytes bytes, _proof bytes, _sigs bytes, _confirmationSig bytes) returns()
func (_RootChain *RootChainSession) ChallengeExit(_cUtxoPos *big.Int, _eUtxoIndex *big.Int, _txBytes []byte, _proof []byte, _sigs []byte, _confirmationSig []byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeExit(&_RootChain.TransactOpts, _cUtxoPos, _eUtxoIndex, _txBytes, _proof, _sigs, _confirmationSig)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x32773ba3.
//
// Solidity: function challengeExit(_cUtxoPos uint256, _eUtxoIndex uint256, _txBytes bytes, _proof bytes, _sigs bytes, _confirmationSig bytes) returns()
func (_RootChain *RootChainTransactorSession) ChallengeExit(_cUtxoPos *big.Int, _eUtxoIndex *big.Int, _txBytes []byte, _proof []byte, _sigs []byte, _confirmationSig []byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeExit(&_RootChain.TransactOpts, _cUtxoPos, _eUtxoIndex, _txBytes, _proof, _sigs, _confirmationSig)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_RootChain *RootChainTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_RootChain *RootChainSession) Deposit() (*types.Transaction, error) {
	return _RootChain.Contract.Deposit(&_RootChain.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_RootChain *RootChainTransactorSession) Deposit() (*types.Transaction, error) {
	return _RootChain.Contract.Deposit(&_RootChain.TransactOpts)
}

// FinalizeExits is a paid mutator transaction binding the contract method 0x7d7b2051.
//
// Solidity: function finalizeExits(_token address) returns()
func (_RootChain *RootChainTransactor) FinalizeExits(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "finalizeExits", _token)
}

// FinalizeExits is a paid mutator transaction binding the contract method 0x7d7b2051.
//
// Solidity: function finalizeExits(_token address) returns()
func (_RootChain *RootChainSession) FinalizeExits(_token common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeExits(&_RootChain.TransactOpts, _token)
}

// FinalizeExits is a paid mutator transaction binding the contract method 0x7d7b2051.
//
// Solidity: function finalizeExits(_token address) returns()
func (_RootChain *RootChainTransactorSession) FinalizeExits(_token common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeExits(&_RootChain.TransactOpts, _token)
}

// StartDepositExit is a paid mutator transaction binding the contract method 0xf82f7a4b.
//
// Solidity: function startDepositExit(_depositPos uint256, _token address, _amount uint256) returns()
func (_RootChain *RootChainTransactor) StartDepositExit(opts *bind.TransactOpts, _depositPos *big.Int, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startDepositExit", _depositPos, _token, _amount)
}

// StartDepositExit is a paid mutator transaction binding the contract method 0xf82f7a4b.
//
// Solidity: function startDepositExit(_depositPos uint256, _token address, _amount uint256) returns()
func (_RootChain *RootChainSession) StartDepositExit(_depositPos *big.Int, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.StartDepositExit(&_RootChain.TransactOpts, _depositPos, _token, _amount)
}

// StartDepositExit is a paid mutator transaction binding the contract method 0xf82f7a4b.
//
// Solidity: function startDepositExit(_depositPos uint256, _token address, _amount uint256) returns()
func (_RootChain *RootChainTransactorSession) StartDepositExit(_depositPos *big.Int, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.StartDepositExit(&_RootChain.TransactOpts, _depositPos, _token, _amount)
}

// StartExit is a paid mutator transaction binding the contract method 0x1c91a6b9.
//
// Solidity: function startExit(_utxoPos uint256, _txBytes bytes, _proof bytes, _sigs bytes) returns()
func (_RootChain *RootChainTransactor) StartExit(opts *bind.TransactOpts, _utxoPos *big.Int, _txBytes []byte, _proof []byte, _sigs []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startExit", _utxoPos, _txBytes, _proof, _sigs)
}

// StartExit is a paid mutator transaction binding the contract method 0x1c91a6b9.
//
// Solidity: function startExit(_utxoPos uint256, _txBytes bytes, _proof bytes, _sigs bytes) returns()
func (_RootChain *RootChainSession) StartExit(_utxoPos *big.Int, _txBytes []byte, _proof []byte, _sigs []byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, _utxoPos, _txBytes, _proof, _sigs)
}

// StartExit is a paid mutator transaction binding the contract method 0x1c91a6b9.
//
// Solidity: function startExit(_utxoPos uint256, _txBytes bytes, _proof bytes, _sigs bytes) returns()
func (_RootChain *RootChainTransactorSession) StartExit(_utxoPos *big.Int, _txBytes []byte, _proof []byte, _sigs []byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, _utxoPos, _txBytes, _proof, _sigs)
}

// StartFeeExit is a paid mutator transaction binding the contract method 0x40b8d53a.
//
// Solidity: function startFeeExit(_token address, _amount uint256) returns()
func (_RootChain *RootChainTransactor) StartFeeExit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startFeeExit", _token, _amount)
}

// StartFeeExit is a paid mutator transaction binding the contract method 0x40b8d53a.
//
// Solidity: function startFeeExit(_token address, _amount uint256) returns()
func (_RootChain *RootChainSession) StartFeeExit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.StartFeeExit(&_RootChain.TransactOpts, _token, _amount)
}

// StartFeeExit is a paid mutator transaction binding the contract method 0x40b8d53a.
//
// Solidity: function startFeeExit(_token address, _amount uint256) returns()
func (_RootChain *RootChainTransactorSession) StartFeeExit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.StartFeeExit(&_RootChain.TransactOpts, _token, _amount)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(_root bytes32) returns()
func (_RootChain *RootChainTransactor) SubmitBlock(opts *bind.TransactOpts, _root [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitBlock", _root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(_root bytes32) returns()
func (_RootChain *RootChainSession) SubmitBlock(_root [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitBlock(&_RootChain.TransactOpts, _root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(_root bytes32) returns()
func (_RootChain *RootChainTransactorSession) SubmitBlock(_root [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitBlock(&_RootChain.TransactOpts, _root)
}

// RootChainBlockSubmittedIterator is returned from FilterBlockSubmitted and is used to iterate over the raw logs and unpacked data for BlockSubmitted events raised by the RootChain contract.
type RootChainBlockSubmittedIterator struct {
	Event *RootChainBlockSubmitted // Event containing the contract specifics and raw log

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
func (it *RootChainBlockSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainBlockSubmitted)
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
		it.Event = new(RootChainBlockSubmitted)
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
func (it *RootChainBlockSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainBlockSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainBlockSubmitted represents a BlockSubmitted event raised by the RootChain contract.
type RootChainBlockSubmitted struct {
	Root      [32]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlockSubmitted is a free log retrieval operation binding the contract event 0x3be20e672590051e92360aa9dd9247e52541ad35d7a0297ff9e0f94e993295e9.
//
// Solidity: e BlockSubmitted(root bytes32, timestamp uint256)
func (_RootChain *RootChainFilterer) FilterBlockSubmitted(opts *bind.FilterOpts) (*RootChainBlockSubmittedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "BlockSubmitted")
	if err != nil {
		return nil, err
	}
	return &RootChainBlockSubmittedIterator{contract: _RootChain.contract, event: "BlockSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockSubmitted is a free log subscription operation binding the contract event 0x3be20e672590051e92360aa9dd9247e52541ad35d7a0297ff9e0f94e993295e9.
//
// Solidity: e BlockSubmitted(root bytes32, timestamp uint256)
func (_RootChain *RootChainFilterer) WatchBlockSubmitted(opts *bind.WatchOpts, sink chan<- *RootChainBlockSubmitted) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "BlockSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainBlockSubmitted)
				if err := _RootChain.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
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

// RootChainDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the RootChain contract.
type RootChainDepositIterator struct {
	Event *RootChainDeposit // Event containing the contract specifics and raw log

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
func (it *RootChainDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainDeposit)
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
		it.Event = new(RootChainDeposit)
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
func (it *RootChainDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainDeposit represents a Deposit event raised by the RootChain contract.
type RootChainDeposit struct {
	Depositor    common.Address
	DepositBlock *big.Int
	Token        common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xd2f8022f659fd9c8c558f30c00fd5ee7038f7cb56da45095c3e0e7d48b3e0c4b.
//
// Solidity: e Deposit(depositor indexed address, depositBlock indexed uint256, token address, amount uint256)
func (_RootChain *RootChainFilterer) FilterDeposit(opts *bind.FilterOpts, depositor []common.Address, depositBlock []*big.Int) (*RootChainDepositIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var depositBlockRule []interface{}
	for _, depositBlockItem := range depositBlock {
		depositBlockRule = append(depositBlockRule, depositBlockItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "Deposit", depositorRule, depositBlockRule)
	if err != nil {
		return nil, err
	}
	return &RootChainDepositIterator{contract: _RootChain.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xd2f8022f659fd9c8c558f30c00fd5ee7038f7cb56da45095c3e0e7d48b3e0c4b.
//
// Solidity: e Deposit(depositor indexed address, depositBlock indexed uint256, token address, amount uint256)
func (_RootChain *RootChainFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *RootChainDeposit, depositor []common.Address, depositBlock []*big.Int) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var depositBlockRule []interface{}
	for _, depositBlockItem := range depositBlock {
		depositBlockRule = append(depositBlockRule, depositBlockItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "Deposit", depositorRule, depositBlockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainDeposit)
				if err := _RootChain.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// RootChainExitStartedIterator is returned from FilterExitStarted and is used to iterate over the raw logs and unpacked data for ExitStarted events raised by the RootChain contract.
type RootChainExitStartedIterator struct {
	Event *RootChainExitStarted // Event containing the contract specifics and raw log

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
func (it *RootChainExitStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainExitStarted)
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
		it.Event = new(RootChainExitStarted)
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
func (it *RootChainExitStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainExitStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainExitStarted represents a ExitStarted event raised by the RootChain contract.
type RootChainExitStarted struct {
	Exitor  common.Address
	UtxoPos *big.Int
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExitStarted is a free log retrieval operation binding the contract event 0xf0537507c0bb9b823a1e4f5522ed9f25512618d4d169c9eb70a5ee2e9fb8bd57.
//
// Solidity: e ExitStarted(exitor indexed address, utxoPos indexed uint256, token address, amount uint256)
func (_RootChain *RootChainFilterer) FilterExitStarted(opts *bind.FilterOpts, exitor []common.Address, utxoPos []*big.Int) (*RootChainExitStartedIterator, error) {

	var exitorRule []interface{}
	for _, exitorItem := range exitor {
		exitorRule = append(exitorRule, exitorItem)
	}
	var utxoPosRule []interface{}
	for _, utxoPosItem := range utxoPos {
		utxoPosRule = append(utxoPosRule, utxoPosItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "ExitStarted", exitorRule, utxoPosRule)
	if err != nil {
		return nil, err
	}
	return &RootChainExitStartedIterator{contract: _RootChain.contract, event: "ExitStarted", logs: logs, sub: sub}, nil
}

// WatchExitStarted is a free log subscription operation binding the contract event 0xf0537507c0bb9b823a1e4f5522ed9f25512618d4d169c9eb70a5ee2e9fb8bd57.
//
// Solidity: e ExitStarted(exitor indexed address, utxoPos indexed uint256, token address, amount uint256)
func (_RootChain *RootChainFilterer) WatchExitStarted(opts *bind.WatchOpts, sink chan<- *RootChainExitStarted, exitor []common.Address, utxoPos []*big.Int) (event.Subscription, error) {

	var exitorRule []interface{}
	for _, exitorItem := range exitor {
		exitorRule = append(exitorRule, exitorItem)
	}
	var utxoPosRule []interface{}
	for _, utxoPosItem := range utxoPos {
		utxoPosRule = append(utxoPosRule, utxoPosItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "ExitStarted", exitorRule, utxoPosRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainExitStarted)
				if err := _RootChain.contract.UnpackLog(event, "ExitStarted", log); err != nil {
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

// RootChainTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the RootChain contract.
type RootChainTokenAddedIterator struct {
	Event *RootChainTokenAdded // Event containing the contract specifics and raw log

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
func (it *RootChainTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainTokenAdded)
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
		it.Event = new(RootChainTokenAdded)
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
func (it *RootChainTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainTokenAdded represents a TokenAdded event raised by the RootChain contract.
type RootChainTokenAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: e TokenAdded(token address)
func (_RootChain *RootChainFilterer) FilterTokenAdded(opts *bind.FilterOpts) (*RootChainTokenAddedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return &RootChainTokenAddedIterator{contract: _RootChain.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: e TokenAdded(token address)
func (_RootChain *RootChainFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *RootChainTokenAdded) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainTokenAdded)
				if err := _RootChain.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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
