# Utils Module
Module of simple, well-tested, and reusable Go utilities. Inspired by [Trevor Sawler](https://www.udemy.com/user/trevor-sawler/)'s "[Building a Module in Go.](https://www.udemy.com/course/building-a-module-in-go-golang)" Udemy course.

Most of the utilities in this module are based on Trevor Sawler's [**Toolbox** module](https://github.com/tsawler/toolbox).

#### The utilities in this module:
##### RandomString()
```go
Utils.RandomString(length int) string
```
Generates a random string of the given `length`.

##### UploadOneFile()
```go
Utils.UploadOneFile(req *http.Request, uploadDir string, rename ...bool) (*UploadedFile, error)
```
A convenience method that calls `UploadFiles()`, but expects only one file to be in the upload.
##### UploadFiles()
```go
Utils.UploadFiles(req *http.Request, uploadDir string, rename ...bool) ([]*UploadedFile, error)
```
Uploads one or more files from a multipart form submission contained within an `http.Request` to the specified `uploadDir` directory. It gives the files a random name. It returns a slice of `UploadedFile` structs, and potentially an `error`. If the optional last parameter is set to `true`, the files won't be renamed.


## Installation

`go get -u github.com/loickreitmann/utils`

## Usage
