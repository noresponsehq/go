package components

import (
	"encoding/json"
	"fmt"
)

// Transport - Where the customer derived this context. Backend is the trusted integration path; browser values are treated as weak hints.
type Transport string

const (
	TransportBackend Transport = "backend"
	TransportBrowser Transport = "browser"
)

func (e Transport) ToPointer() *Transport {
	return &e
}
func (e *Transport) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "backend":
		fallthrough
	case "browser":
		*e = Transport(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Transport: %v", v)
	}
}
