package core

import "plasma-go/util"

type Transaction struct {
	// first input
	blkNum1 int32
	txIndex1 int32
	oIndex1 int32
	sig1 []byte

	// second input
	blkNum2 int32
	txIndex int32
	oIndex2 int32
	sig2 []byte

	// output
	newOwner1 util.Address
	amount1 int32
	newOwner2 util.Address
	amount2 int32
}