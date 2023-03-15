package utils

import (
	"errors"
	"reflect"
)

type OperandDetail struct {
	Name          string
	LeftOperands  int
	RightOperands int
}

func (op *OperandDetail) Execute(operatorType string, left, right, optional interface{}) (bool, error) {
	_, ok := OperatorType[operatorType]
	if !ok {
		return false, errors.New("operator_not_found")
	}

	return false, nil
}

func (op *OperandDetail) ExecuteEq(left, right interface{}) (bool, error) {
	if reflect.TypeOf(left).Kind() != reflect.TypeOf(right).Kind() {
		return false, errors.New("invalid_left_right_operands")
	}
	return left == right, nil
}
