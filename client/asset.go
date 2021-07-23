package client

import (
	"errors"
	"io"
	"io/ioutil"
)

// Asset is the underlying file in markdown and image assets
type Asset struct {
	Path    string
	Hash    string
	Content []byte
}

// NewAsset returns an asset created from a path and reader
func NewAsset(path string, reader io.Reader) (*Asset, error) {
	a := &Asset{
		Path: path,
	}
	hasherReader := newHasherReader(reader)
	bytes, err := ioutil.ReadAll(hasherReader)
	if err != nil {
		return nil, errors.New("failed to read asset content")
	}
	a.Content = bytes
	a.Hash = hasherReader.Hash() // Read entirety of io.Reader, hash is complete
	return a, nil
}
