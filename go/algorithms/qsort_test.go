package algorithms

import "testing"


type IntArray []int

func (a IntArray) equal(b IntArray) bool{
	if len(a) != len(b){
		return false
	}
	if (a == nil) != (b == nil){
		return false
	}
	for i,v := range a {
		if v != b[i]{
			return false
		}
	}
	return true
}

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
	test := IntArray{3,2,1,4,5}
	expect := IntArray{1,2,3,4,5}
	Qsort([]int(test))
	if !test.equal(expect) {
		t.Log(test)
		t.FailNow()
	}
	t.Log(test)
	t.Log("TestQsort1 success")
	//已排序数组
	test = IntArray{1,1,1,2,2,2}
	expect = IntArray{1,1,1,2,2,2}
	Qsort([]int(test))
	if !test.equal(expect) {
		t.Log(test)
		t.FailNow()
	}
	t.Log(test)
	t.Log("TestQsort2 success")
	//已排序数组，倒序
	test = IntArray{2,2,2,1,1,1}
	expect = IntArray{1,1,1,2,2,2}
	Qsort([]int(test))
	if !test.equal(expect) {
		t.Log(test)
		t.FailNow()
	}
	t.Log(test)
	t.Log("TestQsort3 success")
	//空数组，倒序
	test = nil
	expect = nil
	Qsort([]int(test))
	if !test.equal(expect) {
		t.Log(test)
		t.FailNow()
	}
	t.Log(test)
	t.Log("TestQsort4 success")
}