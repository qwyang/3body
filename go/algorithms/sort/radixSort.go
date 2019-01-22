package algorithms

import (
	"container/list"
)

/*
基数排序算法，基数radix:F，数据量：N，数据位数d：8
算法复杂度:O(d(N+N)),属于分配算法。
 */
func radixSort(data []int){
	var radix int = 0x10
	digitNum := 8
	buckets := make([]*list.List,radix)
	for i:=0;i<len(buckets);i++{
		buckets[i] = list.New()
	}
	for i:=0;i<digitNum;i++{
		//radix = 0xF << 4*i
		radix = 0xF << uint(4*i)
		//fmt.Println("??",radix)
		for _,d := range data {//分配
			pos := (d & radix)>> uint(4*i)
			//fmt.Println("##",radix,pos)
			buckets[pos].PushBack(d)
		}
		j:=0
		for _,l := range buckets {//回收
			for l.Len() > 0 {
				data[j] = l.Remove(l.Front()).(int)
				j++
			}
		}
		//fmt.Println(data)
	}
}