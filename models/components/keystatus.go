package components

// KeyStatus - active: the key authenticates. expired: a rotation's overlap window has ended (expires_at). revoked: the key was revoked (revoked_at) and answers 401.
type KeyStatus string

const (
	KeyStatusActive  KeyStatus = "active"
	KeyStatusExpired KeyStatus = "expired"
	KeyStatusRevoked KeyStatus = "revoked"
)

func (e KeyStatus) ToPointer() *KeyStatus {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *KeyStatus) IsExact() bool {
	if e != nil {
		switch *e {
		case "active", "expired", "revoked":
			return true
		}
	}
	return false
}
