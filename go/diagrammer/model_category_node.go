package diagrammer

type ModelCategoryNode interface {
	ModelNode

	CanAddInstance() bool
}
