package main

import (
	"fmt"
	"time"
)

/*
计时器timer，闹钟，到期后执行特定操作。
*/
func Timer() {
	exit := make(chan bool)
	time.AfterFunc(time.Second, func() {
		fmt.Println("execute after 1 second.")
		exit <- true
	})
	<-exit
}

/*
ticker:打点器，每隔一段时间执行一次。
*/
func Ticker() {
	stoper := time.NewTimer(time.Second * 2)
	ticker := time.NewTicker(time.Millisecond * 500)
OuterLoop:
	for {
		select {
		case <-stoper.C:
			break OuterLoop
		case <-ticker.C:
			fmt.Println("ticker run...")
		}
	}
}

func main() {
	Timer()
	Ticker()
}
