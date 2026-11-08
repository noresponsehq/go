package components

import (
	"encoding/json"
	"fmt"
)

type ScreeningRequestPlatform string

const (
	ScreeningRequestPlatformInstagram ScreeningRequestPlatform = "instagram"
	ScreeningRequestPlatformX         ScreeningRequestPlatform = "x"
	ScreeningRequestPlatformBluesky   ScreeningRequestPlatform = "bluesky"
	ScreeningRequestPlatformYoutube   ScreeningRequestPlatform = "youtube"
	ScreeningRequestPlatformTiktok    ScreeningRequestPlatform = "tiktok"
	ScreeningRequestPlatformFacebook  ScreeningRequestPlatform = "facebook"
	ScreeningRequestPlatformThreads   ScreeningRequestPlatform = "threads"
	ScreeningRequestPlatformReddit    ScreeningRequestPlatform = "reddit"
	ScreeningRequestPlatformTelegram  ScreeningRequestPlatform = "telegram"
	ScreeningRequestPlatformLinkedin  ScreeningRequestPlatform = "linkedin"
)

func (e ScreeningRequestPlatform) ToPointer() *ScreeningRequestPlatform {
	return &e
}
func (e *ScreeningRequestPlatform) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "instagram":
		fallthrough
	case "x":
		fallthrough
	case "bluesky":
		fallthrough
	case "youtube":
		fallthrough
	case "tiktok":
		fallthrough
	case "facebook":
		fallthrough
	case "threads":
		fallthrough
	case "reddit":
		fallthrough
	case "telegram":
		fallthrough
	case "linkedin":
		*e = ScreeningRequestPlatform(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ScreeningRequestPlatform: %v", v)
	}
}

// ScreeningRequest - Screens the event's attendees no screening has reached and everyone on watch or escalated.
type ScreeningRequest struct {
	Depth     ScreeningDepth             `json:"depth"`
	Platforms []ScreeningRequestPlatform `json:"platforms"`
}

func (s *ScreeningRequest) GetDepth() ScreeningDepth {
	if s == nil {
		return ScreeningDepth("")
	}
	return s.Depth
}

func (s *ScreeningRequest) GetPlatforms() []ScreeningRequestPlatform {
	if s == nil {
		return []ScreeningRequestPlatform{}
	}
	return s.Platforms
}
