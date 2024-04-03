package diagrammer

type ModelElementNode interface {
	ModelNode

	GetElement() any

	HasCheckboxButton() bool

	// CanBeAddedToDiagram returns true if the model element can be displayed
	CanBeAddedToDiagram() bool
	AddToDiagram()
	RemoveFromDiagram()
}
