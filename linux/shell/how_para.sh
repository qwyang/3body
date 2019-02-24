#!/bin/bash
#程序的文件名是什么
#如果参数个数小于2则告知用户参数数量过少
#共有几个参数
#全部参数内容是什么
#第一个参数是什么
#第二个参数是什么

if [ $# -lt 2 ];then
    echo "param number too few."
else
    printf "total parameter number:%d\n" $#
    printf "all params:%s\n" "$*"
    printf "parm1=%s,param2=%s\n" $1 $2
    shift 2
    printf "total paramter numbers is:%d after shift 2\n" $#
fi
