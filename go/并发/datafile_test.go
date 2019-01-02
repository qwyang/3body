package main

import (
	"os"
	"sync"
	"testing"
)

func TestDatafile(t *testing.T) {
	os.Remove("tmp.txt")
	df,err := NewDataFile("tmp.txt",1)
	if err != nil {
		t.Fatal("NewDataFile Failed.")
	}
	defer df.Close()
	var i int
	var wg sync.WaitGroup

	for i=0;i<100;i++{
		wg.Add(1)
		go func(i uint8) {
			id := i
			data := []byte{i+'a'}
			wsn,err := df.Write(data)
			if err != nil{
				t.Fatal("df.write failed,wsn:",wsn)
			}
			t.Logf("write goroutine id:%d,wsn:%d done,data=%s\n",id,wsn,data)
			wg.Done()
		}(uint8(i))
	}
	for j:=0;j<i;j++{
		wg.Add(1)
		go func(i uint8) {
			id := i
			rsn,data,err := df.Read()
			if err != nil{
				t.Fatal("read go routine id:",id,"read failed,rsn:",rsn)
			}
			t.Logf("read goroutine id:%d,rsn:%d,data=%s\n",id,rsn,data)
			wg.Done()
		}(uint8(j))
	}
	wg.Wait()
}

func TestReadAt(t *testing.T){
	file,err := os.Open("tmp.txt")
	data := make([]byte,10)
	if err == nil{
		n,err := file.ReadAt(data,95)
		t.Log(n,err)
		file.Seek(101,0)
		n,err = file.Read(data)
		t.Log(n,err)
		n,err = file.Read(data)
		t.Log(n,err)
	}
}
