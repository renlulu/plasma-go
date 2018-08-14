package util

import (
	"testing"
	"fmt"
)

func Test_Encode(t *testing.T) {
	encoded := encodeUTXOId(19999,2,2)
	fmt.Println(encoded)
	blkNum,txIndex,oIndex := decodeUTXOId(encoded)
	fmt.Println(blkNum,txIndex,oIndex)
}


