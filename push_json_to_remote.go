package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// PushJSONToRemote sends arbitrary data to a specified URL as JSON, and returns
// the response and status code, or an error if any.
// The client parameter is optional. If none is specified, it uses the standard
// library's http.Client.
func (u *Utils) PushJSONToRemote(uri string, method string, data interface{}, client ...*http.Client) (*http.Response, int, error) {
	// create JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, 0, err
	}

	// check for custom http client
	httpClient := &http.Client{}
	if len(client) > 0 {
		httpClient = client[0]
	}

	// build the request and set the header
	request, err := http.NewRequest(method, uri, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, 0, err
	}
	request.Header.Set("Content-Type", "application/json")

	// call the remote api
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, 0, err
	}
	defer response.Body.Close()

	// send the repsonse pack

	return response, response.StatusCode, nil
}
