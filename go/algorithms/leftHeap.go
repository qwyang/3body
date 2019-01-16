package algorithms
/*
左式堆的实现。
堆的性质：根的key最小（小根堆）/最大（大根堆）。
左式堆的性质：左子树的NPL>=右子树NPL,NPL(Null Path Length).
*/
type LeftHeapNode struct {
	lchild *LeftHeapNode
	rchild *LeftHeapNode
	npl int
	key ElemType
}

type ElemType int

func merge(H1,H2 *LeftHeapNode) *LeftHeapNode{
	if H1.lchild == nil{
		H1.lchild = H2
	}else{
		H1.rchild = Merge(H1.rchild,H2)
		if H1.lchild.npl < H1.rchild.npl {
			H1.lchild,H1.rchild = H1.rchild,H1.lchild
		}
		H1.npl = H1.rchild.npl + 1
	}
	return H1
}

func Merge(H1,H2 *LeftHeapNode) *LeftHeapNode{
	if H1 == nil{
		return H2
	}
	if H2 == nil {
		return H1
	}
	if H1.key > H2.key {
		return merge(H2,H1)
	}
	return merge(H1,H2)
}

func (H *LeftHeapNode)Insert(data ElemType)*LeftHeapNode{
	node := &LeftHeapNode{nil,nil,0,data}
	return Merge(H,node)
}

func (H *LeftHeapNode)DeleteMin()(T *LeftHeapNode, data ElemType){
	data = H.key
	T = Merge(H.lchild,H.rchild)
	return T,data
}