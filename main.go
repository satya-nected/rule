package main

import (
	"fmt"
	"test/utils/parser"
)

func main() {
	expr, err := parser.Parse(Data1)
	fmt.Println(expr, err)
}
