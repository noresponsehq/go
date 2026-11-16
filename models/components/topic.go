package components

type Topic string

const (
	TopicIdentity      Topic = "identity"
	TopicLocation      Topic = "location"
	TopicMovement      Topic = "movement"
	TopicTimeline      Topic = "timeline"
	TopicInterests     Topic = "interests"
	TopicRelationships Topic = "relationships"
	TopicWork          Topic = "work"
	TopicEducation     Topic = "education"
	TopicAccounts      Topic = "accounts"
	TopicCredentials   Topic = "credentials"
	TopicDevices       Topic = "devices"
	TopicAbsence       Topic = "absence"
)

func (e Topic) ToPointer() *Topic {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Topic) IsExact() bool {
	if e != nil {
		switch *e {
		case "identity", "location", "movement", "timeline", "interests", "relationships", "work", "education", "accounts", "credentials", "devices", "absence":
			return true
		}
	}
	return false
}
