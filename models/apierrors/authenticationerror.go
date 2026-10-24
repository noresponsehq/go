package apierrors

import (
	"encoding/json"
	"github.com/noresponsehq/go/v7/models/components"
)

// AuthenticationError - HTTP 401 problem document.
type AuthenticationError struct {
	Type     string  `json:"type"`
	Title    string  `json:"title"`
	Status   int64   `const:"401" json:"status"`
	Detail   string  `json:"detail"`
	Instance *string `json:"instance,omitzero"`
	// Machine codes an HTTP 401 answer can carry.
	Code          components.AuthenticationErrorCode `json:"code"`
	RequestID     string                             `json:"request_id"`
	InvalidParams []components.ProblemInvalidParam   `json:"invalid_params,omitzero"`
}

var _ error = &AuthenticationError{}

func (e *AuthenticationError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
