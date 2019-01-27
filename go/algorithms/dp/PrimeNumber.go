package dp

/*
素数测试:
费马小定理（素数的必要条件）：A^(N-1)%N==1,1<A<N.
witness:求A^(N-1)%N的余数
 */
func witness(a uint32,i uint32,N uint32)uint32{
	if i == 0{return 1}
	r := witness(a,i/2,N)
	if r == 0{return 0}
	var y uint64 = (uint64(r)*uint64(r))%uint64(N)
	if y==1 && r!=1 && r!=N-1{
		return 0
	}
	if i%2 == 0{
		return uint32(y)
	}else{
		y=(y*uint64(a))%uint64(N)
		return uint32(y)
	}
}
func IsPrime(P uint32) bool{
	var A uint32 = uint32(RandInt(2,int(P)))
	return witness(A,P-1,P)==1
}