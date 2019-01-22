package examples

import (
	"fmt"
	"testing"
)
/*
a,b长度不等
a,b长度相等且为0:a是nil，b是[],b是nil,a是[],a,b都是[],a,b都是nil
a,b长度相等，且不为0：a不是nil，b不是nil
 */
func compare(a,b[]byte)bool{
	if len(a) != len(b){
		return false
	}else{
		if (a==nil) != (b==nil){
			return false
		}else{
			for i,e := range a{
				if e!=b[i]{
					return false
				}
			}
		}
	}
	return true
}

func TestHanoi(t *testing.T) {
	hanoi(20,'A','B','C')
	data := hanoi_nr(5,'A','B','C')
	fmt.Println(string(data))
}
