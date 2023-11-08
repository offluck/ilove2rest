package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port         uint16 `yaml:"port"`
	DBURL        string `yaml:"db_url"`
	LoggingLevel string `yaml:"logging_level"`
}

func ReadConfig(filePath string) (Config, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}

	if config.Port == 0 {
		return Config{}, errors.New("Could not read the port")
	}

	if config.DBURL == "" {
		return Config{}, errors.New("Could not read DB URL")
	}

	return config, nil
}
