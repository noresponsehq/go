package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type AggregateAssessmentsResponse struct {
	Headers map[string][]string
	Result  components.AssessmentAggregate
}

func (a *AggregateAssessmentsResponse) GetHeaders() map[string][]string {
	if a == nil {
		return map[string][]string{}
	}
	return a.Headers
}

func (a *AggregateAssessmentsResponse) GetResult() components.AssessmentAggregate {
	if a == nil {
		return components.AssessmentAggregate{}
	}
	return a.Result
}
