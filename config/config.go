package config

import (
	"encoding/json"
	"os"
	"pitempsensor/model"
)

type Config struct {
	Secrets model.Secrets `json:"secrets"`
	ApiURL  string        `json:"api_url"`
}

func LoadConfig(configFile string) (Config, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
