package main

var Data0 = `
{
    "startNode": "1",
    "nodes":	{
		"1": {
		  "nodeType": "group",
		  "operator": "and",
		  "parent": "",
		  "children": ["2", "3"],
		  "siblingIndex": 1
		},
		"2": {
		  "nodeType": "condition",
		  "operator": "gt",
		  "dataType": "numeric",
		  "parent": "1",
		  "leftNode": ["4"],
		  "rightNode": ["5"],
		  "siblingIndex": 1
		},
		"3": {
		  "nodeType": "group",
		  "operator": "or",
		  "parent": "1",
		  "children": ["6", "7"],
		  "siblingIndex": 2
		},
		"4": {
		  "nodeType": "params",
		  "sourceType": "customInput",
		  "dataType": "numeric",
		  "attribute": "cartValue",
		  "parent": "2",
		  "siblingIndex": 1
		},
		"5": {
		  "nodeType": "constant",
		  "dataType": "numeric",
		  "value": 100,
		  "parent": "2",
		  "siblingIndex": 1
		},
		"6": {
		  "nodeType": "condition",
		  "operator": "gt",
		  "dataType": "numeric",
		  "parent": "3",
		  "leftNode": ["8"],
		  "rightNode": ["9"],
		  "siblingIndex": 1
		},
		"7": {
		  "nodeType": "condition",
		  "operator": "eq",
		  "dataType": "string",
		  "parent": "3",
		  "leftNode": ["10"],
		  "rightNode": ["11"],
		  "siblingIndex": 2
		},
		"8": {
		  "nodeType": "params",
		  "sourceType": "customInput",
		  "attribute": "cartValue",
		  "dataType": "numeric",
		  "parent": "6",
		  "siblingIndex": 1
		},
		"9": {
		  "nodeType": "constant",
		  "dataType": "numeric",
		  "value": 200,
		  "parent": "6",
		  "siblingIndex": 1
		},
		"10": {
		  "nodeType": "params",
		  "sourceType": "customInput",
		  "attribute": "itemsCategory",
		  "dataType": "string",
		  "parent": "7",
		  "siblingIndex": 1
		},
		"11": {
		  "nodeType": "constant",
		  "dataType": "string",
		  "value": "jwellery",
		  "parent": "7",
		  "siblingIndex": 1
		}
	}
}
`

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
			"datatype": "numeric",
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
			"sourceType": "customInput",
			"attribute": "CI0",
			"parent": "rule_02",
			"siblingIndex": 1
		},
		"rule_1": {
			"siblingIndex": 1,
			"nodeType": "constant",
			"value": 60,
			"datatype":"numeric",
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
			"sourceType": "customInput",
			"attribute": "CI0",
			"parent": "rule_2",
			"siblingIndex": 2
		},
		"rule_5": {
			"siblingIndex": 1,
			"nodeType": "constant",
			"datatype":"numeric",
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
			"sourceType": "customInput",
			"attribute": "CI0",
			"parent": "rule_02",
			"siblingIndex": 1
		},
		"rule_1": {
			"siblingIndex": 1,
			"nodeType": "constant",
			"value": 60,
			"datatype": "numeric",
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
			"sourceType": "customInput",
			"attribute": "CI0",
			"parent": "rule_2",
			"siblingIndex": 2
		},
		"rule_5": {
			"siblingIndex": 1,
			"nodeType": "constant",
			"value": 90,
			"datatype": "numeric",
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
			"sourceType": "customInput",
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
			"sourceType": "customInput",
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

var Data4 = `
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
				"rule_6",
				"rule_19"
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
			"dataType": "numeric"
		},
		"rule_03": {
			"nodeType": "params",
			"sourceType": "customInput",
			"attribute": "CI0",
			"parent": "rule_02",
			"siblingIndex": 1
		},
		"rule_1": {
			"siblingIndex": 1,
			"nodeType": "constant",
			"value": 60,
			"datatype":"numeric",
			"parent": "rule_1"
		},
		"rule_2": {
			"parent": "rule_01",
			"nodeType": "condition",
			"operator": "lt",
			"dataType": "numeric",
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
			"sourceType": "customInput",
			"attribute": "CI0",
			"parent": "rule_2",
			"siblingIndex": 2
		},
		"rule_5": {
			"siblingIndex": 1,
			"nodeType": "constant",
			"value": 90,
			"datatype":"numeric",
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
			"dataType": "string"
		},
		"rule_7": {
			"nodeType": "params",
			"sourceType": "customInput",
			"attribute": "CI10",
			"parent": "rule_9",
			"siblingIndex": 1
		},
		"rule_11": {
			"siblingIndex": 1,
			"nodeType": "constant",
			"value": "satya",
			"datatype":"string",
			"parent": "rule_11"
		},
		"rule_12": {
			"parent": "rule_6",
			"nodeType": "condition",
			"operator": "eq",
			"dataType": "string",
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
			"sourceType": "customInput",
			"attribute": "CI10",
			"parent": "rule_12",
			"siblingIndex": 2
		},
		"rule_15": {
			"siblingIndex": 1,
			"nodeType": "constant",
			"value": "aman",
			"datatype":"string",
			"parent": "rule_15"
		},
		"rule_19": {
			"parent": "rule_01",
			"nodeType": "condition",
			"operator": "bet",
			"dataType": "numeric",
			"siblingIndex": 4,
			"leftNode": [
				"rule_20"
			],
			"rightNode": [
				"rule_23",
				"rule_24"
			]
		},
		"rule_20": {
			"nodeType": "params",
			"sourceType": "customInput",
			"attribute": "CI0",
			"parent": "rule_19",
			"siblingIndex": 4
		},
		"rule_23": {
			"siblingIndex": 1,
			"nodeType": "constant",
			"value": 4,
			"datatype":"numeric",
			"parent": "rule_23"
		},
		"rule_24": {
			"siblingIndex": 1,
			"nodeType": "constant",
			"value": 10,
			"datatype":"numeric",
			"parent": "rule_24"
		}
	}
}
`
