package config

import (
	"log"
	"sync"
)

type std struct {
	config *Config
	once   sync.Once
}

var stdImpl = std{}

func stdConfig() *Config {
	stdImpl.once.Do(func() {
		config, err := ParseFile("config.yaml")
		if err != nil {
			log.Fatalf("creating config error: %s", err)
		}
		stdImpl.config = config
	})
	return stdImpl.config
}

func Global() *Config {
	return stdConfig()
}
