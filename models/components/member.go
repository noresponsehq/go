package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

type MemberEnvironment string

const (
	MemberEnvironmentLive MemberEnvironment = "live"
	MemberEnvironmentTest MemberEnvironment = "test"
)

func (e MemberEnvironment) ToPointer() *MemberEnvironment {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *MemberEnvironment) IsExact() bool {
	if e != nil {
		switch *e {
		case "live", "test":
			return true
		}
	}
	return false
}

type MemberDeniedAction string

const (
	MemberDeniedActionEvidenceRead  MemberDeniedAction = "evidence.read"
	MemberDeniedActionExportsRead   MemberDeniedAction = "exports.read"
	MemberDeniedActionPolicyPublish MemberDeniedAction = "policy.publish"
)

func (e MemberDeniedAction) ToPointer() *MemberDeniedAction {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *MemberDeniedAction) IsExact() bool {
	if e != nil {
		switch *e {
		case "evidence.read", "exports.read", "policy.publish":
			return true
		}
	}
	return false
}

// Member - One member of a workspace: an external user id holding one role. A removed member loses every permission at once.
type Member struct {
	// Opaque, server-assigned member id: mem_ and a 26-character ULID. It stays the same when a removed member is invited back.
	ID string `json:"id"`
	// The workspace id, which is also its live environment's workspace id; the test environment's is the same id with :test appended. A workspace created through the API is tnt_ and a 26-character ULID.
	WorkspaceID string `json:"workspace_id"`
	// The user's id at the identity provider the console signs people in with: 1 to 256 printable ASCII characters with no space, compared exactly.
	ExternalUserID string `json:"external_user_id"`
	// A built-in capability set. Specialist roles are independent; effective access is returned by GET /v1/workspaces/{workspace_id}/capabilities.
	Role Role `json:"role"`
	// Listed and returned members are active.
	Status    MemberStatus `json:"status"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
	// Allowed data environments. Invites default to both. Omitted on member PATCH preserves the current grants; owners retain both.
	Environments []MemberEnvironment `json:"environments,omitzero"`
	// Independent removals from a built-in preset: sensitive evidence, exports, or active policy publication. Owners retain full authority. Omitted member updates preserve exclusions.
	DeniedActions []MemberDeniedAction `json:"denied_actions,omitzero"`
	// Immutable internal workspace principal, equal to member id, preserved across explicit provider-identity migrations.
	PrincipalID *string `json:"principal_id,omitzero"`
	// Informational display name. Authentication and provider identity use the verified subject binding.
	DisplayName *string `json:"display_name,omitzero"`
	// Unverified contact address for display. Empty clears it; a supplied value must be an address. Identity linking uses the verified provider subject.
	UnverifiedEmail *string `json:"unverified_email,omitzero"`
}

func (m Member) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *Member) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (m *Member) GetID() string {
	if m == nil {
		return ""
	}
	return m.ID
}

func (m *Member) GetWorkspaceID() string {
	if m == nil {
		return ""
	}
	return m.WorkspaceID
}

func (m *Member) GetExternalUserID() string {
	if m == nil {
		return ""
	}
	return m.ExternalUserID
}

func (m *Member) GetRole() Role {
	if m == nil {
		return Role("")
	}
	return m.Role
}

func (m *Member) GetStatus() MemberStatus {
	if m == nil {
		return MemberStatus("")
	}
	return m.Status
}

func (m *Member) GetCreatedAt() string {
	if m == nil {
		return ""
	}
	return m.CreatedAt
}

func (m *Member) GetUpdatedAt() string {
	if m == nil {
		return ""
	}
	return m.UpdatedAt
}

func (m *Member) GetEnvironments() []MemberEnvironment {
	if m == nil {
		return nil
	}
	return m.Environments
}

func (m *Member) GetDeniedActions() []MemberDeniedAction {
	if m == nil {
		return nil
	}
	return m.DeniedActions
}

func (m *Member) GetPrincipalID() *string {
	if m == nil {
		return nil
	}
	return m.PrincipalID
}

func (m *Member) GetDisplayName() *string {
	if m == nil {
		return nil
	}
	return m.DisplayName
}

func (m *Member) GetUnverifiedEmail() *string {
	if m == nil {
		return nil
	}
	return m.UnverifiedEmail
}
