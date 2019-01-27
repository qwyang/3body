package algorithms

import (
	"fmt"
	"testing"
)

func TestOptMatrix(t *testing.T) {
	C:=[]int{50,10,40,30,5}
	M,L:=OptMatrix(C)
	fmt.Println(M)
	fmt.Println(L)
}
