package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// RiskDetails - Detailed view only (response_mode=detailed, which needs the detailed scope); a standard answer leaves it out. Merged observations without provider, source, origin, raw breach record, value or score.
type RiskDetails struct {
	PhoneNames []RiskDetailsPhoneName `json:"phone_names,omitzero"`
	// The matching breach tables, newest first, at most 100.
	Breaches []Breach `json:"breaches,omitzero"`
	// Names found next to the identifiers in breach records, merged, with the number of distinct tables each was seen in.
	BreachNames []RiskDetailsBreachName `json:"breach_names,omitzero"`
	// The identity's dated footprint events, newest first, at most 300: the year and month, what happened (created, last_active, breached), the site, its category, the kind of evidence and the identifier it belongs to. No breach value: dates are coarsened to the month.
	Events []RiskDetailsEvent `json:"events,omitzero"`
}

func (r RiskDetails) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RiskDetails) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (r *RiskDetails) GetPhoneNames() []RiskDetailsPhoneName {
	if r == nil {
		return nil
	}
	return r.PhoneNames
}

func (r *RiskDetails) GetBreaches() []Breach {
	if r == nil {
		return nil
	}
	return r.Breaches
}

func (r *RiskDetails) GetBreachNames() []RiskDetailsBreachName {
	if r == nil {
		return nil
	}
	return r.BreachNames
}

func (r *RiskDetails) GetEvents() []RiskDetailsEvent {
	if r == nil {
		return nil
	}
	return r.Events
}
