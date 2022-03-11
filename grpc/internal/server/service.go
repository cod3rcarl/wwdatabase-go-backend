package server

import (
	"fmt"
	"net"

	"github.com/cod3rcarl/wwdatabase-go-backend/grpc/internal/storage"
	client "github.com/cod3rcarl/wwdatabase-go-backend/grpc/internal/wwdatabase"
	pb "github.com/cod3rcarl/wwdatabase-go-backend/grpc/pkg/wwdatabase"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const ServiceName = "grpc-server"

type Service struct {
	config Config
	server *grpc.Server
	serviceOptions
	pb.UnimplementedWwdatabaseServer
}

type serviceOptions struct {
	loggerOption
	storageOption
	clientOption
}

func NewServer(cfg Config, opts ...Option) *Service {
	srv := grpc.NewServer()

	// register the gRPC server for reflection to expose available endpoints
	reflection.Register(srv)

	svc := &Service{
		config: cfg,
		server: srv,
		serviceOptions: serviceOptions{
			loggerOption: loggerOption{
				logger: zap.NewNop(),
			},
		},
	}

	pb.RegisterWwdatabaseServer(srv, svc)

	return svc.withOptions(opts...)
}

func WithLogger(l *zap.Logger) Option {
	return loggerOption{logger: l}
}

func WithStorage(s storage.ServiceInterface) Option {
	return storageOption{store: s}
}

func WithClient(c client.ServiceInterface) Option {
	return clientOption{c}
}

func (s *Service) Stop() {
	s.server.GracefulStop()
}

func (s *Service) ListenAndServe() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.config.Host, s.config.Port))
	if err != nil {
		return errors.Errorf("failed to listen: %v", err)
	}

	s.logger.Info("starting grpc server", zap.String("address", lis.Addr().String()))

	if err = s.server.Serve(lis); err != nil {
		return errors.Errorf("failed to serve: %v", err)
	}

	return nil
}
