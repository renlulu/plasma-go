package core

type Block struct {
	transactions []Transaction
	number int32
	sig []byte
}