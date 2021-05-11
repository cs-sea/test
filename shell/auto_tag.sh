#!/bin/bash

msg=$1
echo $#
echo $msg

if [ -z $msg ]; then
   echo 'tag msg 必传'
   exit
fi

prevTag=$(git tag --sort=-tag | head -n 1)
datePreTag=$(git tag --sort=-taggerdate | head -n 1)

if [ $prevTag \< $datePreTag ]; then
    prevTag="$datePreTag"
fi

echo 'prevTag is ' $prevTag
array=(${prevTag//./ })

len=${#array[*]}
echo 'tag 级别 is' $len

i=$((len-1))
level=$((len-1))
max=9

while(( $i>0 ))
do
    value="${array[$i]}"
    if [ $value -lt $max ]; then
        let "array[$i]++"
        if [ $i -eq $level ]; then
            break
        fi
        array[$i+1]=0
        break
    fi

    let "i--"
done

newTag=''
for (( i = 0; i <= level; i++ )); do
    newTag="$newTag${array[$i]}"

    if [ "$i" -eq $level ]; then
        break
    fi
    newTag=$newTag"."
done

echo "$newTag"

if [ $newTag == $prevTag ]; then
    echo '请手动生成新的tag, 必须指定m 参数'
    exit
fi

git tag $newTag -m $msg
git push --tag