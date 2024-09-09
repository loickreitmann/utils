package utils

import (
	"fmt"
	"net/http"
)

// ErrorJSON takes an error and optionally an http status code, then generates and
// sends a json formatted error http response. If no status code is passed,
// http.StatusBadRequest is the defualt used.
func (u *Utils) ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	var errorPayload = JSONResponse{
		Error: JSONError{
			Code:    fmt.Sprint(statusCode),
			Message: err.Error(),
		},
		Data: nil,
	}
	return u.WriteJSON(w, statusCode, errorPayload)
}
