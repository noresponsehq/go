package components

// AccountOutcomeRequest - Record a factual report, not an enforcement action. A known outcome replaces the effective report. unknown alone records uncertainty without changing it; unknown with supersedes explicitly retracts it. supersedes must name the current effective report or the request answers 409 outcome_conflict. Reports remain append-only. Idempotency-Key is required. observed_at defaults to receipt time and cannot be over five minutes ahead.
type AccountOutcomeRequest struct {
	// What the provider later learned about the assessed account. confirmed_abuse, account_takeover and chargeback report abuse; false_positive and legitimate withdraw an earlier abuse report of this workspace for the assessment; unknown is recorded only. Outcomes are append-only and the latest one recorded for an assessment wins.
	Outcome    Outcome `json:"outcome"`
	ObservedAt *string `json:"observed_at,omitzero"`
	// Your own opaque reference for the outcome, such as a case or ticket id. Stored as given and not interpreted; keep direct identifiers out of it.
	ExternalRef *string `json:"external_ref,omitzero"`
	// Opaque, server-assigned id of one recorded outcome: 'out_' and a 26-character Crockford base32 ULID. Unique within the workspace.
	Supersedes *string `json:"supersedes,omitzero"`
}

func (a *AccountOutcomeRequest) GetOutcome() Outcome {
	if a == nil {
		return Outcome("")
	}
	return a.Outcome
}

func (a *AccountOutcomeRequest) GetObservedAt() *string {
	if a == nil {
		return nil
	}
	return a.ObservedAt
}

func (a *AccountOutcomeRequest) GetExternalRef() *string {
	if a == nil {
		return nil
	}
	return a.ExternalRef
}

func (a *AccountOutcomeRequest) GetSupersedes() *string {
	if a == nil {
		return nil
	}
	return a.Supersedes
}
