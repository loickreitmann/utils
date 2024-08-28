package utils

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

// table test setup
var uploadTests = []struct {
	name          string
	allowedTypes  []string
	testFiles     []string
	renameFiles   bool
	errorExpected bool
}{
	{
		name:          "allowed, not renamed",
		allowedTypes:  []string{"image/jpeg", "image/png", "image/gif"},
		testFiles:     []string{"./testdata/upload_test.jpg", "./testdata/upload_test.gif", "./testdata/upload_test.png"},
		renameFiles:   false,
		errorExpected: false,
	},
	{
		name:          "allowed, renamed",
		allowedTypes:  []string{"image/jpeg", "image/png", "image/gif"},
		testFiles:     []string{"./testdata/upload_test.gif", "./testdata/upload_test.jpg", "./testdata/upload_test.png"},
		renameFiles:   true,
		errorExpected: false,
	},
	{
		name:          "not allowed",
		allowedTypes:  []string{"image/jpeg", "image/gif"},
		testFiles:     []string{"./testdata/upload_test.png"},
		renameFiles:   false,
		errorExpected: true,
	},
	{
		name:          "all types allowed, not renamed",
		allowedTypes:  []string{},
		testFiles:     []string{"./testdata/upload_test.png", "./testdata/upload_test.gif", "./testdata/upload_test.jpg"},
		renameFiles:   false,
		errorExpected: false,
	},
}

func TestUtils_UploadFiles(t *testing.T) {
	wg := sync.WaitGroup{}
	for _, ut := range uploadTests {
		for _, testFile := range ut.testFiles {
			// ARRANGE
			// set up a pipe to avoid buffering
			pr, pw := io.Pipe()
			writer := multipart.NewWriter(pw)
			wg.Add(1)

			go func() {
				defer writer.Close()
				defer wg.Done()

				// create the form data field 'file'
				part, err := writer.CreateFormFile("file", testFile)
				if err != nil {
					t.Error(err)
				}
				file, err := os.Open(testFile)
				if err != nil {
					t.Error(err)
				}
				defer file.Close()

				img, _, err := image.Decode(file)
				if err != nil {
					t.Error("error decoding image", err)
				}
				ext := filepath.Ext(testFile)
				switch ext {
				case ".png":
					if err = png.Encode(part, img); err != nil {
						t.Error(err)
					}
				case ".gif":
					if err = gif.Encode(part, img, nil); err != nil {
						t.Error(err)
					}
				case ".jpg":
					if err = jpeg.Encode(part, img, nil); err != nil {
						t.Error(err)
					}
				}
			}()

			// read fro the pipe which receives data
			request := httptest.NewRequest("POST", "/", pr)
			request.Header.Add("Content-Type", writer.FormDataContentType())

			// ACT
			testUtils := New()
			testUtils.AllowedTypes = ut.allowedTypes

			uploadedFiles, err := testUtils.UploadFiles(request, "./testdata/uploads/", ut.renameFiles)
			if err != nil {
				t.Error(err)
			}

			// ASSERT
			// Positive test: should not result in an error
			if !ut.errorExpected {
				if err != nil {
					t.Errorf("[%s] error unexpected: %s", ut.name, err.Error())
				}
				expectedFile := fmt.Sprintf("./testdata/uploads/%s", uploadedFiles[0].NewFilename)
				if _, fileErr := os.Stat(expectedFile); os.IsNotExist(fileErr) {
					t.Errorf("[%s] expected file to exist: %s", ut.name, fileErr.Error())
				}
				// clean up
				_ = os.Remove(expectedFile)
			}
			wg.Wait()
		}
	}
}
