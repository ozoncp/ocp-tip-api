package config

import (
	"github.com/kelseyhightower/envconfig"
)

type DbConfig struct {
	DbHost string `default:"localhost"`
	DbPort string `default:"5432"`
	DbUser string `default:"postgres"`
	DbPass string `default:"postgres"`
	DbName string `default:"ocp_tip_api"`
}

type Config struct {
	DbConfig
}

func GetConfig() (*Config, error) {
	var d DbConfig
	err := envconfig.Process("tip", &d)
	if err != nil {
		return nil, err
	}
	return &Config{d}, nil
}
