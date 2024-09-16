package utils_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/loickreitmann/utils"
)

var errorCases = []struct {
	name string
	code int
	err  error
}{
	{name: "forbidden error", code: http.StatusForbidden, err: errors.New("you shall not pass")},
	{name: "not found error", code: http.StatusNotFound, err: errors.New("nothing to see here move along")},
	{name: "service unavailable", code: http.StatusServiceUnavailable, err: errors.New("fuggehdaboutdit")},
}

func TestUtils_ErrorJSON(t *testing.T) {
	var testUtils utils.Utils
	for _, errorCase := range errorCases {
		// ARRANGE
		rr := httptest.NewRecorder()

		// ACT
		err := testUtils.ErrorJSON(rr, errorCase.err, errorCase.code)

		// ASSERT
		if err != nil {
			t.Errorf("[%s]: unexpected error: %v", errorCase.name, err.Error())
		}
		var errorPayload utils.JSONResponse
		decoder := json.NewDecoder(rr.Body)
		err = decoder.Decode(&errorPayload)
		if err != nil {
			t.Errorf("[%s]: unable to decode json: %v", errorCase.name, err.Error())
		}
		if errorPayload.Data != nil {
			t.Errorf("[%s]: unexpected data: %v", errorCase.name, errorPayload.Data)
		}
		if rr.Code != errorCase.code {
			t.Errorf("[%s]: wrong status code; expected %d, got %d", errorCase.name, errorCase.code, rr.Code)
		}
	}
}
