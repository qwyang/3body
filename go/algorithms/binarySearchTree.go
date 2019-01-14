package algorithms

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
			if T.rchild != nil{
				T.key = T.rchild.key
				T.rchild = T.rchild.Delete(T.key)
			}else if T.lchild != nil{
				T.key = T.lchild.key
				T.lchild = T.lchild.Delete(T.key)
			}else{//被删元素是叶子节点
				T = nil
			}
		}
	}
	return T
}

func (T *BinarySearchTree) DeleteMin() *BinarySearchTree{
	if T == nil{
		//
	}else if T.lchild == nil{
		T = T.rchild
	}else{
		T.lchild = T.lchild.DeleteMin()
	}
	return T
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
