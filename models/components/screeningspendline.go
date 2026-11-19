package components

// ScreeningSpendLine - What one implementation cost.
type ScreeningSpendLine struct {
	// The resolver or collector that produced a record, by its catalogue name.
	Implementation string `json:"implementation"`
	// What the implementation did. analyser appears only on spend recorded before the engine stopped analysing posts.
	Role       ScreeningSpendLineRole `json:"role"`
	SpendMinor int64                  `json:"spend_minor"`
}

func (s *ScreeningSpendLine) GetImplementation() string {
	if s == nil {
		return ""
	}
	return s.Implementation
}

func (s *ScreeningSpendLine) GetRole() ScreeningSpendLineRole {
	if s == nil {
		return ScreeningSpendLineRole("")
	}
	return s.Role
}

func (s *ScreeningSpendLine) GetSpendMinor() int64 {
	if s == nil {
		return 0
	}
	return s.SpendMinor
}
