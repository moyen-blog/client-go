package client

import (
	"crypto/sha1"
	"encoding/hex"
	"hash"
	"io"
)

// hasherReader calculates the hash of a byte stream
// As an underlying io.Reader is read from, the hash is updated
type hasherReader struct {
	hash   hash.Hash
	reader io.Reader
}

// newHasherReader creates a new hasherReader from a provided io.Reader
func newHasherReader(r io.Reader) hasherReader {
	hash := sha1.New()
	reader := io.TeeReader(r, hash)
	return hasherReader{hash, reader}
}

// Hash returns the hash value
// Ensure all contents of the underlying io.Reader have been read
func (h hasherReader) Hash() string {
	return hex.EncodeToString(h.hash.Sum(nil))
}

// Read allows hasherReader to conform to io.Reader interface
func (h hasherReader) Read(p []byte) (n int, err error) {
	return h.reader.Read(p)
}
