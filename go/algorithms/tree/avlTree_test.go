package algorithms

import "testing"

func TestAvlTree_Insert(t *testing.T) {
	var T *AvlTree
	T = T.Insert(1)
	t.Log(T.Height())
	T = T.Insert(3)
	t.Log(T.Height())
	T = T.Insert(2)
	t.Log(T.Height())
	T = T.Insert(4)
	t.Log(T.Height())
	T = T.Insert(6)
	t.Log(T.Height())
	T = T.Insert(5)
	t.Log(T.Height())
	T.Traverse()
}
