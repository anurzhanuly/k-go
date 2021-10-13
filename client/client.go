package client

import (
	"context"
	"github.com/sirupsen/logrus"
	pb "go.kolesa-team.org/ba/go-course/rpc"
	"google.golang.org/grpc"
)

type Client struct {
	Addr string
}

func NewClient(addr string) *Client {
	return &Client{
		Addr: addr,
	}
}

func (c *Client) Client(ctx context.Context) (pb.ServiceClient, error) {
	conn, err := grpc.DialContext(ctx, c.Addr, grpc.WithInsecure())

	if err != nil {
		logrus.WithError(err).Error("Фатальная ошибка при подключении клиента")

		return nil, err
	}

	client := pb.NewServiceClient(conn)

	return client, err
}
