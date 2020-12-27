package asset

// Image defines a simple view of an image file
type Image struct {
	Asset Asset
}

// NewImage declares an image file from a file path
// The hash of the file contents is computed
func NewImage(path string) (f *Image, err error) {
	a, err := NewAsset(path)
	if err != nil {
		return nil, err
	}
	f = &Image{
		Asset: *a,
	}
	return f, nil
}
