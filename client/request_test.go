package client

import (
	"errors"
	"net/http"
	"testing"

	"github.com/moyen-blog/sync-dir/client/mocks"
)

func init() {
	DefaultHTTPClient = &mocks.MockHTTPClient // Ensure no real HTTP calls are made
}

func TestRequestSuccess(t *testing.T) {
	token := "testtoken"
	mocks.MockHTTPClient.SetResponse("", 200, nil)
	status, err := request(http.MethodGet, "", token, nil, nil)
	if err != nil {
		t.Error("Should successfully make request")
	}
	if status != 200 {
		t.Errorf("Should return status %d but got %d", 200, status)
	}
	authHeader := mocks.MockHTTPClient.LastRequest.Header.Get("Authorization")
	if authHeader != "Bearer "+token {
		t.Errorf("Should have auth token %s but got %s", token, authHeader)
	}
}

func TestRequestStatusError(t *testing.T) {
	mocks.MockHTTPClient.SetResponse("[]", 500, nil)
	_, err := request(http.MethodGet, "", "", nil, nil)
	if err == nil {
		t.Error("Should fail request with server error")
	}
}

func TestRequestJSONError(t *testing.T) {
	holder := struct{}{}
	mocks.MockHTTPClient.SetResponse("*&(^", 200, nil)
	_, err := request(http.MethodGet, "", "", nil, holder)
	if err == nil {
		t.Error("Should fail request with JSON error")
	}
}

func TestRequestHTTPError(t *testing.T) {
	mocks.MockHTTPClient.SetResponse("", 200, errors.New("deliberate error in TestRequestHTTPError"))
	_, err := request(http.MethodGet, "", "", nil, nil)
	if err == nil {
		t.Error("Should fail request with HTTP error")
	}
}
