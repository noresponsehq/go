package components

type AccountRiskReason struct {
	// A stable reason code. New codes may appear; ignore any you do not know.
	Code   ReasonCode `json:"code"`
	Weight int64      `json:"weight"`
	// The submitted identifier or client context field a reason concerns; absent when it concerns the whole assessment. Client context is a claimed hint, not a device fingerprint.
	Identifier *ReasonIdentifier `json:"identifier,omitzero"`
}

func (a *AccountRiskReason) GetCode() ReasonCode {
	if a == nil {
		return ReasonCode("")
	}
	return a.Code
}

func (a *AccountRiskReason) GetWeight() int64 {
	if a == nil {
		return 0
	}
	return a.Weight
}

func (a *AccountRiskReason) GetIdentifier() *ReasonIdentifier {
	if a == nil {
		return nil
	}
	return a.Identifier
}
