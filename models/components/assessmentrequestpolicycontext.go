package components

type AssessmentRequestPolicyContext struct {
	// The provider flow the assessment gates. Use the top-level flow instead; this member is ignored when the top-level flow is present. Absent means signup.
	Flow *Flow `json:"flow,omitzero"`
	// ISO 3166-1 alpha-2 region of the signup, as the provider determines it.
	Region *string `json:"region,omitzero"`
}

func (a *AssessmentRequestPolicyContext) GetFlow() *Flow {
	if a == nil {
		return nil
	}
	return a.Flow
}

func (a *AssessmentRequestPolicyContext) GetRegion() *string {
	if a == nil {
		return nil
	}
	return a.Region
}
