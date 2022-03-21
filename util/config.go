package util

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	PostgresUser     string `mapstructure:"postgres-user"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB       string `mapstructure:"POSTGRES_DB"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	GRPCServerHost   string `mapstructure:"GRPC_HOST"`
	GRPCServerPort   string `mapstructure:"GRPC_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	var environment string
	environment = "development"
	if err = godotenv.Load(); err != nil {
		environment = "production"
	}

	if environment == "development" {
		viper.AddConfigPath(path)
		viper.SetConfigName("app")
		viper.SetConfigType("env")
	}

	if environment == "production" {
		viper.AddConfigPath(path)
		viper.SetConfigName("prod")
		viper.SetConfigType("env")
	}

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return config, nil
}
