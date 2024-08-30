package utils_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/loickreitmann/utils"
)

func TestUtils_ForceFileDownload(t *testing.T) {
	// ARRANGE
	respRec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	var utilsTest utils.Utils

	// ACT
	utilsTest.ForceFileDownload(respRec, req, "./testdata", "upload_test.gif", "minions.gif")
	resp := respRec.Result()
	defer resp.Body.Close()

	// ASSERT
	if resp.Header["Content-Type"][0] != "image/gif" {
		t.Errorf("wrong content type: expected image/gif, got %s", resp.Header["Content-Type"][0])
	}
	if resp.Header["Content-Length"][0] != "371779" {
		t.Errorf("wrong content length: expected 371779, got %s", resp.Header["Content-Length"][0])
	}
	if resp.Header["Content-Disposition"][0] != "attachment; filename=\"minions.gif\"" {
		t.Errorf("wrong content disposition: expected attachment; filename=\"minions.gif\", got %s", resp.Header["Content-Disposition"][0])
	}
	_, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
}
