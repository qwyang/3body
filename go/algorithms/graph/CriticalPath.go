package graph

import (
	"container/list"
	"fmt"
	"math"
)

/*
关键路径问题。
项目最早结束时间(最长路径问题)，各个活动最晚结束时间，各个活动的松弛时间。
*/

type CriticalPath struct{
	g               GraphInterface
	start           int
	end				int
	preVertex, cost [][]int
	path            []int//关键路径
	early           []int//最早完成时间，按计划时间完成的最好的预期
	latest          []int//最晚完成时间，最差的预期
	slack           []int//每个任务不影响最后结果的可以推迟的时间
}

func NewCriticalPath(g GraphInterface,s,e int)*CriticalPath {
	new := &CriticalPath{
		g: g,
		start:s,
		end:e,
		preVertex:make([][]int,g.VertexNum()),
		cost:make([][]int,g.VertexNum()),
	}
	vertexNum := new.g.VertexNum()
	new.path = make([]int,vertexNum)
	for i:=0;i<vertexNum;i++{
		new.path[i] = -1
	}
	new.early = make([]int,vertexNum)
	for i:=0;i<vertexNum;i++{
		new.early[i] = math.MinInt32
	}
	new.latest = make([]int,vertexNum)
	for i:=0;i<vertexNum;i++{
		new.latest[i] = math.MaxInt32
	}
	new.slack = make([]int,vertexNum)
	new.early[new.start]=0
	return new
}

func (dij *CriticalPath)PrintResult(){
	for i:=0;i<dij.g.VertexNum();i++{
		fmt.Printf("earliest finished time for %d:%d\n",i,dij.early[i])
	}
	for i:=0;i<dij.g.VertexNum();i++{
		fmt.Printf("latest finished time for %d:%d\n",i,dij.latest[i])
	}
	for i:=0;i<dij.g.VertexNum();i++{
		fmt.Printf("slacked time for %d:%d\n",i,dij.slack[i])
	}
}

func (dij *CriticalPath)PrintPath(dst int){
	if dij.early[dst] > math.MinInt32{
		if dij.path[dst] != -1 {
			dij.PrintPath(dij.path[dst])
		}
		fmt.Print(dst," ")
	}else{
		fmt.Println("not path to reach")
	}
}
/*
假设是无圈图，否则进入死循环。
 */
func (dij *CriticalPath)Run(){
	queue := list.New()
	queue.PushBack(dij.start)
	for queue.Len()>0 {//复杂度：O(E)
		v := queue.Remove(queue.Front()).(int)
		vertex := dij.g.GetVertex(v)
		for _, edge :=range vertex.AdjVertexes(){
			w:= edge.Index()
			cvw:= edge.Cost()
			if dij.early[v]+cvw > dij.early[w]{
				dij.early[w] = dij.early[v]+cvw
				dij.path[w] = v
			}
			dij.preVertex[w]=append(dij.preVertex[w], v)
			dij.cost[w]=append(dij.cost[w],cvw)
			queue.PushBack(w)
		}
	}
	dij.latest[dij.end] = dij.early[dij.end]
	dij.slack[dij.end] = 0
	queue.PushBack(dij.end)
	for queue.Len() > 0{
		w := queue.Remove(queue.Front()).(int)
		for i,v := range dij.preVertex[w] {
			cvw := dij.cost[w][i]
			if dij.latest[w] - cvw < dij.latest[v]{
				dij.latest[v] = dij.latest[w] - cvw
				dij.slack[v] = dij.latest[v] - dij.early[v]
			}
			queue.PushBack(v)
		}
	}
}