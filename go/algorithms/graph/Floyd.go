package graph

import (
	"fmt"
	"math"
)

/*
弗洛伊德算法：计算图中两点之间最近的距离和路径。
原理：
1.用V1去更新其他路径时，所有(路径长<=2)且(经过V1)且（比原路径长度<=1）花费小的所有路径都会更新。
2.用V2去更新其他路径时，所有(路径长<=3)且(经过V1，V2)且（比原路径长度<=2）花费小的所有路径都会更新。
k.用Vk去更新其他路径时，所有(路径长<=k)且(经过V1，V2..Vk)且比（原路径长度<=k-1）花费小的所有路径都会更新。
 */
func Floyd(G [][]int)(cost,path[][]int){
	n:=len(G)
	cost = make([][]int,n)
	path = make([][]int,n)
	for i:=0;i<len(G);i++{
		cost[i] = make([]int,n)
		path[i] = make([]int,n)
	}
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			cost[i][j]=G[i][j]
			if cost[i][j] < math.MaxInt32{
				path[i][j]=i
			}
		}
	}
	for v:=0;v<n;v++{
		for i:=0;i<n;i++{
			for j:=i+1;j<n;j++{
				fmt.Printf("Update %d to %d,using %d\n",i,j,v)
				c := cost[i][v]+cost[v][j]
				if c < cost[i][j]{
					cost[i][j] = c
					cost[j][i] = c
					path[i][j] = v
					fmt.Println(i,j,c,v)
				}
			}
		}
	}
	return
}
