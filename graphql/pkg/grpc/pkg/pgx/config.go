package pgx

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	DatabaseDSN      string
	PostgresUSER     string `envconfig:"POSTGRES_USER" required:"true"`
	PostgresPASSWORD string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	PostgresDB       string `envconfig:"POSTGRES_DB" required:"true"`
	PostgresPORT     string `envconfig:"POSTGRES_PORT" required:"true"`
	PostgresHOST     string `envconfig:"POSTGRES_HOST" required:"true"`
}

func ReadConfig() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, errors.Errorf("failed to parse config; error=%v", err)
	}

	cfg.DatabaseDSN = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		cfg.PostgresUSER, cfg.PostgresPASSWORD, cfg.PostgresHOST, cfg.PostgresPORT, cfg.PostgresDB)

	return cfg, nil
}
