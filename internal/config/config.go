package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port         uint16   `yaml:"port"`
	DB           DBConfig `yaml:"db"`
	LoggingLevel string   `yaml:"logging_level"`
}

type DBConfig struct {
	Schema   string `yaml:"schema"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	IP       string `yaml:"ip"`
	Port     uint16 `yaml:"port"`
	DataBase string `yaml:"database"`
	SSL      string `yaml:"ssl"`
}

func (dbc DBConfig) GetDBURL() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s?sslmode=%s",
		dbc.Schema,
		dbc.User,
		dbc.Password,
		dbc.IP,
		dbc.Port,
		dbc.DataBase,
		dbc.SSL,
	)
}

func (dbc DBConfig) isValid() bool {
	if dbc.Schema == "" {
		return false
	}

	if dbc.User == "" {
		return false
	}

	if dbc.Password == "" {
		return false
	}

	if dbc.IP == "" {
		return false
	}

	if dbc.Port == 0 {
		return false
	}

	if dbc.DataBase == "" {
		return false
	}

	if dbc.SSL == "" {
		return false
	}

	return true
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

	if !config.DB.isValid() {
		return Config{}, errors.New("Could not read DB config")
	}

	return config, nil
}
