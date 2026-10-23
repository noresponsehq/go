package components

// Flow - The event an assessment is made for: account_creation (a user creates an account), login (a user returns to an existing account) or purchase (a payment or a credit grant). The account's risk is reported with the decision for this flow.
type Flow string

const (
	FlowAccountCreation Flow = "account_creation"
	FlowLogin           Flow = "login"
	FlowPurchase        Flow = "purchase"
)

func (e Flow) ToPointer() *Flow {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Flow) IsExact() bool {
	if e != nil {
		switch *e {
		case "account_creation", "login", "purchase":
			return true
		}
	}
	return false
}
