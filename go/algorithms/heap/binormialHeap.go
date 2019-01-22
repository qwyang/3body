package algorithms

import (
	"container/list"
	"errors"
	"fmt"
	"math"
)

/*
优先队列：二项堆的实现
*/
type BinormialQueue struct {
	Trees []*BinormialTree
	CurSize int
}

type BinormialTree struct {
	lchild *BinormialTree
	sibling *BinormialTree
	data int
}
func CombineBinormialTree(t1,t2 *BinormialTree) *BinormialTree{
	if t1.data > t2.data {
		t1,t2 = t2,t1
	}
	t2.sibling = t1.lchild
	t1.lchild = t2
	return t1
}
func MergeBinQueue(H1,H2 *BinormialQueue)*BinormialQueue{
	if H1 == nil {
		return H2
	}
	if H2 == nil{
		return H1
	}
	H1.CurSize = H1.CurSize + H2.CurSize
	length:=len(H1.Trees)
	for s:= bits(H1.CurSize);s>length;s--{
		H1.Trees = append(H1.Trees,nil)
		//fmt.Println("expand")
	}
	var t1,t2,carry *BinormialTree
	var flag int
	for i,j:=0,1;j<=H1.CurSize;i,j=i+1,j<<1{
		flag = 0
		t1,t2 = nil,nil
		if i < len(H1.Trees) && H1.Trees[i]!=nil{
			t1 = H1.Trees[i]
			flag += 1
		}
		if i < len(H2.Trees) && H2.Trees[i]!=nil{
			t2 = H2.Trees[i]
			flag += 2
		}
		if carry != nil {
			flag+=4
		}
		switch flag {
		case 0:
		case 1:
		case 2:
			H1.Trees[i] = t2
			H2.Trees[i] = nil
		case 3:
			//fmt.Println(3,t1,t2)
			carry = CombineBinormialTree(t1,t2)
			H1.Trees[i],H2.Trees[i] = nil,nil
		case 4:
			H1.Trees[i] = carry
			carry = nil
		case 5:
			//fmt.Println(5,t1,carry)
			carry = CombineBinormialTree(t1,carry)
			H1.Trees[i] = nil
		case 6:
			carry = CombineBinormialTree(t2,carry)
			H2.Trees[i] = nil
		case 7:
			H1.Trees[i] = carry
			carry = CombineBinormialTree(t1,t2)
			H2.Trees[i] = nil
		}
	}
	H1.Traverse()
	return H1
}

func NewBinorminalQueue(data int)*BinormialQueue{
	new := &BinormialTree{nil,nil,data,}
	return &BinormialQueue{[]*BinormialTree{new},1}
}

func (H *BinormialQueue)Insert(data int)*BinormialQueue{
	if H == nil{
		return NewBinorminalQueue(data)
	}
	return MergeBinQueue(H,NewBinorminalQueue(data))
}

func (H *BinormialQueue)DeleteMin()(minItem int,err error){
	if H == nil || len(H.Trees) == 0{
		return 0,errors.New("empty queue")
	}
	minItem = math.MaxInt32
	var minTreeIndex = -1
	var minTree *BinormialTree
	for i,t := range H.Trees{//O(logN)
		if t!=nil && t.data < minItem {
			minTree = t
			minTreeIndex = i
			minItem = t.data
		}
	}
	deleteTree := minTree.lchild
	deleteBinormialQueue := &BinormialQueue{make([]*BinormialTree,minTreeIndex),1<<uint(minTreeIndex)-1}
	for j:=minTreeIndex-1;j>=0;j--{
		deleteBinormialQueue.Trees[j] = deleteTree
		deleteTree = deleteTree.sibling
		deleteBinormialQueue.Trees[j].sibling = nil
	}
	H.Trees[minTreeIndex] = nil
	H.CurSize = H.CurSize - 1<<uint(minTreeIndex)
	H.Trees = H.Trees[:bits(H.CurSize)]
	H = MergeBinQueue(H,deleteBinormialQueue)
	return
}

func levelTraverse(T *BinormialTree){
	if T!=nil{
		queue := list.New()
		queue.PushBack(T)
		for queue.Len() > 0 {
			T = queue.Remove(queue.Front()).(*BinormialTree)
			fmt.Println(T.data,",")
			if T.lchild != nil{
				queue.PushBack(T.lchild)
			}
			if T.sibling != nil{
				queue.PushBack(T.sibling)
			}
		}
	}
}
func (H *BinormialQueue)Traverse(){
	if H!=nil{
		fmt.Println("total:",H.CurSize)
		for i,T := range H.Trees{
			if T!=nil{
				fmt.Printf("traverse T.Trees[%d]\n",i)
				levelTraverse(T)
				fmt.Println()
			}
		}
	}
}
