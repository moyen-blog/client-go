package client

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestNewHasher(t *testing.T) {
	hash := "da39a3ee5e6b4b0d3255bfef95601890afd80709"
	hasherReader := newHasherReader(strings.NewReader(""))
	bytes, err := ioutil.ReadAll(hasherReader)
	if err != nil {
		t.Error("Should read HasherReader")
	}
	if len(bytes) != 0 {
		t.Errorf("Should read %d bytes from reader but got %d", 0, len(bytes))
	}
	if hasherReader.Hash() != hash {
		t.Errorf("Should have hash %s but got %s", hash, hasherReader.Hash())
	}
}
