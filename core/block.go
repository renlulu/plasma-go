package core

type 	Block struct {
	Transactions []*Transaction 	`json:"transactions"`
	Number       uint64				`json:"number"`
	Sig          []byte				`json:"sig"`
}