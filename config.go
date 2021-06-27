package main

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config holds configuration for syncing local files to the API
type Config struct {
	Username string
	Token    string
	Endpoint string
	ignore   []string
}

// parseIgnoreGlobs loads a list of glob patterns to ignore local files and directories
func parseIgnoreGlobs(dir string) (ignore []string) {
	ignore = append(ignore, ".git")
	path := filepath.Join(dir, ".moyenignore")
	file, err := os.Open(path)
	if err != nil {
		return // File doesn't exist, use default ignore
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ignore = append(ignore, scanner.Text()) // Add to slice of ignored globs
	}
	return
}

// ParseConfigYAML parses .moyenrc configuration file in supplied directory
// Configuration username and token must be specified
func ParseConfigYAML(dir string) (*Config, error) {
	config := &Config{
		ignore: parseIgnoreGlobs(dir),
	}
	configYAML, err := os.ReadFile(filepath.Join(dir, ".moyenrc"))
	if err != nil {
		return nil, errors.New("failed to read configuration JSON")
	}
	err = yaml.Unmarshal(configYAML, &config)
	// err = json.Unmarshal(configJSON, &config)
	if err != nil {
		return nil, errors.New("failed to parse configuration JSON")
	}
	return config, nil
}
