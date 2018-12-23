fn_usage(){
   echo "$0 {dir} {file_suffix}"
   echo "example:$0 /tmp c"
   exit 1
}
if [ $# -ne 2 ];then
   fn_usage
fi
CodeDir=$1
CodeType=$2
total_count=0
space_count=0
note_count=0
tmp_count=0
#find $CodeDir | grep -E "\.$CodeType$"
for file in `find $CodeDir | grep -E "\.$CodeType$"`;do
   if [ -f $file ];then
       tmp_count=`cat $file | wc -l`
       ((total_count+=$tmp_count))
       tmp_count=`cat $file | sed -n '/^$/p' | wc -l`
       ((space_count+=$tmp_count))
       tmp_count=`cat $file | sed -n '/^\s*\/\//p' | wc -l`
       ((note_count+=$tmp_count))
#       echo "// count:$tmp_count"
       if [ "$CodeType" == "sh" -o "$CodeType" == "py" ];then
           tmp_count=`cat $file | sed -n '/^\s*#.*$/p' | wc -l`
           ((note_count+=$tmp_count))
#           echo "# count:$tmp_count"
       fi
       tmp_count=`cat $file | sed -n '/^\s*\/\*.*$/p;/^\s*\*.*$/p' | wc -l`
       ((note_count+=$tmp_count))
#       echo "/**/ count:$tmp_count"
   fi
done
echo "total_count:$total_count"
echo "space_count:$space_count"
echo "note_count:$note_count"
echo "total valid count:$((${total_count}-${space_count}-${note_count}))"
