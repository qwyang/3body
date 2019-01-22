package sort

import "testing"

func TestInsertSort(t *testing.T) {
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
	InsertSort2(stus)
	if !stus.equal(expect){
		t.Errorf("InsertSort2 Test failed, %v",stus)
	}
	t.Logf("InsertSort2 Test success,%v",stus)
}