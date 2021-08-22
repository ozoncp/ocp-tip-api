package config

import (
	"github.com/kelseyhightower/envconfig"
)

type DbConfig struct {
	DbHost string `envconfig:"POSTGRES_HOST" default:"localhost"`
	DbPort string `envconfig:"POSTGRES_PORT" default:"5432"`
	DbUser string `envconfig:"POSTGRES_USER" default:"postgres"`
	DbPass string `envconfig:"POSTGRES_PASSWORD" default:"postgres"`
	DbName string `envconfig:"POSTGRES_DB" default:"ocp_tip_api"`
}

type Config struct {
	Brokers             []string
	JaegerAgentHostPort string `envconfig:"JAEGER_AGENT_HOST_PORT" default:"jaeger:6831"`
	DbConfig
}

func GetConfig() (*Config, error) {
	var c Config
	err := envconfig.Process("tip", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
