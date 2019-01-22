package algorithms

import (
	"fmt"
	"testing"
)

func TestFindKth(t *testing.T) {
	array := randArray(1000000,1000000)
	k := 500
	e,c := FindKth(array,k)
	qsort(array,0,len(array)-1)
	reverse(array)
	fmt.Println(e,c,array[k-1])
}
