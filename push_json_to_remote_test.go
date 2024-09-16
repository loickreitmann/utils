package utils_test

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/loickreitmann/utils"
)

// set up a client that can send back an arbitrary response.
// so we can write tests without actually having a remote API active.
// 1. RoundTripFunc used to satisfy the interface requirements for HTTP client.
type RoundTripFunc func(req *http.Request) (*http.Response, error)

// 2. RoundTrip returns a function that takes a request as a parameter.
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

// 3. NewTestClient takes a parameter function of type RoundTripFunc and returns a pointer to HTTP client.
// This returns a reference to a HTTP client from the standard library, but substitutes default transport
// that's built into the standard library our function.
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var foo = struct {
	Bar string `json:"bar"`
}{Bar: "baz"}

var pushTestCases = []struct {
	name        string
	payload     interface{}
	expectError bool
}{
	{name: "successful push", payload: foo, expectError: false},
	{name: "unsuccessful push due to cyclic data structure", payload: cyclicDataStructure(), expectError: true},
	{name: "response time out", payload: foo, expectError: true},
}

func TestUtils_PushJSONToRemote(t *testing.T) {
	var testUtils utils.Utils
	for _, pushCase := range pushTestCases {
		// ARRANGE
		client := NewTestClient(func(req *http.Request) (*http.Response, error) {
			if pushCase.name == "response time out" {
				// Simulate a timeout by returning a timeout error
				return nil, errors.New("timeout: request canceled while waiting for a response")
			}
			// Default behavior: return a successful response
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString("ok")),
				Header:     make(http.Header),
			}, nil
		})
		// ACT
		_, _, err := testUtils.PushJSONToRemote("https://example.com/some/api", http.MethodPost, pushCase.payload, client)
		// ASSERT
		if !pushCase.expectError && err != nil {
			t.Errorf("[%s]: expected no error but got one: %v", pushCase.name, err)
		}
		if pushCase.expectError && err == nil {
			t.Errorf("[%s]: expected an error but got none", pushCase.name)
		}
	}
}
