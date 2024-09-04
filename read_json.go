package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (u *Utils) ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	if u.MaxJSONReadSize == 0 {
		u.MaxJSONReadSize = mB
	}
	// Read the body from the request
	r.Body = http.MaxBytesReader(w, r.Body, int64(u.MaxJSONReadSize))
	jsonDecoder := json.NewDecoder(r.Body)
	if !u.AllowUnknownFields {
		jsonDecoder.DisallowUnknownFields()
	}

	err := jsonDecoder.Decode(data)
	if err != nil {
		return err
	}

	// check if body contains more than one JSON structure
	err = jsonDecoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("json body must only contain a single object structure")
	}

	return nil
}
