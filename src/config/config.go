package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

func Load() (*ServiceConfig, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var config ServiceConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
