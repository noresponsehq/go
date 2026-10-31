package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/noresponsehq/go/v10/internal/utils"
)

type ActorType string

const (
	ActorTypeUser    ActorType = "user"
	ActorTypeAPIKey  ActorType = "api_key"
	ActorTypeUnknown ActorType = "UNKNOWN"
)

type Actor struct {
	UserActor  *UserActor      `queryParam:"inline" union:"member"`
	KeyActor   *KeyActor       `queryParam:"inline" union:"member"`
	UnknownRaw json.RawMessage `json:"-" union:"unknown"`

	Type ActorType
}

func CreateActorUser(user UserActor) Actor {
	typ := ActorTypeUser

	return Actor{
		UserActor: &user,
		Type:      typ,
	}
}

func CreateActorAPIKey(apiKey KeyActor) Actor {
	typ := ActorTypeAPIKey

	return Actor{
		KeyActor: &apiKey,
		Type:     typ,
	}
}

func CreateActorUnknown(raw json.RawMessage) Actor {
	return Actor{
		UnknownRaw: raw,
		Type:       ActorTypeUnknown,
	}
}

func (u Actor) GetUnknownRaw() json.RawMessage {
	return u.UnknownRaw
}

func (u Actor) IsUnknown() bool {
	return u.Type == ActorTypeUnknown
}

func (u *Actor) UnmarshalJSON(data []byte) (err error) {
	previous := *u
	*u = Actor{}
	defer func() {
		if err != nil {
			*u = previous
		}
	}()

	type discriminator struct {
		Type string `json:"type"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		var syntaxErr *json.SyntaxError
		if errors.As(err, &syntaxErr) {
			return fmt.Errorf("could not unmarshal discriminator: %w", err)
		}
		u.UnknownRaw = json.RawMessage(data)
		u.Type = ActorTypeUnknown
		return nil
	}
	if dis == nil {
		u.UnknownRaw = json.RawMessage(data)
		u.Type = ActorTypeUnknown
		return nil
	}

	switch dis.Type {
	case "user":
		userActor := new(UserActor)
		if err := utils.UnmarshalJSON(data, &userActor, "", true, nil); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == user) type UserActor within Actor: %w", string(data), err)
		}

		u.UserActor = userActor
		u.Type = ActorTypeUser
		return nil
	case "api_key":
		keyActor := new(KeyActor)
		if err := utils.UnmarshalJSON(data, &keyActor, "", true, nil); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == api_key) type KeyActor within Actor: %w", string(data), err)
		}

		u.KeyActor = keyActor
		u.Type = ActorTypeAPIKey
		return nil
	default:
		u.UnknownRaw = json.RawMessage(data)
		u.Type = ActorTypeUnknown
		return nil
	}

}

func (u Actor) MarshalJSON() ([]byte, error) {
	if u.UserActor != nil {
		return utils.MarshalJSON(u.UserActor, "", true)
	}

	if u.KeyActor != nil {
		return utils.MarshalJSON(u.KeyActor, "", true)
	}

	if u.UnknownRaw != nil {
		return json.RawMessage(u.UnknownRaw), nil
	}
	return nil, errors.New("could not marshal union type Actor: all fields are null")
}

// AuditEvent - One privileged workspace action with its actor, target, outcome and time. Successful data changes commit their audit evidence atomically. Sensitive exports record access before returning bytes. Details exclude secrets and subject identifiers.
type AuditEvent struct {
	// Opaque, server-assigned event id: aud_ and a 26-character ULID.
	ID string `json:"id"`
	// The workspace id, which is also its live environment's workspace id; the test environment's is the same id with :test appended. A workspace created through the API is tnt_ and a 26-character ULID.
	WorkspaceID string                   `json:"workspace_id"`
	Environment AuditSettingsEnvironment `json:"environment"`
	// What happened, as resource.verb: workspace.created, member.role_changed, member.removed, invite.created, invite.revoked, invite.accepted, key.rotated, webhook_endpoint.created, webhook_endpoint.deleted, webhook_endpoint.secret_rotated, integration.created, integration.deleted, integration.secret_rotated, list_entry.created, list_entry.deleted, webhook_endpoint.test_sent, webhook_endpoint.verification_requested, webhook_endpoint.delivery_replayed, assessment.created, assessment.outcome_recorded, assessment.deleted, account.outcome_recorded, account.enforcement_recorded, account.deleted, account.exported. A change with no finer name carries the permission it used, such as keys.live.write. New actions may appear; treat an unknown one as a change to its target.
	Action  string                             `json:"action"`
	Outcome CustomerExportRequestFilterOutcome `json:"outcome"`
	Actor   Actor                              `json:"actor"`
	Target  AuditEventTarget                   `json:"target"`
	// At most 8 short facts about the change, keyed by lowercase names: the role a member was given or accepted (role), the invite an accepted member used (invite_id), the key that replaced a rotated one (new_key_id), the assessment a member's change acted on (assessment_id). Absent when there are none.
	Detail              map[string]string `json:"detail,omitzero"`
	CreatedAt           string            `json:"created_at"`
	AffectedEnvironment *Environment      `json:"affected_environment,omitzero"`
	// Validated request correlation id, also returned in X-Request-ID, linking this event to request diagnostics. Absent on historical events without correlation.
	RequestID *string `json:"request_id,omitzero"`
}

func (a AuditEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuditEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AuditEvent) GetID() string {
	if a == nil {
		return ""
	}
	return a.ID
}

func (a *AuditEvent) GetWorkspaceID() string {
	if a == nil {
		return ""
	}
	return a.WorkspaceID
}

func (a *AuditEvent) GetEnvironment() AuditSettingsEnvironment {
	if a == nil {
		return AuditSettingsEnvironment("")
	}
	return a.Environment
}

func (a *AuditEvent) GetAction() string {
	if a == nil {
		return ""
	}
	return a.Action
}

func (a *AuditEvent) GetOutcome() CustomerExportRequestFilterOutcome {
	if a == nil {
		return CustomerExportRequestFilterOutcome("")
	}
	return a.Outcome
}

func (a *AuditEvent) GetActor() Actor {
	if a == nil {
		return Actor{}
	}
	return a.Actor
}

func (a *AuditEvent) GetActorUser() *UserActor {
	return a.GetActor().UserActor
}

func (a *AuditEvent) GetActorAPIKey() *KeyActor {
	return a.GetActor().KeyActor
}

func (a *AuditEvent) GetTarget() AuditEventTarget {
	if a == nil {
		return AuditEventTarget{}
	}
	return a.Target
}

func (a *AuditEvent) GetDetail() map[string]string {
	if a == nil {
		return nil
	}
	return a.Detail
}

func (a *AuditEvent) GetCreatedAt() string {
	if a == nil {
		return ""
	}
	return a.CreatedAt
}

func (a *AuditEvent) GetAffectedEnvironment() *Environment {
	if a == nil {
		return nil
	}
	return a.AffectedEnvironment
}

func (a *AuditEvent) GetRequestID() *string {
	if a == nil {
		return nil
	}
	return a.RequestID
}
