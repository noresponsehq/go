package operations

import (
	"github.com/noresponsehq/go/v8/internal/utils"
	"github.com/noresponsehq/go/v8/models/components"
)

type ListAccountOutcomesRequest struct {
	// The account_id the workspace sends with its assessments, as one path segment (percent-encode '|' and ':' if the client requires it). A value that does not match the account_id shape answers 400 invalid_request.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// Items per page.
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// next_cursor from the previous page, unchanged.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListAccountOutcomesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountOutcomesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListAccountOutcomesRequest) GetAccountID() string {
	if l == nil {
		return ""
	}
	return l.AccountID
}

func (l *ListAccountOutcomesRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListAccountOutcomesRequest) GetCursor() *string {
	if l == nil {
		return nil
	}
	return l.Cursor
}

type ListAccountOutcomesResponse struct {
	Headers map[string][]string
	Result  components.AccountOutcomePage
}

func (l *ListAccountOutcomesResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListAccountOutcomesResponse) GetResult() components.AccountOutcomePage {
	if l == nil {
		return components.AccountOutcomePage{}
	}
	return l.Result
}
