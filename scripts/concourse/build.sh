#!/bin/bash
# Build
go get -v gopkg.in/gin-gonic/gin.v1
GOOS=linux GOARCH=amd64 go build -v -o ../built/app
cp scripts/deploy/Dockerfile ../built
