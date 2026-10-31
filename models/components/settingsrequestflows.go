package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// SettingsRequestFlows - Use account_creation or its deprecated alias signup, not both.
type SettingsRequestFlows struct {
	// A change to one flow. A field left out keeps the value in force. Thresholds set on login or purchase start from the values in force, account_creation's while it inherits them. After the change, 0 < step_up < review <= 100 and block is off or review <= block <= 100.
	AccountCreation *FlowChange `json:"account_creation,omitzero"`
	// A change to one flow. A field left out keeps the value in force. Thresholds set on login or purchase start from the values in force, account_creation's while it inherits them. After the change, 0 < step_up < review <= 100 and block is off or review <= block <= 100.
	Login *FlowChange `json:"login,omitzero"`
	// A change to one flow. A field left out keeps the value in force. Thresholds set on login or purchase start from the values in force, account_creation's while it inherits them. After the change, 0 < step_up < review <= 100 and block is off or review <= block <= 100.
	Purchase *FlowChange `json:"purchase,omitzero"`
	// A change to one flow. A field left out keeps the value in force. Thresholds set on login or purchase start from the values in force, account_creation's while it inherits them. After the change, 0 < step_up < review <= 100 and block is off or review <= block <= 100.
	Signup *FlowChange `json:"signup,omitzero"`
}

func (s SettingsRequestFlows) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SettingsRequestFlows) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (s *SettingsRequestFlows) GetAccountCreation() *FlowChange {
	if s == nil {
		return nil
	}
	return s.AccountCreation
}

func (s *SettingsRequestFlows) GetLogin() *FlowChange {
	if s == nil {
		return nil
	}
	return s.Login
}

func (s *SettingsRequestFlows) GetPurchase() *FlowChange {
	if s == nil {
		return nil
	}
	return s.Purchase
}

func (s *SettingsRequestFlows) GetSignup() *FlowChange {
	if s == nil {
		return nil
	}
	return s.Signup
}
