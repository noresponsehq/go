package components

// Outcome - What the provider later learned about the assessed account. confirmed_abuse, account_takeover and chargeback report abuse; false_positive and legitimate withdraw an earlier abuse report of this tenant for the assessment; unknown is recorded only. Outcomes are append-only and the latest one recorded for an assessment wins.
type Outcome string

const (
	OutcomeConfirmedAbuse  Outcome = "confirmed_abuse"
	OutcomeAccountTakeover Outcome = "account_takeover"
	OutcomeChargeback      Outcome = "chargeback"
	OutcomeFalsePositive   Outcome = "false_positive"
	OutcomeLegitimate      Outcome = "legitimate"
	OutcomeUnknown         Outcome = "unknown"
)

func (e Outcome) ToPointer() *Outcome {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Outcome) IsExact() bool {
	if e != nil {
		switch *e {
		case "confirmed_abuse", "account_takeover", "chargeback", "false_positive", "legitimate", "unknown":
			return true
		}
	}
	return false
}
