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
