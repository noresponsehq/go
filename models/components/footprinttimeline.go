package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// FootprintTimeline - One identifier's dated footprint across every source: the first and last year any event was seen, the years between them, the latest year a source dated activity (omitted when none did), how many distinct sites and which kinds of evidence saw it.
type FootprintTimeline struct {
	FirstSeenYear  int64             `json:"first_seen_year"`
	LastSeenYear   int64             `json:"last_seen_year"`
	SpanYears      int64             `json:"span_years"`
	LastActiveYear *int64            `json:"last_active_year,omitzero"`
	Sites          int64             `json:"sites"`
	Families       []FootprintFamily `json:"families"`
}

func (f FootprintTimeline) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FootprintTimeline) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (f *FootprintTimeline) GetFirstSeenYear() int64 {
	if f == nil {
		return 0
	}
	return f.FirstSeenYear
}

func (f *FootprintTimeline) GetLastSeenYear() int64 {
	if f == nil {
		return 0
	}
	return f.LastSeenYear
}

func (f *FootprintTimeline) GetSpanYears() int64 {
	if f == nil {
		return 0
	}
	return f.SpanYears
}

func (f *FootprintTimeline) GetLastActiveYear() *int64 {
	if f == nil {
		return nil
	}
	return f.LastActiveYear
}

func (f *FootprintTimeline) GetSites() int64 {
	if f == nil {
		return 0
	}
	return f.Sites
}

func (f *FootprintTimeline) GetFamilies() []FootprintFamily {
	if f == nil {
		return []FootprintFamily{}
	}
	return f.Families
}
