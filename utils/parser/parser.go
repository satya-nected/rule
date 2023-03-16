package parser

import (
	"encoding/json"
	"fmt"
	"test/utils/ast"
	"test/utils/token"
)

type Conditions struct {
	StartNode string                `json:"startNode" bson:"startNode,omitempty"`
	Nodes     map[string]NodeDetail `json:"nodes" bson:"nodes,omitempty"`
}

type NodeDetail struct {
	NodeType     string      `json:"nodeType" bson:"nodeType,omitempty"`
	Parent       string      `json:"parent" bson:"parent,omitempty"`
	SiblingIndex int         `json:"siblingIndex" bson:"siblingIndex,omitempty"`
	Name         string      `json:"name" bson:"name,omitempty"`
	Operator     string      `json:"operator" bson:"operator,omitempty"`
	Datatype     string      `json:"datatype" bson:"datatype,omitempty"`
	Children     []string    `json:"children" bson:"children,omitempty"`
	LeftNode     []string    `json:"leftNode" bson:"leftNode,omitempty"`
	RightNode    []string    `json:"rightNode" bson:"rightNode,omitempty"`
	SourceType   string      `json:"sourceType" bson:"sourceType,omitempty"`
	Attribute    string      `json:"attribute" bson:"attribute,omitempty"`
	Query        string      `json:"query" bson:"query,omitempty"`
	Value        interface{} `json:"value" bson:"value,omitempty"`
}

func (c *Conditions) IsValidNode(nodeId string) bool {
	_, ok := c.Nodes[nodeId]
	return ok
}

func Parse(data string) error {
	var conditions Conditions
	json.Unmarshal([]byte(data), &conditions)

	visited := make(map[string]bool)
	fmt.Println(parse(conditions.StartNode, &conditions, visited))
	return nil
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
		nodeDetail NodeDetail
		ok         bool
		groupExp   ast.GroupExpr
	)

	// check valid nodeDetail
	if nodeDetail, ok = conditions.Nodes[nodeId]; !ok {
		return nil, fmt.Errorf("invalid_nodeId_%v", nodeId)
	}

	// check valid group node
	if _token := token.NewToken(nodeDetail.NodeType); _token != token.GROUP {
		return nil, fmt.Errorf("invalid_nodeType_%v", nodeDetail.NodeType)
	}

	// load and check operator
	operatorToken := token.NewToken(nodeDetail.Operator)
	if !(nodeDetail.Operator == "" || operatorToken.IsGroupOperator()) {
		return nil, fmt.Errorf("invalid_operator_%v", nodeDetail.Operator)
	}

	// loop over children
	groupExp.Id = nodeId
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
		nodeDetail        NodeDetail
		ok                bool
		err               error
		lexp, rexp, r2exp ast.Expr
	)

	// check valid nodeDetail
	if nodeDetail, ok = conditions.Nodes[nodeId]; !ok {
		return nil, fmt.Errorf("invalid_nodeId_%v", nodeId)
	}

	// check valid condition node
	if _token := token.NewToken(nodeDetail.NodeType); _token != token.CONDITION {
		return nil, fmt.Errorf("invalid_nodeType_%v_%v", nodeDetail.NodeType, nodeId)
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
		return &ast.UniaryExpr{Id: nodeId, LHS: lexp}, nil
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
		return &ast.BinaryExpr{Id: nodeId, LHS: lexp, RHS: rexp}, nil
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
		return &ast.TerniaryExpr{Id: nodeId, LHS: lexp, RHS: rexp, RHS2: r2exp}, nil
	}
	return nil, fmt.Errorf("invalid_operator_%v_%v", nodeDetail.Operator, nodeId)
}

func parseConstantNode(nodeId string, conditions *Conditions, visited map[string]bool) (ast.Expr, error) {
	fmt.Println("started parseConstantNode ...", nodeId)

	// make visited nodeId
	visited[nodeId] = true

	var (
		nodeDetail NodeDetail
		ok         bool
	)

	// check valid nodeDetail
	if nodeDetail, ok = conditions.Nodes[nodeId]; !ok {
		return nil, fmt.Errorf("invalid_nodeId_%v", nodeId)
	}

	// check valid constant node
	if _token := token.NewToken(nodeDetail.NodeType); _token != token.CONSTANT {
		return nil, fmt.Errorf("invalid_nodeType_%v_%v", nodeDetail.NodeType, nodeId)
	}

	datatypeToken := token.NewToken(nodeDetail.Datatype)
	if !datatypeToken.IsLiteral() {
		return nil, fmt.Errorf("invalid_datatype_%v_%v", nodeDetail.Datatype, nodeId)
	}

	// switch datatypeToken {
	// case token.NUMBER:
	// 	v := nodeDetail.Value.(float64)
	// 	return &ast.NumberLiteral{Id: nodeId, Val: v}, nil
	// }

	return nil, nil
}

func parseParamsNode(nodeId string, _ *Conditions, visited map[string]bool) (ast.Expr, error) {
	fmt.Println("started parseParamsNode ...", nodeId)

	// make visited nodeId
	visited[nodeId] = true

	return nil, nil
}
