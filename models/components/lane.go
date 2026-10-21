package components

// Lane - An evidence lane: accounts (account presence for email and username), phone, breach and ip.
type Lane string

const (
	LaneAccounts Lane = "accounts"
	LanePhone    Lane = "phone"
	LaneBreach   Lane = "breach"
	LaneIP       Lane = "ip"
)

func (e Lane) ToPointer() *Lane {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Lane) IsExact() bool {
	if e != nil {
		switch *e {
		case "accounts", "phone", "breach", "ip":
			return true
		}
	}
	return false
}
