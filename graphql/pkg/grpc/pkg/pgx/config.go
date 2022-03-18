package pgx

import (
	"fmt"
	"log"

	"github.com/cod3rcarl/wwdatabase-go-backend/util"
)

type Config struct {
	DatabaseDSN string
}

func ReadConfig() (c Config, err error) {
	cfg, err := util.LoadConfig("./../../../../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	c.DatabaseDSN = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB)

	return c, nil
}
