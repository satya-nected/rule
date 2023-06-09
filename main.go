package main

import (
	"encoding/json"
	"fmt"
	"rule/utils/evalute"
	"rule/utils/parser"
)

func main() {
	// dataA := [...]string{Data1, Data2, Data3}
	// paramsA := [...]string{Params1, Params2, Params3}
	// var params map[string]map[string]interface{}

	// logger.SetLogLevel(logrus.DebugLevel)

	// for i := 0; i < 3; i++ {
	// 	expr, err := parser.Parse(dataA[i])
	// 	if err != nil {
	// 		fmt.Println("error in parsing ", err)
	// 		return
	// 	}

	// 	if err := json.Unmarshal([]byte(paramsA[i]), &params); err != nil {
	// 		fmt.Println("error in evalute ", err)
	// 		return
	// 	}

	// 	result, err := evalute.Evaluate(expr, params)
	// 	if err != nil {
	// 		fmt.Println("error: ", err)
	// 		return
	// 	}
	// 	fmt.Println("expr: ", expr)
	// 	fmt.Println("params: ", params)
	// 	fmt.Println("result: ", result, "\n\n.")
	// }

	var d map[string]interface{}
	json.Unmarshal([]byte(Data0), &d)

	expr, err := parser.Parse(d)
	if err != nil {
		fmt.Println("error in parsing ", err)
		return
	}

	var params map[string]map[string]interface{}
	if err := json.Unmarshal([]byte(Params0), &params); err != nil {
		fmt.Println("error in unmarshal params ", err)
		return
	}

	fmt.Println(expr, "ss")

	result, err := evalute.Evaluate(expr, params)
	if err != nil {
		fmt.Println("error in evalute ", err)
		return
	}
	fmt.Println("expr: ", expr)
	fmt.Println("params: ", params)
	fmt.Println("result: ", result)

}
