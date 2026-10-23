package operations

import (
	"github.com/noresponsehq/go/v5/models/components"
)

type ListIntegrationsResponse struct {
	Headers map[string][]string
	Result  components.IntegrationList
}

func (l *ListIntegrationsResponse) GetHeaders() map[string][]string {
	if l == nil {
		return map[string][]string{}
	}
	return l.Headers
}

func (l *ListIntegrationsResponse) GetResult() components.IntegrationList {
	if l == nil {
		return components.IntegrationList{}
	}
	return l.Result
}
