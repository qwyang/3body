package examples
/*
a,b两值交换，不用第三者变量算法：
1.a=a+b,b=a-b,a=a-b
2.加法可改为异或操作a=a^b,b=a^b,a=a^b
 */
func swap(a,b int) (int,int) {
 	a=a+b
 	b=a-b
 	a=a-b
 	return a,b
}
func swap2(a,b int) (int,int) {
	a=a^b
	b=a^b
	a=a^b
	return a,b
}