package server

import (
	"log"

	"github.com/cod3rcarl/wwdatabase-go-backend/util"
)

func ReadConfig() (util.Config, error) {
	cfg, err := util.LoadConfig("./../../../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	return cfg, nil
}
