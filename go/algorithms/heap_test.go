package algorithms

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	pq := NewPQueue()
	for i:=0;i<10;i++{
		elem := NewPQelem(i,i)
		pq.Insert(elem)
		if pq.Size() != int64(i+1) {
			t.Fatal("pq.size should = ",i+1)
		}
	}
	//pq.(*pqueue).Traverse()
	for i:=0;i<10;i++{
		elem,ok := pq.DeleteMax()
		if ok {
			//pq.(*pqueue).Traverse()
			if pq.Size() != int64(10-i-1) {
				t.Fatal("pq.size should = ",10-i-1)
			}
			fmt.Println(elem.Priority())
			if elem.Priority() != 10-i-1 {
				t.Fatal("pq.max should = ",10-i-1)
			}
		}
	}
}
