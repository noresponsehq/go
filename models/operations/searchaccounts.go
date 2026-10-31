package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type SearchAccountsResponse struct {
	Headers map[string][]string
	Result  components.AccountDirectoryPage
}

func (s *SearchAccountsResponse) GetHeaders() map[string][]string {
	if s == nil {
		return map[string][]string{}
	}
	return s.Headers
}

func (s *SearchAccountsResponse) GetResult() components.AccountDirectoryPage {
	if s == nil {
		return components.AccountDirectoryPage{}
	}
	return s.Result
}
