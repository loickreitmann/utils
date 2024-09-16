package utils

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

// CrawlLogPaths: given a starting path, it will crawl the directory
// hierachy below that path, and output a log message of each full
// path from the specified starting path.
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
