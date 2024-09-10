package utils_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/loickreitmann/utils"
)

var testFiles = []struct {
	name                string
	fileName            string
	displayName         string
	expectedDisplayName string
	mimeExpected        string
}{
	{name: "gif", fileName: "upload_test.gif", displayName: "mini:ons.gif", expectedDisplayName: "mini-ons.gif", mimeExpected: "image/gif"},
	{name: "mimetype unknown", fileName: "upload_test", displayName: "minions", expectedDisplayName: "minions", mimeExpected: "application/octet-stream"},
}

func TestUtils_ForceFileDownload(t *testing.T) {
	var utilsTest utils.Utils
	for _, testFile := range testFiles {
		// ARRANGE
		respRec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)

		// ACT
		utilsTest.ForceFileDownload(respRec, req, "./testdata", testFile.fileName, testFile.displayName)
		resp := respRec.Result()
		defer resp.Body.Close()

		// ASSERT
		if resp.Header["Content-Type"][0] != testFile.mimeExpected {
			t.Errorf("[%s]: wrong content type: expected %s, got %s", testFile.name, testFile.mimeExpected, resp.Header["Content-Type"][0])
		}
		if resp.Header["Content-Length"][0] != "371779" {
			t.Errorf("[%s]: wrong content length: expected 371779, got %s", testFile.name, resp.Header["Content-Length"][0])
		}
		if resp.Header["Content-Disposition"][0] != fmt.Sprintf("attachment; filename=\"%s\"", testFile.expectedDisplayName) {
			t.Errorf("[%s]: wrong content disposition: expected attachment; filename=\"%s\", got %s", testFile.name, testFile.expectedDisplayName, resp.Header["Content-Disposition"][0])
		}
		_, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}
	}
}
