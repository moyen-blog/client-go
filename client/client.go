package client

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
)

// Client defines an instance of an API client
type Client struct {
	username string
	token    string
	endpoint string
}

// NewClient creates an API client
// Configuration is loaded from .moyenrc (JSON) in CWD
func NewClient() (*Client, error) {
	config, err := parseConfigJSON()
	if err != nil {
		return nil, err
	}
	e, err := url.Parse(config.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse API endpoint %s", e)
	}
	e.Path, e.RawQuery, e.RawQuery, e.User = "", "", "", nil // Ignore all but scheme, host
	e.Host = config.Username + "." + e.Host
	return &Client{
		config.Username,
		config.Token,
		e.String(),
	}, nil
}

func (c *Client) assetEndpoint(p string) string {
	return c.endpoint + "/" + p
}

// GetAssets gets asset paths and hashes for a provided author
// JSON response is decoded into the provided holder
func (c *Client) GetAssets(holder interface{}) error {
	_, err := request("GET", c.endpoint, c.token, nil, holder)
	return err
}

// PutAsset upserts an asset for a provided author
// Used for both creating and updating articles and images
func (c *Client) PutAsset(path string, payload *bytes.Buffer) error {
	if payload == nil {
		return errors.New("Payload can not be nil")
	}
	_, err := request("PUT", c.assetEndpoint(path), c.token, payload, nil)
	return err
}

// DeleteAsset deletes an asset for a provided author
func (c *Client) DeleteAsset(path string) error {
	_, err := request("DELETE", c.assetEndpoint(path), c.token, nil, nil)
	return err
}
