
cd /home/www/golang/src/apollo/cloud/proto/system;
protoc  --go-grpc_out=/home/www/golang/src/apollo/cloud/service/system  --go_out=/home/www/golang/src/apollo/cloud/service/system  dict/system_dict_type.proto;
protoc-go-inject-tag -input=/home/www/golang/src/apollo/cloud/service/system/dict/system_dict_type.pb.go;