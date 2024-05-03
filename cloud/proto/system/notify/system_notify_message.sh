
cd /home/www/golang/src/apollo/cloud/proto/system;
protoc  --go-grpc_out=/home/www/golang/src/apollo/cloud/service/system  --go_out=/home/www/golang/src/apollo/cloud/service/system  notify/system_notify_message.proto;
protoc-go-inject-tag -input=/home/www/golang/src/apollo/cloud/service/system/notify/system_notify_message.pb.go;