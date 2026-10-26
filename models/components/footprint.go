package components

import (
	"github.com/noresponsehq/go/v8/internal/utils"
)

// Footprint - How long and how widely the signup's email and username have existed, and how far independent sources agree on the person behind them, in years and counts only. families are the kinds of evidence that found either identifier. Omitted when nothing found them: a footprint lookup that failed makes no claim, and absence is no_digital_footprint's to state.
type Footprint struct {
	// One identifier's dated footprint across every source: the first and last year any event was seen, the years between them, the latest year a source dated activity (omitted when none did), how many distinct sites and which kinds of evidence saw it.
	Email *FootprintTimeline `json:"email,omitzero"`
	// One identifier's dated footprint across every source: the first and last year any event was seen, the years between them, the latest year a source dated activity (omitted when none did), how many distinct sites and which kinds of evidence saw it.
	Username  *FootprintTimeline  `json:"username,omitzero"`
	Families  []FootprintFamily   `json:"families"`
	Coherence *FootprintCoherence `json:"coherence,omitzero"`
}

func (f Footprint) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *Footprint) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (f *Footprint) GetEmail() *FootprintTimeline {
	if f == nil {
		return nil
	}
	return f.Email
}

func (f *Footprint) GetUsername() *FootprintTimeline {
	if f == nil {
		return nil
	}
	return f.Username
}

func (f *Footprint) GetFamilies() []FootprintFamily {
	if f == nil {
		return []FootprintFamily{}
	}
	return f.Families
}

func (f *Footprint) GetCoherence() *FootprintCoherence {
	if f == nil {
		return nil
	}
	return f.Coherence
}
