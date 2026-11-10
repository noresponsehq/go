package components

import (
	"encoding/json"
	"fmt"
)

type ListAttendeesRole string

const (
	ListAttendeesRoleGuest    ListAttendeesRole = "guest"
	ListAttendeesRoleCrew     ListAttendeesRole = "crew"
	ListAttendeesRoleVendor   ListAttendeesRole = "vendor"
	ListAttendeesRolePress    ListAttendeesRole = "press"
	ListAttendeesRoleUnstated ListAttendeesRole = "unstated"
)

func (e ListAttendeesRole) ToPointer() *ListAttendeesRole {
	return &e
}
func (e *ListAttendeesRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "guest":
		fallthrough
	case "crew":
		fallthrough
	case "vendor":
		fallthrough
	case "press":
		fallthrough
	case "unstated":
		*e = ListAttendeesRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAttendeesRole: %v", v)
	}
}
