package client

import (
	"context"
	"github.com/sirupsen/logrus"
	pb "go.kolesa-team.org/gl/go-course/rpc"
	"google.golang.org/grpc"
)

// Конфиги для клиента gRPC
type GrpcOptions struct {
	ServerAddress string
}

// Структура клиента
type Client struct {
}

// Новый клиент
func NewGrpc(ctx context.Context, config GrpcOptions) (client pb.ServiceClient, err error) {
	connection, err := grpc.DialContext(ctx, config.ServerAddress, grpc.WithInsecure())

	if err != nil {
		logrus.WithError(err).Error("Невозможно установить TCP соединение с gRPC сервером")
		return
	}

	client = pb.NewServiceClient(connection)

	return
}
