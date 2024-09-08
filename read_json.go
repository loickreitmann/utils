package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
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
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		// identify the relevant error
		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly formed json at character %d", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains unterminated json")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect json type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect json type at character %d", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case strings.HasPrefix(err.Error(), "json: unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field")
			return fmt.Errorf("body contains unknown key %s", fieldName)
		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", u.MaxJSONReadSize)
		case errors.As(err, &invalidUnmarshalError):
			return fmt.Errorf("error unmashalling json: %s", err.Error())
		default:
			return err
		}
	}

	// check if body contains more than one JSON structure
	err = jsonDecoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("json body must only contain a single object structure")
	}

	return nil
}
