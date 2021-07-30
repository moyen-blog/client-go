package mocks

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type mockHTTPClient struct {
	LastRequest *http.Request
	response    *http.Response
	err         error
}

func (m *mockHTTPClient) SetResponse(body string, code int, err error) {
	r := ioutil.NopCloser(strings.NewReader(body))
	var response *http.Response = &http.Response{
		StatusCode: code,
		Status:     http.StatusText(code),
		Body:       r,
	}
	m.response = response
	m.err = err
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	r, e := m.response, m.err
	m.response, m.err = nil, nil
	m.LastRequest = req
	return r, e
}

var MockHTTPClient mockHTTPClient = mockHTTPClient{}
