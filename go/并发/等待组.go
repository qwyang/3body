package main

import "sync"

/*
1.计数器：counter int32，counter_mutex
2.条件变量：cond，绑定counter_mutex
 */
type WaitGroup struct {
	counter int
	counter_mutex sync.Mutex
	cond *sync.Cond
}

func NewWaitGroup()*WaitGroup{
	wg := &WaitGroup{}
	wg.cond = sync.NewCond(&wg.counter_mutex)
	return wg
}

func (wg *WaitGroup)Add(c int){
	wg.counter_mutex.Lock()
	wg.counter += c
	if wg.counter == 0 {
		wg.cond.Signal()
	}
	wg.counter_mutex.Unlock()
}
func (wg *WaitGroup)Wait(){
	wg.counter_mutex.Lock()
	for wg.counter != 0 {
		wg.cond.Wait()
	}
	wg.counter_mutex.Unlock()
}
func (wg *WaitGroup)Done(){
	wg.Add(-1)
}