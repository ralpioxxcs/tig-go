package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type User struct {
	Email string
	Name  string
}

type Core struct {
	Bare bool
}

type Config struct {
	Core Core
	User User
}

const configFileName = "config"

func CreateConfigFile(gitDir string, cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(gitDir, configFileName), data, os.ModePerm)
}

func ReadConfigFile(gitDir string) (Config, error) {
	data, err := os.ReadFile(filepath.Join(gitDir, configFileName))
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, nil
	}

  return cfg, nil;
}
