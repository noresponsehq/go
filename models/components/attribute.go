package components

// Attribute - name_gender is the records' gender against the submitted first name's usual one.
type Attribute string

const (
	AttributeBirthDate   Attribute = "birth_date"
	AttributeGender      Attribute = "gender"
	AttributeNationality Attribute = "nationality"
	AttributeNameGender  Attribute = "name_gender"
)

func (e Attribute) ToPointer() *Attribute {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Attribute) IsExact() bool {
	if e != nil {
		switch *e {
		case "birth_date", "gender", "nationality", "name_gender":
			return true
		}
	}
	return false
}
