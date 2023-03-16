package ast

import (
	"fmt"
	"strconv"
	"test/utils/token"
	"time"
)

type Node interface {
	node()
	String() string
}

type Expr interface {
	Node
	expr()
	Args() []string
}

// GroupRef represents a reference to a variable.
type GroupRef struct {
	Id       string
	Children []Expr
}

func (*GroupRef) expr() {}
func (*GroupRef) node() {}
func (*GroupRef) Args() []string {
	args := []string{}
	return args
}
func (r *GroupRef) String() string { return "group" }

// VarRef represents a reference to a variable.
type VarRef struct {
	Id  string
	Val string
}

func (*VarRef) expr() {}
func (*VarRef) node() {}
func (*VarRef) Args() []string {
	args := []string{}
	return args
}
func (r *VarRef) String() string { return QuoteIdent(r.Val) }

// NumberLiteral represents a numeric literal.
type NumberLiteral struct {
	Id  string
	Val float64
}

func (*NumberLiteral) expr() {}
func (*NumberLiteral) node() {}
func (*NumberLiteral) Args() []string {
	args := []string{}
	return args
}
func (l *NumberLiteral) String() string { return strconv.FormatFloat(l.Val, 'f', 3, 64) }

// BooleanLiteral represents a boolean literal.
type BooleanLiteral struct {
	Id  string
	Val bool
}

func (*BooleanLiteral) expr() {}
func (*BooleanLiteral) node() {}
func (*BooleanLiteral) Args() []string {
	args := []string{}
	return args
}
func (l *BooleanLiteral) String() string {
	if l.Val {
		return "true"
	}
	return "false"
}

// StringLiteral represents a string literal.
type StringLiteral struct {
	Id  string
	Val string
}

func (*StringLiteral) expr() {}
func (*StringLiteral) node() {}
func (*StringLiteral) Args() []string {
	args := []string{}
	return args
}
func (l *StringLiteral) String() string { return Quote(l.Val) }

// TimeLiteral represents a point-in-time literal.
type TimeLiteral struct {
	Id  string
	Val time.Time
}

func (*TimeLiteral) expr() {}
func (*TimeLiteral) node() {}
func (*TimeLiteral) Args() []string {
	args := []string{}
	return args
}
func (l *TimeLiteral) String() string { return l.Val.UTC().Format("2006-01-02 15:04:05.999") }

// BinaryExpr represents an operation between two expressions.
type BinaryExpr struct {
	Id  string
	Op  token.Token
	LHS Expr
	RHS Expr
}

func (*BinaryExpr) expr() {}
func (*BinaryExpr) node() {}
func (*BinaryExpr) Args() []string {
	args := []string{}
	return args
}
func (e *BinaryExpr) String() string {
	return fmt.Sprintf("%s %s %s", e.LHS.String(), e.Op, e.RHS.String())
}
