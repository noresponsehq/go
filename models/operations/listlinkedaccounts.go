package operations

import (
	"github.com/noresponsehq/go/v9/internal/utils"
	"github.com/noresponsehq/go/v9/models/components"
)

type ListLinkedAccountsRequest struct {
	// The account_id the workspace sends with its assessments, as one path segment (percent-encode '|' and ':' if the client requires it). A value that does not match the account_id shape answers 400 invalid_request.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// Linked accounts per page.
	Limit *int64 `default:"25" queryParam:"style=form,explode=true,name=limit"`
	// next_cursor from the previous page, unchanged.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListLinkedAccountsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListLinkedAccountsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListLinkedAccountsRequest) GetAccountID() string {
	if l == nil {
		return ""
	}
	return l.AccountID
}

func (l *ListLinkedAccountsRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListLinkedAccountsRequest) GetCursor() *string {
	if l == nil {
		return nil
	}
	return l.Cursor
}

type ListLinkedAccountsResponse struct {
	Headers map[string][]string
	Result  components.LinkedAccountPage

	Next func() (*ListLinkedAccountsResponse, error)
}

func (l *ListLinkedAccountsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListLinkedAccountsResponse) GetResult() components.LinkedAccountPage {
	if l == nil {
		return components.LinkedAccountPage{}
	}
	return l.Result
}
