package util

import (
	"errors"
	"math"
)

type Merkle struct {
	Depth  int
	Root   string
	Leaves []string
	Tree   []Node
}

func MakeMerkle(depth int, leaves []string) (Merkle, error) {
	if depth < 1 {
		return Merkle{}, errors.New("depth should be at least 1")
	}

	leafCount := uint64(math.Pow(float64(depth), 2))
	nullLeaves := makeNullHash(leafCount - uint64(len(leaves)))
	leaves = append(leaves, nullLeaves...)

	m := Merkle{
		Depth:  depth,
		Leaves: leaves,
	}
	nodes := m.createNodes()
	m.createTree(nodes)

	return m, nil
}

func makeNullHash(count uint64) []string {
	var h []string
	for i := uint64(0); i < count; i++ {
		var b = []byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}
		h = append(h, string(b[:]))
	}
	return h
}

func (m *Merkle) createNodes() []Node {
	var n []Node

	for _, l := range m.Leaves {
		n = append(n, MakeNode(l, Node{}, Node{}))
	}

	return n
}

func (m *Merkle) createTree(leaves []Node) {
	if len(leaves) == 1 {
		m.Root = leaves[0].Data
		return
	}

	nextLevel := len(leaves)
	var treeLevel []Node
	for i := 0; i < nextLevel-1; i = i + 2 {
		hash := Sha3(leaves[i].Data + leaves[i+1].Data)
		parentNode := MakeNode(string(hash[:]), leaves[i], leaves[i+1])
		treeLevel = append(treeLevel, parentNode)
	}

	m.Tree = append(m.Tree, treeLevel...)
	m.createTree(treeLevel)
}
