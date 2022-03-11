package client

import (
	"fmt"

	pb "github.com/cod3rcarl/wwdatabase-go-backend/grpc/pkg/wwdatabase"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	GRPCServerHost string `envconfig:"GRPC_HOST" required:"true"`
	GRPCServerPort string `envconfig:"GRPC_PORT" required:"true"`
}

type Client struct {
	logger               *zap.Logger
	cc                   *grpc.ClientConn
	wwdatabaseGRPCClient pb.WwdatabaseClient
}

func NewClient(l *zap.Logger, cfg Config) (Client, error) {
	client := Client{
		logger: l,
	}

	ccon, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.GRPCServerHost, cfg.GRPCServerPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		client.logger.Error(err.Error(), zap.Error(err))

		return client, err
	}

	client.cc = ccon

	client.wwdatabaseGRPCClient = pb.NewWwdatabaseClient(client.cc)

	return client, nil
}

func (c *Client) Close() {
	if err := c.cc.Close(); err != nil {
		c.logger.Error("failed to close client", zap.Error(err))
	}
}