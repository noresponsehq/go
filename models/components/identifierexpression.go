package components

// IdentifierExpression - A bounded AND/OR expression over submitted identifiers. Required scope assessment means the entire expression must match one retained assessment. Scope account permits each term to match a different retained assessment of the same account and is accepted only by account search. Exactly one of term, and or or is allowed at every node. Groups contain 2 to 8 operands, with at most 3 nested group levels and 8 raw term occurrences over the entire expression. A root term is allowed. Duplicate terms with the same kind, match rule, namespace and normalized comparison are idempotent, but count toward the raw limit. Match rules default to exact and preserve the existing identifier contract. Only retained submitted identifiers can satisfy a term. Unknown fields, unsupported scopes and malformed or oversized expressions return 400 invalid_request.
type IdentifierExpression struct {
	// assessment requires one submission to satisfy the complete expression. account permits witnesses across retained submissions of one account and is supported only by POST /v1/accounts/search. No default: select scope explicitly.
	Scope      ScopeEnum                `json:"scope"`
	Expression IdentifierExpressionNode `json:"expression"`
}

func (i *IdentifierExpression) GetScope() ScopeEnum {
	if i == nil {
		return ScopeEnum("")
	}
	return i.Scope
}

func (i *IdentifierExpression) GetExpression() IdentifierExpressionNode {
	if i == nil {
		return IdentifierExpressionNode{}
	}
	return i.Expression
}
