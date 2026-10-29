package components

// OutcomesByOutcome - Current customer-reported labels, not independently established ground truth. Missing keys mean zero. Unknown is an explicit report; absent or expired feedback is unreported.
type OutcomesByOutcome struct {
	ConfirmedAbuse  *int64 `json:"confirmed_abuse,omitzero"`
	AccountTakeover *int64 `json:"account_takeover,omitzero"`
	Chargeback      *int64 `json:"chargeback,omitzero"`
	FalsePositive   *int64 `json:"false_positive,omitzero"`
	Legitimate      *int64 `json:"legitimate,omitzero"`
	Unknown         *int64 `json:"unknown,omitzero"`
}

func (o *OutcomesByOutcome) GetConfirmedAbuse() *int64 {
	if o == nil {
		return nil
	}
	return o.ConfirmedAbuse
}

func (o *OutcomesByOutcome) GetAccountTakeover() *int64 {
	if o == nil {
		return nil
	}
	return o.AccountTakeover
}

func (o *OutcomesByOutcome) GetChargeback() *int64 {
	if o == nil {
		return nil
	}
	return o.Chargeback
}

func (o *OutcomesByOutcome) GetFalsePositive() *int64 {
	if o == nil {
		return nil
	}
	return o.FalsePositive
}

func (o *OutcomesByOutcome) GetLegitimate() *int64 {
	if o == nil {
		return nil
	}
	return o.Legitimate
}

func (o *OutcomesByOutcome) GetUnknown() *int64 {
	if o == nil {
		return nil
	}
	return o.Unknown
}
