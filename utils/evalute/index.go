package evalute

import (
	"fmt"
	"test/utils/ast"
	"test/utils/token"
	"time"
)

func ApplyBinaryOperator(op token.Token, leftValue, rightValue ast.Expr) (ast.Expr, error) {
	switch op {
	case token.EQ:
		return ApplyEQ(leftValue, rightValue)
	}
	return nil, nil
}

func ApplyEQ(leftValue, rightValue ast.Expr) (ast.Expr, error) {
	var (
		as, bs string
		an, bn float64
		ab, bb bool
		err    error
	)
	as, err = getString(leftValue)
	if err == nil {
		bs, err = getString(rightValue)
		if err != nil {
			return falseExpr, fmt.Errorf("error_cmp_str_nstr")
		}
		return &ast.BooleanLiteral{Val: (as == bs)}, nil
	}
	an, err = getNumber(leftValue)
	if err == nil {
		bn, err = getNumber(rightValue)
		if err != nil {
			return falseExpr, fmt.Errorf("error_cmp_num_nnum")
		}
		return &ast.BooleanLiteral{Val: (an == bn)}, nil
	}
	ab, err = getBoolean(leftValue)
	if err == nil {
		bb, err = getBoolean(rightValue)
		if err != nil {
			return falseExpr, fmt.Errorf("error_cmp_bool_nbool")
		}
		return &ast.BooleanLiteral{Val: (ab == bb)}, nil
	}
	return falseExpr, nil
}

func applyAND(leftVal, rightVal ast.Expr) (*ast.BooleanLiteral, error) {
	var (
		a, b bool
		err  error
	)
	a, err = getBoolean(leftVal)
	if err != nil {
		return nil, err
	}
	b, err = getBoolean(rightVal)
	if err != nil {
		return nil, err
	}
	return &ast.BooleanLiteral{Val: (a && b)}, nil
}

func applyOR(leftVal, rightVal ast.Expr) (*ast.BooleanLiteral, error) {
	var (
		a, b bool
		err  error
	)
	a, err = getBoolean(leftVal)
	if err != nil {
		return nil, err
	}
	b, err = getBoolean(rightVal)
	if err != nil {
		return nil, err
	}
	return &ast.BooleanLiteral{Val: (a || b)}, nil
}

// getBoolean performs type assertion and returns boolean value or error
func getBoolean(e ast.Expr) (bool, error) {
	switch n := e.(type) {
	case *ast.BooleanLiteral:
		return n.Val, nil
	default:
		return false, fmt.Errorf("literal_not_boolean: %v", n)
	}
}

// getString performs type assertion and returns string value or error
func getString(e ast.Expr) (string, error) {
	switch n := e.(type) {
	case *ast.StringLiteral:
		return n.Val, nil
	default:
		return "", fmt.Errorf("literal_not_string: %v", n)
	}
}

// getNumber performs type assertion and returns float64 value or error
func getNumber(e ast.Expr) (float64, error) {
	switch n := e.(type) {
	case *ast.NumberLiteral:
		return n.Val, nil
	default:
		return 0, fmt.Errorf("literal_not_number: %v", n)
	}
}

// GetDatetime performs type assertion and returns time value or error
func GetDatetime(e ast.Expr) (time.Time, error) {
	switch n := e.(type) {
	case *ast.TimeLiteral:
		return n.Val, nil
	default:
		return time.Time{}, fmt.Errorf("literal_not_datetime: %v", n)
	}
}
