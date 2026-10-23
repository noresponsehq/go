package operations

import (
	"github.com/noresponsehq/go/v4/models/components"
)

type RotateAPIKeyResponse struct {
	Headers map[string][]string
	Result  components.KeyRotation
}

func (r *RotateAPIKeyResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *RotateAPIKeyResponse) GetResult() components.KeyRotation {
	if r == nil {
		return components.KeyRotation{}
	}
	return r.Result
}
