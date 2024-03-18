
cd /home/www/golang/src/apollo/cloud/proto/system;
protoc  --go-grpc_out=/home/www/golang/src/apollo/cloud/service/system  --go_out=/home/www/golang/src/apollo/cloud/service/system  user/system_user_post.proto;
protoc-go-inject-tag -input=/home/www/golang/src/apollo/cloud/service/system/user/system_user_post.pb.go;