package components

type ProfileStatement struct {
	Topic Topic  `json:"topic"`
	Text  string `json:"text"`
	// confirmed: a linked record or the searched identifier states it directly. likely: a strong inference. possible: a weak inference.
	Certainty Certainty `json:"certainty"`
	// stated: read from a record. inferred: reasoned from several facts. absence: every source answered and none found it.
	Basis Basis `json:"basis"`
	// Other explanations to keep open for an inference.
	Alternatives []string `json:"alternatives"`
	// Every record, post or check the statement rests on.
	Evidence []ProfileCitation `json:"evidence"`
}

func (p *ProfileStatement) GetTopic() Topic {
	if p == nil {
		return Topic("")
	}
	return p.Topic
}

func (p *ProfileStatement) GetText() string {
	if p == nil {
		return ""
	}
	return p.Text
}

func (p *ProfileStatement) GetCertainty() Certainty {
	if p == nil {
		return Certainty("")
	}
	return p.Certainty
}

func (p *ProfileStatement) GetBasis() Basis {
	if p == nil {
		return Basis("")
	}
	return p.Basis
}

func (p *ProfileStatement) GetAlternatives() []string {
	if p == nil {
		return []string{}
	}
	return p.Alternatives
}

func (p *ProfileStatement) GetEvidence() []ProfileCitation {
	if p == nil {
		return []ProfileCitation{}
	}
	return p.Evidence
}
