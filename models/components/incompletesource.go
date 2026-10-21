package components

// IncompleteSource - One evidence lane that did not settle, and why. It names a lane, not a source. A lane that missed its deadline is timed_out rather than not_found.
type IncompleteSource struct {
	// An evidence lane: accounts (account presence for email and username), phone, breach and ip.
	Lane Lane `json:"lane"`
	// Why a lane did not settle. deferred means the lane is asynchronous and still running; a continuation settles it. timed_out means the lane missed its deadline. insufficient_evidence is a settled inconclusive signal rather than an incomplete lane.
	Reason IncompleteReason `json:"reason"`
}

func (i *IncompleteSource) GetLane() Lane {
	if i == nil {
		return Lane("")
	}
	return i.Lane
}

func (i *IncompleteSource) GetReason() IncompleteReason {
	if i == nil {
		return IncompleteReason("")
	}
	return i.Reason
}
