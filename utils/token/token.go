package token

type Token int

const (
	ILLEGAL_TOKEN Token = iota

	// Literals
	literalBegin
	NUMBER
	STRING
	DATETIME
	literalEnd

	// nodeType operator
	nodeTypeBegin
	GROUP
	PARAMS
	CONDITION
	SQLCONDITION
	CONSTANT
	nodeTypeEnd

	// group operator
	groupOperatorBegin
	AND
	OR
	groupOperatorEnd

	// conditional operator begin
	conditionOperatorBegin
	// uniaryOperator
	uniaryOperatorBegin
	EXISTS
	NEXISTS
	ISNULL
	NNULL
	EVEN
	ODD
	TRUE
	FALSE
	uniaryOperatorEnd

	// binaryOperator
	binaryOperatorBegin
	EQ
	NEQ
	GT
	LT
	GTE
	LTE
	CONTAINS
	NCONTAINS
	STARTWITH
	NSTARTWITH
	ENDWITH
	NENDWITH
	binaryOperatorEnd

	// terniaryOperator
	terniaryOperatorBegin
	BETWEEN
	NBETWEEN
	terniaryOperatorEnd
	conditionOperatorEnd
)

var tokenList = [...]string{
	ILLEGAL_TOKEN: "ILLEGAL",

	NUMBER:   "number",
	STRING:   "string",
	DATETIME: "dateTime",

	GROUP:        "group",
	PARAMS:       "params",
	CONDITION:    "condition",
	SQLCONDITION: "sqlCondition",
	CONSTANT:     "constant",

	AND: "and",
	OR:  "or",

	EXISTS:  "ex",
	NEXISTS: "nex",
	ISNULL:  "isNull",
	NNULL:   "notNull",
	EVEN:    "even",
	ODD:     "odd",
	TRUE:    "t",
	FALSE:   "f",

	EQ:         "eq",
	NEQ:        "neq",
	GT:         "gt",
	LT:         "lt",
	GTE:        "gte",
	LTE:        "lte",
	CONTAINS:   "contains",
	NCONTAINS:  "notContains",
	STARTWITH:  "sw",
	NSTARTWITH: "nsw",
	ENDWITH:    "ew",
	NENDWITH:   "new",

	BETWEEN:  "bet",
	NBETWEEN: "nbet",
}

func NewToken(input string) Token {
	switch input {

	case NUMBER.String():
		return NUMBER
	case STRING.String():
		return STRING
	case DATETIME.String():
		return DATETIME

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
	case EVEN.String():
		return EVEN
	case ODD.String():
		return ODD
	case TRUE.String():
		return TRUE
	case FALSE.String():
		return FALSE

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

	case BETWEEN.String():
		return BETWEEN
	case NBETWEEN.String():
		return NBETWEEN
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
func (t Token) IsNodeType() bool {
	return t > nodeTypeBegin && t < nodeTypeEnd
}

// IsGroupOperator returns true for group operator tokens.
func (t Token) IsGroupOperator() bool {
	return t > groupOperatorBegin && t < groupOperatorEnd
}

// IsConditionOperator returns true for condition operator tokens.
func (t Token) IsConditionOperator() bool {
	return t > conditionOperatorBegin && t < conditionOperatorEnd
}

// IsUniaryOperator returns true for uniary operator tokens.
func (t Token) IsUniaryOperator() bool {
	return t > uniaryOperatorBegin && t < uniaryOperatorEnd
}

// IsBinaryOperator returns true for binary operator tokens.
func (t Token) IsBinaryOperator() bool {
	return t > binaryOperatorBegin && t < binaryOperatorEnd
}

// IsTerniaryOperator returns true for terniary operator tokens.
func (t Token) IsTerniaryOperator() bool {
	return t > terniaryOperatorBegin && t < terniaryOperatorEnd
}

// IsLiteral returns true for terniary operator tokens.
func (t Token) IsLiteral() bool {
	return t > literalBegin && t < literalEnd
}
