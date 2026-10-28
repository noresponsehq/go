package components

// CredentialRequest - Associate the SHA-256 of a high-entropy key issued by your service. Send the hash instead of the key. Repeating a hash returns its original association, including its retirement state. A hash associated with another account returns a conflict.
type CredentialRequest struct {
	Sha256 string `json:"sha256"`
}

func (c *CredentialRequest) GetSha256() string {
	if c == nil {
		return ""
	}
	return c.Sha256
}
