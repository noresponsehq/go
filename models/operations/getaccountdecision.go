package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type GetAccountDecisionRequest struct {
	// The account_id the workspace sends with its assessments, as one path segment (percent-encode '|' and ':' if the client requires it). A value that does not match the account_id shape answers 400 invalid_request.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// The canonical decimal risk_version originally served.
	RiskVersion int64 `pathParam:"style=simple,explode=false,name=risk_version"`
}

func (g *GetAccountDecisionRequest) GetAccountID() string {
	if g == nil {
		return ""
	}
	return g.AccountID
}

func (g *GetAccountDecisionRequest) GetRiskVersion() int64 {
	if g == nil {
		return 0
	}
	return g.RiskVersion
}

type GetAccountDecisionResponse struct {
	Headers map[string][]string
	Result  components.AccountDecisionSnapshot
}

func (g *GetAccountDecisionResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetAccountDecisionResponse) GetResult() components.AccountDecisionSnapshot {
	if g == nil {
		return components.AccountDecisionSnapshot{}
	}
	return g.Result
}
