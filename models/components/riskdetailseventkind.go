package components

type RiskDetailsEventKind string

const (
	RiskDetailsEventKindCreated    RiskDetailsEventKind = "created"
	RiskDetailsEventKindLastActive RiskDetailsEventKind = "last_active"
	RiskDetailsEventKindBreached   RiskDetailsEventKind = "breached"
)

func (e RiskDetailsEventKind) ToPointer() *RiskDetailsEventKind {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *RiskDetailsEventKind) IsExact() bool {
	if e != nil {
		switch *e {
		case "created", "last_active", "breached":
			return true
		}
	}
	return false
}
