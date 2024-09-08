package utils_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/loickreitmann/utils"
)

var jsonTestCases = []struct {
	name          string
	json          string
	errorExpected bool
	maxSize       int
	allowUnknowns bool
	asPointer     bool
}{
	{name: "valid json", json: `{"foo": "bar", "buzz": 1}`, errorExpected: false, maxSize: 1024, allowUnknowns: false, asPointer: true},
	{name: "max size not set", json: `{"foo": "bar", "buzz": 1}`, errorExpected: false, maxSize: 0, allowUnknowns: false, asPointer: true},
	{name: "two json structures", json: `{"foo": "bar"}{"buzz": 1}`, errorExpected: true, maxSize: 1024, allowUnknowns: false, asPointer: true},
	{name: "valid json, with ommisisons", json: `{"foo": "bar"}`, errorExpected: false, maxSize: 1024, allowUnknowns: false, asPointer: true},
	{name: "incorrect type", json: `{"foo": "bar", "buzz": "1"}`, errorExpected: true, maxSize: 1024, allowUnknowns: false, asPointer: true},
	{name: "unknown field in json", json: `{"foo": "bar", "busz": 1}`, errorExpected: true, maxSize: 1024, allowUnknowns: false, asPointer: true},
	{name: "empty body", json: ``, errorExpected: true, maxSize: 1024, allowUnknowns: false, asPointer: true},
	{name: "syntax error in json", json: `{"foo": "bar", "buzz": 1"}`, errorExpected: true, maxSize: 1024, allowUnknowns: false, asPointer: true},
	{name: "unexpected end of file", json: `{"foo": "bar", "buzz": 1`, errorExpected: true, maxSize: 1024, allowUnknowns: false, asPointer: true},
	{name: "badly formed json", json: `"foo": "bar", "buzz": 1}`, errorExpected: true, maxSize: 1024, allowUnknowns: false, asPointer: true},
	{name: "file too large", json: `{"foo": "bar", "buzz": 1}`, errorExpected: true, maxSize: 10, allowUnknowns: false, asPointer: true},
	{name: "not json", json: `Hello, there!`, errorExpected: true, maxSize: 1024, allowUnknowns: false, asPointer: true},
	{name: "not as pointer", json: `{"foo": "bar", "buzz": 1}`, errorExpected: true, maxSize: 1024, allowUnknowns: false, asPointer: false},
}

func TestUtils_ReadJSON(t *testing.T) {
	var testUtils utils.Utils

	for _, testCase := range jsonTestCases {
		// ARRANGE
		// set JSON max size
		testUtils.MaxJSONReadSize = testCase.maxSize

		// allow/disallow unknown fields
		testUtils.AllowUnknownFields = testCase.allowUnknowns

		// declare variable to read the decoded json into
		var decodedJSON struct {
			Foo  string `json:"foo"`
			Buzz int    `json:"buzz,omitempty"`
		}

		// ACT
		// create a request with the body
		req, err := http.NewRequest("POST", "/", bytes.NewReader([]byte(testCase.json)))
		if err != nil {
			t.Log("error:", err.Error())
		}

		// create a response recorder
		rr := httptest.NewRecorder()

		if testCase.asPointer {
			err = testUtils.ReadJSON(rr, req, &decodedJSON)
		} else {
			err = testUtils.ReadJSON(rr, req, decodedJSON)
		}
		if err != nil {
			t.Log(err.Error())
		}

		// ASSESS
		if testCase.errorExpected && err == nil {
			t.Errorf("[%s]: error expected, but none received", testCase.name)
		}

		if !testCase.errorExpected && err != nil {
			t.Errorf("[%s]: error not expected, but one received", testCase.name)
		}

		req.Body.Close()
	}
}
