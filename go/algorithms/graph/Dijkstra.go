package graph

import (
	"fmt"
	"math"
)

/*
单源最短路径问题，给定一个图g和顶点s，找出s到达其余顶点的最小距离dist和路径path。
主过程：
1.扫描表格，取出未知节点中距离s最小的顶点v。FindMinDistVertex()。
2.对于v的所有边(v,w),如果dist(start,w)>early(start,v)+early(v,w),更新dist(start,w),path(start,w)=v。
3.重复上述过程。
初始化：
三个数组：known,path,early
辅助函数：printPath(vertex)
*/

type Dijkstra struct{
	known []bool
	path,dist []int
	g GraphInterface
	s int
}

func NewDijkstra(g GraphInterface,s int)*Dijkstra{
	return &Dijkstra{g:g,s:s}
}

func (dij *Dijkstra)initialize(){
	vertexNum := dij.g.VertexNum()
	dij.known = make([]bool,vertexNum)
	dij.path = make([]int,vertexNum)
	for i:=0;i<vertexNum;i++{
		dij.path[i] = -1
	}
	dij.dist = make([]int,vertexNum)
	for i:=0;i<vertexNum;i++{
		dij.dist[i] = math.MaxInt32
	}
	dij.dist[dij.s]=0
}

func (dij *Dijkstra)findMinDistVertex()int{
	min := math.MaxInt32
	min_vertex := -1
	for i:=0;i<dij.g.VertexNum();i++{
		if dij.known[i] == true {continue}
		if dij.dist[i] < min {
			min = dij.dist[i]
			min_vertex = i
		}
	}
	return min_vertex
}

func (dij *Dijkstra)PrintPath(dst int){
	//fmt.Printf("distance from %d to %d:%d\n",dij.start,dst,dij.early[dst])
	if dij.dist[dst] < math.MaxInt32{
		if dij.path[dst] != -1 {
			dij.PrintPath(dij.path[dst])
		}
		fmt.Print(dst," ")
	}else{
		fmt.Println("not path to reach")
	}
}
func (dij *Dijkstra)Run(){
	dij.initialize()
	for i:=0;i<dij.g.VertexNum();i++ {
		v := dij.findMinDistVertex()
		if v==-1{break}
		vertex := dij.g.GetVertex(v)
		for _,tmp:=range vertex.AdjVertexes(){
			w:=tmp.Index()
			cvw:=tmp.Cost()
			if dij.dist[v]+cvw < dij.dist[w]{
				dij.dist[w] = dij.dist[v]+cvw
				dij.path[w] = v
			}
		}
		dij.known[v] = true
	}
}