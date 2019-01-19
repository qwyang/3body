package algorithms

import (
	"fmt"
	"sort"
)

/*
归并排序非递归实现,基于抽象的编程方法。
[a,b,c,d,e,f,g]
算法实现：
1.第一次length=1,递增2*length，两两排序
2.第二次length=2
3.最后一次length<n
*/
type MergeSortInterface interface {
	sort.Interface
	Get(int) interface{}
	Assign(int,interface{})
}
func MergeSort(arr,help MergeSortInterface){
	merger:=func(arr MergeSortInterface,i,j,k,l int,help MergeSortInterface){
		fmt.Println("merge ",i,j,k,l)
		start:=i
		count := 0
		for i<=j&&k<=l{
			if arr.Less(i,k){
				help.Assign(count,arr.Get(i))
				i++
			}else{
				help.Assign(count,arr.Get(k))
				k++
			}
			count++
		}
		for i<=j {
			help.Assign(count,arr.Get(i))
			i++
			count++
		}
		for k<=l {
			help.Assign(count,arr.Get(k))
			k++
			count++
		}
		for i:=0;i<count;i++{
			arr.Assign(start+i,help.Get(i))
			fmt.Print(help.Get(i),",")
		}
		fmt.Println()
	}

	for length:=1;length<arr.Len();length=length<<1{
		var i,j,k,l int
		for i=0;i<arr.Len();i+=2*length{
			k,l=i+length,i+2*length
			j=k-1
			if k < arr.Len(){
				if l >= arr.Len(){
					l=arr.Len()-1
					merger(arr,i,j,k,l,help)
				}else{
					merger(arr,i,j,k,l-1,help)
				}
			}
		}
	}
}
