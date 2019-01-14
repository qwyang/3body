package algorithms

import (
	"math/rand"
	"time"
)

// hash 用于计算给定字符串的哈希值的整数形式。
// 本函数实现了BKDR哈希算法。
func hash(str string) uint64 {
	seed := uint64(13131)
	var hash uint64
	for i := 0; i < len(str); i++ {
		hash = hash*seed + uint64(str[i])
	}
	return (hash & 0x7FFFFFFFFFFFFFFF)
}
/*
随机生成一个[min-max）区域的整数
 */
func RandInt(min,max int) int {
	rand.Seed(time.Now().Unix())
	if max <= min || max == 0{
		return max
	}
	return rand.Intn(max-min) + min
}
/*
生成一个大小为n，元素在[min-max）区域的数组
 */
func randArray(n,max int)[]int{
	var data []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i<n; i++ {
		data=append(data,r.Intn(max))
	}
	return data
}
/*
交换一个数组的两个元素
 */
func swap(array []int,i,j int) {
	array[i],array[j] = array[j],array[i]
}
/*
反转一个数组
 */
 func reverse(array []int){
 	size := len(array)
 	for i:=0;i<size/2;i++{
 		swap(array,i,size-i-1)
	}
 }

/*
判断一个数组是否相等
 */
func IsIntArrayEqual(a,b []int) bool{
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

func Max(max int,params... int) int {
	for _,v := range params {
		if v > max {
			max = v
		}
	}
	return max
}