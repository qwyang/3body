package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func runCmds(cmds []string)io.Reader{//不行，这是为什么？
	var buf1,buf2 bytes.Buffer
	buf_in,buf_out := &buf1,&buf2
	var comand *exec.Cmd
	//fmt.Printf("pid:%d\n",os.Getpid())
	for _,cmd := range cmds{
		args := strings.Fields(cmd)
		fmt.Println(args)
		comand = exec.Command(args[0],args[1:]...)
		comand.Stdin = buf_in
		fmt.Printf("#4,input:\n%s\n",buf_in.String())
		comand.Stdout = buf_out
		err := comand.Start()
		if err != nil{
			panic(fmt.Sprintf("#1:%v,%v\n",args,err))
			return nil
		}
		//fmt.Printf("pid:%d\n",os.Getpid())
		err = comand.Wait()
		if err != nil{
			switch err.(type){
				case *exec.ExitError:
					fmt.Printf("#2:%v,%v,state:%v\n",args,err,comand.ProcessState)
				default:
					panic(fmt.Sprintf("#3:%v,%v,state:%v\n",args,err,comand.ProcessState))
			}
		}
		//fmt.Printf("#4,output:\n%s\n",buf_out.String())
		buf_in,buf_out = buf_out,buf_in

	}
	return buf_in
}

func runCmds2(cmds []*exec.Cmd)io.Reader{
	var buf1,buf2 bytes.Buffer
	buf_in,buf_out := &buf1,&buf2
	var comand *exec.Cmd
	//fmt.Printf("pid:%d\n",os.Getpid())
	for _,comand = range cmds{
		comand.Stdin = buf_in
		//fmt.Printf("#4,input:\n%s\n",buf_in.String())
		comand.Stdout = buf_out
		err := comand.Start()
		if err != nil{
			panic(fmt.Sprintf("#1:%v,%v\n",comand,err))
			return nil
		}
		//fmt.Printf("pid:%d\n",os.Getpid())
		err = comand.Wait()
		if err != nil{
			switch err.(type){
			case *exec.ExitError:
				fmt.Printf("#2:%v,%v,state:%v\n",comand,err,comand.ProcessState)
			default:
				panic(fmt.Sprintf("#3:%v,%v,state:%v\n",comand,err,comand.ProcessState))
			}
		}
		//fmt.Printf("#4,output:\n%s\n",buf_out.String())
		buf_in,buf_out = buf_out,buf_in

	}
	return buf_in
}

func main(){
	//cmds := []string{"ps aux","grep signal","grep -v grep",`grep -v signal.go`,`awk "{print $2}"`}
	cmds := []*exec.Cmd{
		exec.Command("ps", "aux"),
		exec.Command("grep", "signal"),
		exec.Command("grep", "-v", "grep"),
		exec.Command("grep", "-v", "go run"),
		exec.Command("awk", "{print $2}"),
	}
	out := runCmds2(cmds)
	//out := runCmds(cmds)
	data,_ := ioutil.ReadAll(out)
	pid,_ := strconv.Atoi(strings.TrimSpace(string(data)))
	process,_ := os.FindProcess(pid)
	process.Signal(syscall.SIGINT)
}