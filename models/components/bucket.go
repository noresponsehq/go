package components

// Bucket - The period each count covers, in UTC; a week starts on Monday.
type Bucket string

const (
	BucketHour Bucket = "hour"
	BucketDay  Bucket = "day"
	BucketWeek Bucket = "week"
)

func (e Bucket) ToPointer() *Bucket {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Bucket) IsExact() bool {
	if e != nil {
		switch *e {
		case "hour", "day", "week":
			return true
		}
	}
	return false
}
