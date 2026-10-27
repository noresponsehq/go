package components

// ListEntryAction - block decides block for the accounts the entry matches; allow decides allow. Block wins when both match.
type ListEntryAction string

const (
	ListEntryActionBlock ListEntryAction = "block"
	ListEntryActionAllow ListEntryAction = "allow"
)

func (e ListEntryAction) ToPointer() *ListEntryAction {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ListEntryAction) IsExact() bool {
	if e != nil {
		switch *e {
		case "block", "allow":
			return true
		}
	}
	return false
}
