package components

import (
	"encoding/json"
	"fmt"
)

// Platform - Client platform. Closed vocabulary keeps integrations stable.
type Platform string

const (
	PlatformWeb     Platform = "web"
	PlatformIos     Platform = "ios"
	PlatformAndroid Platform = "android"
	PlatformOther   Platform = "other"
)

func (e Platform) ToPointer() *Platform {
	return &e
}
func (e *Platform) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "web":
		fallthrough
	case "ios":
		fallthrough
	case "android":
		fallthrough
	case "other":
		*e = Platform(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Platform: %v", v)
	}
}
