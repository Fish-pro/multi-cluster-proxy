package config

import (
	"fmt"
	"os"
)

type Config struct {
	LogLevel string

	Host string
	Port string

	TlsCert string
	TlsKey  string

	DataSource string

	BasicAuthUser     string
	BasicAuthPassword string
}

func getEnvOrDefault(key string, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}

func NewConfigFromEnv() *Config {
	return &Config{
		LogLevel: getEnvOrDefault("LOG_LEVEL", "info"),
		Host:     getEnvOrDefault("HOST", "0.0.0.0"),
		Port:     getEnvOrDefault("PORT", "9001"),
		TlsCert:  getEnvOrDefault("TLS_CERT", ""),
		TlsKey:   getEnvOrDefault("TLS_KEY", ""),
		DataSource: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			getEnvOrDefault("MYSQL_USERNAME", "root"),
			getEnvOrDefault("MYSQL_PASSWORD", "dangerous"),
			getEnvOrDefault("MYSQL_HOST", "127.0.0.1"),
			getEnvOrDefault("MYSQL_PORT", "3306"),
			getEnvOrDefault("DB_NAME", "edip-dev"),
		),
		BasicAuthUser:     getEnvOrDefault("BASIC_USER", "admin"),
		BasicAuthPassword: getEnvOrDefault("BASIC_PASSWORD", "admin"),
	}
}
