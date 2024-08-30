# Utils Module
Module of simple, well-tested, and reusable Go utilities. Inspired by [Trevor Sawler](https://www.udemy.com/user/trevor-sawler/)'s "[Building a Module in Go.](https://www.udemy.com/course/building-a-module-in-go-golang)" Udemy course.

Most of the utilities in this module are based on Trevor Sawler's [**Toolbox** module](https://github.com/tsawler/toolbox).

---
#### The utilities in this module:
##### 1. RandomString()
```go
func (u *Utils) RandomString(length int) string
```
Generates a random string of the given `length`.

##### 2. UploadOneFile()
```go
func (u *Utils) UploadOneFile(req *http.Request, uploadDir string, rename ...bool) (*UploadedFile, error) 
```
A convenience method that calls `UploadFiles()`, but expects only one file to be in the upload.

##### 3. UploadFiles()
```go
func (u *Utils) UploadFiles(req *http.Request, uploadDir string, rename ...bool) ([]*UploadedFile, error)
```
Uploads one or more files from a multipart form submission contained within an `http.Request` to the specified `uploadDir` directory. It gives the files a random name. It returns a slice of `UploadedFile` structs, and potentially an `error`. If the optional last parameter is set to `true`, the files won't be renamed.

##### 4. MakeDirStructure()
```go
func (u *Utils) MakeDirStructure(pathsToBeMade []string) error
```
MakeDirStructure is a convenience method of the `utils` package which uses `os.MkdirAll` to creates a directory structure based on the slice of path strings provided. It returns nil when all the directories are successfully created, or else returns an error. The permission bits default to 0755, and are used for all directories created. If a path is already a directory, os.MkdirAll does nothing and returns nil, so MakeDirStructure will also return nil. If there's a permission issue encountered for any of the paths, the error reported by os.MkdirAll will be collected, and MakeDirStructure will return all encountered those errors as one.

##### 5. TextToSlug()
The TextToSlug function converts accented characters to their unaccented versions, replaces all non-alphanumeric characters with dashes, trims redundant dashes, and converts the string to lowercase.
This approach makes the slug both URL-friendly and human-readable.
```go
func (u *Utils) TextToSlug(input string) string 
```

---
## Installation

`go get -u github.com/loickreitmann/utils`

---
## Usage
