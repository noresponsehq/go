package operations

import (
	"github.com/noresponsehq/go/v9/internal/utils"
	"github.com/noresponsehq/go/v9/models/components"
)

type ListAccountDecisionsRequest struct {
	// The account_id the workspace sends with its assessments, as one path segment (percent-encode '|' and ':' if the client requires it). A value that does not match the account_id shape answers 400 invalid_request.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// Items per page.
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// next_cursor from the previous page, unchanged.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListAccountDecisionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountDecisionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListAccountDecisionsRequest) GetAccountID() string {
	if l == nil {
		return ""
	}
	return l.AccountID
}

func (l *ListAccountDecisionsRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListAccountDecisionsRequest) GetCursor() *string {
	if l == nil {
		return nil
	}
	return l.Cursor
}

type ListAccountDecisionsResponse struct {
	Headers map[string][]string
	Result  components.AccountDecisionList

	Next func() (*ListAccountDecisionsResponse, error)
}

func (l *ListAccountDecisionsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListAccountDecisionsResponse) GetResult() components.AccountDecisionList {
	if l == nil {
		return components.AccountDecisionList{}
	}
	return l.Result
}
