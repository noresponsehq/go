package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// ScreeningCompletedByBand - Screened attendees by the band of their most severe finding.
type ScreeningCompletedByBand struct {
	Critical int64 `json:"critical"`
	High     int64 `json:"high"`
	Moderate int64 `json:"moderate"`
	Low      int64 `json:"low"`
	Info     int64 `json:"info"`
}

func (s ScreeningCompletedByBand) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ScreeningCompletedByBand) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (s *ScreeningCompletedByBand) GetCritical() int64 {
	if s == nil {
		return 0
	}
	return s.Critical
}

func (s *ScreeningCompletedByBand) GetHigh() int64 {
	if s == nil {
		return 0
	}
	return s.High
}

func (s *ScreeningCompletedByBand) GetModerate() int64 {
	if s == nil {
		return 0
	}
	return s.Moderate
}

func (s *ScreeningCompletedByBand) GetLow() int64 {
	if s == nil {
		return 0
	}
	return s.Low
}

func (s *ScreeningCompletedByBand) GetInfo() int64 {
	if s == nil {
		return 0
	}
	return s.Info
}
