package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "api.localhost:8080"

func authorURL(a string) string {
	return fmt.Sprintf("http://%s.%s", a, baseURL)
}

func authorFileURL(a string, p string) string {
	return fmt.Sprintf("%s/%s", authorURL(a), p)
}

// GetArticles gets article paths and hashes for a provided author
// JSON response is decoded into the provided holder
func GetArticles(author string, token string, holder interface{}) error {
	_, err := request("GET", authorURL(author), token, nil, holder)
	return err
}

// PutArticle upserts an article for a provided author
// Used for both creating and updating articles
func PutArticle(author string, token string, path string, markdown []byte) error {
	_, err := request("PUT", authorFileURL(author, path), token, bytes.NewBuffer(markdown), nil)
	return err
}

// DeleteArticle deletes an article for a provided author
func DeleteArticle(author string, token string, path string) error {
	_, err := request("DELETE", authorFileURL(author, path), token, nil, nil)
	return err
}

func request(method string, url string, token string, payload io.Reader, holder interface{}) (int, error) {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return 0, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("Request failed with status code %d", resp.StatusCode)
	}
	json.NewDecoder(resp.Body).Decode(&holder)
	return resp.StatusCode, nil
}
