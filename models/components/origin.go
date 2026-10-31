package components

import (
	"encoding/json"
	"fmt"
)

// Origin - api (the default) searches the assessments your integration sent. lookup searches the assessments your team ran by hand on the console's Lookup page. The API sets an assessment's origin from the credential that sent it. A lookup is not in any other list, count, webhook event or account.
type Origin string

const (
	OriginAPI    Origin = "api"
	OriginLookup Origin = "lookup"
)

func (e Origin) ToPointer() *Origin {
	return &e
}
func (e *Origin) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api":
		fallthrough
	case "lookup":
		*e = Origin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Origin: %v", v)
	}
}
