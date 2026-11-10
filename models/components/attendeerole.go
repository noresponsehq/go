package components

// AttendeeRole - The attendee's role at the event: guest (a ticket holder), crew, vendor or press. Leave it out when you do not know it: an attendee without one has no role, not guest. A later row naming the same attendee replaces it, and a change of role alone keeps the attendee's screening results.
type AttendeeRole string

const (
	AttendeeRoleGuest  AttendeeRole = "guest"
	AttendeeRoleCrew   AttendeeRole = "crew"
	AttendeeRoleVendor AttendeeRole = "vendor"
	AttendeeRolePress  AttendeeRole = "press"
)

func (e AttendeeRole) ToPointer() *AttendeeRole {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *AttendeeRole) IsExact() bool {
	if e != nil {
		switch *e {
		case "guest", "crew", "vendor", "press":
			return true
		}
	}
	return false
}
