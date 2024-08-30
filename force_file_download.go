package utils

import (
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

// sanitizeFileName ensures that the filename is safe and does not contain any
// special characters
func (u *Utils) sanitizeFileName(name string) string {
	// Replace any characters that might interfere with the header
	// This can be extended to remove more dangerous or undesired characters
	return strings.Map(func(r rune) rune {
		if r == '/' || r == '\\' || r == '"' || r == '\'' || r == ':' {
			return '-'
		}
		return r
	}, name)
}

// ForceFileDownload forces the browser to avoid displaying it in the browser window
// by setting the Content-Disposition header. It also allows specifying a custom
// display name for the downloaded file.
func (u *Utils) ForceFileDownload(w http.ResponseWriter, r *http.Request, fileDir, file, displayName string) {
	// Ensure file paths are safe and use the correct path separator
	fullFilePath := filepath.Join(fileDir, filepath.Clean(file))

	// Attempt to detect the file's MIME type
	mimeType := mime.TypeByExtension(filepath.Ext(file))
	if mimeType == "" {
		mimeType = "application/octet-stream" // Default to binary stream if unknown
	}
	w.Header().Set("Content-Type", mimeType)

	// Set Content-Disposition header for file download with sanitized display name
	displayName = u.sanitizeFileName(displayName)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", displayName))

	// Attempt to serve the file
	http.ServeFile(w, r, fullFilePath)
}
