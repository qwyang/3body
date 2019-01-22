package sort

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

func testRadixSort(n int,b *testing.B){
	data := randArray(n,1<<30)
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i:=0;i<b.N;i++{
		radixSort(data)
	}
	b.Logf("benchmark for randixsort N=%d done\n",n)
}

func BenchmarkRadixSort1024(b *testing.B){
	testRadixSort(1024,b)
}

func BenchmarkRadixSort2048(b *testing.B){
	testRadixSort(2048,b)
}

func BenchmarkRadixSort4096(b *testing.B){
	testRadixSort(4096,b)
}

func BenchmarkRadixSort8192(b *testing.B){
	testRadixSort(8192,b)
}

func BenchmarkRadixSortLarge(b *testing.B){
	testRadixSort(1<<20,b)
}