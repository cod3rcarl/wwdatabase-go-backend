package pgx

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var ErrConnectionString = errors.New("a database DSN must be used, URLs are not allowed")

type Service struct {
	Pool *pgxpool.Pool
	serviceOptions
}

type serviceOptions struct {
	loggerOption
}

func NewPgx(cfg Config, opts ...Option) (*Service, error) {
	svc := &Service{
		serviceOptions: serviceOptions{
			loggerOption: loggerOption{
				Logger: zap.NewNop(),
			},
		},
	}

	svc.withOptions(opts...)

	// # Example DSN
	// user=jack password=secret host=pg.example.com port=5432 dbname=mydb sslmode=verify-ca pool_max_conns=10
	//
	// # Example URL
	// postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca&pool_max_conns=10
	if strings.HasPrefix(cfg.DatabaseDSN, "postgres://") || strings.HasPrefix(cfg.DatabaseDSN, "postgresql://") {
		return nil, ErrConnectionString
	}

	poolConfig, err := pgxpool.ParseConfig(cfg.DatabaseDSN)
	if err != nil {
		return nil, errors.Errorf("failed to parse database config: %v", err)
	}

	// increase log level to warn, otherwise every query is logged at info level
	poolConfig.ConnConfig.Logger = zapadapter.NewLogger(svc.Logger.WithOptions(zap.IncreaseLevel(zap.WarnLevel)))

	svc.Pool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, errors.Errorf("failed to connect to database: %v", err)
	}

	return svc, nil
}

func WithLogger(l *zap.Logger) Option {
	return loggerOption{Logger: l}
}

func (s *Service) Stop() {
	s.Logger.Info("gracefully shutting down pgx service")
	s.Pool.Close()
}
