package main

import (
	"encoding/json"
	"fmt"
	"test/utils/evalute"
	"test/utils/parser"
)

func main() {
	expr, err := parser.Parse(Data1)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	var params map[string]map[string]interface{}
	if err := json.Unmarshal([]byte(Params1), &params); err != nil {
		fmt.Println("error: ", err)
		return
	}

	result, err := evalute.Evaluate(expr, params)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println(expr, result)
}
