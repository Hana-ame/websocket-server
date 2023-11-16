#!/bin/bash
go build -o wss.exe

export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
# export CC="x86_64-linux-musl-gcc"
export CGO_ENABLED=1

go build -o wss.bin