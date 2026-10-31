package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// Consistency - Cross-identifier agreement, computed from facts only. Every member is optional; absent makes no claim.
type Consistency struct {
	// How a submitted identifier agrees with what the evidence holds for the other identifiers. unknown is a definite 'cannot tell' (for example a truncated or too-common breach result); absent means the comparison was not made.
	Name *ConsistencyValue `json:"name,omitzero"`
	// How a submitted identifier agrees with what the evidence holds for the other identifiers. unknown is a definite 'cannot tell' (for example a truncated or too-common breach result); absent means the comparison was not made.
	Phone *ConsistencyValue `json:"phone,omitzero"`
	// How a submitted identifier agrees with what the evidence holds for the other identifiers. unknown is a definite 'cannot tell' (for example a truncated or too-common breach result); absent means the comparison was not made.
	IP *ConsistencyValue `json:"ip,omitzero"`
	// IP country, phone country, breach countries and policy_context.region together.
	Country *Country `json:"country,omitzero"`
}

func (c Consistency) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *Consistency) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (c *Consistency) GetName() *ConsistencyValue {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *Consistency) GetPhone() *ConsistencyValue {
	if c == nil {
		return nil
	}
	return c.Phone
}

func (c *Consistency) GetIP() *ConsistencyValue {
	if c == nil {
		return nil
	}
	return c.IP
}

func (c *Consistency) GetCountry() *Country {
	if c == nil {
		return nil
	}
	return c.Country
}
