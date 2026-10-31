package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

type InviteEnvironment string

const (
	InviteEnvironmentLive InviteEnvironment = "live"
	InviteEnvironmentTest InviteEnvironment = "test"
)

func (e InviteEnvironment) ToPointer() *InviteEnvironment {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *InviteEnvironment) IsExact() bool {
	if e != nil {
		switch *e {
		case "live", "test":
			return true
		}
	}
	return false
}

type InviteDeniedAction string

const (
	InviteDeniedActionEvidenceRead  InviteDeniedAction = "evidence.read"
	InviteDeniedActionExportsRead   InviteDeniedAction = "exports.read"
	InviteDeniedActionPolicyPublish InviteDeniedAction = "policy.publish"
)

func (e InviteDeniedAction) ToPointer() *InviteDeniedAction {
	return &e
}

// IsExact returns true if the value matches a known enum value, false otherwise.
func (e *InviteDeniedAction) IsExact() bool {
	if e != nil {
		switch *e {
		case "evidence.read", "exports.read", "policy.publish":
			return true
		}
	}
	return false
}

// Invite - A pending workspace invitation. Creating or reissuing it returns a token once; lists omit it. The token can be accepted once before expiry and is invalidated on revocation or reissue. Token state does not establish email delivery.
type Invite struct {
	// Opaque, server-assigned invite id: inv_ and a 26-character ULID.
	ID string `json:"id"`
	// The workspace id, which is also its live environment's workspace id; the test environment's is the same id with :test appended. A workspace created through the API is tnt_ and a 26-character ULID.
	WorkspaceID string `json:"workspace_id"`
	// A built-in capability set. Specialist roles are independent; effective access is returned by GET /v1/workspaces/{workspace_id}/capabilities.
	Role Role `json:"role"`
	// Where the console sends the invite. Informational: it is not matched against the accepting user.
	Email *string `json:"email,omitzero"`
	// The user's id at the identity provider the console signs people in with: 1 to 256 printable ASCII characters with no space, compared exactly.
	ExternalUserID *string `json:"external_user_id,omitzero"`
	// The stable member id that issued this invite, or the authenticated live API key prefix for members_manage automation. Reissuing preserves the original creator; the audit event names the reissuing actor.
	CreatedBy string `json:"created_by"`
	CreatedAt string `json:"created_at"`
	ExpiresAt string `json:"expires_at"`
	// The invite token: nrinv_, the invite id's ULID, _ and 43 unpadded base64url characters (32 random bytes). The new member's console sends it to POST /v1/invites/accept.
	Token *string `json:"token,omitzero"`
	// Allowed data environments. Invites default to both. Omitted on member PATCH preserves the current grants; owners retain both.
	Environments []InviteEnvironment `json:"environments,omitzero"`
	// Independent removals from a built-in preset: sensitive evidence, exports, or active policy publication. Owners retain full authority. Omitted member updates preserve exclusions.
	DeniedActions []InviteDeniedAction `json:"denied_actions,omitzero"`
	// Last token reissue, not proof of email delivery.
	ReissuedAt *string `json:"reissued_at,omitzero"`
}

func (i Invite) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *Invite) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (i *Invite) GetID() string {
	if i == nil {
		return ""
	}
	return i.ID
}

func (i *Invite) GetWorkspaceID() string {
	if i == nil {
		return ""
	}
	return i.WorkspaceID
}

func (i *Invite) GetRole() Role {
	if i == nil {
		return Role("")
	}
	return i.Role
}

func (i *Invite) GetEmail() *string {
	if i == nil {
		return nil
	}
	return i.Email
}

func (i *Invite) GetExternalUserID() *string {
	if i == nil {
		return nil
	}
	return i.ExternalUserID
}

func (i *Invite) GetCreatedBy() string {
	if i == nil {
		return ""
	}
	return i.CreatedBy
}

func (i *Invite) GetCreatedAt() string {
	if i == nil {
		return ""
	}
	return i.CreatedAt
}

func (i *Invite) GetExpiresAt() string {
	if i == nil {
		return ""
	}
	return i.ExpiresAt
}

func (i *Invite) GetToken() *string {
	if i == nil {
		return nil
	}
	return i.Token
}

func (i *Invite) GetEnvironments() []InviteEnvironment {
	if i == nil {
		return nil
	}
	return i.Environments
}

func (i *Invite) GetDeniedActions() []InviteDeniedAction {
	if i == nil {
		return nil
	}
	return i.DeniedActions
}

func (i *Invite) GetReissuedAt() *string {
	if i == nil {
		return nil
	}
	return i.ReissuedAt
}
