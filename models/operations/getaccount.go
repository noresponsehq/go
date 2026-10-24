package operations

import (
	"github.com/noresponsehq/go/v6/models/components"
)

type GetAccountRequest struct {
	// The account_id the workspace sends with its assessments, as one path segment (percent-encode '|' and ':' if the client requires it). A value that does not match the account_id shape answers 400 invalid_request.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
}

func (g *GetAccountRequest) GetAccountID() string {
	if g == nil {
		return ""
	}
	return g.AccountID
}

type GetAccountResponse struct {
	Headers map[string][]string
	Result  components.Account
}

func (g *GetAccountResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetAccountResponse) GetResult() components.Account {
	if g == nil {
		return components.Account{}
	}
	return g.Result
}
