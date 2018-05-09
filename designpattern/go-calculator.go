package main

import "strconv"

type Operation interface {
	GetValue() string
}

type OperationAdd struct {
	op1, op2 float64
	result   string
}

type OperationSub struct {
	op1, op2 float64
	result   string
}

func (o *OperationAdd) GetValue() string {
	return strconv.FormatFloat(o.op1+o.op2, 'f', -1, 64)
}

func (o *OperationSub) GetValue() string {
	return strconv.FormatFloat(o.op1-o.op2, 'f', -1, 64)
}

func main() {

}
