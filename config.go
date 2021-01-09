package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
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

// ParseConfigJSON parses .moyenrc configuration file in supplied directory
// Configuration username and token must be specified
func ParseConfigJSON(dir string) (*Config, error) {
	config := &Config{
		ignore: parseIgnoreGlobs(dir),
	}
	configJSON, err := ioutil.ReadFile(filepath.Join(dir, ".moyenrc"))
	err = json.Unmarshal(configJSON, &config)
	if err != nil {
		return nil, errors.New("Failed to parse configuration JSON")
	}
	return config, nil
}
