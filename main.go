package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"test/utils"
)

var data = `
{
	"startNode": "rule_01",
	"nodes": {
		"rule_01": {
			"attribute": "",
			"children": [
			"rule_02"
			],
			"datatype": "",
			"leftNode": null,
			"name": "",
			"nodeType": "group",
			"operator": "",
			"parent": "",
			"query": "",
			"rightNode": null,
			"siblingIndex": 1,
			"sourceType": "",
			"value": null
		},
		"rule_02": {
			"attribute": "",
			"children": null,
			"datatype": "string",
			"leftNode": [
			"rule_03"
			],
			"name": "",
			"nodeType": "condition",
			"operator": "eq",
			"parent": "rule_01",
			"query": "",
			"rightNode": [
			"rule_2"
			],
			"siblingIndex": 1,
			"sourceType": "",
			"value": null
		},
		"rule_03": {
			"attribute": "CI0",
			"children": null,
			"datatype": "",
			"leftNode": null,
			"name": "",
			"nodeType": "params",
			"operator": "",
			"parent": "rule_02",
			"query": "",
			"rightNode": null,
			"siblingIndex": 1,
			"sourceType": "constantInput",
			"value": null
		},
		"rule_2": {
			"attribute": "",
			"children": null,
			"datatype": "",
			"leftNode": null,
			"name": "",
			"nodeType": "constant",
			"operator": "",
			"parent": "rule_2",
			"query": "",
			"rightNode": null,
			"siblingIndex": 1,
			"sourceType": "",
			"value": "23"
		}
	}
}
`

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

func ExecuteGroupNode(nodeId string, conditions *Conditions) (bool, error) {
	fmt.Println("started ExecuteGroupNode ...", nodeId)

	nodeDetail, ok := conditions.Nodes[nodeId]
	if !ok {
		return false, errors.New("invalid_nodeId")
	}
	if nodeDetail.NodeType != "group" {
		return false, errors.New("invalid_nodeType")
	}

	if !(nodeDetail.Operator == "" || nodeDetail.Operator == "and" || nodeDetail.Operator == "or") {
		return false, errors.New("invalid_operator")
	}

	finalAns := true
	operator := nodeDetail.Operator

	if operator == "" {
		operator = "and"
	}

	for _, childNode := range nodeDetail.Children {
		childNodeDetail, ok := conditions.Nodes[childNode]
		if !ok {
			return false, errors.New("invalid_nodeId")
		}
		switch childNodeDetail.NodeType {
		case "group":
			ans, err := ExecuteGroupNode(childNode, conditions)
			if err != nil {
				return false, err
			}

			if operator == "and" && !ans {
				return false, nil
			}

			if operator == "and" {
				finalAns = finalAns && ans
			} else if operator == "or" {
				finalAns = finalAns || ans
			}

		case "condition":
			ans, err := ExecuteConditionNode(childNode, conditions)
			if err != nil {
				return false, err
			}

			if operator == "and" && !ans {
				return false, nil
			}

			if operator == "and" {
				finalAns = finalAns && ans
			} else if operator == "or" {
				finalAns = finalAns || ans
			}

		case "sqlCondition":
			ans, err := ExecuteSqlConditionNode(childNode, conditions)
			if err != nil {
				return false, err
			}

			if operator == "and" && !ans {
				return false, nil
			}

			if operator == "and" {
				finalAns = finalAns && ans
			} else if operator == "or" {
				finalAns = finalAns || ans
			}
		}
	}
	return finalAns, nil
}

func ExecuteParamsNode(nodeId string, conditions *Conditions) {

}

func ExecuteConditionNode(nodeId string, conditions *Conditions) (bool, error) {
	fmt.Println("started ExecuteConditionNode ...", nodeId)

	nodeDetail, ok := conditions.Nodes[nodeId]
	if !ok {
		return false, errors.New("invalid_nodeId")
	}
	if nodeDetail.NodeType != "condition" {
		return false, errors.New("invalid_nodeType")
	}

	operatorType, ok := utils.OperatorType[nodeDetail.Operator]
	if !ok {
		return false, errors.New("invalid_operator")
	}

	fmt.Println(operatorType)

	return true, nil
}

func ExecuteSqlConditionNode(nodeId string, conditions *Conditions) (bool, error) {
	fmt.Println("started ExecuteSqlConditionNode ...", nodeId)

	return true, nil
}

func ExecuteConstantNode(nodeId string, conditions *Conditions) {

}

func main() {
	var conditions Conditions
	json.Unmarshal([]byte(data), &conditions)

	startNode, ok := conditions.Nodes[conditions.StartNode]
	if !ok {
		fmt.Println("invalid start Node")
	}
	if startNode.NodeType == "group" {
		ans, err := ExecuteGroupNode(conditions.StartNode, &conditions)
		fmt.Println(ans, err)
	}

}
