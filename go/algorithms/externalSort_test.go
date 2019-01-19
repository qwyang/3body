package algorithms

import (
	"encoding/binary"
	"os"
	"testing"
)
func generateTmpFile(tmpfile string,n int,max int){
	f, _ := os.OpenFile(tmpfile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer f.Close()
	for _, v := range randArray(n, max) {
		err := binary.Write(f, binary.LittleEndian, int32(v))
		if err != nil {
			panic(err)
		}
	}
}

func TestExternalSort_Sort(t *testing.T) {
	tmpfile := "abc.dat"
	generateTmpFile(tmpfile,33,100)
	printBinaryFileConent(tmpfile)
	es := NewExternalSort(tmpfile,8,3)
	es.Initialize()
	printBinaryFileConent("tmp0.dat","tmp1.dat")
	es.Sort()
	printBinaryFileConent("tmp0.dat")
}
