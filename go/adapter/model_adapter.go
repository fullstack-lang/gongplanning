package adapter

import (
	diagrammer "github.com/fullstack-lang/gongplanning/go/diagrammer"

	gongplanning_models "github.com/fullstack-lang/gongplanning/go/models"
)

type ModelAdapter struct {
	portfolioAdapter  *PortfolioAdapter
	gongplanningStage *gongplanning_models.StageStruct
	rootNodes         []diagrammer.ModelNode
}

func NewModelAdapter(portfolioAdapter *PortfolioAdapter,
	gongplanningStage *gongplanning_models.StageStruct,
) (adapter *ModelAdapter) {
	adapter = new(ModelAdapter)

	adapter.portfolioAdapter = portfolioAdapter
	adapter.gongplanningStage = gongplanningStage

	return
}

var _ diagrammer.Model = &ModelAdapter{}

// GenerateProgeny implements diagrammer.Model.
func (m *ModelAdapter) GenerateProgeny() []diagrammer.ModelNode {
	projectCategoryNode := NewProjectCategoryNode(m.portfolioAdapter, "Projects", m.gongplanningStage)
	m.rootNodes = append(m.rootNodes, projectCategoryNode)

	return m.rootNodes
}

// GetChildren implements diagrammer.Model.
func (m *ModelAdapter) GetChildren() []diagrammer.ModelNode {
	return m.rootNodes
}
