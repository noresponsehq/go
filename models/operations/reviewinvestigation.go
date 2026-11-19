package operations

import (
	"github.com/noresponsehq/go/v16/models/components"
)

type ReviewInvestigationRequest struct {
	InvestigationID string                         `pathParam:"style=simple,explode=false,name=investigation_id"`
	Body            components.PersonReviewRequest `request:"mediaType=application/json"`
}

func (r *ReviewInvestigationRequest) GetInvestigationID() string {
	if r == nil {
		return ""
	}
	return r.InvestigationID
}

func (r *ReviewInvestigationRequest) GetBody() components.PersonReviewRequest {
	if r == nil {
		return components.PersonReviewRequest{}
	}
	return r.Body
}

type ReviewInvestigationResponse struct {
	Headers map[string][]string
	Result  components.Investigation
}

func (r *ReviewInvestigationResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *ReviewInvestigationResponse) GetResult() components.Investigation {
	if r == nil {
		return components.Investigation{}
	}
	return r.Result
}
