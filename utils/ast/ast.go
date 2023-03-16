package ast

import (
	"time"
)

type DataType string

const (
	Unknown = DataType("")
	Number  = DataType("number")
	Boolean = DataType("boolean")
	String  = DataType("string")
	Time    = DataType("time")
)

// InspectDataType returns the data type of a given value.
func InspectDataType(v interface{}) DataType {
	switch v.(type) {
	case float64:
		return Number
	case bool:
		return Boolean
	case string:
		return String
	case time.Time:
		return Time
	default:
		return Unknown
	}
}

type Node interface {
	node()
	String() string
}

type Expr interface {
	Node
	expr()
	Args() []string
}
