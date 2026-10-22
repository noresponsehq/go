package components

import (
	"encoding/json"
	"fmt"
)

// Transport - Where this context was derived. Backend is the trusted integration path; browser values are treated as weak hints. provider_hook is a hosted provider hook (an Auth0 or WorkOS Action) calling at hook_point; it needs hook_point and verified false, and its authentication facts are not scored.
type Transport string

const (
	TransportBackend      Transport = "backend"
	TransportBrowser      Transport = "browser"
	TransportProviderHook Transport = "provider_hook"
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
		fallthrough
	case "provider_hook":
		*e = Transport(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Transport: %v", v)
	}
}
