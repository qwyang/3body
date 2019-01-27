package graph

import (
	"fmt"
	"math"
	"testing"
)
type Graph [][]int
func (g Graph)AddEdge(v1,v2,c int){
	g[v1][v2]=c
	g[v2][v1]=c
}
func NewGraph(n int)(G Graph){
	G=make([][]int,n)
	for i:=0;i<n;i++{
		G[i]=make([]int,n)
	}
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			G[i][j]=math.MaxInt32
			if i==j{
				G[i][j]=0
			}
		}
	}
	return
}
func TestFloyd(t *testing.T) {
	n:=7
	G:=NewGraph(n)
	G.AddEdge(0,1,1)
	G.AddEdge(0,2,1)
	G.AddEdge(0,3,9)
	G.AddEdge(1,3,3)
	G.AddEdge(1,4,1)
	G.AddEdge(2,3,1)
	G.AddEdge(2,5,10)
	//G.AddEdge(3,4,1)
	G.AddEdge(3,5,3)
	G.AddEdge(3,6,5)
	G.AddEdge(4,6,1)
	G.AddEdge(5,6,1)
	cost,path:=Floyd(G)
	fmt.Println(cost)
	fmt.Println(path)
}