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

func (g *Graph) RemoveNode(nodeName NodeName) {
	r := -1
	for i, n := range g.Nodes {
		if n.name == nodeName {
			r = i
		}
	}

	if r > -1 {
		g.Nodes[r] = g.Nodes[len(g.Nodes)-1]
		g.Nodes = g.Nodes[:len(g.Nodes)-1]
	}

	delete(g.Edges, nodeName)

	for n := range g.Edges {
		g.removeEdge(n, nodeName)
	}
}

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

func (g *Graph) RemoveEdge(n1, n2 NodeName) {
	g.removeEdge(n1, n2)
	g.removeEdge(n2, n1)
}

func (g *Graph) removeEdge(m, n NodeName) {
	edges := g.Edges[m]
	r := -1

	for i, edge := range edges {
		if edge.node.name == n {
			r = i
		}
	}

	if r > -1 {
		edges[r] = edges[len(edges)-1]
		g.Edges[m] = edges[:len(edges)-1]
	}
}
