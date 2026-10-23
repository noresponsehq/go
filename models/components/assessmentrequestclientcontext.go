package components

// AssessmentRequestClientContext - Optional context observed by the client of the user. The server parses these hints into coarse features, leaves raw user_agent and opaque ids out of every answer, and weighs a mismatch with the other evidence.
type AssessmentRequestClientContext struct {
	// IANA time-zone name such as Europe/London or America/Los_Angeles. Fixed numeric offsets are not accepted.
	Timezone *string `json:"timezone,omitzero"`
	// BCP 47 locale, for example en-AU. Treat as a presentation hint, not a nationality signal.
	Locale *string `json:"locale,omitzero"`
	// Client platform. Closed vocabulary keeps integrations stable.
	Platform *Platform `json:"platform,omitzero"`
	// Customer application version, without build metadata that contains secrets.
	AppVersion *string `json:"app_version,omitzero"`
	// The raw User-Agent header of the request that the user sent, at most 512 characters. It is parsed into coarse features and left out of every answer.
	UserAgent *string `json:"user_agent,omitzero"`
	// Workspace-scoped opaque identifier your application mints and rotates when the account or device changes. Send an id of your own, not a hardware fingerprint.
	DeviceID *string `json:"device_id,omitzero"`
	// Opaque signup session identifier. It is used for workspace-scoped velocity and replay checks only.
	SessionID *string `json:"session_id,omitzero"`
	// Elapsed time from form start to submit. Values below three seconds are a weak automation hint, weighed with the other evidence.
	FormDurationMs *int64 `json:"form_duration_ms,omitzero"`
	// RFC 3339 timestamp observed by the client. Server ingress time remains authoritative; a large clock skew is only a low-confidence feature.
	ClientObservedAt *string `json:"client_observed_at,omitzero"`
}

func (a *AssessmentRequestClientContext) GetTimezone() *string {
	if a == nil {
		return nil
	}
	return a.Timezone
}

func (a *AssessmentRequestClientContext) GetLocale() *string {
	if a == nil {
		return nil
	}
	return a.Locale
}

func (a *AssessmentRequestClientContext) GetPlatform() *Platform {
	if a == nil {
		return nil
	}
	return a.Platform
}

func (a *AssessmentRequestClientContext) GetAppVersion() *string {
	if a == nil {
		return nil
	}
	return a.AppVersion
}

func (a *AssessmentRequestClientContext) GetUserAgent() *string {
	if a == nil {
		return nil
	}
	return a.UserAgent
}

func (a *AssessmentRequestClientContext) GetDeviceID() *string {
	if a == nil {
		return nil
	}
	return a.DeviceID
}

func (a *AssessmentRequestClientContext) GetSessionID() *string {
	if a == nil {
		return nil
	}
	return a.SessionID
}

func (a *AssessmentRequestClientContext) GetFormDurationMs() *int64 {
	if a == nil {
		return nil
	}
	return a.FormDurationMs
}

func (a *AssessmentRequestClientContext) GetClientObservedAt() *string {
	if a == nil {
		return nil
	}
	return a.ClientObservedAt
}
