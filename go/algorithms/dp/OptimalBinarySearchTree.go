package dp

import (
	"math"
)

/*
动态规划：最优二叉查找树。
测试用例：
P[0.22,0.18,0.20,0.05,0.25,0.02,0.08]
W[a,am,and,egg,if,the,two]
C(i,j)=Min(C(i,k-1)+C(k+1,j)+P(i,j);其中：i<=k<=j;
最终求：C(1,N)
*/
func OptimalBinarySearchTree(P []float32)(C[][]float32,R[][]int){
	n:=len(P)
	C = make([][]float32,n)
	R = make([][]int,n)
	for i:=0;i<n;i++{
		C[i] = make([]float32,n)
		R[i] = make([]int,n)
	}
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			if i==j{
				C[i][i]=P[i]
				R[i][i]=i
			}else {
				C[i][j]=math.MaxFloat32
			}
		}
	}
	sum:=func(i,j int)(s float32){
		for i<=j{
			s+=P[i]
			i++
		}
		return
	}
	for d:=1;d<n;d++{
		for i:=0;i<n-d;i++{
			j:=i+d
			for k:=i;k<=j;k++{
				var cl,cr float32
				if k>i {
					cl=C[i][k-1]
				}
				if k<j{
					cr=C[k+1][j]
				}
				cij:=cl+cr+sum(i,j)
				if cij < C[i][j]{
					C[i][j]=cij
					R[i][j]=k
				}
			}
		}
	}
	return
}
