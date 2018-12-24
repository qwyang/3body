package main

import "fmt"

func StarsSimple(){
	for i:=1;i<=9;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func StarsComplex(){
	N := 9
	for i:=0;i<=N/2;i++ {
		for j:=0;j<N;j++{
			if j>=N/2-i && j<=N/2+i{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func MultiTable99(){
	for i:=1;i<=9;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%dx%d=%02d ",i,j,i*j)
		}
		fmt.Println()
	}
}

func main(){
	StarsSimple()
	StarsComplex()
	MultiTable99()
}
