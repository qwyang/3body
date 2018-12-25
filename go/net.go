package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

/*
echo server:
	listen->accept->handle session:receive,send
client:
	connect->send->receive
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

func main(){
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
