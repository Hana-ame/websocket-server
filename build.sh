#!/bin/bash
go build -o wss.exe

export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
# export CC=aarch64-linux-gnu-gcc

go build -o wss.bin