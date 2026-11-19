package operations

import (
	"github.com/noresponsehq/go/v17/models/components"
)

type JudgeWorkspacePersonRecordRequest struct {
	PersonID string                          `pathParam:"style=simple,explode=false,name=person_id"`
	Body     components.RecordVerdictRequest `request:"mediaType=application/json"`
}

func (j *JudgeWorkspacePersonRecordRequest) GetPersonID() string {
	if j == nil {
		return ""
	}
	return j.PersonID
}

func (j *JudgeWorkspacePersonRecordRequest) GetBody() components.RecordVerdictRequest {
	if j == nil {
		return components.RecordVerdictRequest{}
	}
	return j.Body
}

type JudgeWorkspacePersonRecordResponse struct {
	Headers map[string][]string
	Result  components.RecordVerdict
}

func (j *JudgeWorkspacePersonRecordResponse) GetHeaders() map[string][]string {
	if j == nil {
		return map[string][]string{}
	}
	return j.Headers
}

func (j *JudgeWorkspacePersonRecordResponse) GetResult() components.RecordVerdict {
	if j == nil {
		return components.RecordVerdict{}
	}
	return j.Result
}
