package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"os"
	"strings"
)
/*
buffer io
 */
 func Buffer(){
 	var buf bytes.Buffer
 	buf.WriteString("第一滴血")
 	fmt.Fprintf(&buf,",电影：%s","first blood")
 	fmt.Println(buf.String())
 }
/*
生成一副正弦图300x300 pixel
1.生成一副空图
2.根据(x,y)设置灰度值
 */
func SinImage(filename string){
	const size = 300
	pic := image.NewGray(image.Rect(0,0,size,size))
	for x := 0;x<size;x++{
		for y:=0;y<size;y++{
			pic.SetGray(x,y,color.Gray{255})
		}
	}
	for x:=0;x<size;x++{
		angle:=float64(x)*2*math.Pi/size*2
		y := size/2*math.Sin(angle)
		y = -y + size/2
		pic.SetGray(x,int(y),color.Gray{0})
	}
	file,err := os.Create(filename)
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	err = png.Encode(file,pic)
	if err != nil{
		log.Fatal(err)
	}
}

/*
从ini文件读取指定键值
 */
func GetIniValue(filename,section,key string) string {
	var currSection string
	file,err := os.Open(filename)
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	fin := bufio.NewReader(file)
	for {
		line,err := fin.ReadString('\n')
		if err !=nil{
			if err == io.EOF{
				break
			}
			log.Fatal(err)
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		pair := strings.Split(line,"=")
		switch {
			case line[0] == '[' && line[len(line)-1] == ']':
				currSection = line[1:len(line)-1]
			case len(pair) == 2:
				currkey := strings.TrimSpace(pair[0])
				value := strings.TrimSpace(pair[1])
				if currSection == section && currkey == key {
					return value
				}else{
					fmt.Printf("%s.%s:%s\n",currSection,currkey,value)
				}
			default:
				log.Println("unkonwn options,content:",line)
		}

	}
	return "NotFound"
}
func main(){
	Buffer()
	SinImage("sin.png")
	v := GetIniValue("/home/qwyang/.gitconfig","core","editor")
	fmt.Printf("core.editor:%s\n",v)
}