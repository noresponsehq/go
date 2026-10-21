package components

type RiskReason struct {
	// A stable reason code. New codes may appear; ignore any you do not know.
	Code   ReasonCode `json:"code"`
	Weight int64      `json:"weight"`
	// The submitted identifier or client context field a reason concerns; absent when it concerns the whole assessment. Client context is a claimed hint, not a device fingerprint.
	Identifier *ReasonIdentifier `json:"identifier,omitzero"`
}

func (r *RiskReason) GetCode() ReasonCode {
	if r == nil {
		return ReasonCode("")
	}
	return r.Code
}

func (r *RiskReason) GetWeight() int64 {
	if r == nil {
		return 0
	}
	return r.Weight
}

func (r *RiskReason) GetIdentifier() *ReasonIdentifier {
	if r == nil {
		return nil
	}
	return r.Identifier
}
