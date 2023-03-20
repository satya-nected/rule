package parser

import (
	"encoding/json"
	"fmt"
	"time"

	"rule/utils/ast"
	"rule/utils/token"
)

func Parse(data string) (ast.Expr, error) {
	var conditions Conditions
	if err := json.Unmarshal([]byte(data), &conditions); err != nil {
		return nil, err
	}

	visited := make(map[string]bool)
	return parse(conditions.StartNode, &conditions, visited)
}

func parse(nodeId string, conditions *Conditions, visited map[string]bool) (ast.Expr, error) {
	fmt.Println("parsing start for ", nodeId)

	// check if node is visited
	if _, ok := visited[nodeId]; ok {
		return nil, fmt.Errorf("node_already_visted_%v", nodeId)
	}
	var (
		ans        ast.Expr
		ok         bool
		err        error
		nodeDetail NodeDetail
	)

	if nodeDetail, ok = conditions.Nodes[nodeId]; !ok {
		return nil, fmt.Errorf("invalid_nodeId_%v", nodeId)
	}

	_token := token.NewToken(nodeDetail.NodeType)
	if _token.IsNodeType() {
		switch _token {
		case token.GROUP:
			ans, err = parseGroupNode(nodeId, conditions, visited)
		case token.CONDITION:
			ans, err = parseConditionNode(nodeId, conditions, visited)
		case token.PARAMS:
			ans, err = parseParamsNode(nodeId, conditions, visited)
		case token.CONSTANT:
			ans, err = parseConstantNode(nodeId, conditions, visited)
		default:
			ans, err = nil, fmt.Errorf("invalid_nodeType_%v", nodeDetail.NodeType)
		}
	}
	return ans, err
}

func parseGroupNode(nodeId string, conditions *Conditions, visited map[string]bool) (ast.Expr, error) {
	fmt.Println("started parseGroupNode ...", nodeId)

	// make visited nodeId
	visited[nodeId] = true

	var (
		nodeDetail *NodeDetail
		groupExp   ast.GroupExpr
		err        error
	)

	// check valid nodeDetail
	if nodeDetail, err = conditions.ValidNode(nodeId, token.GROUP); err != nil {
		return nil, err
	}

	// if operator is empty consider it and
	if nodeDetail.Operator == "" {
		nodeDetail.Operator = "and"
	}

	// load and check operator
	operatorToken := token.NewToken(nodeDetail.Operator)
	if !operatorToken.IsGroupOperator() {
		return nil, fmt.Errorf("invalid_operator_%v", nodeDetail.Operator)
	}

	// loop over children
	groupExp.Id = nodeId
	groupExp.Op = operatorToken
	for _, childNode := range nodeDetail.Children {
		expr, err := parse(childNode, conditions, visited)
		if err != nil {
			return nil, err
		}
		groupExp.Children = append(groupExp.Children, expr)
	}
	return &groupExp, nil
}

func parseConditionNode(nodeId string, conditions *Conditions, visited map[string]bool) (ast.Expr, error) {
	fmt.Println("started parseConditionNode ...", nodeId)

	// make visited nodeId
	visited[nodeId] = true

	var (
		nodeDetail        *NodeDetail
		err               error
		lexp, rexp, r2exp ast.Expr
	)

	// check valid nodeDetail
	if nodeDetail, err = conditions.ValidNode(nodeId, token.CONDITION); err != nil {
		return nil, err
	}

	// load and check operator
	operatorToken := token.NewToken(nodeDetail.Operator)
	if !operatorToken.IsConditionOperator() {
		return nil, fmt.Errorf("invalid_operator_%v_%v", nodeDetail.Operator, nodeId)
	}

	if operatorToken.IsUniaryOperator() {
		if len(nodeDetail.LeftNode) != 1 {
			return nil, fmt.Errorf("invalid_leftNode_%v", nodeId)
		}
		// parse left & right node
		if lexp, err = parse(nodeDetail.LeftNode[0], conditions, visited); err != nil {
			return nil, err
		}
		return &ast.UniaryExpr{Id: nodeId, LHS: lexp, Op: operatorToken}, nil
	} else if operatorToken.IsBinaryOperator() {
		if len(nodeDetail.LeftNode) != 1 && len(nodeDetail.RightNode) != 1 {
			return nil, fmt.Errorf("invalid_leftNode_or_rightNode_%v", nodeId)
		}

		// parse left & right node
		if lexp, err = parse(nodeDetail.LeftNode[0], conditions, visited); err != nil {
			return nil, err
		}
		if rexp, err = parse(nodeDetail.RightNode[0], conditions, visited); err != nil {
			return nil, err
		}
		return &ast.BinaryExpr{Id: nodeId, LHS: lexp, RHS: rexp, Op: operatorToken}, nil
	} else if operatorToken.IsTerniaryOperator() {
		if len(nodeDetail.LeftNode) != 1 && len(nodeDetail.RightNode) != 2 {
			return nil, fmt.Errorf("invalid_leftNode_or_rightNode_%v", nodeId)
		}

		// parse left & right node
		if lexp, err = parse(nodeDetail.LeftNode[0], conditions, visited); err != nil {
			return nil, err
		}
		if rexp, err = parse(nodeDetail.RightNode[0], conditions, visited); err != nil {
			return nil, err
		}
		if r2exp, err = parse(nodeDetail.RightNode[1], conditions, visited); err != nil {
			return nil, err
		}
		return &ast.TerniaryExpr{Id: nodeId, LHS: lexp, RHS: rexp, RHS2: r2exp, Op: operatorToken}, nil
	}
	return nil, fmt.Errorf("invalid_operator_%v_%v", nodeDetail.Operator, nodeId)
}

func parseConstantNode(nodeId string, conditions *Conditions, visited map[string]bool) (ast.Expr, error) {
	fmt.Println("started parseConstantNode ...", nodeId)

	// make visited nodeId
	visited[nodeId] = true

	var (
		nodeDetail *NodeDetail
		err        error
	)

	// check valid nodeDetail
	if nodeDetail, err = conditions.ValidNode(nodeId, token.CONSTANT); err != nil {
		return nil, err
	}

	datatypeToken := token.NewToken(nodeDetail.Datatype)
	if !datatypeToken.IsLiteral() {
		return nil, fmt.Errorf("invalid_datatype_%v_%v", nodeDetail.Datatype, nodeId)
	}

	switch datatypeToken {
	case token.NUMBER:
		if ast.InspectDataType(nodeDetail.Value) != ast.Number {
			return nil, fmt.Errorf("invalid_value_%v_%v", nodeDetail.Value, nodeId)
		}
		return &ast.NumberLiteral{Id: nodeId, Val: nodeDetail.Value.(float64)}, nil
	case token.STRING:
		if ast.InspectDataType(nodeDetail.Value) != ast.String {
			return nil, fmt.Errorf("invalid_value_%v_%v", nodeDetail.Value, nodeId)
		}
		return &ast.StringLiteral{Id: nodeId, Val: nodeDetail.Value.(string)}, nil
	case token.DATETIME:
		if ast.InspectDataType(nodeDetail.Value) != ast.Time {
			return nil, fmt.Errorf("invalid_value_%v_%v", nodeDetail.Value, nodeId)
		}
		return &ast.TimeLiteral{Id: nodeId, Val: nodeDetail.Value.(time.Time)}, nil
	}
	return nil, fmt.Errorf("invalid_datatype_%v_%v", nodeDetail.Datatype, nodeId)
}

func parseParamsNode(nodeId string, conditions *Conditions, visited map[string]bool) (ast.Expr, error) {
	fmt.Println("started parseParamsNode ...", nodeId)

	// make visited nodeId
	visited[nodeId] = true

	var (
		nodeDetail *NodeDetail
		err        error
	)

	// check valid nodeDetail
	if nodeDetail, err = conditions.ValidNode(nodeId, token.PARAMS); err != nil {
		return nil, err
	}

	if nodeDetail.SourceType == "" || nodeDetail.Attribute == "" {
		return nil, fmt.Errorf("sourceType_and_attribute_required_%v", nodeId)
	}
	return &ast.VarRef{Id: nodeId, Var: nodeDetail.Attribute, Source: nodeDetail.SourceType}, nil
}
