package main

import (
	"fmt"
	"math"
	"time"
)

type Point struct {
	x, y int
}

/*
Point对象方法，非指针类型接收器
*/
func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

type Vec struct {
	X, Y float32
}

/*
Vec对象方法，非指针类型接收器
*/
func (self Vec) Sub(other Vec) Vec {
	return Vec{self.X - other.X, self.Y - other.Y}
}

func (self Vec) Add(other Vec) Vec {
	return Vec{self.X + other.X, self.Y + other.Y}
}

func (self Vec) DistanceTo(t Vec) float32 {
	d := t.Sub(self)
	r := float32(math.Sqrt(float64(d.X*d.X + d.Y*d.Y)))
	return r
}

func (self Vec) Normalize() Vec {
	r := math.Sqrt(float64(self.X*self.X + self.Y*self.Y))
	if r > 0 {
		return Vec{float32(float64(self.X) / r), float32(float64(self.Y) / r)}
	}
	return Vec{0, 0}
}

func (self Vec) Scale(s float32) Vec {
	return Vec{self.X * s, self.Y * s}
}

type Player struct {
	curPos    Vec
	targetPos Vec
	speed     float32
}

func (p *Player) MoveTo(t Vec) {
	p.targetPos = t
}

func (p *Player) Pos() Vec {
	return p.curPos
}

func (p *Player) Distance() float32 {
	d := p.curPos.DistanceTo(p.targetPos)
	return d
}

func (p *Player) Arrived() bool {
	return p.Distance() < p.speed
}

func (p *Player) Update() {
	if !p.Arrived() {
		dir := p.targetPos.Sub(p.curPos).Normalize()
		step := dir.Scale(p.speed)
		p.curPos = p.curPos.Add(step)
	}
}

/*
实例1：相同类型结构体可以直接比较，不同类型无法比较
*/
func testEqual() {
	a := Point{1, 1}
	b := Point{1, 1}
	//c := Vec{1,1}
	fmt.Printf("%v\n", a == b)
	fmt.Printf("%v\n", &a == &b)
	//fmt.Printf("%v\n",a == c)
}

/*
实例2：矢量游戏，设定目的地和speed，循环直到到达目的地为止
*/
func testVecPlayer() {
	p := Player{curPos: Vec{0, 0}, targetPos: Vec{3, 4}, speed: 1.0}
	for !p.Arrived() {
		p.Update()
		fmt.Printf("curPos:%v,distance:%v\n", p.Pos(), p.Distance())
		time.Sleep(time.Second)
	}
	//fmt.Printf("curPos:%v,distance:%v\n",p.Pos(),p.Distance())
}

var eventMap = make(map[string][]func(interface{}))

func register(event string, handler func(interface{})) {
	handlers := eventMap[event]
	handlers = append(handlers, handler)
	eventMap[event] = handlers
}
func callEvent(event string, data interface{}) {
	handlers := eventMap[event]
	for _, function := range handlers {
		function(data)
	}
}

type EventListener struct{}

func (e *EventListener) onChange(data interface{}) {
	fmt.Printf("method:%v\n", data)
}
func Change(data interface{}) {
	fmt.Printf("function:%v\n", data)
}
func testEventSystem() {
	ins1 := &EventListener{}
	register("Update", ins1.onChange)
	register("Update", Change)
	//fmt.Printf("map:%v\n",eventMap)
	callEvent("Update", time.Now().String())
}

func main() {
	fmt.Printf("--------------testEqual----------------\n")
	testEqual()
	fmt.Printf("--------------testVecPlayer----------------\n")
	testVecPlayer()
	fmt.Printf("--------------testEventSystem----------------\n")
	testEventSystem()
}
