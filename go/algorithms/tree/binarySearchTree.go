package tree

import "fmt"

/*
二叉搜索树的实现
*/
type BinarySearchTree struct{
	lchild,rchild *BinarySearchTree
	key int /*节点保存的数据*/
	count int
	height int /*高度*/
}

func NewBinaryTreeNode(key int) (node *BinarySearchTree) {
	node = &BinarySearchTree{nil,nil,key,1,0}
	return
}

func (T *BinarySearchTree)Height()int{
	if T == nil{
		return -1
	}else{
		return T.height
	}
}
/*
插入值同时更新数高和计数器。
 */
func (T *BinarySearchTree)Insert(key int) *BinarySearchTree {
	if T == nil {
		T = NewBinaryTreeNode(key)
	}else if key < T.key {
		T.lchild = T.lchild.Insert(key)
	}else if key > T.key {
		T.rchild = T.rchild.Insert(key)
	}else{
		T.count++
	}
	T.height = Max(T.lchild.Height(),T.rchild.Height()) + 1
	return T
}
func (T *BinarySearchTree)Delete(key int)*BinarySearchTree{
	if T == nil{
		//
	}else if key < T.key{
		T.lchild = T.lchild.Delete(key)
	}else if key > T.key{
		T.rchild = T.rchild.Delete(key)
	}else{//相等的情况
		T.count--
		if T.count == 0{
			/*有右子树*/
			if T.rchild != nil{
				q := T //q代表最小节点的父节点
				p := T.rchild //p代表最小节点,最小节点不可能有左子树
				for p.lchild != nil {
					q = p
					p = p.lchild
				}
				if p == q.lchild {
					q.lchild = p.rchild
				}else{
					q.rchild = p.rchild
				}
				T.key,T.count = p.key,p.count
			}else{//只有左子树,或者左子树为空
				T = T.lchild
			}
		}
	}
	return T
}

func (T *BinarySearchTree) DeleteMin() (*BinarySearchTree,int){
	var min int
	if T == nil{
		//
	}else if T.lchild == nil{
		T.count--
		min = T.key
		if T.count == 0 {
			T = T.rchild
		}
	}else{
		T.lchild,min = T.lchild.DeleteMin()
	}
	return T,min
}

/*
中序遍历
 */
func (T *BinarySearchTree)Traverse(){
	if T == nil{
		return
	}
	T.lchild.Traverse()
	for i:=0;i<T.count;i++{
		fmt.Print(T.key," ")
	}
	T.rchild.Traverse()
}
