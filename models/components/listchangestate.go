package components

type ListChangeState string

const (
	ListChangeStatePropagating ListChangeState = "propagating"
	ListChangeStateCompleted   ListChangeState = "completed"
)

func (e ListChangeState) ToPointer() *ListChangeState {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ListChangeState) IsExact() bool {
	if e != nil {
		switch *e {
		case "propagating", "completed":
			return true
		}
	}
	return false
}
