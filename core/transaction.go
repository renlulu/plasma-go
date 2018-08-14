package core

import (
	"github.com/renlulu/plasma-go/util"
)

type Transaction struct {
	// first input
	blkNum1 uint64
	txIndex1 uint64
	oIndex1 uint64
	sig1 []byte
	spend1 bool

	// second input
	blkNum2 uint64
	txIndex uint64
	oIndex2 uint64
	sig2 []byte
	spend2 bool

	// output
	newOwner1 util.Address
	amount1 uint64
	newOwner2 util.Address
	amount2 uint64
}