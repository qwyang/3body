package main

import (
	"bytes"
	"fmt"
	"math"
)

/*
1.闭包closure,记忆效应，斐波那契函数
 */
func FibGen()func(int)int{
	var results = []int{1,1}
	return func(n int)int{
		if n <= 0 {
			panic("n must > 0")
		}
		if n > len(results) {
			length := len(results)
			j := n - length
			for a, b := results[length-2], results[length-1]; j > 0; j-- {
				c := a + b
				a = b
				b = c
				results = append(results, b)
			}
		}
		return results[n-1]
	}
}
 /*
 2.可变参数,类型断言
  */
func printType(params ...interface{})string{
	var buf bytes.Buffer
	var paraType string
	var para interface{}
  	for _,para = range params {
  		switch para.(type){
		case bool:
			paraType = "bool"
		case string:
			paraType = "string"
		case int:
			paraType = "int"
		case float32,float64:
			paraType = "float"
		default:
			paraType = "unkonwn"
		}
  		buf.WriteString("value:")
  		value := fmt.Sprintf("%v,",para)
  		buf.WriteString(value)
  		buf.WriteString("type:")
  		buf.WriteString(paraType)
  		buf.WriteString("\n")
	}
  	return buf.String()
}

func main(){
	v := printType(1,2,"abc",math.Pi,true)
	fmt.Println(v)
	fib := FibGen()
	for i:=1;i<=10;i++{
		fmt.Printf("%d.%d\n",i,fib(i))
	}
}