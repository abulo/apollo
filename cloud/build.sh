#!/bin/bash
path=$(cd `dirname $0`; pwd)
# go-bindata -o=initial/view/view.go -pkg=view  ./view/...
gohash=$(git rev-parse HEAD)
date=$(date +"%Y-%m-%d %H:%M:%S")
flags="-X 'main.build=$date' -X 'main.hash=$gohash'"
go build --ldflags="$flags" -o $path/bin/server $path/cmd/server/server.go