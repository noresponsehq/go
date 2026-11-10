package components

type ScreeningSpendLineRole string

const (
	ScreeningSpendLineRoleResolver  ScreeningSpendLineRole = "resolver"
	ScreeningSpendLineRoleCollector ScreeningSpendLineRole = "collector"
	ScreeningSpendLineRoleAnalyser  ScreeningSpendLineRole = "analyser"
)

func (e ScreeningSpendLineRole) ToPointer() *ScreeningSpendLineRole {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ScreeningSpendLineRole) IsExact() bool {
	if e != nil {
		switch *e {
		case "resolver", "collector", "analyser":
			return true
		}
	}
	return false
}
