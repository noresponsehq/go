package apierrors

import (
	"fmt"
	"net/http"
)

type APIStatusError struct {
	Message     string
	StatusCode  int
	Body        string
	RawResponse *http.Response
}

var _ error = &APIStatusError{}

func NewAPIStatusError(message string, statusCode int, body string, httpRes *http.Response) *APIStatusError {
	return &APIStatusError{
		Message:     message,
		StatusCode:  statusCode,
		Body:        body,
		RawResponse: httpRes,
	}
}

func (e *APIStatusError) Error() string {
	body := ""
	if len(e.Body) > 0 {
		body = fmt.Sprintf("\n%s", e.Body)
	}

	return fmt.Sprintf("%s: Status %d%s", e.Message, e.StatusCode, body)
}
