package adapter

import (
	diagrammer "github.com/fullstack-lang/gongplanning/go/diagrammer"

	gongplanning_models "github.com/fullstack-lang/gongplanning/go/models"
)

func NewProjectNode(
	portfolioAdapter *PortfolioAdapter,
	gongStruct *gongplanning_models.Project) (gongStructNode *ProjectNode) {
	gongStructNode = &ProjectNode{
		ElementNodeBase: ElementNodeBase{portfolioAdapter: portfolioAdapter}}

	gongStructNode.gongStruct = gongStruct
	return
}

var _ diagrammer.ModelElementNode = &ProjectNode{}

type ProjectNode struct {
	ElementNodeBase
	gongStruct *gongplanning_models.Project
}

// AddToDiagram implements diagrammer.ModelElementNode.
func (p *ProjectNode) AddToDiagram() {
	panic("unimplemented")
}

// CanBeAddedToDiagram implements diagrammer.ModelElementNode.
// Subtle: this method shadows the method (ElementNodeBase).CanBeAddedToDiagram of ProjectNode.ElementNodeBase.
func (p *ProjectNode) CanBeAddedToDiagram() bool {
	panic("unimplemented")
}

// GenerateProgeny implements diagrammer.ModelElementNode.
func (p *ProjectNode) GenerateProgeny() []diagrammer.ModelNode {
	panic("unimplemented")
}

// GetChildren implements diagrammer.ModelElementNode.
// Subtle: this method shadows the method (ElementNodeBase).GetChildren of ProjectNode.ElementNodeBase.
func (p *ProjectNode) GetChildren() []diagrammer.ModelNode {
	panic("unimplemented")
}

// GetElement implements diagrammer.ModelElementNode.
func (p *ProjectNode) GetElement() any {
	panic("unimplemented")
}

// GetName implements diagrammer.ModelElementNode.
func (p *ProjectNode) GetName() string {
	panic("unimplemented")
}

// GetParent implements diagrammer.ModelElementNode.
// Subtle: this method shadows the method (ElementNodeBase).GetParent of ProjectNode.ElementNodeBase.
func (p *ProjectNode) GetParent() diagrammer.ModelNode {
	panic("unimplemented")
}

// IsExpanded implements diagrammer.ModelElementNode.
// Subtle: this method shadows the method (ElementNodeBase).IsExpanded of ProjectNode.ElementNodeBase.
func (p *ProjectNode) IsExpanded() bool {
	panic("unimplemented")
}

// IsNameEditable implements diagrammer.ModelElementNode.
// Subtle: this method shadows the method (ElementNodeBase).IsNameEditable of ProjectNode.ElementNodeBase.
func (p *ProjectNode) IsNameEditable() bool {
	panic("unimplemented")
}

// RemoveFromDiagram implements diagrammer.ModelElementNode.
func (p *ProjectNode) RemoveFromDiagram() {
	panic("unimplemented")
}

// SetIsExpanded implements diagrammer.ModelElementNode.
// Subtle: this method shadows the method (ElementNodeBase).SetIsExpanded of ProjectNode.ElementNodeBase.
func (p *ProjectNode) SetIsExpanded(isExpanded bool) {
	panic("unimplemented")
}
