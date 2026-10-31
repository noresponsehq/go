package operations

import (
	"github.com/noresponsehq/go/v10/models/components"
)

type GetAssessmentVersionRequest struct {
	// Path segment, matched case-sensitively. A value that does not match the assessment_id shape answers 404, not 400.
	AssessmentID string `pathParam:"style=simple,explode=false,name=assessment_id"`
	// Optional standard or separately authorised detailed customer view; raw is rejected.
	ResponseMode *components.ResponseMode `queryParam:"style=form,explode=true,name=response_mode"`
	// The stored assessment_version, in canonical decimal form. An unavailable or malformed version answers 404.
	Version int64 `pathParam:"style=simple,explode=false,name=version"`
}

func (g *GetAssessmentVersionRequest) GetAssessmentID() string {
	if g == nil {
		return ""
	}
	return g.AssessmentID
}

func (g *GetAssessmentVersionRequest) GetResponseMode() *components.ResponseMode {
	if g == nil {
		return nil
	}
	return g.ResponseMode
}

func (g *GetAssessmentVersionRequest) GetVersion() int64 {
	if g == nil {
		return 0
	}
	return g.Version
}

type GetAssessmentVersionResponse struct {
	Headers map[string][]string
	Result  components.AssessmentRevision
}

func (g *GetAssessmentVersionResponse) GetHeaders() map[string][]string {
	if g == nil {
		return map[string][]string{}
	}
	return g.Headers
}

func (g *GetAssessmentVersionResponse) GetResult() components.AssessmentRevision {
	if g == nil {
		return components.AssessmentRevision{}
	}
	return g.Result
}
