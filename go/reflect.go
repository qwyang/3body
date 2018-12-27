package main

import (
	"fmt"
	"reflect"
)

/*
获取struct变量的类型对象(reflect.Type)信息
声明State接口及其实现*StateInfo，*IdleState
*/
type State interface {
	Name() string
}

type StateInfo struct {
	name string
}

func (s *StateInfo) Name() string {
	return s.name
}

func (s *StateInfo) setName(name string) {
	s.name = name
}

func GetTypeName() {
	type IdleState struct {
		StateInfo
	}
	type RunningState struct {
		StateInfo
	}
	idle := IdleState{}
	run := RunningState{}
	f := func(s State) string {
		var typeOfState reflect.Type = reflect.TypeOf(s)
		return typeOfState.Elem().Name()
	}
	name := f(&idle)
	fmt.Printf("name:%s\n", name)
	idle.setName(name)
	fmt.Printf("state:%v\n", idle.name)
	name = f(&run)
	fmt.Printf("name:%s\n", name)
	idle.setName(name)
	fmt.Printf("state:%v\n", idle.name)
	type Cat struct {
		Name string
		age  int
	}
	c := Cat{}
	tofState := reflect.TypeOf(c)
	for i := 0; i < tofState.NumField(); i++ {
		f := tofState.Field(i)
		fmt.Println(f.Type, f.Name, f.Index, f.Offset)
	}
}

/*
获取struct变量的值对象(reflect.Value)信息
Value可修改条件：1.可寻址（指针类型） 2.可导出(大写字母开头）
*/
func SetValue() {
	type Cat struct {
		Name string
		Age  int
	}
	c := Cat{}
	fmt.Printf("cat:%+v\n", c)
	vofCat := reflect.ValueOf(&c)
	vofCat.Elem().FieldByName("Age").SetInt(100)
	vofCat.Elem().FieldByName("Name").SetString("ketty")
	fmt.Printf("cat:%+v\n", c)
}

func main() {
	SetValue()
	GetTypeName()
}
