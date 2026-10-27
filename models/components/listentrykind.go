package components

// ListEntryKind - email matches one address, whatever the case of its local part. email_domain matches every address of the registrable domain and of its subdomains.
type ListEntryKind string

const (
	ListEntryKindEmail       ListEntryKind = "email"
	ListEntryKindEmailDomain ListEntryKind = "email_domain"
)

func (e ListEntryKind) ToPointer() *ListEntryKind {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *ListEntryKind) IsExact() bool {
	if e != nil {
		switch *e {
		case "email", "email_domain":
			return true
		}
	}
	return false
}
