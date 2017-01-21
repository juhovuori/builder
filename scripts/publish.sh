#!/bin/bash

VERSION="$(git rev-parse HEAD)"
OPTS="--region=eu-central-1 --storage-class=STANDARD_IA --acl=public-read"
FILES="builder builder.hcl"

for FILE in $FILES
do
  echo $FILE
  aws s3 cp "$FILE" "s3://juhovuori/builder/$FILE" $OPTS

  echo $FILE-$VERSION
  aws s3 cp "$FILE" "s3://juhovuori/builder/$FILE-$VERSION" $OPTS
done
