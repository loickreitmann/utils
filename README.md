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
func (u *Utils) UploadOneFile(r *http.Request, uploadDir string, rename ...bool) (*UploadedFile, error) 
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
MakeDirStructure is a convenience method of the `utils` package which uses `os.MkdirAll` to creates a directory structure based on the slice of path strings provided. It returns nil when all the directories are successfully created, or else returns an error. The permission bits default to `0755`, and are used for all directories created. If a path is already a directory, `os.MkdirAll` does nothing and returns `nil`, so MakeDirStructure will also return `nil`. If there's a permission issue encountered for any of the paths, the error reported by `os.MkdirAll` will be collected, and MakeDirStructure will return all encountered those errors as one.

##### 5. TextToSlug()
The TextToSlug function converts accented characters to their unaccented versions, replaces all non-alphanumeric characters with dashes, trims redundant dashes, and converts the string to lowercase.
This approach makes the slug both URL-friendly and human-readable.
```go
func (u *Utils) TextToSlug(input string) string 
```

##### 6. ForceFileDownload()
ForceFileDownload forces the browser to avoid displaying it in the browser window by setting the `Content-Disposition` header. It also allows specifying a custom display name for the downloaded file.
```go
func (u *Utils) ForceFileDownload(w http.ResponseWriter, r *http.Request, fileDir, file, displayName string) 
```

##### 7. ReadJSON()
ReadJSON tries to read the body of a request and converts from json into a go `data` variable.
```go
func (u *Utils) ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error 
```

##### 8. WriteJSON()
WriteJSON takes a response `status` code and arbitrary `data`, then generates and sends json in the http response to the client.
```go
func (u *Utils) WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error
```

##### 9. ErrorJSON()
ErrorJSON takes an `error` and optionally an http `status` code, then generates and sends a json formatted error http response. If no `status` code is passed, `http.StatusBadRequest` is the defualt used.
```go
func (u *Utils) ErrorJSON(w http.ResponseWriter, err error, status ...int) error
```

##### 10. PushJSONToRemote()
PushJSONToRemote sends arbitrary `data` to a specified URL as JSON, and returns the `response` and `status` code, or an `error` if any.
The `client` parameter is optional. If none is specified, it uses the standard library's `http.Client`.
```go
func (u *Utils) PushJSONToRemote(uri string, method string, data interface{}, client ...*http.Client) (*http.Response, int, error)
```

---
## Installation

`go get -u github.com/loickreitmann/utils`

---
## Usage
See the example test files, or the package doumentation on the [Go Package Discovery](https://pkg.go.dev/github.com/loickreitmann/utils) site.
