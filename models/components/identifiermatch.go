package components

// IdentifierMatch - Why this result matched a typed identifier search. value is the original value submitted in assessment_id. Provenance records a customer submission, not a source discovery or verified ownership. Equality is deterministic and has no inferred confidence score. An account result names its newest retained matching assessment, which can differ from latest.
type IdentifierMatch struct {
	Kind  IdentifierSearchKind `json:"kind"`
	Value string               `json:"value"`
	// exact compares the original submitted value; canonical uses identifier-specific normalisation; mailbox explicitly selects the email alias rules used by lists.
	MatchBasis Match `json:"match_basis"`
	// Customer-submitted input. It does not mean the customer or NoResponse verified ownership.
	Provenance Provenance `json:"provenance"`
	// Required for username: the username submitted by this customer in this workspace/environment. It does not name or join a source platform. Omit for email and phone.
	Namespace *Namespace `json:"namespace,omitzero"`
	// Opaque, server-assigned assessment id: 'asm_' and a 26-character Crockford base32 ULID. Unique within the workspace; another workspace's assessment id answers 404.
	AssessmentID string `json:"assessment_id"`
}

func (i *IdentifierMatch) GetKind() IdentifierSearchKind {
	if i == nil {
		return IdentifierSearchKind("")
	}
	return i.Kind
}

func (i *IdentifierMatch) GetValue() string {
	if i == nil {
		return ""
	}
	return i.Value
}

func (i *IdentifierMatch) GetMatchBasis() Match {
	if i == nil {
		return Match("")
	}
	return i.MatchBasis
}

func (i *IdentifierMatch) GetProvenance() Provenance {
	if i == nil {
		return Provenance("")
	}
	return i.Provenance
}

func (i *IdentifierMatch) GetNamespace() *Namespace {
	if i == nil {
		return nil
	}
	return i.Namespace
}

func (i *IdentifierMatch) GetAssessmentID() string {
	if i == nil {
		return ""
	}
	return i.AssessmentID
}
