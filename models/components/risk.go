package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// Risk - The one risk decision for this event. score is 0 (safe) to 100 (risky), set by the reasons that fired under a hand-set, versioned heuristic policy. Each reason carries its direction and impact band; the policy's weights are not published. Every answer with a recommendation carries it, and the top-level recommendation equals risk.recommendation. It names no source or provider.
type Risk struct {
	Score int64 `json:"score"`
	Band  Band  `json:"band"`
	// The recommended action: allow, step_up or review. Your application applies it.
	Recommendation Recommendation `json:"recommendation"`
	// true when the workspace's block threshold is set and score reached it. A provider hook can deny the account creation on it; the provider decides and enforces. It does not change recommendation. It is false while blocking is off, the default, and a body from before the flag existed omits it: read absent as false.
	Block *bool `json:"block,omitzero"`
	// The block threshold this answer used: block is true exactly when score is at least this value. Absent while blocking is off. NoResponse sets it per workspace on request, at or above the review threshold.
	BlockThreshold *int64 `json:"block_threshold,omitzero"`
	// How much of the planned evidence answered; it says nothing about risk.
	Confidence RiskConfidence `json:"confidence"`
	// Version of the hand-authored heuristic policy that produced the recommendation, for example 'risk-2026-10.7'. A workspace threshold override adds a +t<N> suffix, e.g. 'risk-2026-10.7+t3'.
	PolicyVersion string `json:"policy_version"`
	// Labels the v0 policy as hand-authored and not statistically validated.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	policyBasis string `const:"heuristic_unvalidated" json:"policy_basis"`
	// Every reason that fired: risk-raising reasons first, then neutral, then risk-lowering, each strongest impact band first and then by code. The policy's numeric weights are kept private, so a reason carries only its direction and impact band.
	Reasons []RiskReason `json:"reasons"`
	// A few typed facts per identifier. Every fact is optional: absent means the evidence did not answer and makes no claim, while false or 0 is a definite answer. Lanes that do not exist yet leave their members absent.
	Summary *RiskSummary `json:"summary,omitzero"`
	// Counts of planned source-identifier pairs by evidence outcome, deduplicated across signals and incomplete entries. These counts do not measure identity agreement or prove that processing is still running. A lane that was not searched is not an answered source.
	Coverage RiskCoverage `json:"coverage"`
	// Detailed view only (response_mode=detailed, which needs the detailed scope); a standard answer leaves it out. Merged observations without provider, source, origin, raw breach record, value or score.
	Details *RiskDetails `json:"details,omitzero"`
}

func (r Risk) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *Risk) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (r *Risk) GetScore() int64 {
	if r == nil {
		return 0
	}
	return r.Score
}

func (r *Risk) GetBand() Band {
	if r == nil {
		return Band("")
	}
	return r.Band
}

func (r *Risk) GetRecommendation() Recommendation {
	if r == nil {
		return Recommendation("")
	}
	return r.Recommendation
}

func (r *Risk) GetBlock() *bool {
	if r == nil {
		return nil
	}
	return r.Block
}

func (r *Risk) GetBlockThreshold() *int64 {
	if r == nil {
		return nil
	}
	return r.BlockThreshold
}

func (r *Risk) GetConfidence() RiskConfidence {
	if r == nil {
		return RiskConfidence("")
	}
	return r.Confidence
}

func (r *Risk) GetPolicyVersion() string {
	if r == nil {
		return ""
	}
	return r.PolicyVersion
}

func (r *Risk) GetPolicyBasis() string {
	return "heuristic_unvalidated"
}

func (r *Risk) GetReasons() []RiskReason {
	if r == nil {
		return []RiskReason{}
	}
	return r.Reasons
}

func (r *Risk) GetSummary() *RiskSummary {
	if r == nil {
		return nil
	}
	return r.Summary
}

func (r *Risk) GetCoverage() RiskCoverage {
	if r == nil {
		return RiskCoverage{}
	}
	return r.Coverage
}

func (r *Risk) GetDetails() *RiskDetails {
	if r == nil {
		return nil
	}
	return r.Details
}
