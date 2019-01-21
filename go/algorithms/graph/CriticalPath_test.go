package graph

import (
	"fmt"
	"testing"
)

func TestCriticalPath(t *testing.T) {
	g := NewAdjacencyList(DG)
	g.AddEdge("v1", "v2", 1)
	g.AddEdge("v1", "v3", 1)
	g.AddEdge("v1", "v4", 9)
	g.AddEdge("v2", "v4", 1)
	g.AddEdge("v2", "v5", 1)
	g.AddEdge("v3", "v6", 10)
	g.AddEdge("v4", "v3", 1)
	g.AddEdge("v4", "v6", 3)
	g.AddEdge("v4", "v7", 5)
	g.AddEdge("v5", "v4", -10)
	g.AddEdge("v5", "v7", 1)
	g.AddEdge("v7", "v6", 1)
	crit := NewCriticalPath(g, 0,5)
	crit.Run()
	crit.PrintPath(1)
	fmt.Println()
	crit.PrintPath(2)
	fmt.Println()
	crit.PrintPath(3)
	fmt.Println()
	crit.PrintPath(4)
	fmt.Println()
	crit.PrintPath(5)
	fmt.Println()
	crit.PrintResult()
}