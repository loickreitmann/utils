package utils

import (
	"encoding/json"
	"net/http"
)

// WriteJSON takes a response, an httpStatus code, and arbitrary data and generates and sends json in the http response to the client
func (u *Utils) WriteJSON(w http.ResponseWriter, httpStatus int, data interface{}, headers ...http.Header) error {
	output, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header().Set(key, value[0])
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_, err = w.Write(output)
	if err != nil {
		return err
	}
	return nil
}
