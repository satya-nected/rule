package evalute

import (
	"fmt"
	"math"
	"time"

	"rule/utils/ast"
	"rule/utils/token"
)

func applyUniaryOperator(op token.Token, leftValue ast.Expr) (ast.Expr, error) {
	switch op {
	case token.EVEN:
		return applyEVEN(leftValue)
	case token.ODD:
		return applyODD(leftValue)
	case token.TRUE:
		return applyTRUE(leftValue)
	case token.FALSE:
		return applyFALSE(leftValue)
	}
	return falseExpr, fmt.Errorf("invalid_operator_%v", op)
}

func applyBinaryOperator(op token.Token, leftValue, rightValue ast.Expr) (ast.Expr, error) {
	switch op {
	case token.EQ:
		return applyEQ(leftValue, rightValue)
	case token.NEQ:
		return applyNEQ(leftValue, rightValue)
	case token.GT:
		return applyGT(leftValue, rightValue)
	case token.LT:
		return applyLT(leftValue, rightValue)
	case token.GTE:
		return applyGTE(leftValue, rightValue)
	case token.LTE:
		return applyLTE(leftValue, rightValue)
	case token.CONTAINS:
		return applyCONTAINS(leftValue, rightValue)
	case token.NCONTAINS:
		return applyNCONTAINS(leftValue, rightValue)
	case token.STARTWITH:
		return applySTARTWITH(leftValue, rightValue)
	case token.NSTARTWITH:
		return applyNSTARTWITH(leftValue, rightValue)
	case token.ENDWITH:
		return applyENDWITH(leftValue, rightValue)
	case token.NENDWITH:
		return applyNENDWITH(leftValue, rightValue)
	}
	return falseExpr, fmt.Errorf("invalid_operator_%v", op)
}

func applyTerniaryOperator(op token.Token, leftValue, rightValue, rightValue2 ast.Expr) (ast.Expr, error) {
	switch op {
	case token.BETWEEN:
		return applyBETWEEN(leftValue, rightValue, rightValue2)
	case token.NBETWEEN:
		return applyNBETWEEN(leftValue, rightValue, rightValue2)
	}
	return falseExpr, fmt.Errorf("invalid_operator_%v", op)
}

func applyEQ(leftValue, rightValue ast.Expr) (ast.Expr, error) {
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

func applyNEQ(leftValue, rightValue ast.Expr) (ast.Expr, error) {
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

func applyGT(leftValue, rightValue ast.Expr) (ast.Expr, error) {
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

func applyLT(leftValue, rightValue ast.Expr) (ast.Expr, error) {
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

func applyGTE(leftValue, rightValue ast.Expr) (ast.Expr, error) {
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

func applyLTE(leftValue, rightValue ast.Expr) (ast.Expr, error) {
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

func applyCONTAINS(leftValue, rightValue ast.Expr) (ast.Expr, error) {
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

func applyNCONTAINS(leftValue, rightValue ast.Expr) (ast.Expr, error) {
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

func applySTARTWITH(leftValue, rightValue ast.Expr) (ast.Expr, error) {
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

func applyNSTARTWITH(leftValue, rightValue ast.Expr) (ast.Expr, error) {
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

func applyENDWITH(leftValue, rightValue ast.Expr) (ast.Expr, error) {
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

func applyNENDWITH(leftValue, rightValue ast.Expr) (ast.Expr, error) {
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

func applyBETWEEN(leftValue, rightValue, rightValue2 ast.Expr) (*ast.BooleanLiteral, error) {
	var (
		an, bn, cn float64
		err, err2  error
	)
	an, err = getNumber(leftValue)
	if err == nil {
		bn, err2 = getNumber(rightValue)
		if err2 == nil {
			cn, err = getNumber(rightValue2)
			if err != nil {
				return &ast.BooleanLiteral{Val: (an >= bn && an <= cn)}, nil
			}
		}
		return falseExpr, fmt.Errorf("error_cmp_num_nnum")
	}
	return falseExpr, fmt.Errorf("invalid_left_or_rights_operands")
}

func applyNBETWEEN(leftValue, rightValue, rightValue2 ast.Expr) (*ast.BooleanLiteral, error) {
	var (
		an, bn, cn float64
		err, err2  error
	)
	an, err = getNumber(leftValue)
	if err == nil {
		bn, err2 = getNumber(rightValue)
		if err2 == nil {
			cn, err = getNumber(rightValue2)
			if err != nil {
				return &ast.BooleanLiteral{Val: !(an >= bn && an <= cn)}, nil
			}
		}
		return falseExpr, fmt.Errorf("error_cmp_num_nnum")
	}
	return falseExpr, fmt.Errorf("invalid_left_or_rights_operands")
}

func applyEVEN(leftValue ast.Expr) (ast.Expr, error) {
	var (
		an  float64
		err error
	)
	an, err = getNumber(leftValue)
	if err == nil {
		return &ast.BooleanLiteral{Val: (math.Mod(an, 2) == 0)}, nil
	}
	return falseExpr, fmt.Errorf("invalid_left_operands")
}

func applyODD(leftValue ast.Expr) (ast.Expr, error) {
	var (
		an  float64
		err error
	)
	an, err = getNumber(leftValue)
	if err == nil {
		return &ast.BooleanLiteral{Val: (math.Mod(an, 2) != 0)}, nil
	}
	return falseExpr, fmt.Errorf("invalid_left_operands")
}

func applyTRUE(leftVal ast.Expr) (*ast.BooleanLiteral, error) {
	var (
		a, b bool
		err  error
	)
	a, err = getBoolean(leftVal)
	if err == nil {
		return &ast.BooleanLiteral{Val: a}, nil
	}
	return &ast.BooleanLiteral{Val: (a && b)}, nil
}

func applyFALSE(leftVal ast.Expr) (*ast.BooleanLiteral, error) {
	var (
		a, b bool
		err  error
	)
	a, err = getBoolean(leftVal)
	if err == nil {
		return &ast.BooleanLiteral{Val: !a}, nil
	}
	return &ast.BooleanLiteral{Val: (a && b)}, nil
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

func getDatetime(e ast.Expr) (time.Time, error) {
	switch n := e.(type) {
	case *ast.TimeLiteral:
		return n.Val, nil
	default:
		return time.Time{}, fmt.Errorf("literal_not_datetime: %v", n)
	}
}
