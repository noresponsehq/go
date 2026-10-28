package operations

import (
	"github.com/noresponsehq/go/v8/models/components"
)

type RecordAccountOutcomeRequest struct {
	// The account_id the workspace sends with its assessments, as one path segment (percent-encode '|' and ':' if the client requires it). A value that does not match the account_id shape answers 400 invalid_request.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// Opaque client key, 16 to 128 characters of [A-Za-z0-9._:-] (a UUID or ULID fits). Scoped to the workspace principal and bound to the canonical body for 24 hours.
	IdempotencyKey string                           `header:"style=simple,explode=false,name=Idempotency-Key"`
	Body           components.AccountOutcomeRequest `request:"mediaType=application/json"`
}

func (r *RecordAccountOutcomeRequest) GetAccountID() string {
	if r == nil {
		return ""
	}
	return r.AccountID
}

func (r *RecordAccountOutcomeRequest) GetIdempotencyKey() string {
	if r == nil {
		return ""
	}
	return r.IdempotencyKey
}

func (r *RecordAccountOutcomeRequest) GetBody() components.AccountOutcomeRequest {
	if r == nil {
		return components.AccountOutcomeRequest{}
	}
	return r.Body
}

type RecordAccountOutcomeResponse struct {
	Headers map[string][]string
	Result  components.AccountOutcome
}

func (r *RecordAccountOutcomeResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *RecordAccountOutcomeResponse) GetResult() components.AccountOutcome {
	if r == nil {
		return components.AccountOutcome{}
	}
	return r.Result
}
