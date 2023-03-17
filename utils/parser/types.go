package parser

import (
	"fmt"
	"rule/utils/token"
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

func (c *Conditions) ValidNode(nodeId string, tokenType token.Token) (*NodeDetail, error) {
	// check valid nodeDetail
	nodeDetail, ok := c.Nodes[nodeId]
	if !ok {
		return nil, fmt.Errorf("invalid_nodeId_%v", nodeId)
	}

	// check valid node token
	if _token := token.NewToken(nodeDetail.NodeType); _token != tokenType {
		return nil, fmt.Errorf("invalid_nodeType_%v", nodeDetail.NodeType)
	}

	return &nodeDetail, nil
}
