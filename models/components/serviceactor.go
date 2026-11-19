package components

import (
	"github.com/noresponsehq/go/v17/internal/utils"
)

type ServiceActor struct {
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	type_ string `const:"service" json:"type"`
	// The NoResponse service that made the change with no person acting, such as replenisher, which replaces dead platform sessions in NoResponse's own workspace.
	Name string `json:"name"`
}

func (s ServiceActor) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ServiceActor) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (s *ServiceActor) GetType() string {
	return "service"
}

func (s *ServiceActor) GetName() string {
	if s == nil {
		return ""
	}
	return s.Name
}
