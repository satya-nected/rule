package evalute

import (
	"fmt"
	"reflect"

	"rule/logger"
	"rule/utils/ast"
	"rule/utils/token"
)

var (
	falseExpr = &ast.BooleanLiteral{Val: false}
)

func Evaluate(expr ast.Expr, args map[string]map[string]interface{}) (bool, error) {
	if expr == nil {
		return false, fmt.Errorf("expr_is_nil")
	}

	result, err := evaluateSubtree(expr, args)
	if err != nil {
		return false, err
	}
	switch n := result.(type) {
	case *ast.BooleanLiteral:
		return n.Val, nil
	}
	return false, fmt.Errorf("unexpected_result_root_expression: %#v", result)
}

func evaluateSubtree(expr ast.Expr, args map[string]map[string]interface{}) (ast.Expr, error) {
	if expr == nil {
		return falseExpr, fmt.Errorf("expr_is_nil")
	}

	switch expr := expr.(type) {
	case *ast.GroupExpr:
		return evaluteGroupExpr(expr, args)
	case *ast.UniaryExpr:
		return evaluateUniaryExpr(expr, args)
	case *ast.BinaryExpr:
		return evaluateBinaryExpr(expr, args)
	case *ast.TerniaryExpr:
		return evaluateTerniaryExpr(expr, args)
	case *ast.VarRef:
		return evaluateVarRef(expr, args)
	case *ast.NumberLiteral, *ast.StringLiteral, *ast.BooleanLiteral, *ast.TimeLiteral:
		return expr, nil
	}
	return falseExpr, fmt.Errorf("unsupported_expr_%v", expr)
}

func evaluteGroupExpr(expr *ast.GroupExpr, args map[string]map[string]interface{}) (ast.Expr, error) {
	logger.Debugf("group_exp %v", expr)
	ans := falseExpr
	if expr.Op == token.AND {
		ans.Val = true
	} else if expr.Op == token.OR {
		ans.Val = false
	}
	for _, exp := range expr.Children {
		expRes, err := evaluateSubtree(exp, args)
		if err != nil {
			return falseExpr, err
		}
		if expr.Op == token.AND {
			ans, err = applyAND(ans, expRes)
			if err != nil {
				return falseExpr, err
			}
			if !ans.Val {
				return ans, nil
			}
		} else if expr.Op == token.OR {
			ans, err = applyOR(ans, expRes)
			if err != nil {
				return falseExpr, err
			}
			if ans.Val {
				return ans, nil
			}
		} else {
			return falseExpr, fmt.Errorf("invalid_operator_in_group_%v", expr.Id)
		}
	}
	return ans, nil
}

func evaluateUniaryExpr(expr *ast.UniaryExpr, args map[string]map[string]interface{}) (ast.Expr, error) {
	logger.Debugf("uniary_exp_called: %v", expr)
	lv, err := evaluateSubtree(expr.LHS, args)
	if err != nil {
		return falseExpr, err
	}
	return applyUniaryOperator(expr.Op, lv)
}

func evaluateBinaryExpr(expr *ast.BinaryExpr, args map[string]map[string]interface{}) (ast.Expr, error) {
	logger.Debugf("binary_exp_called: %v", expr)
	lv, err := evaluateSubtree(expr.LHS, args)
	if err != nil {
		return falseExpr, err
	}
	rv, err := evaluateSubtree(expr.RHS, args)
	if err != nil {
		return falseExpr, err
	}
	return applyBinaryOperator(expr.Op, lv, rv)
}

func evaluateTerniaryExpr(expr *ast.TerniaryExpr, args map[string]map[string]interface{}) (ast.Expr, error) {
	logger.Debugf("terniary_exp_called: %v", expr)
	lv, err := evaluateSubtree(expr.LHS, args)
	if err != nil {
		return falseExpr, err
	}
	rv, err := evaluateSubtree(expr.RHS, args)
	if err != nil {
		return falseExpr, err
	}
	rv2, err := evaluateSubtree(expr.RHS2, args)
	if err != nil {
		return falseExpr, err
	}
	return applyTerniaryOperator(expr.Op, lv, rv, rv2)
}

func evaluateVarRef(expr *ast.VarRef, args map[string]map[string]interface{}) (ast.Expr, error) {
	logger.Debugf("val_exp %v", expr)
	if _, ok := args[expr.Source]; !ok {
		return falseExpr, fmt.Errorf("source_data_not_found_%v_%v", expr.Source, expr.Var)
	}
	if _, ok := args[expr.Source][expr.Var]; !ok {
		return falseExpr, fmt.Errorf("source_attribute_data_not_found_%v_%v", expr.Source, expr.Var)
	}
	val := args[expr.Source][expr.Var]
	kind := reflect.TypeOf(val).Kind()
	switch kind {
	case reflect.Int:
		return &ast.NumberLiteral{Val: float64(val.(int))}, nil
	case reflect.Int32:
		return &ast.NumberLiteral{Val: float64(val.(int32))}, nil
	case reflect.Int64:
		return &ast.NumberLiteral{Val: float64(val.(int64))}, nil
	case reflect.Float32:
		return &ast.NumberLiteral{Val: float64(val.(float32))}, nil
	case reflect.Float64:
		return &ast.NumberLiteral{Val: float64(val.(float64))}, nil
	case reflect.String:
		return &ast.StringLiteral{Val: val.(string)}, nil
	case reflect.Bool:
		return &ast.BooleanLiteral{Val: val.(bool)}, nil
	}
	return falseExpr, fmt.Errorf("unsupported_argument_%v_%v_type: %v", expr.Source, expr.Var, kind)
}
