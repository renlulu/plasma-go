package core

type Block struct {
	transactions []*Transaction
	number uint64
	sig []byte
}

func GetGenesisBlock() Block {
	return Block{
		number:0,
		sig: []byte("xiaohuo"),
	}
}