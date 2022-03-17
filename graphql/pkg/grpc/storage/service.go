package storage

import (
	"context"

	"github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/models"
	"github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/pkg/pgx"
	"go.uber.org/zap"
)

type ServiceInterface interface {
	WWDatabaseService
	Stop()
}

type WWDatabaseService interface {
	GetAllChampions(ctx context.Context) (models.Champions, error)
	GetChampionListByName(ctx context.Context, name string) (models.Champions, error)
	AddChampion(ctx context.Context, input models.CreateChampionInput) (models.Champion, error)
	GetPreviousChampion(ctx context.Context) (models.Champion, error)
	DeleteChampion(ctx context.Context, id string) (string, error)
}

type Service struct {
	serviceOptions
}

type serviceOptions struct {
	baseOption
	loggerOption
}

func NewStorage(opts ...Option) *Service {
	svc := &Service{}

	return svc.withOptions(opts...)
}

func WithLogger(l *zap.Logger) Option {
	return loggerOption{logger: l}
}

func WithBase(b *pgx.Service) Option {
	return baseOption{b}
}
