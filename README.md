# Utils Module

[![Go Reference](https://pkg.go.dev/badge/github.com/loickreitmann/utils.svg)](https://pkg.go.dev/github.com/loickreitmann/utils) [![Coverage Status](https://coveralls.io/repos/github/loickreitmann/utils/badge.svg?branch=main)](https://coveralls.io/github/loickreitmann/utils?branch=main)

Module of simple, well-tested, and reusable Go utilities. Inspired by [Trevor Sawler](https://www.udemy.com/user/trevor-sawler/)'s "[Building a Module in Go.](https://www.udemy.com/course/building-a-module-in-go-golang)" Udemy course.

Most of the utilities in this module are based on Trevor Sawler's [**Toolbox** module](https://github.com/tsawler/toolbox).

---
## The utilities in this module:
#### RandomString()
RandomString generates and returns a string of a specified `length` of random characters. The characters are randomnly asembled from othe ptional `characterSeed` parameter, or they are sourced from the default set made up of the following characters:
`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_+!@#$%*`
```go
func (u *Utils) RandomString(length int, characterSeed ...[]rune) string
```

#### UploadOneFile()
UploadOneFile is a convenience method that calls `UploadFiles`, but expects only one file to be in the upload.
```go
func (u *Utils) UploadOneFile(r *http.Request, uploadDir string, rename ...bool) (*UploadedFile, error) 
```

#### UploadFiles()
UploadFiles uploads one or more files from a multipart form submission contained within an `http.Request` to the specified `uploadDir` directory. It gives the files a random name. It returns a slice of `UploadedFile` structs, and potentially an `error`. If the optional last parameter is set to `true`, the files won't be renamed.
```go
func (u *Utils) UploadFiles(req *http.Request, uploadDir string, rename ...bool) ([]*UploadedFile, error)
```

#### MakeDirStructure()
MakeDirStructure is a convenience method of the `utils` package which uses `os.MkdirAll` to creates a directory structure based on the slice of path strings provided. It returns nil when all the directories are successfully created, or else returns an error. The permission bits default to `0755`, and are used for all directories created. If a path is already a directory, `os.MkdirAll` does nothing and returns `nil`, so MakeDirStructure will also return `nil`. If there's a permission issue encountered for any of the paths, the error reported by `os.MkdirAll` will be collected, and MakeDirStructure will return all encountered those errors as one.
```go
func (u *Utils) MakeDirStructure(pathsToBeMade []string) error
```

#### CrawlLogPaths()
CrawlLogPaths: given a starting path, it will crawl the directory hierachy below that path, and output a log message of each full path from the specified starting path.
```go
func (u *Utils) CrawlLogPaths(root string) error
```

#### TextToSlug()
The TextToSlug function converts accented characters to their unaccented versions, replaces all non-alphanumeric characters with dashes, trims redundant dashes, and converts the string to lowercase.
This approach makes the slug both URL-friendly and human-readable.
```go
func (u *Utils) TextToSlug(input string) string 
```

#### ForceFileDownload()
ForceFileDownload forces the browser to avoid displaying it in the browser window by setting the `Content-Disposition` header. It also allows specifying a custom display name for the downloaded file.
```go
func (u *Utils) ForceFileDownload(w http.ResponseWriter, r *http.Request, pathToFile, displayName string) 
```

#### ReadJSON()
ReadJSON tries to read the body of a request and converts from json into a go `data` variable.
```go
func (u *Utils) ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error 
```

#### WriteJSON()
WriteJSON takes a response, an `httpStatus` code, and arbitrary `data`, then generates and sends json in the http response to the client.
```go
func (u *Utils) WriteJSON(w http.ResponseWriter, httpStatus int, data interface{}, headers ...http.Header) error
```

#### ErrorJSON()
ErrorJSON takes an `error` and optionally an http `httpStatus` code, then generates and sends a json formatted error http response. If no `httpStatus` code is passed, `http.StatusBadRequest` is the defualt used.
```go
func (u *Utils) ErrorJSON(w http.ResponseWriter, err error, httpStatus ...int) error
```

#### PushJSONToRemote()
PushJSONToRemote sends arbitrary `data` to a specified URL as JSON, and returns the `response` and http `status` code, or an `error` if any.
The `client` parameter is optional. If none is specified, it uses the standard library's `http.Client`.
```go
func (u *Utils) PushJSONToRemote(uri string, method string, data interface{}, client ...*http.Client) (*http.Response, int, error)
```

#### LoadEnvVarsFromFile()
LoadEnvVarsFromFile expects a string represnting the path to the environment variable file.
This approach relies on the environment variables file existing in the file system and being readable.
```go
func (u *Utils) LoadEnvVarsFromFile(filename string) error
```

#### LoadEnvVarsFromEmbed()
LoadEnvVarsFromEmbed expects a string resulting from having uses Go's `//go:embed` directive to import an embedded environment variable file.
This approach relies on the environment variables file being embedded directly in the binary, which might not be ideal for sensitive data in some cases.
```go
func (u *Utils) LoadEnvVarsFromEmbed(goEmbedReadFile string) error
```

#### UniqueRunes()
UniqueRunes takes a string and converts it to a slice of unique runes. This method preserves the order of the first appearance of runes in the input string. Because Go's rune type handles Unicode characters properly, this solution works correctly even with strings containing non-ASCII characters.
```go
func (u *Utils) UniqueRunes(s string) []rune 
```

#### ContainsAllRunes()
ContainsAllRunes checks if all runes in a `subset` slice are contained in `set` slice. This method correctly handles cases where runes are duplicated, both in the set and subset slices. This works with Unicode characters, as Go's rune type properly supports Unicode.
```go
func (u *Utils) ContainsAllRunes(set, subset []rune) bool
```

---
## Installation

`go get -u github.com/loickreitmann/utils`

---
## Usage
See the example test files, or the package doumentation on the [Go Package Discovery](https://pkg.go.dev/github.com/loickreitmann/utils) site.
