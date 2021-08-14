package config

import "github.com/aifaniyi/env"

type Config struct {
	Port string
}

func LoadConfig() *Config {
	return &Config{
		Port: env.LoadString("PORT", ":8000"),
	}
}
