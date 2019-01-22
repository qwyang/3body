package algorithms

import (
	"fmt"
)

const (
	MinBuketSize    =16
	HightLoadFactor =0.75
	LowLoadFactor = 0.25
)
/*
第一层设计：Pair(key,value)
 */
type Pair struct {
	key string
	value interface{}
	hash uint64
	next *Pair
}
/*
第二层设计：桶，链表,带有头节点的链表。
 */
type bucket struct {
	head *Pair
	size uint64
}

func NewBucket()*bucket{
	return &bucket{head:&Pair{},size:0}
}
func (b *bucket)Get(key string) *Pair{
	for pair := b.head.next; pair !=nil; pair = pair.next{
		if pair.key == key{
			return pair
		}
	}
	return nil
}
/*
添加情况：
1.元素存在，替换值
2.元素不存在，表头插入
 */
func (b *bucket)Put(pair *Pair) bool{
	target := b.Get(pair.key)
	if target != nil{
		target.value = pair.value
		return false
	}else{//表头插入
		pair.next = b.head.next
		b.head.next = pair
		b.size++
		return true
	}
}
/*
删除情况：
1.不存在，忽略
2.存在，删除
 */
func (b*bucket)Delete(key string)*Pair{
	target := b.Get(key)
	if target == nil{return nil}
	var p *Pair
	for p=b.head; p!=nil;p=p.next{
		if p.next == target{
			break
		}
	}
	p.next = target.next
	b.size--
	return target
}
/*
第三层设计：Segment
 */
type Map interface {
	Put(key string,elem interface{})
	Get(key string) interface{}
	Delete(key string) interface{}
	Redistribute()
}

type MyMap struct {
	buckets []*bucket
	buckets_len uint64
	total_count uint64
}

func NewMyMap()*MyMap{
	m := &MyMap{}
	m.buckets = make([]*bucket,MinBuketSize)
	m.buckets_len = MinBuketSize
	for i:=0;i<len(m.buckets);i++{
		m.buckets[i] = NewBucket()
	}
	return m
}

func (m *MyMap)Put(key string,elem interface{}){
	keyhash := hash(key)
	m.PutWithHash(key,keyhash,elem)
}

func (m *MyMap)PutWithHash(key string,hash uint64,elem interface{}){
	p := &Pair{key,elem,hash,nil}
	b := m.buckets[hash % m.buckets_len]
	ok := b.Put(p)
	if ok {
		m.total_count++
		m.Redistribute()
	}
	if b.size > 1 {
		fmt.Println("bucket:",hash % m.buckets_len,"size:",b.size)
	}
}

func (m *MyMap)Redistribute(){
    curretFacor := float64(m.total_count)/float64(m.buckets_len)
	if curretFacor >= HightLoadFactor {
		m.buckets_len <<= 1
	}else if curretFacor <= LowLoadFactor && m.buckets_len > MinBuketSize{
		m.buckets_len >>= 1
	}else{
		return
	}

	all_pairs := []*Pair{}
	for _,b := range m.buckets {
		for pair := b.head;pair != nil;pair = pair.next{
			all_pairs = append(all_pairs,pair)
		}
	}
	m.buckets = make([]*bucket,m.buckets_len)
	for i := range m.buckets{
		m.buckets[i] = NewBucket()
	}
	for _,p := range all_pairs{
        b := m.buckets[p.hash % m.buckets_len]
        b.Put(p)
	}
}

func (m *MyMap)Get(key string) interface{}{
	keyhash := hash(key) % m.buckets_len
	b:=m.buckets[keyhash]
	for l := b.head;l!=nil;l=l.next{
		if l.key == key {
			return l.value
		}
	}
	return nil
}
/*
没有该key值
有该key值
 */
func (m *MyMap)Delete(key string)interface{}{
	b := m.buckets[hash(key) % m.buckets_len]
	target := b.Get(key)
	if target == nil{
		return nil
	}
	b.Delete(key)
	m.total_count--
	m.Redistribute()
	return target.value
}

func (m *MyMap)Traverse(){
    fmt.Printf("total count:%d,bucket length:%d\n",m.total_count,m.buckets_len)
    for i, b :=range m.buckets{
		fmt.Printf("bucket %d elemcount:%d\n",i,b.size)
		for pair := b.head.next;pair != nil;pair = pair.next{
			fmt.Println(pair.key,pair.value)
		}
	}
}
