package types

//go:generate jsonenums -type=AnswerLayoutType

type AnswerLayoutType int

const (
	LayoutColumn1 AnswerLayoutType = iota
	LayoutColumn2
	LayoutColumn3
	LayoutColumn4
	LayoutImageColumn4
	LayoutImageColumn3
	LayoutImageColumn2
	LayoutImageColumn1
	LayoutRuler
)
