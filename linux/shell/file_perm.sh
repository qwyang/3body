#!/bin/bash
#input a file and output the file type and permission of user.
#变量要用双引号包起来，否则包含空格的变量展开会导致意外。
read -p "input filename:" -t 10 filename
test -z $filename && echo "must input a filename" && exit 0
test -e $filename || { echo "file not exist";exit 0;}
test -f $filename && filetype="normal file"
test -d $filename && filetype="directory"
test -x $filename && fileperm="exutable"
test -r $filename && fileperm="$fileperm readable"
test -w $filename && fileperm="$fileperm writable"

test -z "$filetype" && filetype="unknown"
echo "filetype:$filetype"
test -z "$fileperm" && fileperm="none"
echo "fileperm:$fileperm"

exit 0
