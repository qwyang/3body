package main

import "fmt"

/*
打印*号组成的直角三角形
*/
func StarsSimple() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

/*
打印*号组成的等腰三角形
*/
func StarsComplex() {
	N := 9
	for i := 0; i <= N/2; i++ {
		for j := 0; j < N; j++ {
			if j >= N/2-i && j <= N/2+i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

/*
打印99乘法表
*/
func MultiTable99() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%02d ", i, j, i*j)
		}
		fmt.Println()
	}
}

func main() {
	StarsSimple()
	StarsComplex()
	MultiTable99()
}
