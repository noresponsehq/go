package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

// Outcomes - Accounts whose current completed assessment was created in the selected window and remains retained at as_of. The effective account feedback head is observed at as_of, including explicit unknown reports. Customer-reported labels and unlabelled accounts have separate counts. Superseded labels are excluded.
type Outcomes struct {
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	population  string `const:"retained_current_accounts_by_current_assessment_created_at" json:"population"`
	Denominator int64  `json:"denominator"`
	Reported    int64  `json:"reported"`
	Unreported  int64  `json:"unreported"`
	// Counts only; a zero denominator means no eligible retained records and does not imply a zero rate.
	LabelCoverage Coverage `json:"label_coverage"`
	// Current customer-reported labels, not independently established ground truth. Missing keys mean zero. Unknown is an explicit report; absent or expired feedback is unreported.
	ByOutcome OutcomesByOutcome `json:"by_outcome"`
}

func (o Outcomes) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *Outcomes) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (o *Outcomes) GetPopulation() string {
	return "retained_current_accounts_by_current_assessment_created_at"
}

func (o *Outcomes) GetDenominator() int64 {
	if o == nil {
		return 0
	}
	return o.Denominator
}

func (o *Outcomes) GetReported() int64 {
	if o == nil {
		return 0
	}
	return o.Reported
}

func (o *Outcomes) GetUnreported() int64 {
	if o == nil {
		return 0
	}
	return o.Unreported
}

func (o *Outcomes) GetLabelCoverage() Coverage {
	if o == nil {
		return Coverage{}
	}
	return o.LabelCoverage
}

func (o *Outcomes) GetByOutcome() OutcomesByOutcome {
	if o == nil {
		return OutcomesByOutcome{}
	}
	return o.ByOutcome
}
