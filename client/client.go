package client

import (
	"bytes"
	"errors"
	"fmt"
)

const baseURL = "api.localhost:8080"

func authorURL(a string) string {
	return fmt.Sprintf("http://%s.%s", a, baseURL)
}

func authorFileURL(a string, p string) string {
	return fmt.Sprintf("%s/%s", authorURL(a), p)
}

// GetAssets gets asset paths and hashes for a provided author
// JSON response is decoded into the provided holder
func GetAssets(author string, token string, holder interface{}) error {
	_, err := request("GET", authorURL(author), token, nil, holder)
	return err
}

// PutAsset upserts an asset for a provided author
// Used for both creating and updating articles and images
func PutAsset(author string, token string, path string, payload *bytes.Buffer) error {
	if payload == nil {
		return errors.New("Payload can not be nil")
	}
	_, err := request("PUT", authorFileURL(author, path), token, payload, nil)
	return err
}

// DeleteAsset deletes an asset for a provided author
func DeleteAsset(author string, token string, path string) error {
	_, err := request("DELETE", authorFileURL(author, path), token, nil, nil)
	return err
}
