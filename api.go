package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "localhost:8080"

// GetArticleState returns the state of an authors articles
func GetArticleState(author string, token string) (result []MarkdownFile, err error) {
	url := fmt.Sprintf("http://%s.%s/", author, baseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	// req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&result)
	return
}
