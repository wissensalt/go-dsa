package main

import (
	"fmt"
)

type Node struct {
	Item interface{}
}

type Graph struct {
	Nodes []*Node
	Edges map[*Node][]*Node
}

type UndirectedGraphImplementation interface {
	AddNode(item interface{}) *Node
	AddEdge(n1, n2 *Node)
	PrintGraph()
}

func (g *Graph) AddNode(item interface{}) *Node {
	newNode := &Node{item}
	g.Nodes = append(g.Nodes, newNode)

	return newNode
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	if g.Edges == nil {
		g.Edges = map[*Node][]*Node{}
	}

	g.Edges[n1] = append(g.Edges[n1], n2)
	g.Edges[n2] = append(g.Edges[n2], n1)
}

func (g *Graph) PrintGraph() {
	for _, node := range g.Nodes {
		fmt.Printf("Node: %v\n", node.Item)
		for _, edge := range g.Edges[node] {
			fmt.Printf("Edge: %v ", edge.Item)
		}
		fmt.Println()
	}
}

func main() {
	graph := Graph{}
	n0 := graph.AddNode(0)
	n1 := graph.AddNode(1)
	n2 := graph.AddNode(2)
	n3 := graph.AddNode(3)
	n4 := graph.AddNode(4)
	n5 := graph.AddNode(5)
	n6 := graph.AddNode(6)

	graph.AddEdge(n0, n1)
	graph.AddEdge(n0, n2)
	graph.AddEdge(n1, n3)
	graph.AddEdge(n2, n4)
	graph.AddEdge(n3, n4)
	graph.AddEdge(n4, n5)
	graph.AddEdge(n5, n6)

	graph.PrintGraph()
}
