package components

type CustomerExportRequestFilterOutcome string

const (
	CustomerExportRequestFilterOutcomeSucceeded CustomerExportRequestFilterOutcome = "succeeded"
	CustomerExportRequestFilterOutcomeDenied    CustomerExportRequestFilterOutcome = "denied"
)

func (e CustomerExportRequestFilterOutcome) ToPointer() *CustomerExportRequestFilterOutcome {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *CustomerExportRequestFilterOutcome) IsExact() bool {
	if e != nil {
		switch *e {
		case "succeeded", "denied":
			return true
		}
	}
	return false
}
