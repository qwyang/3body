package heap

import (
	"fmt"
	"math"
)
/*
插入和删除只需要log2N次方
*/
type PQueue interface {
	Size() int64
	Max() PqElem
	Insert(PqElem) bool
	DeleteMax() (PqElem,bool)
}

type PqElem interface {
	SetPriority(int) bool
	Priority() int
	Value() interface{}
}

type pqelem struct {
	priority int
	value interface{}
}

func NewPQelem(p int,value interface{})*pqelem{
	return &pqelem{p,value}
}

func (e *pqelem) SetPriority(p int) bool{
	e.priority = p
	return true
}

func (e *pqelem) Priority() int{
	return e.priority
}

func (e *pqelem) Value() interface{}{
	return e.value
}

type pqueue struct {
	array []PqElem
}

func (pq *pqueue) Size() int64 {
	return int64(len(pq.array)-1)
}

func (pq *pqueue) Max() PqElem {
	return pq.array[1]
}

func (pq *pqueue) Insert(elem PqElem) bool {
	if elem == nil{	return false}
	position := pq.Size()+1
	pq.array = append(pq.array,elem)
	parent := position/2
	//compare and swap
	for parent > 0 && pq.array[parent].Priority() < pq.array[position].Priority(){
		pq.array[parent],pq.array[position] = pq.array[position],pq.array[parent]
		position = parent
		parent = parent/2
	}
	return true
}

func (pq *pqueue) justify(parent int64){
	for parent < pq.Size() {
		lchild := 2*parent
		rchild := lchild + 1
		var p,l,r int
		p=pq.array[parent].Priority()
		if lchild > pq.Size(){
			l = math.MinInt32
		}else{
			l=pq.array[lchild].Priority()
		}
		if rchild > pq.Size() {
			r = math.MinInt32
		}else{
			r = pq.array[rchild].Priority()
		}

		if p >= l && p >= r{
			break
		}else if l >= r && l > p {
			//fmt.Printf("child priority bigger,parent_index:%d,priority:%d,child_index:%d,priority:%d\n",parent,p,lchild,l)
			pq.array[parent],pq.array[lchild] = pq.array[lchild],pq.array[parent]
			parent = lchild
			//pq.Traverse()
		}else{
			//fmt.Printf("child priority bigger,parent_index:%d,priority:%d,child_index:%d,priority:%d\n",parent,p,rchild,r)
			pq.array[parent],pq.array[rchild] = pq.array[rchild],pq.array[parent]
			parent = rchild
			//pq.Traverse()
		}
	}
}

func (pq *pqueue) DeleteMax() (PqElem,bool){
	size := pq.Size()
	if size <= 0 {
		return nil,false
	}
	max := pq.Max()
	pq.array[1] = pq.array[size]
	pq.array = pq.array[:size]
	//justify
	var parent int64 = 1
	pq.justify(parent)
	return max,true
}

func (pq *pqueue)Traverse(){
	p := []int{}
	for _,elem := range pq.array {
		p = append(p,elem.Priority())
	}
	fmt.Println(p)
}

func NewPQueue() PQueue {
	pq := &pqueue{}
	pq.array = append(pq.array,&pqelem{})
	return pq
}

