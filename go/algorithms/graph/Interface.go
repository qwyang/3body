package graph
type GraphType int
const (
	DG GraphType=1<<iota
	G
)

type AdjNodeInterface interface {
	Cost() int
	Index() int
}

type VertexInterface interface {
	Index() int
	Name() string
	AdjVertexes() []AdjNodeInterface
	AddAdj(AdjNodeInterface)
}

type GraphInterface interface {
	VertexNum() int
	EdegeNum() int
	AddEdge(v1,v2 string,cost int)
	DelEdge(v1,v2 string)
	GetVertex(v int) VertexInterface
	GetVertexIndex(v string)int
}

type GraphInterface2 interface {
	VertexNum() int
	EdegeNum() int
	AddEdge(v1,v2 string,cost int)
	DelEdge(v1,v2 string)
	GetVertexIndex(v string)int
	GetVertexName(index int)string
	GetAdjVertexes(index int)[]int
	GetCost(index1,index2 int)int
}