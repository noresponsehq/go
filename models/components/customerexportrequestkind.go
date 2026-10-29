package components

type CustomerExportRequestKind string

const (
	CustomerExportRequestKindAccount CustomerExportRequestKind = "account"
	CustomerExportRequestKindAudit   CustomerExportRequestKind = "audit"
)

func (e CustomerExportRequestKind) ToPointer() *CustomerExportRequestKind {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *CustomerExportRequestKind) IsExact() bool {
	if e != nil {
		switch *e {
		case "account", "audit":
			return true
		}
	}
	return false
}
