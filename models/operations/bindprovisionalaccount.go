package operations

import (
	"github.com/noresponsehq/go/v8/models/components"
)

type BindProvisionalAccountRequest struct {
	// Path segment, matched case-sensitively. A value that does not match the assessment_id shape answers 404, not 400.
	AssessmentID string                           `pathParam:"style=simple,explode=false,name=assessment_id"`
	Body         components.AccountBindingRequest `request:"mediaType=application/json"`
}

func (b *BindProvisionalAccountRequest) GetAssessmentID() string {
	if b == nil {
		return ""
	}
	return b.AssessmentID
}

func (b *BindProvisionalAccountRequest) GetBody() components.AccountBindingRequest {
	if b == nil {
		return components.AccountBindingRequest{}
	}
	return b.Body
}

type BindProvisionalAccountResponse struct {
	Headers map[string][]string
	Result  components.AccountBinding
}

func (b *BindProvisionalAccountResponse) GetHeaders() map[string][]string {
	if b == nil {
		return map[string][]string{}
	}
	return b.Headers
}

func (b *BindProvisionalAccountResponse) GetResult() components.AccountBinding {
	if b == nil {
		return components.AccountBinding{}
	}
	return b.Result
}
