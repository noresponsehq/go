package operations

import (
	"github.com/noresponsehq/go/v7/internal/utils"
	"github.com/noresponsehq/go/v7/models/components"
)

type ListAccountsRequest struct {
	// next_cursor from an earlier page, unchanged. Absent, the list starts from the oldest change.
	ChangedSince *string `queryParam:"style=form,explode=true,name=changed_since"`
	// Most accounts per page.
	Limit *int64 `default:"50" queryParam:"style=form,explode=true,name=limit"`
	// true lists only accounts with a decision other than allow for some flow and no applied enforcement at their current risk_version.
	Unenforced *bool `default:"false" queryParam:"style=form,explode=true,name=unenforced"`
}

func (l ListAccountsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListAccountsRequest) GetChangedSince() *string {
	if l == nil {
		return nil
	}
	return l.ChangedSince
}

func (l *ListAccountsRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListAccountsRequest) GetUnenforced() *bool {
	if l == nil {
		return nil
	}
	return l.Unenforced
}

type ListAccountsResponse struct {
	Headers map[string][]string
	Result  components.AccountList
}

func (l *ListAccountsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListAccountsResponse) GetResult() components.AccountList {
	if l == nil {
		return components.AccountList{}
	}
	return l.Result
}
