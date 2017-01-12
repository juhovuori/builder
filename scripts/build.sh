#!/bin/bash

export GOPATH=$(pwd)
go get -d github.com/juhovuori/builder
cd src/github.com/juhovuori/builder
make build
mv bin/builder ~
cp builder.hcl ~
cp project.hcl ~
cp version.json ~
killall builder # systemd should restart
