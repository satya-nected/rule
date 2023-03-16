package utils

type OperandDetail struct {
	Name          string
	LeftOperands  int
	RightOperands int
}

var Operator = map[string]string{
	"and": "",
	"or":  "",
}

var NodeType = map[string]string{
	"group":        "",
	"params":       "",
	"condition":    "",
	"sqlCondition": "",
	"constant":     "",
}

var OperatorType = map[string]OperandDetail{
	"ex": {
		Name:          "Exists",
		LeftOperands:  1,
		RightOperands: 0,
	},
	"nex": {
		Name:          "Does Not Exists",
		LeftOperands:  1,
		RightOperands: 0,
	},
	"isNull": {
		Name:          "Is Null",
		LeftOperands:  1,
		RightOperands: 0,
	},
	"notNull": {
		Name:          "Not Null",
		LeftOperands:  1,
		RightOperands: 0,
	},
	"bet": {
		Name:          "Between",
		LeftOperands:  1,
		RightOperands: 2,
	},
	"nbet": {
		Name:          "Not Between",
		LeftOperands:  1,
		RightOperands: 2,
	},
	"eq": {
		Name:          "Equals",
		LeftOperands:  1,
		RightOperands: 1,
	},
	"neq": {
		Name:          "Not Equals",
		LeftOperands:  1,
		RightOperands: 1,
	},
	"gt": {
		Name:          "Greater than",
		LeftOperands:  1,
		RightOperands: 1,
	},
	"lt": {
		Name:          "Less than",
		LeftOperands:  1,
		RightOperands: 1,
	},
	"gte": {
		Name:          "Greater than or equals",
		LeftOperands:  1,
		RightOperands: 1,
	},
	"lte": {
		Name:          "Less than or equals",
		LeftOperands:  1,
		RightOperands: 1,
	},
	"even": {
		Name:          "Is Even",
		LeftOperands:  1,
		RightOperands: 0,
	},
	"odd": {
		Name:          "Is Odd",
		LeftOperands:  1,
		RightOperands: 0,
	},
	"contains": {
		Name:          "Contains",
		LeftOperands:  1,
		RightOperands: 1,
	},
	"notContains": {
		Name:          "Does Not Contains",
		LeftOperands:  1,
		RightOperands: 1,
	},
	"sw": {
		Name:          "Starts With",
		LeftOperands:  1,
		RightOperands: 1,
	},
	"nsw": {
		Name:          "Does Not Starts With",
		LeftOperands:  1,
		RightOperands: 1,
	},
	"ew": {
		Name:          "Ends With",
		LeftOperands:  1,
		RightOperands: 1,
	},
	"new": {
		Name:          "Does Not Ends With",
		LeftOperands:  1,
		RightOperands: 1,
	},
	"t": {
		Name:          "Is True",
		LeftOperands:  1,
		RightOperands: 0,
	},
	"f": {
		Name:          "Is False",
		LeftOperands:  1,
		RightOperands: 0,
	},
}
