package operations

import (
	"github.com/noresponsehq/go/v2/models/components"
)

type DeleteAssessmentRequest struct {
	// Path segment, matched case-sensitively. A value that does not match the assessment_id shape answers 404, not 400.
	AssessmentID string `pathParam:"style=simple,explode=false,name=assessment_id"`
}

func (d *DeleteAssessmentRequest) GetAssessmentID() string {
	if d == nil {
		return ""
	}
	return d.AssessmentID
}

type DeleteAssessmentResponse struct {
	Headers map[string][]string
	Result  components.AssessmentDeletion
}

func (d *DeleteAssessmentResponse) GetHeaders() map[string][]string {
	if d == nil {
		return map[string][]string{}
	}
	return d.Headers
}

func (d *DeleteAssessmentResponse) GetResult() components.AssessmentDeletion {
	if d == nil {
		return components.AssessmentDeletion{}
	}
	return d.Result
}
