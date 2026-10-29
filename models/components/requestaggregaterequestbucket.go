package components

type RequestAggregateRequestBucket string

const (
	RequestAggregateRequestBucketHour RequestAggregateRequestBucket = "hour"
	RequestAggregateRequestBucketDay  RequestAggregateRequestBucket = "day"
)

func (e RequestAggregateRequestBucket) ToPointer() *RequestAggregateRequestBucket {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *RequestAggregateRequestBucket) IsExact() bool {
	if e != nil {
		switch *e {
		case "hour", "day":
			return true
		}
	}
	return false
}
