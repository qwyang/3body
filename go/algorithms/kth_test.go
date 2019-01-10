package algorithms

import (
	"fmt"
	"testing"
)

func TestFindKth(t *testing.T) {
	array := randArray(100,1000)
	k := 50
	e,c := FindKth(array,k)
	qsort(array,0,len(array)-1)
	reverse(array)
	fmt.Println(e,c,array[k-1])
}
