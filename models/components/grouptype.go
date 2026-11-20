package components

// GroupType - What the group's people share: an event, an organisation, a place or another attribute.
type GroupType string

const (
	GroupTypeEvent        GroupType = "event"
	GroupTypeOrganisation GroupType = "organisation"
	GroupTypePlace        GroupType = "place"
	GroupTypeCustom       GroupType = "custom"
)

func (e GroupType) ToPointer() *GroupType {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *GroupType) IsExact() bool {
	if e != nil {
		switch *e {
		case "event", "organisation", "place", "custom":
			return true
		}
	}
	return false
}
