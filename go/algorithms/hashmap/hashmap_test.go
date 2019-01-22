package hashmap

import (
	"testing"
)

func TestHashMap(t *testing.T) {
	m := NewMyMap()
	var key string
	for i:=0;i<16;i++{
		key += "a"
		m.Put(key,i)
	}
	for i:=0;i<10;i++{
		key = "a"
		m.Put(key,9999)
	}
	m.Traverse()
    key=""
	for i:=0;i<16;i++{
		key += "a"
        t.Logf("Get %s:%v\n",key,m.Get(key))
	}
	key=""
	for i:=0;i<16;i++{
		key += "a"
		t.Logf("Delete %s:%v\n",key,m.Delete(key))
	}
	key=""
	for i:=0;i<16;i++{
		key += "a"
		t.Logf("Get %s:%v\n",key,m.Get(key))
	}
	m.Traverse()
}
