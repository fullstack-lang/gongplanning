// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gongplanning/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Predecessor:
		switch reverseField.GongstructName {
		// insertion point
		case "Task":
			switch reverseField.Fieldname {
			case "Predecessors":
				if _task, ok := stage.Task_Predecessors_reverseMap[inst]; ok {
					res = _task.Name
				}
			}
		}

	case *models.Project:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Task:
		switch reverseField.GongstructName {
		// insertion point
		case "Project":
			switch reverseField.Fieldname {
			case "Tasks":
				if _project, ok := stage.Project_Tasks_reverseMap[inst]; ok {
					res = _project.Name
				}
			}
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Predecessor:
		switch reverseField.GongstructName {
		// insertion point
		case "Task":
			switch reverseField.Fieldname {
			case "Predecessors":
				res = stage.Task_Predecessors_reverseMap[inst]
			}
		}

	case *models.Project:
		switch reverseField.GongstructName {
		// insertion point
		}

	case *models.Task:
		switch reverseField.GongstructName {
		// insertion point
		case "Project":
			switch reverseField.Fieldname {
			case "Tasks":
				res = stage.Project_Tasks_reverseMap[inst]
			}
		}

	default:
		_ = inst
	}
	return res
}
