package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

type RiskDetailsBreachName struct {
	Value  string `json:"value"`
	SeenIn int64  `json:"seen_in"`
	// The domains of the breached sites whose tables hold the name, lower case, each once, newest breach first, at most 32: the domain values of breaches[]. Covers the email's and the username's tables together. Absent when none of the tables has a site.
	Sites []string `json:"sites,omitzero"`
	// How many of the tables that hold the name have no site, such as a credential list or a stealer log. Absent when none.
	Unsited *int64 `json:"unsited,omitzero"`
}

func (r RiskDetailsBreachName) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RiskDetailsBreachName) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (r *RiskDetailsBreachName) GetValue() string {
	if r == nil {
		return ""
	}
	return r.Value
}

func (r *RiskDetailsBreachName) GetSeenIn() int64 {
	if r == nil {
		return 0
	}
	return r.SeenIn
}

func (r *RiskDetailsBreachName) GetSites() []string {
	if r == nil {
		return nil
	}
	return r.Sites
}

func (r *RiskDetailsBreachName) GetUnsited() *int64 {
	if r == nil {
		return nil
	}
	return r.Unsited
}
