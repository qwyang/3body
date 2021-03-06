package examples

import (
	"fmt"
	"testing"
)

func TestDecompose(t *testing.T) {
	data := decompose(341)
	fmt.Println(data)
	data = decompose(2*3*5*7*2*13*11*17*19*101)
	fmt.Println(data)
	data = decompose(561)
	fmt.Println(data)
}
func TestDecompose2(t *testing.T) {
	data := decompose2(341)
	fmt.Println(data)
	data = decompose2(2*3*5*7*2*13*11*17*19*101)
	fmt.Println(data)
	data = decompose2(71)
	fmt.Println(data)
}