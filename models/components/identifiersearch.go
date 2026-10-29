package components

// IdentifierSearch - One customer-submitted identifier. Defaults to exact comparison of the original submitted value. canonical uses the supported identifier normalisation without an ownership claim: email preserves local-part case, NFC-normalises it and lowercases the domain without removing dots or tags; phone uses international E.164 without extensions or an inferred region; username uses NFC and Unicode case folding within the customer namespace. mailbox is email-only and explicitly selects existing list rules: case-insensitive, remove +tag, and ignore local-part dots at gmail.com and googlemail.com; googlemail.com maps to gmail.com. Username requires namespace customer. Email and phone forbid namespace. Source-discovered identifiers and platform handles are not searched. Unsupported kinds, namespaces, modes and invalid values return 400. Search equality does not establish verified ownership or merge accounts.
type IdentifierSearch struct {
	Kind IdentifierSearchKind `json:"kind"`
	// The comparison value. Stored original identifiers remain unchanged. An invalid identifier is refused instead of returning an empty result.
	Value string `json:"value"`
	// exact compares the original submitted value; canonical uses identifier-specific normalisation; mailbox explicitly selects the email alias rules used by lists.
	Match *Match `json:"match,omitzero"`
	// Required for username: the username submitted by this customer in this workspace/environment. It does not name or join a source platform. Omit for email and phone.
	Namespace *Namespace `json:"namespace,omitzero"`
}

func (i *IdentifierSearch) GetKind() IdentifierSearchKind {
	if i == nil {
		return IdentifierSearchKind("")
	}
	return i.Kind
}

func (i *IdentifierSearch) GetValue() string {
	if i == nil {
		return ""
	}
	return i.Value
}

func (i *IdentifierSearch) GetMatch() *Match {
	if i == nil {
		return nil
	}
	return i.Match
}

func (i *IdentifierSearch) GetNamespace() *Namespace {
	if i == nil {
		return nil
	}
	return i.Namespace
}
