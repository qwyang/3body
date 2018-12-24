package main

import (
	"fmt"
	"math"
	"time"
)

type Point struct {
	x,y int
}
/*
Point对象方法，非指针类型接收器
 */
func (p Point) Add(q Point) Point{
	return Point{p.x+q.x,p.y+q.y}
}
/*
Point对象方法，指针类型接收器SetX/SetY
 */
func (p *Point) SetX(x int) {
	p.x = x
}

func (p *Point) SetY(y int){
	p.y = y
}

type Vec struct {
	X,Y float32
}

func (self Vec)Sub(other Vec) Vec {
	return Vec{self.X-other.X,self.Y-other.Y}
}

func (self Vec)Add(other Vec) Vec {
	return Vec{self.X+other.X,self.Y+other.Y}
}

func (self Vec)DistanceTo(t Vec)float32{
	d := t.Sub(self)
	r := float32(math.Sqrt(float64(d.X*d.X + d.Y*d.Y)))
	return r
}

func (self Vec)Normalize() Vec{
	r := math.Sqrt(float64(self.X*self.X + self.Y*self.Y))
	if r > 0{
		return Vec{float32(float64(self.X)/r),float32(float64(self.Y)/r)}
	}
	return Vec{0,0}
}

func (self Vec)Scale(s float32) Vec {
	return Vec{self.X*s,self.Y*s}
}

type Player struct {
	curPos Vec
	targetPos Vec
	speed float32
}

func (p *Player)MoveTo(t Vec){
	p.targetPos = t
}

func (p *Player)Pos() Vec{
	return p.curPos
}

func (p *Player)Arrived() bool{
	return p.curPos.DistanceTo(p.targetPos) < p.speed
}

func (p *Player)Update(){
	if !p.Arrived(){
		dir := p.targetPos.Sub(p.curPos).Normalize()
		step := dir.Scale(p.speed)
		p.curPos = p.curPos.Add(step)
	}
}

func testPoint(){
	a := Point{1,1}
	b := Point{1,1}
	fmt.Printf("%v\n",a==b)
	fmt.Printf("%v\n",&a==&b)
	fmt.Printf("%v+%v=%v",a,b,a.Add(b))
}

func main(){
	p := Player{curPos:Vec{0,0},targetPos:Vec{15,5},speed:1.0}
	for !p.Arrived(){
		p.Update()
		fmt.Printf("curPos:%v\n",p.Pos())
		time.Sleep(1e9)
	}
}