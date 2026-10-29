package components

// CustomerExportChunk - Download one authorized export chunk.
type CustomerExportChunk struct {
	ExportID   string `json:"export_id"`
	Index      int64  `json:"index"`
	Generation int64  `json:"generation"`
	Sha256     string `json:"sha256"`
	// Base64-encoded UTF-8 JSON Lines, at most 256 KiB decoded. Each line is a typed customer export record. Decode and concatenate chunks in index order within the same generation.
	Data      string `json:"data"`
	NextChunk *int64 `json:"next_chunk,omitzero"`
}

func (c *CustomerExportChunk) GetExportID() string {
	if c == nil {
		return ""
	}
	return c.ExportID
}

func (c *CustomerExportChunk) GetIndex() int64 {
	if c == nil {
		return 0
	}
	return c.Index
}

func (c *CustomerExportChunk) GetGeneration() int64 {
	if c == nil {
		return 0
	}
	return c.Generation
}

func (c *CustomerExportChunk) GetSha256() string {
	if c == nil {
		return ""
	}
	return c.Sha256
}

func (c *CustomerExportChunk) GetData() string {
	if c == nil {
		return ""
	}
	return c.Data
}

func (c *CustomerExportChunk) GetNextChunk() *int64 {
	if c == nil {
		return nil
	}
	return c.NextChunk
}
