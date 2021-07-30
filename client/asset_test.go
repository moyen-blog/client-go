package client

import (
	"strings"
	"testing"

	"github.com/moyen-blog/client-go/client/mocks"
)

func TestNewAsset(t *testing.T) {
	path := "test.md"
	content := strings.NewReader("")
	hash := "da39a3ee5e6b4b0d3255bfef95601890afd80709"
	a, err := NewAsset(path, content)
	if err != nil {
		t.Error("Should successfully create asset")
	}
	if a.Hash != hash {
		t.Errorf("Should have hash %s but got %s", hash, a.Hash)
	}
}

func TestNewAssetErrorReader(t *testing.T) {
	path := "test.md"
	content := mocks.ErrorReader{}
	_, err := NewAsset(path, content)
	if err == nil {
		t.Error("Should throw error when reading faulty reader")
	}
}
