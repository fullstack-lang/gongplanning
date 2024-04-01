// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongplanning/go/models"
	"github.com/fullstack-lang/gongplanning/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__PredecessorFormCallback(
	predecessor *models.Predecessor,
	probe *Probe,
	formGroup *table.FormGroup,
) (predecessorFormCallback *PredecessorFormCallback) {
	predecessorFormCallback = new(PredecessorFormCallback)
	predecessorFormCallback.probe = probe
	predecessorFormCallback.predecessor = predecessor
	predecessorFormCallback.formGroup = formGroup

	predecessorFormCallback.CreationMode = (predecessor == nil)

	return
}

type PredecessorFormCallback struct {
	predecessor *models.Predecessor

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (predecessorFormCallback *PredecessorFormCallback) OnSave() {

	log.Println("PredecessorFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	predecessorFormCallback.probe.formStage.Checkout()

	if predecessorFormCallback.predecessor == nil {
		predecessorFormCallback.predecessor = new(models.Predecessor).Stage(predecessorFormCallback.probe.stageOfInterest)
	}
	predecessor_ := predecessorFormCallback.predecessor
	_ = predecessor_

	for _, formDiv := range predecessorFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(predecessor_.Name), formDiv)
		case "DependencyType":
			FormDivEnumStringFieldToField(&(predecessor_.DependencyType), formDiv)
		case "Task":
			FormDivSelectFieldToField(&(predecessor_.Task), predecessorFormCallback.probe.stageOfInterest, formDiv)
		case "Task:Predecessors":
			// we need to retrieve the field owner before the change
			var pastTaskOwner *models.Task
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Task"
			rf.Fieldname = "Predecessors"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				predecessorFormCallback.probe.stageOfInterest,
				predecessorFormCallback.probe.backRepoOfInterest,
				predecessor_,
				&rf)

			if reverseFieldOwner != nil {
				pastTaskOwner = reverseFieldOwner.(*models.Task)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastTaskOwner != nil {
					idx := slices.Index(pastTaskOwner.Predecessors, predecessor_)
					pastTaskOwner.Predecessors = slices.Delete(pastTaskOwner.Predecessors, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _task := range *models.GetGongstructInstancesSet[models.Task](predecessorFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _task.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newTaskOwner := _task // we have a match
						if pastTaskOwner != nil {
							if newTaskOwner != pastTaskOwner {
								idx := slices.Index(pastTaskOwner.Predecessors, predecessor_)
								pastTaskOwner.Predecessors = slices.Delete(pastTaskOwner.Predecessors, idx, idx+1)
								newTaskOwner.Predecessors = append(newTaskOwner.Predecessors, predecessor_)
							}
						} else {
							newTaskOwner.Predecessors = append(newTaskOwner.Predecessors, predecessor_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if predecessorFormCallback.formGroup.HasSuppressButtonBeenPressed {
		predecessor_.Unstage(predecessorFormCallback.probe.stageOfInterest)
	}

	predecessorFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Predecessor](
		predecessorFormCallback.probe,
	)
	predecessorFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if predecessorFormCallback.CreationMode || predecessorFormCallback.formGroup.HasSuppressButtonBeenPressed {
		predecessorFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(predecessorFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__PredecessorFormCallback(
			nil,
			predecessorFormCallback.probe,
			newFormGroup,
		)
		predecessor := new(models.Predecessor)
		FillUpForm(predecessor, newFormGroup, predecessorFormCallback.probe)
		predecessorFormCallback.probe.formStage.Commit()
	}

	fillUpTree(predecessorFormCallback.probe)
}
func __gong__New__ProjectFormCallback(
	project *models.Project,
	probe *Probe,
	formGroup *table.FormGroup,
) (projectFormCallback *ProjectFormCallback) {
	projectFormCallback = new(ProjectFormCallback)
	projectFormCallback.probe = probe
	projectFormCallback.project = project
	projectFormCallback.formGroup = formGroup

	projectFormCallback.CreationMode = (project == nil)

	return
}

type ProjectFormCallback struct {
	project *models.Project

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (projectFormCallback *ProjectFormCallback) OnSave() {

	log.Println("ProjectFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	projectFormCallback.probe.formStage.Checkout()

	if projectFormCallback.project == nil {
		projectFormCallback.project = new(models.Project).Stage(projectFormCallback.probe.stageOfInterest)
	}
	project_ := projectFormCallback.project
	_ = project_

	for _, formDiv := range projectFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(project_.Name), formDiv)
		}
	}

	// manage the suppress operation
	if projectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		project_.Unstage(projectFormCallback.probe.stageOfInterest)
	}

	projectFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Project](
		projectFormCallback.probe,
	)
	projectFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if projectFormCallback.CreationMode || projectFormCallback.formGroup.HasSuppressButtonBeenPressed {
		projectFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(projectFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__ProjectFormCallback(
			nil,
			projectFormCallback.probe,
			newFormGroup,
		)
		project := new(models.Project)
		FillUpForm(project, newFormGroup, projectFormCallback.probe)
		projectFormCallback.probe.formStage.Commit()
	}

	fillUpTree(projectFormCallback.probe)
}
func __gong__New__TaskFormCallback(
	task *models.Task,
	probe *Probe,
	formGroup *table.FormGroup,
) (taskFormCallback *TaskFormCallback) {
	taskFormCallback = new(TaskFormCallback)
	taskFormCallback.probe = probe
	taskFormCallback.task = task
	taskFormCallback.formGroup = formGroup

	taskFormCallback.CreationMode = (task == nil)

	return
}

type TaskFormCallback struct {
	task *models.Task

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe

	formGroup *table.FormGroup
}

func (taskFormCallback *TaskFormCallback) OnSave() {

	log.Println("TaskFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	taskFormCallback.probe.formStage.Checkout()

	if taskFormCallback.task == nil {
		taskFormCallback.task = new(models.Task).Stage(taskFormCallback.probe.stageOfInterest)
	}
	task_ := taskFormCallback.task
	_ = task_

	for _, formDiv := range taskFormCallback.formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(task_.Name), formDiv)
		case "Project:Tasks":
			// we need to retrieve the field owner before the change
			var pastProjectOwner *models.Project
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Project"
			rf.Fieldname = "Tasks"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				taskFormCallback.probe.stageOfInterest,
				taskFormCallback.probe.backRepoOfInterest,
				task_,
				&rf)

			if reverseFieldOwner != nil {
				pastProjectOwner = reverseFieldOwner.(*models.Project)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastProjectOwner != nil {
					idx := slices.Index(pastProjectOwner.Tasks, task_)
					pastProjectOwner.Tasks = slices.Delete(pastProjectOwner.Tasks, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _project := range *models.GetGongstructInstancesSet[models.Project](taskFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _project.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newProjectOwner := _project // we have a match
						if pastProjectOwner != nil {
							if newProjectOwner != pastProjectOwner {
								idx := slices.Index(pastProjectOwner.Tasks, task_)
								pastProjectOwner.Tasks = slices.Delete(pastProjectOwner.Tasks, idx, idx+1)
								newProjectOwner.Tasks = append(newProjectOwner.Tasks, task_)
							}
						} else {
							newProjectOwner.Tasks = append(newProjectOwner.Tasks, task_)
						}
					}
				}
			}
		}
	}

	// manage the suppress operation
	if taskFormCallback.formGroup.HasSuppressButtonBeenPressed {
		task_.Unstage(taskFormCallback.probe.stageOfInterest)
	}

	taskFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Task](
		taskFormCallback.probe,
	)
	taskFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if taskFormCallback.CreationMode || taskFormCallback.formGroup.HasSuppressButtonBeenPressed {
		taskFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
		}).Stage(taskFormCallback.probe.formStage)
		newFormGroup.OnSave = __gong__New__TaskFormCallback(
			nil,
			taskFormCallback.probe,
			newFormGroup,
		)
		task := new(models.Task)
		FillUpForm(task, newFormGroup, taskFormCallback.probe)
		taskFormCallback.probe.formStage.Commit()
	}

	fillUpTree(taskFormCallback.probe)
}
