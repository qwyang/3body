package examples

import "math"

/*
质因数分解
24=2*2*2*3,
input: 24
output:[]int{2,2,2,3}
 */
func decompose(n int)(ret []int){
	var i int
	var sqrt int = int(math.Sqrt(float64(n)))
	for i=2;i<=sqrt;i++{//寻找质因数
		if n % i == 0{
			break
		}
	}
	if i > sqrt{//递归终止条件:已是质数不可分解
		return []int{n}
	}else{//非终止条件，继续递归
		x,y := i,n/i
		ret1:=decompose(x)
		ret2:=decompose(y)
		return append(ret1,ret2...)
	}
}

/*
质数分解非递归算法,求整数n的质因数。
1.在范围i=[2,sqrt(n)]查找n的质因数f
2.找到一个就质因数f就将被除数除以该质因数，n=n/f,没找到i++
3.迭代查找，直到i>sqrt(n)
 */

func decompose2(n int)(ret []int){
	ret = []int{}
	var i int = 2
	var sqrt int = int(math.Sqrt(float64(n)))
	for i <= sqrt{//从最小质因数开始找
		if n % i == 0{//找到一个最小质因数
			ret = append(ret,i)
			n = n/i
			sqrt = int(math.Sqrt(float64(n)))
		}else{
			i++
		}
	}
	ret = append(ret,n)
	return
}