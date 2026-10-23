package components

import (
	"encoding/json"
	"fmt"
)

// Depth - How long the synchronous answer waits for sources. quick waits only for the fast sources, those that answer within 500 ms of ingress (or deadline_ms when lower); full waits for the whole budget, up to 3000 ms. Sources still running at the cut continue asynchronously and update the assessment either way.
type Depth string

const (
	DepthQuick Depth = "quick"
	DepthFull  Depth = "full"
)

func (e Depth) ToPointer() *Depth {
	return &e
}
func (e *Depth) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "quick":
		fallthrough
	case "full":
		*e = Depth(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Depth: %v", v)
	}
}
