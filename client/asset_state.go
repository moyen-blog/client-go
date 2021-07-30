package client

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

// AssetStateLocal returns all asset files in the current directory of fsys
// If nil is passed as fsys, os.DirFS(".") will be used
func (c *Client) AssetStateLocal(fsys fs.FS) ([]Asset, error) {
	assets := make([]Asset, 0)
	pattern := regexp.MustCompile(`.\.(?:md|png|jpg|jpeg)$`)
	walk := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err // Error reading file/dir at path
		}
		for _, patternIgnore := range c.ignore { // Skip ignored files and directories
			if match, _ := filepath.Match(patternIgnore, path); match { // Glob file or directory name
				if d.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}
		if pattern.MatchString(path) {
			f, err := fsys.Open(path)
			if err != nil {
				return err
			}
			asset, err := NewAsset(path, f)
			if err != nil {
				return err
			}
			assets = append(assets, *asset)
		}
		return nil
	}
	if fsys == nil { // Default filesystem
		fsys = os.DirFS(".")
	}
	err := fs.WalkDir(fsys, ".", walk)
	return assets, err
}

// AssetStateRemote returns the state of an authors articles
// Essentially a wrapper around GetAssets
func (c *Client) AssetStateRemote() ([]Asset, error) {
	return c.GetAssets()
}
