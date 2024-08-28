package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

/*
MakeDirStructure is a convenience method of the `utils` package which uses `os.MkdirAll` to creates
a directory structure based on the slice of path strings provided. It returns nil when all the
directories are successfully created, or else returns an error. The permission bits default to 0755,
and are used for all directories created. If a path is already a directory, os.MkdirAll does nothing
and returns nil, so MakeDirStructure will also return nil. If there's a permission issue encountered
for any of the paths, the error reported by os.MkdirAll will be collected, and MakeDirStructure will
return all encountered those errors as one.
*/
func (u *Utils) MakeDirStructure(pathsToBeMade []string) error {
	var mode fs.FileMode = 0755
	var errorMsgs string
	for _, pathStr := range pathsToBeMade {
		err := os.MkdirAll(pathStr, mode)
		if err != nil {
			errorMsgs += fmt.Sprintf("could not create %s: %s\n", pathStr, err.Error())
		}
	}
	if errorMsgs != "" {
		return errors.New(errorMsgs)
	}
	return nil
}
