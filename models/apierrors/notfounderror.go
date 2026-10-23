package apierrors

import (
	"encoding/json"
	"github.com/noresponsehq/go/v3/models/components"
)

// NotFoundError - HTTP 404 problem document.
type NotFoundError struct {
	Type     string  `json:"type"`
	Title    string  `json:"title"`
	Status   int64   `const:"404" json:"status"`
	Detail   string  `json:"detail"`
	Instance *string `json:"instance,omitzero"`
	// Machine codes an HTTP 404 answer can carry.
	Code          components.NotFoundErrorCode     `json:"code"`
	RequestID     string                           `json:"request_id"`
	InvalidParams []components.ProblemInvalidParam `json:"invalid_params,omitzero"`
}

var _ error = &NotFoundError{}

func (e *NotFoundError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
