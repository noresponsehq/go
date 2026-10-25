package components

// KeyEnvironment - The environment the key belongs to. A live key assesses with live sources; a test key answers from the deterministic simulator.
type KeyEnvironment string

const (
	KeyEnvironmentLive KeyEnvironment = "live"
	KeyEnvironmentTest KeyEnvironment = "test"
)

func (e KeyEnvironment) ToPointer() *KeyEnvironment {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *KeyEnvironment) IsExact() bool {
	if e != nil {
		switch *e {
		case "live", "test":
			return true
		}
	}
	return false
}
