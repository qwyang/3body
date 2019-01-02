package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
目标：演示注册信号处理和取消注册。信号时进程间通信的唯一异步方法。
常用信号：
	syscall.SIGINT,syscall.SIGQUIT
注册取消函数：
	signal.Notify(ch chan<- os.Signal,sig... os.Signal)
	signal.Stop(ch chan<- os.Signal)
获取进程PID:
	os.Getpid()
Linux命令行发送信号方法：
	kill -SIGQUIT PID
	kill -SIGINT PID
*/
func testSignal(){
	var count int
	ch := make(chan os.Signal,1)
	signal.Notify(ch,syscall.SIGINT,syscall.SIGQUIT)
	for s := range ch {
		if s == syscall.SIGQUIT{
			fmt.Println("received:",syscall.SIGQUIT)
		}
		if s == syscall.SIGINT {
			fmt.Println("received:",syscall.SIGINT)
		}
		count++
		if count >= 5{
			signal.Stop(ch)
			close(ch)
		}
	}
}

func main(){
	fmt.Printf("PID:%d\n",os.Getpid())
	testSignal()
}

