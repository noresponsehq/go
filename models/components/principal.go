package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// Principal - A person the event protects.
type Principal struct {
	// The principal's name.
	Name string `json:"name"`
	// Other names the principal goes by, at most 10.
	Aliases []string `json:"aliases,omitzero"`
}

func (p Principal) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Principal) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *Principal) GetName() string {
	if p == nil {
		return ""
	}
	return p.Name
}

func (p *Principal) GetAliases() []string {
	if p == nil {
		return nil
	}
	return p.Aliases
}
