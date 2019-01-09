package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter,req *http.Request){
	fmt.Fprintf(w,"hello host:%s,path:%s,header:%v",req.Host, req.URL.Path,req.Header)
}

func main(){
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080",nil)
}
