package components

import (
	"github.com/noresponsehq/go/v2/internal/utils"
)

// RiskSummary - A few typed facts per identifier. Every fact is optional: absent means the evidence did not answer and makes no claim, while false or 0 is a definite answer. Lanes that do not exist yet leave their members absent.
type RiskSummary struct {
	Phone    *RiskSummaryPhone    `json:"phone,omitzero"`
	Email    *RiskSummaryEmail    `json:"email,omitzero"`
	Username *RiskSummaryUsername `json:"username,omitzero"`
	// Facts about the submitted IP from the offline database and this tenant's own recent assessments. Every member is optional; absent makes no claim.
	IP *IPSummary `json:"ip,omitzero"`
	// Cross-identifier agreement, computed from facts only. Every member is optional; absent makes no claim.
	Consistency *Consistency `json:"consistency,omitzero"`
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
