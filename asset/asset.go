package asset

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

// Asset is the underlying file in markdown and image assets
type Asset struct {
	Path string
	Hash string
}

// Read allows Asset to conform to the io.Reader interface
func (f Asset) Read(p []byte) (int, error) {
	file, err := os.Open(f.Path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.Read(p)
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
