package main

import "fmt"

/*
状态模式：糖果机。NoQuarter->HasQuarter->Sold->SoldOut-
1.type State interface,状态必须实现的统一接口,每个状态下允许的动作；
2.ConcreteState,具体状态。
3.GumballMachine，通过状态确定行为。
*/
type State interface {
	InsertQuarter()
	EjectQuarter()
	TurnCranker()
	Dispense()
}

type GumballMachine struct{
	curState State
	count int
	NoQuarter *NoQuarterState
	HasQuarter *HasQuarterState
	Sold *SoldState
	SoldOut *SoldOutState
}

func NewGumballMachine(count int) *GumballMachine{
	if count <= 0 {
		panic("cannot filled GumballMachine with zero or negative number of Gumball.")
	}
	gm := &GumballMachine{
		count:count,
	}
	gm.NoQuarter = &NoQuarterState{BaseState{gm}}
	gm.HasQuarter = &HasQuarterState{BaseState{gm}}
	gm.Sold = &SoldState{BaseState{gm}}
	gm.SoldOut = &SoldOutState{BaseState{gm}}
	gm.curState = gm.NoQuarter
	return gm
}

func (gm *GumballMachine)SetState(s State){
	gm.curState = s
}

func (gm *GumballMachine)InsertQuarter(){
	gm.curState.InsertQuarter()
}

func (gm *GumballMachine)EjectQuarter(){
	gm.curState.EjectQuarter()
}

func (gm *GumballMachine)TurnCranker(){
	gm.curState.TurnCranker()
	gm.curState.Dispense()
}

func (gm *GumballMachine)Dispense(){
	gm.count--
	fmt.Println("a Gumball rolling out.")
}

func (gm *GumballMachine)Refill(count int){
	if count <= 0{panic("cannot fill GumballMachine with zero or negative number")}
	if gm.curState == gm.SoldOut {
		gm.count += count
		gm.curState = gm.NoQuarter
	}else if gm.curState == gm.NoQuarter{
		gm.count += count
	}else{
		fmt.Println("wait until soldout or noquater state.")
	}
}

type BaseState struct {
	gm *GumballMachine
}
type NoQuarterState struct {BaseState}
type HasQuarterState struct {BaseState}
type SoldState struct{BaseState}
type SoldOutState struct{BaseState}

func (s *NoQuarterState) InsertQuarter(){
	fmt.Println("Ok,Insert a Quarter.状态转移至HasQuarter")
	s.gm.SetState(s.gm.HasQuarter)
}
func (s *NoQuarterState) EjectQuarter(){
	fmt.Println("NoQurterState，EjectQuarter：无效操作")
}

func (s *NoQuarterState) TurnCranker(){
	fmt.Println("NoQurterState，TurnCranker：无效操作")
}
func (s *NoQuarterState) Dispense(){
	fmt.Println("NoQurterState，Dispense：无效操作")
}


func (s *HasQuarterState) InsertQuarter(){
	fmt.Println("HasQuarterState，InsertQuarter：无效操作")
}
func (s *HasQuarterState) EjectQuarter(){
	fmt.Println("HasQuarter，EjectQuarter，状态转移至NoQuarter")
	s.gm.SetState(s.gm.NoQuarter)
}

func (s *HasQuarterState) TurnCranker(){
	fmt.Println("HasQuarterState，TurnCranker,状态转移至Sold")
	s.gm.SetState(s.gm.Sold)
}
func (s *HasQuarterState) Dispense(){
	fmt.Println("HasQuarterState，Dispense：无效操作")
}


func (s *SoldState) InsertQuarter(){
	fmt.Println("SoldState，InsertQuarter：无效操作")
}
func (s *SoldState) EjectQuarter(){
	fmt.Println("SoldState，EjectQuarter，无效操作")
}

func (s *SoldState) TurnCranker(){
	fmt.Println("SoldState，TurnCranker,无效操作")
}
func (s *SoldState) Dispense(){
	s.gm.Dispense()
	if s.gm.count == 0{
		fmt.Println("SoldState,Dispense,状态转移至SoldOut")
		s.gm.SetState(s.gm.SoldOut)
	}else{
		fmt.Println("SoldState,Dispense,状态转移至NoQuarter")
		s.gm.SetState(s.gm.NoQuarter)
	}
}

func (s *SoldOutState) InsertQuarter(){
	fmt.Println("SoldOutState,Insert a Quarter,无效操作")
}
func (s *SoldOutState) EjectQuarter(){
	fmt.Println("SoldOutState，EjectQuarter，无效操作")
}

func (s *SoldOutState) TurnCranker(){
	fmt.Println("SoldOutState，TurnCranker：无效操作")
}
func (s *SoldOutState) Dispense(){
	fmt.Println("SoldOutState，Dispense：无效操作")
}

func main(){
	gm := NewGumballMachine(3)
	gm.InsertQuarter()
	gm.TurnCranker()
	gm.InsertQuarter()
	gm.TurnCranker()
	gm.InsertQuarter()
	gm.TurnCranker()
	gm.Refill(1)
	gm.InsertQuarter()
	gm.InsertQuarter() //无效操作
	gm.TurnCranker()
	gm.TurnCranker() //无效操作
	gm.EjectQuarter() //无效操作
}