package util

import (
	"testing"
	"fmt"
)

func Test_Encode(t *testing.T) {
	encoded := EncodeUTXOId(19999,2,2)
	fmt.Println(encoded)
	blkNum,txIndex,oIndex := DecodeUTXOId(encoded)
	fmt.Println(blkNum,txIndex,oIndex)
}


