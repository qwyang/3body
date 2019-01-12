package algorithms

import (
	"sort"
	"testing"
)


func TestRandInt(t *testing.T) {
	min := 5
	max := 10
	r := RandInt(min,max)
	if r >= min && r < max {
		t.Log("RandInt success")
	}else{
		t.FailNow()
	}
	t.Log("random:",r)
}

func TestQsort(t *testing.T){
	//normal case
	test := []int{3,2,1,4,5}
	expect := []int{1,2,3,4,5}
	Qsort([]int(test))
	if !IsIntArrayEqual(test,expect) {
		t.Log(test)
		t.FailNow()
	}
	t.Log(test)
	t.Log("TestQsort1 success")
	//已排序数组
	test = []int{1,1,1,2,2,2}
	expect = []int{1,1,1,2,2,2}
	Qsort([]int(test))
	if !IsIntArrayEqual(test,expect) {
		t.Log(test)
		t.FailNow()
	}
	t.Log(test)
	t.Log("TestQsort2 success")
	//已排序数组，倒序
	test = []int{2,2,2,1,1,1}
	expect = []int{1,1,1,2,2,2}
	Qsort([]int(test))
	if !IsIntArrayEqual(test,expect) {
		t.Log(test)
		t.FailNow()
	}
	t.Log(test)
	t.Log("TestQsort3 success")
	//空数组，倒序
	test = nil
	expect = nil
	Qsort([]int(test))
	if !IsIntArrayEqual(test,expect) {
		t.Log(test)
		t.FailNow()
	}
	t.Log(test)
	t.Log("TestQsort4 success")
	array := randArray(100,1000)
	expect = make([]int,100)
	copy(expect,array)
	sort.Ints(expect)
	Qsort(array)
	if !IsIntArrayEqual(array,expect) {
		t.Log(array)
		t.Log(expect)
		t.FailNow()
	}
	t.Log(array)
	t.Log("TestQsort5 success")
}
func testQSort(n int,b *testing.B){
	data := randArray(n,1<<30)
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i:=0;i<b.N;i++{
		Qsort(data)
	}
	b.Logf("benchmark for qsort N=%d done\n",n)
}
func BenchmarkQSort1024(b *testing.B){
	testRadixSort(1024,b)
}

func BenchmarkQSort2048(b *testing.B){
	testRadixSort(2048,b)
}

func BenchmarkQSort4096(b *testing.B){
	testRadixSort(4096,b)
}

func BenchmarkQSort8192(b *testing.B){
	testRadixSort(8192,b)
}

func BenchmarkQSortLarge(b *testing.B){
	testRadixSort(1<<20,b)
}