package components

import (
	"github.com/noresponsehq/go/v7/internal/utils"
)

// RiskSummary - A few typed facts per identifier. Every fact is optional: absent means the evidence did not answer and makes no claim, while false or 0 is a definite answer. Lanes that do not exist yet leave their members absent.
type RiskSummary struct {
	Phone    *RiskSummaryPhone    `json:"phone,omitzero"`
	Email    *RiskSummaryEmail    `json:"email,omitzero"`
	Username *RiskSummaryUsername `json:"username,omitzero"`
	// Facts about the submitted IP from the offline database and this workspace's own recent assessments. Every member is optional; absent makes no claim.
	IP *IPSummary `json:"ip,omitzero"`
	// Cross-identifier agreement, computed from facts only. Every member is optional; absent makes no claim.
	Consistency *Consistency `json:"consistency,omitzero"`
	// How long and how widely the signup's email and username have existed, and how far independent sources agree on the person behind them, in years and counts only. families are the kinds of evidence that found either identifier. Omitted when nothing found them: a footprint lookup that failed makes no claim, and absence is no_digital_footprint's to state.
	Footprint *Footprint `json:"footprint,omitzero"`
}

func (r RiskSummary) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RiskSummary) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (r *RiskSummary) GetPhone() *RiskSummaryPhone {
	if r == nil {
		return nil
	}
	return r.Phone
}

func (r *RiskSummary) GetEmail() *RiskSummaryEmail {
	if r == nil {
		return nil
	}
	return r.Email
}

func (r *RiskSummary) GetUsername() *RiskSummaryUsername {
	if r == nil {
		return nil
	}
	return r.Username
}

func (r *RiskSummary) GetIP() *IPSummary {
	if r == nil {
		return nil
	}
	return r.IP
}

func (r *RiskSummary) GetConsistency() *Consistency {
	if r == nil {
		return nil
	}
	return r.Consistency
}

func (r *RiskSummary) GetFootprint() *Footprint {
	if r == nil {
		return nil
	}
	return r.Footprint
}
