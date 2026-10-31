package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// Publication - Immutable publication provenance remains in policy history after evaluation evidence expires. SHA-256 values identify the frozen manifest, evaluator semantic input, cohort, and canonical result separately.
type Publication struct {
	ManifestSha256  string `json:"manifest_sha256"`
	InputSha256     string `json:"input_sha256"`
	CohortSha256    string `json:"cohort_sha256"`
	ResultSha256    string `json:"result_sha256"`
	EvaluatorSha256 string `json:"evaluator_sha256"`
	EvaluationID    string `json:"evaluation_id"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	evaluatorVersion string `const:"identity-threshold-evaluator-v2" json:"evaluator_version"`
	BuildVersion     string `json:"build_version"`
	CapturedAt       string `json:"captured_at"`
	PublishedAt      string `json:"published_at"`
}

func (p Publication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Publication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *Publication) GetManifestSha256() string {
	if p == nil {
		return ""
	}
	return p.ManifestSha256
}

func (p *Publication) GetInputSha256() string {
	if p == nil {
		return ""
	}
	return p.InputSha256
}

func (p *Publication) GetCohortSha256() string {
	if p == nil {
		return ""
	}
	return p.CohortSha256
}

func (p *Publication) GetResultSha256() string {
	if p == nil {
		return ""
	}
	return p.ResultSha256
}

func (p *Publication) GetEvaluatorSha256() string {
	if p == nil {
		return ""
	}
	return p.EvaluatorSha256
}

func (p *Publication) GetEvaluationID() string {
	if p == nil {
		return ""
	}
	return p.EvaluationID
}

func (p *Publication) GetEvaluatorVersion() string {
	return "identity-threshold-evaluator-v2"
}

func (p *Publication) GetBuildVersion() string {
	if p == nil {
		return ""
	}
	return p.BuildVersion
}

func (p *Publication) GetCapturedAt() string {
	if p == nil {
		return ""
	}
	return p.CapturedAt
}

func (p *Publication) GetPublishedAt() string {
	if p == nil {
		return ""
	}
	return p.PublishedAt
}
