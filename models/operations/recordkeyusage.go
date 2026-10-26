package operations

import (
	"github.com/noresponsehq/go/v8/models/components"
)

type RecordKeyUsageRequest struct {
	// Opaque client key, 16 to 128 characters of [A-Za-z0-9._:-] (a UUID or ULID fits). Scoped to the workspace principal and bound to the canonical body for 24 hours.
	IdempotencyKey string                          `header:"style=simple,explode=false,name=Idempotency-Key"`
	Body           components.KeyUsageBatchRequest `request:"mediaType=application/json"`
}

func (r *RecordKeyUsageRequest) GetIdempotencyKey() string {
	if r == nil {
		return ""
	}
	return r.IdempotencyKey
}

func (r *RecordKeyUsageRequest) GetBody() components.KeyUsageBatchRequest {
	if r == nil {
		return components.KeyUsageBatchRequest{}
	}
	return r.Body
}

type RecordKeyUsageResponse struct {
	Headers map[string][]string
	Result  components.KeyUsageBatch
}

func (r *RecordKeyUsageResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *RecordKeyUsageResponse) GetResult() components.KeyUsageBatch {
	if r == nil {
		return components.KeyUsageBatch{}
	}
	return r.Result
}
