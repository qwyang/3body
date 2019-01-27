package dp

import (
	"fmt"
	"testing"
)

func TestOptimalBinarySearchTree(t *testing.T) {
	P:=[]float32{0.22,0.18,0.20,0.05,0.25,0.02,0.08}
	//W:=[]string{"a","am","and","egg","if","the","two"}
	C,R:=OptimalBinarySearchTree(P)
	fmt.Println(C)
	fmt.Println(R)
}
