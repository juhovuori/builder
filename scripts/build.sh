#!/bin/bash

export GOPATH=$(pwd)
go get github.com/juhovuori/builder
cp bin/builder ~
cp src/github.com/juhovuori/builder/builder.hcl ~
cp src/github.com/juhovuori/builder/project.hcl ~
killall builder # systemd should restart
