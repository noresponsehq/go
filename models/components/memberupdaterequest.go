package components

import (
	"encoding/json"
	"fmt"
	"github.com/noresponsehq/go/v9/internal/utils"
)

type MemberUpdateRequestEnvironment string

const (
	MemberUpdateRequestEnvironmentLive MemberUpdateRequestEnvironment = "live"
	MemberUpdateRequestEnvironmentTest MemberUpdateRequestEnvironment = "test"
)

func (e MemberUpdateRequestEnvironment) ToPointer() *MemberUpdateRequestEnvironment {
	return &e
}
func (e *MemberUpdateRequestEnvironment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "live":
		fallthrough
	case "test":
		*e = MemberUpdateRequestEnvironment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MemberUpdateRequestEnvironment: %v", v)
	}
}

type MemberUpdateRequestDeniedAction string

const (
	MemberUpdateRequestDeniedActionEvidenceRead  MemberUpdateRequestDeniedAction = "evidence.read"
	MemberUpdateRequestDeniedActionExportsRead   MemberUpdateRequestDeniedAction = "exports.read"
	MemberUpdateRequestDeniedActionPolicyPublish MemberUpdateRequestDeniedAction = "policy.publish"
)

func (e MemberUpdateRequestDeniedAction) ToPointer() *MemberUpdateRequestDeniedAction {
	return &e
}
func (e *MemberUpdateRequestDeniedAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "evidence.read":
		fallthrough
	case "exports.read":
		fallthrough
	case "policy.publish":
		*e = MemberUpdateRequestDeniedAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MemberUpdateRequestDeniedAction: %v", v)
	}
}

// MemberUpdateRequest - PATCH /v1/workspaces/{workspace_id}/members/{member_id} body: the member's new role. The workspace's last owner cannot be demoted.
type MemberUpdateRequest struct {
	// A built-in capability set. Specialist roles are independent; effective access is returned by GET /v1/workspaces/{workspace_id}/capabilities.
	Role Role `json:"role"`
	// Allowed data environments. Invites default to both. Omitted on member PATCH preserves the current grants; owners retain both.
	Environments []MemberUpdateRequestEnvironment `json:"environments,omitzero"`
	// Independent removals from a built-in preset: sensitive evidence, exports, or active policy publication. Owners retain full authority. Omitted member updates preserve exclusions.
	DeniedActions []MemberUpdateRequestDeniedAction `json:"denied_actions,omitzero"`
}

func (m MemberUpdateRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MemberUpdateRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (m *MemberUpdateRequest) GetRole() Role {
	if m == nil {
		return Role("")
	}
	return m.Role
}

func (m *MemberUpdateRequest) GetEnvironments() []MemberUpdateRequestEnvironment {
	if m == nil {
		return nil
	}
	return m.Environments
}

func (m *MemberUpdateRequest) GetDeniedActions() []MemberUpdateRequestDeniedAction {
	if m == nil {
		return nil
	}
	return m.DeniedActions
}
