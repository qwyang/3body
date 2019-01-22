package sort

func InsertSort2(data SortInterface){
	for i:=1;i<data.Len();i++{
		for j:=i-1;j>=0 && data.Less(j+1,j);j--{//寻找插入位置，j为小于p的最大下标，j+1为插入位置
			data.Swap(j,j+1)
		}
	}
}
