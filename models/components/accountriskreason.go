package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// AccountRiskReason - One reason that fired. direction is whether it raises or lowers the score, or is neutral (informational, it moves nothing). impact is a coarse band of how much it moves the score: none for a neutral reason, otherwise low, medium or high.
type AccountRiskReason struct {
	// A stable reason code. New codes may appear; ignore any you do not know.
	Code      ReasonCode `json:"code"`
	Direction Direction  `json:"direction"`
	// The selected reason's served impact band. Requires reason.
	Impact Impact `json:"impact"`
	// The submitted identifier or client context field a reason concerns; absent when it concerns the whole assessment. Client context is a claimed hint, not a device fingerprint.
	Identifier *ReasonIdentifier `json:"identifier,omitzero"`
	// The list entry a list_match reason applied. No other reason carries it.
	ListEntryID *string `json:"list_entry_id,omitzero"`
}

func (a AccountRiskReason) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountRiskReason) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (a *AccountRiskReason) GetCode() ReasonCode {
	if a == nil {
		return ReasonCode("")
	}
	return a.Code
}

func (a *AccountRiskReason) GetDirection() Direction {
	if a == nil {
		return Direction("")
	}
	return a.Direction
}

func (a *AccountRiskReason) GetImpact() Impact {
	if a == nil {
		return Impact("")
	}
	return a.Impact
}

func (a *AccountRiskReason) GetIdentifier() *ReasonIdentifier {
	if a == nil {
		return nil
	}
	return a.Identifier
}

func (a *AccountRiskReason) GetListEntryID() *string {
	if a == nil {
		return nil
	}
	return a.ListEntryID
}
