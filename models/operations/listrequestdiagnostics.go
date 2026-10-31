package operations

import (
	"github.com/noresponsehq/go/v10/internal/utils"
	"github.com/noresponsehq/go/v10/models/components"
)

type ListRequestDiagnosticsRequest struct {
	// Inclusive start of a range of at most seven days.
	From string `queryParam:"style=form,explode=true,name=from"`
	// Exclusive end of the range.
	To string `queryParam:"style=form,explode=true,name=to"`
	// Optional filter or paging parameter.
	RequestID *string `queryParam:"style=form,explode=true,name=request_id"`
	// Optional filter or paging parameter.
	Route *string `queryParam:"style=form,explode=true,name=route"`
	// Optional filter or paging parameter.
	Status *int64 `queryParam:"style=form,explode=true,name=status"`
	// Optional filter or paging parameter.
	Limit *int64 `default:"50" queryParam:"style=form,explode=true,name=limit"`
	// Optional filter or paging parameter.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Filter by the public authenticated API key identifier; older unattributed rows do not match.
	APIKeyID *string `queryParam:"style=form,explode=true,name=api_key_id"`
}

func (l ListRequestDiagnosticsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListRequestDiagnosticsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListRequestDiagnosticsRequest) GetFrom() string {
	if l == nil {
		return ""
	}
	return l.From
}

func (l *ListRequestDiagnosticsRequest) GetTo() string {
	if l == nil {
		return ""
	}
	return l.To
}

func (l *ListRequestDiagnosticsRequest) GetRequestID() *string {
	if l == nil {
		return nil
	}
	return l.RequestID
}

func (l *ListRequestDiagnosticsRequest) GetRoute() *string {
	if l == nil {
		return nil
	}
	return l.Route
}

func (l *ListRequestDiagnosticsRequest) GetStatus() *int64 {
	if l == nil {
		return nil
	}
	return l.Status
}

func (l *ListRequestDiagnosticsRequest) GetLimit() *int64 {
	if l == nil {
		return nil
	}
	return l.Limit
}

func (l *ListRequestDiagnosticsRequest) GetCursor() *string {
	if l == nil {
		return nil
	}
	return l.Cursor
}

func (l *ListRequestDiagnosticsRequest) GetAPIKeyID() *string {
	if l == nil {
		return nil
	}
	return l.APIKeyID
}

type ListRequestDiagnosticsResponse struct {
	Headers map[string][]string
	Result  components.RequestDiagnosticPage

	Next func() (*ListRequestDiagnosticsResponse, error)
}

func (l *ListRequestDiagnosticsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListRequestDiagnosticsResponse) GetResult() components.RequestDiagnosticPage {
	if l == nil {
		return components.RequestDiagnosticPage{}
	}
	return l.Result
}
