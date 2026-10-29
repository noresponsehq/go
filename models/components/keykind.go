package components

type KeyKind string

const (
	KeyKindService  KeyKind = "service"
	KeyKindPersonal KeyKind = "personal"
)

func (e KeyKind) ToPointer() *KeyKind {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *KeyKind) IsExact() bool {
	if e != nil {
		switch *e {
		case "service", "personal":
			return true
		}
	}
	return false
}
