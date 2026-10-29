package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type RetireAccountCredentialRequest struct {
	// The account_id the workspace sends with its assessments, as one path segment (percent-encode '|' and ':' if the client requires it). A value that does not match the account_id shape answers 400 invalid_request.
	AccountID    string `pathParam:"style=simple,explode=false,name=account_id"`
	CredentialID string `pathParam:"style=simple,explode=false,name=credential_id"`
}

func (r *RetireAccountCredentialRequest) GetAccountID() string {
	if r == nil {
		return ""
	}
	return r.AccountID
}

func (r *RetireAccountCredentialRequest) GetCredentialID() string {
	if r == nil {
		return ""
	}
	return r.CredentialID
}

type RetireAccountCredentialResponse struct {
	Headers map[string][]string
	Result  components.CredentialAssociation
}

func (r *RetireAccountCredentialResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *RetireAccountCredentialResponse) GetResult() components.CredentialAssociation {
	if r == nil {
		return components.CredentialAssociation{}
	}
	return r.Result
}
