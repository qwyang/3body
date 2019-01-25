package algorithms

import (
	"fmt"
	"testing"
)
/*
典型例子：341=11*31,561=3*11*17
Carmichael数N的定义：导致与N互素的A^(P-1)==1.
 */
func TestWitness(t *testing.T) {
	var P uint32=1<<31-1
	var A uint32 = 2
	fmt.Printf("is P=%d(A=%d) Prime:%d\n",P,A,witness(A,P-1,P))
	P=341
	fmt.Printf("is P=%d(A=%d) Prime:%d\n",P,A,witness(A,P-1,P))
	A=3
	fmt.Printf("is P=%d(A=%d) Prime:%d\n",P,A,witness(A,P-1,P))
	P=561
	A=7
	fmt.Printf("is P=%d(A=%d) Prime:%d\n",P,A,witness(A,P-1,P))
	P=7*7*7*7*7
	A=7
	fmt.Printf("is P=%d(A=%d) Prime:%d\n",P,A,witness(A,P-1,P))
}
func TestIsPrime(t *testing.T) {
	count:=0
	for i:=2;i<100;i++{
		if IsPrime(uint32(i)) {
			count++
			fmt.Printf("is P=%d Prime:%v\n",i,true)
		}
	}
	fmt.Printf("total prime count:%d\n",count)
}