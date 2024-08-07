// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongplanning/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Predecessor:
		predecessorInstance := any(concreteInstance).(*models.Predecessor)
		ret2 := backRepo.BackRepoPredecessor.GetPredecessorDBFromPredecessorPtr(predecessorInstance)
		ret = any(ret2).(*T2)
	case *models.Project:
		projectInstance := any(concreteInstance).(*models.Project)
		ret2 := backRepo.BackRepoProject.GetProjectDBFromProjectPtr(projectInstance)
		ret = any(ret2).(*T2)
	case *models.Task:
		taskInstance := any(concreteInstance).(*models.Task)
		ret2 := backRepo.BackRepoTask.GetTaskDBFromTaskPtr(taskInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Predecessor:
		tmp := GetInstanceDBFromInstance[models.Predecessor, PredecessorDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Project:
		tmp := GetInstanceDBFromInstance[models.Project, ProjectDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Task:
		tmp := GetInstanceDBFromInstance[models.Task, TaskDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Predecessor:
		tmp := GetInstanceDBFromInstance[models.Predecessor, PredecessorDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Project:
		tmp := GetInstanceDBFromInstance[models.Project, ProjectDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Task:
		tmp := GetInstanceDBFromInstance[models.Task, TaskDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
