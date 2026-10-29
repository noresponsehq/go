package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

type PolicyPreviewRequestFlows struct {
	AccountCreation *PolicyPreviewRequestFlowsAccountCreation `json:"account_creation,omitzero"`
	Login           *PolicyPreviewRequestFlowsLogin           `json:"login,omitzero"`
	Purchase        *PolicyPreviewRequestFlowsPurchase        `json:"purchase,omitzero"`
}

func (p PolicyPreviewRequestFlows) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PolicyPreviewRequestFlows) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (p *PolicyPreviewRequestFlows) GetAccountCreation() *PolicyPreviewRequestFlowsAccountCreation {
	if p == nil {
		return nil
	}
	return p.AccountCreation
}

func (p *PolicyPreviewRequestFlows) GetLogin() *PolicyPreviewRequestFlowsLogin {
	if p == nil {
		return nil
	}
	return p.Login
}

func (p *PolicyPreviewRequestFlows) GetPurchase() *PolicyPreviewRequestFlowsPurchase {
	if p == nil {
		return nil
	}
	return p.Purchase
}
