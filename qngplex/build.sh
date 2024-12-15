#!/bin/bash
# Build qngplex and put the binaries in ./bin
OLD_GOBIN=$GOBIN
export GOBIN="`pwd`/bin"

go get qngplex qngplex-server qngplex-prod qngplex-blocknotify
go install qngplex-server qngplex-prod qngplex-blocknotify

rm $GOPATH/src/qngplex -rf
rm $GOPATH/src/qngplex-* -rf
export GOBIN=$OLD_GOBIN