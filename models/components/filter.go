package components

type Filter struct {
	Status              *ReviewStatus `json:"status,omitzero"`
	Priority            *Priority     `json:"priority,omitzero"`
	Assignee            *string       `json:"assignee,omitzero"`
	Unassigned          *bool         `json:"unassigned,omitzero"`
	AssigneePrincipalID *string       `json:"assignee_principal_id,omitzero"`
}

func (f *Filter) GetStatus() *ReviewStatus {
	if f == nil {
		return nil
	}
	return f.Status
}

func (f *Filter) GetPriority() *Priority {
	if f == nil {
		return nil
	}
	return f.Priority
}

func (f *Filter) GetAssignee() *string {
	if f == nil {
		return nil
	}
	return f.Assignee
}

func (f *Filter) GetUnassigned() *bool {
	if f == nil {
		return nil
	}
	return f.Unassigned
}

func (f *Filter) GetAssigneePrincipalID() *string {
	if f == nil {
		return nil
	}
	return f.AssigneePrincipalID
}
