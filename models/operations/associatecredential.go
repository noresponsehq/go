package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type AssociateCredentialRequest struct {
	// The account_id the workspace sends with its assessments, as one path segment (percent-encode '|' and ':' if the client requires it). A value that does not match the account_id shape answers 400 invalid_request.
	AccountID string                       `pathParam:"style=simple,explode=false,name=account_id"`
	Body      components.CredentialRequest `request:"mediaType=application/json"`
}

func (a *AssociateCredentialRequest) GetAccountID() string {
	if a == nil {
		return ""
	}
	return a.AccountID
}

func (a *AssociateCredentialRequest) GetBody() components.CredentialRequest {
	if a == nil {
		return components.CredentialRequest{}
	}
	return a.Body
}

type AssociateCredentialResponse struct {
	Headers map[string][]string
	Result  components.CredentialAssociation
}

func (a *AssociateCredentialResponse) GetHeaders() map[string][]string {
	if a == nil {
		return map[string][]string{}
	}
	return a.Headers
}

func (a *AssociateCredentialResponse) GetResult() components.CredentialAssociation {
	if a == nil {
		return components.CredentialAssociation{}
	}
	return a.Result
}
