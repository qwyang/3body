package main

import (
	"errors"
	"fmt"
)

/*
功能：div函数，当除数为0时返回除数为0的错误。
目的：示例自定义一个错误errors.New
 */
func div(divident,divider int) (int,error){
	if divider == 0{
		return 0,errors.New("zero divider")
	}
	return divident/divider,nil
}
/*
目的：自定义"高级"错误类型，错误信息保护文件名，行号
只要实现error接口既可type error interface{Error()string}
 */
type ParseError struct {
	Filename string
	Line int
}

func (p *ParseError) Error()string{
	return fmt.Sprintf("%s,%d",p.Filename,p.Line)
}

func NewParseError(filename string,lineno int) *ParseError {
	return &ParseError{filename,lineno}
}

/*
目的：演示函数宕机时如何进行捕获恢复
注意: 只能在函数返回时进行recover
*/
func PanicRecover(){
	defer func(){
		err := recover()
		switch err.(type) {
		case *ParseError:
			fmt.Println("recover from a ParseError")
		default:
			fmt.Println("recover from other error")
		}
	}()
	fmt.Printf("PanicRecover,running before panic!\n")
	panic(&ParseError{"error.go",49})
	fmt.Printf("PanicRecover,running after panic!\n")
}


func main(){
	r,err := div(10,0)
	fmt.Println(r,err)
	var e error
	e = NewParseError("error.go",39)
	fmt.Println(e)
	switch e.(type) {
		case *ParseError:
			p := e.(*ParseError)
			fmt.Printf("filename:%s,lineno:%d\n",p.Filename,p.Line)
		default:
			fmt.Println("other error")
	}
	PanicRecover()
}