package bnet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ErrorResponse is the error response structure from the Battle.net API
type ErrorResponse struct {
	Response *http.Response

	Code        string `json:"error"`
	Description string `json:"error_description"`
	Scope       string `json:"scope"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%s: %s", r.Code, r.Description)
}

// CheckError checks for an error in the given response.
func CheckError(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	return errorResponse
}
