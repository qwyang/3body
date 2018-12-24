package algorithms

import (
	"math/rand"
	"testing"
)

func checkSort(data []int)bool{
	for i:=len(data)-1;i>0;i--{
		if data[i] < data[i-1]{
			return false
		}
	}
	return true
}
func TestInsertSort1(t *testing.T) {
	data := rand.Perm(10000)
	InsertSort1(data)
	if checkSort(data){
		t.Logf("TestInsertSort1 Success.")
	}else{
		t.Logf("TestInsertSort1 Failed.")
	}
}
func BenchmarkInsertSort1(b *testing.B) {
	data := rand.Perm(100)
	//b.Logf("array:%v\n",data)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		InsertSort1(data)
	}
	b.StopTimer()
	//b.Logf("array:%v\n",data)
}

func BenchmarkInsertSort2(b *testing.B) {
	data := rand.Perm(10000)
	//b.Logf("array:%v\n",data)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		InsertSort1(data)
	}
	b.StopTimer()
	//b.Logf("array:%v\n",data)
}