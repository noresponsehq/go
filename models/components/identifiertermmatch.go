package components

// IdentifierTermMatch - One satisfied unique term in an identifier expression. term_index is the zero-based depth-first raw leaf position of its first occurrence. Duplicate normalized terms produce one witness. All satisfied unique terms are returned, ordered by term_index, including satisfied terms in an OR branch that was not needed. value preserves the original submitted value in assessment_id. In assessment scope all witnesses belong to the same matching assessment; account search chooses the newest such assessment. In account scope each term uses its newest retained matching assessment, so witnesses may have different assessment IDs. This evidence does not assert ownership or equality between accounts.
type IdentifierTermMatch struct {
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
	// Zero-based depth-first position of the first occurrence of this normalized term in the submitted expression.
	TermIndex int64 `json:"term_index"`
}

func (i *IdentifierTermMatch) GetKind() IdentifierSearchKind {
	if i == nil {
		return IdentifierSearchKind("")
	}
	return i.Kind
}

func (i *IdentifierTermMatch) GetValue() string {
	if i == nil {
		return ""
	}
	return i.Value
}

func (i *IdentifierTermMatch) GetMatchBasis() Match {
	if i == nil {
		return Match("")
	}
	return i.MatchBasis
}

func (i *IdentifierTermMatch) GetProvenance() Provenance {
	if i == nil {
		return Provenance("")
	}
	return i.Provenance
}

func (i *IdentifierTermMatch) GetNamespace() *Namespace {
	if i == nil {
		return nil
	}
	return i.Namespace
}

func (i *IdentifierTermMatch) GetAssessmentID() string {
	if i == nil {
		return ""
	}
	return i.AssessmentID
}

func (i *IdentifierTermMatch) GetTermIndex() int64 {
	if i == nil {
		return 0
	}
	return i.TermIndex
}
