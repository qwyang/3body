package algorithms

import (
	"container/heap"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
)

/*
外部排序算法实现，K路归并排序。
假设：内存每次最多排序M个数字。k路外部临时文件名称为tmp1.dat..tmpk.dat。
*/

type ExternalSort struct {
	filename string
	M int
	K int
	totalRecordCount int
	tmpfiles []string
}

func NewExternalSort(filename string,M int,K int)*ExternalSort{
	es := &ExternalSort{filename,M,K,0,nil}
	for i:=0;i<2*K;i++{
		filename := fmt.Sprintf("tmp%d.dat",i)
		f,err:=os.Create(filename)
		if err!=nil{panic(err)}
		err = f.Close()
		if err!=nil{panic(err)}
		es.tmpfiles = append(es.tmpfiles,filename)
	}
	return es
}

func readRecords(f *os.File,size int)(data []int32,err error){
	//r:=bufio.NewReader(f)
	var num int32
	for i:=0;i<size;i++{
		err = binary.Read(f,binary.LittleEndian,&num)
		if err!=nil{
			return
		}
		data = append(data,num)
	}
	fmt.Println("meet ",err,",last data:",data)
	return
}

func writeRecords(f *os.File,data []int32)(err error){
	for _,num := range data{
		err = binary.Write(f,binary.LittleEndian,num)
		if err!=nil{
			return err
		}
	}
	return
}
type Int32s []int32
func (arr Int32s)Less(i,j int)bool{return arr[i]<arr[j]}
func (arr Int32s)Len()int{return len(arr)}
func (arr Int32s)Swap(i,j int){arr[i],arr[j]=arr[j],arr[i]}
func (s *ExternalSort) Initialize(){
	f,err := os.Open(s.filename)
	if err!=nil{panic(err)}
	defer f.Close()
	dstFiles := s.tmpfiles[:s.K]
	var dstFds []*os.File
	for i := 0;i < s.K;i++ {//打开K路文件准备写数据
		fd,err := os.Create(dstFiles[i])
		if err!=nil{
			panic(err)
		}else{
			dstFds = append(dstFds,fd)
		}
	}
	totalRecordCount:=0
	dstFileNum:=0
	for{//初始化顺串
		data,err := readRecords(f,s.M)
		if err!=nil&&err!=io.EOF{
			panic(err)
		}
		if len(data)>0{
			sort.Sort(Int32s(data))
			err := writeRecords(dstFds[dstFileNum], data)
			if err != nil {
				panic(err)
			}
			totalRecordCount += len(data)
			dstFileNum = (dstFileNum + 1) % s.K
		}
		if err==io.EOF{
			break
		}
	}
	for i := 0;i < s.K;i++ { //关闭文件
		 dstFds[i].Close()
	}
	s.totalRecordCount = totalRecordCount
}
func printBinaryFileConent(files... string){
	for _,tmpfile := range files{
		fmt.Printf("file %s\n",tmpfile)
		f, _ := os.OpenFile(tmpfile, os.O_RDONLY, os.ModePerm)
		var data int32
		for {
			err := binary.Read(f, binary.LittleEndian, &data)
			if err == io.EOF {
				break
			}
			fmt.Print(data," ")
		}
		fmt.Println()
		f.Close()
	}
}
func (s *ExternalSort)Sort(){
	size := s.M
	srcFiles,dstFiles := s.tmpfiles[:s.K],s.tmpfiles[s.K:]
	for size < s.totalRecordCount{//归并序串,每次归并size大小的顺串,一个顺串大小超过记录总数，说明已经完全排好序了
		//fmt.Println("size:",size)
		var srcFds,dstFds []*os.File
		for i := 0;i < s.K;i++ {//打开K路文件准备读/写数据
			fd,err := os.Open(srcFiles[i])
			if err!=nil{
				panic(err)
			}else{
				srcFds = append(srcFds,fd)
			}
			fd,err = os.Create(dstFiles[i])
			if err!=nil{
				panic(err)
			}else{
				dstFds = append(dstFds,fd)
			}
		}
		j := 0
		merge_counter := 0
		for merge_counter*s.K*size < s.totalRecordCount{//一次归并K*size个数据，merge_counter记录归并次数
			MergeToDst(srcFds,size,dstFds[j])//K路文件分别读取size个数据进行归并,一次归并导致一个顺串扩大为：K*size
			merge_counter++
			j = (j+1)%s.K
		}
		for i := 0;i < s.K;i++ {//关闭文件
			err := srcFds[i].Close()
			if err!=nil{panic(err)}
			err = dstFds[i].Close()
			if err!=nil{panic(err)}
		}
		printBinaryFileConent(dstFiles...)
		size = s.K*size//完成一次归并，每个顺串增长K倍
		srcFiles,dstFiles = dstFiles,srcFiles //读写文件互换
	}
}

type LimitedReader struct {
	r io.Reader
	limit int
	index int
	eof int
}

func NewLimitedReader(fd io.Reader,limit int)*LimitedReader{
	return &LimitedReader{fd,limit,0,0}
}

func (lr *LimitedReader)GetInt()(int32,error){
	if lr.eof > 0 {
		return 0,io.EOF
	}
	if lr.index >= lr.limit {
		return 0,errors.New("limited access")
	}
	var data int32
	err := binary.Read(lr.r,binary.LittleEndian,&data)
	if err!=nil{
		if err==io.EOF{
			lr.eof = 1
			return 0,err
		}else{
			panic(err)
		}
	}else{
		lr.index++
	}
	fmt.Println("LimitedReader GetInt",data)
	return data,nil
}

type DataRef struct {
	data          int32
	limitedReader *LimitedReader
}

type MyHeap struct {
	array []*DataRef
}

func NewMyHeap()*MyHeap{
	return &MyHeap{nil}
}
func (h *MyHeap)Len()int{
	return len(h.array)
}
func (h *MyHeap)Swap(i,j int){
	h.array[i],h.array[j] = h.array[j],h.array[i]
}
func (h *MyHeap)Less(i,j int)bool{
	return h.array[i].data < h.array[j].data
}
func (h *MyHeap)Push(data interface{}){
	h.array = append(h.array,data.(*DataRef))
}
func (h *MyHeap)Pop()interface{}{
	size := len(h.array)
	if size > 0{
		n := size-1
		data := h.array[n]
		h.array = h.array[:n]
		return data
	}
	return nil
}
func MergeToDst(srcFds []*os.File,size int,dstFd *os.File){
	h := NewMyHeap()
	heap.Init(h)
	limitedReaders := make([]*LimitedReader,len(srcFds))
	for i:=0;i<len(srcFds);i++{
		limitedReaders[i]=NewLimitedReader(srcFds[i],size)
		data,err := limitedReaders[i].GetInt()
		if err == nil{
			heap.Push(h,&DataRef{data,limitedReaders[i]})
		}
	}
	//w := bufio.NewWriter(dstFd)
	for h.Len() > 0{
		min := heap.Pop(h).(*DataRef)
		fmt.Println("pop ",min.data)
		err := binary.Write(dstFd,binary.LittleEndian,min.data)
		if err!=nil{
			fmt.Println(err)
		}
		data,err := min.limitedReader.GetInt()
		if err!=nil{
			fmt.Println(err)
		}else{
			min.data = data
			heap.Push(h, min)
		}
	}
	//err := w.Flush()
	//if err!=nil{
	//	fmt.Println(err)
	//}
}