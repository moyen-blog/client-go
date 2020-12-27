package asset

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
)

// Asset is the underlying file in markdown and image assets
type Asset struct {
	Path string
	Hash string
}

func (f *Asset) computeHash() (string, error) {
	file, err := os.Open(f.Path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	h := sha1.New()
	_, err = io.Copy(h, file)
	return hex.EncodeToString(h.Sum(nil)), err
}

// Buffer returns a bytes.Buffer of the asset files content
func (f *Asset) Buffer() (*bytes.Buffer, error) {
	b, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(b), nil
}

// NewAsset declares an asset describing markdown and images files
// The hash of the file contents is computed
func NewAsset(path string) (f *Asset, err error) {
	a := &Asset{
		Path: path,
	}
	a.Hash, err = a.computeHash()
	return a, err
}
