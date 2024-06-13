#!/bin/bash
# 当前文件夹
path=$(cd `dirname $0`; pwd);

# echo $path;
# 获取下级目录proto下的所有proto文件
protoFiles=$(ls $path/proto/*.proto);

# 遍历proto文件
for file in $protoFiles ; do
    # echo $file;
    # 获取文件名
    fileName=$(basename $file);
    # echo $fileName;
    # protoc  --go-grpc_out=$path/service  --go_out=$path/service  $currentPath/pagination.proto;
    protoc -I=proto  --go_out=../ --go-grpc_out=../  $fileName;
    # protoc-go-inject-tag -input=$path/service/$(basename $file .proto)/$(basename $file .proto).pb.go;
done

#  遍历 service目录包括子目录下的所有pb.go文件, 但是要将_grpc.pb.go文件排除
pbFiles=$(find $path/service -name "*.pb.go" | grep -v "_grpc.pb.go");
for file in $pbFiles ; do
    protoc-go-inject-tag -input=$file;
done