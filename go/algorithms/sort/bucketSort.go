package algorithms
/*
桶排序，对于值在特定小范围内的数据排序可以使用桶排序方法实现线性时间排序完成。
*/
const BucketSortBucketSize = 100
func BucketSort(data []int){
	buckets := make([]int,100)
	for _,d := range data {//登记
		if d >=0 && d<=BucketSortBucketSize {
			buckets[d]++
		}
	}
	j := 0
	for i := range buckets {//串联
		for buckets[i] != 0 {
			data[j] = i
			j++
			buckets[i]--
		}
	}
}