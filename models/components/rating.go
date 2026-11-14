package components

// Rating - unique: an email or phone, which belongs to one person by design. distinctive: rare enough that a new search on it likely finds one person. common: shared by many people, so a search on it finds strangers. A new search is suggested for unique and distinctive values only; any value can still be searched.
type Rating string

const (
	RatingUnique      Rating = "unique"
	RatingDistinctive Rating = "distinctive"
	RatingCommon      Rating = "common"
)

func (e Rating) ToPointer() *Rating {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Rating) IsExact() bool {
	if e != nil {
		switch *e {
		case "unique", "distinctive", "common":
			return true
		}
	}
	return false
}
