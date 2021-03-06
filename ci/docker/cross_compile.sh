#!/bin/bash

# set go env
export GO111MODULE=auto
export GOPROXY=https://goproxy.io,direct
go mod tidy

# build binary executable
mkdir -p ./bin
GOOS=linux GOARCH=amd64 go build -o ./bin/giotto_gateway_admin

# build docker images
commit=`git rev-parse --short HEAD`
docker build -f ./ci/docker/Dockerfile -t giotto-gateway-admin:$commit .
