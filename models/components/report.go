package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

type ReportScope struct {
}

type Report struct {
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	evaluatorVersion string       `const:"identity-threshold-evaluator-v2" json:"evaluator_version"`
	InputSha256      string       `json:"input_sha256"`
	CohortSha256     string       `json:"cohort_sha256"`
	Status           ResultStatus `json:"status"`
	Candidate        Thresholds   `json:"candidate"`
	Scope            ReportScope  `json:"scope"`
}

func (r Report) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *Report) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (r *Report) GetEvaluatorVersion() string {
	return "identity-threshold-evaluator-v2"
}

func (r *Report) GetInputSha256() string {
	if r == nil {
		return ""
	}
	return r.InputSha256
}

func (r *Report) GetCohortSha256() string {
	if r == nil {
		return ""
	}
	return r.CohortSha256
}

func (r *Report) GetStatus() ResultStatus {
	if r == nil {
		return ResultStatus("")
	}
	return r.Status
}

func (r *Report) GetCandidate() Thresholds {
	if r == nil {
		return Thresholds{}
	}
	return r.Candidate
}

func (r *Report) GetScope() ReportScope {
	if r == nil {
		return ReportScope{}
	}
	return r.Scope
}
