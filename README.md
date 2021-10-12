Устанавливаем protobuf compiler

`sudo snap install protobuf --classic`

Уставнавливаем зависимости и генераторы

`github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway`

`github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2`

`google.golang.org/protobuf/cmd/protoc-gen-go`

`google.golang.org/grpc/cmd/protoc-gen-go-grpc`

Создаем proto file:

`syntax = "proto3";
package rpc;

message StringMessage {
string value = 1;
}

service YourService {
    rpc Echo(StringMessage) returns (StringMessage) {}
}`

protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
rpc/service.proto

protoc -I . --grpc-gateway_out ./gen/go \
--grpc-gateway_opt logtostderr=true \
--grpc-gateway_opt paths=source_relative \
--grpc-gateway_opt generate_unbound_methods=true \
`

protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
--grpc-gateway_out . \
--grpc-gateway_opt logtostderr=true \
--grpc-gateway_opt paths=source_relative \
--grpc-gateway_opt generate_unbound_methods=true \
rpc/service.proto

protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
--grpc-gateway_out . \
--grpc-gateway_opt logtostderr=true \
--grpc-gateway_opt paths=source_relative \
--grpc-gateway_opt generate_unbound_methods=true \
rpc/service.proto

go install \
github.com/bufbuild/buf/cmd/buf@latest \
github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking@latest \
github.com/bufbuild/buf/cmd/protoc-gen-buf-lint@latest

buf mod init
buf mod update