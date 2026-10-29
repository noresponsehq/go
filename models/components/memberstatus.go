package components

// MemberStatus - Listed and returned members are active.
type MemberStatus string

const (
	MemberStatusActive MemberStatus = "active"
)

func (e MemberStatus) ToPointer() *MemberStatus {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *MemberStatus) IsExact() bool {
	if e != nil {
		switch *e {
		case "active":
			return true
		}
	}
	return false
}
