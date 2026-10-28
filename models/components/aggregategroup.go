package components

// AggregateGroup - One group's count in one bucket. It holds each grouped dimension the counted assessments had a value for: a pending first answer has no decision, recommendation or band; an answer that fired no reason, listed no incomplete lane, matched no list entry or whose IP has no known country or ASN leaves that dimension out.
type AggregateGroup struct {
	// An assessment's own decision: block when its stored risk.block is true, otherwise its recommendation. It is not the account's decision or held-admission state.
	Decision *Decision `json:"decision,omitzero"`
	// The event an assessment is made for: account_creation (a user creates an account), login (a user returns to an existing account) or purchase (a payment or a credit grant). The account's risk is reported with the decision for this flow.
	Flow *Flow `json:"flow,omitzero"`
	// The recommended action: allow, step_up or review. Your application applies it.
	Recommendation *Recommendation `json:"recommendation,omitzero"`
	// low 0-29, medium 30-59, high 60-84, critical 85-100.
	Band *Band `json:"band,omitzero"`
	// The action of the list entry that decided the account.
	ListMatch *ListMatch `json:"list_match,omitzero"`
	// A stable reason code. New codes may appear; ignore any you do not know.
	Reason *ReasonCode `json:"reason,omitzero"`
	Impact *Impact     `json:"impact,omitzero"`
	// An evidence lane: accounts (account presence for email and username), phone, breach, domain (the email domain's registration) and ip.
	Lane    *Lane   `json:"lane,omitzero"`
	Country *string `json:"country,omitzero"`
	Asn     *int64  `json:"asn,omitzero"`
	Count   int64   `json:"count"`
}

func (a *AggregateGroup) GetDecision() *Decision {
	if a == nil {
		return nil
	}
	return a.Decision
}

func (a *AggregateGroup) GetFlow() *Flow {
	if a == nil {
		return nil
	}
	return a.Flow
}

func (a *AggregateGroup) GetRecommendation() *Recommendation {
	if a == nil {
		return nil
	}
	return a.Recommendation
}

func (a *AggregateGroup) GetBand() *Band {
	if a == nil {
		return nil
	}
	return a.Band
}

func (a *AggregateGroup) GetListMatch() *ListMatch {
	if a == nil {
		return nil
	}
	return a.ListMatch
}

func (a *AggregateGroup) GetReason() *ReasonCode {
	if a == nil {
		return nil
	}
	return a.Reason
}

func (a *AggregateGroup) GetImpact() *Impact {
	if a == nil {
		return nil
	}
	return a.Impact
}

func (a *AggregateGroup) GetLane() *Lane {
	if a == nil {
		return nil
	}
	return a.Lane
}

func (a *AggregateGroup) GetCountry() *string {
	if a == nil {
		return nil
	}
	return a.Country
}

func (a *AggregateGroup) GetAsn() *int64 {
	if a == nil {
		return nil
	}
	return a.Asn
}

func (a *AggregateGroup) GetCount() int64 {
	if a == nil {
		return 0
	}
	return a.Count
}
