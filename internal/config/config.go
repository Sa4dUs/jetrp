package config

import (
	"os"
	"strings"
)

type Config struct {
	Port     string
	Backends []string
}

func LoadConfig() Config {
	port := getEnv("PROXY_PORT", "8080")
	backends := strings.Split(getEnv("BACKENDS", "http://backend1:8081, http://backend2:8081"), ",")

	return Config{
		Port:     port,
		Backends: backends,
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
