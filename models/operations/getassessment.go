package operations

import (
	"github.com/noresponsehq/go/v7/models/components"
)

type GetAssessmentRequest struct {
	// Path segment, matched case-sensitively. A value that does not match the assessment_id shape answers 404, not 400.
	AssessmentID string `pathParam:"style=simple,explode=false,name=assessment_id"`
	// Optional standard or separately authorised detailed customer view; raw is rejected.
	ResponseMode *components.ResponseMode `queryParam:"style=form,explode=true,name=response_mode"`
}

func (g *GetAssessmentRequest) GetAssessmentID() string {
	if g == nil {
		return ""
	}
	return g.AssessmentID
}

func (g *GetAssessmentRequest) GetResponseMode() *components.ResponseMode {
	if g == nil {
		return nil
	}
	return g.ResponseMode
}

type GetAssessmentResponse struct {
	Headers map[string][]string
	Result  components.Assessment
}

func (g *GetAssessmentResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetAssessmentResponse) GetResult() components.Assessment {
	if g == nil {
		return components.Assessment{}
	}
	return g.Result
}
