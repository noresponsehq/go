package components

// AggregateScoreCount - First answers of one flow whose risk.score is in [score_min, score_max]: 0-4, 5-9 up to 95-100. A pending answer has no score.
type AggregateScoreCount struct {
	// The event an assessment is made for: account_creation (a user creates an account), login (a user returns to an existing account) or purchase (a payment or a credit grant). The account's risk is reported with the decision for this flow.
	Flow Flow `json:"flow"`
	// A multiple of 5.
	ScoreMin int64 `json:"score_min"`
	ScoreMax int64 `json:"score_max"`
	Count    int64 `json:"count"`
}

func (a *AggregateScoreCount) GetFlow() Flow {
	if a == nil {
		return Flow("")
	}
	return a.Flow
}

func (a *AggregateScoreCount) GetScoreMin() int64 {
	if a == nil {
		return 0
	}
	return a.ScoreMin
}

func (a *AggregateScoreCount) GetScoreMax() int64 {
	if a == nil {
		return 0
	}
	return a.ScoreMax
}

func (a *AggregateScoreCount) GetCount() int64 {
	if a == nil {
		return 0
	}
	return a.Count
}
