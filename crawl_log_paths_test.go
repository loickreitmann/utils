package utils_test

import (
	"testing"

	"github.com/loickreitmann/utils"
)

var testPaths = map[string]string{
	"pass": "./testdata",
	"fail": "./notARealPath",
}

func TestUtils_CrawlLogPaths(t *testing.T) {
	// ARRANGE
	var testUtils utils.Utils

	for expectation, testPath := range testPaths {
		// ACT
		err := testUtils.CrawlLogPaths(testPath)

		// ASSESS
		switch expectation {
		case "pass":
			if err != nil {
				t.Errorf("expected to pass, but failed for path %s", testPath)
			}
		case "fail":
			if err == nil {
				t.Errorf("expected to fail, but passed for path %s", testPath)
			}
		}
	}
}
