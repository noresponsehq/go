package operations

import (
	"github.com/noresponsehq/go/v9/models/components"
)

type SearchAssessmentsResponse struct {
	Headers map[string][]string
	Result  components.AssessmentSearchPage
}

func (s *SearchAssessmentsResponse) GetHeaders() map[string][]string {
	if s == nil {
		return map[string][]string{}
	}
	return s.Headers
}

func (s *SearchAssessmentsResponse) GetResult() components.AssessmentSearchPage {
	if s == nil {
		return components.AssessmentSearchPage{}
	}
	return s.Result
}
