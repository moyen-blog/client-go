package main

import (
	"errors"
	"io/fs"
	"os"

	"gopkg.in/yaml.v2"
)

// Config holds configuration for syncing local files to the API
type Config struct {
	Ignore      []string
	Endpoint    string
	credentials struct {
		Username string
		Token    string
	}
}

// parseYAML loads file contents and attempts to unmarshal into destination
func parseYaml(fsys fs.FS, file string, dest interface{}) error {
	data, err := fs.ReadFile(fsys, file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil // File not found; skip it
		}
		return err
	}
	if err := yaml.Unmarshal(data, dest); err != nil {
		return err
	}
	return nil
}

// parseConfig parses expected credentials file and optional config file
func parseConfig(fsys fs.FS) (*Config, error) {
	config := &Config{}

	if err := parseYaml(fsys, ".moyenconfig", &config); err != nil {
		return nil, errors.New("failed to read configuration YAML")
	}
	config.Ignore = append(config.Ignore, ".git") // Always ignore git dir
	if err := parseYaml(fsys, ".moyencredentials", &config.credentials); err != nil {
		return nil, errors.New("failed to read credentials YAML")
	}
	if username := os.Getenv("MOYEN_USERNAME"); username != "" {
		config.credentials.Username = username
	}
	if token := os.Getenv("MOYEN_TOKEN"); token != "" {
		config.credentials.Token = token
	}
	if config.credentials.Username == "" || config.credentials.Token == "" {
		return config, errors.New("username and token are required")
	}
	return config, nil
}
