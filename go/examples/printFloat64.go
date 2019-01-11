package examples

import "fmt"
func printIntPart(intPart int64){
	if intPart >= 10 {//recursive
		printIntPart(intPart/10)
	}
	//base condition
	fmt.Print(intPart%10)
}

func printLittlePart(littlePart float64){
	//base condition: littlePart == 0.0
	d := int64(10*littlePart)
	r := 10*littlePart - float64(d)
	fmt.Print(d)
	if littlePart > 0.0 {
		printLittlePart(r)
	}
}

func printFloat64(real float64){
	intPart := int64(real)
	littlePart := real - float64(intPart)
	printIntPart(intPart)
	fmt.Print(".")
	printLittlePart(littlePart)
}