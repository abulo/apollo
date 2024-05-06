
cd /home/www/golang/src/apollo/cloud/proto/system;
protoc  --go-grpc_out=/home/www/golang/src/apollo/cloud/service/system  --go_out=/home/www/golang/src/apollo/cloud/service/system  mail/system_mail_template.proto;
protoc-go-inject-tag -input=/home/www/golang/src/apollo/cloud/service/system/mail/system_mail_template.pb.go;