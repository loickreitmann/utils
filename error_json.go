package utils

import (
	"fmt"
	"net/http"
)

// ErrorJSON takes an error and optionally an http status code, then generates and
// sends a json formatted error http response. If no status code is passed,
// http.StatusBadRequest is the defualt used.
func (u *Utils) ErrorJSON(w http.ResponseWriter, err error, httpStatus ...int) error {
	statusCode := http.StatusBadRequest
	if len(httpStatus) > 0 {
		statusCode = httpStatus[0]
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
