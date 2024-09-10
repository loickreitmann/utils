package utils_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/loickreitmann/utils"
)

// set up a client that can send back an arbitrary response.
// so we can write tests without actually having a remote API active.
// 1. RoundTripFunc used to satisfy the interface requirements for HTTP client.
type RoundTripFunc func(req *http.Request) *http.Response

// 2. RoundTrip returns a function that takes a request as a parameter.
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// 3. NewTestClient takes a parameter function of type RoundTripFunc and returns a pointer to HTTP client.
// This returns a reference to a HTTP client from the standard library, but substitutes default transport
// that's built into the standard library our function.
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func TestUtils_PushJSONToRemote(t *testing.T) {
	// ARRANGE
	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test Request Parameters
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString("ok")),
			Header:     make(http.Header),
		}
	})

	var testUtils utils.Utils
	var foo struct {
		Bar string `json:"bar"`
	}
	foo.Bar = "baz"

	_, _, err := testUtils.PushJSONToRemote("https://example.com/some/api", http.MethodPost, foo, client)
	if err != nil {
		t.Error("failed to call remote url:", err)
	}
}
