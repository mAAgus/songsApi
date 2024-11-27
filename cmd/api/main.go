package main

import (
	"flag"
	"go/scr/hhruxongs/internal/app/api"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/goloop/env"
)

var (
	configPath    string
	formatConfigs string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
	flag.StringVar(&formatConfigs, "format", ".toml", "format to configs select file")
}

func main() {
	flag.Parse()
	log.Println("Start works!!!")

	config := api.NewConfig()

	if formatConfigs == ".toml" {
		log.Println("Using TOML configuration format")
		_, err := toml.DecodeFile(configPath, config)
		if err != nil {
			log.Printf("Can not find TOML configs file. Using default values: %v", err)
		}
	} else if formatConfigs == ".env" {
		log.Println("Using ENV configuration format")
		err := env.Unmarshal(configPath, config)
		if err != nil {
			log.Printf("Can not find ENV configs file. Using default values: %v", err)
		}
	} else {
		log.Println("Undefined format. Using TOML configuration")
		_, err := toml.DecodeFile(configPath, config)
		if err != nil {
			log.Printf("Can not find configs file. Using default values: %v", err)
		}
	}

	log.Printf("Configuration loaded: %+v", config)

	server := api.New(config)

	log.Println("Starting API server...")
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
