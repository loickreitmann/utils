package utils

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

func (u *Utils) CrawlLogPaths(root string) error {
	return filepath.Walk(root, walkFunc)
}

func walkFunc(path string, info fs.FileInfo, err error) error {
	if err != nil {
		// Handle error while accessing a file or directory
		log.Printf("Error accessing path %s: %v", path, err)
		return err
	}

	// Log the full path of the current file or directory
	fmt.Println(path)

	// Continue walking the directory structure
	return nil
}
