package apierrors

import (
	"encoding/json"
	"github.com/noresponsehq/go/v10/models/components"
)

// BadRequestError - HTTP 400 problem document.
type BadRequestError struct {
	Type     string  `json:"type"`
	Title    string  `json:"title"`
	Status   int64   `const:"400" json:"status"`
	Detail   string  `json:"detail"`
	Instance *string `json:"instance,omitzero"`
	// Machine codes an HTTP 400 answer can carry.
	Code          components.BadRequestErrorCode   `json:"code"`
	RequestID     string                           `json:"request_id"`
	InvalidParams []components.ProblemInvalidParam `json:"invalid_params,omitzero"`
}

var _ error = &BadRequestError{}

func (e *BadRequestError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
