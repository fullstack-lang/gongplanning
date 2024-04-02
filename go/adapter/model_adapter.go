package adapter

import (
	diagrammer "github.com/fullstack-lang/gongplanning/go/diagrammer"
)

type ModelAdapter struct {
	portfolioAdapter *PortfolioAdapter
	rootNodes        []diagrammer.ModelNode
}

func NewModelAdapter(portfolioAdapter *PortfolioAdapter) (adapter *ModelAdapter) {
	adapter = new(ModelAdapter)
	adapter.portfolioAdapter = portfolioAdapter
	return
}

var _ diagrammer.Model = &ModelAdapter{}
