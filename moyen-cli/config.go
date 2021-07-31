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
func parseYaml(fsys fs.FS, file string, required bool, dest interface{}) error {
	data, err := fs.ReadFile(fsys, file)
	if err != nil {
		if !required && errors.Is(err, os.ErrNotExist) {
			return nil // File not required; skip it
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

	if err := parseYaml(fsys, ".moyenconfig", false, &config); err != nil {
		return nil, errors.New("failed to read configuration YAML")
	}
	config.Ignore = append(config.Ignore, ".git") // Always ignore git dir
	if err := parseYaml(fsys, ".moyencredentials", true, &config.credentials); err != nil {
		return nil, errors.New("failed to read credentials YAML")
	}
	return config, nil
}
