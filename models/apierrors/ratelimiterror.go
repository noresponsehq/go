package apierrors

import (
	"encoding/json"
	"github.com/noresponsehq/go/v3/models/components"
)

// RateLimitError - HTTP 429 problem document.
type RateLimitError struct {
	Type     string  `json:"type"`
	Title    string  `json:"title"`
	Status   int64   `const:"429" json:"status"`
	Detail   string  `json:"detail"`
	Instance *string `json:"instance,omitzero"`
	// Machine codes an HTTP 429 answer can carry.
	Code          components.RateLimitErrorCode    `json:"code"`
	RequestID     string                           `json:"request_id"`
	InvalidParams []components.ProblemInvalidParam `json:"invalid_params,omitzero"`
}

var _ error = &RateLimitError{}

func (e *RateLimitError) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
