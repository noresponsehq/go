package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// IPSummary - Facts about the submitted IP from the offline database and this workspace's own recent assessments. Every member is optional; absent makes no claim.
type IPSummary struct {
	// ISO 3166-1 alpha-2 country of the address.
	Country *string `json:"country,omitzero"`
	// The address's network type.
	Type *IPSummaryType `json:"type,omitzero"`
	// Autonomous system number of the network that announces the address, from the offline database. Absent when unknown.
	Asn *int64 `json:"asn,omitzero"`
	// Assessments in this workspace with the same IP in the last 24 hours, this one included. A count only.
	SeenInWorkspace *int64 `json:"seen_in_workspace,omitzero"`
}

func (i IPSummary) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IPSummary) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (i *IPSummary) GetCountry() *string {
	if i == nil {
		return nil
	}
	return i.Country
}

func (i *IPSummary) GetType() *IPSummaryType {
	if i == nil {
		return nil
	}
	return i.Type
}

func (i *IPSummary) GetAsn() *int64 {
	if i == nil {
		return nil
	}
	return i.Asn
}

func (i *IPSummary) GetSeenInWorkspace() *int64 {
	if i == nil {
		return nil
	}
	return i.SeenInWorkspace
}
