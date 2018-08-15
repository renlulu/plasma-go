package child_chain

import (
	"github.com/renlulu/plasma-go/root-chain"
	"github.com/renlulu/plasma-go/core"
)

type ChildChain struct {
	RootChain *root_chain.RootChain
	Chain *core.Chain
	CurrentBlock *core.Block
}

func MakeChildChain(rootChain *root_chain.RootChain) ChildChain {
	childChain := ChildChain{
		rootChain,core.MakeChain(),&core.Block{Number:0},
	}
	return childChain
}

func (chain *ChildChain) SubmitBlock(block core.Block) {
	chain.Chain.AddBlock(&block)
	//TODO 对接公链
	chain.CurrentBlock = &core.Block{Number:chain.Chain.NextChildBlock}
}


func (chain *ChildChain) GetBlock(blkNum uint64) *core.Block {
	return chain.Chain.GetBlock(blkNum)
}

func (chain *ChildChain) GetCurrentBlock() *core.Block {
	return chain.CurrentBlock
}

func (chain *ChildChain) GetTransaction(utxoId uint64) *core.Transaction {
	return chain.Chain.GetTransaction(utxoId)
}

func (chain *ChildChain) GetBlockNum() uint64 {
	return chain.Chain.GetBlockNum()
}

