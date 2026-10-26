package components

// FootprintFamily - The source-neutral kind of evidence: account_presence (an account found on a site), auth_provider (the account at your auth provider), breach (breach and credential tables), breach_abuse (abuse and fraud communities' tables) and code_hosting (a code hosting profile).
type FootprintFamily string

const (
	FootprintFamilyAccountPresence FootprintFamily = "account_presence"
	FootprintFamilyAuthProvider    FootprintFamily = "auth_provider"
	FootprintFamilyBreach          FootprintFamily = "breach"
	FootprintFamilyBreachAbuse     FootprintFamily = "breach_abuse"
	FootprintFamilyCodeHosting     FootprintFamily = "code_hosting"
)

func (e FootprintFamily) ToPointer() *FootprintFamily {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *FootprintFamily) IsExact() bool {
	if e != nil {
		switch *e {
		case "account_presence", "auth_provider", "breach", "breach_abuse", "code_hosting":
			return true
		}
	}
	return false
}
