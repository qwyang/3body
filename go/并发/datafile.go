package main

import (
	"io"
	"log"
	"os"
	"sync"
)

/*
多个读者和写着同时并发进行“分块”读写操作。
读行为：没有读到数据，一直等待下去。
演示：
1.mutex
2.condition
3.atomic
4.Readat
*/
type DataFile interface {
	Read() (rsn int64,data Data,err error)
	Write(data Data)(wsn int64,err error)
	RSN() int64
	WSN() int64
	Datalen() int64
	Close() error
}

type Data []byte
var logger = log.New(os.Stdout,"test_",log.Ldate|log.Ltime|log.Lshortfile)

type datafile struct {
	f *os.File
	fmutex sync.RWMutex
	roffset int64
	woffset int64
	datalen int64
	rmutex sync.Mutex
	wmutex sync.Mutex
	rcond *sync.Cond
}

func NewDataFile(path string,datalen int64)(DataFile,error){
	file,err := os.Create(path)
	if err != nil{
		return nil,err
	}
	df := &datafile{f:file,datalen:datalen}
	df.rcond = sync.NewCond(df.fmutex.RLocker())
	return df,nil
}

func (f *datafile)Read() (rsn int64,data Data,err error){
	f.rmutex.Lock()
	offset := f.roffset
	f.roffset += f.datalen
	f.rmutex.Unlock()
	rsn = offset/f.datalen
	data = make([]byte,f.datalen)
	f.fmutex.RLock()
	defer f.fmutex.RUnlock()
	for{
		n,err := f.f.ReadAt(data,offset)
		if err != nil{
			if err == io.EOF {
				logger.Printf("read rsn:%d, meet EOF\n",rsn)
				f.rcond.Wait()
				continue
			}
			if int64(n) != f.datalen {
				logger.Printf("read:%d less than expected:%d\n",n,f.datalen)
			}
			return offset/f.datalen,data,err
		}
		logger.Printf("read rsn:%d,data=%v\n",rsn,data)
		return offset/f.datalen,data,nil
	}
}

func (f *datafile)Write(data Data)(wsn int64,err error){
	f.rmutex.Lock()
	offset := f.woffset
	f.woffset += f.datalen
	f.rmutex.Unlock()

	wsn = offset/f.datalen

	f.fmutex.Lock()
	defer f.fmutex.Unlock()
	_,err = f.f.WriteAt(data[:f.datalen],offset)
	logger.Printf("write wsn:%d,data=%v\n",wsn,data)
	f.rcond.Broadcast()
	return
}

func (f *datafile) RSN() int64{
	f.rmutex.Lock()
	offset := f.roffset
	f.rmutex.Unlock()
	return offset/f.datalen
}
func (f *datafile) WSN() int64{
	f.rmutex.Lock()
	offset := f.woffset
	f.rmutex.Unlock()
	return offset/f.datalen
}
func (f *datafile) Datalen() int64{
	return f.datalen
}
func (f *datafile)Close() error{
	if f.f == nil{
		return nil
	}
	return f.f.Close()
}
