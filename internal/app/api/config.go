package api

import (
	"go/scr/hhruxongs/storage"
	"log"
)

type Config struct {
	BindAddr    string `env:"bind_addr"`
	LoggerLevel string `env:"logger_level"`
	Storage     *storage.Config
}

func NewConfig() *Config {
	config := &Config{
		BindAddr:    ":8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}

	log.Printf("New API configuration created:\n")
	log.Printf("  BindAddr:    %s\n", config.BindAddr)
	log.Printf("  LoggerLevel: %s\n", config.LoggerLevel)
	log.Printf("  Storage:     %+v\n", config.Storage)

	return config
}
