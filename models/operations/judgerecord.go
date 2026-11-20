package operations

import (
	"github.com/noresponsehq/go/v20/models/components"
)

type JudgeRecordRequest struct {
	GroupID  string                          `pathParam:"style=simple,explode=false,name=group_id"`
	PersonID string                          `pathParam:"style=simple,explode=false,name=person_id"`
	Body     components.RecordVerdictRequest `request:"mediaType=application/json"`
}

func (j *JudgeRecordRequest) GetGroupID() string {
	if j == nil {
		return ""
	}
	return j.GroupID
}

func (j *JudgeRecordRequest) GetPersonID() string {
	if j == nil {
		return ""
	}
	return j.PersonID
}

func (j *JudgeRecordRequest) GetBody() components.RecordVerdictRequest {
	if j == nil {
		return components.RecordVerdictRequest{}
	}
	return j.Body
}

type JudgeRecordResponse struct {
	Headers map[string][]string
	Result  components.RecordVerdict
}

func (j *JudgeRecordResponse) GetHeaders() map[string][]string {
	if j == nil {
		return map[string][]string{}
	}
	return j.Headers
}

func (j *JudgeRecordResponse) GetResult() components.RecordVerdict {
	if j == nil {
		return components.RecordVerdict{}
	}
	return j.Result
}
