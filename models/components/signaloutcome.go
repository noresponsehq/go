package components

// SignalOutcome - Per-signal outcome. A source that timed out, was blocked or errored reports blocked, unavailable or inconclusive with a reason and is listed in incomplete_sources; not_found is a definite absence.
type SignalOutcome string

const (
	SignalOutcomeFound        SignalOutcome = "found"
	SignalOutcomeNotFound     SignalOutcome = "not_found"
	SignalOutcomeInconclusive SignalOutcome = "inconclusive"
	SignalOutcomeBlocked      SignalOutcome = "blocked"
	SignalOutcomeUnavailable  SignalOutcome = "unavailable"
)

func (e SignalOutcome) ToPointer() *SignalOutcome {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *SignalOutcome) IsExact() bool {
	if e != nil {
		switch *e {
		case "found", "not_found", "inconclusive", "blocked", "unavailable":
			return true
		}
	}
	return false
}
