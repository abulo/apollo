#!/bin/bash
path=$(cd `dirname $0`; pwd)
go-bindata-assetfs  -o=initial/asset/asset.go -pkg=asset  resource/...
rm -rf bindata.go