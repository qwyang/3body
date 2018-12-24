/*
string is immutable type.
 */
package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
)

func base64enc(data string)string{
	return base64.StdEncoding.EncodeToString([]byte(data))
}
func base64dec(data string)([]byte,error){
	return base64.StdEncoding.DecodeString(data)
}

func main(){
	a_str := "你好"
	b_str := "hello"
	var buf bytes.Buffer
	buf.WriteString(a_str)
	buf.WriteString(b_str)
	fmt.Fprintf(&buf,",my name is:%s","羊羊羊")
	r := fmt.Sprintf("score:%d",100)
	fmt.Fprintf(&buf,",%s",r)
	fmt.Printf("%v\n%v\n",buf.String(),buf.Bytes())
	b64e := base64enc(buf.String())
	fmt.Printf("%v\n",b64e)
	b64d,err := base64dec(b64e)
	if err == nil{
		fmt.Printf("%s\n",string(b64d))
	}
	pos := strings.Index(buf.String(),"hello")
	fmt.Println(buf.String()[pos:pos+len("hello")])
	v := strings.Split(buf.String(),",")
	for i,str := range v{
		fmt.Printf("%d:%s\n",i,str)
	}
}
