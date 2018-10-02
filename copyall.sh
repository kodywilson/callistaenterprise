#!/bin/bash
export GOOS=linux
export CGO_ENABLED=0

go get;go build -o accountservice-linux-amd64;echo built `pwd`

export GOOS=darwin

docker build -t kodywilson/accountservice .

docker service rm accountservice
docker service create --name=accountservice --replicas=1 -p=6767:6767 kodywilson/accountservice
