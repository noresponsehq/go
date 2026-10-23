package components

// Depth - How long the synchronous answer waits for sources. quick waits only for the fast sources, those that answer within 500 ms of ingress (or deadline_ms when lower); full waits for the whole budget, up to 3000 ms. Sources still running at the cut continue asynchronously and update the assessment either way.
type Depth string

const (
	DepthQuick Depth = "quick"
	DepthFull  Depth = "full"
)

func (e Depth) ToPointer() *Depth {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Depth) IsExact() bool {
	if e != nil {
		switch *e {
		case "quick", "full":
			return true
		}
	}
	return false
}
