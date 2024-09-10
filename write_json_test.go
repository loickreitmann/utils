package utils_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/loickreitmann/utils"
)

type Node struct {
	Value    string
	NextNode *Node
}

var (
	sampleResponseHeaderProps = map[string]string{
		"X-Content-Type-Options": "nosniff",
		"X-Frame-Options":        "deny",
	}

	goodJSON = utils.JSONResponse{
		Error: nil,
		Data: struct {
			Foo string
		}{
			Foo: "bar",
		},
	}

	errorJSON = utils.JSONResponse{
		Error: utils.JSONError{
			Code:    fmt.Sprint(http.StatusForbidden),
			Message: "unauthorized request",
		},
		Data: nil,
	}

	testCases = []struct {
		name          string
		errorExpected bool
		status        int
		payload       utils.JSONResponse
		headerProps   map[string]string
	}{
		{
			name:          "send good payload",
			errorExpected: false,
			status:        http.StatusOK,
			payload:       goodJSON,
			headerProps:   sampleResponseHeaderProps,
		},
		{
			name:          "send error payload",
			errorExpected: false,
			status:        http.StatusForbidden,
			payload:       errorJSON,
			headerProps:   sampleResponseHeaderProps,
		},
		{
			name:          "no custom headers",
			errorExpected: false,
			status:        http.StatusForbidden,
			payload:       goodJSON,
			headerProps:   nil,
		},
		{
			name:          "cyclic data structure to cause marshalling error",
			errorExpected: true,
			status:        http.StatusForbidden,
			payload: func() utils.JSONResponse {
				node1 := &Node{Value: "first"}
				node2 := &Node{Value: "second", NextNode: node1}
				node1.NextNode = node2
				jsonResp := utils.JSONResponse{
					Data: node1,
				}
				return jsonResp
			}(),
			headerProps: nil,
		},
	}
)

func TestUtils_WriteJSON(t *testing.T) {
	var testUtils utils.Utils
	var err error
	for _, testCase := range testCases {
		// ARRANGE
		rr := httptest.NewRecorder()
		headers := make(http.Header)
		for prop, value := range testCase.headerProps {
			headers.Add(prop, value)
		}

		// ACT
		if testCase.headerProps != nil {
			err = testUtils.WriteJSON(rr, testCase.status, testCase.payload, headers)
		} else {
			err = testUtils.WriteJSON(rr, testCase.status, testCase.payload)
		}

		// ASSESS
		if !testCase.errorExpected && err != nil {
			t.Errorf("[%s]: failed to write JSON: %v", testCase.name, err)
		}
		if testCase.errorExpected && err == nil {
			t.Errorf("[%s]: exxpected an error for this test but got none", testCase.name)
		}
	}
}
