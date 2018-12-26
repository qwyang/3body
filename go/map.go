package main

import "fmt"

/*
多键查询，key：结构体的多个部分（name，age），value：整个结构体
*/
type Key struct {
	name string
	age int
}

type P struct {
	name string
	age int
	married bool
}
/*
??结构体作为值有问题，不要协程map[Key]P
 */
var MultiKeyMap = make(map[Key]*P)

func BuildMultiKeyMap(list []*P){
	for _,p := range list{
		key := Key{p.name,p.age}
		MultiKeyMap[key] = p
	}
}

func QueryMultiKeyMap(name string,age int) *P {
	key := Key{name,age}
	v,ok := MultiKeyMap[key]
	if ok {
		return v
	}else{
		return nil
	}
}

func main(){
	l := []*P{
		&P{"a",1,false},
		&P{"b",3,false},
		&P{"aa",40,true},
		&P{"aaa",30,true},
		&P{"bb",30,true},
		}
	BuildMultiKeyMap(l)
	p := QueryMultiKeyMap("b",3)
	fmt.Println(p)
	p = QueryMultiKeyMap("bb",30)
	fmt.Println(p)
	p = QueryMultiKeyMap("aa",40)
	fmt.Println(p)
	type A struct{
		x,y int
	}
}