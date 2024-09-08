package utils

type JSONOptions struct {
	MaxJSONReadSize    int
	AllowUnknownFields bool
}

type JSONError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// JSONResponse is the type used for sending JSON
type JSONResponse struct {
	Error JSONError   `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}
