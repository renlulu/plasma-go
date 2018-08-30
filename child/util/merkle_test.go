package util

import (
	"fmt"
	"testing"
)

func TestMakeMerkle(t *testing.T) {
	s := []string{"	00000000000000000023f353e60b6f08351480251202f49c6b3b7f59777f924d", "0000000000000000001e66ad1a5f78ea6efe8a006d5c6a787b3d21ee575d7154","00000000000000000023f353e60b6f08351480251202f49c6b3b7f59777f924d","00000000000000000023f353e60b6f08351480251202f49c6b3b7f59777f924d"}
	m, _ := MakeMerkle(3, s)
	fmt.Println(m.Depth)
	fmt.Println(len(m.Leaves))
	fmt.Printf("%+v\n", m.Leaves)
	fmt.Println(m.Root)
}
