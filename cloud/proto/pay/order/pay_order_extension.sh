
cd /home/www/golang/src/apollo/cloud/proto/pay;
protoc  --go-grpc_out=/home/www/golang/src/apollo/cloud/service/pay  --go_out=/home/www/golang/src/apollo/cloud/service/pay  order/pay_order_extension.proto;
protoc-go-inject-tag -input=/home/www/golang/src/apollo/cloud/service/pay/order/pay_order_extension.pb.go;