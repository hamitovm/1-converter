package config

import "os"

type Config struct {
	Key string
}

func NewConfig() *Config {
	keyValue := os.Getenv("KEY")
	return &Config{Key: keyValue}
}
