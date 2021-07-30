package client

import (
	"errors"
	"testing"

	"github.com/moyen-blog/client-go/client/mocks"
)

func init() {
	DefaultHTTPClient = &mocks.MockHTTPClient // Ensure no real HTTP calls are made
}

func TestSyncMixed(t *testing.T) {
	username := "testuser"
	token := "testtoken"
	c, err := NewClient(username, token, "", []string{})
	if err != nil {
		t.Error("Should create new client with valid credentials")
	}
	assetsLocal := []Asset{
		{Path: "testLocal.md"},
		{Path: "testIntersection.md", Hash: "A"},
	}
	assetsRemote := []Asset{
		{Path: "testRemote.md"},
		{Path: "testIntersection.md", Hash: "B"},
	}
	diff := c.DiffAssets(assetsLocal, assetsRemote)
	count := 0
	c.Sync(diff, func(a Asset, e error) error {
		count += 1
		return nil
	})
	if count != 3 {
		t.Errorf("Should receive callback %d times but got %d", 3, count)
	}
}

func TestSyncHalt(t *testing.T) {
	username := "testuser"
	token := "testtoken"
	c, err := NewClient(username, token, "", []string{})
	if err != nil {
		t.Error("Should create new client with valid credentials")
	}
	assetsLocal := []Asset{
		{Path: "testLocal.md"},
		{Path: "testIntersection.md", Hash: "A"},
	}
	assetsRemote := []Asset{
		{Path: "testRemote.md"},
		{Path: "testIntersection.md", Hash: "B"},
	}
	diff := c.DiffAssets(assetsLocal, assetsRemote)
	count := 0
	c.Sync(diff, func(a Asset, e error) error {
		count += 1
		return errors.New("Deliberate error to halt synchronization")
	})
	if count != 1 {
		t.Errorf("Should receive callback %d times but got %d", 1, count)
	}
}
