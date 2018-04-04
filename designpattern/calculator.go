package main

import (
	"fmt"
	"strconv"
)

func main() {
	oper, ok := createOperate("+").(Oper)
	if !ok {
		fmt.Println("the program is failure")
	}
	oper.operator1 = 1
	oper.operator2 = 2
	fmt.Println("THE RESULT IS ", oper.GetResult())
}

func createOperate(operate string) interface{} {
	var oper interface{}

	switch operate {
	case "+":
		oper = new(OperationAdd)
		break
	case "-":
		oper = new(OperationSub)
		break
	case "*":
		oper = new(OperationMul)
		break
	case "/":
		oper = new(OperationDiv)
		break
	}

	return oper
}

type Operation interface {
	GetResult() string
}

type Oper struct {
	operator1 float64
	operator2 float64
	result    string
}

func (op *Oper) GetResult() string {
	return ""
}

type OperationAdd struct {
	Oper
}

type OperationSub struct {
	Oper
}

type OperationMul struct {
	Oper
}

type OperationDiv struct {
	Oper
}

func (op *OperationAdd) GetResult() string {
	op.result = strconv.FormatFloat(op.operator1+op.operator2, 'E', -1, 64)
	return op.result
}

func (op *OperationSub) GetResult() string {
	op.result = strconv.FormatFloat(op.operator1-op.operator2, 'E', -1, 64)
	return op.result
}

func (op *OperationMul) GetResult() string {
	op.result = strconv.FormatFloat(op.operator1*op.operator2, 'E', -1, 64)
	return op.result
}

func (op *OperationDiv) GetResult() string {
	if op.operator1 == 0.0 {
		fmt.Println("operator 1 can't be zero")
		return ""
	}
	op.result = strconv.FormatFloat(op.operator1/op.operator2, 'E', -1, 64)
	return op.result
}
