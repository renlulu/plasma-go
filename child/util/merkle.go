package util


type Merkle struct {
	Root string
	Leaves []string
	Tree []Node

}

func (m *Merkle) createNodes() []string {
	var s []string

	for _, l := range m.Leaves {
		s = append(s, l)
	}

	return s
}


func (m *Merkle) createTree(leaves []Node) {
	if len(leaves) == 1 {
		m.Root = leaves[0].Data
	}

	nextLevel := len(leaves)
	var treeLevel []Node
	for i := 0;i < nextLevel; i = i + 2 {
		hash := Sha3(leaves[i].Data + leaves[i + 1].Data)
		parentNode := MakeNode(string(hash[:]),leaves[i],leaves[i + 1])
		treeLevel = append(treeLevel,parentNode)
	}

	m.Tree = append(m.Tree,treeLevel...)
	m.createTree(treeLevel)
}