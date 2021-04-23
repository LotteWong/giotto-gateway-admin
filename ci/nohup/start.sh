#!/bin/bash

# set go env
export GO111MODULE=auto
export GOPROXY=https://goproxy.io,direct
go mod tidy

# build binary executable
mkdir -p ../bin/giotto_gateway_admin
go build -o ../bin/giotto_gateway_admin

# kill processes already started
pkill -9 giotto_gateway_admin

# run admin backgroud server
nohup ../bin/giotto_gateway_admin -config ../config/prod/ >> logs/admin.log 2>&1 &
echo 'nohup ../bin/giotto_gateway_admin -config ./config/prod/ >> logs/admin.log 2>&1 &'
