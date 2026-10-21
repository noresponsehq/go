package components

// AssessmentOutcomeRequest - POST /v1/assessments/{assessment_id}/outcomes body: what the provider later learned about the account it assessed. The assessment is named by the path and the tenant by the API key, so neither is a body member. Idempotency-Key is required and binds the body as it does on POST /v1/assessments.
type AssessmentOutcomeRequest struct {
	// What the provider later learned about the assessed account. confirmed_abuse, account_takeover and chargeback report abuse; false_positive and legitimate withdraw an earlier abuse report of this tenant for the assessment; unknown is recorded only. Outcomes are append-only and the latest one recorded for an assessment wins.
	Outcome Outcome `json:"outcome"`
	// When the provider observed the outcome. Defaults to when NoResponse received it.
	ObservedAt *string `json:"observed_at,omitzero"`
	// Your own opaque reference for the outcome, such as a case or ticket id. Stored as given and not interpreted; keep direct identifiers out of it.
	ExternalRef *string `json:"external_ref,omitzero"`
	// Free text for your own records, at most 500 characters. Stored only; it is left out of scoring and answers. Keep personal data out of it. Control characters other than tab and line breaks are refused.
	Notes *string `json:"notes,omitzero"`
}

func (a *AssessmentOutcomeRequest) GetOutcome() Outcome {
	if a == nil {
		return Outcome("")
	}
	return a.Outcome
}

func (a *AssessmentOutcomeRequest) GetObservedAt() *string {
	if a == nil {
		return nil
	}
	return a.ObservedAt
}

func (a *AssessmentOutcomeRequest) GetExternalRef() *string {
	if a == nil {
		return nil
	}
	return a.ExternalRef
}

func (a *AssessmentOutcomeRequest) GetNotes() *string {
	if a == nil {
		return nil
	}
	return a.Notes
}
