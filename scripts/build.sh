#!/bin/bash

export GOPATH=$(pwd)
PATH=$GOPATH/bin:$PATH

echo Cloning
builder clone src/github.com/juhovuori/builder
cd src/github.com/juhovuori/builder

echo Downloading dependencies
go get github.com/jteeuwen/go-bindata/...
go get -d ./...

echo Building
make build

echo Testing
make test-v 2>&1 | builder add-stage test-results

echo Deploying
mv builder ~
builder shutdown # graceful shutdown => systemd restarts
