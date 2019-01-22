package sort

import (
	"sort"
	"testing"
)


func TestBucketSort(t *testing.T) {
	data := randArray(100,100)
	expect := make([]int,len(data))
	copy(expect,data)
	sort.Ints(expect)
	BucketSort(data)
	if !IsIntArrayEqual(data,expect){
		t.Fatal(data,"not expected.")
	}else{
		t.Log("Passed,data=",data)
	}
}
