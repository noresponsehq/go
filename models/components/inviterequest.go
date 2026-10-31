package components

import (
	"encoding/json"
	"fmt"
	"github.com/noresponsehq/go/v10/internal/utils"
)

type InviteRequestEnvironment string

const (
	InviteRequestEnvironmentLive InviteRequestEnvironment = "live"
	InviteRequestEnvironmentTest InviteRequestEnvironment = "test"
)

func (e InviteRequestEnvironment) ToPointer() *InviteRequestEnvironment {
	return &e
}
func (e *InviteRequestEnvironment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "live":
		fallthrough
	case "test":
		*e = InviteRequestEnvironment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InviteRequestEnvironment: %v", v)
	}
}

type InviteRequestDeniedAction string

const (
	InviteRequestDeniedActionEvidenceRead  InviteRequestDeniedAction = "evidence.read"
	InviteRequestDeniedActionExportsRead   InviteRequestDeniedAction = "exports.read"
	InviteRequestDeniedActionPolicyPublish InviteRequestDeniedAction = "policy.publish"
)

func (e InviteRequestDeniedAction) ToPointer() *InviteRequestDeniedAction {
	return &e
}
func (e *InviteRequestDeniedAction) UnmarshalJSON(data []byte) error {
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
		*e = InviteRequestDeniedAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InviteRequestDeniedAction: %v", v)
	}
}

// InviteRequest - POST /v1/workspaces/{workspace_id}/invites body. Creates an invite for role and returns its token once. email is where the console sends the token and is not checked against the user who accepts. With external_user_id set, only that user can accept.
type InviteRequest struct {
	// A built-in capability set. Specialist roles are independent; effective access is returned by GET /v1/workspaces/{workspace_id}/capabilities.
	Role Role `json:"role"`
	// Where the console sends the invite. Informational: it is not matched against the accepting user.
	Email *string `json:"email,omitzero"`
	// The user's id at the identity provider the console signs people in with: 1 to 256 printable ASCII characters with no space, compared exactly.
	ExternalUserID *string `json:"external_user_id,omitzero"`
	// Allowed data environments. Invites default to both. Omitted on member PATCH preserves the current grants; owners retain both.
	Environments []InviteRequestEnvironment `json:"environments,omitzero"`
	// Independent removals from a built-in preset: sensitive evidence, exports, or active policy publication. Owners retain full authority. Omitted member updates preserve exclusions.
	DeniedActions []InviteRequestDeniedAction `json:"denied_actions,omitzero"`
}

func (i InviteRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *InviteRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (i *InviteRequest) GetRole() Role {
	if i == nil {
		return Role("")
	}
	return i.Role
}

func (i *InviteRequest) GetEmail() *string {
	if i == nil {
		return nil
	}
	return i.Email
}

func (i *InviteRequest) GetExternalUserID() *string {
	if i == nil {
		return nil
	}
	return i.ExternalUserID
}

func (i *InviteRequest) GetEnvironments() []InviteRequestEnvironment {
	if i == nil {
		return nil
	}
	return i.Environments
}

func (i *InviteRequest) GetDeniedActions() []InviteRequestDeniedAction {
	if i == nil {
		return nil
	}
	return i.DeniedActions
}
