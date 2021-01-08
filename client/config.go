package client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

const defaultURL = "https://api.moyen.blog"

type config struct {
	Username string
	Token    string
	Endpoint string
}

// Parse .moyenrc configuration file in CWD
// Configuration username and token must be specified
func parseConfigJSON() (*config, error) {
	c := config{
		Endpoint: defaultURL,
	}
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	configJSON, err := ioutil.ReadFile(filepath.Join(path, ".moyenrc"))
	err = json.Unmarshal(configJSON, &c)
	if err != nil {
		return nil, errors.New("Failed to parse configuration JSON")
	}
	r := regexp.MustCompile(`^\w+$`)
	if !r.MatchString(c.Username) || !r.MatchString(c.Token) {
		return nil, errors.New("Invalid username and/or token in configuration JSON")
	}
	return &c, nil
}
