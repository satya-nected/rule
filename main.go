package main

import (
	"fmt"
	"test/utils/parser"
)

var data1 = `
{
	"startNode": "rule_01",
	"nodes": {
		"rule_01": {
			"attribute": "",
			"children": [],
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
	}
}
`

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
			"datatype": "",
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
			"sourceType": "customInput",
			"value": null
		},
		"rule_2": {
			"attribute": "",
			"children": null,
			"datatype": "number",
			"leftNode": null,
			"name": "",
			"nodeType": "constant",
			"operator": "",
			"parent": "rule_2",
			"query": "",
			"rightNode": null,
			"siblingIndex": 1,
			"sourceType": "",
			"value": 23
		}
	}
}
`

func main() {
	err := parser.Parse(data)
	fmt.Println(err)
}
