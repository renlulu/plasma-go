package core

import (
	"github.com/renlulu/plasma-go/child-chain/util"
	"fmt"
)

type Chain struct {
	Blocks []*Block `json:"blocks"`
	BlockPool map[uint64][]*Block `json:"block_pool"`
	NextDepositBlock uint64 `json:"next_deposit_block"`
	NextChildBlock uint64 `json:"next_child_block"`
	ChildBlockInterval uint64 `json:"child_block_interval"`
	BlockNum uint64 `json:"block_number"` //指向下一个区块的指针，也代表区块数量
}


func MakeChain() *Chain {
	c := Chain{
		Blocks: make([]*Block,10),
		BlockPool: make(map[uint64][]*Block,10),
		NextDepositBlock: 0,
		ChildBlockInterval: 1000,
		BlockNum: 0,
	}

	c.NextChildBlock = c.ChildBlockInterval
	return &c
}

func (chain *Chain) AddBlock(block *Block) bool {
	isNextChildBlock := block.Number == chain.NextChildBlock
	isNextDepositBlock := block.Number == chain.NextDepositBlock

	// block should be add to the head
	if isNextChildBlock || isNextDepositBlock {
		fmt.Printf("block should be add to the head\n")
		// validate
		chain.validateBlock(block)

		// insert
		chain.applyBlock(block)

		// update head state
		if isNextChildBlock {
			chain.NextChildBlock = chain.NextChildBlock + 1
			chain.NextDepositBlock = chain.NextDepositBlock + 1
		} else {
			chain.NextDepositBlock = chain.NextDepositBlock + 1
		}

	} else if block.Number > chain.NextDepositBlock {
		// or
		fmt.Printf("or \n")
		parentBlockNumber := block.Number - 1
		if _,ok := chain.BlockPool[parentBlockNumber]; !ok {
			chain.BlockPool[parentBlockNumber] = make([]*Block,0)
		}

		chain.BlockPool[parentBlockNumber] = append(chain.BlockPool[parentBlockNumber],block)
		return false

	} else {
		// block already exists
		fmt.Printf("block already exists\n")
		return false
	}

	if _,ok := chain.BlockPool[block.Number]; ok {
		for _,blk := range chain.BlockPool[block.Number] {
			chain.AddBlock(blk)
		}
		delete(chain.BlockPool,block.Number)
	}

	return true



}

func (chain *Chain) GetBlockNum() uint64{
	return chain.BlockNum
}


func (chain *Chain) validateBlock(block *Block) {
	//TODO
}


func (chain *Chain) GetBlock(blkNum uint64) *Block {
	return chain.Blocks[blkNum]
}

func (chain *Chain) MarkUTXOSpend(utxoId uint64) {
	_,_,oIndex := util.DecodeUTXOId(utxoId)
	tx := chain.GetTransaction(oIndex)
	if oIndex == 0 {
		tx.Spend1 = true
	} else {
		tx.Spend2 = true
	}
}

func (chain *Chain) GetTransaction(utxoId uint64) *Transaction {
	fmt.Printf("utxo is is %d\n",utxoId)
	blkNum,txindex,_ := util.DecodeUTXOId(utxoId)
	fmt.Printf("blkNum is %d,txindex is %d",blkNum,txindex)
	return chain.Blocks[blkNum].Transactions[txindex]
}

func (chain *Chain) applyTransaction(tx *Transaction) {
	var uxtos []UTXO
	uxtos = append(uxtos,MakeUTXO(tx.BlkNum1,tx.TxIndex1,tx.OIndex1,tx.Sig1,tx.Spend1))
	uxtos = append(uxtos,MakeUTXO(tx.BlkNum2,tx.TxIndex2,tx.OIndex2,tx.Sig2,tx.Spend2))

	for _,utxo := range uxtos {
		if utxo.blkNum == 0 {
			continue
		}
		id := util.EncodeUTXOId(utxo.blkNum,utxo.txIndex,utxo.oIndex)
		chain.MarkUTXOSpend(id)
	}
}

func (chain *Chain) applyBlock(block *Block) {
	for _,tx := range block.Transactions {
		chain.applyTransaction(tx)
	}

	chain.Blocks[chain.BlockNum] = block
	chain.BlockNum = chain.BlockNum + 1
}

