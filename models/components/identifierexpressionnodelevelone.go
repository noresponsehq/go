package components

import (
	"github.com/noresponsehq/go/v9/internal/utils"
)

type IdentifierExpressionNodeLevelOne struct {
	// One customer-submitted identifier. Defaults to exact comparison of the original submitted value. canonical uses the supported identifier normalisation without an ownership claim: email preserves local-part case, NFC-normalises it and lowercases the domain without removing dots or tags; phone uses international E.164 without extensions or an inferred region; username uses NFC and Unicode case folding within the customer namespace. mailbox is email-only and explicitly selects existing list rules: case-insensitive, remove +tag, and ignore local-part dots at gmail.com and googlemail.com; googlemail.com maps to gmail.com. Username requires namespace customer. Email and phone forbid namespace. Source-discovered identifiers and platform handles are not searched. Unsupported kinds, namespaces, modes and invalid values return 400. Search equality does not establish verified ownership or merge accounts.
	Term *IdentifierSearch          `json:"term,omitzero"`
	And  []IdentifierExpressionLeaf `json:"and,omitzero"`
	Or   []IdentifierExpressionLeaf `json:"or,omitzero"`
}

func (i IdentifierExpressionNodeLevelOne) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IdentifierExpressionNodeLevelOne) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, nil); err != nil {
		return err
	}
	return nil
}

func (i *IdentifierExpressionNodeLevelOne) GetTerm() *IdentifierSearch {
	if i == nil {
		return nil
	}
	return i.Term
}

func (i *IdentifierExpressionNodeLevelOne) GetAnd() []IdentifierExpressionLeaf {
	if i == nil {
		return nil
	}
	return i.And
}

func (i *IdentifierExpressionNodeLevelOne) GetOr() []IdentifierExpressionLeaf {
	if i == nil {
		return nil
	}
	return i.Or
}
