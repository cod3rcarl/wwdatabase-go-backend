package server

import (
	"github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/storage"
	ww "github.com/cod3rcarl/wwdatabase-go-backend/graphql/pkg/grpc/wwdatabase"

	"go.uber.org/zap"
)

type Option interface {
	apply(s *Service)
}

type optionFunc func(s *Service)

func (f optionFunc) apply(s *Service) {
	f(s)
}

func Options(opts ...Option) Option {
	return optionFunc(func(s *Service) {
		for _, opt := range opts {
			opt.apply(s)
		}
	})
}

func (s *Service) withOptions(opts ...Option) *Service {
	for _, opt := range opts {
		opt.apply(s)
	}

	return s
}

type loggerOption struct {
	logger *zap.Logger
}

func (o loggerOption) apply(s *Service) {
	s.logger = o.logger
}

type storageOption struct {
	store storage.ServiceInterface
}

func (o storageOption) apply(s *Service) {
	s.store = o.store
}

type clientOption struct {
	wwdatabase ww.ServiceInterface
}

func (o clientOption) apply(s *Service) {
	s.wwdatabase = o.wwdatabase
}
