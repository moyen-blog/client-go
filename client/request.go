package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// request is used internally to make calls to the API
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
		err = fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}
	json.NewDecoder(resp.Body).Decode(&holder)
	return resp.StatusCode, err
}
