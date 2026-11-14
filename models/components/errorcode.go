package components

// ErrorCode - Why a failed search found nothing: its sources did not answer, none could be asked, or it could not run before its attempts were spent.
type ErrorCode string

const (
	ErrorCodeNotAnswered             ErrorCode = "not_answered"
	ErrorCodeNoSource                ErrorCode = "no_source"
	ErrorCodeWorkerAttemptsExhausted ErrorCode = "worker_attempts_exhausted"
)

func (e ErrorCode) ToPointer() *ErrorCode {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ErrorCode) IsExact() bool {
	if e != nil {
		switch *e {
		case "not_answered", "no_source", "worker_attempts_exhausted":
			return true
		}
	}
	return false
}
