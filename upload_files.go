package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// constants defining size increments
const (
	_ = 1 << (iota * 10)
	kB
	mB
	gB
)

type UploadOptions struct {
	MaxUploadFileSize int
	AllowedTypes      []string
}

func (u *Utils) isAllowedType(buffer []byte) bool {
	fileType := http.DetectContentType(buffer)
	if len(u.AllowedTypes) > 0 {
		// AllowedTypes are defined, so check if the current file matches an allowed type
		for _, allowedType := range u.AllowedTypes {
			if strings.EqualFold(fileType, allowedType) {
				return true
			}
		}
	} else {
		// All file types are allowed. Not very safe.
		return true
	}
	return false
}

func (u *Utils) UploadFiles(req *http.Request, uploadDir string, rename ...bool) ([]*UploadedFile, error) {
	renameFiles := true
	if len(rename) > 0 {
		renameFiles = rename[0]
	}

	var uploadedFiles []*UploadedFile
	err := req.ParseMultipartForm(int64(u.MaxUploadFileSize))
	if err != nil {
		return nil, errors.New("the uploaded file is too big")
	}

	uploadedFiles, err = func(uploadedFiles []*UploadedFile) ([]*UploadedFile, error) {
		for _, fileHeaders := range req.MultipartForm.File {
			for _, hdr := range fileHeaders {
				var uploadedFile UploadedFile
				currentFile, err := hdr.Open()
				if err != nil {
					return nil, err
				}
				defer currentFile.Close()

				// read the fist 512 bytes to identify the file type
				buffer := make([]byte, 512)
				_, err = currentFile.Read(buffer)
				if err != nil {
					return nil, err
				}

				// check if the file type is allowed for upload receipt
				isAllowed := u.isAllowedType(buffer)
				if !isAllowed {
					// return nil, errors.New("the uploaded file type is not allowed")
					// instead of failing out, ignore this file, skip out of this iteration and continue
					continue
				}

				_, err = currentFile.Seek(0, 0)
				if err != nil {
					return nil, err
				}

				// set the uploaded file's new file name
				if renameFiles {
					uploadedFile.NewFilename = fmt.Sprintf("%s%s", u.RandomString(32), filepath.Ext(hdr.Filename))
				} else {
					uploadedFile.NewFilename = hdr.Filename
				}

				// prepare the output file
				var outputFile *os.File
				defer outputFile.Close()
				if outputFile, err := os.Create(filepath.Join(uploadDir, uploadedFile.NewFilename)); err != nil {
					return nil, err
				} else {
					fileSize, err := io.Copy(outputFile, currentFile)
					if err != nil {
						return nil, err
					}
					uploadedFile.FileSize = fileSize
				}
				uploadedFiles = append(uploadedFiles, &uploadedFile)

			}
		}
		return uploadedFiles, nil
	}(uploadedFiles)
	if err != nil {
		return uploadedFiles, err
	}

	return uploadedFiles, nil
}
