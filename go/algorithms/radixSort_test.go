package algorithms

import (
	"sort"
	"testing"
)

func TestRadixSort(t *testing.T){
	data := randArray(10,1<<10)
	expect := make([]int,len(data))
	copy(expect,data)
	sort.Ints(expect)
	radixSort(data)
	if !IsIntArrayEqual(data,expect){
		t.Fatal("not expected,data:",data,"expect:",expect)
	}else{
		t.Log("success,data:",data)
	}
}

