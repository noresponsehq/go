package operations

import (
	"github.com/noresponsehq/go/v8/models/components"
)

type RecordEnforcementRequest struct {
	// The account_id the workspace sends with its assessments, as one path segment (percent-encode '|' and ':' if the client requires it). A value that does not match the account_id shape answers 400 invalid_request.
	AccountID string                        `pathParam:"style=simple,explode=false,name=account_id"`
	Body      components.EnforcementRequest `request:"mediaType=application/json"`
}

func (r *RecordEnforcementRequest) GetAccountID() string {
	if r == nil {
		return ""
	}
	return r.AccountID
}

func (r *RecordEnforcementRequest) GetBody() components.EnforcementRequest {
	if r == nil {
		return components.EnforcementRequest{}
	}
	return r.Body
}

type RecordEnforcementResponse struct {
	Headers map[string][]string
	Result  components.Enforcement
}

func (r *RecordEnforcementResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *RecordEnforcementResponse) GetResult() components.Enforcement {
	if r == nil {
		return components.Enforcement{}
	}
	return r.Result
}
