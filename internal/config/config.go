package config

import "os"

type DbConfig struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
}

type Config struct {
	DbConfig
}

func GetConfig() *Config {
	return &Config{
		DbConfig{
			DbHost: getEnv("DB_HOST", "localhost"),
			DbPort: getEnv("DB_PORT", "5432"),
			DbUser: getEnv("DB_USER", "postgres"),
			DbPass: getEnv("DB_PASSWORD", "postgres"),
			DbName: getEnv("DB_NAME", "ocp_tip_api"),
		},
	}
}

func getEnv(key string, defultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defultValue
}
