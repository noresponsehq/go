package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type AggregateRequestsResponse struct {
	Headers map[string][]string
	Result  components.OperationalAggregate
}

func (a *AggregateRequestsResponse) GetHeaders() map[string][]string {
	if a == nil {
		return map[string][]string{}
	}
	return a.Headers
}

func (a *AggregateRequestsResponse) GetResult() components.OperationalAggregate {
	if a == nil {
		return components.OperationalAggregate{}
	}
	return a.Result
}
