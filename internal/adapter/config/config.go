package config

import (
	"os"
)

type Config struct {
  BrokerUrl   string
  DatabaseUrl string
}

func New() (*Config, error) {
  config := &Config{
    BrokerUrl:   os.Getenv("BROKER_URL"),
    DatabaseUrl: os.Getenv("DATABASE_URL"),
  }

	return config, nil
}
