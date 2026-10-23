package components

// Flow - The customer event an assessment is made for: signup (a new account), login (a returning sign-in) or purchase (a payment or credit grant). The account's risk is reported with the decision for this flow.
type Flow string

const (
	FlowSignup   Flow = "signup"
	FlowLogin    Flow = "login"
	FlowPurchase Flow = "purchase"
)

func (e Flow) ToPointer() *Flow {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Flow) IsExact() bool {
	if e != nil {
		switch *e {
		case "signup", "login", "purchase":
			return true
		}
	}
	return false
}
