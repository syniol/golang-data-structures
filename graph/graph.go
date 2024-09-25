package graph

type NodeName string

type Graph struct {
	Nodes []*Node
	Edges map[NodeName][]*Edge
}

type Node struct {
	name NodeName
}

type Edge struct {
	node   *Node
	weight int
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: *new([]*Node),
		Edges: make(map[NodeName][]*Edge),
	}
}

func (g *Graph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) RemoveNode(nodeName NodeName) {}

func (g *Graph) AddEdge(n1, n2 *Node, weight int) {
	g.Edges[n1.name] = append(g.Edges[n1.name], &Edge{
		node:   n2,
		weight: weight,
	})

	g.Edges[n2.name] = append(g.Edges[n2.name], &Edge{
		node:   n1,
		weight: weight,
	})
}

func (g *Graph) RemoveEdge(n1, n2 NodeName) {}
