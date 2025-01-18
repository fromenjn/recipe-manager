package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Config holds all application configuration parameters.
type Config struct {
	ServerPort  string `json:"server_port"`
	RecipesPath string `json:"recipes_path"`
	// Add other fields as needed, e.g. database creds, logging level, etc.
}

// LoadConfig reads a JSON config file from the given path
// and unmarshals it into a Config struct.
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file %s: %w", path, err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(bytes, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	// Optionally, set defaults if fields are empty:
	if cfg.ServerPort == "" {
		cfg.ServerPort = ":8080"
	}
	if cfg.RecipesPath == "" {
		cfg.RecipesPath = "data/recipes"
	}

	return &cfg, nil
}
