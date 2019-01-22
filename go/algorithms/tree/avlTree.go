package tree

import "fmt"

/*
avl tree的实现
*/
type AvlTree struct {
	lchild *AvlTree
	rchild *AvlTree
	height int
	key int
	count int
}

func NewAvlTreeNode(key int)*AvlTree{
	return &AvlTree{nil,nil,0,key,1}
}
/*k1<k2*/
func (T *AvlTree)RotateWithLeft()*AvlTree{
	k2,k1 := T,T.lchild
	k2.lchild = k1.rchild
	k1.rchild = k2
	k2.height = Max(k2.lchild.Height(),k2.rchild.Height())+1
	k1.height = Max(k1.lchild.Height(),k1.rchild.Height())+1
	return k1
}
func (T *AvlTree)RotateWithRight()*AvlTree{
	k1,k2 := T,T.rchild
	k1.rchild = k2.lchild
	k2.lchild = k1
	k1.height = Max(k1.lchild.Height(),k1.rchild.Height())+1
	k2.height = Max(k2.lchild.Height(),k2.rchild.Height())+1
	return k2
}
func (T *AvlTree)RotateWithRightLeft()*AvlTree{
	k3:=T
	k1:=T.lchild
	k3.lchild = k1.RotateWithRight()
	return k3.RotateWithLeft()
}
func (T *AvlTree)RotateWithLeftRight()*AvlTree{
	k1,k3 := T,T.rchild
	k1.rchild = k3.RotateWithLeft()
	return k1.RotateWithRight()
}
func (T *AvlTree)Height()int{
	if T == nil{
		return -1
	}
	return T.height
}
func (T *AvlTree)Insert(key int)*AvlTree{
	if T == nil {
		T = NewAvlTreeNode(key)
	}else if key < T.key {
		T.lchild = T.lchild.Insert(key)
		T.height = Max(T.lchild.Height(),T.rchild.Height())+1
		if T.lchild.Height() - T.rchild.Height() > 1 {
			//todo:调整
			if key < T.lchild.key {
				T = T.RotateWithLeft()
			}
			if key > T.lchild.key {
				T = T.RotateWithRightLeft()
			}
		}
	}else if key > T.key {
		T.rchild = T.rchild.Insert(key)
		T.height = Max(T.lchild.Height(),T.rchild.Height())+1
		if T.rchild.Height() - T.lchild.Height() > 1 {
			//todo:调整
			if key > T.rchild.key {
				T = T.RotateWithRight()
			}
			if key < T.rchild.key {
				T = T.RotateWithLeftRight()
			}
		}
	}else{
		T.count++
	}
	return T
}

/*
中序遍历
 */
func (T *AvlTree)Traverse(){
	if T == nil{
		return
	}
	T.lchild.Traverse()
	for i:=0;i<T.count;i++{
		fmt.Print(T.key," ")
	}
	T.rchild.Traverse()
}