package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const (
	EnvBaseURL = "YOUTRACK_BASE_URL"
	EnvToken   = "YOUTRACK_TOKEN"
)

type Config struct {
	BaseURL string `json:"base_url"`
	Token   string `json:"token"`
}

func DefaultPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("get user config dir: %w", err)
	}

	return filepath.Join(dir, "try_parse_youtrack", "config.json"), nil
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return Config{}, nil
	}
	if err != nil {
		return Config{}, fmt.Errorf("read config: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("parse config: %w", err)
	}

	return cfg, nil
}

func ApplyEnv(cfg Config) Config {
	if baseURL := os.Getenv(EnvBaseURL); baseURL != "" {
		cfg.BaseURL = baseURL
	}
	if token := os.Getenv(EnvToken); token != "" {
		cfg.Token = token
	}

	return cfg
}
