package components

// ListEntryRequest - POST /v1/workspaces/{workspace_id}/lists/{environment} body. Adds one entry to the environment's list and, in the same change, re-decides every account it matches: each sends account.risk_changed when a decision changes. Needs the admin or owner role, or a management key of the environment.
type ListEntryRequest struct {
	// email matches one address, whatever the case of its local part. email_domain matches every address of the registrable domain and of its subdomains.
	Kind ListEntryKind `json:"kind"`
	// An email address for kind email, written as in an assessment request; an LDH domain (A-label for an internationalised one) such as example.com for kind email_domain, which must sit below a public suffix (co.uk alone is refused with 400). NoResponse stores it in canonical form and answers that form.
	Value string `json:"value"`
	// block decides block for the accounts the entry matches; allow decides allow. Block wins when both match.
	Action ListEntryAction `json:"action"`
	// Why the entry exists, for the people who manage the list. It is shown with the entry and decides nothing.
	Note *string `json:"note,omitzero"`
	// Optional expiry. An expired entry stops matching immediately; bounded maintenance removes it and propagates account risk changes.
	ExpiresAt *string `json:"expires_at,omitzero"`
}

func (l *ListEntryRequest) GetKind() ListEntryKind {
	if l == nil {
		return ListEntryKind("")
	}
	return l.Kind
}

func (l *ListEntryRequest) GetValue() string {
	if l == nil {
		return ""
	}
	return l.Value
}

func (l *ListEntryRequest) GetAction() ListEntryAction {
	if l == nil {
		return ListEntryAction("")
	}
	return l.Action
}

func (l *ListEntryRequest) GetNote() *string {
	if l == nil {
		return nil
	}
	return l.Note
}

func (l *ListEntryRequest) GetExpiresAt() *string {
	if l == nil {
		return nil
	}
	return l.ExpiresAt
}
