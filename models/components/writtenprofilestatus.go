package components

// WrittenProfileStatus - queued: the profile is being written. written: summary, statements and leads hold it. failed: it could not be written; error_code says why and no content is returned.
type WrittenProfileStatus string

const (
	WrittenProfileStatusQueued  WrittenProfileStatus = "queued"
	WrittenProfileStatusWritten WrittenProfileStatus = "written"
	WrittenProfileStatusFailed  WrittenProfileStatus = "failed"
)

func (e WrittenProfileStatus) ToPointer() *WrittenProfileStatus {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *WrittenProfileStatus) IsExact() bool {
	if e != nil {
		switch *e {
		case "queued", "written", "failed":
			return true
		}
	}
	return false
}
