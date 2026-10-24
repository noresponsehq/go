package operations

import (
	"github.com/noresponsehq/go/v7/models/components"
)

type DeleteAccountRequest struct {
	// The account_id the workspace sends with its assessments, as one path segment (percent-encode '|' and ':' if the client requires it). A value that does not match the account_id shape answers 400 invalid_request.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
}

func (d *DeleteAccountRequest) GetAccountID() string {
	if d == nil {
		return ""
	}
	return d.AccountID
}

type DeleteAccountResponse struct {
	Headers map[string][]string
	Result  components.AccountDeletion
}

func (d *DeleteAccountResponse) GetHeaders() map[string][]string {
	if d == nil {
		return map[string][]string{}
	}
	return d.Headers
}

func (d *DeleteAccountResponse) GetResult() components.AccountDeletion {
	if d == nil {
		return components.AccountDeletion{}
	}
	return d.Result
}
