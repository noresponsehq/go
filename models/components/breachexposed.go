package components

// BreachExposed - The kind of data a breach table held beside the identifier. password means a plaintext password was present and password_hash means a hash, salt or token was present. The value itself is not stored or returned; the schema has no member for one.
type BreachExposed string

const (
	BreachExposedEmail        BreachExposed = "email"
	BreachExposedUsername     BreachExposed = "username"
	BreachExposedName         BreachExposed = "name"
	BreachExposedPassword     BreachExposed = "password"
	BreachExposedPasswordHash BreachExposed = "password_hash"
	BreachExposedPhone        BreachExposed = "phone"
	BreachExposedAddress      BreachExposed = "address"
	BreachExposedDob          BreachExposed = "dob"
	BreachExposedIP           BreachExposed = "ip"
	BreachExposedPayment      BreachExposed = "payment"
	BreachExposedSecurityQa   BreachExposed = "security_qa"
)

func (e BreachExposed) ToPointer() *BreachExposed {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *BreachExposed) IsExact() bool {
	if e != nil {
		switch *e {
		case "email", "username", "name", "password", "password_hash", "phone", "address", "dob", "ip", "payment", "security_qa":
			return true
		}
	}
	return false
}
