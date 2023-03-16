package main

var Data1 = `
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

var Data2 = `
{
	"startNode": "rule_01",
	"nodes":{
		"rule_01": {
			"operator": "and",
			"parent": "",
			"nodeType": "group",
			"children": [
				"rule_02",
				"rule_2"
			],
			"siblingIndex": 1
		},
		"rule_02": {
			"operator": "gte",
			"parent": "rule_01",
			"nodeType": "condition",
			"siblingIndex": 1,
			"leftNode": [
				"rule_03"
			],
			"rightNode": [
				"rule_1"
			],
			"datatype": "numeric"
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
			"value": 60,
			"datatype":"number",
			"parent": "rule_1"
		},
		"rule_2": {
			"parent": "rule_01",
			"nodeType": "condition",
			"operator": "lt",
			"datatype": "numeric",
			"siblingIndex": 2,
			"leftNode": [
				"rule_3"
			],
			"rightNode": [
				"rule_5"
			]
		},
		"rule_3": {
			"nodeType": "params",
			"sourceType": "constantInput",
			"attribute": "CI0",
			"parent": "rule_2",
			"siblingIndex": 2
		},
		"rule_5": {
			"siblingIndex": 1,
			"nodeType": "constant",
			"datatype":"number",
			"value": 90,
			"parent": "rule_5"
		}
	}
}
`

var Data3 = `
{
	"startNode": "rule_01",
	"nodes": {
		"rule_01": {
			"operator": "and",
			"parent": "",
			"nodeType": "group",
			"children": [
				"rule_02",
				"rule_2",
				"rule_6"
			],
			"siblingIndex": 1
		},
		"rule_02": {
			"operator": "gte",
			"parent": "rule_01",
			"nodeType": "condition",
			"siblingIndex": 1,
			"leftNode": [
				"rule_03"
			],
			"rightNode": [
				"rule_1"
			],
			"datatype": "numeric"
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
			"value": 60,
			"datatype": "number",
			"parent": "rule_1"
		},
		"rule_2": {
			"parent": "rule_01",
			"nodeType": "condition",
			"operator": "lt",
			"datatype": "numeric",
			"siblingIndex": 2,
			"leftNode": [
				"rule_3"
			],
			"rightNode": [
				"rule_5"
			]
		},
		"rule_3": {
			"nodeType": "params",
			"sourceType": "constantInput",
			"attribute": "CI0",
			"parent": "rule_2",
			"siblingIndex": 2
		},
		"rule_5": {
			"siblingIndex": 1,
			"nodeType": "constant",
			"value": 90,
			"datatype": "number",
			"parent": "rule_5"
		},
		"rule_6": {
			"parent": "rule_01",
			"nodeType": "group",
			"operator": "or",
			"children": [
				"rule_9",
				"rule_12"
			],
			"siblingIndex": 3
		},
		"rule_9": {
			"parent": "rule_6",
			"nodeType": "condition",
			"operator": "eq",
			"siblingIndex": 1,
			"leftNode": [
				"rule_7"
			],
			"rightNode": [
				"rule_11"
			],
			"datatype": "string"
		},
		"rule_7": {
			"nodeType": "params",
			"sourceType": "constantInput",
			"attribute": "CI10",
			"parent": "rule_9",
			"siblingIndex": 1
		},
		"rule_11": {
			"siblingIndex": 1,
			"nodeType": "constant",
			"value": "satya",
			"datatype": "string",
			"parent": "rule_11"
		},
		"rule_12": {
			"parent": "rule_6",
			"nodeType": "condition",
			"operator": "eq",
			"datatype": "string",
			"siblingIndex": 2,
			"leftNode": [
				"rule_13"
			],
			"rightNode": [
				"rule_15"
			]
		},
		"rule_13": {
			"nodeType": "params",
			"sourceType": "constantInput",
			"attribute": "CI10",
			"parent": "rule_12",
			"siblingIndex": 2
		},
		"rule_15": {
			"siblingIndex": 1,
			"nodeType": "constant",
			"value": "aman",
			"datatype": "string",
			"parent": "rule_15"
		}
	}
}
`
