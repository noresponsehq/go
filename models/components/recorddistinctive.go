package components

// RecordDistinctive - How distinctive a searchable value is: an email, phone, username, password, hash, person name or IP. Absent for other values, for a field with no value, and when its rating could not be worked out, such as when a count failed. Linkage weighs a password, hash or username by these numbers: the rarer the value, the more a shared one counts.
type RecordDistinctive struct {
	// unique: an email or phone, which belongs to one person by design. distinctive: rare enough that a new search on it likely finds one person. common: shared by many people, so a search on it finds strangers. A new search is suggested for unique and distinctive values only; any value can still be searched.
	Rating Rating `json:"rating"`
	// Why, in plain words, such as "seen 3,861,493 times in breach data", "in 4 breach rows" or "38 bits, about one in 430 billion usernames".
	Reason string `json:"reason"`
	// How many breach rows hold the value, when they were counted. Set for a username, password, hash or person name.
	Rows *int64 `json:"rows,omitzero"`
	// How many times the Pwned Passwords list holds the password, when it was asked. Set for a password only.
	Pwned *int64 `json:"pwned,omitzero"`
	// How surprising the username is, to one decimal place: a character model spells it about once in 2^bits usernames. Set for a username the model can spell.
	Bits *float64 `json:"bits,omitzero"`
}

func (r *RecordDistinctive) GetRating() Rating {
	if r == nil {
		return Rating("")
	}
	return r.Rating
}

func (r *RecordDistinctive) GetReason() string {
	if r == nil {
		return ""
	}
	return r.Reason
}

func (r *RecordDistinctive) GetRows() *int64 {
	if r == nil {
		return nil
	}
	return r.Rows
}

func (r *RecordDistinctive) GetPwned() *int64 {
	if r == nil {
		return nil
	}
	return r.Pwned
}

func (r *RecordDistinctive) GetBits() *float64 {
	if r == nil {
		return nil
	}
	return r.Bits
}
