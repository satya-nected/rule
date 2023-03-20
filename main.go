package main

import (
	"encoding/json"
	"fmt"

	"github.com/satya-nected/rule/utils/evalute"
	"github.com/satya-nected/rule/utils/parser"
)

func main() {
	dataA := [...]string{Data1, Data2, Data3}
	paramsA := [...]string{Params1, Params2, Params3}
	var params map[string]map[string]interface{}

	for i := 0; i < 3; i++ {
		expr, err := parser.Parse(dataA[i])
		if err != nil {
			fmt.Println("error in parsing ", err)
			return
		}

		if err := json.Unmarshal([]byte(paramsA[i]), &params); err != nil {
			fmt.Println("error in evalute ", err)
			return
		}

		result, err := evalute.Evaluate(expr, params)
		if err != nil {
			fmt.Println("error: ", err)
			return
		}
		fmt.Println("expr: ", expr)
		fmt.Println("params: ", params)
		fmt.Println("result: ", result, "\n\n.")
	}

	/*
		expr, err := parser.Parse(Data3)
		if err != nil {
			fmt.Println("error in parsing ", err)
			return
		}

		var params map[string]map[string]interface{}
		if err := json.Unmarshal([]byte(Params3), &params); err != nil {
			fmt.Println("error in unmarshal params ", err)
			return
		}

		result, err := evalute.Evaluate(expr, params)
		if err != nil {
			fmt.Println("error in evalute ", err)
			return
		}
		fmt.Println("expr: ", expr)
		fmt.Println("params: ", params)
		fmt.Println("result: ", result)
	*/
}
