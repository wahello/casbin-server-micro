package main

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	MetricsPort   string `envconfig:"METRICS_PORT" required:"false" default:"8086"`
	Environment   string `envconfig:"ENVIRONMENT" default:"dev"`
	BrokerAddress string `envconfig:"BROKER_ADDRESS" default:"amqp://127.0.0.1:5672"`
	MicroRegistry string `envconfig:"MICRO_REGISTRY" required:"false"`
	CasbinDSN     string `envconfig:"CASBIN_DSN" required:"true"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
