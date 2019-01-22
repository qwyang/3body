package sort

import (
	"testing"
)

type Student struct{
	Name string
	Score uint8
}

type Students []Student

func (self Students) equal(s Students) bool {
	if len(self) != len(s){
		return false
	}
	if (self != nil) != (s != nil) {
		return false
	}
	for i,stu := range self {
		if !stu.equal(&s[i]){
			return false
		}
	}
	return true
}

func (self *Student) equal(s *Student) bool{
	return self.Name == s.Name && self.Score == s.Score
}

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
func TestQsort2(t *testing.T) {
	stus := Students{
		Student{"qw",60},
		Student{"xm",90},
		Student{"aa",90},
		Student{"kt",85},
		Student{"lw",100},
	}
	expect := Students{
		Student{"lw",100},
		Student{"aa",90},
		Student{"xm",90},
		Student{"kt",85},
		Student{"qw",60},
	}
	Qsort2(stus)
	if !stus.equal(expect){
		t.Errorf("Test failed, %v",stus)
	}
	t.Logf("Test success,%v",stus)
}
