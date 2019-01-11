package examples

import (
	"fmt"
	"testing"
)

func TestPow(t *testing.T){
	fmt.Println(Pow(2,60))
	fmt.Println(PowNr(2,60))
}
