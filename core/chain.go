package core

import "github.com/renlulu/plasma-go/util"

type Chain struct {
	Blocks []Block
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

func (chain *Chain) GetTransaction(utxoId uint64) Transaction {
	blkNum,txindex,_ := util.DecodeUTXOId(utxoId)
	return chain.Blocks[blkNum].transactions[txindex]
}