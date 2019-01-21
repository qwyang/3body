package graph

type IndexGen struct {
	next int
}

func NewIndexGen()*IndexGen{
	return &IndexGen{}
}

func (gen *IndexGen)Next()int{
	alloc := gen.next
	gen.next++
	return alloc
}

type AdjacencyList struct {
	vertexNum int
	edgeNum int
	vertexArr []VertexInterface
	nameIndexConverter map[string]int
	indexGen *IndexGen
	graphType GraphType
}
func NewAdjacencyList(t GraphType)*AdjacencyList{
	return &AdjacencyList{graphType:t,indexGen:NewIndexGen(),nameIndexConverter:make(map[string]int)}
}
func (g *AdjacencyList)Type() GraphType{return g.graphType}
func (g *AdjacencyList)VertexNum() int{return g.vertexNum}
func (g *AdjacencyList)EdegeNum() int{return g.edgeNum}
func (g *AdjacencyList)AddEdge(v1,v2 string, cost int){
	if g.GetVertexIndex(v1) == -1{
		index := g.indexGen.Next()
		g.vertexArr = append(g.vertexArr,&Vertex{v1,index,nil})
		g.nameIndexConverter[v1] = index
		g.vertexNum++
	}
	if g.GetVertexIndex(v2) == -1{
		index := g.indexGen.Next()
		g.vertexArr = append(g.vertexArr,&Vertex{v2,index,nil})
		g.nameIndexConverter[v2] = index
		g.vertexNum++
	}
	index1,index2 := g.GetVertexIndex(v1),g.GetVertexIndex(v2)
	//fmt.Println(index1,index2)
	g.GetVertex(index1).AddAdj(&AdjacentVertex{cost,index2})
	g.edgeNum++
	if g.graphType == G {
		g.GetVertex(index2).AddAdj(&AdjacentVertex{cost,index1})
	}
}
func (g *AdjacencyList)DelEdge(v1,v2 string){

}
func (g *AdjacencyList)GetVertex(index int) VertexInterface{
	if index >= g.vertexNum||index < 0{return nil}
	return g.vertexArr[index]
}

func (g *AdjacencyList)GetVertexIndex(v string)int{
	index,ok := g.nameIndexConverter[v]
	if ok{
		return index
	}
	return -1
}

type AdjacentVertex struct {
	cost int
	index int
}

func (n *AdjacentVertex)Cost() int{return n.cost}
func (n *AdjacentVertex)Index() int{return n.index}

type Vertex struct {
	name  string
	index int
	adj  []AdjNodeInterface
}

func (v *Vertex) Index() int{return v.index}
func (v *Vertex) Name() string{return v.name}
func (v *Vertex) AddAdj(n AdjNodeInterface){ v.adj=append(v.adj,n )}
func (v *Vertex) AdjVertexes()[]AdjNodeInterface{
	return v.adj
}