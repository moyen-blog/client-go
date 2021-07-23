package client

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"regexp"
)

const defaultEndpoint = "https://api.moyen.blog"

// Client defines an instance of an API client
type Client struct {
	username string
	token    string
	endpoint string
	ignore   []string
}

// NewClient creates an API client
func NewClient(username, token, endpoint string, ignore []string) (*Client, error) {
	if endpoint == "" {
		endpoint = defaultEndpoint
	}
	r := regexp.MustCompile(`^\w+$`) // Just ensure username and token are present
	if !r.MatchString(username) || !r.MatchString(token) {
		return nil, errors.New("invalid username and/or token")
	}
	e, err := url.Parse(endpoint) // Ensure valid API endpoint
	if err != nil {
		return nil, fmt.Errorf("failed to parse API endpoint %s", e)
	}
	e.Path, e.RawQuery, e.RawQuery, e.User = "", "", "", nil // Ignore all but scheme, host
	e.Host = username + "." + e.Host
	return &Client{
		username,
		token,
		e.String(),
		ignore,
	}, nil
}

// GetAssets gets asset paths and hashes for a provided author
// JSON response is decoded into the provided holder
func (c *Client) GetAssets(holder interface{}) error {
	_, err := request("GET", c.endpoint, c.token, nil, holder)
	return err
}

// PutAsset upserts an asset for a provided author
// Used for both creating and updating articles and images
func (c *Client) PutAsset(path string, payload []byte) error {
	buf := bytes.NewBuffer(payload)
	_, err := request("PUT", c.endpoint+"/"+path, c.token, buf, nil)
	return err
}

// DeleteAsset deletes an asset for a provided author
func (c *Client) DeleteAsset(path string) error {
	_, err := request("DELETE", c.endpoint+"/"+path, c.token, nil, nil)
	return err
}
