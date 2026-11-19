package components

import (
	"github.com/noresponsehq/go/v16/internal/utils"
)

// ScreeningRequest - Screens the group's people no screening has reached.
type ScreeningRequest struct {
	Depth ScreeningDepth `json:"depth"`
	// Deprecated: leave it out. Without it the screening is a full search: every platform a connector serves in the environment when the screening starts, which the screening then lists, each searched only with the identifiers its connector takes. A list sent here still limits the screening until a later contract release removes the field.
	//
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
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
