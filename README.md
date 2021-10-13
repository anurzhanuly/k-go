Устанавливаем protobuf compiler

`sudo snap install protobuf --classic`

Уставнавливаем зависимости и генераторы

`go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
google.golang.org/protobuf/cmd/protoc-gen-go \
google.golang.org/grpc/cmd/protoc-gen-go-grpc \`

Устанавливаем buf

`go install \
github.com/bufbuild/buf/cmd/buf@latest \
github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking@latest \
github.com/bufbuild/buf/cmd/protoc-gen-buf-lint@latest`

Для перегенерации proto файла надо запустить команду:

`buf generate`