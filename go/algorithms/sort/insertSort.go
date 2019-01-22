package sort

func InsertSort1(data []int){
	for i:=1;i<len(data);i++{
		for j:=i-1;j>=0 && data[j+1] < data[j];j--{//寻找插入位置，j为小于p的最大下标，j+1为插入位置
			data[j],data[j+1] = data[j+1],data[j]
		}
	}
}
