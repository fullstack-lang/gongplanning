package adapter

import (
	diagrammer "github.com/fullstack-lang/gongplanning/go/diagrammer"
	gongsvg_models "github.com/fullstack-lang/gongsvg/go/models"
)

type PortfolioAdapter struct {
	gongsvgStage *gongsvg_models.StageStruct
	diagrammer   *diagrammer.Diagrammer
}

func NewPortfolioAdapter(

	svgStage *gongsvg_models.StageStruct,
) (portfolioAdapter *PortfolioAdapter) {

	portfolioAdapter = new(PortfolioAdapter)

	portfolioAdapter.gongsvgStage = svgStage

	return
}

func (portfolioAdapter *PortfolioAdapter) SetDiagrammer(diagrammer *diagrammer.Diagrammer) {
	portfolioAdapter.diagrammer = diagrammer
}

var _ diagrammer.Portfolio = &PortfolioAdapter{}

// AddElementToDiagram implements diagrammer.Portfolio.
func (portfolioAdapter *PortfolioAdapter) AddElementToDiagram(diagrammer.ModelElementNode) map[diagrammer.ModelElementNode]diagrammer.Shape {
	panic("unimplemented")
}

// GeneratesProgeny implements diagrammer.Portfolio.
func (portfolioAdapter *PortfolioAdapter) GeneratesProgeny() []diagrammer.PortfolioNode {
	return nil // to be completed
}

// GetChildren implements diagrammer.Portfolio.
func (portfolioAdapter *PortfolioAdapter) GetChildren() []diagrammer.PortfolioNode {
	return nil // to be completed
}

// GetSelectedPortfolioDiagramNode implements diagrammer.Portfolio.
func (portfolioAdapter *PortfolioAdapter) GetSelectedPortfolioDiagramNode() diagrammer.PortfolioDiagramNode {
	panic("unimplemented")
}

// IsInDrawingMode implements diagrammer.Portfolio.
func (portfolioAdapter *PortfolioAdapter) IsInDrawingMode() bool {
	return false
}

// IsInSelectionMode implements diagrammer.Portfolio.
func (portfolioAdapter *PortfolioAdapter) IsInSelectionMode() bool {
	panic("unimplemented")
}

// RemoveElementFromDiagram implements diagrammer.Portfolio.
func (portfolioAdapter *PortfolioAdapter) RemoveElementFromDiagram(diagrammer.ModelElementNode) map[diagrammer.ModelElementNode]diagrammer.Shape {
	panic("unimplemented")
}
