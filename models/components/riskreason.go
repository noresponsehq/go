package components

import (
	"github.com/noresponsehq/go/v7/internal/utils"
	"github.com/noresponsehq/go/v7/types"
)

// RiskReason - One reason that fired. direction is whether it raises or lowers the score, or is neutral (informational, it moves nothing). impact is a coarse band of how much it moves the score: none for a neutral reason, otherwise low, medium or high.
type RiskReason struct {
	// A stable reason code. New codes may appear; ignore any you do not know.
	Code      ReasonCode `json:"code"`
	Direction Direction  `json:"direction"`
	Impact    Impact     `json:"impact"`
	// The submitted identifier or client context field a reason concerns; absent when it concerns the whole assessment. Client context is a claimed hint, not a device fingerprint.
	Identifier *ReasonIdentifier `json:"identifier,omitzero"`
	// Deprecated and always 0: the policy's weights are not published. It is sent so that SDK versions that require it keep reading reasons. Read direction and impact.
	//
	// Deprecated: This will be removed in a future release, please migrate away from it as soon as possible.
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	weight *int64 `const:"0" json:"weight,omitzero"`
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

func (r *RiskReason) GetWeight() *int64 {
	return types.Pointer[int64](0)
}
