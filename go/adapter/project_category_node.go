package adapter

import (
	"cmp"
	"slices"

	diagrammer "github.com/fullstack-lang/gongplanning/go/diagrammer"

	gongplanning_models "github.com/fullstack-lang/gongplanning/go/models"
)

type ProjectCategoryNode struct {
	ModelCategoryNodeBase
	gongplanningStage *gongplanning_models.StageStruct
}

func NewProjectCategoryNode(
	portfolioAdapter *PortfolioAdapter,
	name string,
	gongplanningStage *gongplanning_models.StageStruct) *ProjectCategoryNode {
	return &ProjectCategoryNode{ModelCategoryNodeBase: ModelCategoryNodeBase{portfolioAdapter: portfolioAdapter, Name: name}}
}

// GenerateProgeny implements diagrammer.Node.
func (categoryNode *ProjectCategoryNode) GenerateProgeny() []diagrammer.ModelNode {

	for project := range *gongplanning_models.GetGongstructInstancesSet[gongplanning_models.Project](
		categoryNode.gongplanningStage) {

		projectNode := NewProjectNode(categoryNode.portfolioAdapter, project)
		projectNode.GenerateProgeny()
		categoryNode.children = append(categoryNode.children, projectNode)
	}

	slices.SortFunc(categoryNode.children, func(a, b diagrammer.ModelNode) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	return categoryNode.children
}
