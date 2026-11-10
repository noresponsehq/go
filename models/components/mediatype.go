package components

type MediaType string

const (
	MediaTypeImage MediaType = "image"
	MediaTypeVideo MediaType = "video"
)

func (e MediaType) ToPointer() *MediaType {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *MediaType) IsExact() bool {
	if e != nil {
		switch *e {
		case "image", "video":
			return true
		}
	}
	return false
}
