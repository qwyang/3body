package graph

import (
	"container/heap"
	"github.com/3body/go/algorithms/disjset"
)

/*
Kruskal最小生成树算法(比prime算法简单)。
对于所有边，选择最小代价的边(采用优先队列），当前仅当该边的顶点不再已选的SPT上。
*/
/*优先队列实现*/
type Edge struct{s,e,cost int}
type EdgeArr struct{edges[]*Edge}
func (arr *EdgeArr)Len()int{return len(arr.edges)}
func (arr *EdgeArr)Less(i,j int)bool{return arr.edges[i].cost < arr.edges[j].cost}
func (arr *EdgeArr)Swap(i,j int){arr.edges[i],arr.edges[j]=arr.edges[j],arr.edges[i]}
func (arr *EdgeArr)Push(d interface{}){arr.edges = append(arr.edges,d.(*Edge))}
func (arr *EdgeArr)Pop()interface{}{n:=arr.Len();d:=arr.edges[n-1];arr.edges=arr.edges[:n-1];return d}
func Kruskal(g GraphInterface)(spt [][]int){
	spt = make([][]int,g.VertexNum())
	sets := disjset.NewDisjSet(g.VertexNum())
	h := &EdgeArr{}
	for i:=0;i<g.VertexNum();i++{
		v:=g.GetVertex(i)
		for _,n:=range v.AdjVertexes(){
			heap.Push(h,&Edge{i,n.Index(),n.Cost()})
		}
	}
	for h.Len()>0{
		edge:=heap.Pop(h).(*Edge)
		s,e := edge.s,edge.e
		s1,s2:=sets.Find(s),sets.Find(e)
		if s1!=s2{
			spt[s] = append(spt[s],e)
			sets.Union(s1,s2)
		}
	}
	return
}
