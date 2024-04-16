package data_structures

type Node struct {
	Nodes []*Node
}

func (n *Node) AddEdge(node *Node) {}

type Graph struct {
	Nodes []*Node
}

func (g *Graph) AddNode() *Node {
	babyNode := &Node{}

	g.Nodes = append(g.Nodes, babyNode)

	return babyNode
}
