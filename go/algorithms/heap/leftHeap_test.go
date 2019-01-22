package algorithms

import "testing"

func TestLeftHeap(t *testing.T) {
	var H *LeftHeapNode
	for i:=0;i<10;i++{
		H = H.Insert(ElemType(i))
		t.Logf("%+v",H)
	}
	var e ElemType
	for i:=0;i<10;i++{
		H,e = H.DeleteMin()
		t.Logf("%+v,%v",H,e)
	}
}