package operations

import (
	"github.com/noresponsehq/go/v8/models/components"
)

type ExportAccountRequest struct {
	// The account_id the workspace sends with its assessments, as one path segment (percent-encode '|' and ':' if the client requires it). A value that does not match the account_id shape answers 400 invalid_request.
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
	// standard returns the consolidated risk response; detailed adds merged provider-neutral facts when the key has detailed scope. There is no raw mode.
	ResponseMode *components.ResponseMode `queryParam:"style=form,explode=true,name=response_mode"`
}

func (e *ExportAccountRequest) GetAccountID() string {
	if e == nil {
		return ""
	}
	return e.AccountID
}

func (e *ExportAccountRequest) GetResponseMode() *components.ResponseMode {
	if e == nil {
		return nil
	}
	return e.ResponseMode
}

type ExportAccountResponse struct {
	Headers map[string][]string
	Result  components.AccountExport
}

func (e *ExportAccountResponse) GetHeaders() map[string][]string {
	if e == nil {
		return map[string][]string{}
	}
	return e.Headers
}

func (e *ExportAccountResponse) GetResult() components.AccountExport {
	if e == nil {
		return components.AccountExport{}
	}
	return e.Result
}
