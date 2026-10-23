package operations

import (
	"github.com/noresponsehq/go/v3/models/components"
)

type QueryAccountAssessmentsResponse struct {
	Headers map[string][]string
	Result  components.AssessmentList

	Next func() (*QueryAccountAssessmentsResponse, error)
}

func (q *QueryAccountAssessmentsResponse) GetHeaders() map[string][]string {
	if q == nil {
		return map[string][]string{}
	}
	return q.Headers
}

func (q *QueryAccountAssessmentsResponse) GetResult() components.AssessmentList {
	if q == nil {
		return components.AssessmentList{}
	}
	return q.Result
}
