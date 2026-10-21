package operations

import (
	"github.com/noresponsehq/go/v2/models/components"
)

type RecordAssessmentOutcomeRequest struct {
	// Path segment, matched case-sensitively. A value that does not match the assessment_id shape answers 404, not 400.
	AssessmentID string `pathParam:"style=simple,explode=false,name=assessment_id"`
	// Opaque client key, 16 to 128 characters of [A-Za-z0-9._:-] (a UUID or ULID fits). Scoped to the tenant principal and bound to the canonical body for 24 hours.
	IdempotencyKey string                              `header:"style=simple,explode=false,name=Idempotency-Key"`
	Body           components.AssessmentOutcomeRequest `request:"mediaType=application/json"`
}

func (r *RecordAssessmentOutcomeRequest) GetAssessmentID() string {
	if r == nil {
		return ""
	}
	return r.AssessmentID
}

func (r *RecordAssessmentOutcomeRequest) GetIdempotencyKey() string {
	if r == nil {
		return ""
	}
	return r.IdempotencyKey
}

func (r *RecordAssessmentOutcomeRequest) GetBody() components.AssessmentOutcomeRequest {
	if r == nil {
		return components.AssessmentOutcomeRequest{}
	}
	return r.Body
}

type RecordAssessmentOutcomeResponse struct {
	Headers map[string][]string
	Result  components.AssessmentOutcome
}

func (r *RecordAssessmentOutcomeResponse) GetHeaders() map[string][]string {
	if r == nil {
		return map[string][]string{}
	}
	return r.Headers
}

func (r *RecordAssessmentOutcomeResponse) GetResult() components.AssessmentOutcome {
	if r == nil {
		return components.AssessmentOutcome{}
	}
	return r.Result
}
