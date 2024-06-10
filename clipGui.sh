#!/bin/bash

index=0
field=()
filepath="/home/matteo/programmazione/GoClipboard/GoHistory/dump.txt"
echo $filepath

#cat clipboardHistory.txt | while read line

while read line
do
	if [ $line == "#" ]; then
		read line1
		b=`echo "$line1" | base64 --decode`
		
		field+=("$b")  
	fi
done <"$filepath"


echo "${field[@]}"

returntosender=$(zenity \
	--list \
	--title "Clipboard History" \
	--column "content" \
	"${field[@]}" \
	--height=800 \
	--width=800
)

printf "$returntosender" | `xclip -sel c -i` 
exit
