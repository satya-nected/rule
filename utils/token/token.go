package token

type Token int

const (
	ILLEGAL_TOKEN Token = iota

	nodeTypeBegin
	GROUP
	PARAMS
	CONDITION
	SQLCONDITION
	CONSTANT
	nodeTypeEnd

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

var tokenList = [...]string{
	ILLEGAL_TOKEN: "ILLEGAL",

	GROUP:        "group",
	PARAMS:       "params",
	CONDITION:    "condition",
	SQLCONDITION: "sqlCondition",
	CONSTANT:     "constant",

	AND: "and",
	OR:  "or",

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

func NewToken(input string) Token {
	switch input {
	case GROUP.String():
		return GROUP
	case PARAMS.String():
		return PARAMS
	case CONDITION.String():
		return CONDITION
	case SQLCONDITION.String():
		return SQLCONDITION
	case CONSTANT.String():
		return CONSTANT
	case AND.String():
		return AND
	case OR.String():
		return OR
	case EXISTS.String():
		return EXISTS
	case NEXISTS.String():
		return NEXISTS
	case ISNULL.String():
		return ISNULL
	case NNULL.String():
		return NNULL
	case BETWEEN.String():
		return BETWEEN
	case NBETWEEN.String():
		return NBETWEEN
	case CONTAINS.String():
		return CONTAINS
	case NCONTAINS.String():
		return NCONTAINS
	case STARTWITH.String():
		return STARTWITH
	case NSTARTWITH.String():
		return NSTARTWITH
	case ENDWITH.String():
		return ENDWITH
	case NENDWITH.String():
		return NENDWITH
	case EVEN.String():
		return EVEN
	case ODD.String():
		return ODD
	case EQ.String():
		return EQ
	case NEQ.String():
		return NEQ
	case GT.String():
		return GT
	case LT.String():
		return LT
	case GTE.String():
		return GTE
	case LTE.String():
		return LTE
	case TRUE.String():
		return TRUE
	case FALSE.String():
		return FALSE
	}
	return ILLEGAL_TOKEN
}

// String returns the string representation of the token.
func (tok Token) String() string {
	if tok >= 0 && tok < Token(len(tokenList)) {
		return tokenList[tok]
	}
	return ""
}

// IsNodeType returns true for node tokens.
func (tok Token) IsNodeType() bool {
	return tok > nodeTypeBegin && tok < nodeTypeEnd
}

// IsGroupOperator returns true for group operator tokens.
func (tok Token) IsGroupOperator() bool {
	return tok > groupOperatorBegin && tok < groupOperatorEnd
}

// IsConditionOperator returns true for condition operator tokens.
func (tok Token) IsConditionOperator() bool {
	return tok > conditionOperatorBegin && tok < conditionOperatorEnd
}
