package examples

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	a,b := swap(1,2)
	fmt.Println(a==2 && b==1)
	a,b = swap2(100,200)
	fmt.Println(a==200 && b==100)
}
