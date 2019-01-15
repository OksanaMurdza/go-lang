#!/bin/sh

go build
: ${TESTFILE=test.tmp}
TEMP=$1
CYCLES=$2
INPUT=$3
: ${TEMP:='*)'}
: ${CYCLES:=1}
: ${INPUT:=`curl https://lurkmore.to/Lisp`}
rm -rf $TESTFILE
i=0
while [ $i -lt $CYCLES ]
do
  echo $INPUT >> $TESTFILE
  i=`expr $i + 1`
done
./go-lang -cpuprofile prof.log -template $TEMP $TESTFILE
