package examples

import (
	"bytes"
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

/*
格式化输入字符串，连续空格整理为1个，运算符和数字之间添加空格
 */
 func normalizeExpr(expr string) string{
	 var buf bytes.Buffer
	 var last rune=' '
	 for _,elem := range expr {
		 switch elem {
		 case '+', '-', '*', '/', '(', ')'://若运算符之间没有空格，则添加一个空格
			if last != ' '{
				fmt.Fprintf(&buf," %c ",elem)
			}else{
				fmt.Fprintf(&buf,"%c ",elem)
			}
			last = ' '
		 case ' '://多个空格整理为一个空格
		 	if last != ' '{
				fmt.Fprintf(&buf,"%c",' ')
				last = ' '
			 }
		 default://数字情况，前面无论是空格还是数字，一样打印
		 	fmt.Fprintf(&buf,"%c",elem)
		 	last = elem
		 }
	 }
	 //字符串末尾是运算符会多添加一个空格
	 return strings.TrimSpace(buf.String())
 }
/*
分割字符串，并且使用Atoi将字符串变为数字，表达式整体整理为[]int格式
 */
func NormalizeExpr(expr string) ([]int,error){
	expr = normalizeExpr(expr)
	data := strings.Split(expr," ")
	var ret []int
	var num int
	var err error
	for _,elem := range data {
		//fmt.Println(elem)
		switch elem {
			case "+","-","*","/","(",")":
				num = int([]rune(elem)[0])
			default:
				num,err = strconv.Atoi(elem)
				if err != nil{
					return nil,err
				}
		}
		ret = append(ret,num)
		//fmt.Println(ret,num)
	}
	return ret,nil
}
/*
将中缀表达式转化为后缀表达式
 */
func infixToPosfix(data []int) (ret []int,err error){
	stack := list.New()
	for _,elem := range data{
		switch elem {
			case '+','-':
				for head:=stack.Front();head!=nil;head = head.Next(){
					operator := head.Value.(int)
					if operator != '(' {
						fmt.Println("pop ",operator)
						stack.Remove(head)
						ret = append(ret,operator)
					}else{
						break
					}
				}
				stack.PushFront(elem)
				fmt.Println("push ",elem)
			case '*','/':
				for head:=stack.Front();head!=nil;head = stack.Front(){
					operator := head.Value.(int)
					if operator == '*'||operator == '/' {
						fmt.Println("pop ",operator)
						stack.Remove(head)
						ret = append(ret,operator)
					}else{
						break
					}
				}
				stack.PushFront(elem)
			case '(':
				stack.PushFront(elem)
			case ')':
				for head:=stack.Front();head!=nil;head = stack.Front(){
					operator := head.Value.(int)
					if operator != '(' {
						stack.Remove(head)
						ret = append(ret,operator)
					}else{
						stack.Remove(head)
						break
					}
				}
			default:
				ret = append(ret,elem)
		}
	}
	for head:=stack.Front();head!=nil;head = stack.Front(){
		operator := head.Value.(int)
		stack.Remove(head)
		fmt.Println("pop ",operator)
		ret = append(ret,operator)
	}
	return
}


func Calculate(expr string) (int,error) {
	data,err := NormalizeExpr(expr)
	if err != nil{
		return 0,err
	}
	data,err = infixToPosfix(data)
	if err != nil{
		return 0,err
	}
	stack := list.New()
	var sum int
	for _,d := range data{
		switch d {
		case '+':
			a:= stack.Remove(stack.Front())
			b:= stack.Remove(stack.Front())
			sum = a.(int)+b.(int)
			stack.PushFront(sum)
		case '-':
			a:= stack.Remove(stack.Front())
			b:= stack.Remove(stack.Front())
			sum = b.(int)-a.(int)
			stack.PushFront(sum)
		case '*':
			a:= stack.Remove(stack.Front())
			b:= stack.Remove(stack.Front())
			sum = b.(int)*a.(int)
			stack.PushFront(sum)
		case '/':
			a:= stack.Remove(stack.Front())
			b:= stack.Remove(stack.Front())
			sum = b.(int)/a.(int)
			stack.PushFront(sum)
		default:
			stack.PushFront(d)
		}
	}
	return sum,nil
}
