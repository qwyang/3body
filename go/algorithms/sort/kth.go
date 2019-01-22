package sort

import "fmt"

/*
在数组array[0-size)查找第k大的元素,数组初始为无序状态。
算法：随机选取一个元素将数组分成两部分子数组，所选元素为第r大。
1.如果r==k,则返回
2.如果r in [k+1-size]，则在分割出来的元素较小的数组继续查找第k大的元素。
3.如果r in [1-k-1]，则在分割出来的元素较大的数组继续寻找k-x大的元素
算法复杂度为O(n)=n+n/2+n/4+...
 */
func divide(array []int,count *int)int{
	oldcount := *count
	p := RandInt(0,len(array))
	pivot := array[p]
	array[p],array[len(array)-1] = array[len(array)-1],array[p]
	r := -1//array[0-r]>=pivot,array[r+1,size-1]<pivot
	//for i:=0;i<len(array);i++{//错误，最后一个数字不要加入比较
	for i:=0;i<len(array)-1;i++{
		if array[i] >= pivot {//大元素往前调,小元素往后
			array[r+1],array[i] = array[i],array[r+1]
			r++
		}
		*count++
	}
	array[r+1],array[len(array)-1] = array[len(array)-1],array[r+1]
	fmt.Printf("r:%d,count:%d\n",r,*count-oldcount)
	return r+1
}
var counter int
func FindKth(array []int,k int) (int,int) {
	r := divide(array,&counter)
	if r+1 == k {
		return array[r],counter
	}else if r+1 > k{
		array = array[0:r]
		return FindKth(array,k)
	}else{
		array = array[r+1:]
		return FindKth(array,k-r-1)
	}
}