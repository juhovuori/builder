#!/bin/bash

FILES="builder builder.hcl project.hcl version.json"

echo Stopping builder
killall builder

echo Copying files
for FILE in $FILES
do
  echo $FILE
  curl -o "$FILE" "https://s3.eu-central-1.amazonaws.com/juhovuori/builder/$FILE"
done

echo Restarting
nohup ./builder server > builder.out 2> builder.err < /dev/null &
echo Done
