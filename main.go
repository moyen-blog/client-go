package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
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

// WalkFilterMarkdown returns all markdown files in the current directory
func WalkFilterMarkdown(dir string) ([]MarkdownFile, error) {
	files := make([]MarkdownFile, 0)
	r := regexp.MustCompile(`.\.md$`)
	walk := func(n string, f os.FileInfo, err error) error {
		if f.IsDir() {
			if f.Name() == ".git" {
				return filepath.SkipDir
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

func main() {
	files, err := WalkFilterMarkdown(".")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Found %d markdown files\n", len(files))
	fmt.Println(files)
}
