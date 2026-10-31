package components

// BreachConduct - What a breach table's rows record against the identifier: chargeback, one or more chargebacks; declined_payment, one or more declined payments; forum_sanction, a forum ban, warning or infraction. A category and not a value: the schema has no member for one.
type BreachConduct string

const (
	BreachConductChargeback      BreachConduct = "chargeback"
	BreachConductDeclinedPayment BreachConduct = "declined_payment"
	BreachConductForumSanction   BreachConduct = "forum_sanction"
)

func (e BreachConduct) ToPointer() *BreachConduct {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *BreachConduct) IsExact() bool {
	if e != nil {
		switch *e {
		case "chargeback", "declined_payment", "forum_sanction":
			return true
		}
	}
	return false
}
