package components

import (
	"encoding/json"
	"fmt"
)

// Environment - The environment whose endpoints receive this event: live, or test for the workspace's test environment.
type Environment string

const (
	EnvironmentLive Environment = "live"
	EnvironmentTest Environment = "test"
)

func (e Environment) ToPointer() *Environment {
	return &e
}
func (e *Environment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "live":
		fallthrough
	case "test":
		*e = Environment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Environment: %v", v)
	}
}
