package components

// WrittenProfileErrorCode - Present only when status is failed. no_writer: no profile writer serves this environment. refused: the model refused. not_written: the answer was not a profile, or no statement survived grounding. uncited_statement: a statement or lead cited a record the person does not hold. not_answered: the writer did not answer. worker_attempts_exhausted: the worker stopped retrying.
type WrittenProfileErrorCode string

const (
	WrittenProfileErrorCodeNoWriter                WrittenProfileErrorCode = "no_writer"
	WrittenProfileErrorCodeRefused                 WrittenProfileErrorCode = "refused"
	WrittenProfileErrorCodeNotWritten              WrittenProfileErrorCode = "not_written"
	WrittenProfileErrorCodeUncitedStatement        WrittenProfileErrorCode = "uncited_statement"
	WrittenProfileErrorCodeNotAnswered             WrittenProfileErrorCode = "not_answered"
	WrittenProfileErrorCodeWorkerAttemptsExhausted WrittenProfileErrorCode = "worker_attempts_exhausted"
)

func (e WrittenProfileErrorCode) ToPointer() *WrittenProfileErrorCode {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *WrittenProfileErrorCode) IsExact() bool {
	if e != nil {
		switch *e {
		case "no_writer", "refused", "not_written", "uncited_statement", "not_answered", "worker_attempts_exhausted":
			return true
		}
	}
	return false
}
