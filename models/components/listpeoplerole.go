package components

import (
	"encoding/json"
	"fmt"
)

type ListPeopleRole string

const (
	ListPeopleRoleGuest    ListPeopleRole = "guest"
	ListPeopleRoleCrew     ListPeopleRole = "crew"
	ListPeopleRoleVendor   ListPeopleRole = "vendor"
	ListPeopleRolePress    ListPeopleRole = "press"
	ListPeopleRoleUnstated ListPeopleRole = "unstated"
)

func (e ListPeopleRole) ToPointer() *ListPeopleRole {
	return &e
}
func (e *ListPeopleRole) UnmarshalJSON(data []byte) error {
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
		*e = ListPeopleRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListPeopleRole: %v", v)
	}
}
