package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

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

func main(){
	SinImage("sin.png")
}