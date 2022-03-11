package client

import (
	"context"

	"github.com/cod3rcarl/wwdatabase-go-backend/grpc/internal/storage"
	pb "github.com/cod3rcarl/wwdatabase-go-backend/grpc/pkg/wwdatabase"

	"go.uber.org/zap"
)

const (
	ServiceName = "wwdatabase"
)

type Service struct {
	serviceOptions
}

type ServiceInterface interface {
	GetAllChampions(ctx context.Context, req *pb.GetChampionsRequest) (*pb.ChampionsList, error)
	AddChampion(ctx context.Context, req *pb.NewChampionData) (*pb.CreateChampionPayload, error)
	DeleteChampion(ctx context.Context, id string) (*pb.DeleteChampionResponse, error)
}

type serviceOptions struct {
	loggerOption
	storageOption
}

func NewService(opts ...Option) *Service {
	svc := &Service{}

	return svc.withOptions(opts...)
}

func WithLogger(l *zap.Logger) Option {
	return loggerOption{logger: l}
}

func WithStorage(s storage.ServiceInterface) Option {
	return storageOption{store: s}
}
