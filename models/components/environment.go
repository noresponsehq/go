package components

type Environment string

const (
	EnvironmentLive      Environment = "live"
	EnvironmentTest      Environment = "test"
	EnvironmentWorkspace Environment = "workspace"
)

func (e Environment) ToPointer() *Environment {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Environment) IsExact() bool {
	if e != nil {
		switch *e {
		case "live", "test", "workspace":
			return true
		}
	}
	return false
}
