#!/bin/bash
for i in $(cat cl)
do
	echo -n "$i "
	cd $i
	go build .	
	cd ..
done
echo
