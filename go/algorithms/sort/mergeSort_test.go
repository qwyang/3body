package sort

import (
	"fmt"
	"testing"
)

type IntMergeSort []int

func (arr IntMergeSort)Len()int{return len(arr)}
func (arr IntMergeSort)Swap(i,j int){arr[i],arr[j]=arr[j],arr[i]}
func (arr IntMergeSort)Less(i,j int)bool{return arr[i]<arr[j]}
func (arr IntMergeSort)Get(i int)interface{}{return arr[i]}
func (arr IntMergeSort)Assign(i int,data interface{}){arr[i]=data.(int)}

func TestMergeSort(t *testing.T){
	data := randArray(10,100)
	fmt.Println(data)
	help := make([]int,len(data))
	MergeSort(IntMergeSort(data),IntMergeSort(help))
	fmt.Println(data)
}
