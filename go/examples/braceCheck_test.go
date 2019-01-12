package examples

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBraceCheck(t *testing.T) {
	str := "{[a+b]*(a+(b+3))}a"
	fmt.Println(BraceCheck(str))
	str = "{[a+b]]*(a+(b+3))}a"
	fmt.Println(BraceCheck(str))
	fmt.Println(reflect.TypeOf('a').Name())
}
