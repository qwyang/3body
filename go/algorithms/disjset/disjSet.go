package disjset
/*
不相交集ADT
 */
type DisjSet []int

type DisjSetInterface interface {
	Union(i,j int)
	Find(i int) int
}

func NewDisjSet(n int)DisjSet{
	s := make([]int,n)
	for i:=0;i<len(s);i++{
		s[i]=-1
	}
	return s
}
/*
高度低的挂到高度高的
 */
func (s DisjSet)Union(r1, r2 int){
	if s[r1] > s[r2]{ //值越大，高度越低
		s[r1] = r2
	}else if s[r1] == s[r2]{
		s[r2] = r1
		s[r1] = s[r1]-1
	}else{
		s[r2] = r1
	}
}
/*
含路径压缩
 */
func (s DisjSet)Find(i int)int{
	if s[i] < 0 {return i}
	s[i] = s.Find(s[i])
	return s[i]
}