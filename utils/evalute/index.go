package evalute

import (
	"fmt"
	"rule/utils/ast"
	"rule/utils/token"
	"time"
)

func ApplyUniaryOperator(op token.Token, leftValue ast.Expr) (ast.Expr, error) {

	return nil, nil
}

func ApplyBinaryOperator(op token.Token, leftValue, rightValue ast.Expr) (ast.Expr, error) {
	switch op {
	case token.EQ:
		return ApplyEQ(leftValue, rightValue)
	case token.NEQ:
		return ApplyNEQ(leftValue, rightValue)
	case token.GT:
		return ApplyGT(leftValue, rightValue)
	case token.LT:
		return ApplyLT(leftValue, rightValue)
	case token.GTE:
		return ApplyGTE(leftValue, rightValue)
	case token.LTE:
		return ApplyLTE(leftValue, rightValue)
	case token.CONTAINS:
		return ApplyCONTAINS(leftValue, rightValue)
	case token.NCONTAINS:
		return ApplyNCONTAINS(leftValue, rightValue)
	case token.STARTWITH:
		return ApplySTARTWITH(leftValue, rightValue)
	case token.NSTARTWITH:
		return ApplyNSTARTWITH(leftValue, rightValue)
	case token.ENDWITH:
		return ApplyENDWITH(leftValue, rightValue)
	case token.NENDWITH:
		return ApplyNENDWITH(leftValue, rightValue)
	}
	return falseExpr, fmt.Errorf("invalid_operator_%v", op)
}

func ApplyTerniaryOperator(op token.Token, leftValue, rightValue, rightVal2 ast.Expr) (ast.Expr, error) {
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
	return falseExpr, fmt.Errorf("invalid_left_or_right_operands")
}

func ApplyNEQ(leftValue, rightValue ast.Expr) (ast.Expr, error) {
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
		return &ast.BooleanLiteral{Val: (as != bs)}, nil
	}
	an, err = getNumber(leftValue)
	if err == nil {
		bn, err = getNumber(rightValue)
		if err != nil {
			return falseExpr, fmt.Errorf("error_cmp_num_nnum")
		}
		return &ast.BooleanLiteral{Val: (an != bn)}, nil
	}
	ab, err = getBoolean(leftValue)
	if err == nil {
		bb, err = getBoolean(rightValue)
		if err != nil {
			return falseExpr, fmt.Errorf("error_cmp_bool_nbool")
		}
		return &ast.BooleanLiteral{Val: (ab != bb)}, nil
	}
	return falseExpr, fmt.Errorf("invalid_left_or_right_operands")
}

func ApplyGT(leftValue, rightValue ast.Expr) (ast.Expr, error) {
	var (
		an, bn float64
		err    error
	)
	an, err = getNumber(leftValue)
	if err == nil {
		bn, err = getNumber(rightValue)
		if err != nil {
			return falseExpr, fmt.Errorf("error_cmp_num_nnum")
		}
		return &ast.BooleanLiteral{Val: (an > bn)}, nil
	}
	return falseExpr, fmt.Errorf("invalid_left_or_right_operands")
}

func ApplyLT(leftValue, rightValue ast.Expr) (ast.Expr, error) {
	var (
		an, bn float64
		err    error
	)
	an, err = getNumber(leftValue)
	if err == nil {
		bn, err = getNumber(rightValue)
		if err != nil {
			return falseExpr, fmt.Errorf("error_cmp_num_nnum")
		}
		return &ast.BooleanLiteral{Val: (an < bn)}, nil
	}
	return falseExpr, fmt.Errorf("invalid_left_or_right_operands")
}

func ApplyGTE(leftValue, rightValue ast.Expr) (ast.Expr, error) {
	var (
		an, bn float64
		err    error
	)
	an, err = getNumber(leftValue)
	if err == nil {
		bn, err = getNumber(rightValue)
		if err != nil {
			return falseExpr, fmt.Errorf("error_cmp_num_nnum")
		}
		return &ast.BooleanLiteral{Val: (an >= bn)}, nil
	}
	return falseExpr, fmt.Errorf("invalid_left_or_right_operands")
}

func ApplyLTE(leftValue, rightValue ast.Expr) (ast.Expr, error) {
	var (
		an, bn float64
		err    error
	)
	an, err = getNumber(leftValue)
	if err == nil {
		bn, err = getNumber(rightValue)
		if err != nil {
			return falseExpr, fmt.Errorf("error_cmp_num_nnum")
		}
		return &ast.BooleanLiteral{Val: (an <= bn)}, nil
	}
	return falseExpr, fmt.Errorf("invalid_left_or_right_operands")
}

func ApplyCONTAINS(leftValue, rightValue ast.Expr) (ast.Expr, error) {
	var (
		an, bn string
		err    error
	)
	an, err = getString(leftValue)
	if err == nil {
		bn, err = getString(rightValue)
		if err != nil {
			return falseExpr, fmt.Errorf("error_cmp_num_nnum")
		}
		return &ast.BooleanLiteral{Val: CheckContainsString(an, bn)}, nil
	}
	return falseExpr, fmt.Errorf("invalid_left_or_right_operands")
}

func ApplyNCONTAINS(leftValue, rightValue ast.Expr) (ast.Expr, error) {
	var (
		an, bn string
		err    error
	)
	an, err = getString(leftValue)
	if err == nil {
		bn, err = getString(rightValue)
		if err != nil {
			return falseExpr, fmt.Errorf("error_cmp_num_nnum")
		}
		return &ast.BooleanLiteral{Val: !CheckContainsString(an, bn)}, nil
	}
	return falseExpr, fmt.Errorf("invalid_left_or_right_operands")
}

func ApplySTARTWITH(leftValue, rightValue ast.Expr) (ast.Expr, error) {
	var (
		an, bn string
		err    error
	)
	an, err = getString(leftValue)
	if err == nil {
		bn, err = getString(rightValue)
		if err != nil {
			return falseExpr, fmt.Errorf("error_cmp_num_nnum")
		}
		return &ast.BooleanLiteral{Val: CheckStartWithString(an, bn)}, nil
	}
	return falseExpr, fmt.Errorf("invalid_left_or_right_operands")
}

func ApplyNSTARTWITH(leftValue, rightValue ast.Expr) (ast.Expr, error) {
	var (
		an, bn string
		err    error
	)
	an, err = getString(leftValue)
	if err == nil {
		bn, err = getString(rightValue)
		if err != nil {
			return falseExpr, fmt.Errorf("error_cmp_num_nnum")
		}
		return &ast.BooleanLiteral{Val: !CheckStartWithString(an, bn)}, nil
	}
	return falseExpr, fmt.Errorf("invalid_left_or_right_operands")
}

func ApplyENDWITH(leftValue, rightValue ast.Expr) (ast.Expr, error) {
	var (
		an, bn string
		err    error
	)
	an, err = getString(leftValue)
	if err == nil {
		bn, err = getString(rightValue)
		if err != nil {
			return falseExpr, fmt.Errorf("error_cmp_num_nnum")
		}
		return &ast.BooleanLiteral{Val: CheckEndWithString(an, bn)}, nil
	}
	return falseExpr, fmt.Errorf("invalid_left_or_right_operands")
}

func ApplyNENDWITH(leftValue, rightValue ast.Expr) (ast.Expr, error) {
	var (
		an, bn string
		err    error
	)
	an, err = getString(leftValue)
	if err == nil {
		bn, err = getString(rightValue)
		if err != nil {
			return falseExpr, fmt.Errorf("error_cmp_num_nnum")
		}
		return &ast.BooleanLiteral{Val: !CheckEndWithString(an, bn)}, nil
	}
	return falseExpr, fmt.Errorf("invalid_left_or_right_operands")
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

func getBoolean(e ast.Expr) (bool, error) {
	switch n := e.(type) {
	case *ast.BooleanLiteral:
		return n.Val, nil
	default:
		return false, fmt.Errorf("literal_not_boolean: %v", n)
	}
}

func getString(e ast.Expr) (string, error) {
	switch n := e.(type) {
	case *ast.StringLiteral:
		return n.Val, nil
	default:
		return "", fmt.Errorf("literal_not_string: %v", n)
	}
}

func getNumber(e ast.Expr) (float64, error) {
	switch n := e.(type) {
	case *ast.NumberLiteral:
		return n.Val, nil
	default:
		return 0, fmt.Errorf("literal_not_number: %v", n)
	}
}

func GetDatetime(e ast.Expr) (time.Time, error) {
	switch n := e.(type) {
	case *ast.TimeLiteral:
		return n.Val, nil
	default:
		return time.Time{}, fmt.Errorf("literal_not_datetime: %v", n)
	}
}
