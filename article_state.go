package main

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/moyen-blog/sync-dir/client"
)

const baseURL = "api.localhost:8080"

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
	err = client.GetArticles(author, token, &result)
	return
}
