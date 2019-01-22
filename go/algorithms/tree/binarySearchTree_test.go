package tree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree_Insert(t *testing.T) {
	data := randArray(10,10)
	fmt.Println(data)
	var T *BinarySearchTree
	/*case2:insert 10 elem*/
	for _,v := range data {
		T = T.Insert(v)
	}
	T.Traverse()
	fmt.Println()
	/*case2:delete elem*/
	for i:=0;i<10;i++ {
		T = T.Delete(i)
	}
	T.Traverse()
	var e int
	/*case3:deleteMin 10 times*/
	for i:=0;i<10;i++ {
		T,e = T.DeleteMin()
		t.Log("delete:",e)
	}
	T.Traverse()
}