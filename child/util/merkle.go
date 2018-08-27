package util

type Merkle struct {
	Leaves []string
}

func (m *Merkle) createNodes() []string {
	var s []string

	for _, l := range m.Leaves {
		s = append(s, l)
	}

	return s
}
