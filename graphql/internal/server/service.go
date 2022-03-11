package server

import (
	"context"
	"net/http"

	"github.com/cod3rcarl/wwdatabase-go-backend/graphql/client"

	"github.com/cod3rcarl/wwdatabase-go-backend/graphql/internal/server/graph/resolvers"

	gqlgenerated "github.com/cod3rcarl/wwdatabase-go-backend/graphql/internal/server/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Config struct {
	Port               string   `envconfig:"PORT" default:"8080"`
	CORSAllowedOrigins []string `envconfig:"CORS_ALLOWED_ORIGINS" default:"*"`
	CORSAllowedHeaders []string `envconfig:"CORS_ALLOWED_HEADERS" default:"Authorization,Content-Type"`
	CORSAllowedMethods []string `envconfig:"CORS_ALLOWED_METHODS" default:"GET,POST,HEAD,OPTIONS"`
}

type Server struct {
	logger  *zap.Logger
	handler http.Handler
	port    string
}

func NewServer(logger *zap.Logger, wwSrv *client.Client, cfg Config) (*Server, error) {
	resolver := &resolvers.Resolver{
		Logger: logger,
		Server: wwSrv,
	}
	srv := handler.NewDefaultServer(gqlgenerated.NewExecutableSchema(gqlgenerated.Config{Resolvers: resolver}))

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		logger.Error("server panic")

		return errors.Errorf("server panic")
	})

	s := Server{
		logger: logger,
		port:   cfg.Port,
	}
	s.handler = s.newRouter(srv, cfg)

	return &s, nil
}

func (s *Server) Start() error {
	s.logger.Info("starting server", zap.String("port", s.port))

	if err := http.ListenAndServe(":"+s.port, s.handler); err != nil {
		return errors.Errorf("server error: %v", err)
	}

	return nil
}
