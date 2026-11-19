package components

// PersonGroup - One group the person belongs to and their role in it.
type PersonGroup struct {
	GroupID   string `json:"group_id"`
	GroupName string `json:"group_name"`
	// The group's optional event start; null for an ongoing group.
	GroupStartsAt *string `json:"group_starts_at"`
	// The person's role in the group: guest, crew, vendor or press, or null when it was not given.
	Role *PersonRole `json:"role"`
	// When the person was added to the group.
	AddedAt string `json:"added_at"`
}

func (p *PersonGroup) GetGroupID() string {
	if p == nil {
		return ""
	}
	return p.GroupID
}

func (p *PersonGroup) GetGroupName() string {
	if p == nil {
		return ""
	}
	return p.GroupName
}

func (p *PersonGroup) GetGroupStartsAt() *string {
	if p == nil {
		return nil
	}
	return p.GroupStartsAt
}

func (p *PersonGroup) GetRole() *PersonRole {
	if p == nil {
		return nil
	}
	return p.Role
}

func (p *PersonGroup) GetAddedAt() string {
	if p == nil {
		return ""
	}
	return p.AddedAt
}
