package asset

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
)

// Markdown defines a simple view of a markdown file
type Markdown struct {
	Asset  Asset
	Images []*Image
}

// scanImages scans the markdown file for local images
func (f *Markdown) scanImages() ([]*Image, error) {
	file, err := os.Open(f.Asset.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	imagePattern := regexp.MustCompile(`!\[[^\[]+\]\((.*?)(?:\s*".*")?\)`)
	schemePattern := regexp.MustCompile(`^[a-zA-Z]{2,20}:\/\/`)
	images := make([]*Image, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := imagePattern.FindAllStringSubmatch(line, -1)
		for _, i := range matches {
			path := i[1]
			if schemePattern.MatchString(path) {
				continue // Remote asset URL
			}
			imagePath := filepath.Join(filepath.Dir(f.Asset.Path), path) // Relative to markdown file
			image, err := NewImage(imagePath)
			if err != nil {
				return images, err
			}
			images = append(images, image)
		}
	}
	return images, nil
}

// NewMarkdown declares a markdown file from a file path
// The hash of the file contents is computed
func NewMarkdown(path string) (f *Markdown, err error) {
	a := Asset{
		Path: path,
	}
	a.Hash, err = a.computeHash()
	if err != nil {
		return nil, err
	}
	f = &Markdown{
		Asset: a,
	}
	f.Images, err = f.scanImages()
	return f, err
}
