// Package of simple and well tested reusable Go utilities.
// Inspired by Trevor Sawler's "Building a Module in Go." Udemy course.
package utils

// Utils is the type used to instantiate the module. Any variable of this type will have access to all
// the methods with the *Utils receiver.
type Utils struct {
	UploadOptions
}

func New() *Utils {
	return &Utils{
		UploadOptions{
			MaxUploadFileSize: gB,
			AllowedTypes:      []string{"image/jpeg", "image/png", "image/gif"},
		},
	}
}
