#!/bin/bash
targetlist=$(grep -r fraanx|cut -d ':' -f 1)

for fname in $targetlist
do
echo file is $fname
sed -i 's/fraanx/fraanx/g' $fname
done

