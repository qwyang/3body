#!/bin/bash
#演示case/esac/function用法

function printit(){
    if test -n "$1" ;then
        echo "you input $1"
    fi
}

input=$1
case "$input" in
    1)
        printit 1
        ;;
    2)
        printit 2
        ;;
    3)
        printit 3
        ;;
    *)
        printit "invalid choice"
        ;;
esac
