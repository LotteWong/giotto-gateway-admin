#!/bin/sh

# update docker images
commit=`git rev-parse --short HEAD`
kubectl set image deployment/giotto-gateway-admin giotto-gateway-admin=giotto-gateway-admin:$commit