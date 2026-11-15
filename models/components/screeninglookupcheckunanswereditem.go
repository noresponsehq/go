package components

type ScreeningLookupCheckUnansweredItem struct {
	// The id of the connector that did not answer: breach for the breach search.
	Connector string `json:"connector"`
	// Why a signal or source did not produce a definite found/not_found. timed_out means the source missed the request deadline; a continuation may still settle it. deferred means the (source, identifier) pair is async-only and was not attempted on the synchronous path; a continuation settles it.
	Reason FailureReason `json:"reason"`
}

func (s *ScreeningLookupCheckUnansweredItem) GetConnector() string {
	if s == nil {
		return ""
	}
	return s.Connector
}

func (s *ScreeningLookupCheckUnansweredItem) GetReason() FailureReason {
	if s == nil {
		return FailureReason("")
	}
	return s.Reason
}
