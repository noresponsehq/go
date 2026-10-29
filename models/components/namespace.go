package components

// Namespace - Required for username: the username submitted by this customer in this workspace/environment. It does not name or join a source platform. Omit for email and phone.
type Namespace string

const (
	NamespaceCustomer Namespace = "customer"
)

func (e Namespace) ToPointer() *Namespace {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Namespace) IsExact() bool {
	if e != nil {
		switch *e {
		case "customer":
			return true
		}
	}
	return false
}
