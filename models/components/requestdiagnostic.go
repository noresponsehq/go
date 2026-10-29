package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// RequestDiagnostic - One authenticated HTTP attempt. An attempt id is independent of the caller correlation id. No raw body, path parameter, credential or source payload is retained. write_complete means the server wrote/flushed its response, not proof the client processed it. Collection is bounded best effort; ii_signup_request_diagnostic_failures_total exposes gaps. Retained for seven days or the shorter environment retention.
type RequestDiagnostic struct {
	ID                 string  `json:"id"`
	RequestID          string  `json:"request_id"`
	Route              string  `json:"route"`
	HTTPStatus         int64   `json:"http_status"`
	DurationMs         int64   `json:"duration_ms"`
	StartedAt          string  `json:"started_at"`
	CompletedAt        string  `json:"completed_at"`
	ExpiresAt          string  `json:"expires_at"`
	ErrorCode          *string `json:"error_code,omitzero"`
	IdempotentReplayed bool    `json:"idempotent_replayed"`
	WriteComplete      bool    `json:"write_complete"`
	// The decision prepared for one HTTP attempt. Its account field preserves the account decision returned for that attempt, even after subsequent changes.
	Decision *DecisionReceipt `json:"decision,omitzero"`
	// Public authenticated API key identifier. Absent for member requests and older unattributed diagnostics.
	APIKeyID *string `json:"api_key_id,omitzero"`
}

func (r RequestDiagnostic) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RequestDiagnostic) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (r *RequestDiagnostic) GetID() string {
	if r == nil {
		return ""
	}
	return r.ID
}

func (r *RequestDiagnostic) GetRequestID() string {
	if r == nil {
		return ""
	}
	return r.RequestID
}

func (r *RequestDiagnostic) GetRoute() string {
	if r == nil {
		return ""
	}
	return r.Route
}

func (r *RequestDiagnostic) GetHTTPStatus() int64 {
	if r == nil {
		return 0
	}
	return r.HTTPStatus
}

func (r *RequestDiagnostic) GetDurationMs() int64 {
	if r == nil {
		return 0
	}
	return r.DurationMs
}

func (r *RequestDiagnostic) GetStartedAt() string {
	if r == nil {
		return ""
	}
	return r.StartedAt
}

func (r *RequestDiagnostic) GetCompletedAt() string {
	if r == nil {
		return ""
	}
	return r.CompletedAt
}

func (r *RequestDiagnostic) GetExpiresAt() string {
	if r == nil {
		return ""
	}
	return r.ExpiresAt
}

func (r *RequestDiagnostic) GetErrorCode() *string {
	if r == nil {
		return nil
	}
	return r.ErrorCode
}

func (r *RequestDiagnostic) GetIdempotentReplayed() bool {
	if r == nil {
		return false
	}
	return r.IdempotentReplayed
}

func (r *RequestDiagnostic) GetWriteComplete() bool {
	if r == nil {
		return false
	}
	return r.WriteComplete
}

func (r *RequestDiagnostic) GetDecision() *DecisionReceipt {
	if r == nil {
		return nil
	}
	return r.Decision
}

func (r *RequestDiagnostic) GetAPIKeyID() *string {
	if r == nil {
		return nil
	}
	return r.APIKeyID
}
