
cd /home/www/golang/src/apollo/cloud/proto/pay;
protoc  --go-grpc_out=/home/www/golang/src/apollo/cloud/service/pay  --go_out=/home/www/golang/src/apollo/cloud/service/pay  notify/pay_notify_log.proto;
protoc-go-inject-tag -input=/home/www/golang/src/apollo/cloud/service/pay/notify/pay_notify_log.pb.go;