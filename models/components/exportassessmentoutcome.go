package components

// ExportAssessmentOutcome - Structured fields of a retained per-assessment outcome. Legacy notes remain private as promised by the existing outcome contract.
type ExportAssessmentOutcome struct {
	// Opaque, server-assigned id of one recorded outcome: 'out_' and a 26-character Crockford base32 ULID. Unique within the workspace.
	ID string `json:"id"`
	// Opaque, server-assigned assessment id: 'asm_' and a 26-character Crockford base32 ULID. Unique within the workspace; another workspace's assessment id answers 404.
	AssessmentID string `json:"assessment_id"`
	// What the provider later learned about the assessed account. confirmed_abuse, account_takeover and chargeback report abuse; false_positive and legitimate withdraw an earlier abuse report of this workspace for the assessment; unknown is recorded only. Outcomes are append-only and the latest one recorded for an assessment wins.
	Outcome    Outcome `json:"outcome"`
	ObservedAt string  `json:"observed_at"`
	ReceivedAt string  `json:"received_at"`
	// Your own opaque reference for the outcome, such as a case or ticket id. Stored as given and not interpreted; keep direct identifiers out of it.
	ExternalRef *string `json:"external_ref,omitzero"`
}

func (e *ExportAssessmentOutcome) GetID() string {
	if e == nil {
		return ""
	}
	return e.ID
}

func (e *ExportAssessmentOutcome) GetAssessmentID() string {
	if e == nil {
		return ""
	}
	return e.AssessmentID
}

func (e *ExportAssessmentOutcome) GetOutcome() Outcome {
	if e == nil {
		return Outcome("")
	}
	return e.Outcome
}

func (e *ExportAssessmentOutcome) GetObservedAt() string {
	if e == nil {
		return ""
	}
	return e.ObservedAt
}

func (e *ExportAssessmentOutcome) GetReceivedAt() string {
	if e == nil {
		return ""
	}
	return e.ReceivedAt
}

func (e *ExportAssessmentOutcome) GetExternalRef() *string {
	if e == nil {
		return nil
	}
	return e.ExternalRef
}
