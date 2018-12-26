package main

import (
	"container/list"
	"fmt"
)
/*
演示：linkedlist的使用，双链表允许从队列前方或后方插入。支持在某元素前/后插入。
插入并打印：1，2，3，3
 */
func main(){
	l := list.New()
	elem :=l.PushBack(2)
	l.PushFront(1)
	elem = l.InsertAfter(3,elem)
	l.InsertBefore(3,elem)
	for elem := l.Front();elem != nil;elem = elem.Next(){
		fmt.Println(elem.Value)
	}
}
