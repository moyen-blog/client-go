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

// cleanURL validates a URL and removes all but scheme and host
func cleanURL(u string) (*url.URL, error) {
	url, err := url.Parse(u)
	if err != nil {
		return nil, fmt.Errorf("failed to parse API endpoint %s", u)
	}
	if url.Host == "" || url.Scheme == "" {
		return nil, errors.New("URL host and scheme can not be empty")
	}
	url.Path, url.RawQuery, url.RawFragment, url.User = "", "", "", nil // Ignore all but scheme, host
	return url, nil
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
	url, err := cleanURL(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse API endpoint %s", endpoint)
	}
	url.Host = username + "." + url.Host
	return &Client{
		username,
		token,
		url.String(),
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
