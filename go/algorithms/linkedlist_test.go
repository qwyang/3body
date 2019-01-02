package algorithms

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := NewLinkedList()
	for i:=0;i<10;i++{
		list.PushBack(i)
	}
	list.Traverse(func(i interface{}) {
		fmt.Print(i," ")
	})
	fmt.Println("listsize:",list.Size())
	for i:=0;i<10;i++{
		list.PopBack()
	}
	list.Traverse(func(i interface{}) {
		fmt.Print(i," ")
	})
	fmt.Println("listsize:",list.Size())
	for i:=0;i<10;i++{
		list.PushFront(i)
	}
	list.Traverse(func(i interface{}) {
		fmt.Print(i," ")
	})
	fmt.Println("listsize:",list.Size())
	for i:=0;i<10;i++{
		list.PopFront()
	}
	list.Traverse(func(i interface{}) {
		fmt.Print(i," ")
	})
	fmt.Println()
	for i:=0;i<10;i+=2{
		e:=list.PushFront(i)
		list.InsertBefore(i+1,e)
	}
	list.Traverse(func(i interface{}) {
		fmt.Print(i," ")
	})
	fmt.Println("listsize:",list.Size())
	for i:=0;i<10;i++{
		list.Remove(i)
	}
	for i:=0;i<10;i+=2{
		e:=list.PushBack(i)
		list.InsertAfter(i+1,e)
	}
	list.Traverse(func(i interface{}) {
		fmt.Print(i," ")
	})
	fmt.Println("listsize:",list.Size())
	for i:=20;i>5;i--{
		list.Remove(i)
	}
	list.Traverse(func(i interface{}) {
		fmt.Print(i," ")
	})
	fmt.Println("listsize:",list.Size())
}
