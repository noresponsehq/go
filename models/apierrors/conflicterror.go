package apierrors

import (
	"encoding/json"
	"github.com/noresponsehq/go/v8/models/components"
)

// ConflictError - HTTP 409 problem document.
type ConflictError struct {
	Type     string  `json:"type"`
	Title    string  `json:"title"`
	Status   int64   `const:"409" json:"status"`
	Detail   string  `json:"detail"`
	Instance *string `json:"instance,omitzero"`
	// Machine codes an HTTP 409 answer can carry.
	Code          components.ConflictErrorCode     `json:"code"`
	RequestID     string                           `json:"request_id"`
	InvalidParams []components.ProblemInvalidParam `json:"invalid_params,omitzero"`
}

var _ error = &ConflictError{}

func (e *ConflictError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
