package main

import (
	"fmt"
	"sync"
	"time"
)

/*
不公平读写锁的实现
RWLock{
	reader_counter
	reader_counter_lock
	file_lock
	RLock()
	RUnlock()
	Lock()
	Unlock()
}
 */

type RWLock struct {
	reader_counter int
	reader_counter_lock sync.Mutex
	file_lock sync.Mutex
}

func (l *RWLock) Lock(){
	l.file_lock.Lock()
}

func (l *RWLock) UnLock(){
	l.file_lock.Unlock()
}
/*
第1种情况：当前已有reader，既reader_counter>0，不阻塞
第2种情况：当前没有reader，既reader_counter==0
	2.1，当前没有writer，不阻塞
	2.2，当前有writer，阻塞
 */
func (l *RWLock) RLock(){
	l.reader_counter_lock.Lock()//有writer时其他reader阻塞于此
	l.reader_counter++
	if l.reader_counter == 1{
		l.file_lock.Lock()//有writer时第一个reader阻塞于此
	}
	l.reader_counter_lock.Unlock()
}
func (l *RWLock) RUnlock(){
	l.reader_counter_lock.Lock()
	l.reader_counter--
	if l.reader_counter == 0{
		l.file_lock.Unlock()
	}
	l.reader_counter_lock.Unlock()
}

func main(){
	wg := NewWaitGroup()
	l := &RWLock{}
	for i:=0;i<3;i++{
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			l.Lock()
			time.Sleep(time.Second)
			fmt.Printf("write %d\n",i)
			l.UnLock()
		}(i)
		go func(i int){
			defer wg.Done()
			l.RLock()
			time.Sleep(time.Second)
			fmt.Printf("read %d\n",i)
			l.RUnlock()
		}(i)
	}
	wg.Wait()
}