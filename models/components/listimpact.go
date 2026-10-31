package components

import (
	"github.com/noresponsehq/go/v10/internal/utils"
)

// ListImpact - List override changes among the newest 5000 matching accounts. These counts describe override inputs. Customer enforcement and eventual outcomes remain separate. The response reports whether the matching population was truncated.
type ListImpact struct {
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	scope            string         `const:"current_list_override_inputs" json:"scope"`
	Kind             ListImpactKind `json:"kind"`
	Value            string         `json:"value"`
	Action           Action         `json:"action"`
	Remove           *string        `json:"remove,omitzero"`
	MatchedAccounts  int64          `json:"matched_accounts"`
	ChangedOverrides int64          `json:"changed_overrides"`
	Truncated        bool           `json:"truncated"`
	//lint:ignore U1000 accessed via reflection for JSON marshaling
	historicalPopulationKnown bool                   `const:"false" json:"historical_population_known"`
	Transitions               []ListImpactTransition `json:"transitions"`
}

func (l ListImpact) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListImpact) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (l *ListImpact) GetScope() string {
	return "current_list_override_inputs"
}

func (l *ListImpact) GetKind() ListImpactKind {
	if l == nil {
		return ListImpactKind("")
	}
	return l.Kind
}

func (l *ListImpact) GetValue() string {
	if l == nil {
		return ""
	}
	return l.Value
}

func (l *ListImpact) GetAction() Action {
	if l == nil {
		return Action("")
	}
	return l.Action
}

func (l *ListImpact) GetRemove() *string {
	if l == nil {
		return nil
	}
	return l.Remove
}

func (l *ListImpact) GetMatchedAccounts() int64 {
	if l == nil {
		return 0
	}
	return l.MatchedAccounts
}

func (l *ListImpact) GetChangedOverrides() int64 {
	if l == nil {
		return 0
	}
	return l.ChangedOverrides
}

func (l *ListImpact) GetTruncated() bool {
	if l == nil {
		return false
	}
	return l.Truncated
}

func (l *ListImpact) GetHistoricalPopulationKnown() bool {
	return false
}

func (l *ListImpact) GetTransitions() []ListImpactTransition {
	if l == nil {
		return []ListImpactTransition{}
	}
	return l.Transitions
}
