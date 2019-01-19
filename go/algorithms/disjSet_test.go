package algorithms

import (
	"fmt"
	"testing"
)
/*
(1,2),(3,4),(3,5),(1,7),(3,6),(8,9),(1,8),(3,10),(3,11),(3,12),(3,13)
 */
type EqualClass struct {
	a,b int
}
func TestNewDisjSet(t *testing.T) {
	s := NewDisjSet(14)
	l:=[]*EqualClass{
		&EqualClass{1,2},
		&EqualClass{3,4},
		&EqualClass{3,5},
		&EqualClass{1,7},
		&EqualClass{3,6},
		&EqualClass{8,9},
		&EqualClass{1,8},
		&EqualClass{3,10},
		&EqualClass{3,11},
		&EqualClass{3,12},
		&EqualClass{3,13},
	}
	for _,eq := range l{
		r1,r2 := s.Find(eq.a),s.Find(eq.b)
		if r1!=r2{
			fmt.Printf("a=%d,b=%d,a.root=%d,b.root=%d\n",eq.a,eq.b,r1,r2)
			s.Union(r1,r2)
		}
		fmt.Printf("s[%d]:%d,s[%d]:%d\n",r1,s[r1],r2,s[r2])
	}

}
