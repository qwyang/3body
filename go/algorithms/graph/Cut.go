package graph

import "fmt"

/*
割点：对于无向联通图，如果删除掉v则不在联通，则v是图的一个割点。
算法：
1.深度遍历，给每个顶点分配编号(num)，记录顶点的父顶点(parent)。
2.递归查找每个顶点所能访问到的最小编号顶点(low),如果某个孩子的low>本节点的num，则本届点是割点。
low(v)的定义：
1.num(v).
2.所有反向边(v,w)的最小num(w).
3.树的所有边(v,w)的最低low(w).
*/
type Cut struct {
	Num []int
	Parent []int
	known []bool
	Low []int
	counter int
	G GraphInterface
}

func NewCut(g GraphInterface)*Cut{
	c := &Cut{}
	c.Num = make([]int,g.VertexNum())
	c.Parent = make([]int,g.VertexNum())
	c.known = make([]bool,g.VertexNum())
	c.Low = make([]int,g.VertexNum())
	c.G = g
	return c
}

func (c *Cut)AssignNum(s int){//O(E)
	c.known[s]=true
	c.counter++
	c.Num[s] = c.counter
	v := c.G.GetVertex(s)
	for _,n:=range v.AdjVertexes(){
		e := n.Index()
		if !c.known[e]{
			c.Parent[e] = s
			c.AssignNum(e)
		}
	}
}
func (c *Cut)AssignLow(s int){
	c.Low[s] = c.Num[s]
	v := c.G.GetVertex(s)
	for _,n:=range v.AdjVertexes(){
		w := n.Index()
		if c.Num[w] > c.Num[s]{
			c.AssignLow(w)
			if c.Low[w] < c.Low[s]{
				c.Low[s] = c.Low[w]
			}
			if c.Low[w] >= c.Num[s]{
				fmt.Printf("%d is a cut point.\n",s)
			}
		}else{
			if c.Parent[w]!=s{//背向边
				if c.Num[w] < c.Low[s] {
					c.Low[s] = c.Num[w]
				}
			}
		}
	}
}
