package sort

type SortInterface interface{
	Less(i,j int) bool
	Swap(i,j int)
	Len() int
}

func partition2(obj SortInterface, s,e int) int {
	//r位置元素小于选定值，r+1位置元素大于等于选定值，初始设置为s-1
	r := s - 1
	for i:=s;i<e;i++{
		if obj.Less(i,e) {
			obj.Swap(r+1,i)
			r++
		}
	}
	obj.Swap(r+1,e)
	return r+1
}

func qsort2(obj SortInterface,s,e int){
	if s >= e {//只有一个元素不用排序，递归终结
		return
	}
	//分割数组后，分别排序
	r := partition2(obj,s,e)
	qsort2(obj,0,r-1)
	qsort2(obj, r+1,e)
}

func Qsort2(obj SortInterface){
	if obj.Len() <= 1 {
		return
	}
	qsort2(obj,0,obj.Len() - 1)
}