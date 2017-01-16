#!/bin/bash

export GOPATH=$(pwd)
PATH=$GOPATH/bin:$PATH
go get github.com/jteeuwen/go-bindata/...
go get -d github.com/juhovuori/builder
cd src/github.com/juhovuori/builder
make build
make test-v
mv builder ~
killall builder # systemd should restart
