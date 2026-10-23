package apierrors

import (
	"encoding/json"
	"github.com/noresponsehq/go/v4/models/components"
)

// PermissionDeniedError - HTTP 403 problem document.
type PermissionDeniedError struct {
	Type     string  `json:"type"`
	Title    string  `json:"title"`
	Status   int64   `const:"403" json:"status"`
	Detail   string  `json:"detail"`
	Instance *string `json:"instance,omitzero"`
	// Machine codes an HTTP 403 answer can carry.
	Code          components.PermissionDeniedErrorCode `json:"code"`
	RequestID     string                               `json:"request_id"`
	InvalidParams []components.ProblemInvalidParam     `json:"invalid_params,omitzero"`
}

var _ error = &PermissionDeniedError{}

func (e *PermissionDeniedError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
