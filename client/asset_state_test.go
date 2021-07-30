package client

import (
	"path"
	"testing"
	"testing/fstest"

	"github.com/moyen-blog/sync-dir/client/mocks"
)

func init() {
	DefaultHTTPClient = &mocks.MockHTTPClient // Ensure no real HTTP calls are made
}

func TestAssetStateRemote(t *testing.T) {
	username := "testusername"
	token := "testtoken"
	c, err := NewClient(username, token, "", []string{})
	if err != nil {
		t.Error("Should create client")
	}
	mocks.MockHTTPClient.SetResponse("[]", 200, nil)
	a, err := c.AssetStateRemote()
	if err != nil {
		t.Error("Should successfully get remote asset state")
	}
	if len(a) != 0 {
		t.Errorf("Should return %d assets but got %d", 0, len(a))
	}
}

func TestAssetStateLocal(t *testing.T) {
	path := "test.md"
	hash := "da39a3ee5e6b4b0d3255bfef95601890afd80709"
	fsys := make(fstest.MapFS)
	fsys[path] = &fstest.MapFile{
		Data: []byte(""),
	}
	username := "testusername"
	token := "testtoken"
	c, err := NewClient(username, token, "", []string{})
	if err != nil {
		t.Error("Should create client")
	}
	a, err := c.AssetStateLocal(fsys)
	if err != nil {
		t.Error("Should successfully get remote asset state")
	}
	if len(a) != 1 {
		t.Errorf("Should return %d assets but got %d", 1, len(a))
	}
	if a[0].Path != path {
		t.Errorf("Should have asset path %s but got %s", path, a[0].Path)
	}
	if a[0].Hash != hash {
		t.Errorf("Should have asset hash %s but got %s", hash, a[0].Hash)
	}
}

func TestAssetStateLocalIgnore(t *testing.T) {
	path := "test.md"
	fsys := make(fstest.MapFS)
	fsys[path] = &fstest.MapFile{
		Data: []byte(""),
	}
	username := "testusername"
	token := "testtoken"
	c, err := NewClient(username, token, "", []string{path})
	if err != nil {
		t.Error("Should create client")
	}
	a, err := c.AssetStateLocal(fsys)
	if err != nil {
		t.Error("Should successfully get remote asset state")
	}
	if len(a) != 0 {
		t.Errorf("Should return %d assets but got %d", 0, len(a))
	}
}

func TestAssetStateLocalIgnoreGlob(t *testing.T) {
	path := "test_glob.md"
	glob := "test_*.md"
	fsys := make(fstest.MapFS)
	fsys[path] = &fstest.MapFile{
		Data: []byte(""),
	}
	username := "testusername"
	token := "testtoken"
	c, err := NewClient(username, token, "", []string{glob})
	if err != nil {
		t.Error("Should create client")
	}
	a, err := c.AssetStateLocal(fsys)
	if err != nil {
		t.Error("Should successfully get remote asset state")
	}
	if len(a) != 0 {
		t.Errorf("Should return %d assets but got %d", 0, len(a))
	}
}

func TestAssetStateLocalIgnoreSubdir(t *testing.T) {
	subdir := "subdir"
	path := path.Join(subdir, "test.md")
	fsys := make(fstest.MapFS)
	fsys[path] = &fstest.MapFile{
		Data: []byte(""),
	}
	username := "testusername"
	token := "testtoken"
	c, err := NewClient(username, token, "", []string{subdir})
	if err != nil {
		t.Error("Should create client")
	}
	a, err := c.AssetStateLocal(fsys)
	if err != nil {
		t.Error("Should successfully get remote asset state")
	}
	if len(a) != 0 {
		t.Errorf("Should return %d assets but got %d", 0, len(a))
	}
}
