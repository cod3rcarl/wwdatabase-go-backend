package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/cod3rcarl/wwdatabase-go-backend/graphql/client"
	server "github.com/cod3rcarl/wwdatabase-go-backend/graphql/internal/server"
	"github.com/joho/godotenv"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type App struct {
	// app control
	ctx      context.Context
	cancel   context.CancelFunc
	shutdown sync.Mutex

	// app dependencies
	logger   *zap.Logger
	server   *server.Server
	wwClient *client.Client
}

const (
	shutdownTimeout = 2 * time.Minute

	shutdownComplete       = 1
	shutdownTimeoutExpired = 2
	forcedExit             = 3
)

func NewApp() *App {
	c, cFn := context.WithCancel(context.Background())

	a := &App{
		ctx:    c,
		cancel: cFn,
	}

	return a
}

func (a *App) WithLogger() *App {
	// disable caller info in the logs
	l, err := zap.NewProduction(zap.WithCaller(false))
	if err != nil {
		log.Fatalf("failed to init logger; error=%v", err)
	}

	a.logger = l

	return a
}

func (a *App) WithClient() *App {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
	if a.logger == nil {
		log.Fatalf("logger is a server dependency and cannot be nil")
	}

	var cfg client.Config

	if err := envconfig.Process("", &cfg); err != nil {
		a.logger.Fatal("failed to parse client config", zap.Error(err))
	}
	c, err := client.NewClient(a.logger, cfg)
	if err != nil {
		log.Fatalf("client is a server dependency and cannot be nil")
	}
	a.wwClient = &c

	return a
}

func (a *App) WithServer() *App {
	if a.logger == nil {
		log.Fatalf("logger is a server dependency and cannot be nil")
	}

	if a.wwClient == nil {
		log.Fatalf("client is a server dependency and cannot be nil")
	}

	var cfg server.Config
	if err := envconfig.Process("", &cfg); err != nil {
		a.logger.Fatal("failed to parse server config", zap.Error(err))
	}

	s, err := server.NewServer(a.logger, a.wwClient, cfg)
	if err != nil {
		a.logger.Fatal("failed to init server", zap.Error(err))
	}

	a.server = s

	return a
}

func (a *App) Start() {
	if a.server != nil {
		if err := a.server.Start(); err != nil {
			a.logger.Error("server error: %v", zap.Error(err))
			a.Shutdown()
		}
	}
}

func (a *App) Shutdown() {
	// Shutdown can be called either when the server exited or an interrupt signal was received. The mutex will make
	// sure that no two goroutines will run the shutdown process at the same time (we hard exit at the end).
	a.shutdown.Lock()

	a.logger.Info("shutting down app")

	// let the ctx chain know that the app finished
	a.cancel()

	// flush the log buffer
	_ = a.logger.Sync()

	a.logger.Info("shutdown complete")

	a.shutdown.Unlock()

	os.Exit(shutdownComplete)
}

func (a *App) WithGracefulShutdown() *App {
	go func() {
		// create interrupt signal listener
		ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

		// block and listen for the interrupt signal
		<-ctx.Done()

		// notify user of shutdown
		a.logger.Info("shutting down gracefully, press Ctrl+C again to force")

		go func() {
			select {
			case <-time.After(shutdownTimeout):
				a.logger.Info("shutdown timeout expired")

				os.Exit(shutdownTimeoutExpired)
			case <-ctx.Done():
				a.logger.Info("forced exit")

				os.Exit(forcedExit)
			}
		}()

		a.Shutdown()
	}()

	return a
}

func (a *App) GetLogger() *zap.Logger {
	return a.logger
}

func (a *App) GetServer() *server.Server {
	return a.server
}
