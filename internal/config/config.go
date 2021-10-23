package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type HackConfig struct {
	ID   string   `yaml:"id"`
	Blur []string `yaml:"blur"`
	Orig []string `yaml:"orig"`
}

type Config struct {
	TelegramBot struct {
		Token string `yaml:"token"`
	} `yaml:"telegram_bot"`

	Instagram struct {
		SessionID string `yaml:"sessionid"`
	} `yaml:"instagram"`

	VK struct {
		Token string `yaml:"token"`
	} `yaml:"vk"`

	YooMoney struct {
		AccessToken string `yaml:"access_token"`
		ClientID    string `yaml:"client_id"`
		RedirectURI string `yaml:"redirect_uri"`
		Scope       string `yaml:"scope"`
	} `yaml:"yoomoney"`

	Store struct {
		DBName string `yaml:"db_name"`
	}

	Hacks []HackConfig `yaml:"hacks"`
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
