package main

import (
	"fmt"
	"test/utils/parser"
)

var data = `
{
	"startNode": "rule_01",
	"nodes": {
		"rule_01": {
			"operator": "",
			"parent": "",
			"nodeType": "group",
			"children": [
				"rule_02"
			],
			"siblingIndex": 1
		},
		"rule_02": {
			"operator": "eq",
			"parent": "rule_01",
			"nodeType": "condition",
			"siblingIndex": 1,
			"leftNode": [
				"rule_03"
			],
			"rightNode": [
				"rule_1"
			],
			"dataType": "numeric"
		},
		"rule_03": {
			"nodeType": "params",
			"sourceType": "constantInput",
			"attribute": "CI0",
			"parent": "rule_02",
			"siblingIndex": 1
		},
		"rule_1": {
			"siblingIndex": 1,
			"nodeType": "constant",
			"value": "50",
			"parent": "rule_1"
		}
	}
}
`

func main() {
	err := parser.Parse(data)
	fmt.Println(err)
}
