package components

type IdentifierSubmission struct {
	// The original value submitted in assessment_id. Under canonical or mailbox matching the two accounts' values can differ.
	Value string `json:"value"`
	// Opaque, server-assigned assessment id: 'asm_' and a 26-character Crockford base32 ULID. Unique within the workspace; another workspace's assessment id answers 404.
	AssessmentID string `json:"assessment_id"`
	// When assessment_id was created.
	SubmittedAt string `json:"submitted_at"`
}

func (i *IdentifierSubmission) GetValue() string {
	if i == nil {
		return ""
	}
	return i.Value
}

func (i *IdentifierSubmission) GetAssessmentID() string {
	if i == nil {
		return ""
	}
	return i.AssessmentID
}

func (i *IdentifierSubmission) GetSubmittedAt() string {
	if i == nil {
		return ""
	}
	return i.SubmittedAt
}
