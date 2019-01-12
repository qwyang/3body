package examples

import "container/list"

/*
括号表达式检查，支持[]{}()
*/
func match(a,b rune)bool{
	if a == '[' && b == ']'{return true}
	if a == '{' && b == '}'{return true}
	if a == '(' && b == ')'{return true}
	return false
}
func BraceCheck(data string)bool{
	stack := list.New()
	for _,c := range data{
		switch c {
			case '[','{','(':
				stack.PushFront(c)
			case ']','}',')':
				before:=stack.Remove(stack.Front())
				if match(before.(rune),c) {
					continue
				}else {
					return false
				}
			default:
		}
	}
	return true
}
