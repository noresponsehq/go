package components

// ListEntry - One entry of an environment's list: an exact email or an email domain that decides an account whatever its score. An account matches when one of its admitted assessments sent the email, or an email of the domain. A block entry decides block for every flow and an allow entry decides allow; when both match, block wins. The account's risk carries the list_match reason with the entry's id, and its score stays the one the assessment's facts give. A held flow still waits for its investigation.
type ListEntry struct {
	// The entry's id: lst_ and a ULID.
	ID string `json:"id"`
	// The workspace id, which is also its live environment's workspace id; the test environment's is the same id with :test appended. A workspace created through the API is tnt_ and a 26-character ULID.
	WorkspaceID string `json:"workspace_id"`
	// The environment the key belongs to. A live key assesses with live sources; a test key answers from the deterministic simulator.
	Environment KeyEnvironment `json:"environment"`
	// email matches one address, whatever the case of its local part. email_domain matches every address of the registrable domain and of its subdomains.
	Kind ListEntryKind `json:"kind"`
	// The value the entry matches, in canonical form: an email lower-cased whole, or a domain lower-cased and taken to its registrable domain (mail.example.co.uk is stored as example.co.uk).
	Value string `json:"value"`
	// block decides block for the accounts the entry matches; allow decides allow. Block wins when both match.
	Action ListEntryAction `json:"action"`
	// Why the entry exists, for the people who manage the list. It is shown with the entry and decides nothing.
	Note      *string `json:"note,omitzero"`
	CreatedAt string  `json:"created_at"`
	// The external user id of the member who added the entry through the console, or the prefix (nrk_<environment>_<key id>) of the management key that added it.
	CreatedBy string `json:"created_by"`
}

func (l *ListEntry) GetID() string {
	if l == nil {
		return ""
	}
	return l.ID
}

func (l *ListEntry) GetWorkspaceID() string {
	if l == nil {
		return ""
	}
	return l.WorkspaceID
}

func (l *ListEntry) GetEnvironment() KeyEnvironment {
	if l == nil {
		return KeyEnvironment("")
	}
	return l.Environment
}

func (l *ListEntry) GetKind() ListEntryKind {
	if l == nil {
		return ListEntryKind("")
	}
	return l.Kind
}

func (l *ListEntry) GetValue() string {
	if l == nil {
		return ""
	}
	return l.Value
}

func (l *ListEntry) GetAction() ListEntryAction {
	if l == nil {
		return ListEntryAction("")
	}
	return l.Action
}

func (l *ListEntry) GetNote() *string {
	if l == nil {
		return nil
	}
	return l.Note
}

func (l *ListEntry) GetCreatedAt() string {
	if l == nil {
		return ""
	}
	return l.CreatedAt
}

func (l *ListEntry) GetCreatedBy() string {
	if l == nil {
		return ""
	}
	return l.CreatedBy
}
