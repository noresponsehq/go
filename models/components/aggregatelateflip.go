package components

// AggregateLateFlip - Changes of the account decision for an assessment's flow after its first answer, as account.risk_changed announces them, counted in the bucket of the assessment's creation. from is the decision announced before (allow before the first announcement); an assessment that changes twice counts twice.
type AggregateLateFlip struct {
	// block when the score reached your block threshold (risk.block of the assessment), otherwise the assessment's recommendation. held only for a flow whose admission is held (settings flows.<flow>.admission), when the account's current assessment was made for that flow and its investigation has not completed (status partial): the final decision arrives as account.risk_changed when the investigation completes, decided by its score, or when the hold's deadline (hold_seconds after the assessment was created) passes first, as the flow's on_expiry action or the decision the evidence so far gives, whichever is stricter. Your application applies it.
	From AccountDecision `json:"from"`
	// block when the score reached your block threshold (risk.block of the assessment), otherwise the assessment's recommendation. held only for a flow whose admission is held (settings flows.<flow>.admission), when the account's current assessment was made for that flow and its investigation has not completed (status partial): the final decision arrives as account.risk_changed when the investigation completes, decided by its score, or when the hold's deadline (hold_seconds after the assessment was created) passes first, as the flow's on_expiry action or the decision the evidence so far gives, whichever is stricter. Your application applies it.
	To AccountDecision `json:"to"`
	// How long after the assessment's creation the change was made.
	Delay Delay `json:"delay"`
	Count int64 `json:"count"`
}

func (a *AggregateLateFlip) GetFrom() AccountDecision {
	if a == nil {
		return AccountDecision("")
	}
	return a.From
}

func (a *AggregateLateFlip) GetTo() AccountDecision {
	if a == nil {
		return AccountDecision("")
	}
	return a.To
}

func (a *AggregateLateFlip) GetDelay() Delay {
	if a == nil {
		return Delay("")
	}
	return a.Delay
}

func (a *AggregateLateFlip) GetCount() int64 {
	if a == nil {
		return 0
	}
	return a.Count
}
