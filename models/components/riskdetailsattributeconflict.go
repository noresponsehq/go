package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

type RiskDetailsAttributeConflict struct {
	// name_gender is the records' gender against the submitted first name's usual one.
	Attribute Attribute `json:"attribute"`
	// The domains of the breached sites whose tables report the attribute, newest breach first, at most 32. Absent when none of the tables has a site.
	Sites []string `json:"sites,omitzero"`
	// How many of the tables reporting the attribute have no site, such as a credential list or a stealer log. Absent when none.
	Unsited *int64 `json:"unsited,omitzero"`
}

func (r RiskDetailsAttributeConflict) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RiskDetailsAttributeConflict) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (r *RiskDetailsAttributeConflict) GetAttribute() Attribute {
	if r == nil {
		return Attribute("")
	}
	return r.Attribute
}

func (r *RiskDetailsAttributeConflict) GetSites() []string {
	if r == nil {
		return nil
	}
	return r.Sites
}

func (r *RiskDetailsAttributeConflict) GetUnsited() *int64 {
	if r == nil {
		return nil
	}
	return r.Unsited
}
