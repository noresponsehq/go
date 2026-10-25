package components

type SettingsFlows struct {
	// One flow's policy in force. A score below step_up is allow, below review step_up, and from review review; from block, when present, the decision is block. admission says when the flow's decision is final.
	AccountCreation FlowSettings `json:"account_creation"`
	// One flow's policy in force. A score below step_up is allow, below review step_up, and from review review; from block, when present, the decision is block. admission says when the flow's decision is final.
	Login FlowSettings `json:"login"`
	// One flow's policy in force. A score below step_up is allow, below review step_up, and from review review; from block, when present, the decision is block. admission says when the flow's decision is final.
	Purchase FlowSettings `json:"purchase"`
}

func (s *SettingsFlows) GetAccountCreation() FlowSettings {
	if s == nil {
		return FlowSettings{}
	}
	return s.AccountCreation
}

func (s *SettingsFlows) GetLogin() FlowSettings {
	if s == nil {
		return FlowSettings{}
	}
	return s.Login
}

func (s *SettingsFlows) GetPurchase() FlowSettings {
	if s == nil {
		return FlowSettings{}
	}
	return s.Purchase
}
