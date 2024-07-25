package data_structures

type GraphNodeName string

type Graph struct {
	Nodes []*GraphNode
	Edges map[GraphNodeName][]*GraphEdge
}

type GraphNode struct {
	name GraphNodeName
}

type GraphEdge struct {
	node   GraphNode
	weight int
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: *new([]*GraphNode),
		Edges: make(map[GraphNodeName][]*GraphEdge),
	}
}

func (g *Graph) AddNode(n *GraphNode) {}

func (g *Graph) AddEdge(n1, n2 *GraphNode, weight int) {}

func (g *Graph) RemoveEdge(n1, n2 GraphNodeName) {}
