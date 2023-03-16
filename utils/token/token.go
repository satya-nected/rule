package token

type Token int
type TokenDetail struct {
	name              string
	leftOperandCount  int
	rightOperandCount int
}

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

var tokenList = [...]TokenDetail{
	ILLEGAL_TOKEN: {"ILLEGAL", 0, 0},

	GROUP:        {"group", -1, -1},
	PARAMS:       {"params", -1, -1},
	CONDITION:    {"condition", -1, -1},
	SQLCONDITION: {"sqlCondition", -1, -1},
	CONSTANT:     {"constant", -1, -1},

	AND: {"and", -1, -1},
	OR:  {"or", -1, -1},

	EXISTS:  {"ex", 1, 0},
	NEXISTS: {"nex", 1, 0},

	ISNULL: {"isNull", 1, 0},
	NNULL:  {"notNull", 1, 0},

	BETWEEN:  {"bet", 1, 2},
	NBETWEEN: {"nbet", 1, 2},

	CONTAINS:  {"contains", 1, 1},
	NCONTAINS: {"notContains", 1, 1},

	STARTWITH:  {"sw", 1, 1},
	NSTARTWITH: {"nsw", 1, 1},

	ENDWITH:  {"ew", 1, 1},
	NENDWITH: {"new", 1, 1},

	EVEN: {"even", 1, 0},
	ODD:  {"odd", 1, 0},

	EQ:  {"eq", 1, 1},
	NEQ: {"neq", 1, 1},
	GT:  {"gt", 1, 1},
	LT:  {"lt", 1, 1},
	GTE: {"gte", 1, 1},
	LTE: {"lte", 1, 1},

	TRUE:  {"t", 1, 0},
	FALSE: {"f", 1, 0},
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
		return tokenList[tok].name
	}
	return ""
}

// GetLeftOperandCount returns count of leftOperands for tokens.
func (t Token) GetLeftOperandCount() int {
	if t >= 0 && t < Token(len(tokenList)) {
		return tokenList[t].leftOperandCount
	}
	return -1
}

// GetRightOperandCount returns count of rightOperands for tokens.
func (t Token) GetRightOperandCount() int {
	if t >= 0 && t < Token(len(tokenList)) {
		return tokenList[t].rightOperandCount
	}
	return -1
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
