package algorithms

import (
	"math"
)

/*
矩阵乘法的顺序安排。
存放矩阵大小的数组：A[50x10]*B[10x40]*C[40x30]*D[30x5]=C[5,10,40,30,5]。
乘法次数：M[1][N]=Min(M[1][k]+M[k][N]+C[0]*C[k]*C[N])。
*/
func OptMatrix(C []int)(M,L [][]int){
	N:=len(C)-1
	M=make([][]int,N+1)
	L=make([][]int,N+1)
	for i := range M{
		M[i] = make([]int,N+1)
		L[i] = make([]int,N+1)
	}
	for i:=1;i<=N;i++{
		for j:=i;j<=N;j++{
			if i==j{
				M[i][j]=0
			}else{
				M[i][j]=math.MaxInt32
			}
		}
	}
	for d:=1;d<N;d++{
		for Left:=1;Left<=N-d;Left++{
			Right:=Left+d
			for k:=Left;k<Right;k++{
				count := M[Left][k]+M[k+1][Right]+C[Left-1]*C[k]*C[Right]
				if count < M[Left][Right]{
					//fmt.Println(count,Left,Right)
					M[Left][Right]=count
					L[Left][Right]=k
				}
			}
		}
	}
	return
}