package components

// AssessmentRequestPolicyContext - Optional policy inputs: region.
type AssessmentRequestPolicyContext struct {
	// ISO 3166-1 alpha-2 region of the signup, as the provider determines it.
	Region *string `json:"region,omitzero"`
}

func (a *AssessmentRequestPolicyContext) GetRegion() *string {
	if a == nil {
		return nil
	}
	return a.Region
}
