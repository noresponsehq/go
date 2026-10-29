package components

type CustomerExportRequestFilter struct {
	Action      *string                             `json:"action,omitzero"`
	Actor       *string                             `json:"actor,omitzero"`
	TargetType  *string                             `json:"target_type,omitzero"`
	TargetID    *string                             `json:"target_id,omitzero"`
	Environment *Environment                        `json:"environment,omitzero"`
	Outcome     *CustomerExportRequestFilterOutcome `json:"outcome,omitzero"`
	Since       *string                             `json:"since,omitzero"`
	Until       *string                             `json:"until,omitzero"`
}

func (c *CustomerExportRequestFilter) GetAction() *string {
	if c == nil {
		return nil
	}
	return c.Action
}

func (c *CustomerExportRequestFilter) GetActor() *string {
	if c == nil {
		return nil
	}
	return c.Actor
}

func (c *CustomerExportRequestFilter) GetTargetType() *string {
	if c == nil {
		return nil
	}
	return c.TargetType
}

func (c *CustomerExportRequestFilter) GetTargetID() *string {
	if c == nil {
		return nil
	}
	return c.TargetID
}

func (c *CustomerExportRequestFilter) GetEnvironment() *Environment {
	if c == nil {
		return nil
	}
	return c.Environment
}

func (c *CustomerExportRequestFilter) GetOutcome() *CustomerExportRequestFilterOutcome {
	if c == nil {
		return nil
	}
	return c.Outcome
}

func (c *CustomerExportRequestFilter) GetSince() *string {
	if c == nil {
		return nil
	}
	return c.Since
}

func (c *CustomerExportRequestFilter) GetUntil() *string {
	if c == nil {
		return nil
	}
	return c.Until
}
