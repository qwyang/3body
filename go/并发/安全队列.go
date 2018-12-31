package main

import (
	"fmt"
	"sync"
)

/*
SafeQueue
*/
type SafeQueue struct {
	queue [10]interface{}
	queue_size int
	r_index int
	p_index int
	queue_mutex sync.Mutex
	queue_p_cond *sync.Cond
	queue_c_cond *sync.Cond
}

func NewSafeQueue()*SafeQueue{
	q := &SafeQueue{}
	q.queue_c_cond = sync.NewCond(&q.queue_mutex)
	q.queue_p_cond = sync.NewCond(&q.queue_mutex)
	return q
}

func (sq *SafeQueue)Put(elem interface{}){
	sq.queue_mutex.Lock()
	for sq.queue_size >= len(sq.queue) {//队列满，生产者等待
		sq.queue_p_cond.Wait()
	}
	sq.queue[sq.p_index] = elem
	sq.p_index = (sq.p_index + 1) % len(sq.queue)
	sq.queue_size++
	sq.queue_c_cond.Signal()//通知消费者有数据了
	sq.queue_mutex.Unlock()
}

func (sq *SafeQueue)Get()interface{}{
	sq.queue_mutex.Lock()
	for sq.queue_size <= 0 {//队列空，消费者等待
		sq.queue_c_cond.Wait()
	}
	elem := sq.queue[sq.r_index]
	sq.r_index = (sq.r_index + 1) % len(sq.queue)
	sq.queue_size--
	sq.queue_p_cond.Signal()//通知生产者有空间了
	sq.queue_mutex.Unlock()
	return elem
}

func main(){
	wg := NewWaitGroup()
	sq := NewSafeQueue()
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i:=0;i<100;i++{
			sq.Put(i)
			fmt.Println("put ",i)
		}
	}()
	go func() {
		defer wg.Done()
		for i:=0;i<100;i++{
			elem := sq.Get()
			fmt.Println("get ",elem)
		}
	}()
	wg.Wait()
}
