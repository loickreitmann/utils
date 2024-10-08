package utils_test

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/loickreitmann/utils"
)

// table test setup
var uploadTests = []struct {
	name          string
	allowedTypes  []string
	testFiles     []string
	renameFiles   bool
	errorExpected bool
	maxUploadSize int
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
		errorExpected: false,
	},
	{
		name:          "file too big",
		allowedTypes:  []string{"image/jpeg", "image/png", "image/gif"},
		testFiles:     []string{"./testdata/upload_test.png"},
		renameFiles:   false,
		errorExpected: true,
		maxUploadSize: 1,
	},
	{
		name:          "all types allowed, not renamed",
		allowedTypes:  []string{},
		testFiles:     []string{"./testdata/upload_test.png", "./testdata/upload_test.gif", "./testdata/upload_test.jpg"},
		renameFiles:   false,
		errorExpected: false,
	},
}

const (
	uploadsRoot     string = "./testdata/uploads"
	uploadsBasePath string = uploadsRoot + "/"
)

func pipeFile(writer *multipart.Writer, testFile string, t *testing.T, wg *sync.WaitGroup) {
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
}

func TestUtils_UploadOneFile(t *testing.T) {
	wg := sync.WaitGroup{}
	for _, ut := range uploadTests {
		// ARRANGE
		if ut.maxUploadSize > 0 {
			continue
		}
		testFile := ut.testFiles[0]
		// set up a pipe to avoid buffering
		pr, pw := io.Pipe()
		writer := multipart.NewWriter(pw)
		wg.Add(1)

		go pipeFile(writer, testFile, t, &wg)

		// read fro the pipe which receives data
		request := httptest.NewRequest(http.MethodPost, "/", pr)
		request.Header.Add("Content-Type", writer.FormDataContentType())

		// ACT
		testUtils := utils.New()
		testUtils.AllowedTypes = ut.allowedTypes

		uploadedFile, err := testUtils.UploadOneFile(request, uploadsBasePath, ut.renameFiles)
		if err != nil {
			t.Error(err)
		}

		// ASSERT
		// Positive test: should not result in an error
		if !ut.errorExpected {
			if err != nil {
				t.Errorf("[%s] error unexpected: %s", ut.name, err.Error())
			}
			expectedFile := fmt.Sprintf("./testdata/uploads/%s", uploadedFile.NewFilename)
			if _, fileErr := os.Stat(expectedFile); os.IsNotExist(fileErr) {
				t.Errorf("[%s] expected file to exist: %s", ut.name, fileErr.Error())
			}
		}
		wg.Wait()
	}
	// CLEAN UP
	_ = os.RemoveAll(uploadsRoot)
}

func TestUtils_UploadFiles(t *testing.T) {
	wg := sync.WaitGroup{}
	for _, ut := range uploadTests {
		for _, testFile := range ut.testFiles {
			// ARRANGE
			// set up a pipe to avoid buffering
			pr, pw := io.Pipe()
			writer := multipart.NewWriter(pw)

			var request *http.Request
			if ut.maxUploadSize > 0 {
				// define a dummy large size request
				largeBody := bytes.Repeat([]byte("a"), 1024)
				request = httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(largeBody))
			} else {
				wg.Add(1)
				go pipeFile(writer, testFile, t, &wg)
				// read from the pipe which receives data
				request = httptest.NewRequest(http.MethodPost, "/", pr)
			}
			request.Header.Add("Content-Type", writer.FormDataContentType())

			// ACT
			testUtils := utils.New()
			testUtils.AllowedTypes = ut.allowedTypes
			if ut.maxUploadSize > 0 {
				testUtils.MaxUploadFileSize = ut.maxUploadSize
			}

			uploadedFiles, err := testUtils.UploadFiles(request, uploadsBasePath, ut.renameFiles)
			// ASSERT
			if !ut.errorExpected {
				if err != nil {
					t.Errorf("[%s] error not expected: %s", ut.name, err.Error())
				}
				expectedFile := fmt.Sprintf("./testdata/uploads/%s", uploadedFiles[0].NewFilename)
				if _, fileErr := os.Stat(expectedFile); os.IsNotExist(fileErr) {
					t.Errorf("[%s] expected file to exist: %s", ut.name, fileErr.Error())
				}
			}
			if ut.errorExpected && err == nil {
				t.Errorf("[%s] error expected but none received", ut.name)
			}
			wg.Wait()
		}
	}
	// CLEAN UP
	_ = os.RemoveAll(uploadsRoot)
}
