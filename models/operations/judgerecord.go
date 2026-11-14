package operations

import (
	"github.com/noresponsehq/go/v12/models/components"
)

type JudgeRecordRequest struct {
	// Your workspace id, as each key lists it in workspace_id. Another workspace, a malformed id and a test environment's id answer 404.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
	// The environment: live or test. Each has its own settings. Any other value answers 400. An API key reads and changes only its own environment's; the other answers 404.
	Environment components.KeyEnvironment       `pathParam:"style=simple,explode=false,name=environment"`
	GroupID     string                          `pathParam:"style=simple,explode=false,name=group_id"`
	PersonID    string                          `pathParam:"style=simple,explode=false,name=person_id"`
	Body        components.RecordVerdictRequest `request:"mediaType=application/json"`
}

func (j *JudgeRecordRequest) GetWorkspaceID() string {
	if j == nil {
		return ""
	}
	return j.WorkspaceID
}

func (j *JudgeRecordRequest) GetEnvironment() components.KeyEnvironment {
	if j == nil {
		return components.KeyEnvironment("")
	}
	return j.Environment
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
