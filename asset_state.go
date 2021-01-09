package main

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/moyen-blog/sync-dir/asset"
	"github.com/moyen-blog/sync-dir/client"
	"github.com/ryanuber/go-glob"
)

func uniqueAssets(assets []asset.Asset) (unique []asset.Asset) {
	keys := make(map[string]bool) // Used to check for previously seen asset paths
	for _, a := range assets {
		if _, hit := keys[a.Path]; !hit {
			keys[a.Path] = true
			unique = append(unique, a)
		}
	}
	return
}

// LocalAssetState returns all asset files in the current directory
func LocalAssetState(dir string, ignore []string) ([]asset.Asset, error) {
	assets := make([]asset.Asset, 0)
	r := regexp.MustCompile(`.\.md$`)
	walk := func(n string, f os.FileInfo, err error) error {
		for _, i := range ignore { // Skip ignored files and directories
			if glob.Glob(i, f.Name()) { // Glob file or directory name
				if f.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
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
	return uniqueAssets(assets), err
}

// RemoteAssetState returns the state of an authors articles
func RemoteAssetState(c *client.Client) (result []asset.Asset, err error) {
	err = c.GetAssets(&result)
	return
}
