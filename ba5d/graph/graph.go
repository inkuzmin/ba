package graph


type Graph struct {
	Nodes []Node
}

type Node struct {
	Label int
	Predecessors []int
	Weights []int
}

func (b *Node) AddLabel(i int) {
	b.Label = i
}

func (b *Node) AddPredecessor(a int, w int) {
	b.Predecessors = append(b.Predecessors, a)
	b.Weights = append(b.Weights, w)
}

func (b *Node) RemovePredecessor(a int) {
	l := len(b.Predecessors)
	for i := 0; i < l; i++ {
		if b.Predecessors[i] == a {
			b.Predecessors = append(b.Predecessors[:i], b.Predecessors[i+1:]...)
			b.Weights = append(b.Weights[:i], b.Weights[i+1:]...)
			i--
			l--
		}
	}
}

func (g *Graph) RemoveNode(a int) {
	l := len(g.Nodes)
	for i := 0; i < l; i++ {
		if g.Nodes[i].Label == a {
			g.Nodes = append(g.Nodes[:i], g.Nodes[i+1:]...)
			i--
			l--
		}
	}
}