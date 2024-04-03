// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongplanning/go/models"
	"github.com/fullstack-lang/gongplanning/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Predecessor:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("DependencyType", instanceWithInferedType.DependencyType, instanceWithInferedType, probe.formStage, formGroup)
		AssociationFieldToForm("Task", instanceWithInferedType.Task, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Task"
			rf.Fieldname = "Predecessors"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Task),
					"Predecessors",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Task, *models.Predecessor](
					nil,
					"Predecessors",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	case *models.Project:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Tasks", instanceWithInferedType, &instanceWithInferedType.Tasks, formGroup, probe)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)

	case *models.Task:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Predecessors", instanceWithInferedType, &instanceWithInferedType.Predecessors, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Project"
			rf.Fieldname = "Tasks"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Project),
					"Tasks",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Project, *models.Task](
					nil,
					"Tasks",
					instanceWithInferedType,
					formGroup,
					probe)
			}
		}

	default:
		_ = instanceWithInferedType
	}
}
