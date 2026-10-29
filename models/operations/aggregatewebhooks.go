package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type AggregateWebhooksResponse struct {
	Headers map[string][]string
	Result  components.OperationalAggregate
}

func (a *AggregateWebhooksResponse) GetHeaders() map[string][]string {
	if a == nil {
		return map[string][]string{}
	}
	return a.Headers
}

func (a *AggregateWebhooksResponse) GetResult() components.OperationalAggregate {
	if a == nil {
		return components.OperationalAggregate{}
	}
	return a.Result
}
