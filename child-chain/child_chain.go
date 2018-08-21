package child_chain

import (
	"github.com/renlulu/plasma-go/child-chain/core"
	"github.com/renlulu/plasma-go/root-chain/artifact"
)

type ChildChain struct {
	RootChain *chain.RootChain
	Chain *core.Chain
	CurrentBlock *core.Block
	operator string
}

func MakeChildChain(rootChain *chain.RootChain, operator string) ChildChain {
	childChain := ChildChain{
		rootChain,core.MakeChain(),&core.Block{Number:0}, operator,
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

