package core

import "github.com/renlulu/plasma-go/util"

type Chain struct {
	Blocks []*Block
}


func (chain *Chain) GetBlock(blkNum uint64) *Block {
	return chain.Blocks[blkNum]
}

func (chain *Chain) MarkUTXOSpend(utxoId uint64) {
	_,_,oIndex := util.DecodeUTXOId(utxoId)
	tx := chain.GetTransaction(oIndex)
	if oIndex == 0 {
		tx.spend1 = true
	} else {
		tx.spend2 = true
	}
}

func (chain *Chain) GetTransaction(utxoId uint64) *Transaction {
	blkNum,txindex,_ := util.DecodeUTXOId(utxoId)
	return chain.Blocks[blkNum].transactions[txindex]
}

func (chain *Chain) applyTransaction(tx *Transaction) {
	var uxtos []UTXO
	uxtos = append(uxtos,MakeUTXO(tx.blkNum1,tx.txIndex1,tx.oIndex1,tx.sig1,tx.spend1))
	uxtos = append(uxtos,MakeUTXO(tx.blkNum2,tx.txIndex2,tx.oIndex2,tx.sig2,tx.spend2))

	for _,utxo := range uxtos {
		if utxo.blkNum == 0 {
			continue
		}
		id := util.EncodeUTXOId(utxo.blkNum,utxo.txIndex,utxo.oIndex)
		chain.MarkUTXOSpend(id)
	}
}

func (chain *Chain) applyBlock(block *Block) {
	for _,tx := range block.transactions {
		chain.applyTransaction(tx)
	}

	chain.Blocks[len(chain.Blocks)] = block
}