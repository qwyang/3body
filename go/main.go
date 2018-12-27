package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/3body/go/algorithms"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type Person struct {
	Name string `json:"student_name"`
	age  int8
}

type Student2 struct {
	Person
	speciality string
}

type Integer int

func (a Integer) compare(b Integer) bool {
	return a < b
}

func (person *Person) PString() string {
	buffer := bytes.Buffer{}
	fmt.Fprintf(&buffer, "name:%s,age:%d", person.Name, person.age)
	return string(buffer.Bytes())
}

func (stu *Student2) String() string {
	buffer := bytes.Buffer{}
	fmt.Fprintf(&buffer, "name:%s,age:%d,speciality:%s", stu.Name, stu.age, stu.speciality)
	return string(buffer.Bytes())
}

var funcMap map[string]func() = make(map[string]func())

func Init() {
	funcMap["run"] = func() {
		fmt.Printf("run\n")
	}
	funcMap["fly"] = func() {
		fmt.Printf("fly\n")
	}
	funcMap["walk"] = func() {
		fmt.Printf("walk\n")
	}
}

type Animal interface {
	Fly()
	Run()
}

type Flyer interface {
	Fly()
}

type Bird struct{}

func (b *Bird) Fly() {
	fmt.Printf("I can fly...\n")
}

func (b *Bird) Run() {
	fmt.Printf("I can run...\n")
}

func Write(quit chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	time.Sleep(2)
	quit <- true
}

type Index struct{}

func (i *Index) ServeHTTP(response http.ResponseWriter, r *http.Request) {
	n, err := response.Write([]byte("index"))
	fmt.Println(n, err)
}

func main() {
	Init()
	param := flag.String("skill", "", "the skill")
	flag.Parse()
	if flag.Parsed() {
		if f, ok := funcMap[*param]; ok {
			f()
		} else {
			fmt.Printf("usage:%s --skill {run|fly|walk}\n", os.Args[0])
		}
	}
	c := "你"
	fmt.Printf("unicode:%x\n", '你')
	for i, v := range []byte(c) {
		fmt.Printf("byte%d:%X\n", i, v)
	}
	stu1 := Student2{Person{"zs", 25}, "math"}
	stu2 := new(Student2)
	stu2.age = 26
	stu2.Name = "ls"
	fmt.Printf("stu1:%#v\n", stu1)
	fmt.Printf("stu2:%#v\n", stu2)
	fmt.Printf("stu1:%s\n", stu1.String())
	fmt.Printf("person:%s\n", stu1.PString())
	var a Integer = Integer(1)
	var b Integer = Integer(2)
	fmt.Printf("a<b:%v\n", a.compare(b))
	bird := new(Bird)
	var animal Animal
	var flyer Flyer
	animal = bird
	animal.Fly()
	animal.Run()
	flyer = animal
	flyer.Fly()
	var v1 interface{} = 6
	switch v := v1.(type) {
	case float64:
		fmt.Printf("float64 value:%v\n", v)
	case int:
		fmt.Printf("int value:%v\n", v)
	default:
		fmt.Printf("unkown value:%v\n", v)
	}
	//quit := make(chan bool)
	//go Write(quit)
	//<- quit
	//
	//select {
	//case <- quit:
	//	fmt.Printf("quit\n")
	//case <- time.After(1e9):
	//	fmt.Printf("timeout 10s\n")
	//}

	//encoding/json
	//github.com/pquerna/ffjson/ffjson
	sm := make(map[string]float64)
	sm["a"] = 1.1
	sm["b"] = 1.2
	ret, err := json.Marshal(sm)
	if err == nil {
		fmt.Printf("json:%s\n", string(ret))
	}
	stu3 := Student2{Person: Person{
		"name", 23,
	},
		speciality: "math",
	}
	ret, err = json.Marshal(stu3)
	if err == nil {
		fmt.Printf("json:%s\n", string(ret))
	}
	var ustu interface{}
	json.Unmarshal(ret, &ustu)
	fmt.Printf("%v\n", ustu)
	md5ins := md5.New()
	md5ins.Write([]byte("abcdefg"))
	md5s := md5ins.Sum([]byte("x"))
	fmt.Printf("md5:%x\n", md5s)
	//net/http
	//http.HandleFunc("/hello",func(response http.ResponseWriter,r *http.Request){
	//	n,err := response.Write([]byte("hello,world"))
	//	fmt.Println(n,err)
	//})
	//err = http.ListenAndServe("127.0.0.1:8888",nil)
	//if err != nil{
	//	fmt.Println("server error.",err)
	//}
	//resp,err := http.Get("http://www.baidu.com")
	//if err!=nil {
	//	fmt.Printf("error:%v\n",err)
	//	os.Exit(1)
	//}
	//defer resp.Body.Close()
	//body,err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	resp, err := http.Post("http://www.baidu.com", "application/x-www-form-urlencoded", strings.NewReader("id=1"))
	//fmt.Printf("%v,%v\n",resp,err)
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%v\n", string(body))
	}
	reg := regexp.MustCompile("[a-zA-Z]{3}")
	results := reg.FindAllString("abc efg ijk", -1)
	fmt.Printf("%v\n", results)

	url := "https://movie.douban.com/subject/24751763/"
	resp, err = http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	rege := regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*)</span>`)
	matches := rege.FindAllStringSubmatch(string(html), -1)
	//fmt.Printf("%v\n",matches)
	for _, match := range matches {
		name := match[1]
		fmt.Printf("find movie name:%s\n", name)
	}
	rege2 := regexp.MustCompile(`<strong class="ll\s*rating_num"\s*property="v:average">(.*)</strong>`)
	matches2 := rege2.FindAllStringSubmatch(string(html), -1)
	for _, match := range matches2 {
		score := match[1]
		fmt.Printf("find movie name:%s\n", score)
	}
	arr := []int{3, 2, 1}
	algorithms.Qsort(arr)
	fmt.Printf("%v\n", arr)
	arr1 := arr[1:]
	arr1[0] = 99
	fmt.Printf("%v\n", arr)
}
