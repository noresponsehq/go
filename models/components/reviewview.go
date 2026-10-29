package components

type ReviewView struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
	// The credential namespace of the creator. An external user ID equal to an API-key label does not grant ownership of that key's private view.
	OwnerType OwnerType `json:"owner_type"`
	Shared    bool      `json:"shared"`
	Version   int64     `json:"version"`
	Filter    Filter    `json:"filter"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	// Immutable member ID or API key ID, interpreted with owner_type. Ownership survives changes to an external user identity.
	OwnerPrincipalID string `json:"owner_principal_id"`
}

func (r *ReviewView) GetID() string {
	if r == nil {
		return ""
	}
	return r.ID
}

func (r *ReviewView) GetName() string {
	if r == nil {
		return ""
	}
	return r.Name
}

func (r *ReviewView) GetOwner() string {
	if r == nil {
		return ""
	}
	return r.Owner
}

func (r *ReviewView) GetOwnerType() OwnerType {
	if r == nil {
		return OwnerType("")
	}
	return r.OwnerType
}

func (r *ReviewView) GetShared() bool {
	if r == nil {
		return false
	}
	return r.Shared
}

func (r *ReviewView) GetVersion() int64 {
	if r == nil {
		return 0
	}
	return r.Version
}

func (r *ReviewView) GetFilter() Filter {
	if r == nil {
		return Filter{}
	}
	return r.Filter
}

func (r *ReviewView) GetCreatedAt() string {
	if r == nil {
		return ""
	}
	return r.CreatedAt
}

func (r *ReviewView) GetUpdatedAt() string {
	if r == nil {
		return ""
	}
	return r.UpdatedAt
}

func (r *ReviewView) GetOwnerPrincipalID() string {
	if r == nil {
		return ""
	}
	return r.OwnerPrincipalID
}
