package main

import (
	"fmt"
	"sort"
)

type Student struct{
	Name string
	Score uint8
}

type Students []Student

//降序排序，成绩相同，按名字先后
func (s Students) Less(i,j int) bool {
	si,sj := s[i].Score,s[j].Score
	si_name,sj_name := s[i].Name,s[j].Name
	if si > sj {
		return true
	}else if si == sj {
		return si_name < sj_name
	}else{
		return false
	}
}

func (s Students) Swap(i,j int) {
	s[i],s[j] = s[j],s[i]
}

func (s Students) Len() int {
	return len(s)
}

func main(){
	stus := Students{
		Student{"qw",60},
		Student{"xm",90},
		Student{"aa",90},
		Student{"kt",85},
		Student{"lw",100},
	}
	sort.Sort(stus)
	fmt.Printf("%v\n",stus)
}