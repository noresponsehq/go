package components

// AssessmentRequestIdentifiers - At least one of email, username or phone. name and ip refine the match and travel beside a lookup identifier; a body with only name or ip is refused with 400.
type AssessmentRequestIdentifiers struct {
	// One address: a 1 to 64 character local part without whitespace, angle brackets, control or invisible characters, '@', and a domain with a letter TLD. Send an internationalised domain in punycode. Nothing is trimmed. The local part keeps its case and the domain is lower-cased; no alias rewriting ('+tag', dots or provider heuristics) is applied.
	Email *string `json:"email,omitzero"`
	// The handle the account is signing up with: 1 to 64 characters, no whitespace, control or invisible characters, and not a credential hash or token. Echoed as received; compared case-insensitively.
	Username *string `json:"username,omitzero"`
	// Strict E.164: '+', then 8 to 15 digits, the first not 0, with no separators. Phone evidence is gathered asynchronously, so a phone-only request answers 202 pending and the phone is listed in incomplete_sources on the synchronous answer.
	Phone *string `json:"phone,omitzero"`
	// The full name as the user typed it, 1 to 128 characters with no control or invisible characters. It is used only to match names in the evidence.
	Name *string `json:"name,omitzero"`
	// The account_creation's client IP address: an IPv4 dotted quad without leading zeros, or an IPv6 address in its RFC 5952 text form (lower case, no leading zeros, the longest zero run compressed). Send an IPv4 client as a dotted quad, not IPv4-mapped. No zone, port or prefix. The server refuses a non-canonical IPv6 spelling, and every private, loopback, link-local, multicast, documentation or otherwise reserved address (the IANA special-purpose registries), with 400: only a public address says anything about a signup.
	IP *string `json:"ip,omitzero"`
}

func (a *AssessmentRequestIdentifiers) GetEmail() *string {
	if a == nil {
		return nil
	}
	return a.Email
}

func (a *AssessmentRequestIdentifiers) GetUsername() *string {
	if a == nil {
		return nil
	}
	return a.Username
}

func (a *AssessmentRequestIdentifiers) GetPhone() *string {
	if a == nil {
		return nil
	}
	return a.Phone
}

func (a *AssessmentRequestIdentifiers) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *AssessmentRequestIdentifiers) GetIP() *string {
	if a == nil {
		return nil
	}
	return a.IP
}
