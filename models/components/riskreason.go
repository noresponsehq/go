package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// RiskReason - One reason that fired. direction is whether it raises or lowers the score, or is neutral (informational, it moves nothing). impact is a coarse band of how much it moves the score: none for a neutral reason, otherwise low, medium or high.
type RiskReason struct {
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

func (r RiskReason) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RiskReason) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (r *RiskReason) GetCode() ReasonCode {
	if r == nil {
		return ReasonCode("")
	}
	return r.Code
}

func (r *RiskReason) GetDirection() Direction {
	if r == nil {
		return Direction("")
	}
	return r.Direction
}

func (r *RiskReason) GetImpact() Impact {
	if r == nil {
		return Impact("")
	}
	return r.Impact
}

func (r *RiskReason) GetIdentifier() *ReasonIdentifier {
	if r == nil {
		return nil
	}
	return r.Identifier
}

func (r *RiskReason) GetListEntryID() *string {
	if r == nil {
		return nil
	}
	return r.ListEntryID
}
