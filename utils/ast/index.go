package ast

import (
	"regexp"
	"strings"
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

// Quote returns a quoted string.
func Quote(s string) string {
	return `"` + strings.NewReplacer("\n", `\n`, `\`, `\\`, `"`, `\"`).Replace(s) + `"`
}

func QuoteIdent(s string) string {
	if s == "" || regexp.MustCompile(`[^a-zA-Z_.]`).MatchString(s) {
		return Quote(s)
	}
	return s
}
