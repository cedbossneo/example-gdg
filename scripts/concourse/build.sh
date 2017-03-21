#!/bin/bash
# Build
go get -v gopkg.in/gin-gonic/gin.v1
go build -v -o ../built/app
cp scripts/deploy/Dockerfile ../built
