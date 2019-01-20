package graph

import (
	"fmt"
	"testing"
)

func TestTopSort1(t *testing.T) {
	g := NewAdjacencyList(DG)
	g.AddEdge("v1","v2",1)
	g.AddEdge("v1","v3",1)
	g.AddEdge("v1","v4",1)
	g.AddEdge("v2","v4",1)
	g.AddEdge("v2","v5",1)
	g.AddEdge("v3","v6",1)
	g.AddEdge("v4","v3",1)
	g.AddEdge("v4","v6",1)
	g.AddEdge("v4","v7",1)
	g.AddEdge("v5","v4",1)
	g.AddEdge("v5","v7",1)
	g.AddEdge("v7","v6",1)
	r := TopSort1(g)
	fmt.Println(r)
	r = TopSort2(g)
	fmt.Println(r)
}
