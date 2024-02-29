#!/bin/bash
path=$(cd `dirname $0`; pwd)
# go-bindata -o=initial/view/view.go -pkg=view  ./view/...
gohash=$(git rev-parse HEAD)
date=$(date +"%Y-%m-%d %H:%M:%S")
flags="-X 'main.BuildTime=$date' -X 'main.BuildVersion=$gohash'"
go build --ldflags="$flags" -o $path/bin/hertz $path/cmd/hertz/hertz.go
go build --ldflags="$flags" -o $path/bin/grpc $path/cmd/grpc/grpc.go