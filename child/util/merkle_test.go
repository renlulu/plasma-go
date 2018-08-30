package util

import (
	"fmt"
	"testing"
)

func TestMakeMerkle(t *testing.T) {
	s := []string{"1", "2"}
	m, _ := MakeMerkle(2, s)
	fmt.Println(m.Depth)
	fmt.Println(len(m.Leaves))
	fmt.Printf("%+v\n", m.Leaves)
	fmt.Println(m.Root)
}
