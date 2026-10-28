package components

// AggregateScoreOutcome - First answers of one flow whose risk.score is in [score_min, score_max] and whose assessment's latest reported outcome (the latest received) is outcome. Counted when the first answer is counted and moved when a later outcome wins; outcomes reported before the first answer was counted are counted with it. An outcome reported before this count existed is counted only once another outcome is reported for the assessment.
type AggregateScoreOutcome struct {
	// The event an assessment is made for: account_creation (a user creates an account), login (a user returns to an existing account) or purchase (a payment or a credit grant). The account's risk is reported with the decision for this flow.
	Flow Flow `json:"flow"`
	// A multiple of 5.
	ScoreMin int64 `json:"score_min"`
	ScoreMax int64 `json:"score_max"`
	// What the provider later learned about the assessed account. confirmed_abuse, account_takeover and chargeback report abuse; false_positive and legitimate withdraw an earlier abuse report of this workspace for the assessment; unknown is recorded only. Outcomes are append-only and the latest one recorded for an assessment wins.
	Outcome Outcome `json:"outcome"`
	Count   int64   `json:"count"`
}

func (a *AggregateScoreOutcome) GetFlow() Flow {
	if a == nil {
		return Flow("")
	}
	return a.Flow
}

func (a *AggregateScoreOutcome) GetScoreMin() int64 {
	if a == nil {
		return 0
	}
	return a.ScoreMin
}

func (a *AggregateScoreOutcome) GetScoreMax() int64 {
	if a == nil {
		return 0
	}
	return a.ScoreMax
}

func (a *AggregateScoreOutcome) GetOutcome() Outcome {
	if a == nil {
		return Outcome("")
	}
	return a.Outcome
}

func (a *AggregateScoreOutcome) GetCount() int64 {
	if a == nil {
		return 0
	}
	return a.Count
}
