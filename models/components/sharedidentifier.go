package components

type SharedIdentifier struct {
	Kind IdentifierSearchKind `json:"kind"`
	// exact compares the original submitted value; canonical uses identifier-specific normalisation; mailbox explicitly selects the email alias rules used by lists.
	MatchBasis Match `json:"match_basis"`
	// Customer-submitted input. It does not mean the customer or NoResponse verified ownership.
	Provenance Provenance `json:"provenance"`
	// Required for username: the username submitted by this customer in this workspace/environment. It does not name or join a source platform. Omit for email and phone.
	Namespace     *Namespace           `json:"namespace,omitzero"`
	Account       IdentifierSubmission `json:"account"`
	LinkedAccount IdentifierSubmission `json:"linked_account"`
}

func (s *SharedIdentifier) GetKind() IdentifierSearchKind {
	if s == nil {
		return IdentifierSearchKind("")
	}
	return s.Kind
}

func (s *SharedIdentifier) GetMatchBasis() Match {
	if s == nil {
		return Match("")
	}
	return s.MatchBasis
}

func (s *SharedIdentifier) GetProvenance() Provenance {
	if s == nil {
		return Provenance("")
	}
	return s.Provenance
}

func (s *SharedIdentifier) GetNamespace() *Namespace {
	if s == nil {
		return nil
	}
	return s.Namespace
}

func (s *SharedIdentifier) GetAccount() IdentifierSubmission {
	if s == nil {
		return IdentifierSubmission{}
	}
	return s.Account
}

func (s *SharedIdentifier) GetLinkedAccount() IdentifierSubmission {
	if s == nil {
		return IdentifierSubmission{}
	}
	return s.LinkedAccount
}
