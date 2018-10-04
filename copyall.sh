#!/bin/bash
export GOOS=linux
export CGO_ENABLED=0

go get;go build -o accountservice-linux-amd64;echo built `pwd`

cd healthchecker;go get;go build -o healthchecker-linux-amd64;echo built `pwd`;cd ..

export GOOS=darwin

cp healthchecker/healthchecker-linux-amd64 .

docker build -t kodywilson/accountservice .

docker service rm accountservice
docker service create --name=accountservice --replicas=1 --network=my_network -p=6767:6767 kodywilson/accountservice
