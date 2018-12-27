package main

import "fmt"
/*
切片和原切片/数组是共享空间的。
 */
func main(){
	buf := make([]byte,20)
	buf2 := buf[:2]
	buf2[0]='a'
	fmt.Println(buf)
	fmt.Println(buf2)
	array := [3]int{1,2,3}
	arr := array[:1]
	arr[0] = 3
	fmt.Println(array)
	fmt.Println(arr)
}
