package components

// Bucket - The period each count covers: an hour, or a day or a week (from Monday) in time_zone. Hour buckets are whole UTC hours in any time zone.
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
