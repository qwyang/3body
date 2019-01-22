package algorithms

func partition(arr []int, s,e int) int {
	pivot := RandInt(s,e)
	v := arr[pivot]
	swap(arr,pivot,e)
	r := s - 1 //r位置元素小于选定值，r+1位置元素大于等于选定值，初始设置为s-1
	for i:=s;i<e;i++{
		if arr[i] < v {
			swap(arr,r+1,i)
			r++
		}
	}
	swap(arr,r+1,e)
	return r+1
}

func qsort(arr []int,s,e int){
	if s >= e {//只有一个元素不用排序，递归终结
		return
	}
	//分割数组后，分别排序
	r := partition(arr,s,e)
	qsort(arr,s,r-1)
	qsort(arr, r+1,e)
}

func Qsort(arr []int){
	if len(arr) <= 1 {
		return
	}
	qsort(arr,0,len(arr) - 1)
}