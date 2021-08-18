package config

import (
	"os"
)

type Config struct {
	LogLevel string

	Host string
	Port string

	TlsCert string
	TlsKey  string
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
	}
}
