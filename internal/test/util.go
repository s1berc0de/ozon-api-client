package test

import (
	"net/http"
)

const (
	ApiKey   = "my-api-key"
	ClientID = "my-client-id"
)

var _ http.RoundTripper = (*RoundTripFunc)(nil)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn http.RoundTripper) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}
