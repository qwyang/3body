#!/bin/bash
#1到100之和的计算
#分别使用while，for循环
#1.for循环
sum=0
for ((i=1;i<=100;i=i+1));do
    sum=$((sum+i))
done
echo $sum
sum=0
for i in {1..100};do
    sum=$((sum+i))
done
echo $sum
sum=0
for i in $(seq 1 100);do
    sum=$((sum+i))
done
echo $sum
#2.while循环
sum=0
i=0
while [ "$i" -le 100 ];do
    sum=$((sum+i))
    i=$((i+1))
done
echo $sum
