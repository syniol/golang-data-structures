package graph

import (
	"testing"
)

func TestGraph_AddNode(t *testing.T) {
	sut := NewGraph()
	sut.AddNode(&Node{
		name: "Black Node",
	})

	if len(sut.Nodes) != 1 {
		t.Error("it was expecting one node inside the Graph Object")
	}
}

func TestGraph_RemoveNode(t *testing.T) {
	sut := NewGraph()
	const removedNodeName = "First Node"

	sut.AddNode(&Node{name: removedNodeName})
	sut.AddNode(&Node{name: "Second Node"})

	sut.RemoveNode(removedNodeName)

	if len(sut.Nodes) != 1 {
		t.Error("it was expecting only one node")
	}

	if sut.Edges[removedNodeName] != nil {
		t.Error("it was expecting to have no edge for removed Node")
	}
}

func TestGraph_AddEdge(t *testing.T) {
	sut := NewGraph()

	firstNode := &Node{
		name: "First Node",
	}

	secondNode := &Node{
		name: "Second Node",
	}

	sut.AddNode(firstNode)
	sut.AddNode(secondNode)

	sut.AddEdge(firstNode, secondNode, 5)

	if len(sut.Edges[firstNode.name]) != 1 {
		t.Error("It was expecting an edge")
	}

	if len(sut.Edges[secondNode.name]) != 1 {
		t.Error("It was expecting an edge")
	}

	if sut.Edges[secondNode.name][0].node != firstNode {
		t.Error("it was expecting an edge from second node to first node")
	}

	if sut.Edges[firstNode.name][0].node != secondNode {
		t.Error("it was expecting an edge from first node to second node")
	}
}
