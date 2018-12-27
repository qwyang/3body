package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
echo server:
	listen->accept->handle session:receive,send
client:
	telnet:connect->send->receive
*/

func EchoServer(handler func(conn net.Conn,quit chan bool),quit chan bool){
	l,err := net.Listen("tcp","127.0.0.1:8888")
	if err != nil{
		log.Fatal(err)
	}
	defer l.Close()
	for{
		c,err := l.Accept()
		if err!=nil{
			log.Println(err)
			continue
		}
		go handler(c,quit)
	}
}

func EchoTest(){
	quit := make(chan bool)
	go EchoServer(func(conn net.Conn,quit chan bool){
		defer conn.Close()
		r := bufio.NewReader(conn)
		for {
			str,err := r.ReadString('\n')
			if err!=nil {
				goto ERROR
			}
			if strings.TrimSpace(str) == "@close" {
				break
			}else if strings.TrimSpace(str) == "@shutdown"{
				quit <- true
				break
			}else{
				_,err := conn.Write([]byte(str))
				if err != nil{
					goto ERROR
				}
			}
		}
	ERROR:
		return
	},quit)
	//wait
	<- quit
}
/*
目的：TCP粘包问题测试
设计：Packet{Size,Body},connector（客户端）/acceptor（服务器），辅助函数：writePacket，readPacket
测试:TestCount:10000，address:127.0.0.1:8010
流程:acceptor routine接收数据，connector routine发送数据，main routine等待服务器结束。
 */
 type Packet struct {
 	Size uint16
 	Body []byte
 }
 /*
 先获取Packet大小，根据大小分配空间，再读取Packet内容。ReadFull一直等待数据填充完成为止，否则不返回。
  */
 func readPacket(r io.Reader)(*Packet,error){
 	var s = make([]byte,2)
 	_,err := io.ReadFull(r,s)
 	if err != nil{return nil,err}
 	p := &Packet{}
 	err = binary.Read(bytes.NewReader(s),binary.LittleEndian,&p.Size)
 	if err!=nil{return nil,err}
 	p.Body = make([]byte,p.Size)
 	_,err = io.ReadFull(r,p.Body)
 	if err!=nil{return nil,err}
 	return p,nil
 }
 /*
 利用bytes.Buffer将packet整理成字节流，然后发送到socket
  */
 func writePacket(w io.Writer,data[]byte)error{
	buf := bytes.Buffer{}
	err := binary.Write(&buf,binary.LittleEndian,uint16(len(data)))
	if err!=nil {return err}
	buf.Write(data)
	_,err = w.Write(buf.Bytes())
	if err != nil{return err}
	return nil
 }
/*
Acceptor:Stop/Wait/Start
 */
 type Acceptor struct{
 	l net.Listener
 	wg sync.WaitGroup
 	onSessionData func([]byte) bool
 }

 func NewAcceptor()*Acceptor{
 	server := new(Acceptor)
 	return server
 }

func handleSession(r io.ReadCloser,callback func([]byte)bool){
	defer r.Close()
	for {
		p,err := readPacket(r)
		if err!=nil || !callback(p.Body){break}
	}
}

func (s *Acceptor)Start(address string){
	var err error
	s.l,err = net.Listen("tcp",address)
	if err != nil{panic("server start failed.")	}
	s.wg.Add(1)
	defer s.wg.Done()
	for {
		conn,err := s.l.Accept()
		if err != nil{
			log.Printf("server:%v\n",err)
			break
		}
		go handleSession(conn,s.onSessionData)
	}
	log.Printf("Acceptor finished.\n")
}

func (s *Acceptor) Close(){
	s.l.Close()
}

 func (s *Acceptor)Wait(){
 	s.wg.Wait()
 }
/*
Connector:SendPackets
 */
func Connector(address string,sendCount int){
	conn,err:=net.Dial("tcp",address)
	if err!=nil{log.Printf("Net.Dial err:%v\n",err);return;}
	defer conn.Close()
	for i:=1;i<=sendCount;i++{
		data := strconv.Itoa(i)
		err := writePacket(conn,[]byte(data))
		if err != nil{
			log.Printf("WritePacket err:%v\n",err)
			break
		}
	}
	log.Printf("SendCount:%d\n",sendCount)
}


func main(){
	var SendTimes = 10000
	server := NewAcceptor()
	server.onSessionData = func(data []byte) bool {
		n,err := strconv.Atoi(string(data))
		if err != nil{return false}
		if n == SendTimes {
			log.Printf("RecivedTimes %d Ok, Closing Server Now\n",n)
			server.Close()
			return false
		}
		//log.Println("received data",n)
		return true
	}
	address := "127.0.0.1:8010"
	go server.Start(address)
	time.Sleep(time.Second*2)
	go Connector(address,SendTimes)
	server.Wait()
	log.Printf("main routine exit.\n")
}
