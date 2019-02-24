#!/bin/bash
#计算剩余退休时间
#输入一个退休时间，正则校验，转为秒。
#减去当前时间
#计算剩余天数
read -p "input date:" date_in
date_m=$(echo $date_in | grep '[0-9]\{8\}')
if [ -z "$date_m" ];then
    echo "wrong input, please input a date."
    exit 0
fi
date_s=$(date -d "$date_m" +%s)
date_n=$(date +%s)
seconds=$((date_s-date_n))
days=$((seconds/60/60/24))
echo "remaining days $days"
