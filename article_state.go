package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

const baseURL = "localhost:8080"

var ignore = [...]string{".git"}

// LocalArticleState returns all markdown files in the current directory
// SHA1 hash is computed from file contents
func LocalArticleState(dir string) ([]MarkdownFile, error) {
	files := make([]MarkdownFile, 0)
	r := regexp.MustCompile(`.\.md$`)
	walk := func(n string, f os.FileInfo, err error) error {
		if f.IsDir() {
			for _, i := range ignore {
				if f.Name() == i {
					return filepath.SkipDir
				}
			}
			return nil
		}
		if r.MatchString(n) {
			m, err := NewMarkdownFile(n)
			if err != nil {
				return err
			}
			files = append(files, *m)
		}
		return nil
	}
	err := filepath.Walk(dir, walk)
	return files, err
}

// RemoteArticleState returns the state of an authors articles
func RemoteArticleState(author string, token string) (result []MarkdownFile, err error) {
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
