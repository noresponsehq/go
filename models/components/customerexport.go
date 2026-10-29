package components

// CustomerExport - Request an asynchronous customer export.
type CustomerExport struct {
	ID          string                    `json:"id"`
	Kind        CustomerExportRequestKind `json:"kind"`
	Environment Environment               `json:"environment"`
	// standard returns the consolidated risk response; detailed adds merged provider-neutral facts when the key has detailed scope. There is no raw mode.
	ResponseMode  ResponseMode         `json:"response_mode"`
	Status        CustomerExportStatus `json:"status"`
	CreatedAt     string               `json:"created_at"`
	SnapshotAt    *string              `json:"snapshot_at,omitzero"`
	ExpiresAt     string               `json:"expires_at"`
	FinishedAt    *string              `json:"finished_at,omitzero"`
	Records       int64                `json:"records"`
	Bytes         int64                `json:"bytes"`
	Chunks        int64                `json:"chunks"`
	Attempts      int64                `json:"attempts"`
	Generation    int64                `json:"generation"`
	ErrorCode     *string              `json:"error_code,omitzero"`
	RequestSha256 string               `json:"request_sha256"`
}

func (c *CustomerExport) GetID() string {
	if c == nil {
		return ""
	}
	return c.ID
}

func (c *CustomerExport) GetKind() CustomerExportRequestKind {
	if c == nil {
		return CustomerExportRequestKind("")
	}
	return c.Kind
}

func (c *CustomerExport) GetEnvironment() Environment {
	if c == nil {
		return Environment("")
	}
	return c.Environment
}

func (c *CustomerExport) GetResponseMode() ResponseMode {
	if c == nil {
		return ResponseMode("")
	}
	return c.ResponseMode
}

func (c *CustomerExport) GetStatus() CustomerExportStatus {
	if c == nil {
		return CustomerExportStatus("")
	}
	return c.Status
}

func (c *CustomerExport) GetCreatedAt() string {
	if c == nil {
		return ""
	}
	return c.CreatedAt
}

func (c *CustomerExport) GetSnapshotAt() *string {
	if c == nil {
		return nil
	}
	return c.SnapshotAt
}

func (c *CustomerExport) GetExpiresAt() string {
	if c == nil {
		return ""
	}
	return c.ExpiresAt
}

func (c *CustomerExport) GetFinishedAt() *string {
	if c == nil {
		return nil
	}
	return c.FinishedAt
}

func (c *CustomerExport) GetRecords() int64 {
	if c == nil {
		return 0
	}
	return c.Records
}

func (c *CustomerExport) GetBytes() int64 {
	if c == nil {
		return 0
	}
	return c.Bytes
}

func (c *CustomerExport) GetChunks() int64 {
	if c == nil {
		return 0
	}
	return c.Chunks
}

func (c *CustomerExport) GetAttempts() int64 {
	if c == nil {
		return 0
	}
	return c.Attempts
}

func (c *CustomerExport) GetGeneration() int64 {
	if c == nil {
		return 0
	}
	return c.Generation
}

func (c *CustomerExport) GetErrorCode() *string {
	if c == nil {
		return nil
	}
	return c.ErrorCode
}

func (c *CustomerExport) GetRequestSha256() string {
	if c == nil {
		return ""
	}
	return c.RequestSha256
}
