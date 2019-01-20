package graph

import "container/list"

func TopSort1(g GraphInterface)[]string{
	indegree := make([]int,g.VertexNum())
	for count:=0;count<g.VertexNum();count++{
		vertex:=g.GetVertex(count)
		for _,node := range vertex.AdjVertexes(){
			indegree[node.Index()]++
		}
	}
	//fmt.Println(indegree)
	findNewVertexOfIndegreeZero:=func()int{
		for i,degree := range indegree{
			if degree == 0{
				indegree[i] = -1
				return i
			}
		}
		return -1
	}
	sorted := []string{}
	for count:=0;count<g.VertexNum();count++{
		vertex_index := findNewVertexOfIndegreeZero()
		if vertex_index==-1{
			panic("graph has cycle")
		}
		vertex:=g.GetVertex(vertex_index)
		sorted = append(sorted, vertex.Name())
		for _,node := range vertex.AdjVertexes(){
			indegree[node.Index()]--
		}
	}
	return sorted
}


func TopSort2(g GraphInterface)[]string{
	indegree := make([]int,g.VertexNum())
	for count:=0;count<g.VertexNum();count++{
		vertex:=g.GetVertex(count)
		for _,node := range vertex.AdjVertexes(){
			indegree[node.Index()]++
		}
	}
	sorted := []string{}
	queue := list.New()
	for i,e := range indegree{
		if e == 0{
			queue.PushBack(e)
			indegree[i]--
		}
	}
	for queue.Len()>0{
		vertex_index := queue.Remove(queue.Front()).(int)
		vertex:=g.GetVertex(vertex_index)
		sorted = append(sorted, vertex.Name())
		for _,node := range vertex.AdjVertexes(){
			indegree[node.Index()]--
			if indegree[node.Index()] == 0{
				queue.PushBack(node.Index())
				indegree[node.Index()]--
			}
		}
	}
	if len(sorted) != g.VertexNum() {
		panic("has a cicle")
	}
	return sorted
}