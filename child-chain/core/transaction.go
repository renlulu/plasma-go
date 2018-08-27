package core

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-ethereum/rlp"
	"github.com/renlulu/plasma-go/child-chain/util"
)

type Transaction struct {
	// first input
	BlkNum1  uint64 `json:"blkNum1"`
	TxIndex1 uint64 `json:"txIndex1"`
	OIndex1  uint64 `json:"oIndex1"`
	Sig1     []byte `json:"sig1"`
	Spend1   bool

	// second input
	BlkNum2  uint64 `json:"blkNum2"`
	TxIndex2 uint64 `json:"txIndex2"`
	OIndex2  uint64 `json:"oIndex2"`
	Sig2     []byte `json:"sig2"`
	Spend2   bool

	// output
	NewOwner1 common.Address `json:"newOwner1"`
	Amount1   uint64         `json:"amount1"`
	NewOwner2 common.Address `json:"newOwner2"`
	Amount2   uint64         `json:"amount2"`
}

type UnsignedTx struct {
	// first input
	BlkNum1  uint64 `json:"blkNum1"`
	TxIndex1 uint64 `json:"txIndex1"`
	OIndex1  uint64 `json:"oIndex1"`

	// second input
	BlkNum2  uint64 `json:"blkNum2"`
	TxIndex2 uint64 `json:"txIndex2"`
	OIndex2  uint64 `json:"oIndex2"`

	// output
	NewOwner1 common.Address `json:"newOwner1"`
	Amount1   uint64         `json:"amount1"`
	NewOwner2 common.Address `json:"newOwner2"`
	Amount2   uint64         `json:"amount2"`
}

func (transaction *Transaction) UnsignedTransaction() UnsignedTx {
	return UnsignedTx{
		BlkNum1:   transaction.BlkNum1,
		TxIndex1:  transaction.TxIndex1,
		OIndex1:   transaction.OIndex1,
		BlkNum2:   transaction.BlkNum2,
		TxIndex2:  transaction.TxIndex2,
		OIndex2:   transaction.OIndex2,
		NewOwner1: transaction.NewOwner1,
		Amount1:   transaction.Amount1,
		NewOwner2: transaction.NewOwner2,
		Amount2:   transaction.Amount2,
	}
}

func MakeTransaction(address common.Address, amount uint64) *Transaction {
	return &Transaction{
		NewOwner1: address, Amount1: amount,
	}
}

type UTXO struct {
	blkNum  uint64
	txIndex uint64
	oIndex  uint64
	sig     []byte
	spend   bool
}

func MakeUTXO(blkNum uint64, txIndex uint64, oIndex uint64, sig []byte, spend bool) UTXO {
	utxo := UTXO{
		blkNum:  blkNum,
		txIndex: txIndex,
		oIndex:  oIndex,
		sig:     sig,
		spend:   spend,
	}
	return utxo
}

func (tx *Transaction) Encode() []byte {
	b, e := rlp.EncodeToBytes(tx.UnsignedTransaction())
	if e != nil {
		return nil
	}
	return b
}

func (tx *Transaction) MerkleHash() [32]byte {
	seed := string(tx.Encode()) + string(tx.Sig1) + string(tx.Sig2)
	return util.Sha3(seed)
}
