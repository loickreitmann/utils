// Package of simple and well tested reusable Go utilities.
// Inspired by Trevor Sawler's "Building a Module in Go." Udemy course.
package utils

// Utils is the type used to instantiate the module. Any variable of this type will
// have access to all the methods with the *Utils receiver.
type Utils struct {
	UploadOptions
	JSONOptions
}

// Constructor for Utils struct that defines default options necessary when using
// the UploadFiles and UploadOneFile methods.
// MaxUploadFileSize: 1GB.
// AllowedTypes: JPG, PNG, and GIF
func New() *Utils {
	return &Utils{
		UploadOptions{
			MaxUploadFileSize: gB,
			AllowedTypes:      []string{"image/jpeg", "image/png", "image/gif"},
		},
		JSONOptions{
			MaxJSONReadSize:    mB,
			AllowUnknownFields: false,
		},
	}
}
