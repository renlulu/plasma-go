package util

type Node struct {
	Data string
	Left Node
	Right Node
}

func MakeNode(data string, left Node, right Node) Node {
	return Node{
		Data:data,
		Left:left,
		Right:right,
	}
}

