package examples

import "fmt"

/*
递归求x的n次幂，pow(x,n)=pow(x*x,n/2)
 */
func Pow(x,n int64) int64 {
	if n == 0 {return 1}
	if n & 1 == 0{
		return Pow(x*x,n/2)
	}else{
		return Pow(x*x,n/2)*x
	}
}
/*
非递归求x的n次幂，pow(x,n)=pow(x*x,n/2)
算法：
1.先求出低次幂并保存[x,x^2,x^4...x^n/2]
2.拼接:x^59=x^(32+16+8+2+1)
*/
func PowNr(x,n int64) int64{
	var ret int64 = 1
	var factor []int64
	var f int64 = x
	var count int64 = 1
	for count <= n {//O(logN)
		factor = append(factor,f)
		f = f*f
		count = count << 1
	}
	fmt.Println(factor)
	i := len(factor) - 1
	var index int64
	for n > 0{//O(logN)
		index = 1<<uint(i)
		if n >= index {
			n = n - index
			ret = ret * factor[i]
			fmt.Print(index,",")
		}else {
			i--
		}
	}
	return ret
}