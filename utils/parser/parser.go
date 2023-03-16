package parser

import (
	"encoding/json"
	"fmt"
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

func parse(nodeId string, conditions *Conditions, visited map[string]bool) (bool, error) {
	fmt.Println("parsing start for ", nodeId)

	// check if node is visited
	if _, ok := visited[nodeId]; ok {
		return false, fmt.Errorf("node_already_visted_%v", nodeId)
	}
	var (
		ans, ok    bool
		err        error
		nodeDetail NodeDetail
	)

	if nodeDetail, ok = conditions.Nodes[nodeId]; !ok {
		return false, fmt.Errorf("invalid_nodeId_%v", nodeId)
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
			ans, err = false, fmt.Errorf("invalid_nodeType_%v", nodeDetail.NodeType)
		}
	}
	return ans, err
}

func parseGroupNode(nodeId string, conditions *Conditions, visited map[string]bool) (bool, error) {
	fmt.Println("started parseGroupNode ...", nodeId)

	// make visited nodeId
	visited[nodeId] = true

	var (
		nodeDetail NodeDetail
		ok         bool
	)

	// check valid nodeDetail
	if nodeDetail, ok = conditions.Nodes[nodeId]; !ok {
		return false, fmt.Errorf("invalid_nodeId_%v", nodeId)
	}

	// check valid group node
	if _token := token.NewToken(nodeDetail.NodeType); _token != token.GROUP {
		return false, fmt.Errorf("invalid_nodeType_%v", nodeDetail.NodeType)
	}

	// load and check operator
	operatorToken := token.NewToken(nodeDetail.Operator)
	if !(nodeDetail.Operator == "" || operatorToken.IsGroupOperator()) {
		return false, fmt.Errorf("invalid_operator_%v", nodeDetail.Operator)
	}

	// loop over children
	for _, childNode := range nodeDetail.Children {
		_, err := parse(childNode, conditions, visited)
		if err != nil {
			return false, err
		}
	}
	return false, nil
}

func parseConditionNode(nodeId string, conditions *Conditions, visited map[string]bool) (bool, error) {
	fmt.Println("started parseConditionNode ...", nodeId)

	// make visited nodeId
	visited[nodeId] = true

	var (
		nodeDetail NodeDetail
		ok         bool
	)

	// check valid nodeDetail
	if nodeDetail, ok = conditions.Nodes[nodeId]; !ok {
		return false, fmt.Errorf("invalid_nodeId_%v", nodeId)
	}

	// check valid condition node
	if _token := token.NewToken(nodeDetail.NodeType); _token != token.CONDITION {
		return false, fmt.Errorf("invalid_nodeType_%v", nodeDetail.NodeType)
	}

	// load and check operator
	operatorToken := token.NewToken(nodeDetail.Operator)
	if !operatorToken.IsConditionOperator() {
		return false, fmt.Errorf("invalid_operator_%v", nodeDetail.Operator)
	}

	switch operatorToken {
	case token.EQ:
		parse(nodeDetail.LeftNode[0], conditions, visited)
		parse(nodeDetail.RightNode[0], conditions, visited)

	}

	return true, nil
}

func parseParamsNode(nodeId string, _ *Conditions, visited map[string]bool) (bool, error) {
	fmt.Println("started parseParamsNode ...", nodeId)

	// make visited nodeId
	visited[nodeId] = true

	return false, nil
}

func parseConstantNode(nodeId string, _ *Conditions, visited map[string]bool) (bool, error) {
	fmt.Println("started parseConstantNode ...", nodeId)

	// make visited nodeId
	visited[nodeId] = true

	return false, nil
}
