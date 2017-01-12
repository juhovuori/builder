#!/bin/bash

FILES="builder builder.hcl project.hcl version.json"


mkdir -p deploy

echo Copying files
for FILE in $FILES
do
  echo $FILE
  curl -o "deploy/$FILE" "https://s3.eu-central-1.amazonaws.com/juhovuori/builder/$FILE"
done

chmod 755 deploy/builder

mv deploy/* .
rmdir deploy

echo Restart builder
killall builder

