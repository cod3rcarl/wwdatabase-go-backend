package main

import (
	"fmt"
	"log"

	"github.com/cod3rcarl/wwdatabase-go-backend/grpc/internal/server"
	"github.com/cod3rcarl/wwdatabase-go-backend/grpc/internal/storage"
	ww "github.com/cod3rcarl/wwdatabase-go-backend/grpc/internal/wwdatabase"
	"github.com/cod3rcarl/wwdatabase-go-backend/grpc/pkg/logger"
	"github.com/cod3rcarl/wwdatabase-go-backend/grpc/pkg/pgx"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func CreateLogger() *zap.Logger {
	cfg, err := logger.ReadConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to read logger config, error=%v", err))
	}

	l, err := logger.NewLogger(cfg)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to create logger, error=%v", err))
	}

	return l
}

func CreateStorage(l *zap.Logger) *storage.Service {
	cfg, err := pgx.ReadConfig()
	if err != nil {
		l.Fatal(
			"failed to read config",
			zap.String("service", "pgx-service"),
			zap.Error(err),
		)
	}

	base, err := pgx.NewPgx(cfg)
	if err != nil {
		l.Fatal(
			"failed to create service",
			zap.String("service", "pgx-service"),
			zap.Error(err),
		)
	}

	return storage.NewStorage(storage.WithBase(base), storage.WithLogger(l))
}

func CreateGrpcServer(
	l *zap.Logger,
	storeService storage.ServiceInterface,
	c ww.ServiceInterface,
) *server.Service {
	cfg, err := server.ReadConfig()
	if err != nil {
		l.Fatal(
			"failed to read config",
			zap.String("service", server.ServiceName),
			zap.Error(err),
		)
	}

	return server.NewServer(
		cfg,
		server.WithLogger(l),
		server.WithStorage(storeService),
		server.WithClient(c),
	)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	l := CreateLogger()
	store := CreateStorage(l)

	wwdatabaseService := ww.NewService(
		ww.WithLogger(l),
		ww.WithStorage(store),
	)
	grpcSrv := CreateGrpcServer(l, store, wwdatabaseService)

	if err := grpcSrv.ListenAndServe(); err != nil {
		l.Fatal("error running grpc server server", zap.Error(err))
	}
}
