package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("post", "http://www.163.com/", strings.NewReader("a=b"))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", "mozilla/5.0")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	buf := bytes.Buffer{}
	req.Write(&buf)
	fmt.Printf("%v######", buf.String())

	defer resp.Body.Close()
}
