package client

import (
	"context"
	"github.com/sirupsen/logrus"
	pb "go.kolesa-team.org/ba/go-course/rpc"
	"google.golang.org/grpc"
)

func NewClient(ctx context.Context, addr string) (pb.ServiceClient, error) {
	conn, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())

	if err != nil {
		logrus.WithError(err).Error("Фатальная ошибка при подключении клиента")

		return nil, err
	}

	client := pb.NewServiceClient(conn)

	return client, err
}
