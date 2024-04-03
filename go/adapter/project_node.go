package adapter

import (
	diagrammer "github.com/fullstack-lang/gongplanning/go/diagrammer"

	gongplanning_models "github.com/fullstack-lang/gongplanning/go/models"
)

func NewProjectNode(
	portfolioAdapter *PortfolioAdapter,
	gongplanningStage *gongplanning_models.StageStruct,
	project *gongplanning_models.Project) (gongStructNode *ProjectNode) {

	return &ProjectNode{
		ElementNodeBase:   ElementNodeBase{portfolioAdapter: portfolioAdapter},
		gongplanningStage: gongplanningStage,
		project:           project,
	}
}

var _ diagrammer.ModelElementNode = &ProjectNode{}

type ProjectNode struct {
	ElementNodeBase
	gongplanningStage *gongplanning_models.StageStruct
	project           *gongplanning_models.Project
}

// HasCheckboxButton implements diagrammer.ModelElementNode.
func (p *ProjectNode) HasCheckboxButton() bool {
	return false
}

// GenerateProgeny
func (p *ProjectNode) GenerateProgeny() []diagrammer.ModelNode {
	return nil
}

// GetChildren
// Subtle: this method shadows the method (ElementNodeBase).GetChildren of ProjectNode.ElementNodeBase.
func (p *ProjectNode) GetChildren() []diagrammer.ModelNode {
	return nil
}

// GetElement
func (p *ProjectNode) GetElement() any {
	return p.project
}

// GetName
func (p *ProjectNode) GetName() string {
	return p.project.Name
}

// GetParent
// Subtle: this method shadows the method (ElementNodeBase).GetParent of ProjectNode.ElementNodeBase.
func (p *ProjectNode) GetParent() diagrammer.ModelNode {
	panic("unimplemented")
}

// IsExpanded
// Subtle: this method shadows the method (ElementNodeBase).IsExpanded of ProjectNode.ElementNodeBase.
func (p *ProjectNode) IsExpanded() bool {
	return p.project.IsExpanded
}

// SetIsExpanded
// Subtle: this method shadows the method (ElementNodeBase).SetIsExpanded of ProjectNode.ElementNodeBase.
func (p *ProjectNode) SetIsExpanded(isExpanded bool) {
	p.project.IsExpanded = true
	p.gongplanningStage.Commit()

}

// IsNameEditable
// Subtle: this method shadows the method (ElementNodeBase).IsNameEditable of ProjectNode.ElementNodeBase.
func (p *ProjectNode) IsNameEditable() bool {
	panic("unimplemented")
}

// AddToDiagram
func (p *ProjectNode) AddToDiagram() {
	panic("unimplemented")
}

// CanBeAddedToDiagram
// Subtle: this method shadows the method (ElementNodeBase).CanBeAddedToDiagram of ProjectNode.ElementNodeBase.
func (p *ProjectNode) CanBeAddedToDiagram() bool {
	panic("unimplemented")
}

// RemoveFromDiagram
func (p *ProjectNode) RemoveFromDiagram() {
	panic("unimplemented")
}
