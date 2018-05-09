package main

import (
	"fmt"
	"strconv"
)

func main() {
	oper := createOperate("+")
	fmt.Println("The result is ", oper.GetResult(3.5, 4.4))
}

func createOperate(operate string) Operation {
	var oper Operation

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
	GetResult(op1, op2 float64) string
}

type OperationAdd struct {
}

type OperationSub struct {
}

type OperationMul struct {
}

type OperationDiv struct {
}

func (op *OperationAdd) GetResult(op1, op2 float64) string {
	return strconv.FormatFloat(op1+op2, 'f', -1, 64)
}

func (op *OperationSub) GetResult(op1, op2 float64) string {
	return strconv.FormatFloat(op1-op2, 'f', -1, 64)
}

func (op *OperationMul) GetResult(op1, op2 float64) string {
	return strconv.FormatFloat(op1*op2, 'f', -1, 64)
}

func (op *OperationDiv) GetResult(op1, op2 float64) string {
	if op1 == 0.0 {
		fmt.Println("operator 1 can't be zero")
		return ""
	}
	return strconv.FormatFloat(op1/op2, 'f', -1, 64)
}
