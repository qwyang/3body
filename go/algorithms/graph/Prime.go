package graph

import (
	"math"
)

/*
最小生成树prime算法。
对于集合S中的顶点u，选择cost(u,v)最小的且v不在S中。一次循环选择一个顶点。
*/
func Prime(g GraphInterface)(spt [][]int){
	known:=make([]bool,g.VertexNum())
	cost := make([]int,g.VertexNum())
	for i:=0;i<len(cost);i++{cost[i]=math.MaxInt32}
	path := make([]int,g.VertexNum())
	spt = make([][]int,g.VertexNum())
	//初始化
	cost[0] = 0;known[0]=true
	vertex := g.GetVertex(0)
	for _,n := range vertex.AdjVertexes(){
		cuv := n.Cost()
		v := n.Index()
		if cuv < cost[v] {
			cost[v] = cuv
			path[v] = 0
		}
	}
	for{
		u := -1
		min := math.MaxInt32
		for i,flag := range known{//O(V^2）
			if flag == true{continue}
			if cost[i] < min{min = cost[i];	u = i}
		}
		if u == -1 {break}
		spt[path[u]]=append(spt[path[u]], u)
		known[u] = true
		vertex := g.GetVertex(u)
		for _,n := range vertex.AdjVertexes(){//O(E)
			cuv := n.Cost()
			v := n.Index()
			if cuv < cost[v] {
				cost[v] = cuv
				path[v] = u
			}
		}
	}
	return
}
