package components

// AssessmentOutcome - The 202 answer of POST /v1/assessments/{assessment_id}/outcomes. A replay with the same Idempotency-Key and body answers the same record.
type AssessmentOutcome struct {
	// Opaque, server-assigned id of one recorded outcome: 'out_' and a 26-character Crockford base32 ULID. Unique within the tenant.
	ID string `json:"id"`
	// Opaque, server-assigned assessment id: 'asm_' and a 26-character Crockford base32 ULID. Unique within the tenant; another tenant's assessment id answers 404.
	AssessmentID string `json:"assessment_id"`
	// What the provider later learned about the assessed account. confirmed_abuse, account_takeover and chargeback report abuse; false_positive and legitimate withdraw an earlier abuse report of this tenant for the assessment; unknown is recorded only. Outcomes are append-only and the latest one recorded for an assessment wins.
	Outcome Outcome `json:"outcome"`
	// When NoResponse first recorded the outcome.
	ReceivedAt string `json:"received_at"`
}

func (a *AssessmentOutcome) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *AssessmentOutcome) GetAssessmentID() string {
	if a == nil {
		return ""
	}
	return a.AssessmentID
}

func (a *AssessmentOutcome) GetOutcome() Outcome {
	if a == nil {
		return Outcome("")
	}
	return a.Outcome
}

func (a *AssessmentOutcome) GetReceivedAt() string {
	if a == nil {
		return ""
	}
	return a.ReceivedAt
}
