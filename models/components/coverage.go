package components

// Coverage - Counts only; a zero denominator means no eligible retained records and does not imply a zero rate.
type Coverage struct {
	Numerator   int64 `json:"numerator"`
	Denominator int64 `json:"denominator"`
}

func (c *Coverage) GetNumerator() int64 {
	if c == nil {
		return 0
	}
	return c.Numerator
}

func (c *Coverage) GetDenominator() int64 {
	if c == nil {
		return 0
	}
	return c.Denominator
}
