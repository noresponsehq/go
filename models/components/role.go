package components

// Role - A built-in capability set. Specialist roles are independent; effective access is returned by GET /v1/workspaces/{workspace_id}/capabilities.
type Role string

const (
	RoleOwner        Role = "owner"
	RoleAdmin        Role = "admin"
	RoleDeveloper    Role = "developer"
	RoleViewer       Role = "viewer"
	RoleInvestigator Role = "investigator"
	RoleAuditor      Role = "auditor"
	RoleBilling      Role = "billing"
)

func (e Role) ToPointer() *Role {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *Role) IsExact() bool {
	if e != nil {
		switch *e {
		case "owner", "admin", "developer", "viewer", "investigator", "auditor", "billing":
			return true
		}
	}
	return false
}
