package core

type Block struct {
	transactions []Transaction
	number uint64
	sig []byte
}