package core

type Block struct {
	Transactions []*Transaction 	`json:"transactions"`
	Number       uint64				`json:"number"`
	Sig          []byte				`json:"sig"`
}

func MakeBlock(txs []*Transaction, number uint64) Block {
	return Block{
		Transactions:txs,Number:number,
	}
}