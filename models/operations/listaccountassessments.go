package operations

import (
	"github.com/noresponsehq/go/v7/internal/utils"
	"github.com/noresponsehq/go/v7/models/components"
)

type ListAccountAssessmentsRequest struct {
	// The account_id the workspace sends with its assessments, as one path segment (percent-encode '|' and ':' if the client requires it). A value that does not match the account_id shape answers 400 invalid_request.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// Items per page.
	Limit *int64 `default:"20" queryParam:"style=form,explode=true,name=limit"`
	// next_cursor from the previous page, unchanged.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Optional standard or separately authorised detailed customer view; raw is rejected.
	ResponseMode *components.ResponseMode `queryParam:"style=form,explode=true,name=response_mode"`
}

func (l ListAccountAssessmentsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountAssessmentsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListAccountAssessmentsRequest) GetAccountID() string {
	if l == nil {
		return ""
	}
	return l.AccountID
}

func (l *ListAccountAssessmentsRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListAccountAssessmentsRequest) GetCursor() *string {
	if l == nil {
		return nil
	}
	return l.Cursor
}

func (l *ListAccountAssessmentsRequest) GetResponseMode() *components.ResponseMode {
	if l == nil {
		return nil
	}
	return l.ResponseMode
}

type ListAccountAssessmentsResponse struct {
	Headers map[string][]string
	Result  components.AssessmentList

	Next func() (*ListAccountAssessmentsResponse, error)
}

func (l *ListAccountAssessmentsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListAccountAssessmentsResponse) GetResult() components.AssessmentList {
	if l == nil {
		return components.AssessmentList{}
	}
	return l.Result
}
