package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var DefaultHTTPClient HTTPClient = http.DefaultClient

// request is used internally to make calls to the API
func request(method string, url string, token string, payload io.Reader, dest interface{}) (int, error) {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return 0, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := DefaultHTTPClient.Do(req)
	if err != nil {
		return 0, err
	}
	if resp == nil {
		return 0, errors.New("request returned nil response")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}
	if err == nil && dest != nil {
		err = json.NewDecoder(resp.Body).Decode(&dest)
	}
	return resp.StatusCode, err
}
