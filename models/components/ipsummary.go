package components

// IPSummary - Facts about the submitted IP from the offline database and this tenant's own recent assessments. Every member is optional; absent makes no claim.
type IPSummary struct {
	// ISO 3166-1 alpha-2 country of the address.
	Country *string `json:"country,omitzero"`
	// The address's network type.
	Type *IPSummaryType `json:"type,omitzero"`
	// Assessments in this tenant with the same IP in the last 24 hours, this one included. A count only.
	SeenInTenant *int64 `json:"seen_in_tenant,omitzero"`
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

func (i *IPSummary) GetSeenInTenant() *int64 {
	if i == nil {
		return nil
	}
	return i.SeenInTenant
}
