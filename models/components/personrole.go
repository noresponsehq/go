package components

// PersonRole - The person's role in the group: guest (a ticket holder), crew, vendor or press. Leave it out when you do not know it: a person without one has no role, not guest. A later row naming the same person replaces it, and a change of role alone keeps the person's screening results.
type PersonRole string

const (
	PersonRoleGuest  PersonRole = "guest"
	PersonRoleCrew   PersonRole = "crew"
	PersonRoleVendor PersonRole = "vendor"
	PersonRolePress  PersonRole = "press"
)

func (e PersonRole) ToPointer() *PersonRole {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *PersonRole) IsExact() bool {
	if e != nil {
		switch *e {
		case "guest", "crew", "vendor", "press":
			return true
		}
	}
	return false
}
