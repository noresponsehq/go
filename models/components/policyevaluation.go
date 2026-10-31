package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

type PolicyEvaluation struct {
	ID            string                   `json:"id"`
	Environment   AuditSettingsEnvironment `json:"environment"`
	Status        PolicyEvaluationStatus   `json:"status"`
	PolicyVersion int64                    `json:"policy_version"`
	ScoringPolicy string                   `json:"scoring_policy"`
	CreatedAt     string                   `json:"created_at"`
	ExpiresAt     string                   `json:"expires_at"`
	CapturedAt    *string                  `json:"captured_at,omitzero"`
	FinishedAt    *string                  `json:"finished_at,omitzero"`
	Accounts      int64                    `json:"accounts"`
	Requests      int64                    `json:"requests"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	evaluatorVersion       string        `const:"identity-threshold-evaluator-v2" json:"evaluator_version"`
	EvaluatorSha256        string        `json:"evaluator_sha256"`
	BuildVersion           *string       `json:"build_version,omitzero"`
	ResultStatus           *ResultStatus `json:"result_status,omitzero"`
	ErrorCode              *string       `json:"error_code,omitzero"`
	PublishedPolicyVersion *int64        `json:"published_policy_version,omitzero"`
	ManifestSha256         *string       `json:"manifest_sha256,omitzero"`
	InputSha256            *string       `json:"input_sha256,omitzero"`
	CohortSha256           *string       `json:"cohort_sha256,omitzero"`
	ResultSha256           *string       `json:"result_sha256,omitzero"`
}

func (p PolicyEvaluation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PolicyEvaluation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PolicyEvaluation) GetID() string {
	if p == nil {
		return ""
	}
	return p.ID
}

func (p *PolicyEvaluation) GetEnvironment() AuditSettingsEnvironment {
	if p == nil {
		return AuditSettingsEnvironment("")
	}
	return p.Environment
}

func (p *PolicyEvaluation) GetStatus() PolicyEvaluationStatus {
	if p == nil {
		return PolicyEvaluationStatus("")
	}
	return p.Status
}

func (p *PolicyEvaluation) GetPolicyVersion() int64 {
	if p == nil {
		return 0
	}
	return p.PolicyVersion
}

func (p *PolicyEvaluation) GetScoringPolicy() string {
	if p == nil {
		return ""
	}
	return p.ScoringPolicy
}

func (p *PolicyEvaluation) GetCreatedAt() string {
	if p == nil {
		return ""
	}
	return p.CreatedAt
}

func (p *PolicyEvaluation) GetExpiresAt() string {
	if p == nil {
		return ""
	}
	return p.ExpiresAt
}

func (p *PolicyEvaluation) GetCapturedAt() *string {
	if p == nil {
		return nil
	}
	return p.CapturedAt
}

func (p *PolicyEvaluation) GetFinishedAt() *string {
	if p == nil {
		return nil
	}
	return p.FinishedAt
}

func (p *PolicyEvaluation) GetAccounts() int64 {
	if p == nil {
		return 0
	}
	return p.Accounts
}

func (p *PolicyEvaluation) GetRequests() int64 {
	if p == nil {
		return 0
	}
	return p.Requests
}

func (p *PolicyEvaluation) GetEvaluatorVersion() string {
	return "identity-threshold-evaluator-v2"
}

func (p *PolicyEvaluation) GetEvaluatorSha256() string {
	if p == nil {
		return ""
	}
	return p.EvaluatorSha256
}

func (p *PolicyEvaluation) GetBuildVersion() *string {
	if p == nil {
		return nil
	}
	return p.BuildVersion
}

func (p *PolicyEvaluation) GetResultStatus() *ResultStatus {
	if p == nil {
		return nil
	}
	return p.ResultStatus
}

func (p *PolicyEvaluation) GetErrorCode() *string {
	if p == nil {
		return nil
	}
	return p.ErrorCode
}

func (p *PolicyEvaluation) GetPublishedPolicyVersion() *int64 {
	if p == nil {
		return nil
	}
	return p.PublishedPolicyVersion
}

func (p *PolicyEvaluation) GetManifestSha256() *string {
	if p == nil {
		return nil
	}
	return p.ManifestSha256
}

func (p *PolicyEvaluation) GetInputSha256() *string {
	if p == nil {
		return nil
	}
	return p.InputSha256
}

func (p *PolicyEvaluation) GetCohortSha256() *string {
	if p == nil {
		return nil
	}
	return p.CohortSha256
}

func (p *PolicyEvaluation) GetResultSha256() *string {
	if p == nil {
		return nil
	}
	return p.ResultSha256
}
