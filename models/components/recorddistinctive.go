package components

// RecordDistinctive - How distinctive a searchable value is: an email, phone, username, password, hash, person name or IP. Absent for other values, and when its rating could not be worked out, such as when a count failed.
type RecordDistinctive struct {
	// unique: an email or phone, which belongs to one person by design. distinctive: rare enough that a new search on it likely finds one person. common: shared by many people, so a search on it finds strangers. A new search is suggested for unique and distinctive values only; any value can still be searched.
	Rating Rating `json:"rating"`
	// Why, in plain words, such as "seen 3,861,493 times in breach data", "in 4 breach rows" or "38 bits, about one in 430 billion usernames".
	Reason string `json:"reason"`
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
