package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

/*
3个任务同时进行网络访问，主线程等待任务完成。
waitgroup作为任务和主线程同步手段，任务之间没有任何关联。
*/
func WaitGroup() {
	var wg sync.WaitGroup
	urls := []string{
		"http://www.github.com/",
		"https://www.qiniu.com/",
		"https://www.golangtc.com/",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			//content,_ := ioutil.ReadAll(resp.Body)
			//fmt.Println(string(content))
			defer resp.Body.Close()
		}(url)
	}
	wg.Wait()
}

/*
10个协程同时做sum++运算，sum变量由mutex保护,wg通知main线程任务完成。
任务之间相互联系：对同一变量进行操作。
使用Mutex(互斥锁)和原子加操作两种方式实现共享数据保护。
WaitGroup(等待组)实现任务同步。
*/
func ShareProtect() {
	var (
		sum   int32
		mutex sync.Mutex
		wg    sync.WaitGroup
	)
	f := func() {
		for i := 0; i < 10000; i++ {
			mutex.Lock()
			sum++
			mutex.Unlock()
		}
		wg.Done()
	}
	f2 := func() {
		for i := 0; i < 10000; i++ {
			atomic.AddInt32(&sum, 1)
		}
		wg.Done()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f()
	}
	wg.Wait()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f2()
	}
	wg.Wait()
	fmt.Printf("final sum=%d\n", sum)
}

/*
2个协程同时打印数据，协程1打印1-1000，协程2打印1001-2000,查看并发执行效果。
同步机制：
main协程等待任务完成，任务协程需要通知main协程任务结束。使用channel实现任务的同步。channel相当于线程安全的队列。
goroutine和线程：
goroutine行为类似线程会进行任务切换，但是是由Go runtime进行调度，而不是由操作系统进行调度。
goroutine和协程：
不同于协程，协程是单线程运行，不主动出让cpu不会进行切换。
*/
func TwoGoRoutine(cpu int) {
	var quit = make(chan int)
	f := func(min, max int, q chan int) {
		for i := min; i <= max; i++ {
			fmt.Println(i)
		}
		q <- min
		fmt.Printf("Goroutine end return:%d\n", min)
	}
	runtime.GOMAXPROCS(cpu)
	go f(1, 100, quit)
	go f(101, 200, quit)
	r := <-quit
	fmt.Printf("routine end return:%d\n", r)
	r = <-quit
	fmt.Printf("routine end return:%d\n", r)
}

/*
并发打印，channel既当数据通道，也当信号通道。一般来讲数据通道和信号通道是分开的。
go rutine之间共享的数据是channel
生产者-消费者模式，消费者死循环等待数据到来，如同服务器一般。
两者运行时交替调度执行，因为读写channel是阻塞性操作。
*/
func GoPrint() {
	ch := make(chan int)
	go func() { //consumer
		for {
			d := <-ch
			if d != 0 {
				fmt.Println("Consum data:", d)
			} else {
				break
			}
		}
		//通知主协程任务完成。
		ch <- 0
	}()
	for i := 10; i > 0; i-- {
		fmt.Println("Generating a new data:", i)
		ch <- i
	}
	//通知消费者停止打印
	ch <- 0
	//等待消费者任务结束
	<-ch
}

/*
使用多通道复用（select）模拟RPC远程调用
后台gorutine运行服务器程序
客户端发出请求后等待响应
客户端输出响应结果。
*/
func RPCSim() {
	ch := make(chan string)
	RPCClient := func(ch chan string, req string) (string, error) {
		ch <- req
		select {
		case data := <-ch:
			return data, nil
		case <-time.After(time.Second * 2):
			return "", errors.New("timeout")
		}
	}
	RPCServer := func(ch chan string) { //server
		for {
			data := <-ch
			fmt.Println("server received from client:", data)
			time.Sleep(time.Second * 2)
			ch <- "roger"
		}
	}
	go RPCServer(ch)
	/*以下是客户端代码*/
	data, err := RPCClient(ch, "hi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client received from server:", data)
	}
}

func main() {
	//WaitGroup()
	//fmt.Printf("GoMaxProcs:%d\n",runtime.NumCPU())
	//TwoGoRoutine(runtime.NumCPU())
	//ShareProtect()
	//GoPrint()
	RPCSim()
	//var quit = make(chan int)
	//quit <- 0//stop here forever
}
