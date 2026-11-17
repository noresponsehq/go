package components

import (
	"github.com/noresponsehq/go/v15/internal/utils"
)

// ScreeningRequest - Screens the group's people no screening has reached and everyone on watch or escalated.
type ScreeningRequest struct {
	Depth ScreeningDepth `json:"depth"`
	// The platforms to search. Leave it out for a full search: every platform a connector serves in the environment when the screening starts, which the screening then lists.
	Platforms []string `json:"platforms,omitzero"`
}

func (s ScreeningRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ScreeningRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (s *ScreeningRequest) GetDepth() ScreeningDepth {
	if s == nil {
		return ScreeningDepth("")
	}
	return s.Depth
}

func (s *ScreeningRequest) GetPlatforms() []string {
	if s == nil {
		return nil
	}
	return s.Platforms
}
