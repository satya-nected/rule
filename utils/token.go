package utils

type Token int

const (
	ILLEGAL Token = iota

	GROUP

	groupOperatorBegin
	AND
	OR
	groupOperatorEnd

	conditionOperatorBegin
	EXISTS
	NEXISTS
	ISNULL
	NNULL
	BETWEEN
	NBETWEEN
	CONTAINS
	NCONTAINS
	STARTWITH
	NSTARTWITH
	ENDWITH
	NENDWITH
	EVEN
	ODD
	EQ
	NEQ
	GT
	LT
	GTE
	LTE
	TRUE
	FALSE
	conditionOperatorEnd
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",

	GROUP: "group",
	AND:   "and",
	OR:    "or",

	EXISTS:  "ex",
	NEXISTS: "nex",

	ISNULL: "isNull",
	NNULL:  "notNull",

	BETWEEN:  "bet",
	NBETWEEN: "nbet",

	CONTAINS:  "contains",
	NCONTAINS: "notContains",

	STARTWITH:  "sw",
	NSTARTWITH: "nsw",

	ENDWITH:  "ew",
	NENDWITH: "new",

	EVEN: "even",
	ODD:  "odd",

	EQ:  "eq",
	NEQ: "neq",
	GT:  "gt",
	LT:  "lt",
	GTE: "gte",
	LTE: "lte",

	TRUE:  "t",
	FALSE: "f",
}

// String returns the string representation of the token.
func (tok Token) String() string {
	if tok >= 0 && tok < Token(len(tokens)) {
		return tokens[tok]
	}
	return ""
}

// IsGroupOperator returns true for group operator tokens.
func (tok Token) IsGroupOperator() bool {
	return tok > groupOperatorBegin && tok < groupOperatorEnd
}

// IsConditionOperator returns true for condition operator tokens.
func (tok Token) IsConditionOperator() bool {
	return tok > conditionOperatorBegin && tok < conditionOperatorEnd
}
