package main

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/moyen-blog/sync-dir/asset"
	"github.com/moyen-blog/sync-dir/client"
)

const baseURL = "api.localhost:8080"

var ignore = [...]string{".git"}

// LocalAssetState returns all asset files in the current directory
func LocalAssetState(dir string) ([]asset.Asset, error) {
	assets := make([]asset.Asset, 0)
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
			m, err := asset.NewMarkdown(n)
			if err != nil {
				return err
			}
			assets = append(assets, m.Asset)
			for _, i := range m.Images {
				assets = append(assets, i.Asset)
			}
		}
		return nil
	}
	err := filepath.Walk(dir, walk)
	return assets, err
}

// RemoteAssetState returns the state of an authors articles
func RemoteAssetState(author string, token string) (result []asset.Asset, err error) {
	err = client.GetAssets(author, token, &result)
	return
}
