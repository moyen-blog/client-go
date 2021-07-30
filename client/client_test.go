package client

import (
	"fmt"
	"testing"

	"github.com/moyen-blog/client-go/client/mocks"
)

func init() {
	DefaultHTTPClient = &mocks.MockHTTPClient // Ensure no real HTTP calls are made
}

func TestClientCreate(t *testing.T) {
	username := "testusername"
	token := "testtoken"
	expectedEndpoint := "https://testusername.api.moyen.blog"
	c, err := NewClient(username, token, "", []string{})
	if err != nil {
		t.Error("Should create new client")
	}
	if c.username != username {
		t.Errorf("Should have client username %s but got %s", username, c.username)
	}
	if c.token != token {
		t.Errorf("Should have client token %s but got %s", token, c.token)
	}
	if c.endpoint != expectedEndpoint {
		t.Errorf("Should have client endpoint %s but got %s", defaultEndpoint, expectedEndpoint)
	}
}

func TestClientBadUsername(t *testing.T) {
	username := ""
	token := "testtoken"
	_, err := NewClient(username, token, "", []string{})
	if err == nil {
		t.Error("Should throw error for bad username")
	}
}

func TestClientBadToken(t *testing.T) {
	username := "testusername"
	token := ""
	_, err := NewClient(username, token, "", []string{})
	if err == nil {
		t.Error("Should throw error for bad token")
	}
}

func TestClientCustomEndpoint(t *testing.T) {
	username := "testusername"
	token := "testtoken"
	endpoint := "https://goodurl.com/goodpath"
	_, err := NewClient(username, token, endpoint, []string{})
	if err != nil {
		t.Error("Should create client with custom endpoint")
	}
}

func TestClientBadEndpoint(t *testing.T) {
	username := "testusername"
	token := "testtoken"
	endpoint := "badurl"
	_, err := NewClient(username, token, endpoint, []string{})
	if err == nil {
		t.Error("Should throw error for bad endpoint")
	}
}

func TestClientGetAssets(t *testing.T) {
	username := "testusername"
	token := "testtoken"
	c, err := NewClient(username, token, "", []string{})
	if err != nil {
		t.Error("Should create client")
	}
	path := "test.md"
	hash := "0123456789ABCDEF0123456789ABCDEF01234567"
	json := fmt.Sprintf(`[{ "path": "%s", "hash": "%s" }]`, path, hash)
	mocks.MockHTTPClient.SetResponse(json, 200, nil)
	a, err := c.GetAssets()
	if err != nil {
		t.Error("Should successfully fetch assets")
	}
	if len(a) != 1 {
		t.Errorf("Should return %d assets but got %d", 1, len(a))
	}
	if a[0].Path != path {
		t.Errorf("Should have path %s but got %s", path, a[0].Path)
	}
	if a[0].Hash != hash {
		t.Errorf("Should have hash %s but got %s", hash, a[0].Hash)
	}
}

func TestClientGetAssetsStatusError(t *testing.T) {
	username := "testusername"
	token := "testtoken"
	c, err := NewClient(username, token, "", []string{})
	if err != nil {
		t.Error("Should create client")
	}
	mocks.MockHTTPClient.SetResponse("[]", 500, nil)
	_, err = c.GetAssets()
	if err == nil {
		t.Error("Should fail to fetch assets with server error")
	}
}

func TestClientPutAsset(t *testing.T) {
	username := "testusername"
	token := "testtoken"
	c, err := NewClient(username, token, "", []string{})
	if err != nil {
		t.Error("Should create client")
	}
	path := "test.md"
	content := []byte("test")
	mocks.MockHTTPClient.SetResponse("", 200, nil)
	err = c.PutAsset(path, content)
	if err != nil {
		t.Error("Should successfully put asset")
	}
}

func TestClientDeleteAsset(t *testing.T) {
	username := "testusername"
	token := "testtoken"
	c, err := NewClient(username, token, "", []string{})
	if err != nil {
		t.Error("Should create client")
	}
	path := "test.md"
	mocks.MockHTTPClient.SetResponse("", 200, nil)
	err = c.DeleteAsset(path)
	if err != nil {
		t.Error("Should successfully put asset")
	}
}
