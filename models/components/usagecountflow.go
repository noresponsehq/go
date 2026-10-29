package components

// UsageCountFlow - assessments only.
type UsageCountFlow string

const (
	UsageCountFlowAccountCreation UsageCountFlow = "account_creation"
	UsageCountFlowLogin           UsageCountFlow = "login"
	UsageCountFlowPurchase        UsageCountFlow = "purchase"
)

func (e UsageCountFlow) ToPointer() *UsageCountFlow {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *UsageCountFlow) IsExact() bool {
	if e != nil {
		switch *e {
		case "account_creation", "login", "purchase":
			return true
		}
	}
	return false
}
