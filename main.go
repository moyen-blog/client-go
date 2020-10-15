package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// MarkdownFile defines a simple view of a markdown file
type MarkdownFile struct {
	Path    string
	Content string
	Hash    string
}

func (f *MarkdownFile) computeHash() {}

// WalkFilterMarkdown returns all markdown files in the current directory
func WalkFilterMarkdown(dir string) {
	r := regexp.MustCompile(`.\.md$`)
	walk := func(n string, f os.FileInfo, err error) error {
		if f.IsDir() {
			if f.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		if r.MatchString(n) {
			fmt.Printf("Markdown file: %s\n", n)
		} else {
			fmt.Printf("Other file: %s\n", n)
		}
		return nil
	}
	filepath.Walk(dir, walk)
}

func main() {
	WalkFilterMarkdown(".")
}
