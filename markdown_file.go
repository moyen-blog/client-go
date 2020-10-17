package main

import (
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
)

// MarkdownFile defines a simple view of a markdown file
type MarkdownFile struct {
	Path string
	Hash string
}

func (f *MarkdownFile) getContent() ([]byte, error) {
	return ioutil.ReadFile(f.Path)
}

func (f *MarkdownFile) computeHash() (string, error) {
	content, err := f.getContent()
	if err != nil {
		return "", err
	}
	h := sha1.New()
	h.Write(content)
	return hex.EncodeToString(h.Sum(nil)), nil
}

// NewMarkdownFile declares a markdown file from a file path
// The hash of the file contents is computed
// To save memory, the file contents themselves are not stored in the struct
func NewMarkdownFile(path string) (m *MarkdownFile, err error) {
	m = &MarkdownFile{
		Path: path,
	}
	m.Hash, err = m.computeHash()
	return m, err
}
