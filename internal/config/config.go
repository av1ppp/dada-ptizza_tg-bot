package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	TelegramBot struct {
		Token string `yaml:"token"`
	} `yaml:"telegram_bot"`

	VK struct {
		Token string `yaml:"token"`
	} `yaml:"vk"`
}

// ParseFile - ...
func ParseFile(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
