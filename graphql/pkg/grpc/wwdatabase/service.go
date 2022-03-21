package client

import (
	"context"

	pb "github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/pkg/wwdatabase"
	"github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/storage"

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
	GetChampionByOrderNumber(ctx context.Context, req *pb.ChampionNumber) (*pb.ChampionResponse, error)
	GetChampionListByName(ctx context.Context, name string) (*pb.ChampionsList, error)
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
