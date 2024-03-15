// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1BlockStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"number\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint64\"},{\"astId\":1001,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"timestamp\",\"offset\":8,\"slot\":\"0\",\"type\":\"t_uint64\"},{\"astId\":1002,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"basefee\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"hash\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_bytes32\"},{\"astId\":1004,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"sequenceNumber\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_uint64\"},{\"astId\":1005,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"blobBaseFeeScalar\",\"offset\":8,\"slot\":\"3\",\"type\":\"t_uint32\"},{\"astId\":1006,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"baseFeeScalar\",\"offset\":12,\"slot\":\"3\",\"type\":\"t_uint32\"},{\"astId\":1007,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_bytes32\"},{\"astId\":1008,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"l1FeeOverhead\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"l1FeeScalar\",\"offset\":0,\"slot\":\"6\",\"type\":\"t_uint256\"},{\"astId\":1010,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"blobBaseFee\",\"offset\":0,\"slot\":\"7\",\"type\":\"t_uint256\"},{\"astId\":1011,\"contract\":\"src/L2/L1Block.sol:L1Block\",\"label\":\"dependencySet\",\"offset\":0,\"slot\":\"8\",\"type\":\"t_array(t_uint256)dyn_storage\"}],\"types\":{\"t_array(t_uint256)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"uint256[]\",\"numberOfBytes\":\"32\",\"base\":\"t_uint256\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"}}}"

var L1BlockStorageLayout = new(solc.StorageLayout)

var L1BlockDeployedBin = "0x608060405234801561001057600080fd5b50600436106101515760003560e01c80638381f58a116100cd578063c598591811610081578063e591b28211610066578063e591b282146102f8578063e81b2c6d14610338578063f82061401461034157600080fd5b8063c5985918146102b5578063e38bbc32146102d557600080fd5b80639689fb4a116100b25780639689fb4a146102795780639e8c49661461028c578063b80777ea1461029557600080fd5b80638381f58a1461025c5780638b239f731461027057600080fd5b80635cf249691161012457806364ca23ef1161010957806364ca23ef146101f657806368d5dca614610223578063760ee04d1461025457600080fd5b80635cf24969146101d85780635eb30fa3146101e157600080fd5b8063015d8eb91461015657806309bd5a601461016b578063440a5e201461018757806354fd4d501461018f575b600080fd5b610169610164366004610689565b61034a565b005b61017460025481565b6040519081526020015b60405180910390f35b610169610489565b6101cb6040518060400160405280600581526020017f312e332e3000000000000000000000000000000000000000000000000000000081525081565b60405161017e91906106fb565b61017460015481565b60085460405160ff909116815260200161017e565b60035461020a9067ffffffffffffffff1681565b60405167ffffffffffffffff909116815260200161017e565b60035461023f9068010000000000000000900463ffffffff1681565b60405163ffffffff909116815260200161017e565b6101696104fc565b60005461020a9067ffffffffffffffff1681565b61017460055481565b61017461028736600461076e565b6105e5565b61017460065481565b60005461020a9068010000000000000000900467ffffffffffffffff1681565b60035461023f906c01000000000000000000000000900463ffffffff1681565b6102e86102e336600461076e565b610606565b604051901515815260200161017e565b61031373deaddeaddeaddeaddeaddeaddeaddeaddead000181565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161017e565b61017460045481565b61017460075481565b3373deaddeaddeaddeaddeaddeaddeaddeaddead0001146103f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603b60248201527f4c31426c6f636b3a206f6e6c7920746865206465706f7369746f72206163636f60448201527f756e742063616e20736574204c3120626c6f636b2076616c7565730000000000606482015260840160405180910390fd5b6000805467ffffffffffffffff98891668010000000000000000027fffffffffffffffffffffffffffffffff00000000000000000000000000000000909116998916999099179890981790975560019490945560029290925560038054919094167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009190911617909255600491909155600555600655565b7f3cc50b45000000000000000000000000000000000000000000000000000000003373deaddeaddeaddeaddeaddeaddeaddeaddead0001146104cf57806000526004601cfd5b60043560801c60035560143560801c60005560243560015560443560075560643560025560843560045550565b7f3cc50b45000000000000000000000000000000000000000000000000000000007f44165b6a000000000000000000000000000000000000000000000000000000003373deaddeaddeaddeaddeaddeaddeaddeaddead00011461056357816000526004601cfd5b60043560801c60035560143560801c60005560243560015560443560075560643560025560843560045560a43560f81c366020820260a50118156105ab57816000526004601cfd5b6008600052602060002060005b828110156105de5760088054600190810190915560a560208302013583830155016105b8565b5050505050565b600881815481106105f557600080fd5b600091825260209091200154905081565b600046820361061757506001919050565b60005b60085481101561066357826008828154811061063857610638610787565b9060005260206000200154036106515750600192915050565b8061065b816107b6565b91505061061a565b50600092915050565b803567ffffffffffffffff8116811461068457600080fd5b919050565b600080600080600080600080610100898b0312156106a657600080fd5b6106af8961066c565b97506106bd60208a0161066c565b965060408901359550606089013594506106d960808a0161066c565b979a969950949793969560a0850135955060c08501359460e001359350915050565b600060208083528351808285015260005b818110156107285785810183015185820160400152820161070c565b8181111561073a576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60006020828403121561078057600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361080e577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(L1BlockStorageLayoutJSON), L1BlockStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1Block"] = L1BlockStorageLayout
	deployedBytecodes["L1Block"] = L1BlockDeployedBin
	immutableReferences["L1Block"] = false
}
