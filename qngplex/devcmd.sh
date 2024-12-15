#!/bin/bash
# Copy packages in $GOPATH, and run the given command directly
# e.g. ./devcmd server to run ./qngplex-server
rm $GOPATH/src/qngplex -rf
rm $GOPATH/src/qngplex-* -rf
cp -r ./pkg $GOPATH/src/qngplex
cp -r ./cmd/qngplex-$1 $GOPATH/src/qngplex-$1
go get qngplex qngplex-$1
if [ $? -eq 0 ]; then
    go run cmd/qngplex-$1/qngplex-$1.go
fi