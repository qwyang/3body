package examples

import (
	"bytes"
	"container/list"
	"fmt"
)

/*
汉诺塔：从src柱移动n个圆盘到dst柱，一次只能移动一个圆盘，大圆盘不能放在小圆盘上面，可借助mid柱。
 */
 func move(n int,src,dst rune){
 	fmt.Printf("move %dth from %c to %c\n",n,src,dst)
 }

 func wrap(fn func(int,rune,rune))func (int,rune,rune){
 	var count int
 	return func(n int,src,dst rune){
 		count++
 		fmt.Printf("step %d:",count)
 		fn(n,src,dst)
	}
 }

 var Move func(int,rune,rune)

func hanoi(n int,src,mid,dst rune){
	if n==0{
		//递归结束条件
	}else{
		//移动n-1个圆盘到mid
		hanoi(n-1,src,dst,mid)
		//移动最下面一个圆盘到dst
		Move(n,src,dst)
		//移动n-1个圆盘到dst
		hanoi(n-1,mid,src,dst)
	}
}

type StackElem struct {
	fn string
	nth int
	n int
	src,mid,dst rune
}

func hanoi_nr(n int,src,mid,dst rune) (steps []byte){
	stack := list.New()
	e := &StackElem{fn:"hanoi",n:n,src:src,mid:mid,dst:dst}
	stack.PushBack(e)
	var count int
	var buf bytes.Buffer
	for stack.Len() > 0 {
		back := stack.Back()
		elem := back.Value.(*StackElem)
		stack.Remove(back)
		fn,n,nth,src,mid,dst := elem.fn,elem.n,elem.nth,elem.src,elem.mid,elem.dst
		if fn == "hanoi" {
			if n == 0 {
				//退栈，什么也不做
			}else{//后进先出，步骤倒序
				//第三步：把n-1个圆盘从辅助柱移动到目的柱
				e := &StackElem{fn:"hanoi",n:n-1,src:mid,mid:src,dst:dst}
				stack.PushBack(e)
				//第二步：把最后一个圆盘移动到目的柱
				e = &StackElem{fn:"move",nth:n,src:src,mid:mid,dst:dst}
				stack.PushBack(e)
				//第一步：把n-1个圆盘移动到辅助柱
				e = &StackElem{fn:"hanoi",n:n-1,src:src,mid:dst,dst:mid}
				stack.PushBack(e)
			}
		}else{
			count++
			fmt.Fprintf(&buf,"step %d:move %dth from %c to %c\n",count,nth,src,dst)
		}
	}
	return buf.Bytes()
}

func init(){
	Move=wrap(move)
}