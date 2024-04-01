package models

// StackName -
type StackName string

// values for TableName
const (
	StackNameDefault        StackName = "gongplanning-default"
	StackNameTreeAtTheLeft  StackName = "gongplanning-tree-left"
	StackNameSVGAtTheCenter StackName = "gongplanning-svg-center"
	StackNameFormAtTheRight StackName = "gongplanning-form-right"
)
