package core

import (
	"testing"
	"fmt"
)

func Test_GetGenesisBlock(t *testing.T) {
	gBlock := GetGenesisBlock()
	fmt.Println(gBlock)
}
