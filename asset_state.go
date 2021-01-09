package main

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/moyen-blog/sync-dir/asset"
	"github.com/moyen-blog/sync-dir/client"
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
	walk := func(path string, f os.FileInfo, err error) error {
		relative, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		for _, i := range ignore { // Skip ignored files and directories
			if match, _ := filepath.Match(i, relative); match { // Glob file or directory name
				if f.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}
		if r.MatchString(path) {
			markdown, err := asset.NewMarkdown(relative)
			if err != nil {
				return err
			}
			assets = append(assets, markdown.Asset)
			for _, image := range markdown.Images {
				assets = append(assets, image.Asset)
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
