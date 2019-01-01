package algorithms

/*
基于接口的抽象设计
 */
type LinkedListElem interface {
	Next() LinkedListElem
	Before() LinkedListElem
	SetNext(LinkedListElem) bool
	SetBefore(LinkedListElem) bool
	Value() interface{}
	SetValue(interface{}) bool
}

type linkedListElem struct {
	next LinkedListElem
	before LinkedListElem
	value interface{}
}

func NewLinkedListElem(value interface{})*linkedListElem{
	e := &linkedListElem{value:value}
	return e
}

func (e *linkedListElem)Next()LinkedListElem{
	return e.next
}

func (e *linkedListElem)SetNext(n LinkedListElem)bool{
	e.next = n
	return true
}

func (e *linkedListElem)Before()LinkedListElem{
	return e.before
}

func (e *linkedListElem)SetBefore(b LinkedListElem)bool{
	e.before = b
	return true
}

func (e *linkedListElem)SetValue(v interface{})bool{
	e.value = v
	return true
}

func (e *linkedListElem)Value()interface{}{
	return e.value
}

/*
双链表，而非循环链表
 */

type LinkedList interface {
	Head() LinkedListElem
	Tail() LinkedListElem
	Size()uint64
	PushFront(v interface{}) LinkedListElem
	PopFront()interface{}
	PushBack(v interface{}) LinkedListElem
	PopBack()interface{}
	InsertAfter(interface{},LinkedListElem)LinkedListElem
	InsertBefore(interface{},LinkedListElem)LinkedListElem
	Remove(interface{})LinkedListElem
	Traverse()
}

type linkedList struct {
	head LinkedListElem
	tail LinkedListElem
	size uint64
}

func NewLinkedList()*linkedList{
	h := NewLinkedListElem(nil)
	t := NewLinkedListElem(nil)
	h.SetBefore(nil)
	h.SetNext(t)
	t.SetBefore(h)
	t.SetNext(nil)
	list := &linkedList{h,t,0}
	return list
}

func (list *linkedList) Head() LinkedListElem {
	if list.size == 0 {
		return nil
	}
	return list.head.Next()
}

func (list *linkedList) Tail() LinkedListElem {
	if list.size == 0 {
		return nil
	}
	return list.tail.Before()
}

func (list *linkedList) Size() uint64 {
	return list.size
}

func (list *linkedList)PushFront(v interface{})LinkedListElem{
	if v == nil{
		return nil
	}
	e := NewLinkedListElem(v)
	a := list.head.Next()
	b := list.head
	e.SetNext(a)
	e.SetBefore(b)
	b.SetNext(e)
	a.SetBefore(e)
	list.size++
	return e
}

func (list *linkedList)PopFront()interface{}{
	if list.size == 0{
		return nil
	}
	e := list.Head()
	a := e.Next()
	b := e.Before()
	b.SetNext(a)
	a.SetBefore(b)
	list.size--
	return e.Value()
}

func (list *linkedList)PushBack(v interface{})LinkedListElem{
	if v == nil{
		return nil
	}
	e := NewLinkedListElem(v)
	a := list.tail
	b := a.Before()
	//fmt.Printf("e:%v,a:%v,b:%v\n",e,a,b)
	a.SetBefore(e)
	e.SetNext(a)
	b.SetNext(e)
	e.SetBefore(b)
	list.size++
	return e
}

func (list *linkedList)PopBack()interface{}{
	if list.size == 0{
		return nil
	}
	e := list.Tail()
	a := e.Next()
	b := e.Before()
	b.SetNext(a)
	a.SetBefore(b)
	list.size--
	return e.Value()
}
func (list *linkedList)InsertAfter(v interface{},b LinkedListElem)LinkedListElem{
	if v == nil{
		return nil
	}
	e := NewLinkedListElem(v)
	a := b.Next()
	b.SetNext(e)
	e.SetBefore(b)
	a.SetBefore(e)
	e.SetNext(a)
	list.size++
	return e
}
func (list *linkedList)InsertBefore(v interface{},a LinkedListElem)LinkedListElem{
	if v == nil{
		return nil
	}
	e := NewLinkedListElem(v)
	b := a.Before()
	b.SetNext(e)
	e.SetBefore(b)
	a.SetBefore(e)
	e.SetNext(a)
	list.size++
	return e
}
func (list *linkedList) Remove(v interface{})LinkedListElem{
	var e LinkedListElem
	for e=list.Head();e!=nil;e=e.Next(){
		if e.Value() == v {
			break
		}
	}
	if e == nil{
		return nil
	}
	b := e.Before()
	a := e.Next()
	b.SetNext(a)
	a.SetBefore(b)
	list.size--
	return e
}

func (list *linkedList)Traverse(f func(interface{})){
	for e:=list.head.Next();e!=list.tail;e=e.Next(){
		f(e.Value())
	}
}