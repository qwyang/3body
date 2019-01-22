package heap

import "testing"

func TestBinormialHeap(t *testing.T) {
	var H *BinormialQueue
	for i:=0;i<9;i++{
		H = H.Insert(i)
	}

	for i:=0;i<10;i++{
		H,e := H.DeleteMin()
		t.Logf("%+v,%v",H,e)
	}
}
